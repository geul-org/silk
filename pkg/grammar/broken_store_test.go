//ff:type feature=grammar type=model
//ff:what brokenStore — a store whose document reads fail
package grammar

import "github.com/geul-org/gogeul/pkg/pack"

type brokenStore struct {
	*pack.Archive
}
