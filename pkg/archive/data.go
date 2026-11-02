//ff:type feature=archive type=model
//ff:what Data — one data pack of a set: object number and hash
package archive

type Data struct {
	Obj  uint32
	Hash string
}
