//ff:func feature=seg type=helper control=sequence
//ff:what row — test helper: a partition row of a 2-byte sort key, document and offset
package seg

// row builds sort key k (2 bytes) + doc + off.
func row(k string, doc, off uint32) []byte {
	return AppendSeq([]byte(k), doc, off)
}
