// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"context"
	"defaultgroup"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
	"sync"
)

// ServiceServer is a fake server for instances of the defaultgroup.ServiceClient type.
type ServiceServer struct {
	// BarServer contains the fakes for client BarClient
	BarServer BarServer

	// FooServer contains the fakes for client FooClient
	FooServer FooServer

	// QuxServer contains the fakes for client QuxClient
	QuxServer QuxServer

	// One is the fake for method ServiceClient.One
	// HTTP status codes to indicate success: http.StatusNoContent
	One func(ctx context.Context, options *defaultgroup.ServiceClientOneOptions) (resp azfake.Responder[defaultgroup.ServiceClientOneResponse], errResp azfake.ErrorResponder)

	// Two is the fake for method ServiceClient.Two
	// HTTP status codes to indicate success: http.StatusNoContent
	Two func(ctx context.Context, options *defaultgroup.ServiceClientTwoOptions) (resp azfake.Responder[defaultgroup.ServiceClientTwoResponse], errResp azfake.ErrorResponder)
}

// NewServiceServerTransport creates a new instance of ServiceServerTransport with the provided implementation.
// The returned ServiceServerTransport instance is connected to an instance of defaultgroup.ServiceClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServiceServerTransport(srv *ServiceServer) *ServiceServerTransport {
	return &ServiceServerTransport{srv: srv}
}

// ServiceServerTransport connects instances of defaultgroup.ServiceClient to instances of ServiceServer.
// Don't use this type directly, use NewServiceServerTransport instead.
type ServiceServerTransport struct {
	srv         *ServiceServer
	trMu        sync.Mutex
	trBarServer *BarServerTransport
	trFooServer *FooServerTransport
	trQuxServer *QuxServerTransport
}

// Do implements the policy.Transporter interface for ServiceServerTransport.
func (s *ServiceServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	if client := method[:strings.Index(method, ".")]; client != "ServiceClient" {
		return s.dispatchToClientFake(req, client)
	}
	return s.dispatchToMethodFake(req, method)
}

func (s *ServiceServerTransport) dispatchToClientFake(req *http.Request, client string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch client {
	case "BarClient":
		initServer(&s.trMu, &s.trBarServer, func() *BarServerTransport {
			return NewBarServerTransport(&s.srv.BarServer)
		})
		resp, err = s.trBarServer.Do(req)
	case "FooClient":
		initServer(&s.trMu, &s.trFooServer, func() *FooServerTransport {
			return NewFooServerTransport(&s.srv.FooServer)
		})
		resp, err = s.trFooServer.Do(req)
	case "QuxClient":
		initServer(&s.trMu, &s.trQuxServer, func() *QuxServerTransport {
			return NewQuxServerTransport(&s.srv.QuxServer)
		})
		resp, err = s.trQuxServer.Do(req)
	default:
		err = fmt.Errorf("unhandled client %s", client)
	}

	return resp, err
}

func (s *ServiceServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch method {
	case "ServiceClient.One":
		resp, err = s.dispatchOne(req)
	case "ServiceClient.Two":
		resp, err = s.dispatchTwo(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	return resp, err
}

func (s *ServiceServerTransport) dispatchOne(req *http.Request) (*http.Response, error) {
	if s.srv.One == nil {
		return nil, &nonRetriableError{errors.New("fake for method One not implemented")}
	}
	respr, errRespr := s.srv.One(req.Context(), nil)
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

func (s *ServiceServerTransport) dispatchTwo(req *http.Request) (*http.Response, error) {
	if s.srv.Two == nil {
		return nil, &nonRetriableError{errors.New("fake for method Two not implemented")}
	}
	respr, errRespr := s.srv.Two(req.Context(), nil)
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
