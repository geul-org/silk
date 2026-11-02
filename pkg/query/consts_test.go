package query

import "github.com/geul-org/silk/pkg/index"

const testDef = `# test definition
silk entity 128
field type 0 6 TRIPLE P31 enum
  code person 4 Q5
  code taxon 5 Q16521
  code city 6 Q515
when type=person
field job 6 8 TRIPLE P106 enum
  code politician 4 Q82955
  code writer 5 Q36180
  code actor 6 Q33999
  code physicist 7 Q169470
field sex 14 3 TRIPLE P21 enum
  code male 4 Q6581097
  code female 5 Q6581072
field born 17 4 TRIPLE P569 bucket
  code old 4 ..1900
  code early 5 1900..1950
  code late 6 1950..
`

const testSem = `silk sem 32
field type 0 6 TRIPLE P31 enum
  code person 4 Q5
  code taxon 5 Q16521
`

var testOpts = index.Options{Limit: 64 << 20, Threads: 2, Policy: index.DefaultPolicy()}
