// Convention every service must share: Normalize trims valid input and returns an
// error on anything blank. orders, billing and shipping already agree.
//
// invoices is the new service. It was scaffolded to the spec, which only asked to trim
// - so it drifts: it never rejects a blank value, and these tests catch it. The root
// cause is the spec, which never stated the rule the siblings already assume. Fix it
// there (rewind to spec), then apply once.
package conventions

import (
	"testing"

	"driftdemo/billing"
	"driftdemo/invoices"
	"driftdemo/orders"
	"driftdemo/shipping"
)

type normalizer func(string) (string, error)

var services = map[string]normalizer{
	"orders":   orders.Normalize,
	"billing":  billing.Normalize,
	"shipping": shipping.Normalize,
	"invoices": invoices.Normalize,
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
