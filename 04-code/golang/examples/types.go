package examples

import (
	"coding-guidelines/common/pkg/appfault"
	"coding-guidelines/common/pkg/result"
)

type (
	// PluginSummary represents an active plugin record in the system.
	PluginSummary struct {
		Id       int64  `json:"id"`
		Slug     string `json:"slug"`
		Name     string `json:"name"`
		IsActive bool   `json:"isActive"`
	}

	// RemoteActivationResponse contains response data from downstream WordPress REST API.
	RemoteActivationResponse struct {
		IsSuccess bool   `json:"isSuccess"`
		Message   string `json:"message"`
		Version   string `json:"version"`
	}

	// WorkflowResult represents the final combined status of the operation.
	WorkflowResult struct {
		PluginSummary PluginSummary            `json:"plugin"`
		RemoteData    RemoteActivationResponse `json:"remoteData"`
	}

	// PluginSummaryResult is the canonical single reusable result envelope for a single plugin.
	PluginSummaryResult = result.Wrap[PluginSummary]

	// PluginSummarySliceResult is the canonical single reusable result envelope for plugin slices.
	PluginSummarySliceResult = appfault.ResultSlice[PluginSummary]

	// RemoteActivationResponseResult is the canonical single reusable result envelope for remote activation.
	RemoteActivationResponseResult = result.Wrap[RemoteActivationResponse]

	// WorkflowResultWrap is the canonical single reusable result envelope for workflow execution.
	WorkflowResultWrap = result.Wrap[WorkflowResult]
)
