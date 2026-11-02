package sig

const defText = `# test definition
silk entity 64
field type 0 6 TRIPLE P31 enum
  code person 4 Q5 Q15632617
  code city 5 Q515
field alive 6 1 TRIPLE P570 flag
when type=person
field job 7 8 TRIPLE P106 enum
  code politician 4 Q82955
  code writer 5 Q36180
field area 15 6 TRIPLE P27 prefix
  code asia 0b01
  code korea 0b0100 Q884
  code japan 0b0101 Q17
field born 21 4 TRIPLE P569 bucket
  code old 4 ..1800
  code mid 5 1800..1900
  code new 6 1900..
when type=city
field pop 7 3 TRIPLE P1082 enum
  code big 4 Q1
`
