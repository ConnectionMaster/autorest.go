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

// ExpressRouteConnectionsClient contains the methods for the ExpressRouteConnections group.
// Don't use this type directly, use NewExpressRouteConnectionsClient() instead.
type ExpressRouteConnectionsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewExpressRouteConnectionsClient creates a new instance of ExpressRouteConnectionsClient with the specified values.
func NewExpressRouteConnectionsClient(con *armcore.Connection, subscriptionID string) *ExpressRouteConnectionsClient {
	return &ExpressRouteConnectionsClient{con: con, subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates a connection between an ExpressRoute gateway and an ExpressRoute circuit.
// If the operation fails it returns the *CloudError error type.
func (client *ExpressRouteConnectionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, expressRouteGatewayName string, connectionName string, putExpressRouteConnectionParameters ExpressRouteConnection, options *ExpressRouteConnectionsBeginCreateOrUpdateOptions) (ExpressRouteConnectionsCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, expressRouteGatewayName, connectionName, putExpressRouteConnectionParameters, options)
	if err != nil {
		return ExpressRouteConnectionsCreateOrUpdatePollerResponse{}, err
	}
	result := ExpressRouteConnectionsCreateOrUpdatePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("ExpressRouteConnectionsClient.CreateOrUpdate", "azure-async-operation", resp, client.con.Pipeline(), client.createOrUpdateHandleError)
	if err != nil {
		return ExpressRouteConnectionsCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &ExpressRouteConnectionsCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new ExpressRouteConnectionsCreateOrUpdatePoller from the specified resume token.
// token - The value must come from a previous call to ExpressRouteConnectionsCreateOrUpdatePoller.ResumeToken().
func (client *ExpressRouteConnectionsClient) ResumeCreateOrUpdate(ctx context.Context, token string) (ExpressRouteConnectionsCreateOrUpdatePollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("ExpressRouteConnectionsClient.CreateOrUpdate", token, client.con.Pipeline(), client.createOrUpdateHandleError)
	if err != nil {
		return ExpressRouteConnectionsCreateOrUpdatePollerResponse{}, err
	}
	poller := &ExpressRouteConnectionsCreateOrUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return ExpressRouteConnectionsCreateOrUpdatePollerResponse{}, err
	}
	result := ExpressRouteConnectionsCreateOrUpdatePollerResponse{
		Poller:      poller,
		RawResponse: resp,
	}
	result.Poller = poller
	return result, nil
}

// CreateOrUpdate - Creates a connection between an ExpressRoute gateway and an ExpressRoute circuit.
// If the operation fails it returns the *CloudError error type.
func (client *ExpressRouteConnectionsClient) createOrUpdate(ctx context.Context, resourceGroupName string, expressRouteGatewayName string, connectionName string, putExpressRouteConnectionParameters ExpressRouteConnection, options *ExpressRouteConnectionsBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, expressRouteGatewayName, connectionName, putExpressRouteConnectionParameters, options)
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
func (client *ExpressRouteConnectionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, expressRouteGatewayName string, connectionName string, putExpressRouteConnectionParameters ExpressRouteConnection, options *ExpressRouteConnectionsBeginCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteGateways/{expressRouteGatewayName}/expressRouteConnections/{connectionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if expressRouteGatewayName == "" {
		return nil, errors.New("parameter expressRouteGatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{expressRouteGatewayName}", url.PathEscape(expressRouteGatewayName))
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
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
	return req, req.MarshalAsJSON(putExpressRouteConnectionParameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *ExpressRouteConnectionsClient) createOrUpdateHandleError(resp *azcore.Response) error {
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

// BeginDelete - Deletes a connection to a ExpressRoute circuit.
// If the operation fails it returns the *CloudError error type.
func (client *ExpressRouteConnectionsClient) BeginDelete(ctx context.Context, resourceGroupName string, expressRouteGatewayName string, connectionName string, options *ExpressRouteConnectionsBeginDeleteOptions) (ExpressRouteConnectionsDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, expressRouteGatewayName, connectionName, options)
	if err != nil {
		return ExpressRouteConnectionsDeletePollerResponse{}, err
	}
	result := ExpressRouteConnectionsDeletePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("ExpressRouteConnectionsClient.Delete", "location", resp, client.con.Pipeline(), client.deleteHandleError)
	if err != nil {
		return ExpressRouteConnectionsDeletePollerResponse{}, err
	}
	result.Poller = &ExpressRouteConnectionsDeletePoller{
		pt: pt,
	}
	return result, nil
}

// ResumeDelete creates a new ExpressRouteConnectionsDeletePoller from the specified resume token.
// token - The value must come from a previous call to ExpressRouteConnectionsDeletePoller.ResumeToken().
func (client *ExpressRouteConnectionsClient) ResumeDelete(ctx context.Context, token string) (ExpressRouteConnectionsDeletePollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("ExpressRouteConnectionsClient.Delete", token, client.con.Pipeline(), client.deleteHandleError)
	if err != nil {
		return ExpressRouteConnectionsDeletePollerResponse{}, err
	}
	poller := &ExpressRouteConnectionsDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return ExpressRouteConnectionsDeletePollerResponse{}, err
	}
	result := ExpressRouteConnectionsDeletePollerResponse{
		Poller:      poller,
		RawResponse: resp,
	}
	result.Poller = poller
	return result, nil
}

// Delete - Deletes a connection to a ExpressRoute circuit.
// If the operation fails it returns the *CloudError error type.
func (client *ExpressRouteConnectionsClient) deleteOperation(ctx context.Context, resourceGroupName string, expressRouteGatewayName string, connectionName string, options *ExpressRouteConnectionsBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, expressRouteGatewayName, connectionName, options)
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
func (client *ExpressRouteConnectionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, expressRouteGatewayName string, connectionName string, options *ExpressRouteConnectionsBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteGateways/{expressRouteGatewayName}/expressRouteConnections/{connectionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if expressRouteGatewayName == "" {
		return nil, errors.New("parameter expressRouteGatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{expressRouteGatewayName}", url.PathEscape(expressRouteGatewayName))
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
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
func (client *ExpressRouteConnectionsClient) deleteHandleError(resp *azcore.Response) error {
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

// Get - Gets the specified ExpressRouteConnection.
// If the operation fails it returns the *CloudError error type.
func (client *ExpressRouteConnectionsClient) Get(ctx context.Context, resourceGroupName string, expressRouteGatewayName string, connectionName string, options *ExpressRouteConnectionsGetOptions) (ExpressRouteConnectionsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, expressRouteGatewayName, connectionName, options)
	if err != nil {
		return ExpressRouteConnectionsGetResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ExpressRouteConnectionsGetResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ExpressRouteConnectionsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ExpressRouteConnectionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, expressRouteGatewayName string, connectionName string, options *ExpressRouteConnectionsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteGateways/{expressRouteGatewayName}/expressRouteConnections/{connectionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if expressRouteGatewayName == "" {
		return nil, errors.New("parameter expressRouteGatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{expressRouteGatewayName}", url.PathEscape(expressRouteGatewayName))
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
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
func (client *ExpressRouteConnectionsClient) getHandleResponse(resp *azcore.Response) (ExpressRouteConnectionsGetResponse, error) {
	result := ExpressRouteConnectionsGetResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.ExpressRouteConnection); err != nil {
		return ExpressRouteConnectionsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ExpressRouteConnectionsClient) getHandleError(resp *azcore.Response) error {
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

// List - Lists ExpressRouteConnections.
// If the operation fails it returns the *CloudError error type.
func (client *ExpressRouteConnectionsClient) List(ctx context.Context, resourceGroupName string, expressRouteGatewayName string, options *ExpressRouteConnectionsListOptions) (ExpressRouteConnectionsListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, expressRouteGatewayName, options)
	if err != nil {
		return ExpressRouteConnectionsListResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ExpressRouteConnectionsListResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ExpressRouteConnectionsListResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *ExpressRouteConnectionsClient) listCreateRequest(ctx context.Context, resourceGroupName string, expressRouteGatewayName string, options *ExpressRouteConnectionsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteGateways/{expressRouteGatewayName}/expressRouteConnections"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if expressRouteGatewayName == "" {
		return nil, errors.New("parameter expressRouteGatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{expressRouteGatewayName}", url.PathEscape(expressRouteGatewayName))
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
func (client *ExpressRouteConnectionsClient) listHandleResponse(resp *azcore.Response) (ExpressRouteConnectionsListResponse, error) {
	result := ExpressRouteConnectionsListResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.ExpressRouteConnectionList); err != nil {
		return ExpressRouteConnectionsListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *ExpressRouteConnectionsClient) listHandleError(resp *azcore.Response) error {
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
