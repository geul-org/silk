//ff:func feature=sig type=decoder control=selection
//ff:what reservedCode — the reserved code a condition value names
package sig

func reservedCode(v string) (uint64, bool) {
	switch v {
	case "none":
		return CodeNone, true
	case "novalue":
		return CodeNoValue, true
	case "somevalue":
		return CodeSomeValue, true
	case "other":
		return CodeOther, true
	}
	return 0, false
}
