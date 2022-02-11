//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package nonstringenumgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// FloatClient contains the methods for the Float group.
// Don't use this type directly, use NewFloatClient() instead.
type FloatClient struct {
	pl runtime.Pipeline
}

// NewFloatClient creates a new instance of FloatClient with the specified values.
// options - pass nil to accept the default values.
func NewFloatClient(options *azcore.ClientOptions) *FloatClient {
	if options == nil {
		options = &azcore.ClientOptions{}
	}
	client := &FloatClient{
		pl: runtime.NewPipeline(moduleName, moduleVersion, runtime.PipelineOptions{}, options),
	}
	return client
}

// Get - Get a float enum
// If the operation fails it returns an *azcore.ResponseError type.
// options - FloatClientGetOptions contains the optional parameters for the FloatClient.Get method.
func (client *FloatClient) Get(ctx context.Context, options *FloatClientGetOptions) (FloatClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, options)
	if err != nil {
		return FloatClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FloatClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FloatClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *FloatClient) getCreateRequest(ctx context.Context, options *FloatClientGetOptions) (*policy.Request, error) {
	urlPath := "/nonStringEnums/float/get"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *FloatClient) getHandleResponse(resp *http.Response) (FloatClientGetResponse, error) {
	result := FloatClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return FloatClientGetResponse{}, err
	}
	return result, nil
}

// Put - Put a float enum
// If the operation fails it returns an *azcore.ResponseError type.
// options - FloatClientPutOptions contains the optional parameters for the FloatClient.Put method.
func (client *FloatClient) Put(ctx context.Context, options *FloatClientPutOptions) (FloatClientPutResponse, error) {
	req, err := client.putCreateRequest(ctx, options)
	if err != nil {
		return FloatClientPutResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FloatClientPutResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FloatClientPutResponse{}, runtime.NewResponseError(resp)
	}
	return client.putHandleResponse(resp)
}

// putCreateRequest creates the Put request.
func (client *FloatClient) putCreateRequest(ctx context.Context, options *FloatClientPutOptions) (*policy.Request, error) {
	urlPath := "/nonStringEnums/float/put"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	if options != nil && options.Input != nil {
		return req, runtime.MarshalAsJSON(req, *options.Input)
	}
	return req, nil
}

// putHandleResponse handles the Put response.
func (client *FloatClient) putHandleResponse(resp *http.Response) (FloatClientPutResponse, error) {
	result := FloatClientPutResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return FloatClientPutResponse{}, err
	}
	return result, nil
}
