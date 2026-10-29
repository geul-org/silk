//ff:func feature=index type=helper control=sequence
//ff:what TestOptionsBatch — DefaultBatch when unset, else the given value
package index

import "testing"

func TestOptionsBatch(t *testing.T) {
	if (Options{}).batch() != DefaultBatch || (Options{Batch: 5}).batch() != 5 {
		t.Fatal("batch")
	}
}
