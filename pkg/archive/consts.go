// Package archive opens a GEUL pack archive for indexing and querying: its
// sets at their latest generation, their data packs, a step-by-step replay of
// each pack's inventory, and single documents by location. Nothing here holds
// a whole set's paths in memory; zstd and raw records take the same path.
package archive

import "errors"

const (
	manifestMagic = "GEULMANI"
	manifestHead  = 64 // fixed manifest header (GEUL pack format §4.3)
	sectionSize   = 24 // section table entry
	objectSize    = 48 // OBJECTS entry
	secObjects    = 0x0001
)

var (
	// ErrNotPacked means the directory holds no pack set.
	ErrNotPacked = errors.New("silk/archive: not a pack archive")
	// ErrManifest means a pack manifest breaks its format.
	ErrManifest = errors.New("silk/archive: bad pack manifest")
)
