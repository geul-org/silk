//ff:type feature=sig type=model
//ff:what Pattern — a signature test: sig AND Mask equals Want, word by word
package sig

// Pattern accepts a signature s when s[i]&Mask[i] == Want[i] for every word.
type Pattern struct {
	Mask []uint64
	Want []uint64
}
