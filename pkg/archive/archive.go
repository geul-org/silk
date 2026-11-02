//ff:type feature=archive type=model
//ff:what Archive — an opened pack archive: its `.well-known/geul/` directory and its sets at their latest generation
package archive

type Archive struct {
	Root string
	Sets []*Set
}
