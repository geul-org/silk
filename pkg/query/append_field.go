//ff:func feature=query type=helper control=sequence
//ff:what appendField — add a field condition unless its value is nil (any)
package query

// appendField skips wildcards.
func appendField(fs []FieldCond, path []int, v []byte) []FieldCond {
	if v == nil {
		return fs
	}
	return append(fs, FieldCond{Path: path, Value: v})
}
