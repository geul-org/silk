//ff:func feature=index type=builder control=sequence
//ff:what startManifest — an empty manifest continuing the generation and file numbers of the index already in dir (if any)
package index

// startManifest keeps numbers unique across rebuilds.
func startManifest(dir string) Manifest {
	cur, err := CurrentManifest(dir)
	if err != nil {
		return Manifest{Version: manifestVersion}
	}
	return Manifest{Version: manifestVersion, Generation: cur.Generation, NextFile: cur.NextFile}
}
