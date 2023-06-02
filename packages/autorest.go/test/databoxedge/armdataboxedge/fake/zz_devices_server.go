//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package fake

import (
	"armdataboxedge"
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"net/http"
	"regexp"
)

// DevicesServer is a fake server for instances of the armdataboxedge.DevicesClient type.
type DevicesServer struct {
	// CreateOrUpdate is the fake for method DevicesClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK
	CreateOrUpdate func(ctx context.Context, deviceName string, resourceGroupName string, dataBoxEdgeDevice armdataboxedge.Device, options *armdataboxedge.DevicesClientCreateOrUpdateOptions) (resp azfake.Responder[armdataboxedge.DevicesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginCreateOrUpdateSecuritySettings is the fake for method DevicesClient.BeginCreateOrUpdateSecuritySettings
	// HTTP status codes to indicate success: http.StatusAccepted, http.StatusNoContent
	BeginCreateOrUpdateSecuritySettings func(ctx context.Context, deviceName string, resourceGroupName string, securitySettings armdataboxedge.SecuritySettings, options *armdataboxedge.DevicesClientBeginCreateOrUpdateSecuritySettingsOptions) (resp azfake.PollerResponder[armdataboxedge.DevicesClientCreateOrUpdateSecuritySettingsResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method DevicesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, deviceName string, resourceGroupName string, options *armdataboxedge.DevicesClientBeginDeleteOptions) (resp azfake.PollerResponder[armdataboxedge.DevicesClientDeleteResponse], errResp azfake.ErrorResponder)

	// BeginDownloadUpdates is the fake for method DevicesClient.BeginDownloadUpdates
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginDownloadUpdates func(ctx context.Context, deviceName string, resourceGroupName string, options *armdataboxedge.DevicesClientBeginDownloadUpdatesOptions) (resp azfake.PollerResponder[armdataboxedge.DevicesClientDownloadUpdatesResponse], errResp azfake.ErrorResponder)

	// GenerateCertificate is the fake for method DevicesClient.GenerateCertificate
	// HTTP status codes to indicate success: http.StatusOK
	GenerateCertificate func(ctx context.Context, deviceName string, resourceGroupName string, options *armdataboxedge.DevicesClientGenerateCertificateOptions) (resp azfake.Responder[armdataboxedge.DevicesClientGenerateCertificateResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method DevicesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, deviceName string, resourceGroupName string, options *armdataboxedge.DevicesClientGetOptions) (resp azfake.Responder[armdataboxedge.DevicesClientGetResponse], errResp azfake.ErrorResponder)

	// GetExtendedInformation is the fake for method DevicesClient.GetExtendedInformation
	// HTTP status codes to indicate success: http.StatusOK
	GetExtendedInformation func(ctx context.Context, deviceName string, resourceGroupName string, options *armdataboxedge.DevicesClientGetExtendedInformationOptions) (resp azfake.Responder[armdataboxedge.DevicesClientGetExtendedInformationResponse], errResp azfake.ErrorResponder)

	// GetNetworkSettings is the fake for method DevicesClient.GetNetworkSettings
	// HTTP status codes to indicate success: http.StatusOK
	GetNetworkSettings func(ctx context.Context, deviceName string, resourceGroupName string, options *armdataboxedge.DevicesClientGetNetworkSettingsOptions) (resp azfake.Responder[armdataboxedge.DevicesClientGetNetworkSettingsResponse], errResp azfake.ErrorResponder)

	// GetUpdateSummary is the fake for method DevicesClient.GetUpdateSummary
	// HTTP status codes to indicate success: http.StatusOK
	GetUpdateSummary func(ctx context.Context, deviceName string, resourceGroupName string, options *armdataboxedge.DevicesClientGetUpdateSummaryOptions) (resp azfake.Responder[armdataboxedge.DevicesClientGetUpdateSummaryResponse], errResp azfake.ErrorResponder)

	// BeginInstallUpdates is the fake for method DevicesClient.BeginInstallUpdates
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginInstallUpdates func(ctx context.Context, deviceName string, resourceGroupName string, options *armdataboxedge.DevicesClientBeginInstallUpdatesOptions) (resp azfake.PollerResponder[armdataboxedge.DevicesClientInstallUpdatesResponse], errResp azfake.ErrorResponder)

	// NewListByResourceGroupPager is the fake for method DevicesClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armdataboxedge.DevicesClientListByResourceGroupOptions) (resp azfake.PagerResponder[armdataboxedge.DevicesClientListByResourceGroupResponse])

	// NewListBySubscriptionPager is the fake for method DevicesClient.NewListBySubscriptionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySubscriptionPager func(options *armdataboxedge.DevicesClientListBySubscriptionOptions) (resp azfake.PagerResponder[armdataboxedge.DevicesClientListBySubscriptionResponse])

	// BeginScanForUpdates is the fake for method DevicesClient.BeginScanForUpdates
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginScanForUpdates func(ctx context.Context, deviceName string, resourceGroupName string, options *armdataboxedge.DevicesClientBeginScanForUpdatesOptions) (resp azfake.PollerResponder[armdataboxedge.DevicesClientScanForUpdatesResponse], errResp azfake.ErrorResponder)

	// Update is the fake for method DevicesClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, deviceName string, resourceGroupName string, parameters armdataboxedge.DevicePatch, options *armdataboxedge.DevicesClientUpdateOptions) (resp azfake.Responder[armdataboxedge.DevicesClientUpdateResponse], errResp azfake.ErrorResponder)

	// UpdateExtendedInformation is the fake for method DevicesClient.UpdateExtendedInformation
	// HTTP status codes to indicate success: http.StatusOK
	UpdateExtendedInformation func(ctx context.Context, deviceName string, resourceGroupName string, parameters armdataboxedge.DeviceExtendedInfoPatch, options *armdataboxedge.DevicesClientUpdateExtendedInformationOptions) (resp azfake.Responder[armdataboxedge.DevicesClientUpdateExtendedInformationResponse], errResp azfake.ErrorResponder)

	// UploadCertificate is the fake for method DevicesClient.UploadCertificate
	// HTTP status codes to indicate success: http.StatusOK
	UploadCertificate func(ctx context.Context, deviceName string, resourceGroupName string, parameters armdataboxedge.UploadCertificateRequest, options *armdataboxedge.DevicesClientUploadCertificateOptions) (resp azfake.Responder[armdataboxedge.DevicesClientUploadCertificateResponse], errResp azfake.ErrorResponder)
}

// NewDevicesServerTransport creates a new instance of DevicesServerTransport with the provided implementation.
// The returned DevicesServerTransport instance is connected to an instance of armdataboxedge.DevicesClient by way of the
// undefined.Transporter field.
func NewDevicesServerTransport(srv *DevicesServer) *DevicesServerTransport {
	return &DevicesServerTransport{srv: srv}
}

// DevicesServerTransport connects instances of armdataboxedge.DevicesClient to instances of DevicesServer.
// Don't use this type directly, use NewDevicesServerTransport instead.
type DevicesServerTransport struct {
	srv                                 *DevicesServer
	beginCreateOrUpdateSecuritySettings *azfake.PollerResponder[armdataboxedge.DevicesClientCreateOrUpdateSecuritySettingsResponse]
	beginDelete                         *azfake.PollerResponder[armdataboxedge.DevicesClientDeleteResponse]
	beginDownloadUpdates                *azfake.PollerResponder[armdataboxedge.DevicesClientDownloadUpdatesResponse]
	beginInstallUpdates                 *azfake.PollerResponder[armdataboxedge.DevicesClientInstallUpdatesResponse]
	newListByResourceGroupPager         *azfake.PagerResponder[armdataboxedge.DevicesClientListByResourceGroupResponse]
	newListBySubscriptionPager          *azfake.PagerResponder[armdataboxedge.DevicesClientListBySubscriptionResponse]
	beginScanForUpdates                 *azfake.PollerResponder[armdataboxedge.DevicesClientScanForUpdatesResponse]
}

// Do implements the policy.Transporter interface for DevicesServerTransport.
func (d *DevicesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "DevicesClient.CreateOrUpdate":
		resp, err = d.dispatchCreateOrUpdate(req)
	case "DevicesClient.BeginCreateOrUpdateSecuritySettings":
		resp, err = d.dispatchBeginCreateOrUpdateSecuritySettings(req)
	case "DevicesClient.BeginDelete":
		resp, err = d.dispatchBeginDelete(req)
	case "DevicesClient.BeginDownloadUpdates":
		resp, err = d.dispatchBeginDownloadUpdates(req)
	case "DevicesClient.GenerateCertificate":
		resp, err = d.dispatchGenerateCertificate(req)
	case "DevicesClient.Get":
		resp, err = d.dispatchGet(req)
	case "DevicesClient.GetExtendedInformation":
		resp, err = d.dispatchGetExtendedInformation(req)
	case "DevicesClient.GetNetworkSettings":
		resp, err = d.dispatchGetNetworkSettings(req)
	case "DevicesClient.GetUpdateSummary":
		resp, err = d.dispatchGetUpdateSummary(req)
	case "DevicesClient.BeginInstallUpdates":
		resp, err = d.dispatchBeginInstallUpdates(req)
	case "DevicesClient.NewListByResourceGroupPager":
		resp, err = d.dispatchNewListByResourceGroupPager(req)
	case "DevicesClient.NewListBySubscriptionPager":
		resp, err = d.dispatchNewListBySubscriptionPager(req)
	case "DevicesClient.BeginScanForUpdates":
		resp, err = d.dispatchBeginScanForUpdates(req)
	case "DevicesClient.Update":
		resp, err = d.dispatchUpdate(req)
	case "DevicesClient.UpdateExtendedInformation":
		resp, err = d.dispatchUpdateExtendedInformation(req)
	case "DevicesClient.UploadCertificate":
		resp, err = d.dispatchUploadCertificate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (d *DevicesServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if d.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("method CreateOrUpdate not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/(?P<deviceName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armdataboxedge.Device](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.CreateOrUpdate(req.Context(), matches[regex.SubexpIndex("deviceName")], matches[regex.SubexpIndex("resourceGroupName")], body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Device, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DevicesServerTransport) dispatchBeginCreateOrUpdateSecuritySettings(req *http.Request) (*http.Response, error) {
	if d.srv.BeginCreateOrUpdateSecuritySettings == nil {
		return nil, &nonRetriableError{errors.New("method BeginCreateOrUpdateSecuritySettings not implemented")}
	}
	if d.beginCreateOrUpdateSecuritySettings == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/(?P<deviceName>[a-zA-Z0-9-_]+)/securitySettings/default/update"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armdataboxedge.SecuritySettings](req)
		if err != nil {
			return nil, err
		}
		respr, errRespr := d.srv.BeginCreateOrUpdateSecuritySettings(req.Context(), matches[regex.SubexpIndex("deviceName")], matches[regex.SubexpIndex("resourceGroupName")], body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		d.beginCreateOrUpdateSecuritySettings = &respr
	}

	resp, err := server.PollerResponderNext(d.beginCreateOrUpdateSecuritySettings, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(d.beginCreateOrUpdateSecuritySettings) {
		d.beginCreateOrUpdateSecuritySettings = nil
	}

	return resp, nil
}

func (d *DevicesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if d.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("method BeginDelete not implemented")}
	}
	if d.beginDelete == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/(?P<deviceName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		respr, errRespr := d.srv.BeginDelete(req.Context(), matches[regex.SubexpIndex("deviceName")], matches[regex.SubexpIndex("resourceGroupName")], nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		d.beginDelete = &respr
	}

	resp, err := server.PollerResponderNext(d.beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(d.beginDelete) {
		d.beginDelete = nil
	}

	return resp, nil
}

func (d *DevicesServerTransport) dispatchBeginDownloadUpdates(req *http.Request) (*http.Response, error) {
	if d.srv.BeginDownloadUpdates == nil {
		return nil, &nonRetriableError{errors.New("method BeginDownloadUpdates not implemented")}
	}
	if d.beginDownloadUpdates == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/(?P<deviceName>[a-zA-Z0-9-_]+)/downloadUpdates"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		respr, errRespr := d.srv.BeginDownloadUpdates(req.Context(), matches[regex.SubexpIndex("deviceName")], matches[regex.SubexpIndex("resourceGroupName")], nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		d.beginDownloadUpdates = &respr
	}

	resp, err := server.PollerResponderNext(d.beginDownloadUpdates, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(d.beginDownloadUpdates) {
		d.beginDownloadUpdates = nil
	}

	return resp, nil
}

func (d *DevicesServerTransport) dispatchGenerateCertificate(req *http.Request) (*http.Response, error) {
	if d.srv.GenerateCertificate == nil {
		return nil, &nonRetriableError{errors.New("method GenerateCertificate not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/(?P<deviceName>[a-zA-Z0-9-_]+)/generateCertificate"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := d.srv.GenerateCertificate(req.Context(), matches[regex.SubexpIndex("deviceName")], matches[regex.SubexpIndex("resourceGroupName")], nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).GenerateCertResponse, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DevicesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if d.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("method Get not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/(?P<deviceName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := d.srv.Get(req.Context(), matches[regex.SubexpIndex("deviceName")], matches[regex.SubexpIndex("resourceGroupName")], nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Device, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DevicesServerTransport) dispatchGetExtendedInformation(req *http.Request) (*http.Response, error) {
	if d.srv.GetExtendedInformation == nil {
		return nil, &nonRetriableError{errors.New("method GetExtendedInformation not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/(?P<deviceName>[a-zA-Z0-9-_]+)/getExtendedInformation"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := d.srv.GetExtendedInformation(req.Context(), matches[regex.SubexpIndex("deviceName")], matches[regex.SubexpIndex("resourceGroupName")], nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DeviceExtendedInfo, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DevicesServerTransport) dispatchGetNetworkSettings(req *http.Request) (*http.Response, error) {
	if d.srv.GetNetworkSettings == nil {
		return nil, &nonRetriableError{errors.New("method GetNetworkSettings not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/(?P<deviceName>[a-zA-Z0-9-_]+)/networkSettings/default"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := d.srv.GetNetworkSettings(req.Context(), matches[regex.SubexpIndex("deviceName")], matches[regex.SubexpIndex("resourceGroupName")], nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).NetworkSettings, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DevicesServerTransport) dispatchGetUpdateSummary(req *http.Request) (*http.Response, error) {
	if d.srv.GetUpdateSummary == nil {
		return nil, &nonRetriableError{errors.New("method GetUpdateSummary not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/(?P<deviceName>[a-zA-Z0-9-_]+)/updateSummary/default"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := d.srv.GetUpdateSummary(req.Context(), matches[regex.SubexpIndex("deviceName")], matches[regex.SubexpIndex("resourceGroupName")], nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).UpdateSummary, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DevicesServerTransport) dispatchBeginInstallUpdates(req *http.Request) (*http.Response, error) {
	if d.srv.BeginInstallUpdates == nil {
		return nil, &nonRetriableError{errors.New("method BeginInstallUpdates not implemented")}
	}
	if d.beginInstallUpdates == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/(?P<deviceName>[a-zA-Z0-9-_]+)/installUpdates"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		respr, errRespr := d.srv.BeginInstallUpdates(req.Context(), matches[regex.SubexpIndex("deviceName")], matches[regex.SubexpIndex("resourceGroupName")], nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		d.beginInstallUpdates = &respr
	}

	resp, err := server.PollerResponderNext(d.beginInstallUpdates, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(d.beginInstallUpdates) {
		d.beginInstallUpdates = nil
	}

	return resp, nil
}

func (d *DevicesServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if d.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("method NewListByResourceGroupPager not implemented")}
	}
	if d.newListByResourceGroupPager == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		expandParam := getOptional(qp.Get("$expand"))
		var options *armdataboxedge.DevicesClientListByResourceGroupOptions
		if expandParam != nil {
			options = &armdataboxedge.DevicesClientListByResourceGroupOptions{
				Expand: expandParam,
			}
		}
		resp := d.srv.NewListByResourceGroupPager(matches[regex.SubexpIndex("resourceGroupName")], options)
		d.newListByResourceGroupPager = &resp
		server.PagerResponderInjectNextLinks(d.newListByResourceGroupPager, req, func(page *armdataboxedge.DevicesClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(d.newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(d.newListByResourceGroupPager) {
		d.newListByResourceGroupPager = nil
	}
	return resp, nil
}

func (d *DevicesServerTransport) dispatchNewListBySubscriptionPager(req *http.Request) (*http.Response, error) {
	if d.srv.NewListBySubscriptionPager == nil {
		return nil, &nonRetriableError{errors.New("method NewListBySubscriptionPager not implemented")}
	}
	if d.newListBySubscriptionPager == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		expandParam := getOptional(qp.Get("$expand"))
		var options *armdataboxedge.DevicesClientListBySubscriptionOptions
		if expandParam != nil {
			options = &armdataboxedge.DevicesClientListBySubscriptionOptions{
				Expand: expandParam,
			}
		}
		resp := d.srv.NewListBySubscriptionPager(options)
		d.newListBySubscriptionPager = &resp
		server.PagerResponderInjectNextLinks(d.newListBySubscriptionPager, req, func(page *armdataboxedge.DevicesClientListBySubscriptionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(d.newListBySubscriptionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(d.newListBySubscriptionPager) {
		d.newListBySubscriptionPager = nil
	}
	return resp, nil
}

func (d *DevicesServerTransport) dispatchBeginScanForUpdates(req *http.Request) (*http.Response, error) {
	if d.srv.BeginScanForUpdates == nil {
		return nil, &nonRetriableError{errors.New("method BeginScanForUpdates not implemented")}
	}
	if d.beginScanForUpdates == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/(?P<deviceName>[a-zA-Z0-9-_]+)/scanForUpdates"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		respr, errRespr := d.srv.BeginScanForUpdates(req.Context(), matches[regex.SubexpIndex("deviceName")], matches[regex.SubexpIndex("resourceGroupName")], nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		d.beginScanForUpdates = &respr
	}

	resp, err := server.PollerResponderNext(d.beginScanForUpdates, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(d.beginScanForUpdates) {
		d.beginScanForUpdates = nil
	}

	return resp, nil
}

func (d *DevicesServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if d.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("method Update not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/(?P<deviceName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armdataboxedge.DevicePatch](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.Update(req.Context(), matches[regex.SubexpIndex("deviceName")], matches[regex.SubexpIndex("resourceGroupName")], body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Device, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DevicesServerTransport) dispatchUpdateExtendedInformation(req *http.Request) (*http.Response, error) {
	if d.srv.UpdateExtendedInformation == nil {
		return nil, &nonRetriableError{errors.New("method UpdateExtendedInformation not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/(?P<deviceName>[a-zA-Z0-9-_]+)/updateExtendedInformation"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armdataboxedge.DeviceExtendedInfoPatch](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.UpdateExtendedInformation(req.Context(), matches[regex.SubexpIndex("deviceName")], matches[regex.SubexpIndex("resourceGroupName")], body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DeviceExtendedInfo, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DevicesServerTransport) dispatchUploadCertificate(req *http.Request) (*http.Response, error) {
	if d.srv.UploadCertificate == nil {
		return nil, &nonRetriableError{errors.New("method UploadCertificate not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/(?P<deviceName>[a-zA-Z0-9-_]+)/uploadCertificate"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armdataboxedge.UploadCertificateRequest](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.UploadCertificate(req.Context(), matches[regex.SubexpIndex("deviceName")], matches[regex.SubexpIndex("resourceGroupName")], body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).UploadCertificateResponse, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}