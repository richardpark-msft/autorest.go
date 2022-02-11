//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package paramgroupinggroup

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// ParameterGroupingClient contains the methods for the ParameterGrouping group.
// Don't use this type directly, use NewParameterGroupingClient() instead.
type ParameterGroupingClient struct {
	pl runtime.Pipeline
}

// NewParameterGroupingClient creates a new instance of ParameterGroupingClient with the specified values.
// options - pass nil to accept the default values.
func NewParameterGroupingClient(options *azcore.ClientOptions) *ParameterGroupingClient {
	if options == nil {
		options = &azcore.ClientOptions{}
	}
	client := &ParameterGroupingClient{
		pl: runtime.NewPipeline(moduleName, moduleVersion, runtime.PipelineOptions{}, options),
	}
	return client
}

// PostMultiParamGroups - Post parameters from multiple different parameter groups
// If the operation fails it returns an *azcore.ResponseError type.
// FirstParameterGroup - FirstParameterGroup contains a group of parameters for the ParameterGroupingClient.PostMultiParamGroups
// method.
// ParameterGroupingClientPostMultiParamGroupsSecondParamGroup - ParameterGroupingClientPostMultiParamGroupsSecondParamGroup
// contains a group of parameters for the ParameterGroupingClient.PostMultiParamGroups method.
func (client *ParameterGroupingClient) PostMultiParamGroups(ctx context.Context, firstParameterGroup *FirstParameterGroup, parameterGroupingClientPostMultiParamGroupsSecondParamGroup *ParameterGroupingClientPostMultiParamGroupsSecondParamGroup) (ParameterGroupingClientPostMultiParamGroupsResponse, error) {
	req, err := client.postMultiParamGroupsCreateRequest(ctx, firstParameterGroup, parameterGroupingClientPostMultiParamGroupsSecondParamGroup)
	if err != nil {
		return ParameterGroupingClientPostMultiParamGroupsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ParameterGroupingClientPostMultiParamGroupsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ParameterGroupingClientPostMultiParamGroupsResponse{}, runtime.NewResponseError(resp)
	}
	return ParameterGroupingClientPostMultiParamGroupsResponse{}, nil
}

