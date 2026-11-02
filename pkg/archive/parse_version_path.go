//ff:func feature=archive type=decoder control=sequence
//ff:what parseVersionPath — `<Issuer>/<16-digit ID>/<8-digit version>.geul` as a Step; any other path (versions.geul, h/, named) is false
package archive

import "strings"

func parseVersionPath(p string) (Step, bool) {
	parts := strings.Split(p, "/")
	if len(parts) != 3 || !strings.HasPrefix(parts[0], "0x") || !strings.HasSuffix(parts[2], ".geul") {
		return Step{}, false
	}
	id, ok := parseHexWidth(parts[1], 16)
	if !ok {
		return Step{}, false
	}
	v, ok := parseHexWidth(strings.TrimSuffix(parts[2], ".geul"), 8)
	if !ok || v == 0 {
		return Step{}, false
	}
	return Step{Issuer: parts[0], ID: id, Version: uint32(v)}, true
}
