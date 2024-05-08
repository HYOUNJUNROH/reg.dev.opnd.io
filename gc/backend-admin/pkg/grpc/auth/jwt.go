package grpc_auth

// reference : https://github.com/labstack/echo/blob/v4.10.2/middleware/jwt.go

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"git.dev.opnd.io/gc/backend-admin/pkg/logger"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type ValuesExtractor func(ctx context.Context) ([]string, error)
type SecurityInterceptor func(ctx context.Context, FullMethod string) (context.Context, error)
type Skipper func(ctx context.Context, FullMethod string) bool

func DefaultSkipper(ctx context.Context, FullMethod string) bool {
	return false
}

type (
	JWTConfig struct {
		Skipper          Skipper
		SigningKey       interface{}
		SigningKeys      map[string]interface{}
		SigningMethod    string
		ContextKey       string
		Claims           jwt.Claims
		TokenLookup      string
		TokenLookupFuncs []ValuesExtractor
		AuthScheme       string
		KeyFunc          jwt.Keyfunc
		ParseTokenFunc   func(auth string, ctx context.Context) (interface{}, error)
	}
)

// Algorithms
const (
	AlgorithmHS256 = "HS256"
)

// Errors
var (
	ErrJWTMissing = status.Error(codes.InvalidArgument, "missing or malformed jwt")
	ErrJWTInvalid = status.Error(codes.Unauthenticated, "invalid or expired jwt")
)

const (
	// extractorLimit is arbitrary number to limit values extractor can return. this limits possible resource exhaustion
	// attack vector
	extractorLimit = 20
)

var errHeaderExtractorValueMissing = errors.New("missing value in request header")
var errHeaderExtractorValueInvalid = errors.New("invalid value in request header")
var errCookieExtractorValueMissing = errors.New("missing value in cookies")

var (
	// DefaultJWTConfig is the default JWT auth middleware config.
	DefaultJWTConfig = JWTConfig{
		Skipper:          DefaultSkipper,
		SigningMethod:    AlgorithmHS256,
		ContextKey:       "user",
		TokenLookup:      "header:" + echo.HeaderAuthorization,
		TokenLookupFuncs: nil,
		AuthScheme:       "Bearer",
		Claims:           jwt.MapClaims{},
		KeyFunc:          nil,
	}
)

func CreateExtractors(lookups string) ([]ValuesExtractor, error) {
	return createExtractors(lookups, "")
}

func createExtractors(lookups string, authScheme string) ([]ValuesExtractor, error) {
	if lookups == "" {
		return nil, nil
	}
	sources := strings.Split(lookups, ",")
	var extractors = make([]ValuesExtractor, 0)
	for _, source := range sources {
		parts := strings.Split(source, ":")
		if len(parts) < 2 {
			return nil, fmt.Errorf("extractor source for lookup could not be split into needed parts: %v", source)
		}

		switch parts[0] {
		case "cookie":
			extractors = append(extractors, valuesFromCookie(parts[1]))
		case "header":
			prefix := ""
			if len(parts) > 2 {
				prefix = parts[2]
			} else if authScheme != "" && parts[1] == echo.HeaderAuthorization {
				// backwards compatibility for JWT and KeyAuth:
				// * we only apply this fix to Authorization as header we use and uses prefixes like "Bearer <token-value>" etc
				// * previously header extractor assumed that auth-scheme/prefix had a space as suffix we need to retain that
				//   behaviour for default values and Authorization header.
				prefix = authScheme
				if !strings.HasSuffix(prefix, " ") {
					prefix += " "
				}
			}
			extractors = append(extractors, valuesFromHeader(parts[1], prefix))
		}
	}
	return extractors, nil
}

// valuesFromHeader returns a functions that extracts values from the request header.
// valuePrefix is parameter to remove first part (prefix) of the extracted value. This is useful if header value has static
// prefix like `Authorization: <auth-scheme> <authorisation-parameters>` where part that we want to remove is `<auth-scheme> `
// note the space at the end. In case of basic authentication `Authorization: Basic <credentials>` prefix we want to remove
// is `Basic `. In case of JWT tokens `Authorization: Bearer <token>` prefix is `Bearer `.
// If prefix is left empty the whole value is returned.
func valuesFromHeader(header string, valuePrefix string) ValuesExtractor {
	prefixLen := len(valuePrefix)
	// standard library parses http.Request header keys in canonical form but we may provide something else so fix this
	header = strings.ToLower(header)
	return func(ctx context.Context) ([]string, error) {
		values, ok := getHeaderFromMetaData(ctx, header)
		if !ok || (values != nil && len(values) == 0) {
			return nil, errHeaderExtractorValueMissing
		}

		result := make([]string, 0)
		for i, value := range values {
			if prefixLen == 0 {
				result = append(result, value)
				if i >= extractorLimit-1 {
					break
				}
				continue
			}
			if len(value) > prefixLen && strings.EqualFold(value[:prefixLen], valuePrefix) {
				result = append(result, value[prefixLen:])
				if i >= extractorLimit-1 {
					break
				}
			}
		}

		if len(result) == 0 {
			if prefixLen > 0 {
				return nil, errHeaderExtractorValueInvalid
			}
			return nil, errHeaderExtractorValueMissing
		}
		return result, nil
	}
}

