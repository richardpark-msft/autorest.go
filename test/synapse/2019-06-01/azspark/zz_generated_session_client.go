//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azspark

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type sessionClient struct {
	endpoint string
	pl       runtime.Pipeline
}

// newSessionClient creates a new instance of sessionClient with the specified values.
// endpoint - The workspace development endpoint, for example https://myworkspace.dev.azuresynapse.net.
// livyAPIVersion - Valid api-version for the request.
// sparkPoolName - Name of the spark pool.
// pl - the pipeline used for sending requests and handling responses.
func newSessionClient(endpoint string, livyAPIVersion *string, sparkPoolName string, pl runtime.Pipeline) *sessionClient {
	hostURL := "{endpoint}/livyApi/versions/{livyApiVersion}/sparkPools/{sparkPoolName}"
	hostURL = strings.ReplaceAll(hostURL, "{endpoint}", endpoint)
	if livyAPIVersion == nil {
		defaultValue := "2019-11-01-preview"
		livyAPIVersion = &defaultValue
	}
	hostURL = strings.ReplaceAll(hostURL, "{livyApiVersion}", *livyAPIVersion)
	hostURL = strings.ReplaceAll(hostURL, "{sparkPoolName}", sparkPoolName)
	client := &sessionClient{
		endpoint: hostURL,
		pl:       pl,
	}
	return client
}

// CancelSparkSession - Cancels a running spark session.
// If the operation fails it returns an *azcore.ResponseError type.
// sessionID - Identifier for the session.
// options - sessionClientCancelSparkSessionOptions contains the optional parameters for the sessionClient.CancelSparkSession
// method.
func (client *sessionClient) CancelSparkSession(ctx context.Context, sessionID int32, options *sessionClientCancelSparkSessionOptions) (sessionClientCancelSparkSessionResponse, error) {
	req, err := client.cancelSparkSessionCreateRequest(ctx, sessionID, options)
	if err != nil {
		return sessionClientCancelSparkSessionResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return sessionClientCancelSparkSessionResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return sessionClientCancelSparkSessionResponse{}, runtime.NewResponseError(resp)
	}
	return sessionClientCancelSparkSessionResponse{}, nil
}

