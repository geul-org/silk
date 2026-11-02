//ff:func feature=cli type=decoder control=iteration dimension=1
//ff:what parseArgs — flags may stand between positional words; returns the positional words
package cli

func (f *flags) parseArgs(args []string) ([]string, error) {
	var pos []string
	for {
		if err := f.set.Parse(args); err != nil {
			return nil, err
		}
		if f.set.NArg() == 0 {
			return pos, nil
		}
		pos = append(pos, f.set.Arg(0))
		args = f.set.Args()[1:]
	}
}
