//ff:type feature=seg type=model
//ff:what cursor — a merge input: one segment walked partition by partition
package seg

type cursor struct {
	f    *File
	part uint64
	key  []byte
	reg  Region
	i    int
}
