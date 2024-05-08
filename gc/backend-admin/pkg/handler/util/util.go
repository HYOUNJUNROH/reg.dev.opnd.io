package util

import (
  "context"
  "strconv"
  "strings"

  "git.dev.opnd.io/gc/backend-admin/pkg/config/db"
  "git.dev.opnd.io/gc/backend-admin/pkg/model"
  "github.com/golang-jwt/jwt"
  "github.com/labstack/echo/v4"
  uuid "github.com/satori/go.uuid"
  "github.com/volatiletech/null/v9"
)

func GetUserFromID(id uuid.UUID) (*model.Users, error) {
  var user model.Users
  err := db.DB.Model(&model.Users{}).Where("id = ?", id).First(&user).Error
  if err != nil {
    return nil, err
  }
  return &user, nil
}

func GetUserFromJWT(token *jwt.Token) (*model.Users, error) {
  claims := token.Claims.(jwt.MapClaims)
  if id, err := uuid.FromString(claims["sub"].(string)); err != nil {
    return nil, err
  } else {
    return &model.Users{
      ID:   id,
      Role: null.StringFrom(claims["role"].(string)),
    }, nil
    return GetUserFromID(id)
  }
}

func GetUserFromEchoContext(c echo.Context) (*model.Users, error) {
  user := c.Get("user").(*jwt.Token)
  return GetUserFromJWT(user)
}

func GetUserContextFromEchoContext(c echo.Context, ctx context.Context) context.Context {
  user := c.Get("user").(*jwt.Token)
  return context.WithValue(ctx, "user", user)
}

func GetPageAndLimit(c echo.Context) (int, int, error) {
  var err error
  var page, limit int
  pageStr := c.QueryParam("page")
  if strings.TrimSpace(pageStr) != "" {
    if page, err = strconv.Atoi(pageStr); err != nil {
      //fmt.Println(fmt.Sprintf("[%s] Bind error: %s", tag, err.Error()))
      return 0, 0, err
    }
  }
  limitStr := c.QueryParam("limit")
  if strings.TrimSpace(limitStr) != "" {
    if limit, err = strconv.Atoi(limitStr); err != nil {
      //fmt.Println(fmt.Sprintf("[%s] Bind error: %s", tag, err.Error()))
      return 0, 0, err
    }
  }
  return page, limit, nil
}

func GetIdFromPath(c echo.Context) (int, error) {
  idStr := c.Param("id")
  if id, err := strconv.Atoi(idStr); err != nil {
    return 0, err
  } else {
    return id, nil
  }
}
