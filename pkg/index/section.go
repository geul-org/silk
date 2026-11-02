//ff:type feature=index type=model
//ff:what Section — one bundle of a declaration: its sorted runs and append array
package index

type Section struct {
	Name    string   `json:"name"`
	Decl    string   `json:"decl"`
	Bundle  int      `json:"bundle"`
	Pending bool     `json:"pending,omitempty"`
	Runs    []SegRef `json:"runs,omitempty"`
	Appends []SegRef `json:"appends,omitempty"`
}
