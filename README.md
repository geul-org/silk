# silk

SILK (Symbolic Index for LLM Knowledge) is a search engine for GEUL archives.
Its core is a sorted array whose keys carry meaning in their bits: search is
stepwise narrowing, a jump to a sorted range at each step, and only the
candidates that remain are read from their documents.

Status: **early development.** Formats may change before a first release.

Four kinds of section share one segment format (fixed-width rows in byte order,
mapped with `mmap`, never compressed):

| Kind | Finds | Row |
|---|---|---|
| index | records | key slots W, document 4 B, offset 4 B |
| silk | meaning | meaning bits, document 4 B (two bundles: value order, document order) |
| text | strings | first 8 text bytes, document 4 B, QID 4 B |
| composite | both | first 8 text bytes, meaning 4 B, QID 4 B, document 4 B |

Archives are GEUL pack sets, raw or zstd, read through
[gogeul](https://github.com/geul-org/gogeul). Layouts, names and languages come
from the archive's own grammar set `0x0000` (or `-grammar DIR`).

## Install

```sh
go install github.com/geul-org/silk/cmd/silkg@latest
```

Linux and other Unix systems (`mmap`).

## Use

`create` declares, `index update` builds what is declared or new, and every
writer needs a memory limit: there is no default.

```sh
silkg create index triple TRIPLE -part predicate -key object -width 6 ARCHIVE
silkg create text label LABEL ARCHIVE
silkg create silk entity -def entity.silk ARCHIVE
silkg create composite label-sem LABEL -def meaning.silk ARCHIVE
silkg list ARCHIVE
silkg drop NAME ARCHIVE

silkg index update -mem 2G ARCHIVE        # new documents, sections declared since
silkg index build -mem 2G ARCHIVE         # every section from scratch
silkg index compact -mem 2G ARCHIVE       # merges that are due
silkg index compact -mem 2G -full ARCHIVE # one run per section
silkg index compact -mem 2G -idle-cpu 0.3 -idle-mem 0.3 ARCHIVE
silkg index verify -mem 2G ARCHIVE        # against a fresh build
silkg index stat ARCHIVE

silkg query TRIPLE predicate=P106 object=Q169470 ARCHIVE
silkg query -keys TRIPLE predicate=P106 object=Q169470 ARCHIVE
silkg query LABEL object.language=@ko 'object.bytes="김치"' ARCHIVE
silkg query LABEL 'object.bytes^="김치"' ARCHIVE
silkg query entity job=politician area=asia ARCHIVE
silkg query label-sem object.language=@en 'object.bytes="Paris"' type=city ARCHIVE
silkg query LABEL 'object.bytes="Paris"' entity type=city ARCHIVE
silkg query -batch ARCHIVE < queries.txt
silkg query -explain -scan ... ARCHIVE
```

A word without `=` opens a condition: a TAG, a declared section, or a silk.
Conditions must hold in one document. Output is the document path, `#Edge Id`,
`~` for a row that is only a candidate, and `?` for a value the statements
leave unknown. `-keys` reads rows only and never opens a document.

`-index DIR` keeps the index elsewhere (default `ARCHIVE/.silk`). A query no
section can answer is refused with the `create index` that would answer it;
`-scan` reads every live document instead.

Documents: [docs/segments.md](docs/segments.md),
[docs/index.md](docs/index.md), [docs/silk.md](docs/silk.md),
[docs/query.md](docs/query.md).

## License

CC0 1.0 for our own work; see LICENSE and NOTICE.
