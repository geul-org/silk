//ff:func feature=index type=reader control=sequence
//ff:what Index.Section — a built bundle of a declaration; nil when it is not built
package index

func (ix *Index) Section(name string, bundle int) *Sec {
	ss := ix.secs[name]
	if bundle >= len(ss) {
		return nil
	}
	return ss[bundle]
}
