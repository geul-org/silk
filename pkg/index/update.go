//ff:func feature=index type=builder control=sequence
//ff:what Update — bring an index to the archive's current generations: read only new or moved versions into appended segments, mark removed versions dead, rebuild tags whose spec changed
package index

import (
	"context"
	"path/filepath"
	"time"

	"github.com/park-jun-woo/silk/pkg/archive"
	"github.com/park-jun-woo/silk/pkg/spec"
)

// Update is incremental: unchanged versions are not read (except for tags
// whose spec changed). With o.Compact the due merges run afterwards.
func Update(ctx context.Context, a *archive.Archive, dir string, sp spec.Spec, o Options) (Report, error) {
	t0 := time.Now()
	if err := sp.Validate(); err != nil {
		return Report{}, err
	}
	unlock, err := lockDir(dir)
	if err != nil {
		return Report{}, err
	}
	defer unlock()
	m, err := CurrentManifest(dir)
	if err != nil {
		return Report{}, err
	}
	old, err := loadDocs(dir, m)
	if err != nil {
		return Report{}, err
	}
	live, err := readBitmap(filepath.Join(dir, m.Live))
	if err != nil {
		return Report{}, err
	}
	docs, err := liveDocs(ctx, a)
	if err != nil {
		return Report{}, err
	}
	d, rebuild := diffDocs(m, old, live, docs), rebuildMask(m.Sections, sp)
	if len(d.added)+len(d.killed) == 0 && !anyTrue(rebuild) && len(m.Sections) == len(sp.Tags) && checkSets(a, m.Sets) == nil {
		return Report{Generation: m.Generation, Unchanged: true, Docs: live.Count(), Kept: len(d.kept), Elapsed: time.Since(t0)}, nil
	}
	rep, err := applyDiff(dir, &m, sp, a, old, live, d, rebuild, o)
	if err != nil {
		return rep, err
	}
	return finishWriter(ctx, dir, &m, o, rep, t0)
}
