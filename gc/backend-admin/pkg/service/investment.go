package service

import (
  "errors"
  "fmt"
  "git.dev.opnd.io/gc/backend-admin/pkg/model/request"
  "git.dev.opnd.io/gc/backend-admin/pkg/model/response"
  "git.dev.opnd.io/gc/backend-admin/pkg/util/db_middleware"
  "github.com/labstack/echo/v4"
  "gorm.io/gorm"
)

func GetInvestmentsForInvestmentList(page, limit int, ctx echo.Context) ([]response.Investment, error) {
  queryStr := `
   select
     i.id as id,
     i.invest_id as invest_id,
     i.created_at as created_at,
     i.start_date as start_date,
     i.status as status,
     (select count(*) from user_investments ui where ui.investment_id = i.id) as user_count,
     i.current_invest as current_invest,
     i.max_invest as max_invest,
     i.title as title
   from investments i
   where i.disabled_at is null
   order by i.id
   offset %d limit %d
   `
  queryStr = fmt.Sprintf(queryStr, (page-1)*limit, limit)
  //fmt.Println("[GetInvestmentsForInvestmentList] queryStr: ", queryStr)
  return getInvestments(queryStr, ctx)
}

func GetInvestmentsForInvestmentListCount(ctx echo.Context) (int, error) {
  queryStr := `
   select
     count(*) as count
   from investments i
   `
  return getInvestmentsCount(queryStr, ctx)
}

func GetInvestmentsForUserList(page, limit int, userId int, ctx echo.Context) ([]response.Investment, error) {
  queryStr := `
    select
      i.id as id,
      i.invest_id as invest_id,
      i.created_at as created_at,
      i.start_date as start_date,
      i.status as status,
      ui.amount as amount,
      ui.paid_status as paid_status,
      i.title as title
    from investments i
    left join user_investments ui on i.id = ui.investment_id
    where ui.created_by = %d
    and i.disabled_at is null
    order by i.id
    offset %d limit %d
    `
  queryStr = fmt.Sprintf(queryStr, userId, (page-1)*limit, limit)
  //fmt.Println("[GetInvestmentsForUserList] queryStr: ", queryStr)
  return getInvestments(queryStr, ctx)
}

func GetInvestmentsForUserListCount(userId int, ctx echo.Context) (int, error) {
  queryStr := `
    select
      count(*) as count
    from investments i
    left join user_investments ui on i.id = ui.investment_id
    where ui.created_by = %d
    `
  queryStr = fmt.Sprintf(queryStr, userId)
  return getInvestmentsCount(queryStr, ctx)
}

func CancelUserInvestment(investmentId int, userIds request.UserIds, c echo.Context) ([]response.User, error) {
  var ids []int
  for _, userId := range userIds.UserIds {
    ids = append(ids, userId.ID)
  }

  if len(ids) == 0 {
    return nil, errors.New("no user ids provided")
  }

  tx := db_middleware.GetDBFromEcho(c).Begin()

  // 납부가 여부 체크
  queryStr := `
    select
      u.id as id,
      u.user_id as user_id,
      u.name as name,
      u.phone1 as phone1,
      u.phone2 as phone2,
      ui.amount as amount,
      ui.paid_status as paid_status
    from users u
    left join user_investments ui on u.id = ui.created_by
    where ui.investment_id = ?
    and ui.created_by IN ?
    order by u.created_at
    `
  var targetUsers []response.User
  if ret := tx.Raw(queryStr, investmentId, ids).Find(&targetUsers).Error; ret != nil {
    if !errors.Is(ret, gorm.ErrRecordNotFound) {
      tx.Rollback()
      return nil, ret
    }
  }

  //db_middleware.GetDBFromContext(ctx).Table("user_investments").Where("investment_id = ? and created_by = ?", investmentId, userId).Delete(&response.UserInvestment{})

  // 납부 안된 회원만 삭제처리
  ret := tx.Unscoped().Exec("delete from user_investments where investment_id = ? and created_by IN ? and paid_status != 'paid'", investmentId, ids)
  if ret.Error != nil {
    tx.Rollback()
    return nil, ret.Error
  }

  // 취소된 사람의 구좌수에 맞게 현재 구좌수를 업데이트
  queryStr = `
  select COALESCE(sum(ui.amount), 0) from user_investments ui
  where ui.investment_id = ? and ui.created_by NOT IN ?
  `
  var amount int
  if ret := tx.Raw(queryStr, investmentId, ids).First(&amount).Error; ret != nil {
    if !errors.Is(ret, gorm.ErrRecordNotFound) {
      tx.Rollback()
      return nil, ret
    }
  }

  if ret := tx.Table("investments").Where("id = ?", investmentId).Update("current_invest", amount); ret.Error != nil {
    tx.Rollback()
    return nil, ret.Error
  }

  tx.Commit()

  var resultUsers []response.User
  for _, targetUser := range targetUsers {
    if targetUser.PaidStatus == "paid" {
      targetUser.CancelStatus = "faild"
    } else {
      targetUser.CancelStatus = "success"
    }
    targetUser.PaidStatus = ""
    resultUsers = append(resultUsers, targetUser)
  }

  return resultUsers, nil
}

func getInvestments(queryStr string, ctx echo.Context) ([]response.Investment, error) {
  var investments []response.Investment
  if err := db_middleware.
    GetDBFromEcho(ctx).
    Raw(queryStr).
    Scan(&investments).
    Error; err != nil {
    if errors.Is(err, gorm.ErrRecordNotFound) {
      return nil, nil
    } else {
      return nil, err
    }
  }
  return investments, nil
}

func getInvestmentsCount(queryStr string, ctx echo.Context) (int, error) {
  var count = 0
  if err := db_middleware.
    GetDBFromEcho(ctx).
    Raw(queryStr).
    Scan(&count).
    Error; err != nil {
    if errors.Is(err, gorm.ErrRecordNotFound) {
      return count, nil
    } else {
      return count, err
    }
  }
  return count, nil
}

func ConfirmDeposit(investmentId int, userIds request.UserIds, c echo.Context) ([]response.User, error) {
  var ids []int
  for _, userId := range userIds.UserIds {
    ids = append(ids, userId.ID)
  }

  if len(ids) == 0 {
    return nil, errors.New("no user ids provided")
  }

  // 납부 여부 체크
  queryStr := `
    select
      u.id as id,
      u.user_id as user_id,
      u.name as name,
      u.phone1 as phone1,
      u.phone2 as phone2,
      ui.amount as amount,
      ui.paid_status as paid_status
    from users u
    left join user_investments ui on u.id = ui.created_by
    where ui.investment_id = ?
    and ui.created_by IN ?
    order by u.created_at
    `
  var targetUsers []response.User
  if ret := db_middleware.GetDBFromEcho(c).Raw(queryStr, investmentId, ids).Find(&targetUsers).Error; ret != nil {
    if !errors.Is(ret, gorm.ErrRecordNotFound) {
      return nil, ret
    }
  }

  // 납부 안된 회원만 처리
  //db_middleware.GetDBFromContext(ctx).Table("user_investments").Where("investment_id = ? and created_by = ?", investmentId, userId).Delete(&response.UserInvestment{})
  ret := db_middleware.GetDBFromEcho(c).Exec("update user_investments set paid_status = 'paid' where investment_id = ? and created_by IN ? and paid_status != 'paid'", investmentId, ids)
  if ret.Error != nil {
    return nil, ret.Error
  }

  var resultUsers []response.User
  for _, targetUser := range targetUsers {
    targetUser.PaidStatus = "paid"
    resultUsers = append(resultUsers, targetUser)
  }

  return resultUsers, nil
}
