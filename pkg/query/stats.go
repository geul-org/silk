//ff:type feature=query type=model
//ff:what Stats — what a run touched: segments, rows examined, rows of dead versions, candidates, records that passed the exact check
package query

// Stats counts a run's work over all conditions.
type Stats struct {
	Segments   int
	Rows       uint64
	Dead       uint64
	Candidates uint64
	Verified   uint64
}
