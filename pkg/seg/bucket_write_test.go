//ff:func feature=seg type=writer control=iteration dimension=1
//ff:what TestBucketWrite — sorted and appended segments verify; a bad path or a key of the wrong width fails
package seg

import (
	"os"
	"path/filepath"
	"testing"
)

func TestBucketWrite(t *testing.T) {
	for _, sorted := range []bool{true, false} {
		f := sampleSeg(t, sorted)
		if f.Info.Sorted != sorted || f.Verify() != nil || f.Info.Rows+f.Info.Other != 7 {
			t.Fatal(f.Info)
		}
	}
	d := t.TempDir()
	os.WriteFile(filepath.Join(d, "f"), nil, 0o644)
	b := NewBucket(Info{PartWidth: 2, SortWidth: 2}, 1)
	if _, err := b.Write(filepath.Join(d, "f", "x.sseg"), true); err == nil {
		t.Fatal("bad path")
	}
	b.Add([]byte("a"), row("k1", 1, 0))
	if _, err := b.Write(filepath.Join(d, "y.sseg"), true); err == nil {
		t.Fatal("key width")
	}
	if _, err := os.Stat(filepath.Join(d, "y.sseg.tmp")); !os.IsNotExist(err) {
		t.Fatal("temporary file kept")
	}
}
