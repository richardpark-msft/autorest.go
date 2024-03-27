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

// CollectionFormatServer is a fake server for instances of the collectionfmtgroup.CollectionFormatClient type.
type CollectionFormatServer struct {
	// HeaderServer contains the fakes for client HeaderClient
	HeaderServer HeaderServer

	// QueryServer contains the fakes for client QueryClient
	QueryServer QueryServer
}

// NewCollectionFormatServerTransport creates a new instance of CollectionFormatServerTransport with the provided implementation.
// The returned CollectionFormatServerTransport instance is connected to an instance of collectionfmtgroup.CollectionFormatClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewCollectionFormatServerTransport(srv *CollectionFormatServer) *CollectionFormatServerTransport {
	return &CollectionFormatServerTransport{srv: srv}
}

// CollectionFormatServerTransport connects instances of collectionfmtgroup.CollectionFormatClient to instances of CollectionFormatServer.
// Don't use this type directly, use NewCollectionFormatServerTransport instead.
type CollectionFormatServerTransport struct {
	srv            *CollectionFormatServer
	trMu           sync.Mutex
	trHeaderServer *HeaderServerTransport
	trQueryServer  *QueryServerTransport
}

// Do implements the policy.Transporter interface for CollectionFormatServerTransport.
func (c *CollectionFormatServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return c.dispatchToClientFake(req, method[:strings.Index(method, ".")])
}

func (c *CollectionFormatServerTransport) dispatchToClientFake(req *http.Request, client string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch client {
	case "HeaderClient":
		initServer(&c.trMu, &c.trHeaderServer, func() *HeaderServerTransport {
			return NewHeaderServerTransport(&c.srv.HeaderServer)
		})
		resp, err = c.trHeaderServer.Do(req)
	case "QueryClient":
		initServer(&c.trMu, &c.trQueryServer, func() *QueryServerTransport {
			return NewQueryServerTransport(&c.srv.QueryServer)
		})
		resp, err = c.trQueryServer.Do(req)
	default:
		err = fmt.Errorf("unhandled client %s", client)
	}

	return resp, err
}
