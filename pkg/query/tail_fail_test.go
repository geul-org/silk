//ff:type feature=query type=model
//ff:what tailFail — a ReaderAt over map bytes that fails reads at or after an offset (the off[] table)
package query

// tailFail serves b below from and fails from on.
type tailFail struct {
	b    []byte
	from int64
}
