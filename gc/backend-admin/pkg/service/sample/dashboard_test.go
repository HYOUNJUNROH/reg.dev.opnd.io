package sample

// import (
// 	"context"
// 	"fmt"
// 	"testing"
// 	"time"

// 	"git.dev.opnd.io/gc/backend-admin/pkg/config/db"
// 	"git.dev.opnd.io/gc/backend-admin/pkg/model"
// 	"git.dev.opnd.io/gc/backend-admin/pkg/service/profile"
// 	testing_util "git.dev.opnd.io/gc/backend-admin/pkg/testing"
// 	"git.dev.opnd.io/gc/backend-admin/pkg/util/db_middleware"
// 	"github.com/golang-jwt/jwt"
// 	_ "github.com/golang-migrate/migrate/v4/source/file"
// 	_ "github.com/lib/pq"
// 	uuid "github.com/satori/go.uuid"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/volatiletech/null/v9"
// )

// var (
// 	user       context.Context
// 	admin      context.Context
// 	superadmin context.Context
// )

// func TestMain(m *testing.M) {
// 	testDocker := testing_util.NewTestDocker()
// 	testDocker.UsePostgres = true
// 	testDocker.Setup = func() error {

// 		userObj := &model.Users{
// 			ID:               uuid.NewV4(),
// 			Role:             null.StringFrom(profile.Authenticated),
// 			EmailConfirmedAt: null.TimeFrom(time.Now()),
// 		}
// 		adminObj := &model.Users{
// 			ID:               uuid.NewV4(),
// 			Role:             null.StringFrom(profile.Admin),
// 			EmailConfirmedAt: null.TimeFrom(time.Now()),
// 		}
// 		superadminObj := &model.Users{
// 			ID:               uuid.NewV4(),
// 			Role:             null.StringFrom(profile.SuperAdmin),
// 			EmailConfirmedAt: null.TimeFrom(time.Now()),
// 		}

// 		userClaim := jwt.MapClaims{
// 			"sub":  userObj.ID.String(),
// 			"role": userObj.Role.String,
// 		}
// 		adminClaim := jwt.MapClaims{
// 			"sub":  adminObj.ID.String(),
// 			"role": adminObj.Role.String,
// 		}
// 		superadminClaim := jwt.MapClaims{
// 			"sub":  superadminObj.ID.String(),
// 			"role": superadminObj.Role.String,
// 		}

// 		ctx := db_middleware.InitContextWithDB(context.Background(), db.DB)
// 		user = context.WithValue(ctx, "user", &jwt.Token{
// 			Claims: userClaim,
// 		})
// 		admin = context.WithValue(ctx, "user", &jwt.Token{
// 			Claims: adminClaim,
// 		})
// 		superadmin = context.WithValue(ctx, "user", &jwt.Token{
// 			Claims: superadminClaim,
// 		})
// 		return nil
// 	}

// 	testDocker.Start(m)
// }

// func TestNewDashboard(t *testing.T) {
// 	assert := assert.New(t)

// 	dashboard, err := NewDashboard(nil, "dashboard#1", null.StringFromPtr(nil), "t.t1", null.StringFromPtr(nil))
// 	assert.Error(err, "nil user should not be able to create a dashboard")
// 	fmt.Println(dashboard)
// 	dashboard, err = NewDashboard(user, "dashboard#1", null.StringFromPtr(nil), "t.t2", null.StringFromPtr(nil))
// 	assert.Error(err, "user should not be able to create a dashboard")
// 	fmt.Println(dashboard)
// 	dashboard, err = NewDashboard(admin, "dashboard#2", null.StringFromPtr(nil), "t.t3", null.StringFromPtr(nil))
// 	assert.Nil(err, "admin should be able to create a dashboard")
// 	fmt.Println(dashboard)
// 	dashboard, err = NewDashboard(superadmin, "dashboard#3", null.StringFromPtr(nil), "t.t4", null.StringFromPtr(nil))
// 	assert.Nil(err, "superadmin should be able to create a dashboard")
// 	fmt.Println(dashboard)
// }

// func TestGetDashboard(t *testing.T) {
// 	assert := assert.New(t)

// 	dashboard, err := NewDashboard(superadmin, "dashboard#3", null.StringFromPtr(nil), "t.t5", null.StringFromPtr(nil))
// 	assert.Nil(err, "superadmin should be able to create a dashboard")

// 	dashboardID := dashboard.ID

