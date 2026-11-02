//ff:func feature=cli type=builder control=sequence
//ff:what newFlags — a flag set that reports to stderr
package cli

import (
	"flag"
	"io"
	"runtime"

	"github.com/geul-org/silk/pkg/index"
)

func newFlags(name string, stderr io.Writer) *flags {
	p := index.DefaultPolicy()
	f := &flags{set: flag.NewFlagSet(name, flag.ContinueOnError)}
	s := f.set
	s.SetOutput(stderr)
	s.StringVar(&f.index, "index", "", "index directory (default ARCHIVE/.silk)")
	s.StringVar(&f.grammar, "grammar", "", "pack set directory of Issuer 0x0000 (default the archive's)")
	s.StringVar(&f.part, "part", "", "partition field path")
	s.StringVar(&f.key, "key", "", "key field paths, comma separated")
	s.StringVar(&f.width, "width", "", "key widths in bytes, comma separated")
	s.StringVar(&f.def, "def", "", "silk definition file")
	s.StringVar(&f.mem, "mem", "", "memory limit (required for writers), e.g. 2G")
	s.StringVar(&f.cache, "doc-cache", "0", "decoded document cache size")
	s.StringVar(&f.name, "index-name", "", "use this section")
	s.IntVar(&f.threads, "threads", runtime.NumCPU(), "worker threads")
	s.IntVar(&f.runs, "runs-per-tier", p.RunsPerTier, "runs of one size tier before they merge")
	s.Uint64Var(&f.appendRows, "append-rows", p.AppendRows, "append array bound in rows")
	s.Float64Var(&f.stale, "stale-ratio", p.StaleRatio, "dead row share that forces a major merge")
	s.Float64Var(&f.idleCPU, "idle-cpu", 0, "wait until this CPU share is free")
	s.Float64Var(&f.idleMem, "idle-mem", 0, "wait until this memory share is free")
	s.BoolVar(&f.full, "full", false, "merge every section to one run")
	s.BoolVar(&f.docs, "docs", false, "print documents with Edge Ids (default)")
	s.BoolVar(&f.records, "records", false, "print records")
	s.BoolVar(&f.keys, "keys", false, "print rows without reading documents")
	s.BoolVar(&f.explain, "explain", false, "print the plan")
	s.BoolVar(&f.scan, "scan", false, "read every live document instead of a section")
	s.BoolVar(&f.batch, "batch", false, "one query per standard input line")
	return f
}
