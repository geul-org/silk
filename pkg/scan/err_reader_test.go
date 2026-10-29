//ff:type feature=scan type=model
//ff:what errReader — a ReaderAt that always fails
package scan

// errReader fails every read.
type errReader struct{}
