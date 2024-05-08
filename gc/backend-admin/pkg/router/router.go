package router

import (
  "context"
  "fmt"
  "git.dev.opnd.io/gc/backend-admin/pkg/handler/adm"
  "git.dev.opnd.io/gc/backend-admin/pkg/logger"
  "git.dev.opnd.io/gc/backend-admin/pkg/service"
  "git.dev.opnd.io/gc/backend-admin/pkg/util/db_middleware"
  "net/http"
  "regexp"
  "time"

  "git.dev.opnd.io/gc/backend-admin/pkg/config"
  "git.dev.opnd.io/gc/backend-admin/pkg/config/db"
  "git.dev.opnd.io/gc/backend-admin/pkg/config/kvstore"
  "github.com/hellofresh/health-go/v4"
  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
  "google.golang.org/grpc"
)

//func RouterSetup(e *echo.Echo, grpcServer *grpc.Server) {
//  e.IPExtractor = echo.ExtractIPFromRealIPHeader(echo.TrustIPRange(mustParseCIDR(config.Config.TrustedCIDR)))
//  e.IPExtractor = echo.ExtractIPFromXFFHeader(echo.TrustIPRange(mustParseCIDR(config.Config.TrustedCIDR)))
//
//  matchString := `(^((https?)|(capacitor)|(ionic)):\/\/` + "localhost" + `$)|(^https?:\/\/.*` + config.ExternalHost + `$)|(^https?:\/\/.*` + config.ExternalMobileHost + `$)`
//  logger.Logger.Info("CORS : ", matchString)
//
//  allowOrigin := func(origin string) (bool, error) {
//    return regexp.MatchString(matchString, origin)
//  }
//
//  if config.IsDevelopment() {
//    e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
//      // Access-Control-Allow-Credentials
//      AllowCredentials: true,
//      AllowOriginFunc: func(origin string) (bool, error) {
//        return true, nil
//      },
//      AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
//    }))
//  } else {
//    e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
//      // Access-Control-Allow-Credentials
//      AllowCredentials: true,
//      AllowOriginFunc:  allowOrigin,
//      AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
//    }))
//  }
//
//  // disable grpc
//  if config.CertKeyPair != nil {
//    e.Use(GrpcMiddleware(grpcServer))
//  }
//
//  // register grpc-gateway
//  mux := runtime.NewServeMux()
//  {
//    //ctx := context.WithValue(context.Background(), db_middleware.DBContextKey, db.DB)
//    //server := &grpc_impl.GcServerImpl{}
//    //api.RegisterGcServer(grpcServer, server)
//    //api.RegisterGcHandlerServer(ctx, mux, server)
//  }
//
//  jwtConfig := middleware.JWTConfig{
//    TokenLookup: "cookie:sb-access-token,header:Authorization:Bearer ",
//    SigningKey:  []byte(config.Config.GoTrue.JwtSecret),
//  }
//
//  jwtGroup := e.Group("")
//  jwtGroup.Use(middleware.JWTWithConfig(jwtConfig))
//
//  // register grpc-gateway
//  jwtGroup.Any("/v1/*", echo.WrapHandler(mux))
//
//  // register base handler
//  SetupBaseHandler(e, grpcServer)
//
//  // register grpc-web
//  wrappedGrpc := grpcweb.WrapServer(grpcServer)
//  e.Any("/*", echo.WrapHandler(wrappedGrpc))
//}

