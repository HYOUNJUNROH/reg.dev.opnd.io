package logger

import (
	"os"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// var Logger *zap.Logger
var logger *zap.Logger
var Logger *zap.SugaredLogger

func Initialize(production bool) {
	var config zapcore.EncoderConfig
	var defaultLogLevel zapcore.Level
	if production {
		config = zap.NewProductionEncoderConfig()
		defaultLogLevel = zapcore.DebugLevel
	} else {
		config = zap.NewDevelopmentEncoderConfig()
		defaultLogLevel = zapcore.InfoLevel
	}
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(config)
	// fileEncoder := zapcore.NewJSONEncoder(config)
	// logFile, _ := os.OpenFile("text.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// writer := zapcore.AddSync(logFile)
	core := zapcore.NewTee(
		// zapcore.NewCore(fileEncoder, writer, defaultLogLevel),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), defaultLogLevel),
	)
	logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	Logger = logger.Sugar()
}

// gist.github.com/ndrewnee/6187a01427b9203b9f11ca5864b8a60d
// ZapLogger is an example of echo middleware that logs requests using logger "zap"
func ZapLogger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			err := next(c)
			if err != nil {
				c.Error(err)
			}

			req := c.Request()
			res := c.Response()

			id := req.Header.Get(echo.HeaderXRequestID)
			if id == "" {
				id = res.Header().Get(echo.HeaderXRequestID)
			}

			fields := []zapcore.Field{
				zap.Int("status", res.Status),
				zap.String("latency", time.Since(start).String()),
				zap.String("id", id),
				zap.String("method", req.Method),
				zap.String("uri", req.RequestURI),
				zap.String("host", req.Host),
				zap.String("remote_ip", c.RealIP()),
			}

			n := res.Status
			if strings.HasPrefix(req.RequestURI, "/health") {
				switch {
				case n >= 500:
					logger.Error("Server error", fields...)
				case n >= 400:
					logger.Warn("Client error", fields...)
				case n >= 300:
					logger.Info("Redirection", fields...)
				default:
					// logger.Info("Success", fields...)
				}
			} else {
				switch {
				case n >= 500:
					logger.Error("Server error", fields...)
				case n >= 400:
					logger.Warn("Client error", fields...)
				case n >= 300:
					logger.Info("Redirection", fields...)
				default:
					logger.Info("Success", fields...)
				}
			}
			return nil
		}
	}
}
