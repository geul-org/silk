//ff:type feature=index type=model
//ff:what VerifyReport — what Verify checked: generation, sections, segments, live documents, live rows
package index

// VerifyReport counts a successful verification.
type VerifyReport struct {
	Generation uint64
	Sections   int
	Segments   int
	Docs       int
	Rows       uint64
}
