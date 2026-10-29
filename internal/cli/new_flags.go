//ff:func feature=cli type=helper control=sequence
//ff:what newFlags — a flag set writing errors to stderr, with the -index flag every command takes
package cli

import (
	"flag"
	"io"
)

// newFlags returns the set and the -index value.
func newFlags(name string, stderr io.Writer) (*flag.FlagSet, *string) {
	fs := flag.NewFlagSet("silk "+name, flag.ContinueOnError)
	fs.SetOutput(stderr)
	dir := fs.String("index", "", "index directory (default ARCHIVE/.silk)")
	return fs, dir
}
