//ff:type feature=seg type=model
//ff:what cursor — a sequential reader over one Group's rows, chunk by chunk
package seg

// cursor yields rows (partition slot stripped) of g.
type cursor struct {
	g    Group
	buf  []byte
	i    int
	read uint64
}
