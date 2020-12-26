package transport

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/mustafa-korkmaz/goapitemplate/pkg/api/upload"
	"github.com/mustafa-korkmaz/goapitemplate/pkg/viewmodel/request"
)

// HTTP represents upload http transport service
type HTTP struct {
	svc upload.Service
}

// New creates new upload http service with valid api versions
func New(svc upload.Service, groups ...*echo.Group) {
	h := HTTP{svc}
	v1 := groups[0].Group("/uploads")

	//define /V1/uploads methods
	v1.POST("", h.save)
}

func (h *HTTP) save(c echo.Context) error {

	model := new(request.Upload)

	if err := c.Bind(model); err != nil {
		return err
	}

	var resp = h.svc.Save(model)

	return c.JSON(http.StatusOK, resp)
}
