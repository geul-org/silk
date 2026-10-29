// Package query answers SILK conditions over an index. A condition names a
// TAG kind and field values (a TRIPLE's Predicate and Object, a LABEL's
// language, ...); one document must meet every condition. The planner maps
// each condition onto its section's spec: a field that is the partition key
// picks the partition, a field that is the sort key is binary-searched in
// sorted runs (appended segments are scanned), and every candidate record is
// then decoded with gogeul and checked field by field. Conditions are joined
// by document with a merge join on sequence numbers.
package query

import "errors"

var (
	// ErrSyntax means a condition could not be read.
	ErrSyntax = errors.New("silk/query: bad condition")
	// ErrVerify means a candidate record is not a record of the condition's
	// kind (a bug or a damaged index).
	ErrVerify = errors.New("silk/query: exact check failed")
	// ErrNoSection means the index spec does not index the condition's kind.
	ErrNoSection = errors.New("silk/query: the index does not cover this kind")
)
