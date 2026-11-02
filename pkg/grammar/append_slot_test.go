//ff:func feature=grammar type=scanner control=sequence
//ff:what TestAppendSlot — records and raw fields pad; rest fields end in 00 whole, FF cut; missing is zero
package grammar

import (
	"bytes"
	"testing"
)

func TestAppendSlot(t *testing.T) {
	if s := AppendSlot([]byte{9}, Val{B: []byte{1, 2}}, true, 4); !bytes.Equal(s, []byte{9, 1, 2, 0, 0}) {
		t.Fatal(s)
	}
	if s := AppendSlot(nil, Val{B: []byte{1, 2, 3}}, true, 2); !bytes.Equal(s, []byte{1, 2}) {
		t.Fatal(s)
	}
	if s := AppendSlot(nil, Val{}, false, 2); !bytes.Equal(s, []byte{0, 0}) {
		t.Fatal(s)
	}
	if s := AppendSlot(nil, Val{B: []byte{1}, Rest: true}, true, 3); !bytes.Equal(s, []byte{1, 0, 0}) {
		t.Fatal(s)
	}
	if s := AppendSlot(nil, Val{B: []byte{1, 2, 3}, Rest: true}, true, 3); !bytes.Equal(s, []byte{1, 2, 0xFF}) {
		t.Fatal(s)
	}
	if SlotExact(Val{B: []byte{1, 2, 3}}, 2) || !SlotExact(Val{B: []byte{1}}, 2) {
		t.Fatal("record")
	}
	if SlotExact(Val{Rest: true}, 2) || SlotExact(Val{B: []byte{1, 2}, Rest: true}, 2) || !SlotExact(Val{B: []byte{1}, Rest: true}, 2) {
		t.Fatal("rest")
	}
}
