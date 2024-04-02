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
	"projectednamegroup"
)

// ProjectedNamePropertyServer is a fake server for instances of the projectednamegroup.ProjectedNamePropertyClient type.
type ProjectedNamePropertyServer struct {
	// Client is the fake for method ProjectedNamePropertyClient.Client
	// HTTP status codes to indicate success: http.StatusNoContent
	Client func(ctx context.Context, clientProjectedNameModel projectednamegroup.ClientProjectedNameModel, options *projectednamegroup.ProjectedNamePropertyClientClientOptions) (resp azfake.Responder[projectednamegroup.ProjectedNamePropertyClientClientResponse], errResp azfake.ErrorResponder)

	// JSON is the fake for method ProjectedNamePropertyClient.JSON
	// HTTP status codes to indicate success: http.StatusNoContent
	JSON func(ctx context.Context, jsonProjectedNameModel projectednamegroup.JSONProjectedNameModel, options *projectednamegroup.ProjectedNamePropertyClientJSONOptions) (resp azfake.Responder[projectednamegroup.ProjectedNamePropertyClientJSONResponse], errResp azfake.ErrorResponder)

	// JSONAndClient is the fake for method ProjectedNamePropertyClient.JSONAndClient
	// HTTP status codes to indicate success: http.StatusNoContent
	JSONAndClient func(ctx context.Context, jsonAndClientProjectedNameModel projectednamegroup.JSONAndClientProjectedNameModel, options *projectednamegroup.ProjectedNamePropertyClientJSONAndClientOptions) (resp azfake.Responder[projectednamegroup.ProjectedNamePropertyClientJSONAndClientResponse], errResp azfake.ErrorResponder)

	// Language is the fake for method ProjectedNamePropertyClient.Language
	// HTTP status codes to indicate success: http.StatusNoContent
	Language func(ctx context.Context, languageProjectedNameModel projectednamegroup.LanguageProjectedNameModel, options *projectednamegroup.ProjectedNamePropertyClientLanguageOptions) (resp azfake.Responder[projectednamegroup.ProjectedNamePropertyClientLanguageResponse], errResp azfake.ErrorResponder)
}

// NewProjectedNamePropertyServerTransport creates a new instance of ProjectedNamePropertyServerTransport with the provided implementation.
// The returned ProjectedNamePropertyServerTransport instance is connected to an instance of projectednamegroup.ProjectedNamePropertyClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewProjectedNamePropertyServerTransport(srv *ProjectedNamePropertyServer) *ProjectedNamePropertyServerTransport {
	return &ProjectedNamePropertyServerTransport{srv: srv}
}

// ProjectedNamePropertyServerTransport connects instances of projectednamegroup.ProjectedNamePropertyClient to instances of ProjectedNamePropertyServer.
// Don't use this type directly, use NewProjectedNamePropertyServerTransport instead.
type ProjectedNamePropertyServerTransport struct {
	srv *ProjectedNamePropertyServer
}

// Do implements the policy.Transporter interface for ProjectedNamePropertyServerTransport.
func (p *ProjectedNamePropertyServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return p.dispatchToMethodFake(req, method)
}

func (p *ProjectedNamePropertyServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch method {
	case "ProjectedNamePropertyClient.Client":
		resp, err = p.dispatchClient(req)
	case "ProjectedNamePropertyClient.JSON":
		resp, err = p.dispatchJSON(req)
	case "ProjectedNamePropertyClient.JSONAndClient":
		resp, err = p.dispatchJSONAndClient(req)
	case "ProjectedNamePropertyClient.Language":
		resp, err = p.dispatchLanguage(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	return resp, err
}

func (p *ProjectedNamePropertyServerTransport) dispatchClient(req *http.Request) (*http.Response, error) {
	if p.srv.Client == nil {
		return nil, &nonRetriableError{errors.New("fake for method Client not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[projectednamegroup.ClientProjectedNameModel](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Client(req.Context(), body, nil)
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

func (p *ProjectedNamePropertyServerTransport) dispatchJSON(req *http.Request) (*http.Response, error) {
	if p.srv.JSON == nil {
		return nil, &nonRetriableError{errors.New("fake for method JSON not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[projectednamegroup.JSONProjectedNameModel](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.JSON(req.Context(), body, nil)
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

func (p *ProjectedNamePropertyServerTransport) dispatchJSONAndClient(req *http.Request) (*http.Response, error) {
	if p.srv.JSONAndClient == nil {
		return nil, &nonRetriableError{errors.New("fake for method JSONAndClient not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[projectednamegroup.JSONAndClientProjectedNameModel](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.JSONAndClient(req.Context(), body, nil)
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

func (p *ProjectedNamePropertyServerTransport) dispatchLanguage(req *http.Request) (*http.Response, error) {
	if p.srv.Language == nil {
		return nil, &nonRetriableError{errors.New("fake for method Language not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[projectednamegroup.LanguageProjectedNameModel](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Language(req.Context(), body, nil)
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
