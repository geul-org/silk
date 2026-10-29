# silk

SILK is a query engine for GEUL archives. It answers questions such as "which
documents state that someone's occupation (P106) is physicist (Q169470)?"
over large sets of GEUL documents by scanning compact sorted record maps — and,
later, entity signatures — instead of decoding every document.

Status: **early development.** The record map, the multi-threaded scan and
TRIPLE queries work; entity signatures exist as a library but are not yet
built into indexes. Formats may change before a first release.

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
silk index verify ARCHIVE
silk query -triple P106=Q169470 ARCHIVE      # documents with Edge Ids
silk query -count -triple P31=Q5 -triple P106=Q169470 ARCHIVE
silk query -explain -triple P106=Q169470 ARCHIVE
```

A condition is `PRED=OBJ` with tokens `Qn`, `Pn`, `#n` (LOCAL) or `*` (any).
Several conditions must all hold in one document. Results are printed as
`<Issuer>/<document>/<version>.geul #<Edge Id>,…`. `-threads N` sets the
scanning goroutines (default 8, chosen to reach memory bandwidth while
leaving a desktop usable); `-index DIR` puts the index elsewhere.

On 16,589 Wikidata documents (271 MB, 7.18 million records) the index is
29 MB and builds in about 1.3 s; `-triple P106=Q169470` scans all 2.24
million TRIPLE records in about 20 ms with 8 threads.

## How it works

GEUL records start with a 16-bit TAG WORD that gives the record's kind,
codec and length. A decoder can count and skip records without a codebook.
SILK builds on that:

- **Sorted record map.** For each data pack and Issuer, SILK stores the byte
  offset (u32) of every top-level record, sorted by (Kind·Codec, offset),
  plus a small key range table and a document table — about 4 bytes per
  record. The TRIPLE records of a pack form one contiguous range, so a TRIPLE
  query reads exactly those offsets. Documents themselves are never
  reordered: their bytes, content hashes and signatures stay as published.
  Format: [docs/recmap.md](docs/recmap.md).
- **Fused scan.** Offsets are split across goroutines; each reads the data
  pack in windows and, in one pass, frames each record from its TAG WORD and
  compares the Predicate and Object record bytes. GEUL's canonical form makes
  equal records equal bytes. Speed comes from the fused loop and threads, not
  the language.
- **Exact check.** Every hit is decoded with gogeul and its Edge Id
  extracted; documents must meet every condition.
- **Entity signatures (next).** A fixed-width bit string per entity (16, 64,
  128… bits) whose fields the user declares — enumeration codes, one-hot and
  Bloom bits, prefix codes, range buckets — as an approximate first filter
  before the record map scan. The schema will be declared as a GEUL document.

Design choices measured on the sample above: sorting records inside documents
by codec saved only 0.7–3% after zstd and broke statement locality, so the
order lives in the map, not the documents. A per-record document-order map
(`tag[]`) is not kept; walking a document follows TAG WORDs.

## Layout

| Path | Contents |
|---|---|
| `pkg/archive` | Open a pack archive: sets, data packs, version documents per pack (raw only) |
| `pkg/recmap` | Record map: build, file format, read with `ReadAt`, verify |
| `pkg/index` | Index directory: build every map, catalog, open against current generations, verify |
| `pkg/scan` | Multi-threaded fused scan over a map range with a record matcher |
| `pkg/sig` | Entity signature schema, encoder, mask/pattern scan (enum and one-hot fields so far) |
| `pkg/query` | Conditions → plan (signature filter, map scan, exact check) → results |
| `cmd/silk`, `internal/cli` | Command line |

## License

The project's own work is dedicated to the public domain under
[CC0 1.0](LICENSE). Third-party components keep their licenses; see
[NOTICE](NOTICE).
