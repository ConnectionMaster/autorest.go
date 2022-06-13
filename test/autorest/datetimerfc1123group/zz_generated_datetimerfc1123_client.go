//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package datetimerfc1123group

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"time"
)

// Datetimerfc1123Client contains the methods for the Datetimerfc1123 group.
// Don't use this type directly, use NewDatetimerfc1123Client() instead.
type Datetimerfc1123Client struct {
	pl runtime.Pipeline
}

// NewDatetimerfc1123Client creates a new instance of Datetimerfc1123Client with the specified values.
// pl - the pipeline used for sending requests and handling responses.
func NewDatetimerfc1123Client(pl runtime.Pipeline) *Datetimerfc1123Client {
	client := &Datetimerfc1123Client{
		pl: pl,
	}
	return client
}

// GetInvalid - Get invalid datetime value
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 1.0.0
// options - Datetimerfc1123ClientGetInvalidOptions contains the optional parameters for the Datetimerfc1123Client.GetInvalid
// method.
func (client *Datetimerfc1123Client) GetInvalid(ctx context.Context, options *Datetimerfc1123ClientGetInvalidOptions) (Datetimerfc1123ClientGetInvalidResponse, error) {
	req, err := client.getInvalidCreateRequest(ctx, options)
	if err != nil {
		return Datetimerfc1123ClientGetInvalidResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return Datetimerfc1123ClientGetInvalidResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return Datetimerfc1123ClientGetInvalidResponse{}, runtime.NewResponseError(resp)
	}
	return client.getInvalidHandleResponse(resp)
}

