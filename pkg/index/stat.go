//ff:type feature=index type=model
//ff:what Stat — an index as reported
package index

type Stat struct {
	Generation uint64
	Docs       uint32
	Live       uint64
	MaxDoc     uint64
	Sections   []SectionStat
}
