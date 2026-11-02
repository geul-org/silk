//ff:type feature=index type=model
//ff:what task — one merge: a section's chosen runs and appends
package index

type task struct {
	sec     *Section
	runs    []SegRef
	appends []SegRef
}
