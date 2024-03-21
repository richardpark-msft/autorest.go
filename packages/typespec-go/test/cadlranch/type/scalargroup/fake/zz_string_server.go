// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"scalargroup"
)

// StringServer is a fake server for instances of the scalargroup.StringClient type.
type StringServer struct {
	// Get is the fake for method StringClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, options *scalargroup.StringClientGetOptions) (resp azfake.Responder[scalargroup.StringClientGetResponse], errResp azfake.ErrorResponder)

	// Put is the fake for method StringClient.Put
	// HTTP status codes to indicate success: http.StatusNoContent
	Put func(ctx context.Context, body string, options *scalargroup.StringClientPutOptions) (resp azfake.Responder[scalargroup.StringClientPutResponse], errResp azfake.ErrorResponder)
}

// NewStringServerTransport creates a new instance of StringServerTransport with the provided implementation.
// The returned StringServerTransport instance is connected to an instance of scalargroup.StringClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewStringServerTransport(srv *StringServer) *StringServerTransport {
	return &StringServerTransport{srv: srv}
}

// StringServerTransport connects instances of scalargroup.StringClient to instances of StringServer.
// Don't use this type directly, use NewStringServerTransport instead.
type StringServerTransport struct {
	srv *StringServer
}

// Do implements the policy.Transporter interface for StringServerTransport.
func (s *StringServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "StringClient.Get":
		resp, err = s.dispatchGet(req)
	case "StringClient.Put":
		resp, err = s.dispatchPut(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *StringServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if s.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	respr, errRespr := s.srv.Get(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Value, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *StringServerTransport) dispatchPut(req *http.Request) (*http.Response, error) {
	if s.srv.Put == nil {
		return nil, &nonRetriableError{errors.New("fake for method Put not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[string](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Put(req.Context(), body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
