//ff:type feature=query type=model
//ff:what Result — one version document meeting every condition, with the Edge Ids of its matching TRIPLEs
package query

// Result names `<Issuer>/<Doc>/<Version>.geul` and the matching Edge Ids
// (per condition, in record order).
type Result struct {
	Issuer  string
	Doc     uint64
	Version uint32
	Edges   []uint64
}
