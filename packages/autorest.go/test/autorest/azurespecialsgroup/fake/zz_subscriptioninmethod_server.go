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
	"generatortests/azurespecialsgroup"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"regexp"
)

// SubscriptionInMethodServer is a fake server for instances of the azurespecialsgroup.SubscriptionInMethodClient type.
type SubscriptionInMethodServer struct {
	// PostMethodLocalNull is the fake for method SubscriptionInMethodClient.PostMethodLocalNull
	// HTTP status codes to indicate success: http.StatusOK
	PostMethodLocalNull func(ctx context.Context, subscriptionID string, options *azurespecialsgroup.SubscriptionInMethodClientPostMethodLocalNullOptions) (resp azfake.Responder[azurespecialsgroup.SubscriptionInMethodClientPostMethodLocalNullResponse], errResp azfake.ErrorResponder)

	// PostMethodLocalValid is the fake for method SubscriptionInMethodClient.PostMethodLocalValid
	// HTTP status codes to indicate success: http.StatusOK
	PostMethodLocalValid func(ctx context.Context, subscriptionID string, options *azurespecialsgroup.SubscriptionInMethodClientPostMethodLocalValidOptions) (resp azfake.Responder[azurespecialsgroup.SubscriptionInMethodClientPostMethodLocalValidResponse], errResp azfake.ErrorResponder)

	// PostPathLocalValid is the fake for method SubscriptionInMethodClient.PostPathLocalValid
	// HTTP status codes to indicate success: http.StatusOK
	PostPathLocalValid func(ctx context.Context, subscriptionID string, options *azurespecialsgroup.SubscriptionInMethodClientPostPathLocalValidOptions) (resp azfake.Responder[azurespecialsgroup.SubscriptionInMethodClientPostPathLocalValidResponse], errResp azfake.ErrorResponder)

	// PostSwaggerLocalValid is the fake for method SubscriptionInMethodClient.PostSwaggerLocalValid
	// HTTP status codes to indicate success: http.StatusOK
	PostSwaggerLocalValid func(ctx context.Context, subscriptionID string, options *azurespecialsgroup.SubscriptionInMethodClientPostSwaggerLocalValidOptions) (resp azfake.Responder[azurespecialsgroup.SubscriptionInMethodClientPostSwaggerLocalValidResponse], errResp azfake.ErrorResponder)
}

// NewSubscriptionInMethodServerTransport creates a new instance of SubscriptionInMethodServerTransport with the provided implementation.
// The returned SubscriptionInMethodServerTransport instance is connected to an instance of azurespecialsgroup.SubscriptionInMethodClient by way of the
// undefined.Transporter field.
func NewSubscriptionInMethodServerTransport(srv *SubscriptionInMethodServer) *SubscriptionInMethodServerTransport {
	return &SubscriptionInMethodServerTransport{srv: srv}
}

// SubscriptionInMethodServerTransport connects instances of azurespecialsgroup.SubscriptionInMethodClient to instances of SubscriptionInMethodServer.
// Don't use this type directly, use NewSubscriptionInMethodServerTransport instead.
type SubscriptionInMethodServerTransport struct {
	srv *SubscriptionInMethodServer
}

// Do implements the policy.Transporter interface for SubscriptionInMethodServerTransport.
func (s *SubscriptionInMethodServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "SubscriptionInMethodClient.PostMethodLocalNull":
		resp, err = s.dispatchPostMethodLocalNull(req)
	case "SubscriptionInMethodClient.PostMethodLocalValid":
		resp, err = s.dispatchPostMethodLocalValid(req)
	case "SubscriptionInMethodClient.PostPathLocalValid":
		resp, err = s.dispatchPostPathLocalValid(req)
	case "SubscriptionInMethodClient.PostSwaggerLocalValid":
		resp, err = s.dispatchPostSwaggerLocalValid(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *SubscriptionInMethodServerTransport) dispatchPostMethodLocalNull(req *http.Request) (*http.Response, error) {
	if s.srv.PostMethodLocalNull == nil {
		return nil, &nonRetriableError{errors.New("method PostMethodLocalNull not implemented")}
	}
	const regexStr = "/azurespecials/subscriptionId/method/string/none/path/local/null/(?P<subscriptionId>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := s.srv.PostMethodLocalNull(req.Context(), matches[regex.SubexpIndex("subscriptionId")], nil)
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

func (s *SubscriptionInMethodServerTransport) dispatchPostMethodLocalValid(req *http.Request) (*http.Response, error) {
	if s.srv.PostMethodLocalValid == nil {
		return nil, &nonRetriableError{errors.New("method PostMethodLocalValid not implemented")}
	}
	const regexStr = "/azurespecials/subscriptionId/method/string/none/path/local/1234-5678-9012-3456/(?P<subscriptionId>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := s.srv.PostMethodLocalValid(req.Context(), matches[regex.SubexpIndex("subscriptionId")], nil)
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

func (s *SubscriptionInMethodServerTransport) dispatchPostPathLocalValid(req *http.Request) (*http.Response, error) {
	if s.srv.PostPathLocalValid == nil {
		return nil, &nonRetriableError{errors.New("method PostPathLocalValid not implemented")}
	}
	const regexStr = "/azurespecials/subscriptionId/path/string/none/path/local/1234-5678-9012-3456/(?P<subscriptionId>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := s.srv.PostPathLocalValid(req.Context(), matches[regex.SubexpIndex("subscriptionId")], nil)
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

func (s *SubscriptionInMethodServerTransport) dispatchPostSwaggerLocalValid(req *http.Request) (*http.Response, error) {
	if s.srv.PostSwaggerLocalValid == nil {
		return nil, &nonRetriableError{errors.New("method PostSwaggerLocalValid not implemented")}
	}
	const regexStr = "/azurespecials/subscriptionId/swagger/string/none/path/local/1234-5678-9012-3456/(?P<subscriptionId>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := s.srv.PostSwaggerLocalValid(req.Context(), matches[regex.SubexpIndex("subscriptionId")], nil)
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
