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

// DatetimeServer is a fake server for instances of the datetimegroup.DatetimeClient type.
type DatetimeServer struct {
	// HeaderServer contains the fakes for client HeaderClient
	HeaderServer HeaderServer

	// PropertyServer contains the fakes for client PropertyClient
	PropertyServer PropertyServer

	// QueryServer contains the fakes for client QueryClient
	QueryServer QueryServer

	// ResponseHeaderServer contains the fakes for client ResponseHeaderClient
	ResponseHeaderServer ResponseHeaderServer
}

// NewDatetimeServerTransport creates a new instance of DatetimeServerTransport with the provided implementation.
// The returned DatetimeServerTransport instance is connected to an instance of datetimegroup.DatetimeClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewDatetimeServerTransport(srv *DatetimeServer) *DatetimeServerTransport {
	return &DatetimeServerTransport{srv: srv}
}

// DatetimeServerTransport connects instances of datetimegroup.DatetimeClient to instances of DatetimeServer.
// Don't use this type directly, use NewDatetimeServerTransport instead.
type DatetimeServerTransport struct {
	srv                    *DatetimeServer
	trMu                   sync.Mutex
	trHeaderServer         *HeaderServerTransport
	trPropertyServer       *PropertyServerTransport
	trQueryServer          *QueryServerTransport
	trResponseHeaderServer *ResponseHeaderServerTransport
}

// Do implements the policy.Transporter interface for DatetimeServerTransport.
func (d *DatetimeServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return d.dispatchToClientFake(req, method[:strings.Index(method, ".")])
}

func (d *DatetimeServerTransport) dispatchToClientFake(req *http.Request, client string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch client {
	case "HeaderClient":
		initServer(&d.trMu, &d.trHeaderServer, func() *HeaderServerTransport {
			return NewHeaderServerTransport(&d.srv.HeaderServer)
		})
		resp, err = d.trHeaderServer.Do(req)
	case "PropertyClient":
		initServer(&d.trMu, &d.trPropertyServer, func() *PropertyServerTransport {
			return NewPropertyServerTransport(&d.srv.PropertyServer)
		})
		resp, err = d.trPropertyServer.Do(req)
	case "QueryClient":
		initServer(&d.trMu, &d.trQueryServer, func() *QueryServerTransport {
			return NewQueryServerTransport(&d.srv.QueryServer)
		})
		resp, err = d.trQueryServer.Do(req)
	case "ResponseHeaderClient":
		initServer(&d.trMu, &d.trResponseHeaderServer, func() *ResponseHeaderServerTransport {
			return NewResponseHeaderServerTransport(&d.srv.ResponseHeaderServer)
		})
		resp, err = d.trResponseHeaderServer.Do(req)
	default:
		err = fmt.Errorf("unhandled client %s", client)
	}

	return resp, err
}
