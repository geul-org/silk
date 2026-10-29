// Package cli is the silk command line: argument parsing and output only.
// Building, updating, compacting and querying the index are pkg/index and
// pkg/query.
package cli

// usageText is printed by `silk help` and on a usage error.
const usageText = `silk — query engine for GEUL pack archives (early development)

Usage:
  silk index build   [-index DIR] [-spec FILE] [-threads N] ARCHIVE
  silk index update  [-index DIR] [-spec FILE] [-compact] ARCHIVE   (also: add)
  silk index compact [-index DIR] [-full] [-idle] [-dry] [-threads N] ARCHIVE
  silk index verify  [-index DIR] ARCHIVE
  silk index stat    [-index DIR] ARCHIVE
  silk index spec    (print the default index specification)
  silk query [-index DIR] [-threads N] [-count] [-explain]
             (-triple P=O | -match KIND:PATH=VALUE[,PATH=VALUE] | -label LANG)... ARCHIVE

ARCHIVE is a data folder holding .well-known/geul/ (or that directory) packed
with records stored raw: GEULCOMPRESS=none geul pack build -o ARCHIVE SRC.
The index defaults to ARCHIVE/.silk. Without -spec, build uses the default
spec and update keeps the index's spec. A -triple condition is PRED=OBJ with
tokens Qn, Pn, #n or * (any); -match names an Edge kind (TRIPLE, LABEL, ...)
and field paths (1, 2.0) with tokens Qn, Pn, #n, @LANG or 0x<record hex>;
-label LANG is -match LABEL:2.0=@LANG. All conditions must hold in one
document. compact -idle waits for an idle machine (load and memory leave 30%
free) and pauses between partitions while it is busy; Ctrl-C saves a
checkpoint that the next compact resumes.
`
