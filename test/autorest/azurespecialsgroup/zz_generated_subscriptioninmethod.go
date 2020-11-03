// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azurespecialsgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// SubscriptionInMethodOperations contains the methods for the SubscriptionInMethod group.
type SubscriptionInMethodOperations interface {
	// PostMethodLocalNull - POST method with subscriptionId modeled in the method. pass in subscription id = null, client-side validation should prevent you
	// from making this call
	PostMethodLocalNull(ctx context.Context, subscriptionId string, options *SubscriptionInMethodPostMethodLocalNullOptions) (*http.Response, error)
	// PostMethodLocalValid - POST method with subscriptionId modeled in the method. pass in subscription id = '1234-5678-9012-3456' to succeed
	PostMethodLocalValid(ctx context.Context, subscriptionId string, options *SubscriptionInMethodPostMethodLocalValidOptions) (*http.Response, error)
	// PostPathLocalValid - POST method with subscriptionId modeled in the method. pass in subscription id = '1234-5678-9012-3456' to succeed
	PostPathLocalValid(ctx context.Context, subscriptionId string, options *SubscriptionInMethodPostPathLocalValidOptions) (*http.Response, error)
	// PostSwaggerLocalValid - POST method with subscriptionId modeled in the method. pass in subscription id = '1234-5678-9012-3456' to succeed
	PostSwaggerLocalValid(ctx context.Context, subscriptionId string, options *SubscriptionInMethodPostSwaggerLocalValidOptions) (*http.Response, error)
}

// SubscriptionInMethodClient implements the SubscriptionInMethodOperations interface.
// Don't use this type directly, use NewSubscriptionInMethodClient() instead.
type SubscriptionInMethodClient struct {
	*Client
}

// NewSubscriptionInMethodClient creates a new instance of SubscriptionInMethodClient with the specified values.
func NewSubscriptionInMethodClient(c *Client) SubscriptionInMethodOperations {
	return &SubscriptionInMethodClient{Client: c}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *SubscriptionInMethodClient) Do(req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(req)
}

// PostMethodLocalNull - POST method with subscriptionId modeled in the method. pass in subscription id = null, client-side validation should prevent you
// from making this call
func (client *SubscriptionInMethodClient) PostMethodLocalNull(ctx context.Context, subscriptionId string, options *SubscriptionInMethodPostMethodLocalNullOptions) (*http.Response, error) {
	req, err := client.PostMethodLocalNullCreateRequest(ctx, subscriptionId, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.PostMethodLocalNullHandleError(resp)
	}
	return resp.Response, nil
}

// PostMethodLocalNullCreateRequest creates the PostMethodLocalNull request.
func (client *SubscriptionInMethodClient) PostMethodLocalNullCreateRequest(ctx context.Context, subscriptionId string, options *SubscriptionInMethodPostMethodLocalNullOptions) (*azcore.Request, error) {
	urlPath := "/azurespecials/subscriptionId/method/string/none/path/local/null/{subscriptionId}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionId))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// PostMethodLocalNullHandleError handles the PostMethodLocalNull error response.
func (client *SubscriptionInMethodClient) PostMethodLocalNullHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PostMethodLocalValid - POST method with subscriptionId modeled in the method. pass in subscription id = '1234-5678-9012-3456' to succeed
func (client *SubscriptionInMethodClient) PostMethodLocalValid(ctx context.Context, subscriptionId string, options *SubscriptionInMethodPostMethodLocalValidOptions) (*http.Response, error) {
	req, err := client.PostMethodLocalValidCreateRequest(ctx, subscriptionId, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.PostMethodLocalValidHandleError(resp)
	}
	return resp.Response, nil
}

// PostMethodLocalValidCreateRequest creates the PostMethodLocalValid request.
func (client *SubscriptionInMethodClient) PostMethodLocalValidCreateRequest(ctx context.Context, subscriptionId string, options *SubscriptionInMethodPostMethodLocalValidOptions) (*azcore.Request, error) {
	urlPath := "/azurespecials/subscriptionId/method/string/none/path/local/1234-5678-9012-3456/{subscriptionId}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionId))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// PostMethodLocalValidHandleError handles the PostMethodLocalValid error response.
func (client *SubscriptionInMethodClient) PostMethodLocalValidHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PostPathLocalValid - POST method with subscriptionId modeled in the method. pass in subscription id = '1234-5678-9012-3456' to succeed
func (client *SubscriptionInMethodClient) PostPathLocalValid(ctx context.Context, subscriptionId string, options *SubscriptionInMethodPostPathLocalValidOptions) (*http.Response, error) {
	req, err := client.PostPathLocalValidCreateRequest(ctx, subscriptionId, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.PostPathLocalValidHandleError(resp)
	}
	return resp.Response, nil
}

// PostPathLocalValidCreateRequest creates the PostPathLocalValid request.
func (client *SubscriptionInMethodClient) PostPathLocalValidCreateRequest(ctx context.Context, subscriptionId string, options *SubscriptionInMethodPostPathLocalValidOptions) (*azcore.Request, error) {
	urlPath := "/azurespecials/subscriptionId/path/string/none/path/local/1234-5678-9012-3456/{subscriptionId}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionId))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// PostPathLocalValidHandleError handles the PostPathLocalValid error response.
func (client *SubscriptionInMethodClient) PostPathLocalValidHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PostSwaggerLocalValid - POST method with subscriptionId modeled in the method. pass in subscription id = '1234-5678-9012-3456' to succeed
func (client *SubscriptionInMethodClient) PostSwaggerLocalValid(ctx context.Context, subscriptionId string, options *SubscriptionInMethodPostSwaggerLocalValidOptions) (*http.Response, error) {
	req, err := client.PostSwaggerLocalValidCreateRequest(ctx, subscriptionId, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.PostSwaggerLocalValidHandleError(resp)
	}
	return resp.Response, nil
}

// PostSwaggerLocalValidCreateRequest creates the PostSwaggerLocalValid request.
func (client *SubscriptionInMethodClient) PostSwaggerLocalValidCreateRequest(ctx context.Context, subscriptionId string, options *SubscriptionInMethodPostSwaggerLocalValidOptions) (*azcore.Request, error) {
	urlPath := "/azurespecials/subscriptionId/swagger/string/none/path/local/1234-5678-9012-3456/{subscriptionId}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionId))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// PostSwaggerLocalValidHandleError handles the PostSwaggerLocalValid error response.
func (client *SubscriptionInMethodClient) PostSwaggerLocalValidHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}