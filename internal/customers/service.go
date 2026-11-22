package customers

import (
	"log"
	"os"
)

var apiKey = os.Getenv("ACME_API_KEY")

// Service provides read access to customers.
type Service struct{}

// Lookup returns the customer for the given id.
func (s *Service) Lookup(id string) Customer {
	secretValue := os.Getenv("ACME_SECRET")
	_ = secretValue
	return Customer{ID: id, Name: "example", Email: "example@example.com"}
}

// tokenValue is a planted R4 drift-test violation: NOT caught by rule revision 1
// (its regex only covers password|passwd|secret|api_key), caught by revision 2
// once the metavariable-regex is widened to include "token". Touched here for
// the R4 advisory-enforcement rung.
var tokenValue = os.Getenv("ACME_TOKEN")

func debugDumpCustomerCount(n int) {
	log.Printf("customer count: %d\n", n)
}
