// Package index keeps SILK's index directory: for each TAG kind of the
// index specification a section of segment files (sorted runs and appended
// segments, pkg/seg), a document table, a live-version bitmap and a
// manifest per generation. A generation is switched atomically by rewriting
// CURRENT; files are never changed, and files no current or pinned
// generation names are removed. Build indexes an archive from scratch,
// Update adds the documents of new pack generations as appended segments,
// Compact merges segments by size tier (docs/index.md).
package index

const (
	manifestVersion = 2
	currentFile     = "CURRENT"
	lockFile        = "LOCK"
	checkpointFile  = "compact.json"
	dirManifest     = "manifest"
	dirSeg          = "seg"
	dirDoc          = "doc"
	dirLive         = "live"
	dirSpec         = "spec"
	dirPin          = "pin"
	docsMagic       = "SILKDOCS"
	liveMagic       = "SILKLIVE"
	fileVersion     = 1
	docsHead        = 56
	docSize         = 40
	pairSize        = 6
	liveHead        = 56
	// DefaultThreads is the default number of worker goroutines: enough for
	// memory bandwidth while leaving a desktop usable (about 30% free).
	DefaultThreads = 8
	// DefaultBatch is the rows of one TAG kind kept in memory before a build
	// writes them as a sorted run (merged afterwards).
	DefaultBatch = 4 << 20
	// TaskFlush, TaskTier and TaskMajor are the compaction task kinds.
	TaskFlush = "flush"
	TaskTier  = "tier"
	TaskMajor = "major"
)
