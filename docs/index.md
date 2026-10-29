# Index directory (version 2)

`silk index build ARCHIVE` writes `ARCHIVE/.silk/` (or `-index DIR`):

| Path | Contents |
|---|---|
| `CURRENT` | the current generation number |
| `manifest/<gen>.json` | one manifest per generation (16 hex digits) |
| `spec/<sha256>.json` | the index specification by hash (docs/spec.md) |
| `seg/<n>.sseg` | segments (docs/segments.md) |
| `doc/<n>.sdoc` | document table files |
| `live/<n>.sbit` | live-version bitmap of a generation |
| `pin/<gen>-<pid>-<time>` | generations in use by readers |
| `LOCK` | process ID of the one writer (build, update, compact) |
| `compact.json` | checkpoints of paused merges |

File numbers `<n>` come from the manifest's counter and are never reused.

## Manifest

```json
{
  "version": 2, "generation": 7, "spec": "<sha256>",
  "sets": [{"name": "0x00010000", "generation": 3, "manifest": "<sha256>"}],
  "packs": [{"set": "0x00010000", "hash": "<sha256>"}],
  "issuers": ["0x00010000"],
  "docs": [{"file": "doc/…01.sdoc", "first": 0, "count": 16696}],
  "live": "live/…09.sbit",
  "sections": [{"name": "TRIPLE", "tag": 4106, "hash": "<tag sha256>",
                "runs": [{"file": "seg/…", "rows": 2242955, "bytes": 32643686}],
                "appends": [], "dead": 0}],
  "next_file": 9
}
```

- `sets` are the archive set generations the generation covers; a reader on
  an archive that moved on gets "run silk index update".
- `packs` and `issuers` are append-only lists the document table numbers
  into.
- A section's `dead` counts rows (in its runs and appended segments) whose
  document version is no longer live. Merges drop them.

## Document table (`.sdoc`)

Header (56 bytes, big-endian): magic `SILKDOCS`, version u16 = 1, reserved
u16, first sequence number u32, entry count u32, pair count u32, SHA-256 of
the body. Body: 40-byte entries, then 6-byte pairs.

| Entry offset | Size | Field |
|---:|---:|---|
| 0 | 4 | pack number |
| 4 | 4 | Issuer number |
| 8 | 8 | document offset in the data pack |
| 16 | 4 | document length |
| 20 | 4 | version ID |
| 24 | 8 | document ID |
| 32 | 4 | first pair |
| 36 | 2 | pair count |
| 38 | 2 | reserved, 0 |

A pair is (TAG WORD u16, record count u32) for every Kind·Codec present in the
document, whether indexed or not: when the version dies, each section adds
the records of its kind to `dead`, and a section added later by a spec change
can still count.

Each update appends one file; sequence numbers continue. A sequence number
identifies a document *version at a place* (set, pack, offset, length): a
version rewritten into a new pack generation gets a new sequence number.

## Live bitmap (`.sbit`)

Header (56 bytes): magic `SILKLIVE`, version u16 = 1, reserved u16, bit count
u32, reserved 8 bytes, SHA-256 of the bits. Bit `i` (byte `i/8`, bit `i%8`)
is set when sequence `i` is live. Each generation that changes liveness writes
a new file.

## Generations, pins and collection

A writer takes `LOCK` (created exclusively; a lock whose process is gone is
taken over), writes new immutable files, writes `manifest/<gen+1>.json` and
then replaces `CURRENT` (temporary file, sync, rename) — the atomic switch.
After a commit it removes every file that neither the current generation, nor
a pinned generation, nor a paused merge (`compact.json`) needs.

A reader pins the generation it opens: it reads `CURRENT`, creates
`pin/<gen>-<pid>-<time>`, and reads `CURRENT` again; if it changed, it drops
the pin and retries. A writer that commits and then collects therefore either
sees the pin or the reader sees the new generation. Closing the index removes
the pin; pins of processes that no longer exist are removed by the collector.
Old segments thus stay readable until every query that started on their
generation has finished.

## Build

Every live version document of the archive gets a sequence number in archive
order and is checked with gogeul (`core.SplitDocument`). Its top-level records
are framed by TAG WORDs; each record of an indexed kind becomes a row. Rows are
kept per tag up to a batch size (4 Mi rows) and written as sorted runs; a tag
with several runs is merged into one.

## Update (incremental)

`silk index update` (alias `add`) compares the archive's live versions with
the live document table by identity (Issuer, document, version) and place
(set, pack, offset, length):

- unchanged versions are not read;
- new or moved versions are read and their rows written as **appended
  segments** (one per section), with new sequence numbers;
- versions no longer in the archive (or moved) are cleared in the bitmap, and
  their record counts added to each section's `dead`.

Pack sets only ever add objects, so a generation's data packs stay readable;
only the documents of new objects are new. A tag whose spec hash changed is
rebuilt from every live version (sorted runs). With `-compact`, the due merges
run afterwards (docs/compaction.md).

## Queries

A condition fixes fields of one Edge kind (`-triple P=O` is TRIPLE fields 1
and 3; `-match KIND:PATH=VALUE`; `-label LANG` is LABEL `2.0`). The planner
uses the field at the tag's partition path as the partition key and the one
at its sort path as the sort key; other fields are left to the exact check.
Each segment of the section is searched (binary search in runs, scan in
appended segments), rows of dead versions are skipped, each candidate record
is read from its data pack, decoded with gogeul and compared, and the Edge Id
taken. Several conditions are joined by document with a merge join on
sequence numbers.

## Verification

`silk index verify` opens the current generation against the archive and
checks: every segment's body hash, row order and footer (TAG, widths, spec
hash, sorted exactly for runs); that the live document table equals the
archive's live versions; that each section's live rows equal the rows rebuilt
from the documents; and that each section's dead count equals its rows of
dead versions.

## Limits (to do)

- Build and update hold one batch of rows per tag in memory, and the
  other-partition rows of a merge stay in memory; both need spilling for
  archives far larger than memory.
- The document table is read whole at update and query time, and the bitmap
  is rewritten each generation; dead document entries are never reclaimed.
- Document parsing in build and update is single-threaded.
