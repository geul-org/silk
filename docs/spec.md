# Index specification (`spec.json`, version 1)

An index specification says, for each TAG kind (Kind·Codec) SILK indexes,
which record field partitions the entries and which field orders them.
Nothing about predicates, languages or Wikidata is hard-coded in SILK: the
default spec below is only a default.

```json
{
  "version": 1,
  "tags": [
    {"name": "TRIPLE", "kind": 2, "codec": 10,
     "partition": {"path": [1], "width": 6},
     "sort": {"path": [3], "width": 6},
     "min_partition": 256}
  ]
}
```

| Field | Meaning |
|---|---|
| `name` | Name of the tag (the Edge layout name by convention) |
| `kind`, `codec` | Kind (2 = Edge) and direct Codec of the records |
| `partition` | Partition key; omitted = one partition for the whole kind |
| `sort` | Sort key inside a partition; omitted = rows ordered by document and offset only |
| `min_partition` | Partitions with fewer entries go to the segment's *other* partition |

## Keys

A key is a record field at a **fixed position** and a slot **width** in bytes
(even, 2–32). `path` is `[field]` — the n-th embedded record of the Edge,
field 0 being the Edge Id — or `[field, 0]`, the first embedded record of a
UTF8-LANGUAGE-TEXT field (its language record). The slot holds the leading
`width` bytes of the field record, zero padded. A record of at most `width`
bytes is kept whole; since its TAG WORD carries its length, equal slots mean
equal records. A longer record keeps a prefix, so a slot match is only a
candidate; every candidate is decoded and compared field by field (exact
check). A record whose field cannot be framed (for example a LABEL whose
Object is a LOCAL Reference, which needs a REFERENCE-NODE to resolve) gets a
zero slot and is never returned for a language condition.

Keys are framed by TAG WORDs alone, without a codebook. Validation therefore
rejects every position that would need one:

- the TAG kind must have a layout SILK knows (below); Issuer-format codecs and
  Kinds other than Edge have none;
- `path[0]` must be one of the layout's fixed record fields (a raw word such as
  a TRIPLE's Rank, or anything after it, is not);
- a second step is allowed only into a UTF8-LANGUAGE-TEXT field, and only `0`;
- a **partition key** must be a bounded field: a Reference (or slot marker) or
  a language record. A field that may hold a Value of any length (a TRIPLE
  Object, a whole text) can be a sort key but not a partition key.

Known layouts (from the grammar pages, `edge/*.md`):

| Codec | Name | Fixed record fields |
|---|---|---|
| `0x00` | REFERENCE-NODE | Edge Id, Subject, Object (any) |
| `0x04` | LABEL | Edge Id, Subject, Object (language text) |
| `0x05` | ALIAS | Edge Id, Subject, Object (language text) |
| `0x06` | DESCRIPTION | Edge Id, Subject, Object (language text) |
| `0x0A` | TRIPLE | Edge Id, Predicate, Subject, Object (any); Rank is a raw word |
| `0x0B` | QUALIFIER | Edge Id, Property, Target, Value (any) |
| `0x20` | SOURCE | Edge Id, Target, Origin |

## Default spec (Wikidata documents)

`silk index spec` prints it.

| Tag | Partition | Sort | Min |
|---|---|---|---:|
| TRIPLE | Predicate `[1]`, 6 B | Object `[3]`, 6 B | 256 |
| QUALIFIER | Property `[1]`, 6 B | Value `[3]`, 6 B | 256 |
| LABEL, ALIAS, DESCRIPTION | language `[2,0]`, 16 B | Subject `[1]`, 6 B | 256 |
| SOURCE | — | Origin `[2]`, 6 B | 0 |
| REFERENCE-NODE | — | Object `[2]`, 16 B | 0 |

Six bytes hold a Reference TAG WORD and an ID of up to two words (Q and P
numbers below 2³²). Sixteen bytes hold short BCP 47 tags whole.

## Hashes

The spec's hash is the SHA-256 of its canonical JSON (`Spec.Encode`). The
manifest names it and the index keeps a copy as `spec/<hash>.json`. Each
section records the hash of its own tag, and each segment footer carries it.
When `silk index update -spec FILE` changes a tag, only that tag's section is
rebuilt (from every live document); other sections keep their segments and
receive appended rows as usual. A tag removed from the spec loses its section.

## TODO

- Declare the spec as a GEUL document (like codec-spec documents) instead of
  JSON, and name it by its document address and content hash.
- Read field layouts from codec-spec documents (DECLARE·FIELD) instead of the
  built-in table, so new Edge codecs can carry keys.
