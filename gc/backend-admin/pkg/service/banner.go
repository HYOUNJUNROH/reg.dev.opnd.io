package service

import (
  "errors"
  "git.dev.opnd.io/gc/backend-admin/pkg/model"
  "git.dev.opnd.io/gc/backend-admin/pkg/model/request"
  "git.dev.opnd.io/gc/backend-admin/pkg/util/db_middleware"
  "github.com/labstack/echo/v4"
  "gorm.io/gorm"
)

func GetBanners(ctx echo.Context) ([]model.Banner, error) {
  var banners []model.Banner
  if err := db_middleware.GetDBFromEcho(ctx).Table("banner").Find(&banners).Error; err != nil {
    if errors.Is(err, gorm.ErrRecordNotFound) {
      return nil, nil
    } else {
      return nil, err
    }
  }
  return banners, nil
}

func PostBanner(banner model.Banner, ctx echo.Context) error {
  if err := db_middleware.GetDBFromEcho(ctx).Table("banner").Save(&banner).Error; err != nil {
    return err
  }
  return nil
}

func PostPriorityBanner(priorities request.BannerPriorities, ctx echo.Context) error {
  for _, priority := range priorities.BannerPriorities {
    if err := db_middleware.GetDBFromEcho(ctx).Table("banner").Where("id = ?", priority.Id).Update("priority", priority.Priority).Error; err != nil {
      return err
    }
  }
  return nil
}

func DeleteBanner(bannerId int, ctx echo.Context) error {
  if err := db_middleware.GetDBFromEcho(ctx).Table("banner").Delete(&model.Banner{}, bannerId).Error; err != nil {
    return err
  }
  return nil
}
