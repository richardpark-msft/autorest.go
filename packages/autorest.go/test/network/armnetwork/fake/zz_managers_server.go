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
	"strconv"
)

// ManagersServer is a fake server for instances of the armnetwork.ManagersClient type.
type ManagersServer struct {
	// CreateOrUpdate is the fake for method ManagersClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, networkManagerName string, parameters armnetwork.Manager, options *armnetwork.ManagersClientCreateOrUpdateOptions) (resp azfake.Responder[armnetwork.ManagersClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method ManagersClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, networkManagerName string, options *armnetwork.ManagersClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetwork.ManagersClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ManagersClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, networkManagerName string, options *armnetwork.ManagersClientGetOptions) (resp azfake.Responder[armnetwork.ManagersClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method ManagersClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, options *armnetwork.ManagersClientListOptions) (resp azfake.PagerResponder[armnetwork.ManagersClientListResponse])

	// NewListBySubscriptionPager is the fake for method ManagersClient.NewListBySubscriptionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySubscriptionPager func(options *armnetwork.ManagersClientListBySubscriptionOptions) (resp azfake.PagerResponder[armnetwork.ManagersClientListBySubscriptionResponse])

	// Patch is the fake for method ManagersClient.Patch
	// HTTP status codes to indicate success: http.StatusOK
	Patch func(ctx context.Context, resourceGroupName string, networkManagerName string, parameters armnetwork.PatchObject, options *armnetwork.ManagersClientPatchOptions) (resp azfake.Responder[armnetwork.ManagersClientPatchResponse], errResp azfake.ErrorResponder)
}

// NewManagersServerTransport creates a new instance of ManagersServerTransport with the provided implementation.
// The returned ManagersServerTransport instance is connected to an instance of armnetwork.ManagersClient by way of the
// undefined.Transporter field.
func NewManagersServerTransport(srv *ManagersServer) *ManagersServerTransport {
	return &ManagersServerTransport{srv: srv}
}

// ManagersServerTransport connects instances of armnetwork.ManagersClient to instances of ManagersServer.
// Don't use this type directly, use NewManagersServerTransport instead.
type ManagersServerTransport struct {
	srv                        *ManagersServer
	beginDelete                *azfake.PollerResponder[armnetwork.ManagersClientDeleteResponse]
	newListPager               *azfake.PagerResponder[armnetwork.ManagersClientListResponse]
	newListBySubscriptionPager *azfake.PagerResponder[armnetwork.ManagersClientListBySubscriptionResponse]
}

// Do implements the policy.Transporter interface for ManagersServerTransport.
func (m *ManagersServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ManagersClient.CreateOrUpdate":
		resp, err = m.dispatchCreateOrUpdate(req)
	case "ManagersClient.BeginDelete":
		resp, err = m.dispatchBeginDelete(req)
	case "ManagersClient.Get":
		resp, err = m.dispatchGet(req)
	case "ManagersClient.NewListPager":
		resp, err = m.dispatchNewListPager(req)
	case "ManagersClient.NewListBySubscriptionPager":
		resp, err = m.dispatchNewListBySubscriptionPager(req)
	case "ManagersClient.Patch":
		resp, err = m.dispatchPatch(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (m *ManagersServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if m.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("method CreateOrUpdate not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/networkManagers/(?P<networkManagerName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armnetwork.Manager](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := m.srv.CreateOrUpdate(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("networkManagerName")], body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Manager, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *ManagersServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if m.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("method BeginDelete not implemented")}
	}
	if m.beginDelete == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/networkManagers/(?P<networkManagerName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		forceParam, err := parseOptional(qp.Get("force"), strconv.ParseBool)
		if err != nil {
			return nil, err
		}
		var options *armnetwork.ManagersClientBeginDeleteOptions
		if forceParam != nil {
			options = &armnetwork.ManagersClientBeginDeleteOptions{
				Force: forceParam,
			}
		}
		respr, errRespr := m.srv.BeginDelete(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("networkManagerName")], options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		m.beginDelete = &respr
	}

	resp, err := server.PollerResponderNext(m.beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(m.beginDelete) {
		m.beginDelete = nil
	}

	return resp, nil
}

func (m *ManagersServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if m.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("method Get not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/networkManagers/(?P<networkManagerName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := m.srv.Get(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("networkManagerName")], nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Manager, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *ManagersServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if m.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("method NewListPager not implemented")}
	}
	if m.newListPager == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/networkManagers"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
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
		skipTokenParam := getOptional(qp.Get("$skipToken"))
		var options *armnetwork.ManagersClientListOptions
		if topParam != nil || skipTokenParam != nil {
			options = &armnetwork.ManagersClientListOptions{
				Top:       topParam,
				SkipToken: skipTokenParam,
			}
		}
		resp := m.srv.NewListPager(matches[regex.SubexpIndex("resourceGroupName")], options)
		m.newListPager = &resp
		server.PagerResponderInjectNextLinks(m.newListPager, req, func(page *armnetwork.ManagersClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(m.newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(m.newListPager) {
		m.newListPager = nil
	}
	return resp, nil
}

func (m *ManagersServerTransport) dispatchNewListBySubscriptionPager(req *http.Request) (*http.Response, error) {
	if m.srv.NewListBySubscriptionPager == nil {
		return nil, &nonRetriableError{errors.New("method NewListBySubscriptionPager not implemented")}
	}
	if m.newListBySubscriptionPager == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/networkManagers"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
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
		skipTokenParam := getOptional(qp.Get("$skipToken"))
		var options *armnetwork.ManagersClientListBySubscriptionOptions
		if topParam != nil || skipTokenParam != nil {
			options = &armnetwork.ManagersClientListBySubscriptionOptions{
				Top:       topParam,
				SkipToken: skipTokenParam,
			}
		}
		resp := m.srv.NewListBySubscriptionPager(options)
		m.newListBySubscriptionPager = &resp
		server.PagerResponderInjectNextLinks(m.newListBySubscriptionPager, req, func(page *armnetwork.ManagersClientListBySubscriptionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(m.newListBySubscriptionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(m.newListBySubscriptionPager) {
		m.newListBySubscriptionPager = nil
	}
	return resp, nil
}

func (m *ManagersServerTransport) dispatchPatch(req *http.Request) (*http.Response, error) {
	if m.srv.Patch == nil {
		return nil, &nonRetriableError{errors.New("method Patch not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/networkManagers/(?P<networkManagerName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armnetwork.PatchObject](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := m.srv.Patch(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("networkManagerName")], body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Manager, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
