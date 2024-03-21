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

// BooleanValueServer is a fake server for instances of the arraygroup.BooleanValueClient type.
type BooleanValueServer struct {
	// Get is the fake for method BooleanValueClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, options *arraygroup.BooleanValueClientGetOptions) (resp azfake.Responder[arraygroup.BooleanValueClientGetResponse], errResp azfake.ErrorResponder)

	// Put is the fake for method BooleanValueClient.Put
	// HTTP status codes to indicate success: http.StatusNoContent
	Put func(ctx context.Context, body []bool, options *arraygroup.BooleanValueClientPutOptions) (resp azfake.Responder[arraygroup.BooleanValueClientPutResponse], errResp azfake.ErrorResponder)
}

// NewBooleanValueServerTransport creates a new instance of BooleanValueServerTransport with the provided implementation.
// The returned BooleanValueServerTransport instance is connected to an instance of arraygroup.BooleanValueClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewBooleanValueServerTransport(srv *BooleanValueServer) *BooleanValueServerTransport {
	return &BooleanValueServerTransport{srv: srv}
}

// BooleanValueServerTransport connects instances of arraygroup.BooleanValueClient to instances of BooleanValueServer.
// Don't use this type directly, use NewBooleanValueServerTransport instead.
type BooleanValueServerTransport struct {
	srv *BooleanValueServer
}

// Do implements the policy.Transporter interface for BooleanValueServerTransport.
func (b *BooleanValueServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "BooleanValueClient.Get":
		resp, err = b.dispatchGet(req)
	case "BooleanValueClient.Put":
		resp, err = b.dispatchPut(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (b *BooleanValueServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
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

func (b *BooleanValueServerTransport) dispatchPut(req *http.Request) (*http.Response, error) {
	if b.srv.Put == nil {
		return nil, &nonRetriableError{errors.New("fake for method Put not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[[]bool](req)
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
