//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package azurespecialsgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// APIVersionLocalClient contains the methods for the APIVersionLocal group.
// Don't use this type directly, use NewAPIVersionLocalClient() instead.
type APIVersionLocalClient struct {
	pl runtime.Pipeline
}

// NewAPIVersionLocalClient creates a new instance of APIVersionLocalClient with the specified values.
// pl - the pipeline used for sending requests and handling responses.
func NewAPIVersionLocalClient(pl runtime.Pipeline) *APIVersionLocalClient {
	client := &APIVersionLocalClient{
		pl: pl,
	}
	return client
}

// GetMethodLocalNull - Get method with api-version modeled in the method. pass in api-version = null to succeed
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2015-07-01-preview
// options - APIVersionLocalClientGetMethodLocalNullOptions contains the optional parameters for the APIVersionLocalClient.GetMethodLocalNull
// method.
func (client *APIVersionLocalClient) GetMethodLocalNull(ctx context.Context, options *APIVersionLocalClientGetMethodLocalNullOptions) (APIVersionLocalClientGetMethodLocalNullResponse, error) {
	req, err := client.getMethodLocalNullCreateRequest(ctx, options)
	if err != nil {
		return APIVersionLocalClientGetMethodLocalNullResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return APIVersionLocalClientGetMethodLocalNullResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return APIVersionLocalClientGetMethodLocalNullResponse{}, runtime.NewResponseError(resp)
	}
	return APIVersionLocalClientGetMethodLocalNullResponse{}, nil
}

// getMethodLocalNullCreateRequest creates the GetMethodLocalNull request.
func (client *APIVersionLocalClient) getMethodLocalNullCreateRequest(ctx context.Context, options *APIVersionLocalClientGetMethodLocalNullOptions) (*policy.Request, error) {
	urlPath := "/azurespecials/apiVersion/method/string/none/query/local/null"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.APIVersion != nil {
		reqQP.Set("api-version", *options.APIVersion)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// GetMethodLocalValid - Get method with api-version modeled in the method. pass in api-version = '2.0' to succeed
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2015-07-01-preview
// options - APIVersionLocalClientGetMethodLocalValidOptions contains the optional parameters for the APIVersionLocalClient.GetMethodLocalValid
// method.
func (client *APIVersionLocalClient) GetMethodLocalValid(ctx context.Context, options *APIVersionLocalClientGetMethodLocalValidOptions) (APIVersionLocalClientGetMethodLocalValidResponse, error) {
	req, err := client.getMethodLocalValidCreateRequest(ctx, options)
	if err != nil {
		return APIVersionLocalClientGetMethodLocalValidResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return APIVersionLocalClientGetMethodLocalValidResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return APIVersionLocalClientGetMethodLocalValidResponse{}, runtime.NewResponseError(resp)
	}
	return APIVersionLocalClientGetMethodLocalValidResponse{}, nil
}

// getMethodLocalValidCreateRequest creates the GetMethodLocalValid request.
func (client *APIVersionLocalClient) getMethodLocalValidCreateRequest(ctx context.Context, options *APIVersionLocalClientGetMethodLocalValidOptions) (*policy.Request, error) {
	urlPath := "/azurespecials/apiVersion/method/string/none/query/local/2.0"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2.0")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// GetPathLocalValid - Get method with api-version modeled in the method. pass in api-version = '2.0' to succeed
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2015-07-01-preview
// options - APIVersionLocalClientGetPathLocalValidOptions contains the optional parameters for the APIVersionLocalClient.GetPathLocalValid
// method.
func (client *APIVersionLocalClient) GetPathLocalValid(ctx context.Context, options *APIVersionLocalClientGetPathLocalValidOptions) (APIVersionLocalClientGetPathLocalValidResponse, error) {
	req, err := client.getPathLocalValidCreateRequest(ctx, options)
	if err != nil {
		return APIVersionLocalClientGetPathLocalValidResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return APIVersionLocalClientGetPathLocalValidResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return APIVersionLocalClientGetPathLocalValidResponse{}, runtime.NewResponseError(resp)
	}
	return APIVersionLocalClientGetPathLocalValidResponse{}, nil
}

// getPathLocalValidCreateRequest creates the GetPathLocalValid request.
func (client *APIVersionLocalClient) getPathLocalValidCreateRequest(ctx context.Context, options *APIVersionLocalClientGetPathLocalValidOptions) (*policy.Request, error) {
	urlPath := "/azurespecials/apiVersion/path/string/none/query/local/2.0"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2.0")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// GetSwaggerLocalValid - Get method with api-version modeled in the method. pass in api-version = '2.0' to succeed
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2015-07-01-preview
// options - APIVersionLocalClientGetSwaggerLocalValidOptions contains the optional parameters for the APIVersionLocalClient.GetSwaggerLocalValid
// method.
func (client *APIVersionLocalClient) GetSwaggerLocalValid(ctx context.Context, options *APIVersionLocalClientGetSwaggerLocalValidOptions) (APIVersionLocalClientGetSwaggerLocalValidResponse, error) {
	req, err := client.getSwaggerLocalValidCreateRequest(ctx, options)
	if err != nil {
		return APIVersionLocalClientGetSwaggerLocalValidResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return APIVersionLocalClientGetSwaggerLocalValidResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return APIVersionLocalClientGetSwaggerLocalValidResponse{}, runtime.NewResponseError(resp)
	}
	return APIVersionLocalClientGetSwaggerLocalValidResponse{}, nil
}

// getSwaggerLocalValidCreateRequest creates the GetSwaggerLocalValid request.
func (client *APIVersionLocalClient) getSwaggerLocalValidCreateRequest(ctx context.Context, options *APIVersionLocalClientGetSwaggerLocalValidOptions) (*policy.Request, error) {
	urlPath := "/azurespecials/apiVersion/swagger/string/none/query/local/2.0"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2.0")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}
