package selector

import wf "go.temporal.io/sdk/workflow"

type Selector[T any] interface {
	Select(wf.Context) (T, bool)
	HasPending() bool
}
