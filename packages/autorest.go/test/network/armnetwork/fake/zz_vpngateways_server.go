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
	"reflect"
	"regexp"
)

// VPNGatewaysServer is a fake server for instances of the armnetwork.VPNGatewaysClient type.
type VPNGatewaysServer struct {
	// BeginCreateOrUpdate is the fake for method VPNGatewaysClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, gatewayName string, vpnGatewayParameters armnetwork.VPNGateway, options *armnetwork.VPNGatewaysClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armnetwork.VPNGatewaysClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method VPNGatewaysClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, gatewayName string, options *armnetwork.VPNGatewaysClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetwork.VPNGatewaysClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method VPNGatewaysClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, gatewayName string, options *armnetwork.VPNGatewaysClientGetOptions) (resp azfake.Responder[armnetwork.VPNGatewaysClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method VPNGatewaysClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armnetwork.VPNGatewaysClientListOptions) (resp azfake.PagerResponder[armnetwork.VPNGatewaysClientListResponse])

	// NewListByResourceGroupPager is the fake for method VPNGatewaysClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armnetwork.VPNGatewaysClientListByResourceGroupOptions) (resp azfake.PagerResponder[armnetwork.VPNGatewaysClientListByResourceGroupResponse])

	// BeginReset is the fake for method VPNGatewaysClient.BeginReset
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginReset func(ctx context.Context, resourceGroupName string, gatewayName string, options *armnetwork.VPNGatewaysClientBeginResetOptions) (resp azfake.PollerResponder[armnetwork.VPNGatewaysClientResetResponse], errResp azfake.ErrorResponder)

	// BeginStartPacketCapture is the fake for method VPNGatewaysClient.BeginStartPacketCapture
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginStartPacketCapture func(ctx context.Context, resourceGroupName string, gatewayName string, options *armnetwork.VPNGatewaysClientBeginStartPacketCaptureOptions) (resp azfake.PollerResponder[armnetwork.VPNGatewaysClientStartPacketCaptureResponse], errResp azfake.ErrorResponder)

	// BeginStopPacketCapture is the fake for method VPNGatewaysClient.BeginStopPacketCapture
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginStopPacketCapture func(ctx context.Context, resourceGroupName string, gatewayName string, options *armnetwork.VPNGatewaysClientBeginStopPacketCaptureOptions) (resp azfake.PollerResponder[armnetwork.VPNGatewaysClientStopPacketCaptureResponse], errResp azfake.ErrorResponder)

	// BeginUpdateTags is the fake for method VPNGatewaysClient.BeginUpdateTags
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdateTags func(ctx context.Context, resourceGroupName string, gatewayName string, vpnGatewayParameters armnetwork.TagsObject, options *armnetwork.VPNGatewaysClientBeginUpdateTagsOptions) (resp azfake.PollerResponder[armnetwork.VPNGatewaysClientUpdateTagsResponse], errResp azfake.ErrorResponder)
}

// NewVPNGatewaysServerTransport creates a new instance of VPNGatewaysServerTransport with the provided implementation.
// The returned VPNGatewaysServerTransport instance is connected to an instance of armnetwork.VPNGatewaysClient by way of the
// undefined.Transporter field.
func NewVPNGatewaysServerTransport(srv *VPNGatewaysServer) *VPNGatewaysServerTransport {
	return &VPNGatewaysServerTransport{srv: srv}
}

