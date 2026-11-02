//ff:func feature=grammar type=scanner control=sequence
//ff:what TestDocLoadRefs — only REFERENCE-NODEs with both fields bind
package grammar

import (
	"bytes"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestDocLoadRefs(t *testing.T) {
	g := testGrammar(t)
	text := RecordOf(core.EncodeBytesValue(core.CodecUTF8TextEven, []byte("ab")))
	d, _ := Frame(bytes.Join([][]byte{edgeOf(0x00, tok("#1"), tok("#5"), text), edgeOf(0x00, tok("#2"), tok("#6")), testLabel()}, nil))
	d.loadRefs(g)
	if len(d.refs) != 1 || !bytes.Equal(d.refs[string(tok("#5"))], text) {
		t.Fatal(d.refs)
	}
}
