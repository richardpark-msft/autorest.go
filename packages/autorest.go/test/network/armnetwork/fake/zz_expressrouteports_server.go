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
	"net/url"
	"regexp"
)

// ExpressRoutePortsServer is a fake server for instances of the armnetwork.ExpressRoutePortsClient type.
type ExpressRoutePortsServer struct {
	// BeginCreateOrUpdate is the fake for method ExpressRoutePortsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, expressRoutePortName string, parameters armnetwork.ExpressRoutePort, options *armnetwork.ExpressRoutePortsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armnetwork.ExpressRoutePortsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method ExpressRoutePortsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, expressRoutePortName string, options *armnetwork.ExpressRoutePortsClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetwork.ExpressRoutePortsClientDeleteResponse], errResp azfake.ErrorResponder)

	// GenerateLOA is the fake for method ExpressRoutePortsClient.GenerateLOA
	// HTTP status codes to indicate success: http.StatusOK
	GenerateLOA func(ctx context.Context, resourceGroupName string, expressRoutePortName string, request armnetwork.GenerateExpressRoutePortsLOARequest, options *armnetwork.ExpressRoutePortsClientGenerateLOAOptions) (resp azfake.Responder[armnetwork.ExpressRoutePortsClientGenerateLOAResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ExpressRoutePortsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, expressRoutePortName string, options *armnetwork.ExpressRoutePortsClientGetOptions) (resp azfake.Responder[armnetwork.ExpressRoutePortsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method ExpressRoutePortsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armnetwork.ExpressRoutePortsClientListOptions) (resp azfake.PagerResponder[armnetwork.ExpressRoutePortsClientListResponse])

	// NewListByResourceGroupPager is the fake for method ExpressRoutePortsClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armnetwork.ExpressRoutePortsClientListByResourceGroupOptions) (resp azfake.PagerResponder[armnetwork.ExpressRoutePortsClientListByResourceGroupResponse])

	// UpdateTags is the fake for method ExpressRoutePortsClient.UpdateTags
	// HTTP status codes to indicate success: http.StatusOK
	UpdateTags func(ctx context.Context, resourceGroupName string, expressRoutePortName string, parameters armnetwork.TagsObject, options *armnetwork.ExpressRoutePortsClientUpdateTagsOptions) (resp azfake.Responder[armnetwork.ExpressRoutePortsClientUpdateTagsResponse], errResp azfake.ErrorResponder)
}

// NewExpressRoutePortsServerTransport creates a new instance of ExpressRoutePortsServerTransport with the provided implementation.
// The returned ExpressRoutePortsServerTransport instance is connected to an instance of armnetwork.ExpressRoutePortsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewExpressRoutePortsServerTransport(srv *ExpressRoutePortsServer) *ExpressRoutePortsServerTransport {
	return &ExpressRoutePortsServerTransport{
		srv:                         srv,
		beginCreateOrUpdate:         newTracker[azfake.PollerResponder[armnetwork.ExpressRoutePortsClientCreateOrUpdateResponse]](),
		beginDelete:                 newTracker[azfake.PollerResponder[armnetwork.ExpressRoutePortsClientDeleteResponse]](),
		newListPager:                newTracker[azfake.PagerResponder[armnetwork.ExpressRoutePortsClientListResponse]](),
		newListByResourceGroupPager: newTracker[azfake.PagerResponder[armnetwork.ExpressRoutePortsClientListByResourceGroupResponse]](),
	}
}

// ExpressRoutePortsServerTransport connects instances of armnetwork.ExpressRoutePortsClient to instances of ExpressRoutePortsServer.
// Don't use this type directly, use NewExpressRoutePortsServerTransport instead.
type ExpressRoutePortsServerTransport struct {
	srv                         *ExpressRoutePortsServer
	beginCreateOrUpdate         *tracker[azfake.PollerResponder[armnetwork.ExpressRoutePortsClientCreateOrUpdateResponse]]
	beginDelete                 *tracker[azfake.PollerResponder[armnetwork.ExpressRoutePortsClientDeleteResponse]]
	newListPager                *tracker[azfake.PagerResponder[armnetwork.ExpressRoutePortsClientListResponse]]
	newListByResourceGroupPager *tracker[azfake.PagerResponder[armnetwork.ExpressRoutePortsClientListByResourceGroupResponse]]
}

// Do implements the policy.Transporter interface for ExpressRoutePortsServerTransport.
func (e *ExpressRoutePortsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return e.dispatchToMethodFake(req, method)
}

func (e *ExpressRoutePortsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch method {
	case "ExpressRoutePortsClient.BeginCreateOrUpdate":
		resp, err = e.dispatchBeginCreateOrUpdate(req)
	case "ExpressRoutePortsClient.BeginDelete":
		resp, err = e.dispatchBeginDelete(req)
	case "ExpressRoutePortsClient.GenerateLOA":
		resp, err = e.dispatchGenerateLOA(req)
	case "ExpressRoutePortsClient.Get":
		resp, err = e.dispatchGet(req)
	case "ExpressRoutePortsClient.NewListPager":
		resp, err = e.dispatchNewListPager(req)
	case "ExpressRoutePortsClient.NewListByResourceGroupPager":
		resp, err = e.dispatchNewListByResourceGroupPager(req)
	case "ExpressRoutePortsClient.UpdateTags":
		resp, err = e.dispatchUpdateTags(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	return resp, err
}

func (e *ExpressRoutePortsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if e.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := e.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/ExpressRoutePorts/(?P<expressRoutePortName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.ExpressRoutePort](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		expressRoutePortNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("expressRoutePortName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := e.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, expressRoutePortNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		e.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		e.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		e.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (e *ExpressRoutePortsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if e.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := e.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/ExpressRoutePorts/(?P<expressRoutePortName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		expressRoutePortNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("expressRoutePortName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := e.srv.BeginDelete(req.Context(), resourceGroupNameParam, expressRoutePortNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		e.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		e.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		e.beginDelete.remove(req)
	}

	return resp, nil
}

func (e *ExpressRoutePortsServerTransport) dispatchGenerateLOA(req *http.Request) (*http.Response, error) {
	if e.srv.GenerateLOA == nil {
		return nil, &nonRetriableError{errors.New("fake for method GenerateLOA not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/expressRoutePorts/(?P<expressRoutePortName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/generateLoa`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armnetwork.GenerateExpressRoutePortsLOARequest](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	expressRoutePortNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("expressRoutePortName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.GenerateLOA(req.Context(), resourceGroupNameParam, expressRoutePortNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).GenerateExpressRoutePortsLOAResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *ExpressRoutePortsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if e.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/ExpressRoutePorts/(?P<expressRoutePortName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	expressRoutePortNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("expressRoutePortName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.Get(req.Context(), resourceGroupNameParam, expressRoutePortNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ExpressRoutePort, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *ExpressRoutePortsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if e.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := e.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/ExpressRoutePorts`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := e.srv.NewListPager(nil)
		newListPager = &resp
		e.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armnetwork.ExpressRoutePortsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		e.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		e.newListPager.remove(req)
	}
	return resp, nil
}

func (e *ExpressRoutePortsServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if e.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := e.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/ExpressRoutePorts`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := e.srv.NewListByResourceGroupPager(resourceGroupNameParam, nil)
		newListByResourceGroupPager = &resp
		e.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armnetwork.ExpressRoutePortsClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		e.newListByResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourceGroupPager) {
		e.newListByResourceGroupPager.remove(req)
	}
	return resp, nil
}

func (e *ExpressRoutePortsServerTransport) dispatchUpdateTags(req *http.Request) (*http.Response, error) {
	if e.srv.UpdateTags == nil {
		return nil, &nonRetriableError{errors.New("fake for method UpdateTags not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/ExpressRoutePorts/(?P<expressRoutePortName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armnetwork.TagsObject](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	expressRoutePortNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("expressRoutePortName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.UpdateTags(req.Context(), resourceGroupNameParam, expressRoutePortNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ExpressRoutePort, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
