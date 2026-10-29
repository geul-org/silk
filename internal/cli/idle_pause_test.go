//ff:func feature=cli type=helper control=sequence
//ff:what TestIdlePause — explains the rule when the system can be read and says it will not wait when it cannot; the hook returns at once on an idle or unreadable system
package cli

import (
	"context"
	"strings"
	"testing"

	"github.com/geul-org/silk/pkg/idle"
)

func TestIdlePause(t *testing.T) {
	var b strings.Builder
	quiet := func() (idle.Sample, error) { return idle.Sample{CPUs: 8, MemTotal: 10, MemAvail: 9}, nil }
	if err := idlePause(2, &b, quiet)(context.Background()); err != nil || !strings.Contains(b.String(), "8 CPUs") {
		t.Fatal(err, b.String())
	}
	b.Reset()
	none := func() (idle.Sample, error) { return idle.Sample{}, idle.ErrUnsupported }
	if err := idlePause(2, &b, none)(context.Background()); err != nil || !strings.Contains(b.String(), "without waiting") {
		t.Fatal(err, b.String())
	}
}
