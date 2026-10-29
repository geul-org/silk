//ff:func feature=index type=rule control=sequence
//ff:what TestVerifyOne — a built map verifies; missing map, another identity, missing pack and changed documents fail
package index

import (
	"context"
	"errors"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/silk/pkg/recmap"
)

func TestVerifyOne(t *testing.T) {
	a, dir := fixtureArchive(t)
	ps, _ := a.Sets[0].Packs(context.Background())
	tg := Target{Pack: ps[0], Issuer: "0x00010000", Docs: ps[0].Docs}
	p := filepath.Join(dir, "m.srm")
	buildOne(p, tg)
	if err := verifyOne(p, tg); err != nil {
		t.Fatal(err)
	}
	if err := verifyOne(filepath.Join(dir, "none.srm"), tg); err == nil {
		t.Fatal("missing map")
	}
	other := tg
	other.Issuer = "0x0001"
	if err := verifyOne(p, other); !errors.Is(err, ErrStale) {
		t.Fatal(err)
	}
	gone := tg
	gone.Pack.Path = filepath.Join(dir, "none.gpk")
	if err := verifyOne(p, gone); err == nil {
		t.Fatal("missing pack")
	}
	fewer := tg
	fewer.Docs = tg.Docs[:2]
	if err := verifyOne(p, fewer); !errors.Is(err, recmap.ErrMismatch) {
		t.Fatal(err)
	}
}
