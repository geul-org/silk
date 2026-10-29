//ff:func feature=scan type=scanner control=iteration dimension=1
//ff:what scanChunk — the fused loop of one goroutine: frame each offset's record from the window and keep the ones the Matcher accepts
package scan

// scanChunk tests the records at offs (ascending) in one pass.
func scanChunk(w *window, offs []uint32, m Matcher) ([]Hit, error) {
	var out []Hit
	for _, off := range offs {
		rec, err := w.record(uint64(off))
		if err != nil {
			return nil, err
		}
		if m.Match(rec) {
			out = append(out, Hit{Off: off, Rec: append([]byte(nil), rec...)})
		}
	}
	return out, nil
}
