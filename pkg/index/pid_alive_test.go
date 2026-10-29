//ff:func feature=index type=helper control=sequence
//ff:what TestPidAlive — this process is alive; 0 and an unused process ID are not
package index

import (
	"os"
	"testing"
)

func TestPidAlive(t *testing.T) {
	if !pidAlive(os.Getpid()) || pidAlive(0) || pidAlive(1<<30) {
		t.Fatal("pidAlive")
	}
}
