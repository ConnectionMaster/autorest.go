//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package azartifacts

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

type bigDataPoolsClient struct {
	endpoint string
	pl       runtime.Pipeline
}

// newBigDataPoolsClient creates a new instance of bigDataPoolsClient with the specified values.
// endpoint - The workspace development endpoint, for example https://myworkspace.dev.azuresynapse.net.
// pl - the pipeline used for sending requests and handling responses.
func newBigDataPoolsClient(endpoint string, pl runtime.Pipeline) *bigDataPoolsClient {
	client := &bigDataPoolsClient{
		endpoint: endpoint,
		pl:       pl,
	}
	return client
}

// Get - Get Big Data Pool
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-06-01-preview
// bigDataPoolName - The Big Data Pool name
// options - bigDataPoolsClientGetOptions contains the optional parameters for the bigDataPoolsClient.Get method.
func (client *bigDataPoolsClient) Get(ctx context.Context, bigDataPoolName string, options *bigDataPoolsClientGetOptions) (BigDataPoolsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, bigDataPoolName, options)
	if err != nil {
		return BigDataPoolsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BigDataPoolsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BigDataPoolsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *bigDataPoolsClient) getCreateRequest(ctx context.Context, bigDataPoolName string, options *bigDataPoolsClientGetOptions) (*policy.Request, error) {
	urlPath := "/bigDataPools/{bigDataPoolName}"
	if bigDataPoolName == "" {
		return nil, errors.New("parameter bigDataPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{bigDataPoolName}", url.PathEscape(bigDataPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *bigDataPoolsClient) getHandleResponse(resp *http.Response) (BigDataPoolsClientGetResponse, error) {
	result := BigDataPoolsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BigDataPoolResourceInfo); err != nil {
		return BigDataPoolsClientGetResponse{}, err
	}
	return result, nil
}

// List - List Big Data Pools
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-06-01-preview
// options - bigDataPoolsClientListOptions contains the optional parameters for the bigDataPoolsClient.List method.
func (client *bigDataPoolsClient) List(ctx context.Context, options *bigDataPoolsClientListOptions) (BigDataPoolsClientListResponse, error) {
	req, err := client.listCreateRequest(ctx, options)
	if err != nil {
		return BigDataPoolsClientListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BigDataPoolsClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BigDataPoolsClientListResponse{}, runtime.NewResponseError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *bigDataPoolsClient) listCreateRequest(ctx context.Context, options *bigDataPoolsClientListOptions) (*policy.Request, error) {
	urlPath := "/bigDataPools"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *bigDataPoolsClient) listHandleResponse(resp *http.Response) (BigDataPoolsClientListResponse, error) {
	result := BigDataPoolsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BigDataPoolResourceInfoListResult); err != nil {
		return BigDataPoolsClientListResponse{}, err
	}
	return result, nil
}
