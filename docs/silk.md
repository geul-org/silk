# SILK definition

A definition states every field, bit position, width, code number and value.
Nothing is assigned: no frequencies, no automatic numbering, no table
references. A value the definition does not list is `other`.

```
silk entity 128
field type 0 6 TRIPLE P31 enum
  code person 4 Q5 Q15632617
  code city 5 Q515
field alive 6 1 TRIPLE P570 flag
when type=person
field job 7 8 TRIPLE P106 enum
  code politician 4 Q82955
field area 15 6 TRIPLE P27 prefix
  code asia 0b01
  code korea 0b0100 Q884
field born 21 4 TRIPLE P569 bucket
  code old 4 ..1800
  code mid 5 1800..1900
  code new 6 1900..
```

(Fields here are an example. The first real silk is a user decision.)

| Line | Meaning |
|---|---|
| `silk NAME BITS` | name and row width in bits |
| `field NAME OFFSET WIDTH TRIPLE PREDICATE KIND` | a field; after `when`, part of that block |
| `code NAME NUMBER VALUE…` | enum: one number for one or more values |
| `code NAME 0bBITS [VALUE…]` | prefix: leading bits; without values only a group for queries |
| `code NAME NUMBER LO..HI` | bucket: a half-open range, either end open |
| `when FIELD=CODE` | following fields exist only under that code of a common enum field |

`flag` is 1 bit and has no codes: 1 for any statement but novalue.

The parser refuses: a field past the declared bits, two fields sharing a bit
(common fields with each block), a code that does not fit its width, a code
overlapping the reserved codes, a repeated name, number, value or range, a
field without codes, when blocks on two fields, and any line it does not know.
Bits may be left unused.

## Reserved codes

| Code | Name | Meaning |
|---:|---|---|
| 0 | `none` | no statement |
| 1 | `novalue` | the statement says none |
| 2 | `somevalue` | a value exists, unknown |
| 3 | `other` | outside the definition |

Deprecated statements are left out.

## Rows

Statements group by Subject. Each entity has a row per combination of its codes:
a person with two jobs has two rows, a neighbour in both places. A date spanning
two buckets spreads the same way. Past `MaxExpand` (3) rows for one entity, only
one field spreads, the when field if it already picked several blocks, and the
others hold `other`, which every value condition keeps as a candidate. No row is
written for an entity with no defined statement.

The composite section's meaning half is one row per entity: a multi-valued field
is `other`, nothing spreads.

## Conditions

`field=value`, where value is a code name, a value token, a year for a bucket, or
a reserved name. A row matches on the exact code or on `somevalue`/`other`.
Exact rows answer without a document. Candidates go to the document, where the
statements decide: yes, no, or unknown (`?`).

A field under `when` needs that code; rows whose when field is `somevalue` or
`other` carry no such field and stay candidates, and the document shows them as
unknown when the statement is there.

## Jumps

A pattern is a list of (offset, width, allowed codes). `Pattern.seek` computes
the smallest bit string at or above a row that can match; the scan binary-searches
to it. Fixed leading fields are one jump; a condition on a later field hops
through the combinations of the fields before it.
