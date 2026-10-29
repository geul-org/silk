//ff:type feature=index type=model
//ff:what Section — one TAG kind's segments: name, TAG WORD, tag spec hash, sorted runs, appended segments, dead rows they still hold
package index

// Section holds the index entries of one spec tag. Dead counts rows whose
// document version is no longer live; merges drop them.
type Section struct {
	Name    string   `json:"name"`
	Tag     uint16   `json:"tag"`
	Hash    string   `json:"hash"`
	Runs    []SegRef `json:"runs"`
	Appends []SegRef `json:"appends"`
	Dead    uint64   `json:"dead"`
}
