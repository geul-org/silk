// Package query compiles conditions to field predicates and silk patterns,
// plans each on the best section, narrows candidates through rows, and opens
// only the documents an exact check still needs. See docs/query.md.
package query

import "errors"

const (
	// Result units.
	ModeDocs    = "docs"
	ModeRecords = "records"
	ModeKeys    = "keys"
)

var (
	// ErrSyntax means a condition cannot be read.
	ErrSyntax = errors.New("silk/query: bad condition")
	// ErrNoPlan means no section can answer and -scan was not given.
	ErrNoPlan = errors.New("silk/query: no section answers this condition")
	// ErrNotCovered means -keys was asked of a section that does not cover the condition.
	ErrNotCovered = errors.New("silk/query: section does not cover the condition; -keys cannot answer")
)
