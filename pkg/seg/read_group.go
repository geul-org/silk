//ff:func feature=seg type=reader control=iteration dimension=1
//ff:what readGroup — all rows of a group, partition slot stripped, as one buffer
package seg

// readGroup drains a cursor.
func readGroup(g Group) ([]byte, error) {
	c := &cursor{g: g}
	var out []byte
	for {
		row, ok, err := c.next()
		if !ok {
			return out, err
		}
		out = append(out, row...)
	}
}
