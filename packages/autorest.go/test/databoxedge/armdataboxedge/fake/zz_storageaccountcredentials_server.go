//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"armdataboxedge"
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

// StorageAccountCredentialsServer is a fake server for instances of the armdataboxedge.StorageAccountCredentialsClient type.
type StorageAccountCredentialsServer struct {
	// BeginCreateOrUpdate is the fake for method StorageAccountCredentialsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginCreateOrUpdate func(ctx context.Context, deviceName string, name string, resourceGroupName string, storageAccountCredential armdataboxedge.StorageAccountCredential, options *armdataboxedge.StorageAccountCredentialsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armdataboxedge.StorageAccountCredentialsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method StorageAccountCredentialsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, deviceName string, name string, resourceGroupName string, options *armdataboxedge.StorageAccountCredentialsClientBeginDeleteOptions) (resp azfake.PollerResponder[armdataboxedge.StorageAccountCredentialsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method StorageAccountCredentialsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, deviceName string, name string, resourceGroupName string, options *armdataboxedge.StorageAccountCredentialsClientGetOptions) (resp azfake.Responder[armdataboxedge.StorageAccountCredentialsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByDataBoxEdgeDevicePager is the fake for method StorageAccountCredentialsClient.NewListByDataBoxEdgeDevicePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByDataBoxEdgeDevicePager func(deviceName string, resourceGroupName string, options *armdataboxedge.StorageAccountCredentialsClientListByDataBoxEdgeDeviceOptions) (resp azfake.PagerResponder[armdataboxedge.StorageAccountCredentialsClientListByDataBoxEdgeDeviceResponse])
}

// NewStorageAccountCredentialsServerTransport creates a new instance of StorageAccountCredentialsServerTransport with the provided implementation.
// The returned StorageAccountCredentialsServerTransport instance is connected to an instance of armdataboxedge.StorageAccountCredentialsClient by way of the
// undefined.Transporter field.
func NewStorageAccountCredentialsServerTransport(srv *StorageAccountCredentialsServer) *StorageAccountCredentialsServerTransport {
	return &StorageAccountCredentialsServerTransport{srv: srv}
}

// StorageAccountCredentialsServerTransport connects instances of armdataboxedge.StorageAccountCredentialsClient to instances of StorageAccountCredentialsServer.
// Don't use this type directly, use NewStorageAccountCredentialsServerTransport instead.
type StorageAccountCredentialsServerTransport struct {
	srv                             *StorageAccountCredentialsServer
	beginCreateOrUpdate             *azfake.PollerResponder[armdataboxedge.StorageAccountCredentialsClientCreateOrUpdateResponse]
	beginDelete                     *azfake.PollerResponder[armdataboxedge.StorageAccountCredentialsClientDeleteResponse]
	newListByDataBoxEdgeDevicePager *azfake.PagerResponder[armdataboxedge.StorageAccountCredentialsClientListByDataBoxEdgeDeviceResponse]
}

// Do implements the policy.Transporter interface for StorageAccountCredentialsServerTransport.
func (s *StorageAccountCredentialsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "StorageAccountCredentialsClient.BeginCreateOrUpdate":
		resp, err = s.dispatchBeginCreateOrUpdate(req)
	case "StorageAccountCredentialsClient.BeginDelete":
		resp, err = s.dispatchBeginDelete(req)
	case "StorageAccountCredentialsClient.Get":
		resp, err = s.dispatchGet(req)
	case "StorageAccountCredentialsClient.NewListByDataBoxEdgeDevicePager":
		resp, err = s.dispatchNewListByDataBoxEdgeDevicePager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *StorageAccountCredentialsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if s.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("method BeginCreateOrUpdate not implemented")}
	}
	if s.beginCreateOrUpdate == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/(?P<deviceName>[a-zA-Z0-9-_]+)/storageAccountCredentials/(?P<name>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armdataboxedge.StorageAccountCredential](req)
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginCreateOrUpdate(req.Context(), matches[regex.SubexpIndex("deviceName")], matches[regex.SubexpIndex("name")], matches[regex.SubexpIndex("resourceGroupName")], body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		s.beginCreateOrUpdate = &respr
	}

	resp, err := server.PollerResponderNext(s.beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(s.beginCreateOrUpdate) {
		s.beginCreateOrUpdate = nil
	}

	return resp, nil
}

func (s *StorageAccountCredentialsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if s.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("method BeginDelete not implemented")}
	}
	if s.beginDelete == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/(?P<deviceName>[a-zA-Z0-9-_]+)/storageAccountCredentials/(?P<name>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		respr, errRespr := s.srv.BeginDelete(req.Context(), matches[regex.SubexpIndex("deviceName")], matches[regex.SubexpIndex("name")], matches[regex.SubexpIndex("resourceGroupName")], nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		s.beginDelete = &respr
	}

	resp, err := server.PollerResponderNext(s.beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(s.beginDelete) {
		s.beginDelete = nil
	}

	return resp, nil
}

func (s *StorageAccountCredentialsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if s.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("method Get not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/(?P<deviceName>[a-zA-Z0-9-_]+)/storageAccountCredentials/(?P<name>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := s.srv.Get(req.Context(), matches[regex.SubexpIndex("deviceName")], matches[regex.SubexpIndex("name")], matches[regex.SubexpIndex("resourceGroupName")], nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).StorageAccountCredential, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *StorageAccountCredentialsServerTransport) dispatchNewListByDataBoxEdgeDevicePager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListByDataBoxEdgeDevicePager == nil {
		return nil, &nonRetriableError{errors.New("method NewListByDataBoxEdgeDevicePager not implemented")}
	}
	if s.newListByDataBoxEdgeDevicePager == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/(?P<deviceName>[a-zA-Z0-9-_]+)/storageAccountCredentials"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := s.srv.NewListByDataBoxEdgeDevicePager(matches[regex.SubexpIndex("deviceName")], matches[regex.SubexpIndex("resourceGroupName")], nil)
		s.newListByDataBoxEdgeDevicePager = &resp
		server.PagerResponderInjectNextLinks(s.newListByDataBoxEdgeDevicePager, req, func(page *armdataboxedge.StorageAccountCredentialsClientListByDataBoxEdgeDeviceResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(s.newListByDataBoxEdgeDevicePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(s.newListByDataBoxEdgeDevicePager) {
		s.newListByDataBoxEdgeDevicePager = nil
	}
	return resp, nil
}
