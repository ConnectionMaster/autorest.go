//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package azurespecialsgroup

// APIVersionDefaultClientGetMethodGlobalNotProvidedValidOptions contains the optional parameters for the APIVersionDefaultClient.GetMethodGlobalNotProvidedValid
// method.
type APIVersionDefaultClientGetMethodGlobalNotProvidedValidOptions struct {
	// placeholder for future optional parameters
}

// APIVersionDefaultClientGetMethodGlobalValidOptions contains the optional parameters for the APIVersionDefaultClient.GetMethodGlobalValid
// method.
type APIVersionDefaultClientGetMethodGlobalValidOptions struct {
	// placeholder for future optional parameters
}

// APIVersionDefaultClientGetPathGlobalValidOptions contains the optional parameters for the APIVersionDefaultClient.GetPathGlobalValid
// method.
type APIVersionDefaultClientGetPathGlobalValidOptions struct {
	// placeholder for future optional parameters
}

// APIVersionDefaultClientGetSwaggerGlobalValidOptions contains the optional parameters for the APIVersionDefaultClient.GetSwaggerGlobalValid
// method.
type APIVersionDefaultClientGetSwaggerGlobalValidOptions struct {
	// placeholder for future optional parameters
}

// APIVersionLocalClientGetMethodLocalNullOptions contains the optional parameters for the APIVersionLocalClient.GetMethodLocalNull
// method.
type APIVersionLocalClientGetMethodLocalNullOptions struct {
	// This should appear as a method parameter, use value null, this should result in no serialized parameter
	APIVersion *string
}

// APIVersionLocalClientGetMethodLocalValidOptions contains the optional parameters for the APIVersionLocalClient.GetMethodLocalValid
// method.
type APIVersionLocalClientGetMethodLocalValidOptions struct {
	// placeholder for future optional parameters
}

// APIVersionLocalClientGetPathLocalValidOptions contains the optional parameters for the APIVersionLocalClient.GetPathLocalValid
// method.
type APIVersionLocalClientGetPathLocalValidOptions struct {
	// placeholder for future optional parameters
}

// APIVersionLocalClientGetSwaggerLocalValidOptions contains the optional parameters for the APIVersionLocalClient.GetSwaggerLocalValid
// method.
type APIVersionLocalClientGetSwaggerLocalValidOptions struct {
	// placeholder for future optional parameters
}

type Error struct {
	// REQUIRED
	ConstantID *int32  `json:"constantId,omitempty"`
	Message    *string `json:"message,omitempty"`
	Status     *int32  `json:"status,omitempty"`
}

// HeaderClientCustomNamedRequestIDHeadOptions contains the optional parameters for the HeaderClient.CustomNamedRequestIDHead
// method.
type HeaderClientCustomNamedRequestIDHeadOptions struct {
	// placeholder for future optional parameters
}

// HeaderClientCustomNamedRequestIDOptions contains the optional parameters for the HeaderClient.CustomNamedRequestID method.
type HeaderClientCustomNamedRequestIDOptions struct {
	// placeholder for future optional parameters
}

// HeaderClientCustomNamedRequestIDParamGroupingOptions contains the optional parameters for the HeaderClient.CustomNamedRequestIDParamGrouping
// method.
type HeaderClientCustomNamedRequestIDParamGroupingOptions struct {
	// placeholder for future optional parameters
}

// HeaderClientCustomNamedRequestIDParamGroupingParameters contains a group of parameters for the HeaderClient.CustomNamedRequestIDParamGrouping
// method.
type HeaderClientCustomNamedRequestIDParamGroupingParameters struct {
	// The fooRequestId
	FooClientRequestID string
}

// ODataClientGetWithFilterOptions contains the optional parameters for the ODataClient.GetWithFilter method.
type ODataClientGetWithFilterOptions struct {
	// The filter parameter with value '$filter=id gt 5 and name eq 'foo''.
	Filter *string
	// The orderby parameter with value id.
	Orderby *string
	// The top parameter with value 10.
	Top *int32
}

