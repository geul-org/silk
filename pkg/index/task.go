//ff:type feature=index type=model
//ff:what Task — one merge of a section: kind (flush, tier, major), input segments, output file, minimum partition of its tag
package index

// Task replaces Inputs of section Section with one sorted run at Output.
type Task struct {
	Section int
	Name    string
	Kind    string
	Inputs  []SegRef
	Output  string
	Min     int
}
