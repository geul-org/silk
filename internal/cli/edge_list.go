//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what edgeList — Edge Ids as `#7,#12`
package cli

import (
	"strconv"
	"strings"
)

// edgeList formats ids.
func edgeList(ids []uint64) string {
	parts := make([]string, len(ids))
	for i, id := range ids {
		parts[i] = "#" + strconv.FormatUint(id, 10)
	}
	return strings.Join(parts, ",")
}
