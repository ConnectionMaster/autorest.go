package integergroup

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/date"
	"net/http"
)

// IntGroupClient is the test Infrastructure for AutoRest
type IntGroupClient struct {
	ManagementClient
}

// NewIntGroupClient creates an instance of the IntGroupClient client.
func NewIntGroupClient() IntGroupClient {
	return NewIntGroupClientWithBaseURI(DefaultBaseURI)
}

// NewIntGroupClientWithBaseURI creates an instance of the IntGroupClient client.
func NewIntGroupClientWithBaseURI(baseURI string) IntGroupClient {
	return IntGroupClient{NewWithBaseURI(baseURI)}
}

// GetInvalid get invalid Int value
func (client IntGroupClient) GetInvalid() (result Int32, err error) {
	req, err := client.GetInvalidPreparer()
	if err != nil {
		err = autorest.NewErrorWithError(err, "integergroup.IntGroupClient", "GetInvalid", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetInvalidSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "integergroup.IntGroupClient", "GetInvalid", resp, "Failure sending request")
		return
	}

	result, err = client.GetInvalidResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integergroup.IntGroupClient", "GetInvalid", resp, "Failure responding to request")
	}

	return
}

// GetInvalidPreparer prepares the GetInvalid request.
func (client IntGroupClient) GetInvalidPreparer() (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/int/invalid"))
	return preparer.Prepare(&http.Request{})
}

// GetInvalidSender sends the GetInvalid request. The method will close the
// http.Response Body if it receives an error.
func (client IntGroupClient) GetInvalidSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetInvalidResponder handles the response to the GetInvalid request. The method always
// closes the http.Response Body.
func (client IntGroupClient) GetInvalidResponder(resp *http.Response) (result Int32, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetInvalidUnixTime get invalid Unix time value
func (client IntGroupClient) GetInvalidUnixTime() (result UnixTime, err error) {
	req, err := client.GetInvalidUnixTimePreparer()
	if err != nil {
		err = autorest.NewErrorWithError(err, "integergroup.IntGroupClient", "GetInvalidUnixTime", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetInvalidUnixTimeSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "integergroup.IntGroupClient", "GetInvalidUnixTime", resp, "Failure sending request")
		return
	}

	result, err = client.GetInvalidUnixTimeResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integergroup.IntGroupClient", "GetInvalidUnixTime", resp, "Failure responding to request")
	}

	return
}

// GetInvalidUnixTimePreparer prepares the GetInvalidUnixTime request.
func (client IntGroupClient) GetInvalidUnixTimePreparer() (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/int/invalidunixtime"))
	return preparer.Prepare(&http.Request{})
}

// GetInvalidUnixTimeSender sends the GetInvalidUnixTime request. The method will close the
// http.Response Body if it receives an error.
func (client IntGroupClient) GetInvalidUnixTimeSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetInvalidUnixTimeResponder handles the response to the GetInvalidUnixTime request. The method always
// closes the http.Response Body.
func (client IntGroupClient) GetInvalidUnixTimeResponder(resp *http.Response) (result UnixTime, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetNull get null Int value
func (client IntGroupClient) GetNull() (result Int32, err error) {
	req, err := client.GetNullPreparer()
	if err != nil {
		err = autorest.NewErrorWithError(err, "integergroup.IntGroupClient", "GetNull", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetNullSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "integergroup.IntGroupClient", "GetNull", resp, "Failure sending request")
		return
	}

	result, err = client.GetNullResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integergroup.IntGroupClient", "GetNull", resp, "Failure responding to request")
	}

	return
}

// GetNullPreparer prepares the GetNull request.
func (client IntGroupClient) GetNullPreparer() (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/int/null"))
	return preparer.Prepare(&http.Request{})
}

// GetNullSender sends the GetNull request. The method will close the
// http.Response Body if it receives an error.
func (client IntGroupClient) GetNullSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetNullResponder handles the response to the GetNull request. The method always
// closes the http.Response Body.
func (client IntGroupClient) GetNullResponder(resp *http.Response) (result Int32, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetNullUnixTime get null Unix time value
func (client IntGroupClient) GetNullUnixTime() (result UnixTime, err error) {
	req, err := client.GetNullUnixTimePreparer()
	if err != nil {
		err = autorest.NewErrorWithError(err, "integergroup.IntGroupClient", "GetNullUnixTime", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetNullUnixTimeSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "integergroup.IntGroupClient", "GetNullUnixTime", resp, "Failure sending request")
		return
	}

	result, err = client.GetNullUnixTimeResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integergroup.IntGroupClient", "GetNullUnixTime", resp, "Failure responding to request")
	}

	return
}

