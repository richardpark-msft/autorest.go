// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"context"
	"dictionarygroup"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// UnknownValueServer is a fake server for instances of the dictionarygroup.UnknownValueClient type.
type UnknownValueServer struct {
	// Get is the fake for method UnknownValueClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, options *dictionarygroup.UnknownValueClientGetOptions) (resp azfake.Responder[dictionarygroup.UnknownValueClientGetResponse], errResp azfake.ErrorResponder)

	// Put is the fake for method UnknownValueClient.Put
	// HTTP status codes to indicate success: http.StatusNoContent
	Put func(ctx context.Context, body map[string]any, options *dictionarygroup.UnknownValueClientPutOptions) (resp azfake.Responder[dictionarygroup.UnknownValueClientPutResponse], errResp azfake.ErrorResponder)
}

// NewUnknownValueServerTransport creates a new instance of UnknownValueServerTransport with the provided implementation.
// The returned UnknownValueServerTransport instance is connected to an instance of dictionarygroup.UnknownValueClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewUnknownValueServerTransport(srv *UnknownValueServer) *UnknownValueServerTransport {
	return &UnknownValueServerTransport{srv: srv}
}

// UnknownValueServerTransport connects instances of dictionarygroup.UnknownValueClient to instances of UnknownValueServer.
// Don't use this type directly, use NewUnknownValueServerTransport instead.
type UnknownValueServerTransport struct {
	srv *UnknownValueServer
}

// Do implements the policy.Transporter interface for UnknownValueServerTransport.
func (u *UnknownValueServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return u.dispatchToMethodFake(req, method)
}

func (u *UnknownValueServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch method {
	case "UnknownValueClient.Get":
		resp, err = u.dispatchGet(req)
	case "UnknownValueClient.Put":
		resp, err = u.dispatchPut(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	return resp, err
}

func (u *UnknownValueServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if u.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	respr, errRespr := u.srv.Get(req.Context(), nil)
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

func (u *UnknownValueServerTransport) dispatchPut(req *http.Request) (*http.Response, error) {
	if u.srv.Put == nil {
		return nil, &nonRetriableError{errors.New("fake for method Put not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[map[string]any](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := u.srv.Put(req.Context(), body, nil)
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
