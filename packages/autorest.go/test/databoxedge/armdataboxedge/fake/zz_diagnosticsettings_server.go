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
	"net/http"
	"regexp"
)

// DiagnosticSettingsServer is a fake server for instances of the armdataboxedge.DiagnosticSettingsClient type.
type DiagnosticSettingsServer struct {
	// GetDiagnosticProactiveLogCollectionSettings is the fake for method DiagnosticSettingsClient.GetDiagnosticProactiveLogCollectionSettings
	// HTTP status codes to indicate success: http.StatusOK
	GetDiagnosticProactiveLogCollectionSettings func(ctx context.Context, deviceName string, resourceGroupName string, options *armdataboxedge.DiagnosticSettingsClientGetDiagnosticProactiveLogCollectionSettingsOptions) (resp azfake.Responder[armdataboxedge.DiagnosticSettingsClientGetDiagnosticProactiveLogCollectionSettingsResponse], errResp azfake.ErrorResponder)

	// GetDiagnosticRemoteSupportSettings is the fake for method DiagnosticSettingsClient.GetDiagnosticRemoteSupportSettings
	// HTTP status codes to indicate success: http.StatusOK
	GetDiagnosticRemoteSupportSettings func(ctx context.Context, deviceName string, resourceGroupName string, options *armdataboxedge.DiagnosticSettingsClientGetDiagnosticRemoteSupportSettingsOptions) (resp azfake.Responder[armdataboxedge.DiagnosticSettingsClientGetDiagnosticRemoteSupportSettingsResponse], errResp azfake.ErrorResponder)

	// BeginUpdateDiagnosticProactiveLogCollectionSettings is the fake for method DiagnosticSettingsClient.BeginUpdateDiagnosticProactiveLogCollectionSettings
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdateDiagnosticProactiveLogCollectionSettings func(ctx context.Context, deviceName string, resourceGroupName string, proactiveLogCollectionSettings armdataboxedge.DiagnosticProactiveLogCollectionSettings, options *armdataboxedge.DiagnosticSettingsClientBeginUpdateDiagnosticProactiveLogCollectionSettingsOptions) (resp azfake.PollerResponder[armdataboxedge.DiagnosticSettingsClientUpdateDiagnosticProactiveLogCollectionSettingsResponse], errResp azfake.ErrorResponder)

	// BeginUpdateDiagnosticRemoteSupportSettings is the fake for method DiagnosticSettingsClient.BeginUpdateDiagnosticRemoteSupportSettings
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdateDiagnosticRemoteSupportSettings func(ctx context.Context, deviceName string, resourceGroupName string, diagnosticRemoteSupportSettings armdataboxedge.DiagnosticRemoteSupportSettings, options *armdataboxedge.DiagnosticSettingsClientBeginUpdateDiagnosticRemoteSupportSettingsOptions) (resp azfake.PollerResponder[armdataboxedge.DiagnosticSettingsClientUpdateDiagnosticRemoteSupportSettingsResponse], errResp azfake.ErrorResponder)
}

// NewDiagnosticSettingsServerTransport creates a new instance of DiagnosticSettingsServerTransport with the provided implementation.
// The returned DiagnosticSettingsServerTransport instance is connected to an instance of armdataboxedge.DiagnosticSettingsClient by way of the
// undefined.Transporter field.
func NewDiagnosticSettingsServerTransport(srv *DiagnosticSettingsServer) *DiagnosticSettingsServerTransport {
	return &DiagnosticSettingsServerTransport{srv: srv}
}

// DiagnosticSettingsServerTransport connects instances of armdataboxedge.DiagnosticSettingsClient to instances of DiagnosticSettingsServer.
// Don't use this type directly, use NewDiagnosticSettingsServerTransport instead.
type DiagnosticSettingsServerTransport struct {
	srv                                                 *DiagnosticSettingsServer
	beginUpdateDiagnosticProactiveLogCollectionSettings *azfake.PollerResponder[armdataboxedge.DiagnosticSettingsClientUpdateDiagnosticProactiveLogCollectionSettingsResponse]
	beginUpdateDiagnosticRemoteSupportSettings          *azfake.PollerResponder[armdataboxedge.DiagnosticSettingsClientUpdateDiagnosticRemoteSupportSettingsResponse]
}

