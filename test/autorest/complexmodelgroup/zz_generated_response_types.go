// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package complexmodelgroup

import "net/http"

// CatalogArrayResponse is the response envelope for operations that return a CatalogArray type.
type CatalogArrayResponse struct {
	CatalogArray *CatalogArray

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// CatalogDictionaryResponse is the response envelope for operations that return a CatalogDictionary type.
type CatalogDictionaryResponse struct {
	CatalogDictionary *CatalogDictionary

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}