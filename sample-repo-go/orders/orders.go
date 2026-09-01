// Package orders - part of the consistent baseline (reject blank, trim the rest).
package orders

import (
	"errors"
	"strings"
)

func Normalize(x string) (string, error) {
	t := strings.TrimSpace(x)
	if t == "" {
		return "", errors.New("input must not be blank")
	}
	return t, nil
}
