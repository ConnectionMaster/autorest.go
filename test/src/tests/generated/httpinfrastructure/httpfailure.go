package httpinfrastructuregroup

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// HTTPFailureClient is the test Infrastructure for AutoRest
type HTTPFailureClient struct {
	ManagementClient
}

// NewHTTPFailureClient creates an instance of the HTTPFailureClient client.
func NewHTTPFailureClient() HTTPFailureClient {
	return NewHTTPFailureClientWithBaseURI(DefaultBaseURI)
}

// NewHTTPFailureClientWithBaseURI creates an instance of the HTTPFailureClient client.
func NewHTTPFailureClientWithBaseURI(baseURI string) HTTPFailureClient {
	return HTTPFailureClient{NewWithBaseURI(baseURI)}
}

// GetEmptyError get empty error form server
func (client HTTPFailureClient) GetEmptyError() (result Bool, err error) {
	req, err := client.GetEmptyErrorPreparer()
	if err != nil {
		err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPFailureClient", "GetEmptyError", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetEmptyErrorSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPFailureClient", "GetEmptyError", resp, "Failure sending request")
		return
	}

	result, err = client.GetEmptyErrorResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPFailureClient", "GetEmptyError", resp, "Failure responding to request")
	}

	return
}

// GetEmptyErrorPreparer prepares the GetEmptyError request.
func (client HTTPFailureClient) GetEmptyErrorPreparer() (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/http/failure/emptybody/error"))
	return preparer.Prepare(&http.Request{})
}

// GetEmptyErrorSender sends the GetEmptyError request. The method will close the
// http.Response Body if it receives an error.
func (client HTTPFailureClient) GetEmptyErrorSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetEmptyErrorResponder handles the response to the GetEmptyError request. The method always
// closes the http.Response Body.
func (client HTTPFailureClient) GetEmptyErrorResponder(resp *http.Response) (result Bool, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetNoModelEmpty get empty response from server
func (client HTTPFailureClient) GetNoModelEmpty() (result Bool, err error) {
	req, err := client.GetNoModelEmptyPreparer()
	if err != nil {
		err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPFailureClient", "GetNoModelEmpty", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetNoModelEmptySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPFailureClient", "GetNoModelEmpty", resp, "Failure sending request")
		return
	}

	result, err = client.GetNoModelEmptyResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPFailureClient", "GetNoModelEmpty", resp, "Failure responding to request")
	}

	return
}

// GetNoModelEmptyPreparer prepares the GetNoModelEmpty request.
func (client HTTPFailureClient) GetNoModelEmptyPreparer() (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/http/failure/nomodel/empty"))
	return preparer.Prepare(&http.Request{})
}

// GetNoModelEmptySender sends the GetNoModelEmpty request. The method will close the
// http.Response Body if it receives an error.
func (client HTTPFailureClient) GetNoModelEmptySender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetNoModelEmptyResponder handles the response to the GetNoModelEmpty request. The method always
// closes the http.Response Body.
func (client HTTPFailureClient) GetNoModelEmptyResponder(resp *http.Response) (result Bool, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetNoModelError get empty error form server
func (client HTTPFailureClient) GetNoModelError() (result Bool, err error) {
	req, err := client.GetNoModelErrorPreparer()
	if err != nil {
		err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPFailureClient", "GetNoModelError", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetNoModelErrorSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPFailureClient", "GetNoModelError", resp, "Failure sending request")
		return
	}

	result, err = client.GetNoModelErrorResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPFailureClient", "GetNoModelError", resp, "Failure responding to request")
	}

	return
}

// GetNoModelErrorPreparer prepares the GetNoModelError request.
func (client HTTPFailureClient) GetNoModelErrorPreparer() (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/http/failure/nomodel/error"))
	return preparer.Prepare(&http.Request{})
}

// GetNoModelErrorSender sends the GetNoModelError request. The method will close the
// http.Response Body if it receives an error.
func (client HTTPFailureClient) GetNoModelErrorSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetNoModelErrorResponder handles the response to the GetNoModelError request. The method always
// closes the http.Response Body.
func (client HTTPFailureClient) GetNoModelErrorResponder(resp *http.Response) (result Bool, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
