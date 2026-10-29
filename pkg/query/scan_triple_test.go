//ff:func feature=query type=scanner control=sequence
//ff:what TestScanTriple — hits per document over all parts; a cancelled context and a failing part stop
package query

import (
	"context"
	"errors"
	"testing"

	"github.com/park-jun-woo/silk/pkg/scan"
)

func TestScanTriple(t *testing.T) {
	ix := fixtureIndex(t)
	var st Stats
	got, err := scanTriple(context.Background(), ix, mustTriple(t, "P31=Q5"), scan.Options{}, &st)
	if err != nil || len(got) != 2 || st.Scanned == 0 {
		t.Fatal(got, err)
	}
	ctx, stop := context.WithCancel(context.Background())
	stop()
	if _, err := scanTriple(ctx, ix, mustTriple(t, "P31=Q5"), scan.Options{}, &st); !errors.Is(err, context.Canceled) {
		t.Fatal(err)
	}
	ix.Parts[0].Data.Close()
	if _, err := scanTriple(context.Background(), ix, mustTriple(t, "P31=Q5"), scan.Options{}, &st); err == nil {
		t.Fatal("closed data file must fail")
	}
}
