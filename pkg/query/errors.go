package query

import "errors"

var (
	// ErrSyntax means a condition could not be read.
	ErrSyntax = errors.New("silk/query: bad condition")
	// ErrVerify means a record the scan accepted failed the exact check
	// (a bug or a record the byte matcher misread).
	ErrVerify = errors.New("silk/query: exact check failed")
)
