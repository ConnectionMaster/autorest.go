// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package lrogroup

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// LRORetrysClient contains the methods for the LRORetrys group.
// Don't use this type directly, use NewLRORetrysClient() instead.
type LRORetrysClient struct {
	con *Connection
}

// NewLRORetrysClient creates a new instance of LRORetrysClient with the specified values.
func NewLRORetrysClient(con *Connection) *LRORetrysClient {
	return &LRORetrysClient{con: con}
}

// BeginDelete202Retry200 - Long running delete request, service returns a 500, then a 202 to the initial request. Polls return this value until the last
// poll returns a ‘200’ with ProvisioningState=’Succeeded’
// If the operation fails it returns the *CloudError error type.
func (client *LRORetrysClient) BeginDelete202Retry200(ctx context.Context, options *LRORetrysBeginDelete202Retry200Options) (LRORetrysDelete202Retry200PollerResponse, error) {
	resp, err := client.delete202Retry200(ctx, options)
	if err != nil {
		return LRORetrysDelete202Retry200PollerResponse{}, err
	}
	result := LRORetrysDelete202Retry200PollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("LRORetrysClient.Delete202Retry200", "", resp, client.con.Pipeline(), client.delete202Retry200HandleError)
	if err != nil {
		return LRORetrysDelete202Retry200PollerResponse{}, err
	}
	result.Poller = &LRORetrysDelete202Retry200Poller{
		pt: pt,
	}
	return result, nil
}

// ResumeDelete202Retry200 creates a new LRORetrysDelete202Retry200Poller from the specified resume token.
// token - The value must come from a previous call to LRORetrysDelete202Retry200Poller.ResumeToken().
func (client *LRORetrysClient) ResumeDelete202Retry200(ctx context.Context, token string) (LRORetrysDelete202Retry200PollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("LRORetrysClient.Delete202Retry200", token, client.con.Pipeline(), client.delete202Retry200HandleError)
	if err != nil {
		return LRORetrysDelete202Retry200PollerResponse{}, err
	}
	poller := &LRORetrysDelete202Retry200Poller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return LRORetrysDelete202Retry200PollerResponse{}, err
	}
	result := LRORetrysDelete202Retry200PollerResponse{
		Poller:      poller,
		RawResponse: resp,
	}
	result.Poller = poller
	return result, nil
}

// Delete202Retry200 - Long running delete request, service returns a 500, then a 202 to the initial request. Polls return this value until the last poll
// returns a ‘200’ with ProvisioningState=’Succeeded’
// If the operation fails it returns the *CloudError error type.
func (client *LRORetrysClient) delete202Retry200(ctx context.Context, options *LRORetrysBeginDelete202Retry200Options) (*azcore.Response, error) {
	req, err := client.delete202Retry200CreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusAccepted) {
		return nil, client.delete202Retry200HandleError(resp)
	}
	return resp, nil
}

