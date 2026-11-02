//ff:type feature=grammar type=model
//ff:what Leaf — a resolved path: its numbers and the shape of the field it ends on
package grammar

type Leaf struct {
	Path  Path
	Field Field
}
