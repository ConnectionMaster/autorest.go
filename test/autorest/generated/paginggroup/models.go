// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package paginggroup

import "net/http"

type OdataProductResult struct {
	OdataNextLink *string    `json:"odata.nextLink,omitempty"`
	Values        *[]Product `json:"values,omitempty"`
}

// OdataProductResultResponse is the response envelope for operations that return a OdataProductResult type.
type OdataProductResultResponse struct {
	OdataProductResult *OdataProductResult

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

type OperationResult struct {
	// The status of the request
	Status *OperationResultStatus `undefined:"status"`
}

// PagingGetMultiplePagesLroOptions contains the optional parameters for the Paging.GetMultiplePagesLro method.
type PagingGetMultiplePagesLroOptions struct {
	ClientRequestId *string
	// Sets the maximum number of items to return in the response.
	Maxresults *int32
	// Sets the maximum time that the server can spend processing the request, in seconds. The default is 30 seconds.
	Timeout *int32
}

// PagingGetMultiplePagesOptions contains the optional parameters for the Paging.GetMultiplePages method.
type PagingGetMultiplePagesOptions struct {
	ClientRequestId *string
	// Sets the maximum number of items to return in the response.
	Maxresults *int32
	// Sets the maximum time that the server can spend processing the request, in seconds. The default is 30 seconds.
	Timeout *int32
}

// PagingGetMultiplePagesWithOffsetOptions contains the optional parameters for the Paging.GetMultiplePagesWithOffset method.
type PagingGetMultiplePagesWithOffsetOptions struct {
	ClientRequestId *string
	// Sets the maximum number of items to return in the response.
	Maxresults *int32
	// Sets the maximum time that the server can spend processing the request, in seconds. The default is 30 seconds.
	Timeout *int32
}

// PagingGetOdataMultiplePagesOptions contains the optional parameters for the Paging.GetOdataMultiplePages method.
type PagingGetOdataMultiplePagesOptions struct {
	ClientRequestId *string
	// Sets the maximum number of items to return in the response.
	Maxresults *int32
	// Sets the maximum time that the server can spend processing the request, in seconds. The default is 30 seconds.
	Timeout *int32
}

type Product struct {
	Properties *ProductProperties `json:"properties,omitempty"`
}

type ProductProperties struct {
	ID   *int32  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type ProductResult struct {
	NextLink *string    `json:"nextLink,omitempty"`
	Values   *[]Product `json:"values,omitempty"`
}

// ProductResultResponse is the response envelope for operations that return a ProductResult type.
type ProductResultResponse struct {
	ProductResult *ProductResult

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

type ProductResultValue struct {
	NextLink *string    `json:"nextLink,omitempty"`
	Value    *[]Product `json:"value,omitempty"`
}

// ProductResultValueResponse is the response envelope for operations that return a ProductResultValue type.
type ProductResultValueResponse struct {
	ProductResultValue *ProductResultValue

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}