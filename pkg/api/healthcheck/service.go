package healthcheck

// Service represents healthcheck application interface
type Service interface {
	Get(string) error
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

// New creates new healthCheck application service
func New() *HealthCheck {
	return new(HealthCheck)
}
