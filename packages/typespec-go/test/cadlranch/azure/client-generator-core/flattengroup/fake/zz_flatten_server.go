// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"context"
	"errors"
	"flattengroup"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// FlattenServer is a fake server for instances of the flattengroup.FlattenClient type.
type FlattenServer struct {
	// PutFlattenModel is the fake for method FlattenClient.PutFlattenModel
	// HTTP status codes to indicate success: http.StatusOK
	PutFlattenModel func(ctx context.Context, input flattengroup.FlattenModel, options *flattengroup.FlattenClientPutFlattenModelOptions) (resp azfake.Responder[flattengroup.FlattenClientPutFlattenModelResponse], errResp azfake.ErrorResponder)

	// PutNestedFlattenModel is the fake for method FlattenClient.PutNestedFlattenModel
	// HTTP status codes to indicate success: http.StatusOK
	PutNestedFlattenModel func(ctx context.Context, input flattengroup.NestedFlattenModel, options *flattengroup.FlattenClientPutNestedFlattenModelOptions) (resp azfake.Responder[flattengroup.FlattenClientPutNestedFlattenModelResponse], errResp azfake.ErrorResponder)
}

// NewFlattenServerTransport creates a new instance of FlattenServerTransport with the provided implementation.
// The returned FlattenServerTransport instance is connected to an instance of flattengroup.FlattenClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewFlattenServerTransport(srv *FlattenServer) *FlattenServerTransport {
	return &FlattenServerTransport{srv: srv}
}

// FlattenServerTransport connects instances of flattengroup.FlattenClient to instances of FlattenServer.
// Don't use this type directly, use NewFlattenServerTransport instead.
type FlattenServerTransport struct {
	srv *FlattenServer
}

// Do implements the policy.Transporter interface for FlattenServerTransport.
func (f *FlattenServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return f.dispatchToMethodFake(req, method)
}

func (f *FlattenServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch method {
	case "FlattenClient.PutFlattenModel":
		resp, err = f.dispatchPutFlattenModel(req)
	case "FlattenClient.PutNestedFlattenModel":
		resp, err = f.dispatchPutNestedFlattenModel(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	return resp, err
}

func (f *FlattenServerTransport) dispatchPutFlattenModel(req *http.Request) (*http.Response, error) {
	if f.srv.PutFlattenModel == nil {
		return nil, &nonRetriableError{errors.New("fake for method PutFlattenModel not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[flattengroup.FlattenModel](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := f.srv.PutFlattenModel(req.Context(), body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).FlattenModel, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (f *FlattenServerTransport) dispatchPutNestedFlattenModel(req *http.Request) (*http.Response, error) {
	if f.srv.PutNestedFlattenModel == nil {
		return nil, &nonRetriableError{errors.New("fake for method PutNestedFlattenModel not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[flattengroup.NestedFlattenModel](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := f.srv.PutNestedFlattenModel(req.Context(), body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).NestedFlattenModel, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
