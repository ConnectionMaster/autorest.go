//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armconsumption

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ReservationsSummariesClient contains the methods for the ReservationsSummaries group.
// Don't use this type directly, use NewReservationsSummariesClient() instead.
type ReservationsSummariesClient struct {
	host string
	pl   runtime.Pipeline
}

// NewReservationsSummariesClient creates a new instance of ReservationsSummariesClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewReservationsSummariesClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*ReservationsSummariesClient, error) {
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
	client := &ReservationsSummariesClient{
		host: ep,
		pl:   pl,
	}
	return client, nil
}

// NewListPager - Lists the reservations summaries for the defined scope daily or monthly grain.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-10-01
// scope - The scope associated with reservations summaries operations. This includes '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}'
// for BillingAccount scope (legacy), and
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for BillingProfile
// scope (modern).
// grain - Can be daily or monthly
// options - ReservationsSummariesClientListOptions contains the optional parameters for the ReservationsSummariesClient.List
// method.
func (client *ReservationsSummariesClient) NewListPager(scope string, grain Datagrain, options *ReservationsSummariesClientListOptions) *runtime.Pager[ReservationsSummariesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ReservationsSummariesClientListResponse]{
		More: func(page ReservationsSummariesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ReservationsSummariesClientListResponse) (ReservationsSummariesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, scope, grain, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ReservationsSummariesClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ReservationsSummariesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ReservationsSummariesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ReservationsSummariesClient) listCreateRequest(ctx context.Context, scope string, grain Datagrain, options *ReservationsSummariesClientListOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Consumption/reservationSummaries"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("grain", string(grain))
	if options != nil && options.StartDate != nil {
		reqQP.Set("startDate", *options.StartDate)
	}
	if options != nil && options.EndDate != nil {
		reqQP.Set("endDate", *options.EndDate)
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.ReservationID != nil {
		reqQP.Set("reservationId", *options.ReservationID)
	}
	if options != nil && options.ReservationOrderID != nil {
		reqQP.Set("reservationOrderId", *options.ReservationOrderID)
	}
	reqQP.Set("api-version", "2019-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ReservationsSummariesClient) listHandleResponse(resp *http.Response) (ReservationsSummariesClientListResponse, error) {
	result := ReservationsSummariesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ReservationSummariesListResult); err != nil {
		return ReservationsSummariesClientListResponse{}, err
	}
	return result, nil
}

// NewListByReservationOrderPager - Lists the reservations summaries for daily or monthly grain.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-10-01
// reservationOrderID - Order Id of the reservation
// grain - Can be daily or monthly
// options - ReservationsSummariesClientListByReservationOrderOptions contains the optional parameters for the ReservationsSummariesClient.ListByReservationOrder
// method.
func (client *ReservationsSummariesClient) NewListByReservationOrderPager(reservationOrderID string, grain Datagrain, options *ReservationsSummariesClientListByReservationOrderOptions) *runtime.Pager[ReservationsSummariesClientListByReservationOrderResponse] {
	return runtime.NewPager(runtime.PagingHandler[ReservationsSummariesClientListByReservationOrderResponse]{
		More: func(page ReservationsSummariesClientListByReservationOrderResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ReservationsSummariesClientListByReservationOrderResponse) (ReservationsSummariesClientListByReservationOrderResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByReservationOrderCreateRequest(ctx, reservationOrderID, grain, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ReservationsSummariesClientListByReservationOrderResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ReservationsSummariesClientListByReservationOrderResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ReservationsSummariesClientListByReservationOrderResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByReservationOrderHandleResponse(resp)
		},
	})
}

// listByReservationOrderCreateRequest creates the ListByReservationOrder request.
func (client *ReservationsSummariesClient) listByReservationOrderCreateRequest(ctx context.Context, reservationOrderID string, grain Datagrain, options *ReservationsSummariesClientListByReservationOrderOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Capacity/reservationorders/{reservationOrderId}/providers/Microsoft.Consumption/reservationSummaries"
	if reservationOrderID == "" {
		return nil, errors.New("parameter reservationOrderID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{reservationOrderId}", url.PathEscape(reservationOrderID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("grain", string(grain))
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2019-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByReservationOrderHandleResponse handles the ListByReservationOrder response.
func (client *ReservationsSummariesClient) listByReservationOrderHandleResponse(resp *http.Response) (ReservationsSummariesClientListByReservationOrderResponse, error) {
	result := ReservationsSummariesClientListByReservationOrderResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ReservationSummariesListResult); err != nil {
		return ReservationsSummariesClientListByReservationOrderResponse{}, err
	}
	return result, nil
}

// NewListByReservationOrderAndReservationPager - Lists the reservations summaries for daily or monthly grain.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-10-01
// reservationOrderID - Order Id of the reservation
// reservationID - Id of the reservation
// grain - Can be daily or monthly
// options - ReservationsSummariesClientListByReservationOrderAndReservationOptions contains the optional parameters for the
// ReservationsSummariesClient.ListByReservationOrderAndReservation method.
func (client *ReservationsSummariesClient) NewListByReservationOrderAndReservationPager(reservationOrderID string, reservationID string, grain Datagrain, options *ReservationsSummariesClientListByReservationOrderAndReservationOptions) *runtime.Pager[ReservationsSummariesClientListByReservationOrderAndReservationResponse] {
	return runtime.NewPager(runtime.PagingHandler[ReservationsSummariesClientListByReservationOrderAndReservationResponse]{
		More: func(page ReservationsSummariesClientListByReservationOrderAndReservationResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ReservationsSummariesClientListByReservationOrderAndReservationResponse) (ReservationsSummariesClientListByReservationOrderAndReservationResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByReservationOrderAndReservationCreateRequest(ctx, reservationOrderID, reservationID, grain, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ReservationsSummariesClientListByReservationOrderAndReservationResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ReservationsSummariesClientListByReservationOrderAndReservationResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ReservationsSummariesClientListByReservationOrderAndReservationResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByReservationOrderAndReservationHandleResponse(resp)
		},
	})
}

// listByReservationOrderAndReservationCreateRequest creates the ListByReservationOrderAndReservation request.
func (client *ReservationsSummariesClient) listByReservationOrderAndReservationCreateRequest(ctx context.Context, reservationOrderID string, reservationID string, grain Datagrain, options *ReservationsSummariesClientListByReservationOrderAndReservationOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Capacity/reservationorders/{reservationOrderId}/reservations/{reservationId}/providers/Microsoft.Consumption/reservationSummaries"
	if reservationOrderID == "" {
		return nil, errors.New("parameter reservationOrderID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{reservationOrderId}", url.PathEscape(reservationOrderID))
	if reservationID == "" {
		return nil, errors.New("parameter reservationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{reservationId}", url.PathEscape(reservationID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("grain", string(grain))
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2019-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByReservationOrderAndReservationHandleResponse handles the ListByReservationOrderAndReservation response.
func (client *ReservationsSummariesClient) listByReservationOrderAndReservationHandleResponse(resp *http.Response) (ReservationsSummariesClientListByReservationOrderAndReservationResponse, error) {
	result := ReservationsSummariesClientListByReservationOrderAndReservationResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ReservationSummariesListResult); err != nil {
		return ReservationsSummariesClientListByReservationOrderAndReservationResponse{}, err
	}
	return result, nil
}
