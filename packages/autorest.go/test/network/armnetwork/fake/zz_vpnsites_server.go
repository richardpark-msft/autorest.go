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

// VPNSitesServer is a fake server for instances of the armnetwork.VPNSitesClient type.
type VPNSitesServer struct {
	// BeginCreateOrUpdate is the fake for method VPNSitesClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, vpnSiteName string, vpnSiteParameters armnetwork.VPNSite, options *armnetwork.VPNSitesClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armnetwork.VPNSitesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method VPNSitesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, vpnSiteName string, options *armnetwork.VPNSitesClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetwork.VPNSitesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method VPNSitesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, vpnSiteName string, options *armnetwork.VPNSitesClientGetOptions) (resp azfake.Responder[armnetwork.VPNSitesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method VPNSitesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armnetwork.VPNSitesClientListOptions) (resp azfake.PagerResponder[armnetwork.VPNSitesClientListResponse])

	// NewListByResourceGroupPager is the fake for method VPNSitesClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armnetwork.VPNSitesClientListByResourceGroupOptions) (resp azfake.PagerResponder[armnetwork.VPNSitesClientListByResourceGroupResponse])

	// UpdateTags is the fake for method VPNSitesClient.UpdateTags
	// HTTP status codes to indicate success: http.StatusOK
	UpdateTags func(ctx context.Context, resourceGroupName string, vpnSiteName string, vpnSiteParameters armnetwork.TagsObject, options *armnetwork.VPNSitesClientUpdateTagsOptions) (resp azfake.Responder[armnetwork.VPNSitesClientUpdateTagsResponse], errResp azfake.ErrorResponder)
}

// NewVPNSitesServerTransport creates a new instance of VPNSitesServerTransport with the provided implementation.
// The returned VPNSitesServerTransport instance is connected to an instance of armnetwork.VPNSitesClient by way of the
// undefined.Transporter field.
func NewVPNSitesServerTransport(srv *VPNSitesServer) *VPNSitesServerTransport {
	return &VPNSitesServerTransport{srv: srv}
}

// VPNSitesServerTransport connects instances of armnetwork.VPNSitesClient to instances of VPNSitesServer.
// Don't use this type directly, use NewVPNSitesServerTransport instead.
type VPNSitesServerTransport struct {
	srv                         *VPNSitesServer
	beginCreateOrUpdate         *azfake.PollerResponder[armnetwork.VPNSitesClientCreateOrUpdateResponse]
	beginDelete                 *azfake.PollerResponder[armnetwork.VPNSitesClientDeleteResponse]
	newListPager                *azfake.PagerResponder[armnetwork.VPNSitesClientListResponse]
	newListByResourceGroupPager *azfake.PagerResponder[armnetwork.VPNSitesClientListByResourceGroupResponse]
}

// Do implements the policy.Transporter interface for VPNSitesServerTransport.
func (v *VPNSitesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "VPNSitesClient.BeginCreateOrUpdate":
		resp, err = v.dispatchBeginCreateOrUpdate(req)
	case "VPNSitesClient.BeginDelete":
		resp, err = v.dispatchBeginDelete(req)
	case "VPNSitesClient.Get":
		resp, err = v.dispatchGet(req)
	case "VPNSitesClient.NewListPager":
		resp, err = v.dispatchNewListPager(req)
	case "VPNSitesClient.NewListByResourceGroupPager":
		resp, err = v.dispatchNewListByResourceGroupPager(req)
	case "VPNSitesClient.UpdateTags":
		resp, err = v.dispatchUpdateTags(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (v *VPNSitesServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if v.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("method BeginCreateOrUpdate not implemented")}
	}
	if v.beginCreateOrUpdate == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/vpnSites/(?P<vpnSiteName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.VPNSite](req)
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginCreateOrUpdate(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("vpnSiteName")], body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		v.beginCreateOrUpdate = &respr
	}

	resp, err := server.PollerResponderNext(v.beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(v.beginCreateOrUpdate) {
		v.beginCreateOrUpdate = nil
	}

	return resp, nil
}

func (v *VPNSitesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if v.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("method BeginDelete not implemented")}
	}
	if v.beginDelete == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/vpnSites/(?P<vpnSiteName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		respr, errRespr := v.srv.BeginDelete(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("vpnSiteName")], nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		v.beginDelete = &respr
	}

	resp, err := server.PollerResponderNext(v.beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(v.beginDelete) {
		v.beginDelete = nil
	}

	return resp, nil
}

func (v *VPNSitesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if v.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("method Get not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/vpnSites/(?P<vpnSiteName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := v.srv.Get(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("vpnSiteName")], nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).VPNSite, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VPNSitesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if v.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("method NewListPager not implemented")}
	}
	if v.newListPager == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/vpnSites"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := v.srv.NewListPager(nil)
		v.newListPager = &resp
		server.PagerResponderInjectNextLinks(v.newListPager, req, func(page *armnetwork.VPNSitesClientListResponse, createLink func() string) {
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

func (v *VPNSitesServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if v.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("method NewListByResourceGroupPager not implemented")}
	}
	if v.newListByResourceGroupPager == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/vpnSites"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := v.srv.NewListByResourceGroupPager(matches[regex.SubexpIndex("resourceGroupName")], nil)
		v.newListByResourceGroupPager = &resp
		server.PagerResponderInjectNextLinks(v.newListByResourceGroupPager, req, func(page *armnetwork.VPNSitesClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(v.newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(v.newListByResourceGroupPager) {
		v.newListByResourceGroupPager = nil
	}
	return resp, nil
}

func (v *VPNSitesServerTransport) dispatchUpdateTags(req *http.Request) (*http.Response, error) {
	if v.srv.UpdateTags == nil {
		return nil, &nonRetriableError{errors.New("method UpdateTags not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/vpnSites/(?P<vpnSiteName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armnetwork.TagsObject](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := v.srv.UpdateTags(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("vpnSiteName")], body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).VPNSite, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
