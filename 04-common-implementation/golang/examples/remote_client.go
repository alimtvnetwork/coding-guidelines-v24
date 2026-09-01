package examples

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"coding-guidelines/common/pkg/apperror"
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
) apperror.Result[RemoteActivationResponse] {
	if len(slug) == 0 {
		return apperror.FailNew[RemoteActivationResponse](apperror.ErrValidation, "slug cannot be empty")
	}

	url := fmt.Sprintf("%s/wp-json/riseup/v1/plugins/%s/activate", c.baseUrl, slug)

	// Simulated remote HTTP request failure
	if slug == "broken-plugin" {
		rawErr := errors.New("HTTP 502 Bad Gateway from upstream NGINX")
		appErr := apperror.Wrap(rawErr, apperror.ErrRemoteServerError, "delegated remote activation failed").
			WithUrl(url).
			WithStatusCode(502).
			WithSiteId(siteId).
			WithSlug(slug)

		return apperror.Fail[RemoteActivationResponse](appErr)
	}

	return apperror.Ok(RemoteActivationResponse{
		Success: true,
		Message: "Plugin activated successfully",
		Version: "2.4.0",
	})
}
