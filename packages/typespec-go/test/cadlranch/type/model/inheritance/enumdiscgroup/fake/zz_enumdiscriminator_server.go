// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"context"
	"enumdiscgroup"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// EnumDiscriminatorServer is a fake server for instances of the enumdiscgroup.EnumDiscriminatorClient type.
type EnumDiscriminatorServer struct {
	// GetExtensibleModel is the fake for method EnumDiscriminatorClient.GetExtensibleModel
	// HTTP status codes to indicate success: http.StatusOK
	GetExtensibleModel func(ctx context.Context, options *enumdiscgroup.EnumDiscriminatorClientGetExtensibleModelOptions) (resp azfake.Responder[enumdiscgroup.EnumDiscriminatorClientGetExtensibleModelResponse], errResp azfake.ErrorResponder)

	// GetExtensibleModelMissingDiscriminator is the fake for method EnumDiscriminatorClient.GetExtensibleModelMissingDiscriminator
	// HTTP status codes to indicate success: http.StatusOK
	GetExtensibleModelMissingDiscriminator func(ctx context.Context, options *enumdiscgroup.EnumDiscriminatorClientGetExtensibleModelMissingDiscriminatorOptions) (resp azfake.Responder[enumdiscgroup.EnumDiscriminatorClientGetExtensibleModelMissingDiscriminatorResponse], errResp azfake.ErrorResponder)

	// GetExtensibleModelWrongDiscriminator is the fake for method EnumDiscriminatorClient.GetExtensibleModelWrongDiscriminator
	// HTTP status codes to indicate success: http.StatusOK
	GetExtensibleModelWrongDiscriminator func(ctx context.Context, options *enumdiscgroup.EnumDiscriminatorClientGetExtensibleModelWrongDiscriminatorOptions) (resp azfake.Responder[enumdiscgroup.EnumDiscriminatorClientGetExtensibleModelWrongDiscriminatorResponse], errResp azfake.ErrorResponder)

	// GetFixedModel is the fake for method EnumDiscriminatorClient.GetFixedModel
	// HTTP status codes to indicate success: http.StatusOK
	GetFixedModel func(ctx context.Context, options *enumdiscgroup.EnumDiscriminatorClientGetFixedModelOptions) (resp azfake.Responder[enumdiscgroup.EnumDiscriminatorClientGetFixedModelResponse], errResp azfake.ErrorResponder)

	// GetFixedModelMissingDiscriminator is the fake for method EnumDiscriminatorClient.GetFixedModelMissingDiscriminator
	// HTTP status codes to indicate success: http.StatusOK
	GetFixedModelMissingDiscriminator func(ctx context.Context, options *enumdiscgroup.EnumDiscriminatorClientGetFixedModelMissingDiscriminatorOptions) (resp azfake.Responder[enumdiscgroup.EnumDiscriminatorClientGetFixedModelMissingDiscriminatorResponse], errResp azfake.ErrorResponder)

	// GetFixedModelWrongDiscriminator is the fake for method EnumDiscriminatorClient.GetFixedModelWrongDiscriminator
	// HTTP status codes to indicate success: http.StatusOK
	GetFixedModelWrongDiscriminator func(ctx context.Context, options *enumdiscgroup.EnumDiscriminatorClientGetFixedModelWrongDiscriminatorOptions) (resp azfake.Responder[enumdiscgroup.EnumDiscriminatorClientGetFixedModelWrongDiscriminatorResponse], errResp azfake.ErrorResponder)

	// PutExtensibleModel is the fake for method EnumDiscriminatorClient.PutExtensibleModel
	// HTTP status codes to indicate success: http.StatusNoContent
	PutExtensibleModel func(ctx context.Context, input enumdiscgroup.DogClassification, options *enumdiscgroup.EnumDiscriminatorClientPutExtensibleModelOptions) (resp azfake.Responder[enumdiscgroup.EnumDiscriminatorClientPutExtensibleModelResponse], errResp azfake.ErrorResponder)

	// PutFixedModel is the fake for method EnumDiscriminatorClient.PutFixedModel
	// HTTP status codes to indicate success: http.StatusNoContent
	PutFixedModel func(ctx context.Context, input enumdiscgroup.SnakeClassification, options *enumdiscgroup.EnumDiscriminatorClientPutFixedModelOptions) (resp azfake.Responder[enumdiscgroup.EnumDiscriminatorClientPutFixedModelResponse], errResp azfake.ErrorResponder)
}

// NewEnumDiscriminatorServerTransport creates a new instance of EnumDiscriminatorServerTransport with the provided implementation.
// The returned EnumDiscriminatorServerTransport instance is connected to an instance of enumdiscgroup.EnumDiscriminatorClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewEnumDiscriminatorServerTransport(srv *EnumDiscriminatorServer) *EnumDiscriminatorServerTransport {
	return &EnumDiscriminatorServerTransport{srv: srv}
}

