//ff:func feature=index type=reader control=sequence
//ff:what readCheckpoints — the paused merges of compact.json (none when the file is missing or unreadable)
package index

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// readCheckpoints loads dir/compact.json.
func readCheckpoints(dir string) []checkpoint {
	b, err := os.ReadFile(filepath.Join(dir, checkpointFile))
	if err != nil {
		return nil
	}
	var cs []checkpoint
	if json.Unmarshal(b, &cs) != nil {
		return nil
	}
	return cs
}
