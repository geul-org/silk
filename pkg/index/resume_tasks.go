//ff:func feature=index type=builder control=iteration dimension=1
//ff:what resumeTasks — tasks for the checkpoints whose section still holds exactly their inputs; stale checkpoints lose their unfinished output
package index

import (
	"os"
	"path/filepath"
)

// resumeTasks also moves m.NextFile past every output a checkpoint holds.
func resumeTasks(dir string, m *Manifest, cs []checkpoint) ([]Task, map[string]*checkpoint) {
	var ts []Task
	by := map[string]*checkpoint{}
	for _, c := range cs {
		i, refs, ok := findInputs(m.Sections, c.Section, c.Inputs)
		if !ok {
			os.Remove(filepath.Join(dir, c.Output+".tmp"))
			continue
		}
		m.NextFile = max(m.NextFile, fileNumber(c.Output))
		ts = append(ts, Task{Section: i, Name: c.Section, Kind: c.Kind, Inputs: refs, Output: c.Output, Min: c.State.Min})
		by[c.Output] = &c
	}
	return ts, by
}
