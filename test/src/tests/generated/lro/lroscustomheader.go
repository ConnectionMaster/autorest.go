package lrogroup

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

// LROsCustomHeaderClient is the long-running Operation for AutoRest
type LROsCustomHeaderClient struct {
	ManagementClient
}

// NewLROsCustomHeaderClient creates an instance of the LROsCustomHeaderClient client.
func NewLROsCustomHeaderClient() LROsCustomHeaderClient {
	return NewLROsCustomHeaderClientWithBaseURI(DefaultBaseURI)
}

// NewLROsCustomHeaderClientWithBaseURI creates an instance of the LROsCustomHeaderClient client.
func NewLROsCustomHeaderClientWithBaseURI(baseURI string) LROsCustomHeaderClient {
	return LROsCustomHeaderClient{NewWithBaseURI(baseURI)}
}

// Post202Retry200 x-ms-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 is required message header for all
// requests. Long running post request, service returns a 202 to the initial request, with 'Location' and 'Retry-After'
// headers, Polls return a 200 with a response body after success
//
// product is product to put
func (client LROsCustomHeaderClient) Post202Retry200(product *Product) (result LROsCustomHeaderPost202Retry200Future, err error) {
	req, err := client.Post202Retry200Preparer(product)
	if err != nil {
		err = autorest.NewErrorWithError(err, "lrogroup.LROsCustomHeaderClient", "Post202Retry200", nil, "Failure preparing request")
		return
	}

	result, err = client.Post202Retry200Sender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "lrogroup.LROsCustomHeaderClient", "Post202Retry200", nil, "Failure sending request'result.Response()'")
		return
	}

	return
}

// Post202Retry200Preparer prepares the Post202Retry200 request.
func (client LROsCustomHeaderClient) Post202Retry200Preparer(product *Product) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/lro/customheader/post/202/retry/200"))
	if product != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(product))
	}
	return preparer.Prepare(&http.Request{})
}

// Post202Retry200Sender sends the Post202Retry200 request. The method will close the
// http.Response Body if it receives an error.
func (client LROsCustomHeaderClient) Post202Retry200Sender(req *http.Request) (LROsCustomHeaderPost202Retry200Future, error) {
	sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	future := azure.NewFuture(req)
	_, err := future.Done(sender)
	f := LROsCustomHeaderPost202Retry200Future{Future: future}
	return f, err
}

// Post202Retry200Responder handles the response to the Post202Retry200 request. The method always
// closes the http.Response Body.
func (client LROsCustomHeaderClient) Post202Retry200Responder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// PostAsyncRetrySucceeded x-ms-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 is required message header for
// all requests. Long running post request, service returns a 202 to the initial request, with an entity that contains
// ProvisioningState=’Creating’. Poll the endpoint indicated in the Azure-AsyncOperation header for operation status
//
// product is product to put
func (client LROsCustomHeaderClient) PostAsyncRetrySucceeded(product *Product) (result LROsCustomHeaderPostAsyncRetrySucceededFuture, err error) {
	req, err := client.PostAsyncRetrySucceededPreparer(product)
	if err != nil {
		err = autorest.NewErrorWithError(err, "lrogroup.LROsCustomHeaderClient", "PostAsyncRetrySucceeded", nil, "Failure preparing request")
		return
	}

	result, err = client.PostAsyncRetrySucceededSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "lrogroup.LROsCustomHeaderClient", "PostAsyncRetrySucceeded", nil, "Failure sending request'result.Response()'")
		return
	}

	return
}

// PostAsyncRetrySucceededPreparer prepares the PostAsyncRetrySucceeded request.
func (client LROsCustomHeaderClient) PostAsyncRetrySucceededPreparer(product *Product) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/lro/customheader/postasync/retry/succeeded"))
	if product != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(product))
	}
	return preparer.Prepare(&http.Request{})
}

// PostAsyncRetrySucceededSender sends the PostAsyncRetrySucceeded request. The method will close the
// http.Response Body if it receives an error.
func (client LROsCustomHeaderClient) PostAsyncRetrySucceededSender(req *http.Request) (LROsCustomHeaderPostAsyncRetrySucceededFuture, error) {
	sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	future := azure.NewFuture(req)
	_, err := future.Done(sender)
	f := LROsCustomHeaderPostAsyncRetrySucceededFuture{Future: future}
	return f, err
}

