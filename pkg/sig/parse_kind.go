//ff:func feature=sig type=decoder control=selection
//ff:what parseKind — a kind word
package sig

import "fmt"

func parseKind(s string) (Kind, error) {
	switch s {
	case "enum":
		return KindEnum, nil
	case "flag":
		return KindFlag, nil
	case "prefix":
		return KindPrefix, nil
	case "bucket":
		return KindBucket, nil
	}
	return 0, fmt.Errorf("kind %q", s)
}
