//ff:func feature=index type=rule control=sequence
//ff:what TestIndexCheckSection — a section's live rows equal the expected rows; a kind missing from the spec, a footer that disagrees, missing rows and a wrong dead count fail
package index

import (
	"context"
	"errors"
	"testing"
)

func TestIndexCheckSection(t *testing.T) {
	a, _, dir := built(t)
	ix, _ := Open(a, dir)
	defer ix.Close()
	docs, _ := liveDocs(context.Background(), a)
	want, _ := expectRows(ix, docs, []uint32{0, 1, 2, 3})
	if n, err := ix.checkSection(0, want[0]); err != nil || n == 0 {
		t.Fatal(err)
	}
	if _, err := ix.checkSection(0, want[0][20:]); !errors.Is(err, ErrMismatch) {
		t.Fatal(err)
	}
	ix.M.Sections[0].Dead = 1
	if _, err := ix.checkSection(0, want[0]); !errors.Is(err, ErrMismatch) {
		t.Fatal(err)
	}
	ix.M.Sections[0].Dead = 0
	ix.Spec.Tags[0].Sort.Width = 8
	if _, err := ix.checkSection(0, want[0]); !errors.Is(err, ErrMismatch) {
		t.Fatal(err)
	}
	ix.M.Sections[0].Tag = 0x0801
	if _, err := ix.checkSection(0, want[0]); !errors.Is(err, ErrMismatch) {
		t.Fatal(err)
	}
}
