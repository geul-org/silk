//ff:func feature=sig type=scanner control=iteration dimension=1
//ff:what TestFieldAdd — fit, reserved, clash and value checks
package sig

import "testing"

func TestFieldAdd(t *testing.T) {
	f := &Field{Name: "f", Width: 4, byValue: map[string]*Entry{}}
	v := []byte{1}
	if err := f.add(&Entry{Name: "a", Code: 4, Len: 4, Values: [][]byte{v}}); err != nil || f.byValue["\x01"] == nil {
		t.Fatal(err)
	}
	for _, e := range []*Entry{{Name: "x", Code: 16, Len: 4}, {Name: "x", Code: 1, Len: 8}, {Name: "x", Code: 3, Len: 4}, {Name: "a", Code: 5, Len: 4}, {Name: "x", Code: 4, Len: 4}, {Name: "x", Code: 6, Len: 4, Values: [][]byte{v}}} {
		if err := f.add(e); err == nil {
			t.Fatal(e)
		}
	}
	w := &Field{Width: 64, byValue: map[string]*Entry{}}
	if err := w.add(&Entry{Name: "b", Code: 1 << 60, Len: 64}); err != nil {
		t.Fatal(err)
	}
}
