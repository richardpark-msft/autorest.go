//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azspark

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// BatchClient contains the methods for the SparkBatch group.
// Don't use this type directly, use a constructor function instead.
type BatchClient struct {
	internal       *azcore.Client
	endpoint       string
	livyAPIVersion string
	sparkPoolName  string
}

// CancelSparkBatchJob - Cancels a running spark batch job.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-12-01
//   - batchID - Identifier for the batch job.
//   - options - BatchClientCancelSparkBatchJobOptions contains the optional parameters for the BatchClient.CancelSparkBatchJob
//     method.
func (client *BatchClient) CancelSparkBatchJob(ctx context.Context, batchID int32, options *BatchClientCancelSparkBatchJobOptions) (BatchClientCancelSparkBatchJobResponse, error) {
	var err error
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "BatchClient.CancelSparkBatchJob")
	req, err := client.cancelSparkBatchJobCreateRequest(ctx, batchID, options)
	if err != nil {
		return BatchClientCancelSparkBatchJobResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BatchClientCancelSparkBatchJobResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return BatchClientCancelSparkBatchJobResponse{}, err
	}
	return BatchClientCancelSparkBatchJobResponse{}, nil
}

// cancelSparkBatchJobCreateRequest creates the CancelSparkBatchJob request.
func (client *BatchClient) cancelSparkBatchJobCreateRequest(ctx context.Context, batchID int32, options *BatchClientCancelSparkBatchJobOptions) (*policy.Request, error) {
	urlPath := "/livyApi/versions/{livyApiVersion}/sparkPools/{sparkPoolName}/batches/{batchId}"
	urlPath = strings.ReplaceAll(urlPath, "{livyApiVersion}", client.livyAPIVersion)
	urlPath = strings.ReplaceAll(urlPath, "{sparkPoolName}", client.sparkPoolName)
	urlPath = strings.ReplaceAll(urlPath, "{batchId}", url.PathEscape(strconv.FormatInt(int64(batchID), 10)))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

// CreateSparkBatchJob - Create new spark batch job.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-12-01
//   - sparkBatchJobOptions - Livy compatible batch job request payload.
//   - options - BatchClientCreateSparkBatchJobOptions contains the optional parameters for the BatchClient.CreateSparkBatchJob
//     method.
func (client *BatchClient) CreateSparkBatchJob(ctx context.Context, sparkBatchJobOptions BatchJobOptions, options *BatchClientCreateSparkBatchJobOptions) (BatchClientCreateSparkBatchJobResponse, error) {
	var err error
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "BatchClient.CreateSparkBatchJob")
	req, err := client.createSparkBatchJobCreateRequest(ctx, sparkBatchJobOptions, options)
	if err != nil {
		return BatchClientCreateSparkBatchJobResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BatchClientCreateSparkBatchJobResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return BatchClientCreateSparkBatchJobResponse{}, err
	}
	resp, err := client.createSparkBatchJobHandleResponse(httpResp)
	return resp, err
}

// createSparkBatchJobCreateRequest creates the CreateSparkBatchJob request.
func (client *BatchClient) createSparkBatchJobCreateRequest(ctx context.Context, sparkBatchJobOptions BatchJobOptions, options *BatchClientCreateSparkBatchJobOptions) (*policy.Request, error) {
	urlPath := "/livyApi/versions/{livyApiVersion}/sparkPools/{sparkPoolName}/batches"
	urlPath = strings.ReplaceAll(urlPath, "{livyApiVersion}", client.livyAPIVersion)
	urlPath = strings.ReplaceAll(urlPath, "{sparkPoolName}", client.sparkPoolName)
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Detailed != nil {
		reqQP.Set("detailed", strconv.FormatBool(*options.Detailed))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, sparkBatchJobOptions); err != nil {
		return nil, err
	}
	return req, nil
}

// createSparkBatchJobHandleResponse handles the CreateSparkBatchJob response.
func (client *BatchClient) createSparkBatchJobHandleResponse(resp *http.Response) (BatchClientCreateSparkBatchJobResponse, error) {
	result := BatchClientCreateSparkBatchJobResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BatchJob); err != nil {
		return BatchClientCreateSparkBatchJobResponse{}, err
	}
	return result, nil
}

