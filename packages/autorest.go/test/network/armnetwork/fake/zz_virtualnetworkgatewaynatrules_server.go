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

// VirtualNetworkGatewayNatRulesServer is a fake server for instances of the armnetwork.VirtualNetworkGatewayNatRulesClient type.
type VirtualNetworkGatewayNatRulesServer struct {
	// BeginCreateOrUpdate is the fake for method VirtualNetworkGatewayNatRulesClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, virtualNetworkGatewayName string, natRuleName string, natRuleParameters armnetwork.VirtualNetworkGatewayNatRule, options *armnetwork.VirtualNetworkGatewayNatRulesClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armnetwork.VirtualNetworkGatewayNatRulesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method VirtualNetworkGatewayNatRulesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, virtualNetworkGatewayName string, natRuleName string, options *armnetwork.VirtualNetworkGatewayNatRulesClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetwork.VirtualNetworkGatewayNatRulesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method VirtualNetworkGatewayNatRulesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, virtualNetworkGatewayName string, natRuleName string, options *armnetwork.VirtualNetworkGatewayNatRulesClientGetOptions) (resp azfake.Responder[armnetwork.VirtualNetworkGatewayNatRulesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByVirtualNetworkGatewayPager is the fake for method VirtualNetworkGatewayNatRulesClient.NewListByVirtualNetworkGatewayPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByVirtualNetworkGatewayPager func(resourceGroupName string, virtualNetworkGatewayName string, options *armnetwork.VirtualNetworkGatewayNatRulesClientListByVirtualNetworkGatewayOptions) (resp azfake.PagerResponder[armnetwork.VirtualNetworkGatewayNatRulesClientListByVirtualNetworkGatewayResponse])
}

// NewVirtualNetworkGatewayNatRulesServerTransport creates a new instance of VirtualNetworkGatewayNatRulesServerTransport with the provided implementation.
// The returned VirtualNetworkGatewayNatRulesServerTransport instance is connected to an instance of armnetwork.VirtualNetworkGatewayNatRulesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewVirtualNetworkGatewayNatRulesServerTransport(srv *VirtualNetworkGatewayNatRulesServer) *VirtualNetworkGatewayNatRulesServerTransport {
	return &VirtualNetworkGatewayNatRulesServerTransport{
		srv:                                 srv,
		beginCreateOrUpdate:                 newTracker[azfake.PollerResponder[armnetwork.VirtualNetworkGatewayNatRulesClientCreateOrUpdateResponse]](),
		beginDelete:                         newTracker[azfake.PollerResponder[armnetwork.VirtualNetworkGatewayNatRulesClientDeleteResponse]](),
		newListByVirtualNetworkGatewayPager: newTracker[azfake.PagerResponder[armnetwork.VirtualNetworkGatewayNatRulesClientListByVirtualNetworkGatewayResponse]](),
	}
}

// VirtualNetworkGatewayNatRulesServerTransport connects instances of armnetwork.VirtualNetworkGatewayNatRulesClient to instances of VirtualNetworkGatewayNatRulesServer.
// Don't use this type directly, use NewVirtualNetworkGatewayNatRulesServerTransport instead.
type VirtualNetworkGatewayNatRulesServerTransport struct {
	srv                                 *VirtualNetworkGatewayNatRulesServer
	beginCreateOrUpdate                 *tracker[azfake.PollerResponder[armnetwork.VirtualNetworkGatewayNatRulesClientCreateOrUpdateResponse]]
	beginDelete                         *tracker[azfake.PollerResponder[armnetwork.VirtualNetworkGatewayNatRulesClientDeleteResponse]]
	newListByVirtualNetworkGatewayPager *tracker[azfake.PagerResponder[armnetwork.VirtualNetworkGatewayNatRulesClientListByVirtualNetworkGatewayResponse]]
}

// Do implements the policy.Transporter interface for VirtualNetworkGatewayNatRulesServerTransport.
func (v *VirtualNetworkGatewayNatRulesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return v.dispatchToMethodFake(req, method)
}

func (v *VirtualNetworkGatewayNatRulesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch method {
	case "VirtualNetworkGatewayNatRulesClient.BeginCreateOrUpdate":
		resp, err = v.dispatchBeginCreateOrUpdate(req)
	case "VirtualNetworkGatewayNatRulesClient.BeginDelete":
		resp, err = v.dispatchBeginDelete(req)
	case "VirtualNetworkGatewayNatRulesClient.Get":
		resp, err = v.dispatchGet(req)
	case "VirtualNetworkGatewayNatRulesClient.NewListByVirtualNetworkGatewayPager":
		resp, err = v.dispatchNewListByVirtualNetworkGatewayPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	return resp, err
}

func (v *VirtualNetworkGatewayNatRulesServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if v.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := v.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/virtualNetworkGateways/(?P<virtualNetworkGatewayName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/natRules/(?P<natRuleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.VirtualNetworkGatewayNatRule](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		virtualNetworkGatewayNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("virtualNetworkGatewayName")])
		if err != nil {
			return nil, err
		}
		natRuleNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("natRuleName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, virtualNetworkGatewayNameParam, natRuleNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		v.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		v.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		v.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (v *VirtualNetworkGatewayNatRulesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if v.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := v.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/virtualNetworkGateways/(?P<virtualNetworkGatewayName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/natRules/(?P<natRuleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		virtualNetworkGatewayNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("virtualNetworkGatewayName")])
		if err != nil {
			return nil, err
		}
		natRuleNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("natRuleName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginDelete(req.Context(), resourceGroupNameParam, virtualNetworkGatewayNameParam, natRuleNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		v.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		v.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		v.beginDelete.remove(req)
	}

	return resp, nil
}

func (v *VirtualNetworkGatewayNatRulesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if v.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/virtualNetworkGateways/(?P<virtualNetworkGatewayName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/natRules/(?P<natRuleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	virtualNetworkGatewayNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("virtualNetworkGatewayName")])
	if err != nil {
		return nil, err
	}
	natRuleNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("natRuleName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := v.srv.Get(req.Context(), resourceGroupNameParam, virtualNetworkGatewayNameParam, natRuleNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).VirtualNetworkGatewayNatRule, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VirtualNetworkGatewayNatRulesServerTransport) dispatchNewListByVirtualNetworkGatewayPager(req *http.Request) (*http.Response, error) {
	if v.srv.NewListByVirtualNetworkGatewayPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByVirtualNetworkGatewayPager not implemented")}
	}
	newListByVirtualNetworkGatewayPager := v.newListByVirtualNetworkGatewayPager.get(req)
	if newListByVirtualNetworkGatewayPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/virtualNetworkGateways/(?P<virtualNetworkGatewayName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/natRules`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		virtualNetworkGatewayNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("virtualNetworkGatewayName")])
		if err != nil {
			return nil, err
		}
		resp := v.srv.NewListByVirtualNetworkGatewayPager(resourceGroupNameParam, virtualNetworkGatewayNameParam, nil)
		newListByVirtualNetworkGatewayPager = &resp
		v.newListByVirtualNetworkGatewayPager.add(req, newListByVirtualNetworkGatewayPager)
		server.PagerResponderInjectNextLinks(newListByVirtualNetworkGatewayPager, req, func(page *armnetwork.VirtualNetworkGatewayNatRulesClientListByVirtualNetworkGatewayResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByVirtualNetworkGatewayPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		v.newListByVirtualNetworkGatewayPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByVirtualNetworkGatewayPager) {
		v.newListByVirtualNetworkGatewayPager.remove(req)
	}
	return resp, nil
}
