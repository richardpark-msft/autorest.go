//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"azspark"
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"regexp"
	"strconv"
)

// SessionServer is a fake server for instances of the azspark.SessionClient type.
type SessionServer struct {
	// CancelSparkSession is the fake for method SessionClient.CancelSparkSession
	// HTTP status codes to indicate success: http.StatusOK
	CancelSparkSession func(ctx context.Context, sessionID int32, options *azspark.SessionClientCancelSparkSessionOptions) (resp azfake.Responder[azspark.SessionClientCancelSparkSessionResponse], errResp azfake.ErrorResponder)

	// CancelSparkStatement is the fake for method SessionClient.CancelSparkStatement
	// HTTP status codes to indicate success: http.StatusOK
	CancelSparkStatement func(ctx context.Context, sessionID int32, statementID int32, options *azspark.SessionClientCancelSparkStatementOptions) (resp azfake.Responder[azspark.SessionClientCancelSparkStatementResponse], errResp azfake.ErrorResponder)

	// CreateSparkSession is the fake for method SessionClient.CreateSparkSession
	// HTTP status codes to indicate success: http.StatusOK
	CreateSparkSession func(ctx context.Context, sparkSessionOptions azspark.SessionOptions, options *azspark.SessionClientCreateSparkSessionOptions) (resp azfake.Responder[azspark.SessionClientCreateSparkSessionResponse], errResp azfake.ErrorResponder)

	// CreateSparkStatement is the fake for method SessionClient.CreateSparkStatement
	// HTTP status codes to indicate success: http.StatusOK
	CreateSparkStatement func(ctx context.Context, sessionID int32, sparkStatementOptions azspark.StatementOptions, options *azspark.SessionClientCreateSparkStatementOptions) (resp azfake.Responder[azspark.SessionClientCreateSparkStatementResponse], errResp azfake.ErrorResponder)

	// GetSparkSession is the fake for method SessionClient.GetSparkSession
	// HTTP status codes to indicate success: http.StatusOK
	GetSparkSession func(ctx context.Context, sessionID int32, options *azspark.SessionClientGetSparkSessionOptions) (resp azfake.Responder[azspark.SessionClientGetSparkSessionResponse], errResp azfake.ErrorResponder)

	// GetSparkSessions is the fake for method SessionClient.GetSparkSessions
	// HTTP status codes to indicate success: http.StatusOK
	GetSparkSessions func(ctx context.Context, options *azspark.SessionClientGetSparkSessionsOptions) (resp azfake.Responder[azspark.SessionClientGetSparkSessionsResponse], errResp azfake.ErrorResponder)

	// GetSparkStatement is the fake for method SessionClient.GetSparkStatement
	// HTTP status codes to indicate success: http.StatusOK
	GetSparkStatement func(ctx context.Context, sessionID int32, statementID int32, options *azspark.SessionClientGetSparkStatementOptions) (resp azfake.Responder[azspark.SessionClientGetSparkStatementResponse], errResp azfake.ErrorResponder)

	// GetSparkStatements is the fake for method SessionClient.GetSparkStatements
	// HTTP status codes to indicate success: http.StatusOK
	GetSparkStatements func(ctx context.Context, sessionID int32, options *azspark.SessionClientGetSparkStatementsOptions) (resp azfake.Responder[azspark.SessionClientGetSparkStatementsResponse], errResp azfake.ErrorResponder)

	// ResetSparkSessionTimeout is the fake for method SessionClient.ResetSparkSessionTimeout
	// HTTP status codes to indicate success: http.StatusOK
	ResetSparkSessionTimeout func(ctx context.Context, sessionID int32, options *azspark.SessionClientResetSparkSessionTimeoutOptions) (resp azfake.Responder[azspark.SessionClientResetSparkSessionTimeoutResponse], errResp azfake.ErrorResponder)
}

