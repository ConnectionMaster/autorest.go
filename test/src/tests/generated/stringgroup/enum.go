package stringgroup

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// EnumClient is the test Infrastructure for AutoRest Swagger BAT
type EnumClient struct {
	BaseClient
}

// NewEnumClient creates an instance of the EnumClient client.
func NewEnumClient() EnumClient {
	return NewEnumClientWithBaseURI(DefaultBaseURI)
}

// NewEnumClientWithBaseURI creates an instance of the EnumClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewEnumClientWithBaseURI(baseURI string) EnumClient {
	return EnumClient{NewWithBaseURI(baseURI)}
}

// GetNotExpandable get enum value 'red color' from enumeration of 'red color', 'green-color', 'blue_color'.
func (client EnumClient) GetNotExpandable(ctx context.Context) (result StringModel, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EnumClient.GetNotExpandable")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetNotExpandablePreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "stringgroup.EnumClient", "GetNotExpandable", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetNotExpandableSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "stringgroup.EnumClient", "GetNotExpandable", resp, "Failure sending request")
		return
	}

	result, err = client.GetNotExpandableResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "stringgroup.EnumClient", "GetNotExpandable", resp, "Failure responding to request")
	}

	return
}

// GetNotExpandablePreparer prepares the GetNotExpandable request.
func (client EnumClient) GetNotExpandablePreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/string/enum/notExpandable"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetNotExpandableSender sends the GetNotExpandable request. The method will close the
// http.Response Body if it receives an error.
func (client EnumClient) GetNotExpandableSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetNotExpandableResponder handles the response to the GetNotExpandable request. The method always
// closes the http.Response Body.
func (client EnumClient) GetNotExpandableResponder(resp *http.Response) (result StringModel, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetReferenced get enum value 'red color' from enumeration of 'red color', 'green-color', 'blue_color'.
func (client EnumClient) GetReferenced(ctx context.Context) (result StringModel, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EnumClient.GetReferenced")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetReferencedPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "stringgroup.EnumClient", "GetReferenced", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetReferencedSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "stringgroup.EnumClient", "GetReferenced", resp, "Failure sending request")
		return
	}

	result, err = client.GetReferencedResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "stringgroup.EnumClient", "GetReferenced", resp, "Failure responding to request")
	}

	return
}

// GetReferencedPreparer prepares the GetReferenced request.
func (client EnumClient) GetReferencedPreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/string/enum/Referenced"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetReferencedSender sends the GetReferenced request. The method will close the
// http.Response Body if it receives an error.
func (client EnumClient) GetReferencedSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetReferencedResponder handles the response to the GetReferenced request. The method always
// closes the http.Response Body.
func (client EnumClient) GetReferencedResponder(resp *http.Response) (result StringModel, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetReferencedConstant get value 'green-color' from the constant.
func (client EnumClient) GetReferencedConstant(ctx context.Context) (result RefColorConstant, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EnumClient.GetReferencedConstant")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetReferencedConstantPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "stringgroup.EnumClient", "GetReferencedConstant", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetReferencedConstantSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "stringgroup.EnumClient", "GetReferencedConstant", resp, "Failure sending request")
		return
	}

	result, err = client.GetReferencedConstantResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "stringgroup.EnumClient", "GetReferencedConstant", resp, "Failure responding to request")
	}

	return
}

// GetReferencedConstantPreparer prepares the GetReferencedConstant request.
func (client EnumClient) GetReferencedConstantPreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/string/enum/ReferencedConstant"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetReferencedConstantSender sends the GetReferencedConstant request. The method will close the
// http.Response Body if it receives an error.
func (client EnumClient) GetReferencedConstantSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetReferencedConstantResponder handles the response to the GetReferencedConstant request. The method always
// closes the http.Response Body.
func (client EnumClient) GetReferencedConstantResponder(resp *http.Response) (result RefColorConstant, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// PutNotExpandable sends value 'red color' from enumeration of 'red color', 'green-color', 'blue_color'
func (client EnumClient) PutNotExpandable(ctx context.Context, stringBody Colors) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EnumClient.PutNotExpandable")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.PutNotExpandablePreparer(ctx, stringBody)
	if err != nil {
		err = autorest.NewErrorWithError(err, "stringgroup.EnumClient", "PutNotExpandable", nil, "Failure preparing request")
		return
	}

	resp, err := client.PutNotExpandableSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "stringgroup.EnumClient", "PutNotExpandable", resp, "Failure sending request")
		return
	}

	result, err = client.PutNotExpandableResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "stringgroup.EnumClient", "PutNotExpandable", resp, "Failure responding to request")
	}

	return
}

