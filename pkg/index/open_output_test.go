//ff:func feature=index type=writer control=sequence
//ff:what TestOpenOutput — a new sorted run from the input's footer, or the checkpointed writer with its last key and bytes
package index

import (
	"path/filepath"
	"testing"

	"github.com/geul-org/silk/pkg/seg"
)

func TestOpenOutput(t *testing.T) {
	d := t.TempDir()
	in := seg.Info{Tag: 0x100A, PartWidth: 6, SortWidth: 6, Spec: [32]byte{3}}
	w, from, base, err := openOutput(d, Task{Output: "seg/a.sseg", Min: 2}, in, nil)
	if err != nil || from != nil || base != 0 || w.Min != 2 {
		t.Fatal(err)
	}
	st, _ := w.Suspend()
	w2, from2, base2, err := openOutput(d, Task{}, in, &checkpoint{Last: []byte("k"), State: st})
	if err != nil || string(from2) != "k" || base2 != st.Written || w2.Path != filepath.Join(d, "seg/a.sseg") {
		t.Fatal(err)
	}
	w2.Abort()
}
