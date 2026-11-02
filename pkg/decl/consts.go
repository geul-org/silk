// Package decl holds section declarations: what kind of index, on which TAGs
// and fields, and how one document's records become its rows. A declaration
// is plain data; its SHA-256 goes into every segment it produces.
package decl

import "errors"

const (
	// Kinds of declaration.
	KindIndex     = "index"
	KindText      = "text"
	KindSilk      = "silk"
	KindComposite = "composite"

	// TextWidth is the text key of TEXT and composite rows.
	TextWidth = 8
	// LangWidth is the language partition key of TEXT and composite rows.
	LangWidth = 16
	// SIDXMeaning is the bit width of a composite row's meaning half.
	SIDXMeaning = 32
)

// ErrDecl means a declaration cannot be built.
var ErrDecl = errors.New("silk/decl: bad declaration")
