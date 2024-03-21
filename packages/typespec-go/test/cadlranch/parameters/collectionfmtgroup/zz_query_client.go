// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package collectionfmtgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
)

// QueryClient contains the methods for the Parameters.CollectionFormat namespace.
// Don't use this type directly, use [CollectionFormatClient.NewQueryClient] instead.
type QueryClient struct {
	internal *azcore.Client
}

// - colors - Possible values for colors are [blue,red,green]
// - options - QueryClientCSVOptions contains the optional parameters for the QueryClient.CSV method.
func (client *QueryClient) CSV(ctx context.Context, colors []string, options *QueryClientCSVOptions) (QueryClientCSVResponse, error) {
	var err error
	const operationName = "QueryClient.CSV"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.csvCreateRequest(ctx, colors, options)
	if err != nil {
		return QueryClientCSVResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return QueryClientCSVResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return QueryClientCSVResponse{}, err
	}
	return QueryClientCSVResponse{}, nil
}

// csvCreateRequest creates the CSV request.
func (client *QueryClient) csvCreateRequest(ctx context.Context, colors []string, options *QueryClientCSVOptions) (*policy.Request, error) {
	urlPath := "/parameters/collection-format/query/csv"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("colors", strings.Join(colors, ","))
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// - colors - Possible values for colors are [blue,red,green]
// - options - QueryClientMultiOptions contains the optional parameters for the QueryClient.Multi method.
func (client *QueryClient) Multi(ctx context.Context, colors []string, options *QueryClientMultiOptions) (QueryClientMultiResponse, error) {
	var err error
	const operationName = "QueryClient.Multi"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.multiCreateRequest(ctx, colors, options)
	if err != nil {
		return QueryClientMultiResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return QueryClientMultiResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return QueryClientMultiResponse{}, err
	}
	return QueryClientMultiResponse{}, nil
}

// multiCreateRequest creates the Multi request.
func (client *QueryClient) multiCreateRequest(ctx context.Context, colors []string, options *QueryClientMultiOptions) (*policy.Request, error) {
	urlPath := "/parameters/collection-format/query/multi"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	for _, qv := range colors {
		reqQP.Add("colors", qv)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// - colors - Possible values for colors are [blue,red,green]
// - options - QueryClientPipesOptions contains the optional parameters for the QueryClient.Pipes method.
func (client *QueryClient) Pipes(ctx context.Context, colors []string, options *QueryClientPipesOptions) (QueryClientPipesResponse, error) {
	var err error
	const operationName = "QueryClient.Pipes"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.pipesCreateRequest(ctx, colors, options)
	if err != nil {
		return QueryClientPipesResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return QueryClientPipesResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return QueryClientPipesResponse{}, err
	}
	return QueryClientPipesResponse{}, nil
}

// pipesCreateRequest creates the Pipes request.
func (client *QueryClient) pipesCreateRequest(ctx context.Context, colors []string, options *QueryClientPipesOptions) (*policy.Request, error) {
	urlPath := "/parameters/collection-format/query/pipes"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("colors", strings.Join(colors, "|"))
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// - colors - Possible values for colors are [blue,red,green]
// - options - QueryClientSsvOptions contains the optional parameters for the QueryClient.Ssv method.
func (client *QueryClient) Ssv(ctx context.Context, colors []string, options *QueryClientSsvOptions) (QueryClientSsvResponse, error) {
	var err error
	const operationName = "QueryClient.Ssv"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.ssvCreateRequest(ctx, colors, options)
	if err != nil {
		return QueryClientSsvResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return QueryClientSsvResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return QueryClientSsvResponse{}, err
	}
	return QueryClientSsvResponse{}, nil
}

// ssvCreateRequest creates the Ssv request.
func (client *QueryClient) ssvCreateRequest(ctx context.Context, colors []string, options *QueryClientSsvOptions) (*policy.Request, error) {
	urlPath := "/parameters/collection-format/query/ssv"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("colors", strings.Join(colors, " "))
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// - colors - Possible values for colors are [blue,red,green]
// - options - QueryClientTsvOptions contains the optional parameters for the QueryClient.Tsv method.
func (client *QueryClient) Tsv(ctx context.Context, colors []string, options *QueryClientTsvOptions) (QueryClientTsvResponse, error) {
	var err error
	const operationName = "QueryClient.Tsv"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.tsvCreateRequest(ctx, colors, options)
	if err != nil {
		return QueryClientTsvResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return QueryClientTsvResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return QueryClientTsvResponse{}, err
	}
	return QueryClientTsvResponse{}, nil
}

// tsvCreateRequest creates the Tsv request.
func (client *QueryClient) tsvCreateRequest(ctx context.Context, colors []string, options *QueryClientTsvOptions) (*policy.Request, error) {
	urlPath := "/parameters/collection-format/query/tsv"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("colors", strings.Join(colors, "\t"))
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}
