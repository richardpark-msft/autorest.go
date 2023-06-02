//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package fake

import (
	"azblob"
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// ServiceServer is a fake server for instances of the azblob.ServiceClient type.
type ServiceServer struct {
	// FilterBlobs is the fake for method ServiceClient.FilterBlobs
	// HTTP status codes to indicate success: http.StatusOK
	FilterBlobs func(ctx context.Context, comp azblob.Enum10, options *azblob.ServiceClientFilterBlobsOptions) (resp azfake.Responder[azblob.ServiceClientFilterBlobsResponse], errResp azfake.ErrorResponder)

	// GetAccountInfo is the fake for method ServiceClient.GetAccountInfo
	// HTTP status codes to indicate success: http.StatusOK
	GetAccountInfo func(ctx context.Context, restype azblob.Enum8, comp azblob.Enum1, options *azblob.ServiceClientGetAccountInfoOptions) (resp azfake.Responder[azblob.ServiceClientGetAccountInfoResponse], errResp azfake.ErrorResponder)

	// GetProperties is the fake for method ServiceClient.GetProperties
	// HTTP status codes to indicate success: http.StatusOK
	GetProperties func(ctx context.Context, restype azblob.Enum0, comp azblob.Enum1, options *azblob.ServiceClientGetPropertiesOptions) (resp azfake.Responder[azblob.ServiceClientGetPropertiesResponse], errResp azfake.ErrorResponder)

	// GetStatistics is the fake for method ServiceClient.GetStatistics
	// HTTP status codes to indicate success: http.StatusOK
	GetStatistics func(ctx context.Context, restype azblob.Enum0, comp azblob.Enum3, options *azblob.ServiceClientGetStatisticsOptions) (resp azfake.Responder[azblob.ServiceClientGetStatisticsResponse], errResp azfake.ErrorResponder)

	// GetUserDelegationKey is the fake for method ServiceClient.GetUserDelegationKey
	// HTTP status codes to indicate success: http.StatusOK
	GetUserDelegationKey func(ctx context.Context, restype azblob.Enum0, comp azblob.Enum7, keyInfo azblob.KeyInfo, options *azblob.ServiceClientGetUserDelegationKeyOptions) (resp azfake.Responder[azblob.ServiceClientGetUserDelegationKeyResponse], errResp azfake.ErrorResponder)

	// NewListContainersSegmentPager is the fake for method ServiceClient.NewListContainersSegmentPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListContainersSegmentPager func(comp azblob.Enum5, options *azblob.ServiceClientListContainersSegmentOptions) (resp azfake.PagerResponder[azblob.ServiceClientListContainersSegmentResponse])

	// SetProperties is the fake for method ServiceClient.SetProperties
	// HTTP status codes to indicate success: http.StatusAccepted
	SetProperties func(ctx context.Context, restype azblob.Enum0, comp azblob.Enum1, storageServiceProperties azblob.StorageServiceProperties, options *azblob.ServiceClientSetPropertiesOptions) (resp azfake.Responder[azblob.ServiceClientSetPropertiesResponse], errResp azfake.ErrorResponder)

	// SubmitBatch is the fake for method ServiceClient.SubmitBatch
	// HTTP status codes to indicate success: http.StatusOK
	SubmitBatch func(ctx context.Context, comp azblob.Enum9, contentLength int64, multipartContentType string, body io.ReadSeekCloser, options *azblob.ServiceClientSubmitBatchOptions) (resp azfake.Responder[azblob.ServiceClientSubmitBatchResponse], errResp azfake.ErrorResponder)
}

// NewServiceServerTransport creates a new instance of ServiceServerTransport with the provided implementation.
// The returned ServiceServerTransport instance is connected to an instance of azblob.ServiceClient by way of the
// undefined.Transporter field.
func NewServiceServerTransport(srv *ServiceServer) *ServiceServerTransport {
	return &ServiceServerTransport{srv: srv}
}

// ServiceServerTransport connects instances of azblob.ServiceClient to instances of ServiceServer.
// Don't use this type directly, use NewServiceServerTransport instead.
type ServiceServerTransport struct {
	srv                           *ServiceServer
	newListContainersSegmentPager *azfake.PagerResponder[azblob.ServiceClientListContainersSegmentResponse]
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
	case "ServiceClient.FilterBlobs":
		resp, err = s.dispatchFilterBlobs(req)
	case "ServiceClient.GetAccountInfo":
		resp, err = s.dispatchGetAccountInfo(req)
	case "ServiceClient.GetProperties":
		resp, err = s.dispatchGetProperties(req)
	case "ServiceClient.GetStatistics":
		resp, err = s.dispatchGetStatistics(req)
	case "ServiceClient.GetUserDelegationKey":
		resp, err = s.dispatchGetUserDelegationKey(req)
	case "ServiceClient.NewListContainersSegmentPager":
		resp, err = s.dispatchNewListContainersSegmentPager(req)
	case "ServiceClient.SetProperties":
		resp, err = s.dispatchSetProperties(req)
	case "ServiceClient.SubmitBatch":
		resp, err = s.dispatchSubmitBatch(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *ServiceServerTransport) dispatchFilterBlobs(req *http.Request) (*http.Response, error) {
	if s.srv.FilterBlobs == nil {
		return nil, &nonRetriableError{errors.New("method FilterBlobs not implemented")}
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
	whereParam := getOptional(qp.Get("where"))
	markerParam := getOptional(qp.Get("marker"))
	maxresultsParam, err := parseOptional(qp.Get("maxresults"), func(v string) (int32, error) {
		p, parseErr := strconv.ParseInt(v, 10, 32)
		if parseErr != nil {
			return 0, parseErr
		}
		return int32(p), nil
	})
	if err != nil {
		return nil, err
	}
	elements := strings.Split(qp.Get("include"), ",")
	includeParam := make([]azblob.FilterBlobsIncludeItem, len(elements))
	for i := 0; i < len(elements); i++ {
		includeParam[i] = azblob.FilterBlobsIncludeItem(elements[i])
	}
	var options *azblob.ServiceClientFilterBlobsOptions
	if timeoutParam != nil || requestIDParam != nil || whereParam != nil || markerParam != nil || maxresultsParam != nil || len(includeParam) > 0 {
		options = &azblob.ServiceClientFilterBlobsOptions{
			Timeout:    timeoutParam,
			RequestID:  requestIDParam,
			Where:      whereParam,
			Marker:     markerParam,
			Maxresults: maxresultsParam,
			Include:    includeParam,
		}
	}
	respr, errRespr := s.srv.FilterBlobs(req.Context(), azblob.Enum10(qp.Get("comp")), options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsXML(respContent, server.GetResponse(respr).FilterBlobSegment, req)
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

func (s *ServiceServerTransport) dispatchGetAccountInfo(req *http.Request) (*http.Response, error) {
	if s.srv.GetAccountInfo == nil {
		return nil, &nonRetriableError{errors.New("method GetAccountInfo not implemented")}
	}
	qp := req.URL.Query()
	respr, errRespr := s.srv.GetAccountInfo(req.Context(), azblob.Enum8(qp.Get("restype")), azblob.Enum1(qp.Get("comp")), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
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
	if val := server.GetResponse(respr).Date; val != nil {
		resp.Header.Set("Date", val.Format(time.RFC1123))
	}
	if val := server.GetResponse(respr).SKUName; val != nil {
		resp.Header.Set("x-ms-sku-name", string(*val))
	}
	if val := server.GetResponse(respr).AccountKind; val != nil {
		resp.Header.Set("x-ms-account-kind", string(*val))
	}
	if val := server.GetResponse(respr).IsHierarchicalNamespaceEnabled; val != nil {
		resp.Header.Set("x-ms-is-hns-enabled", strconv.FormatBool(*val))
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
	var options *azblob.ServiceClientGetPropertiesOptions
	if timeoutParam != nil || requestIDParam != nil {
		options = &azblob.ServiceClientGetPropertiesOptions{
			Timeout:   timeoutParam,
			RequestID: requestIDParam,
		}
	}
	respr, errRespr := s.srv.GetProperties(req.Context(), azblob.Enum0(qp.Get("restype")), azblob.Enum1(qp.Get("comp")), options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsXML(respContent, server.GetResponse(respr).StorageServiceProperties, req)
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
	var options *azblob.ServiceClientGetStatisticsOptions
	if timeoutParam != nil || requestIDParam != nil {
		options = &azblob.ServiceClientGetStatisticsOptions{
			Timeout:   timeoutParam,
			RequestID: requestIDParam,
		}
	}
	respr, errRespr := s.srv.GetStatistics(req.Context(), azblob.Enum0(qp.Get("restype")), azblob.Enum3(qp.Get("comp")), options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsXML(respContent, server.GetResponse(respr).StorageServiceStats, req)
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

func (s *ServiceServerTransport) dispatchGetUserDelegationKey(req *http.Request) (*http.Response, error) {
	if s.srv.GetUserDelegationKey == nil {
		return nil, &nonRetriableError{errors.New("method GetUserDelegationKey not implemented")}
	}
	qp := req.URL.Query()
	body, err := server.UnmarshalRequestAsXML[azblob.KeyInfo](req)
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
	var options *azblob.ServiceClientGetUserDelegationKeyOptions
	if timeoutParam != nil || requestIDParam != nil {
		options = &azblob.ServiceClientGetUserDelegationKeyOptions{
			Timeout:   timeoutParam,
			RequestID: requestIDParam,
		}
	}
	respr, errRespr := s.srv.GetUserDelegationKey(req.Context(), azblob.Enum0(qp.Get("restype")), azblob.Enum7(qp.Get("comp")), body, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsXML(respContent, server.GetResponse(respr).UserDelegationKey, req)
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

func (s *ServiceServerTransport) dispatchNewListContainersSegmentPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListContainersSegmentPager == nil {
		return nil, &nonRetriableError{errors.New("method NewListContainersSegmentPager not implemented")}
	}
	if s.newListContainersSegmentPager == nil {
		qp := req.URL.Query()
		prefixParam := getOptional(qp.Get("prefix"))
		markerParam := getOptional(qp.Get("marker"))
		maxresultsParam, err := parseOptional(qp.Get("maxresults"), func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		elements := strings.Split(qp.Get("include"), ",")
		includeParam := make([]azblob.ListContainersIncludeType, len(elements))
		for i := 0; i < len(elements); i++ {
			includeParam[i] = azblob.ListContainersIncludeType(elements[i])
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
		var options *azblob.ServiceClientListContainersSegmentOptions
		if prefixParam != nil || markerParam != nil || maxresultsParam != nil || len(includeParam) > 0 || timeoutParam != nil || requestIDParam != nil {
			options = &azblob.ServiceClientListContainersSegmentOptions{
				Prefix:     prefixParam,
				Marker:     markerParam,
				Maxresults: maxresultsParam,
				Include:    includeParam,
				Timeout:    timeoutParam,
				RequestID:  requestIDParam,
			}
		}
		resp := s.srv.NewListContainersSegmentPager(azblob.Enum5(qp.Get("comp")), options)
		s.newListContainersSegmentPager = &resp
		server.PagerResponderInjectNextLinks(s.newListContainersSegmentPager, req, func(page *azblob.ServiceClientListContainersSegmentResponse, createLink func() string) {
			page.NextMarker = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(s.newListContainersSegmentPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(s.newListContainersSegmentPager) {
		s.newListContainersSegmentPager = nil
	}
	return resp, nil
}

func (s *ServiceServerTransport) dispatchSetProperties(req *http.Request) (*http.Response, error) {
	if s.srv.SetProperties == nil {
		return nil, &nonRetriableError{errors.New("method SetProperties not implemented")}
	}
	qp := req.URL.Query()
	body, err := server.UnmarshalRequestAsXML[azblob.StorageServiceProperties](req)
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
	var options *azblob.ServiceClientSetPropertiesOptions
	if timeoutParam != nil || requestIDParam != nil {
		options = &azblob.ServiceClientSetPropertiesOptions{
			Timeout:   timeoutParam,
			RequestID: requestIDParam,
		}
	}
	respr, errRespr := s.srv.SetProperties(req.Context(), azblob.Enum0(qp.Get("restype")), azblob.Enum1(qp.Get("comp")), body, options)
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

func (s *ServiceServerTransport) dispatchSubmitBatch(req *http.Request) (*http.Response, error) {
	if s.srv.SubmitBatch == nil {
		return nil, &nonRetriableError{errors.New("method SubmitBatch not implemented")}
	}
	qp := req.URL.Query()
	contentLengthParam, err := strconv.ParseInt(getHeaderValue(req.Header, "Content-Length"), 10, 64)
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
	var options *azblob.ServiceClientSubmitBatchOptions
	if timeoutParam != nil || requestIDParam != nil {
		options = &azblob.ServiceClientSubmitBatchOptions{
			Timeout:   timeoutParam,
			RequestID: requestIDParam,
		}
	}
	respr, errRespr := s.srv.SubmitBatch(req.Context(), azblob.Enum9(qp.Get("comp")), contentLengthParam, getHeaderValue(req.Header, "Content-Type"), req.Body.(io.ReadSeekCloser), options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, &server.ResponseOptions{
		Body:        server.GetResponse(respr).Body,
		ContentType: "application/octet-stream",
	})
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).ContentType; val != nil {
		resp.Header.Set("Content-Type", *val)
	}
	if val := server.GetResponse(respr).RequestID; val != nil {
		resp.Header.Set("x-ms-request-id", *val)
	}
	if val := server.GetResponse(respr).Version; val != nil {
		resp.Header.Set("x-ms-version", *val)
	}
	return resp, nil
}