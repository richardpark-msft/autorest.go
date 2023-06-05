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

// VirtualMachineRunCommandsServer is a fake server for instances of the armcompute.VirtualMachineRunCommandsClient type.
type VirtualMachineRunCommandsServer struct {
	// BeginCreateOrUpdate is the fake for method VirtualMachineRunCommandsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, vmName string, runCommandName string, runCommand armcompute.VirtualMachineRunCommand, options *armcompute.VirtualMachineRunCommandsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armcompute.VirtualMachineRunCommandsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method VirtualMachineRunCommandsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, vmName string, runCommandName string, options *armcompute.VirtualMachineRunCommandsClientBeginDeleteOptions) (resp azfake.PollerResponder[armcompute.VirtualMachineRunCommandsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method VirtualMachineRunCommandsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, location string, commandID string, options *armcompute.VirtualMachineRunCommandsClientGetOptions) (resp azfake.Responder[armcompute.VirtualMachineRunCommandsClientGetResponse], errResp azfake.ErrorResponder)

	// GetByVirtualMachine is the fake for method VirtualMachineRunCommandsClient.GetByVirtualMachine
	// HTTP status codes to indicate success: http.StatusOK
	GetByVirtualMachine func(ctx context.Context, resourceGroupName string, vmName string, runCommandName string, options *armcompute.VirtualMachineRunCommandsClientGetByVirtualMachineOptions) (resp azfake.Responder[armcompute.VirtualMachineRunCommandsClientGetByVirtualMachineResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method VirtualMachineRunCommandsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(location string, options *armcompute.VirtualMachineRunCommandsClientListOptions) (resp azfake.PagerResponder[armcompute.VirtualMachineRunCommandsClientListResponse])

	// NewListByVirtualMachinePager is the fake for method VirtualMachineRunCommandsClient.NewListByVirtualMachinePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByVirtualMachinePager func(resourceGroupName string, vmName string, options *armcompute.VirtualMachineRunCommandsClientListByVirtualMachineOptions) (resp azfake.PagerResponder[armcompute.VirtualMachineRunCommandsClientListByVirtualMachineResponse])

	// BeginUpdate is the fake for method VirtualMachineRunCommandsClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK
	BeginUpdate func(ctx context.Context, resourceGroupName string, vmName string, runCommandName string, runCommand armcompute.VirtualMachineRunCommandUpdate, options *armcompute.VirtualMachineRunCommandsClientBeginUpdateOptions) (resp azfake.PollerResponder[armcompute.VirtualMachineRunCommandsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewVirtualMachineRunCommandsServerTransport creates a new instance of VirtualMachineRunCommandsServerTransport with the provided implementation.
// The returned VirtualMachineRunCommandsServerTransport instance is connected to an instance of armcompute.VirtualMachineRunCommandsClient by way of the
// undefined.Transporter field.
func NewVirtualMachineRunCommandsServerTransport(srv *VirtualMachineRunCommandsServer) *VirtualMachineRunCommandsServerTransport {
	return &VirtualMachineRunCommandsServerTransport{srv: srv}
}

// VirtualMachineRunCommandsServerTransport connects instances of armcompute.VirtualMachineRunCommandsClient to instances of VirtualMachineRunCommandsServer.
// Don't use this type directly, use NewVirtualMachineRunCommandsServerTransport instead.
type VirtualMachineRunCommandsServerTransport struct {
	srv                          *VirtualMachineRunCommandsServer
	beginCreateOrUpdate          *azfake.PollerResponder[armcompute.VirtualMachineRunCommandsClientCreateOrUpdateResponse]
	beginDelete                  *azfake.PollerResponder[armcompute.VirtualMachineRunCommandsClientDeleteResponse]
	newListPager                 *azfake.PagerResponder[armcompute.VirtualMachineRunCommandsClientListResponse]
	newListByVirtualMachinePager *azfake.PagerResponder[armcompute.VirtualMachineRunCommandsClientListByVirtualMachineResponse]
	beginUpdate                  *azfake.PollerResponder[armcompute.VirtualMachineRunCommandsClientUpdateResponse]
}

// Do implements the policy.Transporter interface for VirtualMachineRunCommandsServerTransport.
func (v *VirtualMachineRunCommandsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "VirtualMachineRunCommandsClient.BeginCreateOrUpdate":
		resp, err = v.dispatchBeginCreateOrUpdate(req)
	case "VirtualMachineRunCommandsClient.BeginDelete":
		resp, err = v.dispatchBeginDelete(req)
	case "VirtualMachineRunCommandsClient.Get":
		resp, err = v.dispatchGet(req)
	case "VirtualMachineRunCommandsClient.GetByVirtualMachine":
		resp, err = v.dispatchGetByVirtualMachine(req)
	case "VirtualMachineRunCommandsClient.NewListPager":
		resp, err = v.dispatchNewListPager(req)
	case "VirtualMachineRunCommandsClient.NewListByVirtualMachinePager":
		resp, err = v.dispatchNewListByVirtualMachinePager(req)
	case "VirtualMachineRunCommandsClient.BeginUpdate":
		resp, err = v.dispatchBeginUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (v *VirtualMachineRunCommandsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if v.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("method BeginCreateOrUpdate not implemented")}
	}
	if v.beginCreateOrUpdate == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/virtualMachines/(?P<vmName>[a-zA-Z0-9-_]+)/runCommands/(?P<runCommandName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcompute.VirtualMachineRunCommand](req)
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginCreateOrUpdate(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("vmName")], matches[regex.SubexpIndex("runCommandName")], body, nil)
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

func (v *VirtualMachineRunCommandsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if v.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("method BeginDelete not implemented")}
	}
	if v.beginDelete == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/virtualMachines/(?P<vmName>[a-zA-Z0-9-_]+)/runCommands/(?P<runCommandName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		respr, errRespr := v.srv.BeginDelete(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("vmName")], matches[regex.SubexpIndex("runCommandName")], nil)
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

func (v *VirtualMachineRunCommandsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if v.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("method Get not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/locations/(?P<location>[a-zA-Z0-9-_]+)/runCommands/(?P<commandId>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := v.srv.Get(req.Context(), matches[regex.SubexpIndex("location")], matches[regex.SubexpIndex("commandId")], nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).RunCommandDocument, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VirtualMachineRunCommandsServerTransport) dispatchGetByVirtualMachine(req *http.Request) (*http.Response, error) {
	if v.srv.GetByVirtualMachine == nil {
		return nil, &nonRetriableError{errors.New("method GetByVirtualMachine not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/virtualMachines/(?P<vmName>[a-zA-Z0-9-_]+)/runCommands/(?P<runCommandName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	expandParam := getOptional(qp.Get("$expand"))
	var options *armcompute.VirtualMachineRunCommandsClientGetByVirtualMachineOptions
	if expandParam != nil {
		options = &armcompute.VirtualMachineRunCommandsClientGetByVirtualMachineOptions{
			Expand: expandParam,
		}
	}
	respr, errRespr := v.srv.GetByVirtualMachine(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("vmName")], matches[regex.SubexpIndex("runCommandName")], options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).VirtualMachineRunCommand, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VirtualMachineRunCommandsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if v.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("method NewListPager not implemented")}
	}
	if v.newListPager == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/locations/(?P<location>[a-zA-Z0-9-_]+)/runCommands"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := v.srv.NewListPager(matches[regex.SubexpIndex("location")], nil)
		v.newListPager = &resp
		server.PagerResponderInjectNextLinks(v.newListPager, req, func(page *armcompute.VirtualMachineRunCommandsClientListResponse, createLink func() string) {
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

func (v *VirtualMachineRunCommandsServerTransport) dispatchNewListByVirtualMachinePager(req *http.Request) (*http.Response, error) {
	if v.srv.NewListByVirtualMachinePager == nil {
		return nil, &nonRetriableError{errors.New("method NewListByVirtualMachinePager not implemented")}
	}
	if v.newListByVirtualMachinePager == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/virtualMachines/(?P<vmName>[a-zA-Z0-9-_]+)/runCommands"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		expandParam := getOptional(qp.Get("$expand"))
		var options *armcompute.VirtualMachineRunCommandsClientListByVirtualMachineOptions
		if expandParam != nil {
			options = &armcompute.VirtualMachineRunCommandsClientListByVirtualMachineOptions{
				Expand: expandParam,
			}
		}
		resp := v.srv.NewListByVirtualMachinePager(matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("vmName")], options)
		v.newListByVirtualMachinePager = &resp
		server.PagerResponderInjectNextLinks(v.newListByVirtualMachinePager, req, func(page *armcompute.VirtualMachineRunCommandsClientListByVirtualMachineResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(v.newListByVirtualMachinePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(v.newListByVirtualMachinePager) {
		v.newListByVirtualMachinePager = nil
	}
	return resp, nil
}

func (v *VirtualMachineRunCommandsServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if v.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("method BeginUpdate not implemented")}
	}
	if v.beginUpdate == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/virtualMachines/(?P<vmName>[a-zA-Z0-9-_]+)/runCommands/(?P<runCommandName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcompute.VirtualMachineRunCommandUpdate](req)
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginUpdate(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("vmName")], matches[regex.SubexpIndex("runCommandName")], body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		v.beginUpdate = &respr
	}

	resp, err := server.PollerResponderNext(v.beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PollerResponderMore(v.beginUpdate) {
		v.beginUpdate = nil
	}

	return resp, nil
}
