//ff:func feature=index type=builder control=sequence
//ff:what DefaultPolicy — flush appended rows above 10%, merge everything above 8 segments or 30% dead rows, merge 4 runs of one size tier
package index

// DefaultPolicy is the compaction policy of the design discussion.
func DefaultPolicy() Policy {
	return Policy{AppendRatio: 0.10, MaxRuns: 8, StaleRatio: 0.30, Fanout: 4}
}
