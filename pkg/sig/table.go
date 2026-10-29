//ff:type feature=sig type=model
//ff:what Table — a column of signatures, Words uint64 words per row, rows back to back
package sig

// Table holds signatures row-major; row i is Data[i*Words : (i+1)*Words].
type Table struct {
	Words int
	Data  []uint64
}
