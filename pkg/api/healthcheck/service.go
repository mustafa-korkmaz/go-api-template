package healthcheck

import (
	"github.com/mustafa-korkmaz/goapitemplate/pkg/model"
)

// Service represents healthcheck application interface
type Service interface {
	Get(string) error
	GetPagedList(model.PagedListRequest) (model.PagedListResponse, error)
	Post(string) error
}

// HealthCheck represents healthCheck application service
type HealthCheck struct {
	foo int
	bar string
}

// Get checks the api status
func (hc *HealthCheck) Get(value string) error {
	return nil
}

// Post checks the api status
func (hc *HealthCheck) Post(value string) error {
	return nil
}

//GetPagedList returns a new pagedListResponse by the given criterias
func (hc *HealthCheck) GetPagedList(req model.PagedListRequest) (model.PagedListResponse, error) {
	var numbers = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	var resp = model.PagedListResponse{
		RecordsTotal: req.Length,
		Items:        make([]interface{}, len(numbers)),
	}

	for i, n := range numbers {
		resp.Items[i] = n
	}
	return resp, nil
}

// New creates new healthCheck application service
func New() *HealthCheck {
	return new(HealthCheck)
}
