//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package urlgroup

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

// PathItemsClient contains the methods for the PathItems group.
// Don't use this type directly, use a constructor function instead.
type PathItemsClient struct {
	internal          *azcore.Client
	globalStringPath  string
	globalStringQuery *string
}

// GetAllWithValues - send globalStringPath='globalStringPath', pathItemStringPath='pathItemStringPath', localStringPath='localStringPath',
// globalStringQuery='globalStringQuery', pathItemStringQuery='pathItemStringQuery',
// localStringQuery='localStringQuery'
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - pathItemStringPath - A string value 'pathItemStringPath' that appears in the path
//   - localStringPath - should contain value 'localStringPath'
//   - options - PathItemsClientGetAllWithValuesOptions contains the optional parameters for the PathItemsClient.GetAllWithValues
//     method.
func (client *PathItemsClient) GetAllWithValues(ctx context.Context, pathItemStringPath string, localStringPath string, options *PathItemsClientGetAllWithValuesOptions) (PathItemsClientGetAllWithValuesResponse, error) {
	var err error
	const operationName = "PathItemsClient.GetAllWithValues"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getAllWithValuesCreateRequest(ctx, pathItemStringPath, localStringPath, options)
	if err != nil {
		return PathItemsClientGetAllWithValuesResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PathItemsClientGetAllWithValuesResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PathItemsClientGetAllWithValuesResponse{}, err
	}
	return PathItemsClientGetAllWithValuesResponse{}, nil
}

