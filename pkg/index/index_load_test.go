//ff:func feature=index type=reader control=iteration dimension=1
//ff:what TestIndexLoad — every file a generation needs is checked: manifest, archive sets, spec, document table, bitmap, pack sets, segments
package index

import (
	"os"
	"path/filepath"
	"testing"
)

func TestIndexLoad(t *testing.T) {
	if err := openBroken(t, func(string, Manifest) {}); err != nil {
		t.Fatal(err)
	}
	breaks := map[string]func(dir string, m Manifest){
		"manifest": func(dir string, m Manifest) { os.Remove(filepath.Join(dir, manifestName(m.Generation))) },
		"sets":     func(dir string, m Manifest) { m.Sets[0].Generation = 99; rewriteManifest(dir, m) },
		"spec":     func(dir string, m Manifest) { os.Remove(filepath.Join(dir, specName(m.Spec))) },
		"docs":     func(dir string, m Manifest) { os.Remove(filepath.Join(dir, m.Docs[0].File)) },
		"live":     func(dir string, m Manifest) { os.Remove(filepath.Join(dir, m.Live)) },
		"packs":    func(dir string, m Manifest) { m.Packs[0].Set = "0x7777"; rewriteManifest(dir, m) },
		"segments": func(dir string, m Manifest) { os.Remove(filepath.Join(dir, m.Sections[3].Runs[0].File)) },
	}
	for name, fn := range breaks {
		if err := openBroken(t, fn); err == nil {
			t.Fatal(name)
		}
	}
}
