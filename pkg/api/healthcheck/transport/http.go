package transport

import (
	"net/http"

	"github.com/mustafa-korkmaz/goapitemplate/pkg/api/healthcheck"

	"github.com/labstack/echo"
)

// HTTP represents healthcheck http transport service
type HTTP struct {
	svc healthcheck.Service
}

// NewHTTP creates new healthcheck http service
func NewHTTP(svc healthcheck.Service, er *echo.Group) {
	h := HTTP{svc}
	pr := er.Group("/healthcheck")

	pr.GET("/", h.get)
}

func (h *HTTP) get(c echo.Context) error {

	p := new(HealthCheckReq)
	if err := h.svc.Get(p.Value); err != nil {
		return err
	}

	var resp = HealthCheckResp{}
	resp.Result.Value = p.Value
	return c.JSON(http.StatusOK, resp)
}

// HealthCheckReq represents body of HealthCheck request.
type HealthCheckReq struct {
	Value string `json:"value" validate:"required,min=8"`
}

// HealthCheckResp represents body of HealthCheck response.
type HealthCheckResp struct {
	Result struct {
		Value string `json:"value"`
	} `json:"result"`
}
