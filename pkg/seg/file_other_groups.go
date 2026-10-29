//ff:func feature=seg type=reader control=iteration dimension=1
//ff:what File.otherGroups — split the other partition's rows into one Group per partition key (its rows are grouped by key)
package seg

import "bytes"

// otherGroups reads the other partition once and appends its groups to out.
func (f *File) otherGroups(out []Group) ([]Group, error) {
	rg := f.otherRegion()
	b, err := f.readRows(rg, 0, rg.n)
	if err != nil {
		return nil, err
	}
	pw := f.Info.PartWidth
	start := uint64(0)
	for i := uint64(1); i <= rg.n; i++ {
		if i < rg.n && bytes.Equal(b[i*uint64(rg.w):i*uint64(rg.w)+uint64(pw)], b[start*uint64(rg.w):start*uint64(rg.w)+uint64(pw)]) {
			continue
		}
		key := append([]byte{}, b[start*uint64(rg.w):start*uint64(rg.w)+uint64(pw)]...)
		out = append(out, Group{Key: key, f: f, rg: region{off: rg.off + int64(start)*int64(rg.w), n: i - start, w: rg.w}, strip: pw})
		start = i
	}
	return out, nil
}
