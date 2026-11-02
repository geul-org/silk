//ff:func feature=query type=helper control=iteration dimension=1
//ff:what randomQueries — conditions over the corpus: field equalities, text, silk and joins
package query

import (
	"math/rand"
	"strconv"
)

func randomQueries(r *rand.Rand, n int) [][]string {
	vals := []string{"Q82955", "Q36180", "Q33999", "Q169470", "Q98"}
	types := []string{"Q5", "Q16521", "Q515", "Q99"}
	names := []string{"Paris", "Par", "Paris Hilton", "Parisian culture", "김치", "김치찌개", "대한민국", "Ab", "Abcdefgh", "Abcdefghi", "Abcdefghij", "없음"}
	silks := []string{"job=physicist", "job=writer", "job=other", "sex=male", "born=1931", "born=old", "type=person", "type=taxon", "type=somevalue"}
	var out [][]string
	for range n {
		var q []string
		switch r.Intn(7) {
		case 0:
			q = []string{"TRIPLE", "predicate=P106", "object=" + vals[r.Intn(len(vals))]}
		case 1:
			q = []string{"TRIPLE", "predicate=P31", "object=" + types[r.Intn(len(types))], "TRIPLE", "predicate=P106", "object=" + vals[r.Intn(len(vals))]}
		case 2:
			q = []string{"LABEL", "object.language=@" + []string{"en", "ko", "de"}[r.Intn(3)], "object.bytes=" + strconv.Quote(names[r.Intn(len(names))])}
		case 3:
			q = []string{"ALIAS", "object.bytes=" + strconv.Quote(names[r.Intn(len(names))])}
		case 4:
			q = []string{"entity", silks[r.Intn(len(silks))]}
		case 6:
			q = []string{"label-sem", "object.language=@en", "object.bytes^=" + strconv.Quote(names[r.Intn(4)][:3]), "type=" + []string{"person", "taxon", "somevalue"}[r.Intn(3)]}
		case 5:
			q = []string{"LABEL", "object.language=@en", "object.bytes^=" + strconv.Quote(names[r.Intn(4)][:3]), "entity", silks[r.Intn(len(silks))]}
		}
		out = append(out, q)
	}
	return out
}
