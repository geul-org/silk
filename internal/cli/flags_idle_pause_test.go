//ff:func feature=cli type=scanner control=sequence
//ff:what TestFlagsIdlePause — none set is no hook; half set is refused; a full rule waits
package cli

import (
	"context"
	"errors"
	"testing"

	"github.com/geul-org/silk/pkg/idle"
)

func TestFlagsIdlePause(t *testing.T) {
	ctx := context.Background()
	if p, err := (&flags{}).idlePause(ctx); p != nil || err != nil {
		t.Fatal(err)
	}
	if _, err := (&flags{idleCPU: .5}).idlePause(ctx); !errors.Is(err, idle.ErrRule) {
		t.Fatal(err)
	}
	p, err := (&flags{idleCPU: .01, idleMem: .01}).idlePause(ctx)
	if err != nil || p == nil {
		t.Fatal(err)
	}
}
