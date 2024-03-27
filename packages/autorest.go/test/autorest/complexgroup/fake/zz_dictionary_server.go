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

// DictionaryServer is a fake server for instances of the complexgroup.DictionaryClient type.
type DictionaryServer struct {
	// GetEmpty is the fake for method DictionaryClient.GetEmpty
	// HTTP status codes to indicate success: http.StatusOK
	GetEmpty func(ctx context.Context, options *complexgroup.DictionaryClientGetEmptyOptions) (resp azfake.Responder[complexgroup.DictionaryClientGetEmptyResponse], errResp azfake.ErrorResponder)

	// GetNotProvided is the fake for method DictionaryClient.GetNotProvided
	// HTTP status codes to indicate success: http.StatusOK
	GetNotProvided func(ctx context.Context, options *complexgroup.DictionaryClientGetNotProvidedOptions) (resp azfake.Responder[complexgroup.DictionaryClientGetNotProvidedResponse], errResp azfake.ErrorResponder)

	// GetNull is the fake for method DictionaryClient.GetNull
	// HTTP status codes to indicate success: http.StatusOK
	GetNull func(ctx context.Context, options *complexgroup.DictionaryClientGetNullOptions) (resp azfake.Responder[complexgroup.DictionaryClientGetNullResponse], errResp azfake.ErrorResponder)

	// GetValid is the fake for method DictionaryClient.GetValid
	// HTTP status codes to indicate success: http.StatusOK
	GetValid func(ctx context.Context, options *complexgroup.DictionaryClientGetValidOptions) (resp azfake.Responder[complexgroup.DictionaryClientGetValidResponse], errResp azfake.ErrorResponder)

	// PutEmpty is the fake for method DictionaryClient.PutEmpty
	// HTTP status codes to indicate success: http.StatusOK
	PutEmpty func(ctx context.Context, complexBody complexgroup.DictionaryWrapper, options *complexgroup.DictionaryClientPutEmptyOptions) (resp azfake.Responder[complexgroup.DictionaryClientPutEmptyResponse], errResp azfake.ErrorResponder)

	// PutValid is the fake for method DictionaryClient.PutValid
	// HTTP status codes to indicate success: http.StatusOK
	PutValid func(ctx context.Context, complexBody complexgroup.DictionaryWrapper, options *complexgroup.DictionaryClientPutValidOptions) (resp azfake.Responder[complexgroup.DictionaryClientPutValidResponse], errResp azfake.ErrorResponder)
}

// NewDictionaryServerTransport creates a new instance of DictionaryServerTransport with the provided implementation.
// The returned DictionaryServerTransport instance is connected to an instance of complexgroup.DictionaryClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewDictionaryServerTransport(srv *DictionaryServer) *DictionaryServerTransport {
	return &DictionaryServerTransport{srv: srv}
}

// DictionaryServerTransport connects instances of complexgroup.DictionaryClient to instances of DictionaryServer.
// Don't use this type directly, use NewDictionaryServerTransport instead.
type DictionaryServerTransport struct {
	srv *DictionaryServer
}

// Do implements the policy.Transporter interface for DictionaryServerTransport.
func (d *DictionaryServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return d.dispatchToMethodFake(req, method)
}

func (d *DictionaryServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch method {
	case "DictionaryClient.GetEmpty":
		resp, err = d.dispatchGetEmpty(req)
	case "DictionaryClient.GetNotProvided":
		resp, err = d.dispatchGetNotProvided(req)
	case "DictionaryClient.GetNull":
		resp, err = d.dispatchGetNull(req)
	case "DictionaryClient.GetValid":
		resp, err = d.dispatchGetValid(req)
	case "DictionaryClient.PutEmpty":
		resp, err = d.dispatchPutEmpty(req)
	case "DictionaryClient.PutValid":
		resp, err = d.dispatchPutValid(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	return resp, err
}

func (d *DictionaryServerTransport) dispatchGetEmpty(req *http.Request) (*http.Response, error) {
	if d.srv.GetEmpty == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetEmpty not implemented")}
	}
	respr, errRespr := d.srv.GetEmpty(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DictionaryWrapper, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DictionaryServerTransport) dispatchGetNotProvided(req *http.Request) (*http.Response, error) {
	if d.srv.GetNotProvided == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetNotProvided not implemented")}
	}
	respr, errRespr := d.srv.GetNotProvided(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DictionaryWrapper, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DictionaryServerTransport) dispatchGetNull(req *http.Request) (*http.Response, error) {
	if d.srv.GetNull == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetNull not implemented")}
	}
	respr, errRespr := d.srv.GetNull(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DictionaryWrapper, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DictionaryServerTransport) dispatchGetValid(req *http.Request) (*http.Response, error) {
	if d.srv.GetValid == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetValid not implemented")}
	}
	respr, errRespr := d.srv.GetValid(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DictionaryWrapper, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DictionaryServerTransport) dispatchPutEmpty(req *http.Request) (*http.Response, error) {
	if d.srv.PutEmpty == nil {
		return nil, &nonRetriableError{errors.New("fake for method PutEmpty not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[complexgroup.DictionaryWrapper](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.PutEmpty(req.Context(), body, nil)
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

func (d *DictionaryServerTransport) dispatchPutValid(req *http.Request) (*http.Response, error) {
	if d.srv.PutValid == nil {
		return nil, &nonRetriableError{errors.New("fake for method PutValid not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[complexgroup.DictionaryWrapper](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.PutValid(req.Context(), body, nil)
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
