//ff:func feature=fixture type=helper control=sequence
//ff:what fails — whether fn stops on Fatal
package fixture

import "testing"

func fails(t *testing.T, fn func(tb testing.TB)) bool {
	t.Helper()
	r := &recorder{TB: t}
	done := make(chan struct{})
	go func() {
		defer close(done)
		fn(r)
	}()
	<-done
	return r.failed
}
