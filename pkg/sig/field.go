//ff:type feature=sig type=model
//ff:what Field — one declared bit field: position, width, source predicate, kind and its codes
package sig

type Field struct {
	Name    string
	Off     int
	Width   int
	Pred    []byte
	Kind    Kind
	Codes   []*Entry
	byValue map[string]*Entry
}
