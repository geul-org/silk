//ff:func feature=index type=scanner control=sequence
//ff:what TestUpdaterTargetOf — one target per declaration, an unreadable one fails
package index

import "testing"

func TestUpdaterTargetOf(t *testing.T) {
	_, dir, a := built(t, 1)
	u := &updater{ar: a, dir: dir, o: testOpts, m: current(t, dir), tmp: t.TempDir()}
	g, _ := loadGrammar(dir, u.m.Grammar)
	by := map[string]*target{}
	x, err := u.targetOf(by, &u.m.Sections[0], g)
	if err != nil {
		t.Fatal(err)
	}
	y, err := u.targetOf(by, &u.m.Sections[0], g)
	if err != nil || x != y {
		t.Fatal(err)
	}
	if _, err := u.targetOf(by, &Section{Decl: "none"}, g); err == nil {
		t.Fatal("no declaration")
	}
}
