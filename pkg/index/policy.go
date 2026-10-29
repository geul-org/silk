//ff:type feature=index type=model
//ff:what Policy — compaction triggers: appended share, run count, dead share, and the size ratio of a tier
package index

// Policy decides when a section is merged (docs/compaction.md). Zero values
// take DefaultPolicy's.
type Policy struct {
	AppendRatio float64 `json:"append_ratio"`
	MaxRuns     int     `json:"max_runs"`
	StaleRatio  float64 `json:"stale_ratio"`
	Fanout      int     `json:"fanout"`
}
