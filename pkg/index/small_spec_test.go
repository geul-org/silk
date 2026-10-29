//ff:func feature=index type=helper control=iteration dimension=1
//ff:what smallSpec — test helper: the default spec with minimum partition 2, so the small test archive has table partitions and other partitions
package index

import "github.com/park-jun-woo/silk/pkg/spec"

// smallSpec suits four documents.
func smallSpec() spec.Spec {
	s := spec.Default()
	for i := range s.Tags {
		s.Tags[i].MinPartition = 2
	}
	return s
}
