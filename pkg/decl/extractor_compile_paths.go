//ff:func feature=decl type=builder control=iteration dimension=1
//ff:what Extractor.compilePaths — resolve the partition and key paths on the first TAG
package decl

func (x *Extractor) compilePaths() error {
	if x.d.Part != "" {
		l, err := x.g.Resolve(x.d.Tags[0], x.d.Part)
		if err != nil {
			return err
		}
		x.part = &l
	}
	for _, k := range x.d.Keys {
		l, err := x.g.Resolve(x.d.Tags[0], k)
		if err != nil {
			return err
		}
		x.keys = append(x.keys, l)
	}
	return nil
}
