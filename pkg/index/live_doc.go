//ff:type feature=index type=model
//ff:what liveDoc — a live version document of the archive: its set, data pack (hash and file) and place in the pack
package index

import "github.com/geul-org/silk/pkg/archive"

// liveDoc is one document an index must cover.
type liveDoc struct {
	Set  string
	Pack string
	Path string
	Doc  archive.Doc
}
