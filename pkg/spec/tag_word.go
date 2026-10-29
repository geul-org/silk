//ff:func feature=spec type=helper control=sequence
//ff:what TagWord — Kind and Codec as a TAG WORD with only those bits set (Kind bits 12-11, Codec bits 6-0)
package spec

// TagWord is gogeul core.TargetWord(kind, codec) for direct codecs.
func TagWord(kind, codec uint8) uint16 {
	return uint16(kind&3)<<11 | uint16(codec&0x7F)
}
