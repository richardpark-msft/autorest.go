//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package fake

import (
	"aztables"
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strconv"
	"time"
)

// ServiceServer is a fake server for instances of the aztables.ServiceClient type.
type ServiceServer struct {
	// GetProperties is the fake for method ServiceClient.GetProperties
	// HTTP status codes to indicate success: http.StatusOK
	GetProperties func(ctx context.Context, options *aztables.ServiceClientGetPropertiesOptions) (resp azfake.Responder[aztables.ServiceClientGetPropertiesResponse], errResp azfake.ErrorResponder)

	// GetStatistics is the fake for method ServiceClient.GetStatistics
	// HTTP status codes to indicate success: http.StatusOK
	GetStatistics func(ctx context.Context, options *aztables.ServiceClientGetStatisticsOptions) (resp azfake.Responder[aztables.ServiceClientGetStatisticsResponse], errResp azfake.ErrorResponder)

	// SetProperties is the fake for method ServiceClient.SetProperties
	// HTTP status codes to indicate success: http.StatusAccepted
	SetProperties func(ctx context.Context, tableServiceProperties aztables.ServiceProperties, options *aztables.ServiceClientSetPropertiesOptions) (resp azfake.Responder[aztables.ServiceClientSetPropertiesResponse], errResp azfake.ErrorResponder)
}

// NewServiceServerTransport creates a new instance of ServiceServerTransport with the provided implementation.
// The returned ServiceServerTransport instance is connected to an instance of aztables.ServiceClient by way of the
// undefined.Transporter field.
func NewServiceServerTransport(srv *ServiceServer) *ServiceServerTransport {
	return &ServiceServerTransport{srv: srv}
}

// ServiceServerTransport connects instances of aztables.ServiceClient to instances of ServiceServer.
// Don't use this type directly, use NewServiceServerTransport instead.
type ServiceServerTransport struct {
	srv *ServiceServer
}

// Do implements the policy.Transporter interface for ServiceServerTransport.
func (s *ServiceServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ServiceClient.GetProperties":
		resp, err = s.dispatchGetProperties(req)
	case "ServiceClient.GetStatistics":
		resp, err = s.dispatchGetStatistics(req)
	case "ServiceClient.SetProperties":
		resp, err = s.dispatchSetProperties(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *ServiceServerTransport) dispatchGetProperties(req *http.Request) (*http.Response, error) {
	if s.srv.GetProperties == nil {
		return nil, &nonRetriableError{errors.New("method GetProperties not implemented")}
	}
	qp := req.URL.Query()
	timeoutParam, err := parseOptional(qp.Get("timeout"), func(v string) (int32, error) {
		p, parseErr := strconv.ParseInt(v, 10, 32)
		if parseErr != nil {
			return 0, parseErr
		}
		return int32(p), nil
	})
	if err != nil {
		return nil, err
	}
	requestIDParam := getOptional(getHeaderValue(req.Header, "x-ms-client-request-id"))
	var options *aztables.ServiceClientGetPropertiesOptions
	if timeoutParam != nil || requestIDParam != nil {
		options = &aztables.ServiceClientGetPropertiesOptions{
			Timeout:   timeoutParam,
			RequestID: requestIDParam,
		}
	}
	respr, errRespr := s.srv.GetProperties(req.Context(), options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsXML(respContent, server.GetResponse(respr).ServiceProperties, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).ClientRequestID; val != nil {
		resp.Header.Set("x-ms-client-request-id", *val)
	}
	if val := server.GetResponse(respr).RequestID; val != nil {
		resp.Header.Set("x-ms-request-id", *val)
	}
	if val := server.GetResponse(respr).Version; val != nil {
		resp.Header.Set("x-ms-version", *val)
	}
	return resp, nil
}

func (s *ServiceServerTransport) dispatchGetStatistics(req *http.Request) (*http.Response, error) {
	if s.srv.GetStatistics == nil {
		return nil, &nonRetriableError{errors.New("method GetStatistics not implemented")}
	}
	qp := req.URL.Query()
	timeoutParam, err := parseOptional(qp.Get("timeout"), func(v string) (int32, error) {
		p, parseErr := strconv.ParseInt(v, 10, 32)
		if parseErr != nil {
			return 0, parseErr
		}
		return int32(p), nil
	})
	if err != nil {
		return nil, err
	}
	requestIDParam := getOptional(getHeaderValue(req.Header, "x-ms-client-request-id"))
	var options *aztables.ServiceClientGetStatisticsOptions
	if timeoutParam != nil || requestIDParam != nil {
		options = &aztables.ServiceClientGetStatisticsOptions{
			Timeout:   timeoutParam,
			RequestID: requestIDParam,
		}
	}
	respr, errRespr := s.srv.GetStatistics(req.Context(), options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsXML(respContent, server.GetResponse(respr).ServiceStats, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).ClientRequestID; val != nil {
		resp.Header.Set("x-ms-client-request-id", *val)
	}
	if val := server.GetResponse(respr).RequestID; val != nil {
		resp.Header.Set("x-ms-request-id", *val)
	}
	if val := server.GetResponse(respr).Version; val != nil {
		resp.Header.Set("x-ms-version", *val)
	}
	if val := server.GetResponse(respr).Date; val != nil {
		resp.Header.Set("Date", val.Format(time.RFC1123))
	}
	return resp, nil
}

func (s *ServiceServerTransport) dispatchSetProperties(req *http.Request) (*http.Response, error) {
	if s.srv.SetProperties == nil {
		return nil, &nonRetriableError{errors.New("method SetProperties not implemented")}
	}
	qp := req.URL.Query()
	body, err := server.UnmarshalRequestAsXML[aztables.ServiceProperties](req)
	if err != nil {
		return nil, err
	}
	timeoutParam, err := parseOptional(qp.Get("timeout"), func(v string) (int32, error) {
		p, parseErr := strconv.ParseInt(v, 10, 32)
		if parseErr != nil {
			return 0, parseErr
		}
		return int32(p), nil
	})
	if err != nil {
		return nil, err
	}
	requestIDParam := getOptional(getHeaderValue(req.Header, "x-ms-client-request-id"))
	var options *aztables.ServiceClientSetPropertiesOptions
	if timeoutParam != nil || requestIDParam != nil {
		options = &aztables.ServiceClientSetPropertiesOptions{
			Timeout:   timeoutParam,
			RequestID: requestIDParam,
		}
	}
	respr, errRespr := s.srv.SetProperties(req.Context(), body, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusAccepted}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).ClientRequestID; val != nil {
		resp.Header.Set("x-ms-client-request-id", *val)
	}
	if val := server.GetResponse(respr).RequestID; val != nil {
		resp.Header.Set("x-ms-request-id", *val)
	}
	if val := server.GetResponse(respr).Version; val != nil {
		resp.Header.Set("x-ms-version", *val)
	}
	return resp, nil
}