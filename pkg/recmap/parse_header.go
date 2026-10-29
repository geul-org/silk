//ff:func feature=recmap type=decoder control=sequence
//ff:what parseHeader — the 176-byte header back to a Header; wrong magic, version, Issuer length or nonzero padding is ErrFormat
package recmap

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
)

// parseHeader reads b[:headerSize].
func parseHeader(b []byte) (Header, error) {
	le := binary.LittleEndian
	if len(b) < headerSize || string(b[:8]) != magic || le.Uint16(b[8:]) != version {
		return Header{}, fmt.Errorf("%w: header", ErrFormat)
	}
	n := int(le.Uint16(b[10:]))
	if (n != 2 && n != 4 && n != 8 && n != 32) || !bytes.Equal(b[112+n:144], make([]byte, maxIssuer-n)) {
		return Header{}, fmt.Errorf("%w: Issuer", ErrFormat)
	}
	h := Header{Keys: le.Uint32(b[12:]), Records: le.Uint64(b[16:]), Docs: le.Uint64(b[24:]),
		Generation: le.Uint64(b[32:]), PackSize: le.Uint64(b[40:]),
		Pack: hex.EncodeToString(b[48:80]), Manifest: hex.EncodeToString(b[80:112]),
		Issuer: "0x" + hex.EncodeToString(b[112:112+n])}
	copy(h.Sum[:], b[144:176])
	return h, nil
}
