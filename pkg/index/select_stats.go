//ff:type feature=index type=model
//ff:what SelectStats — what a selection touched: segments, rows examined, rows of dead versions skipped, live rows returned
package index

// SelectStats counts one Index.Select.
type SelectStats struct {
	Segments int
	Rows     uint64
	Dead     uint64
	Hits     uint64
}
