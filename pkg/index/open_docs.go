//ff:func feature=index type=reader control=iteration dimension=1
//ff:what openDocs — map a manifest's document table files
package index

import (
	"path/filepath"

	"github.com/geul-org/silk/pkg/seg"
)

func openDocs(dir string, m *Manifest) (*DocTable, error) {
	t := &DocTable{}
	for _, d := range m.Docs {
		b, err := seg.Map(filepath.Join(dir, d.File))
		if err != nil {
			t.Close()
			return nil, err
		}
		t.files = append(t.files, docMap{first: d.First, count: d.Count, raw: b})
	}
	return t, nil
}
