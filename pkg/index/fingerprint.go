//ff:type feature=index type=model
//ff:what fingerprint — an order-free fingerprint: row count and the sum of row hashes
package index

type fingerprint struct {
	n, sum uint64
}
