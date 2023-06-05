//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	"generatortests/stringgroup"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// StringServer is a fake server for instances of the stringgroup.StringClient type.
type StringServer struct {
	// GetBase64Encoded is the fake for method StringClient.GetBase64Encoded
	// HTTP status codes to indicate success: http.StatusOK
	GetBase64Encoded func(ctx context.Context, options *stringgroup.StringClientGetBase64EncodedOptions) (resp azfake.Responder[stringgroup.StringClientGetBase64EncodedResponse], errResp azfake.ErrorResponder)

	// GetBase64URLEncoded is the fake for method StringClient.GetBase64URLEncoded
	// HTTP status codes to indicate success: http.StatusOK
	GetBase64URLEncoded func(ctx context.Context, options *stringgroup.StringClientGetBase64URLEncodedOptions) (resp azfake.Responder[stringgroup.StringClientGetBase64URLEncodedResponse], errResp azfake.ErrorResponder)

	// GetEmpty is the fake for method StringClient.GetEmpty
	// HTTP status codes to indicate success: http.StatusOK
	GetEmpty func(ctx context.Context, options *stringgroup.StringClientGetEmptyOptions) (resp azfake.Responder[stringgroup.StringClientGetEmptyResponse], errResp azfake.ErrorResponder)

	// GetMBCS is the fake for method StringClient.GetMBCS
	// HTTP status codes to indicate success: http.StatusOK
	GetMBCS func(ctx context.Context, options *stringgroup.StringClientGetMBCSOptions) (resp azfake.Responder[stringgroup.StringClientGetMBCSResponse], errResp azfake.ErrorResponder)

	// GetNotProvided is the fake for method StringClient.GetNotProvided
	// HTTP status codes to indicate success: http.StatusOK
	GetNotProvided func(ctx context.Context, options *stringgroup.StringClientGetNotProvidedOptions) (resp azfake.Responder[stringgroup.StringClientGetNotProvidedResponse], errResp azfake.ErrorResponder)

	// GetNull is the fake for method StringClient.GetNull
	// HTTP status codes to indicate success: http.StatusOK
	GetNull func(ctx context.Context, options *stringgroup.StringClientGetNullOptions) (resp azfake.Responder[stringgroup.StringClientGetNullResponse], errResp azfake.ErrorResponder)

	// GetNullBase64URLEncoded is the fake for method StringClient.GetNullBase64URLEncoded
	// HTTP status codes to indicate success: http.StatusOK
	GetNullBase64URLEncoded func(ctx context.Context, options *stringgroup.StringClientGetNullBase64URLEncodedOptions) (resp azfake.Responder[stringgroup.StringClientGetNullBase64URLEncodedResponse], errResp azfake.ErrorResponder)

	// GetWhitespace is the fake for method StringClient.GetWhitespace
	// HTTP status codes to indicate success: http.StatusOK
	GetWhitespace func(ctx context.Context, options *stringgroup.StringClientGetWhitespaceOptions) (resp azfake.Responder[stringgroup.StringClientGetWhitespaceResponse], errResp azfake.ErrorResponder)

	// PutBase64URLEncoded is the fake for method StringClient.PutBase64URLEncoded
	// HTTP status codes to indicate success: http.StatusOK
	PutBase64URLEncoded func(ctx context.Context, stringBody []byte, options *stringgroup.StringClientPutBase64URLEncodedOptions) (resp azfake.Responder[stringgroup.StringClientPutBase64URLEncodedResponse], errResp azfake.ErrorResponder)

	// PutEmpty is the fake for method StringClient.PutEmpty
	// HTTP status codes to indicate success: http.StatusOK
	PutEmpty func(ctx context.Context, options *stringgroup.StringClientPutEmptyOptions) (resp azfake.Responder[stringgroup.StringClientPutEmptyResponse], errResp azfake.ErrorResponder)

	// PutMBCS is the fake for method StringClient.PutMBCS
	// HTTP status codes to indicate success: http.StatusOK
	PutMBCS func(ctx context.Context, options *stringgroup.StringClientPutMBCSOptions) (resp azfake.Responder[stringgroup.StringClientPutMBCSResponse], errResp azfake.ErrorResponder)

	// PutNull is the fake for method StringClient.PutNull
	// HTTP status codes to indicate success: http.StatusOK
	PutNull func(ctx context.Context, stringBody string, options *stringgroup.StringClientPutNullOptions) (resp azfake.Responder[stringgroup.StringClientPutNullResponse], errResp azfake.ErrorResponder)

	// PutWhitespace is the fake for method StringClient.PutWhitespace
	// HTTP status codes to indicate success: http.StatusOK
	PutWhitespace func(ctx context.Context, options *stringgroup.StringClientPutWhitespaceOptions) (resp azfake.Responder[stringgroup.StringClientPutWhitespaceResponse], errResp azfake.ErrorResponder)
}

// NewStringServerTransport creates a new instance of StringServerTransport with the provided implementation.
// The returned StringServerTransport instance is connected to an instance of stringgroup.StringClient by way of the
// undefined.Transporter field.
func NewStringServerTransport(srv *StringServer) *StringServerTransport {
	return &StringServerTransport{srv: srv}
}

// StringServerTransport connects instances of stringgroup.StringClient to instances of StringServer.
// Don't use this type directly, use NewStringServerTransport instead.
type StringServerTransport struct {
	srv *StringServer
}

