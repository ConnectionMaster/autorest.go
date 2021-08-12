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

// PacketCapturesClient contains the methods for the PacketCaptures group.
// Don't use this type directly, use NewPacketCapturesClient() instead.
type PacketCapturesClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewPacketCapturesClient creates a new instance of PacketCapturesClient with the specified values.
func NewPacketCapturesClient(con *armcore.Connection, subscriptionID string) *PacketCapturesClient {
	return &PacketCapturesClient{con: con, subscriptionID: subscriptionID}
}

// BeginCreate - Create and start a packet capture on the specified VM.
// If the operation fails it returns the *ErrorResponse error type.
func (client *PacketCapturesClient) BeginCreate(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string, parameters PacketCapture, options *PacketCapturesBeginCreateOptions) (PacketCapturesCreatePollerResponse, error) {
	resp, err := client.create(ctx, resourceGroupName, networkWatcherName, packetCaptureName, parameters, options)
	if err != nil {
		return PacketCapturesCreatePollerResponse{}, err
	}
	result := PacketCapturesCreatePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("PacketCapturesClient.Create", "azure-async-operation", resp, client.con.Pipeline(), client.createHandleError)
	if err != nil {
		return PacketCapturesCreatePollerResponse{}, err
	}
	result.Poller = &PacketCapturesCreatePoller{
		pt: pt,
	}
	return result, nil
}

// ResumeCreate creates a new PacketCapturesCreatePoller from the specified resume token.
// token - The value must come from a previous call to PacketCapturesCreatePoller.ResumeToken().
func (client *PacketCapturesClient) ResumeCreate(ctx context.Context, token string) (PacketCapturesCreatePollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("PacketCapturesClient.Create", token, client.con.Pipeline(), client.createHandleError)
	if err != nil {
		return PacketCapturesCreatePollerResponse{}, err
	}
	poller := &PacketCapturesCreatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return PacketCapturesCreatePollerResponse{}, err
	}
	result := PacketCapturesCreatePollerResponse{
		Poller:      poller,
		RawResponse: resp,
	}
	result.Poller = poller
	return result, nil
}

