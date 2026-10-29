//ff:func feature=seg type=helper control=sequence
//ff:what sampleSeg — test helper: a segment with partitions aa (3 rows) and cc (2) and bb, dd (1 row each) in the other partition (min 2)
package seg

import (
	"path/filepath"
	"testing"
)

// sampleSeg writes and opens the sample, sorted or appended.
func sampleSeg(t testing.TB, sorted bool) *File {
	t.Helper()
	b := NewBucket(Info{Tag: 0x100A, PartWidth: 2, SortWidth: 2}, 2)
	b.Add([]byte("aa"), row("x2", 3, 9))
	b.Add([]byte("aa"), row("x1", 1, 4))
	b.Add([]byte("aa"), row("x1", 2, 4))
	b.Add([]byte("bb"), row("y1", 5, 0))
	b.Add([]byte("cc"), row("z1", 1, 8))
	b.Add([]byte("cc"), row("z0", 4, 2))
	b.Add([]byte("dd"), row("w1", 6, 0))
	p := filepath.Join(t.TempDir(), "s.sseg")
	if _, err := b.Write(p, sorted); err != nil {
		t.Fatal(err)
	}
	f, err := Open(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { f.Close() })
	return f
}
