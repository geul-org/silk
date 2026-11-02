//ff:func feature=index type=builder control=sequence
//ff:what updater.ingest — read the documents the sections need and turn them into rows
package index

import "context"

func (u *updater) ingest(ctx context.Context, r *Report) error {
	ts, err := u.targets()
	if err != nil || len(ts) == 0 {
		return err
	}
	docs, err := openDocs(u.dir, &u.m)
	if err != nil {
		return err
	}
	defer docs.Close()
	g := &ingester{u: u, ts: ts, budget: u.o.Limit / 2}
	if err := g.run(ctx, docs); err != nil {
		return err
	}
	r.ReadDocs = g.read
	u.m.MaxDoc = max(u.m.MaxDoc, g.maxDoc)
	return u.finishTargets(ts, r)
}
