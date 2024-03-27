// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
	"sync"
)

// BytesServer is a fake server for instances of the bytesgroup.BytesClient type.
type BytesServer struct {
	// HeaderServer contains the fakes for client HeaderClient
	HeaderServer HeaderServer

	// PropertyServer contains the fakes for client PropertyClient
	PropertyServer PropertyServer

	// QueryServer contains the fakes for client QueryClient
	QueryServer QueryServer

	// RequestBodyServer contains the fakes for client RequestBodyClient
	RequestBodyServer RequestBodyServer

	// ResponseBodyServer contains the fakes for client ResponseBodyClient
	ResponseBodyServer ResponseBodyServer
}

// NewBytesServerTransport creates a new instance of BytesServerTransport with the provided implementation.
// The returned BytesServerTransport instance is connected to an instance of bytesgroup.BytesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewBytesServerTransport(srv *BytesServer) *BytesServerTransport {
	return &BytesServerTransport{srv: srv}
}

// BytesServerTransport connects instances of bytesgroup.BytesClient to instances of BytesServer.
// Don't use this type directly, use NewBytesServerTransport instead.
type BytesServerTransport struct {
	srv                  *BytesServer
	trMu                 sync.Mutex
	trHeaderServer       *HeaderServerTransport
	trPropertyServer     *PropertyServerTransport
	trQueryServer        *QueryServerTransport
	trRequestBodyServer  *RequestBodyServerTransport
	trResponseBodyServer *ResponseBodyServerTransport
}

// Do implements the policy.Transporter interface for BytesServerTransport.
func (b *BytesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return b.dispatchToClientFake(req, method[:strings.Index(method, ".")])
}

func (b *BytesServerTransport) dispatchToClientFake(req *http.Request, client string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch client {
	case "HeaderClient":
		initServer(&b.trMu, &b.trHeaderServer, func() *HeaderServerTransport {
			return NewHeaderServerTransport(&b.srv.HeaderServer)
		})
		resp, err = b.trHeaderServer.Do(req)
	case "PropertyClient":
		initServer(&b.trMu, &b.trPropertyServer, func() *PropertyServerTransport {
			return NewPropertyServerTransport(&b.srv.PropertyServer)
		})
		resp, err = b.trPropertyServer.Do(req)
	case "QueryClient":
		initServer(&b.trMu, &b.trQueryServer, func() *QueryServerTransport {
			return NewQueryServerTransport(&b.srv.QueryServer)
		})
		resp, err = b.trQueryServer.Do(req)
	case "RequestBodyClient":
		initServer(&b.trMu, &b.trRequestBodyServer, func() *RequestBodyServerTransport {
			return NewRequestBodyServerTransport(&b.srv.RequestBodyServer)
		})
		resp, err = b.trRequestBodyServer.Do(req)
	case "ResponseBodyClient":
		initServer(&b.trMu, &b.trResponseBodyServer, func() *ResponseBodyServerTransport {
			return NewResponseBodyServerTransport(&b.srv.ResponseBodyServer)
		})
		resp, err = b.trResponseBodyServer.Do(req)
	default:
		err = fmt.Errorf("unhandled client %s", client)
	}

	return resp, err
}
