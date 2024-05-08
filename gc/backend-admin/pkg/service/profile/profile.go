package profile

import (
	"context"

	"git.dev.opnd.io/gc/backend-admin/pkg/model"
	"git.dev.opnd.io/gc/backend-admin/pkg/util/db_middleware"
	uuid "github.com/satori/go.uuid"
)

const (
	Authenticated            = "authenticated"
	Admin                    = "Admin"
	SuperAdmin               = "SuperAdmin"
	AccessLevelNone          = 0
	AccessLevelAuthenticated = 1
	AccessLevelAdmin         = 2
	AccessLevelSuperAdmin    = 3
)

func GetUser(ctx context.Context, userID uuid.UUID) (*model.Users, error) {
	user := &model.Users{}
	if err := db_middleware.GetDBFromContext(ctx).Model(&model.Users{}).Where("id = ? and confirmed_at is not null", userID).First(user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func IsUserAdmin(user *model.Users) bool {
	if user == nil {
		return false
		// return nil, errors.New("user is nil")
	}

	if !user.Role.Valid {
		return false
	}

	if user.Role.String != Admin && user.Role.String != SuperAdmin {
		return false
		// return nil, errors.New("user is not admin or superadmin")
	}
	return true
}

func IsUserSuperAdmin(user *model.Users) bool {
	if user == nil {
		return false
		// return nil, errors.New("user is nil")
	}

	if !user.Role.Valid {
		return false
	}

	if user.Role.String != SuperAdmin {
		return false
		// return nil, errors.New("user is not admin or superadmin")
	}
	return true
}

func GetRoleLevel(user *model.Users) int32 {
	if user == nil {
		return 0
	}

	if !user.Role.Valid {
		return 0
	}

	// superadmin
	if user.Role.String == SuperAdmin {
		return 3
	}
	// admin
	if user.Role.String == Admin {
		return 2
	}
	// authenticated user
	return 1
}
