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

// AzureFirewallsClient contains the methods for the AzureFirewalls group.
// Don't use this type directly, use NewAzureFirewallsClient() instead.
type AzureFirewallsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewAzureFirewallsClient creates a new instance of AzureFirewallsClient with the specified values.
func NewAzureFirewallsClient(con *armcore.Connection, subscriptionID string) *AzureFirewallsClient {
	return &AzureFirewallsClient{con: con, subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates or updates the specified Azure Firewall.
// If the operation fails it returns the *CloudError error type.
func (client *AzureFirewallsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, azureFirewallName string, parameters AzureFirewall, options *AzureFirewallsBeginCreateOrUpdateOptions) (AzureFirewallsCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, azureFirewallName, parameters, options)
	if err != nil {
		return AzureFirewallsCreateOrUpdatePollerResponse{}, err
	}
	result := AzureFirewallsCreateOrUpdatePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("AzureFirewallsClient.CreateOrUpdate", "azure-async-operation", resp, client.con.Pipeline(), client.createOrUpdateHandleError)
	if err != nil {
		return AzureFirewallsCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &AzureFirewallsCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new AzureFirewallsCreateOrUpdatePoller from the specified resume token.
// token - The value must come from a previous call to AzureFirewallsCreateOrUpdatePoller.ResumeToken().
func (client *AzureFirewallsClient) ResumeCreateOrUpdate(ctx context.Context, token string) (AzureFirewallsCreateOrUpdatePollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("AzureFirewallsClient.CreateOrUpdate", token, client.con.Pipeline(), client.createOrUpdateHandleError)
	if err != nil {
		return AzureFirewallsCreateOrUpdatePollerResponse{}, err
	}
	poller := &AzureFirewallsCreateOrUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return AzureFirewallsCreateOrUpdatePollerResponse{}, err
	}
	result := AzureFirewallsCreateOrUpdatePollerResponse{
		Poller:      poller,
		RawResponse: resp,
	}
	result.Poller = poller
	return result, nil
}

// CreateOrUpdate - Creates or updates the specified Azure Firewall.
// If the operation fails it returns the *CloudError error type.
func (client *AzureFirewallsClient) createOrUpdate(ctx context.Context, resourceGroupName string, azureFirewallName string, parameters AzureFirewall, options *AzureFirewallsBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, azureFirewallName, parameters, options)
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
func (client *AzureFirewallsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, azureFirewallName string, parameters AzureFirewall, options *AzureFirewallsBeginCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/azureFirewalls/{azureFirewallName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if azureFirewallName == "" {
		return nil, errors.New("parameter azureFirewallName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{azureFirewallName}", url.PathEscape(azureFirewallName))
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
func (client *AzureFirewallsClient) createOrUpdateHandleError(resp *azcore.Response) error {
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

// BeginDelete - Deletes the specified Azure Firewall.
// If the operation fails it returns the *CloudError error type.
func (client *AzureFirewallsClient) BeginDelete(ctx context.Context, resourceGroupName string, azureFirewallName string, options *AzureFirewallsBeginDeleteOptions) (AzureFirewallsDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, azureFirewallName, options)
	if err != nil {
		return AzureFirewallsDeletePollerResponse{}, err
	}
	result := AzureFirewallsDeletePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("AzureFirewallsClient.Delete", "location", resp, client.con.Pipeline(), client.deleteHandleError)
	if err != nil {
		return AzureFirewallsDeletePollerResponse{}, err
	}
	result.Poller = &AzureFirewallsDeletePoller{
		pt: pt,
	}
	return result, nil
}

// ResumeDelete creates a new AzureFirewallsDeletePoller from the specified resume token.
// token - The value must come from a previous call to AzureFirewallsDeletePoller.ResumeToken().
func (client *AzureFirewallsClient) ResumeDelete(ctx context.Context, token string) (AzureFirewallsDeletePollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("AzureFirewallsClient.Delete", token, client.con.Pipeline(), client.deleteHandleError)
	if err != nil {
		return AzureFirewallsDeletePollerResponse{}, err
	}
	poller := &AzureFirewallsDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return AzureFirewallsDeletePollerResponse{}, err
	}
	result := AzureFirewallsDeletePollerResponse{
		Poller:      poller,
		RawResponse: resp,
	}
	result.Poller = poller
	return result, nil
}

// Delete - Deletes the specified Azure Firewall.
// If the operation fails it returns the *CloudError error type.
func (client *AzureFirewallsClient) deleteOperation(ctx context.Context, resourceGroupName string, azureFirewallName string, options *AzureFirewallsBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, azureFirewallName, options)
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
func (client *AzureFirewallsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, azureFirewallName string, options *AzureFirewallsBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/azureFirewalls/{azureFirewallName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if azureFirewallName == "" {
		return nil, errors.New("parameter azureFirewallName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{azureFirewallName}", url.PathEscape(azureFirewallName))
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
func (client *AzureFirewallsClient) deleteHandleError(resp *azcore.Response) error {
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

// Get - Gets the specified Azure Firewall.
// If the operation fails it returns the *CloudError error type.
func (client *AzureFirewallsClient) Get(ctx context.Context, resourceGroupName string, azureFirewallName string, options *AzureFirewallsGetOptions) (AzureFirewallsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, azureFirewallName, options)
	if err != nil {
		return AzureFirewallsGetResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return AzureFirewallsGetResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return AzureFirewallsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AzureFirewallsClient) getCreateRequest(ctx context.Context, resourceGroupName string, azureFirewallName string, options *AzureFirewallsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/azureFirewalls/{azureFirewallName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if azureFirewallName == "" {
		return nil, errors.New("parameter azureFirewallName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{azureFirewallName}", url.PathEscape(azureFirewallName))
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

// getHandleResponse handles the Get response.
func (client *AzureFirewallsClient) getHandleResponse(resp *azcore.Response) (AzureFirewallsGetResponse, error) {
	result := AzureFirewallsGetResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.AzureFirewall); err != nil {
		return AzureFirewallsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *AzureFirewallsClient) getHandleError(resp *azcore.Response) error {
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

// List - Lists all Azure Firewalls in a resource group.
// If the operation fails it returns the *CloudError error type.
func (client *AzureFirewallsClient) List(resourceGroupName string, options *AzureFirewallsListOptions) *AzureFirewallsListPager {
	return &AzureFirewallsListPager{
		client: client,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp AzureFirewallsListResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.AzureFirewallListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *AzureFirewallsClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *AzureFirewallsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/azureFirewalls"
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
func (client *AzureFirewallsClient) listHandleResponse(resp *azcore.Response) (AzureFirewallsListResponse, error) {
	result := AzureFirewallsListResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.AzureFirewallListResult); err != nil {
		return AzureFirewallsListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *AzureFirewallsClient) listHandleError(resp *azcore.Response) error {
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

// ListAll - Gets all the Azure Firewalls in a subscription.
// If the operation fails it returns the *CloudError error type.
func (client *AzureFirewallsClient) ListAll(options *AzureFirewallsListAllOptions) *AzureFirewallsListAllPager {
	return &AzureFirewallsListAllPager{
		client: client,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listAllCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp AzureFirewallsListAllResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.AzureFirewallListResult.NextLink)
		},
	}
}

// listAllCreateRequest creates the ListAll request.
func (client *AzureFirewallsClient) listAllCreateRequest(ctx context.Context, options *AzureFirewallsListAllOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/azureFirewalls"
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
func (client *AzureFirewallsClient) listAllHandleResponse(resp *azcore.Response) (AzureFirewallsListAllResponse, error) {
	result := AzureFirewallsListAllResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.AzureFirewallListResult); err != nil {
		return AzureFirewallsListAllResponse{}, err
	}
	return result, nil
}

// listAllHandleError handles the ListAll error response.
func (client *AzureFirewallsClient) listAllHandleError(resp *azcore.Response) error {
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

// BeginUpdateTags - Updates tags of an Azure Firewall resource.
// If the operation fails it returns the *CloudError error type.
func (client *AzureFirewallsClient) BeginUpdateTags(ctx context.Context, resourceGroupName string, azureFirewallName string, parameters TagsObject, options *AzureFirewallsBeginUpdateTagsOptions) (AzureFirewallsUpdateTagsPollerResponse, error) {
	resp, err := client.updateTags(ctx, resourceGroupName, azureFirewallName, parameters, options)
	if err != nil {
		return AzureFirewallsUpdateTagsPollerResponse{}, err
	}
	result := AzureFirewallsUpdateTagsPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("AzureFirewallsClient.UpdateTags", "azure-async-operation", resp, client.con.Pipeline(), client.updateTagsHandleError)
	if err != nil {
		return AzureFirewallsUpdateTagsPollerResponse{}, err
	}
	result.Poller = &AzureFirewallsUpdateTagsPoller{
		pt: pt,
	}
	return result, nil
}

// ResumeUpdateTags creates a new AzureFirewallsUpdateTagsPoller from the specified resume token.
// token - The value must come from a previous call to AzureFirewallsUpdateTagsPoller.ResumeToken().
func (client *AzureFirewallsClient) ResumeUpdateTags(ctx context.Context, token string) (AzureFirewallsUpdateTagsPollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("AzureFirewallsClient.UpdateTags", token, client.con.Pipeline(), client.updateTagsHandleError)
	if err != nil {
		return AzureFirewallsUpdateTagsPollerResponse{}, err
	}
	poller := &AzureFirewallsUpdateTagsPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return AzureFirewallsUpdateTagsPollerResponse{}, err
	}
	result := AzureFirewallsUpdateTagsPollerResponse{
		Poller:      poller,
		RawResponse: resp,
	}
	result.Poller = poller
	return result, nil
}

// UpdateTags - Updates tags of an Azure Firewall resource.
// If the operation fails it returns the *CloudError error type.
func (client *AzureFirewallsClient) updateTags(ctx context.Context, resourceGroupName string, azureFirewallName string, parameters TagsObject, options *AzureFirewallsBeginUpdateTagsOptions) (*azcore.Response, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, azureFirewallName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.updateTagsHandleError(resp)
	}
	return resp, nil
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *AzureFirewallsClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, azureFirewallName string, parameters TagsObject, options *AzureFirewallsBeginUpdateTagsOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/azureFirewalls/{azureFirewallName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if azureFirewallName == "" {
		return nil, errors.New("parameter azureFirewallName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{azureFirewallName}", url.PathEscape(azureFirewallName))
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

// updateTagsHandleError handles the UpdateTags error response.
func (client *AzureFirewallsClient) updateTagsHandleError(resp *azcore.Response) error {
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
