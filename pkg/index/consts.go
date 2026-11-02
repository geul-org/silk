// Package index is a SILK index directory: generations of a manifest over
// declarations, a document table, a live bitmap and segments. Build and update
// read packs in pieces under a memory limit the caller sets; compaction
// merges size-tiered runs; a reader pins a generation and maps its files.
// See docs/index.md.
package index

import "errors"

const (
	manifestVersion = 1
	currentFile     = "CURRENT"
	lockFile        = "LOCK"
	dirManifest     = "manifest"
	dirDecl         = "decl"
	dirGrammar      = "grammar"
	dirSeg          = "seg"
	dirDoc          = "doc"
	dirLive         = "live"
	dirPin          = "pin"
	dirTmp          = "tmp"

	docMagic  = "SILKDOCS"
	liveMagic = "SILKLIVE"
	docSize   = 40
	tailSize  = 56
	pathKey   = 16
	encDead   = 0xFF

	minFanIn   = 2
	maxFanIn   = 256
	fanInBytes = 1 << 20
	staleProbe = 4096
)

var (
	// ErrStale means the index was built from other set generations than the archive now has.
	ErrStale = errors.New("silk/index: index does not match the archive (run silkg index update)")
	// ErrNoIndex means the directory holds no index yet.
	ErrNoIndex = errors.New("silk/index: no index (run silkg create, then silkg index build)")
	// ErrFormat means an index file breaks its format.
	ErrFormat = errors.New("silk/index: bad index file")
	// ErrLocked means another process is writing the index.
	ErrLocked = errors.New("silk/index: index is locked by another writer")
	// ErrLimit means the operation needs a memory limit from the caller.
	ErrLimit = errors.New("silk/index: a memory limit is required")
	// ErrMismatch means verification found the index differs from the archive.
	ErrMismatch = errors.New("silk/index: index does not match its documents")
	// ErrPending means a section was declared but not built yet.
	ErrPending = errors.New("silk/index: section not built (run silkg index update)")
	// ErrExists means the name is already declared.
	ErrExists = errors.New("silk/index: name already declared")
	// ErrUnknown means no such section is declared.
	ErrUnknown = errors.New("silk/index: no such section")
)
