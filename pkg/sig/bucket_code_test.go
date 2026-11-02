//ff:func feature=sig type=scanner control=iteration dimension=1
//ff:what TestBucketCode — a number and one range; bad forms fail
package sig

import (
	"math"
	"testing"
)

func TestBucketCode(t *testing.T) {
	if e, err := bucketCode("n", "4", []string{"1800..1900"}, 4); err != nil || e.Lo != 1800 || e.Hi != 1900 {
		t.Fatal(e, err)
	}
	if e, err := bucketCode("n", "4", []string{"..1900"}, 4); err != nil || e.Lo != math.MinInt64 {
		t.Fatal(e, err)
	}
	for _, r := range [][]string{nil, {"a", "b"}, {"x..y"}, {"5..1"}, {"1900"}, {".."}} {
		if _, err := bucketCode("n", "4", r, 4); err == nil {
			t.Fatal(r)
		}
	}
	if _, err := bucketCode("n", "x", []string{"1..2"}, 4); err == nil {
		t.Fatal("number")
	}
}
