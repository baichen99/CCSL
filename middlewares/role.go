package middlewares

import (
	"ccsl/configs"
	"ccsl/utils"
	"errors"
	"fmt"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

type roleConfig struct {
	Roles []string
}

// RoleMiddleware middleware used to check the JWT's role
type RoleMiddleware struct {
	roleConfig roleConfig
}

// NewRoleMiddleware returns a new user role middleware with given config
func NewRoleMiddleware(cfg ...roleConfig) *RoleMiddleware {
	var c roleConfig
	if len(cfg) == 0 {
		c = roleConfig{}
	} else {
		c = cfg[0]
	}
	return &RoleMiddleware{roleConfig: c}
}

// CheckUserRole will check the user role
func (m *RoleMiddleware) CheckUserRole(ctx context.Context) (err error) {
	_, role := GetJWTParams(ctx)
	// If user role is not super user and not defined in role config
	if HasPermisson(role, configs.RoleStudent) && !utils.StringsUnion(m.roleConfig.Roles, role) {
		err = fmt.Errorf("role '%s' does not have sufficient permissions", role)
		utils.SetError(ctx, iris.StatusForbidden, err.Error(), errors.New("RoleError"))
		return
	}
	return nil
}

// Serve is http handler
func (m *RoleMiddleware) Serve(ctx context.Context) {
	if err := m.CheckUserRole(ctx); err != nil {
		ctx.StopExecution()
		return
	}
	ctx.Next()
}

// HasPermisson checks if user has role
func HasPermisson(userRoles []string, checkRole string) bool {
	return utils.StringsContains(userRoles, checkRole) != -1
}