// PostAsyncRetrySucceededResponder handles the response to the PostAsyncRetrySucceeded request. The method always
// closes the http.Response Body.
func (client LROsCustomHeaderClient) PostAsyncRetrySucceededResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Put201CreatingSucceeded200 x-ms-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 is required message header
// for all requests. Long running put request, service returns a 201 to the initial request, with an entity that
// contains ProvisioningState=’Creating’.  Polls return this value until the last poll returns a ‘200’ with
// ProvisioningState=’Succeeded’
//
// product is product to put
func (client LROsCustomHeaderClient) Put201CreatingSucceeded200(product *Product) (result LROsCustomHeaderPut201CreatingSucceeded200Future, err error) {
	req, err := client.Put201CreatingSucceeded200Preparer(product)
	if err != nil {
		err = autorest.NewErrorWithError(err, "lrogroup.LROsCustomHeaderClient", "Put201CreatingSucceeded200", nil, "Failure preparing request")
		return
	}

	result, err = client.Put201CreatingSucceeded200Sender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "lrogroup.LROsCustomHeaderClient", "Put201CreatingSucceeded200", nil, "Failure sending request'result.Response()'")
		return
	}

	return
}

// Put201CreatingSucceeded200Preparer prepares the Put201CreatingSucceeded200 request.
func (client LROsCustomHeaderClient) Put201CreatingSucceeded200Preparer(product *Product) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/lro/customheader/put/201/creating/succeeded/200"))
	if product != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(product))
	}
	return preparer.Prepare(&http.Request{})
}

// Put201CreatingSucceeded200Sender sends the Put201CreatingSucceeded200 request. The method will close the
// http.Response Body if it receives an error.
func (client LROsCustomHeaderClient) Put201CreatingSucceeded200Sender(req *http.Request) (LROsCustomHeaderPut201CreatingSucceeded200Future, error) {
	sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	future := azure.NewFuture(req)
	_, err := future.Done(sender)
	f := LROsCustomHeaderPut201CreatingSucceeded200Future{Future: future}
	return f, err
}

// Put201CreatingSucceeded200Responder handles the response to the Put201CreatingSucceeded200 request. The method always
// closes the http.Response Body.
func (client LROsCustomHeaderClient) Put201CreatingSucceeded200Responder(resp *http.Response) (result Product, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// PutAsyncRetrySucceeded x-ms-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 is required message header for
// all requests. Long running put request, service returns a 200 to the initial request, with an entity that contains
// ProvisioningState=’Creating’. Poll the endpoint indicated in the Azure-AsyncOperation header for operation status
//
// product is product to put
func (client LROsCustomHeaderClient) PutAsyncRetrySucceeded(product *Product) (result LROsCustomHeaderPutAsyncRetrySucceededFuture, err error) {
	req, err := client.PutAsyncRetrySucceededPreparer(product)
	if err != nil {
		err = autorest.NewErrorWithError(err, "lrogroup.LROsCustomHeaderClient", "PutAsyncRetrySucceeded", nil, "Failure preparing request")
		return
	}

	result, err = client.PutAsyncRetrySucceededSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "lrogroup.LROsCustomHeaderClient", "PutAsyncRetrySucceeded", nil, "Failure sending request'result.Response()'")
		return
	}

	return
}

// PutAsyncRetrySucceededPreparer prepares the PutAsyncRetrySucceeded request.
func (client LROsCustomHeaderClient) PutAsyncRetrySucceededPreparer(product *Product) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/lro/customheader/putasync/retry/succeeded"))
	if product != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(product))
	}
	return preparer.Prepare(&http.Request{})
}

// PutAsyncRetrySucceededSender sends the PutAsyncRetrySucceeded request. The method will close the
// http.Response Body if it receives an error.
func (client LROsCustomHeaderClient) PutAsyncRetrySucceededSender(req *http.Request) (LROsCustomHeaderPutAsyncRetrySucceededFuture, error) {
	sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	future := azure.NewFuture(req)
	_, err := future.Done(sender)
	f := LROsCustomHeaderPutAsyncRetrySucceededFuture{Future: future}
	return f, err
}

// PutAsyncRetrySucceededResponder handles the response to the PutAsyncRetrySucceeded request. The method always
// closes the http.Response Body.
func (client LROsCustomHeaderClient) PutAsyncRetrySucceededResponder(resp *http.Response) (result Product, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
