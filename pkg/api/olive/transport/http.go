package transport

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/mustafa-korkmaz/goapitemplate/pkg/api/olive"
)

// HTTP represents olive http transport service
type HTTP struct {
	svc olive.Service
}

// NewHTTP creates new olive http service with valid api versions
func NewHTTP(svc olive.Service, groups ...*echo.Group) {
	h := HTTP{svc}
	v1 := groups[0].Group("/olive")

	//define /V1/olive methods
	v1.GET("/:id", h.get)
	v1.GET("/count", h.count)
}

func (h *HTTP) get(c echo.Context) error {

	id := c.Param("id")

	var resp = h.svc.Get(id)

	return c.JSON(http.StatusOK, resp)
}

func (h *HTTP) count(c echo.Context) error {

	var resp = h.svc.Count()

	return c.JSON(http.StatusOK, resp)
}
