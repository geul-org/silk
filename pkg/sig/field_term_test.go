//ff:func feature=sig type=scanner control=iteration dimension=1
//ff:what TestFieldTerm — values, reserved names, flags, values outside the definition
package sig

import "testing"

func TestFieldTerm(t *testing.T) {
	d := testDef(t)
	job, _ := d.field("job")
	born, _ := d.field("born")
	x, err := job.term("Q36180")
	if err != nil || len(x.Exact) != 1 || x.Exact[0].Lo != 5 || len(x.Spans) != 2 {
		t.Fatal(x, err)
	}
	x, _ = job.term("Q777")
	if len(x.Exact) != 0 || len(x.Spans) != 1 {
		t.Fatal(x)
	}
	x, _ = job.term("nope")
	if len(x.Exact) != 0 {
		t.Fatal(x)
	}
	for _, s := range []string{"none", "novalue", "somevalue"} {
		x, _ := job.term(s)
		if len(x.Spans) < 1 || len(x.Exact) != 1 {
			t.Fatal(s, x)
		}
	}
	x, _ = job.term("other")
	if len(x.Exact) != 0 || len(x.Truth) != 1 {
		t.Fatal(x)
	}
	if _, ok := reservedCode("x"); ok {
		t.Fatal("reserved")
	}
	x, _ = born.term("old")
	if x.Exact[0].Lo != 4 {
		t.Fatal(x)
	}
	if s, _ := born.bucketSpans("1850"); s[0].Lo != 5 {
		t.Fatal(s)
	}
	gap, _ := Parse("silk s 8\nfield a 0 4 TRIPLE P1 bucket\n  code x 4 1800..1900\n")
	if s, err := gap.Common[0].bucketSpans("2000"); s != nil || err != nil {
		t.Fatal(s, err)
	}
	fl := d.common("alive")
	if x, err := fl.term("1"); err != nil || len(x.Spans) != 1 {
		t.Fatal(x, err)
	}
	if _, err := fl.term("9"); err == nil {
		t.Fatal("flag")
	}
	if s := job.tableSpans("writer"); s[0].Lo != 5 {
		t.Fatal(s)
	}
}
