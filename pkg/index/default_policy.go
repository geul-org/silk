//ff:func feature=index type=builder control=sequence
//ff:what DefaultPolicy — the design's values: 10,000 append rows, 8 runs per tier, 30% stale
package index

func DefaultPolicy() Policy {
	return Policy{AppendRows: 10000, RunsPerTier: 8, StaleRatio: 0.30}
}
