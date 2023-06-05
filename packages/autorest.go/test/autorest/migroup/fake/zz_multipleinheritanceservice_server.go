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
	"generatortests/migroup"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// MultipleInheritanceServiceServer is a fake server for instances of the migroup.MultipleInheritanceServiceClient type.
type MultipleInheritanceServiceServer struct {
	// GetCat is the fake for method MultipleInheritanceServiceClient.GetCat
	// HTTP status codes to indicate success: http.StatusOK
	GetCat func(ctx context.Context, options *migroup.MultipleInheritanceServiceClientGetCatOptions) (resp azfake.Responder[migroup.MultipleInheritanceServiceClientGetCatResponse], errResp azfake.ErrorResponder)

	// GetFeline is the fake for method MultipleInheritanceServiceClient.GetFeline
	// HTTP status codes to indicate success: http.StatusOK
	GetFeline func(ctx context.Context, options *migroup.MultipleInheritanceServiceClientGetFelineOptions) (resp azfake.Responder[migroup.MultipleInheritanceServiceClientGetFelineResponse], errResp azfake.ErrorResponder)

	// GetHorse is the fake for method MultipleInheritanceServiceClient.GetHorse
	// HTTP status codes to indicate success: http.StatusOK
	GetHorse func(ctx context.Context, options *migroup.MultipleInheritanceServiceClientGetHorseOptions) (resp azfake.Responder[migroup.MultipleInheritanceServiceClientGetHorseResponse], errResp azfake.ErrorResponder)

	// GetKitten is the fake for method MultipleInheritanceServiceClient.GetKitten
	// HTTP status codes to indicate success: http.StatusOK
	GetKitten func(ctx context.Context, options *migroup.MultipleInheritanceServiceClientGetKittenOptions) (resp azfake.Responder[migroup.MultipleInheritanceServiceClientGetKittenResponse], errResp azfake.ErrorResponder)

	// GetPet is the fake for method MultipleInheritanceServiceClient.GetPet
	// HTTP status codes to indicate success: http.StatusOK
	GetPet func(ctx context.Context, options *migroup.MultipleInheritanceServiceClientGetPetOptions) (resp azfake.Responder[migroup.MultipleInheritanceServiceClientGetPetResponse], errResp azfake.ErrorResponder)

	// PutCat is the fake for method MultipleInheritanceServiceClient.PutCat
	// HTTP status codes to indicate success: http.StatusOK
	PutCat func(ctx context.Context, cat migroup.Cat, options *migroup.MultipleInheritanceServiceClientPutCatOptions) (resp azfake.Responder[migroup.MultipleInheritanceServiceClientPutCatResponse], errResp azfake.ErrorResponder)

	// PutFeline is the fake for method MultipleInheritanceServiceClient.PutFeline
	// HTTP status codes to indicate success: http.StatusOK
	PutFeline func(ctx context.Context, feline migroup.Feline, options *migroup.MultipleInheritanceServiceClientPutFelineOptions) (resp azfake.Responder[migroup.MultipleInheritanceServiceClientPutFelineResponse], errResp azfake.ErrorResponder)

	// PutHorse is the fake for method MultipleInheritanceServiceClient.PutHorse
	// HTTP status codes to indicate success: http.StatusOK
	PutHorse func(ctx context.Context, horse migroup.Horse, options *migroup.MultipleInheritanceServiceClientPutHorseOptions) (resp azfake.Responder[migroup.MultipleInheritanceServiceClientPutHorseResponse], errResp azfake.ErrorResponder)

	// PutKitten is the fake for method MultipleInheritanceServiceClient.PutKitten
	// HTTP status codes to indicate success: http.StatusOK
	PutKitten func(ctx context.Context, kitten migroup.Kitten, options *migroup.MultipleInheritanceServiceClientPutKittenOptions) (resp azfake.Responder[migroup.MultipleInheritanceServiceClientPutKittenResponse], errResp azfake.ErrorResponder)

	// PutPet is the fake for method MultipleInheritanceServiceClient.PutPet
	// HTTP status codes to indicate success: http.StatusOK
	PutPet func(ctx context.Context, pet migroup.Pet, options *migroup.MultipleInheritanceServiceClientPutPetOptions) (resp azfake.Responder[migroup.MultipleInheritanceServiceClientPutPetResponse], errResp azfake.ErrorResponder)
}

// NewMultipleInheritanceServiceServerTransport creates a new instance of MultipleInheritanceServiceServerTransport with the provided implementation.
// The returned MultipleInheritanceServiceServerTransport instance is connected to an instance of migroup.MultipleInheritanceServiceClient by way of the
// undefined.Transporter field.
func NewMultipleInheritanceServiceServerTransport(srv *MultipleInheritanceServiceServer) *MultipleInheritanceServiceServerTransport {
	return &MultipleInheritanceServiceServerTransport{srv: srv}
}

