package router

import (
	"fmt"
	"net"
	"net/http"

	"git.dev.opnd.io/gc/backend-admin/pkg/logger"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
)

func mustParseCIDR(s string) *net.IPNet {
	_, IPNet, err := net.ParseCIDR(s)
	if err != nil {
		logger.Logger.Fatal(err)
	}
	return IPNet
}

func NewIPFilter(cidr *net.IPNet) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if cidr != nil && !cidr.Contains(net.ParseIP(c.RealIP())) {
				return echo.NewHTTPError(http.StatusUnauthorized,
					fmt.Sprintf("IP address %s not allowed", c.RealIP()))
			}
			return next(c)
		}
	}
}

func GrpcMiddleware(grpcServer *grpc.Server) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			if c.Request().Header.Get("Content-Type") == "application/grpc" {
				grpcServer.ServeHTTP(c.Response().Writer, c.Request())
				return nil
			}
			return next(c)
		}
	}
}
