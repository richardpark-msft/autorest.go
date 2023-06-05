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

// ApplicationGatewaysServer is a fake server for instances of the armnetwork.ApplicationGatewaysClient type.
type ApplicationGatewaysServer struct {
	// BeginBackendHealth is the fake for method ApplicationGatewaysClient.BeginBackendHealth
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginBackendHealth func(ctx context.Context, resourceGroupName string, applicationGatewayName string, options *armnetwork.ApplicationGatewaysClientBeginBackendHealthOptions) (resp azfake.PollerResponder[armnetwork.ApplicationGatewaysClientBackendHealthResponse], errResp azfake.ErrorResponder)

	// BeginBackendHealthOnDemand is the fake for method ApplicationGatewaysClient.BeginBackendHealthOnDemand
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginBackendHealthOnDemand func(ctx context.Context, resourceGroupName string, applicationGatewayName string, probeRequest armnetwork.ApplicationGatewayOnDemandProbe, options *armnetwork.ApplicationGatewaysClientBeginBackendHealthOnDemandOptions) (resp azfake.PollerResponder[armnetwork.ApplicationGatewaysClientBackendHealthOnDemandResponse], errResp azfake.ErrorResponder)

	// BeginCreateOrUpdate is the fake for method ApplicationGatewaysClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, applicationGatewayName string, parameters armnetwork.ApplicationGateway, options *armnetwork.ApplicationGatewaysClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armnetwork.ApplicationGatewaysClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method ApplicationGatewaysClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, applicationGatewayName string, options *armnetwork.ApplicationGatewaysClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetwork.ApplicationGatewaysClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ApplicationGatewaysClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, applicationGatewayName string, options *armnetwork.ApplicationGatewaysClientGetOptions) (resp azfake.Responder[armnetwork.ApplicationGatewaysClientGetResponse], errResp azfake.ErrorResponder)

	// GetSSLPredefinedPolicy is the fake for method ApplicationGatewaysClient.GetSSLPredefinedPolicy
	// HTTP status codes to indicate success: http.StatusOK
	GetSSLPredefinedPolicy func(ctx context.Context, predefinedPolicyName string, options *armnetwork.ApplicationGatewaysClientGetSSLPredefinedPolicyOptions) (resp azfake.Responder[armnetwork.ApplicationGatewaysClientGetSSLPredefinedPolicyResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method ApplicationGatewaysClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, options *armnetwork.ApplicationGatewaysClientListOptions) (resp azfake.PagerResponder[armnetwork.ApplicationGatewaysClientListResponse])

	// NewListAllPager is the fake for method ApplicationGatewaysClient.NewListAllPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListAllPager func(options *armnetwork.ApplicationGatewaysClientListAllOptions) (resp azfake.PagerResponder[armnetwork.ApplicationGatewaysClientListAllResponse])

	// ListAvailableRequestHeaders is the fake for method ApplicationGatewaysClient.ListAvailableRequestHeaders
	// HTTP status codes to indicate success: http.StatusOK
	ListAvailableRequestHeaders func(ctx context.Context, options *armnetwork.ApplicationGatewaysClientListAvailableRequestHeadersOptions) (resp azfake.Responder[armnetwork.ApplicationGatewaysClientListAvailableRequestHeadersResponse], errResp azfake.ErrorResponder)

	// ListAvailableResponseHeaders is the fake for method ApplicationGatewaysClient.ListAvailableResponseHeaders
	// HTTP status codes to indicate success: http.StatusOK
	ListAvailableResponseHeaders func(ctx context.Context, options *armnetwork.ApplicationGatewaysClientListAvailableResponseHeadersOptions) (resp azfake.Responder[armnetwork.ApplicationGatewaysClientListAvailableResponseHeadersResponse], errResp azfake.ErrorResponder)

	// ListAvailableSSLOptions is the fake for method ApplicationGatewaysClient.ListAvailableSSLOptions
	// HTTP status codes to indicate success: http.StatusOK
	ListAvailableSSLOptions func(ctx context.Context, options *armnetwork.ApplicationGatewaysClientListAvailableSSLOptionsOptions) (resp azfake.Responder[armnetwork.ApplicationGatewaysClientListAvailableSSLOptionsResponse], errResp azfake.ErrorResponder)

	// NewListAvailableSSLPredefinedPoliciesPager is the fake for method ApplicationGatewaysClient.NewListAvailableSSLPredefinedPoliciesPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListAvailableSSLPredefinedPoliciesPager func(options *armnetwork.ApplicationGatewaysClientListAvailableSSLPredefinedPoliciesOptions) (resp azfake.PagerResponder[armnetwork.ApplicationGatewaysClientListAvailableSSLPredefinedPoliciesResponse])

	// ListAvailableServerVariables is the fake for method ApplicationGatewaysClient.ListAvailableServerVariables
	// HTTP status codes to indicate success: http.StatusOK
	ListAvailableServerVariables func(ctx context.Context, options *armnetwork.ApplicationGatewaysClientListAvailableServerVariablesOptions) (resp azfake.Responder[armnetwork.ApplicationGatewaysClientListAvailableServerVariablesResponse], errResp azfake.ErrorResponder)

	// ListAvailableWafRuleSets is the fake for method ApplicationGatewaysClient.ListAvailableWafRuleSets
	// HTTP status codes to indicate success: http.StatusOK
	ListAvailableWafRuleSets func(ctx context.Context, options *armnetwork.ApplicationGatewaysClientListAvailableWafRuleSetsOptions) (resp azfake.Responder[armnetwork.ApplicationGatewaysClientListAvailableWafRuleSetsResponse], errResp azfake.ErrorResponder)

	// BeginStart is the fake for method ApplicationGatewaysClient.BeginStart
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginStart func(ctx context.Context, resourceGroupName string, applicationGatewayName string, options *armnetwork.ApplicationGatewaysClientBeginStartOptions) (resp azfake.PollerResponder[armnetwork.ApplicationGatewaysClientStartResponse], errResp azfake.ErrorResponder)

	// BeginStop is the fake for method ApplicationGatewaysClient.BeginStop
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginStop func(ctx context.Context, resourceGroupName string, applicationGatewayName string, options *armnetwork.ApplicationGatewaysClientBeginStopOptions) (resp azfake.PollerResponder[armnetwork.ApplicationGatewaysClientStopResponse], errResp azfake.ErrorResponder)

	// UpdateTags is the fake for method ApplicationGatewaysClient.UpdateTags
	// HTTP status codes to indicate success: http.StatusOK
	UpdateTags func(ctx context.Context, resourceGroupName string, applicationGatewayName string, parameters armnetwork.TagsObject, options *armnetwork.ApplicationGatewaysClientUpdateTagsOptions) (resp azfake.Responder[armnetwork.ApplicationGatewaysClientUpdateTagsResponse], errResp azfake.ErrorResponder)
}

// NewApplicationGatewaysServerTransport creates a new instance of ApplicationGatewaysServerTransport with the provided implementation.
// The returned ApplicationGatewaysServerTransport instance is connected to an instance of armnetwork.ApplicationGatewaysClient by way of the
// undefined.Transporter field.
func NewApplicationGatewaysServerTransport(srv *ApplicationGatewaysServer) *ApplicationGatewaysServerTransport {
	return &ApplicationGatewaysServerTransport{srv: srv}
}

// ApplicationGatewaysServerTransport connects instances of armnetwork.ApplicationGatewaysClient to instances of ApplicationGatewaysServer.
// Don't use this type directly, use NewApplicationGatewaysServerTransport instead.
type ApplicationGatewaysServerTransport struct {
	srv                                        *ApplicationGatewaysServer
	beginBackendHealth                         *azfake.PollerResponder[armnetwork.ApplicationGatewaysClientBackendHealthResponse]
	beginBackendHealthOnDemand                 *azfake.PollerResponder[armnetwork.ApplicationGatewaysClientBackendHealthOnDemandResponse]
	beginCreateOrUpdate                        *azfake.PollerResponder[armnetwork.ApplicationGatewaysClientCreateOrUpdateResponse]
	beginDelete                                *azfake.PollerResponder[armnetwork.ApplicationGatewaysClientDeleteResponse]
	newListPager                               *azfake.PagerResponder[armnetwork.ApplicationGatewaysClientListResponse]
	newListAllPager                            *azfake.PagerResponder[armnetwork.ApplicationGatewaysClientListAllResponse]
	newListAvailableSSLPredefinedPoliciesPager *azfake.PagerResponder[armnetwork.ApplicationGatewaysClientListAvailableSSLPredefinedPoliciesResponse]
	beginStart                                 *azfake.PollerResponder[armnetwork.ApplicationGatewaysClientStartResponse]
	beginStop                                  *azfake.PollerResponder[armnetwork.ApplicationGatewaysClientStopResponse]
}

// Do implements the policy.Transporter interface for ApplicationGatewaysServerTransport.
func (a *ApplicationGatewaysServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ApplicationGatewaysClient.BeginBackendHealth":
		resp, err = a.dispatchBeginBackendHealth(req)
	case "ApplicationGatewaysClient.BeginBackendHealthOnDemand":
		resp, err = a.dispatchBeginBackendHealthOnDemand(req)
	case "ApplicationGatewaysClient.BeginCreateOrUpdate":
		resp, err = a.dispatchBeginCreateOrUpdate(req)
	case "ApplicationGatewaysClient.BeginDelete":
		resp, err = a.dispatchBeginDelete(req)
	case "ApplicationGatewaysClient.Get":
		resp, err = a.dispatchGet(req)
	case "ApplicationGatewaysClient.GetSSLPredefinedPolicy":
		resp, err = a.dispatchGetSSLPredefinedPolicy(req)
	case "ApplicationGatewaysClient.NewListPager":
		resp, err = a.dispatchNewListPager(req)
	case "ApplicationGatewaysClient.NewListAllPager":
		resp, err = a.dispatchNewListAllPager(req)
	case "ApplicationGatewaysClient.ListAvailableRequestHeaders":
		resp, err = a.dispatchListAvailableRequestHeaders(req)
	case "ApplicationGatewaysClient.ListAvailableResponseHeaders":
		resp, err = a.dispatchListAvailableResponseHeaders(req)
	case "ApplicationGatewaysClient.ListAvailableSSLOptions":
		resp, err = a.dispatchListAvailableSSLOptions(req)
	case "ApplicationGatewaysClient.NewListAvailableSSLPredefinedPoliciesPager":
		resp, err = a.dispatchNewListAvailableSSLPredefinedPoliciesPager(req)
	case "ApplicationGatewaysClient.ListAvailableServerVariables":
		resp, err = a.dispatchListAvailableServerVariables(req)
	case "ApplicationGatewaysClient.ListAvailableWafRuleSets":
		resp, err = a.dispatchListAvailableWafRuleSets(req)
	case "ApplicationGatewaysClient.BeginStart":
		resp, err = a.dispatchBeginStart(req)
	case "ApplicationGatewaysClient.BeginStop":
		resp, err = a.dispatchBeginStop(req)
	case "ApplicationGatewaysClient.UpdateTags":
		resp, err = a.dispatchUpdateTags(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *ApplicationGatewaysServerTransport) dispatchBeginBackendHealth(req *http.Request) (*http.Response, error) {
	if a.srv.BeginBackendHealth == nil {
		return nil, &nonRetriableError{errors.New("method BeginBackendHealth not implemented")}
	}
	if a.beginBackendHealth == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/applicationGateways/(?P<applicationGatewayName>[a-zA-Z0-9-_]+)/backendhealth"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		expandParam := getOptional(qp.Get("$expand"))
		var options *armnetwork.ApplicationGatewaysClientBeginBackendHealthOptions
		if expandParam != nil {
			options = &armnetwork.ApplicationGatewaysClientBeginBackendHealthOptions{
				Expand: expandParam,
			}
		}
		respr, errRespr := a.srv.BeginBackendHealth(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("applicationGatewayName")], options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		a.beginBackendHealth = &respr
	}

	resp, err := server.PollerResponderNext(a.beginBackendHealth, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(a.beginBackendHealth) {
		a.beginBackendHealth = nil
	}

	return resp, nil
}

func (a *ApplicationGatewaysServerTransport) dispatchBeginBackendHealthOnDemand(req *http.Request) (*http.Response, error) {
	if a.srv.BeginBackendHealthOnDemand == nil {
		return nil, &nonRetriableError{errors.New("method BeginBackendHealthOnDemand not implemented")}
	}
	if a.beginBackendHealthOnDemand == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/applicationGateways/(?P<applicationGatewayName>[a-zA-Z0-9-_]+)/getBackendHealthOnDemand"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		body, err := server.UnmarshalRequestAsJSON[armnetwork.ApplicationGatewayOnDemandProbe](req)
		if err != nil {
			return nil, err
		}
		expandParam := getOptional(qp.Get("$expand"))
		var options *armnetwork.ApplicationGatewaysClientBeginBackendHealthOnDemandOptions
		if expandParam != nil {
			options = &armnetwork.ApplicationGatewaysClientBeginBackendHealthOnDemandOptions{
				Expand: expandParam,
			}
		}
		respr, errRespr := a.srv.BeginBackendHealthOnDemand(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("applicationGatewayName")], body, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		a.beginBackendHealthOnDemand = &respr
	}

	resp, err := server.PollerResponderNext(a.beginBackendHealthOnDemand, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(a.beginBackendHealthOnDemand) {
		a.beginBackendHealthOnDemand = nil
	}

	return resp, nil
}

func (a *ApplicationGatewaysServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if a.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("method BeginCreateOrUpdate not implemented")}
	}
	if a.beginCreateOrUpdate == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/applicationGateways/(?P<applicationGatewayName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.ApplicationGateway](req)
		if err != nil {
			return nil, err
		}
		respr, errRespr := a.srv.BeginCreateOrUpdate(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("applicationGatewayName")], body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		a.beginCreateOrUpdate = &respr
	}

	resp, err := server.PollerResponderNext(a.beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(a.beginCreateOrUpdate) {
		a.beginCreateOrUpdate = nil
	}

	return resp, nil
}

