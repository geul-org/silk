//ff:func feature=index type=helper control=sequence
//ff:what Manifest.DocCount — the document numbers given so far
package index

func (m *Manifest) DocCount() uint32 {
	if len(m.Docs) == 0 {
		return 0
	}
	l := m.Docs[len(m.Docs)-1]
	return l.First + l.Count
}
