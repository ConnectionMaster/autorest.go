//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package aztables

import "time"

// ClientCreateResponse contains the response from method Client.Create.
type ClientCreateResponse struct {
	Response
	// ClientRequestID contains the information returned from the x-ms-client-request-id header response.
	ClientRequestID *string

	// Date contains the information returned from the Date header response.
	Date *time.Time

	// PreferenceApplied contains the information returned from the Preference-Applied header response.
	PreferenceApplied *string

	// RequestID contains the information returned from the x-ms-request-id header response.
	RequestID *string

	// Version contains the information returned from the x-ms-version header response.
	Version *string
}

// ClientDeleteEntityResponse contains the response from method Client.DeleteEntity.
type ClientDeleteEntityResponse struct {
	// ClientRequestID contains the information returned from the x-ms-client-request-id header response.
	ClientRequestID *string

	// Date contains the information returned from the Date header response.
	Date *time.Time

	// RequestID contains the information returned from the x-ms-request-id header response.
	RequestID *string

	// Version contains the information returned from the x-ms-version header response.
	Version *string
}

// ClientDeleteResponse contains the response from method Client.Delete.
type ClientDeleteResponse struct {
	// ClientRequestID contains the information returned from the x-ms-client-request-id header response.
	ClientRequestID *string

	// Date contains the information returned from the Date header response.
	Date *time.Time

	// RequestID contains the information returned from the x-ms-request-id header response.
	RequestID *string

	// Version contains the information returned from the x-ms-version header response.
	Version *string
}

// ClientGetAccessPolicyResponse contains the response from method Client.GetAccessPolicy.
type ClientGetAccessPolicyResponse struct {
	// ClientRequestID contains the information returned from the x-ms-client-request-id header response.
	ClientRequestID *string `xml:"ClientRequestID"`

	// Date contains the information returned from the Date header response.
	Date *time.Time `xml:"Date"`

	// RequestID contains the information returned from the x-ms-request-id header response.
	RequestID *string `xml:"RequestID"`

	// A collection of signed identifiers.
	SignedIdentifiers []*SignedIdentifier `xml:"SignedIdentifier"`

	// Version contains the information returned from the x-ms-version header response.
	Version *string `xml:"Version"`
}

// ClientInsertEntityResponse contains the response from method Client.InsertEntity.
type ClientInsertEntityResponse struct {
	// ClientRequestID contains the information returned from the x-ms-client-request-id header response.
	ClientRequestID *string

	// ContentType contains the information returned from the Content-Type header response.
	ContentType *string

	// Date contains the information returned from the Date header response.
	Date *time.Time

	// ETag contains the information returned from the ETag header response.
	ETag *string

	// PreferenceApplied contains the information returned from the Preference-Applied header response.
	PreferenceApplied *string

	// RequestID contains the information returned from the x-ms-request-id header response.
	RequestID *string

	// The other properties of the table entity.
	Value map[string]interface{}

	// Version contains the information returned from the x-ms-version header response.
	Version *string
}

// ClientMergeEntityResponse contains the response from method Client.MergeEntity.
type ClientMergeEntityResponse struct {
	// ClientRequestID contains the information returned from the x-ms-client-request-id header response.
	ClientRequestID *string

	// Date contains the information returned from the Date header response.
	Date *time.Time

	// ETag contains the information returned from the ETag header response.
	ETag *string

	// RequestID contains the information returned from the x-ms-request-id header response.
	RequestID *string

	// Version contains the information returned from the x-ms-version header response.
	Version *string
}

// ClientQueryEntitiesResponse contains the response from method Client.QueryEntities.
type ClientQueryEntitiesResponse struct {
	EntityQueryResponse
	// ClientRequestID contains the information returned from the x-ms-client-request-id header response.
	ClientRequestID *string

	// Date contains the information returned from the Date header response.
	Date *time.Time

	// RequestID contains the information returned from the x-ms-request-id header response.
	RequestID *string

	// Version contains the information returned from the x-ms-version header response.
	Version *string

	// XMSContinuationNextPartitionKey contains the information returned from the x-ms-continuation-NextPartitionKey header response.
	XMSContinuationNextPartitionKey *string

	// XMSContinuationNextRowKey contains the information returned from the x-ms-continuation-NextRowKey header response.
	XMSContinuationNextRowKey *string
}

