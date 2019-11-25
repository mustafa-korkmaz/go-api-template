package transport

import (
	"net/http"

	"github.com/mustafa-korkmaz/goapitemplate/pkg/enum"

	"github.com/labstack/echo"
	"github.com/mustafa-korkmaz/goapitemplate/pkg/api/healthcheck"
	"github.com/mustafa-korkmaz/goapitemplate/pkg/model"
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
	hr.POST("/paginationtest", h.getPagedList)
	hr.POST("", h.post)
}

func (h *HTTP) get(c echo.Context) error {

	val := c.Param("value")

	if err := h.svc.Get(val); err != nil {
		return err
	}

	var resp = model.APIResponse{}
	resp.Code = enum.ResponseCode.Success

	resp.Data = struct {
		AwesomeCars []string `json:"awesome_cars"`
		Value       string   `json:"value"`
	}{
		[]string{"Wv", "Jaguar", "Tesla"},
		val,
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *HTTP) getPagedList(c echo.Context) error {

	req := model.PagedListRequest{}
	if err := c.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}

	var resp = model.APIResponse{}
	resp.Code = enum.ResponseCode.Success

	var pagedListResp, err = h.svc.GetPagedList(req)

	if err != nil {
		return err
	}

	resp.Data = pagedListResp

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

	var resp = model.APIResponse{}

	resp.Code = enum.ResponseCode.Success

	resp.Data = struct {
		AwesomePhones []string `json:"awesome_phones"`
		Value         string   `json:"value"`
	}{
		[]string{"iPhone", "Samsung", "Huawei"},
		req.Value,
	}

	return c.JSON(http.StatusOK, resp)
}

// HealthCheckReq represents body of HealthCheck request.
type HealthCheckReq struct {
	Value string `json:"value" validate:"required,min=8"`
}
