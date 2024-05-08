package adm

import (
  "fmt"
  "git.dev.opnd.io/gc/backend-admin/pkg/handler/util"
  "git.dev.opnd.io/gc/backend-admin/pkg/model/request"
  "git.dev.opnd.io/gc/backend-admin/pkg/model/response"
  "git.dev.opnd.io/gc/backend-admin/pkg/service"
  "github.com/labstack/echo/v4"
  "net/http"
)

func GetInvestmentsForInvestmentList(c echo.Context) error {
  page, limit, err := util.GetPageAndLimit(c)
  if err != nil {
    return err
  }
  resData := response.Data{}
  investments, err := service.GetInvestmentsForInvestmentList(page, limit, c)
  if err != nil {
    resData.Success = false
    resData.Data = err.Error()
    return c.JSON(http.StatusOK, resData)
  }
  resData.Success = true
  resData.Data = response.Investments{Investments: investments}
  return c.JSON(http.StatusOK, resData)
}

func GetInvestmentsForInvestmentListCount(c echo.Context) error {
  resData := response.Data{}
  count, err := service.GetInvestmentsForInvestmentListCount(c)
  if err != nil {
    resData.Success = false
    resData.Data = err.Error()
    return c.JSON(http.StatusOK, resData)
  }
  resData.Success = true
  resData.Data = count
  return c.JSON(http.StatusOK, resData)
}

func GetInvestmentsForUserList(c echo.Context) error {
  page, limit, err := util.GetPageAndLimit(c)
  if err != nil {
    return err
  }
  resData := response.Data{}
  userId, err := util.GetIdFromPath(c)
  if err != nil {
    resData.Success = false
    resData.Data = err.Error()
    return c.JSON(http.StatusOK, resData)
  }
  investments, err := service.GetInvestmentsForUserList(page, limit, userId, c)
  if err != nil {
    resData.Success = false
    resData.Data = err.Error()
    return c.JSON(http.StatusOK, resData)
  }
  resData.Success = true
  resData.Data = response.Investments{Investments: investments}
  return c.JSON(http.StatusOK, resData)
}

func GetInvestmentsForUserListCount(c echo.Context) error {
  resData := response.Data{}
  userId, err := util.GetIdFromPath(c)
  if err != nil {
    resData.Success = false
    resData.Data = err.Error()
    return c.JSON(http.StatusOK, resData)
  }
  count, err := service.GetInvestmentsForUserListCount(userId, c)
  if err != nil {
    resData.Success = false
    resData.Data = err.Error()
    return c.JSON(http.StatusOK, resData)
  }
  resData.Success = true
  resData.Data = count
  return c.JSON(http.StatusOK, resData)
}

func CancelUserInvestment(c echo.Context) error {

  var ids request.UserIds
  if err := c.Bind(&ids); err != nil {
    fmt.Println("[AddAgency] Bind error: ", err.Error())
    //ioutil.ReadAll(c.Request().Body)
    //fmt.Println("[AddAgency] body: ", )
    return c.NoContent(http.StatusBadRequest)
  }

  resData := response.Data{}
  investmentId, err := util.GetIdFromPath(c)
  if err != nil {
    resData.Success = false
    resData.Data = err.Error()
    return c.JSON(http.StatusOK, resData)
  }
  users, err := service.CancelUserInvestment(investmentId, ids, c)
  if err != nil {
    resData.Success = false
    resData.Data = err.Error()
    return c.JSON(http.StatusOK, resData)
  }
  resData.Success = true
  resData.Data = response.Users{Users: users}
  return c.JSON(http.StatusOK, resData)
}

func ConfirmDeposit(c echo.Context) error {
  var ids request.UserIds
  if err := c.Bind(&ids); err != nil {
    fmt.Println("[AddAgency] Bind error: ", err.Error())
    //ioutil.ReadAll(c.Request().Body)
    //fmt.Println("[AddAgency] body: ", )
    return c.NoContent(http.StatusBadRequest)
  }

  resData := response.Data{}
  investmentId, err := util.GetIdFromPath(c)
  if err != nil {
    resData.Success = false
    resData.Data = err.Error()
    return c.JSON(http.StatusOK, resData)
  }
  users, err := service.ConfirmDeposit(investmentId, ids, c)
  if err != nil {
    resData.Success = false
    resData.Data = err.Error()
    return c.JSON(http.StatusOK, resData)
  }
  resData.Success = true
  resData.Data = response.Users{Users: users}
  return c.JSON(http.StatusOK, resData)
}
