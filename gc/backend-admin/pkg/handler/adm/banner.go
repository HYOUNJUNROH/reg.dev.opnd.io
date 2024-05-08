package adm

import (
  "fmt"
  "git.dev.opnd.io/gc/backend-admin/pkg/model"
  "git.dev.opnd.io/gc/backend-admin/pkg/model/request"
  "git.dev.opnd.io/gc/backend-admin/pkg/model/response"
  "git.dev.opnd.io/gc/backend-admin/pkg/service"
  "github.com/google/uuid"
  "github.com/labstack/echo/v4"
  "io"
  "net/http"
  "os"
  "path/filepath"
  "strconv"
)

// /usr/src/app/uploads

func GetBanners(c echo.Context) error {
  resData := response.Data{}
  banners, err := service.GetBanners(c)

  if err != nil {
    resData.Success = false
    resData.Data = err.Error()
    return c.JSON(http.StatusOK, resData)
  }
  resData.Success = true
  resData.Data = model.Banners{Banners: banners}
  return c.JSON(http.StatusOK, resData)
}

func PostBanner(c echo.Context) error {
  tag := "PostBanner"
  resData := response.Data{}

  _id := c.FormValue("id")
  id, err := strconv.Atoi(_id)
  if err != nil {
    id = 0
  }
  createdAt := c.FormValue("created_at")
  updatedAt := c.FormValue("updated_at")
  description := c.FormValue("description")
  image := c.FormValue("image")
  _priority := c.FormValue("priority")
  priority, err := strconv.Atoi(_priority)
  if err != nil {
    fmt.Println(fmt.Sprintf("[%s] priority 변환 error: %s", tag, err.Error()))
    resData.Success = false
    resData.Data = err.Error()
    return c.JSON(http.StatusOK, resData)
  }

  {
    form, err := c.MultipartForm()
    if err != nil {
      fmt.Println(fmt.Sprintf("[%s] priority 변환 error: %s", tag, err.Error()))
      resData.Success = false
      resData.Data = err.Error()
      return c.JSON(http.StatusOK, resData)
    }
    defer func() {
      err = form.RemoveAll()
      if err != nil {
        fmt.Println(fmt.Sprintf("[%s] 임시 Form 파일 삭제 error: %s", tag, err.Error()))
      }
    }()

    for _, file := range form.File {
      src, err := file[0].Open()
      if err != nil {
        fmt.Println(fmt.Sprintf("[%s] 폼 파일 열기 error: %s", tag, err.Error()))
        resData.Success = false
        resData.Data = err.Error()
        return c.JSON(http.StatusOK, resData)
      }
      //goland:noinspection GoDeferInLoop
      defer func() {
        _ = src.Close()
      }()

      // 파일 경로
      // /usr/src/app/uploads
      ext := filepath.Ext(file[0].Filename)
      u := uuid.New()
      fileName := u.String() + ext
      //fullPath := filepath.Join("./test", fileName)
      fullPath := filepath.Join("/usr/src/app/uploads", fileName)
      image = fileName
      // Destination
      dst, err := os.Create(fullPath)
      if err != nil {
        resData.Success = false
        resData.Data = err.Error()
        return c.JSON(http.StatusOK, resData)
      }
      //goland:noinspection GoDeferInLoop
      defer func() {
        _ = dst.Close()
      }()

      // Copy
      if _, err = io.Copy(dst, src); err != nil {
        resData.Success = false
        resData.Data = err.Error()
        return c.JSON(http.StatusOK, resData)
      }
    }
  }

  {
    banner := model.Banner{
      Description: description,
      Image:       image,
      Priority:    priority,
    }
    if id != 0 {
      banner.Id = id
    }
    if createdAt != "" {
      banner.CreatedAt = createdAt
    }
    if updatedAt != "" {
      banner.UpdatedAt = updatedAt
    }

    err := service.PostBanner(banner, c)
    if err != nil {
      resData.Success = false
      resData.Data = err.Error()
      return c.JSON(http.StatusOK, resData)
    }
  }

  resData.Success = true
  return c.JSON(http.StatusOK, resData)
}

func PostPriorityBanner(c echo.Context) error {
  resData := response.Data{}

  var bannerPriorities request.BannerPriorities
  err := c.Bind(&bannerPriorities)
  if err != nil {
    resData.Success = false
    resData.Data = err.Error()
    return c.JSON(http.StatusOK, resData)
  }

  err = service.PostPriorityBanner(bannerPriorities, c)
  if err != nil {
    resData.Success = false
    resData.Data = err.Error()
    return c.JSON(http.StatusOK, resData)
  }

  resData.Success = true
  return c.JSON(http.StatusOK, resData)
}

func DeleteBanner(c echo.Context) error {
  resData := response.Data{}

  _id := c.FormValue("id")
  id, err := strconv.Atoi(_id)
  if err != nil {
    resData.Success = false
    resData.Data = err.Error()
    return c.JSON(http.StatusOK, resData)
  }

  err = service.DeleteBanner(id, c)
  if err != nil {
    resData.Success = false
    resData.Data = err.Error()
    return c.JSON(http.StatusOK, resData)
  }

  resData.Success = true
  return c.JSON(http.StatusOK, resData)
}
