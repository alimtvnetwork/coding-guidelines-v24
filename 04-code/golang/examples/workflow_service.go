package examples

import (
	"context"

	"coding-guidelines/common/pkg/apperror"
	"coding-guidelines/common/pkg/logger"
)

// WorkflowResult represents the final combined status of the operation.
type WorkflowResult struct {
	PluginSummary PluginSummary            `json:"plugin"`
	RemoteData    RemoteActivationResponse `json:"remoteData"`
}

// PluginWorkflowService orchestrates multi-step business logic across boundaries.
type PluginWorkflowService struct {
	repo   *PluginRepository
	client *WordPressClient
	log    *logger.Logger
}

// NewPluginWorkflowService constructs a workflow service.
func NewPluginWorkflowService(
	repo *PluginRepository,
	client *WordPressClient,
	log *logger.Logger,
) *PluginWorkflowService {
	return &PluginWorkflowService{
		repo:   repo,
		client: client,
		log:    log,
	}
}

// ActivateWorkflow executes end-to-end activation with zero redundant error re-wrapping.
func (s *PluginWorkflowService) ActivateWorkflow(
	ctx context.Context,
	siteId int64,
	pluginId int64,
) apperror.Result[WorkflowResult] {
	// Step 1: Query Database
	pluginRes := s.repo.FindById(ctx, pluginId)
	if pluginRes.HasError() {
		// Zero re-wrapping: propagate the existing AppError directly
		s.log.LogError(pluginRes.AppError())

		return apperror.Fail[WorkflowResult](pluginRes.AppError())
	}

	plugin := pluginRes.Value()

	// Step 2: Call Remote Delegated Endpoint
	remoteRes := s.client.ActivateRemotePlugin(ctx, siteId, plugin.Slug)
	if remoteRes.HasError() {
		// Propagate existing AppError enriched with site context
		s.log.LogError(remoteRes.AppError())

		return apperror.Fail[WorkflowResult](remoteRes.AppError().WithSiteId(siteId))
	}

	// Step 3: Return Combined Success
	s.log.Info("Workflow completed successfully for plugin: " + plugin.Slug)

	return apperror.Ok(WorkflowResult{
		PluginSummary: plugin,
		RemoteData:    remoteRes.Value(),
	})
}
