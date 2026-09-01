// Package billing - same convention as orders, different style.
package billing

import (
	"errors"
	"strings"
)

func Normalize(x string) (string, error) {
	if strings.TrimSpace(x) == "" {
		return "", errors.New("input must not be blank")
	}
	return strings.TrimSpace(x), nil
}
