//ff:func feature=index type=writer control=sequence
//ff:what writeCheckpoints — save paused merges to compact.json (removed when there are none); bytes written
package index

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// writeCheckpoints replaces dir/compact.json.
func writeCheckpoints(dir string, cs []checkpoint) (int64, error) {
	p := filepath.Join(dir, checkpointFile)
	if len(cs) == 0 {
		os.Remove(p)
		return 0, nil
	}
	b, _ := json.MarshalIndent(cs, "", "  ")
	return writeAtomic(p, b)
}
