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

// MediaTypeServer is a fake server for instances of the mediatypegroup.MediaTypeClient type.
type MediaTypeServer struct {
	// StringBodyServer contains the fakes for client StringBodyClient
	StringBodyServer StringBodyServer
}

// NewMediaTypeServerTransport creates a new instance of MediaTypeServerTransport with the provided implementation.
// The returned MediaTypeServerTransport instance is connected to an instance of mediatypegroup.MediaTypeClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewMediaTypeServerTransport(srv *MediaTypeServer) *MediaTypeServerTransport {
	return &MediaTypeServerTransport{srv: srv}
}

// MediaTypeServerTransport connects instances of mediatypegroup.MediaTypeClient to instances of MediaTypeServer.
// Don't use this type directly, use NewMediaTypeServerTransport instead.
type MediaTypeServerTransport struct {
	srv                *MediaTypeServer
	trMu               sync.Mutex
	trStringBodyServer *StringBodyServerTransport
}

// Do implements the policy.Transporter interface for MediaTypeServerTransport.
func (m *MediaTypeServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return m.dispatchToClientFake(req, method[:strings.Index(method, ".")])
}

func (m *MediaTypeServerTransport) dispatchToClientFake(req *http.Request, client string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch client {
	case "StringBodyClient":
		initServer(&m.trMu, &m.trStringBodyServer, func() *StringBodyServerTransport {
			return NewStringBodyServerTransport(&m.srv.StringBodyServer)
		})
		resp, err = m.trStringBodyServer.Do(req)
	default:
		err = fmt.Errorf("unhandled client %s", client)
	}

	return resp, err
}
