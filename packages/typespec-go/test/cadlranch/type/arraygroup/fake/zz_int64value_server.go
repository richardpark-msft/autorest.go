// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"arraygroup"
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// Int64ValueServer is a fake server for instances of the arraygroup.Int64ValueClient type.
type Int64ValueServer struct {
	// Get is the fake for method Int64ValueClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, options *arraygroup.Int64ValueClientGetOptions) (resp azfake.Responder[arraygroup.Int64ValueClientGetResponse], errResp azfake.ErrorResponder)

	// Put is the fake for method Int64ValueClient.Put
	// HTTP status codes to indicate success: http.StatusNoContent
	Put func(ctx context.Context, body []int64, options *arraygroup.Int64ValueClientPutOptions) (resp azfake.Responder[arraygroup.Int64ValueClientPutResponse], errResp azfake.ErrorResponder)
}

// NewInt64ValueServerTransport creates a new instance of Int64ValueServerTransport with the provided implementation.
// The returned Int64ValueServerTransport instance is connected to an instance of arraygroup.Int64ValueClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewInt64ValueServerTransport(srv *Int64ValueServer) *Int64ValueServerTransport {
	return &Int64ValueServerTransport{srv: srv}
}

// Int64ValueServerTransport connects instances of arraygroup.Int64ValueClient to instances of Int64ValueServer.
// Don't use this type directly, use NewInt64ValueServerTransport instead.
type Int64ValueServerTransport struct {
	srv *Int64ValueServer
}

// Do implements the policy.Transporter interface for Int64ValueServerTransport.
func (i *Int64ValueServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "Int64ValueClient.Get":
		resp, err = i.dispatchGet(req)
	case "Int64ValueClient.Put":
		resp, err = i.dispatchPut(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (i *Int64ValueServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if i.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	respr, errRespr := i.srv.Get(req.Context(), nil)
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

func (i *Int64ValueServerTransport) dispatchPut(req *http.Request) (*http.Response, error) {
	if i.srv.Put == nil {
		return nil, &nonRetriableError{errors.New("fake for method Put not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[[]int64](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.Put(req.Context(), body, nil)
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
