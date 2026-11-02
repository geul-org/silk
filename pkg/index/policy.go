//ff:type feature=index type=model
//ff:what Policy — when segments merge: AppendRows bounds the append array in rows, RunsPerTier the runs of one size tier, StaleRatio the dead rows of a section
package index

type Policy struct {
	AppendRows  uint64
	RunsPerTier int
	StaleRatio  float64
}
