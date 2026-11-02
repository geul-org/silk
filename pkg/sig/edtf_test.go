//ff:func feature=sig type=helper control=sequence
//ff:what edtf — an EDTF record
package sig

import (
	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/silk/pkg/grammar"
)

func edtf(s string) []byte {
	return grammar.RecordOf(core.EncodeBytesValue(core.CodecEDTFEven, []byte(s)))
}
