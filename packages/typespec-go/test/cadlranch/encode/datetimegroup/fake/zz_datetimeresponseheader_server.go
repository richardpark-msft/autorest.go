// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"context"
	"datetimegroup"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"time"
)

// DatetimeResponseHeaderServer is a fake server for instances of the datetimegroup.DatetimeResponseHeaderClient type.
type DatetimeResponseHeaderServer struct {
	// Default is the fake for method DatetimeResponseHeaderClient.Default
	// HTTP status codes to indicate success: http.StatusNoContent
	Default func(ctx context.Context, options *datetimegroup.DatetimeResponseHeaderClientDefaultOptions) (resp azfake.Responder[datetimegroup.DatetimeResponseHeaderClientDefaultResponse], errResp azfake.ErrorResponder)

	// RFC3339 is the fake for method DatetimeResponseHeaderClient.RFC3339
	// HTTP status codes to indicate success: http.StatusNoContent
	RFC3339 func(ctx context.Context, options *datetimegroup.DatetimeResponseHeaderClientRFC3339Options) (resp azfake.Responder[datetimegroup.DatetimeResponseHeaderClientRFC3339Response], errResp azfake.ErrorResponder)

	// RFC7231 is the fake for method DatetimeResponseHeaderClient.RFC7231
	// HTTP status codes to indicate success: http.StatusNoContent
	RFC7231 func(ctx context.Context, options *datetimegroup.DatetimeResponseHeaderClientRFC7231Options) (resp azfake.Responder[datetimegroup.DatetimeResponseHeaderClientRFC7231Response], errResp azfake.ErrorResponder)

	// UnixTimestamp is the fake for method DatetimeResponseHeaderClient.UnixTimestamp
	// HTTP status codes to indicate success: http.StatusNoContent
	UnixTimestamp func(ctx context.Context, options *datetimegroup.DatetimeResponseHeaderClientUnixTimestampOptions) (resp azfake.Responder[datetimegroup.DatetimeResponseHeaderClientUnixTimestampResponse], errResp azfake.ErrorResponder)
}

// NewDatetimeResponseHeaderServerTransport creates a new instance of DatetimeResponseHeaderServerTransport with the provided implementation.
// The returned DatetimeResponseHeaderServerTransport instance is connected to an instance of datetimegroup.DatetimeResponseHeaderClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewDatetimeResponseHeaderServerTransport(srv *DatetimeResponseHeaderServer) *DatetimeResponseHeaderServerTransport {
	return &DatetimeResponseHeaderServerTransport{srv: srv}
}

// DatetimeResponseHeaderServerTransport connects instances of datetimegroup.DatetimeResponseHeaderClient to instances of DatetimeResponseHeaderServer.
// Don't use this type directly, use NewDatetimeResponseHeaderServerTransport instead.
type DatetimeResponseHeaderServerTransport struct {
	srv *DatetimeResponseHeaderServer
}

// Do implements the policy.Transporter interface for DatetimeResponseHeaderServerTransport.
func (d *DatetimeResponseHeaderServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return d.dispatchToMethodFake(req, method)
}

func (d *DatetimeResponseHeaderServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch method {
	case "DatetimeResponseHeaderClient.Default":
		resp, err = d.dispatchDefault(req)
	case "DatetimeResponseHeaderClient.RFC3339":
		resp, err = d.dispatchRFC3339(req)
	case "DatetimeResponseHeaderClient.RFC7231":
		resp, err = d.dispatchRFC7231(req)
	case "DatetimeResponseHeaderClient.UnixTimestamp":
		resp, err = d.dispatchUnixTimestamp(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	return resp, err
}

func (d *DatetimeResponseHeaderServerTransport) dispatchDefault(req *http.Request) (*http.Response, error) {
	if d.srv.Default == nil {
		return nil, &nonRetriableError{errors.New("fake for method Default not implemented")}
	}
	respr, errRespr := d.srv.Default(req.Context(), nil)
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
	if val := server.GetResponse(respr).Value; val != nil {
		resp.Header.Set("value", val.Format(time.RFC1123))
	}
	return resp, nil
}

func (d *DatetimeResponseHeaderServerTransport) dispatchRFC3339(req *http.Request) (*http.Response, error) {
	if d.srv.RFC3339 == nil {
		return nil, &nonRetriableError{errors.New("fake for method RFC3339 not implemented")}
	}
	respr, errRespr := d.srv.RFC3339(req.Context(), nil)
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
	if val := server.GetResponse(respr).Value; val != nil {
		resp.Header.Set("value", val.Format(time.RFC3339Nano))
	}
	return resp, nil
}

func (d *DatetimeResponseHeaderServerTransport) dispatchRFC7231(req *http.Request) (*http.Response, error) {
	if d.srv.RFC7231 == nil {
		return nil, &nonRetriableError{errors.New("fake for method RFC7231 not implemented")}
	}
	respr, errRespr := d.srv.RFC7231(req.Context(), nil)
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
	if val := server.GetResponse(respr).Value; val != nil {
		resp.Header.Set("value", val.Format(time.RFC1123))
	}
	return resp, nil
}

func (d *DatetimeResponseHeaderServerTransport) dispatchUnixTimestamp(req *http.Request) (*http.Response, error) {
	if d.srv.UnixTimestamp == nil {
		return nil, &nonRetriableError{errors.New("fake for method UnixTimestamp not implemented")}
	}
	respr, errRespr := d.srv.UnixTimestamp(req.Context(), nil)
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
	if val := server.GetResponse(respr).Value; val != nil {
		resp.Header.Set("value", timeUnix(*val).String())
	}
	return resp, nil
}