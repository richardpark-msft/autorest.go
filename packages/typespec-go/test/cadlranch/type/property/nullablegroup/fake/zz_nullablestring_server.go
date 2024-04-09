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
	"nullablegroup"
)

// NullableStringServer is a fake server for instances of the nullablegroup.NullableStringClient type.
type NullableStringServer struct {
	// GetNonNull is the fake for method NullableStringClient.GetNonNull
	// HTTP status codes to indicate success: http.StatusOK
	GetNonNull func(ctx context.Context, options *nullablegroup.NullableStringClientGetNonNullOptions) (resp azfake.Responder[nullablegroup.NullableStringClientGetNonNullResponse], errResp azfake.ErrorResponder)

	// GetNull is the fake for method NullableStringClient.GetNull
	// HTTP status codes to indicate success: http.StatusOK
	GetNull func(ctx context.Context, options *nullablegroup.NullableStringClientGetNullOptions) (resp azfake.Responder[nullablegroup.NullableStringClientGetNullResponse], errResp azfake.ErrorResponder)

	// PatchNonNull is the fake for method NullableStringClient.PatchNonNull
	// HTTP status codes to indicate success: http.StatusNoContent
	PatchNonNull func(ctx context.Context, body nullablegroup.StringProperty, options *nullablegroup.NullableStringClientPatchNonNullOptions) (resp azfake.Responder[nullablegroup.NullableStringClientPatchNonNullResponse], errResp azfake.ErrorResponder)

	// PatchNull is the fake for method NullableStringClient.PatchNull
	// HTTP status codes to indicate success: http.StatusNoContent
	PatchNull func(ctx context.Context, body nullablegroup.StringProperty, options *nullablegroup.NullableStringClientPatchNullOptions) (resp azfake.Responder[nullablegroup.NullableStringClientPatchNullResponse], errResp azfake.ErrorResponder)
}

// NewNullableStringServerTransport creates a new instance of NullableStringServerTransport with the provided implementation.
// The returned NullableStringServerTransport instance is connected to an instance of nullablegroup.NullableStringClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewNullableStringServerTransport(srv *NullableStringServer) *NullableStringServerTransport {
	return &NullableStringServerTransport{srv: srv}
}

// NullableStringServerTransport connects instances of nullablegroup.NullableStringClient to instances of NullableStringServer.
// Don't use this type directly, use NewNullableStringServerTransport instead.
type NullableStringServerTransport struct {
	srv *NullableStringServer
}

// Do implements the policy.Transporter interface for NullableStringServerTransport.
func (n *NullableStringServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return n.dispatchToMethodFake(req, method)
}

func (n *NullableStringServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch method {
	case "NullableStringClient.GetNonNull":
		resp, err = n.dispatchGetNonNull(req)
	case "NullableStringClient.GetNull":
		resp, err = n.dispatchGetNull(req)
	case "NullableStringClient.PatchNonNull":
		resp, err = n.dispatchPatchNonNull(req)
	case "NullableStringClient.PatchNull":
		resp, err = n.dispatchPatchNull(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	return resp, err
}

func (n *NullableStringServerTransport) dispatchGetNonNull(req *http.Request) (*http.Response, error) {
	if n.srv.GetNonNull == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetNonNull not implemented")}
	}
	respr, errRespr := n.srv.GetNonNull(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).StringProperty, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (n *NullableStringServerTransport) dispatchGetNull(req *http.Request) (*http.Response, error) {
	if n.srv.GetNull == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetNull not implemented")}
	}
	respr, errRespr := n.srv.GetNull(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).StringProperty, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (n *NullableStringServerTransport) dispatchPatchNonNull(req *http.Request) (*http.Response, error) {
	if n.srv.PatchNonNull == nil {
		return nil, &nonRetriableError{errors.New("fake for method PatchNonNull not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[nullablegroup.StringProperty](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := n.srv.PatchNonNull(req.Context(), body, nil)
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

func (n *NullableStringServerTransport) dispatchPatchNull(req *http.Request) (*http.Response, error) {
	if n.srv.PatchNull == nil {
		return nil, &nonRetriableError{errors.New("fake for method PatchNull not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[nullablegroup.StringProperty](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := n.srv.PatchNull(req.Context(), body, nil)
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