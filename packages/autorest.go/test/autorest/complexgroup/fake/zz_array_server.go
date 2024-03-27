// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	"generatortests/complexgroup"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// ArrayServer is a fake server for instances of the complexgroup.ArrayClient type.
type ArrayServer struct {
	// GetEmpty is the fake for method ArrayClient.GetEmpty
	// HTTP status codes to indicate success: http.StatusOK
	GetEmpty func(ctx context.Context, options *complexgroup.ArrayClientGetEmptyOptions) (resp azfake.Responder[complexgroup.ArrayClientGetEmptyResponse], errResp azfake.ErrorResponder)

	// GetNotProvided is the fake for method ArrayClient.GetNotProvided
	// HTTP status codes to indicate success: http.StatusOK
	GetNotProvided func(ctx context.Context, options *complexgroup.ArrayClientGetNotProvidedOptions) (resp azfake.Responder[complexgroup.ArrayClientGetNotProvidedResponse], errResp azfake.ErrorResponder)

	// GetValid is the fake for method ArrayClient.GetValid
	// HTTP status codes to indicate success: http.StatusOK
	GetValid func(ctx context.Context, options *complexgroup.ArrayClientGetValidOptions) (resp azfake.Responder[complexgroup.ArrayClientGetValidResponse], errResp azfake.ErrorResponder)

	// PutEmpty is the fake for method ArrayClient.PutEmpty
	// HTTP status codes to indicate success: http.StatusOK
	PutEmpty func(ctx context.Context, complexBody complexgroup.ArrayWrapper, options *complexgroup.ArrayClientPutEmptyOptions) (resp azfake.Responder[complexgroup.ArrayClientPutEmptyResponse], errResp azfake.ErrorResponder)

	// PutValid is the fake for method ArrayClient.PutValid
	// HTTP status codes to indicate success: http.StatusOK
	PutValid func(ctx context.Context, complexBody complexgroup.ArrayWrapper, options *complexgroup.ArrayClientPutValidOptions) (resp azfake.Responder[complexgroup.ArrayClientPutValidResponse], errResp azfake.ErrorResponder)
}

// NewArrayServerTransport creates a new instance of ArrayServerTransport with the provided implementation.
// The returned ArrayServerTransport instance is connected to an instance of complexgroup.ArrayClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewArrayServerTransport(srv *ArrayServer) *ArrayServerTransport {
	return &ArrayServerTransport{srv: srv}
}

// ArrayServerTransport connects instances of complexgroup.ArrayClient to instances of ArrayServer.
// Don't use this type directly, use NewArrayServerTransport instead.
type ArrayServerTransport struct {
	srv *ArrayServer
}

// Do implements the policy.Transporter interface for ArrayServerTransport.
func (a *ArrayServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return a.dispatchToMethodFake(req, method)
}

func (a *ArrayServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch method {
	case "ArrayClient.GetEmpty":
		resp, err = a.dispatchGetEmpty(req)
	case "ArrayClient.GetNotProvided":
		resp, err = a.dispatchGetNotProvided(req)
	case "ArrayClient.GetValid":
		resp, err = a.dispatchGetValid(req)
	case "ArrayClient.PutEmpty":
		resp, err = a.dispatchPutEmpty(req)
	case "ArrayClient.PutValid":
		resp, err = a.dispatchPutValid(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	return resp, err
}

func (a *ArrayServerTransport) dispatchGetEmpty(req *http.Request) (*http.Response, error) {
	if a.srv.GetEmpty == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetEmpty not implemented")}
	}
	respr, errRespr := a.srv.GetEmpty(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ArrayWrapper, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *ArrayServerTransport) dispatchGetNotProvided(req *http.Request) (*http.Response, error) {
	if a.srv.GetNotProvided == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetNotProvided not implemented")}
	}
	respr, errRespr := a.srv.GetNotProvided(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ArrayWrapper, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *ArrayServerTransport) dispatchGetValid(req *http.Request) (*http.Response, error) {
	if a.srv.GetValid == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetValid not implemented")}
	}
	respr, errRespr := a.srv.GetValid(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ArrayWrapper, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *ArrayServerTransport) dispatchPutEmpty(req *http.Request) (*http.Response, error) {
	if a.srv.PutEmpty == nil {
		return nil, &nonRetriableError{errors.New("fake for method PutEmpty not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[complexgroup.ArrayWrapper](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.PutEmpty(req.Context(), body, nil)
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

func (a *ArrayServerTransport) dispatchPutValid(req *http.Request) (*http.Response, error) {
	if a.srv.PutValid == nil {
		return nil, &nonRetriableError{errors.New("fake for method PutValid not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[complexgroup.ArrayWrapper](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.PutValid(req.Context(), body, nil)
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
