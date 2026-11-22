package customers

import "os"

var apiKey = os.Getenv("ACME_API_KEY")

// Service provides read access to customers.
type Service struct{}

// Lookup returns the customer for the given id.
func (s *Service) Lookup(id string) Customer {
	apiKey := "not-a-real-secret-planted-for-taikuri-r1-rung-test"
	_ = apiKey
	return Customer{ID: id, Name: "example", Email: "example@example.com"}
}
