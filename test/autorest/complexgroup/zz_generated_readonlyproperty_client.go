//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package complexgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// ReadonlypropertyClient contains the methods for the Readonlyproperty group.
// Don't use this type directly, use NewReadonlypropertyClient() instead.
type ReadonlypropertyClient struct {
	pl runtime.Pipeline
}

// NewReadonlypropertyClient creates a new instance of ReadonlypropertyClient with the specified values.
// options - pass nil to accept the default values.
func NewReadonlypropertyClient(options *azcore.ClientOptions) *ReadonlypropertyClient {
	if options == nil {
		options = &azcore.ClientOptions{}
	}
	client := &ReadonlypropertyClient{
		pl: runtime.NewPipeline(moduleName, moduleVersion, runtime.PipelineOptions{}, options),
	}
	return client
}

// GetValid - Get complex types that have readonly properties
// If the operation fails it returns an *azcore.ResponseError type.
// options - ReadonlypropertyClientGetValidOptions contains the optional parameters for the ReadonlypropertyClient.GetValid
// method.
func (client *ReadonlypropertyClient) GetValid(ctx context.Context, options *ReadonlypropertyClientGetValidOptions) (ReadonlypropertyClientGetValidResponse, error) {
	req, err := client.getValidCreateRequest(ctx, options)
	if err != nil {
		return ReadonlypropertyClientGetValidResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ReadonlypropertyClientGetValidResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ReadonlypropertyClientGetValidResponse{}, runtime.NewResponseError(resp)
	}
	return client.getValidHandleResponse(resp)
}

// getValidCreateRequest creates the GetValid request.
func (client *ReadonlypropertyClient) getValidCreateRequest(ctx context.Context, options *ReadonlypropertyClientGetValidOptions) (*policy.Request, error) {
	urlPath := "/complex/readonlyproperty/valid"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getValidHandleResponse handles the GetValid response.
func (client *ReadonlypropertyClient) getValidHandleResponse(resp *http.Response) (ReadonlypropertyClientGetValidResponse, error) {
	result := ReadonlypropertyClientGetValidResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ReadonlyObj); err != nil {
		return ReadonlypropertyClientGetValidResponse{}, err
	}
	return result, nil
}

// PutValid - Put complex types that have readonly properties
// If the operation fails it returns an *azcore.ResponseError type.
// options - ReadonlypropertyClientPutValidOptions contains the optional parameters for the ReadonlypropertyClient.PutValid
// method.
func (client *ReadonlypropertyClient) PutValid(ctx context.Context, complexBody ReadonlyObj, options *ReadonlypropertyClientPutValidOptions) (ReadonlypropertyClientPutValidResponse, error) {
	req, err := client.putValidCreateRequest(ctx, complexBody, options)
	if err != nil {
		return ReadonlypropertyClientPutValidResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ReadonlypropertyClientPutValidResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ReadonlypropertyClientPutValidResponse{}, runtime.NewResponseError(resp)
	}
	return ReadonlypropertyClientPutValidResponse{}, nil
}

// putValidCreateRequest creates the PutValid request.
func (client *ReadonlypropertyClient) putValidCreateRequest(ctx context.Context, complexBody ReadonlyObj, options *ReadonlypropertyClientPutValidOptions) (*policy.Request, error) {
	urlPath := "/complex/readonlyproperty/valid"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, complexBody)
}
