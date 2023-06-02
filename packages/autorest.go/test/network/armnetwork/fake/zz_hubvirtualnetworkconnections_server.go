//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package fake

import (
	"armnetwork"
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

// HubVirtualNetworkConnectionsServer is a fake server for instances of the armnetwork.HubVirtualNetworkConnectionsClient type.
type HubVirtualNetworkConnectionsServer struct {
	// BeginCreateOrUpdate is the fake for method HubVirtualNetworkConnectionsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, virtualHubName string, connectionName string, hubVirtualNetworkConnectionParameters armnetwork.HubVirtualNetworkConnection, options *armnetwork.HubVirtualNetworkConnectionsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armnetwork.HubVirtualNetworkConnectionsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method HubVirtualNetworkConnectionsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, virtualHubName string, connectionName string, options *armnetwork.HubVirtualNetworkConnectionsClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetwork.HubVirtualNetworkConnectionsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method HubVirtualNetworkConnectionsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, virtualHubName string, connectionName string, options *armnetwork.HubVirtualNetworkConnectionsClientGetOptions) (resp azfake.Responder[armnetwork.HubVirtualNetworkConnectionsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method HubVirtualNetworkConnectionsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, virtualHubName string, options *armnetwork.HubVirtualNetworkConnectionsClientListOptions) (resp azfake.PagerResponder[armnetwork.HubVirtualNetworkConnectionsClientListResponse])
}

// NewHubVirtualNetworkConnectionsServerTransport creates a new instance of HubVirtualNetworkConnectionsServerTransport with the provided implementation.
// The returned HubVirtualNetworkConnectionsServerTransport instance is connected to an instance of armnetwork.HubVirtualNetworkConnectionsClient by way of the
// undefined.Transporter field.
func NewHubVirtualNetworkConnectionsServerTransport(srv *HubVirtualNetworkConnectionsServer) *HubVirtualNetworkConnectionsServerTransport {
	return &HubVirtualNetworkConnectionsServerTransport{srv: srv}
}

// HubVirtualNetworkConnectionsServerTransport connects instances of armnetwork.HubVirtualNetworkConnectionsClient to instances of HubVirtualNetworkConnectionsServer.
// Don't use this type directly, use NewHubVirtualNetworkConnectionsServerTransport instead.
type HubVirtualNetworkConnectionsServerTransport struct {
	srv                 *HubVirtualNetworkConnectionsServer
	beginCreateOrUpdate *azfake.PollerResponder[armnetwork.HubVirtualNetworkConnectionsClientCreateOrUpdateResponse]
	beginDelete         *azfake.PollerResponder[armnetwork.HubVirtualNetworkConnectionsClientDeleteResponse]
	newListPager        *azfake.PagerResponder[armnetwork.HubVirtualNetworkConnectionsClientListResponse]
}

// Do implements the policy.Transporter interface for HubVirtualNetworkConnectionsServerTransport.
func (h *HubVirtualNetworkConnectionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "HubVirtualNetworkConnectionsClient.BeginCreateOrUpdate":
		resp, err = h.dispatchBeginCreateOrUpdate(req)
	case "HubVirtualNetworkConnectionsClient.BeginDelete":
		resp, err = h.dispatchBeginDelete(req)
	case "HubVirtualNetworkConnectionsClient.Get":
		resp, err = h.dispatchGet(req)
	case "HubVirtualNetworkConnectionsClient.NewListPager":
		resp, err = h.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h *HubVirtualNetworkConnectionsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if h.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("method BeginCreateOrUpdate not implemented")}
	}
	if h.beginCreateOrUpdate == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/virtualHubs/(?P<virtualHubName>[a-zA-Z0-9-_]+)/hubVirtualNetworkConnections/(?P<connectionName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.HubVirtualNetworkConnection](req)
		if err != nil {
			return nil, err
		}
		respr, errRespr := h.srv.BeginCreateOrUpdate(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("virtualHubName")], matches[regex.SubexpIndex("connectionName")], body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		h.beginCreateOrUpdate = &respr
	}

	resp, err := server.PollerResponderNext(h.beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(h.beginCreateOrUpdate) {
		h.beginCreateOrUpdate = nil
	}

	return resp, nil
}

func (h *HubVirtualNetworkConnectionsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if h.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("method BeginDelete not implemented")}
	}
	if h.beginDelete == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/virtualHubs/(?P<virtualHubName>[a-zA-Z0-9-_]+)/hubVirtualNetworkConnections/(?P<connectionName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		respr, errRespr := h.srv.BeginDelete(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("virtualHubName")], matches[regex.SubexpIndex("connectionName")], nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		h.beginDelete = &respr
	}

	resp, err := server.PollerResponderNext(h.beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(h.beginDelete) {
		h.beginDelete = nil
	}

	return resp, nil
}

func (h *HubVirtualNetworkConnectionsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if h.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("method Get not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/virtualHubs/(?P<virtualHubName>[a-zA-Z0-9-_]+)/hubVirtualNetworkConnections/(?P<connectionName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := h.srv.Get(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("virtualHubName")], matches[regex.SubexpIndex("connectionName")], nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).HubVirtualNetworkConnection, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *HubVirtualNetworkConnectionsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if h.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("method NewListPager not implemented")}
	}
	if h.newListPager == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/virtualHubs/(?P<virtualHubName>[a-zA-Z0-9-_]+)/hubVirtualNetworkConnections"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := h.srv.NewListPager(matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("virtualHubName")], nil)
		h.newListPager = &resp
		server.PagerResponderInjectNextLinks(h.newListPager, req, func(page *armnetwork.HubVirtualNetworkConnectionsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(h.newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(h.newListPager) {
		h.newListPager = nil
	}
	return resp, nil
}