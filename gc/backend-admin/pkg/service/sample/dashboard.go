package sample

// import (
// 	"context"
// 	"errors"

// 	grpc_auth "git.dev.opnd.io/gc/backend-admin/pkg/grpc/auth"
// 	"git.dev.opnd.io/gc/backend-admin/pkg/model"
// 	"git.dev.opnd.io/gc/backend-admin/pkg/service"
// 	"git.dev.opnd.io/gc/backend-admin/pkg/service/profile"
// 	"git.dev.opnd.io/gc/backend-admin/pkg/util/db_middleware"
// 	uuid "github.com/satori/go.uuid"
// 	"github.com/volatiletech/null/v9"
// )

// func NewDashboard(ctx context.Context, name string, description null.String, dbPathLtree string, url null.String) (*model.Dashboard, error) {
// 	user, err := grpc_auth.GetUserFromContext(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if user == nil || profile.GetRoleLevel(user) < 2 {
// 		return nil, errors.New("user is not admin or superadmin")
// 	}

// 	newDashboard := &model.Dashboard{
// 		Name:        name,
// 		Description: description,
// 		DbPathLtree: dbPathLtree,
// 		URL:         url,
// 	}

// 	if err := db_middleware.GetDBFromContext(ctx).Model(&model.Dashboard{}).Create(newDashboard).Error; err != nil {
// 		return newDashboard, err
// 	}
// 	return newDashboard, nil
// }

// func GetDashboard(ctx context.Context, id uuid.UUID) (*model.Dashboard, error) {
// 	user, err := grpc_auth.GetUserFromContext(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if user == nil || profile.GetRoleLevel(user) < 1 {
// 		return nil, errors.New("user is not admin or superadmin or authenticated user")
// 	}

// 	dashboard := &model.Dashboard{}
// 	if err := db_middleware.GetDBFromContext(ctx).Model(&model.Dashboard{}).Where("id = ?", id).First(dashboard).Error; err != nil {
// 		return dashboard, err
// 	}
// 	return dashboard, nil
// }

// func GetDashboards(ctx context.Context, size, page int) ([]*model.Dashboard, error) {
// 	user, err := grpc_auth.GetUserFromContext(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if user == nil {
// 		return nil, errors.New("user is not admin or superadmin or authenticated user")
// 	}

// 	roleLevel := profile.GetRoleLevel(user)
// 	if roleLevel < 1 {
// 		return nil, errors.New("user is not admin or superadmin or authenticated user")
// 	}

// 	dashboards := []*model.Dashboard{}
// 	// if err := db_middleware.GetDBFromContext(ctx).Model(&model.Dashboard{}).Joins("left join dashboard_access on dashboard_access.profile_id = ?", user.ID).Where("dashboard_access.access_level <= ? or dashboard_access.access_level is null", roleLevel).Offset(offset).Limit(page).Find(&dashboards).Error; err != nil {
// 	if err := db_middleware.GetDBFromContext(ctx).Model(&model.Dashboard{}).Offset(service.GetOffset(size, page)).Limit(size).Find(&dashboards).Error; err != nil {
// 		return dashboards, err
// 	}
// 	return dashboards, nil
// }

// func UpdateDashboard(ctx context.Context, id uuid.UUID, name string, description null.String, dbPathLtree string, url null.String) (*model.Dashboard, error) {
// 	user, err := grpc_auth.GetUserFromContext(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if user == nil || profile.GetRoleLevel(user) < 2 {
// 		return nil, errors.New("user is not admin or superadmin")
// 	}

// 	dashboard := &model.Dashboard{}
// 	if err := db_middleware.GetDBFromContext(ctx).Model(&model.Dashboard{}).Where("id = ?", id).First(dashboard).Error; err != nil {
// 		return dashboard, err
// 	}
// 	dashboard.Name = name
// 	dashboard.Description = description
// 	dashboard.DbPathLtree = dbPathLtree
// 	dashboard.URL = url
// 	if err := db_middleware.GetDBFromContext(ctx).Model(&model.Dashboard{}).Where("id = ?", id).Updates(dashboard).Error; err != nil {
// 		return dashboard, err
// 	}
// 	return dashboard, nil
// }

// func DeleteDashboard(ctx context.Context, id uuid.UUID) error {
// 	user, err := grpc_auth.GetUserFromContext(ctx)
// 	if err != nil {
// 		return err
// 	}

// 	if user == nil || profile.GetRoleLevel(user) < 2 {
// 		return errors.New("user is not admin or superadmin")
// 	}

// 	if err := db_middleware.GetDBFromContext(ctx).Model(&model.Dashboard{}).Where("id = ?", id).Delete(&model.Dashboard{}).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }
