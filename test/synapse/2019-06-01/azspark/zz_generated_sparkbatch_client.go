//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azspark

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type sparkBatchClient struct {
	con *connection
}

// CancelSparkBatchJob - Cancels a running spark batch job.
// If the operation fails it returns a generic error.
func (client *sparkBatchClient) CancelSparkBatchJob(ctx context.Context, batchID int32, options *SparkBatchCancelSparkBatchJobOptions) (SparkBatchCancelSparkBatchJobResponse, error) {
	req, err := client.cancelSparkBatchJobCreateRequest(ctx, batchID, options)
	if err != nil {
		return SparkBatchCancelSparkBatchJobResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return SparkBatchCancelSparkBatchJobResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SparkBatchCancelSparkBatchJobResponse{}, client.cancelSparkBatchJobHandleError(resp)
	}
	return SparkBatchCancelSparkBatchJobResponse{RawResponse: resp}, nil
}

// cancelSparkBatchJobCreateRequest creates the CancelSparkBatchJob request.
func (client *sparkBatchClient) cancelSparkBatchJobCreateRequest(ctx context.Context, batchID int32, options *SparkBatchCancelSparkBatchJobOptions) (*policy.Request, error) {
	urlPath := "/batches/{batchId}"
	urlPath = strings.ReplaceAll(urlPath, "{batchId}", url.PathEscape(strconv.FormatInt(int64(batchID), 10)))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

// cancelSparkBatchJobHandleError handles the CancelSparkBatchJob error response.
func (client *sparkBatchClient) cancelSparkBatchJobHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// CreateSparkBatchJob - Create new spark batch job.
// If the operation fails it returns a generic error.
func (client *sparkBatchClient) CreateSparkBatchJob(ctx context.Context, sparkBatchJobOptions SparkBatchJobOptions, options *SparkBatchCreateSparkBatchJobOptions) (SparkBatchCreateSparkBatchJobResponse, error) {
	req, err := client.createSparkBatchJobCreateRequest(ctx, sparkBatchJobOptions, options)
	if err != nil {
		return SparkBatchCreateSparkBatchJobResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return SparkBatchCreateSparkBatchJobResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SparkBatchCreateSparkBatchJobResponse{}, client.createSparkBatchJobHandleError(resp)
	}
	return client.createSparkBatchJobHandleResponse(resp)
}

// createSparkBatchJobCreateRequest creates the CreateSparkBatchJob request.
func (client *sparkBatchClient) createSparkBatchJobCreateRequest(ctx context.Context, sparkBatchJobOptions SparkBatchJobOptions, options *SparkBatchCreateSparkBatchJobOptions) (*policy.Request, error) {
	urlPath := "/batches"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Detailed != nil {
		reqQP.Set("detailed", strconv.FormatBool(*options.Detailed))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, sparkBatchJobOptions)
}

// createSparkBatchJobHandleResponse handles the CreateSparkBatchJob response.
func (client *sparkBatchClient) createSparkBatchJobHandleResponse(resp *http.Response) (SparkBatchCreateSparkBatchJobResponse, error) {
	result := SparkBatchCreateSparkBatchJobResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SparkBatchJob); err != nil {
		return SparkBatchCreateSparkBatchJobResponse{}, err
	}
	return result, nil
}

// createSparkBatchJobHandleError handles the CreateSparkBatchJob error response.
func (client *sparkBatchClient) createSparkBatchJobHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// GetSparkBatchJob - Gets a single spark batch job.
// If the operation fails it returns a generic error.
func (client *sparkBatchClient) GetSparkBatchJob(ctx context.Context, batchID int32, options *SparkBatchGetSparkBatchJobOptions) (SparkBatchGetSparkBatchJobResponse, error) {
	req, err := client.getSparkBatchJobCreateRequest(ctx, batchID, options)
	if err != nil {
		return SparkBatchGetSparkBatchJobResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return SparkBatchGetSparkBatchJobResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SparkBatchGetSparkBatchJobResponse{}, client.getSparkBatchJobHandleError(resp)
	}
	return client.getSparkBatchJobHandleResponse(resp)
}

// getSparkBatchJobCreateRequest creates the GetSparkBatchJob request.
func (client *sparkBatchClient) getSparkBatchJobCreateRequest(ctx context.Context, batchID int32, options *SparkBatchGetSparkBatchJobOptions) (*policy.Request, error) {
	urlPath := "/batches/{batchId}"
	urlPath = strings.ReplaceAll(urlPath, "{batchId}", url.PathEscape(strconv.FormatInt(int64(batchID), 10)))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.con.Endpoint(), urlPath))
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

// getSparkBatchJobHandleResponse handles the GetSparkBatchJob response.
func (client *sparkBatchClient) getSparkBatchJobHandleResponse(resp *http.Response) (SparkBatchGetSparkBatchJobResponse, error) {
	result := SparkBatchGetSparkBatchJobResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SparkBatchJob); err != nil {
		return SparkBatchGetSparkBatchJobResponse{}, err
	}
	return result, nil
}

// getSparkBatchJobHandleError handles the GetSparkBatchJob error response.
func (client *sparkBatchClient) getSparkBatchJobHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// GetSparkBatchJobs - List all spark batch jobs which are running under a particular spark pool.
// If the operation fails it returns a generic error.
func (client *sparkBatchClient) GetSparkBatchJobs(ctx context.Context, options *SparkBatchGetSparkBatchJobsOptions) (SparkBatchGetSparkBatchJobsResponse, error) {
	req, err := client.getSparkBatchJobsCreateRequest(ctx, options)
	if err != nil {
		return SparkBatchGetSparkBatchJobsResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return SparkBatchGetSparkBatchJobsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SparkBatchGetSparkBatchJobsResponse{}, client.getSparkBatchJobsHandleError(resp)
	}
	return client.getSparkBatchJobsHandleResponse(resp)
}

// getSparkBatchJobsCreateRequest creates the GetSparkBatchJobs request.
func (client *sparkBatchClient) getSparkBatchJobsCreateRequest(ctx context.Context, options *SparkBatchGetSparkBatchJobsOptions) (*policy.Request, error) {
	urlPath := "/batches"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.con.Endpoint(), urlPath))
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

// getSparkBatchJobsHandleResponse handles the GetSparkBatchJobs response.
func (client *sparkBatchClient) getSparkBatchJobsHandleResponse(resp *http.Response) (SparkBatchGetSparkBatchJobsResponse, error) {
	result := SparkBatchGetSparkBatchJobsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SparkBatchJobCollection); err != nil {
		return SparkBatchGetSparkBatchJobsResponse{}, err
	}
	return result, nil
}

// getSparkBatchJobsHandleError handles the GetSparkBatchJobs error response.
func (client *sparkBatchClient) getSparkBatchJobsHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
