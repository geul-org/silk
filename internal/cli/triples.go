//ff:type feature=cli type=model
//ff:what triples — the repeatable -triple flag: conditions in the order given
package cli

// triples implements flag.Value.
type triples []string
