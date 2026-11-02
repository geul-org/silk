//ff:type feature=index type=model
//ff:what SectionStat — one section as reported: its segments and an estimate of its dead rows
package index

type SectionStat struct {
	Name    string
	Kind    string
	Bundle  int
	Pending bool
	Rows    uint64
	Bytes   uint64
	Runs    int
	Appends uint64
	Stale   float64
}
