package index

import "errors"

var (
	// ErrStale means the index was built from other set generations than the
	// archive now has: run silkg index update.
	ErrStale = errors.New("silk/index: index does not match the archive (run silkg index update)")
	// ErrNoIndex means the directory holds no index yet: run silkg index build.
	ErrNoIndex = errors.New("silk/index: no index (run silkg index build)")
	// ErrFormat means an index file breaks its format.
	ErrFormat = errors.New("silk/index: bad index file")
	// ErrLocked means another process is writing the index.
	ErrLocked = errors.New("silk/index: index is locked by another writer")
	// ErrPaused means compaction stopped at a checkpoint; running it again resumes.
	ErrPaused = errors.New("silk/index: compaction paused (run silkg index compact to resume)")
	// ErrMismatch means verification found the index differs from the archive.
	ErrMismatch = errors.New("silk/index: index does not match its documents")
)
