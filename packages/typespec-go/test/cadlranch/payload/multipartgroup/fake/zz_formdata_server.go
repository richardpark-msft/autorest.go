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
	"multipartgroup"
	"net/http"
)

// FormDataServer is a fake server for instances of the multipartgroup.FormDataClient type.
type FormDataServer struct {
	// Basic is the fake for method FormDataClient.Basic
	// HTTP status codes to indicate success: http.StatusNoContent
	Basic func(ctx context.Context, body multipartgroup.MultiPartRequest, options *multipartgroup.FormDataClientBasicOptions) (resp azfake.Responder[multipartgroup.FormDataClientBasicResponse], errResp azfake.ErrorResponder)

	// BinaryArrayParts is the fake for method FormDataClient.BinaryArrayParts
	// HTTP status codes to indicate success: http.StatusNoContent
	BinaryArrayParts func(ctx context.Context, body multipartgroup.BinaryArrayPartsRequest, options *multipartgroup.FormDataClientBinaryArrayPartsOptions) (resp azfake.Responder[multipartgroup.FormDataClientBinaryArrayPartsResponse], errResp azfake.ErrorResponder)

	// CheckFileNameAndContentType is the fake for method FormDataClient.CheckFileNameAndContentType
	// HTTP status codes to indicate success: http.StatusNoContent
	CheckFileNameAndContentType func(ctx context.Context, body multipartgroup.MultiPartRequest, options *multipartgroup.FormDataClientCheckFileNameAndContentTypeOptions) (resp azfake.Responder[multipartgroup.FormDataClientCheckFileNameAndContentTypeResponse], errResp azfake.ErrorResponder)

	// Complex is the fake for method FormDataClient.Complex
	// HTTP status codes to indicate success: http.StatusNoContent
	Complex func(ctx context.Context, body multipartgroup.ComplexPartsRequest, options *multipartgroup.FormDataClientComplexOptions) (resp azfake.Responder[multipartgroup.FormDataClientComplexResponse], errResp azfake.ErrorResponder)

	// JSONArrayParts is the fake for method FormDataClient.JSONArrayParts
	// HTTP status codes to indicate success: http.StatusNoContent
	JSONArrayParts func(ctx context.Context, body multipartgroup.JSONArrayPartsRequest, options *multipartgroup.FormDataClientJSONArrayPartsOptions) (resp azfake.Responder[multipartgroup.FormDataClientJSONArrayPartsResponse], errResp azfake.ErrorResponder)

	// JSONPart is the fake for method FormDataClient.JSONPart
	// HTTP status codes to indicate success: http.StatusNoContent
	JSONPart func(ctx context.Context, body multipartgroup.JSONPartRequest, options *multipartgroup.FormDataClientJSONPartOptions) (resp azfake.Responder[multipartgroup.FormDataClientJSONPartResponse], errResp azfake.ErrorResponder)

	// MultiBinaryParts is the fake for method FormDataClient.MultiBinaryParts
	// HTTP status codes to indicate success: http.StatusNoContent
	MultiBinaryParts func(ctx context.Context, body multipartgroup.MultiBinaryPartsRequest, options *multipartgroup.FormDataClientMultiBinaryPartsOptions) (resp azfake.Responder[multipartgroup.FormDataClientMultiBinaryPartsResponse], errResp azfake.ErrorResponder)
}

// NewFormDataServerTransport creates a new instance of FormDataServerTransport with the provided implementation.
// The returned FormDataServerTransport instance is connected to an instance of multipartgroup.FormDataClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewFormDataServerTransport(srv *FormDataServer) *FormDataServerTransport {
	return &FormDataServerTransport{srv: srv}
}

// FormDataServerTransport connects instances of multipartgroup.FormDataClient to instances of FormDataServer.
// Don't use this type directly, use NewFormDataServerTransport instead.
type FormDataServerTransport struct {
	srv *FormDataServer
}

// Do implements the policy.Transporter interface for FormDataServerTransport.
func (f *FormDataServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "FormDataClient.Basic":
		resp, err = f.dispatchBasic(req)
	case "FormDataClient.BinaryArrayParts":
		resp, err = f.dispatchBinaryArrayParts(req)
	case "FormDataClient.CheckFileNameAndContentType":
		resp, err = f.dispatchCheckFileNameAndContentType(req)
	case "FormDataClient.Complex":
		resp, err = f.dispatchComplex(req)
	case "FormDataClient.JSONArrayParts":
		resp, err = f.dispatchJSONArrayParts(req)
	case "FormDataClient.JSONPart":
		resp, err = f.dispatchJSONPart(req)
	case "FormDataClient.MultiBinaryParts":
		resp, err = f.dispatchMultiBinaryParts(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (f *FormDataServerTransport) dispatchBasic(req *http.Request) (*http.Response, error) {
	if f.srv.Basic == nil {
		return nil, &nonRetriableError{errors.New("fake for method Basic not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[multipartgroup.MultiPartRequest](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := f.srv.Basic(req.Context(), body, nil)
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

func (f *FormDataServerTransport) dispatchBinaryArrayParts(req *http.Request) (*http.Response, error) {
	if f.srv.BinaryArrayParts == nil {
		return nil, &nonRetriableError{errors.New("fake for method BinaryArrayParts not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[multipartgroup.BinaryArrayPartsRequest](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := f.srv.BinaryArrayParts(req.Context(), body, nil)
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

func (f *FormDataServerTransport) dispatchCheckFileNameAndContentType(req *http.Request) (*http.Response, error) {
	if f.srv.CheckFileNameAndContentType == nil {
		return nil, &nonRetriableError{errors.New("fake for method CheckFileNameAndContentType not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[multipartgroup.MultiPartRequest](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := f.srv.CheckFileNameAndContentType(req.Context(), body, nil)
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

func (f *FormDataServerTransport) dispatchComplex(req *http.Request) (*http.Response, error) {
	if f.srv.Complex == nil {
		return nil, &nonRetriableError{errors.New("fake for method Complex not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[multipartgroup.ComplexPartsRequest](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := f.srv.Complex(req.Context(), body, nil)
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

func (f *FormDataServerTransport) dispatchJSONArrayParts(req *http.Request) (*http.Response, error) {
	if f.srv.JSONArrayParts == nil {
		return nil, &nonRetriableError{errors.New("fake for method JSONArrayParts not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[multipartgroup.JSONArrayPartsRequest](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := f.srv.JSONArrayParts(req.Context(), body, nil)
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

func (f *FormDataServerTransport) dispatchJSONPart(req *http.Request) (*http.Response, error) {
	if f.srv.JSONPart == nil {
		return nil, &nonRetriableError{errors.New("fake for method JSONPart not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[multipartgroup.JSONPartRequest](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := f.srv.JSONPart(req.Context(), body, nil)
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

func (f *FormDataServerTransport) dispatchMultiBinaryParts(req *http.Request) (*http.Response, error) {
	if f.srv.MultiBinaryParts == nil {
		return nil, &nonRetriableError{errors.New("fake for method MultiBinaryParts not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[multipartgroup.MultiBinaryPartsRequest](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := f.srv.MultiBinaryParts(req.Context(), body, nil)
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