// delete202Retry200CreateRequest creates the Delete202Retry200 request.
func (client *LRORetrysClient) delete202Retry200CreateRequest(ctx context.Context, options *LRORetrysBeginDelete202Retry200Options) (*azcore.Request, error) {
	urlPath := "/lro/retryerror/delete/202/retry/200"
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// delete202Retry200HandleError handles the Delete202Retry200 error response.
func (client *LRORetrysClient) delete202Retry200HandleError(resp *azcore.Response) error {
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

// BeginDeleteAsyncRelativeRetrySucceeded - Long running delete request, service returns a 500, then a 202 to the initial request. Poll the endpoint indicated
// in the Azure-AsyncOperation header for operation status
// If the operation fails it returns the *CloudError error type.
func (client *LRORetrysClient) BeginDeleteAsyncRelativeRetrySucceeded(ctx context.Context, options *LRORetrysBeginDeleteAsyncRelativeRetrySucceededOptions) (LRORetrysDeleteAsyncRelativeRetrySucceededPollerResponse, error) {
	resp, err := client.deleteAsyncRelativeRetrySucceeded(ctx, options)
	if err != nil {
		return LRORetrysDeleteAsyncRelativeRetrySucceededPollerResponse{}, err
	}
	result := LRORetrysDeleteAsyncRelativeRetrySucceededPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("LRORetrysClient.DeleteAsyncRelativeRetrySucceeded", "", resp, client.con.Pipeline(), client.deleteAsyncRelativeRetrySucceededHandleError)
	if err != nil {
		return LRORetrysDeleteAsyncRelativeRetrySucceededPollerResponse{}, err
	}
	result.Poller = &LRORetrysDeleteAsyncRelativeRetrySucceededPoller{
		pt: pt,
	}
	return result, nil
}

// ResumeDeleteAsyncRelativeRetrySucceeded creates a new LRORetrysDeleteAsyncRelativeRetrySucceededPoller from the specified resume token.
// token - The value must come from a previous call to LRORetrysDeleteAsyncRelativeRetrySucceededPoller.ResumeToken().
func (client *LRORetrysClient) ResumeDeleteAsyncRelativeRetrySucceeded(ctx context.Context, token string) (LRORetrysDeleteAsyncRelativeRetrySucceededPollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("LRORetrysClient.DeleteAsyncRelativeRetrySucceeded", token, client.con.Pipeline(), client.deleteAsyncRelativeRetrySucceededHandleError)
	if err != nil {
		return LRORetrysDeleteAsyncRelativeRetrySucceededPollerResponse{}, err
	}
	poller := &LRORetrysDeleteAsyncRelativeRetrySucceededPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return LRORetrysDeleteAsyncRelativeRetrySucceededPollerResponse{}, err
	}
	result := LRORetrysDeleteAsyncRelativeRetrySucceededPollerResponse{
		Poller:      poller,
		RawResponse: resp,
	}
	result.Poller = poller
	return result, nil
}

// DeleteAsyncRelativeRetrySucceeded - Long running delete request, service returns a 500, then a 202 to the initial request. Poll the endpoint indicated
// in the Azure-AsyncOperation header for operation status
// If the operation fails it returns the *CloudError error type.
func (client *LRORetrysClient) deleteAsyncRelativeRetrySucceeded(ctx context.Context, options *LRORetrysBeginDeleteAsyncRelativeRetrySucceededOptions) (*azcore.Response, error) {
	req, err := client.deleteAsyncRelativeRetrySucceededCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusAccepted) {
		return nil, client.deleteAsyncRelativeRetrySucceededHandleError(resp)
	}
	return resp, nil
}

