//ff:func feature=index type=builder control=sequence
//ff:what mergeFiles — merge some temporary runs into out and remove them
package index

import "github.com/geul-org/silk/pkg/seg"

func mergeFiles(paths []string, info seg.Info, hk seg.Hooks, out string) (string, error) {
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
	if _, err := w.Finish(); err != nil {
		return "", err
	}
	removeAll(paths)
	return out, nil
}