// NewSessionServerTransport creates a new instance of SessionServerTransport with the provided implementation.
// The returned SessionServerTransport instance is connected to an instance of azspark.SessionClient by way of the
// undefined.Transporter field.
func NewSessionServerTransport(srv *SessionServer) *SessionServerTransport {
	return &SessionServerTransport{srv: srv}
}

// SessionServerTransport connects instances of azspark.SessionClient to instances of SessionServer.
// Don't use this type directly, use NewSessionServerTransport instead.
type SessionServerTransport struct {
	srv *SessionServer
}

// Do implements the policy.Transporter interface for SessionServerTransport.
func (s *SessionServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "SessionClient.CancelSparkSession":
		resp, err = s.dispatchCancelSparkSession(req)
	case "SessionClient.CancelSparkStatement":
		resp, err = s.dispatchCancelSparkStatement(req)
	case "SessionClient.CreateSparkSession":
		resp, err = s.dispatchCreateSparkSession(req)
	case "SessionClient.CreateSparkStatement":
		resp, err = s.dispatchCreateSparkStatement(req)
	case "SessionClient.GetSparkSession":
		resp, err = s.dispatchGetSparkSession(req)
	case "SessionClient.GetSparkSessions":
		resp, err = s.dispatchGetSparkSessions(req)
	case "SessionClient.GetSparkStatement":
		resp, err = s.dispatchGetSparkStatement(req)
	case "SessionClient.GetSparkStatements":
		resp, err = s.dispatchGetSparkStatements(req)
	case "SessionClient.ResetSparkSessionTimeout":
		resp, err = s.dispatchResetSparkSessionTimeout(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *SessionServerTransport) dispatchCancelSparkSession(req *http.Request) (*http.Response, error) {
	if s.srv.CancelSparkSession == nil {
		return nil, &nonRetriableError{errors.New("method CancelSparkSession not implemented")}
	}
	const regexStr = "/livyApi/versions/(?P<livyApiVersion>[a-zA-Z0-9-_]+)/sparkPools/(?P<sparkPoolName>[a-zA-Z0-9-_]+)/sessions/(?P<sessionId>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	sessionIDParam, err := parseWithCast(matches[regex.SubexpIndex("sessionId")], func(v string) (int32, error) {
		p, parseErr := strconv.ParseInt(v, 10, 32)
		if parseErr != nil {
			return 0, parseErr
		}
		return int32(p), nil
	})
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.CancelSparkSession(req.Context(), int32(sessionIDParam), nil)
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
	return resp, nil
}

