//ff:type feature=archive type=model
//ff:what Archive — an opened pack archive: its `.well-known/geul/` directory and its sets at their latest generation
package archive

// Archive is a pack archive found under Root (a `.well-known/geul/`
// directory), one Set per pack set.
type Archive struct {
	Root string
	Sets []Set
}