// deleteAsyncRelativeRetrySucceededCreateRequest creates the DeleteAsyncRelativeRetrySucceeded request.
func (client *LRORetrysClient) deleteAsyncRelativeRetrySucceededCreateRequest(ctx context.Context, options *LRORetrysBeginDeleteAsyncRelativeRetrySucceededOptions) (*azcore.Request, error) {
	urlPath := "/lro/retryerror/deleteasync/retry/succeeded"
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteAsyncRelativeRetrySucceededHandleError handles the DeleteAsyncRelativeRetrySucceeded error response.
func (client *LRORetrysClient) deleteAsyncRelativeRetrySucceededHandleError(resp *azcore.Response) error {
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

// BeginDeleteProvisioning202Accepted200Succeeded - Long running delete request, service returns a 500, then a 202 to the initial request, with an entity
// that contains ProvisioningState=’Accepted’. Polls return this value until the last poll returns a
// ‘200’ with ProvisioningState=’Succeeded’
// If the operation fails it returns the *CloudError error type.
func (client *LRORetrysClient) BeginDeleteProvisioning202Accepted200Succeeded(ctx context.Context, options *LRORetrysBeginDeleteProvisioning202Accepted200SucceededOptions) (LRORetrysDeleteProvisioning202Accepted200SucceededPollerResponse, error) {
	resp, err := client.deleteProvisioning202Accepted200Succeeded(ctx, options)
	if err != nil {
		return LRORetrysDeleteProvisioning202Accepted200SucceededPollerResponse{}, err
	}
	result := LRORetrysDeleteProvisioning202Accepted200SucceededPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("LRORetrysClient.DeleteProvisioning202Accepted200Succeeded", "", resp, client.con.Pipeline(), client.deleteProvisioning202Accepted200SucceededHandleError)
	if err != nil {
		return LRORetrysDeleteProvisioning202Accepted200SucceededPollerResponse{}, err
	}
	result.Poller = &LRORetrysDeleteProvisioning202Accepted200SucceededPoller{
		pt: pt,
	}
	return result, nil
}

// ResumeDeleteProvisioning202Accepted200Succeeded creates a new LRORetrysDeleteProvisioning202Accepted200SucceededPoller from the specified resume token.
// token - The value must come from a previous call to LRORetrysDeleteProvisioning202Accepted200SucceededPoller.ResumeToken().
func (client *LRORetrysClient) ResumeDeleteProvisioning202Accepted200Succeeded(ctx context.Context, token string) (LRORetrysDeleteProvisioning202Accepted200SucceededPollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("LRORetrysClient.DeleteProvisioning202Accepted200Succeeded", token, client.con.Pipeline(), client.deleteProvisioning202Accepted200SucceededHandleError)
	if err != nil {
		return LRORetrysDeleteProvisioning202Accepted200SucceededPollerResponse{}, err
	}
	poller := &LRORetrysDeleteProvisioning202Accepted200SucceededPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return LRORetrysDeleteProvisioning202Accepted200SucceededPollerResponse{}, err
	}
	result := LRORetrysDeleteProvisioning202Accepted200SucceededPollerResponse{
		Poller:      poller,
		RawResponse: resp,
	}
	result.Poller = poller
	return result, nil
}

// DeleteProvisioning202Accepted200Succeeded - Long running delete request, service returns a 500, then a 202 to the initial request, with an entity that
// contains ProvisioningState=’Accepted’. Polls return this value until the last poll returns a
// ‘200’ with ProvisioningState=’Succeeded’
// If the operation fails it returns the *CloudError error type.
func (client *LRORetrysClient) deleteProvisioning202Accepted200Succeeded(ctx context.Context, options *LRORetrysBeginDeleteProvisioning202Accepted200SucceededOptions) (*azcore.Response, error) {
	req, err := client.deleteProvisioning202Accepted200SucceededCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.deleteProvisioning202Accepted200SucceededHandleError(resp)
	}
	return resp, nil
}

// deleteProvisioning202Accepted200SucceededCreateRequest creates the DeleteProvisioning202Accepted200Succeeded request.
func (client *LRORetrysClient) deleteProvisioning202Accepted200SucceededCreateRequest(ctx context.Context, options *LRORetrysBeginDeleteProvisioning202Accepted200SucceededOptions) (*azcore.Request, error) {
	urlPath := "/lro/retryerror/delete/provisioning/202/accepted/200/succeeded"
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteProvisioning202Accepted200SucceededHandleError handles the DeleteProvisioning202Accepted200Succeeded error response.
func (client *LRORetrysClient) deleteProvisioning202Accepted200SucceededHandleError(resp *azcore.Response) error {
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

// BeginPost202Retry200 - Long running post request, service returns a 500, then a 202 to the initial request, with 'Location' and 'Retry-After' headers,
// Polls return a 200 with a response body after success
// If the operation fails it returns the *CloudError error type.
func (client *LRORetrysClient) BeginPost202Retry200(ctx context.Context, options *LRORetrysBeginPost202Retry200Options) (LRORetrysPost202Retry200PollerResponse, error) {
	resp, err := client.post202Retry200(ctx, options)
	if err != nil {
		return LRORetrysPost202Retry200PollerResponse{}, err
	}
	result := LRORetrysPost202Retry200PollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("LRORetrysClient.Post202Retry200", "", resp, client.con.Pipeline(), client.post202Retry200HandleError)
	if err != nil {
		return LRORetrysPost202Retry200PollerResponse{}, err
	}
	result.Poller = &LRORetrysPost202Retry200Poller{
		pt: pt,
	}
	return result, nil
}

// ResumePost202Retry200 creates a new LRORetrysPost202Retry200Poller from the specified resume token.
// token - The value must come from a previous call to LRORetrysPost202Retry200Poller.ResumeToken().
func (client *LRORetrysClient) ResumePost202Retry200(ctx context.Context, token string) (LRORetrysPost202Retry200PollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("LRORetrysClient.Post202Retry200", token, client.con.Pipeline(), client.post202Retry200HandleError)
	if err != nil {
		return LRORetrysPost202Retry200PollerResponse{}, err
	}
	poller := &LRORetrysPost202Retry200Poller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return LRORetrysPost202Retry200PollerResponse{}, err
	}
	result := LRORetrysPost202Retry200PollerResponse{
		Poller:      poller,
		RawResponse: resp,
	}
	result.Poller = poller
	return result, nil
}

// Post202Retry200 - Long running post request, service returns a 500, then a 202 to the initial request, with 'Location' and 'Retry-After' headers, Polls
// return a 200 with a response body after success
// If the operation fails it returns the *CloudError error type.
func (client *LRORetrysClient) post202Retry200(ctx context.Context, options *LRORetrysBeginPost202Retry200Options) (*azcore.Response, error) {
	req, err := client.post202Retry200CreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusAccepted) {
		return nil, client.post202Retry200HandleError(resp)
	}
	return resp, nil
}

