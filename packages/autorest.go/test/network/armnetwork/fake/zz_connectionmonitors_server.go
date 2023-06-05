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
	"net/http"
	"regexp"
)

// ConnectionMonitorsServer is a fake server for instances of the armnetwork.ConnectionMonitorsClient type.
type ConnectionMonitorsServer struct {
	// BeginCreateOrUpdate is the fake for method ConnectionMonitorsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, parameters armnetwork.ConnectionMonitor, options *armnetwork.ConnectionMonitorsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armnetwork.ConnectionMonitorsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method ConnectionMonitorsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, options *armnetwork.ConnectionMonitorsClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetwork.ConnectionMonitorsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ConnectionMonitorsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, options *armnetwork.ConnectionMonitorsClientGetOptions) (resp azfake.Responder[armnetwork.ConnectionMonitorsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method ConnectionMonitorsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, networkWatcherName string, options *armnetwork.ConnectionMonitorsClientListOptions) (resp azfake.PagerResponder[armnetwork.ConnectionMonitorsClientListResponse])

	// BeginQuery is the fake for method ConnectionMonitorsClient.BeginQuery
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginQuery func(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, options *armnetwork.ConnectionMonitorsClientBeginQueryOptions) (resp azfake.PollerResponder[armnetwork.ConnectionMonitorsClientQueryResponse], errResp azfake.ErrorResponder)

	// BeginStart is the fake for method ConnectionMonitorsClient.BeginStart
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginStart func(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, options *armnetwork.ConnectionMonitorsClientBeginStartOptions) (resp azfake.PollerResponder[armnetwork.ConnectionMonitorsClientStartResponse], errResp azfake.ErrorResponder)

	// BeginStop is the fake for method ConnectionMonitorsClient.BeginStop
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginStop func(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, options *armnetwork.ConnectionMonitorsClientBeginStopOptions) (resp azfake.PollerResponder[armnetwork.ConnectionMonitorsClientStopResponse], errResp azfake.ErrorResponder)

	// UpdateTags is the fake for method ConnectionMonitorsClient.UpdateTags
	// HTTP status codes to indicate success: http.StatusOK
	UpdateTags func(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, parameters armnetwork.TagsObject, options *armnetwork.ConnectionMonitorsClientUpdateTagsOptions) (resp azfake.Responder[armnetwork.ConnectionMonitorsClientUpdateTagsResponse], errResp azfake.ErrorResponder)
}

// NewConnectionMonitorsServerTransport creates a new instance of ConnectionMonitorsServerTransport with the provided implementation.
// The returned ConnectionMonitorsServerTransport instance is connected to an instance of armnetwork.ConnectionMonitorsClient by way of the
// undefined.Transporter field.
func NewConnectionMonitorsServerTransport(srv *ConnectionMonitorsServer) *ConnectionMonitorsServerTransport {
	return &ConnectionMonitorsServerTransport{srv: srv}
}

// ConnectionMonitorsServerTransport connects instances of armnetwork.ConnectionMonitorsClient to instances of ConnectionMonitorsServer.
// Don't use this type directly, use NewConnectionMonitorsServerTransport instead.
type ConnectionMonitorsServerTransport struct {
	srv                 *ConnectionMonitorsServer
	beginCreateOrUpdate *azfake.PollerResponder[armnetwork.ConnectionMonitorsClientCreateOrUpdateResponse]
	beginDelete         *azfake.PollerResponder[armnetwork.ConnectionMonitorsClientDeleteResponse]
	newListPager        *azfake.PagerResponder[armnetwork.ConnectionMonitorsClientListResponse]
	beginQuery          *azfake.PollerResponder[armnetwork.ConnectionMonitorsClientQueryResponse]
	beginStart          *azfake.PollerResponder[armnetwork.ConnectionMonitorsClientStartResponse]
	beginStop           *azfake.PollerResponder[armnetwork.ConnectionMonitorsClientStopResponse]
}

// Do implements the policy.Transporter interface for ConnectionMonitorsServerTransport.
func (c *ConnectionMonitorsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ConnectionMonitorsClient.BeginCreateOrUpdate":
		resp, err = c.dispatchBeginCreateOrUpdate(req)
	case "ConnectionMonitorsClient.BeginDelete":
		resp, err = c.dispatchBeginDelete(req)
	case "ConnectionMonitorsClient.Get":
		resp, err = c.dispatchGet(req)
	case "ConnectionMonitorsClient.NewListPager":
		resp, err = c.dispatchNewListPager(req)
	case "ConnectionMonitorsClient.BeginQuery":
		resp, err = c.dispatchBeginQuery(req)
	case "ConnectionMonitorsClient.BeginStart":
		resp, err = c.dispatchBeginStart(req)
	case "ConnectionMonitorsClient.BeginStop":
		resp, err = c.dispatchBeginStop(req)
	case "ConnectionMonitorsClient.UpdateTags":
		resp, err = c.dispatchUpdateTags(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *ConnectionMonitorsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if c.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("method BeginCreateOrUpdate not implemented")}
	}
	if c.beginCreateOrUpdate == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/networkWatchers/(?P<networkWatcherName>[a-zA-Z0-9-_]+)/connectionMonitors/(?P<connectionMonitorName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		body, err := server.UnmarshalRequestAsJSON[armnetwork.ConnectionMonitor](req)
		if err != nil {
			return nil, err
		}
		migrateParam := getOptional(qp.Get("migrate"))
		var options *armnetwork.ConnectionMonitorsClientBeginCreateOrUpdateOptions
		if migrateParam != nil {
			options = &armnetwork.ConnectionMonitorsClientBeginCreateOrUpdateOptions{
				Migrate: migrateParam,
			}
		}
		respr, errRespr := c.srv.BeginCreateOrUpdate(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("networkWatcherName")], matches[regex.SubexpIndex("connectionMonitorName")], body, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		c.beginCreateOrUpdate = &respr
	}

	resp, err := server.PollerResponderNext(c.beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(c.beginCreateOrUpdate) {
		c.beginCreateOrUpdate = nil
	}

	return resp, nil
}

