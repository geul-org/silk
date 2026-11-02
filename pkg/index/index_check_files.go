//ff:func feature=index type=rule control=iteration dimension=1
//ff:what Index.checkFiles — hashes and row order of the segments and tables
package index

func (ix *Index) checkFiles() error {
	if err := ix.Docs.check(ix.Dir, &ix.Manifest); err != nil {
		return err
	}
	for _, ss := range ix.secs {
		if err := checkSecs(ss); err != nil {
			return err
		}
	}
	return nil
}
