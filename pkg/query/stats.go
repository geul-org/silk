//ff:type feature=query type=model
//ff:what Stats — what a run touched: maps, records scanned, records the scan accepted
package query

// Stats counts a run's work.
type Stats struct {
	Maps    int
	Scanned int
	Hits    int
}
