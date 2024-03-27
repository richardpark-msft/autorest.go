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
	"specialwordsgroup"
)

// ModelPropertiesServer is a fake server for instances of the specialwordsgroup.ModelPropertiesClient type.
type ModelPropertiesServer struct {
	// SameAsModel is the fake for method ModelPropertiesClient.SameAsModel
	// HTTP status codes to indicate success: http.StatusNoContent
	SameAsModel func(ctx context.Context, body specialwordsgroup.SameAsModel, options *specialwordsgroup.ModelPropertiesClientSameAsModelOptions) (resp azfake.Responder[specialwordsgroup.ModelPropertiesClientSameAsModelResponse], errResp azfake.ErrorResponder)
}

// NewModelPropertiesServerTransport creates a new instance of ModelPropertiesServerTransport with the provided implementation.
// The returned ModelPropertiesServerTransport instance is connected to an instance of specialwordsgroup.ModelPropertiesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewModelPropertiesServerTransport(srv *ModelPropertiesServer) *ModelPropertiesServerTransport {
	return &ModelPropertiesServerTransport{srv: srv}
}

// ModelPropertiesServerTransport connects instances of specialwordsgroup.ModelPropertiesClient to instances of ModelPropertiesServer.
// Don't use this type directly, use NewModelPropertiesServerTransport instead.
type ModelPropertiesServerTransport struct {
	srv *ModelPropertiesServer
}

// Do implements the policy.Transporter interface for ModelPropertiesServerTransport.
func (m *ModelPropertiesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return m.dispatchToMethodFake(req, method)
}

func (m *ModelPropertiesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch method {
	case "ModelPropertiesClient.SameAsModel":
		resp, err = m.dispatchSameAsModel(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	return resp, err
}

func (m *ModelPropertiesServerTransport) dispatchSameAsModel(req *http.Request) (*http.Response, error) {
	if m.srv.SameAsModel == nil {
		return nil, &nonRetriableError{errors.New("fake for method SameAsModel not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[specialwordsgroup.SameAsModel](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := m.srv.SameAsModel(req.Context(), body, nil)
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
