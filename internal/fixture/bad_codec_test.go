//ff:type feature=fixture type=model
//ff:what badCodec — a pack.Codec that always fails
package fixture

// badCodec makes every pack write fail.
type badCodec struct{}
