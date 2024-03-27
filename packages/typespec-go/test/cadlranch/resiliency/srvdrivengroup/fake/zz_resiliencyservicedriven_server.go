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
	"net/url"
	"srvdrivengroup"
)

// ResiliencyServiceDrivenServer is a fake server for instances of the srvdrivengroup.ResiliencyServiceDrivenClient type.
type ResiliencyServiceDrivenServer struct {
	// AddOperation is the fake for method ResiliencyServiceDrivenClient.AddOperation
	// HTTP status codes to indicate success: http.StatusNoContent
	AddOperation func(ctx context.Context, options *srvdrivengroup.ResiliencyServiceDrivenClientAddOperationOptions) (resp azfake.Responder[srvdrivengroup.ResiliencyServiceDrivenClientAddOperationResponse], errResp azfake.ErrorResponder)

	// FromNone is the fake for method ResiliencyServiceDrivenClient.FromNone
	// HTTP status codes to indicate success: http.StatusNoContent
	FromNone func(ctx context.Context, options *srvdrivengroup.ResiliencyServiceDrivenClientFromNoneOptions) (resp azfake.Responder[srvdrivengroup.ResiliencyServiceDrivenClientFromNoneResponse], errResp azfake.ErrorResponder)

	// FromOneOptional is the fake for method ResiliencyServiceDrivenClient.FromOneOptional
	// HTTP status codes to indicate success: http.StatusNoContent
	FromOneOptional func(ctx context.Context, options *srvdrivengroup.ResiliencyServiceDrivenClientFromOneOptionalOptions) (resp azfake.Responder[srvdrivengroup.ResiliencyServiceDrivenClientFromOneOptionalResponse], errResp azfake.ErrorResponder)

	// FromOneRequired is the fake for method ResiliencyServiceDrivenClient.FromOneRequired
	// HTTP status codes to indicate success: http.StatusNoContent
	FromOneRequired func(ctx context.Context, parameter string, options *srvdrivengroup.ResiliencyServiceDrivenClientFromOneRequiredOptions) (resp azfake.Responder[srvdrivengroup.ResiliencyServiceDrivenClientFromOneRequiredResponse], errResp azfake.ErrorResponder)
}

// NewResiliencyServiceDrivenServerTransport creates a new instance of ResiliencyServiceDrivenServerTransport with the provided implementation.
// The returned ResiliencyServiceDrivenServerTransport instance is connected to an instance of srvdrivengroup.ResiliencyServiceDrivenClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewResiliencyServiceDrivenServerTransport(srv *ResiliencyServiceDrivenServer) *ResiliencyServiceDrivenServerTransport {
	return &ResiliencyServiceDrivenServerTransport{srv: srv}
}

// ResiliencyServiceDrivenServerTransport connects instances of srvdrivengroup.ResiliencyServiceDrivenClient to instances of ResiliencyServiceDrivenServer.
// Don't use this type directly, use NewResiliencyServiceDrivenServerTransport instead.
type ResiliencyServiceDrivenServerTransport struct {
	srv *ResiliencyServiceDrivenServer
}

// Do implements the policy.Transporter interface for ResiliencyServiceDrivenServerTransport.
func (r *ResiliencyServiceDrivenServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return r.dispatchToMethodFake(req, method)
}

func (r *ResiliencyServiceDrivenServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch method {
	case "ResiliencyServiceDrivenClient.AddOperation":
		resp, err = r.dispatchAddOperation(req)
	case "ResiliencyServiceDrivenClient.FromNone":
		resp, err = r.dispatchFromNone(req)
	case "ResiliencyServiceDrivenClient.FromOneOptional":
		resp, err = r.dispatchFromOneOptional(req)
	case "ResiliencyServiceDrivenClient.FromOneRequired":
		resp, err = r.dispatchFromOneRequired(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	return resp, err
}

func (r *ResiliencyServiceDrivenServerTransport) dispatchAddOperation(req *http.Request) (*http.Response, error) {
	if r.srv.AddOperation == nil {
		return nil, &nonRetriableError{errors.New("fake for method AddOperation not implemented")}
	}
	respr, errRespr := r.srv.AddOperation(req.Context(), nil)
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

func (r *ResiliencyServiceDrivenServerTransport) dispatchFromNone(req *http.Request) (*http.Response, error) {
	if r.srv.FromNone == nil {
		return nil, &nonRetriableError{errors.New("fake for method FromNone not implemented")}
	}
	qp := req.URL.Query()
	newParameterUnescaped, err := url.QueryUnescape(qp.Get("new-parameter"))
	if err != nil {
		return nil, err
	}
	newParameterParam := getOptional(newParameterUnescaped)
	var options *srvdrivengroup.ResiliencyServiceDrivenClientFromNoneOptions
	if newParameterParam != nil {
		options = &srvdrivengroup.ResiliencyServiceDrivenClientFromNoneOptions{
			NewParameter: newParameterParam,
		}
	}
	respr, errRespr := r.srv.FromNone(req.Context(), options)
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

func (r *ResiliencyServiceDrivenServerTransport) dispatchFromOneOptional(req *http.Request) (*http.Response, error) {
	if r.srv.FromOneOptional == nil {
		return nil, &nonRetriableError{errors.New("fake for method FromOneOptional not implemented")}
	}
	qp := req.URL.Query()
	parameterUnescaped, err := url.QueryUnescape(qp.Get("parameter"))
	if err != nil {
		return nil, err
	}
	parameterParam := getOptional(parameterUnescaped)
	newParameterUnescaped, err := url.QueryUnescape(qp.Get("new-parameter"))
	if err != nil {
		return nil, err
	}
	newParameterParam := getOptional(newParameterUnescaped)
	var options *srvdrivengroup.ResiliencyServiceDrivenClientFromOneOptionalOptions
	if parameterParam != nil || newParameterParam != nil {
		options = &srvdrivengroup.ResiliencyServiceDrivenClientFromOneOptionalOptions{
			Parameter:    parameterParam,
			NewParameter: newParameterParam,
		}
	}
	respr, errRespr := r.srv.FromOneOptional(req.Context(), options)
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

func (r *ResiliencyServiceDrivenServerTransport) dispatchFromOneRequired(req *http.Request) (*http.Response, error) {
	if r.srv.FromOneRequired == nil {
		return nil, &nonRetriableError{errors.New("fake for method FromOneRequired not implemented")}
	}
	qp := req.URL.Query()
	parameterParam, err := url.QueryUnescape(qp.Get("parameter"))
	if err != nil {
		return nil, err
	}
	newParameterUnescaped, err := url.QueryUnescape(qp.Get("new-parameter"))
	if err != nil {
		return nil, err
	}
	newParameterParam := getOptional(newParameterUnescaped)
	var options *srvdrivengroup.ResiliencyServiceDrivenClientFromOneRequiredOptions
	if newParameterParam != nil {
		options = &srvdrivengroup.ResiliencyServiceDrivenClientFromOneRequiredOptions{
			NewParameter: newParameterParam,
		}
	}
	respr, errRespr := r.srv.FromOneRequired(req.Context(), parameterParam, options)
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
