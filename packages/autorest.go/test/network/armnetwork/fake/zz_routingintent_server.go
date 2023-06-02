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

// RoutingIntentServer is a fake server for instances of the armnetwork.RoutingIntentClient type.
type RoutingIntentServer struct {
	// BeginCreateOrUpdate is the fake for method RoutingIntentClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, virtualHubName string, routingIntentName string, routingIntentParameters armnetwork.RoutingIntent, options *armnetwork.RoutingIntentClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armnetwork.RoutingIntentClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method RoutingIntentClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, virtualHubName string, routingIntentName string, options *armnetwork.RoutingIntentClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetwork.RoutingIntentClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method RoutingIntentClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, virtualHubName string, routingIntentName string, options *armnetwork.RoutingIntentClientGetOptions) (resp azfake.Responder[armnetwork.RoutingIntentClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method RoutingIntentClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, virtualHubName string, options *armnetwork.RoutingIntentClientListOptions) (resp azfake.PagerResponder[armnetwork.RoutingIntentClientListResponse])
}

// NewRoutingIntentServerTransport creates a new instance of RoutingIntentServerTransport with the provided implementation.
// The returned RoutingIntentServerTransport instance is connected to an instance of armnetwork.RoutingIntentClient by way of the
// undefined.Transporter field.
func NewRoutingIntentServerTransport(srv *RoutingIntentServer) *RoutingIntentServerTransport {
	return &RoutingIntentServerTransport{srv: srv}
}

// RoutingIntentServerTransport connects instances of armnetwork.RoutingIntentClient to instances of RoutingIntentServer.
// Don't use this type directly, use NewRoutingIntentServerTransport instead.
type RoutingIntentServerTransport struct {
	srv                 *RoutingIntentServer
	beginCreateOrUpdate *azfake.PollerResponder[armnetwork.RoutingIntentClientCreateOrUpdateResponse]
	beginDelete         *azfake.PollerResponder[armnetwork.RoutingIntentClientDeleteResponse]
	newListPager        *azfake.PagerResponder[armnetwork.RoutingIntentClientListResponse]
}

// Do implements the policy.Transporter interface for RoutingIntentServerTransport.
func (r *RoutingIntentServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "RoutingIntentClient.BeginCreateOrUpdate":
		resp, err = r.dispatchBeginCreateOrUpdate(req)
	case "RoutingIntentClient.BeginDelete":
		resp, err = r.dispatchBeginDelete(req)
	case "RoutingIntentClient.Get":
		resp, err = r.dispatchGet(req)
	case "RoutingIntentClient.NewListPager":
		resp, err = r.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *RoutingIntentServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if r.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("method BeginCreateOrUpdate not implemented")}
	}
	if r.beginCreateOrUpdate == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/virtualHubs/(?P<virtualHubName>[a-zA-Z0-9-_]+)/routingIntent/(?P<routingIntentName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.RoutingIntent](req)
		if err != nil {
			return nil, err
		}
		respr, errRespr := r.srv.BeginCreateOrUpdate(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("virtualHubName")], matches[regex.SubexpIndex("routingIntentName")], body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		r.beginCreateOrUpdate = &respr
	}

	resp, err := server.PollerResponderNext(r.beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(r.beginCreateOrUpdate) {
		r.beginCreateOrUpdate = nil
	}

	return resp, nil
}

func (r *RoutingIntentServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if r.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("method BeginDelete not implemented")}
	}
	if r.beginDelete == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/virtualHubs/(?P<virtualHubName>[a-zA-Z0-9-_]+)/routingIntent/(?P<routingIntentName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		respr, errRespr := r.srv.BeginDelete(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("virtualHubName")], matches[regex.SubexpIndex("routingIntentName")], nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		r.beginDelete = &respr
	}

	resp, err := server.PollerResponderNext(r.beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(r.beginDelete) {
		r.beginDelete = nil
	}

	return resp, nil
}

func (r *RoutingIntentServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if r.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("method Get not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/virtualHubs/(?P<virtualHubName>[a-zA-Z0-9-_]+)/routingIntent/(?P<routingIntentName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := r.srv.Get(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("virtualHubName")], matches[regex.SubexpIndex("routingIntentName")], nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).RoutingIntent, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *RoutingIntentServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if r.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("method NewListPager not implemented")}
	}
	if r.newListPager == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/virtualHubs/(?P<virtualHubName>[a-zA-Z0-9-_]+)/routingIntent"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := r.srv.NewListPager(matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("virtualHubName")], nil)
		r.newListPager = &resp
		server.PagerResponderInjectNextLinks(r.newListPager, req, func(page *armnetwork.RoutingIntentClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(r.newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(r.newListPager) {
		r.newListPager = nil
	}
	return resp, nil
}