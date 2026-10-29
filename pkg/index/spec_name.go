//ff:func feature=index type=helper control=sequence
//ff:what specName — the relative path of a spec stored by its hash
package index

// specName is spec/<sha256>.json.
func specName(hash string) string {
	return dirSpec + "/" + hash + ".json"
}
