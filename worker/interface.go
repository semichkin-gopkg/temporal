package worker

import "github.com/semichkin-gopkg/temporal/alias"

type Worker interface {
	Run() error
}

type ExecutionRegistry = any // WorkflowsRegistry or ActivitiesRegistry

type WorkflowsRegistry interface {
	GetWorkflows() []alias.Workflow
}

type ActivitiesRegistry interface {
	GetActivities() []alias.Activity
}