// post202Retry200CreateRequest creates the Post202Retry200 request.
func (client *LRORetrysClient) post202Retry200CreateRequest(ctx context.Context, options *LRORetrysBeginPost202Retry200Options) (*azcore.Request, error) {
	urlPath := "/lro/retryerror/post/202/retry/200"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	if options != nil && options.Product != nil {
		return req, req.MarshalAsJSON(*options.Product)
	}
	return req, nil
}

// post202Retry200HandleError handles the Post202Retry200 error response.
func (client *LRORetrysClient) post202Retry200HandleError(resp *azcore.Response) error {
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

// BeginPostAsyncRelativeRetrySucceeded - Long running post request, service returns a 500, then a 202 to the initial request, with an entity that contains
// ProvisioningState=’Creating’. Poll the endpoint indicated in the Azure-AsyncOperation
// header for operation status
// If the operation fails it returns the *CloudError error type.
func (client *LRORetrysClient) BeginPostAsyncRelativeRetrySucceeded(ctx context.Context, options *LRORetrysBeginPostAsyncRelativeRetrySucceededOptions) (LRORetrysPostAsyncRelativeRetrySucceededPollerResponse, error) {
	resp, err := client.postAsyncRelativeRetrySucceeded(ctx, options)
	if err != nil {
		return LRORetrysPostAsyncRelativeRetrySucceededPollerResponse{}, err
	}
	result := LRORetrysPostAsyncRelativeRetrySucceededPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("LRORetrysClient.PostAsyncRelativeRetrySucceeded", "", resp, client.con.Pipeline(), client.postAsyncRelativeRetrySucceededHandleError)
	if err != nil {
		return LRORetrysPostAsyncRelativeRetrySucceededPollerResponse{}, err
	}
	result.Poller = &LRORetrysPostAsyncRelativeRetrySucceededPoller{
		pt: pt,
	}
	return result, nil
}

// ResumePostAsyncRelativeRetrySucceeded creates a new LRORetrysPostAsyncRelativeRetrySucceededPoller from the specified resume token.
// token - The value must come from a previous call to LRORetrysPostAsyncRelativeRetrySucceededPoller.ResumeToken().
func (client *LRORetrysClient) ResumePostAsyncRelativeRetrySucceeded(ctx context.Context, token string) (LRORetrysPostAsyncRelativeRetrySucceededPollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("LRORetrysClient.PostAsyncRelativeRetrySucceeded", token, client.con.Pipeline(), client.postAsyncRelativeRetrySucceededHandleError)
	if err != nil {
		return LRORetrysPostAsyncRelativeRetrySucceededPollerResponse{}, err
	}
	poller := &LRORetrysPostAsyncRelativeRetrySucceededPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return LRORetrysPostAsyncRelativeRetrySucceededPollerResponse{}, err
	}
	result := LRORetrysPostAsyncRelativeRetrySucceededPollerResponse{
		Poller:      poller,
		RawResponse: resp,
	}
	result.Poller = poller
	return result, nil
}

// PostAsyncRelativeRetrySucceeded - Long running post request, service returns a 500, then a 202 to the initial request, with an entity that contains ProvisioningState=’Creating’.
// Poll the endpoint indicated in the Azure-AsyncOperation
// header for operation status
// If the operation fails it returns the *CloudError error type.
func (client *LRORetrysClient) postAsyncRelativeRetrySucceeded(ctx context.Context, options *LRORetrysBeginPostAsyncRelativeRetrySucceededOptions) (*azcore.Response, error) {
	req, err := client.postAsyncRelativeRetrySucceededCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusAccepted) {
		return nil, client.postAsyncRelativeRetrySucceededHandleError(resp)
	}
	return resp, nil
}

