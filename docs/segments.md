# Segment format (`.sseg`, version 1)

A segment holds the rows of one bundle of one section. It is written once,
through `<name>.tmp`, and never changed. Integers are big-endian.

| Part | Size |
|---|---:|
| rows, partition by partition | `Rows · RowWidth` |
| partition table: key, first row u64, row count u64 | `Parts · (PartWidth+16)` |
| footer | 128 |

The file size must equal the sum exactly. Rows ascend as bytes inside a
partition; partition keys ascend strictly. A section without partitions has one
entry with an empty key. While writing, the table goes to `<name>.tbl`, so a
high-cardinality partition key costs no memory.

## Footer

| Offset | Size | Field |
|---:|---:|---|
| 0 | 8 | magic `SILKSEG2` |
| 8 | 2 | version = 1 |
| 10 | 2 | `PartWidth` |
| 12 | 2 | `RowWidth` |
| 14 | 2 | `DocOff`, where the document number lies in a row |
| 16 | 8 | `Parts` |
| 24 | 8 | `Rows` |
| 32 | 32 | SHA-256 of the declaration |
| 64 | 32 | SHA-256 of rows and table |
| 96 | 24 | reserved, 0 |
| 120 | 8 | magic `SILKSEG2` |

## Rows by kind

| Kind | Row | `DocOff` |
|---|---|---:|
| index | key slots, document u32, offset u32 | sum of slot widths |
| silk, value order | bits, document u32 | bits/8 |
| silk, document order | document u32, bits | 0 |
| text | text 8, document u32, QID u32 | 8 |
| composite | text 8, meaning 4, QID u32, document u32 | 16 |
| paths (internal) | Issuer u32, document ID u64, version u32, document u32 | 16 |

The document number is a row of the document table; the offset is the record's
byte offset in the *unpacked* document.

The QID word holds the Subject's Reference codec (WIKIDATA-QID 4, WIKIDATA-PID
5) in its top 4 bits and the ID in the low 28. Anything else, including an ID
that does not fit, is 0 and the Subject is read from the document. The grammar
has no Lexeme Reference, so Lexemes are 0.

Rows without an offset are made distinct per document, so byte-equal rows of
one document appear once.

## Key slots

| Field | Slot |
|---|---|
| record | first W bytes, zero padded |
| raw, n words | 2n bytes |
| rest of a record | first W-1 bytes, zero padded, last byte 00 whole or FF cut |
| missing | all zero |

Text keys of text and composite rows are the plain first 8 bytes, zero padded.
The grammar allows U+0000 in text, so a text of 8 bytes or fewer is not known
whole from its key: rows of these sections are always candidates, and the
answer comes from the document.

## Lookup

`File.Part` is a binary search of the table; `Region.Prefix` two binary searches
of the rows. Readers map the file; nothing is loaded.

## Writing and merging

`Writer` streams rows in order. `Buffer` holds rows in memory and sorts them
(an index of offsets, 4 bytes a row). `Merge` is a k-way merge that drops rows
of dead documents and calls a pause hook every 2^20 rows.
