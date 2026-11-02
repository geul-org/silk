//ff:type feature=sig type=model
//ff:what Entry — one declared code: name, bits and their length, the values it stands for, and for a bucket its half-open range
package sig

type Entry struct {
	Name   string
	Code   uint64
	Len    int
	Values [][]byte
	Lo, Hi int64
}