// PutNotExpandablePreparer prepares the PutNotExpandable request.
func (client EnumClient) PutNotExpandablePreparer(ctx context.Context, stringBody Colors) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/string/enum/notExpandable"),
		autorest.WithJSON(stringBody))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PutNotExpandableSender sends the PutNotExpandable request. The method will close the
// http.Response Body if it receives an error.
func (client EnumClient) PutNotExpandableSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// PutNotExpandableResponder handles the response to the PutNotExpandable request. The method always
// closes the http.Response Body.
func (client EnumClient) PutNotExpandableResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// PutReferenced sends value 'red color' from enumeration of 'red color', 'green-color', 'blue_color'
func (client EnumClient) PutReferenced(ctx context.Context, enumStringBody Colors) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EnumClient.PutReferenced")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.PutReferencedPreparer(ctx, enumStringBody)
	if err != nil {
		err = autorest.NewErrorWithError(err, "stringgroup.EnumClient", "PutReferenced", nil, "Failure preparing request")
		return
	}

	resp, err := client.PutReferencedSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "stringgroup.EnumClient", "PutReferenced", resp, "Failure sending request")
		return
	}

	result, err = client.PutReferencedResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "stringgroup.EnumClient", "PutReferenced", resp, "Failure responding to request")
	}

	return
}

// PutReferencedPreparer prepares the PutReferenced request.
func (client EnumClient) PutReferencedPreparer(ctx context.Context, enumStringBody Colors) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/string/enum/Referenced"),
		autorest.WithJSON(enumStringBody))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PutReferencedSender sends the PutReferenced request. The method will close the
// http.Response Body if it receives an error.
func (client EnumClient) PutReferencedSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// PutReferencedResponder handles the response to the PutReferenced request. The method always
// closes the http.Response Body.
func (client EnumClient) PutReferencedResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// PutReferencedConstant sends value 'green-color' from a constant
func (client EnumClient) PutReferencedConstant(ctx context.Context, enumStringBody RefColorConstant) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EnumClient.PutReferencedConstant")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: enumStringBody,
			Constraints: []validation.Constraint{{Target: "enumStringBody.ColorConstant", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("stringgroup.EnumClient", "PutReferencedConstant", err.Error())
	}

	req, err := client.PutReferencedConstantPreparer(ctx, enumStringBody)
	if err != nil {
		err = autorest.NewErrorWithError(err, "stringgroup.EnumClient", "PutReferencedConstant", nil, "Failure preparing request")
		return
	}

	resp, err := client.PutReferencedConstantSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "stringgroup.EnumClient", "PutReferencedConstant", resp, "Failure sending request")
		return
	}

	result, err = client.PutReferencedConstantResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "stringgroup.EnumClient", "PutReferencedConstant", resp, "Failure responding to request")
	}

	return
}

// PutReferencedConstantPreparer prepares the PutReferencedConstant request.
func (client EnumClient) PutReferencedConstantPreparer(ctx context.Context, enumStringBody RefColorConstant) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/string/enum/ReferencedConstant"),
		autorest.WithJSON(enumStringBody))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PutReferencedConstantSender sends the PutReferencedConstant request. The method will close the
// http.Response Body if it receives an error.
func (client EnumClient) PutReferencedConstantSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// PutReferencedConstantResponder handles the response to the PutReferencedConstant request. The method always
// closes the http.Response Body.
func (client EnumClient) PutReferencedConstantResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}