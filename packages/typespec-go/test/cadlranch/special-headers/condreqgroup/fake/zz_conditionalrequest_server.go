// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"condreqgroup"
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// ConditionalRequestServer is a fake server for instances of the condreqgroup.ConditionalRequestClient type.
type ConditionalRequestServer struct {
	// PostIfMatch is the fake for method ConditionalRequestClient.PostIfMatch
	// HTTP status codes to indicate success: http.StatusNoContent
	PostIfMatch func(ctx context.Context, options *condreqgroup.ConditionalRequestClientPostIfMatchOptions) (resp azfake.Responder[condreqgroup.ConditionalRequestClientPostIfMatchResponse], errResp azfake.ErrorResponder)

	// PostIfNoneMatch is the fake for method ConditionalRequestClient.PostIfNoneMatch
	// HTTP status codes to indicate success: http.StatusNoContent
	PostIfNoneMatch func(ctx context.Context, options *condreqgroup.ConditionalRequestClientPostIfNoneMatchOptions) (resp azfake.Responder[condreqgroup.ConditionalRequestClientPostIfNoneMatchResponse], errResp azfake.ErrorResponder)
}

// NewConditionalRequestServerTransport creates a new instance of ConditionalRequestServerTransport with the provided implementation.
// The returned ConditionalRequestServerTransport instance is connected to an instance of condreqgroup.ConditionalRequestClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewConditionalRequestServerTransport(srv *ConditionalRequestServer) *ConditionalRequestServerTransport {
	return &ConditionalRequestServerTransport{srv: srv}
}

// ConditionalRequestServerTransport connects instances of condreqgroup.ConditionalRequestClient to instances of ConditionalRequestServer.
// Don't use this type directly, use NewConditionalRequestServerTransport instead.
type ConditionalRequestServerTransport struct {
	srv *ConditionalRequestServer
}

// Do implements the policy.Transporter interface for ConditionalRequestServerTransport.
func (c *ConditionalRequestServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return c.dispatchToMethodFake(req, method)
}

func (c *ConditionalRequestServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch method {
	case "ConditionalRequestClient.PostIfMatch":
		resp, err = c.dispatchPostIfMatch(req)
	case "ConditionalRequestClient.PostIfNoneMatch":
		resp, err = c.dispatchPostIfNoneMatch(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	return resp, err
}

func (c *ConditionalRequestServerTransport) dispatchPostIfMatch(req *http.Request) (*http.Response, error) {
	if c.srv.PostIfMatch == nil {
		return nil, &nonRetriableError{errors.New("fake for method PostIfMatch not implemented")}
	}
	ifMatchParam := getOptional(getHeaderValue(req.Header, "If-Match"))
	var options *condreqgroup.ConditionalRequestClientPostIfMatchOptions
	if ifMatchParam != nil {
		options = &condreqgroup.ConditionalRequestClientPostIfMatchOptions{
			IfMatch: ifMatchParam,
		}
	}
	respr, errRespr := c.srv.PostIfMatch(req.Context(), options)
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

func (c *ConditionalRequestServerTransport) dispatchPostIfNoneMatch(req *http.Request) (*http.Response, error) {
	if c.srv.PostIfNoneMatch == nil {
		return nil, &nonRetriableError{errors.New("fake for method PostIfNoneMatch not implemented")}
	}
	ifNoneMatchParam := getOptional(getHeaderValue(req.Header, "If-None-Match"))
	var options *condreqgroup.ConditionalRequestClientPostIfNoneMatchOptions
	if ifNoneMatchParam != nil {
		options = &condreqgroup.ConditionalRequestClientPostIfNoneMatchOptions{
			IfNoneMatch: ifNoneMatchParam,
		}
	}
	respr, errRespr := c.srv.PostIfNoneMatch(req.Context(), options)
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
