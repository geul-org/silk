//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what labelMatches — -label languages as -match conditions on the LABEL language record
package cli

// labelMatches maps L to LABEL:2.0=@L.
func labelMatches(ls []string) []string {
	out := make([]string, 0, len(ls))
	for _, l := range ls {
		out = append(out, "LABEL:2.0=@"+l)
	}
	return out
}
