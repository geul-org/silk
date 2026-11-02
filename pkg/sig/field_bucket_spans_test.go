//ff:func feature=sig type=scanner control=sequence
//ff:what TestFieldBucketSpans — by name, by year, outside every bucket, not a number
package sig

import (
	"errors"
	"testing"
)

func TestFieldBucketSpans(t *testing.T) {
	d := testDef(t)
	f, _ := d.field("born")
	if s, _ := f.bucketSpans("mid"); s[0].Lo != 5 {
		t.Fatal(s)
	}
	if s, _ := f.bucketSpans("1950"); s[0].Lo != 6 {
		t.Fatal(s)
	}
	if _, err := f.bucketSpans("x"); !errors.Is(err, ErrCond) {
		t.Fatal(err)
	}
}
