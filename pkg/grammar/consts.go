// Package grammar reads the GEUL grammar a SILK index needs from the GEUL
// Issuer set `0x0000`: every codec's fields, its name, and the language
// registry. It also frames documents and records at the byte level, walks
// field paths, cuts key slots and turns gogeul notation tokens into bytes.
package grammar

import "errors"

const (
	kindReference = 0
	kindValue     = 1
	kindEdge      = 2

	// RestWidth is the declared width of a field holding the remaining words.
	RestWidth = 65535

	// Language, text and slot marker target words (grammar value/text, reference).
	TagLocal     = 0x0008
	TagBCP47     = 0x0846
	TagShort     = 0x0848
	TagText      = 0x0840
	TagTextOdd   = 0x0841
	TagLangText  = 0x084A
	TagLangOdd   = 0x084B
	TagTriple    = 0x100A
	TagRefNode   = 0x1000
	markerUnused = 0
)

var (
	// ErrSyntax means a path or value token cannot be read.
	ErrSyntax = errors.New("silk/grammar: bad syntax")
	// ErrNoGrammar means no codec spec documents were found.
	ErrNoGrammar = errors.New("silk/grammar: no codec spec documents")
	// ErrFrame means a document's records do not frame.
	ErrFrame = errors.New("silk/grammar: record does not frame")
)
