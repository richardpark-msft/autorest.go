//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package datetimegroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
	"time"
)

// QueryClient contains the methods for the Encode.Datetime group.
// Don't use this type directly, use a constructor function instead.
type QueryClient struct {
	internal *azcore.Client
}

func (client *QueryClient) Default(ctx context.Context, value time.Time, options *QueryClientDefaultOptions) (QueryClientDefaultResponse, error) {
	var err error
	req, err := client.defaultCreateRequest(ctx, value, options)
	if err != nil {
		return QueryClientDefaultResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return QueryClientDefaultResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return QueryClientDefaultResponse{}, err
	}
	return QueryClientDefaultResponse{}, nil
}

// defaultCreateRequest creates the Default request.
func (client *QueryClient) defaultCreateRequest(ctx context.Context, value time.Time, options *QueryClientDefaultOptions) (*policy.Request, error) {
	urlPath := "/encode/datetime/query/default"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("value", value.Format(time.RFC3339Nano))
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

func (client *QueryClient) RFC3339(ctx context.Context, value time.Time, options *QueryClientRFC3339Options) (QueryClientRFC3339Response, error) {
	var err error
	req, err := client.rfc3339CreateRequest(ctx, value, options)
	if err != nil {
		return QueryClientRFC3339Response{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return QueryClientRFC3339Response{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return QueryClientRFC3339Response{}, err
	}
	return QueryClientRFC3339Response{}, nil
}

// rfc3339CreateRequest creates the RFC3339 request.
func (client *QueryClient) rfc3339CreateRequest(ctx context.Context, value time.Time, options *QueryClientRFC3339Options) (*policy.Request, error) {
	urlPath := "/encode/datetime/query/rfc3339"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("value", value.Format(time.RFC3339Nano))
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

func (client *QueryClient) RFC7231(ctx context.Context, value time.Time, options *QueryClientRFC7231Options) (QueryClientRFC7231Response, error) {
	var err error
	req, err := client.rfc7231CreateRequest(ctx, value, options)
	if err != nil {
		return QueryClientRFC7231Response{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return QueryClientRFC7231Response{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return QueryClientRFC7231Response{}, err
	}
	return QueryClientRFC7231Response{}, nil
}

// rfc7231CreateRequest creates the RFC7231 request.
func (client *QueryClient) rfc7231CreateRequest(ctx context.Context, value time.Time, options *QueryClientRFC7231Options) (*policy.Request, error) {
	urlPath := "/encode/datetime/query/rfc7231"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("value", value.Format(time.RFC1123))
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

func (client *QueryClient) UnixTimestamp(ctx context.Context, value time.Time, options *QueryClientUnixTimestampOptions) (QueryClientUnixTimestampResponse, error) {
	var err error
	req, err := client.unixTimestampCreateRequest(ctx, value, options)
	if err != nil {
		return QueryClientUnixTimestampResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return QueryClientUnixTimestampResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return QueryClientUnixTimestampResponse{}, err
	}
	return QueryClientUnixTimestampResponse{}, nil
}

// unixTimestampCreateRequest creates the UnixTimestamp request.
func (client *QueryClient) unixTimestampCreateRequest(ctx context.Context, value time.Time, options *QueryClientUnixTimestampOptions) (*policy.Request, error) {
	urlPath := "/encode/datetime/query/unix-timestamp"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("value", timeUnix(value).String())
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

func (client *QueryClient) UnixTimestampArray(ctx context.Context, value []time.Time, options *QueryClientUnixTimestampArrayOptions) (QueryClientUnixTimestampArrayResponse, error) {
	var err error
	req, err := client.unixTimestampArrayCreateRequest(ctx, value, options)
	if err != nil {
		return QueryClientUnixTimestampArrayResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return QueryClientUnixTimestampArrayResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return QueryClientUnixTimestampArrayResponse{}, err
	}
	return QueryClientUnixTimestampArrayResponse{}, nil
}

// unixTimestampArrayCreateRequest creates the UnixTimestampArray request.
func (client *QueryClient) unixTimestampArrayCreateRequest(ctx context.Context, value []time.Time, options *QueryClientUnixTimestampArrayOptions) (*policy.Request, error) {
	urlPath := "/encode/datetime/query/unix-timestamp-array"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("value", strings.Join(func() []string {
		encodedValue := make([]string, len(value))
		for i := 0; i < len(value); i++ {
			encodedValue[i] = timeUnix(value[i]).String()
		}
		return encodedValue
	}(), ","))
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}
