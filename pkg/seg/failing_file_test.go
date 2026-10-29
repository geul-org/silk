//ff:func feature=seg type=helper control=sequence
//ff:what failing — test helper: a copy of a File whose reads fail after n successful ReadAt calls
package seg

// failing shares f's footer and table.
func failing(f *File, n int) *File {
	c := *f
	c.r = &failReader{r: f.r, n: n}
	return &c
}
