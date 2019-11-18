package healthcheck

// Service represents healthcheck application interface
type Service interface {
	Get(string) error
}

// Get checks the api status
func Get(value string) error {
	return nil
}
