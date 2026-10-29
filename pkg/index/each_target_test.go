//ff:func feature=index type=helper control=sequence
//ff:what TestEachTarget — the fixture has one target with four documents; Packs and fn errors stop the walk
package index

import (
	"context"
	"errors"
	"testing"

	"github.com/park-jun-woo/silk/pkg/archive"
)

func TestEachTarget(t *testing.T) {
	ctx := context.Background()
	a, _ := fixtureArchive(t)
	var got []Target
	if err := eachTarget(ctx, a, func(x Target) error { got = append(got, x); return nil }); err != nil || len(got) != 1 || len(got[0].Docs) != 4 {
		t.Fatal(got, err)
	}
	stop := errors.New("stop")
	if err := eachTarget(ctx, a, func(Target) error { return stop }); err != stop {
		t.Fatal(err)
	}
	bad := &archive.Archive{Sets: []archive.Set{{Dir: t.TempDir(), Manifest: "none"}}}
	if err := eachTarget(ctx, bad, func(Target) error { return nil }); err == nil {
		t.Fatal("missing manifest must fail")
	}
}