// ClientQueryEntityWithPartitionAndRowKeyResponse contains the response from method Client.QueryEntityWithPartitionAndRowKey.
type ClientQueryEntityWithPartitionAndRowKeyResponse struct {
	// ClientRequestID contains the information returned from the x-ms-client-request-id header response.
	ClientRequestID *string

	// Date contains the information returned from the Date header response.
	Date *time.Time

	// ETag contains the information returned from the ETag header response.
	ETag *string

	// RequestID contains the information returned from the x-ms-request-id header response.
	RequestID *string

	// The other properties of the table entity.
	Value map[string]interface{}

	// Version contains the information returned from the x-ms-version header response.
	Version *string

	// XMSContinuationNextPartitionKey contains the information returned from the x-ms-continuation-NextPartitionKey header response.
	XMSContinuationNextPartitionKey *string

	// XMSContinuationNextRowKey contains the information returned from the x-ms-continuation-NextRowKey header response.
	XMSContinuationNextRowKey *string
}

// ClientQueryResponse contains the response from method Client.Query.
type ClientQueryResponse struct {
	QueryResponse
	// ClientRequestID contains the information returned from the x-ms-client-request-id header response.
	ClientRequestID *string

	// Date contains the information returned from the Date header response.
	Date *time.Time

	// RequestID contains the information returned from the x-ms-request-id header response.
	RequestID *string

	// Version contains the information returned from the x-ms-version header response.
	Version *string

	// XMSContinuationNextTableName contains the information returned from the x-ms-continuation-NextTableName header response.
	XMSContinuationNextTableName *string
}

// ClientSetAccessPolicyResponse contains the response from method Client.SetAccessPolicy.
type ClientSetAccessPolicyResponse struct {
	// ClientRequestID contains the information returned from the x-ms-client-request-id header response.
	ClientRequestID *string

	// Date contains the information returned from the Date header response.
	Date *time.Time

	// RequestID contains the information returned from the x-ms-request-id header response.
	RequestID *string

	// Version contains the information returned from the x-ms-version header response.
	Version *string
}

// ClientUpdateEntityResponse contains the response from method Client.UpdateEntity.
type ClientUpdateEntityResponse struct {
	// ClientRequestID contains the information returned from the x-ms-client-request-id header response.
	ClientRequestID *string

	// Date contains the information returned from the Date header response.
	Date *time.Time

	// ETag contains the information returned from the ETag header response.
	ETag *string

	// RequestID contains the information returned from the x-ms-request-id header response.
	RequestID *string

	// Version contains the information returned from the x-ms-version header response.
	Version *string
}

// ServiceClientGetPropertiesResponse contains the response from method ServiceClient.GetProperties.
type ServiceClientGetPropertiesResponse struct {
	ServiceProperties
	// ClientRequestID contains the information returned from the x-ms-client-request-id header response.
	ClientRequestID *string `xml:"ClientRequestID"`

	// RequestID contains the information returned from the x-ms-request-id header response.
	RequestID *string `xml:"RequestID"`

	// Version contains the information returned from the x-ms-version header response.
	Version *string `xml:"Version"`
}

// ServiceClientGetStatisticsResponse contains the response from method ServiceClient.GetStatistics.
type ServiceClientGetStatisticsResponse struct {
	ServiceStats
	// ClientRequestID contains the information returned from the x-ms-client-request-id header response.
	ClientRequestID *string `xml:"ClientRequestID"`

	// Date contains the information returned from the Date header response.
	Date *time.Time `xml:"Date"`

	// RequestID contains the information returned from the x-ms-request-id header response.
	RequestID *string `xml:"RequestID"`

	// Version contains the information returned from the x-ms-version header response.
	Version *string `xml:"Version"`
}

// ServiceClientSetPropertiesResponse contains the response from method ServiceClient.SetProperties.
type ServiceClientSetPropertiesResponse struct {
	// ClientRequestID contains the information returned from the x-ms-client-request-id header response.
	ClientRequestID *string

	// RequestID contains the information returned from the x-ms-request-id header response.
	RequestID *string

	// Version contains the information returned from the x-ms-version header response.
	Version *string
}
