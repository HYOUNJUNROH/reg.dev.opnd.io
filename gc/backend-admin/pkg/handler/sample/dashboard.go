package sample

// import (
// 	"context"
// 	"net/http"
// 	"strconv"

// 	"git.dev.opnd.io/gc/backend-admin/pkg/handler/util"
// 	sample "git.dev.opnd.io/gc/backend-admin/pkg/service/sample"
// 	"github.com/labstack/echo/v4"
// 	uuid "github.com/satori/go.uuid"
// 	"github.com/volatiletech/null/v9"
// )

// func NewDashboard(c echo.Context) error {
// 	_, err := util.GetUserFromEchoContext(c)
// 	if err != nil {
// 		return c.JSON(http.StatusUnauthorized, err.Error())
// 	}

// 	type Request struct {
// 		Name        string
// 		Description null.String
// 		DbPathLtree string
// 		Url         null.String
// 	}

// 	request := &Request{}
// 	if err := c.Bind(request); err != nil {
// 		return c.JSON(http.StatusBadRequest, err.Error())
// 	}

// 	dashboard, err := sample.NewDashboard(util.GetUserContextFromEchoContext(c, context.Background()), request.Name, request.Description, request.DbPathLtree, request.Url)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, err.Error())
// 	}

// 	return c.JSON(http.StatusOK, dashboard)
// }
