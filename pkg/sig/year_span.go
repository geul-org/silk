//ff:func feature=sig type=decoder control=selection
//ff:what yearSpan — the years a value covers: GREGORIAN-DATE, EDTF, GEDCOM-DATE or an integer
package sig

import (
	"encoding/binary"
	"strings"
	"time"

	"github.com/geul-org/silk/pkg/grammar"
)

func yearSpan(o []byte) (int64, int64, bool) {
	if len(o) < 2 {
		return 0, 0, false
	}
	t := grammar.Target(o)
	switch t {
	case tagGregorian:
		if len(o) != 6 {
			return 0, 0, false
		}
		days := int64(int32(binary.BigEndian.Uint32(o[2:])))
		y := int64(time.Unix(days*86400, 0).UTC().Year())
		return y, y, true
	case tagEDTF, tagEDTFOdd:
		return edtfYears(strings.TrimRight(string(o[2:]), "\x00"))
	case tagGEDCOM, tagGEDCOMOdd:
		return gedcomYears(strings.TrimRight(string(o[2:]), "\x00"))
	case tagSigned, tagUnsigned:
		return intYears(o[2:], t == tagSigned)
	}
	return 0, 0, false
}
