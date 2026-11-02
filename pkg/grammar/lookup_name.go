//ff:func feature=grammar type=decoder control=iteration dimension=1
//ff:what lookupName — a field by name; its number counts an Edge's Edge Id
package grammar

func lookupName(fs []Field, name string, edge int) (int, Field, bool) {
	for i, f := range fs {
		if f.Name == name {
			return i + edge, f, true
		}
	}
	return 0, Field{}, false
}
