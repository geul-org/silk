//ff:func feature=index type=helper control=sequence
//ff:what manifestName — a generation's manifest path
package index

import "fmt"

func manifestName(gen uint64) string {
	return fmt.Sprintf("%s/%016x.json", dirManifest, gen)
}
