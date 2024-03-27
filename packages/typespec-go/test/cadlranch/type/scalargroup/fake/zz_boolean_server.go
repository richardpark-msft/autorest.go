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

// BooleanServer is a fake server for instances of the scalargroup.BooleanClient type.
type BooleanServer struct {
	// Get is the fake for method BooleanClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, options *scalargroup.BooleanClientGetOptions) (resp azfake.Responder[scalargroup.BooleanClientGetResponse], errResp azfake.ErrorResponder)

	// Put is the fake for method BooleanClient.Put
	// HTTP status codes to indicate success: http.StatusNoContent
	Put func(ctx context.Context, body bool, options *scalargroup.BooleanClientPutOptions) (resp azfake.Responder[scalargroup.BooleanClientPutResponse], errResp azfake.ErrorResponder)
}

// NewBooleanServerTransport creates a new instance of BooleanServerTransport with the provided implementation.
// The returned BooleanServerTransport instance is connected to an instance of scalargroup.BooleanClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewBooleanServerTransport(srv *BooleanServer) *BooleanServerTransport {
	return &BooleanServerTransport{srv: srv}
}

// BooleanServerTransport connects instances of scalargroup.BooleanClient to instances of BooleanServer.
// Don't use this type directly, use NewBooleanServerTransport instead.
type BooleanServerTransport struct {
	srv *BooleanServer
}

// Do implements the policy.Transporter interface for BooleanServerTransport.
func (b *BooleanServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return b.dispatchToMethodFake(req, method)
}

func (b *BooleanServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch method {
	case "BooleanClient.Get":
		resp, err = b.dispatchGet(req)
	case "BooleanClient.Put":
		resp, err = b.dispatchPut(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	return resp, err
}

func (b *BooleanServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if b.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	respr, errRespr := b.srv.Get(req.Context(), nil)
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

func (b *BooleanServerTransport) dispatchPut(req *http.Request) (*http.Response, error) {
	if b.srv.Put == nil {
		return nil, &nonRetriableError{errors.New("fake for method Put not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[bool](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := b.srv.Put(req.Context(), body, nil)
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