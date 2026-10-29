//ff:func feature=index type=builder control=iteration dimension=1
//ff:what addKept — re-index kept versions for the rebuilt tags only (their sequence numbers stay); nothing is read when no tag is rebuilt
package index

// addKept reads kept versions only when some tag's section is rebuilt.
func addKept(g *ingest, kept []keptDoc, rebuild []bool) error {
	if !anyTrue(rebuild) {
		return nil
	}
	for _, k := range kept {
		if _, err := g.doc(k.doc, k.seq, rebuild); err != nil {
			return err
		}
	}
	return nil
}
