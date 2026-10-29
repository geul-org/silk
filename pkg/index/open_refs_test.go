//ff:func feature=index type=reader control=sequence
//ff:what TestOpenRefs — opens each referenced segment; a missing one fails
package index

import (
	"testing"
)

func TestOpenRefs(t *testing.T) {
	_, _, dir := built(t)
	m, _ := CurrentManifest(dir)
	rs := append([]SegRef{}, m.Sections[0].Runs...)
	fs, err := openRefs(dir, rs)
	if err != nil || len(fs) != 1 {
		t.Fatal(err)
	}
	fs[0].Close()
	if _, err := openRefs(dir, append(rs, SegRef{File: "seg/none"})); err == nil {
		t.Fatal("missing")
	}
}
