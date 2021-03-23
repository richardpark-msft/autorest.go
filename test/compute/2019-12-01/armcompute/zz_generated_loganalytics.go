// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// LogAnalyticsClient contains the methods for the LogAnalytics group.
// Don't use this type directly, use NewLogAnalyticsClient() instead.
type LogAnalyticsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewLogAnalyticsClient creates a new instance of LogAnalyticsClient with the specified values.
func NewLogAnalyticsClient(con *armcore.Connection, subscriptionID string) *LogAnalyticsClient {
	return &LogAnalyticsClient{con: con, subscriptionID: subscriptionID}
}

// BeginExportRequestRateByInterval - Export logs that show Api requests made by this subscription in the given time window to show throttling activities.
func (client *LogAnalyticsClient) BeginExportRequestRateByInterval(ctx context.Context, location string, parameters RequestRateByIntervalInput, options *LogAnalyticsBeginExportRequestRateByIntervalOptions) (LogAnalyticsOperationResultPollerResponse, error) {
	resp, err := client.exportRequestRateByInterval(ctx, location, parameters, options)
	if err != nil {
		return LogAnalyticsOperationResultPollerResponse{}, err
	}
	result := LogAnalyticsOperationResultPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("LogAnalyticsClient.ExportRequestRateByInterval", "azure-async-operation", resp, client.exportRequestRateByIntervalHandleError)
	if err != nil {
		return LogAnalyticsOperationResultPollerResponse{}, err
	}
	poller := &logAnalyticsOperationResultPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (LogAnalyticsOperationResultResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeExportRequestRateByInterval creates a new LogAnalyticsOperationResultPoller from the specified resume token.
// token - The value must come from a previous call to LogAnalyticsOperationResultPoller.ResumeToken().
func (client *LogAnalyticsClient) ResumeExportRequestRateByInterval(token string) (LogAnalyticsOperationResultPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("LogAnalyticsClient.ExportRequestRateByInterval", token, client.exportRequestRateByIntervalHandleError)
	if err != nil {
		return nil, err
	}
	return &logAnalyticsOperationResultPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// ExportRequestRateByInterval - Export logs that show Api requests made by this subscription in the given time window to show throttling activities.
func (client *LogAnalyticsClient) exportRequestRateByInterval(ctx context.Context, location string, parameters RequestRateByIntervalInput, options *LogAnalyticsBeginExportRequestRateByIntervalOptions) (*azcore.Response, error) {
	req, err := client.exportRequestRateByIntervalCreateRequest(ctx, location, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.exportRequestRateByIntervalHandleError(resp)
	}
	return resp, nil
}

// exportRequestRateByIntervalCreateRequest creates the ExportRequestRateByInterval request.
func (client *LogAnalyticsClient) exportRequestRateByIntervalCreateRequest(ctx context.Context, location string, parameters RequestRateByIntervalInput, options *LogAnalyticsBeginExportRequestRateByIntervalOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/logAnalytics/apiAccess/getRequestRateByInterval"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-12-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// exportRequestRateByIntervalHandleResponse handles the ExportRequestRateByInterval response.
func (client *LogAnalyticsClient) exportRequestRateByIntervalHandleResponse(resp *azcore.Response) (LogAnalyticsOperationResultResponse, error) {
	var val *LogAnalyticsOperationResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return LogAnalyticsOperationResultResponse{}, err
	}
	return LogAnalyticsOperationResultResponse{RawResponse: resp.Response, LogAnalyticsOperationResult: val}, nil
}

// exportRequestRateByIntervalHandleError handles the ExportRequestRateByInterval error response.
func (client *LogAnalyticsClient) exportRequestRateByIntervalHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// BeginExportThrottledRequests - Export logs that show total throttled Api requests for this subscription in the given time window.
func (client *LogAnalyticsClient) BeginExportThrottledRequests(ctx context.Context, location string, parameters ThrottledRequestsInput, options *LogAnalyticsBeginExportThrottledRequestsOptions) (LogAnalyticsOperationResultPollerResponse, error) {
	resp, err := client.exportThrottledRequests(ctx, location, parameters, options)
	if err != nil {
		return LogAnalyticsOperationResultPollerResponse{}, err
	}
	result := LogAnalyticsOperationResultPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("LogAnalyticsClient.ExportThrottledRequests", "azure-async-operation", resp, client.exportThrottledRequestsHandleError)
	if err != nil {
		return LogAnalyticsOperationResultPollerResponse{}, err
	}
	poller := &logAnalyticsOperationResultPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (LogAnalyticsOperationResultResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeExportThrottledRequests creates a new LogAnalyticsOperationResultPoller from the specified resume token.
// token - The value must come from a previous call to LogAnalyticsOperationResultPoller.ResumeToken().
func (client *LogAnalyticsClient) ResumeExportThrottledRequests(token string) (LogAnalyticsOperationResultPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("LogAnalyticsClient.ExportThrottledRequests", token, client.exportThrottledRequestsHandleError)
	if err != nil {
		return nil, err
	}
	return &logAnalyticsOperationResultPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// ExportThrottledRequests - Export logs that show total throttled Api requests for this subscription in the given time window.
func (client *LogAnalyticsClient) exportThrottledRequests(ctx context.Context, location string, parameters ThrottledRequestsInput, options *LogAnalyticsBeginExportThrottledRequestsOptions) (*azcore.Response, error) {
	req, err := client.exportThrottledRequestsCreateRequest(ctx, location, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.exportThrottledRequestsHandleError(resp)
	}
	return resp, nil
}

// exportThrottledRequestsCreateRequest creates the ExportThrottledRequests request.
func (client *LogAnalyticsClient) exportThrottledRequestsCreateRequest(ctx context.Context, location string, parameters ThrottledRequestsInput, options *LogAnalyticsBeginExportThrottledRequestsOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/logAnalytics/apiAccess/getThrottledRequests"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-12-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// exportThrottledRequestsHandleResponse handles the ExportThrottledRequests response.
func (client *LogAnalyticsClient) exportThrottledRequestsHandleResponse(resp *azcore.Response) (LogAnalyticsOperationResultResponse, error) {
	var val *LogAnalyticsOperationResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return LogAnalyticsOperationResultResponse{}, err
	}
	return LogAnalyticsOperationResultResponse{RawResponse: resp.Response, LogAnalyticsOperationResult: val}, nil
}

// exportThrottledRequestsHandleError handles the ExportThrottledRequests error response.
func (client *LogAnalyticsClient) exportThrottledRequestsHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}
