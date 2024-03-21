// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"bytesgroup"
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// HeaderServer is a fake server for instances of the bytesgroup.HeaderClient type.
type HeaderServer struct {
	// Base64 is the fake for method HeaderClient.Base64
	// HTTP status codes to indicate success: http.StatusNoContent
	Base64 func(ctx context.Context, value []byte, options *bytesgroup.HeaderClientBase64Options) (resp azfake.Responder[bytesgroup.HeaderClientBase64Response], errResp azfake.ErrorResponder)

	// Base64URL is the fake for method HeaderClient.Base64URL
	// HTTP status codes to indicate success: http.StatusNoContent
	Base64URL func(ctx context.Context, value []byte, options *bytesgroup.HeaderClientBase64URLOptions) (resp azfake.Responder[bytesgroup.HeaderClientBase64URLResponse], errResp azfake.ErrorResponder)

	// Base64URLArray is the fake for method HeaderClient.Base64URLArray
	// HTTP status codes to indicate success: http.StatusNoContent
	Base64URLArray func(ctx context.Context, value [][]byte, options *bytesgroup.HeaderClientBase64URLArrayOptions) (resp azfake.Responder[bytesgroup.HeaderClientBase64URLArrayResponse], errResp azfake.ErrorResponder)

	// Default is the fake for method HeaderClient.Default
	// HTTP status codes to indicate success: http.StatusNoContent
	Default func(ctx context.Context, value []byte, options *bytesgroup.HeaderClientDefaultOptions) (resp azfake.Responder[bytesgroup.HeaderClientDefaultResponse], errResp azfake.ErrorResponder)
}

// NewHeaderServerTransport creates a new instance of HeaderServerTransport with the provided implementation.
// The returned HeaderServerTransport instance is connected to an instance of bytesgroup.HeaderClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewHeaderServerTransport(srv *HeaderServer) *HeaderServerTransport {
	return &HeaderServerTransport{srv: srv}
}

// HeaderServerTransport connects instances of bytesgroup.HeaderClient to instances of HeaderServer.
// Don't use this type directly, use NewHeaderServerTransport instead.
type HeaderServerTransport struct {
	srv *HeaderServer
}

// Do implements the policy.Transporter interface for HeaderServerTransport.
func (h *HeaderServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "HeaderClient.Base64":
		resp, err = h.dispatchBase64(req)
	case "HeaderClient.Base64URL":
		resp, err = h.dispatchBase64URL(req)
	case "HeaderClient.Base64URLArray":
		resp, err = h.dispatchBase64URLArray(req)
	case "HeaderClient.Default":
		resp, err = h.dispatchDefault(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h *HeaderServerTransport) dispatchBase64(req *http.Request) (*http.Response, error) {
	if h.srv.Base64 == nil {
		return nil, &nonRetriableError{errors.New("fake for method Base64 not implemented")}
	}
	valueParam, err := base64.StdEncoding.DecodeString(getHeaderValue(req.Header, "value"))
	if err != nil {
		return nil, err
	}
	respr, errRespr := h.srv.Base64(req.Context(), valueParam, nil)
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

func (h *HeaderServerTransport) dispatchBase64URL(req *http.Request) (*http.Response, error) {
	if h.srv.Base64URL == nil {
		return nil, &nonRetriableError{errors.New("fake for method Base64URL not implemented")}
	}
	valueParam, err := base64.URLEncoding.DecodeString(getHeaderValue(req.Header, "value"))
	if err != nil {
		return nil, err
	}
	respr, errRespr := h.srv.Base64URL(req.Context(), valueParam, nil)
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

func (h *HeaderServerTransport) dispatchBase64URLArray(req *http.Request) (*http.Response, error) {
	if h.srv.Base64URLArray == nil {
		return nil, &nonRetriableError{errors.New("fake for method Base64URLArray not implemented")}
	}
	valueElements := splitHelper(getHeaderValue(req.Header, "value"), ",")
	valueParam := make([][]byte, len(valueElements))
	for i := 0; i < len(valueElements); i++ {
		parsedURL, parseErr := base64.URLEncoding.DecodeString(valueElements[i])
		if parseErr != nil {
			return nil, parseErr
		}
		valueParam[i] = []byte(parsedURL)
	}
	respr, errRespr := h.srv.Base64URLArray(req.Context(), valueParam, nil)
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

func (h *HeaderServerTransport) dispatchDefault(req *http.Request) (*http.Response, error) {
	if h.srv.Default == nil {
		return nil, &nonRetriableError{errors.New("fake for method Default not implemented")}
	}
	valueParam, err := base64.StdEncoding.DecodeString(getHeaderValue(req.Header, "value"))
	if err != nil {
		return nil, err
	}
	respr, errRespr := h.srv.Default(req.Context(), valueParam, nil)
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
