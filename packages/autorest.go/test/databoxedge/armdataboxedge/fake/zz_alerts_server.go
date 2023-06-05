//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

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

// AlertsServer is a fake server for instances of the armdataboxedge.AlertsClient type.
type AlertsServer struct {
	// Get is the fake for method AlertsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, deviceName string, name string, resourceGroupName string, options *armdataboxedge.AlertsClientGetOptions) (resp azfake.Responder[armdataboxedge.AlertsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByDataBoxEdgeDevicePager is the fake for method AlertsClient.NewListByDataBoxEdgeDevicePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByDataBoxEdgeDevicePager func(deviceName string, resourceGroupName string, options *armdataboxedge.AlertsClientListByDataBoxEdgeDeviceOptions) (resp azfake.PagerResponder[armdataboxedge.AlertsClientListByDataBoxEdgeDeviceResponse])
}

// NewAlertsServerTransport creates a new instance of AlertsServerTransport with the provided implementation.
// The returned AlertsServerTransport instance is connected to an instance of armdataboxedge.AlertsClient by way of the
// undefined.Transporter field.
func NewAlertsServerTransport(srv *AlertsServer) *AlertsServerTransport {
	return &AlertsServerTransport{srv: srv}
}

// AlertsServerTransport connects instances of armdataboxedge.AlertsClient to instances of AlertsServer.
// Don't use this type directly, use NewAlertsServerTransport instead.
type AlertsServerTransport struct {
	srv                             *AlertsServer
	newListByDataBoxEdgeDevicePager *azfake.PagerResponder[armdataboxedge.AlertsClientListByDataBoxEdgeDeviceResponse]
}

// Do implements the policy.Transporter interface for AlertsServerTransport.
func (a *AlertsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "AlertsClient.Get":
		resp, err = a.dispatchGet(req)
	case "AlertsClient.NewListByDataBoxEdgeDevicePager":
		resp, err = a.dispatchNewListByDataBoxEdgeDevicePager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *AlertsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if a.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("method Get not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/(?P<deviceName>[a-zA-Z0-9-_]+)/alerts/(?P<name>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := a.srv.Get(req.Context(), matches[regex.SubexpIndex("deviceName")], matches[regex.SubexpIndex("name")], matches[regex.SubexpIndex("resourceGroupName")], nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Alert, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AlertsServerTransport) dispatchNewListByDataBoxEdgeDevicePager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListByDataBoxEdgeDevicePager == nil {
		return nil, &nonRetriableError{errors.New("method NewListByDataBoxEdgeDevicePager not implemented")}
	}
	if a.newListByDataBoxEdgeDevicePager == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/(?P<deviceName>[a-zA-Z0-9-_]+)/alerts"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := a.srv.NewListByDataBoxEdgeDevicePager(matches[regex.SubexpIndex("deviceName")], matches[regex.SubexpIndex("resourceGroupName")], nil)
		a.newListByDataBoxEdgeDevicePager = &resp
		server.PagerResponderInjectNextLinks(a.newListByDataBoxEdgeDevicePager, req, func(page *armdataboxedge.AlertsClientListByDataBoxEdgeDeviceResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(a.newListByDataBoxEdgeDevicePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(a.newListByDataBoxEdgeDevicePager) {
		a.newListByDataBoxEdgeDevicePager = nil
	}
	return resp, nil
}
