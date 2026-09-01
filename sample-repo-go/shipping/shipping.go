// Package shipping - a third take that happens to match the convention the team
// should have agreed on all along: reject blank input loudly. The others drifted.
package shipping

import (
	"errors"
	"strings"
)

// Normalize trims the input and rejects anything blank.
func Normalize(x string) (string, error) {
	t := strings.TrimSpace(x)
	if t == "" {
		return "", errors.New("input must not be blank")
	}
	return t, nil
}
