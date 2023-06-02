//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package fake

import (
	"armcompute"
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

// DiskAccessesServer is a fake server for instances of the armcompute.DiskAccessesClient type.
type DiskAccessesServer struct {
	// BeginCreateOrUpdate is the fake for method DiskAccessesClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, diskAccessName string, diskAccess armcompute.DiskAccess, options *armcompute.DiskAccessesClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armcompute.DiskAccessesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method DiskAccessesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, diskAccessName string, options *armcompute.DiskAccessesClientBeginDeleteOptions) (resp azfake.PollerResponder[armcompute.DiskAccessesClientDeleteResponse], errResp azfake.ErrorResponder)

	// BeginDeleteAPrivateEndpointConnection is the fake for method DiskAccessesClient.BeginDeleteAPrivateEndpointConnection
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDeleteAPrivateEndpointConnection func(ctx context.Context, resourceGroupName string, diskAccessName string, privateEndpointConnectionName string, options *armcompute.DiskAccessesClientBeginDeleteAPrivateEndpointConnectionOptions) (resp azfake.PollerResponder[armcompute.DiskAccessesClientDeleteAPrivateEndpointConnectionResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method DiskAccessesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, diskAccessName string, options *armcompute.DiskAccessesClientGetOptions) (resp azfake.Responder[armcompute.DiskAccessesClientGetResponse], errResp azfake.ErrorResponder)

	// GetAPrivateEndpointConnection is the fake for method DiskAccessesClient.GetAPrivateEndpointConnection
	// HTTP status codes to indicate success: http.StatusOK
	GetAPrivateEndpointConnection func(ctx context.Context, resourceGroupName string, diskAccessName string, privateEndpointConnectionName string, options *armcompute.DiskAccessesClientGetAPrivateEndpointConnectionOptions) (resp azfake.Responder[armcompute.DiskAccessesClientGetAPrivateEndpointConnectionResponse], errResp azfake.ErrorResponder)

	// GetPrivateLinkResources is the fake for method DiskAccessesClient.GetPrivateLinkResources
	// HTTP status codes to indicate success: http.StatusOK
	GetPrivateLinkResources func(ctx context.Context, resourceGroupName string, diskAccessName string, options *armcompute.DiskAccessesClientGetPrivateLinkResourcesOptions) (resp azfake.Responder[armcompute.DiskAccessesClientGetPrivateLinkResourcesResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method DiskAccessesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armcompute.DiskAccessesClientListOptions) (resp azfake.PagerResponder[armcompute.DiskAccessesClientListResponse])

	// NewListByResourceGroupPager is the fake for method DiskAccessesClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armcompute.DiskAccessesClientListByResourceGroupOptions) (resp azfake.PagerResponder[armcompute.DiskAccessesClientListByResourceGroupResponse])

	// NewListPrivateEndpointConnectionsPager is the fake for method DiskAccessesClient.NewListPrivateEndpointConnectionsPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPrivateEndpointConnectionsPager func(resourceGroupName string, diskAccessName string, options *armcompute.DiskAccessesClientListPrivateEndpointConnectionsOptions) (resp azfake.PagerResponder[armcompute.DiskAccessesClientListPrivateEndpointConnectionsResponse])

	// BeginUpdate is the fake for method DiskAccessesClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, diskAccessName string, diskAccess armcompute.DiskAccessUpdate, options *armcompute.DiskAccessesClientBeginUpdateOptions) (resp azfake.PollerResponder[armcompute.DiskAccessesClientUpdateResponse], errResp azfake.ErrorResponder)

	// BeginUpdateAPrivateEndpointConnection is the fake for method DiskAccessesClient.BeginUpdateAPrivateEndpointConnection
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdateAPrivateEndpointConnection func(ctx context.Context, resourceGroupName string, diskAccessName string, privateEndpointConnectionName string, privateEndpointConnection armcompute.PrivateEndpointConnection, options *armcompute.DiskAccessesClientBeginUpdateAPrivateEndpointConnectionOptions) (resp azfake.PollerResponder[armcompute.DiskAccessesClientUpdateAPrivateEndpointConnectionResponse], errResp azfake.ErrorResponder)
}

// NewDiskAccessesServerTransport creates a new instance of DiskAccessesServerTransport with the provided implementation.
// The returned DiskAccessesServerTransport instance is connected to an instance of armcompute.DiskAccessesClient by way of the
// undefined.Transporter field.
func NewDiskAccessesServerTransport(srv *DiskAccessesServer) *DiskAccessesServerTransport {
	return &DiskAccessesServerTransport{srv: srv}
}

// DiskAccessesServerTransport connects instances of armcompute.DiskAccessesClient to instances of DiskAccessesServer.
// Don't use this type directly, use NewDiskAccessesServerTransport instead.
type DiskAccessesServerTransport struct {
	srv                                    *DiskAccessesServer
	beginCreateOrUpdate                    *azfake.PollerResponder[armcompute.DiskAccessesClientCreateOrUpdateResponse]
	beginDelete                            *azfake.PollerResponder[armcompute.DiskAccessesClientDeleteResponse]
	beginDeleteAPrivateEndpointConnection  *azfake.PollerResponder[armcompute.DiskAccessesClientDeleteAPrivateEndpointConnectionResponse]
	newListPager                           *azfake.PagerResponder[armcompute.DiskAccessesClientListResponse]
	newListByResourceGroupPager            *azfake.PagerResponder[armcompute.DiskAccessesClientListByResourceGroupResponse]
	newListPrivateEndpointConnectionsPager *azfake.PagerResponder[armcompute.DiskAccessesClientListPrivateEndpointConnectionsResponse]
	beginUpdate                            *azfake.PollerResponder[armcompute.DiskAccessesClientUpdateResponse]
	beginUpdateAPrivateEndpointConnection  *azfake.PollerResponder[armcompute.DiskAccessesClientUpdateAPrivateEndpointConnectionResponse]
}

// Do implements the policy.Transporter interface for DiskAccessesServerTransport.
func (d *DiskAccessesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "DiskAccessesClient.BeginCreateOrUpdate":
		resp, err = d.dispatchBeginCreateOrUpdate(req)
	case "DiskAccessesClient.BeginDelete":
		resp, err = d.dispatchBeginDelete(req)
	case "DiskAccessesClient.BeginDeleteAPrivateEndpointConnection":
		resp, err = d.dispatchBeginDeleteAPrivateEndpointConnection(req)
	case "DiskAccessesClient.Get":
		resp, err = d.dispatchGet(req)
	case "DiskAccessesClient.GetAPrivateEndpointConnection":
		resp, err = d.dispatchGetAPrivateEndpointConnection(req)
	case "DiskAccessesClient.GetPrivateLinkResources":
		resp, err = d.dispatchGetPrivateLinkResources(req)
	case "DiskAccessesClient.NewListPager":
		resp, err = d.dispatchNewListPager(req)
	case "DiskAccessesClient.NewListByResourceGroupPager":
		resp, err = d.dispatchNewListByResourceGroupPager(req)
	case "DiskAccessesClient.NewListPrivateEndpointConnectionsPager":
		resp, err = d.dispatchNewListPrivateEndpointConnectionsPager(req)
	case "DiskAccessesClient.BeginUpdate":
		resp, err = d.dispatchBeginUpdate(req)
	case "DiskAccessesClient.BeginUpdateAPrivateEndpointConnection":
		resp, err = d.dispatchBeginUpdateAPrivateEndpointConnection(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (d *DiskAccessesServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if d.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("method BeginCreateOrUpdate not implemented")}
	}
	if d.beginCreateOrUpdate == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/diskAccesses/(?P<diskAccessName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcompute.DiskAccess](req)
		if err != nil {
			return nil, err
		}
		respr, errRespr := d.srv.BeginCreateOrUpdate(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("diskAccessName")], body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		d.beginCreateOrUpdate = &respr
	}

	resp, err := server.PollerResponderNext(d.beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(d.beginCreateOrUpdate) {
		d.beginCreateOrUpdate = nil
	}

	return resp, nil
}

func (d *DiskAccessesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if d.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("method BeginDelete not implemented")}
	}
	if d.beginDelete == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/diskAccesses/(?P<diskAccessName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		respr, errRespr := d.srv.BeginDelete(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("diskAccessName")], nil)
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

func (d *DiskAccessesServerTransport) dispatchBeginDeleteAPrivateEndpointConnection(req *http.Request) (*http.Response, error) {
	if d.srv.BeginDeleteAPrivateEndpointConnection == nil {
		return nil, &nonRetriableError{errors.New("method BeginDeleteAPrivateEndpointConnection not implemented")}
	}
	if d.beginDeleteAPrivateEndpointConnection == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/diskAccesses/(?P<diskAccessName>[a-zA-Z0-9-_]+)/privateEndpointConnections/(?P<privateEndpointConnectionName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		respr, errRespr := d.srv.BeginDeleteAPrivateEndpointConnection(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("diskAccessName")], matches[regex.SubexpIndex("privateEndpointConnectionName")], nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		d.beginDeleteAPrivateEndpointConnection = &respr
	}

	resp, err := server.PollerResponderNext(d.beginDeleteAPrivateEndpointConnection, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(d.beginDeleteAPrivateEndpointConnection) {
		d.beginDeleteAPrivateEndpointConnection = nil
	}

	return resp, nil
}

func (d *DiskAccessesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if d.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("method Get not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/diskAccesses/(?P<diskAccessName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := d.srv.Get(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("diskAccessName")], nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DiskAccess, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DiskAccessesServerTransport) dispatchGetAPrivateEndpointConnection(req *http.Request) (*http.Response, error) {
	if d.srv.GetAPrivateEndpointConnection == nil {
		return nil, &nonRetriableError{errors.New("method GetAPrivateEndpointConnection not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/diskAccesses/(?P<diskAccessName>[a-zA-Z0-9-_]+)/privateEndpointConnections/(?P<privateEndpointConnectionName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := d.srv.GetAPrivateEndpointConnection(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("diskAccessName")], matches[regex.SubexpIndex("privateEndpointConnectionName")], nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PrivateEndpointConnection, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DiskAccessesServerTransport) dispatchGetPrivateLinkResources(req *http.Request) (*http.Response, error) {
	if d.srv.GetPrivateLinkResources == nil {
		return nil, &nonRetriableError{errors.New("method GetPrivateLinkResources not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/diskAccesses/(?P<diskAccessName>[a-zA-Z0-9-_]+)/privateLinkResources"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := d.srv.GetPrivateLinkResources(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("diskAccessName")], nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PrivateLinkResourceListResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DiskAccessesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if d.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("method NewListPager not implemented")}
	}
	if d.newListPager == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/diskAccesses"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := d.srv.NewListPager(nil)
		d.newListPager = &resp
		server.PagerResponderInjectNextLinks(d.newListPager, req, func(page *armcompute.DiskAccessesClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(d.newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(d.newListPager) {
		d.newListPager = nil
	}
	return resp, nil
}

func (d *DiskAccessesServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if d.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("method NewListByResourceGroupPager not implemented")}
	}
	if d.newListByResourceGroupPager == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/diskAccesses"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := d.srv.NewListByResourceGroupPager(matches[regex.SubexpIndex("resourceGroupName")], nil)
		d.newListByResourceGroupPager = &resp
		server.PagerResponderInjectNextLinks(d.newListByResourceGroupPager, req, func(page *armcompute.DiskAccessesClientListByResourceGroupResponse, createLink func() string) {
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

func (d *DiskAccessesServerTransport) dispatchNewListPrivateEndpointConnectionsPager(req *http.Request) (*http.Response, error) {
	if d.srv.NewListPrivateEndpointConnectionsPager == nil {
		return nil, &nonRetriableError{errors.New("method NewListPrivateEndpointConnectionsPager not implemented")}
	}
	if d.newListPrivateEndpointConnectionsPager == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/diskAccesses/(?P<diskAccessName>[a-zA-Z0-9-_]+)/privateEndpointConnections"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := d.srv.NewListPrivateEndpointConnectionsPager(matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("diskAccessName")], nil)
		d.newListPrivateEndpointConnectionsPager = &resp
		server.PagerResponderInjectNextLinks(d.newListPrivateEndpointConnectionsPager, req, func(page *armcompute.DiskAccessesClientListPrivateEndpointConnectionsResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(d.newListPrivateEndpointConnectionsPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(d.newListPrivateEndpointConnectionsPager) {
		d.newListPrivateEndpointConnectionsPager = nil
	}
	return resp, nil
}

func (d *DiskAccessesServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if d.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("method BeginUpdate not implemented")}
	}
	if d.beginUpdate == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/diskAccesses/(?P<diskAccessName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcompute.DiskAccessUpdate](req)
		if err != nil {
			return nil, err
		}
		respr, errRespr := d.srv.BeginUpdate(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("diskAccessName")], body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		d.beginUpdate = &respr
	}

	resp, err := server.PollerResponderNext(d.beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(d.beginUpdate) {
		d.beginUpdate = nil
	}

	return resp, nil
}

func (d *DiskAccessesServerTransport) dispatchBeginUpdateAPrivateEndpointConnection(req *http.Request) (*http.Response, error) {
	if d.srv.BeginUpdateAPrivateEndpointConnection == nil {
		return nil, &nonRetriableError{errors.New("method BeginUpdateAPrivateEndpointConnection not implemented")}
	}
	if d.beginUpdateAPrivateEndpointConnection == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/diskAccesses/(?P<diskAccessName>[a-zA-Z0-9-_]+)/privateEndpointConnections/(?P<privateEndpointConnectionName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcompute.PrivateEndpointConnection](req)
		if err != nil {
			return nil, err
		}
		respr, errRespr := d.srv.BeginUpdateAPrivateEndpointConnection(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("diskAccessName")], matches[regex.SubexpIndex("privateEndpointConnectionName")], body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		d.beginUpdateAPrivateEndpointConnection = &respr
	}

	resp, err := server.PollerResponderNext(d.beginUpdateAPrivateEndpointConnection, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(d.beginUpdateAPrivateEndpointConnection) {
		d.beginUpdateAPrivateEndpointConnection = nil
	}

	return resp, nil
}