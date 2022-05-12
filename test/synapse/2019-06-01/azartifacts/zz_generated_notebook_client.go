//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

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

type notebookClient struct {
	endpoint string
	pl       runtime.Pipeline
}

// newNotebookClient creates a new instance of notebookClient with the specified values.
// endpoint - The workspace development endpoint, for example https://myworkspace.dev.azuresynapse.net.
// pl - the pipeline used for sending requests and handling responses.
func newNotebookClient(endpoint string, pl runtime.Pipeline) *notebookClient {
	client := &notebookClient{
		endpoint: endpoint,
		pl:       pl,
	}
	return client
}

// BeginCreateOrUpdateNotebook - Creates or updates a Note Book.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-06-01-preview
// notebookName - The notebook name.
// notebook - Note book resource definition.
// options - notebookClientBeginCreateOrUpdateNotebookOptions contains the optional parameters for the notebookClient.BeginCreateOrUpdateNotebook
// method.
func (client *notebookClient) BeginCreateOrUpdateNotebook(ctx context.Context, notebookName string, notebook NotebookResource, options *notebookClientBeginCreateOrUpdateNotebookOptions) (*runtime.Poller[notebookClientCreateOrUpdateNotebookResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdateNotebook(ctx, notebookName, notebook, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[notebookClientCreateOrUpdateNotebookResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[notebookClientCreateOrUpdateNotebookResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdateNotebook - Creates or updates a Note Book.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-06-01-preview
func (client *notebookClient) createOrUpdateNotebook(ctx context.Context, notebookName string, notebook NotebookResource, options *notebookClientBeginCreateOrUpdateNotebookOptions) (*http.Response, error) {
	req, err := client.createOrUpdateNotebookCreateRequest(ctx, notebookName, notebook, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateNotebookCreateRequest creates the CreateOrUpdateNotebook request.
func (client *notebookClient) createOrUpdateNotebookCreateRequest(ctx context.Context, notebookName string, notebook NotebookResource, options *notebookClientBeginCreateOrUpdateNotebookOptions) (*policy.Request, error) {
	urlPath := "/notebooks/{notebookName}"
	if notebookName == "" {
		return nil, errors.New("parameter notebookName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{notebookName}", url.PathEscape(notebookName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, notebook)
}

// BeginDeleteNotebook - Deletes a Note book.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-06-01-preview
// notebookName - The notebook name.
// options - notebookClientBeginDeleteNotebookOptions contains the optional parameters for the notebookClient.BeginDeleteNotebook
// method.
func (client *notebookClient) BeginDeleteNotebook(ctx context.Context, notebookName string, options *notebookClientBeginDeleteNotebookOptions) (*runtime.Poller[notebookClientDeleteNotebookResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteNotebook(ctx, notebookName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[notebookClientDeleteNotebookResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[notebookClientDeleteNotebookResponse](options.ResumeToken, client.pl, nil)
	}
}

// DeleteNotebook - Deletes a Note book.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-06-01-preview
func (client *notebookClient) deleteNotebook(ctx context.Context, notebookName string, options *notebookClientBeginDeleteNotebookOptions) (*http.Response, error) {
	req, err := client.deleteNotebookCreateRequest(ctx, notebookName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteNotebookCreateRequest creates the DeleteNotebook request.
func (client *notebookClient) deleteNotebookCreateRequest(ctx context.Context, notebookName string, options *notebookClientBeginDeleteNotebookOptions) (*policy.Request, error) {
	urlPath := "/notebooks/{notebookName}"
	if notebookName == "" {
		return nil, errors.New("parameter notebookName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{notebookName}", url.PathEscape(notebookName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// GetNotebook - Gets a Note Book.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-06-01-preview
// notebookName - The notebook name.
// options - notebookClientGetNotebookOptions contains the optional parameters for the notebookClient.GetNotebook method.
func (client *notebookClient) GetNotebook(ctx context.Context, notebookName string, options *notebookClientGetNotebookOptions) (notebookClientGetNotebookResponse, error) {
	req, err := client.getNotebookCreateRequest(ctx, notebookName, options)
	if err != nil {
		return notebookClientGetNotebookResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return notebookClientGetNotebookResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNotModified) {
		return notebookClientGetNotebookResponse{}, runtime.NewResponseError(resp)
	}
	return client.getNotebookHandleResponse(resp)
}

// getNotebookCreateRequest creates the GetNotebook request.
func (client *notebookClient) getNotebookCreateRequest(ctx context.Context, notebookName string, options *notebookClientGetNotebookOptions) (*policy.Request, error) {
	urlPath := "/notebooks/{notebookName}"
	if notebookName == "" {
		return nil, errors.New("parameter notebookName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{notebookName}", url.PathEscape(notebookName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfNoneMatch != nil {
		req.Raw().Header["If-None-Match"] = []string{*options.IfNoneMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getNotebookHandleResponse handles the GetNotebook response.
func (client *notebookClient) getNotebookHandleResponse(resp *http.Response) (notebookClientGetNotebookResponse, error) {
	result := notebookClientGetNotebookResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NotebookResource); err != nil {
		return notebookClientGetNotebookResponse{}, err
	}
	return result, nil
}

// NewGetNotebookSummaryByWorkSpacePager - Lists a summary of Notebooks.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-06-01-preview
// options - notebookClientGetNotebookSummaryByWorkSpaceOptions contains the optional parameters for the notebookClient.GetNotebookSummaryByWorkSpace
// method.
func (client *notebookClient) NewGetNotebookSummaryByWorkSpacePager(options *notebookClientGetNotebookSummaryByWorkSpaceOptions) *runtime.Pager[notebookClientGetNotebookSummaryByWorkSpaceResponse] {
	return runtime.NewPager(runtime.PagingHandler[notebookClientGetNotebookSummaryByWorkSpaceResponse]{
		More: func(page notebookClientGetNotebookSummaryByWorkSpaceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *notebookClientGetNotebookSummaryByWorkSpaceResponse) (notebookClientGetNotebookSummaryByWorkSpaceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.getNotebookSummaryByWorkSpaceCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return notebookClientGetNotebookSummaryByWorkSpaceResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return notebookClientGetNotebookSummaryByWorkSpaceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return notebookClientGetNotebookSummaryByWorkSpaceResponse{}, runtime.NewResponseError(resp)
			}
			return client.getNotebookSummaryByWorkSpaceHandleResponse(resp)
		},
	})
}

// getNotebookSummaryByWorkSpaceCreateRequest creates the GetNotebookSummaryByWorkSpace request.
func (client *notebookClient) getNotebookSummaryByWorkSpaceCreateRequest(ctx context.Context, options *notebookClientGetNotebookSummaryByWorkSpaceOptions) (*policy.Request, error) {
	urlPath := "/notebooks/summary"
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

// getNotebookSummaryByWorkSpaceHandleResponse handles the GetNotebookSummaryByWorkSpace response.
func (client *notebookClient) getNotebookSummaryByWorkSpaceHandleResponse(resp *http.Response) (notebookClientGetNotebookSummaryByWorkSpaceResponse, error) {
	result := notebookClientGetNotebookSummaryByWorkSpaceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NotebookListResponse); err != nil {
		return notebookClientGetNotebookSummaryByWorkSpaceResponse{}, err
	}
	return result, nil
}

// NewGetNotebooksByWorkspacePager - Lists Notebooks.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-06-01-preview
// options - notebookClientGetNotebooksByWorkspaceOptions contains the optional parameters for the notebookClient.GetNotebooksByWorkspace
// method.
func (client *notebookClient) NewGetNotebooksByWorkspacePager(options *notebookClientGetNotebooksByWorkspaceOptions) *runtime.Pager[notebookClientGetNotebooksByWorkspaceResponse] {
	return runtime.NewPager(runtime.PagingHandler[notebookClientGetNotebooksByWorkspaceResponse]{
		More: func(page notebookClientGetNotebooksByWorkspaceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *notebookClientGetNotebooksByWorkspaceResponse) (notebookClientGetNotebooksByWorkspaceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.getNotebooksByWorkspaceCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return notebookClientGetNotebooksByWorkspaceResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return notebookClientGetNotebooksByWorkspaceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return notebookClientGetNotebooksByWorkspaceResponse{}, runtime.NewResponseError(resp)
			}
			return client.getNotebooksByWorkspaceHandleResponse(resp)
		},
	})
}

// getNotebooksByWorkspaceCreateRequest creates the GetNotebooksByWorkspace request.
func (client *notebookClient) getNotebooksByWorkspaceCreateRequest(ctx context.Context, options *notebookClientGetNotebooksByWorkspaceOptions) (*policy.Request, error) {
	urlPath := "/notebooks"
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

// getNotebooksByWorkspaceHandleResponse handles the GetNotebooksByWorkspace response.
func (client *notebookClient) getNotebooksByWorkspaceHandleResponse(resp *http.Response) (notebookClientGetNotebooksByWorkspaceResponse, error) {
	result := notebookClientGetNotebooksByWorkspaceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NotebookListResponse); err != nil {
		return notebookClientGetNotebooksByWorkspaceResponse{}, err
	}
	return result, nil
}

// BeginRenameNotebook - Renames a notebook.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-06-01-preview
// notebookName - The notebook name.
// request - proposed new name.
// options - notebookClientBeginRenameNotebookOptions contains the optional parameters for the notebookClient.BeginRenameNotebook
// method.
func (client *notebookClient) BeginRenameNotebook(ctx context.Context, notebookName string, request ArtifactRenameRequest, options *notebookClientBeginRenameNotebookOptions) (*runtime.Poller[notebookClientRenameNotebookResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.renameNotebook(ctx, notebookName, request, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[notebookClientRenameNotebookResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[notebookClientRenameNotebookResponse](options.ResumeToken, client.pl, nil)
	}
}

// RenameNotebook - Renames a notebook.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-06-01-preview
func (client *notebookClient) renameNotebook(ctx context.Context, notebookName string, request ArtifactRenameRequest, options *notebookClientBeginRenameNotebookOptions) (*http.Response, error) {
	req, err := client.renameNotebookCreateRequest(ctx, notebookName, request, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// renameNotebookCreateRequest creates the RenameNotebook request.
func (client *notebookClient) renameNotebookCreateRequest(ctx context.Context, notebookName string, request ArtifactRenameRequest, options *notebookClientBeginRenameNotebookOptions) (*policy.Request, error) {
	urlPath := "/notebooks/{notebookName}/rename"
	if notebookName == "" {
		return nil, errors.New("parameter notebookName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{notebookName}", url.PathEscape(notebookName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, request)
}
