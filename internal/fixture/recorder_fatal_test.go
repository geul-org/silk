//ff:func feature=fixture type=helper control=sequence
//ff:what recorder.Fatal — note the failure instead of stopping the test
package fixture

// Fatal records that it was called.
func (r *recorder) Fatal(...any) {
	r.failed = true
}
