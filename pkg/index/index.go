//ff:type feature=index type=model
//ff:what Index — one pinned generation opened for reading
package index

import (
	"github.com/geul-org/silk/pkg/archive"
	"github.com/geul-org/silk/pkg/decl"
	"github.com/geul-org/silk/pkg/grammar"
)

type Index struct {
	Dir      string
	Manifest Manifest
	Grammar  *grammar.Grammar
	Live     Bitmap
	Docs     *DocTable
	Decls    map[string]decl.Decl
	ar       *archive.Archive
	pin      string
	secs     map[string][]*Sec
	cache    *docCache
}
