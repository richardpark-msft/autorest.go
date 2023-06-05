//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"azartifacts"
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"net/http"
	"regexp"
)

// SQLScriptServer is a fake server for instances of the azartifacts.SQLScriptClient type.
type SQLScriptServer struct {
	// BeginCreateOrUpdateSQLScript is the fake for method SQLScriptClient.BeginCreateOrUpdateSQLScript
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginCreateOrUpdateSQLScript func(ctx context.Context, sqlScriptName string, sqlScript azartifacts.SQLScriptResource, options *azartifacts.SQLScriptClientBeginCreateOrUpdateSQLScriptOptions) (resp azfake.PollerResponder[azartifacts.SQLScriptClientCreateOrUpdateSQLScriptResponse], errResp azfake.ErrorResponder)

	// BeginDeleteSQLScript is the fake for method SQLScriptClient.BeginDeleteSQLScript
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDeleteSQLScript func(ctx context.Context, sqlScriptName string, options *azartifacts.SQLScriptClientBeginDeleteSQLScriptOptions) (resp azfake.PollerResponder[azartifacts.SQLScriptClientDeleteSQLScriptResponse], errResp azfake.ErrorResponder)

	// GetSQLScript is the fake for method SQLScriptClient.GetSQLScript
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNotModified
	GetSQLScript func(ctx context.Context, sqlScriptName string, options *azartifacts.SQLScriptClientGetSQLScriptOptions) (resp azfake.Responder[azartifacts.SQLScriptClientGetSQLScriptResponse], errResp azfake.ErrorResponder)

	// NewGetSQLScriptsByWorkspacePager is the fake for method SQLScriptClient.NewGetSQLScriptsByWorkspacePager
	// HTTP status codes to indicate success: http.StatusOK
	NewGetSQLScriptsByWorkspacePager func(options *azartifacts.SQLScriptClientGetSQLScriptsByWorkspaceOptions) (resp azfake.PagerResponder[azartifacts.SQLScriptClientGetSQLScriptsByWorkspaceResponse])

	// BeginRenameSQLScript is the fake for method SQLScriptClient.BeginRenameSQLScript
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginRenameSQLScript func(ctx context.Context, sqlScriptName string, request azartifacts.ArtifactRenameRequest, options *azartifacts.SQLScriptClientBeginRenameSQLScriptOptions) (resp azfake.PollerResponder[azartifacts.SQLScriptClientRenameSQLScriptResponse], errResp azfake.ErrorResponder)
}

// NewSQLScriptServerTransport creates a new instance of SQLScriptServerTransport with the provided implementation.
// The returned SQLScriptServerTransport instance is connected to an instance of azartifacts.SQLScriptClient by way of the
// undefined.Transporter field.
func NewSQLScriptServerTransport(srv *SQLScriptServer) *SQLScriptServerTransport {
	return &SQLScriptServerTransport{srv: srv}
}

// SQLScriptServerTransport connects instances of azartifacts.SQLScriptClient to instances of SQLScriptServer.
// Don't use this type directly, use NewSQLScriptServerTransport instead.
type SQLScriptServerTransport struct {
	srv                              *SQLScriptServer
	beginCreateOrUpdateSQLScript     *azfake.PollerResponder[azartifacts.SQLScriptClientCreateOrUpdateSQLScriptResponse]
	beginDeleteSQLScript             *azfake.PollerResponder[azartifacts.SQLScriptClientDeleteSQLScriptResponse]
	newGetSQLScriptsByWorkspacePager *azfake.PagerResponder[azartifacts.SQLScriptClientGetSQLScriptsByWorkspaceResponse]
	beginRenameSQLScript             *azfake.PollerResponder[azartifacts.SQLScriptClientRenameSQLScriptResponse]
}