func (s *SessionServerTransport) dispatchCancelSparkStatement(req *http.Request) (*http.Response, error) {
	if s.srv.CancelSparkStatement == nil {
		return nil, &nonRetriableError{errors.New("method CancelSparkStatement not implemented")}
	}
	const regexStr = "/livyApi/versions/(?P<livyApiVersion>[a-zA-Z0-9-_]+)/sparkPools/(?P<sparkPoolName>[a-zA-Z0-9-_]+)/sessions/(?P<sessionId>[a-zA-Z0-9-_]+)/statements/(?P<statementId>[a-zA-Z0-9-_]+)/cancel"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	sessionIDParam, err := parseWithCast(matches[regex.SubexpIndex("sessionId")], func(v string) (int32, error) {
		p, parseErr := strconv.ParseInt(v, 10, 32)
		if parseErr != nil {
			return 0, parseErr
		}
		return int32(p), nil
	})
	if err != nil {
		return nil, err
	}
	statementIDParam, err := parseWithCast(matches[regex.SubexpIndex("statementId")], func(v string) (int32, error) {
		p, parseErr := strconv.ParseInt(v, 10, 32)
		if parseErr != nil {
			return 0, parseErr
		}
		return int32(p), nil
	})
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.CancelSparkStatement(req.Context(), int32(sessionIDParam), int32(statementIDParam), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).StatementCancellationResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SessionServerTransport) dispatchCreateSparkSession(req *http.Request) (*http.Response, error) {
	if s.srv.CreateSparkSession == nil {
		return nil, &nonRetriableError{errors.New("method CreateSparkSession not implemented")}
	}
	const regexStr = "/livyApi/versions/(?P<livyApiVersion>[a-zA-Z0-9-_]+)/sparkPools/(?P<sparkPoolName>[a-zA-Z0-9-_]+)/sessions"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	body, err := server.UnmarshalRequestAsJSON[azspark.SessionOptions](req)
	if err != nil {
		return nil, err
	}
	detailedParam, err := parseOptional(qp.Get("detailed"), strconv.ParseBool)
	if err != nil {
		return nil, err
	}
	var options *azspark.SessionClientCreateSparkSessionOptions
	if detailedParam != nil {
		options = &azspark.SessionClientCreateSparkSessionOptions{
			Detailed: detailedParam,
		}
	}
	respr, errRespr := s.srv.CreateSparkSession(req.Context(), body, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Session, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SessionServerTransport) dispatchCreateSparkStatement(req *http.Request) (*http.Response, error) {
	if s.srv.CreateSparkStatement == nil {
		return nil, &nonRetriableError{errors.New("method CreateSparkStatement not implemented")}
	}
	const regexStr = "/livyApi/versions/(?P<livyApiVersion>[a-zA-Z0-9-_]+)/sparkPools/(?P<sparkPoolName>[a-zA-Z0-9-_]+)/sessions/(?P<sessionId>[a-zA-Z0-9-_]+)/statements"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[azspark.StatementOptions](req)
	if err != nil {
		return nil, err
	}
	sessionIDParam, err := parseWithCast(matches[regex.SubexpIndex("sessionId")], func(v string) (int32, error) {
		p, parseErr := strconv.ParseInt(v, 10, 32)
		if parseErr != nil {
			return 0, parseErr
		}
		return int32(p), nil
	})
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.CreateSparkStatement(req.Context(), int32(sessionIDParam), body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Statement, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SessionServerTransport) dispatchGetSparkSession(req *http.Request) (*http.Response, error) {
	if s.srv.GetSparkSession == nil {
		return nil, &nonRetriableError{errors.New("method GetSparkSession not implemented")}
	}
	const regexStr = "/livyApi/versions/(?P<livyApiVersion>[a-zA-Z0-9-_]+)/sparkPools/(?P<sparkPoolName>[a-zA-Z0-9-_]+)/sessions/(?P<sessionId>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	sessionIDParam, err := parseWithCast(matches[regex.SubexpIndex("sessionId")], func(v string) (int32, error) {
		p, parseErr := strconv.ParseInt(v, 10, 32)
		if parseErr != nil {
			return 0, parseErr
		}
		return int32(p), nil
	})
	if err != nil {
		return nil, err
	}
	detailedParam, err := parseOptional(qp.Get("detailed"), strconv.ParseBool)
	if err != nil {
		return nil, err
	}
	var options *azspark.SessionClientGetSparkSessionOptions
	if detailedParam != nil {
		options = &azspark.SessionClientGetSparkSessionOptions{
			Detailed: detailedParam,
		}
	}
	respr, errRespr := s.srv.GetSparkSession(req.Context(), int32(sessionIDParam), options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Session, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SessionServerTransport) dispatchGetSparkSessions(req *http.Request) (*http.Response, error) {
	if s.srv.GetSparkSessions == nil {
		return nil, &nonRetriableError{errors.New("method GetSparkSessions not implemented")}
	}
	const regexStr = "/livyApi/versions/(?P<livyApiVersion>[a-zA-Z0-9-_]+)/sparkPools/(?P<sparkPoolName>[a-zA-Z0-9-_]+)/sessions"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	fromParam, err := parseOptional(qp.Get("from"), func(v string) (int32, error) {
		p, parseErr := strconv.ParseInt(v, 10, 32)
		if parseErr != nil {
			return 0, parseErr
		}
		return int32(p), nil
	})
	if err != nil {
		return nil, err
	}
	sizeParam, err := parseOptional(qp.Get("size"), func(v string) (int32, error) {
		p, parseErr := strconv.ParseInt(v, 10, 32)
		if parseErr != nil {
			return 0, parseErr
		}
		return int32(p), nil
	})
	if err != nil {
		return nil, err
	}
	detailedParam, err := parseOptional(qp.Get("detailed"), strconv.ParseBool)
	if err != nil {
		return nil, err
	}
	var options *azspark.SessionClientGetSparkSessionsOptions
	if fromParam != nil || sizeParam != nil || detailedParam != nil {
		options = &azspark.SessionClientGetSparkSessionsOptions{
			From:     fromParam,
			Size:     sizeParam,
			Detailed: detailedParam,
		}
	}
	respr, errRespr := s.srv.GetSparkSessions(req.Context(), options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SessionCollection, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SessionServerTransport) dispatchGetSparkStatement(req *http.Request) (*http.Response, error) {
	if s.srv.GetSparkStatement == nil {
		return nil, &nonRetriableError{errors.New("method GetSparkStatement not implemented")}
	}
	const regexStr = "/livyApi/versions/(?P<livyApiVersion>[a-zA-Z0-9-_]+)/sparkPools/(?P<sparkPoolName>[a-zA-Z0-9-_]+)/sessions/(?P<sessionId>[a-zA-Z0-9-_]+)/statements/(?P<statementId>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	sessionIDParam, err := parseWithCast(matches[regex.SubexpIndex("sessionId")], func(v string) (int32, error) {
		p, parseErr := strconv.ParseInt(v, 10, 32)
		if parseErr != nil {
			return 0, parseErr
		}
		return int32(p), nil
	})
	if err != nil {
		return nil, err
	}
	statementIDParam, err := parseWithCast(matches[regex.SubexpIndex("statementId")], func(v string) (int32, error) {
		p, parseErr := strconv.ParseInt(v, 10, 32)
		if parseErr != nil {
			return 0, parseErr
		}
		return int32(p), nil
	})
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.GetSparkStatement(req.Context(), int32(sessionIDParam), int32(statementIDParam), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Statement, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SessionServerTransport) dispatchGetSparkStatements(req *http.Request) (*http.Response, error) {
	if s.srv.GetSparkStatements == nil {
		return nil, &nonRetriableError{errors.New("method GetSparkStatements not implemented")}
	}
	const regexStr = "/livyApi/versions/(?P<livyApiVersion>[a-zA-Z0-9-_]+)/sparkPools/(?P<sparkPoolName>[a-zA-Z0-9-_]+)/sessions/(?P<sessionId>[a-zA-Z0-9-_]+)/statements"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	sessionIDParam, err := parseWithCast(matches[regex.SubexpIndex("sessionId")], func(v string) (int32, error) {
		p, parseErr := strconv.ParseInt(v, 10, 32)
		if parseErr != nil {
			return 0, parseErr
		}
		return int32(p), nil
	})
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.GetSparkStatements(req.Context(), int32(sessionIDParam), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).StatementCollection, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SessionServerTransport) dispatchResetSparkSessionTimeout(req *http.Request) (*http.Response, error) {
	if s.srv.ResetSparkSessionTimeout == nil {
		return nil, &nonRetriableError{errors.New("method ResetSparkSessionTimeout not implemented")}
	}
	const regexStr = "/livyApi/versions/(?P<livyApiVersion>[a-zA-Z0-9-_]+)/sparkPools/(?P<sparkPoolName>[a-zA-Z0-9-_]+)/sessions/(?P<sessionId>[a-zA-Z0-9-_]+)/reset-timeout"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	sessionIDParam, err := parseWithCast(matches[regex.SubexpIndex("sessionId")], func(v string) (int32, error) {
		p, parseErr := strconv.ParseInt(v, 10, 32)
		if parseErr != nil {
			return 0, parseErr
		}
		return int32(p), nil
	})
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.ResetSparkSessionTimeout(req.Context(), int32(sessionIDParam), nil)
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
	return resp, nil
}
