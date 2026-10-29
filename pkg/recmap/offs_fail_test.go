//ff:type feature=recmap type=model
//ff:what offsFail — a ReaderAt that serves every read but the one starting at off[] (the read Verify makes after hashing)
package recmap

// offsFail fails reads that start exactly at at.
type offsFail struct {
	b  []byte
	at int64
}