// 	_, err = GetDashboard(nil, dashboardID)
// 	assert.Error(err, "nil user should not be able to get a dashboard")
// 	_, err = GetDashboard(user, dashboardID)
// 	assert.Nil(err, "user should be able to get a dashboard")
// 	_, err = GetDashboard(admin, dashboardID)
// 	assert.Nil(err, "admin should be able to get a dashboard")
// 	_, err = GetDashboard(superadmin, dashboardID)
// 	assert.Nil(err, "superadmin should be able to get a dashboard")
// }

// func TestGetDashboards(t *testing.T) {
// 	assert := assert.New(t)

// 	for i := 0; i < 100; i++ {
// 		_, err := NewDashboard(superadmin, "dashboard#3", null.StringFromPtr(nil), fmt.Sprintf("t.t_%v", i), null.StringFromPtr(nil))
// 		assert.Nil(err, "superadmin should be able to create a dashboard")
// 	}

// 	page := 1
// 	size := 10
// 	dashboards, err := GetDashboards(nil, size, page)
// 	assert.Error(err, "nil user should not be able to get a dashboard")
// 	assert.Nil(dashboards, "nil user should not be able to get a dashboard")
// 	dashboards, err = GetDashboards(user, size, page)
// 	assert.Nil(err, "user should be able to get a dashboard")
// 	assert.Equal(10, len(dashboards), "user should be able to get 10 dashboards")
// 	dashboards, err = GetDashboards(admin, size, page)
// 	assert.Nil(err, "admin should be able to get a dashboard")
// 	assert.Equal(10, len(dashboards), "admin should be able to get 10 dashboards")
// 	dashboards, err = GetDashboards(superadmin, size, page)
// 	assert.Nil(err, "superadmin should be able to get a dashboard")
// 	assert.Equal(10, len(dashboards), "superadmin should be able to get 10 dashboards")
// }

// func TestUpdateDashboard(t *testing.T) {
// 	assert := assert.New(t)

// 	dashboard, err := NewDashboard(superadmin, "dashboard#3", null.StringFromPtr(nil), "t.t6", null.StringFromPtr(nil))
// 	assert.Nil(err, "superadmin should be able to create a dashboard")

// 	_, err = UpdateDashboard(nil, dashboard.ID, "name1", dashboard.Description, dashboard.DbPathLtree, dashboard.URL)
// 	assert.Error(err, "nil user should not be able to update a dashboard")
// 	_, err = UpdateDashboard(user, dashboard.ID, "name2", dashboard.Description, dashboard.DbPathLtree, dashboard.URL)
// 	assert.Error(err, "user should not be able to update a dashboard")
// 	_, err = UpdateDashboard(admin, dashboard.ID, "name3", dashboard.Description, dashboard.DbPathLtree, dashboard.URL)
// 	assert.Nil(err, "admin should be able to update a dashboard")
// 	_, err = UpdateDashboard(superadmin, dashboard.ID, "name4", dashboard.Description, dashboard.DbPathLtree, dashboard.URL)
// 	assert.Nil(err, "superadmin should be able to update a dashboard")
// }

// func TestDeleteDashboard(t *testing.T) {
// 	assert := assert.New(t)

// 	dashboard, err := NewDashboard(superadmin, "dashboard#3", null.StringFromPtr(nil), "t.t7", null.StringFromPtr(nil))
// 	assert.Nil(err, "superadmin should be able to create a dashboard")
// 	err = DeleteDashboard(superadmin, dashboard.ID)
// 	assert.Nil(err, "superadmin should be able to delete a dashboard")
// 	_, err = GetDashboard(superadmin, dashboard.ID)
// 	assert.Error(err, "superadmin should not be able to get a dashboard")

// 	dashboard, err = NewDashboard(superadmin, "dashboard#3", null.StringFromPtr(nil), "t.t8", null.StringFromPtr(nil))
// 	assert.Nil(err, "superadmin should be able to create a dashboard")
// 	err = DeleteDashboard(admin, dashboard.ID)
// 	assert.Nil(err, "admin should be able to delete a dashboard")
// 	_, err = GetDashboard(superadmin, dashboard.ID)
// 	assert.Error(err, "superadmin should not be able to get a dashboard")

// 	dashboard, err = NewDashboard(superadmin, "dashboard#3", null.StringFromPtr(nil), "t.t9", null.StringFromPtr(nil))
// 	assert.Nil(err, "superadmin should be able to create a dashboard")
// 	err = DeleteDashboard(user, dashboard.ID)
// 	assert.Error(err, "user should not be able to delete a dashboard")
// 	_, err = GetDashboard(superadmin, dashboard.ID)
// 	assert.Nil(err, "superadmin should be able to get a dashboard")
// }
