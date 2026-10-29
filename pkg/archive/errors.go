package archive

import "errors"

var (
	// ErrNotPacked means the directory holds no pack set (a plain directory
	// archive must be packed first: geul pack build).
	ErrNotPacked = errors.New("silk/archive: not a pack archive")
	// ErrCompressed means a version document is stored zstd-compressed. Record
	// maps point into raw data pack bytes, so the archive must be packed with
	// GEULCOMPRESS=none.
	ErrCompressed = errors.New("silk/archive: compressed records (pack with GEULCOMPRESS=none)")
	// ErrManifest means manifest bytes this package reads break the pack format.
	ErrManifest = errors.New("silk/archive: bad manifest")
)
