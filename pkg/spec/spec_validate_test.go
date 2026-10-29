//ff:func feature=spec type=rule control=iteration dimension=1
//ff:what TestSpecValidate — the default spec passes; wrong version, no tags, a bad tag, a repeated name or Kind·Codec fail
package spec

import "testing"

func TestSpecValidate(t *testing.T) {
	if err := Default().Validate(); err != nil {
		t.Fatal(err)
	}
	d := Default()
	dupName := Default()
	dupName.Tags[1].Name = "TRIPLE"
	dupKind := Default()
	dupKind.Tags[1].Codec = 0x0A
	dupKind.Tags[1].Partition = Key{Path: []int{1}, Width: 6}
	badTag := Default()
	badTag.Tags[0].Codec = 0x7F
	for i, s := range []Spec{{Version: 2, Tags: d.Tags}, {Version: Version}, badTag, dupName, dupKind} {
		if s.Validate() == nil {
			t.Fatalf("case %d passed", i)
		}
	}
}
