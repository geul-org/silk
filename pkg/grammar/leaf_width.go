//ff:func feature=grammar type=helper control=selection
//ff:what Leaf.Width — the default slot width of the field: rest 32, raw its bytes, language record 16, other record 6
package grammar

import "slices"

func (l Leaf) Width() int {
	switch {
	case l.Field.Width == RestWidth:
		return 32
	case l.Field.Width > 0:
		return 2 * l.Field.Width
	case slices.Contains(l.Field.Codecs, TagShort):
		return 16
	}
	return 6
}
