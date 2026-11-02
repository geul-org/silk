//ff:func feature=fixture type=scanner control=sequence
//ff:what TestPack — documents pack raw and zstd; Repack adds a generation; the corpus is reproducible
package fixture

import (
	"context"
	"reflect"
	"testing"

	"github.com/geul-org/gogeul/pkg/pack"
	"github.com/geul-org/silk/pkg/archive"
)

func TestPack(t *testing.T) {
	ctx := context.Background()
	docs := Corpus(1, 40)
	if !reflect.DeepEqual(docs, Corpus(1, 40)) || len(Grammar(t).Fields) == 0 {
		t.Fatal("corpus or grammar")
	}
	z, err := pack.ParseCodec("zstd")
	if err != nil {
		t.Fatal(err)
	}
	data := Pack(t, docs, z)
	a, err := archive.Open(ctx, data)
	if err != nil || len(a.Sets) != 1 || a.Sets[0].Generation != 1 {
		t.Fatal(err)
	}
	Repack(t, data, docs[1:], nil)
	a, _ = archive.Open(ctx, data)
	if a.Sets[0].Generation != 2 {
		t.Fatal(a.Sets[0].Generation)
	}
}
