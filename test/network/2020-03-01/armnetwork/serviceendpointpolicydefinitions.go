// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"
)

// ServiceEndpointPolicyDefinitionsOperations contains the methods for the ServiceEndpointPolicyDefinitions group.
type ServiceEndpointPolicyDefinitionsOperations interface {
	// BeginCreateOrUpdate - Creates or updates a service endpoint policy definition in the specified service endpoint policy.
	BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serviceEndpointPolicyName string, serviceEndpointPolicyDefinitionName string, serviceEndpointPolicyDefinitions ServiceEndpointPolicyDefinition) (*ServiceEndpointPolicyDefinitionPollerResponse, error)
	// ResumeCreateOrUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeCreateOrUpdate(token string) (ServiceEndpointPolicyDefinitionPoller, error)
	// BeginDelete - Deletes the specified ServiceEndpoint policy definitions.
	BeginDelete(ctx context.Context, resourceGroupName string, serviceEndpointPolicyName string, serviceEndpointPolicyDefinitionName string) (*HTTPPollerResponse, error)
	// ResumeDelete - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDelete(token string) (HTTPPoller, error)
	// Get - Get the specified service endpoint policy definitions from service endpoint policy.
	Get(ctx context.Context, resourceGroupName string, serviceEndpointPolicyName string, serviceEndpointPolicyDefinitionName string) (*ServiceEndpointPolicyDefinitionResponse, error)
	// ListByResourceGroup - Gets all service endpoint policy definitions in a service end point policy.
	ListByResourceGroup(resourceGroupName string, serviceEndpointPolicyName string) (ServiceEndpointPolicyDefinitionListResultPager, error)
}

// serviceEndpointPolicyDefinitionsOperations implements the ServiceEndpointPolicyDefinitionsOperations interface.
type serviceEndpointPolicyDefinitionsOperations struct {
	*Client
	subscriptionID string
}

