// Package cli is the silkg command line: argument handling and printing only.
package cli

const (
	exitOK    = 0
	exitError = 1
	exitUsage = 2
	usage     = `usage:
  silkg create index NAME TAG [-part PATH] [-key PATH,…] [-width W,…] [-index DIR] [-grammar DIR] ARCHIVE
  silkg create text NAME TAG [-index DIR] [-grammar DIR] ARCHIVE
  silkg create silk NAME -def FILE [-index DIR] [-grammar DIR] ARCHIVE
  silkg create composite NAME TAG -def FILE [-index DIR] [-grammar DIR] ARCHIVE
  silkg drop NAME [-index DIR] ARCHIVE
  silkg list [-index DIR] ARCHIVE
  silkg index build|update -mem SIZE [-threads N] [-append-rows N] [-index DIR] ARCHIVE
  silkg index compact -mem SIZE [-full] [-runs-per-tier N] [-stale-ratio R] [-idle-cpu R -idle-mem R] [-index DIR] ARCHIVE
  silkg index verify -mem SIZE [-threads N] [-index DIR] ARCHIVE
  silkg index stat [-index DIR] ARCHIVE
  silkg query [-index-name NAME] [-docs|-records|-keys] [-explain] [-scan] [-batch] [-doc-cache SIZE] [-index DIR] COND… ARCHIVE
`
)
