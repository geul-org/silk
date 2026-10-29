//ff:func feature=sig type=helper control=sequence
//ff:what TestSchemaField — a field by name; an unknown name is absent
package sig

import "testing"

func TestSchemaField(t *testing.T) {
	if f, ok := sampleSchema().field("P106"); !ok || f.Kind != OneHot {
		t.Fatal(f, ok)
	}
	if _, ok := sampleSchema().field("x"); ok {
		t.Fatal("x")
	}
}
