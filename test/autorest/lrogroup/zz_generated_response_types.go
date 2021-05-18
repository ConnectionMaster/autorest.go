// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package lrogroup

import (
	"context"
	"net/http"
	"time"
)

// HTTPPollerResponse contains the asynchronous HTTP response from the call to the service endpoint.
type HTTPPollerResponse struct {
	// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received
	PollUntilDone func(ctx context.Context, frequency time.Duration) (*http.Response, error)

	// Poller contains an initialized poller.
	Poller HTTPPoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ProductArrayPollerResponse is the response envelope for operations that asynchronously return a []*Product type.
type ProductArrayPollerResponse struct {
	// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received
	PollUntilDone func(ctx context.Context, frequency time.Duration) (ProductArrayResponse, error)

	// Poller contains an initialized poller.
	Poller ProductArrayPoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ProductArrayResponse is the response envelope for operations that return a []*Product type.
type ProductArrayResponse struct {
	// Array of Product
	ProductArray []*Product

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ProductPollerResponse is the response envelope for operations that asynchronously return a Product type.
type ProductPollerResponse struct {
	// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received
	PollUntilDone func(ctx context.Context, frequency time.Duration) (ProductResponse, error)

	// Poller contains an initialized poller.
	Poller ProductPoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ProductResponse is the response envelope for operations that return a Product type.
type ProductResponse struct {
	Product *Product

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SKUPollerResponse is the response envelope for operations that asynchronously return a SKU type.
type SKUPollerResponse struct {
	// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received
	PollUntilDone func(ctx context.Context, frequency time.Duration) (SKUResponse, error)

	// Poller contains an initialized poller.
	Poller SKUPoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SKUResponse is the response envelope for operations that return a SKU type.
type SKUResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
	SKU         *SKU
}

// SubProductPollerResponse is the response envelope for operations that asynchronously return a SubProduct type.
type SubProductPollerResponse struct {
	// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received
	PollUntilDone func(ctx context.Context, frequency time.Duration) (SubProductResponse, error)

	// Poller contains an initialized poller.
	Poller SubProductPoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SubProductResponse is the response envelope for operations that return a SubProduct type.
type SubProductResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
	SubProduct  *SubProduct
}