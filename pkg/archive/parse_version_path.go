//ff:func feature=archive type=decoder control=sequence
//ff:what parseVersionPath — `<Issuer>/<16-digit ID>/<8-digit version>.geul` as a Doc; any other path (versions.geul, h/, named) is false
package archive

import "strings"

// parseVersionPath reads an archive path in canonical form (publishing.md
// "ID 표기"). Version 0 is not a stored version.
func parseVersionPath(p string) (Doc, bool) {
	parts := strings.Split(p, "/")
	if len(parts) != 3 || !strings.HasPrefix(parts[0], "0x") || !strings.HasSuffix(parts[2], ".geul") {
		return Doc{}, false
	}
	id, ok := parseHexWidth(parts[1], 16)
	if !ok {
		return Doc{}, false
	}
	v, ok := parseHexWidth(strings.TrimSuffix(parts[2], ".geul"), 8)
	if !ok || v == 0 {
		return Doc{}, false
	}
	return Doc{Issuer: parts[0], ID: id, Version: uint32(v)}, true
}
