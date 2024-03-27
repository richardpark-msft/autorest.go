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
	"net/url"
	"regexp"
)

// FirewallPolicyIdpsSignaturesOverridesServer is a fake server for instances of the armnetwork.FirewallPolicyIdpsSignaturesOverridesClient type.
type FirewallPolicyIdpsSignaturesOverridesServer struct {
	// Get is the fake for method FirewallPolicyIdpsSignaturesOverridesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, firewallPolicyName string, options *armnetwork.FirewallPolicyIdpsSignaturesOverridesClientGetOptions) (resp azfake.Responder[armnetwork.FirewallPolicyIdpsSignaturesOverridesClientGetResponse], errResp azfake.ErrorResponder)

	// List is the fake for method FirewallPolicyIdpsSignaturesOverridesClient.List
	// HTTP status codes to indicate success: http.StatusOK
	List func(ctx context.Context, resourceGroupName string, firewallPolicyName string, options *armnetwork.FirewallPolicyIdpsSignaturesOverridesClientListOptions) (resp azfake.Responder[armnetwork.FirewallPolicyIdpsSignaturesOverridesClientListResponse], errResp azfake.ErrorResponder)

	// Patch is the fake for method FirewallPolicyIdpsSignaturesOverridesClient.Patch
	// HTTP status codes to indicate success: http.StatusOK
	Patch func(ctx context.Context, resourceGroupName string, firewallPolicyName string, parameters armnetwork.SignaturesOverrides, options *armnetwork.FirewallPolicyIdpsSignaturesOverridesClientPatchOptions) (resp azfake.Responder[armnetwork.FirewallPolicyIdpsSignaturesOverridesClientPatchResponse], errResp azfake.ErrorResponder)

	// Put is the fake for method FirewallPolicyIdpsSignaturesOverridesClient.Put
	// HTTP status codes to indicate success: http.StatusOK
	Put func(ctx context.Context, resourceGroupName string, firewallPolicyName string, parameters armnetwork.SignaturesOverrides, options *armnetwork.FirewallPolicyIdpsSignaturesOverridesClientPutOptions) (resp azfake.Responder[armnetwork.FirewallPolicyIdpsSignaturesOverridesClientPutResponse], errResp azfake.ErrorResponder)
}

// NewFirewallPolicyIdpsSignaturesOverridesServerTransport creates a new instance of FirewallPolicyIdpsSignaturesOverridesServerTransport with the provided implementation.
// The returned FirewallPolicyIdpsSignaturesOverridesServerTransport instance is connected to an instance of armnetwork.FirewallPolicyIdpsSignaturesOverridesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewFirewallPolicyIdpsSignaturesOverridesServerTransport(srv *FirewallPolicyIdpsSignaturesOverridesServer) *FirewallPolicyIdpsSignaturesOverridesServerTransport {
	return &FirewallPolicyIdpsSignaturesOverridesServerTransport{srv: srv}
}

// FirewallPolicyIdpsSignaturesOverridesServerTransport connects instances of armnetwork.FirewallPolicyIdpsSignaturesOverridesClient to instances of FirewallPolicyIdpsSignaturesOverridesServer.
// Don't use this type directly, use NewFirewallPolicyIdpsSignaturesOverridesServerTransport instead.
type FirewallPolicyIdpsSignaturesOverridesServerTransport struct {
	srv *FirewallPolicyIdpsSignaturesOverridesServer
}

// Do implements the policy.Transporter interface for FirewallPolicyIdpsSignaturesOverridesServerTransport.
func (f *FirewallPolicyIdpsSignaturesOverridesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return f.dispatchToMethodFake(req, method)
}

func (f *FirewallPolicyIdpsSignaturesOverridesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch method {
	case "FirewallPolicyIdpsSignaturesOverridesClient.Get":
		resp, err = f.dispatchGet(req)
	case "FirewallPolicyIdpsSignaturesOverridesClient.List":
		resp, err = f.dispatchList(req)
	case "FirewallPolicyIdpsSignaturesOverridesClient.Patch":
		resp, err = f.dispatchPatch(req)
	case "FirewallPolicyIdpsSignaturesOverridesClient.Put":
		resp, err = f.dispatchPut(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	return resp, err
}

func (f *FirewallPolicyIdpsSignaturesOverridesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if f.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/firewallPolicies/(?P<firewallPolicyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/signatureOverrides/default`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	firewallPolicyNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("firewallPolicyName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := f.srv.Get(req.Context(), resourceGroupNameParam, firewallPolicyNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SignaturesOverrides, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (f *FirewallPolicyIdpsSignaturesOverridesServerTransport) dispatchList(req *http.Request) (*http.Response, error) {
	if f.srv.List == nil {
		return nil, &nonRetriableError{errors.New("fake for method List not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/firewallPolicies/(?P<firewallPolicyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/signatureOverrides`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	firewallPolicyNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("firewallPolicyName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := f.srv.List(req.Context(), resourceGroupNameParam, firewallPolicyNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SignaturesOverridesList, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (f *FirewallPolicyIdpsSignaturesOverridesServerTransport) dispatchPatch(req *http.Request) (*http.Response, error) {
	if f.srv.Patch == nil {
		return nil, &nonRetriableError{errors.New("fake for method Patch not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/firewallPolicies/(?P<firewallPolicyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/signatureOverrides/default`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armnetwork.SignaturesOverrides](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	firewallPolicyNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("firewallPolicyName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := f.srv.Patch(req.Context(), resourceGroupNameParam, firewallPolicyNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SignaturesOverrides, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (f *FirewallPolicyIdpsSignaturesOverridesServerTransport) dispatchPut(req *http.Request) (*http.Response, error) {
	if f.srv.Put == nil {
		return nil, &nonRetriableError{errors.New("fake for method Put not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/firewallPolicies/(?P<firewallPolicyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/signatureOverrides/default`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armnetwork.SignaturesOverrides](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	firewallPolicyNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("firewallPolicyName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := f.srv.Put(req.Context(), resourceGroupNameParam, firewallPolicyNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SignaturesOverrides, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
