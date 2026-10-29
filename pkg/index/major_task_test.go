//ff:func feature=index type=helper control=sequence
//ff:what majorTask — test helper: a major merge of section 0 (TRIPLE) of the current manifest with a reserved output, and the live bitmap
package index

import (
	"path/filepath"
	"testing"
)

// majorTask also returns the manifest it was planned on.
func majorTask(t testing.TB, dir string) (Task, Bitmap, Manifest) {
	t.Helper()
	m, err := CurrentManifest(dir)
	if err != nil {
		t.Fatal(err)
	}
	live, _ := readBitmap(filepath.Join(dir, m.Live))
	s := m.Sections[0]
	task := Task{Section: 0, Name: s.Name, Kind: TaskMajor, Inputs: append(append([]SegRef{}, s.Runs...), s.Appends...), Min: 2}
	task.Output = m.newName(dirSeg, ".sseg")
	return task, live, m
}
