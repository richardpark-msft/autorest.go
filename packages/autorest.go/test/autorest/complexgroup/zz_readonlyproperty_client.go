//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
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
// Don't use this type directly, use a constructor function instead.
type ReadonlypropertyClient struct {
	internal *azcore.Client
}

// GetValid - Get complex types that have readonly properties
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2016-02-29
//   - options - ReadonlypropertyClientGetValidOptions contains the optional parameters for the ReadonlypropertyClient.GetValid
//     method.
func (client *ReadonlypropertyClient) GetValid(ctx context.Context, options *ReadonlypropertyClientGetValidOptions) (ReadonlypropertyClientGetValidResponse, error) {
	var err error
	const operationName = "ReadonlypropertyClient.GetValid"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getValidCreateRequest(ctx, options)
	if err != nil {
		return ReadonlypropertyClientGetValidResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ReadonlypropertyClientGetValidResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ReadonlypropertyClientGetValidResponse{}, err
	}
	resp, err := client.getValidHandleResponse(httpResp)
	return resp, err
}

// getValidCreateRequest creates the GetValid request.
func (client *ReadonlypropertyClient) getValidCreateRequest(ctx context.Context, options *ReadonlypropertyClientGetValidOptions) (*policy.Request, error) {
	urlPath := "/complex/readonlyproperty/valid"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
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
//
// Generated from API version 2016-02-29
//   - options - ReadonlypropertyClientPutValidOptions contains the optional parameters for the ReadonlypropertyClient.PutValid
//     method.
func (client *ReadonlypropertyClient) PutValid(ctx context.Context, complexBody ReadonlyObj, options *ReadonlypropertyClientPutValidOptions) (ReadonlypropertyClientPutValidResponse, error) {
	var err error
	const operationName = "ReadonlypropertyClient.PutValid"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putValidCreateRequest(ctx, complexBody, options)
	if err != nil {
		return ReadonlypropertyClientPutValidResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ReadonlypropertyClientPutValidResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ReadonlypropertyClientPutValidResponse{}, err
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
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, complexBody); err != nil {
		return nil, err
	}
	return req, nil
}
