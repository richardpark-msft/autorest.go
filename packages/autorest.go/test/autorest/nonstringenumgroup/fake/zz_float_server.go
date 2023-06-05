//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	"generatortests/nonstringenumgroup"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// FloatServer is a fake server for instances of the nonstringenumgroup.FloatClient type.
type FloatServer struct {
	// Get is the fake for method FloatClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, options *nonstringenumgroup.FloatClientGetOptions) (resp azfake.Responder[nonstringenumgroup.FloatClientGetResponse], errResp azfake.ErrorResponder)

	// Put is the fake for method FloatClient.Put
	// HTTP status codes to indicate success: http.StatusOK
	Put func(ctx context.Context, input nonstringenumgroup.FloatEnum, options *nonstringenumgroup.FloatClientPutOptions) (resp azfake.Responder[nonstringenumgroup.FloatClientPutResponse], errResp azfake.ErrorResponder)
}

// NewFloatServerTransport creates a new instance of FloatServerTransport with the provided implementation.
// The returned FloatServerTransport instance is connected to an instance of nonstringenumgroup.FloatClient by way of the
// undefined.Transporter field.
func NewFloatServerTransport(srv *FloatServer) *FloatServerTransport {
	return &FloatServerTransport{srv: srv}
}

// FloatServerTransport connects instances of nonstringenumgroup.FloatClient to instances of FloatServer.
// Don't use this type directly, use NewFloatServerTransport instead.
type FloatServerTransport struct {
	srv *FloatServer
}

// Do implements the policy.Transporter interface for FloatServerTransport.
func (f *FloatServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "FloatClient.Get":
		resp, err = f.dispatchGet(req)
	case "FloatClient.Put":
		resp, err = f.dispatchPut(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (f *FloatServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if f.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("method Get not implemented")}
	}
	respr, errRespr := f.srv.Get(req.Context(), nil)
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

func (f *FloatServerTransport) dispatchPut(req *http.Request) (*http.Response, error) {
	if f.srv.Put == nil {
		return nil, &nonRetriableError{errors.New("method Put not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[nonstringenumgroup.FloatEnum](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := f.srv.Put(req.Context(), body, nil)
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
