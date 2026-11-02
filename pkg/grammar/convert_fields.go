//ff:func feature=grammar type=builder control=iteration dimension=1
//ff:what convertFields — gogeul's field list as JSON-friendly fields
package grammar

import "github.com/geul-org/gogeul/pkg/codecspec"

func convertFields(fs []codecspec.Field) []Field {
	out := make([]Field, len(fs))
	for i, f := range fs {
		out[i] = Field{Name: f.Name, Width: f.Width}
		for _, c := range f.Codecs {
			out[i].Codecs = append(out[i].Codecs, uint16(c))
		}
	}
	return out
}