func SetupBaseHandler(e *echo.Echo, grpcServer *grpc.Server) {
  //*
  matchString := `(^((https?)|(capacitor)|(ionic)):\/\/` + "localhost" + `$)|(^https?:\/\/.*` + config.ExternalHost + `$)|(^https?:\/\/.*` + config.ExternalMobileHost + `$)`
  logger.Logger.Info("CORS : ", matchString)

  allowOrigin := func(origin string) (bool, error) {
    return regexp.MatchString(matchString, origin)
  }
  if config.IsDevelopment() {
    e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
      // Access-Control-Allow-Credentials
      AllowCredentials: true,
      AllowOriginFunc: func(origin string) (bool, error) {
        return true, nil
      },
      AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
    }))
  } else {
    e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
      // Access-Control-Allow-Credentials
      AllowCredentials: true,
      AllowOriginFunc:  allowOrigin,
      AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
    }))
  }
  //*/

  h, _ := health.New(health.WithChecks(health.Config{
    Name:      "postgresql",
    Timeout:   time.Second * 5,
    SkipOnErr: true,
    Check: func(ctx context.Context) (checkErr error) {
      if db.DB == nil {
        return nil
      }
      checkErr = nil
      d, err := db.DB.DB()
      if err != nil {
        checkErr = fmt.Errorf("PostgreSQL health check failed on get instance: %w", err)
        return
      }
      err = d.PingContext(ctx)
      if err != nil {
        checkErr = fmt.Errorf("PostgreSQL health check failed on ping: %w", err)
        return
      }
      rows, err := d.QueryContext(ctx, `SELECT VERSION()`)
      if err != nil {
        checkErr = fmt.Errorf("PostgreSQL health check failed on select: %w", err)
        return
      }
      defer func() {
        // override checkErr only if there were no other errors
        if err = rows.Close(); err != nil && checkErr == nil {
          checkErr = fmt.Errorf("PostgreSQL health check failed on rows closing: %w", err)
        }
      }()
      return nil
    }}, health.Config{
    Name: "redis",
    Check: func(ctx context.Context) error {
      if kvstore.RedisClient == nil {
        return nil
      }
      pong, err := kvstore.RedisClient.Ping(ctx).Result()
      if err != nil {
        return fmt.Errorf("redis ping failed: %w", err)
      }

      if pong != "PONG" {
        return fmt.Errorf("unexpected response for redis ping: %q", pong)
      }
      return nil
    },
  },
  ))

  healthHandler := func(c echo.Context) error {
    h.HandlerFunc(c.Response(), c.Request())
    return nil
  }

  // 외부에서 health 체크 할 수 없게 제한.
  trustedMiddleware := NewIPFilter(mustParseCIDR(config.Config.TrustedCIDR))
  e.GET("/health", trustedMiddleware(healthHandler))

  //// 이미지 프록시, nCloud의 CORS 이슈로 작성.
  //
  //g := e.Group("/data")
  //g.Use(middleware.Logger())
  //
  //// 간단한 s3 전달 핸들러
  //g.GET("/*", func(c echo.Context) error {
  //  path := strings.TrimPrefix(c.Request().URL.Path, "/api/data/")
  //  if stream, err := s3.Client.GetStream(path); err != nil {
  //    logger.Logger.Error(err)
  //    return c.String(http.StatusNotFound, "not found")
  //  } else {
  //    defer stream.Close()
  //    c.Response().Header().Set(echo.HeaderContentType, mime.TypeByExtension(filepath.Ext(path)))
  //    c.Response().WriteHeader(http.StatusOK)
  //    _, err = io.Copy(c.Response(), stream)
  //    if err != nil {
  //      logger.Logger.Error(err)
  //    }
  //  }
  //  return nil
  //})

  admGroup := e.Group("/adm")
  admGroup.Use(middleware.Logger())
  admGroup.Use(db_middleware.DBMiddleware(db.DB))

  authGroup := admGroup.Group("/auth")
  authGroup.POST("/login", adm.Login)
  //authGroup.GET("/logout", adm.Logout)
  authGroup.POST("/update-password", adm.ChangePassword, service.VerifyToken)

  userGroup := admGroup.Group("/users", service.VerifyToken)
  userGroup.GET("", adm.GetUsersForUserList)
  userGroup.GET("/count", adm.GetUsersForUserListCount)
  userGroup.GET("/:id/investments", adm.GetInvestmentsForUserList)
  userGroup.GET("/:id/investments/count", adm.GetInvestmentsForUserListCount)

  investmentGroup := admGroup.Group("/investments", service.VerifyToken)
  investmentGroup.GET("", adm.GetInvestmentsForInvestmentList)
  investmentGroup.GET("/count", adm.GetInvestmentsForInvestmentListCount)
  investmentGroup.GET("/:id/users", adm.GetUsersForInvestmentList)
  investmentGroup.GET("/:id/users/count", adm.GetUsersForInvestmentListCount)
  investmentGroup.POST("/:id/users/cancel", adm.CancelUserInvestment)
  investmentGroup.POST("/:id/users/deposit", adm.ConfirmDeposit)

  bannerGroup := admGroup.Group("/banners", service.VerifyToken)
  bannerGroup.GET("", adm.GetBanners)
  bannerGroup.POST("", adm.PostBanner)
  bannerGroup.POST("/priority", adm.PostPriorityBanner)
  bannerGroup.DELETE("/:id", adm.DeleteBanner)

  imageStatic := admGroup.Group("/images")
  imageStatic.Static("/", "/usr/src/app/uploads")

  // 버젼 출력
  e.GET("/api/version", config.GetVersion)
}