// EnumDiscriminatorServerTransport connects instances of enumdiscgroup.EnumDiscriminatorClient to instances of EnumDiscriminatorServer.
// Don't use this type directly, use NewEnumDiscriminatorServerTransport instead.
type EnumDiscriminatorServerTransport struct {
	srv *EnumDiscriminatorServer
}

// Do implements the policy.Transporter interface for EnumDiscriminatorServerTransport.
func (e *EnumDiscriminatorServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return e.dispatchToMethodFake(req, method)
}

func (e *EnumDiscriminatorServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch method {
	case "EnumDiscriminatorClient.GetExtensibleModel":
		resp, err = e.dispatchGetExtensibleModel(req)
	case "EnumDiscriminatorClient.GetExtensibleModelMissingDiscriminator":
		resp, err = e.dispatchGetExtensibleModelMissingDiscriminator(req)
	case "EnumDiscriminatorClient.GetExtensibleModelWrongDiscriminator":
		resp, err = e.dispatchGetExtensibleModelWrongDiscriminator(req)
	case "EnumDiscriminatorClient.GetFixedModel":
		resp, err = e.dispatchGetFixedModel(req)
	case "EnumDiscriminatorClient.GetFixedModelMissingDiscriminator":
		resp, err = e.dispatchGetFixedModelMissingDiscriminator(req)
	case "EnumDiscriminatorClient.GetFixedModelWrongDiscriminator":
		resp, err = e.dispatchGetFixedModelWrongDiscriminator(req)
	case "EnumDiscriminatorClient.PutExtensibleModel":
		resp, err = e.dispatchPutExtensibleModel(req)
	case "EnumDiscriminatorClient.PutFixedModel":
		resp, err = e.dispatchPutFixedModel(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	return resp, err
}

func (e *EnumDiscriminatorServerTransport) dispatchGetExtensibleModel(req *http.Request) (*http.Response, error) {
	if e.srv.GetExtensibleModel == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetExtensibleModel not implemented")}
	}
	respr, errRespr := e.srv.GetExtensibleModel(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DogClassification, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *EnumDiscriminatorServerTransport) dispatchGetExtensibleModelMissingDiscriminator(req *http.Request) (*http.Response, error) {
	if e.srv.GetExtensibleModelMissingDiscriminator == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetExtensibleModelMissingDiscriminator not implemented")}
	}
	respr, errRespr := e.srv.GetExtensibleModelMissingDiscriminator(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DogClassification, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *EnumDiscriminatorServerTransport) dispatchGetExtensibleModelWrongDiscriminator(req *http.Request) (*http.Response, error) {
	if e.srv.GetExtensibleModelWrongDiscriminator == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetExtensibleModelWrongDiscriminator not implemented")}
	}
	respr, errRespr := e.srv.GetExtensibleModelWrongDiscriminator(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DogClassification, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *EnumDiscriminatorServerTransport) dispatchGetFixedModel(req *http.Request) (*http.Response, error) {
	if e.srv.GetFixedModel == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetFixedModel not implemented")}
	}
	respr, errRespr := e.srv.GetFixedModel(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SnakeClassification, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *EnumDiscriminatorServerTransport) dispatchGetFixedModelMissingDiscriminator(req *http.Request) (*http.Response, error) {
	if e.srv.GetFixedModelMissingDiscriminator == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetFixedModelMissingDiscriminator not implemented")}
	}
	respr, errRespr := e.srv.GetFixedModelMissingDiscriminator(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SnakeClassification, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *EnumDiscriminatorServerTransport) dispatchGetFixedModelWrongDiscriminator(req *http.Request) (*http.Response, error) {
	if e.srv.GetFixedModelWrongDiscriminator == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetFixedModelWrongDiscriminator not implemented")}
	}
	respr, errRespr := e.srv.GetFixedModelWrongDiscriminator(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SnakeClassification, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *EnumDiscriminatorServerTransport) dispatchPutExtensibleModel(req *http.Request) (*http.Response, error) {
	if e.srv.PutExtensibleModel == nil {
		return nil, &nonRetriableError{errors.New("fake for method PutExtensibleModel not implemented")}
	}
	raw, err := readRequestBody(req)
	if err != nil {
		return nil, err
	}
	body, err := unmarshalDogClassification(raw)
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.PutExtensibleModel(req.Context(), body, nil)
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

func (e *EnumDiscriminatorServerTransport) dispatchPutFixedModel(req *http.Request) (*http.Response, error) {
	if e.srv.PutFixedModel == nil {
		return nil, &nonRetriableError{errors.New("fake for method PutFixedModel not implemented")}
	}
	raw, err := readRequestBody(req)
	if err != nil {
		return nil, err
	}
	body, err := unmarshalSnakeClassification(raw)
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.PutFixedModel(req.Context(), body, nil)
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