// CreateOrUpdate - Creates or updates a service endpoint policy definition in the specified service endpoint policy.
func (client *serviceEndpointPolicyDefinitionsOperations) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serviceEndpointPolicyName string, serviceEndpointPolicyDefinitionName string, serviceEndpointPolicyDefinitions ServiceEndpointPolicyDefinition) (*ServiceEndpointPolicyDefinitionPollerResponse, error) {
	req, err := client.createOrUpdateCreateRequest(resourceGroupName, serviceEndpointPolicyName, serviceEndpointPolicyDefinitionName, serviceEndpointPolicyDefinitions)
	if err != nil {
		return nil, err
	}
	// send the first request to initialize the poller
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.createOrUpdateHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	pt, err := createPollingTracker("serviceEndpointPolicyDefinitionsOperations.CreateOrUpdate", "azure-async-operation", resp, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &serviceEndpointPolicyDefinitionPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*ServiceEndpointPolicyDefinitionResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *serviceEndpointPolicyDefinitionsOperations) ResumeCreateOrUpdate(token string) (ServiceEndpointPolicyDefinitionPoller, error) {
	pt, err := resumePollingTracker("serviceEndpointPolicyDefinitionsOperations.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &serviceEndpointPolicyDefinitionPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *serviceEndpointPolicyDefinitionsOperations) createOrUpdateCreateRequest(resourceGroupName string, serviceEndpointPolicyName string, serviceEndpointPolicyDefinitionName string, serviceEndpointPolicyDefinitions ServiceEndpointPolicyDefinition) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/serviceEndpointPolicies/{serviceEndpointPolicyName}/serviceEndpointPolicyDefinitions/{serviceEndpointPolicyDefinitionName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{serviceEndpointPolicyName}", url.PathEscape(serviceEndpointPolicyName))
	urlPath = strings.ReplaceAll(urlPath, "{serviceEndpointPolicyDefinitionName}", url.PathEscape(serviceEndpointPolicyDefinitionName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(path.Join(client.u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPut, *u)
	return req, req.MarshalAsJSON(serviceEndpointPolicyDefinitions)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *serviceEndpointPolicyDefinitionsOperations) createOrUpdateHandleResponse(resp *azcore.Response) (*ServiceEndpointPolicyDefinitionPollerResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated, http.StatusNoContent) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return &ServiceEndpointPolicyDefinitionPollerResponse{RawResponse: resp.Response}, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *serviceEndpointPolicyDefinitionsOperations) createOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Delete - Deletes the specified ServiceEndpoint policy definitions.
func (client *serviceEndpointPolicyDefinitionsOperations) BeginDelete(ctx context.Context, resourceGroupName string, serviceEndpointPolicyName string, serviceEndpointPolicyDefinitionName string) (*HTTPPollerResponse, error) {
	req, err := client.deleteCreateRequest(resourceGroupName, serviceEndpointPolicyName, serviceEndpointPolicyDefinitionName)
	if err != nil {
		return nil, err
	}
	// send the first request to initialize the poller
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.deleteHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	pt, err := createPollingTracker("serviceEndpointPolicyDefinitionsOperations.Delete", "location", resp, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	poller := &httpPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *serviceEndpointPolicyDefinitionsOperations) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := resumePollingTracker("serviceEndpointPolicyDefinitionsOperations.Delete", token, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *serviceEndpointPolicyDefinitionsOperations) deleteCreateRequest(resourceGroupName string, serviceEndpointPolicyName string, serviceEndpointPolicyDefinitionName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/serviceEndpointPolicies/{serviceEndpointPolicyName}/serviceEndpointPolicyDefinitions/{serviceEndpointPolicyDefinitionName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{serviceEndpointPolicyName}", url.PathEscape(serviceEndpointPolicyName))
	urlPath = strings.ReplaceAll(urlPath, "{serviceEndpointPolicyDefinitionName}", url.PathEscape(serviceEndpointPolicyDefinitionName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(path.Join(client.u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodDelete, *u)
	return req, nil
}

// deleteHandleResponse handles the Delete response.
func (client *serviceEndpointPolicyDefinitionsOperations) deleteHandleResponse(resp *azcore.Response) (*HTTPPollerResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return &HTTPPollerResponse{RawResponse: resp.Response}, nil
}

// deleteHandleError handles the Delete error response.
func (client *serviceEndpointPolicyDefinitionsOperations) deleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Get - Get the specified service endpoint policy definitions from service endpoint policy.
func (client *serviceEndpointPolicyDefinitionsOperations) Get(ctx context.Context, resourceGroupName string, serviceEndpointPolicyName string, serviceEndpointPolicyDefinitionName string) (*ServiceEndpointPolicyDefinitionResponse, error) {
	req, err := client.getCreateRequest(resourceGroupName, serviceEndpointPolicyName, serviceEndpointPolicyDefinitionName)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getCreateRequest creates the Get request.
func (client *serviceEndpointPolicyDefinitionsOperations) getCreateRequest(resourceGroupName string, serviceEndpointPolicyName string, serviceEndpointPolicyDefinitionName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/serviceEndpointPolicies/{serviceEndpointPolicyName}/serviceEndpointPolicyDefinitions/{serviceEndpointPolicyDefinitionName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{serviceEndpointPolicyName}", url.PathEscape(serviceEndpointPolicyName))
	urlPath = strings.ReplaceAll(urlPath, "{serviceEndpointPolicyDefinitionName}", url.PathEscape(serviceEndpointPolicyDefinitionName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(path.Join(client.u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *serviceEndpointPolicyDefinitionsOperations) getHandleResponse(resp *azcore.Response) (*ServiceEndpointPolicyDefinitionResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getHandleError(resp)
	}
	result := ServiceEndpointPolicyDefinitionResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ServiceEndpointPolicyDefinition)
}

// getHandleError handles the Get error response.
func (client *serviceEndpointPolicyDefinitionsOperations) getHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// ListByResourceGroup - Gets all service endpoint policy definitions in a service end point policy.
func (client *serviceEndpointPolicyDefinitionsOperations) ListByResourceGroup(resourceGroupName string, serviceEndpointPolicyName string) (ServiceEndpointPolicyDefinitionListResultPager, error) {
	req, err := client.listByResourceGroupCreateRequest(resourceGroupName, serviceEndpointPolicyName)
	if err != nil {
		return nil, err
	}
	return &serviceEndpointPolicyDefinitionListResultPager{
		pipeline:  client.p,
		request:   req,
		responder: client.listByResourceGroupHandleResponse,
		advancer: func(resp *ServiceEndpointPolicyDefinitionListResultResponse) (*azcore.Request, error) {
			u, err := url.Parse(*resp.ServiceEndpointPolicyDefinitionListResult.NextLink)
			if err != nil {
				return nil, fmt.Errorf("invalid NextLink: %w", err)
			}
			if u.Scheme == "" {
				return nil, fmt.Errorf("no scheme detected in NextLink %s", *resp.ServiceEndpointPolicyDefinitionListResult.NextLink)
			}
			return azcore.NewRequest(http.MethodGet, *u), nil
		},
	}, nil
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *serviceEndpointPolicyDefinitionsOperations) listByResourceGroupCreateRequest(resourceGroupName string, serviceEndpointPolicyName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/serviceEndpointPolicies/{serviceEndpointPolicyName}/serviceEndpointPolicyDefinitions"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{serviceEndpointPolicyName}", url.PathEscape(serviceEndpointPolicyName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(path.Join(client.u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *serviceEndpointPolicyDefinitionsOperations) listByResourceGroupHandleResponse(resp *azcore.Response) (*ServiceEndpointPolicyDefinitionListResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.listByResourceGroupHandleError(resp)
	}
	result := ServiceEndpointPolicyDefinitionListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ServiceEndpointPolicyDefinitionListResult)
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *serviceEndpointPolicyDefinitionsOperations) listByResourceGroupHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}