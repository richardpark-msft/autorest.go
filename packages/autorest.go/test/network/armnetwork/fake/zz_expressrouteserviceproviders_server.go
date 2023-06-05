//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"armnetwork"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"net/http"
	"regexp"
)

// ExpressRouteServiceProvidersServer is a fake server for instances of the armnetwork.ExpressRouteServiceProvidersClient type.
type ExpressRouteServiceProvidersServer struct {
	// NewListPager is the fake for method ExpressRouteServiceProvidersClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armnetwork.ExpressRouteServiceProvidersClientListOptions) (resp azfake.PagerResponder[armnetwork.ExpressRouteServiceProvidersClientListResponse])
}

// NewExpressRouteServiceProvidersServerTransport creates a new instance of ExpressRouteServiceProvidersServerTransport with the provided implementation.
// The returned ExpressRouteServiceProvidersServerTransport instance is connected to an instance of armnetwork.ExpressRouteServiceProvidersClient by way of the
// undefined.Transporter field.
func NewExpressRouteServiceProvidersServerTransport(srv *ExpressRouteServiceProvidersServer) *ExpressRouteServiceProvidersServerTransport {
	return &ExpressRouteServiceProvidersServerTransport{srv: srv}
}

// ExpressRouteServiceProvidersServerTransport connects instances of armnetwork.ExpressRouteServiceProvidersClient to instances of ExpressRouteServiceProvidersServer.
// Don't use this type directly, use NewExpressRouteServiceProvidersServerTransport instead.
type ExpressRouteServiceProvidersServerTransport struct {
	srv          *ExpressRouteServiceProvidersServer
	newListPager *azfake.PagerResponder[armnetwork.ExpressRouteServiceProvidersClientListResponse]
}

// Do implements the policy.Transporter interface for ExpressRouteServiceProvidersServerTransport.
func (e *ExpressRouteServiceProvidersServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ExpressRouteServiceProvidersClient.NewListPager":
		resp, err = e.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (e *ExpressRouteServiceProvidersServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if e.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("method NewListPager not implemented")}
	}
	if e.newListPager == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/expressRouteServiceProviders"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := e.srv.NewListPager(nil)
		e.newListPager = &resp
		server.PagerResponderInjectNextLinks(e.newListPager, req, func(page *armnetwork.ExpressRouteServiceProvidersClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(e.newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(e.newListPager) {
		e.newListPager = nil
	}
	return resp, nil
}
