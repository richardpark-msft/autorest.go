// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"armdataboxedge/v2"
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

// OrdersServer is a fake server for instances of the armdataboxedge.OrdersClient type.
type OrdersServer struct {
	// BeginCreateOrUpdate is the fake for method OrdersClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginCreateOrUpdate func(ctx context.Context, deviceName string, resourceGroupName string, order armdataboxedge.Order, options *armdataboxedge.OrdersClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armdataboxedge.OrdersClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method OrdersClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, deviceName string, resourceGroupName string, options *armdataboxedge.OrdersClientBeginDeleteOptions) (resp azfake.PollerResponder[armdataboxedge.OrdersClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method OrdersClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, deviceName string, resourceGroupName string, options *armdataboxedge.OrdersClientGetOptions) (resp azfake.Responder[armdataboxedge.OrdersClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByDataBoxEdgeDevicePager is the fake for method OrdersClient.NewListByDataBoxEdgeDevicePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByDataBoxEdgeDevicePager func(deviceName string, resourceGroupName string, options *armdataboxedge.OrdersClientListByDataBoxEdgeDeviceOptions) (resp azfake.PagerResponder[armdataboxedge.OrdersClientListByDataBoxEdgeDeviceResponse])

	// ListDCAccessCode is the fake for method OrdersClient.ListDCAccessCode
	// HTTP status codes to indicate success: http.StatusOK
	ListDCAccessCode func(ctx context.Context, deviceName string, resourceGroupName string, options *armdataboxedge.OrdersClientListDCAccessCodeOptions) (resp azfake.Responder[armdataboxedge.OrdersClientListDCAccessCodeResponse], errResp azfake.ErrorResponder)
}

// NewOrdersServerTransport creates a new instance of OrdersServerTransport with the provided implementation.
// The returned OrdersServerTransport instance is connected to an instance of armdataboxedge.OrdersClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewOrdersServerTransport(srv *OrdersServer) *OrdersServerTransport {
	return &OrdersServerTransport{
		srv:                             srv,
		beginCreateOrUpdate:             newTracker[azfake.PollerResponder[armdataboxedge.OrdersClientCreateOrUpdateResponse]](),
		beginDelete:                     newTracker[azfake.PollerResponder[armdataboxedge.OrdersClientDeleteResponse]](),
		newListByDataBoxEdgeDevicePager: newTracker[azfake.PagerResponder[armdataboxedge.OrdersClientListByDataBoxEdgeDeviceResponse]](),
	}
}

// OrdersServerTransport connects instances of armdataboxedge.OrdersClient to instances of OrdersServer.
// Don't use this type directly, use NewOrdersServerTransport instead.
type OrdersServerTransport struct {
	srv                             *OrdersServer
	beginCreateOrUpdate             *tracker[azfake.PollerResponder[armdataboxedge.OrdersClientCreateOrUpdateResponse]]
	beginDelete                     *tracker[azfake.PollerResponder[armdataboxedge.OrdersClientDeleteResponse]]
	newListByDataBoxEdgeDevicePager *tracker[azfake.PagerResponder[armdataboxedge.OrdersClientListByDataBoxEdgeDeviceResponse]]
}

// Do implements the policy.Transporter interface for OrdersServerTransport.
func (o *OrdersServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return o.dispatchToMethodFake(req, method)
}

func (o *OrdersServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch method {
	case "OrdersClient.BeginCreateOrUpdate":
		resp, err = o.dispatchBeginCreateOrUpdate(req)
	case "OrdersClient.BeginDelete":
		resp, err = o.dispatchBeginDelete(req)
	case "OrdersClient.Get":
		resp, err = o.dispatchGet(req)
	case "OrdersClient.NewListByDataBoxEdgeDevicePager":
		resp, err = o.dispatchNewListByDataBoxEdgeDevicePager(req)
	case "OrdersClient.ListDCAccessCode":
		resp, err = o.dispatchListDCAccessCode(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	return resp, err
}

func (o *OrdersServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if o.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := o.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataBoxEdge/dataBoxEdgeDevices/(?P<deviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/orders/default`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armdataboxedge.Order](req)
		if err != nil {
			return nil, err
		}
		deviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deviceName")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := o.srv.BeginCreateOrUpdate(req.Context(), deviceNameParam, resourceGroupNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		o.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		o.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		o.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (o *OrdersServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if o.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := o.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataBoxEdge/dataBoxEdgeDevices/(?P<deviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/orders/default`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		deviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deviceName")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := o.srv.BeginDelete(req.Context(), deviceNameParam, resourceGroupNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		o.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		o.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		o.beginDelete.remove(req)
	}

	return resp, nil
}

func (o *OrdersServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if o.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataBoxEdge/dataBoxEdgeDevices/(?P<deviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/orders/default`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	deviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deviceName")])
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := o.srv.Get(req.Context(), deviceNameParam, resourceGroupNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Order, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (o *OrdersServerTransport) dispatchNewListByDataBoxEdgeDevicePager(req *http.Request) (*http.Response, error) {
	if o.srv.NewListByDataBoxEdgeDevicePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByDataBoxEdgeDevicePager not implemented")}
	}
	newListByDataBoxEdgeDevicePager := o.newListByDataBoxEdgeDevicePager.get(req)
	if newListByDataBoxEdgeDevicePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataBoxEdge/dataBoxEdgeDevices/(?P<deviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/orders`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		deviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deviceName")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := o.srv.NewListByDataBoxEdgeDevicePager(deviceNameParam, resourceGroupNameParam, nil)
		newListByDataBoxEdgeDevicePager = &resp
		o.newListByDataBoxEdgeDevicePager.add(req, newListByDataBoxEdgeDevicePager)
		server.PagerResponderInjectNextLinks(newListByDataBoxEdgeDevicePager, req, func(page *armdataboxedge.OrdersClientListByDataBoxEdgeDeviceResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByDataBoxEdgeDevicePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		o.newListByDataBoxEdgeDevicePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByDataBoxEdgeDevicePager) {
		o.newListByDataBoxEdgeDevicePager.remove(req)
	}
	return resp, nil
}

func (o *OrdersServerTransport) dispatchListDCAccessCode(req *http.Request) (*http.Response, error) {
	if o.srv.ListDCAccessCode == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListDCAccessCode not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataBoxEdge/dataBoxEdgeDevices/(?P<deviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/orders/default/listDCAccessCode`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	deviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deviceName")])
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := o.srv.ListDCAccessCode(req.Context(), deviceNameParam, resourceGroupNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DCAccessCode, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
