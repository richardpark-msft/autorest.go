// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	"generatortests/azurespecialsgroup"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"regexp"
)

// SkipURLEncodingServer is a fake server for instances of the azurespecialsgroup.SkipURLEncodingClient type.
type SkipURLEncodingServer struct {
	// GetMethodPathValid is the fake for method SkipURLEncodingClient.GetMethodPathValid
	// HTTP status codes to indicate success: http.StatusOK
	GetMethodPathValid func(ctx context.Context, unencodedPathParam string, options *azurespecialsgroup.SkipURLEncodingClientGetMethodPathValidOptions) (resp azfake.Responder[azurespecialsgroup.SkipURLEncodingClientGetMethodPathValidResponse], errResp azfake.ErrorResponder)

	// GetMethodQueryNull is the fake for method SkipURLEncodingClient.GetMethodQueryNull
	// HTTP status codes to indicate success: http.StatusOK
	GetMethodQueryNull func(ctx context.Context, options *azurespecialsgroup.SkipURLEncodingClientGetMethodQueryNullOptions) (resp azfake.Responder[azurespecialsgroup.SkipURLEncodingClientGetMethodQueryNullResponse], errResp azfake.ErrorResponder)

	// GetMethodQueryValid is the fake for method SkipURLEncodingClient.GetMethodQueryValid
	// HTTP status codes to indicate success: http.StatusOK
	GetMethodQueryValid func(ctx context.Context, q1 string, options *azurespecialsgroup.SkipURLEncodingClientGetMethodQueryValidOptions) (resp azfake.Responder[azurespecialsgroup.SkipURLEncodingClientGetMethodQueryValidResponse], errResp azfake.ErrorResponder)

	// GetPathQueryValid is the fake for method SkipURLEncodingClient.GetPathQueryValid
	// HTTP status codes to indicate success: http.StatusOK
	GetPathQueryValid func(ctx context.Context, q1 string, options *azurespecialsgroup.SkipURLEncodingClientGetPathQueryValidOptions) (resp azfake.Responder[azurespecialsgroup.SkipURLEncodingClientGetPathQueryValidResponse], errResp azfake.ErrorResponder)

	// GetPathValid is the fake for method SkipURLEncodingClient.GetPathValid
	// HTTP status codes to indicate success: http.StatusOK
	GetPathValid func(ctx context.Context, unencodedPathParam string, options *azurespecialsgroup.SkipURLEncodingClientGetPathValidOptions) (resp azfake.Responder[azurespecialsgroup.SkipURLEncodingClientGetPathValidResponse], errResp azfake.ErrorResponder)

	// GetSwaggerPathValid is the fake for method SkipURLEncodingClient.GetSwaggerPathValid
	// HTTP status codes to indicate success: http.StatusOK
	GetSwaggerPathValid func(ctx context.Context, options *azurespecialsgroup.SkipURLEncodingClientGetSwaggerPathValidOptions) (resp azfake.Responder[azurespecialsgroup.SkipURLEncodingClientGetSwaggerPathValidResponse], errResp azfake.ErrorResponder)

	// GetSwaggerQueryValid is the fake for method SkipURLEncodingClient.GetSwaggerQueryValid
	// HTTP status codes to indicate success: http.StatusOK
	GetSwaggerQueryValid func(ctx context.Context, options *azurespecialsgroup.SkipURLEncodingClientGetSwaggerQueryValidOptions) (resp azfake.Responder[azurespecialsgroup.SkipURLEncodingClientGetSwaggerQueryValidResponse], errResp azfake.ErrorResponder)
}

// NewSkipURLEncodingServerTransport creates a new instance of SkipURLEncodingServerTransport with the provided implementation.
// The returned SkipURLEncodingServerTransport instance is connected to an instance of azurespecialsgroup.SkipURLEncodingClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewSkipURLEncodingServerTransport(srv *SkipURLEncodingServer) *SkipURLEncodingServerTransport {
	return &SkipURLEncodingServerTransport{srv: srv}
}

// SkipURLEncodingServerTransport connects instances of azurespecialsgroup.SkipURLEncodingClient to instances of SkipURLEncodingServer.
// Don't use this type directly, use NewSkipURLEncodingServerTransport instead.
type SkipURLEncodingServerTransport struct {
	srv *SkipURLEncodingServer
}

// Do implements the policy.Transporter interface for SkipURLEncodingServerTransport.
func (s *SkipURLEncodingServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return s.dispatchToMethodFake(req, method)
}

