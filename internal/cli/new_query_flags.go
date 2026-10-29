//ff:func feature=cli type=helper control=sequence
//ff:what newQueryFlags — the flag set of `silk query`, writing errors to stderr
package cli

import (
	"flag"
	"io"

	"github.com/park-jun-woo/silk/pkg/index"
)

// newQueryFlags binds every query flag to f.
func newQueryFlags(stderr io.Writer, f *queryFlags) *flag.FlagSet {
	fs := flag.NewFlagSet("silk query", flag.ContinueOnError)
	fs.SetOutput(stderr)
	fs.StringVar(&f.dir, "index", "", "index directory (default ARCHIVE/.silk)")
	fs.IntVar(&f.threads, "threads", index.DefaultThreads, "checking goroutines")
	fs.BoolVar(&f.count, "count", false, "print only the counts")
	fs.BoolVar(&f.explain, "explain", false, "print the plan instead of running it")
	fs.Var(&f.triples, "triple", "condition PRED=OBJ (repeatable)")
	fs.Var(&f.matches, "match", "condition KIND:PATH=VALUE[,PATH=VALUE] (repeatable)")
	fs.Var(&f.labels, "label", "documents with a LABEL in language LANG (repeatable)")
	return fs
}
