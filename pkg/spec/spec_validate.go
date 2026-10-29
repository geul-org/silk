//ff:func feature=spec type=rule control=iteration dimension=1
//ff:what Spec.Validate — format version, at least one tag, every tag valid, names and Kind·Codec unique
package spec

import "fmt"

// Validate checks the whole specification.
func (s Spec) Validate() error {
	if s.Version != Version || len(s.Tags) == 0 {
		return fmt.Errorf("%w: version %d with %d tags", ErrSpec, s.Version, len(s.Tags))
	}
	names, kinds := map[string]bool{}, map[uint16]bool{}
	for _, t := range s.Tags {
		if err := checkTag(t); err != nil {
			return err
		}
		w := TagWord(t.Kind, t.Codec)
		if names[t.Name] || kinds[w] {
			return fmt.Errorf("%w: %s declared twice", ErrSpec, t.Name)
		}
		names[t.Name], kinds[w] = true, true
	}
	return nil
}
