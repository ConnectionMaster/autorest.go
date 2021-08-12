// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azblob

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type serviceClient struct {
	con *connection
}

// FilterBlobs - The Filter Blobs operation enables callers to list blobs across all containers whose tags match a given search expression. Filter blobs
// searches across all containers within a storage account but can
// be scoped within the expression to a single container.
// If the operation fails it returns the *StorageError error type.
func (client *serviceClient) FilterBlobs(ctx context.Context, options *ServiceFilterBlobsOptions) (ServiceFilterBlobsResponse, error) {
	req, err := client.filterBlobsCreateRequest(ctx, options)
	if err != nil {
		return ServiceFilterBlobsResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ServiceFilterBlobsResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ServiceFilterBlobsResponse{}, client.filterBlobsHandleError(resp)
	}
	return client.filterBlobsHandleResponse(resp)
}

// filterBlobsCreateRequest creates the FilterBlobs request.
func (client *serviceClient) filterBlobsCreateRequest(ctx context.Context, options *ServiceFilterBlobsOptions) (*azcore.Request, error) {
	req, err := azcore.NewRequest(ctx, http.MethodGet, client.con.Endpoint())
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("comp", "blobs")
	if options != nil && options.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	if options != nil && options.Where != nil {
		reqQP.Set("where", *options.Where)
	}
	if options != nil && options.Marker != nil {
		reqQP.Set("marker", *options.Marker)
	}
	if options != nil && options.Maxresults != nil {
		reqQP.Set("maxresults", strconv.FormatInt(int64(*options.Maxresults), 10))
	}
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("x-ms-version", "2020-06-12")
	if options != nil && options.RequestID != nil {
		req.Header.Set("x-ms-client-request-id", *options.RequestID)
	}
	req.Header.Set("Accept", "application/xml")
	return req, nil
}

// filterBlobsHandleResponse handles the FilterBlobs response.
func (client *serviceClient) filterBlobsHandleResponse(resp *azcore.Response) (ServiceFilterBlobsResponse, error) {
	result := ServiceFilterBlobsResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return ServiceFilterBlobsResponse{}, err
		}
		result.Date = &date
	}
	if err := resp.UnmarshalAsXML(&result.FilterBlobSegment); err != nil {
		return ServiceFilterBlobsResponse{}, err
	}
	return result, nil
}

// filterBlobsHandleError handles the FilterBlobs error response.
func (client *serviceClient) filterBlobsHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := StorageError{raw: string(body)}
	if err := resp.UnmarshalAsXML(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// GetAccountInfo - Returns the sku name and account kind
// If the operation fails it returns the *StorageError error type.
func (client *serviceClient) GetAccountInfo(ctx context.Context, options *ServiceGetAccountInfoOptions) (ServiceGetAccountInfoResponse, error) {
	req, err := client.getAccountInfoCreateRequest(ctx, options)
	if err != nil {
		return ServiceGetAccountInfoResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ServiceGetAccountInfoResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ServiceGetAccountInfoResponse{}, client.getAccountInfoHandleError(resp)
	}
	return client.getAccountInfoHandleResponse(resp)
}

// getAccountInfoCreateRequest creates the GetAccountInfo request.
func (client *serviceClient) getAccountInfoCreateRequest(ctx context.Context, options *ServiceGetAccountInfoOptions) (*azcore.Request, error) {
	req, err := azcore.NewRequest(ctx, http.MethodGet, client.con.Endpoint())
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("restype", "account")
	reqQP.Set("comp", "properties")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("x-ms-version", "2020-06-12")
	req.Header.Set("Accept", "application/xml")
	return req, nil
}

// getAccountInfoHandleResponse handles the GetAccountInfo response.
func (client *serviceClient) getAccountInfoHandleResponse(resp *azcore.Response) (ServiceGetAccountInfoResponse, error) {
	result := ServiceGetAccountInfoResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return ServiceGetAccountInfoResponse{}, err
		}
		result.Date = &date
	}
	if val := resp.Header.Get("x-ms-sku-name"); val != "" {
		result.SKUName = (*SKUName)(&val)
	}
	if val := resp.Header.Get("x-ms-account-kind"); val != "" {
		result.AccountKind = (*AccountKind)(&val)
	}
	if val := resp.Header.Get("x-ms-is-hns-enabled"); val != "" {
		isHierarchicalNamespaceEnabled, err := strconv.ParseBool(val)
		if err != nil {
			return ServiceGetAccountInfoResponse{}, err
		}
		result.IsHierarchicalNamespaceEnabled = &isHierarchicalNamespaceEnabled
	}
	return result, nil
}