// GetNullUnixTimePreparer prepares the GetNullUnixTime request.
func (client IntGroupClient) GetNullUnixTimePreparer() (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/int/nullunixtime"))
	return preparer.Prepare(&http.Request{})
}

// GetNullUnixTimeSender sends the GetNullUnixTime request. The method will close the
// http.Response Body if it receives an error.
func (client IntGroupClient) GetNullUnixTimeSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetNullUnixTimeResponder handles the response to the GetNullUnixTime request. The method always
// closes the http.Response Body.
func (client IntGroupClient) GetNullUnixTimeResponder(resp *http.Response) (result UnixTime, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetOverflowInt32 get overflow Int32 value
func (client IntGroupClient) GetOverflowInt32() (result Int32, err error) {
	req, err := client.GetOverflowInt32Preparer()
	if err != nil {
		err = autorest.NewErrorWithError(err, "integergroup.IntGroupClient", "GetOverflowInt32", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetOverflowInt32Sender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "integergroup.IntGroupClient", "GetOverflowInt32", resp, "Failure sending request")
		return
	}

	result, err = client.GetOverflowInt32Responder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integergroup.IntGroupClient", "GetOverflowInt32", resp, "Failure responding to request")
	}

	return
}

// GetOverflowInt32Preparer prepares the GetOverflowInt32 request.
func (client IntGroupClient) GetOverflowInt32Preparer() (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/int/overflowint32"))
	return preparer.Prepare(&http.Request{})
}

// GetOverflowInt32Sender sends the GetOverflowInt32 request. The method will close the
// http.Response Body if it receives an error.
func (client IntGroupClient) GetOverflowInt32Sender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetOverflowInt32Responder handles the response to the GetOverflowInt32 request. The method always
// closes the http.Response Body.
func (client IntGroupClient) GetOverflowInt32Responder(resp *http.Response) (result Int32, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetOverflowInt64 get overflow Int64 value
func (client IntGroupClient) GetOverflowInt64() (result Int64, err error) {
	req, err := client.GetOverflowInt64Preparer()
	if err != nil {
		err = autorest.NewErrorWithError(err, "integergroup.IntGroupClient", "GetOverflowInt64", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetOverflowInt64Sender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "integergroup.IntGroupClient", "GetOverflowInt64", resp, "Failure sending request")
		return
	}

	result, err = client.GetOverflowInt64Responder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integergroup.IntGroupClient", "GetOverflowInt64", resp, "Failure responding to request")
	}

	return
}

// GetOverflowInt64Preparer prepares the GetOverflowInt64 request.
func (client IntGroupClient) GetOverflowInt64Preparer() (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/int/overflowint64"))
	return preparer.Prepare(&http.Request{})
}

// GetOverflowInt64Sender sends the GetOverflowInt64 request. The method will close the
// http.Response Body if it receives an error.
func (client IntGroupClient) GetOverflowInt64Sender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetOverflowInt64Responder handles the response to the GetOverflowInt64 request. The method always
// closes the http.Response Body.
func (client IntGroupClient) GetOverflowInt64Responder(resp *http.Response) (result Int64, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetUnderflowInt32 get underflow Int32 value
func (client IntGroupClient) GetUnderflowInt32() (result Int32, err error) {
	req, err := client.GetUnderflowInt32Preparer()
	if err != nil {
		err = autorest.NewErrorWithError(err, "integergroup.IntGroupClient", "GetUnderflowInt32", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetUnderflowInt32Sender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "integergroup.IntGroupClient", "GetUnderflowInt32", resp, "Failure sending request")
		return
	}

	result, err = client.GetUnderflowInt32Responder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integergroup.IntGroupClient", "GetUnderflowInt32", resp, "Failure responding to request")
	}

	return
}

// GetUnderflowInt32Preparer prepares the GetUnderflowInt32 request.
func (client IntGroupClient) GetUnderflowInt32Preparer() (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/int/underflowint32"))
	return preparer.Prepare(&http.Request{})
}

