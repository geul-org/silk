//ff:func feature=index type=reader control=sequence
//ff:what TestOpenGroups — the inputs' groups in key order; a missing input fails
package index

import (
	"bytes"
	"testing"

	"github.com/geul-org/silk/pkg/seg"
)

func TestOpenGroups(t *testing.T) {
	_, _, dir := updated(t)
	task, _, _ := majorTask(t, dir)
	fs, gs, err := openGroups(dir, task.Inputs)
	defer closeSegs([][]*seg.File{fs})
	if err != nil || len(fs) != 2 || len(gs) < 2 || bytes.Compare(gs[0].Key, gs[len(gs)-1].Key) > 0 {
		t.Fatal(err)
	}
	if _, _, err := openGroups(dir, []SegRef{{File: "seg/none"}}); err == nil {
		t.Fatal("missing input")
	}
}
