//ff:func feature=seg type=matcher control=iteration dimension=1
//ff:what emitChunk — emit the rows of a chunk whose bytes at `at` equal prefix (every row when cmp is false)
package seg

import "bytes"

// emitChunk walks rows of width w.
func emitChunk(b []byte, w int, prefix []byte, at int, cmp bool, emit func(doc, off uint32) error) error {
	for ; len(b) >= w; b = b[w:] {
		if cmp && !bytes.Equal(b[at:at+len(prefix)], prefix) {
			continue
		}
		if err := emit(SeqOf(b[:w])); err != nil {
			return err
		}
	}
	return nil
}
