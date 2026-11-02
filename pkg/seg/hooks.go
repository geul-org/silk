//ff:type feature=seg type=model
//ff:what Hooks — what a merge asks its caller: which documents live, and a pause point every so many rows
package seg

type Hooks struct {
	Live  func(doc uint32) bool
	Pause func() error
}
