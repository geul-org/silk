//ff:func feature=index type=builder control=sequence
//ff:what updater.advance — replay inventories, make rows, commit
package index

import "context"

func (u *updater) advance(ctx context.Context) (Report, error) {
	var r Report
	if err := u.loadLive(); err != nil {
		return r, err
	}
	u.first = u.m.DocCount()
	if err := u.syncDocs(ctx); err != nil {
		return r, err
	}
	r.NewDocs = u.m.DocCount() - u.first
	if err := u.ingest(ctx, &r); err != nil {
		return r, err
	}
	return r, u.save()
}