// Do implements the policy.Transporter interface for DiagnosticSettingsServerTransport.
func (d *DiagnosticSettingsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "DiagnosticSettingsClient.GetDiagnosticProactiveLogCollectionSettings":
		resp, err = d.dispatchGetDiagnosticProactiveLogCollectionSettings(req)
	case "DiagnosticSettingsClient.GetDiagnosticRemoteSupportSettings":
		resp, err = d.dispatchGetDiagnosticRemoteSupportSettings(req)
	case "DiagnosticSettingsClient.BeginUpdateDiagnosticProactiveLogCollectionSettings":
		resp, err = d.dispatchBeginUpdateDiagnosticProactiveLogCollectionSettings(req)
	case "DiagnosticSettingsClient.BeginUpdateDiagnosticRemoteSupportSettings":
		resp, err = d.dispatchBeginUpdateDiagnosticRemoteSupportSettings(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (d *DiagnosticSettingsServerTransport) dispatchGetDiagnosticProactiveLogCollectionSettings(req *http.Request) (*http.Response, error) {
	if d.srv.GetDiagnosticProactiveLogCollectionSettings == nil {
		return nil, &nonRetriableError{errors.New("method GetDiagnosticProactiveLogCollectionSettings not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/(?P<deviceName>[a-zA-Z0-9-_]+)/diagnosticProactiveLogCollectionSettings/default"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := d.srv.GetDiagnosticProactiveLogCollectionSettings(req.Context(), matches[regex.SubexpIndex("deviceName")], matches[regex.SubexpIndex("resourceGroupName")], nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DiagnosticProactiveLogCollectionSettings, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DiagnosticSettingsServerTransport) dispatchGetDiagnosticRemoteSupportSettings(req *http.Request) (*http.Response, error) {
	if d.srv.GetDiagnosticRemoteSupportSettings == nil {
		return nil, &nonRetriableError{errors.New("method GetDiagnosticRemoteSupportSettings not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/(?P<deviceName>[a-zA-Z0-9-_]+)/diagnosticRemoteSupportSettings/default"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := d.srv.GetDiagnosticRemoteSupportSettings(req.Context(), matches[regex.SubexpIndex("deviceName")], matches[regex.SubexpIndex("resourceGroupName")], nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DiagnosticRemoteSupportSettings, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DiagnosticSettingsServerTransport) dispatchBeginUpdateDiagnosticProactiveLogCollectionSettings(req *http.Request) (*http.Response, error) {
	if d.srv.BeginUpdateDiagnosticProactiveLogCollectionSettings == nil {
		return nil, &nonRetriableError{errors.New("method BeginUpdateDiagnosticProactiveLogCollectionSettings not implemented")}
	}
	if d.beginUpdateDiagnosticProactiveLogCollectionSettings == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/(?P<deviceName>[a-zA-Z0-9-_]+)/diagnosticProactiveLogCollectionSettings/default"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armdataboxedge.DiagnosticProactiveLogCollectionSettings](req)
		if err != nil {
			return nil, err
		}
		respr, errRespr := d.srv.BeginUpdateDiagnosticProactiveLogCollectionSettings(req.Context(), matches[regex.SubexpIndex("deviceName")], matches[regex.SubexpIndex("resourceGroupName")], body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		d.beginUpdateDiagnosticProactiveLogCollectionSettings = &respr
	}

	resp, err := server.PollerResponderNext(d.beginUpdateDiagnosticProactiveLogCollectionSettings, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(d.beginUpdateDiagnosticProactiveLogCollectionSettings) {
		d.beginUpdateDiagnosticProactiveLogCollectionSettings = nil
	}

	return resp, nil
}

func (d *DiagnosticSettingsServerTransport) dispatchBeginUpdateDiagnosticRemoteSupportSettings(req *http.Request) (*http.Response, error) {
	if d.srv.BeginUpdateDiagnosticRemoteSupportSettings == nil {
		return nil, &nonRetriableError{errors.New("method BeginUpdateDiagnosticRemoteSupportSettings not implemented")}
	}
	if d.beginUpdateDiagnosticRemoteSupportSettings == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/(?P<deviceName>[a-zA-Z0-9-_]+)/diagnosticRemoteSupportSettings/default"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armdataboxedge.DiagnosticRemoteSupportSettings](req)
		if err != nil {
			return nil, err
		}
		respr, errRespr := d.srv.BeginUpdateDiagnosticRemoteSupportSettings(req.Context(), matches[regex.SubexpIndex("deviceName")], matches[regex.SubexpIndex("resourceGroupName")], body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		d.beginUpdateDiagnosticRemoteSupportSettings = &respr
	}

	resp, err := server.PollerResponderNext(d.beginUpdateDiagnosticRemoteSupportSettings, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(d.beginUpdateDiagnosticRemoteSupportSettings) {
		d.beginUpdateDiagnosticRemoteSupportSettings = nil
	}

	return resp, nil
}