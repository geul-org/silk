package index

import "errors"

var (
	// ErrStale means the index does not match the archive's current
	// generations (another set list, generation or map list): rebuild it.
	ErrStale = errors.New("silk/index: index does not match the archive (run silk index build)")
	// ErrCatalog means index.json is unreadable or of another format version.
	ErrCatalog = errors.New("silk/index: bad index.json")
)
