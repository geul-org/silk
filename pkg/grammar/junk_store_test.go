//ff:type feature=grammar type=model
//ff:what junkStore — a store whose latest version is not a document but whose list still is
package grammar

import "github.com/geul-org/gogeul/pkg/pack"

type junkStore struct {
	*pack.Archive
}
