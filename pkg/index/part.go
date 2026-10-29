//ff:type feature=index type=model
//ff:what Part — one opened map with its data pack file, ready to scan
package index

import (
	"os"

	"github.com/park-jun-woo/silk/pkg/recmap"
)

// Part pairs a record map with the data pack it points into.
type Part struct {
	Map  *recmap.File
	Data *os.File
}
