//ff:func feature=index type=builder control=sequence
//ff:what Build — index every live document of an archive from scratch: sorted runs per tag (merged into one), document table, live bitmap, a new generation
package index

import (
	"context"
	"time"

	"github.com/geul-org/silk/pkg/archive"
	"github.com/geul-org/silk/pkg/spec"
)

// Build replaces whatever dir held (old files are collected once no reader
// pins them). Rows are written in batches of o.Batch per tag and merged.
func Build(ctx context.Context, a *archive.Archive, dir string, sp spec.Spec, o Options) (Report, error) {
	t0 := time.Now()
	if err := sp.Validate(); err != nil {
		return Report{}, err
	}
	unlock, err := lockDir(dir)
	if err != nil {
		return Report{}, err
	}
	defer unlock()
	docs, err := liveDocs(ctx, a)
	if err != nil {
		return Report{}, err
	}
	m := startManifest(dir)
	g := newIngest(dir, &m, sp, fill(len(sp.Tags), true), o.batch())
	entries, err := addDocs(g, docs, 0)
	if err != nil {
		return Report{}, err
	}
	if err := g.flush(); err != nil {
		return Report{}, err
	}
	m.Sections = nextSections(nil, sp, fill(len(sp.Tags), true), g.out)
	live := allLive(uint32(len(entries)))
	rep := Report{Docs: len(entries), New: len(entries), NewBytes: g.read, Read: g.read, Written: g.written}
	if err := writeGeneration(dir, &m, sp, a, entries, live, &rep); err != nil {
		return rep, err
	}
	return finishWriter(ctx, dir, &m, Options{Threads: o.Threads, Full: true, Pause: o.Pause}, rep, t0)
}
