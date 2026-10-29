//ff:func feature=seg type=writer control=sequence
//ff:what ResumeWriter — reopen a checkpointed segment: cut the temporary file back to the saved length and restore table, other rows and hash
package seg

import (
	"bufio"
	"crypto/sha256"
	"encoding"
	"errors"
	"fmt"
	"os"
)

// ResumeWriter continues st; bytes written after the checkpoint are dropped.
func ResumeWriter(st State) (*Writer, error) {
	f, err := os.OpenFile(st.Path+".tmp", os.O_WRONLY|os.O_APPEND, 0o644)
	if err != nil {
		return nil, err
	}
	fi, err := f.Stat()
	if err != nil || fi.Size() < st.Written {
		f.Close()
		return nil, fmt.Errorf("%w: %s.tmp is shorter than its checkpoint", ErrFormat, st.Path)
	}
	h := sha256.New()
	if err := errors.Join(h.(encoding.BinaryUnmarshaler).UnmarshalBinary(st.Hash), f.Truncate(st.Written)); err != nil {
		f.Close()
		return nil, err
	}
	return &Writer{Path: st.Path, Min: st.Min, info: st.Info, f: f, bw: bufio.NewWriterSize(f, 1<<20), h: h,
		parts: st.Parts, other: st.Other, key: st.Key, written: st.Written}, nil
}
