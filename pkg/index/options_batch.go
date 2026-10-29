//ff:func feature=index type=helper control=sequence
//ff:what Options.batch — rows per tag held in memory before a segment is written (DefaultBatch when unset)
package index

// batch resolves o.Batch.
func (o Options) batch() uint64 {
	if o.Batch == 0 {
		return DefaultBatch
	}
	return o.Batch
}
