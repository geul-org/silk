//ff:type feature=index type=model
//ff:what syncer — one inventory replay: the table being written, the identity rows on their way to a sorted run
package index

type syncer struct {
	u    *updater
	dw   *docWriter
	sp   *spill
	next uint32
	pack uint32
	iss  map[string]uint32
}
