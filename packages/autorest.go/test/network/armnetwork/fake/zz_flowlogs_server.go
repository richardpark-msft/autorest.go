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

// FlowLogsServer is a fake server for instances of the armnetwork.FlowLogsClient type.
type FlowLogsServer struct {
	// BeginCreateOrUpdate is the fake for method FlowLogsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, networkWatcherName string, flowLogName string, parameters armnetwork.FlowLog, options *armnetwork.FlowLogsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armnetwork.FlowLogsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method FlowLogsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, networkWatcherName string, flowLogName string, options *armnetwork.FlowLogsClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetwork.FlowLogsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method FlowLogsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, networkWatcherName string, flowLogName string, options *armnetwork.FlowLogsClientGetOptions) (resp azfake.Responder[armnetwork.FlowLogsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method FlowLogsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, networkWatcherName string, options *armnetwork.FlowLogsClientListOptions) (resp azfake.PagerResponder[armnetwork.FlowLogsClientListResponse])

	// UpdateTags is the fake for method FlowLogsClient.UpdateTags
	// HTTP status codes to indicate success: http.StatusOK
	UpdateTags func(ctx context.Context, resourceGroupName string, networkWatcherName string, flowLogName string, parameters armnetwork.TagsObject, options *armnetwork.FlowLogsClientUpdateTagsOptions) (resp azfake.Responder[armnetwork.FlowLogsClientUpdateTagsResponse], errResp azfake.ErrorResponder)
}

// NewFlowLogsServerTransport creates a new instance of FlowLogsServerTransport with the provided implementation.
// The returned FlowLogsServerTransport instance is connected to an instance of armnetwork.FlowLogsClient by way of the
// undefined.Transporter field.
func NewFlowLogsServerTransport(srv *FlowLogsServer) *FlowLogsServerTransport {
	return &FlowLogsServerTransport{srv: srv}
}

// FlowLogsServerTransport connects instances of armnetwork.FlowLogsClient to instances of FlowLogsServer.
// Don't use this type directly, use NewFlowLogsServerTransport instead.
type FlowLogsServerTransport struct {
	srv                 *FlowLogsServer
	beginCreateOrUpdate *azfake.PollerResponder[armnetwork.FlowLogsClientCreateOrUpdateResponse]
	beginDelete         *azfake.PollerResponder[armnetwork.FlowLogsClientDeleteResponse]
	newListPager        *azfake.PagerResponder[armnetwork.FlowLogsClientListResponse]
}

// Do implements the policy.Transporter interface for FlowLogsServerTransport.
func (f *FlowLogsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "FlowLogsClient.BeginCreateOrUpdate":
		resp, err = f.dispatchBeginCreateOrUpdate(req)
	case "FlowLogsClient.BeginDelete":
		resp, err = f.dispatchBeginDelete(req)
	case "FlowLogsClient.Get":
		resp, err = f.dispatchGet(req)
	case "FlowLogsClient.NewListPager":
		resp, err = f.dispatchNewListPager(req)
	case "FlowLogsClient.UpdateTags":
		resp, err = f.dispatchUpdateTags(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (f *FlowLogsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if f.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("method BeginCreateOrUpdate not implemented")}
	}
	if f.beginCreateOrUpdate == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/networkWatchers/(?P<networkWatcherName>[a-zA-Z0-9-_]+)/flowLogs/(?P<flowLogName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.FlowLog](req)
		if err != nil {
			return nil, err
		}
		respr, errRespr := f.srv.BeginCreateOrUpdate(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("networkWatcherName")], matches[regex.SubexpIndex("flowLogName")], body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		f.beginCreateOrUpdate = &respr
	}

	resp, err := server.PollerResponderNext(f.beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(f.beginCreateOrUpdate) {
		f.beginCreateOrUpdate = nil
	}

	return resp, nil
}

func (f *FlowLogsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if f.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("method BeginDelete not implemented")}
	}
	if f.beginDelete == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/networkWatchers/(?P<networkWatcherName>[a-zA-Z0-9-_]+)/flowLogs/(?P<flowLogName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		respr, errRespr := f.srv.BeginDelete(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("networkWatcherName")], matches[regex.SubexpIndex("flowLogName")], nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		f.beginDelete = &respr
	}

	resp, err := server.PollerResponderNext(f.beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(f.beginDelete) {
		f.beginDelete = nil
	}

	return resp, nil
}

func (f *FlowLogsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if f.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("method Get not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/networkWatchers/(?P<networkWatcherName>[a-zA-Z0-9-_]+)/flowLogs/(?P<flowLogName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := f.srv.Get(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("networkWatcherName")], matches[regex.SubexpIndex("flowLogName")], nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).FlowLog, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (f *FlowLogsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if f.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("method NewListPager not implemented")}
	}
	if f.newListPager == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/networkWatchers/(?P<networkWatcherName>[a-zA-Z0-9-_]+)/flowLogs"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := f.srv.NewListPager(matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("networkWatcherName")], nil)
		f.newListPager = &resp
		server.PagerResponderInjectNextLinks(f.newListPager, req, func(page *armnetwork.FlowLogsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(f.newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(f.newListPager) {
		f.newListPager = nil
	}
	return resp, nil
}

func (f *FlowLogsServerTransport) dispatchUpdateTags(req *http.Request) (*http.Response, error) {
	if f.srv.UpdateTags == nil {
		return nil, &nonRetriableError{errors.New("method UpdateTags not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/networkWatchers/(?P<networkWatcherName>[a-zA-Z0-9-_]+)/flowLogs/(?P<flowLogName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armnetwork.TagsObject](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := f.srv.UpdateTags(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("networkWatcherName")], matches[regex.SubexpIndex("flowLogName")], body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).FlowLog, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}