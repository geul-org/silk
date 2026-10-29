# silk

SILK is a query engine for GEUL archives. It answers questions such as "which
documents state that someone's occupation (P106) is physicist (Q169470)?"
over large sets of GEUL documents from a partitioned, sorted index of their
records — and, later, entity signatures — instead of decoding every document.

Status: **early development.** Partitioned indexes, incremental updates,
size-tiered compaction and queries on any indexed Edge field work; entity
signatures exist as a library but are not yet built into indexes. Formats may
change before a first release.

## Install

```sh
go install github.com/park-jun-woo/silk/cmd/silk@latest
```

SILK reads GEUL pack archives through
[gogeul](https://github.com/geul-org/gogeul). The archive must store records
raw (uncompressed), because the index points into data pack bytes:

```sh
GEULCOMPRESS=none geul pack build -o ARCHIVE SRC   # SRC: a directory archive
```

A zstd-compressed archive is rejected with an error asking for an
uncompressed pack. zstd is for distribution and storage; the search tier uses
raw packs.

## Use

```sh
silk index build ARCHIVE                     # writes ARCHIVE/.silk/
silk index update ARCHIVE                    # after new pack generations
silk index compact ARCHIVE                   # merges that are due (-full, -idle)
silk index verify ARCHIVE
silk index stat ARCHIVE
silk index spec > spec.json                  # the default index specification
silk query -triple P106=Q169470 ARCHIVE      # documents with Edge Ids
silk query -count -triple P31=Q5 -triple P106=Q169470 ARCHIVE
silk query -label ko -match LABEL:1=Q42 ARCHIVE
silk query -explain -triple P106=Q169470 ARCHIVE
```

`-triple PRED=OBJ` takes tokens `Qn`, `Pn`, `#n` (LOCAL) or `*` (any).
`-match KIND:PATH=VALUE[,PATH=VALUE]` fixes fields of any indexed Edge kind
(`1` is the first field after the Edge Id, `2.0` the language of a
language-tagged text) with tokens `Qn`, `Pn`, `#n`, `@lang` or `0x<record
hex>`; `-label LANG` is `-match LABEL:2.0=@LANG`. All conditions must hold in
one document. Results are printed as `<Issuer>/<document>/<version>.geul
#<Edge Id>,…`. `-threads N` (default 8) bounds worker goroutines, leaving a
desktop usable; `-index DIR` puts the index elsewhere.

On 16,696 Wikidata documents (271 MB, 7.07 million indexed records) a build
takes about 3 s and writes 113 MB. `-triple P106=Q169470` answers in about
0.13 ms, `-triple P31=Q5` (1,579 documents) in about 1 ms, where the earlier
full scan of the TRIPLE records took about 20 ms. An update with 4,217
changed documents takes 1.8 s; query results after updates and compactions
equal those of a fresh build.

## How it works

GEUL records start with a 16-bit TAG WORD that gives the record's kind,
codec and length; the fields of an Edge are themselves records. SILK frames
records and their fields from TAG WORDs alone, without a codebook.

- **Index specification.** For each TAG kind the spec declares a partition
  key and a sort key — record fields at fixed positions (TRIPLE: Predicate
  and Object; LABEL: language and Subject) — and a minimum partition size.
  Positions that need the codebook, and unbounded fields as partition keys,
  are rejected. Nothing is hard-coded: [docs/spec.md](docs/spec.md).
- **Segments.** Each kind's entries live in segments: rows of (sort key
  slot, document, offset) grouped by partition, small partitions together in
  an "other" partition. Sorted runs are binary-searched; appended segments
  from updates are scanned. [docs/segments.md](docs/segments.md).
- **Generations.** A manifest per generation lists the segments, the
  document table and a live-version bitmap; switching `CURRENT` is atomic,
  readers pin their generation, and old files go when no one reads them.
  Updates read only new or moved document versions and mark replaced ones
  dead. [docs/index.md](docs/index.md).
- **Compaction.** Appended segments become small runs, runs of similar size
  merge in O(n), and the base run joins when there are too many runs or too
  many dead rows; merges pause and resume at partition boundaries and can
  wait for an idle machine. [docs/compaction.md](docs/compaction.md).
- **Exact check.** Every candidate is decoded with gogeul and compared field
  by field, and its Edge Id read; conditions are joined by document.
- **Documents are not reordered.** The index is a sorted view; the GEUL
  bytes, their content hashes and signatures stay as published.
- **Entity signatures (next).** A fixed-width bit string per entity whose
  fields the user declares, as an approximate first filter.

## Layout

| Path | Contents |
|---|---|
| `pkg/archive` | Open a pack archive: sets, data packs, version documents per pack (raw only) |
| `pkg/spec` | Index specification: tags, keys, known Edge layouts, validation, key extraction by TAG WORD framing |
| `pkg/seg` | Segment files: writer (with checkpoints), reader, binary search, scan, k-way merge |
| `pkg/index` | Index directory: manifest generations, document table, live bitmap, build, update, compaction, pins and collection, verification |
| `pkg/idle` | Idle test from load average and available memory (30% left free) |
| `pkg/sig` | Entity signature schema, encoder, mask/pattern scan (enum and one-hot fields so far) |
| `pkg/query` | Conditions → plan (partition and sort keys, exact check, document join) → results |
| `cmd/silk`, `internal/cli` | Command line |

## License

silk is released under the [MIT License](LICENSE). Third-party components keep
their licenses; see [NOTICE](NOTICE).
