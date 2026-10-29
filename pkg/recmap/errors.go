package recmap

import "errors"

var (
	// ErrFormat means map file bytes break the format (docs/recmap.md).
	ErrFormat = errors.New("silk/recmap: bad map file")
	// ErrMismatch means a map does not describe the pack bytes it names.
	ErrMismatch = errors.New("silk/recmap: map does not match its pack")
	// ErrDocument means a document in a pack is not one well-formed GEUL document.
	ErrDocument = errors.New("silk/recmap: bad document")
)
