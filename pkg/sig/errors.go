package sig

import "errors"

var (
	// ErrSchema means a schema declaration is inconsistent.
	ErrSchema = errors.New("silk/sig: bad schema")
	// ErrValue means values or conditions do not fit the schema.
	ErrValue = errors.New("silk/sig: bad value")
	// ErrUnsupported means a field kind this version does not encode yet.
	ErrUnsupported = errors.New("silk/sig: field kind not implemented")
)
