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

// GalleriesServer is a fake server for instances of the armcompute.GalleriesClient type.
type GalleriesServer struct {
	// BeginCreateOrUpdate is the fake for method GalleriesClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, galleryName string, gallery armcompute.Gallery, options *armcompute.GalleriesClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armcompute.GalleriesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method GalleriesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, galleryName string, options *armcompute.GalleriesClientBeginDeleteOptions) (resp azfake.PollerResponder[armcompute.GalleriesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method GalleriesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, galleryName string, options *armcompute.GalleriesClientGetOptions) (resp azfake.Responder[armcompute.GalleriesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method GalleriesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armcompute.GalleriesClientListOptions) (resp azfake.PagerResponder[armcompute.GalleriesClientListResponse])

	// NewListByResourceGroupPager is the fake for method GalleriesClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armcompute.GalleriesClientListByResourceGroupOptions) (resp azfake.PagerResponder[armcompute.GalleriesClientListByResourceGroupResponse])

	// BeginUpdate is the fake for method GalleriesClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK
	BeginUpdate func(ctx context.Context, resourceGroupName string, galleryName string, gallery armcompute.GalleryUpdate, options *armcompute.GalleriesClientBeginUpdateOptions) (resp azfake.PollerResponder[armcompute.GalleriesClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewGalleriesServerTransport creates a new instance of GalleriesServerTransport with the provided implementation.
// The returned GalleriesServerTransport instance is connected to an instance of armcompute.GalleriesClient by way of the
// undefined.Transporter field.
func NewGalleriesServerTransport(srv *GalleriesServer) *GalleriesServerTransport {
	return &GalleriesServerTransport{srv: srv}
}

// GalleriesServerTransport connects instances of armcompute.GalleriesClient to instances of GalleriesServer.
// Don't use this type directly, use NewGalleriesServerTransport instead.
type GalleriesServerTransport struct {
	srv                         *GalleriesServer
	beginCreateOrUpdate         *azfake.PollerResponder[armcompute.GalleriesClientCreateOrUpdateResponse]
	beginDelete                 *azfake.PollerResponder[armcompute.GalleriesClientDeleteResponse]
	newListPager                *azfake.PagerResponder[armcompute.GalleriesClientListResponse]
	newListByResourceGroupPager *azfake.PagerResponder[armcompute.GalleriesClientListByResourceGroupResponse]
	beginUpdate                 *azfake.PollerResponder[armcompute.GalleriesClientUpdateResponse]
}

// Do implements the policy.Transporter interface for GalleriesServerTransport.
func (g *GalleriesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "GalleriesClient.BeginCreateOrUpdate":
		resp, err = g.dispatchBeginCreateOrUpdate(req)
	case "GalleriesClient.BeginDelete":
		resp, err = g.dispatchBeginDelete(req)
	case "GalleriesClient.Get":
		resp, err = g.dispatchGet(req)
	case "GalleriesClient.NewListPager":
		resp, err = g.dispatchNewListPager(req)
	case "GalleriesClient.NewListByResourceGroupPager":
		resp, err = g.dispatchNewListByResourceGroupPager(req)
	case "GalleriesClient.BeginUpdate":
		resp, err = g.dispatchBeginUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (g *GalleriesServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if g.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("method BeginCreateOrUpdate not implemented")}
	}
	if g.beginCreateOrUpdate == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/galleries/(?P<galleryName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcompute.Gallery](req)
		if err != nil {
			return nil, err
		}
		respr, errRespr := g.srv.BeginCreateOrUpdate(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("galleryName")], body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		g.beginCreateOrUpdate = &respr
	}

	resp, err := server.PollerResponderNext(g.beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(g.beginCreateOrUpdate) {
		g.beginCreateOrUpdate = nil
	}

	return resp, nil
}

func (g *GalleriesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if g.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("method BeginDelete not implemented")}
	}
	if g.beginDelete == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/galleries/(?P<galleryName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		respr, errRespr := g.srv.BeginDelete(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("galleryName")], nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		g.beginDelete = &respr
	}

	resp, err := server.PollerResponderNext(g.beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(g.beginDelete) {
		g.beginDelete = nil
	}

	return resp, nil
}

func (g *GalleriesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if g.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("method Get not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/galleries/(?P<galleryName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	selectParam := getOptional(armcompute.SelectPermissions(qp.Get("$select")))
	expandParam := getOptional(armcompute.GalleryExpandParams(qp.Get("$expand")))
	var options *armcompute.GalleriesClientGetOptions
	if selectParam != nil || expandParam != nil {
		options = &armcompute.GalleriesClientGetOptions{
			Select: selectParam,
			Expand: expandParam,
		}
	}
	respr, errRespr := g.srv.Get(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("galleryName")], options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Gallery, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (g *GalleriesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if g.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("method NewListPager not implemented")}
	}
	if g.newListPager == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/galleries"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := g.srv.NewListPager(nil)
		g.newListPager = &resp
		server.PagerResponderInjectNextLinks(g.newListPager, req, func(page *armcompute.GalleriesClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(g.newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(g.newListPager) {
		g.newListPager = nil
	}
	return resp, nil
}

func (g *GalleriesServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if g.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("method NewListByResourceGroupPager not implemented")}
	}
	if g.newListByResourceGroupPager == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/galleries"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := g.srv.NewListByResourceGroupPager(matches[regex.SubexpIndex("resourceGroupName")], nil)
		g.newListByResourceGroupPager = &resp
		server.PagerResponderInjectNextLinks(g.newListByResourceGroupPager, req, func(page *armcompute.GalleriesClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(g.newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(g.newListByResourceGroupPager) {
		g.newListByResourceGroupPager = nil
	}
	return resp, nil
}

func (g *GalleriesServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if g.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("method BeginUpdate not implemented")}
	}
	if g.beginUpdate == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/galleries/(?P<galleryName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcompute.GalleryUpdate](req)
		if err != nil {
			return nil, err
		}
		respr, errRespr := g.srv.BeginUpdate(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("galleryName")], body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		g.beginUpdate = &respr
	}

	resp, err := server.PollerResponderNext(g.beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PollerResponderMore(g.beginUpdate) {
		g.beginUpdate = nil
	}

	return resp, nil
}