// getAccountInfoHandleError handles the GetAccountInfo error response.
func (client *serviceClient) getAccountInfoHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := StorageError{raw: string(body)}
	if err := resp.UnmarshalAsXML(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// GetProperties - gets the properties of a storage account's Blob service, including properties for Storage Analytics and CORS (Cross-Origin Resource Sharing)
// rules.
// If the operation fails it returns the *StorageError error type.
func (client *serviceClient) GetProperties(ctx context.Context, options *ServiceGetPropertiesOptions) (ServiceGetPropertiesResponse, error) {
	req, err := client.getPropertiesCreateRequest(ctx, options)
	if err != nil {
		return ServiceGetPropertiesResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ServiceGetPropertiesResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ServiceGetPropertiesResponse{}, client.getPropertiesHandleError(resp)
	}
	return client.getPropertiesHandleResponse(resp)
}

// getPropertiesCreateRequest creates the GetProperties request.
func (client *serviceClient) getPropertiesCreateRequest(ctx context.Context, options *ServiceGetPropertiesOptions) (*azcore.Request, error) {
	req, err := azcore.NewRequest(ctx, http.MethodGet, client.con.Endpoint())
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("restype", "service")
	reqQP.Set("comp", "properties")
	if options != nil && options.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("x-ms-version", "2020-06-12")
	if options != nil && options.RequestID != nil {
		req.Header.Set("x-ms-client-request-id", *options.RequestID)
	}
	req.Header.Set("Accept", "application/xml")
	return req, nil
}

// getPropertiesHandleResponse handles the GetProperties response.
func (client *serviceClient) getPropertiesHandleResponse(resp *azcore.Response) (ServiceGetPropertiesResponse, error) {
	result := ServiceGetPropertiesResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if err := resp.UnmarshalAsXML(&result.StorageServiceProperties); err != nil {
		return ServiceGetPropertiesResponse{}, err
	}
	return result, nil
}

// getPropertiesHandleError handles the GetProperties error response.
func (client *serviceClient) getPropertiesHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := StorageError{raw: string(body)}
	if err := resp.UnmarshalAsXML(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// GetStatistics - Retrieves statistics related to replication for the Blob service. It is only available on the secondary location endpoint when read-access
// geo-redundant replication is enabled for the storage account.
// If the operation fails it returns the *StorageError error type.
func (client *serviceClient) GetStatistics(ctx context.Context, options *ServiceGetStatisticsOptions) (ServiceGetStatisticsResponse, error) {
	req, err := client.getStatisticsCreateRequest(ctx, options)
	if err != nil {
		return ServiceGetStatisticsResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ServiceGetStatisticsResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ServiceGetStatisticsResponse{}, client.getStatisticsHandleError(resp)
	}
	return client.getStatisticsHandleResponse(resp)
}

// getStatisticsCreateRequest creates the GetStatistics request.
func (client *serviceClient) getStatisticsCreateRequest(ctx context.Context, options *ServiceGetStatisticsOptions) (*azcore.Request, error) {
	req, err := azcore.NewRequest(ctx, http.MethodGet, client.con.Endpoint())
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("restype", "service")
	reqQP.Set("comp", "stats")
	if options != nil && options.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("x-ms-version", "2020-06-12")
	if options != nil && options.RequestID != nil {
		req.Header.Set("x-ms-client-request-id", *options.RequestID)
	}
	req.Header.Set("Accept", "application/xml")
	return req, nil
}

// getStatisticsHandleResponse handles the GetStatistics response.
func (client *serviceClient) getStatisticsHandleResponse(resp *azcore.Response) (ServiceGetStatisticsResponse, error) {
	result := ServiceGetStatisticsResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return ServiceGetStatisticsResponse{}, err
		}
		result.Date = &date
	}
	if err := resp.UnmarshalAsXML(&result.StorageServiceStats); err != nil {
		return ServiceGetStatisticsResponse{}, err
	}
	return result, nil
}

