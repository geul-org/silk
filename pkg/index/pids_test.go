//ff:func feature=index type=scanner control=sequence
//ff:what TestPids — process checks: ours is alive, 0 is not, a free number is not
package index

import (
	"os"
	"testing"
)

func TestPids(t *testing.T) {
	if !pidAlive(os.Getpid()) || pidAlive(0) || pidAlive(1<<30) {
		t.Fatal("pidAlive")
	}
}
