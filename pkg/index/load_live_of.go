//ff:func feature=index type=reader control=sequence
//ff:what loadLiveOf — a manifest's bitmap; empty before the first build
package index

import "path/filepath"

func loadLiveOf(dir string, m Manifest) (Bitmap, error) {
	if m.Live == "" {
		return Bitmap{}, nil
	}
	return readBitmap(filepath.Join(dir, m.Live))
}
