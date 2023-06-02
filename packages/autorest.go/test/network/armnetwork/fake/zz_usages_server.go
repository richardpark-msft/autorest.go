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
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"net/http"
	"regexp"
)

// UsagesServer is a fake server for instances of the armnetwork.UsagesClient type.
type UsagesServer struct {
	// NewListPager is the fake for method UsagesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(location string, options *armnetwork.UsagesClientListOptions) (resp azfake.PagerResponder[armnetwork.UsagesClientListResponse])
}

// NewUsagesServerTransport creates a new instance of UsagesServerTransport with the provided implementation.
// The returned UsagesServerTransport instance is connected to an instance of armnetwork.UsagesClient by way of the
// undefined.Transporter field.
func NewUsagesServerTransport(srv *UsagesServer) *UsagesServerTransport {
	return &UsagesServerTransport{srv: srv}
}

// UsagesServerTransport connects instances of armnetwork.UsagesClient to instances of UsagesServer.
// Don't use this type directly, use NewUsagesServerTransport instead.
type UsagesServerTransport struct {
	srv          *UsagesServer
	newListPager *azfake.PagerResponder[armnetwork.UsagesClientListResponse]
}

// Do implements the policy.Transporter interface for UsagesServerTransport.
func (u *UsagesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "UsagesClient.NewListPager":
		resp, err = u.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (u *UsagesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if u.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("method NewListPager not implemented")}
	}
	if u.newListPager == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/locations/(?P<location>[a-zA-Z0-9-_]+)/usages"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := u.srv.NewListPager(matches[regex.SubexpIndex("location")], nil)
		u.newListPager = &resp
		server.PagerResponderInjectNextLinks(u.newListPager, req, func(page *armnetwork.UsagesClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(u.newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(u.newListPager) {
		u.newListPager = nil
	}
	return resp, nil
}