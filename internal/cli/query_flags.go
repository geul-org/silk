//ff:type feature=cli type=model
//ff:what queryFlags — the flags of `silk query`: index directory, threads, output mode, conditions in the order given
package cli

// queryFlags holds parsed values.
type queryFlags struct {
	dir     string
	threads int
	count   bool
	explain bool
	triples triples
	matches triples
	labels  triples
}
