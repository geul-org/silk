//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what printReport — a writer run's report: generation, documents, bytes read and written, write amplification, merges, time
package cli

import (
	"fmt"
	"io"
	"time"

	"github.com/park-jun-woo/silk/pkg/index"
)

// printReport writes one summary line and one line per merge.
func printReport(w io.Writer, what string, r index.Report) {
	wa := "-"
	if r.NewBytes > 0 {
		wa = fmt.Sprintf("%.3f", float64(r.Written)/float64(r.NewBytes))
	}
	fmt.Fprintf(w, "silk: %s: generation %d; documents %d (new %d, kept %d, removed %d); read %d B (new %d B); written %d B; write amplification %s; %v\n",
		what, r.Generation, r.Docs, r.New, r.Kept, r.Killed, r.Read, r.NewBytes, r.Written, wa, r.Elapsed.Round(time.Millisecond))
	for _, t := range r.Tasks {
		fmt.Fprintf(w, "silk:   merge %s\n", t)
	}
	if r.Unchanged {
		fmt.Fprintln(w, "silk: index already matches the archive")
	}
	if r.Paused {
		fmt.Fprintln(w, "silk: compaction paused at a checkpoint; run silk index compact to resume")
	}
}
