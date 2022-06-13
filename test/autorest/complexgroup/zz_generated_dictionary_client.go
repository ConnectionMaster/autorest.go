//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package complexgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// DictionaryClient contains the methods for the Dictionary group.
// Don't use this type directly, use NewDictionaryClient() instead.
type DictionaryClient struct {
	pl runtime.Pipeline
}

// NewDictionaryClient creates a new instance of DictionaryClient with the specified values.
// pl - the pipeline used for sending requests and handling responses.
func NewDictionaryClient(pl runtime.Pipeline) *DictionaryClient {
	client := &DictionaryClient{
		pl: pl,
	}
	return client
}

// GetEmpty - Get complex types with dictionary property which is empty
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2016-02-29
// options - DictionaryClientGetEmptyOptions contains the optional parameters for the DictionaryClient.GetEmpty method.
func (client *DictionaryClient) GetEmpty(ctx context.Context, options *DictionaryClientGetEmptyOptions) (DictionaryClientGetEmptyResponse, error) {
	req, err := client.getEmptyCreateRequest(ctx, options)
	if err != nil {
		return DictionaryClientGetEmptyResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DictionaryClientGetEmptyResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DictionaryClientGetEmptyResponse{}, runtime.NewResponseError(resp)
	}
	return client.getEmptyHandleResponse(resp)
}

// getEmptyCreateRequest creates the GetEmpty request.
func (client *DictionaryClient) getEmptyCreateRequest(ctx context.Context, options *DictionaryClientGetEmptyOptions) (*policy.Request, error) {
	urlPath := "/complex/dictionary/typed/empty"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getEmptyHandleResponse handles the GetEmpty response.
func (client *DictionaryClient) getEmptyHandleResponse(resp *http.Response) (DictionaryClientGetEmptyResponse, error) {
	result := DictionaryClientGetEmptyResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DictionaryWrapper); err != nil {
		return DictionaryClientGetEmptyResponse{}, err
	}
	return result, nil
}

// GetNotProvided - Get complex types with dictionary property while server doesn't provide a response payload
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2016-02-29
// options - DictionaryClientGetNotProvidedOptions contains the optional parameters for the DictionaryClient.GetNotProvided
// method.
func (client *DictionaryClient) GetNotProvided(ctx context.Context, options *DictionaryClientGetNotProvidedOptions) (DictionaryClientGetNotProvidedResponse, error) {
	req, err := client.getNotProvidedCreateRequest(ctx, options)
	if err != nil {
		return DictionaryClientGetNotProvidedResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DictionaryClientGetNotProvidedResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DictionaryClientGetNotProvidedResponse{}, runtime.NewResponseError(resp)
	}
	return client.getNotProvidedHandleResponse(resp)
}

// getNotProvidedCreateRequest creates the GetNotProvided request.
func (client *DictionaryClient) getNotProvidedCreateRequest(ctx context.Context, options *DictionaryClientGetNotProvidedOptions) (*policy.Request, error) {
	urlPath := "/complex/dictionary/typed/notprovided"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getNotProvidedHandleResponse handles the GetNotProvided response.
func (client *DictionaryClient) getNotProvidedHandleResponse(resp *http.Response) (DictionaryClientGetNotProvidedResponse, error) {
	result := DictionaryClientGetNotProvidedResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DictionaryWrapper); err != nil {
		return DictionaryClientGetNotProvidedResponse{}, err
	}
	return result, nil
}

// GetNull - Get complex types with dictionary property which is null
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2016-02-29
// options - DictionaryClientGetNullOptions contains the optional parameters for the DictionaryClient.GetNull method.
func (client *DictionaryClient) GetNull(ctx context.Context, options *DictionaryClientGetNullOptions) (DictionaryClientGetNullResponse, error) {
	req, err := client.getNullCreateRequest(ctx, options)
	if err != nil {
		return DictionaryClientGetNullResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DictionaryClientGetNullResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DictionaryClientGetNullResponse{}, runtime.NewResponseError(resp)
	}
	return client.getNullHandleResponse(resp)
}