// postMultiParamGroupsCreateRequest creates the PostMultiParamGroups request.
func (client *ParameterGroupingClient) postMultiParamGroupsCreateRequest(ctx context.Context, firstParameterGroup *FirstParameterGroup, parameterGroupingClientPostMultiParamGroupsSecondParamGroup *ParameterGroupingClientPostMultiParamGroupsSecondParamGroup) (*policy.Request, error) {
	urlPath := "/parameterGrouping/postMultipleParameterGroups"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if firstParameterGroup != nil && firstParameterGroup.QueryOne != nil {
		reqQP.Set("query-one", strconv.FormatInt(int64(*firstParameterGroup.QueryOne), 10))
	}
	if parameterGroupingClientPostMultiParamGroupsSecondParamGroup != nil && parameterGroupingClientPostMultiParamGroupsSecondParamGroup.QueryTwo != nil {
		reqQP.Set("query-two", strconv.FormatInt(int64(*parameterGroupingClientPostMultiParamGroupsSecondParamGroup.QueryTwo), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	if firstParameterGroup != nil && firstParameterGroup.HeaderOne != nil {
		req.Raw().Header.Set("header-one", *firstParameterGroup.HeaderOne)
	}
	if parameterGroupingClientPostMultiParamGroupsSecondParamGroup != nil && parameterGroupingClientPostMultiParamGroupsSecondParamGroup.HeaderTwo != nil {
		req.Raw().Header.Set("header-two", *parameterGroupingClientPostMultiParamGroupsSecondParamGroup.HeaderTwo)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// PostOptional - Post a bunch of optional parameters grouped
// If the operation fails it returns an *azcore.ResponseError type.
// options - ParameterGroupingClientPostOptionalParameters contains a group of parameters for the ParameterGroupingClient.PostOptional
// method.
func (client *ParameterGroupingClient) PostOptional(ctx context.Context, options *ParameterGroupingClientPostOptionalParameters) (ParameterGroupingClientPostOptionalResponse, error) {
	req, err := client.postOptionalCreateRequest(ctx, options)
	if err != nil {
		return ParameterGroupingClientPostOptionalResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ParameterGroupingClientPostOptionalResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ParameterGroupingClientPostOptionalResponse{}, runtime.NewResponseError(resp)
	}
	return ParameterGroupingClientPostOptionalResponse{}, nil
}

// postOptionalCreateRequest creates the PostOptional request.
func (client *ParameterGroupingClient) postOptionalCreateRequest(ctx context.Context, options *ParameterGroupingClientPostOptionalParameters) (*policy.Request, error) {
	urlPath := "/parameterGrouping/postOptional"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Query != nil {
		reqQP.Set("query", strconv.FormatInt(int64(*options.Query), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.CustomHeader != nil {
		req.Raw().Header.Set("customHeader", *options.CustomHeader)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// PostRequired - Post a bunch of required parameters grouped
// If the operation fails it returns an *azcore.ResponseError type.
// ParameterGroupingClientPostRequiredParameters - ParameterGroupingClientPostRequiredParameters contains a group of parameters
// for the ParameterGroupingClient.PostRequired method.
func (client *ParameterGroupingClient) PostRequired(ctx context.Context, parameterGroupingClientPostRequiredParameters ParameterGroupingClientPostRequiredParameters) (ParameterGroupingClientPostRequiredResponse, error) {
	req, err := client.postRequiredCreateRequest(ctx, parameterGroupingClientPostRequiredParameters)
	if err != nil {
		return ParameterGroupingClientPostRequiredResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ParameterGroupingClientPostRequiredResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ParameterGroupingClientPostRequiredResponse{}, runtime.NewResponseError(resp)
	}
	return ParameterGroupingClientPostRequiredResponse{}, nil
}

// postRequiredCreateRequest creates the PostRequired request.
func (client *ParameterGroupingClient) postRequiredCreateRequest(ctx context.Context, parameterGroupingClientPostRequiredParameters ParameterGroupingClientPostRequiredParameters) (*policy.Request, error) {
	urlPath := "/parameterGrouping/postRequired/{path}"
	if parameterGroupingClientPostRequiredParameters.Path == "" {
		return nil, errors.New("parameter parameterGroupingClientPostRequiredParameters.Path cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{path}", url.PathEscape(parameterGroupingClientPostRequiredParameters.Path))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if parameterGroupingClientPostRequiredParameters.Query != nil {
		reqQP.Set("query", strconv.FormatInt(int64(*parameterGroupingClientPostRequiredParameters.Query), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	if parameterGroupingClientPostRequiredParameters.CustomHeader != nil {
		req.Raw().Header.Set("customHeader", *parameterGroupingClientPostRequiredParameters.CustomHeader)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameterGroupingClientPostRequiredParameters.Body)
}

// PostReservedWords - Post a grouped parameters with reserved words
// If the operation fails it returns an *azcore.ResponseError type.
// options - ParameterGroupingClientPostReservedWordsParameters contains a group of parameters for the ParameterGroupingClient.PostReservedWords
// method.
func (client *ParameterGroupingClient) PostReservedWords(ctx context.Context, options *ParameterGroupingClientPostReservedWordsParameters) (ParameterGroupingClientPostReservedWordsResponse, error) {
	req, err := client.postReservedWordsCreateRequest(ctx, options)
	if err != nil {
		return ParameterGroupingClientPostReservedWordsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ParameterGroupingClientPostReservedWordsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ParameterGroupingClientPostReservedWordsResponse{}, runtime.NewResponseError(resp)
	}
	return ParameterGroupingClientPostReservedWordsResponse{}, nil
}

// postReservedWordsCreateRequest creates the PostReservedWords request.
func (client *ParameterGroupingClient) postReservedWordsCreateRequest(ctx context.Context, options *ParameterGroupingClientPostReservedWordsParameters) (*policy.Request, error) {
	urlPath := "/parameterGrouping/postReservedWords"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.From != nil {
		reqQP.Set("from", *options.From)
	}
	if options != nil && options.Accept != nil {
		reqQP.Set("accept", *options.Accept)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// PostSharedParameterGroupObject - Post parameters with a shared parameter group object
// If the operation fails it returns an *azcore.ResponseError type.
// options - FirstParameterGroup contains a group of parameters for the ParameterGroupingClient.PostMultiParamGroups method.
func (client *ParameterGroupingClient) PostSharedParameterGroupObject(ctx context.Context, options *FirstParameterGroup) (ParameterGroupingClientPostSharedParameterGroupObjectResponse, error) {
	req, err := client.postSharedParameterGroupObjectCreateRequest(ctx, options)
	if err != nil {
		return ParameterGroupingClientPostSharedParameterGroupObjectResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ParameterGroupingClientPostSharedParameterGroupObjectResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ParameterGroupingClientPostSharedParameterGroupObjectResponse{}, runtime.NewResponseError(resp)
	}
	return ParameterGroupingClientPostSharedParameterGroupObjectResponse{}, nil
}

// postSharedParameterGroupObjectCreateRequest creates the PostSharedParameterGroupObject request.
func (client *ParameterGroupingClient) postSharedParameterGroupObjectCreateRequest(ctx context.Context, options *FirstParameterGroup) (*policy.Request, error) {
	urlPath := "/parameterGrouping/sharedParameterGroupObject"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.QueryOne != nil {
		reqQP.Set("query-one", strconv.FormatInt(int64(*options.QueryOne), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.HeaderOne != nil {
		req.Raw().Header.Set("header-one", *options.HeaderOne)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}
