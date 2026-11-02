//ff:func feature=index type=reader control=sequence
//ff:what currentManifest — the current manifest, or an empty one for a new index
package index

import "errors"

func currentManifest(dir string) (Manifest, error) {
	g, err := readCurrent(dir)
	if errors.Is(err, ErrNoIndex) {
		return Manifest{Version: manifestVersion, Paths: Section{Name: "paths"}}, nil
	}
	if err != nil {
		return Manifest{}, err
	}
	return readManifest(dir, g)
}
