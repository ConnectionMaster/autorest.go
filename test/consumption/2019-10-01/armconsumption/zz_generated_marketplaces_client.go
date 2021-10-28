//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armconsumption

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strconv"
	"strings"
)

// MarketplacesClient contains the methods for the Marketplaces group.
// Don't use this type directly, use NewMarketplacesClient() instead.
type MarketplacesClient struct {
	ep string
	pl runtime.Pipeline
}

// NewMarketplacesClient creates a new instance of MarketplacesClient with the specified values.
func NewMarketplacesClient(credential azcore.TokenCredential, options *arm.ClientOptions) *MarketplacesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &MarketplacesClient{ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// List - Lists the marketplaces for a scope at the defined scope. Marketplaces are available via this API only for May 1, 2014 or later.
// If the operation fails it returns the *ErrorResponse error type.
func (client *MarketplacesClient) List(scope string, options *MarketplacesListOptions) *MarketplacesListPager {
	return &MarketplacesListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, scope, options)
		},
		advancer: func(ctx context.Context, resp MarketplacesListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.MarketplacesListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *MarketplacesClient) listCreateRequest(ctx context.Context, scope string, options *MarketplacesListOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Consumption/marketplaces"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Skiptoken != nil {
		reqQP.Set("$skiptoken", *options.Skiptoken)
	}
	reqQP.Set("api-version", "2019-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *MarketplacesClient) listHandleResponse(resp *http.Response) (MarketplacesListResponse, error) {
	result := MarketplacesListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.MarketplacesListResult); err != nil {
		return MarketplacesListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *MarketplacesClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
