// Package cli is the silk command line: argument parsing and output only.
// Building, opening and querying the index are pkg/index and pkg/query.
package cli

// usageText is printed by `silk help` and on a usage error.
const usageText = `silk — query engine for GEUL pack archives (early development)

Usage:
  silk index build  [-index DIR] ARCHIVE
  silk index verify [-index DIR] ARCHIVE
  silk query [-index DIR] [-threads N] [-count] [-explain] -triple P=O [-triple P=O ...] ARCHIVE

ARCHIVE is a data folder holding .well-known/geul/ (or that directory) packed
with records stored raw: GEULCOMPRESS=none geul pack build -o ARCHIVE SRC.
The index defaults to ARCHIVE/.silk. A -triple condition is PRED=OBJ with
tokens Qn, Pn, #n or * (any), e.g. -triple P106=Q169470; several conditions
must all hold in one document.
`
