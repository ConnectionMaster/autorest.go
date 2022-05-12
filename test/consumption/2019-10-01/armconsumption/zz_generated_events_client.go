//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armconsumption

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
)

// EventsClient contains the methods for the Events group.
// Don't use this type directly, use NewEventsClient() instead.
type EventsClient struct {
	host string
	pl   runtime.Pipeline
}

// NewEventsClient creates a new instance of EventsClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewEventsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*EventsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &EventsClient{
		host: ep,
		pl:   pl,
	}
	return client, nil
}

// NewListPager - Lists the events by billingAccountId and billingProfileId for given start and end date.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-10-01
// startDate - Start date
// endDate - End date
// scope - The scope associated with events operations. This includes '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfile/{billingProfileId}'
// for Billing Profile scope, and
// 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/customers/{customerId}' specific for partners.
// options - EventsClientListOptions contains the optional parameters for the EventsClient.List method.
func (client *EventsClient) NewListPager(startDate string, endDate string, scope string, options *EventsClientListOptions) *runtime.Pager[EventsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[EventsClientListResponse]{
		More: func(page EventsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *EventsClientListResponse) (EventsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, startDate, endDate, scope, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return EventsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return EventsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return EventsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *EventsClient) listCreateRequest(ctx context.Context, startDate string, endDate string, scope string, options *EventsClientListOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Consumption/events"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-10-01")
	reqQP.Set("startDate", startDate)
	reqQP.Set("endDate", endDate)
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *EventsClient) listHandleResponse(resp *http.Response) (EventsClientListResponse, error) {
	result := EventsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Events); err != nil {
		return EventsClientListResponse{}, err
	}
	return result, nil
}
