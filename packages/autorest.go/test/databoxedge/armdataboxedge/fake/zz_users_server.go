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

// UsersServer is a fake server for instances of the armdataboxedge.UsersClient type.
type UsersServer struct {
	// BeginCreateOrUpdate is the fake for method UsersClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginCreateOrUpdate func(ctx context.Context, deviceName string, name string, resourceGroupName string, userParam armdataboxedge.User, options *armdataboxedge.UsersClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armdataboxedge.UsersClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method UsersClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, deviceName string, name string, resourceGroupName string, options *armdataboxedge.UsersClientBeginDeleteOptions) (resp azfake.PollerResponder[armdataboxedge.UsersClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method UsersClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, deviceName string, name string, resourceGroupName string, options *armdataboxedge.UsersClientGetOptions) (resp azfake.Responder[armdataboxedge.UsersClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByDataBoxEdgeDevicePager is the fake for method UsersClient.NewListByDataBoxEdgeDevicePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByDataBoxEdgeDevicePager func(deviceName string, resourceGroupName string, options *armdataboxedge.UsersClientListByDataBoxEdgeDeviceOptions) (resp azfake.PagerResponder[armdataboxedge.UsersClientListByDataBoxEdgeDeviceResponse])
}

// NewUsersServerTransport creates a new instance of UsersServerTransport with the provided implementation.
// The returned UsersServerTransport instance is connected to an instance of armdataboxedge.UsersClient by way of the
// undefined.Transporter field.
func NewUsersServerTransport(srv *UsersServer) *UsersServerTransport {
	return &UsersServerTransport{srv: srv}
}

// UsersServerTransport connects instances of armdataboxedge.UsersClient to instances of UsersServer.
// Don't use this type directly, use NewUsersServerTransport instead.
type UsersServerTransport struct {
	srv                             *UsersServer
	beginCreateOrUpdate             *azfake.PollerResponder[armdataboxedge.UsersClientCreateOrUpdateResponse]
	beginDelete                     *azfake.PollerResponder[armdataboxedge.UsersClientDeleteResponse]
	newListByDataBoxEdgeDevicePager *azfake.PagerResponder[armdataboxedge.UsersClientListByDataBoxEdgeDeviceResponse]
}

// Do implements the policy.Transporter interface for UsersServerTransport.
func (u *UsersServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "UsersClient.BeginCreateOrUpdate":
		resp, err = u.dispatchBeginCreateOrUpdate(req)
	case "UsersClient.BeginDelete":
		resp, err = u.dispatchBeginDelete(req)
	case "UsersClient.Get":
		resp, err = u.dispatchGet(req)
	case "UsersClient.NewListByDataBoxEdgeDevicePager":
		resp, err = u.dispatchNewListByDataBoxEdgeDevicePager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (u *UsersServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if u.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("method BeginCreateOrUpdate not implemented")}
	}
	if u.beginCreateOrUpdate == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/(?P<deviceName>[a-zA-Z0-9-_]+)/users/(?P<name>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armdataboxedge.User](req)
		if err != nil {
			return nil, err
		}
		respr, errRespr := u.srv.BeginCreateOrUpdate(req.Context(), matches[regex.SubexpIndex("deviceName")], matches[regex.SubexpIndex("name")], matches[regex.SubexpIndex("resourceGroupName")], body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		u.beginCreateOrUpdate = &respr
	}

	resp, err := server.PollerResponderNext(u.beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(u.beginCreateOrUpdate) {
		u.beginCreateOrUpdate = nil
	}

	return resp, nil
}

func (u *UsersServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if u.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("method BeginDelete not implemented")}
	}
	if u.beginDelete == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/(?P<deviceName>[a-zA-Z0-9-_]+)/users/(?P<name>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		respr, errRespr := u.srv.BeginDelete(req.Context(), matches[regex.SubexpIndex("deviceName")], matches[regex.SubexpIndex("name")], matches[regex.SubexpIndex("resourceGroupName")], nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		u.beginDelete = &respr
	}

	resp, err := server.PollerResponderNext(u.beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(u.beginDelete) {
		u.beginDelete = nil
	}

	return resp, nil
}

func (u *UsersServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if u.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("method Get not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/(?P<deviceName>[a-zA-Z0-9-_]+)/users/(?P<name>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := u.srv.Get(req.Context(), matches[regex.SubexpIndex("deviceName")], matches[regex.SubexpIndex("name")], matches[regex.SubexpIndex("resourceGroupName")], nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).User, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (u *UsersServerTransport) dispatchNewListByDataBoxEdgeDevicePager(req *http.Request) (*http.Response, error) {
	if u.srv.NewListByDataBoxEdgeDevicePager == nil {
		return nil, &nonRetriableError{errors.New("method NewListByDataBoxEdgeDevicePager not implemented")}
	}
	if u.newListByDataBoxEdgeDevicePager == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/(?P<deviceName>[a-zA-Z0-9-_]+)/users"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		filterParam := getOptional(qp.Get("$filter"))
		var options *armdataboxedge.UsersClientListByDataBoxEdgeDeviceOptions
		if filterParam != nil {
			options = &armdataboxedge.UsersClientListByDataBoxEdgeDeviceOptions{
				Filter: filterParam,
			}
		}
		resp := u.srv.NewListByDataBoxEdgeDevicePager(matches[regex.SubexpIndex("deviceName")], matches[regex.SubexpIndex("resourceGroupName")], options)
		u.newListByDataBoxEdgeDevicePager = &resp
		server.PagerResponderInjectNextLinks(u.newListByDataBoxEdgeDevicePager, req, func(page *armdataboxedge.UsersClientListByDataBoxEdgeDeviceResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(u.newListByDataBoxEdgeDevicePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(u.newListByDataBoxEdgeDevicePager) {
		u.newListByDataBoxEdgeDevicePager = nil
	}
	return resp, nil
}
