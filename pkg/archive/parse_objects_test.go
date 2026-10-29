//ff:func feature=archive type=decoder control=sequence
//ff:what TestParseObjects — a real manifest's OBJECTS (index then data pack); short, wrong magic and missing sections fail
package archive

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/geul-org/silk/internal/fixture"
)

func TestParseObjects(t *testing.T) {
	a, _ := Open(context.Background(), fixture.Packed(t, nil))
	s := a.Sets[0]
	b, _ := os.ReadFile(filepath.Join(s.Dir, "m", s.Manifest+".gpm"))
	objs, err := parseObjects(b)
	if err != nil || len(objs) == 0 || len(objs[0].Hash) != 64 {
		t.Fatal(objs, err)
	}
	if _, err := parseObjects(b[:10]); !errors.Is(err, ErrManifest) {
		t.Fatal(err)
	}
	bad := append([]byte("XEULMANI"), b[8:]...)
	if _, err := parseObjects(bad); !errors.Is(err, ErrManifest) {
		t.Fatal(err)
	}
	head := append([]byte(nil), b[:64]...)
	if _, err := parseObjects(head); !errors.Is(err, ErrManifest) {
		t.Fatal(err)
	}
}
