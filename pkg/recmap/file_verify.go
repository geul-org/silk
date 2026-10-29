//ff:func feature=recmap type=rule control=sequence
//ff:what File.Verify — body SHA-256 matches the header, and rebuilding from the pack and the given documents gives the same tables
package recmap

import (
	"crypto/sha256"
	"fmt"
	"io"
	"slices"

	"github.com/park-jun-woo/silk/pkg/archive"
)

// Verify checks f against the data pack (data, f.Header.PackSize) and the
// documents the archive now places in it (for f's Issuer).
func (f *File) Verify(data io.ReaderAt, docs []archive.Doc) error {
	h := sha256.New()
	if _, err := io.Copy(h, io.NewSectionReader(f.r, headerSize, int64(bodyEnd(f.Header)-headerSize))); err != nil {
		return fmt.Errorf("%w: %v", ErrFormat, err)
	}
	if [32]byte(h.Sum(nil)) != f.Header.Sum {
		return fmt.Errorf("%w: body SHA-256", ErrFormat)
	}
	want, err := Build(data, f.Header.PackSize, docs)
	if err != nil {
		return err
	}
	offs, err := f.offs(0, f.Header.Records)
	if err != nil {
		return err
	}
	if !slices.Equal(f.Keys, want.Keys) || !slices.Equal(f.Docs, want.Docs) || !slices.Equal(offs, want.Offs) {
		return fmt.Errorf("%w: %s", ErrMismatch, f.Header.Pack)
	}
	return nil
}
