//ff:type feature=cli type=model
//ff:what indexFlags — the flags of `silk index` subcommands
package cli

// indexFlags holds parsed values; not every subcommand uses every flag.
type indexFlags struct {
	dir     string
	spec    string
	threads int
	compact bool
	full    bool
	idle    bool
	dry     bool
}
