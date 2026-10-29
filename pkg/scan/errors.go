package scan

import "errors"

// ErrRecord means a map offset does not frame a record inside the data pack.
var ErrRecord = errors.New("silk/scan: no record at offset")
