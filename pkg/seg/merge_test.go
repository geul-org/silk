//ff:func feature=seg type=builder control=sequence
//ff:what TestMerge — a sorted and an appended segment merge into one sorted run without dead rows; resume skips finished keys; after() and writer errors stop
package seg

import (
	"errors"
	"io"
	"path/filepath"
	"slices"
	"testing"
)

func TestMerge(t *testing.T) {
	a, b := sampleSeg(t, true), sampleSeg(t, false)
	ga, _ := a.Groups()
	gb, _ := b.Groups()
	gs := append(ga, gb...)
	SortGroups(gs)
	p := filepath.Join(t.TempDir(), "m.sseg")
	w, _ := NewWriter(p, Info{PartWidth: 2, SortWidth: 2, Sorted: true}, 3)
	var keys []string
	dropped, err := Merge(gs, w, func(d uint32) bool { return d != 2 }, nil, func(k []byte) error { keys = append(keys, string(k)); return nil })
	w.Finish()
	m, _ := Open(p)
	defer m.Close()
	if err != nil || dropped != 2 || !slices.Equal(keys, []string{"aa", "bb", "cc", "dd"}) || m.Info.Rows != 8 || m.Info.Other != 4 || m.Verify() != nil {
		t.Fatal(err, dropped, keys, m.Info)
	}
	w2, _ := NewWriter(filepath.Join(t.TempDir(), "r.sseg"), Info{PartWidth: 2, SortWidth: 2, Sorted: true}, 1)
	defer w2.Abort()
	keys = nil
	Merge(gs, w2, allLive, []byte("bb"), func(k []byte) error { keys = append(keys, string(k)); return nil })
	if !slices.Equal(keys, []string{"cc", "dd"}) {
		t.Fatal(keys)
	}
	if _, err := Merge(gs, w2, allLive, nil, nil); !errors.Is(err, ErrOrder) {
		t.Fatal(err)
	}
	w3, _ := NewWriter(filepath.Join(t.TempDir(), "s.sseg"), Info{PartWidth: 2, SortWidth: 2, Sorted: true}, 1)
	defer w3.Abort()
	if _, err := Merge(gs, w3, allLive, nil, func([]byte) error { return io.EOF }); !errors.Is(err, io.EOF) {
		t.Fatal(err)
	}
}