// Do implements the policy.Transporter interface for StringServerTransport.
func (s *StringServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "StringClient.GetBase64Encoded":
		resp, err = s.dispatchGetBase64Encoded(req)
	case "StringClient.GetBase64URLEncoded":
		resp, err = s.dispatchGetBase64URLEncoded(req)
	case "StringClient.GetEmpty":
		resp, err = s.dispatchGetEmpty(req)
	case "StringClient.GetMBCS":
		resp, err = s.dispatchGetMBCS(req)
	case "StringClient.GetNotProvided":
		resp, err = s.dispatchGetNotProvided(req)
	case "StringClient.GetNull":
		resp, err = s.dispatchGetNull(req)
	case "StringClient.GetNullBase64URLEncoded":
		resp, err = s.dispatchGetNullBase64URLEncoded(req)
	case "StringClient.GetWhitespace":
		resp, err = s.dispatchGetWhitespace(req)
	case "StringClient.PutBase64URLEncoded":
		resp, err = s.dispatchPutBase64URLEncoded(req)
	case "StringClient.PutEmpty":
		resp, err = s.dispatchPutEmpty(req)
	case "StringClient.PutMBCS":
		resp, err = s.dispatchPutMBCS(req)
	case "StringClient.PutNull":
		resp, err = s.dispatchPutNull(req)
	case "StringClient.PutWhitespace":
		resp, err = s.dispatchPutWhitespace(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *StringServerTransport) dispatchGetBase64Encoded(req *http.Request) (*http.Response, error) {
	if s.srv.GetBase64Encoded == nil {
		return nil, &nonRetriableError{errors.New("method GetBase64Encoded not implemented")}
	}
	respr, errRespr := s.srv.GetBase64Encoded(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsByteArray(respContent, server.GetResponse(respr).Value, runtime.Base64StdFormat, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *StringServerTransport) dispatchGetBase64URLEncoded(req *http.Request) (*http.Response, error) {
	if s.srv.GetBase64URLEncoded == nil {
		return nil, &nonRetriableError{errors.New("method GetBase64URLEncoded not implemented")}
	}
	respr, errRespr := s.srv.GetBase64URLEncoded(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsByteArray(respContent, server.GetResponse(respr).Value, runtime.Base64URLFormat, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *StringServerTransport) dispatchGetEmpty(req *http.Request) (*http.Response, error) {
	if s.srv.GetEmpty == nil {
		return nil, &nonRetriableError{errors.New("method GetEmpty not implemented")}
	}
	respr, errRespr := s.srv.GetEmpty(req.Context(), nil)
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

func (s *StringServerTransport) dispatchGetMBCS(req *http.Request) (*http.Response, error) {
	if s.srv.GetMBCS == nil {
		return nil, &nonRetriableError{errors.New("method GetMBCS not implemented")}
	}
	respr, errRespr := s.srv.GetMBCS(req.Context(), nil)
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

func (s *StringServerTransport) dispatchGetNotProvided(req *http.Request) (*http.Response, error) {
	if s.srv.GetNotProvided == nil {
		return nil, &nonRetriableError{errors.New("method GetNotProvided not implemented")}
	}
	respr, errRespr := s.srv.GetNotProvided(req.Context(), nil)
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

func (s *StringServerTransport) dispatchGetNull(req *http.Request) (*http.Response, error) {
	if s.srv.GetNull == nil {
		return nil, &nonRetriableError{errors.New("method GetNull not implemented")}
	}
	respr, errRespr := s.srv.GetNull(req.Context(), nil)
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

func (s *StringServerTransport) dispatchGetNullBase64URLEncoded(req *http.Request) (*http.Response, error) {
	if s.srv.GetNullBase64URLEncoded == nil {
		return nil, &nonRetriableError{errors.New("method GetNullBase64URLEncoded not implemented")}
	}
	respr, errRespr := s.srv.GetNullBase64URLEncoded(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsByteArray(respContent, server.GetResponse(respr).Value, runtime.Base64URLFormat, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *StringServerTransport) dispatchGetWhitespace(req *http.Request) (*http.Response, error) {
	if s.srv.GetWhitespace == nil {
		return nil, &nonRetriableError{errors.New("method GetWhitespace not implemented")}
	}
	respr, errRespr := s.srv.GetWhitespace(req.Context(), nil)
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

func (s *StringServerTransport) dispatchPutBase64URLEncoded(req *http.Request) (*http.Response, error) {
	if s.srv.PutBase64URLEncoded == nil {
		return nil, &nonRetriableError{errors.New("method PutBase64URLEncoded not implemented")}
	}
	body, err := server.UnmarshalRequestAsByteArray(req, runtime.Base64URLFormat)
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.PutBase64URLEncoded(req.Context(), body, nil)
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

func (s *StringServerTransport) dispatchPutEmpty(req *http.Request) (*http.Response, error) {
	if s.srv.PutEmpty == nil {
		return nil, &nonRetriableError{errors.New("method PutEmpty not implemented")}
	}
	respr, errRespr := s.srv.PutEmpty(req.Context(), nil)
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

func (s *StringServerTransport) dispatchPutMBCS(req *http.Request) (*http.Response, error) {
	if s.srv.PutMBCS == nil {
		return nil, &nonRetriableError{errors.New("method PutMBCS not implemented")}
	}
	respr, errRespr := s.srv.PutMBCS(req.Context(), nil)
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

func (s *StringServerTransport) dispatchPutNull(req *http.Request) (*http.Response, error) {
	if s.srv.PutNull == nil {
		return nil, &nonRetriableError{errors.New("method PutNull not implemented")}
	}
	body, err := server.UnmarshalRequestAsText(req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.PutNull(req.Context(), body, nil)
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

func (s *StringServerTransport) dispatchPutWhitespace(req *http.Request) (*http.Response, error) {
	if s.srv.PutWhitespace == nil {
		return nil, &nonRetriableError{errors.New("method PutWhitespace not implemented")}
	}
	respr, errRespr := s.srv.PutWhitespace(req.Context(), nil)
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
