// Package scan runs SILK's fused scan loop: given the pack offsets of one
// key's records (a record map range), it reads the data pack in windows with
// ReadAt, frames each record from its TAG WORD and tests it with a Matcher in
// the same pass, over several goroutines. No intermediate arrays are built.
package scan

const (
	// DefaultThreads is the default number of scanning goroutines: enough to
	// reach memory bandwidth on a desktop CPU while leaving the machine usable.
	DefaultThreads = 8
	// DefaultWindow is the default read size of one window (4 MiB).
	DefaultWindow = 4 << 20
	// maxHead is the longest TAG WORD plus Length extension (3 words).
	maxHead = 6
)
