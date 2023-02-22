package alias

import (
	"github.com/semichkin-gopkg/conf"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/workflow"
)

type (
	StartWorkflowUpdater = conf.Updater[client.StartWorkflowOptions]
	ChildWorkflowUpdater = conf.Updater[workflow.ChildWorkflowOptions]
	ActivityUpdater      = conf.Updater[workflow.ActivityOptions]

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