// postAsyncRelativeRetrySucceededCreateRequest creates the PostAsyncRelativeRetrySucceeded request.
func (client *LRORetrysClient) postAsyncRelativeRetrySucceededCreateRequest(ctx context.Context, options *LRORetrysBeginPostAsyncRelativeRetrySucceededOptions) (*azcore.Request, error) {
	urlPath := "/lro/retryerror/postasync/retry/succeeded"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	if options != nil && options.Product != nil {
		return req, req.MarshalAsJSON(*options.Product)
	}
	return req, nil
}

// postAsyncRelativeRetrySucceededHandleError handles the PostAsyncRelativeRetrySucceeded error response.
func (client *LRORetrysClient) postAsyncRelativeRetrySucceededHandleError(resp *azcore.Response) error {
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

// BeginPut201CreatingSucceeded200 - Long running put request, service returns a 500, then a 201 to the initial request, with an entity that contains ProvisioningState=’Creating’.
// Polls return this value until the last poll returns a
// ‘200’ with ProvisioningState=’Succeeded’
// If the operation fails it returns the *CloudError error type.
func (client *LRORetrysClient) BeginPut201CreatingSucceeded200(ctx context.Context, options *LRORetrysBeginPut201CreatingSucceeded200Options) (LRORetrysPut201CreatingSucceeded200PollerResponse, error) {
	resp, err := client.put201CreatingSucceeded200(ctx, options)
	if err != nil {
		return LRORetrysPut201CreatingSucceeded200PollerResponse{}, err
	}
	result := LRORetrysPut201CreatingSucceeded200PollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("LRORetrysClient.Put201CreatingSucceeded200", "", resp, client.con.Pipeline(), client.put201CreatingSucceeded200HandleError)
	if err != nil {
		return LRORetrysPut201CreatingSucceeded200PollerResponse{}, err
	}
	result.Poller = &LRORetrysPut201CreatingSucceeded200Poller{
		pt: pt,
	}
	return result, nil
}

// ResumePut201CreatingSucceeded200 creates a new LRORetrysPut201CreatingSucceeded200Poller from the specified resume token.
// token - The value must come from a previous call to LRORetrysPut201CreatingSucceeded200Poller.ResumeToken().
func (client *LRORetrysClient) ResumePut201CreatingSucceeded200(ctx context.Context, token string) (LRORetrysPut201CreatingSucceeded200PollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("LRORetrysClient.Put201CreatingSucceeded200", token, client.con.Pipeline(), client.put201CreatingSucceeded200HandleError)
	if err != nil {
		return LRORetrysPut201CreatingSucceeded200PollerResponse{}, err
	}
	poller := &LRORetrysPut201CreatingSucceeded200Poller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return LRORetrysPut201CreatingSucceeded200PollerResponse{}, err
	}
	result := LRORetrysPut201CreatingSucceeded200PollerResponse{
		Poller:      poller,
		RawResponse: resp,
	}
	result.Poller = poller
	return result, nil
}

// Put201CreatingSucceeded200 - Long running put request, service returns a 500, then a 201 to the initial request, with an entity that contains ProvisioningState=’Creating’.
// Polls return this value until the last poll returns a
// ‘200’ with ProvisioningState=’Succeeded’
// If the operation fails it returns the *CloudError error type.
func (client *LRORetrysClient) put201CreatingSucceeded200(ctx context.Context, options *LRORetrysBeginPut201CreatingSucceeded200Options) (*azcore.Response, error) {
	req, err := client.put201CreatingSucceeded200CreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.put201CreatingSucceeded200HandleError(resp)
	}
	return resp, nil
}

// put201CreatingSucceeded200CreateRequest creates the Put201CreatingSucceeded200 request.
func (client *LRORetrysClient) put201CreatingSucceeded200CreateRequest(ctx context.Context, options *LRORetrysBeginPut201CreatingSucceeded200Options) (*azcore.Request, error) {
	urlPath := "/lro/retryerror/put/201/creating/succeeded/200"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	if options != nil && options.Product != nil {
		return req, req.MarshalAsJSON(*options.Product)
	}
	return req, nil
}

