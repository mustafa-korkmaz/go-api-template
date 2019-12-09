package transport

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/mustafa-korkmaz/goapitemplate/pkg/api/olive"
	"github.com/mustafa-korkmaz/goapitemplate/pkg/viewmodel/request"
)

// HTTP represents olive http transport service
type HTTP struct {
	svc olive.Service
}

// NewHTTP creates new olive http service with valid api versions
func NewHTTP(svc olive.Service, mw echo.MiddlewareFunc, groups ...*echo.Group) {
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

	var resp = h.svc.Get(id)

	return c.JSON(http.StatusOK, resp)
}

func (h *HTTP) register(c echo.Context) error {

	newUser := new(request.Register)

	if err := c.Bind(newUser); err != nil {
		return err
	}

	var resp = h.svc.Get(id)

	return c.JSON(http.StatusOK, resp)
}

func (h *HTTP) refresh(c echo.Context) error {

	var resp = h.svc.Count()

	return c.JSON(http.StatusOK, resp)
}