// cancelSparkSessionCreateRequest creates the CancelSparkSession request.
func (client *sessionClient) cancelSparkSessionCreateRequest(ctx context.Context, sessionID int32, options *sessionClientCancelSparkSessionOptions) (*policy.Request, error) {
	urlPath := "/sessions/{sessionId}"
	urlPath = strings.ReplaceAll(urlPath, "{sessionId}", url.PathEscape(strconv.FormatInt(int64(sessionID), 10)))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

// CancelSparkStatement - Kill a statement within a session.
// If the operation fails it returns an *azcore.ResponseError type.
// sessionID - Identifier for the session.
// statementID - Identifier for the statement.
// options - sessionClientCancelSparkStatementOptions contains the optional parameters for the sessionClient.CancelSparkStatement
// method.
func (client *sessionClient) CancelSparkStatement(ctx context.Context, sessionID int32, statementID int32, options *sessionClientCancelSparkStatementOptions) (sessionClientCancelSparkStatementResponse, error) {
	req, err := client.cancelSparkStatementCreateRequest(ctx, sessionID, statementID, options)
	if err != nil {
		return sessionClientCancelSparkStatementResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return sessionClientCancelSparkStatementResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return sessionClientCancelSparkStatementResponse{}, runtime.NewResponseError(resp)
	}
	return client.cancelSparkStatementHandleResponse(resp)
}

// cancelSparkStatementCreateRequest creates the CancelSparkStatement request.
func (client *sessionClient) cancelSparkStatementCreateRequest(ctx context.Context, sessionID int32, statementID int32, options *sessionClientCancelSparkStatementOptions) (*policy.Request, error) {
	urlPath := "/sessions/{sessionId}/statements/{statementId}/cancel"
	urlPath = strings.ReplaceAll(urlPath, "{sessionId}", url.PathEscape(strconv.FormatInt(int64(sessionID), 10)))
	urlPath = strings.ReplaceAll(urlPath, "{statementId}", url.PathEscape(strconv.FormatInt(int64(statementID), 10)))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// cancelSparkStatementHandleResponse handles the CancelSparkStatement response.
func (client *sessionClient) cancelSparkStatementHandleResponse(resp *http.Response) (sessionClientCancelSparkStatementResponse, error) {
	result := sessionClientCancelSparkStatementResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.StatementCancellationResult); err != nil {
		return sessionClientCancelSparkStatementResponse{}, err
	}
	return result, nil
}

// CreateSparkSession - Create new spark session.
// If the operation fails it returns an *azcore.ResponseError type.
// sparkSessionOptions - Livy compatible batch job request payload.
// options - sessionClientCreateSparkSessionOptions contains the optional parameters for the sessionClient.CreateSparkSession
// method.
func (client *sessionClient) CreateSparkSession(ctx context.Context, sparkSessionOptions SessionOptions, options *sessionClientCreateSparkSessionOptions) (sessionClientCreateSparkSessionResponse, error) {
	req, err := client.createSparkSessionCreateRequest(ctx, sparkSessionOptions, options)
	if err != nil {
		return sessionClientCreateSparkSessionResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return sessionClientCreateSparkSessionResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return sessionClientCreateSparkSessionResponse{}, runtime.NewResponseError(resp)
	}
	return client.createSparkSessionHandleResponse(resp)
}

// createSparkSessionCreateRequest creates the CreateSparkSession request.
func (client *sessionClient) createSparkSessionCreateRequest(ctx context.Context, sparkSessionOptions SessionOptions, options *sessionClientCreateSparkSessionOptions) (*policy.Request, error) {
	urlPath := "/sessions"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Detailed != nil {
		reqQP.Set("detailed", strconv.FormatBool(*options.Detailed))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, sparkSessionOptions)
}

// createSparkSessionHandleResponse handles the CreateSparkSession response.
func (client *sessionClient) createSparkSessionHandleResponse(resp *http.Response) (sessionClientCreateSparkSessionResponse, error) {
	result := sessionClientCreateSparkSessionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Session); err != nil {
		return sessionClientCreateSparkSessionResponse{}, err
	}
	return result, nil
}

// CreateSparkStatement - Create statement within a spark session.
// If the operation fails it returns an *azcore.ResponseError type.
// sessionID - Identifier for the session.
// sparkStatementOptions - Livy compatible batch job request payload.
// options - sessionClientCreateSparkStatementOptions contains the optional parameters for the sessionClient.CreateSparkStatement
// method.
func (client *sessionClient) CreateSparkStatement(ctx context.Context, sessionID int32, sparkStatementOptions StatementOptions, options *sessionClientCreateSparkStatementOptions) (sessionClientCreateSparkStatementResponse, error) {
	req, err := client.createSparkStatementCreateRequest(ctx, sessionID, sparkStatementOptions, options)
	if err != nil {
		return sessionClientCreateSparkStatementResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return sessionClientCreateSparkStatementResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return sessionClientCreateSparkStatementResponse{}, runtime.NewResponseError(resp)
	}
	return client.createSparkStatementHandleResponse(resp)
}

// createSparkStatementCreateRequest creates the CreateSparkStatement request.
func (client *sessionClient) createSparkStatementCreateRequest(ctx context.Context, sessionID int32, sparkStatementOptions StatementOptions, options *sessionClientCreateSparkStatementOptions) (*policy.Request, error) {
	urlPath := "/sessions/{sessionId}/statements"
	urlPath = strings.ReplaceAll(urlPath, "{sessionId}", url.PathEscape(strconv.FormatInt(int64(sessionID), 10)))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, sparkStatementOptions)
}

// createSparkStatementHandleResponse handles the CreateSparkStatement response.
func (client *sessionClient) createSparkStatementHandleResponse(resp *http.Response) (sessionClientCreateSparkStatementResponse, error) {
	result := sessionClientCreateSparkStatementResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Statement); err != nil {
		return sessionClientCreateSparkStatementResponse{}, err
	}
	return result, nil
}

