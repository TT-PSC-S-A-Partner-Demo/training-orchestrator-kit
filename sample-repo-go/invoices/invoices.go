// Package invoices - the new service, scaffolded to the spec, which only asked to trim.
// It drifts from the sibling convention: it never rejects a blank value.
package invoices

import "strings"

func Normalize(x string) (string, error) {
	return strings.TrimSpace(x), nil // spec-faithful, but blank is not rejected
}
