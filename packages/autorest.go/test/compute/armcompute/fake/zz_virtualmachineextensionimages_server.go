//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"armcompute"
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"regexp"
	"strconv"
)

// VirtualMachineExtensionImagesServer is a fake server for instances of the armcompute.VirtualMachineExtensionImagesClient type.
type VirtualMachineExtensionImagesServer struct {
	// Get is the fake for method VirtualMachineExtensionImagesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, location string, publisherName string, typeParam string, version string, options *armcompute.VirtualMachineExtensionImagesClientGetOptions) (resp azfake.Responder[armcompute.VirtualMachineExtensionImagesClientGetResponse], errResp azfake.ErrorResponder)

	// ListTypes is the fake for method VirtualMachineExtensionImagesClient.ListTypes
	// HTTP status codes to indicate success: http.StatusOK
	ListTypes func(ctx context.Context, location string, publisherName string, options *armcompute.VirtualMachineExtensionImagesClientListTypesOptions) (resp azfake.Responder[armcompute.VirtualMachineExtensionImagesClientListTypesResponse], errResp azfake.ErrorResponder)

	// ListVersions is the fake for method VirtualMachineExtensionImagesClient.ListVersions
	// HTTP status codes to indicate success: http.StatusOK
	ListVersions func(ctx context.Context, location string, publisherName string, typeParam string, options *armcompute.VirtualMachineExtensionImagesClientListVersionsOptions) (resp azfake.Responder[armcompute.VirtualMachineExtensionImagesClientListVersionsResponse], errResp azfake.ErrorResponder)
}

// NewVirtualMachineExtensionImagesServerTransport creates a new instance of VirtualMachineExtensionImagesServerTransport with the provided implementation.
// The returned VirtualMachineExtensionImagesServerTransport instance is connected to an instance of armcompute.VirtualMachineExtensionImagesClient by way of the
// undefined.Transporter field.
func NewVirtualMachineExtensionImagesServerTransport(srv *VirtualMachineExtensionImagesServer) *VirtualMachineExtensionImagesServerTransport {
	return &VirtualMachineExtensionImagesServerTransport{srv: srv}
}

// VirtualMachineExtensionImagesServerTransport connects instances of armcompute.VirtualMachineExtensionImagesClient to instances of VirtualMachineExtensionImagesServer.
// Don't use this type directly, use NewVirtualMachineExtensionImagesServerTransport instead.
type VirtualMachineExtensionImagesServerTransport struct {
	srv *VirtualMachineExtensionImagesServer
}

// Do implements the policy.Transporter interface for VirtualMachineExtensionImagesServerTransport.
func (v *VirtualMachineExtensionImagesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "VirtualMachineExtensionImagesClient.Get":
		resp, err = v.dispatchGet(req)
	case "VirtualMachineExtensionImagesClient.ListTypes":
		resp, err = v.dispatchListTypes(req)
	case "VirtualMachineExtensionImagesClient.ListVersions":
		resp, err = v.dispatchListVersions(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (v *VirtualMachineExtensionImagesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if v.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("method Get not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/locations/(?P<location>[a-zA-Z0-9-_]+)/publishers/(?P<publisherName>[a-zA-Z0-9-_]+)/artifacttypes/vmextension/types/(?P<type>[a-zA-Z0-9-_]+)/versions/(?P<version>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := v.srv.Get(req.Context(), matches[regex.SubexpIndex("location")], matches[regex.SubexpIndex("publisherName")], matches[regex.SubexpIndex("type")], matches[regex.SubexpIndex("version")], nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).VirtualMachineExtensionImage, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VirtualMachineExtensionImagesServerTransport) dispatchListTypes(req *http.Request) (*http.Response, error) {
	if v.srv.ListTypes == nil {
		return nil, &nonRetriableError{errors.New("method ListTypes not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/locations/(?P<location>[a-zA-Z0-9-_]+)/publishers/(?P<publisherName>[a-zA-Z0-9-_]+)/artifacttypes/vmextension/types"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := v.srv.ListTypes(req.Context(), matches[regex.SubexpIndex("location")], matches[regex.SubexpIndex("publisherName")], nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).VirtualMachineExtensionImageArray, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VirtualMachineExtensionImagesServerTransport) dispatchListVersions(req *http.Request) (*http.Response, error) {
	if v.srv.ListVersions == nil {
		return nil, &nonRetriableError{errors.New("method ListVersions not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/locations/(?P<location>[a-zA-Z0-9-_]+)/publishers/(?P<publisherName>[a-zA-Z0-9-_]+)/artifacttypes/vmextension/types/(?P<type>[a-zA-Z0-9-_]+)/versions"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	filterParam := getOptional(qp.Get("$filter"))
	topParam, err := parseOptional(qp.Get("$top"), func(v string) (int32, error) {
		p, parseErr := strconv.ParseInt(v, 10, 32)
		if parseErr != nil {
			return 0, parseErr
		}
		return int32(p), nil
	})
	if err != nil {
		return nil, err
	}
	orderbyParam := getOptional(qp.Get("$orderby"))
	var options *armcompute.VirtualMachineExtensionImagesClientListVersionsOptions
	if filterParam != nil || topParam != nil || orderbyParam != nil {
		options = &armcompute.VirtualMachineExtensionImagesClientListVersionsOptions{
			Filter:  filterParam,
			Top:     topParam,
			Orderby: orderbyParam,
		}
	}
	respr, errRespr := v.srv.ListVersions(req.Context(), matches[regex.SubexpIndex("location")], matches[regex.SubexpIndex("publisherName")], matches[regex.SubexpIndex("type")], options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).VirtualMachineExtensionImageArray, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
