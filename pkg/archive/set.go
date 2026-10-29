//ff:type feature=archive type=model
//ff:what Set — one pack set: Issuer folder name ("" at the top), directory, manifest SHA-256 hex and generation
package archive

// Set is one pack set (`latest.txt`, `m/`, `o/`) at the generation it had
// when the archive was opened.
type Set struct {
	Name       string
	Dir        string
	Manifest   string
	Generation uint64
}