// getStatisticsHandleError handles the GetStatistics error response.
func (client *serviceClient) getStatisticsHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := StorageError{raw: string(body)}
	if err := resp.UnmarshalAsXML(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// GetUserDelegationKey - Retrieves a user delegation key for the Blob service. This is only a valid operation when using bearer token authentication.
// If the operation fails it returns the *StorageError error type.
func (client *serviceClient) GetUserDelegationKey(ctx context.Context, keyInfo KeyInfo, options *ServiceGetUserDelegationKeyOptions) (ServiceGetUserDelegationKeyResponse, error) {
	req, err := client.getUserDelegationKeyCreateRequest(ctx, keyInfo, options)
	if err != nil {
		return ServiceGetUserDelegationKeyResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ServiceGetUserDelegationKeyResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ServiceGetUserDelegationKeyResponse{}, client.getUserDelegationKeyHandleError(resp)
	}
	return client.getUserDelegationKeyHandleResponse(resp)
}

// getUserDelegationKeyCreateRequest creates the GetUserDelegationKey request.
func (client *serviceClient) getUserDelegationKeyCreateRequest(ctx context.Context, keyInfo KeyInfo, options *ServiceGetUserDelegationKeyOptions) (*azcore.Request, error) {
	req, err := azcore.NewRequest(ctx, http.MethodPost, client.con.Endpoint())
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("restype", "service")
	reqQP.Set("comp", "userdelegationkey")
	if options != nil && options.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("x-ms-version", "2020-06-12")
	if options != nil && options.RequestID != nil {
		req.Header.Set("x-ms-client-request-id", *options.RequestID)
	}
	req.Header.Set("Accept", "application/xml")
	return req, req.MarshalAsXML(keyInfo)
}

// getUserDelegationKeyHandleResponse handles the GetUserDelegationKey response.
func (client *serviceClient) getUserDelegationKeyHandleResponse(resp *azcore.Response) (ServiceGetUserDelegationKeyResponse, error) {
	result := ServiceGetUserDelegationKeyResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return ServiceGetUserDelegationKeyResponse{}, err
		}
		result.Date = &date
	}
	if err := resp.UnmarshalAsXML(&result.UserDelegationKey); err != nil {
		return ServiceGetUserDelegationKeyResponse{}, err
	}
	return result, nil
}

// getUserDelegationKeyHandleError handles the GetUserDelegationKey error response.
func (client *serviceClient) getUserDelegationKeyHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := StorageError{raw: string(body)}
	if err := resp.UnmarshalAsXML(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// ListContainersSegment - The List Containers Segment operation returns a list of the containers under the specified account
// If the operation fails it returns the *StorageError error type.
func (client *serviceClient) ListContainersSegment(options *ServiceListContainersSegmentOptions) *ServiceListContainersSegmentPager {
	return &ServiceListContainersSegmentPager{
		client: client,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listContainersSegmentCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp ServiceListContainersSegmentResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ListContainersSegmentResponse.NextMarker)
		},
	}
}

// listContainersSegmentCreateRequest creates the ListContainersSegment request.
func (client *serviceClient) listContainersSegmentCreateRequest(ctx context.Context, options *ServiceListContainersSegmentOptions) (*azcore.Request, error) {
	req, err := azcore.NewRequest(ctx, http.MethodGet, client.con.Endpoint())
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("comp", "list")
	if options != nil && options.Prefix != nil {
		reqQP.Set("prefix", *options.Prefix)
	}
	if options != nil && options.Marker != nil {
		reqQP.Set("marker", *options.Marker)
	}
	if options != nil && options.Maxresults != nil {
		reqQP.Set("maxresults", strconv.FormatInt(int64(*options.Maxresults), 10))
	}
	if options != nil && options.Include != nil {
		reqQP.Set("include", strings.Join(strings.Fields(strings.Trim(fmt.Sprint(options.Include), "[]")), ","))
	}
	if options != nil && options.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("x-ms-version", "2020-06-12")
	if options != nil && options.RequestID != nil {
		req.Header.Set("x-ms-client-request-id", *options.RequestID)
	}
	req.Header.Set("Accept", "application/xml")
	return req, nil
}

// listContainersSegmentHandleResponse handles the ListContainersSegment response.
func (client *serviceClient) listContainersSegmentHandleResponse(resp *azcore.Response) (ServiceListContainersSegmentResponse, error) {
	result := ServiceListContainersSegmentResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if err := resp.UnmarshalAsXML(&result.ListContainersSegmentResponse); err != nil {
		return ServiceListContainersSegmentResponse{}, err
	}
	return result, nil
}

