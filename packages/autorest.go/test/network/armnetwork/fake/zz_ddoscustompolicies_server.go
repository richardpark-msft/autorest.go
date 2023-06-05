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

// DdosCustomPoliciesServer is a fake server for instances of the armnetwork.DdosCustomPoliciesClient type.
type DdosCustomPoliciesServer struct {
	// BeginCreateOrUpdate is the fake for method DdosCustomPoliciesClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, ddosCustomPolicyName string, parameters armnetwork.DdosCustomPolicy, options *armnetwork.DdosCustomPoliciesClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armnetwork.DdosCustomPoliciesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method DdosCustomPoliciesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, ddosCustomPolicyName string, options *armnetwork.DdosCustomPoliciesClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetwork.DdosCustomPoliciesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method DdosCustomPoliciesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, ddosCustomPolicyName string, options *armnetwork.DdosCustomPoliciesClientGetOptions) (resp azfake.Responder[armnetwork.DdosCustomPoliciesClientGetResponse], errResp azfake.ErrorResponder)

	// UpdateTags is the fake for method DdosCustomPoliciesClient.UpdateTags
	// HTTP status codes to indicate success: http.StatusOK
	UpdateTags func(ctx context.Context, resourceGroupName string, ddosCustomPolicyName string, parameters armnetwork.TagsObject, options *armnetwork.DdosCustomPoliciesClientUpdateTagsOptions) (resp azfake.Responder[armnetwork.DdosCustomPoliciesClientUpdateTagsResponse], errResp azfake.ErrorResponder)
}

// NewDdosCustomPoliciesServerTransport creates a new instance of DdosCustomPoliciesServerTransport with the provided implementation.
// The returned DdosCustomPoliciesServerTransport instance is connected to an instance of armnetwork.DdosCustomPoliciesClient by way of the
// undefined.Transporter field.
func NewDdosCustomPoliciesServerTransport(srv *DdosCustomPoliciesServer) *DdosCustomPoliciesServerTransport {
	return &DdosCustomPoliciesServerTransport{srv: srv}
}

// DdosCustomPoliciesServerTransport connects instances of armnetwork.DdosCustomPoliciesClient to instances of DdosCustomPoliciesServer.
// Don't use this type directly, use NewDdosCustomPoliciesServerTransport instead.
type DdosCustomPoliciesServerTransport struct {
	srv                 *DdosCustomPoliciesServer
	beginCreateOrUpdate *azfake.PollerResponder[armnetwork.DdosCustomPoliciesClientCreateOrUpdateResponse]
	beginDelete         *azfake.PollerResponder[armnetwork.DdosCustomPoliciesClientDeleteResponse]
}

// Do implements the policy.Transporter interface for DdosCustomPoliciesServerTransport.
func (d *DdosCustomPoliciesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "DdosCustomPoliciesClient.BeginCreateOrUpdate":
		resp, err = d.dispatchBeginCreateOrUpdate(req)
	case "DdosCustomPoliciesClient.BeginDelete":
		resp, err = d.dispatchBeginDelete(req)
	case "DdosCustomPoliciesClient.Get":
		resp, err = d.dispatchGet(req)
	case "DdosCustomPoliciesClient.UpdateTags":
		resp, err = d.dispatchUpdateTags(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (d *DdosCustomPoliciesServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if d.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("method BeginCreateOrUpdate not implemented")}
	}
	if d.beginCreateOrUpdate == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/ddosCustomPolicies/(?P<ddosCustomPolicyName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.DdosCustomPolicy](req)
		if err != nil {
			return nil, err
		}
		respr, errRespr := d.srv.BeginCreateOrUpdate(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("ddosCustomPolicyName")], body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		d.beginCreateOrUpdate = &respr
	}

	resp, err := server.PollerResponderNext(d.beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(d.beginCreateOrUpdate) {
		d.beginCreateOrUpdate = nil
	}

	return resp, nil
}

func (d *DdosCustomPoliciesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if d.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("method BeginDelete not implemented")}
	}
	if d.beginDelete == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/ddosCustomPolicies/(?P<ddosCustomPolicyName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		respr, errRespr := d.srv.BeginDelete(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("ddosCustomPolicyName")], nil)
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

func (d *DdosCustomPoliciesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if d.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("method Get not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/ddosCustomPolicies/(?P<ddosCustomPolicyName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := d.srv.Get(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("ddosCustomPolicyName")], nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DdosCustomPolicy, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DdosCustomPoliciesServerTransport) dispatchUpdateTags(req *http.Request) (*http.Response, error) {
	if d.srv.UpdateTags == nil {
		return nil, &nonRetriableError{errors.New("method UpdateTags not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/ddosCustomPolicies/(?P<ddosCustomPolicyName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armnetwork.TagsObject](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.UpdateTags(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("ddosCustomPolicyName")], body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DdosCustomPolicy, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
