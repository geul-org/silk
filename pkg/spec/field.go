//ff:func feature=spec type=decoder control=iteration dimension=1
//ff:what Field — the bytes of the record field at a key path, framed by TAG WORDs alone; false when the record has no such field
package spec

// Field walks rec's embedded records to path[0]; a second step reads the
// language record that starts a UTF8-LANGUAGE-TEXT field (subRecord).
func Field(rec []byte, path []int) ([]byte, bool) {
	p, s := headLen(rec), 0
	for i := 0; i <= path[0]; i++ {
		var ok bool
		s = p
		if p, ok = nextField(rec, p); !ok {
			return nil, false
		}
	}
	if len(path) == 1 {
		return rec[s:p], true
	}
	return subRecord(rec[s:p])
}
