package service

import (
  "errors"
  "fmt"
  "git.dev.opnd.io/gc/backend-admin/pkg/model/response"
  "git.dev.opnd.io/gc/backend-admin/pkg/util/db_middleware"
  "github.com/golang-jwt/jwt/v4"
  "github.com/labstack/echo/v4"
  "golang.org/x/crypto/bcrypt"
  "gorm.io/gorm"
  "net/http"
  "strconv"
  "strings"
  "time"
)

var hmacSampleSecret = []byte("T4FB8PH7tFhfGUXFxM3uMVvv")

func Login(c echo.Context) error {
  email := c.FormValue("email")
  password := c.FormValue("password")

  queryStr := `
    select
      u.id as id,
      u.password as password,
      u.role as role,
      u.activated as activated,
      u.user_id as user_id,
      u.gender as gender,
      u.name as name,
      u.phone1 as phone1,
      u.phone2 as phone2
    from users u
    where u.user_id = '%s'
    ;
    `
  queryStr = fmt.Sprintf(queryStr, email)
  //fmt.Println(queryStr)
  var users []response.User
  if err := db_middleware.GetDBFromEcho(c).Raw(queryStr).Find(&users).Error; err != nil {
    if errors.Is(err, gorm.ErrRecordNotFound) {
      return c.JSON(http.StatusOK, response.Data{Success: false, Data: "Invalid email or password"})
    } else {
      return c.JSON(http.StatusInternalServerError, response.Data{Success: false, Data: err.Error()})
    }
  }

  if len(users) == 0 {
    return c.JSON(http.StatusOK, response.Data{Success: false, Data: "Invalid email or password"})
  }

  if !users[0].Activated {
    return c.JSON(http.StatusOK, response.Data{Success: false, Data: "Invalid email or password"})
  }

  hashedPassword := users[0].Password
  if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
    return c.JSON(http.StatusOK, response.Data{Success: false, Data: "Invalid email or password"})
  }

  users[0].Password = ""
  var err error
  users[0].Token, err = jwtSign(users[0])
  if err != nil {
    return c.JSON(http.StatusInternalServerError, response.Data{Success: false, Data: err.Error()})
  }
  return c.JSON(http.StatusOK, response.Data{Success: true, Data: users[0]})
}

func ChangePassword(c echo.Context) error {
  //userId, err := util.GetIdFromPath(c)
  admin := c.Get("admin")
  if admin == nil {
    return c.JSON(http.StatusInternalServerError, response.Data{Success: false, Data: "Invalid user"})
  }
  adminUser := admin.(response.User)
  oldPassword := c.FormValue("old_password")
  newPassword := c.FormValue("new_password")

  queryStr := `
    select
      u.password as password
    from users u
    where u.id = %d
    ;
    `
  queryStr = fmt.Sprintf(queryStr, adminUser.ID)
  //fmt.Println(queryStr)
  var users []response.User
  if err := db_middleware.GetDBFromEcho(c).Raw(queryStr).Find(&users).Error; err != nil {
    if errors.Is(err, gorm.ErrRecordNotFound) {
      return c.JSON(http.StatusOK, response.Data{Success: false, Data: "Invalid user"})
    } else {
      return c.JSON(http.StatusInternalServerError, response.Data{Success: false, Data: err.Error()})
    }
  }

  hashedPassword := users[0].Password
  if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(oldPassword)); err != nil {
    return c.JSON(http.StatusOK, response.Data{Success: false, Data: "Invalid old password"})
  }

  hashedNewPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
  if err != nil {
    return c.JSON(http.StatusInternalServerError, response.Data{Success: false, Data: err.Error()})
  }

  queryStr = `
    update users
    set password = '%s'
    where id = %d
    ;
    `
  queryStr = fmt.Sprintf(queryStr, hashedNewPassword, adminUser.ID)
  //fmt.Println(queryStr)
  if err := db_middleware.GetDBFromEcho(c).Exec(queryStr).Error; err != nil {
    return c.JSON(http.StatusInternalServerError, response.Data{Success: false, Data: err.Error()})
  }

  return c.JSON(http.StatusOK, response.Data{Success: true, Data: true})
}

func jwtSign(user response.User) (string, error) {

  //  return JWT.sign(
  //  {
  //  id: user.id,
  //    email: user.email,
  //    full_name: user.full_name,
  //    role: user.role,
  //    expiredDate: new Date().getTime() + JWT_EXPIRED_PERIOD,
  //  },
  //  JWT_SECRET,
  //);

  token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
    "id":          user.ID,
    "email":       user.UserID,
    "full_name":   user.Name,
    "role":        user.Role,
    "expiredDate": time.Now().Add(time.Hour * 24 * 3).Unix(),
  })

  // Sign and get the complete encoded token as a string using the secret
  return token.SignedString(hmacSampleSecret)
}

func jwtVerify(token string) (jwt.Claims, error) {
  key := func(token *jwt.Token) (interface{}, error) {
    if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
      return nil, errors.New("unexpected signing method")
    }
    return hmacSampleSecret, nil
  }
  parsedToken, err := jwt.Parse(token, key)
  if err != nil {
    return nil, errors.New("Invalid session. ")
  }
  if parsedToken.Valid {
    return parsedToken.Claims, nil
  } else {
    return nil, errors.New("Invalid session. ")
  }
}

func VerifyToken(next echo.HandlerFunc) echo.HandlerFunc {
  return func(c echo.Context) error {
    token := c.Request().Header.Get("Authorization")
    if token == "" {
      return c.JSON(http.StatusUnauthorized, response.Data{Success: false, Data: "Unauthorized"})
    }

    var bearer = strings.Split(token, " ")
    //const bearerToken = bearer[0] === 'Bearer' ? bearer[1] : bearer[0];
    var bearerToken string
    if bearer[0] == "Bearer" {
      bearerToken = bearer[1]
    } else {
      bearerToken = bearer[0]
    }

    if verifiedToken, err := jwtVerify(bearerToken); err != nil {
      return c.JSON(http.StatusUnauthorized, response.Data{Success: false, Data: "Unauthorized"})
    } else {
      fmt.Println(verifiedToken)
      id_ := fmt.Sprintf("%v", verifiedToken.(jwt.MapClaims)["id"])
      id, err := strconv.Atoi(id_)
      if err != nil {
        return c.JSON(http.StatusUnauthorized, response.Data{Success: false, Data: "Unauthorized"})
      }
      var users []response.User
      if err := db_middleware.
        GetDBFromEcho(c).
        Raw(getAdminQuery(id)).
        Scan(&users).
        Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
          return c.JSON(http.StatusUnauthorized, response.Data{Success: false, Data: "Unauthorized"})
        } else {
          return c.JSON(http.StatusInternalServerError, response.Data{Success: false, Data: err.Error()})
        }
      }
      if users[0].Role != "admin" {
        return c.JSON(http.StatusUnauthorized, response.Data{Success: false, Data: "Unauthorized"})
      }
      c.Set("admin", users[0])
    }
    return next(c)
  }
}

func getAdminQuery(id int) string {
  queryStr := `
    select
      id,
      role,
      created_at,
      user_id,
      gender,
      name,
      phone1,
      phone2
    from users
    where id = %d
   `
  return fmt.Sprintf(queryStr, id)
}
