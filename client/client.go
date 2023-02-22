package client

import (
	"context"
	"github.com/semichkin-gopkg/conf"
	"github.com/semichkin-gopkg/temporal/alias"
	"go.temporal.io/sdk/client"
	wf "go.temporal.io/sdk/workflow"
)

type Configuration struct {
	DefaultStartWorkflowUpdaters []alias.StartWorkflowUpdater
	DefaultChildWorkflowUpdaters []alias.ChildWorkflowUpdater
	DefaultActivityUpdater       []alias.ActivityUpdater
}

type ImplClient struct {
	sdkClient     client.Client
	configuration Configuration
}

func NewClient(
	sdkClient client.Client,
	updaters ...conf.Updater[Configuration],
) *ImplClient {
	return &ImplClient{
		sdkClient:     sdkClient,
		configuration: conf.NewBuilder[Configuration]().Append(updaters...).Build(),
	}
}

func (c *ImplClient) RunWorkflow(
	ctx context.Context,
	workflow alias.Workflow,
	params alias.WorkflowParams,
	updaters ...alias.StartWorkflowUpdater,
) (client.WorkflowRun, alias.TemporalServiceError) {
	configuration := conf.NewBuilder[client.StartWorkflowOptions]().
		Append(c.configuration.DefaultStartWorkflowUpdaters...).
		Append(updaters...).
		Build()

	return c.sdkClient.ExecuteWorkflow(ctx, configuration, workflow, params)
}

func (c *ImplClient) RunWorkflowWithSignal(
	ctx context.Context,
	workflowID alias.WorkflowID,
	signalName alias.WorkflowSignalName,
	signalArgs alias.WorkflowSignalPayload,
	workflow alias.Workflow,
	params alias.WorkflowParams,
	updaters ...alias.StartWorkflowUpdater,
) (client.WorkflowRun, alias.TemporalServiceError) {
	configuration := conf.NewBuilder[client.StartWorkflowOptions]().
		Append(c.configuration.DefaultStartWorkflowUpdaters...).
		Append(updaters...).
		Build()

	return c.sdkClient.SignalWithStartWorkflow(
		ctx,
		workflowID,
		signalName,
		signalArgs,
		configuration,
		workflow,
		params,
	)
}

func (c *ImplClient) ExecuteWorkflow(
	ctx context.Context,
	workflow alias.Workflow,
	params alias.WorkflowParams,
	resultPtr alias.ExecutionResultPtr,
	updaters ...alias.StartWorkflowUpdater,
) alias.TemporalExecutionError {
	run, err := c.RunWorkflow(ctx, workflow, params, updaters...)
	if err != nil {
		return err
	}

	return run.Get(ctx, resultPtr)
}

func (c *ImplClient) RunWorkflowAsChild(
	ctx wf.Context,
	workflow alias.Workflow,
	params alias.WorkflowParams,
	updaters ...alias.ChildWorkflowUpdater,
) wf.ChildWorkflowFuture {
	configuration := conf.NewBuilder[wf.ChildWorkflowOptions]().
		Append(c.configuration.DefaultChildWorkflowUpdaters...).
		Append(updaters...).
		Build()

	ctx = wf.WithChildOptions(ctx, configuration)
	return wf.ExecuteChildWorkflow(ctx, workflow, params)
}

func (c *ImplClient) RunWorkflowAsChildWithWaitExecutionStart(
	ctx wf.Context,
	workflow alias.Workflow,
	params alias.WorkflowParams,
	updaters ...alias.ChildWorkflowUpdater,
) (wf.Execution, error) {
	future := c.RunWorkflowAsChild(ctx, workflow, params, updaters...)
	var execution wf.Execution
	return execution, future.GetChildWorkflowExecution().Get(ctx, &execution)
}

func (c *ImplClient) ExecuteWorkflowAsChild(
	ctx wf.Context,
	workflow alias.Workflow,
	params alias.WorkflowParams,
	resultPtr alias.ExecutionResultPtr,
	updaters ...alias.ChildWorkflowUpdater,
) alias.TemporalExecutionError {
	return c.RunWorkflowAsChild(ctx, workflow, params, updaters...).Get(ctx, resultPtr)
}

func (c *ImplClient) RunActivity(
	ctx wf.Context,
	activity alias.Activity,
	params alias.ActivityParams,
	updaters ...alias.ActivityUpdater,
) wf.Future {
	configuration := conf.NewBuilder[wf.ActivityOptions]().
		Append(c.configuration.DefaultActivityUpdater...).
		Append(updaters...).
		Build()

	ctx = wf.WithActivityOptions(ctx, configuration)
	return wf.ExecuteActivity(ctx, activity, params)
}

func (c *ImplClient) ExecuteActivity(
	ctx wf.Context,
	activity alias.Activity,
	params alias.ActivityParams,
	resultPtr alias.ExecutionResultPtr,
	updaters ...alias.ActivityUpdater,
) alias.TemporalExecutionError {
	return c.RunActivity(ctx, activity, params, updaters...).Get(ctx, resultPtr)
}

func (c *ImplClient) SignalWorkflow(
	workflowID alias.WorkflowID,
	runID alias.WorkflowRunID,
	name alias.WorkflowSignalName,
	payload alias.WorkflowSignalPayload,
) alias.TemporalServiceError {
	return c.sdkClient.SignalWorkflow(
		context.Background(),
		workflowID,
		runID,
		name,
		payload,
	)
}

func (c *ImplClient) TerminateWorkflow(
	workflowID alias.WorkflowID,
	runID alias.WorkflowRunID,
) error {
	return c.sdkClient.TerminateWorkflow(
		context.Background(),
		workflowID,
		runID,
		"",
	)
}
