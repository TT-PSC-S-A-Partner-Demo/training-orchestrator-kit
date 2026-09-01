// Conventions the three services are supposed to share.
//
// Same contract everywhere: Normalize trims valid input, and returns an error on
// anything blank. Today orders and billing drift from the "reject blank" rule that
// shipping follows. These tests make the drift visible.
//
// The root cause is not a bug in any one service - the rule was never written down,
// so three authors each guessed. Fix it in the spec, then apply once.
package conventions

import (
	"testing"

	"driftdemo/billing"
	"driftdemo/orders"
	"driftdemo/shipping"
)

type normalizer func(string) (string, error)

var services = map[string]normalizer{
	"orders":   orders.Normalize,
	"billing":  billing.Normalize,
	"shipping": shipping.Normalize,
}

func TestValidInputIsTrimmed(t *testing.T) {
	for name, normalize := range services {
		got, err := normalize("  hello  ")
		if err != nil || got != "hello" {
			t.Errorf("%s: Normalize(%q) = %q, %v; want %q, nil", name, "  hello  ", got, err, "hello")
		}
	}
}

func TestBlankIsRejected(t *testing.T) {
	for _, in := range []string{"", " ", "   ", "\t"} {
		for name, normalize := range services {
			if _, err := normalize(in); err == nil {
				t.Errorf("%s: Normalize(%q) returned no error; want an error", name, in)
			}
		}
	}
}
