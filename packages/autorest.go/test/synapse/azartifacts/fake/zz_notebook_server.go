//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

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

// NotebookServer is a fake server for instances of the azartifacts.NotebookClient type.
type NotebookServer struct {
	// BeginCreateOrUpdateNotebook is the fake for method NotebookClient.BeginCreateOrUpdateNotebook
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginCreateOrUpdateNotebook func(ctx context.Context, notebookName string, notebook azartifacts.NotebookResource, options *azartifacts.NotebookClientBeginCreateOrUpdateNotebookOptions) (resp azfake.PollerResponder[azartifacts.NotebookClientCreateOrUpdateNotebookResponse], errResp azfake.ErrorResponder)

	// BeginDeleteNotebook is the fake for method NotebookClient.BeginDeleteNotebook
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDeleteNotebook func(ctx context.Context, notebookName string, options *azartifacts.NotebookClientBeginDeleteNotebookOptions) (resp azfake.PollerResponder[azartifacts.NotebookClientDeleteNotebookResponse], errResp azfake.ErrorResponder)

	// GetNotebook is the fake for method NotebookClient.GetNotebook
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNotModified
	GetNotebook func(ctx context.Context, notebookName string, options *azartifacts.NotebookClientGetNotebookOptions) (resp azfake.Responder[azartifacts.NotebookClientGetNotebookResponse], errResp azfake.ErrorResponder)

	// NewGetNotebookSummaryByWorkSpacePager is the fake for method NotebookClient.NewGetNotebookSummaryByWorkSpacePager
	// HTTP status codes to indicate success: http.StatusOK
	NewGetNotebookSummaryByWorkSpacePager func(options *azartifacts.NotebookClientGetNotebookSummaryByWorkSpaceOptions) (resp azfake.PagerResponder[azartifacts.NotebookClientGetNotebookSummaryByWorkSpaceResponse])

	// NewGetNotebooksByWorkspacePager is the fake for method NotebookClient.NewGetNotebooksByWorkspacePager
	// HTTP status codes to indicate success: http.StatusOK
	NewGetNotebooksByWorkspacePager func(options *azartifacts.NotebookClientGetNotebooksByWorkspaceOptions) (resp azfake.PagerResponder[azartifacts.NotebookClientGetNotebooksByWorkspaceResponse])

	// BeginRenameNotebook is the fake for method NotebookClient.BeginRenameNotebook
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginRenameNotebook func(ctx context.Context, notebookName string, request azartifacts.ArtifactRenameRequest, options *azartifacts.NotebookClientBeginRenameNotebookOptions) (resp azfake.PollerResponder[azartifacts.NotebookClientRenameNotebookResponse], errResp azfake.ErrorResponder)
}

// NewNotebookServerTransport creates a new instance of NotebookServerTransport with the provided implementation.
// The returned NotebookServerTransport instance is connected to an instance of azartifacts.NotebookClient by way of the
// undefined.Transporter field.
func NewNotebookServerTransport(srv *NotebookServer) *NotebookServerTransport {
	return &NotebookServerTransport{srv: srv}
}

// NotebookServerTransport connects instances of azartifacts.NotebookClient to instances of NotebookServer.
// Don't use this type directly, use NewNotebookServerTransport instead.
type NotebookServerTransport struct {
	srv                                   *NotebookServer
	beginCreateOrUpdateNotebook           *azfake.PollerResponder[azartifacts.NotebookClientCreateOrUpdateNotebookResponse]
	beginDeleteNotebook                   *azfake.PollerResponder[azartifacts.NotebookClientDeleteNotebookResponse]
	newGetNotebookSummaryByWorkSpacePager *azfake.PagerResponder[azartifacts.NotebookClientGetNotebookSummaryByWorkSpaceResponse]
	newGetNotebooksByWorkspacePager       *azfake.PagerResponder[azartifacts.NotebookClientGetNotebooksByWorkspaceResponse]
	beginRenameNotebook                   *azfake.PollerResponder[azartifacts.NotebookClientRenameNotebookResponse]
}

