//ff:func feature=sig type=encoder control=iteration dimension=1
//ff:what Schema.Encode — one entity's signature from its values per field name
package sig

// Encode builds a Words()-long signature; a field missing from vals stays 0.
func (s Schema) Encode(vals map[string][]uint64) ([]uint64, error) {
	out := make([]uint64, s.Words())
	for _, f := range s.Fields {
		if err := encodeField(out, f, vals[f.Name]); err != nil {
			return nil, err
		}
	}
	return out, nil
}
