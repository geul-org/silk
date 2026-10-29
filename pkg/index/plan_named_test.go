//ff:func feature=index type=builder control=sequence
//ff:what TestPlanNamed — due tasks get distinct output names and their tag's minimum partition
package index

import "testing"

func TestPlanNamed(t *testing.T) {
	_, _, dir := updated(t)
	m, _ := CurrentManifest(dir)
	ts := planNamed(&m, smallSpec(), Options{Full: true})
	if len(ts) != 7 || ts[0].Output == ts[1].Output || ts[0].Min != 2 || ts[0].Output == "" {
		t.Fatal(ts)
	}
}
