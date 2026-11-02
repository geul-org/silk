//ff:func feature=archive type=scanner control=sequence
//ff:what TestSetSteps — data packs replay version steps in order; versions.geul-like paths are skipped; records read back; object lookup
package archive

import (
	"context"
	"errors"
	"testing"
)

func TestSetSteps(t *testing.T) {
	_, a := testArchive(t)
	ctx := context.Background()
	s := a.Sets[0]
	ds := s.Data()
	if len(ds) != 1 {
		t.Fatal(ds)
	}
	if o, ok := s.ObjectOf(ds[0].Hash); !ok || o != ds[0].Obj {
		t.Fatal(o, ok)
	}
	if _, ok := s.ObjectOf("nope"); ok {
		t.Fatal("no such pack")
	}
	var ids []uint64
	var first Step
	err := s.Steps(ctx, ds[0].Obj, func(st Step) error {
		first = st
		ids = append(ids, st.ID)
		return nil
	})
	if err != nil || len(ids) != 2 {
		t.Fatal(ids, err)
	}
	b, err := s.Record(ctx, first.Loc)
	if err != nil || len(b) == 0 {
		t.Fatal(err)
	}
	boom := errors.New("stop")
	if err := s.Steps(ctx, ds[0].Obj, func(Step) error { return boom }); !errors.Is(err, boom) {
		t.Fatal(err)
	}
}
