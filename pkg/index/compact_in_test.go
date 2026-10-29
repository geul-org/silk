//ff:func feature=index type=writer control=iteration dimension=1
//ff:what TestCompactIn — a paused round saves a checkpoint, the next call resumes it and finishes; a missing spec, a missing bitmap and a failing round fail
package index

import (
	"context"
	"errors"
	"io"
	"os"
	"path/filepath"
	"sync/atomic"
	"testing"
)

func TestCompactIn(t *testing.T) {
	_, _, dir := updated(t)
	ctx := context.Background()
	m, _ := CurrentManifest(dir)
	var calls atomic.Int32
	o := Options{Full: true, Pause: func(context.Context) error {
		if calls.Add(1) == 1 {
			return nil
		}
		return io.EOF
	}}
	r, err := compact(ctx, dir, &m, o)
	if !errors.Is(err, ErrPaused) || !r.Paused || len(readCheckpoints(dir)) != 7 {
		t.Fatal(err, r)
	}
	if _, err := compact(ctx, dir, &m, Options{Pause: func(context.Context) error { return io.EOF }}); !errors.Is(err, ErrPaused) {
		t.Fatal("pause before a round", err)
	}
	r, err = compact(ctx, dir, &m, Options{})
	if err != nil || len(r.Tasks) != 7 || len(readCheckpoints(dir)) != 0 {
		t.Fatal(err, r)
	}
	for name, f := range map[string]string{"spec": SpecFile(dir, m), "live": filepath.Join(dir, m.Live)} {
		b, _ := os.ReadFile(f)
		os.Remove(f)
		if _, err := compact(ctx, dir, &m, Options{}); err == nil {
			t.Fatal(name)
		}
		os.WriteFile(f, b, 0o644)
	}
	_, _, dir2 := updated(t)
	m2, _ := CurrentManifest(dir2)
	os.Remove(filepath.Join(dir2, m2.Sections[0].Appends[0].File))
	if _, err := compact(ctx, dir2, &m2, Options{Full: true}); err == nil {
		t.Fatal("failing round")
	}
}
