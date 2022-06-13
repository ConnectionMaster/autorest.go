//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package stringgroup

// EnumClientGetNotExpandableResponse contains the response from method EnumClient.GetNotExpandable.
type EnumClientGetNotExpandableResponse struct {
	// Referenced Color Enum Description.
	Value *Colors
}

// EnumClientGetReferencedConstantResponse contains the response from method EnumClient.GetReferencedConstant.
type EnumClientGetReferencedConstantResponse struct {
	RefColorConstant
}

// EnumClientGetReferencedResponse contains the response from method EnumClient.GetReferenced.
type EnumClientGetReferencedResponse struct {
	// Referenced Color Enum Description.
	Value *Colors
}

// EnumClientPutNotExpandableResponse contains the response from method EnumClient.PutNotExpandable.
type EnumClientPutNotExpandableResponse struct {
	// placeholder for future response values
}

// EnumClientPutReferencedConstantResponse contains the response from method EnumClient.PutReferencedConstant.
type EnumClientPutReferencedConstantResponse struct {
	// placeholder for future response values
}

// EnumClientPutReferencedResponse contains the response from method EnumClient.PutReferenced.
type EnumClientPutReferencedResponse struct {
	// placeholder for future response values
}

// StringClientGetBase64EncodedResponse contains the response from method StringClient.GetBase64Encoded.
type StringClientGetBase64EncodedResponse struct {
	Value []byte
}

// StringClientGetBase64URLEncodedResponse contains the response from method StringClient.GetBase64URLEncoded.
type StringClientGetBase64URLEncodedResponse struct {
	Value []byte
}

// StringClientGetEmptyResponse contains the response from method StringClient.GetEmpty.
type StringClientGetEmptyResponse struct {
	// simple string
	Value *string
}

// StringClientGetMBCSResponse contains the response from method StringClient.GetMBCS.
type StringClientGetMBCSResponse struct {
	// simple string
	Value *string
}

// StringClientGetNotProvidedResponse contains the response from method StringClient.GetNotProvided.
type StringClientGetNotProvidedResponse struct {
	Value *string
}

// StringClientGetNullBase64URLEncodedResponse contains the response from method StringClient.GetNullBase64URLEncoded.
type StringClientGetNullBase64URLEncodedResponse struct {
	Value []byte
}

// StringClientGetNullResponse contains the response from method StringClient.GetNull.
type StringClientGetNullResponse struct {
	Value *string
}

// StringClientGetWhitespaceResponse contains the response from method StringClient.GetWhitespace.
type StringClientGetWhitespaceResponse struct {
	// simple string
	Value *string
}

// StringClientPutBase64URLEncodedResponse contains the response from method StringClient.PutBase64URLEncoded.
type StringClientPutBase64URLEncodedResponse struct {
	// placeholder for future response values
}

// StringClientPutEmptyResponse contains the response from method StringClient.PutEmpty.
type StringClientPutEmptyResponse struct {
	// placeholder for future response values
}

// StringClientPutMBCSResponse contains the response from method StringClient.PutMBCS.
type StringClientPutMBCSResponse struct {
	// placeholder for future response values
}

// StringClientPutNullResponse contains the response from method StringClient.PutNull.
type StringClientPutNullResponse struct {
	// placeholder for future response values
}

// StringClientPutWhitespaceResponse contains the response from method StringClient.PutWhitespace.
type StringClientPutWhitespaceResponse struct {
	// placeholder for future response values
}