type ODataFilter struct {
	ID   *int32  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// SkipURLEncodingClientGetMethodPathValidOptions contains the optional parameters for the SkipURLEncodingClient.GetMethodPathValid
// method.
type SkipURLEncodingClientGetMethodPathValidOptions struct {
	// placeholder for future optional parameters
}

// SkipURLEncodingClientGetMethodQueryNullOptions contains the optional parameters for the SkipURLEncodingClient.GetMethodQueryNull
// method.
type SkipURLEncodingClientGetMethodQueryNullOptions struct {
	// Unencoded query parameter with value null
	Q1 *string
}

// SkipURLEncodingClientGetMethodQueryValidOptions contains the optional parameters for the SkipURLEncodingClient.GetMethodQueryValid
// method.
type SkipURLEncodingClientGetMethodQueryValidOptions struct {
	// placeholder for future optional parameters
}

// SkipURLEncodingClientGetPathQueryValidOptions contains the optional parameters for the SkipURLEncodingClient.GetPathQueryValid
// method.
type SkipURLEncodingClientGetPathQueryValidOptions struct {
	// placeholder for future optional parameters
}

// SkipURLEncodingClientGetPathValidOptions contains the optional parameters for the SkipURLEncodingClient.GetPathValid method.
type SkipURLEncodingClientGetPathValidOptions struct {
	// placeholder for future optional parameters
}

// SkipURLEncodingClientGetSwaggerPathValidOptions contains the optional parameters for the SkipURLEncodingClient.GetSwaggerPathValid
// method.
type SkipURLEncodingClientGetSwaggerPathValidOptions struct {
	// placeholder for future optional parameters
}

// SkipURLEncodingClientGetSwaggerQueryValidOptions contains the optional parameters for the SkipURLEncodingClient.GetSwaggerQueryValid
// method.
type SkipURLEncodingClientGetSwaggerQueryValidOptions struct {
	// placeholder for future optional parameters
}

// SubscriptionInCredentialsClientPostMethodGlobalNotProvidedValidOptions contains the optional parameters for the SubscriptionInCredentialsClient.PostMethodGlobalNotProvidedValid
// method.
type SubscriptionInCredentialsClientPostMethodGlobalNotProvidedValidOptions struct {
	// placeholder for future optional parameters
}

// SubscriptionInCredentialsClientPostMethodGlobalNullOptions contains the optional parameters for the SubscriptionInCredentialsClient.PostMethodGlobalNull
// method.
type SubscriptionInCredentialsClientPostMethodGlobalNullOptions struct {
	// placeholder for future optional parameters
}

// SubscriptionInCredentialsClientPostMethodGlobalValidOptions contains the optional parameters for the SubscriptionInCredentialsClient.PostMethodGlobalValid
// method.
type SubscriptionInCredentialsClientPostMethodGlobalValidOptions struct {
	// placeholder for future optional parameters
}

// SubscriptionInCredentialsClientPostPathGlobalValidOptions contains the optional parameters for the SubscriptionInCredentialsClient.PostPathGlobalValid
// method.
type SubscriptionInCredentialsClientPostPathGlobalValidOptions struct {
	// placeholder for future optional parameters
}

// SubscriptionInCredentialsClientPostSwaggerGlobalValidOptions contains the optional parameters for the SubscriptionInCredentialsClient.PostSwaggerGlobalValid
// method.
type SubscriptionInCredentialsClientPostSwaggerGlobalValidOptions struct {
	// placeholder for future optional parameters
}

// SubscriptionInMethodClientPostMethodLocalNullOptions contains the optional parameters for the SubscriptionInMethodClient.PostMethodLocalNull
// method.
type SubscriptionInMethodClientPostMethodLocalNullOptions struct {
	// placeholder for future optional parameters
}

// SubscriptionInMethodClientPostMethodLocalValidOptions contains the optional parameters for the SubscriptionInMethodClient.PostMethodLocalValid
// method.
type SubscriptionInMethodClientPostMethodLocalValidOptions struct {
	// placeholder for future optional parameters
}

// SubscriptionInMethodClientPostPathLocalValidOptions contains the optional parameters for the SubscriptionInMethodClient.PostPathLocalValid
// method.
type SubscriptionInMethodClientPostPathLocalValidOptions struct {
	// placeholder for future optional parameters
}

// SubscriptionInMethodClientPostSwaggerLocalValidOptions contains the optional parameters for the SubscriptionInMethodClient.PostSwaggerLocalValid
// method.
type SubscriptionInMethodClientPostSwaggerLocalValidOptions struct {
	// placeholder for future optional parameters
}

// XMSClientRequestIDClientGetOptions contains the optional parameters for the XMSClientRequestIDClient.Get method.
type XMSClientRequestIDClientGetOptions struct {
	// placeholder for future optional parameters
}

// XMSClientRequestIDClientParamGetOptions contains the optional parameters for the XMSClientRequestIDClient.ParamGet method.
type XMSClientRequestIDClientParamGetOptions struct {
	// placeholder for future optional parameters
}
