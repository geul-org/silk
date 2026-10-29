//ff:func feature=scan type=scanner control=iteration dimension=1
//ff:what TestRecords — many records over several goroutines give the same ordered hits as one; a bad offset fails
package scan

import (
	"bytes"
	"errors"
	"testing"
)

func TestRecords(t *testing.T) {
	var data []byte
	var offs []uint32
	_, p, o := tripleRec(0, 106, 1, 169470)
	for i := range 1000 {
		obj := uint64(5)
		if i%7 == 0 {
			obj = 169470
		}
		r, _, _ := tripleRec(uint64(i), 106, uint64(i), obj)
		offs = append(offs, uint32(len(data)))
		data = append(data, r...)
	}
	one, err := Records(bytes.NewReader(data), uint64(len(data)), offs, Triple{Pred: p, Obj: o}, Options{Threads: 1, Window: 64})
	many, err2 := Records(bytes.NewReader(data), uint64(len(data)), offs, Triple{Pred: p, Obj: o}, Options{Threads: 5})
	if err != nil || err2 != nil || len(one) != 143 || len(many) != len(one) || many[142].Off != one[142].Off {
		t.Fatal(len(one), len(many), err, err2)
	}
	if _, err := Records(bytes.NewReader(data), uint64(len(data)), []uint32{uint32(len(data))}, Triple{}, Options{}); !errors.Is(err, ErrRecord) {
		t.Fatal(err)
	}
}
