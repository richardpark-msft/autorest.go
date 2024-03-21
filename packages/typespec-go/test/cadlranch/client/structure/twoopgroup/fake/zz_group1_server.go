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
	"twoopgroup"
)

// Group1Server is a fake server for instances of the twoopgroup.Group1Client type.
type Group1Server struct {
	// Four is the fake for method Group1Client.Four
	// HTTP status codes to indicate success: http.StatusNoContent
	Four func(ctx context.Context, options *twoopgroup.Group1ClientFourOptions) (resp azfake.Responder[twoopgroup.Group1ClientFourResponse], errResp azfake.ErrorResponder)

	// One is the fake for method Group1Client.One
	// HTTP status codes to indicate success: http.StatusNoContent
	One func(ctx context.Context, options *twoopgroup.Group1ClientOneOptions) (resp azfake.Responder[twoopgroup.Group1ClientOneResponse], errResp azfake.ErrorResponder)

	// Three is the fake for method Group1Client.Three
	// HTTP status codes to indicate success: http.StatusNoContent
	Three func(ctx context.Context, options *twoopgroup.Group1ClientThreeOptions) (resp azfake.Responder[twoopgroup.Group1ClientThreeResponse], errResp azfake.ErrorResponder)
}

// NewGroup1ServerTransport creates a new instance of Group1ServerTransport with the provided implementation.
// The returned Group1ServerTransport instance is connected to an instance of twoopgroup.Group1Client via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewGroup1ServerTransport(srv *Group1Server) *Group1ServerTransport {
	return &Group1ServerTransport{srv: srv}
}

// Group1ServerTransport connects instances of twoopgroup.Group1Client to instances of Group1Server.
// Don't use this type directly, use NewGroup1ServerTransport instead.
type Group1ServerTransport struct {
	srv *Group1Server
}

// Do implements the policy.Transporter interface for Group1ServerTransport.
func (g *Group1ServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "Group1Client.Four":
		resp, err = g.dispatchFour(req)
	case "Group1Client.One":
		resp, err = g.dispatchOne(req)
	case "Group1Client.Three":
		resp, err = g.dispatchThree(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (g *Group1ServerTransport) dispatchFour(req *http.Request) (*http.Response, error) {
	if g.srv.Four == nil {
		return nil, &nonRetriableError{errors.New("fake for method Four not implemented")}
	}
	respr, errRespr := g.srv.Four(req.Context(), nil)
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

func (g *Group1ServerTransport) dispatchOne(req *http.Request) (*http.Response, error) {
	if g.srv.One == nil {
		return nil, &nonRetriableError{errors.New("fake for method One not implemented")}
	}
	respr, errRespr := g.srv.One(req.Context(), nil)
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

func (g *Group1ServerTransport) dispatchThree(req *http.Request) (*http.Response, error) {
	if g.srv.Three == nil {
		return nil, &nonRetriableError{errors.New("fake for method Three not implemented")}
	}
	respr, errRespr := g.srv.Three(req.Context(), nil)
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
