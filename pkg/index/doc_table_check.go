//ff:func feature=index type=rule control=iteration dimension=1
//ff:what DocTable.check — each file's tail hash
package index

import "fmt"

func (t *DocTable) check(dir string, m *Manifest) error {
	for i, f := range t.files {
		first, count, _, err := readTail(m.Docs[i].File, f.raw, docMagic)
		if err == nil && (first != f.first || count != f.count) {
			err = fmt.Errorf("%w: %s counts", ErrFormat, m.Docs[i].File)
		}
		if err != nil {
			return err
		}
	}
	return nil
}
