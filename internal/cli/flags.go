//ff:type feature=cli type=model
//ff:what flags — every command's flags; each command reads the ones it takes
package cli

import "flag"

type flags struct {
	set                       *flag.FlagSet
	index, grammar, part, key string
	width, def, mem, cache    string
	name                      string
	threads, runs             int
	appendRows                uint64
	stale, idleCPU, idleMem   float64
	full, docs, records, keys bool
	explain, scan, batch      bool
}
