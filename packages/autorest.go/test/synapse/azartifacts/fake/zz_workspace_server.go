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
	"net/http"
)

// WorkspaceServer is a fake server for instances of the azartifacts.WorkspaceClient type.
type WorkspaceServer struct {
	// Get is the fake for method WorkspaceClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, options *azartifacts.WorkspaceClientGetOptions) (resp azfake.Responder[azartifacts.WorkspaceClientGetResponse], errResp azfake.ErrorResponder)
}

// NewWorkspaceServerTransport creates a new instance of WorkspaceServerTransport with the provided implementation.
// The returned WorkspaceServerTransport instance is connected to an instance of azartifacts.WorkspaceClient by way of the
// undefined.Transporter field.
func NewWorkspaceServerTransport(srv *WorkspaceServer) *WorkspaceServerTransport {
	return &WorkspaceServerTransport{srv: srv}
}

// WorkspaceServerTransport connects instances of azartifacts.WorkspaceClient to instances of WorkspaceServer.
// Don't use this type directly, use NewWorkspaceServerTransport instead.
type WorkspaceServerTransport struct {
	srv *WorkspaceServer
}

// Do implements the policy.Transporter interface for WorkspaceServerTransport.
func (w *WorkspaceServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "WorkspaceClient.Get":
		resp, err = w.dispatchGet(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (w *WorkspaceServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if w.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("method Get not implemented")}
	}
	respr, errRespr := w.srv.Get(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Workspace, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
