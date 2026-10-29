//ff:func feature=index type=scanner control=iteration dimension=1
//ff:what Index.Select — a selection over every segment of a section (binary search in runs, scan of appended segments), rows of dead versions skipped
package index

import "github.com/geul-org/silk/pkg/seg"

// Select emits live (document, offset) pairs; the same record never appears
// twice because a document version is indexed once.
func (ix *Index) Select(sec int, sel seg.Sel, emit func(doc, off uint32) error) (SelectStats, error) {
	var st SelectStats
	keep := func(doc, off uint32) error {
		if !ix.Live.Get(doc) {
			st.Dead++
			return nil
		}
		st.Hits++
		return emit(doc, off)
	}
	for _, f := range ix.Segs[sec] {
		n, err := f.Select(sel, keep)
		st.Segments++
		st.Rows += n
		if err != nil {
			return st, err
		}
	}
	return st, nil
}
