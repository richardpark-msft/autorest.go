//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"azartifacts"
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"regexp"
)

// IntegrationRuntimesServer is a fake server for instances of the azartifacts.IntegrationRuntimesClient type.
type IntegrationRuntimesServer struct {
	// Get is the fake for method IntegrationRuntimesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, integrationRuntimeName string, options *azartifacts.IntegrationRuntimesClientGetOptions) (resp azfake.Responder[azartifacts.IntegrationRuntimesClientGetResponse], errResp azfake.ErrorResponder)

	// List is the fake for method IntegrationRuntimesClient.List
	// HTTP status codes to indicate success: http.StatusOK
	List func(ctx context.Context, options *azartifacts.IntegrationRuntimesClientListOptions) (resp azfake.Responder[azartifacts.IntegrationRuntimesClientListResponse], errResp azfake.ErrorResponder)
}

// NewIntegrationRuntimesServerTransport creates a new instance of IntegrationRuntimesServerTransport with the provided implementation.
// The returned IntegrationRuntimesServerTransport instance is connected to an instance of azartifacts.IntegrationRuntimesClient by way of the
// undefined.Transporter field.
func NewIntegrationRuntimesServerTransport(srv *IntegrationRuntimesServer) *IntegrationRuntimesServerTransport {
	return &IntegrationRuntimesServerTransport{srv: srv}
}

// IntegrationRuntimesServerTransport connects instances of azartifacts.IntegrationRuntimesClient to instances of IntegrationRuntimesServer.
// Don't use this type directly, use NewIntegrationRuntimesServerTransport instead.
type IntegrationRuntimesServerTransport struct {
	srv *IntegrationRuntimesServer
}

// Do implements the policy.Transporter interface for IntegrationRuntimesServerTransport.
func (i *IntegrationRuntimesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "IntegrationRuntimesClient.Get":
		resp, err = i.dispatchGet(req)
	case "IntegrationRuntimesClient.List":
		resp, err = i.dispatchList(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (i *IntegrationRuntimesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if i.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("method Get not implemented")}
	}
	const regexStr = "/integrationRuntimes/(?P<integrationRuntimeName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := i.srv.Get(req.Context(), matches[regex.SubexpIndex("integrationRuntimeName")], nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).IntegrationRuntimeResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *IntegrationRuntimesServerTransport) dispatchList(req *http.Request) (*http.Response, error) {
	if i.srv.List == nil {
		return nil, &nonRetriableError{errors.New("method List not implemented")}
	}
	respr, errRespr := i.srv.List(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).IntegrationRuntimeListResponse, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
