//ff:func feature=index type=reader control=sequence
//ff:what TestIndexClose — closes every part; a closed map or a closed data file is reported
package index

import (
	"context"
	"testing"
)

func TestIndexClose(t *testing.T) {
	a, dir := fixtureArchive(t)
	Build(context.Background(), a, dir)
	ix, _ := Open(a, dir)
	if err := ix.Close(); err != nil {
		t.Fatal(err)
	}
	if err := ix.Close(); err == nil {
		t.Fatal("closed map must fail")
	}
	ix2, _ := Open(a, dir)
	ix2.Parts[0].Data.Close()
	if err := ix2.Close(); err == nil {
		t.Fatal("closed data file must fail")
	}
}
