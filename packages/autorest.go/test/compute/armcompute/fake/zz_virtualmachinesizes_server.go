//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"armcompute"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"regexp"
)

// VirtualMachineSizesServer is a fake server for instances of the armcompute.VirtualMachineSizesClient type.
type VirtualMachineSizesServer struct {
	// NewListPager is the fake for method VirtualMachineSizesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(location string, options *armcompute.VirtualMachineSizesClientListOptions) (resp azfake.PagerResponder[armcompute.VirtualMachineSizesClientListResponse])
}

// NewVirtualMachineSizesServerTransport creates a new instance of VirtualMachineSizesServerTransport with the provided implementation.
// The returned VirtualMachineSizesServerTransport instance is connected to an instance of armcompute.VirtualMachineSizesClient by way of the
// undefined.Transporter field.
func NewVirtualMachineSizesServerTransport(srv *VirtualMachineSizesServer) *VirtualMachineSizesServerTransport {
	return &VirtualMachineSizesServerTransport{srv: srv}
}

// VirtualMachineSizesServerTransport connects instances of armcompute.VirtualMachineSizesClient to instances of VirtualMachineSizesServer.
// Don't use this type directly, use NewVirtualMachineSizesServerTransport instead.
type VirtualMachineSizesServerTransport struct {
	srv          *VirtualMachineSizesServer
	newListPager *azfake.PagerResponder[armcompute.VirtualMachineSizesClientListResponse]
}

// Do implements the policy.Transporter interface for VirtualMachineSizesServerTransport.
func (v *VirtualMachineSizesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "VirtualMachineSizesClient.NewListPager":
		resp, err = v.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (v *VirtualMachineSizesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if v.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("method NewListPager not implemented")}
	}
	if v.newListPager == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/locations/(?P<location>[a-zA-Z0-9-_]+)/vmSizes"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := v.srv.NewListPager(matches[regex.SubexpIndex("location")], nil)
		v.newListPager = &resp
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
