package adm

import (
  "git.dev.opnd.io/gc/backend-admin/pkg/handler/util"
  "git.dev.opnd.io/gc/backend-admin/pkg/model/response"
  "git.dev.opnd.io/gc/backend-admin/pkg/service"
  "github.com/labstack/echo/v4"
  "net/http"
)

func GetUsersForUserList(c echo.Context) error {
  page, limit, err := util.GetPageAndLimit(c)
  if err != nil {
    return err
  }
  resData := response.Data{}
  users, err := service.GetUsersForUserList(page, limit, c)
  if err != nil {
    resData.Success = false
    resData.Data = err.Error()
    return c.JSON(http.StatusOK, resData)
  }
  resData.Success = true
  resData.Data = response.Users{Users: users}
  return c.JSON(http.StatusOK, resData)
}

func GetUsersForUserListCount(c echo.Context) error {
  resData := response.Data{}
  count, err := service.GetUsersForUserListCount(c)
  if err != nil {
    resData.Success = false
    resData.Data = err.Error()
    return c.JSON(http.StatusOK, resData)
  }
  resData.Success = true
  resData.Data = count
  return c.JSON(http.StatusOK, resData)
}

func GetUsersForInvestmentList(c echo.Context) error {
  page, limit, err := util.GetPageAndLimit(c)
  if err != nil {
    return err
  }
  resData := response.Data{}
  investmentId, err := util.GetIdFromPath(c)
  if err != nil {
    resData.Success = false
    resData.Data = err.Error()
    return c.JSON(http.StatusOK, resData)
  }
  users, err := service.GetUsersForInvestmentList(page, limit, investmentId, c)
  if err != nil {
    resData.Success = false
    resData.Data = err.Error()
    return c.JSON(http.StatusOK, resData)
  }
  resData.Success = true
  resData.Data = response.Users{Users: users}
  return c.JSON(http.StatusOK, resData)
}

func GetUsersForInvestmentListCount(c echo.Context) error {
  resData := response.Data{}
  investmentId, err := util.GetIdFromPath(c)
  if err != nil {
    resData.Success = false
    resData.Data = err.Error()
    return c.JSON(http.StatusOK, resData)
  }
  count, err := service.GetUsersForInvestmentListCount(investmentId, c)
  if err != nil {
    resData.Success = false
    resData.Data = err.Error()
    return c.JSON(http.StatusOK, resData)
  }
  resData.Success = true
  resData.Data = count
  return c.JSON(http.StatusOK, resData)
}