func (a *ApplicationGatewaysServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if a.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("method BeginDelete not implemented")}
	}
	if a.beginDelete == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/applicationGateways/(?P<applicationGatewayName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		respr, errRespr := a.srv.BeginDelete(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("applicationGatewayName")], nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		a.beginDelete = &respr
	}

	resp, err := server.PollerResponderNext(a.beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(a.beginDelete) {
		a.beginDelete = nil
	}

	return resp, nil
}

func (a *ApplicationGatewaysServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if a.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("method Get not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/applicationGateways/(?P<applicationGatewayName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := a.srv.Get(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("applicationGatewayName")], nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ApplicationGateway, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *ApplicationGatewaysServerTransport) dispatchGetSSLPredefinedPolicy(req *http.Request) (*http.Response, error) {
	if a.srv.GetSSLPredefinedPolicy == nil {
		return nil, &nonRetriableError{errors.New("method GetSSLPredefinedPolicy not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/applicationGatewayAvailableSslOptions/default/predefinedPolicies/(?P<predefinedPolicyName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := a.srv.GetSSLPredefinedPolicy(req.Context(), matches[regex.SubexpIndex("predefinedPolicyName")], nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ApplicationGatewaySSLPredefinedPolicy, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *ApplicationGatewaysServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("method NewListPager not implemented")}
	}
	if a.newListPager == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/applicationGateways"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := a.srv.NewListPager(matches[regex.SubexpIndex("resourceGroupName")], nil)
		a.newListPager = &resp
		server.PagerResponderInjectNextLinks(a.newListPager, req, func(page *armnetwork.ApplicationGatewaysClientListResponse, createLink func() string) {
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

func (a *ApplicationGatewaysServerTransport) dispatchNewListAllPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListAllPager == nil {
		return nil, &nonRetriableError{errors.New("method NewListAllPager not implemented")}
	}
	if a.newListAllPager == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/applicationGateways"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := a.srv.NewListAllPager(nil)
		a.newListAllPager = &resp
		server.PagerResponderInjectNextLinks(a.newListAllPager, req, func(page *armnetwork.ApplicationGatewaysClientListAllResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(a.newListAllPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(a.newListAllPager) {
		a.newListAllPager = nil
	}
	return resp, nil
}

func (a *ApplicationGatewaysServerTransport) dispatchListAvailableRequestHeaders(req *http.Request) (*http.Response, error) {
	if a.srv.ListAvailableRequestHeaders == nil {
		return nil, &nonRetriableError{errors.New("method ListAvailableRequestHeaders not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/applicationGatewayAvailableRequestHeaders"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := a.srv.ListAvailableRequestHeaders(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).StringArray, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *ApplicationGatewaysServerTransport) dispatchListAvailableResponseHeaders(req *http.Request) (*http.Response, error) {
	if a.srv.ListAvailableResponseHeaders == nil {
		return nil, &nonRetriableError{errors.New("method ListAvailableResponseHeaders not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/applicationGatewayAvailableResponseHeaders"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := a.srv.ListAvailableResponseHeaders(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).StringArray, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *ApplicationGatewaysServerTransport) dispatchListAvailableSSLOptions(req *http.Request) (*http.Response, error) {
	if a.srv.ListAvailableSSLOptions == nil {
		return nil, &nonRetriableError{errors.New("method ListAvailableSSLOptions not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/applicationGatewayAvailableSslOptions/default"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := a.srv.ListAvailableSSLOptions(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ApplicationGatewayAvailableSSLOptions, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *ApplicationGatewaysServerTransport) dispatchNewListAvailableSSLPredefinedPoliciesPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListAvailableSSLPredefinedPoliciesPager == nil {
		return nil, &nonRetriableError{errors.New("method NewListAvailableSSLPredefinedPoliciesPager not implemented")}
	}
	if a.newListAvailableSSLPredefinedPoliciesPager == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/applicationGatewayAvailableSslOptions/default/predefinedPolicies"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := a.srv.NewListAvailableSSLPredefinedPoliciesPager(nil)
		a.newListAvailableSSLPredefinedPoliciesPager = &resp
		server.PagerResponderInjectNextLinks(a.newListAvailableSSLPredefinedPoliciesPager, req, func(page *armnetwork.ApplicationGatewaysClientListAvailableSSLPredefinedPoliciesResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(a.newListAvailableSSLPredefinedPoliciesPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(a.newListAvailableSSLPredefinedPoliciesPager) {
		a.newListAvailableSSLPredefinedPoliciesPager = nil
	}
	return resp, nil
}

func (a *ApplicationGatewaysServerTransport) dispatchListAvailableServerVariables(req *http.Request) (*http.Response, error) {
	if a.srv.ListAvailableServerVariables == nil {
		return nil, &nonRetriableError{errors.New("method ListAvailableServerVariables not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/applicationGatewayAvailableServerVariables"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := a.srv.ListAvailableServerVariables(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).StringArray, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *ApplicationGatewaysServerTransport) dispatchListAvailableWafRuleSets(req *http.Request) (*http.Response, error) {
	if a.srv.ListAvailableWafRuleSets == nil {
		return nil, &nonRetriableError{errors.New("method ListAvailableWafRuleSets not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/applicationGatewayAvailableWafRuleSets"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := a.srv.ListAvailableWafRuleSets(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ApplicationGatewayAvailableWafRuleSetsResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *ApplicationGatewaysServerTransport) dispatchBeginStart(req *http.Request) (*http.Response, error) {
	if a.srv.BeginStart == nil {
		return nil, &nonRetriableError{errors.New("method BeginStart not implemented")}
	}
	if a.beginStart == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/applicationGateways/(?P<applicationGatewayName>[a-zA-Z0-9-_]+)/start"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		respr, errRespr := a.srv.BeginStart(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("applicationGatewayName")], nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		a.beginStart = &respr
	}

	resp, err := server.PollerResponderNext(a.beginStart, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(a.beginStart) {
		a.beginStart = nil
	}

	return resp, nil
}

func (a *ApplicationGatewaysServerTransport) dispatchBeginStop(req *http.Request) (*http.Response, error) {
	if a.srv.BeginStop == nil {
		return nil, &nonRetriableError{errors.New("method BeginStop not implemented")}
	}
	if a.beginStop == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/applicationGateways/(?P<applicationGatewayName>[a-zA-Z0-9-_]+)/stop"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		respr, errRespr := a.srv.BeginStop(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("applicationGatewayName")], nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		a.beginStop = &respr
	}

	resp, err := server.PollerResponderNext(a.beginStop, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(a.beginStop) {
		a.beginStop = nil
	}

	return resp, nil
}

func (a *ApplicationGatewaysServerTransport) dispatchUpdateTags(req *http.Request) (*http.Response, error) {
	if a.srv.UpdateTags == nil {
		return nil, &nonRetriableError{errors.New("method UpdateTags not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/applicationGateways/(?P<applicationGatewayName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armnetwork.TagsObject](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.UpdateTags(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("applicationGatewayName")], body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ApplicationGateway, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