// GetSparkSession - Gets a single spark session.
// If the operation fails it returns an *azcore.ResponseError type.
// sessionID - Identifier for the session.
// options - sessionClientGetSparkSessionOptions contains the optional parameters for the sessionClient.GetSparkSession method.
func (client *sessionClient) GetSparkSession(ctx context.Context, sessionID int32, options *sessionClientGetSparkSessionOptions) (sessionClientGetSparkSessionResponse, error) {
	req, err := client.getSparkSessionCreateRequest(ctx, sessionID, options)
	if err != nil {
		return sessionClientGetSparkSessionResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return sessionClientGetSparkSessionResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return sessionClientGetSparkSessionResponse{}, runtime.NewResponseError(resp)
	}
	return client.getSparkSessionHandleResponse(resp)
}

// getSparkSessionCreateRequest creates the GetSparkSession request.
func (client *sessionClient) getSparkSessionCreateRequest(ctx context.Context, sessionID int32, options *sessionClientGetSparkSessionOptions) (*policy.Request, error) {
	urlPath := "/sessions/{sessionId}"
	urlPath = strings.ReplaceAll(urlPath, "{sessionId}", url.PathEscape(strconv.FormatInt(int64(sessionID), 10)))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Detailed != nil {
		reqQP.Set("detailed", strconv.FormatBool(*options.Detailed))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getSparkSessionHandleResponse handles the GetSparkSession response.
func (client *sessionClient) getSparkSessionHandleResponse(resp *http.Response) (sessionClientGetSparkSessionResponse, error) {
	result := sessionClientGetSparkSessionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Session); err != nil {
		return sessionClientGetSparkSessionResponse{}, err
	}
	return result, nil
}

// GetSparkSessions - List all spark sessions which are running under a particular spark pool.
// If the operation fails it returns an *azcore.ResponseError type.
// options - sessionClientGetSparkSessionsOptions contains the optional parameters for the sessionClient.GetSparkSessions
// method.
func (client *sessionClient) GetSparkSessions(ctx context.Context, options *sessionClientGetSparkSessionsOptions) (sessionClientGetSparkSessionsResponse, error) {
	req, err := client.getSparkSessionsCreateRequest(ctx, options)
	if err != nil {
		return sessionClientGetSparkSessionsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return sessionClientGetSparkSessionsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return sessionClientGetSparkSessionsResponse{}, runtime.NewResponseError(resp)
	}
	return client.getSparkSessionsHandleResponse(resp)
}

// getSparkSessionsCreateRequest creates the GetSparkSessions request.
func (client *sessionClient) getSparkSessionsCreateRequest(ctx context.Context, options *sessionClientGetSparkSessionsOptions) (*policy.Request, error) {
	urlPath := "/sessions"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.From != nil {
		reqQP.Set("from", strconv.FormatInt(int64(*options.From), 10))
	}
	if options != nil && options.Size != nil {
		reqQP.Set("size", strconv.FormatInt(int64(*options.Size), 10))
	}
	if options != nil && options.Detailed != nil {
		reqQP.Set("detailed", strconv.FormatBool(*options.Detailed))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getSparkSessionsHandleResponse handles the GetSparkSessions response.
func (client *sessionClient) getSparkSessionsHandleResponse(resp *http.Response) (sessionClientGetSparkSessionsResponse, error) {
	result := sessionClientGetSparkSessionsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SessionCollection); err != nil {
		return sessionClientGetSparkSessionsResponse{}, err
	}
	return result, nil
}

// GetSparkStatement - Gets a single statement within a spark session.
// If the operation fails it returns an *azcore.ResponseError type.
// sessionID - Identifier for the session.
// statementID - Identifier for the statement.
// options - sessionClientGetSparkStatementOptions contains the optional parameters for the sessionClient.GetSparkStatement
// method.
func (client *sessionClient) GetSparkStatement(ctx context.Context, sessionID int32, statementID int32, options *sessionClientGetSparkStatementOptions) (sessionClientGetSparkStatementResponse, error) {
	req, err := client.getSparkStatementCreateRequest(ctx, sessionID, statementID, options)
	if err != nil {
		return sessionClientGetSparkStatementResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return sessionClientGetSparkStatementResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return sessionClientGetSparkStatementResponse{}, runtime.NewResponseError(resp)
	}
	return client.getSparkStatementHandleResponse(resp)
}

// getSparkStatementCreateRequest creates the GetSparkStatement request.
func (client *sessionClient) getSparkStatementCreateRequest(ctx context.Context, sessionID int32, statementID int32, options *sessionClientGetSparkStatementOptions) (*policy.Request, error) {
	urlPath := "/sessions/{sessionId}/statements/{statementId}"
	urlPath = strings.ReplaceAll(urlPath, "{sessionId}", url.PathEscape(strconv.FormatInt(int64(sessionID), 10)))
	urlPath = strings.ReplaceAll(urlPath, "{statementId}", url.PathEscape(strconv.FormatInt(int64(statementID), 10)))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getSparkStatementHandleResponse handles the GetSparkStatement response.
func (client *sessionClient) getSparkStatementHandleResponse(resp *http.Response) (sessionClientGetSparkStatementResponse, error) {
	result := sessionClientGetSparkStatementResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Statement); err != nil {
		return sessionClientGetSparkStatementResponse{}, err
	}
	return result, nil
}

// GetSparkStatements - Gets a list of statements within a spark session.
// If the operation fails it returns an *azcore.ResponseError type.
// sessionID - Identifier for the session.
// options - sessionClientGetSparkStatementsOptions contains the optional parameters for the sessionClient.GetSparkStatements
// method.
func (client *sessionClient) GetSparkStatements(ctx context.Context, sessionID int32, options *sessionClientGetSparkStatementsOptions) (sessionClientGetSparkStatementsResponse, error) {
	req, err := client.getSparkStatementsCreateRequest(ctx, sessionID, options)
	if err != nil {
		return sessionClientGetSparkStatementsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return sessionClientGetSparkStatementsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return sessionClientGetSparkStatementsResponse{}, runtime.NewResponseError(resp)
	}
	return client.getSparkStatementsHandleResponse(resp)
}

// getSparkStatementsCreateRequest creates the GetSparkStatements request.
func (client *sessionClient) getSparkStatementsCreateRequest(ctx context.Context, sessionID int32, options *sessionClientGetSparkStatementsOptions) (*policy.Request, error) {
	urlPath := "/sessions/{sessionId}/statements"
	urlPath = strings.ReplaceAll(urlPath, "{sessionId}", url.PathEscape(strconv.FormatInt(int64(sessionID), 10)))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getSparkStatementsHandleResponse handles the GetSparkStatements response.
func (client *sessionClient) getSparkStatementsHandleResponse(resp *http.Response) (sessionClientGetSparkStatementsResponse, error) {
	result := sessionClientGetSparkStatementsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.StatementCollection); err != nil {
		return sessionClientGetSparkStatementsResponse{}, err
	}
	return result, nil
}

// ResetSparkSessionTimeout - Sends a keep alive call to the current session to reset the session timeout.
// If the operation fails it returns an *azcore.ResponseError type.
// sessionID - Identifier for the session.
// options - sessionClientResetSparkSessionTimeoutOptions contains the optional parameters for the sessionClient.ResetSparkSessionTimeout
// method.
func (client *sessionClient) ResetSparkSessionTimeout(ctx context.Context, sessionID int32, options *sessionClientResetSparkSessionTimeoutOptions) (sessionClientResetSparkSessionTimeoutResponse, error) {
	req, err := client.resetSparkSessionTimeoutCreateRequest(ctx, sessionID, options)
	if err != nil {
		return sessionClientResetSparkSessionTimeoutResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return sessionClientResetSparkSessionTimeoutResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return sessionClientResetSparkSessionTimeoutResponse{}, runtime.NewResponseError(resp)
	}
	return sessionClientResetSparkSessionTimeoutResponse{}, nil
}

// resetSparkSessionTimeoutCreateRequest creates the ResetSparkSessionTimeout request.
func (client *sessionClient) resetSparkSessionTimeoutCreateRequest(ctx context.Context, sessionID int32, options *sessionClientResetSparkSessionTimeoutOptions) (*policy.Request, error) {
	urlPath := "/sessions/{sessionId}/reset-timeout"
	urlPath = strings.ReplaceAll(urlPath, "{sessionId}", url.PathEscape(strconv.FormatInt(int64(sessionID), 10)))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}
