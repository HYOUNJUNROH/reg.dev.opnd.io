package db_middleware

import (
  "context"
  "time"

  "github.com/labstack/echo/v4"
  "google.golang.org/grpc"
  "gorm.io/gorm"
)

type ContextKeyType string

const DBContextKey ContextKeyType = "__db"

func DBMiddleware(db *gorm.DB) echo.MiddlewareFunc {
  return func(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
      c.Set(string(DBContextKey), db)
      return next(c)
    }
  }
}

//func DBStreamInterceptor(db *gorm.DB) grpc.StreamServerInterceptor {
//	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
//		return handler(srv, grpc_auth.NewWrappedStream(ss, context.WithValue(ss.Context(), DBContextKey, db)))
//	}
//}

func DBUnaryInterceptor(db *gorm.DB) grpc.UnaryServerInterceptor {
  return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
    return handler(context.WithValue(ctx, DBContextKey, db), req)
  }
}

func GetDBFromContext(ctx context.Context) *gorm.DB {
  dbInstance := ctx.Value(DBContextKey).(*gorm.DB)
  return dbInstance
}

func GetDBFromEcho(c echo.Context) *gorm.DB {
  dbInstance := c.Get(string(DBContextKey)).(*gorm.DB)
  return dbInstance
}

func InitContextWithDB(ctx context.Context, db *gorm.DB) context.Context {
  return context.WithValue(ctx, DBContextKey, db)
}

func WithNow(ctx context.Context, now func() time.Time) context.Context {
  return context.WithValue(ctx, "now", now)
}
