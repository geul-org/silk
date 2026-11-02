//ff:type feature=index type=model
//ff:what updater — one build or update run: the manifest being advanced, its liveness and temporary directory
package index

import (
	"sync/atomic"

	"github.com/geul-org/silk/pkg/archive"
)

type updater struct {
	ar    *archive.Archive
	dir   string
	o     Options
	m     Manifest
	tmp   string
	n     atomic.Uint64
	live  Bitmap
	first uint32
}
