//ff:func feature=index type=builder control=sequence
//ff:what TestExpectRows — rows per tag, sorted, matching the section sizes; an unreadable or malformed document fails
package index

import (
	"context"
	"testing"
)

func TestExpectRows(t *testing.T) {
	a, _, dir := built(t)
	ix, _ := Open(a, dir)
	defer ix.Close()
	docs, _ := liveDocs(context.Background(), a)
	want, err := expectRows(ix, docs, []uint32{0, 1, 2, 3})
	if err != nil || uint64(len(want[0])/(6+6+8)) != rowsOf(ix.M.Sections[0].Runs) {
		t.Fatal(err, len(want[0]))
	}
	bad := append([]liveDoc{}, docs...)
	bad[1].Path = "/nonexistent"
	if _, err := expectRows(ix, bad, []uint32{0, 1, 2, 3}); err == nil {
		t.Fatal("unreadable")
	}
	bad[1] = docs[1]
	bad[1].Doc.Off += 2
	if _, err := expectRows(ix, bad, []uint32{0, 1, 2, 3}); err == nil {
		t.Fatal("malformed")
	}
}
