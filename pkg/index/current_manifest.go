//ff:func feature=index type=reader control=sequence
//ff:what CurrentManifest — the manifest CURRENT names; ErrNoIndex when the directory has no index
package index

// CurrentManifest reads the live generation of dir.
func CurrentManifest(dir string) (Manifest, error) {
	g, err := readCurrent(dir)
	if err != nil {
		return Manifest{}, err
	}
	return readManifest(dir, g)
}