func (c *ConnectionMonitorsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if c.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("method BeginDelete not implemented")}
	}
	if c.beginDelete == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/networkWatchers/(?P<networkWatcherName>[a-zA-Z0-9-_]+)/connectionMonitors/(?P<connectionMonitorName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		respr, errRespr := c.srv.BeginDelete(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("networkWatcherName")], matches[regex.SubexpIndex("connectionMonitorName")], nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		c.beginDelete = &respr
	}

	resp, err := server.PollerResponderNext(c.beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(c.beginDelete) {
		c.beginDelete = nil
	}

	return resp, nil
}

func (c *ConnectionMonitorsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if c.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("method Get not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/networkWatchers/(?P<networkWatcherName>[a-zA-Z0-9-_]+)/connectionMonitors/(?P<connectionMonitorName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := c.srv.Get(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("networkWatcherName")], matches[regex.SubexpIndex("connectionMonitorName")], nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ConnectionMonitorResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ConnectionMonitorsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("method NewListPager not implemented")}
	}
	if c.newListPager == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/networkWatchers/(?P<networkWatcherName>[a-zA-Z0-9-_]+)/connectionMonitors"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := c.srv.NewListPager(matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("networkWatcherName")], nil)
		c.newListPager = &resp
	}
	resp, err := server.PagerResponderNext(c.newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(c.newListPager) {
		c.newListPager = nil
	}
	return resp, nil
}

func (c *ConnectionMonitorsServerTransport) dispatchBeginQuery(req *http.Request) (*http.Response, error) {
	if c.srv.BeginQuery == nil {
		return nil, &nonRetriableError{errors.New("method BeginQuery not implemented")}
	}
	if c.beginQuery == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/networkWatchers/(?P<networkWatcherName>[a-zA-Z0-9-_]+)/connectionMonitors/(?P<connectionMonitorName>[a-zA-Z0-9-_]+)/query"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		respr, errRespr := c.srv.BeginQuery(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("networkWatcherName")], matches[regex.SubexpIndex("connectionMonitorName")], nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		c.beginQuery = &respr
	}

	resp, err := server.PollerResponderNext(c.beginQuery, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(c.beginQuery) {
		c.beginQuery = nil
	}

	return resp, nil
}

func (c *ConnectionMonitorsServerTransport) dispatchBeginStart(req *http.Request) (*http.Response, error) {
	if c.srv.BeginStart == nil {
		return nil, &nonRetriableError{errors.New("method BeginStart not implemented")}
	}
	if c.beginStart == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/networkWatchers/(?P<networkWatcherName>[a-zA-Z0-9-_]+)/connectionMonitors/(?P<connectionMonitorName>[a-zA-Z0-9-_]+)/start"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		respr, errRespr := c.srv.BeginStart(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("networkWatcherName")], matches[regex.SubexpIndex("connectionMonitorName")], nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		c.beginStart = &respr
	}

	resp, err := server.PollerResponderNext(c.beginStart, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(c.beginStart) {
		c.beginStart = nil
	}

	return resp, nil
}

func (c *ConnectionMonitorsServerTransport) dispatchBeginStop(req *http.Request) (*http.Response, error) {
	if c.srv.BeginStop == nil {
		return nil, &nonRetriableError{errors.New("method BeginStop not implemented")}
	}
	if c.beginStop == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/networkWatchers/(?P<networkWatcherName>[a-zA-Z0-9-_]+)/connectionMonitors/(?P<connectionMonitorName>[a-zA-Z0-9-_]+)/stop"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		respr, errRespr := c.srv.BeginStop(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("networkWatcherName")], matches[regex.SubexpIndex("connectionMonitorName")], nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		c.beginStop = &respr
	}

	resp, err := server.PollerResponderNext(c.beginStop, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(c.beginStop) {
		c.beginStop = nil
	}

	return resp, nil
}

func (c *ConnectionMonitorsServerTransport) dispatchUpdateTags(req *http.Request) (*http.Response, error) {
	if c.srv.UpdateTags == nil {
		return nil, &nonRetriableError{errors.New("method UpdateTags not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/networkWatchers/(?P<networkWatcherName>[a-zA-Z0-9-_]+)/connectionMonitors/(?P<connectionMonitorName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armnetwork.TagsObject](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.UpdateTags(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("networkWatcherName")], matches[regex.SubexpIndex("connectionMonitorName")], body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ConnectionMonitorResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
