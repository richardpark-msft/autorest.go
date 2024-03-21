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
	"renamedopgroup"
)

// GroupServer is a fake server for instances of the renamedopgroup.GroupClient type.
type GroupServer struct {
	// RenamedFour is the fake for method GroupClient.RenamedFour
	// HTTP status codes to indicate success: http.StatusNoContent
	RenamedFour func(ctx context.Context, options *renamedopgroup.GroupClientRenamedFourOptions) (resp azfake.Responder[renamedopgroup.GroupClientRenamedFourResponse], errResp azfake.ErrorResponder)

	// RenamedSix is the fake for method GroupClient.RenamedSix
	// HTTP status codes to indicate success: http.StatusNoContent
	RenamedSix func(ctx context.Context, options *renamedopgroup.GroupClientRenamedSixOptions) (resp azfake.Responder[renamedopgroup.GroupClientRenamedSixResponse], errResp azfake.ErrorResponder)

	// RenamedTwo is the fake for method GroupClient.RenamedTwo
	// HTTP status codes to indicate success: http.StatusNoContent
	RenamedTwo func(ctx context.Context, options *renamedopgroup.GroupClientRenamedTwoOptions) (resp azfake.Responder[renamedopgroup.GroupClientRenamedTwoResponse], errResp azfake.ErrorResponder)
}

// NewGroupServerTransport creates a new instance of GroupServerTransport with the provided implementation.
// The returned GroupServerTransport instance is connected to an instance of renamedopgroup.GroupClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewGroupServerTransport(srv *GroupServer) *GroupServerTransport {
	return &GroupServerTransport{srv: srv}
}

// GroupServerTransport connects instances of renamedopgroup.GroupClient to instances of GroupServer.
// Don't use this type directly, use NewGroupServerTransport instead.
type GroupServerTransport struct {
	srv *GroupServer
}

// Do implements the policy.Transporter interface for GroupServerTransport.
func (g *GroupServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "GroupClient.RenamedFour":
		resp, err = g.dispatchRenamedFour(req)
	case "GroupClient.RenamedSix":
		resp, err = g.dispatchRenamedSix(req)
	case "GroupClient.RenamedTwo":
		resp, err = g.dispatchRenamedTwo(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (g *GroupServerTransport) dispatchRenamedFour(req *http.Request) (*http.Response, error) {
	if g.srv.RenamedFour == nil {
		return nil, &nonRetriableError{errors.New("fake for method RenamedFour not implemented")}
	}
	respr, errRespr := g.srv.RenamedFour(req.Context(), nil)
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

func (g *GroupServerTransport) dispatchRenamedSix(req *http.Request) (*http.Response, error) {
	if g.srv.RenamedSix == nil {
		return nil, &nonRetriableError{errors.New("fake for method RenamedSix not implemented")}
	}
	respr, errRespr := g.srv.RenamedSix(req.Context(), nil)
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

func (g *GroupServerTransport) dispatchRenamedTwo(req *http.Request) (*http.Response, error) {
	if g.srv.RenamedTwo == nil {
		return nil, &nonRetriableError{errors.New("fake for method RenamedTwo not implemented")}
	}
	respr, errRespr := g.srv.RenamedTwo(req.Context(), nil)
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
