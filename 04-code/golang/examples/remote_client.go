package examples

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"coding-guidelines/common/pkg/apperror"
	"coding-guidelines/common/pkg/result"
)

// RemoteActivationResponse contains response data from downstream WordPress REST API.
type RemoteActivationResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Version string `json:"version"`
}

// WordPressClient handles proxy communication with remote WordPress instances.
type WordPressClient struct {
	baseUrl    string
	httpClient *http.Client
}

// NewWordPressClient creates a client pointing to a remote instance.
func NewWordPressClient(baseUrl string) *WordPressClient {
	return &WordPressClient{
		baseUrl:    baseUrl,
		httpClient: http.DefaultClient,
	}
}

// ActivateRemotePlugin dispatches an activation request to the remote endpoint.
func (c *WordPressClient) ActivateRemotePlugin(
	ctx context.Context,
	siteId int64,
	slug string,
) result.Result[RemoteActivationResponse] {
	if len(slug) == 0 {
		return result.NewFailureWithType[RemoteActivationResponse]("E1001", "slug cannot be empty", "wp.client")
	}

	url := fmt.Sprintf("%s/wp-json/riseup/v1/plugins/%s/activate", c.baseUrl, slug)

	// Simulated remote HTTP request failure
	if slug == "broken-plugin" {
		rawErr := errors.New("HTTP 502 Bad Gateway from upstream NGINX")
		fault := apperror.WrapWithDetails(rawErr, "wp.client.activate", "E3003", "delegated remote activation failed", "wp.client", apperror.ErrorTypeExecution, apperror.SeverityError, nil).
			WithUrl(url).
			WithStatusCode(502).
			WithSiteId(siteId).
			WithSlug(slug)

		return result.FailureResult[RemoteActivationResponse](fault)
	}

	return result.SuccessResult(RemoteActivationResponse{
		Success: true,
		Message: "Plugin activated successfully",
		Version: "2.4.0",
	})
}
