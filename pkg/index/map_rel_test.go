//ff:func feature=index type=helper control=sequence
//ff:what TestMapRel — `<Issuer>/<pack>.srm`
package index

import (
	"testing"

	"github.com/park-jun-woo/silk/pkg/archive"
)

func TestMapRel(t *testing.T) {
	if got := mapRel(Target{Issuer: "0x0001", Pack: archive.Pack{Hash: "ab"}}); got != "0x0001/ab.srm" {
		t.Fatal(got)
	}
}
