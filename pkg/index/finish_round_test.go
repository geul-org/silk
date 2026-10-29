//ff:func feature=index type=writer control=sequence
//ff:what TestFinishRound — finished merges are applied and committed, paused ones saved, failed ones reported
package index

import (
	"errors"
	"io"
	"testing"
)

func TestFinishRound(t *testing.T) {
	d := t.TempDir()
	m := Manifest{Sections: []Section{{Name: "A", Runs: []SegRef{{File: "a"}}}, {Name: "B"}}}
	outs := []roundOut{
		{res: taskResult{Task: Task{Section: 0, Name: "A", Kind: TaskMajor, Inputs: []SegRef{{File: "a"}}}, Out: SegRef{File: "b", Rows: 2}, Written: 10}},
		{res: taskResult{Checkpoint: &checkpoint{Section: "B"}, Written: 3}},
		{err: io.EOF},
	}
	var rep Report
	err := finishRound(d, &m, outs, &rep)
	if !errors.Is(err, ErrPaused) || !errors.Is(err, io.EOF) || m.Generation != 1 || m.Sections[0].Runs[0].File != "b" || len(rep.Tasks) != 1 || rep.Written < 13 {
		t.Fatal(err, m, rep)
	}
}