func (s *SkipURLEncodingServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch method {
	case "SkipURLEncodingClient.GetMethodPathValid":
		resp, err = s.dispatchGetMethodPathValid(req)
	case "SkipURLEncodingClient.GetMethodQueryNull":
		resp, err = s.dispatchGetMethodQueryNull(req)
	case "SkipURLEncodingClient.GetMethodQueryValid":
		resp, err = s.dispatchGetMethodQueryValid(req)
	case "SkipURLEncodingClient.GetPathQueryValid":
		resp, err = s.dispatchGetPathQueryValid(req)
	case "SkipURLEncodingClient.GetPathValid":
		resp, err = s.dispatchGetPathValid(req)
	case "SkipURLEncodingClient.GetSwaggerPathValid":
		resp, err = s.dispatchGetSwaggerPathValid(req)
	case "SkipURLEncodingClient.GetSwaggerQueryValid":
		resp, err = s.dispatchGetSwaggerQueryValid(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	return resp, err
}

func (s *SkipURLEncodingServerTransport) dispatchGetMethodPathValid(req *http.Request) (*http.Response, error) {
	if s.srv.GetMethodPathValid == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetMethodPathValid not implemented")}
	}
	const regexStr = `/azurespecials/skipUrlEncoding/method/path/valid/(?P<unencodedPathParam>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	unencodedPathParamParam, err := url.PathUnescape(matches[regex.SubexpIndex("unencodedPathParam")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.GetMethodPathValid(req.Context(), unencodedPathParamParam, nil)
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

func (s *SkipURLEncodingServerTransport) dispatchGetMethodQueryNull(req *http.Request) (*http.Response, error) {
	if s.srv.GetMethodQueryNull == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetMethodQueryNull not implemented")}
	}
	qp := req.URL.Query()
	q1Unescaped, err := url.QueryUnescape(qp.Get("q1"))
	if err != nil {
		return nil, err
	}
	q1Param := getOptional(q1Unescaped)
	var options *azurespecialsgroup.SkipURLEncodingClientGetMethodQueryNullOptions
	if q1Param != nil {
		options = &azurespecialsgroup.SkipURLEncodingClientGetMethodQueryNullOptions{
			Q1: q1Param,
		}
	}
	respr, errRespr := s.srv.GetMethodQueryNull(req.Context(), options)
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

func (s *SkipURLEncodingServerTransport) dispatchGetMethodQueryValid(req *http.Request) (*http.Response, error) {
	if s.srv.GetMethodQueryValid == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetMethodQueryValid not implemented")}
	}
	qp := req.URL.Query()
	q1Param, err := url.QueryUnescape(qp.Get("q1"))
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.GetMethodQueryValid(req.Context(), q1Param, nil)
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

func (s *SkipURLEncodingServerTransport) dispatchGetPathQueryValid(req *http.Request) (*http.Response, error) {
	if s.srv.GetPathQueryValid == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetPathQueryValid not implemented")}
	}
	qp := req.URL.Query()
	q1Param, err := url.QueryUnescape(qp.Get("q1"))
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.GetPathQueryValid(req.Context(), q1Param, nil)
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

func (s *SkipURLEncodingServerTransport) dispatchGetPathValid(req *http.Request) (*http.Response, error) {
	if s.srv.GetPathValid == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetPathValid not implemented")}
	}
	const regexStr = `/azurespecials/skipUrlEncoding/path/path/valid/(?P<unencodedPathParam>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	unencodedPathParamParam, err := url.PathUnescape(matches[regex.SubexpIndex("unencodedPathParam")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.GetPathValid(req.Context(), unencodedPathParamParam, nil)
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

func (s *SkipURLEncodingServerTransport) dispatchGetSwaggerPathValid(req *http.Request) (*http.Response, error) {
	if s.srv.GetSwaggerPathValid == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetSwaggerPathValid not implemented")}
	}
	respr, errRespr := s.srv.GetSwaggerPathValid(req.Context(), nil)
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

func (s *SkipURLEncodingServerTransport) dispatchGetSwaggerQueryValid(req *http.Request) (*http.Response, error) {
	if s.srv.GetSwaggerQueryValid == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetSwaggerQueryValid not implemented")}
	}
	respr, errRespr := s.srv.GetSwaggerQueryValid(req.Context(), nil)
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
