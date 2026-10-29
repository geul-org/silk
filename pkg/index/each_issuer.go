//ff:func feature=index type=helper control=iteration dimension=1
//ff:what eachIssuer — call fn once per Issuer of a data pack's documents (Issuers in order)
package index

import "github.com/park-jun-woo/silk/pkg/archive"

// eachIssuer splits the pack's documents by Issuer.
func eachIssuer(p archive.Pack, fn func(Target) error) error {
	issuers, by := byIssuer(p.Docs)
	for _, iss := range issuers {
		if err := fn(Target{Pack: p, Issuer: iss, Docs: by[iss]}); err != nil {
			return err
		}
	}
	return nil
}