// Do implements the policy.Transporter interface for SQLScriptServerTransport.
func (s *SQLScriptServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "SQLScriptClient.BeginCreateOrUpdateSQLScript":
		resp, err = s.dispatchBeginCreateOrUpdateSQLScript(req)
	case "SQLScriptClient.BeginDeleteSQLScript":
		resp, err = s.dispatchBeginDeleteSQLScript(req)
	case "SQLScriptClient.GetSQLScript":
		resp, err = s.dispatchGetSQLScript(req)
	case "SQLScriptClient.NewGetSQLScriptsByWorkspacePager":
		resp, err = s.dispatchNewGetSQLScriptsByWorkspacePager(req)
	case "SQLScriptClient.BeginRenameSQLScript":
		resp, err = s.dispatchBeginRenameSQLScript(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *SQLScriptServerTransport) dispatchBeginCreateOrUpdateSQLScript(req *http.Request) (*http.Response, error) {
	if s.srv.BeginCreateOrUpdateSQLScript == nil {
		return nil, &nonRetriableError{errors.New("method BeginCreateOrUpdateSQLScript not implemented")}
	}
	if s.beginCreateOrUpdateSQLScript == nil {
		const regexStr = "/sqlScripts/(?P<sqlScriptName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[azartifacts.SQLScriptResource](req)
		if err != nil {
			return nil, err
		}
		ifMatchParam := getOptional(getHeaderValue(req.Header, "If-Match"))
		var options *azartifacts.SQLScriptClientBeginCreateOrUpdateSQLScriptOptions
		if ifMatchParam != nil {
			options = &azartifacts.SQLScriptClientBeginCreateOrUpdateSQLScriptOptions{
				IfMatch: ifMatchParam,
			}
		}
		respr, errRespr := s.srv.BeginCreateOrUpdateSQLScript(req.Context(), matches[regex.SubexpIndex("sqlScriptName")], body, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		s.beginCreateOrUpdateSQLScript = &respr
	}

	resp, err := server.PollerResponderNext(s.beginCreateOrUpdateSQLScript, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(s.beginCreateOrUpdateSQLScript) {
		s.beginCreateOrUpdateSQLScript = nil
	}

	return resp, nil
}

func (s *SQLScriptServerTransport) dispatchBeginDeleteSQLScript(req *http.Request) (*http.Response, error) {
	if s.srv.BeginDeleteSQLScript == nil {
		return nil, &nonRetriableError{errors.New("method BeginDeleteSQLScript not implemented")}
	}
	if s.beginDeleteSQLScript == nil {
		const regexStr = "/sqlScripts/(?P<sqlScriptName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		respr, errRespr := s.srv.BeginDeleteSQLScript(req.Context(), matches[regex.SubexpIndex("sqlScriptName")], nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		s.beginDeleteSQLScript = &respr
	}

	resp, err := server.PollerResponderNext(s.beginDeleteSQLScript, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(s.beginDeleteSQLScript) {
		s.beginDeleteSQLScript = nil
	}

	return resp, nil
}

func (s *SQLScriptServerTransport) dispatchGetSQLScript(req *http.Request) (*http.Response, error) {
	if s.srv.GetSQLScript == nil {
		return nil, &nonRetriableError{errors.New("method GetSQLScript not implemented")}
	}
	const regexStr = "/sqlScripts/(?P<sqlScriptName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	ifNoneMatchParam := getOptional(getHeaderValue(req.Header, "If-None-Match"))
	var options *azartifacts.SQLScriptClientGetSQLScriptOptions
	if ifNoneMatchParam != nil {
		options = &azartifacts.SQLScriptClientGetSQLScriptOptions{
			IfNoneMatch: ifNoneMatchParam,
		}
	}
	respr, errRespr := s.srv.GetSQLScript(req.Context(), matches[regex.SubexpIndex("sqlScriptName")], options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNotModified}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNotModified", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SQLScriptResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SQLScriptServerTransport) dispatchNewGetSQLScriptsByWorkspacePager(req *http.Request) (*http.Response, error) {
	if s.srv.NewGetSQLScriptsByWorkspacePager == nil {
		return nil, &nonRetriableError{errors.New("method NewGetSQLScriptsByWorkspacePager not implemented")}
	}
	if s.newGetSQLScriptsByWorkspacePager == nil {
		resp := s.srv.NewGetSQLScriptsByWorkspacePager(nil)
		s.newGetSQLScriptsByWorkspacePager = &resp
		server.PagerResponderInjectNextLinks(s.newGetSQLScriptsByWorkspacePager, req, func(page *azartifacts.SQLScriptClientGetSQLScriptsByWorkspaceResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(s.newGetSQLScriptsByWorkspacePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(s.newGetSQLScriptsByWorkspacePager) {
		s.newGetSQLScriptsByWorkspacePager = nil
	}
	return resp, nil
}

func (s *SQLScriptServerTransport) dispatchBeginRenameSQLScript(req *http.Request) (*http.Response, error) {
	if s.srv.BeginRenameSQLScript == nil {
		return nil, &nonRetriableError{errors.New("method BeginRenameSQLScript not implemented")}
	}
	if s.beginRenameSQLScript == nil {
		const regexStr = "/sqlScripts/(?P<sqlScriptName>[a-zA-Z0-9-_]+)/rename"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[azartifacts.ArtifactRenameRequest](req)
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginRenameSQLScript(req.Context(), matches[regex.SubexpIndex("sqlScriptName")], body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		s.beginRenameSQLScript = &respr
	}

	resp, err := server.PollerResponderNext(s.beginRenameSQLScript, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(s.beginRenameSQLScript) {
		s.beginRenameSQLScript = nil
	}

	return resp, nil
}
