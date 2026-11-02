//ff:func feature=index type=builder control=sequence
//ff:what mergeKept — merge segments into out without removing the inputs
package index

import "github.com/geul-org/silk/pkg/seg"

func mergeKept(paths []string, info seg.Info, hk seg.Hooks, out string) (string, error) {
	in, err := openSegs(paths)
	if err != nil {
		return "", err
	}
	defer closeSegs(in)
	w, err := seg.NewWriter(out, info)
	if err != nil {
		return "", err
	}
	if _, err := seg.Merge(w, in, hk); err != nil {
		w.Abort()
		return "", err
	}
	_, err = w.Finish()
	return out, err
}
