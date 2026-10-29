//ff:func feature=spec type=rule control=sequence
//ff:what checkTag — a tag has a name, a layout SILK knows, a non-negative minimum partition and valid partition and sort keys
package spec

import "fmt"

// checkTag validates one tag.
func checkTag(t Tag) error {
	l, ok := LayoutOf(t.Kind, t.Codec)
	if !ok {
		return fmt.Errorf("%w: %q (Kind %d, Codec 0x%02x) has no fixed layout SILK can read without the codebook", ErrSpec, t.Name, t.Kind, t.Codec)
	}
	if t.Name == "" || t.MinPartition < 0 {
		return fmt.Errorf("%w: %s needs a name and min_partition >= 0", ErrSpec, l.Name)
	}
	if err := CheckKey(l, t.Partition, true); err != nil {
		return err
	}
	return CheckKey(l, t.Sort, false)
}