// GetSparkBatchJob - Gets a single spark batch job.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-12-01
//   - batchID - Identifier for the batch job.
//   - options - BatchClientGetSparkBatchJobOptions contains the optional parameters for the BatchClient.GetSparkBatchJob method.
func (client *BatchClient) GetSparkBatchJob(ctx context.Context, batchID int32, options *BatchClientGetSparkBatchJobOptions) (BatchClientGetSparkBatchJobResponse, error) {
	var err error
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "BatchClient.GetSparkBatchJob")
	req, err := client.getSparkBatchJobCreateRequest(ctx, batchID, options)
	if err != nil {
		return BatchClientGetSparkBatchJobResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BatchClientGetSparkBatchJobResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return BatchClientGetSparkBatchJobResponse{}, err
	}
	resp, err := client.getSparkBatchJobHandleResponse(httpResp)
	return resp, err
}

// getSparkBatchJobCreateRequest creates the GetSparkBatchJob request.
func (client *BatchClient) getSparkBatchJobCreateRequest(ctx context.Context, batchID int32, options *BatchClientGetSparkBatchJobOptions) (*policy.Request, error) {
	urlPath := "/livyApi/versions/{livyApiVersion}/sparkPools/{sparkPoolName}/batches/{batchId}"
	urlPath = strings.ReplaceAll(urlPath, "{livyApiVersion}", client.livyAPIVersion)
	urlPath = strings.ReplaceAll(urlPath, "{sparkPoolName}", client.sparkPoolName)
	urlPath = strings.ReplaceAll(urlPath, "{batchId}", url.PathEscape(strconv.FormatInt(int64(batchID), 10)))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Detailed != nil {
		reqQP.Set("detailed", strconv.FormatBool(*options.Detailed))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getSparkBatchJobHandleResponse handles the GetSparkBatchJob response.
func (client *BatchClient) getSparkBatchJobHandleResponse(resp *http.Response) (BatchClientGetSparkBatchJobResponse, error) {
	result := BatchClientGetSparkBatchJobResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BatchJob); err != nil {
		return BatchClientGetSparkBatchJobResponse{}, err
	}
	return result, nil
}

// GetSparkBatchJobs - List all spark batch jobs which are running under a particular spark pool.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-12-01
//   - options - BatchClientGetSparkBatchJobsOptions contains the optional parameters for the BatchClient.GetSparkBatchJobs method.
func (client *BatchClient) GetSparkBatchJobs(ctx context.Context, options *BatchClientGetSparkBatchJobsOptions) (BatchClientGetSparkBatchJobsResponse, error) {
	var err error
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "BatchClient.GetSparkBatchJobs")
	req, err := client.getSparkBatchJobsCreateRequest(ctx, options)
	if err != nil {
		return BatchClientGetSparkBatchJobsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BatchClientGetSparkBatchJobsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return BatchClientGetSparkBatchJobsResponse{}, err
	}
	resp, err := client.getSparkBatchJobsHandleResponse(httpResp)
	return resp, err
}

// getSparkBatchJobsCreateRequest creates the GetSparkBatchJobs request.
func (client *BatchClient) getSparkBatchJobsCreateRequest(ctx context.Context, options *BatchClientGetSparkBatchJobsOptions) (*policy.Request, error) {
	urlPath := "/livyApi/versions/{livyApiVersion}/sparkPools/{sparkPoolName}/batches"
	urlPath = strings.ReplaceAll(urlPath, "{livyApiVersion}", client.livyAPIVersion)
	urlPath = strings.ReplaceAll(urlPath, "{sparkPoolName}", client.sparkPoolName)
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
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getSparkBatchJobsHandleResponse handles the GetSparkBatchJobs response.
func (client *BatchClient) getSparkBatchJobsHandleResponse(resp *http.Response) (BatchClientGetSparkBatchJobsResponse, error) {
	result := BatchClientGetSparkBatchJobsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BatchJobCollection); err != nil {
		return BatchClientGetSparkBatchJobsResponse{}, err
	}
	return result, nil
}
