//ff:func feature=index type=helper control=sequence
//ff:what firstSeg — test helper: the first run of the first section of the current generation
package index

// firstSeg names a segment to damage.
func firstSeg(dir string) string {
	m, _ := CurrentManifest(dir)
	return m.Sections[0].Runs[0].File
}
