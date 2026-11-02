//ff:func feature=fixture type=helper control=sequence
//ff:what recorder.Fatal — note the failure and stop the goroutine
package fixture

import "runtime"

func (r *recorder) Fatal(...any) {
	r.failed = true
	runtime.Goexit()
}
