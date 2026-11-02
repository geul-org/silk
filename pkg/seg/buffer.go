//ff:type feature=seg type=model
//ff:what Buffer — rows held in memory until sorted and written: partition key and row side by side
package seg

type Buffer struct {
	pw, rw int
	data   []byte
	idx    []uint32
}