// Do implements the policy.Transporter interface for NotebookServerTransport.
func (n *NotebookServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "NotebookClient.BeginCreateOrUpdateNotebook":
		resp, err = n.dispatchBeginCreateOrUpdateNotebook(req)
	case "NotebookClient.BeginDeleteNotebook":
		resp, err = n.dispatchBeginDeleteNotebook(req)
	case "NotebookClient.GetNotebook":
		resp, err = n.dispatchGetNotebook(req)
	case "NotebookClient.NewGetNotebookSummaryByWorkSpacePager":
		resp, err = n.dispatchNewGetNotebookSummaryByWorkSpacePager(req)
	case "NotebookClient.NewGetNotebooksByWorkspacePager":
		resp, err = n.dispatchNewGetNotebooksByWorkspacePager(req)
	case "NotebookClient.BeginRenameNotebook":
		resp, err = n.dispatchBeginRenameNotebook(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (n *NotebookServerTransport) dispatchBeginCreateOrUpdateNotebook(req *http.Request) (*http.Response, error) {
	if n.srv.BeginCreateOrUpdateNotebook == nil {
		return nil, &nonRetriableError{errors.New("method BeginCreateOrUpdateNotebook not implemented")}
	}
	if n.beginCreateOrUpdateNotebook == nil {
		const regexStr = "/notebooks/(?P<notebookName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[azartifacts.NotebookResource](req)
		if err != nil {
			return nil, err
		}
		ifMatchParam := getOptional(getHeaderValue(req.Header, "If-Match"))
		var options *azartifacts.NotebookClientBeginCreateOrUpdateNotebookOptions
		if ifMatchParam != nil {
			options = &azartifacts.NotebookClientBeginCreateOrUpdateNotebookOptions{
				IfMatch: ifMatchParam,
			}
		}
		respr, errRespr := n.srv.BeginCreateOrUpdateNotebook(req.Context(), matches[regex.SubexpIndex("notebookName")], body, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		n.beginCreateOrUpdateNotebook = &respr
	}

	resp, err := server.PollerResponderNext(n.beginCreateOrUpdateNotebook, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(n.beginCreateOrUpdateNotebook) {
		n.beginCreateOrUpdateNotebook = nil
	}

	return resp, nil
}

func (n *NotebookServerTransport) dispatchBeginDeleteNotebook(req *http.Request) (*http.Response, error) {
	if n.srv.BeginDeleteNotebook == nil {
		return nil, &nonRetriableError{errors.New("method BeginDeleteNotebook not implemented")}
	}
	if n.beginDeleteNotebook == nil {
		const regexStr = "/notebooks/(?P<notebookName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		respr, errRespr := n.srv.BeginDeleteNotebook(req.Context(), matches[regex.SubexpIndex("notebookName")], nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		n.beginDeleteNotebook = &respr
	}

	resp, err := server.PollerResponderNext(n.beginDeleteNotebook, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(n.beginDeleteNotebook) {
		n.beginDeleteNotebook = nil
	}

	return resp, nil
}

func (n *NotebookServerTransport) dispatchGetNotebook(req *http.Request) (*http.Response, error) {
	if n.srv.GetNotebook == nil {
		return nil, &nonRetriableError{errors.New("method GetNotebook not implemented")}
	}
	const regexStr = "/notebooks/(?P<notebookName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	ifNoneMatchParam := getOptional(getHeaderValue(req.Header, "If-None-Match"))
	var options *azartifacts.NotebookClientGetNotebookOptions
	if ifNoneMatchParam != nil {
		options = &azartifacts.NotebookClientGetNotebookOptions{
			IfNoneMatch: ifNoneMatchParam,
		}
	}
	respr, errRespr := n.srv.GetNotebook(req.Context(), matches[regex.SubexpIndex("notebookName")], options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNotModified}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNotModified", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).NotebookResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (n *NotebookServerTransport) dispatchNewGetNotebookSummaryByWorkSpacePager(req *http.Request) (*http.Response, error) {
	if n.srv.NewGetNotebookSummaryByWorkSpacePager == nil {
		return nil, &nonRetriableError{errors.New("method NewGetNotebookSummaryByWorkSpacePager not implemented")}
	}
	if n.newGetNotebookSummaryByWorkSpacePager == nil {
		resp := n.srv.NewGetNotebookSummaryByWorkSpacePager(nil)
		n.newGetNotebookSummaryByWorkSpacePager = &resp
		server.PagerResponderInjectNextLinks(n.newGetNotebookSummaryByWorkSpacePager, req, func(page *azartifacts.NotebookClientGetNotebookSummaryByWorkSpaceResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(n.newGetNotebookSummaryByWorkSpacePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(n.newGetNotebookSummaryByWorkSpacePager) {
		n.newGetNotebookSummaryByWorkSpacePager = nil
	}
	return resp, nil
}

func (n *NotebookServerTransport) dispatchNewGetNotebooksByWorkspacePager(req *http.Request) (*http.Response, error) {
	if n.srv.NewGetNotebooksByWorkspacePager == nil {
		return nil, &nonRetriableError{errors.New("method NewGetNotebooksByWorkspacePager not implemented")}
	}
	if n.newGetNotebooksByWorkspacePager == nil {
		resp := n.srv.NewGetNotebooksByWorkspacePager(nil)
		n.newGetNotebooksByWorkspacePager = &resp
		server.PagerResponderInjectNextLinks(n.newGetNotebooksByWorkspacePager, req, func(page *azartifacts.NotebookClientGetNotebooksByWorkspaceResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(n.newGetNotebooksByWorkspacePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(n.newGetNotebooksByWorkspacePager) {
		n.newGetNotebooksByWorkspacePager = nil
	}
	return resp, nil
}

func (n *NotebookServerTransport) dispatchBeginRenameNotebook(req *http.Request) (*http.Response, error) {
	if n.srv.BeginRenameNotebook == nil {
		return nil, &nonRetriableError{errors.New("method BeginRenameNotebook not implemented")}
	}
	if n.beginRenameNotebook == nil {
		const regexStr = "/notebooks/(?P<notebookName>[a-zA-Z0-9-_]+)/rename"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[azartifacts.ArtifactRenameRequest](req)
		if err != nil {
			return nil, err
		}
		respr, errRespr := n.srv.BeginRenameNotebook(req.Context(), matches[regex.SubexpIndex("notebookName")], body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		n.beginRenameNotebook = &respr
	}

	resp, err := server.PollerResponderNext(n.beginRenameNotebook, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(n.beginRenameNotebook) {
		n.beginRenameNotebook = nil
	}

	return resp, nil
}