// put201CreatingSucceeded200HandleError handles the Put201CreatingSucceeded200 error response.
func (client *LRORetrysClient) put201CreatingSucceeded200HandleError(resp *azcore.Response) error {
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

// BeginPutAsyncRelativeRetrySucceeded - Long running put request, service returns a 500, then a 200 to the initial request, with an entity that contains
// ProvisioningState=’Creating’. Poll the endpoint indicated in the Azure-AsyncOperation
// header for operation status
// If the operation fails it returns the *CloudError error type.
func (client *LRORetrysClient) BeginPutAsyncRelativeRetrySucceeded(ctx context.Context, options *LRORetrysBeginPutAsyncRelativeRetrySucceededOptions) (LRORetrysPutAsyncRelativeRetrySucceededPollerResponse, error) {
	resp, err := client.putAsyncRelativeRetrySucceeded(ctx, options)
	if err != nil {
		return LRORetrysPutAsyncRelativeRetrySucceededPollerResponse{}, err
	}
	result := LRORetrysPutAsyncRelativeRetrySucceededPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("LRORetrysClient.PutAsyncRelativeRetrySucceeded", "", resp, client.con.Pipeline(), client.putAsyncRelativeRetrySucceededHandleError)
	if err != nil {
		return LRORetrysPutAsyncRelativeRetrySucceededPollerResponse{}, err
	}
	result.Poller = &LRORetrysPutAsyncRelativeRetrySucceededPoller{
		pt: pt,
	}
	return result, nil
}

// ResumePutAsyncRelativeRetrySucceeded creates a new LRORetrysPutAsyncRelativeRetrySucceededPoller from the specified resume token.
// token - The value must come from a previous call to LRORetrysPutAsyncRelativeRetrySucceededPoller.ResumeToken().
func (client *LRORetrysClient) ResumePutAsyncRelativeRetrySucceeded(ctx context.Context, token string) (LRORetrysPutAsyncRelativeRetrySucceededPollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("LRORetrysClient.PutAsyncRelativeRetrySucceeded", token, client.con.Pipeline(), client.putAsyncRelativeRetrySucceededHandleError)
	if err != nil {
		return LRORetrysPutAsyncRelativeRetrySucceededPollerResponse{}, err
	}
	poller := &LRORetrysPutAsyncRelativeRetrySucceededPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return LRORetrysPutAsyncRelativeRetrySucceededPollerResponse{}, err
	}
	result := LRORetrysPutAsyncRelativeRetrySucceededPollerResponse{
		Poller:      poller,
		RawResponse: resp,
	}
	result.Poller = poller
	return result, nil
}

// PutAsyncRelativeRetrySucceeded - Long running put request, service returns a 500, then a 200 to the initial request, with an entity that contains ProvisioningState=’Creating’.
// Poll the endpoint indicated in the Azure-AsyncOperation
// header for operation status
// If the operation fails it returns the *CloudError error type.
func (client *LRORetrysClient) putAsyncRelativeRetrySucceeded(ctx context.Context, options *LRORetrysBeginPutAsyncRelativeRetrySucceededOptions) (*azcore.Response, error) {
	req, err := client.putAsyncRelativeRetrySucceededCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.putAsyncRelativeRetrySucceededHandleError(resp)
	}
	return resp, nil
}

// putAsyncRelativeRetrySucceededCreateRequest creates the PutAsyncRelativeRetrySucceeded request.
func (client *LRORetrysClient) putAsyncRelativeRetrySucceededCreateRequest(ctx context.Context, options *LRORetrysBeginPutAsyncRelativeRetrySucceededOptions) (*azcore.Request, error) {
	urlPath := "/lro/retryerror/putasync/retry/succeeded"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	if options != nil && options.Product != nil {
		return req, req.MarshalAsJSON(*options.Product)
	}
	return req, nil
}

// putAsyncRelativeRetrySucceededHandleError handles the PutAsyncRelativeRetrySucceeded error response.
func (client *LRORetrysClient) putAsyncRelativeRetrySucceededHandleError(resp *azcore.Response) error {
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
