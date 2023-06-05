//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azartifacts

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// IntegrationRuntimesClient contains the methods for the IntegrationRuntimes group.
// Don't use this type directly, use a constructor function instead.
type IntegrationRuntimesClient struct {
	internal *azcore.Client
	endpoint string
}

// Get - Get Integration Runtime
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-12-01
//   - integrationRuntimeName - The Integration Runtime name
//   - options - IntegrationRuntimesClientGetOptions contains the optional parameters for the IntegrationRuntimesClient.Get method.
func (client *IntegrationRuntimesClient) Get(ctx context.Context, integrationRuntimeName string, options *IntegrationRuntimesClientGetOptions) (IntegrationRuntimesClientGetResponse, error) {
	var err error
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "IntegrationRuntimesClient.Get")
	req, err := client.getCreateRequest(ctx, integrationRuntimeName, options)
	if err != nil {
		return IntegrationRuntimesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IntegrationRuntimesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return IntegrationRuntimesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *IntegrationRuntimesClient) getCreateRequest(ctx context.Context, integrationRuntimeName string, options *IntegrationRuntimesClientGetOptions) (*policy.Request, error) {
	urlPath := "/integrationRuntimes/{integrationRuntimeName}"
	if integrationRuntimeName == "" {
		return nil, errors.New("parameter integrationRuntimeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationRuntimeName}", url.PathEscape(integrationRuntimeName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *IntegrationRuntimesClient) getHandleResponse(resp *http.Response) (IntegrationRuntimesClientGetResponse, error) {
	result := IntegrationRuntimesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IntegrationRuntimeResource); err != nil {
		return IntegrationRuntimesClientGetResponse{}, err
	}
	return result, nil
}

// List - List Integration Runtimes
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-12-01
//   - options - IntegrationRuntimesClientListOptions contains the optional parameters for the IntegrationRuntimesClient.List
//     method.
func (client *IntegrationRuntimesClient) List(ctx context.Context, options *IntegrationRuntimesClientListOptions) (IntegrationRuntimesClientListResponse, error) {
	var err error
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "IntegrationRuntimesClient.List")
	req, err := client.listCreateRequest(ctx, options)
	if err != nil {
		return IntegrationRuntimesClientListResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IntegrationRuntimesClientListResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return IntegrationRuntimesClientListResponse{}, err
	}
	resp, err := client.listHandleResponse(httpResp)
	return resp, err
}

// listCreateRequest creates the List request.
func (client *IntegrationRuntimesClient) listCreateRequest(ctx context.Context, options *IntegrationRuntimesClientListOptions) (*policy.Request, error) {
	urlPath := "/integrationRuntimes"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *IntegrationRuntimesClient) listHandleResponse(resp *http.Response) (IntegrationRuntimesClientListResponse, error) {
	result := IntegrationRuntimesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IntegrationRuntimeListResponse); err != nil {
		return IntegrationRuntimesClientListResponse{}, err
	}
	return result, nil
}