// VPNGatewaysServerTransport connects instances of armnetwork.VPNGatewaysClient to instances of VPNGatewaysServer.
// Don't use this type directly, use NewVPNGatewaysServerTransport instead.
type VPNGatewaysServerTransport struct {
	srv                         *VPNGatewaysServer
	beginCreateOrUpdate         *azfake.PollerResponder[armnetwork.VPNGatewaysClientCreateOrUpdateResponse]
	beginDelete                 *azfake.PollerResponder[armnetwork.VPNGatewaysClientDeleteResponse]
	newListPager                *azfake.PagerResponder[armnetwork.VPNGatewaysClientListResponse]
	newListByResourceGroupPager *azfake.PagerResponder[armnetwork.VPNGatewaysClientListByResourceGroupResponse]
	beginReset                  *azfake.PollerResponder[armnetwork.VPNGatewaysClientResetResponse]
	beginStartPacketCapture     *azfake.PollerResponder[armnetwork.VPNGatewaysClientStartPacketCaptureResponse]
	beginStopPacketCapture      *azfake.PollerResponder[armnetwork.VPNGatewaysClientStopPacketCaptureResponse]
	beginUpdateTags             *azfake.PollerResponder[armnetwork.VPNGatewaysClientUpdateTagsResponse]
}

// Do implements the policy.Transporter interface for VPNGatewaysServerTransport.
func (v *VPNGatewaysServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "VPNGatewaysClient.BeginCreateOrUpdate":
		resp, err = v.dispatchBeginCreateOrUpdate(req)
	case "VPNGatewaysClient.BeginDelete":
		resp, err = v.dispatchBeginDelete(req)
	case "VPNGatewaysClient.Get":
		resp, err = v.dispatchGet(req)
	case "VPNGatewaysClient.NewListPager":
		resp, err = v.dispatchNewListPager(req)
	case "VPNGatewaysClient.NewListByResourceGroupPager":
		resp, err = v.dispatchNewListByResourceGroupPager(req)
	case "VPNGatewaysClient.BeginReset":
		resp, err = v.dispatchBeginReset(req)
	case "VPNGatewaysClient.BeginStartPacketCapture":
		resp, err = v.dispatchBeginStartPacketCapture(req)
	case "VPNGatewaysClient.BeginStopPacketCapture":
		resp, err = v.dispatchBeginStopPacketCapture(req)
	case "VPNGatewaysClient.BeginUpdateTags":
		resp, err = v.dispatchBeginUpdateTags(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (v *VPNGatewaysServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if v.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("method BeginCreateOrUpdate not implemented")}
	}
	if v.beginCreateOrUpdate == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/vpnGateways/(?P<gatewayName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.VPNGateway](req)
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginCreateOrUpdate(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("gatewayName")], body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		v.beginCreateOrUpdate = &respr
	}

	resp, err := server.PollerResponderNext(v.beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(v.beginCreateOrUpdate) {
		v.beginCreateOrUpdate = nil
	}

	return resp, nil
}

func (v *VPNGatewaysServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if v.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("method BeginDelete not implemented")}
	}
	if v.beginDelete == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/vpnGateways/(?P<gatewayName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		respr, errRespr := v.srv.BeginDelete(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("gatewayName")], nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		v.beginDelete = &respr
	}

	resp, err := server.PollerResponderNext(v.beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(v.beginDelete) {
		v.beginDelete = nil
	}

	return resp, nil
}

