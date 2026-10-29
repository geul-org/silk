//ff:func feature=index type=reader control=sequence
//ff:what TestIndexRecord — the DOC-ID record at offset 0 frames; unknown documents, offsets past the end, missing packs, short reads and bad framing fail
package index

import (
	"testing"
)

func TestIndexRecord(t *testing.T) {
	a, _, dir := built(t)
	ix, _ := Open(a, dir)
	defer ix.Close()
	if rec, err := ix.Record(0, 0); err != nil || len(rec) != 4 {
		t.Fatalf("%x %v", rec, err)
	}
	if _, err := ix.Record(99, 0); err == nil {
		t.Fatal("unknown document")
	}
	if _, err := ix.Record(0, ix.Docs[0].Len); err == nil {
		t.Fatal("offset past the end")
	}
	if _, err := ix.Record(0, 1); err == nil {
		t.Fatal("odd offset frames nothing")
	}
	d := ix.Docs[0]
	ix.Docs[0].Off = 1 << 40
	if _, err := ix.Record(0, 0); err == nil {
		t.Fatal("head read past the pack")
	}
	ix.Docs[0] = d
	ix.Docs[0].Len = 2
	if _, err := ix.Record(0, 0); err == nil {
		t.Fatal("record runs past the document")
	}
	ix.Docs = append(ix.Docs, DocEntry{Pack: 0, Off: uint64(fileSize(ix.paths[0])) - 2, Len: 8})
	if _, err := ix.Record(uint32(len(ix.Docs)-1), 0); err == nil {
		t.Fatal("record read past the pack")
	}
	ix.packs[0], ix.paths[0] = nil, "/nonexistent"
	if _, err := ix.Record(0, 0); err == nil {
		t.Fatal("missing pack")
	}
}
