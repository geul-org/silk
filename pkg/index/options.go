//ff:type feature=index type=model
//ff:what Options — a writer's limits: memory limit in bytes (required), worker threads, merge policy, pause hook
package index

type Options struct {
	Limit   int64
	Threads int
	Policy  Policy
	Pause   func() error
}
