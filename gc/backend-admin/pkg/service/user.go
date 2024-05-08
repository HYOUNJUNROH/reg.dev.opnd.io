package service

import (
  "errors"
  "fmt"
  "git.dev.opnd.io/gc/backend-admin/pkg/model/response"
  "git.dev.opnd.io/gc/backend-admin/pkg/util/db_middleware"
  "github.com/labstack/echo/v4"
  "gorm.io/gorm"
)

func GetUsersForUserList(page, limit int, ctx echo.Context) ([]response.User, error) {
  // response.Users
  queryStr := `
    select
      id,
      created_at,
      user_id,
      gender,
      name,
      phone1,
      phone2
    from users
    order by created_at
    offset %d limit %d;
    `
  queryStr = fmt.Sprintf(queryStr, (page-1)*limit, limit)

  return getUsers(queryStr, ctx)
}

func GetUsersForUserListCount(ctx echo.Context) (int, error) {
  queryStr := `
    select
      count(*) as count
    from users
    `
  return getUsersCount(queryStr, ctx)
}

func GetUsersForInvestmentList(page, limit int, investmentId int, ctx echo.Context) ([]response.User, error) {
  // response.Users
  queryStr := `
    select
      u.id as id,
      u.user_id as user_id,
      u.gender as gender,
      u.name as name,
      u.phone1 as phone1,
      u.phone2 as phone2,
      ui.amount as amount,
      ui.paid_status as paid_status
    from users u
    left join user_investments ui on u.id = ui.created_by
    where ui.investment_id = %d
    order by u.created_at
    offset %d limit %d;
    `
  queryStr = fmt.Sprintf(queryStr, investmentId, (page-1)*limit, limit)
  return getUsers(queryStr, ctx)
}

func GetUsersForInvestmentListCount(investmentId int, ctx echo.Context) (int, error) {
  queryStr := `
    select
      count(*) as count
    from users u
    left join user_investments ui on u.id = ui.created_by
    where ui.investment_id = %d
    `
  queryStr = fmt.Sprintf(queryStr, investmentId)
  return getUsersCount(queryStr, ctx)
}

func getUsers(queryStr string, ctx echo.Context) ([]response.User, error) {
  var users []response.User
  if err := db_middleware.GetDBFromEcho(ctx).Raw(queryStr).Scan(&users).Error; err != nil {
    if errors.Is(err, gorm.ErrRecordNotFound) {
      return nil, nil
    } else {
      return nil, err
    }
  }
  return users, nil
}

func getUsersCount(queryStr string, ctx echo.Context) (int, error) {
  var count = 0
  if err := db_middleware.GetDBFromEcho(ctx).Raw(queryStr).Scan(&count).Error; err != nil {
    if errors.Is(err, gorm.ErrRecordNotFound) {
      return count, nil
    } else {
      return count, err
    }
  }
  return count, nil

}
