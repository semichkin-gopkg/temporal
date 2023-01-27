package client

import (
	"context"
	"github.com/semichkin-gopkg/temporal/alias"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/workflow"
)

type Client interface {
	ScheduleWorkflow(
		context.Context,
		alias.Workflow,
		alias.WorkflowParams,
		...alias.StartWorkflowUpdater,
	) (client.WorkflowRun, alias.TemporalServiceError)

	ExecuteWorkflow(
		context.Context,
		alias.Workflow,
		alias.WorkflowParams,
		alias.ExecutionResultPtr,
		...alias.StartWorkflowUpdater,
	) alias.TemporalExecutionError

	ScheduleWorkflowWithSignal(
		context.Context,
		alias.WorkflowID,
		alias.WorkflowSignalName,
		alias.WorkflowSignalPayload,
		alias.Workflow,
		alias.WorkflowParams,
		...alias.StartWorkflowUpdater,
	) (client.WorkflowRun, alias.TemporalServiceError)

	ScheduleChildWorkflow(
		workflow.Context,
		alias.Workflow,
		alias.WorkflowParams,
		...alias.ChildWorkflowUpdater,
	) workflow.ChildWorkflowFuture

	ScheduleChildWorkflowWithWaitExecutionStart(
		workflow.Context,
		alias.Workflow,
		alias.WorkflowParams,
		...alias.ChildWorkflowUpdater,
	) (workflow.Execution, error)

	ExecuteChildWorkflow(
		workflow.Context,
		alias.Workflow,
		alias.WorkflowParams,
		alias.ExecutionResultPtr,
		...alias.ChildWorkflowUpdater,
	) alias.TemporalExecutionError

	ScheduleActivity(
		workflow.Context,
		alias.Activity,
		alias.ActivityParams,
		...alias.ActivityUpdater,
	) workflow.Future

	ExecuteActivity(
		workflow.Context,
		alias.Activity,
		alias.ActivityParams,
		alias.ExecutionResultPtr,
		...alias.ActivityUpdater,
	) alias.TemporalExecutionError

	SignalWorkflow(
		alias.WorkflowID,
		alias.WorkflowRunID,
		alias.WorkflowSignalName,
		alias.WorkflowSignalPayload,
	) alias.TemporalServiceError

	TerminateWorkflow(
		alias.WorkflowID,
		alias.WorkflowRunID,
	) alias.TemporalServiceError
}