// listContainersSegmentHandleError handles the ListContainersSegment error response.
func (client *serviceClient) listContainersSegmentHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := StorageError{raw: string(body)}
	if err := resp.UnmarshalAsXML(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// SetProperties - Sets properties for a storage account's Blob service endpoint, including properties for Storage Analytics and CORS (Cross-Origin Resource
// Sharing) rules
// If the operation fails it returns the *StorageError error type.
func (client *serviceClient) SetProperties(ctx context.Context, storageServiceProperties StorageServiceProperties, options *ServiceSetPropertiesOptions) (ServiceSetPropertiesResponse, error) {
	req, err := client.setPropertiesCreateRequest(ctx, storageServiceProperties, options)
	if err != nil {
		return ServiceSetPropertiesResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ServiceSetPropertiesResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusAccepted) {
		return ServiceSetPropertiesResponse{}, client.setPropertiesHandleError(resp)
	}
	return client.setPropertiesHandleResponse(resp)
}

// setPropertiesCreateRequest creates the SetProperties request.
func (client *serviceClient) setPropertiesCreateRequest(ctx context.Context, storageServiceProperties StorageServiceProperties, options *ServiceSetPropertiesOptions) (*azcore.Request, error) {
	req, err := azcore.NewRequest(ctx, http.MethodPut, client.con.Endpoint())
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("restype", "service")
	reqQP.Set("comp", "properties")
	if options != nil && options.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("x-ms-version", "2020-06-12")
	if options != nil && options.RequestID != nil {
		req.Header.Set("x-ms-client-request-id", *options.RequestID)
	}
	req.Header.Set("Accept", "application/xml")
	return req, req.MarshalAsXML(storageServiceProperties)
}

// setPropertiesHandleResponse handles the SetProperties response.
func (client *serviceClient) setPropertiesHandleResponse(resp *azcore.Response) (ServiceSetPropertiesResponse, error) {
	result := ServiceSetPropertiesResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	return result, nil
}

// setPropertiesHandleError handles the SetProperties error response.
func (client *serviceClient) setPropertiesHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := StorageError{raw: string(body)}
	if err := resp.UnmarshalAsXML(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// SubmitBatch - The Batch operation allows multiple API calls to be embedded into a single HTTP request.
// If the operation fails it returns the *StorageError error type.
func (client *serviceClient) SubmitBatch(ctx context.Context, contentLength int64, multipartContentType string, body azcore.ReadSeekCloser, options *ServiceSubmitBatchOptions) (ServiceSubmitBatchResponse, error) {
	req, err := client.submitBatchCreateRequest(ctx, contentLength, multipartContentType, body, options)
	if err != nil {
		return ServiceSubmitBatchResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ServiceSubmitBatchResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ServiceSubmitBatchResponse{}, client.submitBatchHandleError(resp)
	}
	return client.submitBatchHandleResponse(resp)
}

// submitBatchCreateRequest creates the SubmitBatch request.
func (client *serviceClient) submitBatchCreateRequest(ctx context.Context, contentLength int64, multipartContentType string, body azcore.ReadSeekCloser, options *ServiceSubmitBatchOptions) (*azcore.Request, error) {
	req, err := azcore.NewRequest(ctx, http.MethodPost, client.con.Endpoint())
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("comp", "batch")
	if options != nil && options.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	req.URL.RawQuery = reqQP.Encode()
	req.SkipBodyDownload()
	req.Header.Set("Content-Length", strconv.FormatInt(contentLength, 10))
	req.Header.Set("Content-Type", multipartContentType)
	req.Header.Set("x-ms-version", "2020-06-12")
	if options != nil && options.RequestID != nil {
		req.Header.Set("x-ms-client-request-id", *options.RequestID)
	}
	req.Header.Set("Accept", "application/xml")
	return req, req.MarshalAsXML(body)
}

// submitBatchHandleResponse handles the SubmitBatch response.
func (client *serviceClient) submitBatchHandleResponse(resp *azcore.Response) (ServiceSubmitBatchResponse, error) {
	result := ServiceSubmitBatchResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("Content-Type"); val != "" {
		result.ContentType = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	return result, nil
}

// submitBatchHandleError handles the SubmitBatch error response.
func (client *serviceClient) submitBatchHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := StorageError{raw: string(body)}
	if err := resp.UnmarshalAsXML(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}
