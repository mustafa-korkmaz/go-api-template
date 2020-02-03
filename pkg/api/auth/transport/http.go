package transport

import (
	"github.com/mustafa-korkmaz/goapitemplate/pkg/api/auth"

	"github.com/labstack/echo"
	"github.com/mustafa-korkmaz/goapitemplate/pkg/viewmodel/request"
)

// HTTP represents olive http transport service
type HTTP struct {
	svc auth.Service
}

// New creates new olive http service with valid api versions
func New(svc auth.Service, mw echo.MiddlewareFunc, groups ...*echo.Group) {
	h := HTTP{svc}
	v1 := groups[0]

	//define /V1/auth methods
	v1.GET("/refresh", h.refresh, mw)
	v1.POST("/login", h.login)
	v1.POST("/register", h.register)
}

func (h *HTTP) login(c echo.Context) error {

	cred := new(request.Login)

	if err := c.Bind(cred); err != nil {
		return err
	}

	var resp = h.svc.Authenticate(cred.UsernameOrEmail, cred.Password)

	return c.JSON(resp.GetStatusCode(), resp)
}

func (h *HTTP) register(c echo.Context) error {

	registerRequest := new(request.Register)

	if err := c.Bind(registerRequest); err != nil {
		return err
	}

	var resp = h.svc.Register(registerRequest)

	return c.JSON(resp.GetStatusCode(), resp)
}

func (h *HTTP) refresh(c echo.Context) error {

	var resp = h.svc.Refresh(c)

	return c.JSON(resp.GetStatusCode(), resp)
}
