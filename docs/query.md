# Queries

## Conditions

```
Query = Cond+            all in one document
Cond  = TAG | section | silk, then predicates
```

Words are read in order. One without `=` opens a condition. `path=value` adds
an equality, `path^=value` a prefix on a text key (`field` paths: `predicate`,
`object.language`, `object.bytes`, or numbers `1`, `2.1`, with 0 the Edge Id).

Values are gogeul notation: `Qn`, `Pn`, `#n`, `"text"`, `"text"@tag`,
`"text"@@tag` (SHORT-CODE only), `@tag`, `0x…`, and decimals for raw fields.
`@tag` stands for both language forms, since the grammar calls them the same
language. Those become one predicate with two byte strings.

Compared by fields, not by record bytes: a field's byte position varies with
Length extensions and Edge Id width.

## Planning

For a TAG condition, the section with the same TAGs and most fixed fields wins
(partition plus the leading keys in order), then fewer rows. A partition not
fixed means every partition. No section is refused with a `create index`
suggestion, unless `-scan`. A section named as the condition, or by
`-index-name`, is used as is.

A condition on an index is **covered** when every predicate lies in the key and
fits its slot whole; the rows then answer alone.

## Running

1. The first condition selects rows: partitions, then a prefix or a silk
   pattern, in each run and append segment. Dead documents are skipped.
2. A composite section's meaning conditions filter in the same range.
3. A later row condition joins by document. A later silk condition reads the
   document-order bundle of each candidate document.
4. What rows cannot decide goes to the documents, each opened once, per
   condition: text and meaning from the statements, `other`/`somevalue` rows.

Offsets are into the unpacked document; zstd documents unpack whole through
`Reader.Record`. A byte-bounded LRU (`-doc-cache`, default none) keeps them
across a batch.

## Results

| Option | Prints | Documents opened |
|---|---|---|
| `-docs` (default) | path, `#Edge Id` | candidates, and for the Edge Id |
| `-records` | path, offset, record hex | same |
| `-keys` | path, row hex, QID | none; refused on an uncovered index |

`-batch` reads one query per input line and prefixes each output line with the
line number. `query.RunBatch` is the same in Go.

`-scan` applies every condition to every live document with no section. It is
the oracle of the tests.
