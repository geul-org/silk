# Segment format (`.sseg`, version 1)

A segment holds index entries (rows) of **one TAG kind**, grouped by partition
key. It is written once — through `<name>.tmp`, synced, renamed — and never
changed. Integers are big-endian (like GEUL); comparing rows as bytes orders
them by (sort slot, document sequence number, record offset).

## Rows

| Region | Row |
|---|---|
| partition rows | sort slot (`SortWidth` B) · document u32 · offset u32 |
| other rows | partition slot (`PartWidth` B) · sort slot · document u32 · offset u32 |

The document is a sequence number of the index's document table (one per
indexed document *version*); the offset is the record's byte offset inside
that document. A partition with at least `min_partition` rows has its own
table entry; smaller ones are kept in the *other* partition, grouped by
partition key in ascending order. A key is either in the table or in the
other partition, never both.

A **sorted run** (flag bit 0 set) has every partition's rows, and the other
rows, in ascending byte order: a lookup is a binary search with one `ReadAt`
per probe. An **appended segment** keeps rows in arrival order inside each
partition (partitions and other-partition keys are still in key order): a
lookup scans the partition.

## Layout

| Part | Size |
|---|---:|
| partition rows, partition by partition in key order | `Rows · (SortWidth+8)` |
| other rows | `Other · (PartWidth+SortWidth+8)` |
| partition table: key slot, start u64, count u64 | `Parts · (PartWidth+16)` |
| footer | 128 |

The file size must equal the sum exactly. Table entries ascend strictly by
key; their row ranges start at 0, are contiguous, non-empty and cover all
partition rows.

### Footer (128 bytes)

| Offset | Size | Field |
|---:|---:|---|
| 0 | 8 | magic `SILKSEG1` |
| 8 | 2 | version = 1 |
| 10 | 2 | flags: bit 0 = sorted run |
| 12 | 2 | TAG WORD with only Kind and Codec bits (`tag & 0x187F`) |
| 14 | 2 | `PartWidth` (0 = one partition with the empty key) |
| 16 | 2 | `SortWidth` (0 = no sort key) |
| 18 | 2 | reserved, 0 |
| 20 | 4 | `Parts`, table entries |
| 24 | 8 | `Rows`, partition rows |
| 32 | 8 | `Other`, other rows |
| 40 | 32 | SHA-256 of the tag's spec (`Tag.Hash`) |
| 72 | 32 | SHA-256 of every byte before the footer |
| 104 | 16 | reserved, 0 |
| 120 | 8 | magic `SILKSEG1` |

## Lookup

`Select(part, sort)`:

- with a partition key: its table partition, or else its slice of the other
  partition (prefix = partition slot ‖ sort slot);
- without: every table partition, then the other partition compared at the
  sort slot's offset;
- with a sort key, sorted runs binary-search the prefix; appended segments
  compare each row.

Rows whose document version is not live are skipped by the caller (the live
bitmap, docs/index.md).

## Writing and merging

The writer takes partitions in key order (`Begin`, `Add`, `End`). It buffers a
partition until it reaches `min_partition` rows, then streams it; a partition
that ends smaller goes to the other partition. Merging (`Merge`) walks the
groups of all inputs in key order — each input's table partitions and its
other-partition keys — and, per key, streams a k-way merge of sorted inputs
or gathers and sorts in memory when an input is an appended segment; rows of
dead versions are dropped. The writer can be checkpointed between partitions
(`Checkpoint`/`Suspend`: bytes written, table so far, other rows, SHA-256
state) and resumed in another process (`ResumeWriter`, which cuts the
temporary file back to the checkpoint).