// GetUnderflowInt32Sender sends the GetUnderflowInt32 request. The method will close the
// http.Response Body if it receives an error.
func (client IntGroupClient) GetUnderflowInt32Sender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetUnderflowInt32Responder handles the response to the GetUnderflowInt32 request. The method always
// closes the http.Response Body.
func (client IntGroupClient) GetUnderflowInt32Responder(resp *http.Response) (result Int32, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetUnderflowInt64 get underflow Int64 value
func (client IntGroupClient) GetUnderflowInt64() (result Int64, err error) {
	req, err := client.GetUnderflowInt64Preparer()
	if err != nil {
		err = autorest.NewErrorWithError(err, "integergroup.IntGroupClient", "GetUnderflowInt64", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetUnderflowInt64Sender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "integergroup.IntGroupClient", "GetUnderflowInt64", resp, "Failure sending request")
		return
	}

	result, err = client.GetUnderflowInt64Responder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integergroup.IntGroupClient", "GetUnderflowInt64", resp, "Failure responding to request")
	}

	return
}

// GetUnderflowInt64Preparer prepares the GetUnderflowInt64 request.
func (client IntGroupClient) GetUnderflowInt64Preparer() (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/int/underflowint64"))
	return preparer.Prepare(&http.Request{})
}

// GetUnderflowInt64Sender sends the GetUnderflowInt64 request. The method will close the
// http.Response Body if it receives an error.
func (client IntGroupClient) GetUnderflowInt64Sender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetUnderflowInt64Responder handles the response to the GetUnderflowInt64 request. The method always
// closes the http.Response Body.
func (client IntGroupClient) GetUnderflowInt64Responder(resp *http.Response) (result Int64, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetUnixTime get datetime encoded as Unix time value
func (client IntGroupClient) GetUnixTime() (result UnixTime, err error) {
	req, err := client.GetUnixTimePreparer()
	if err != nil {
		err = autorest.NewErrorWithError(err, "integergroup.IntGroupClient", "GetUnixTime", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetUnixTimeSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "integergroup.IntGroupClient", "GetUnixTime", resp, "Failure sending request")
		return
	}

	result, err = client.GetUnixTimeResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integergroup.IntGroupClient", "GetUnixTime", resp, "Failure responding to request")
	}

	return
}

// GetUnixTimePreparer prepares the GetUnixTime request.
func (client IntGroupClient) GetUnixTimePreparer() (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/int/unixtime"))
	return preparer.Prepare(&http.Request{})
}

// GetUnixTimeSender sends the GetUnixTime request. The method will close the
// http.Response Body if it receives an error.
func (client IntGroupClient) GetUnixTimeSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetUnixTimeResponder handles the response to the GetUnixTime request. The method always
// closes the http.Response Body.
func (client IntGroupClient) GetUnixTimeResponder(resp *http.Response) (result UnixTime, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// PutMax32 put max int32 value
//
func (client IntGroupClient) PutMax32(intBody int32) (result autorest.Response, err error) {
	req, err := client.PutMax32Preparer(intBody)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integergroup.IntGroupClient", "PutMax32", nil, "Failure preparing request")
		return
	}

	resp, err := client.PutMax32Sender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "integergroup.IntGroupClient", "PutMax32", resp, "Failure sending request")
		return
	}

	result, err = client.PutMax32Responder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integergroup.IntGroupClient", "PutMax32", resp, "Failure responding to request")
	}

	return
}

// PutMax32Preparer prepares the PutMax32 request.
func (client IntGroupClient) PutMax32Preparer(intBody int32) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/int/max/32"),
		autorest.WithJSON(intBody))
	return preparer.Prepare(&http.Request{})
}

// PutMax32Sender sends the PutMax32 request. The method will close the
// http.Response Body if it receives an error.
func (client IntGroupClient) PutMax32Sender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// PutMax32Responder handles the response to the PutMax32 request. The method always
// closes the http.Response Body.
func (client IntGroupClient) PutMax32Responder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// PutMax64 put max int64 value
//
func (client IntGroupClient) PutMax64(intBody int64) (result autorest.Response, err error) {
	req, err := client.PutMax64Preparer(intBody)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integergroup.IntGroupClient", "PutMax64", nil, "Failure preparing request")
		return
	}

	resp, err := client.PutMax64Sender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "integergroup.IntGroupClient", "PutMax64", resp, "Failure sending request")
		return
	}

	result, err = client.PutMax64Responder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integergroup.IntGroupClient", "PutMax64", resp, "Failure responding to request")
	}

	return
}

