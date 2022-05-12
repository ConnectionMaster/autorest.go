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

// CreditsClient contains the methods for the Credits group.
// Don't use this type directly, use NewCreditsClient() instead.
type CreditsClient struct {
	host string
	pl   runtime.Pipeline
}

// NewCreditsClient creates a new instance of CreditsClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewCreditsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*CreditsClient, error) {
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
	client := &CreditsClient{
		host: ep,
		pl:   pl,
	}
	return client, nil
}

// Get - The credit summary by billingAccountId and billingProfileId.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-10-01
// scope - The scope associated with credits operations. This includes '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfile/{billingProfileId}'
// for Billing Profile scope, and
// 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/customers/{customerId}' specific for partners.
// options - CreditsClientGetOptions contains the optional parameters for the CreditsClient.Get method.
func (client *CreditsClient) Get(ctx context.Context, scope string, options *CreditsClientGetOptions) (CreditsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, scope, options)
	if err != nil {
		return CreditsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CreditsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CreditsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *CreditsClient) getCreateRequest(ctx context.Context, scope string, options *CreditsClientGetOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Consumption/credits/balanceSummary"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *CreditsClient) getHandleResponse(resp *http.Response) (CreditsClientGetResponse, error) {
	result := CreditsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CreditSummary); err != nil {
		return CreditsClientGetResponse{}, err
	}
	return result, nil
}
