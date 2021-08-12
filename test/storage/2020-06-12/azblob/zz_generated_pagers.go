// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azblob

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"reflect"
)

// ContainerListBlobFlatSegmentPager provides operations for iterating over paged responses.
type ContainerListBlobFlatSegmentPager struct {
	client    *containerClient
	current   ContainerListBlobFlatSegmentResponse
	err       error
	requester func(context.Context) (*azcore.Request, error)
	advancer  func(context.Context, ContainerListBlobFlatSegmentResponse) (*azcore.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ContainerListBlobFlatSegmentPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ContainerListBlobFlatSegmentPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ListBlobsFlatSegmentResponse.NextMarker == nil || len(*p.current.ListBlobsFlatSegmentResponse.NextMarker) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(http.StatusOK) {
		p.err = p.client.listBlobFlatSegmentHandleError(resp)
		return false
	}
	result, err := p.client.listBlobFlatSegmentHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ContainerListBlobFlatSegmentResponse page.
func (p *ContainerListBlobFlatSegmentPager) PageResponse() ContainerListBlobFlatSegmentResponse {
	return p.current
}

// ContainerListBlobHierarchySegmentPager provides operations for iterating over paged responses.
type ContainerListBlobHierarchySegmentPager struct {
	client    *containerClient
	current   ContainerListBlobHierarchySegmentResponse
	err       error
	requester func(context.Context) (*azcore.Request, error)
	advancer  func(context.Context, ContainerListBlobHierarchySegmentResponse) (*azcore.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ContainerListBlobHierarchySegmentPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ContainerListBlobHierarchySegmentPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ListBlobsHierarchySegmentResponse.NextMarker == nil || len(*p.current.ListBlobsHierarchySegmentResponse.NextMarker) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(http.StatusOK) {
		p.err = p.client.listBlobHierarchySegmentHandleError(resp)
		return false
	}
	result, err := p.client.listBlobHierarchySegmentHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ContainerListBlobHierarchySegmentResponse page.
func (p *ContainerListBlobHierarchySegmentPager) PageResponse() ContainerListBlobHierarchySegmentResponse {
	return p.current
}

// ServiceListContainersSegmentPager provides operations for iterating over paged responses.
type ServiceListContainersSegmentPager struct {
	client    *serviceClient
	current   ServiceListContainersSegmentResponse
	err       error
	requester func(context.Context) (*azcore.Request, error)
	advancer  func(context.Context, ServiceListContainersSegmentResponse) (*azcore.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ServiceListContainersSegmentPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ServiceListContainersSegmentPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ListContainersSegmentResponse.NextMarker == nil || len(*p.current.ListContainersSegmentResponse.NextMarker) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(http.StatusOK) {
		p.err = p.client.listContainersSegmentHandleError(resp)
		return false
	}
	result, err := p.client.listContainersSegmentHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ServiceListContainersSegmentResponse page.
func (p *ServiceListContainersSegmentPager) PageResponse() ServiceListContainersSegmentResponse {
	return p.current
}