// PutMax64Preparer prepares the PutMax64 request.
func (client IntGroupClient) PutMax64Preparer(intBody int64) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/int/max/64"),
		autorest.WithJSON(intBody))
	return preparer.Prepare(&http.Request{})
}

// PutMax64Sender sends the PutMax64 request. The method will close the
// http.Response Body if it receives an error.
func (client IntGroupClient) PutMax64Sender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// PutMax64Responder handles the response to the PutMax64 request. The method always
// closes the http.Response Body.
func (client IntGroupClient) PutMax64Responder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// PutMin32 put min int32 value
//
func (client IntGroupClient) PutMin32(intBody int32) (result autorest.Response, err error) {
	req, err := client.PutMin32Preparer(intBody)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integergroup.IntGroupClient", "PutMin32", nil, "Failure preparing request")
		return
	}

	resp, err := client.PutMin32Sender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "integergroup.IntGroupClient", "PutMin32", resp, "Failure sending request")
		return
	}

	result, err = client.PutMin32Responder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integergroup.IntGroupClient", "PutMin32", resp, "Failure responding to request")
	}

	return
}

// PutMin32Preparer prepares the PutMin32 request.
func (client IntGroupClient) PutMin32Preparer(intBody int32) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/int/min/32"),
		autorest.WithJSON(intBody))
	return preparer.Prepare(&http.Request{})
}

// PutMin32Sender sends the PutMin32 request. The method will close the
// http.Response Body if it receives an error.
func (client IntGroupClient) PutMin32Sender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// PutMin32Responder handles the response to the PutMin32 request. The method always
// closes the http.Response Body.
func (client IntGroupClient) PutMin32Responder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// PutMin64 put min int64 value
//
func (client IntGroupClient) PutMin64(intBody int64) (result autorest.Response, err error) {
	req, err := client.PutMin64Preparer(intBody)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integergroup.IntGroupClient", "PutMin64", nil, "Failure preparing request")
		return
	}

	resp, err := client.PutMin64Sender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "integergroup.IntGroupClient", "PutMin64", resp, "Failure sending request")
		return
	}

	result, err = client.PutMin64Responder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integergroup.IntGroupClient", "PutMin64", resp, "Failure responding to request")
	}

	return
}

// PutMin64Preparer prepares the PutMin64 request.
func (client IntGroupClient) PutMin64Preparer(intBody int64) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/int/min/64"),
		autorest.WithJSON(intBody))
	return preparer.Prepare(&http.Request{})
}

// PutMin64Sender sends the PutMin64 request. The method will close the
// http.Response Body if it receives an error.
func (client IntGroupClient) PutMin64Sender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// PutMin64Responder handles the response to the PutMin64 request. The method always
// closes the http.Response Body.
func (client IntGroupClient) PutMin64Responder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// PutUnixTimeDate put datetime encoded as Unix time
//
func (client IntGroupClient) PutUnixTimeDate(intBody date.UnixTime) (result autorest.Response, err error) {
	req, err := client.PutUnixTimeDatePreparer(intBody)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integergroup.IntGroupClient", "PutUnixTimeDate", nil, "Failure preparing request")
		return
	}

	resp, err := client.PutUnixTimeDateSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "integergroup.IntGroupClient", "PutUnixTimeDate", resp, "Failure sending request")
		return
	}

	result, err = client.PutUnixTimeDateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integergroup.IntGroupClient", "PutUnixTimeDate", resp, "Failure responding to request")
	}

	return
}

// PutUnixTimeDatePreparer prepares the PutUnixTimeDate request.
func (client IntGroupClient) PutUnixTimeDatePreparer(intBody date.UnixTime) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/int/unixtime"),
		autorest.WithJSON(intBody))
	return preparer.Prepare(&http.Request{})
}

// PutUnixTimeDateSender sends the PutUnixTimeDate request. The method will close the
// http.Response Body if it receives an error.
func (client IntGroupClient) PutUnixTimeDateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// PutUnixTimeDateResponder handles the response to the PutUnixTimeDate request. The method always
// closes the http.Response Body.
func (client IntGroupClient) PutUnixTimeDateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}
