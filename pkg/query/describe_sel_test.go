//ff:func feature=query type=helper control=sequence
//ff:what TestDescribeSel — one partition / every partition, sort key search / every row
package query

import "testing"

func TestDescribeSel(t *testing.T) {
	if s := describeSel(CondPlan{ByPart: true, BySort: true}); s != "one partition, sort key search" {
		t.Fatal(s)
	}
	if s := describeSel(CondPlan{}); s != "every partition, every row" {
		t.Fatal(s)
	}
}