// Create - Create and start a packet capture on the specified VM.
// If the operation fails it returns the *ErrorResponse error type.
func (client *PacketCapturesClient) create(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string, parameters PacketCapture, options *PacketCapturesBeginCreateOptions) (*azcore.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, networkWatcherName, packetCaptureName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusCreated) {
		return nil, client.createHandleError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *PacketCapturesClient) createCreateRequest(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string, parameters PacketCapture, options *PacketCapturesBeginCreateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/packetCaptures/{packetCaptureName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkWatcherName == "" {
		return nil, errors.New("parameter networkWatcherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkWatcherName}", url.PathEscape(networkWatcherName))
	if packetCaptureName == "" {
		return nil, errors.New("parameter packetCaptureName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packetCaptureName}", url.PathEscape(packetCaptureName))
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

// createHandleError handles the Create error response.
func (client *PacketCapturesClient) createHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// BeginDelete - Deletes the specified packet capture session.
// If the operation fails it returns the *ErrorResponse error type.
func (client *PacketCapturesClient) BeginDelete(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string, options *PacketCapturesBeginDeleteOptions) (PacketCapturesDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, networkWatcherName, packetCaptureName, options)
	if err != nil {
		return PacketCapturesDeletePollerResponse{}, err
	}
	result := PacketCapturesDeletePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("PacketCapturesClient.Delete", "location", resp, client.con.Pipeline(), client.deleteHandleError)
	if err != nil {
		return PacketCapturesDeletePollerResponse{}, err
	}
	result.Poller = &PacketCapturesDeletePoller{
		pt: pt,
	}
	return result, nil
}

// ResumeDelete creates a new PacketCapturesDeletePoller from the specified resume token.
// token - The value must come from a previous call to PacketCapturesDeletePoller.ResumeToken().
func (client *PacketCapturesClient) ResumeDelete(ctx context.Context, token string) (PacketCapturesDeletePollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("PacketCapturesClient.Delete", token, client.con.Pipeline(), client.deleteHandleError)
	if err != nil {
		return PacketCapturesDeletePollerResponse{}, err
	}
	poller := &PacketCapturesDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return PacketCapturesDeletePollerResponse{}, err
	}
	result := PacketCapturesDeletePollerResponse{
		Poller:      poller,
		RawResponse: resp,
	}
	result.Poller = poller
	return result, nil
}

// Delete - Deletes the specified packet capture session.
// If the operation fails it returns the *ErrorResponse error type.
func (client *PacketCapturesClient) deleteOperation(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string, options *PacketCapturesBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, networkWatcherName, packetCaptureName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *PacketCapturesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string, options *PacketCapturesBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/packetCaptures/{packetCaptureName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkWatcherName == "" {
		return nil, errors.New("parameter networkWatcherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkWatcherName}", url.PathEscape(networkWatcherName))
	if packetCaptureName == "" {
		return nil, errors.New("parameter packetCaptureName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packetCaptureName}", url.PathEscape(packetCaptureName))
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
func (client *PacketCapturesClient) deleteHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Get - Gets a packet capture session by name.
// If the operation fails it returns the *ErrorResponse error type.
func (client *PacketCapturesClient) Get(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string, options *PacketCapturesGetOptions) (PacketCapturesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, networkWatcherName, packetCaptureName, options)
	if err != nil {
		return PacketCapturesGetResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return PacketCapturesGetResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return PacketCapturesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PacketCapturesClient) getCreateRequest(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string, options *PacketCapturesGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/packetCaptures/{packetCaptureName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkWatcherName == "" {
		return nil, errors.New("parameter networkWatcherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkWatcherName}", url.PathEscape(networkWatcherName))
	if packetCaptureName == "" {
		return nil, errors.New("parameter packetCaptureName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packetCaptureName}", url.PathEscape(packetCaptureName))
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
func (client *PacketCapturesClient) getHandleResponse(resp *azcore.Response) (PacketCapturesGetResponse, error) {
	result := PacketCapturesGetResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.PacketCaptureResult); err != nil {
		return PacketCapturesGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *PacketCapturesClient) getHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// BeginGetStatus - Query the status of a running packet capture session.
// If the operation fails it returns the *ErrorResponse error type.
func (client *PacketCapturesClient) BeginGetStatus(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string, options *PacketCapturesBeginGetStatusOptions) (PacketCapturesGetStatusPollerResponse, error) {
	resp, err := client.getStatus(ctx, resourceGroupName, networkWatcherName, packetCaptureName, options)
	if err != nil {
		return PacketCapturesGetStatusPollerResponse{}, err
	}
	result := PacketCapturesGetStatusPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("PacketCapturesClient.GetStatus", "location", resp, client.con.Pipeline(), client.getStatusHandleError)
	if err != nil {
		return PacketCapturesGetStatusPollerResponse{}, err
	}
	result.Poller = &PacketCapturesGetStatusPoller{
		pt: pt,
	}
	return result, nil
}

// ResumeGetStatus creates a new PacketCapturesGetStatusPoller from the specified resume token.
// token - The value must come from a previous call to PacketCapturesGetStatusPoller.ResumeToken().
func (client *PacketCapturesClient) ResumeGetStatus(ctx context.Context, token string) (PacketCapturesGetStatusPollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("PacketCapturesClient.GetStatus", token, client.con.Pipeline(), client.getStatusHandleError)
	if err != nil {
		return PacketCapturesGetStatusPollerResponse{}, err
	}
	poller := &PacketCapturesGetStatusPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return PacketCapturesGetStatusPollerResponse{}, err
	}
	result := PacketCapturesGetStatusPollerResponse{
		Poller:      poller,
		RawResponse: resp,
	}
	result.Poller = poller
	return result, nil
}

// GetStatus - Query the status of a running packet capture session.
// If the operation fails it returns the *ErrorResponse error type.
func (client *PacketCapturesClient) getStatus(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string, options *PacketCapturesBeginGetStatusOptions) (*azcore.Response, error) {
	req, err := client.getStatusCreateRequest(ctx, resourceGroupName, networkWatcherName, packetCaptureName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.getStatusHandleError(resp)
	}
	return resp, nil
}

// getStatusCreateRequest creates the GetStatus request.
func (client *PacketCapturesClient) getStatusCreateRequest(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string, options *PacketCapturesBeginGetStatusOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/packetCaptures/{packetCaptureName}/queryStatus"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkWatcherName == "" {
		return nil, errors.New("parameter networkWatcherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkWatcherName}", url.PathEscape(networkWatcherName))
	if packetCaptureName == "" {
		return nil, errors.New("parameter packetCaptureName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packetCaptureName}", url.PathEscape(packetCaptureName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
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

// getStatusHandleError handles the GetStatus error response.
func (client *PacketCapturesClient) getStatusHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// List - Lists all packet capture sessions within the specified resource group.
// If the operation fails it returns the *ErrorResponse error type.
func (client *PacketCapturesClient) List(ctx context.Context, resourceGroupName string, networkWatcherName string, options *PacketCapturesListOptions) (PacketCapturesListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, networkWatcherName, options)
	if err != nil {
		return PacketCapturesListResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return PacketCapturesListResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return PacketCapturesListResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *PacketCapturesClient) listCreateRequest(ctx context.Context, resourceGroupName string, networkWatcherName string, options *PacketCapturesListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/packetCaptures"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkWatcherName == "" {
		return nil, errors.New("parameter networkWatcherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkWatcherName}", url.PathEscape(networkWatcherName))
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
func (client *PacketCapturesClient) listHandleResponse(resp *azcore.Response) (PacketCapturesListResponse, error) {
	result := PacketCapturesListResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.PacketCaptureListResult); err != nil {
		return PacketCapturesListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *PacketCapturesClient) listHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// BeginStop - Stops a specified packet capture session.
// If the operation fails it returns the *ErrorResponse error type.
func (client *PacketCapturesClient) BeginStop(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string, options *PacketCapturesBeginStopOptions) (PacketCapturesStopPollerResponse, error) {
	resp, err := client.stop(ctx, resourceGroupName, networkWatcherName, packetCaptureName, options)
	if err != nil {
		return PacketCapturesStopPollerResponse{}, err
	}
	result := PacketCapturesStopPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("PacketCapturesClient.Stop", "location", resp, client.con.Pipeline(), client.stopHandleError)
	if err != nil {
		return PacketCapturesStopPollerResponse{}, err
	}
	result.Poller = &PacketCapturesStopPoller{
		pt: pt,
	}
	return result, nil
}

// ResumeStop creates a new PacketCapturesStopPoller from the specified resume token.
// token - The value must come from a previous call to PacketCapturesStopPoller.ResumeToken().
func (client *PacketCapturesClient) ResumeStop(ctx context.Context, token string) (PacketCapturesStopPollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("PacketCapturesClient.Stop", token, client.con.Pipeline(), client.stopHandleError)
	if err != nil {
		return PacketCapturesStopPollerResponse{}, err
	}
	poller := &PacketCapturesStopPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return PacketCapturesStopPollerResponse{}, err
	}
	result := PacketCapturesStopPollerResponse{
		Poller:      poller,
		RawResponse: resp,
	}
	result.Poller = poller
	return result, nil
}

// Stop - Stops a specified packet capture session.
// If the operation fails it returns the *ErrorResponse error type.
func (client *PacketCapturesClient) stop(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string, options *PacketCapturesBeginStopOptions) (*azcore.Response, error) {
	req, err := client.stopCreateRequest(ctx, resourceGroupName, networkWatcherName, packetCaptureName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.stopHandleError(resp)
	}
	return resp, nil
}

// stopCreateRequest creates the Stop request.
func (client *PacketCapturesClient) stopCreateRequest(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string, options *PacketCapturesBeginStopOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/packetCaptures/{packetCaptureName}/stop"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkWatcherName == "" {
		return nil, errors.New("parameter networkWatcherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkWatcherName}", url.PathEscape(networkWatcherName))
	if packetCaptureName == "" {
		return nil, errors.New("parameter packetCaptureName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packetCaptureName}", url.PathEscape(packetCaptureName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
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

// stopHandleError handles the Stop error response.
func (client *PacketCapturesClient) stopHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}