func (v *VPNGatewaysServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if v.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("method Get not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/vpnGateways/(?P<gatewayName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := v.srv.Get(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("gatewayName")], nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).VPNGateway, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VPNGatewaysServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if v.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("method NewListPager not implemented")}
	}
	if v.newListPager == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/vpnGateways"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := v.srv.NewListPager(nil)
		v.newListPager = &resp
		server.PagerResponderInjectNextLinks(v.newListPager, req, func(page *armnetwork.VPNGatewaysClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(v.newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(v.newListPager) {
		v.newListPager = nil
	}
	return resp, nil
}

func (v *VPNGatewaysServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if v.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("method NewListByResourceGroupPager not implemented")}
	}
	if v.newListByResourceGroupPager == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/vpnGateways"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := v.srv.NewListByResourceGroupPager(matches[regex.SubexpIndex("resourceGroupName")], nil)
		v.newListByResourceGroupPager = &resp
		server.PagerResponderInjectNextLinks(v.newListByResourceGroupPager, req, func(page *armnetwork.VPNGatewaysClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(v.newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(v.newListByResourceGroupPager) {
		v.newListByResourceGroupPager = nil
	}
	return resp, nil
}

func (v *VPNGatewaysServerTransport) dispatchBeginReset(req *http.Request) (*http.Response, error) {
	if v.srv.BeginReset == nil {
		return nil, &nonRetriableError{errors.New("method BeginReset not implemented")}
	}
	if v.beginReset == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/vpnGateways/(?P<gatewayName>[a-zA-Z0-9-_]+)/reset"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		iPConfigurationIDParam := getOptional(qp.Get("ipConfigurationId"))
		var options *armnetwork.VPNGatewaysClientBeginResetOptions
		if iPConfigurationIDParam != nil {
			options = &armnetwork.VPNGatewaysClientBeginResetOptions{
				IPConfigurationID: iPConfigurationIDParam,
			}
		}
		respr, errRespr := v.srv.BeginReset(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("gatewayName")], options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		v.beginReset = &respr
	}

	resp, err := server.PollerResponderNext(v.beginReset, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(v.beginReset) {
		v.beginReset = nil
	}

	return resp, nil
}

func (v *VPNGatewaysServerTransport) dispatchBeginStartPacketCapture(req *http.Request) (*http.Response, error) {
	if v.srv.BeginStartPacketCapture == nil {
		return nil, &nonRetriableError{errors.New("method BeginStartPacketCapture not implemented")}
	}
	if v.beginStartPacketCapture == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/vpnGateways/(?P<gatewayName>[a-zA-Z0-9-_]+)/startpacketcapture"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.VPNGatewayPacketCaptureStartParameters](req)
		if err != nil {
			return nil, err
		}
		var options *armnetwork.VPNGatewaysClientBeginStartPacketCaptureOptions
		if !reflect.ValueOf(body).IsZero() {
			options = &armnetwork.VPNGatewaysClientBeginStartPacketCaptureOptions{
				Parameters: &body,
			}
		}
		respr, errRespr := v.srv.BeginStartPacketCapture(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("gatewayName")], options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		v.beginStartPacketCapture = &respr
	}

	resp, err := server.PollerResponderNext(v.beginStartPacketCapture, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(v.beginStartPacketCapture) {
		v.beginStartPacketCapture = nil
	}

	return resp, nil
}

func (v *VPNGatewaysServerTransport) dispatchBeginStopPacketCapture(req *http.Request) (*http.Response, error) {
	if v.srv.BeginStopPacketCapture == nil {
		return nil, &nonRetriableError{errors.New("method BeginStopPacketCapture not implemented")}
	}
	if v.beginStopPacketCapture == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/vpnGateways/(?P<gatewayName>[a-zA-Z0-9-_]+)/stoppacketcapture"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.VPNGatewayPacketCaptureStopParameters](req)
		if err != nil {
			return nil, err
		}
		var options *armnetwork.VPNGatewaysClientBeginStopPacketCaptureOptions
		if !reflect.ValueOf(body).IsZero() {
			options = &armnetwork.VPNGatewaysClientBeginStopPacketCaptureOptions{
				Parameters: &body,
			}
		}
		respr, errRespr := v.srv.BeginStopPacketCapture(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("gatewayName")], options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		v.beginStopPacketCapture = &respr
	}

	resp, err := server.PollerResponderNext(v.beginStopPacketCapture, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(v.beginStopPacketCapture) {
		v.beginStopPacketCapture = nil
	}

	return resp, nil
}

func (v *VPNGatewaysServerTransport) dispatchBeginUpdateTags(req *http.Request) (*http.Response, error) {
	if v.srv.BeginUpdateTags == nil {
		return nil, &nonRetriableError{errors.New("method BeginUpdateTags not implemented")}
	}
	if v.beginUpdateTags == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/vpnGateways/(?P<gatewayName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.TagsObject](req)
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginUpdateTags(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("gatewayName")], body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		v.beginUpdateTags = &respr
	}

	resp, err := server.PollerResponderNext(v.beginUpdateTags, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(v.beginUpdateTags) {
		v.beginUpdateTags = nil
	}

	return resp, nil
}
