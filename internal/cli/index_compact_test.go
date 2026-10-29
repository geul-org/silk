//ff:func feature=cli type=command control=sequence
//ff:what TestIndexCompact — -dry lists, a cancelled context pauses (exit 3), -idle and -full compact, a missing index fails
package cli

import (
	"context"
	"io"
	"path/filepath"
	"strings"
	"testing"

	"github.com/geul-org/silk/internal/fixture"
)

func TestIndexCompact(t *testing.T) {
	data := packedIndex(t)
	fixture.Evolve(t, data)
	call("index update " + data)
	dir := filepath.Join(data, ".silk")
	var o strings.Builder
	if code := indexCompact(context.Background(), dir, indexFlags{dry: true}, &o, io.Discard); code != 0 || !strings.Contains(o.String(), "due:") {
		t.Fatal(code, o.String())
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if code := indexCompact(ctx, dir, indexFlags{full: true}, io.Discard, io.Discard); code != 3 {
		t.Fatal("cancelled", code)
	}
	if code := indexCompact(context.Background(), dir, indexFlags{full: true, idle: true, threads: 1}, io.Discard, io.Discard); code != 0 {
		t.Fatal("idle full", code)
	}
	if code := indexCompact(context.Background(), t.TempDir(), indexFlags{}, io.Discard, io.Discard); code != 1 {
		t.Fatal("no index", code)
	}
}
