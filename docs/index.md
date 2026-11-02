# Index directory (version 1)

`ARCHIVE/.silk/`, or `-index DIR`:

| Path | Contents |
|---|---|
| `CURRENT` | the current generation |
| `manifest/<gen>.json` | one manifest per generation |
| `decl/<sha256>.json` | a declaration, named by its hash |
| `grammar/<sha256>.json` | the layouts, names and languages fixed at the first `create` |
| `seg/<n>.sseg` | segments (segments.md) |
| `doc/<n>.sdoc` | document table files |
| `live/<n>.sbit` | the live bitmap |
| `pin/<gen>-<pid>-<time>` | generations in use by readers |
| `LOCK` | process ID of the one writer |
| `tmp/` | runs in progress |

Every change is one new manifest and a rename of `CURRENT`. The collector
removes what neither the current nor a pinned generation names, so a reader
keeps its files until it closes.

## Document table

Each number is one *step* of a pack inventory, in replay order. 40-byte
entries, big-endian: pack u32, Issuer u32, offset u64, stored length u64,
document ID u64, version u32, encoding u8, 3 reserved. Encoding 0xFF marks a
deletion step, which is never live. The tail is magic `SILKDOCS`, first u32,
count u32, 8 reserved, SHA-256 of the entries. Packs are stored by hash, and
the object number is found in the open generation, so a location stays valid
as a set grows.

The live bitmap has bit `i%8` of byte `i/8` set when number `i` is live, with
the same tail (`SILKLIVE`).

## Build and update

`index update` advances an index; `index build` first forgets everything built
and keeps the declarations. Then:

1. **Replay.** For each set, each data pack not yet seen, in order: every
   version step gets the next number, a table entry and a row in the internal
   *paths* section (identity, number). No path map is held in memory, and no
   document is read. Paths rows spill into sorted runs like any section.
2. **Kill.** Walking the new rows in identity order, every number but the last
   of a path dies, and so does every number older segments hold for that path.
   So a rewritten path (no deletion entry, only a later step) and a deletion
   both remove the old rows. A pack that left the set kills its documents;
   whatever was rewritten comes back through the new pack.
3. **Rows.** Workers read documents one at a time (`Reader.Record`, raw or
   zstd) and hand rows to per-section buffers. New documents go to every
   section; a section declared since gets all live documents. When the buffers
   reach half of the limit, each becomes a sorted run.
4. **Merge and place.** Runs merge at most `limit/1 MiB` at a time (2 to 256),
   level by level. Rows sort as bytes, so the result does not depend on the
   limit or the thread count: the output is byte for byte the same.

`-mem` is required; the run also sets Go's soft memory limit to it. The largest
unpacked document is recorded (`index stat`). A document is read whole, so the
limit must exceed it; gogeul's `Reader.Record` has no streaming form.

## Append array and compaction

Policy fields (`-append-rows`, `-runs-per-tier`, `-stale-ratio`):

| Field | Default | Meaning |
|---|---:|---|
| `AppendRows` | 10,000 | bound of the append array, in rows |
| `RunsPerTier` | 8 | runs of one size tier (⌊log₈ rows⌋) before they merge |
| `StaleRatio` | 0.30 | dead row share that forces a major merge |

An update's rows join the append array while it stays under `AppendRows`;
reaching it merges the array and the new rows into a small sorted run at write
time. Append segments are sorted too, so a query binary-searches them rather
than scanning.

`index compact` merges, per section and in rounds: everything when the section
is past `StaleRatio` or with `-full`, else the lowest tier holding `RunsPerTier`
runs. Merges drop dead rows. The dead share is estimated from up to 4,096 evenly
spaced probes of each segment, so no per-section count has to be maintained.

`-idle-cpu` and `-idle-mem` are the free shares the user wants kept. They have
no default. The merge waits at pause points until both hold.

## Verification

`index verify` checks every hash and row order, builds a fresh index in `tmp/`,
and compares live documents and each section's live rows with document numbers
read as paths.
