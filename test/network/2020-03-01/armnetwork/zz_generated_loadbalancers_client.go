// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// LoadBalancersClient contains the methods for the LoadBalancers group.
// Don't use this type directly, use NewLoadBalancersClient() instead.
type LoadBalancersClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewLoadBalancersClient creates a new instance of LoadBalancersClient with the specified values.
func NewLoadBalancersClient(con *armcore.Connection, subscriptionID string) *LoadBalancersClient {
	return &LoadBalancersClient{con: con, subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates or updates a load balancer.
// If the operation fails it returns the *CloudError error type.
func (client *LoadBalancersClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, loadBalancerName string, parameters LoadBalancer, options *LoadBalancersBeginCreateOrUpdateOptions) (LoadBalancersCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, loadBalancerName, parameters, options)
	if err != nil {
		return LoadBalancersCreateOrUpdatePollerResponse{}, err
	}
	result := LoadBalancersCreateOrUpdatePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("LoadBalancersClient.CreateOrUpdate", "azure-async-operation", resp, client.con.Pipeline(), client.createOrUpdateHandleError)
	if err != nil {
		return LoadBalancersCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &LoadBalancersCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new LoadBalancersCreateOrUpdatePoller from the specified resume token.
// token - The value must come from a previous call to LoadBalancersCreateOrUpdatePoller.ResumeToken().
func (client *LoadBalancersClient) ResumeCreateOrUpdate(ctx context.Context, token string) (LoadBalancersCreateOrUpdatePollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("LoadBalancersClient.CreateOrUpdate", token, client.con.Pipeline(), client.createOrUpdateHandleError)
	if err != nil {
		return LoadBalancersCreateOrUpdatePollerResponse{}, err
	}
	poller := &LoadBalancersCreateOrUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return LoadBalancersCreateOrUpdatePollerResponse{}, err
	}
	result := LoadBalancersCreateOrUpdatePollerResponse{
		Poller:      poller,
		RawResponse: resp,
	}
	result.Poller = poller
	return result, nil
}

// CreateOrUpdate - Creates or updates a load balancer.
// If the operation fails it returns the *CloudError error type.
func (client *LoadBalancersClient) createOrUpdate(ctx context.Context, resourceGroupName string, loadBalancerName string, parameters LoadBalancer, options *LoadBalancersBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, loadBalancerName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *LoadBalancersClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, loadBalancerName string, parameters LoadBalancer, options *LoadBalancersBeginCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if loadBalancerName == "" {
		return nil, errors.New("parameter loadBalancerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancerName}", url.PathEscape(loadBalancerName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *LoadBalancersClient) createOrUpdateHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// BeginDelete - Deletes the specified load balancer.
// If the operation fails it returns the *CloudError error type.
func (client *LoadBalancersClient) BeginDelete(ctx context.Context, resourceGroupName string, loadBalancerName string, options *LoadBalancersBeginDeleteOptions) (LoadBalancersDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, loadBalancerName, options)
	if err != nil {
		return LoadBalancersDeletePollerResponse{}, err
	}
	result := LoadBalancersDeletePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("LoadBalancersClient.Delete", "location", resp, client.con.Pipeline(), client.deleteHandleError)
	if err != nil {
		return LoadBalancersDeletePollerResponse{}, err
	}
	result.Poller = &LoadBalancersDeletePoller{
		pt: pt,
	}
	return result, nil
}

// ResumeDelete creates a new LoadBalancersDeletePoller from the specified resume token.
// token - The value must come from a previous call to LoadBalancersDeletePoller.ResumeToken().
func (client *LoadBalancersClient) ResumeDelete(ctx context.Context, token string) (LoadBalancersDeletePollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("LoadBalancersClient.Delete", token, client.con.Pipeline(), client.deleteHandleError)
	if err != nil {
		return LoadBalancersDeletePollerResponse{}, err
	}
	poller := &LoadBalancersDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return LoadBalancersDeletePollerResponse{}, err
	}
	result := LoadBalancersDeletePollerResponse{
		Poller:      poller,
		RawResponse: resp,
	}
	result.Poller = poller
	return result, nil
}

// Delete - Deletes the specified load balancer.
// If the operation fails it returns the *CloudError error type.
func (client *LoadBalancersClient) deleteOperation(ctx context.Context, resourceGroupName string, loadBalancerName string, options *LoadBalancersBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, loadBalancerName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *LoadBalancersClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, loadBalancerName string, options *LoadBalancersBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if loadBalancerName == "" {
		return nil, errors.New("parameter loadBalancerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancerName}", url.PathEscape(loadBalancerName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *LoadBalancersClient) deleteHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Get - Gets the specified load balancer.
// If the operation fails it returns the *CloudError error type.
func (client *LoadBalancersClient) Get(ctx context.Context, resourceGroupName string, loadBalancerName string, options *LoadBalancersGetOptions) (LoadBalancersGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, loadBalancerName, options)
	if err != nil {
		return LoadBalancersGetResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return LoadBalancersGetResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return LoadBalancersGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *LoadBalancersClient) getCreateRequest(ctx context.Context, resourceGroupName string, loadBalancerName string, options *LoadBalancersGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if loadBalancerName == "" {
		return nil, errors.New("parameter loadBalancerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancerName}", url.PathEscape(loadBalancerName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *LoadBalancersClient) getHandleResponse(resp *azcore.Response) (LoadBalancersGetResponse, error) {
	result := LoadBalancersGetResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.LoadBalancer); err != nil {
		return LoadBalancersGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *LoadBalancersClient) getHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// List - Gets all the load balancers in a resource group.
// If the operation fails it returns the *CloudError error type.
func (client *LoadBalancersClient) List(resourceGroupName string, options *LoadBalancersListOptions) *LoadBalancersListPager {
	return &LoadBalancersListPager{
		client: client,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp LoadBalancersListResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.LoadBalancerListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *LoadBalancersClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *LoadBalancersListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *LoadBalancersClient) listHandleResponse(resp *azcore.Response) (LoadBalancersListResponse, error) {
	result := LoadBalancersListResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.LoadBalancerListResult); err != nil {
		return LoadBalancersListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *LoadBalancersClient) listHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// ListAll - Gets all the load balancers in a subscription.
// If the operation fails it returns the *CloudError error type.
func (client *LoadBalancersClient) ListAll(options *LoadBalancersListAllOptions) *LoadBalancersListAllPager {
	return &LoadBalancersListAllPager{
		client: client,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listAllCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp LoadBalancersListAllResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.LoadBalancerListResult.NextLink)
		},
	}
}

// listAllCreateRequest creates the ListAll request.
func (client *LoadBalancersClient) listAllCreateRequest(ctx context.Context, options *LoadBalancersListAllOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/loadBalancers"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listAllHandleResponse handles the ListAll response.
func (client *LoadBalancersClient) listAllHandleResponse(resp *azcore.Response) (LoadBalancersListAllResponse, error) {
	result := LoadBalancersListAllResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.LoadBalancerListResult); err != nil {
		return LoadBalancersListAllResponse{}, err
	}
	return result, nil
}

// listAllHandleError handles the ListAll error response.
func (client *LoadBalancersClient) listAllHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// UpdateTags - Updates a load balancer tags.
// If the operation fails it returns the *CloudError error type.
func (client *LoadBalancersClient) UpdateTags(ctx context.Context, resourceGroupName string, loadBalancerName string, parameters TagsObject, options *LoadBalancersUpdateTagsOptions) (LoadBalancersUpdateTagsResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, loadBalancerName, parameters, options)
	if err != nil {
		return LoadBalancersUpdateTagsResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return LoadBalancersUpdateTagsResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return LoadBalancersUpdateTagsResponse{}, client.updateTagsHandleError(resp)
	}
	return client.updateTagsHandleResponse(resp)
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *LoadBalancersClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, loadBalancerName string, parameters TagsObject, options *LoadBalancersUpdateTagsOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if loadBalancerName == "" {
		return nil, errors.New("parameter loadBalancerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancerName}", url.PathEscape(loadBalancerName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *LoadBalancersClient) updateTagsHandleResponse(resp *azcore.Response) (LoadBalancersUpdateTagsResponse, error) {
	result := LoadBalancersUpdateTagsResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.LoadBalancer); err != nil {
		return LoadBalancersUpdateTagsResponse{}, err
	}
	return result, nil
}

// updateTagsHandleError handles the UpdateTags error response.
func (client *LoadBalancersClient) updateTagsHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}
