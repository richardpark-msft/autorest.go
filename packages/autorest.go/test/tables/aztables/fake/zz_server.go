//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

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
	"reflect"
	"regexp"
	"strconv"
	"time"
)

// Server is a fake server for instances of the aztables.Client type.
type Server struct {
	// Create is the fake for method Client.Create
	// HTTP status codes to indicate success: http.StatusCreated, http.StatusNoContent
	Create func(ctx context.Context, tableProperties aztables.Properties, options *aztables.ClientCreateOptions) (resp azfake.Responder[aztables.ClientCreateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method Client.Delete
	// HTTP status codes to indicate success: http.StatusNoContent
	Delete func(ctx context.Context, table string, options *aztables.ClientDeleteOptions) (resp azfake.Responder[aztables.ClientDeleteResponse], errResp azfake.ErrorResponder)

	// DeleteEntity is the fake for method Client.DeleteEntity
	// HTTP status codes to indicate success: http.StatusNoContent
	DeleteEntity func(ctx context.Context, table string, partitionKey string, rowKey string, ifMatch string, options *aztables.ClientDeleteEntityOptions) (resp azfake.Responder[aztables.ClientDeleteEntityResponse], errResp azfake.ErrorResponder)

	// GetAccessPolicy is the fake for method Client.GetAccessPolicy
	// HTTP status codes to indicate success: http.StatusOK
	GetAccessPolicy func(ctx context.Context, table string, options *aztables.ClientGetAccessPolicyOptions) (resp azfake.Responder[aztables.ClientGetAccessPolicyResponse], errResp azfake.ErrorResponder)

	// InsertEntity is the fake for method Client.InsertEntity
	// HTTP status codes to indicate success: http.StatusCreated, http.StatusNoContent
	InsertEntity func(ctx context.Context, table string, options *aztables.ClientInsertEntityOptions) (resp azfake.Responder[aztables.ClientInsertEntityResponse], errResp azfake.ErrorResponder)

	// MergeEntity is the fake for method Client.MergeEntity
	// HTTP status codes to indicate success: http.StatusNoContent
	MergeEntity func(ctx context.Context, table string, partitionKey string, rowKey string, tableEntityProperties map[string]any, options *aztables.ClientMergeEntityOptions) (resp azfake.Responder[aztables.ClientMergeEntityResponse], errResp azfake.ErrorResponder)

	// Query is the fake for method Client.Query
	// HTTP status codes to indicate success: http.StatusOK
	Query func(ctx context.Context, options *aztables.ClientQueryOptions) (resp azfake.Responder[aztables.ClientQueryResponse], errResp azfake.ErrorResponder)

	// QueryEntities is the fake for method Client.QueryEntities
	// HTTP status codes to indicate success: http.StatusOK
	QueryEntities func(ctx context.Context, table string, options *aztables.ClientQueryEntitiesOptions) (resp azfake.Responder[aztables.ClientQueryEntitiesResponse], errResp azfake.ErrorResponder)

	// QueryEntityWithPartitionAndRowKey is the fake for method Client.QueryEntityWithPartitionAndRowKey
	// HTTP status codes to indicate success: http.StatusOK
	QueryEntityWithPartitionAndRowKey func(ctx context.Context, table string, partitionKey string, rowKey string, options *aztables.ClientQueryEntityWithPartitionAndRowKeyOptions) (resp azfake.Responder[aztables.ClientQueryEntityWithPartitionAndRowKeyResponse], errResp azfake.ErrorResponder)

	// SetAccessPolicy is the fake for method Client.SetAccessPolicy
	// HTTP status codes to indicate success: http.StatusNoContent
	SetAccessPolicy func(ctx context.Context, table string, tableACL []*aztables.SignedIdentifier, options *aztables.ClientSetAccessPolicyOptions) (resp azfake.Responder[aztables.ClientSetAccessPolicyResponse], errResp azfake.ErrorResponder)

	// UpdateEntity is the fake for method Client.UpdateEntity
	// HTTP status codes to indicate success: http.StatusNoContent
	UpdateEntity func(ctx context.Context, table string, partitionKey string, rowKey string, tableEntityProperties map[string]any, options *aztables.ClientUpdateEntityOptions) (resp azfake.Responder[aztables.ClientUpdateEntityResponse], errResp azfake.ErrorResponder)
}

// NewServerTransport creates a new instance of ServerTransport with the provided implementation.
// The returned ServerTransport instance is connected to an instance of aztables.Client by way of the
// undefined.Transporter field.
func NewServerTransport(srv *Server) *ServerTransport {
	return &ServerTransport{srv: srv}
}

// ServerTransport connects instances of aztables.Client to instances of Server.
// Don't use this type directly, use NewServerTransport instead.
type ServerTransport struct {
	srv *Server
}

// Do implements the policy.Transporter interface for ServerTransport.
func (s *ServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "Client.Create":
		resp, err = s.dispatchCreate(req)
	case "Client.Delete":
		resp, err = s.dispatchDelete(req)
	case "Client.DeleteEntity":
		resp, err = s.dispatchDeleteEntity(req)
	case "Client.GetAccessPolicy":
		resp, err = s.dispatchGetAccessPolicy(req)
	case "Client.InsertEntity":
		resp, err = s.dispatchInsertEntity(req)
	case "Client.MergeEntity":
		resp, err = s.dispatchMergeEntity(req)
	case "Client.Query":
		resp, err = s.dispatchQuery(req)
	case "Client.QueryEntities":
		resp, err = s.dispatchQueryEntities(req)
	case "Client.QueryEntityWithPartitionAndRowKey":
		resp, err = s.dispatchQueryEntityWithPartitionAndRowKey(req)
	case "Client.SetAccessPolicy":
		resp, err = s.dispatchSetAccessPolicy(req)
	case "Client.UpdateEntity":
		resp, err = s.dispatchUpdateEntity(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *ServerTransport) dispatchCreate(req *http.Request) (*http.Response, error) {
	if s.srv.Create == nil {
		return nil, &nonRetriableError{errors.New("method Create not implemented")}
	}
	qp := req.URL.Query()
	body, err := server.UnmarshalRequestAsJSON[aztables.Properties](req)
	if err != nil {
		return nil, err
	}
	requestIDParam := getOptional(getHeaderValue(req.Header, "x-ms-client-request-id"))
	formatParam := getOptional(aztables.ODataMetadataFormat(qp.Get("$format")))
	responsePreferenceParam := getOptional(aztables.ResponseFormat(getHeaderValue(req.Header, "Prefer")))
	var options *aztables.ClientCreateOptions
	if requestIDParam != nil || formatParam != nil || responsePreferenceParam != nil {
		options = &aztables.ClientCreateOptions{
			RequestID:          requestIDParam,
			Format:             formatParam,
			ResponsePreference: responsePreferenceParam,
		}
	}
	respr, errRespr := s.srv.Create(req.Context(), body, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusCreated, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusCreated, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Response, req)
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
	if val := server.GetResponse(respr).PreferenceApplied; val != nil {
		resp.Header.Set("Preference-Applied", *val)
	}
	return resp, nil
}

func (s *ServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if s.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("method Delete not implemented")}
	}
	const regexStr = "/Tables('(?P<table>[a-zA-Z0-9-_]+)')"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	requestIDParam := getOptional(getHeaderValue(req.Header, "x-ms-client-request-id"))
	var options *aztables.ClientDeleteOptions
	if requestIDParam != nil {
		options = &aztables.ClientDeleteOptions{
			RequestID: requestIDParam,
		}
	}
	respr, errRespr := s.srv.Delete(req.Context(), matches[regex.SubexpIndex("table")], options)
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

func (s *ServerTransport) dispatchDeleteEntity(req *http.Request) (*http.Response, error) {
	if s.srv.DeleteEntity == nil {
		return nil, &nonRetriableError{errors.New("method DeleteEntity not implemented")}
	}
	const regexStr = "/(?P<table>[a-zA-Z0-9-_]+)(PartitionKey='(?P<partitionKey>[a-zA-Z0-9-_]+)',RowKey='(?P<rowKey>[a-zA-Z0-9-_]+)')"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
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
	formatParam := getOptional(aztables.ODataMetadataFormat(qp.Get("$format")))
	var options *aztables.ClientDeleteEntityOptions
	if timeoutParam != nil || requestIDParam != nil || formatParam != nil {
		options = &aztables.ClientDeleteEntityOptions{
			Timeout:   timeoutParam,
			RequestID: requestIDParam,
			Format:    formatParam,
		}
	}
	respr, errRespr := s.srv.DeleteEntity(req.Context(), matches[regex.SubexpIndex("table")], matches[regex.SubexpIndex("partitionKey")], matches[regex.SubexpIndex("rowKey")], getHeaderValue(req.Header, "If-Match"), options)
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

func (s *ServerTransport) dispatchGetAccessPolicy(req *http.Request) (*http.Response, error) {
	if s.srv.GetAccessPolicy == nil {
		return nil, &nonRetriableError{errors.New("method GetAccessPolicy not implemented")}
	}
	const regexStr = "/(?P<table>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
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
	var options *aztables.ClientGetAccessPolicyOptions
	if timeoutParam != nil || requestIDParam != nil {
		options = &aztables.ClientGetAccessPolicyOptions{
			Timeout:   timeoutParam,
			RequestID: requestIDParam,
		}
	}
	respr, errRespr := s.srv.GetAccessPolicy(req.Context(), matches[regex.SubexpIndex("table")], options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsXML(respContent, server.GetResponse(respr), req)
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

func (s *ServerTransport) dispatchInsertEntity(req *http.Request) (*http.Response, error) {
	if s.srv.InsertEntity == nil {
		return nil, &nonRetriableError{errors.New("method InsertEntity not implemented")}
	}
	const regexStr = "/(?P<table>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	body, err := server.UnmarshalRequestAsJSON[map[string]any](req)
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
	formatParam := getOptional(aztables.ODataMetadataFormat(qp.Get("$format")))
	responsePreferenceParam := getOptional(aztables.ResponseFormat(getHeaderValue(req.Header, "Prefer")))
	var options *aztables.ClientInsertEntityOptions
	if timeoutParam != nil || requestIDParam != nil || formatParam != nil || responsePreferenceParam != nil || !reflect.ValueOf(body).IsZero() {
		options = &aztables.ClientInsertEntityOptions{
			Timeout:               timeoutParam,
			RequestID:             requestIDParam,
			Format:                formatParam,
			ResponsePreference:    responsePreferenceParam,
			TableEntityProperties: body,
		}
	}
	respr, errRespr := s.srv.InsertEntity(req.Context(), matches[regex.SubexpIndex("table")], options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusCreated, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusCreated, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Value, req)
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
	if val := server.GetResponse(respr).ETag; val != nil {
		resp.Header.Set("ETag", *val)
	}
	if val := server.GetResponse(respr).PreferenceApplied; val != nil {
		resp.Header.Set("Preference-Applied", *val)
	}
	if val := server.GetResponse(respr).ContentType; val != nil {
		resp.Header.Set("Content-Type", *val)
	}
	return resp, nil
}

func (s *ServerTransport) dispatchMergeEntity(req *http.Request) (*http.Response, error) {
	if s.srv.MergeEntity == nil {
		return nil, &nonRetriableError{errors.New("method MergeEntity not implemented")}
	}
	const regexStr = "/(?P<table>[a-zA-Z0-9-_]+)(PartitionKey='(?P<partitionKey>[a-zA-Z0-9-_]+)',RowKey='(?P<rowKey>[a-zA-Z0-9-_]+)')"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	body, err := server.UnmarshalRequestAsJSON[map[string]any](req)
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
	formatParam := getOptional(aztables.ODataMetadataFormat(qp.Get("$format")))
	ifMatchParam := getOptional(getHeaderValue(req.Header, "If-Match"))
	var options *aztables.ClientMergeEntityOptions
	if timeoutParam != nil || requestIDParam != nil || formatParam != nil || ifMatchParam != nil {
		options = &aztables.ClientMergeEntityOptions{
			Timeout:   timeoutParam,
			RequestID: requestIDParam,
			Format:    formatParam,
			IfMatch:   ifMatchParam,
		}
	}
	respr, errRespr := s.srv.MergeEntity(req.Context(), matches[regex.SubexpIndex("table")], matches[regex.SubexpIndex("partitionKey")], matches[regex.SubexpIndex("rowKey")], body, options)
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
	if val := server.GetResponse(respr).ETag; val != nil {
		resp.Header.Set("ETag", *val)
	}
	return resp, nil
}

func (s *ServerTransport) dispatchQuery(req *http.Request) (*http.Response, error) {
	if s.srv.Query == nil {
		return nil, &nonRetriableError{errors.New("method Query not implemented")}
	}
	qp := req.URL.Query()
	requestIDParam := getOptional(getHeaderValue(req.Header, "x-ms-client-request-id"))
	formatParam := getOptional(aztables.ODataMetadataFormat(qp.Get("$format")))
	topParam, err := parseOptional(qp.Get("$top"), func(v string) (int32, error) {
		p, parseErr := strconv.ParseInt(v, 10, 32)
		if parseErr != nil {
			return 0, parseErr
		}
		return int32(p), nil
	})
	if err != nil {
		return nil, err
	}
	selectParam := getOptional(qp.Get("$select"))
	filterParam := getOptional(qp.Get("$filter"))
	nextTableNameParam := getOptional(qp.Get("NextTableName"))
	var options *aztables.ClientQueryOptions
	if requestIDParam != nil || formatParam != nil || topParam != nil || selectParam != nil || filterParam != nil || nextTableNameParam != nil {
		options = &aztables.ClientQueryOptions{
			RequestID:     requestIDParam,
			Format:        formatParam,
			Top:           topParam,
			Select:        selectParam,
			Filter:        filterParam,
			NextTableName: nextTableNameParam,
		}
	}
	respr, errRespr := s.srv.Query(req.Context(), options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).QueryResponse, req)
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
	if val := server.GetResponse(respr).XMSContinuationNextTableName; val != nil {
		resp.Header.Set("x-ms-continuation-NextTableName", *val)
	}
	return resp, nil
}

func (s *ServerTransport) dispatchQueryEntities(req *http.Request) (*http.Response, error) {
	if s.srv.QueryEntities == nil {
		return nil, &nonRetriableError{errors.New("method QueryEntities not implemented")}
	}
	const regexStr = "/(?P<table>[a-zA-Z0-9-_]+)()"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
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
	formatParam := getOptional(aztables.ODataMetadataFormat(qp.Get("$format")))
	topParam, err := parseOptional(qp.Get("$top"), func(v string) (int32, error) {
		p, parseErr := strconv.ParseInt(v, 10, 32)
		if parseErr != nil {
			return 0, parseErr
		}
		return int32(p), nil
	})
	if err != nil {
		return nil, err
	}
	selectParam := getOptional(qp.Get("$select"))
	filterParam := getOptional(qp.Get("$filter"))
	nextPartitionKeyParam := getOptional(qp.Get("NextPartitionKey"))
	nextRowKeyParam := getOptional(qp.Get("NextRowKey"))
	var options *aztables.ClientQueryEntitiesOptions
	if timeoutParam != nil || requestIDParam != nil || formatParam != nil || topParam != nil || selectParam != nil || filterParam != nil || nextPartitionKeyParam != nil || nextRowKeyParam != nil {
		options = &aztables.ClientQueryEntitiesOptions{
			Timeout:          timeoutParam,
			RequestID:        requestIDParam,
			Format:           formatParam,
			Top:              topParam,
			Select:           selectParam,
			Filter:           filterParam,
			NextPartitionKey: nextPartitionKeyParam,
			NextRowKey:       nextRowKeyParam,
		}
	}
	respr, errRespr := s.srv.QueryEntities(req.Context(), matches[regex.SubexpIndex("table")], options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).EntityQueryResponse, req)
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
	if val := server.GetResponse(respr).XMSContinuationNextPartitionKey; val != nil {
		resp.Header.Set("x-ms-continuation-NextPartitionKey", *val)
	}
	if val := server.GetResponse(respr).XMSContinuationNextRowKey; val != nil {
		resp.Header.Set("x-ms-continuation-NextRowKey", *val)
	}
	return resp, nil
}

func (s *ServerTransport) dispatchQueryEntityWithPartitionAndRowKey(req *http.Request) (*http.Response, error) {
	if s.srv.QueryEntityWithPartitionAndRowKey == nil {
		return nil, &nonRetriableError{errors.New("method QueryEntityWithPartitionAndRowKey not implemented")}
	}
	const regexStr = "/(?P<table>[a-zA-Z0-9-_]+)(PartitionKey='(?P<partitionKey>[a-zA-Z0-9-_]+)',RowKey='(?P<rowKey>[a-zA-Z0-9-_]+)')"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
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
	formatParam := getOptional(aztables.ODataMetadataFormat(qp.Get("$format")))
	selectParam := getOptional(qp.Get("$select"))
	filterParam := getOptional(qp.Get("$filter"))
	var options *aztables.ClientQueryEntityWithPartitionAndRowKeyOptions
	if timeoutParam != nil || requestIDParam != nil || formatParam != nil || selectParam != nil || filterParam != nil {
		options = &aztables.ClientQueryEntityWithPartitionAndRowKeyOptions{
			Timeout:   timeoutParam,
			RequestID: requestIDParam,
			Format:    formatParam,
			Select:    selectParam,
			Filter:    filterParam,
		}
	}
	respr, errRespr := s.srv.QueryEntityWithPartitionAndRowKey(req.Context(), matches[regex.SubexpIndex("table")], matches[regex.SubexpIndex("partitionKey")], matches[regex.SubexpIndex("rowKey")], options)
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
	if val := server.GetResponse(respr).ETag; val != nil {
		resp.Header.Set("ETag", *val)
	}
	if val := server.GetResponse(respr).XMSContinuationNextPartitionKey; val != nil {
		resp.Header.Set("x-ms-continuation-NextPartitionKey", *val)
	}
	if val := server.GetResponse(respr).XMSContinuationNextRowKey; val != nil {
		resp.Header.Set("x-ms-continuation-NextRowKey", *val)
	}
	return resp, nil
}

func (s *ServerTransport) dispatchSetAccessPolicy(req *http.Request) (*http.Response, error) {
	if s.srv.SetAccessPolicy == nil {
		return nil, &nonRetriableError{errors.New("method SetAccessPolicy not implemented")}
	}
	const regexStr = "/(?P<table>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	body, err := server.UnmarshalRequestAsXML[[]*aztables.SignedIdentifier](req)
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
	var options *aztables.ClientSetAccessPolicyOptions
	if timeoutParam != nil || requestIDParam != nil {
		options = &aztables.ClientSetAccessPolicyOptions{
			Timeout:   timeoutParam,
			RequestID: requestIDParam,
		}
	}
	respr, errRespr := s.srv.SetAccessPolicy(req.Context(), matches[regex.SubexpIndex("table")], body, options)
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

func (s *ServerTransport) dispatchUpdateEntity(req *http.Request) (*http.Response, error) {
	if s.srv.UpdateEntity == nil {
		return nil, &nonRetriableError{errors.New("method UpdateEntity not implemented")}
	}
	const regexStr = "/(?P<table>[a-zA-Z0-9-_]+)(PartitionKey='(?P<partitionKey>[a-zA-Z0-9-_]+)',RowKey='(?P<rowKey>[a-zA-Z0-9-_]+)')"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	body, err := server.UnmarshalRequestAsJSON[map[string]any](req)
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
	formatParam := getOptional(aztables.ODataMetadataFormat(qp.Get("$format")))
	ifMatchParam := getOptional(getHeaderValue(req.Header, "If-Match"))
	var options *aztables.ClientUpdateEntityOptions
	if timeoutParam != nil || requestIDParam != nil || formatParam != nil || ifMatchParam != nil {
		options = &aztables.ClientUpdateEntityOptions{
			Timeout:   timeoutParam,
			RequestID: requestIDParam,
			Format:    formatParam,
			IfMatch:   ifMatchParam,
		}
	}
	respr, errRespr := s.srv.UpdateEntity(req.Context(), matches[regex.SubexpIndex("table")], matches[regex.SubexpIndex("partitionKey")], matches[regex.SubexpIndex("rowKey")], body, options)
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
	if val := server.GetResponse(respr).ETag; val != nil {
		resp.Header.Set("ETag", *val)
	}
	return resp, nil
}
