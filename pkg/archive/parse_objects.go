//ff:func feature=archive type=decoder control=sequence
//ff:what parseObjects — the OBJECTS section of manifest bytes (pack format §4.3); the manifest's other sections are pack's business
//ff:why gogeul pkg/pack keeps object names unexported; this small reader is a candidate to move there (Reader.Objects).
package archive

import (
	"encoding/binary"
	"fmt"
)

// parseObjects checks the manifest magic and reads OBJECTS. Reading the rest
// and verifying the manifest hash is pack.Open's job.
func parseObjects(b []byte) ([]Object, error) {
	if len(b) < manifestHead || string(b[:8]) != manifestMagic {
		return nil, fmt.Errorf("%w: no manifest header", ErrManifest)
	}
	sec, err := sectionOf(b, int(binary.LittleEndian.Uint32(b[12:])), secObjects)
	if err != nil {
		return nil, err
	}
	return readObjects(sec)
}
