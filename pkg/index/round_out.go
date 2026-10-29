//ff:type feature=index type=model
//ff:what roundOut — one task's outcome in a compaction round: its result and error
package index

// roundOut pairs a taskResult with its error.
type roundOut struct {
	res taskResult
	err error
}
