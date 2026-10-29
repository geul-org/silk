// Package archive opens a GEUL pack archive for indexing: its sets, their data
// packs as local files and the version documents each data pack stores. It
// reads through gogeul's pkg/pack and accepts only records stored raw, since
// a record map points into data pack bytes.
package archive

const (
	manifestMagic = "GEULMANI"
	manifestHead  = 64 // fixed manifest header (GEUL pack format §4.3)
	sectionSize   = 24 // section table entry
	objectSize    = 48 // OBJECTS entry
	secObjects    = 0x0001
	kindData      = 1
	encRaw        = 0
)
