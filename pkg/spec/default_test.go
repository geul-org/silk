//ff:func feature=spec type=builder control=sequence
//ff:what TestDefault — the default spec is valid, covers the seven Wikidata Edge kinds, partitions TRIPLE by Predicate and LABEL by language
package spec

import (
	"slices"
	"testing"
)

func TestDefault(t *testing.T) {
	s := Default()
	if err := s.Validate(); err != nil || len(s.Tags) != 7 {
		t.Fatal(err, len(s.Tags))
	}
	tr, _ := s.Find(0x100A)
	lb, _ := s.Find(0x1004)
	if !slices.Equal(tr.Partition.Path, []int{1}) || !slices.Equal(tr.Sort.Path, []int{3}) || !slices.Equal(lb.Partition.Path, []int{2, 0}) {
		t.Fatal(tr, lb)
	}
}
