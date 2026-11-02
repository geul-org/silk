// Package sig is the SILK definition: the bit fields of a semantic row, the
// reserved codes, the encoder that spreads several values into several rows,
// and the query patterns that jump through a sorted array of such rows.
// Field order is bit order and so sort order. See docs/silk.md.
package sig

import "errors"

const (
	// Reserved codes of every enum, prefix and bucket field.
	CodeNone      = 0 // no statement
	CodeNoValue   = 1 // novalue
	CodeSomeValue = 2 // somevalue
	CodeOther     = 3 // outside the list
	firstCode     = 4

	// MaxExpand bounds the rows of one entity before spreading is limited to one field.
	MaxExpand = 3

	wordNoValue    = "\x00\x02"
	wordSomeValue  = "\x00\x01"
	rankDeprecated = 2
)

// ErrDef means a definition cannot be read.
var ErrDef = errors.New("silk/sig: bad definition")

// ErrCond means a query condition does not fit the definition.
var ErrCond = errors.New("silk/sig: bad condition")

const (
	KindEnum Kind = iota
	KindFlag
	KindPrefix
	KindBucket
)

const (
	tagSigned    = 0x0801
	tagUnsigned  = 0x0802
	tagGregorian = 0x0820
	tagEDTF      = 0x0824
	tagEDTFOdd   = 0x0825
	tagGEDCOM    = 0x0826
	tagGEDCOMOdd = 0x0827
)
