//ff:func feature=index type=helper control=sequence
//ff:what Manifest.nextName — the path of a new file under sub
package index

import "fmt"

func (m *Manifest) nextName(sub, ext string) string {
	m.NextFile++
	return fmt.Sprintf("%s/%016x.%s", sub, m.NextFile, ext)
}
