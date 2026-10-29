//ff:func feature=index type=helper control=sequence
//ff:what manifestName — the relative path of a generation's manifest
package index

import "fmt"

// manifestName is manifest/<generation, 16 hex digits>.json.
func manifestName(gen uint64) string {
	return fmt.Sprintf("%s/%016x.json", dirManifest, gen)
}