// MultipleInheritanceServiceServerTransport connects instances of migroup.MultipleInheritanceServiceClient to instances of MultipleInheritanceServiceServer.
// Don't use this type directly, use NewMultipleInheritanceServiceServerTransport instead.
type MultipleInheritanceServiceServerTransport struct {
	srv *MultipleInheritanceServiceServer
}

// Do implements the policy.Transporter interface for MultipleInheritanceServiceServerTransport.
func (m *MultipleInheritanceServiceServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "MultipleInheritanceServiceClient.GetCat":
		resp, err = m.dispatchGetCat(req)
	case "MultipleInheritanceServiceClient.GetFeline":
		resp, err = m.dispatchGetFeline(req)
	case "MultipleInheritanceServiceClient.GetHorse":
		resp, err = m.dispatchGetHorse(req)
	case "MultipleInheritanceServiceClient.GetKitten":
		resp, err = m.dispatchGetKitten(req)
	case "MultipleInheritanceServiceClient.GetPet":
		resp, err = m.dispatchGetPet(req)
	case "MultipleInheritanceServiceClient.PutCat":
		resp, err = m.dispatchPutCat(req)
	case "MultipleInheritanceServiceClient.PutFeline":
		resp, err = m.dispatchPutFeline(req)
	case "MultipleInheritanceServiceClient.PutHorse":
		resp, err = m.dispatchPutHorse(req)
	case "MultipleInheritanceServiceClient.PutKitten":
		resp, err = m.dispatchPutKitten(req)
	case "MultipleInheritanceServiceClient.PutPet":
		resp, err = m.dispatchPutPet(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (m *MultipleInheritanceServiceServerTransport) dispatchGetCat(req *http.Request) (*http.Response, error) {
	if m.srv.GetCat == nil {
		return nil, &nonRetriableError{errors.New("method GetCat not implemented")}
	}
	respr, errRespr := m.srv.GetCat(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Cat, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *MultipleInheritanceServiceServerTransport) dispatchGetFeline(req *http.Request) (*http.Response, error) {
	if m.srv.GetFeline == nil {
		return nil, &nonRetriableError{errors.New("method GetFeline not implemented")}
	}
	respr, errRespr := m.srv.GetFeline(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Feline, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *MultipleInheritanceServiceServerTransport) dispatchGetHorse(req *http.Request) (*http.Response, error) {
	if m.srv.GetHorse == nil {
		return nil, &nonRetriableError{errors.New("method GetHorse not implemented")}
	}
	respr, errRespr := m.srv.GetHorse(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Horse, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *MultipleInheritanceServiceServerTransport) dispatchGetKitten(req *http.Request) (*http.Response, error) {
	if m.srv.GetKitten == nil {
		return nil, &nonRetriableError{errors.New("method GetKitten not implemented")}
	}
	respr, errRespr := m.srv.GetKitten(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Kitten, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *MultipleInheritanceServiceServerTransport) dispatchGetPet(req *http.Request) (*http.Response, error) {
	if m.srv.GetPet == nil {
		return nil, &nonRetriableError{errors.New("method GetPet not implemented")}
	}
	respr, errRespr := m.srv.GetPet(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Pet, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *MultipleInheritanceServiceServerTransport) dispatchPutCat(req *http.Request) (*http.Response, error) {
	if m.srv.PutCat == nil {
		return nil, &nonRetriableError{errors.New("method PutCat not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[migroup.Cat](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := m.srv.PutCat(req.Context(), body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Value, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *MultipleInheritanceServiceServerTransport) dispatchPutFeline(req *http.Request) (*http.Response, error) {
	if m.srv.PutFeline == nil {
		return nil, &nonRetriableError{errors.New("method PutFeline not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[migroup.Feline](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := m.srv.PutFeline(req.Context(), body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Value, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *MultipleInheritanceServiceServerTransport) dispatchPutHorse(req *http.Request) (*http.Response, error) {
	if m.srv.PutHorse == nil {
		return nil, &nonRetriableError{errors.New("method PutHorse not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[migroup.Horse](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := m.srv.PutHorse(req.Context(), body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Value, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *MultipleInheritanceServiceServerTransport) dispatchPutKitten(req *http.Request) (*http.Response, error) {
	if m.srv.PutKitten == nil {
		return nil, &nonRetriableError{errors.New("method PutKitten not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[migroup.Kitten](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := m.srv.PutKitten(req.Context(), body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Value, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *MultipleInheritanceServiceServerTransport) dispatchPutPet(req *http.Request) (*http.Response, error) {
	if m.srv.PutPet == nil {
		return nil, &nonRetriableError{errors.New("method PutPet not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[migroup.Pet](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := m.srv.PutPet(req.Context(), body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Value, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
