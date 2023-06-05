//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

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

// VirtualHubBgpConnectionsServer is a fake server for instances of the armnetwork.VirtualHubBgpConnectionsClient type.
type VirtualHubBgpConnectionsServer struct {
	// NewListPager is the fake for method VirtualHubBgpConnectionsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, virtualHubName string, options *armnetwork.VirtualHubBgpConnectionsClientListOptions) (resp azfake.PagerResponder[armnetwork.VirtualHubBgpConnectionsClientListResponse])

	// BeginListAdvertisedRoutes is the fake for method VirtualHubBgpConnectionsClient.BeginListAdvertisedRoutes
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginListAdvertisedRoutes func(ctx context.Context, resourceGroupName string, hubName string, connectionName string, options *armnetwork.VirtualHubBgpConnectionsClientBeginListAdvertisedRoutesOptions) (resp azfake.PollerResponder[armnetwork.VirtualHubBgpConnectionsClientListAdvertisedRoutesResponse], errResp azfake.ErrorResponder)

	// BeginListLearnedRoutes is the fake for method VirtualHubBgpConnectionsClient.BeginListLearnedRoutes
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginListLearnedRoutes func(ctx context.Context, resourceGroupName string, hubName string, connectionName string, options *armnetwork.VirtualHubBgpConnectionsClientBeginListLearnedRoutesOptions) (resp azfake.PollerResponder[armnetwork.VirtualHubBgpConnectionsClientListLearnedRoutesResponse], errResp azfake.ErrorResponder)
}

// NewVirtualHubBgpConnectionsServerTransport creates a new instance of VirtualHubBgpConnectionsServerTransport with the provided implementation.
// The returned VirtualHubBgpConnectionsServerTransport instance is connected to an instance of armnetwork.VirtualHubBgpConnectionsClient by way of the
// undefined.Transporter field.
func NewVirtualHubBgpConnectionsServerTransport(srv *VirtualHubBgpConnectionsServer) *VirtualHubBgpConnectionsServerTransport {
	return &VirtualHubBgpConnectionsServerTransport{srv: srv}
}

// VirtualHubBgpConnectionsServerTransport connects instances of armnetwork.VirtualHubBgpConnectionsClient to instances of VirtualHubBgpConnectionsServer.
// Don't use this type directly, use NewVirtualHubBgpConnectionsServerTransport instead.
type VirtualHubBgpConnectionsServerTransport struct {
	srv                       *VirtualHubBgpConnectionsServer
	newListPager              *azfake.PagerResponder[armnetwork.VirtualHubBgpConnectionsClientListResponse]
	beginListAdvertisedRoutes *azfake.PollerResponder[armnetwork.VirtualHubBgpConnectionsClientListAdvertisedRoutesResponse]
	beginListLearnedRoutes    *azfake.PollerResponder[armnetwork.VirtualHubBgpConnectionsClientListLearnedRoutesResponse]
}

// Do implements the policy.Transporter interface for VirtualHubBgpConnectionsServerTransport.
func (v *VirtualHubBgpConnectionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "VirtualHubBgpConnectionsClient.NewListPager":
		resp, err = v.dispatchNewListPager(req)
	case "VirtualHubBgpConnectionsClient.BeginListAdvertisedRoutes":
		resp, err = v.dispatchBeginListAdvertisedRoutes(req)
	case "VirtualHubBgpConnectionsClient.BeginListLearnedRoutes":
		resp, err = v.dispatchBeginListLearnedRoutes(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (v *VirtualHubBgpConnectionsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if v.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("method NewListPager not implemented")}
	}
	if v.newListPager == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/virtualHubs/(?P<virtualHubName>[a-zA-Z0-9-_]+)/bgpConnections"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := v.srv.NewListPager(matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("virtualHubName")], nil)
		v.newListPager = &resp
		server.PagerResponderInjectNextLinks(v.newListPager, req, func(page *armnetwork.VirtualHubBgpConnectionsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(v.newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(v.newListPager) {
		v.newListPager = nil
	}
	return resp, nil
}

func (v *VirtualHubBgpConnectionsServerTransport) dispatchBeginListAdvertisedRoutes(req *http.Request) (*http.Response, error) {
	if v.srv.BeginListAdvertisedRoutes == nil {
		return nil, &nonRetriableError{errors.New("method BeginListAdvertisedRoutes not implemented")}
	}
	if v.beginListAdvertisedRoutes == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/virtualHubs/(?P<hubName>[a-zA-Z0-9-_]+)/bgpConnections/(?P<connectionName>[a-zA-Z0-9-_]+)/advertisedRoutes"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		respr, errRespr := v.srv.BeginListAdvertisedRoutes(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("hubName")], matches[regex.SubexpIndex("connectionName")], nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		v.beginListAdvertisedRoutes = &respr
	}

	resp, err := server.PollerResponderNext(v.beginListAdvertisedRoutes, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(v.beginListAdvertisedRoutes) {
		v.beginListAdvertisedRoutes = nil
	}

	return resp, nil
}

func (v *VirtualHubBgpConnectionsServerTransport) dispatchBeginListLearnedRoutes(req *http.Request) (*http.Response, error) {
	if v.srv.BeginListLearnedRoutes == nil {
		return nil, &nonRetriableError{errors.New("method BeginListLearnedRoutes not implemented")}
	}
	if v.beginListLearnedRoutes == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/virtualHubs/(?P<hubName>[a-zA-Z0-9-_]+)/bgpConnections/(?P<connectionName>[a-zA-Z0-9-_]+)/learnedRoutes"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		respr, errRespr := v.srv.BeginListLearnedRoutes(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("hubName")], matches[regex.SubexpIndex("connectionName")], nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		v.beginListLearnedRoutes = &respr
	}

	resp, err := server.PollerResponderNext(v.beginListLearnedRoutes, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(v.beginListLearnedRoutes) {
		v.beginListLearnedRoutes = nil
	}

	return resp, nil
}
