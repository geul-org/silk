//ff:type feature=archive type=model
//ff:what Pack — one data pack of a set as a local file with the live version documents it stores, in offset order
package archive

// Pack is a data pack object (`o/<Hash>.gpk`) of Set and the version
// documents of the set's generation that lie in it, by ascending Off.
type Pack struct {
	Set  Set
	Hash string
	Path string
	Size uint64
	Docs []Doc
}
