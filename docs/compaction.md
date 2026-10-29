# Compaction

Updates add small appended segments and leave rows of replaced versions in
place. Compaction merges segments by size tier, so a row is rewritten a
logarithmic number of times instead of once per update, and the index is
never sorted again from scratch.

## Policy

Checked per section, in order (the first that applies is the section's task;
values are `Policy` fields):

| Task | Trigger (default) | Inputs → output |
|---|---|---|
| major | segments (runs + appended) > `MaxRuns` (8), or dead rows / rows > `StaleRatio` (30%), or `-full` | every segment of the section → one run without dead rows |
| flush | appended rows / rows > `AppendRatio` (10%) | the appended segments → one small sorted run |
| tier | at least `Fanout` (4) runs, other than the largest (the base), in one size tier ⌊log₄ rows⌋ | those runs → one run |

The base run joins only major merges ("now and then with the base run").
Every merge is O(n): sorted inputs are k-way merged partition by partition;
only appended rows, which are few, are sorted in memory. Dead rows are dropped
by every merge, and the section's `dead` count drops by exactly what was
removed.

## Snapshot safety

Each round runs at most one task per section (up to `-threads` sections at a
time), writes each output as a new immutable segment, and commits one new
generation with all finished outputs in place of their inputs. Readers that
pinned an older generation keep reading its segments; the collector removes
them only when no pin names that generation (docs/index.md).

## Pausing, resuming and idle work

- Between partitions a merge calls the pause hook and checks its context.
  When either says stop, the merge flushes its output, records a checkpoint
  (bytes written, partition table so far, other rows, SHA-256 state, last
  finished key, dead rows dropped) in `compact.json` and returns
  `ErrPaused` (exit code 3). Finished merges of the round are still
  committed.
- The next `silk index compact` resumes each checkpoint whose inputs are
  still exactly in their section: it cuts the temporary file back to the
  checkpoint and continues after the last key. Stale checkpoints lose their
  temporary file.
- `silk index compact` merges what is due now (`-full`: one run per section);
  Ctrl-C stops at a checkpoint.
- `silk index compact -idle` waits for an idle machine before each round and
  pauses between partitions while it is busy. Idle means
  `(1-minute load average + threads) / CPUs <= 0.70` and
  `MemAvailable >= 30% of MemTotal` (from `/proc/loadavg` and `/proc/meminfo`),
  i.e. about 30% of CPU and memory stay free. Without `/proc` it does not
  wait. Under WSL the load average covers only the Linux VM, not Windows
  programs. Run it from cron or a timer for background maintenance.

## Measured (Wikidata Q1–Q10000 and properties, 16,696 documents, 271 MB)

| Step | New GEUL bytes | Index bytes written | Time |
|---|---:|---:|---:|
| build A | 271.5 MB | 113.1 MB | 3.0–3.3 s |
| update A→B (122 versions) | 5.4 MB | 3.0 MB | 0.8 s |
| update B→C (4,217 versions) | 126.8 MB | 56.6 MB | 1.8 s |
| compact (6 major, 1 flush) | — | 107.6 MB | 0.5 s |
| compact -full (ALIAS major) | — | 6.3 MB | 0.06 s |

Query results after each update and each compaction equal those of an index
built from scratch on the same archive generation.
