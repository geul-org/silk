//ff:type feature=seg type=model
//ff:what Region — rows of one partition: a byte slice of fixed-width sorted rows
package seg

type Region struct {
	B []byte
	W int
}
