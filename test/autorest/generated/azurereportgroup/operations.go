// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azurereportgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// Operations contains the methods for the Operations group.
type Operations interface {
	// GetReport - Get test coverage report
	GetReport(ctx context.Context, operationsGetReportOptions *OperationsGetReportOptions) (*MapOfInt32Response, error)
}

// operations implements the Operations interface.
type operations struct {
	*Client
	Host string
}

// GetReport - Get test coverage report
func (client *operations) GetReport(ctx context.Context, operationsGetReportOptions *OperationsGetReportOptions) (*MapOfInt32Response, error) {
	req, err := client.getReportCreateRequest(operationsGetReportOptions)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getReportHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getReportCreateRequest creates the GetReport request.
func (client *operations) getReportCreateRequest(operationsGetReportOptions *OperationsGetReportOptions) (*azcore.Request, error) {
	urlPath := "/report/azure"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	if operationsGetReportOptions != nil && operationsGetReportOptions.Qualifier != nil {
		query.Set("qualifier", *operationsGetReportOptions.Qualifier)
	}
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getReportHandleResponse handles the GetReport response.
func (client *operations) getReportHandleResponse(resp *azcore.Response) (*MapOfInt32Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := MapOfInt32Response{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}