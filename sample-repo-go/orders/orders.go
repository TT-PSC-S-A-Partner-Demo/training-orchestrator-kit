// Package orders - one of three services that each "normalize" input its own way.
// Convention drift: this author assumed input is always meaningful, never blank.
package orders

import "strings"

// Normalize trims the input. It never complains about a blank value.
func Normalize(x string) (string, error) {
	return strings.TrimSpace(x), nil // blank -> "" with no error
}
