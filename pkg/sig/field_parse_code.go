//ff:func feature=sig type=decoder control=selection
//ff:what Field.parseCode — enum `NUMBER VALUE…`, prefix `0bBITS VALUE…`, bucket `NUMBER LO..HI`; a flag has no codes
package sig

import "errors"

func (f *Field) parseCode(name, code string, rest []string) (*Entry, error) {
	switch f.Kind {
	case KindEnum:
		return enumCode(name, code, rest, f.Width)
	case KindPrefix:
		return prefixCode(name, code, rest)
	case KindBucket:
		return bucketCode(name, code, rest, f.Width)
	}
	return nil, errors.New("a flag has no codes")
}
