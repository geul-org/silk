//ff:func feature=cli type=command control=iteration dimension=1
//ff:what indexStat — per section the runs, appended segments, rows, dead rows and bytes of the current generation, and the merges the default policy finds due
package cli

import (
	"fmt"
	"io"

	"github.com/park-jun-woo/silk/pkg/index"
)

// indexStat reads the manifest only.
func indexStat(dir string, stdout, stderr io.Writer) int {
	m, err := index.CurrentManifest(dir)
	if err != nil {
		return fail(stderr, err)
	}
	fmt.Fprintf(stdout, "generation %d, spec %s, %d document files\n", m.Generation, m.Spec[:12], len(m.Docs))
	total := int64(0)
	for _, s := range m.Sections {
		b := segBytes(s)
		total += b
		fmt.Fprintf(stdout, "%-15s runs %d appended %d rows %d dead %d bytes %d\n", s.Name, len(s.Runs), len(s.Appends), segRows(s), s.Dead, b)
	}
	fmt.Fprintf(stdout, "segments %d bytes\n", total)
	for _, t := range index.Plan(m, index.DefaultPolicy(), false) {
		fmt.Fprintf(stdout, "due: %s %s (%d segments)\n", t.Name, t.Kind, len(t.Inputs))
	}
	return 0
}