// getNullCreateRequest creates the GetNull request.
func (client *DictionaryClient) getNullCreateRequest(ctx context.Context, options *DictionaryClientGetNullOptions) (*policy.Request, error) {
	urlPath := "/complex/dictionary/typed/null"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getNullHandleResponse handles the GetNull response.
func (client *DictionaryClient) getNullHandleResponse(resp *http.Response) (DictionaryClientGetNullResponse, error) {
	result := DictionaryClientGetNullResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DictionaryWrapper); err != nil {
		return DictionaryClientGetNullResponse{}, err
	}
	return result, nil
}

// GetValid - Get complex types with dictionary property
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2016-02-29
// options - DictionaryClientGetValidOptions contains the optional parameters for the DictionaryClient.GetValid method.
func (client *DictionaryClient) GetValid(ctx context.Context, options *DictionaryClientGetValidOptions) (DictionaryClientGetValidResponse, error) {
	req, err := client.getValidCreateRequest(ctx, options)
	if err != nil {
		return DictionaryClientGetValidResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DictionaryClientGetValidResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DictionaryClientGetValidResponse{}, runtime.NewResponseError(resp)
	}
	return client.getValidHandleResponse(resp)
}

// getValidCreateRequest creates the GetValid request.
func (client *DictionaryClient) getValidCreateRequest(ctx context.Context, options *DictionaryClientGetValidOptions) (*policy.Request, error) {
	urlPath := "/complex/dictionary/typed/valid"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getValidHandleResponse handles the GetValid response.
func (client *DictionaryClient) getValidHandleResponse(resp *http.Response) (DictionaryClientGetValidResponse, error) {
	result := DictionaryClientGetValidResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DictionaryWrapper); err != nil {
		return DictionaryClientGetValidResponse{}, err
	}
	return result, nil
}

// PutEmpty - Put complex types with dictionary property which is empty
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2016-02-29
// complexBody - Please put an empty dictionary
// options - DictionaryClientPutEmptyOptions contains the optional parameters for the DictionaryClient.PutEmpty method.
func (client *DictionaryClient) PutEmpty(ctx context.Context, complexBody DictionaryWrapper, options *DictionaryClientPutEmptyOptions) (DictionaryClientPutEmptyResponse, error) {
	req, err := client.putEmptyCreateRequest(ctx, complexBody, options)
	if err != nil {
		return DictionaryClientPutEmptyResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DictionaryClientPutEmptyResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DictionaryClientPutEmptyResponse{}, runtime.NewResponseError(resp)
	}
	return DictionaryClientPutEmptyResponse{}, nil
}

// putEmptyCreateRequest creates the PutEmpty request.
func (client *DictionaryClient) putEmptyCreateRequest(ctx context.Context, complexBody DictionaryWrapper, options *DictionaryClientPutEmptyOptions) (*policy.Request, error) {
	urlPath := "/complex/dictionary/typed/empty"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, complexBody)
}

// PutValid - Put complex types with dictionary property
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2016-02-29
// complexBody - Please put a dictionary with 5 key-value pairs: "txt":"notepad", "bmp":"mspaint", "xls":"excel", "exe":"",
// "":null
// options - DictionaryClientPutValidOptions contains the optional parameters for the DictionaryClient.PutValid method.
func (client *DictionaryClient) PutValid(ctx context.Context, complexBody DictionaryWrapper, options *DictionaryClientPutValidOptions) (DictionaryClientPutValidResponse, error) {
	req, err := client.putValidCreateRequest(ctx, complexBody, options)
	if err != nil {
		return DictionaryClientPutValidResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DictionaryClientPutValidResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DictionaryClientPutValidResponse{}, runtime.NewResponseError(resp)
	}
	return DictionaryClientPutValidResponse{}, nil
}

// putValidCreateRequest creates the PutValid request.
func (client *DictionaryClient) putValidCreateRequest(ctx context.Context, complexBody DictionaryWrapper, options *DictionaryClientPutValidOptions) (*policy.Request, error) {
	urlPath := "/complex/dictionary/typed/valid"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, complexBody)
}
