# Record map format (`.srm`, version 1)

A record map indexes the version documents of **one Issuer** stored in **one
data pack** (`o/<sha256>.gpk`) of a GEUL pack set. For every top-level record
of those documents it keeps the record's byte offset in the data pack, sorted
by the record's key and then by offset.

- **Key:** the record's TAG WORD with only the Kind and Codec bits kept
  (`tag & 0x187F`, gogeul `core.TargetWord(kind, codec)`). Sorting by key is
  sorting by (Kind, Codec). All TRIPLE records of a pack are one contiguous
  range (key `0x100A`).
- **Documents are not reordered.** The map is a sorted view; the GEUL bytes,
  their content hashes and signatures stay as published. Scanning documents
  in their own order needs no map: TAG WORDs give every record's length.
- **Deterministic:** the same pack set generation always gives the same map
  bytes. Maps are built only from records stored raw (uncompressed pack sets,
  `GEULCOMPRESS=none`); offsets point into the data pack bytes.

## Layout

All integers are little-endian (like the pack format). Hashes are raw 32-byte
SHA-256 values. Issuer IDs are GEUL words, big-endian.

| Offset | Size | Field |
|---:|---:|---|
| 0 | 8 | magic `SILKRMAP` |
| 8 | 2 | format version = 1 |
| 10 | 2 | Issuer length in bytes: 2, 4, 8 or 32 (1, 2, 4 or 16 words) |
| 12 | 4 | key range count `K` |
| 16 | 8 | record count `N` |
| 24 | 8 | document count `D` |
| 32 | 8 | pack set generation |
| 40 | 8 | data pack size in bytes (≤ 2³²−1) |
| 48 | 32 | data pack SHA-256 (its object name) |
| 80 | 32 | manifest SHA-256 of the generation |
| 112 | 32 | Issuer ID, zero padded |
| 144 | 32 | SHA-256 of the body (every byte after offset 176) |
| 176 | 24·K | key table |
| … | 24·D | document table |
| … | 4·N | `off[]` |

The file size must equal `176 + 24·K + 24·D + 4·N` exactly.

### Key table entry (24 bytes)

| Offset | Size | Field |
|---:|---:|---|
| 0 | 2 | key |
| 2 | 6 | reserved, 0 |
| 8 | 8 | start index in `off[]` |
| 16 | 8 | count (≥ 1) |

Keys strictly ascend. Ranges are contiguous from 0 and cover all `N` entries.

### Document table entry (24 bytes)

| Offset | Size | Field |
|---:|---:|---|
| 0 | 4 | document offset in the data pack |
| 4 | 4 | document length in bytes |
| 8 | 8 | document ID |
| 16 | 4 | version ID |
| 20 | 4 | reserved, 0 |

Entries ascend by offset and do not overlap. The document holding a record
offset is found by binary search.

### `off[]`

`N` u32 offsets of records in the data pack, sorted by (key, offset).

## Index directory

An index directory holds `<Issuer>/<data pack sha256>.srm` for every Issuer
and data pack of an archive, and `index.json`:

```json
{
  "version": 1,
  "sets": [{"name": "0x00010000", "generation": 1, "manifest": "<sha256>"}],
  "maps": ["0x00010000/<sha256>.srm"]
}
```

`sets` must equal the archive's current sets (name, generation, manifest, in
order) or the index is stale and must be rebuilt. `index.json` is written
after every map; map files no longer listed are removed.

## Verification

`silk index verify` checks that `index.json` matches the archive, that every
map exists and names the expected Issuer, generation, manifest and pack, that
its body hash matches, and that rebuilding the map from the pack gives the
same key table, document table and `off[]`.
