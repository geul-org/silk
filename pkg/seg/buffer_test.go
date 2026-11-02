//ff:func feature=seg type=scanner control=iteration dimension=1
//ff:what TestBuffer — rows sort by partition and row, duplicates can go, memory is counted, reset and free empty it
package seg

import (
	"bytes"
	"errors"
	"path/filepath"
	"testing"
)

func TestBuffer(t *testing.T) {
	b := NewBuffer(testInfo())
	for _, r := range []struct {
		k string
		r []byte
	}{{"pb", row(0, 1, 1)}, {"pa", row(0, 2, 2)}, {"pa", row(0, 1, 3)}, {"pa", row(0, 1, 3)}} {
		b.Add([]byte(r.k), r.r)
	}
	o := NewBuffer(testInfo())
	o.Absorb(b)
	if b.Rows() != 0 || o.Rows() != 4 || o.Size() != 4*8+16 {
		t.Fatal(b.Rows(), o.Rows(), o.Size())
	}
	o.sort()
	o.sort()
	if !bytes.Equal(o.at(o.idx[0])[2:], row(0, 1, 3)) {
		t.Fatal(o.at(o.idx[0]))
	}
	o.Unique()
	if o.Rows() != 3 {
		t.Fatal(o.Rows())
	}
	p := filepath.Join(t.TempDir(), "s")
	w, _ := NewWriter(p, testInfo())
	if err := o.WriteTo(w); err != nil {
		t.Fatal(err)
	}
	w.Finish()
	f, _ := Open(p)
	defer f.Close()
	if f.Rows != 3 || f.Parts != 2 || f.Verify() != nil {
		t.Fatal(f.Rows)
	}
	w2, _ := NewWriter(filepath.Join(t.TempDir(), "x"), Info{RowWidth: 8})
	defer w2.Abort()
	if err := o.WriteTo(w2); !errors.Is(err, ErrFormat) {
		t.Fatal(err)
	}
	o.Reset()
	if o.Rows() != 0 {
		t.Fatal("reset")
	}
	o.Add([]byte("pa"), row(0, 0, 0))
	o.Free()
	if o.Rows() != 0 || o.idx != nil {
		t.Fatal("free")
	}
}
