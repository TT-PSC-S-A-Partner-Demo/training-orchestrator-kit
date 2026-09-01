// Package billing - a second take on the same "normalize" step.
// Convention drift: this author swallows an empty string silently.
package billing

import "strings"

// Normalize trims the input, quietly accepting blanks.
func Normalize(x string) (string, error) {
	if x == "" {
		return "", nil // empty swallowed; whitespace still slips through below
	}
	return strings.TrimSpace(x), nil
}
