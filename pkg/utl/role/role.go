//Package role manages role-based authorizations
package role

import (
	"github.com/labstack/echo"
	"github.com/mustafa-korkmaz/goapitemplate/pkg/enum"
	"github.com/mustafa-korkmaz/goapitemplate/pkg/model"
)

// User returns user data stored in jwt token
func User(c echo.Context) *model.AuthUser {
	id := c.Get("id").(string)
	username := c.Get("u").(string)
	email := c.Get("e").(string)
	accessLevel := c.Get("al").(enum.AccessLevelType)
	return &model.AuthUser{
		ID:          id,
		Username:    username,
		Email:       email,
		AccessLevel: accessLevel,
	}
}

// AuthorizeRole authorizes request by AccessRole
func AuthorizeRole(c echo.Context, al enum.AccessLevelType) bool {
	accessLevel := c.Get("al").(enum.AccessLevelType)
	return accessLevel > al
}

func isAdmin(c echo.Context) bool {
	accessLevel := c.Get("al").(enum.AccessLevelType)
	return accessLevel == enum.AccessLevel.Admin
}

// AuthorizeUser checks whether the request userId and the given id the same
// This can be used as a control e.g, while deleting a record which does not belong the request user
func AuthorizeUser(c echo.Context, ownerID string) bool {

	// to allow admin
	// if isAdmin(c) {
	// 	return true
	// }

	return c.Get("id").(string) == ownerID
}
