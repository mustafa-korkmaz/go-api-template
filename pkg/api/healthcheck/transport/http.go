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
	hr := er.Group("/healthcheck")

	hr.GET("/:value", h.get)
	hr.POST("", h.post)
}

func (h *HTTP) get(c echo.Context) error {

	val := c.Param("value")

	if err := h.svc.Get(val); err != nil {
		return err
	}

	var resp = HealthCheckResp{}
	resp.Result.Value = val
	return c.JSON(http.StatusOK, resp)
}

func (h *HTTP) post(c echo.Context) error {

	req := HealthCheckReq{}
	if err := c.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}

	if err := h.svc.Post(req.Value); err != nil {
		return err
	}

	var resp = HealthCheckResp{}
	resp.Result.Value = req.Value
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
