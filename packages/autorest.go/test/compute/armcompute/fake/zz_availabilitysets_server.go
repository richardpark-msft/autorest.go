//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

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

// AvailabilitySetsServer is a fake server for instances of the armcompute.AvailabilitySetsClient type.
type AvailabilitySetsServer struct {
	// CreateOrUpdate is the fake for method AvailabilitySetsClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, availabilitySetName string, parameters armcompute.AvailabilitySet, options *armcompute.AvailabilitySetsClientCreateOrUpdateOptions) (resp azfake.Responder[armcompute.AvailabilitySetsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method AvailabilitySetsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, availabilitySetName string, options *armcompute.AvailabilitySetsClientDeleteOptions) (resp azfake.Responder[armcompute.AvailabilitySetsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method AvailabilitySetsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, availabilitySetName string, options *armcompute.AvailabilitySetsClientGetOptions) (resp azfake.Responder[armcompute.AvailabilitySetsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method AvailabilitySetsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, options *armcompute.AvailabilitySetsClientListOptions) (resp azfake.PagerResponder[armcompute.AvailabilitySetsClientListResponse])

	// NewListAvailableSizesPager is the fake for method AvailabilitySetsClient.NewListAvailableSizesPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListAvailableSizesPager func(resourceGroupName string, availabilitySetName string, options *armcompute.AvailabilitySetsClientListAvailableSizesOptions) (resp azfake.PagerResponder[armcompute.AvailabilitySetsClientListAvailableSizesResponse])

	// NewListBySubscriptionPager is the fake for method AvailabilitySetsClient.NewListBySubscriptionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySubscriptionPager func(options *armcompute.AvailabilitySetsClientListBySubscriptionOptions) (resp azfake.PagerResponder[armcompute.AvailabilitySetsClientListBySubscriptionResponse])

	// Update is the fake for method AvailabilitySetsClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, availabilitySetName string, parameters armcompute.AvailabilitySetUpdate, options *armcompute.AvailabilitySetsClientUpdateOptions) (resp azfake.Responder[armcompute.AvailabilitySetsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewAvailabilitySetsServerTransport creates a new instance of AvailabilitySetsServerTransport with the provided implementation.
// The returned AvailabilitySetsServerTransport instance is connected to an instance of armcompute.AvailabilitySetsClient by way of the
// undefined.Transporter field.
func NewAvailabilitySetsServerTransport(srv *AvailabilitySetsServer) *AvailabilitySetsServerTransport {
	return &AvailabilitySetsServerTransport{srv: srv}
}

// AvailabilitySetsServerTransport connects instances of armcompute.AvailabilitySetsClient to instances of AvailabilitySetsServer.
// Don't use this type directly, use NewAvailabilitySetsServerTransport instead.
type AvailabilitySetsServerTransport struct {
	srv                        *AvailabilitySetsServer
	newListPager               *azfake.PagerResponder[armcompute.AvailabilitySetsClientListResponse]
	newListAvailableSizesPager *azfake.PagerResponder[armcompute.AvailabilitySetsClientListAvailableSizesResponse]
	newListBySubscriptionPager *azfake.PagerResponder[armcompute.AvailabilitySetsClientListBySubscriptionResponse]
}

// Do implements the policy.Transporter interface for AvailabilitySetsServerTransport.
func (a *AvailabilitySetsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "AvailabilitySetsClient.CreateOrUpdate":
		resp, err = a.dispatchCreateOrUpdate(req)
	case "AvailabilitySetsClient.Delete":
		resp, err = a.dispatchDelete(req)
	case "AvailabilitySetsClient.Get":
		resp, err = a.dispatchGet(req)
	case "AvailabilitySetsClient.NewListPager":
		resp, err = a.dispatchNewListPager(req)
	case "AvailabilitySetsClient.NewListAvailableSizesPager":
		resp, err = a.dispatchNewListAvailableSizesPager(req)
	case "AvailabilitySetsClient.NewListBySubscriptionPager":
		resp, err = a.dispatchNewListBySubscriptionPager(req)
	case "AvailabilitySetsClient.Update":
		resp, err = a.dispatchUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *AvailabilitySetsServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if a.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("method CreateOrUpdate not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/availabilitySets/(?P<availabilitySetName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armcompute.AvailabilitySet](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.CreateOrUpdate(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("availabilitySetName")], body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AvailabilitySet, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AvailabilitySetsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if a.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("method Delete not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/availabilitySets/(?P<availabilitySetName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := a.srv.Delete(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("availabilitySetName")], nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AvailabilitySetsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if a.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("method Get not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/availabilitySets/(?P<availabilitySetName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := a.srv.Get(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("availabilitySetName")], nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AvailabilitySet, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AvailabilitySetsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("method NewListPager not implemented")}
	}
	if a.newListPager == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/availabilitySets"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := a.srv.NewListPager(matches[regex.SubexpIndex("resourceGroupName")], nil)
		a.newListPager = &resp
		server.PagerResponderInjectNextLinks(a.newListPager, req, func(page *armcompute.AvailabilitySetsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(a.newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(a.newListPager) {
		a.newListPager = nil
	}
	return resp, nil
}

func (a *AvailabilitySetsServerTransport) dispatchNewListAvailableSizesPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListAvailableSizesPager == nil {
		return nil, &nonRetriableError{errors.New("method NewListAvailableSizesPager not implemented")}
	}
	if a.newListAvailableSizesPager == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/availabilitySets/(?P<availabilitySetName>[a-zA-Z0-9-_]+)/vmSizes"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := a.srv.NewListAvailableSizesPager(matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("availabilitySetName")], nil)
		a.newListAvailableSizesPager = &resp
	}
	resp, err := server.PagerResponderNext(a.newListAvailableSizesPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(a.newListAvailableSizesPager) {
		a.newListAvailableSizesPager = nil
	}
	return resp, nil
}

func (a *AvailabilitySetsServerTransport) dispatchNewListBySubscriptionPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListBySubscriptionPager == nil {
		return nil, &nonRetriableError{errors.New("method NewListBySubscriptionPager not implemented")}
	}
	if a.newListBySubscriptionPager == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/availabilitySets"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		expandParam := getOptional(qp.Get("$expand"))
		var options *armcompute.AvailabilitySetsClientListBySubscriptionOptions
		if expandParam != nil {
			options = &armcompute.AvailabilitySetsClientListBySubscriptionOptions{
				Expand: expandParam,
			}
		}
		resp := a.srv.NewListBySubscriptionPager(options)
		a.newListBySubscriptionPager = &resp
		server.PagerResponderInjectNextLinks(a.newListBySubscriptionPager, req, func(page *armcompute.AvailabilitySetsClientListBySubscriptionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(a.newListBySubscriptionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(a.newListBySubscriptionPager) {
		a.newListBySubscriptionPager = nil
	}
	return resp, nil
}

func (a *AvailabilitySetsServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if a.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("method Update not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/availabilitySets/(?P<availabilitySetName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armcompute.AvailabilitySetUpdate](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Update(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("availabilitySetName")], body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AvailabilitySet, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
