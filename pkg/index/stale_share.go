//ff:func feature=index type=reader control=iteration dimension=1
//ff:what staleShare — the share of dead rows, from evenly spaced probes of each segment
package index

func staleShare(dir string, s Section, live Bitmap) (float64, error) {
	in, err := openSegs(segPaths(dir, s))
	if err != nil {
		return 0, err
	}
	defer closeSegs(in)
	var dead, seen float64
	for _, f := range in {
		d, n := probeDead(f, live)
		w := float64(f.Rows) / max(float64(n), 1)
		dead, seen = dead+float64(d)*w, seen+float64(n)*w
	}
	if seen == 0 {
		return 0, nil
	}
	return dead / seen, nil
}
