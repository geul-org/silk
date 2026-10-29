// Package index keeps SILK's index directory: one record map per data pack
// and Issuer of a pack archive (`<dir>/<Issuer>/<pack SHA-256>.srm`) and a
// catalog (`<dir>/index.json`) naming the set generations they were built
// from and the maps. It builds them, opens them against the archive's current
// generations and verifies them.
package index

const (
	mapExt         = ".srm"
	catalogFile    = "index.json"
	catalogVersion = 1
)
