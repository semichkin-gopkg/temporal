package alias

import (
	"github.com/semichkin-gopkg/configurator"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/workflow"
)

type (
	StartWorkflowUpdater = configurator.Updater[client.StartWorkflowOptions]
	ChildWorkflowUpdater = configurator.Updater[workflow.ChildWorkflowOptions]
	ActivityUpdater      = configurator.Updater[workflow.ActivityOptions]

	Execution          = any
	ExecutionResultPtr = any

	Workflow              = Execution
	WorkflowParams        = any
	WorkflowID            = string
	WorkflowRunID         = string
	WorkflowSignalName    = string
	WorkflowSignalPayload = any

	Activity       = Execution
	ActivityParams = any

	TemporalServiceError   = error
	TemporalExecutionError = error
)
