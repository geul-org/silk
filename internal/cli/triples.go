//ff:type feature=cli type=model
//ff:what triples — a repeatable string flag (-triple, -match, -label): values in the order given
package cli

// triples implements flag.Value.
type triples []string
