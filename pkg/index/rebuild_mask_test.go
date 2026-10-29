//ff:func feature=index type=rule control=sequence
//ff:what TestRebuildMask — unchanged tags are kept; a changed key width or a new tag is rebuilt
package index

import (
	"slices"
	"testing"
)

func TestRebuildMask(t *testing.T) {
	sp := smallSpec()
	secs := nextSections(nil, sp, fill(7, true), make([][]SegRef, 7))
	if m := rebuildMask(secs, sp); slices.Contains(m, true) {
		t.Fatal(m)
	}
	sp.Tags[1].Sort.Width = 8
	if m := rebuildMask(secs[:6], sp); !m[1] || !m[6] || m[0] {
		t.Fatal(m)
	}
}
