//ff:type feature=recmap type=model
//ff:what KeyRange — one key (Kind·Codec TAG WORD) and the slice [Start, Start+Count) of off[] holding its records
package recmap

// KeyRange locates the records of one key in the sorted off[].
type KeyRange struct {
	Key   uint16
	Start uint64
	Count uint64
}
