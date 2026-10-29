//ff:func feature=index type=helper control=iteration dimension=1
//ff:what TestRefs — a manifest's own file, spec, document tables, bitmap, runs and appended segments are kept
package index

import "testing"

func TestRefs(t *testing.T) {
	m := Manifest{Generation: 2, Spec: "h", Live: "live/1", Docs: []DocFile{{File: "doc/1"}},
		Sections: []Section{{Runs: []SegRef{{File: "seg/1"}}, Appends: []SegRef{{File: "seg/2"}}}}}
	k := map[string]bool{}
	refs(m, k)
	for _, f := range []string{manifestName(2), "spec/h.json", "live/1", "doc/1", "seg/1", "seg/2"} {
		if !k[f] {
			t.Fatal(f)
		}
	}
}
