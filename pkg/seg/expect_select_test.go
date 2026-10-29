//ff:func feature=seg type=helper control=sequence
//ff:what expectSelect — test helper: Select's emitted "doc:off" pairs (sorted) must equal want
package seg

import (
	"slices"
	"testing"
)

// expectSelect fails t with the case label when the selection differs.
func expectSelect(t *testing.T, f *File, sel Sel, want []string, label string) {
	t.Helper()
	emit, got := collect()
	_, err := f.Select(sel, emit)
	slices.Sort(*got)
	if err != nil || !slices.Equal(*got, want) {
		t.Fatal(label, err, *got)
	}
}
