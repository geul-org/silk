//ff:func feature=cli type=helper control=sequence
//ff:what newIndexFlags — the flag set of an index subcommand, writing errors to stderr
package cli

import (
	"flag"
	"io"

	"github.com/park-jun-woo/silk/pkg/index"
)

// newIndexFlags binds every index flag to f.
func newIndexFlags(name string, stderr io.Writer, f *indexFlags) *flag.FlagSet {
	fs := flag.NewFlagSet("silk index "+name, flag.ContinueOnError)
	fs.SetOutput(stderr)
	fs.StringVar(&f.dir, "index", "", "index directory (default ARCHIVE/.silk)")
	fs.StringVar(&f.spec, "spec", "", "index specification JSON (build: default spec; update: the index's own)")
	fs.IntVar(&f.threads, "threads", index.DefaultThreads, "worker goroutines")
	fs.BoolVar(&f.compact, "compact", false, "update: run due merges afterwards")
	fs.BoolVar(&f.full, "full", false, "compact: merge each section into one sorted run")
	fs.BoolVar(&f.idle, "idle", false, "compact: wait for an idle machine and pause while it is busy")
	fs.BoolVar(&f.dry, "dry", false, "compact: print the due merges only")
	return fs
}
