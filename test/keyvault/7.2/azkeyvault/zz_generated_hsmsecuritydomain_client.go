// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azkeyvault

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"strings"
	"time"
)

// HSMSecurityDomainClient contains the methods for the HSMSecurityDomain group.
// Don't use this type directly, use NewHSMSecurityDomainClient() instead.
type HSMSecurityDomainClient struct {
	con *Connection
}

// NewHSMSecurityDomainClient creates a new instance of HSMSecurityDomainClient with the specified values.
func NewHSMSecurityDomainClient(con *Connection) *HSMSecurityDomainClient {
	return &HSMSecurityDomainClient{con: con}
}

// BeginDownload - Retrieves the Security Domain from the managed HSM. Calling this endpoint can be used to activate a provisioned managed HSM resource.
// If the operation fails it returns the *KeyVaultError error type.
func (client *HSMSecurityDomainClient) BeginDownload(ctx context.Context, vaultBaseURL string, certificateInfoObject CertificateInfoObject, options *HSMSecurityDomainBeginDownloadOptions) (SecurityDomainObjectPollerResponse, error) {
	resp, err := client.download(ctx, vaultBaseURL, certificateInfoObject, options)
	if err != nil {
		return SecurityDomainObjectPollerResponse{}, err
	}
	result := SecurityDomainObjectPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := azcore.NewLROPoller("HSMSecurityDomainClient.Download", resp, client.con.Pipeline(), client.downloadHandleError)
	if err != nil {
		return SecurityDomainObjectPollerResponse{}, err
	}
	poller := &securityDomainObjectPoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (SecurityDomainObjectResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeDownload creates a new SecurityDomainObjectPoller from the specified resume token.
// token - The value must come from a previous call to SecurityDomainObjectPoller.ResumeToken().
func (client *HSMSecurityDomainClient) ResumeDownload(ctx context.Context, token string) (SecurityDomainObjectPollerResponse, error) {
	pt, err := azcore.NewLROPollerFromResumeToken("HSMSecurityDomainClient.Download", token, client.con.Pipeline(), client.downloadHandleError)
	if err != nil {
		return SecurityDomainObjectPollerResponse{}, err
	}
	poller := &securityDomainObjectPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return SecurityDomainObjectPollerResponse{}, err
	}
	result := SecurityDomainObjectPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (SecurityDomainObjectResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Download - Retrieves the Security Domain from the managed HSM. Calling this endpoint can be used to activate a provisioned managed HSM resource.
// If the operation fails it returns the *KeyVaultError error type.
func (client *HSMSecurityDomainClient) download(ctx context.Context, vaultBaseURL string, certificateInfoObject CertificateInfoObject, options *HSMSecurityDomainBeginDownloadOptions) (*azcore.Response, error) {
	req, err := client.downloadCreateRequest(ctx, vaultBaseURL, certificateInfoObject, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusAccepted) {
		return nil, client.downloadHandleError(resp)
	}
	return resp, nil
}

// downloadCreateRequest creates the Download request.
func (client *HSMSecurityDomainClient) downloadCreateRequest(ctx context.Context, vaultBaseURL string, certificateInfoObject CertificateInfoObject, options *HSMSecurityDomainBeginDownloadOptions) (*azcore.Request, error) {
	host := "{vaultBaseUrl}"
	host = strings.ReplaceAll(host, "{vaultBaseUrl}", vaultBaseURL)
	urlPath := "/securitydomain/download"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "7.2")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(certificateInfoObject)
}

// downloadHandleError handles the Download error response.
func (client *HSMSecurityDomainClient) downloadHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := KeyVaultError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// DownloadPending - Retrieves the Security Domain download operation status
// If the operation fails it returns the *KeyVaultError error type.
func (client *HSMSecurityDomainClient) DownloadPending(ctx context.Context, vaultBaseURL string, options *HSMSecurityDomainDownloadPendingOptions) (SecurityDomainOperationStatusResponse, error) {
	req, err := client.downloadPendingCreateRequest(ctx, vaultBaseURL, options)
	if err != nil {
		return SecurityDomainOperationStatusResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return SecurityDomainOperationStatusResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return SecurityDomainOperationStatusResponse{}, client.downloadPendingHandleError(resp)
	}
	return client.downloadPendingHandleResponse(resp)
}

// downloadPendingCreateRequest creates the DownloadPending request.
func (client *HSMSecurityDomainClient) downloadPendingCreateRequest(ctx context.Context, vaultBaseURL string, options *HSMSecurityDomainDownloadPendingOptions) (*azcore.Request, error) {
	host := "{vaultBaseUrl}"
	host = strings.ReplaceAll(host, "{vaultBaseUrl}", vaultBaseURL)
	urlPath := "/securitydomain/download/pending"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// downloadPendingHandleResponse handles the DownloadPending response.
func (client *HSMSecurityDomainClient) downloadPendingHandleResponse(resp *azcore.Response) (SecurityDomainOperationStatusResponse, error) {
	var val *SecurityDomainOperationStatus
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return SecurityDomainOperationStatusResponse{}, err
	}
	return SecurityDomainOperationStatusResponse{RawResponse: resp.Response, SecurityDomainOperationStatus: val}, nil
}

// downloadPendingHandleError handles the DownloadPending error response.
func (client *HSMSecurityDomainClient) downloadPendingHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := KeyVaultError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// TransferKey - Retrieve Security Domain transfer key
// If the operation fails it returns the *KeyVaultError error type.
func (client *HSMSecurityDomainClient) TransferKey(ctx context.Context, vaultBaseURL string, options *HSMSecurityDomainTransferKeyOptions) (TransferKeyResponse, error) {
	req, err := client.transferKeyCreateRequest(ctx, vaultBaseURL, options)
	if err != nil {
		return TransferKeyResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return TransferKeyResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return TransferKeyResponse{}, client.transferKeyHandleError(resp)
	}
	return client.transferKeyHandleResponse(resp)
}

// transferKeyCreateRequest creates the TransferKey request.
func (client *HSMSecurityDomainClient) transferKeyCreateRequest(ctx context.Context, vaultBaseURL string, options *HSMSecurityDomainTransferKeyOptions) (*azcore.Request, error) {
	host := "{vaultBaseUrl}"
	host = strings.ReplaceAll(host, "{vaultBaseUrl}", vaultBaseURL)
	urlPath := "/securitydomain/upload"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "7.2")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// transferKeyHandleResponse handles the TransferKey response.
func (client *HSMSecurityDomainClient) transferKeyHandleResponse(resp *azcore.Response) (TransferKeyResponse, error) {
	var val *TransferKey
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return TransferKeyResponse{}, err
	}
	return TransferKeyResponse{RawResponse: resp.Response, TransferKey: val}, nil
}

// transferKeyHandleError handles the TransferKey error response.
func (client *HSMSecurityDomainClient) transferKeyHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := KeyVaultError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// BeginUpload - Restore the provided Security Domain.
// If the operation fails it returns the *KeyVaultError error type.
func (client *HSMSecurityDomainClient) BeginUpload(ctx context.Context, vaultBaseURL string, securityDomain SecurityDomainObject, options *HSMSecurityDomainBeginUploadOptions) (SecurityDomainOperationStatusPollerResponse, error) {
	resp, err := client.upload(ctx, vaultBaseURL, securityDomain, options)
	if err != nil {
		return SecurityDomainOperationStatusPollerResponse{}, err
	}
	result := SecurityDomainOperationStatusPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := azcore.NewLROPoller("HSMSecurityDomainClient.Upload", resp, client.con.Pipeline(), client.uploadHandleError)
	if err != nil {
		return SecurityDomainOperationStatusPollerResponse{}, err
	}
	poller := &securityDomainOperationStatusPoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (SecurityDomainOperationStatusResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeUpload creates a new SecurityDomainOperationStatusPoller from the specified resume token.
// token - The value must come from a previous call to SecurityDomainOperationStatusPoller.ResumeToken().
func (client *HSMSecurityDomainClient) ResumeUpload(ctx context.Context, token string) (SecurityDomainOperationStatusPollerResponse, error) {
	pt, err := azcore.NewLROPollerFromResumeToken("HSMSecurityDomainClient.Upload", token, client.con.Pipeline(), client.uploadHandleError)
	if err != nil {
		return SecurityDomainOperationStatusPollerResponse{}, err
	}
	poller := &securityDomainOperationStatusPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return SecurityDomainOperationStatusPollerResponse{}, err
	}
	result := SecurityDomainOperationStatusPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (SecurityDomainOperationStatusResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Upload - Restore the provided Security Domain.
// If the operation fails it returns the *KeyVaultError error type.
func (client *HSMSecurityDomainClient) upload(ctx context.Context, vaultBaseURL string, securityDomain SecurityDomainObject, options *HSMSecurityDomainBeginUploadOptions) (*azcore.Response, error) {
	req, err := client.uploadCreateRequest(ctx, vaultBaseURL, securityDomain, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusAccepted, http.StatusNoContent) {
		return nil, client.uploadHandleError(resp)
	}
	return resp, nil
}

// uploadCreateRequest creates the Upload request.
func (client *HSMSecurityDomainClient) uploadCreateRequest(ctx context.Context, vaultBaseURL string, securityDomain SecurityDomainObject, options *HSMSecurityDomainBeginUploadOptions) (*azcore.Request, error) {
	host := "{vaultBaseUrl}"
	host = strings.ReplaceAll(host, "{vaultBaseUrl}", vaultBaseURL)
	urlPath := "/securitydomain/upload"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(securityDomain)
}

// uploadHandleError handles the Upload error response.
func (client *HSMSecurityDomainClient) uploadHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := KeyVaultError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// UploadPending - Get Security Domain upload operation status
// If the operation fails it returns the *KeyVaultError error type.
func (client *HSMSecurityDomainClient) UploadPending(ctx context.Context, vaultBaseURL string, options *HSMSecurityDomainUploadPendingOptions) (SecurityDomainOperationStatusResponse, error) {
	req, err := client.uploadPendingCreateRequest(ctx, vaultBaseURL, options)
	if err != nil {
		return SecurityDomainOperationStatusResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return SecurityDomainOperationStatusResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return SecurityDomainOperationStatusResponse{}, client.uploadPendingHandleError(resp)
	}
	return client.uploadPendingHandleResponse(resp)
}

// uploadPendingCreateRequest creates the UploadPending request.
func (client *HSMSecurityDomainClient) uploadPendingCreateRequest(ctx context.Context, vaultBaseURL string, options *HSMSecurityDomainUploadPendingOptions) (*azcore.Request, error) {
	host := "{vaultBaseUrl}"
	host = strings.ReplaceAll(host, "{vaultBaseUrl}", vaultBaseURL)
	urlPath := "/securitydomain/upload/pending"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// uploadPendingHandleResponse handles the UploadPending response.
func (client *HSMSecurityDomainClient) uploadPendingHandleResponse(resp *azcore.Response) (SecurityDomainOperationStatusResponse, error) {
	var val *SecurityDomainOperationStatus
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return SecurityDomainOperationStatusResponse{}, err
	}
	return SecurityDomainOperationStatusResponse{RawResponse: resp.Response, SecurityDomainOperationStatus: val}, nil
}

// uploadPendingHandleError handles the UploadPending error response.
func (client *HSMSecurityDomainClient) uploadPendingHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := KeyVaultError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}