// getInvalidCreateRequest creates the GetInvalid request.
func (client *Datetimerfc1123Client) getInvalidCreateRequest(ctx context.Context, options *Datetimerfc1123ClientGetInvalidOptions) (*policy.Request, error) {
	urlPath := "/datetimerfc1123/invalid"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getInvalidHandleResponse handles the GetInvalid response.
func (client *Datetimerfc1123Client) getInvalidHandleResponse(resp *http.Response) (Datetimerfc1123ClientGetInvalidResponse, error) {
	result := Datetimerfc1123ClientGetInvalidResponse{}
	var aux *timeRFC1123
	if err := runtime.UnmarshalAsJSON(resp, &aux); err != nil {
		return Datetimerfc1123ClientGetInvalidResponse{}, err
	}
	result.Value = (*time.Time)(aux)
	return result, nil
}

// GetNull - Get null datetime value
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 1.0.0
// options - Datetimerfc1123ClientGetNullOptions contains the optional parameters for the Datetimerfc1123Client.GetNull method.
func (client *Datetimerfc1123Client) GetNull(ctx context.Context, options *Datetimerfc1123ClientGetNullOptions) (Datetimerfc1123ClientGetNullResponse, error) {
	req, err := client.getNullCreateRequest(ctx, options)
	if err != nil {
		return Datetimerfc1123ClientGetNullResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return Datetimerfc1123ClientGetNullResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return Datetimerfc1123ClientGetNullResponse{}, runtime.NewResponseError(resp)
	}
	return client.getNullHandleResponse(resp)
}

// getNullCreateRequest creates the GetNull request.
func (client *Datetimerfc1123Client) getNullCreateRequest(ctx context.Context, options *Datetimerfc1123ClientGetNullOptions) (*policy.Request, error) {
	urlPath := "/datetimerfc1123/null"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getNullHandleResponse handles the GetNull response.
func (client *Datetimerfc1123Client) getNullHandleResponse(resp *http.Response) (Datetimerfc1123ClientGetNullResponse, error) {
	result := Datetimerfc1123ClientGetNullResponse{}
	var aux *timeRFC1123
	if err := runtime.UnmarshalAsJSON(resp, &aux); err != nil {
		return Datetimerfc1123ClientGetNullResponse{}, err
	}
	result.Value = (*time.Time)(aux)
	return result, nil
}

// GetOverflow - Get overflow datetime value
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 1.0.0
// options - Datetimerfc1123ClientGetOverflowOptions contains the optional parameters for the Datetimerfc1123Client.GetOverflow
// method.
func (client *Datetimerfc1123Client) GetOverflow(ctx context.Context, options *Datetimerfc1123ClientGetOverflowOptions) (Datetimerfc1123ClientGetOverflowResponse, error) {
	req, err := client.getOverflowCreateRequest(ctx, options)
	if err != nil {
		return Datetimerfc1123ClientGetOverflowResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return Datetimerfc1123ClientGetOverflowResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return Datetimerfc1123ClientGetOverflowResponse{}, runtime.NewResponseError(resp)
	}
	return client.getOverflowHandleResponse(resp)
}

// getOverflowCreateRequest creates the GetOverflow request.
func (client *Datetimerfc1123Client) getOverflowCreateRequest(ctx context.Context, options *Datetimerfc1123ClientGetOverflowOptions) (*policy.Request, error) {
	urlPath := "/datetimerfc1123/overflow"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getOverflowHandleResponse handles the GetOverflow response.
func (client *Datetimerfc1123Client) getOverflowHandleResponse(resp *http.Response) (Datetimerfc1123ClientGetOverflowResponse, error) {
	result := Datetimerfc1123ClientGetOverflowResponse{}
	var aux *timeRFC1123
	if err := runtime.UnmarshalAsJSON(resp, &aux); err != nil {
		return Datetimerfc1123ClientGetOverflowResponse{}, err
	}
	result.Value = (*time.Time)(aux)
	return result, nil
}

// GetUTCLowercaseMaxDateTime - Get max datetime value fri, 31 dec 9999 23:59:59 gmt
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 1.0.0
// options - Datetimerfc1123ClientGetUTCLowercaseMaxDateTimeOptions contains the optional parameters for the Datetimerfc1123Client.GetUTCLowercaseMaxDateTime
// method.
func (client *Datetimerfc1123Client) GetUTCLowercaseMaxDateTime(ctx context.Context, options *Datetimerfc1123ClientGetUTCLowercaseMaxDateTimeOptions) (Datetimerfc1123ClientGetUTCLowercaseMaxDateTimeResponse, error) {
	req, err := client.getUTCLowercaseMaxDateTimeCreateRequest(ctx, options)
	if err != nil {
		return Datetimerfc1123ClientGetUTCLowercaseMaxDateTimeResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return Datetimerfc1123ClientGetUTCLowercaseMaxDateTimeResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return Datetimerfc1123ClientGetUTCLowercaseMaxDateTimeResponse{}, runtime.NewResponseError(resp)
	}
	return client.getUTCLowercaseMaxDateTimeHandleResponse(resp)
}

// getUTCLowercaseMaxDateTimeCreateRequest creates the GetUTCLowercaseMaxDateTime request.
func (client *Datetimerfc1123Client) getUTCLowercaseMaxDateTimeCreateRequest(ctx context.Context, options *Datetimerfc1123ClientGetUTCLowercaseMaxDateTimeOptions) (*policy.Request, error) {
	urlPath := "/datetimerfc1123/max/lowercase"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getUTCLowercaseMaxDateTimeHandleResponse handles the GetUTCLowercaseMaxDateTime response.
func (client *Datetimerfc1123Client) getUTCLowercaseMaxDateTimeHandleResponse(resp *http.Response) (Datetimerfc1123ClientGetUTCLowercaseMaxDateTimeResponse, error) {
	result := Datetimerfc1123ClientGetUTCLowercaseMaxDateTimeResponse{}
	var aux *timeRFC1123
	if err := runtime.UnmarshalAsJSON(resp, &aux); err != nil {
		return Datetimerfc1123ClientGetUTCLowercaseMaxDateTimeResponse{}, err
	}
	result.Value = (*time.Time)(aux)
	return result, nil
}

// GetUTCMinDateTime - Get min datetime value Mon, 1 Jan 0001 00:00:00 GMT
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 1.0.0
// options - Datetimerfc1123ClientGetUTCMinDateTimeOptions contains the optional parameters for the Datetimerfc1123Client.GetUTCMinDateTime
// method.
func (client *Datetimerfc1123Client) GetUTCMinDateTime(ctx context.Context, options *Datetimerfc1123ClientGetUTCMinDateTimeOptions) (Datetimerfc1123ClientGetUTCMinDateTimeResponse, error) {
	req, err := client.getUTCMinDateTimeCreateRequest(ctx, options)
	if err != nil {
		return Datetimerfc1123ClientGetUTCMinDateTimeResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return Datetimerfc1123ClientGetUTCMinDateTimeResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return Datetimerfc1123ClientGetUTCMinDateTimeResponse{}, runtime.NewResponseError(resp)
	}
	return client.getUTCMinDateTimeHandleResponse(resp)
}

// getUTCMinDateTimeCreateRequest creates the GetUTCMinDateTime request.
func (client *Datetimerfc1123Client) getUTCMinDateTimeCreateRequest(ctx context.Context, options *Datetimerfc1123ClientGetUTCMinDateTimeOptions) (*policy.Request, error) {
	urlPath := "/datetimerfc1123/min"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getUTCMinDateTimeHandleResponse handles the GetUTCMinDateTime response.
func (client *Datetimerfc1123Client) getUTCMinDateTimeHandleResponse(resp *http.Response) (Datetimerfc1123ClientGetUTCMinDateTimeResponse, error) {
	result := Datetimerfc1123ClientGetUTCMinDateTimeResponse{}
	var aux *timeRFC1123
	if err := runtime.UnmarshalAsJSON(resp, &aux); err != nil {
		return Datetimerfc1123ClientGetUTCMinDateTimeResponse{}, err
	}
	result.Value = (*time.Time)(aux)
	return result, nil
}

// GetUTCUppercaseMaxDateTime - Get max datetime value FRI, 31 DEC 9999 23:59:59 GMT
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 1.0.0
// options - Datetimerfc1123ClientGetUTCUppercaseMaxDateTimeOptions contains the optional parameters for the Datetimerfc1123Client.GetUTCUppercaseMaxDateTime
// method.
func (client *Datetimerfc1123Client) GetUTCUppercaseMaxDateTime(ctx context.Context, options *Datetimerfc1123ClientGetUTCUppercaseMaxDateTimeOptions) (Datetimerfc1123ClientGetUTCUppercaseMaxDateTimeResponse, error) {
	req, err := client.getUTCUppercaseMaxDateTimeCreateRequest(ctx, options)
	if err != nil {
		return Datetimerfc1123ClientGetUTCUppercaseMaxDateTimeResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return Datetimerfc1123ClientGetUTCUppercaseMaxDateTimeResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return Datetimerfc1123ClientGetUTCUppercaseMaxDateTimeResponse{}, runtime.NewResponseError(resp)
	}
	return client.getUTCUppercaseMaxDateTimeHandleResponse(resp)
}

// getUTCUppercaseMaxDateTimeCreateRequest creates the GetUTCUppercaseMaxDateTime request.
func (client *Datetimerfc1123Client) getUTCUppercaseMaxDateTimeCreateRequest(ctx context.Context, options *Datetimerfc1123ClientGetUTCUppercaseMaxDateTimeOptions) (*policy.Request, error) {
	urlPath := "/datetimerfc1123/max/uppercase"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getUTCUppercaseMaxDateTimeHandleResponse handles the GetUTCUppercaseMaxDateTime response.
func (client *Datetimerfc1123Client) getUTCUppercaseMaxDateTimeHandleResponse(resp *http.Response) (Datetimerfc1123ClientGetUTCUppercaseMaxDateTimeResponse, error) {
	result := Datetimerfc1123ClientGetUTCUppercaseMaxDateTimeResponse{}
	var aux *timeRFC1123
	if err := runtime.UnmarshalAsJSON(resp, &aux); err != nil {
		return Datetimerfc1123ClientGetUTCUppercaseMaxDateTimeResponse{}, err
	}
	result.Value = (*time.Time)(aux)
	return result, nil
}

// GetUnderflow - Get underflow datetime value
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 1.0.0
// options - Datetimerfc1123ClientGetUnderflowOptions contains the optional parameters for the Datetimerfc1123Client.GetUnderflow
// method.
func (client *Datetimerfc1123Client) GetUnderflow(ctx context.Context, options *Datetimerfc1123ClientGetUnderflowOptions) (Datetimerfc1123ClientGetUnderflowResponse, error) {
	req, err := client.getUnderflowCreateRequest(ctx, options)
	if err != nil {
		return Datetimerfc1123ClientGetUnderflowResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return Datetimerfc1123ClientGetUnderflowResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return Datetimerfc1123ClientGetUnderflowResponse{}, runtime.NewResponseError(resp)
	}
	return client.getUnderflowHandleResponse(resp)
}

// getUnderflowCreateRequest creates the GetUnderflow request.
func (client *Datetimerfc1123Client) getUnderflowCreateRequest(ctx context.Context, options *Datetimerfc1123ClientGetUnderflowOptions) (*policy.Request, error) {
	urlPath := "/datetimerfc1123/underflow"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getUnderflowHandleResponse handles the GetUnderflow response.
func (client *Datetimerfc1123Client) getUnderflowHandleResponse(resp *http.Response) (Datetimerfc1123ClientGetUnderflowResponse, error) {
	result := Datetimerfc1123ClientGetUnderflowResponse{}
	var aux *timeRFC1123
	if err := runtime.UnmarshalAsJSON(resp, &aux); err != nil {
		return Datetimerfc1123ClientGetUnderflowResponse{}, err
	}
	result.Value = (*time.Time)(aux)
	return result, nil
}

// PutUTCMaxDateTime - Put max datetime value Fri, 31 Dec 9999 23:59:59 GMT
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 1.0.0
// datetimeBody - datetime body
// options - Datetimerfc1123ClientPutUTCMaxDateTimeOptions contains the optional parameters for the Datetimerfc1123Client.PutUTCMaxDateTime
// method.
func (client *Datetimerfc1123Client) PutUTCMaxDateTime(ctx context.Context, datetimeBody time.Time, options *Datetimerfc1123ClientPutUTCMaxDateTimeOptions) (Datetimerfc1123ClientPutUTCMaxDateTimeResponse, error) {
	req, err := client.putUTCMaxDateTimeCreateRequest(ctx, datetimeBody, options)
	if err != nil {
		return Datetimerfc1123ClientPutUTCMaxDateTimeResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return Datetimerfc1123ClientPutUTCMaxDateTimeResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return Datetimerfc1123ClientPutUTCMaxDateTimeResponse{}, runtime.NewResponseError(resp)
	}
	return Datetimerfc1123ClientPutUTCMaxDateTimeResponse{}, nil
}

// putUTCMaxDateTimeCreateRequest creates the PutUTCMaxDateTime request.
func (client *Datetimerfc1123Client) putUTCMaxDateTimeCreateRequest(ctx context.Context, datetimeBody time.Time, options *Datetimerfc1123ClientPutUTCMaxDateTimeOptions) (*policy.Request, error) {
	urlPath := "/datetimerfc1123/max"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	aux := timeRFC1123(datetimeBody)
	return req, runtime.MarshalAsJSON(req, aux)
}

// PutUTCMinDateTime - Put min datetime value Mon, 1 Jan 0001 00:00:00 GMT
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 1.0.0
// datetimeBody - datetime body
// options - Datetimerfc1123ClientPutUTCMinDateTimeOptions contains the optional parameters for the Datetimerfc1123Client.PutUTCMinDateTime
// method.
func (client *Datetimerfc1123Client) PutUTCMinDateTime(ctx context.Context, datetimeBody time.Time, options *Datetimerfc1123ClientPutUTCMinDateTimeOptions) (Datetimerfc1123ClientPutUTCMinDateTimeResponse, error) {
	req, err := client.putUTCMinDateTimeCreateRequest(ctx, datetimeBody, options)
	if err != nil {
		return Datetimerfc1123ClientPutUTCMinDateTimeResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return Datetimerfc1123ClientPutUTCMinDateTimeResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return Datetimerfc1123ClientPutUTCMinDateTimeResponse{}, runtime.NewResponseError(resp)
	}
	return Datetimerfc1123ClientPutUTCMinDateTimeResponse{}, nil
}

// putUTCMinDateTimeCreateRequest creates the PutUTCMinDateTime request.
func (client *Datetimerfc1123Client) putUTCMinDateTimeCreateRequest(ctx context.Context, datetimeBody time.Time, options *Datetimerfc1123ClientPutUTCMinDateTimeOptions) (*policy.Request, error) {
	urlPath := "/datetimerfc1123/min"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	aux := timeRFC1123(datetimeBody)
	return req, runtime.MarshalAsJSON(req, aux)
}
