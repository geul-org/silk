//ff:type feature=index type=model
//ff:what taskResult — a finished or paused merge: its task, output run, dead rows dropped, bytes written, checkpoint when paused
package index

// taskResult is what mergeTask returns.
type taskResult struct {
	Task       Task
	Out        SegRef
	Dropped    uint64
	Written    int64
	Checkpoint *checkpoint
}
