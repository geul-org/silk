//ff:func feature=cli type=builder control=sequence
//ff:what flags.options — writer options; the memory limit must be given
package cli

import "github.com/geul-org/silk/pkg/index"

func (f *flags) options() (index.Options, error) {
	if f.mem == "" {
		return index.Options{}, index.ErrLimit
	}
	n, err := parseSize(f.mem)
	if err != nil {
		return index.Options{}, err
	}
	p := index.Policy{AppendRows: f.appendRows, RunsPerTier: f.runs, StaleRatio: f.stale}
	return index.Options{Limit: n, Threads: f.threads, Policy: p}, nil
}
