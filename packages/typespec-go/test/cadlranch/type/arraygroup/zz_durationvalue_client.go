// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package arraygroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// DurationValueClient - Array of duration values
// Don't use this type directly, use [ArrayClient.NewDurationValueClient] instead.
type DurationValueClient struct {
	internal *azcore.Client
}

// - options - DurationValueClientGetOptions contains the optional parameters for the DurationValueClient.Get method.
func (client *DurationValueClient) Get(ctx context.Context, options *DurationValueClientGetOptions) (DurationValueClientGetResponse, error) {
	var err error
	const operationName = "DurationValueClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, options)
	if err != nil {
		return DurationValueClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DurationValueClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DurationValueClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *DurationValueClient) getCreateRequest(ctx context.Context, options *DurationValueClientGetOptions) (*policy.Request, error) {
	urlPath := "/type/array/duration"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DurationValueClient) getHandleResponse(resp *http.Response) (DurationValueClientGetResponse, error) {
	result := DurationValueClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return DurationValueClientGetResponse{}, err
	}
	return result, nil
}

// - options - DurationValueClientPutOptions contains the optional parameters for the DurationValueClient.Put method.
func (client *DurationValueClient) Put(ctx context.Context, body []string, options *DurationValueClientPutOptions) (DurationValueClientPutResponse, error) {
	var err error
	const operationName = "DurationValueClient.Put"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putCreateRequest(ctx, body, options)
	if err != nil {
		return DurationValueClientPutResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DurationValueClientPutResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return DurationValueClientPutResponse{}, err
	}
	return DurationValueClientPutResponse{}, nil
}

// putCreateRequest creates the Put request.
func (client *DurationValueClient) putCreateRequest(ctx context.Context, body []string, options *DurationValueClientPutOptions) (*policy.Request, error) {
	urlPath := "/type/array/duration"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}
