package worker

import (
	wk "go.temporal.io/sdk/worker"
)

type ImplWorker struct {
	sdkWorker  wk.Worker
	registries []ExecutionRegistry
}

func NewWorker(
	sdkWorker wk.Worker,
	registries ...ExecutionRegistry,
) *ImplWorker {
	return &ImplWorker{
		sdkWorker:  sdkWorker,
		registries: registries,
	}
}

func (w *ImplWorker) Run() error {
	for _, registry := range w.registries {
		if reg, ok := registry.(WorkflowsRegistry); ok {
			for _, workflow := range reg.GetWorkflows() {
				if workflow != nil {
					w.sdkWorker.RegisterWorkflow(workflow)
				}
			}
		}

		if reg, ok := registry.(ActivitiesRegistry); ok {
			for _, activity := range reg.GetActivities() {
				if activity != nil {
					w.sdkWorker.RegisterActivity(activity)
				}
			}
		}
	}

	return w.sdkWorker.Run(wk.InterruptCh())
}
