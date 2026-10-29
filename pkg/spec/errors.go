package spec

import "errors"

// ErrSpec means a specification breaks the format or names a key position
// SILK cannot read without the codebook.
var ErrSpec = errors.New("silk/spec: bad index specification")
