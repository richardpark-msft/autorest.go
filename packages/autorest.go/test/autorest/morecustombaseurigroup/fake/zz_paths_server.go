//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package fake

import (
	"context"
	"errors"
	"fmt"
	"generatortests/morecustombaseurigroup"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"regexp"
)

// PathsServer is a fake server for instances of the morecustombaseurigroup.PathsClient type.
type PathsServer struct {
	// GetEmpty is the fake for method PathsClient.GetEmpty
	// HTTP status codes to indicate success: http.StatusOK
	GetEmpty func(ctx context.Context, host string, keyName string, options *morecustombaseurigroup.PathsClientGetEmptyOptions) (resp azfake.Responder[morecustombaseurigroup.PathsClientGetEmptyResponse], errResp azfake.ErrorResponder)
}

// NewPathsServerTransport creates a new instance of PathsServerTransport with the provided implementation.
// The returned PathsServerTransport instance is connected to an instance of morecustombaseurigroup.PathsClient by way of the
// undefined.Transporter field.
func NewPathsServerTransport(srv *PathsServer) *PathsServerTransport {
	return &PathsServerTransport{srv: srv}
}

// PathsServerTransport connects instances of morecustombaseurigroup.PathsClient to instances of PathsServer.
// Don't use this type directly, use NewPathsServerTransport instead.
type PathsServerTransport struct {
	srv *PathsServer
}

// Do implements the policy.Transporter interface for PathsServerTransport.
func (p *PathsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "PathsClient.GetEmpty":
		resp, err = p.dispatchGetEmpty(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *PathsServerTransport) dispatchGetEmpty(req *http.Request) (*http.Response, error) {
	if p.srv.GetEmpty == nil {
		return nil, &nonRetriableError{errors.New("method GetEmpty not implemented")}
	}
	const regexStr = "/customuri/(?P<subscriptionId>[a-zA-Z0-9-_]+)/(?P<keyName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	keyVersionParam := getOptional(qp.Get("keyVersion"))
	var options *morecustombaseurigroup.PathsClientGetEmptyOptions
	if keyVersionParam != nil {
		options = &morecustombaseurigroup.PathsClientGetEmptyOptions{
			KeyVersion: keyVersionParam,
		}
	}
	respr, errRespr := p.srv.GetEmpty(req.Context(), req.URL.Host, matches[regex.SubexpIndex("keyName")], options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}