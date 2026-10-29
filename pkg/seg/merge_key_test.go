//ff:func feature=seg type=builder control=sequence
//ff:what TestMergeKey — sorted groups stream, an appended group is gathered and sorted; a key out of order fails
package seg

import (
	"path/filepath"
	"testing"
)

func TestMergeKey(t *testing.T) {
	s, u := sampleSeg(t, true), sampleSeg(t, false)
	gs, _ := s.Groups()
	gu, _ := u.Groups()
	w, _ := NewWriter(filepath.Join(t.TempDir(), "m.sseg"), Info{PartWidth: 2, SortWidth: 2, Sorted: true}, 1)
	defer w.Abort()
	if d, err := mergeKey([]Group{gs[0], gs[0]}, w, allLive); err != nil || d != 0 || w.parts[0].Count != 6 {
		t.Fatal(err, w.parts)
	}
	if _, err := mergeKey([]Group{gs[1], gu[1]}, w, allLive); err != nil || w.parts[1].Count != 4 {
		t.Fatal(err, w.parts)
	}
	if _, err := mergeKey([]Group{gs[0]}, w, allLive); err == nil {
		t.Fatal("key out of order")
	}
}