// getAllWithValuesCreateRequest creates the GetAllWithValues request.
func (client *PathItemsClient) getAllWithValuesCreateRequest(ctx context.Context, pathItemStringPath string, localStringPath string, options *PathItemsClientGetAllWithValuesOptions) (*policy.Request, error) {
	urlPath := "/pathitem/nullable/globalStringPath/{globalStringPath}/pathItemStringPath/{pathItemStringPath}/localStringPath/{localStringPath}/globalStringQuery/pathItemStringQuery/localStringQuery"
	if pathItemStringPath == "" {
		return nil, errors.New("parameter pathItemStringPath cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pathItemStringPath}", url.PathEscape(pathItemStringPath))
	if client.globalStringPath == "" {
		return nil, errors.New("parameter client.globalStringPath cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{globalStringPath}", url.PathEscape(client.globalStringPath))
	if localStringPath == "" {
		return nil, errors.New("parameter localStringPath cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{localStringPath}", url.PathEscape(localStringPath))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.PathItemStringQuery != nil {
		reqQP.Set("pathItemStringQuery", *options.PathItemStringQuery)
	}
	if client.globalStringQuery != nil {
		reqQP.Set("globalStringQuery", *client.globalStringQuery)
	}
	if options != nil && options.LocalStringQuery != nil {
		reqQP.Set("localStringQuery", *options.LocalStringQuery)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// GetGlobalAndLocalQueryNull - send globalStringPath=globalStringPath, pathItemStringPath='pathItemStringPath', localStringPath='localStringPath',
// globalStringQuery=null, pathItemStringQuery='pathItemStringQuery',
// localStringQuery=null
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - pathItemStringPath - A string value 'pathItemStringPath' that appears in the path
//   - localStringPath - should contain value 'localStringPath'
//   - options - PathItemsClientGetGlobalAndLocalQueryNullOptions contains the optional parameters for the PathItemsClient.GetGlobalAndLocalQueryNull
//     method.
func (client *PathItemsClient) GetGlobalAndLocalQueryNull(ctx context.Context, pathItemStringPath string, localStringPath string, options *PathItemsClientGetGlobalAndLocalQueryNullOptions) (PathItemsClientGetGlobalAndLocalQueryNullResponse, error) {
	var err error
	const operationName = "PathItemsClient.GetGlobalAndLocalQueryNull"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getGlobalAndLocalQueryNullCreateRequest(ctx, pathItemStringPath, localStringPath, options)
	if err != nil {
		return PathItemsClientGetGlobalAndLocalQueryNullResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PathItemsClientGetGlobalAndLocalQueryNullResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PathItemsClientGetGlobalAndLocalQueryNullResponse{}, err
	}
	return PathItemsClientGetGlobalAndLocalQueryNullResponse{}, nil
}

// getGlobalAndLocalQueryNullCreateRequest creates the GetGlobalAndLocalQueryNull request.
func (client *PathItemsClient) getGlobalAndLocalQueryNullCreateRequest(ctx context.Context, pathItemStringPath string, localStringPath string, options *PathItemsClientGetGlobalAndLocalQueryNullOptions) (*policy.Request, error) {
	urlPath := "/pathitem/nullable/globalStringPath/{globalStringPath}/pathItemStringPath/{pathItemStringPath}/localStringPath/{localStringPath}/null/pathItemStringQuery/null"
	if pathItemStringPath == "" {
		return nil, errors.New("parameter pathItemStringPath cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pathItemStringPath}", url.PathEscape(pathItemStringPath))
	if client.globalStringPath == "" {
		return nil, errors.New("parameter client.globalStringPath cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{globalStringPath}", url.PathEscape(client.globalStringPath))
	if localStringPath == "" {
		return nil, errors.New("parameter localStringPath cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{localStringPath}", url.PathEscape(localStringPath))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.PathItemStringQuery != nil {
		reqQP.Set("pathItemStringQuery", *options.PathItemStringQuery)
	}
	if client.globalStringQuery != nil {
		reqQP.Set("globalStringQuery", *client.globalStringQuery)
	}
	if options != nil && options.LocalStringQuery != nil {
		reqQP.Set("localStringQuery", *options.LocalStringQuery)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// GetGlobalQueryNull - send globalStringPath='globalStringPath', pathItemStringPath='pathItemStringPath', localStringPath='localStringPath',
// globalStringQuery=null, pathItemStringQuery='pathItemStringQuery',
// localStringQuery='localStringQuery'
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - pathItemStringPath - A string value 'pathItemStringPath' that appears in the path
//   - localStringPath - should contain value 'localStringPath'
//   - options - PathItemsClientGetGlobalQueryNullOptions contains the optional parameters for the PathItemsClient.GetGlobalQueryNull
//     method.
func (client *PathItemsClient) GetGlobalQueryNull(ctx context.Context, pathItemStringPath string, localStringPath string, options *PathItemsClientGetGlobalQueryNullOptions) (PathItemsClientGetGlobalQueryNullResponse, error) {
	var err error
	const operationName = "PathItemsClient.GetGlobalQueryNull"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getGlobalQueryNullCreateRequest(ctx, pathItemStringPath, localStringPath, options)
	if err != nil {
		return PathItemsClientGetGlobalQueryNullResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PathItemsClientGetGlobalQueryNullResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PathItemsClientGetGlobalQueryNullResponse{}, err
	}
	return PathItemsClientGetGlobalQueryNullResponse{}, nil
}

// getGlobalQueryNullCreateRequest creates the GetGlobalQueryNull request.
func (client *PathItemsClient) getGlobalQueryNullCreateRequest(ctx context.Context, pathItemStringPath string, localStringPath string, options *PathItemsClientGetGlobalQueryNullOptions) (*policy.Request, error) {
	urlPath := "/pathitem/nullable/globalStringPath/{globalStringPath}/pathItemStringPath/{pathItemStringPath}/localStringPath/{localStringPath}/null/pathItemStringQuery/localStringQuery"
	if pathItemStringPath == "" {
		return nil, errors.New("parameter pathItemStringPath cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pathItemStringPath}", url.PathEscape(pathItemStringPath))
	if client.globalStringPath == "" {
		return nil, errors.New("parameter client.globalStringPath cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{globalStringPath}", url.PathEscape(client.globalStringPath))
	if localStringPath == "" {
		return nil, errors.New("parameter localStringPath cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{localStringPath}", url.PathEscape(localStringPath))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.PathItemStringQuery != nil {
		reqQP.Set("pathItemStringQuery", *options.PathItemStringQuery)
	}
	if client.globalStringQuery != nil {
		reqQP.Set("globalStringQuery", *client.globalStringQuery)
	}
	if options != nil && options.LocalStringQuery != nil {
		reqQP.Set("localStringQuery", *options.LocalStringQuery)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// GetLocalPathItemQueryNull - send globalStringPath='globalStringPath', pathItemStringPath='pathItemStringPath', localStringPath='localStringPath',
// globalStringQuery='globalStringQuery', pathItemStringQuery=null,
// localStringQuery=null
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - pathItemStringPath - A string value 'pathItemStringPath' that appears in the path
//   - localStringPath - should contain value 'localStringPath'
//   - options - PathItemsClientGetLocalPathItemQueryNullOptions contains the optional parameters for the PathItemsClient.GetLocalPathItemQueryNull
//     method.
func (client *PathItemsClient) GetLocalPathItemQueryNull(ctx context.Context, pathItemStringPath string, localStringPath string, options *PathItemsClientGetLocalPathItemQueryNullOptions) (PathItemsClientGetLocalPathItemQueryNullResponse, error) {
	var err error
	const operationName = "PathItemsClient.GetLocalPathItemQueryNull"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getLocalPathItemQueryNullCreateRequest(ctx, pathItemStringPath, localStringPath, options)
	if err != nil {
		return PathItemsClientGetLocalPathItemQueryNullResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PathItemsClientGetLocalPathItemQueryNullResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PathItemsClientGetLocalPathItemQueryNullResponse{}, err
	}
	return PathItemsClientGetLocalPathItemQueryNullResponse{}, nil
}

// getLocalPathItemQueryNullCreateRequest creates the GetLocalPathItemQueryNull request.
func (client *PathItemsClient) getLocalPathItemQueryNullCreateRequest(ctx context.Context, pathItemStringPath string, localStringPath string, options *PathItemsClientGetLocalPathItemQueryNullOptions) (*policy.Request, error) {
	urlPath := "/pathitem/nullable/globalStringPath/{globalStringPath}/pathItemStringPath/{pathItemStringPath}/localStringPath/{localStringPath}/globalStringQuery/null/null"
	if pathItemStringPath == "" {
		return nil, errors.New("parameter pathItemStringPath cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pathItemStringPath}", url.PathEscape(pathItemStringPath))
	if client.globalStringPath == "" {
		return nil, errors.New("parameter client.globalStringPath cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{globalStringPath}", url.PathEscape(client.globalStringPath))
	if localStringPath == "" {
		return nil, errors.New("parameter localStringPath cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{localStringPath}", url.PathEscape(localStringPath))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.PathItemStringQuery != nil {
		reqQP.Set("pathItemStringQuery", *options.PathItemStringQuery)
	}
	if client.globalStringQuery != nil {
		reqQP.Set("globalStringQuery", *client.globalStringQuery)
	}
	if options != nil && options.LocalStringQuery != nil {
		reqQP.Set("localStringQuery", *options.LocalStringQuery)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}
