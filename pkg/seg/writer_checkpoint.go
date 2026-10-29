//ff:func feature=seg type=writer control=sequence
//ff:what Writer.Checkpoint — flush and sync the temporary file and capture the Writer's state; only between partitions
package seg

import (
	"encoding"
	"errors"
	"fmt"
)

// Checkpoint returns a State to resume from; the Writer stays usable.
func (w *Writer) Checkpoint() (State, error) {
	if w.open {
		return State{}, fmt.Errorf("%w: checkpoint inside a partition", ErrOrder)
	}
	if err := errors.Join(w.bw.Flush(), w.f.Sync()); err != nil {
		return State{}, err
	}
	hs, _ := w.h.(encoding.BinaryMarshaler).MarshalBinary()
	return State{Path: w.Path, Min: w.Min, Info: w.info, Parts: append([]Part{}, w.parts...), Other: append([]byte{}, w.other...),
		Key: w.key, Hash: hs, Written: w.written}, nil
}
