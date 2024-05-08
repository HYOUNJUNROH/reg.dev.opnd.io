package adm

import (
  "git.dev.opnd.io/gc/backend-admin/pkg/service"
  "github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
  return service.Login(c)
}

//func Logout(c echo.Context) error {
//  return c.NoContent(http.StatusOK)
//}

func ChangePassword(c echo.Context) error {
  return service.ChangePassword(c)
}