// valuesFromCookie returns a function that extracts values from the named cookie.
func valuesFromCookie(name string) ValuesExtractor {
	return func(ctx context.Context) ([]string, error) {
		cookies, ok := getCookiesFromMetaData(ctx)
		if !ok || (cookies != nil && len(cookies) == 0) {
			return nil, errCookieExtractorValueMissing
		}

		result := make([]string, 0)
		for i, cookie := range cookies {
			if name == cookie.Name {
				result = append(result, cookie.Value)
				if i >= extractorLimit-1 {
					break
				}
			}
		}
		if len(result) == 0 {
			return nil, errCookieExtractorValueMissing
		}
		return result, nil
	}
}

func getCookiesFromMetaData(ctx context.Context) ([]*http.Cookie, bool) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, false
	}
	auth, ok := md["cookie"]
	if !ok {
		auth, ok = md["grpcgateway-cookie"]
	}
	logger.Logger.Info(md)
	logger.Logger.Info(auth)
	if len(auth) == 0 || !ok {
		return nil, false
	}

	header := http.Header{}

	for i := 0; i < len(auth); i++ {
		header.Add("Cookie", auth[i])
	}
	request := http.Request{Header: header}

	return request.Cookies(), true
}

func getHeaderFromMetaData(ctx context.Context, header string) ([]string, bool) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, false
	}
	auth, ok := md[header]
	if !ok {
		auth, ok = md["grpcgateway-"+header]
	}
	logger.Logger.Info(md)
	logger.Logger.Info(auth)
	if len(auth) == 0 || !ok {
		return nil, false
	}
	return auth, true
}

func (config *JWTConfig) SecurityInterceptor() SecurityInterceptor {
	// Defaults
	if config.Skipper == nil {
		config.Skipper = DefaultJWTConfig.Skipper
	}
	if config.SigningKey == nil && len(config.SigningKeys) == 0 && config.KeyFunc == nil && config.ParseTokenFunc == nil {
		panic("echo: jwt middleware requires signing key")
	}
	if config.SigningMethod == "" {
		config.SigningMethod = DefaultJWTConfig.SigningMethod
	}
	if config.ContextKey == "" {
		config.ContextKey = DefaultJWTConfig.ContextKey
	}
	if config.Claims == nil {
		config.Claims = DefaultJWTConfig.Claims
	}
	if config.TokenLookup == "" && len(config.TokenLookupFuncs) == 0 {
		config.TokenLookup = DefaultJWTConfig.TokenLookup
	}
	if config.AuthScheme == "" {
		config.AuthScheme = DefaultJWTConfig.AuthScheme
	}
	if config.KeyFunc == nil {
		config.KeyFunc = config.defaultKeyFunc
	}
	if config.ParseTokenFunc == nil {
		config.ParseTokenFunc = config.defaultParseToken
	}

	extractors, err := createExtractors(config.TokenLookup, config.AuthScheme)
	if err != nil {
		panic(err)
	}
	if len(config.TokenLookupFuncs) > 0 {
		extractors = append(config.TokenLookupFuncs, extractors...)
	}

	return func(ctx context.Context, FullMethod string) (context.Context, error) {
		if config.Skipper(ctx, FullMethod) {
			return ctx, nil
		}

		var lastExtractorErr error
		var lastTokenErr error
		logger.Logger.Info(extractors)
		for _, extractor := range extractors {
			auths, err := extractor(ctx)
			if err != nil {
				lastExtractorErr = ErrJWTMissing // backwards compatibility: all extraction errors are same (unlike KeyAuth)
				continue
			}
			for _, auth := range auths {
				token, err := config.ParseTokenFunc(auth, ctx)
				logger.Logger.Info(err)
				if err != nil {
					lastTokenErr = err
					continue
				}
				// Store user information from token into context.
				return context.WithValue(ctx, config.ContextKey, token), nil
			}
		}

		// we are here only when we did not successfully extract or parse any of the tokens
		err := lastTokenErr
		if err == nil { // prioritize token errors over extracting errors
			err = lastExtractorErr
		}
		return ctx, err
	}
}

func (config *JWTConfig) defaultParseToken(auth string, ctx context.Context) (interface{}, error) {
	var token *jwt.Token
	var err error
	// Issue #647, #656
	if _, ok := config.Claims.(jwt.MapClaims); ok {
		token, err = jwt.Parse(auth, config.KeyFunc)
	} else {
		t := reflect.ValueOf(config.Claims).Type().Elem()
		claims := reflect.New(t).Interface().(jwt.Claims)
		token, err = jwt.ParseWithClaims(auth, claims, config.KeyFunc)
	}
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	return token, nil
}

// defaultKeyFunc returns a signing key of the given token.
func (config *JWTConfig) defaultKeyFunc(t *jwt.Token) (interface{}, error) {
	// Check the signing method
	if t.Method.Alg() != config.SigningMethod {
		return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
	}
	if len(config.SigningKeys) > 0 {
		if kid, ok := t.Header["kid"].(string); ok {
			if key, ok := config.SigningKeys[kid]; ok {
				return key, nil
			}
		}
		return nil, fmt.Errorf("unexpected jwt key id=%v", t.Header["kid"])
	}

	return config.SigningKey, nil
}
