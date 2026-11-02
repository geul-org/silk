//ff:func feature=cli type=scanner control=sequence
//ff:what TestCommands — the whole cycle: declare, build, stat, list, every query form, update, compact, verify, drop
package cli

import (
	"bytes"
	"context"
	"strings"
	"testing"
)

func TestCommands(t *testing.T) {
	data, ix := project(t)
	cmd := func(want int, a ...string) string {
		t.Helper()
		c, o, e := silkg("", append(a[:len(a):len(a)], append(ix, data)...)...)
		if c != want {
			t.Fatal(a, c, e)
		}
		return o + e
	}
	if o := cmd(0, "list"); !strings.Contains(o, "pending") {
		t.Fatal(o)
	}
	if o := cmd(1, "index", "build"); !strings.Contains(o, "memory limit") {
		t.Fatal(o)
	}
	if o := cmd(0, "index", "build", "-mem", "64M", "-threads", "2"); !strings.Contains(o, "new documents 1") {
		t.Fatal(o)
	}
	if o := cmd(0, "index", "update", "-mem", "64M"); !strings.Contains(o, "new documents 0") {
		t.Fatal(o)
	}
	if o := cmd(0, "index", "stat"); !strings.Contains(o, "live 1") || !strings.Contains(o, "built") {
		t.Fatal(o)
	}
	if o := cmd(0, "query", "-explain", "TRIPLE", "predicate=P31", "object=Q5"); !strings.Contains(o, "#1") || !strings.Contains(o, "1 hits") {
		t.Fatal(o)
	}
	if o := cmd(0, "query", "-records", "TRIPLE", "predicate=P31", "object=Q5"); !strings.Contains(o, "+") {
		t.Fatal(o)
	}
	if o := cmd(0, "query", "-keys", "LABEL", `object.bytes="Albert"`); !strings.Contains(o, "~") || !strings.Contains(o, "qid=40000001") {
		t.Fatal(o)
	}
	if o := cmd(0, "query", "-scan", "LABEL", `object.bytes="Albert"`); !strings.Contains(o, "#2") {
		t.Fatal(o)
	}
	if o := cmd(0, "query", "entity", "type=person"); !strings.Contains(o, ".geul") {
		t.Fatal(o)
	}
	cmd(1, "query", "-doc-cache", "x", "TRIPLE")
	cmd(1, "query", "NOPE")
	cmd(1, "query", "-keys", "TRIPLE", "subject=Q1")
	if o := cmd(0, "index", "compact", "-mem", "64M", "-full"); !strings.Contains(o, "merges") {
		t.Fatal(o)
	}
	if o := cmd(0, "index", "compact", "-mem", "64M", "-idle-cpu", "0.01", "-idle-mem", "0.01"); !strings.Contains(o, "merges") {
		t.Fatal(o)
	}
	cmd(1, "index", "compact", "-mem", "64M", "-idle-cpu", "2")
	cmd(1, "index", "compact")
	if o := cmd(0, "index", "verify", "-mem", "64M"); !strings.Contains(o, "ok") {
		t.Fatal(o)
	}
	cmd(1, "index", "verify")
	cmd(0, "drop", "sem")
	if o := cmd(0, "list"); strings.Contains(o, "sem") {
		t.Fatal(o)
	}
	cmd(1, "drop", "nope")
	cmd(1, "index", "nope")
	var o, e bytes.Buffer
	in := "TRIPLE predicate=P31 object=Q5\nTRIPLE predicate=P31 object=Q9\n"
	if c := Run(context.Background(), append(append([]string{"query", "-batch"}, ix...), data), strings.NewReader(in), &o, &e); c != 0 || !strings.HasPrefix(o.String(), "1\t") || strings.Contains(o.String(), "2\t") {
		t.Fatal(c, o.String(), e.String())
	}
	if c := Run(context.Background(), append(append([]string{"query", "-batch"}, ix...), data), strings.NewReader("NOPE\n"), &o, &e); c != 1 {
		t.Fatal(c)
	}
}
