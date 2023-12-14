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

// HeaderClient contains the methods for the Encode.Datetime group.
// Don't use this type directly, use a constructor function instead.
type HeaderClient struct {
	internal *azcore.Client
}

func (client *HeaderClient) Default(ctx context.Context, value time.Time, options *HeaderClientDefaultOptions) (HeaderClientDefaultResponse, error) {
	var err error
	req, err := client.defaultCreateRequest(ctx, value, options)
	if err != nil {
		return HeaderClientDefaultResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HeaderClientDefaultResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return HeaderClientDefaultResponse{}, err
	}
	return HeaderClientDefaultResponse{}, nil
}

// defaultCreateRequest creates the Default request.
func (client *HeaderClient) defaultCreateRequest(ctx context.Context, value time.Time, options *HeaderClientDefaultOptions) (*policy.Request, error) {
	urlPath := "/encode/datetime/header/default"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["value"] = []string{value.Format(time.RFC1123)}
	return req, nil
}

func (client *HeaderClient) RFC3339(ctx context.Context, value time.Time, options *HeaderClientRFC3339Options) (HeaderClientRFC3339Response, error) {
	var err error
	req, err := client.rfc3339CreateRequest(ctx, value, options)
	if err != nil {
		return HeaderClientRFC3339Response{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HeaderClientRFC3339Response{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return HeaderClientRFC3339Response{}, err
	}
	return HeaderClientRFC3339Response{}, nil
}

// rfc3339CreateRequest creates the RFC3339 request.
func (client *HeaderClient) rfc3339CreateRequest(ctx context.Context, value time.Time, options *HeaderClientRFC3339Options) (*policy.Request, error) {
	urlPath := "/encode/datetime/header/rfc3339"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["value"] = []string{value.Format(time.RFC3339Nano)}
	return req, nil
}

func (client *HeaderClient) RFC7231(ctx context.Context, value time.Time, options *HeaderClientRFC7231Options) (HeaderClientRFC7231Response, error) {
	var err error
	req, err := client.rfc7231CreateRequest(ctx, value, options)
	if err != nil {
		return HeaderClientRFC7231Response{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HeaderClientRFC7231Response{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return HeaderClientRFC7231Response{}, err
	}
	return HeaderClientRFC7231Response{}, nil
}

// rfc7231CreateRequest creates the RFC7231 request.
func (client *HeaderClient) rfc7231CreateRequest(ctx context.Context, value time.Time, options *HeaderClientRFC7231Options) (*policy.Request, error) {
	urlPath := "/encode/datetime/header/rfc7231"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["value"] = []string{value.Format(time.RFC1123)}
	return req, nil
}

func (client *HeaderClient) UnixTimestamp(ctx context.Context, value time.Time, options *HeaderClientUnixTimestampOptions) (HeaderClientUnixTimestampResponse, error) {
	var err error
	req, err := client.unixTimestampCreateRequest(ctx, value, options)
	if err != nil {
		return HeaderClientUnixTimestampResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HeaderClientUnixTimestampResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return HeaderClientUnixTimestampResponse{}, err
	}
	return HeaderClientUnixTimestampResponse{}, nil
}

// unixTimestampCreateRequest creates the UnixTimestamp request.
func (client *HeaderClient) unixTimestampCreateRequest(ctx context.Context, value time.Time, options *HeaderClientUnixTimestampOptions) (*policy.Request, error) {
	urlPath := "/encode/datetime/header/unix-timestamp"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["value"] = []string{timeUnix(value).String()}
	return req, nil
}

func (client *HeaderClient) UnixTimestampArray(ctx context.Context, value []time.Time, options *HeaderClientUnixTimestampArrayOptions) (HeaderClientUnixTimestampArrayResponse, error) {
	var err error
	req, err := client.unixTimestampArrayCreateRequest(ctx, value, options)
	if err != nil {
		return HeaderClientUnixTimestampArrayResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HeaderClientUnixTimestampArrayResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return HeaderClientUnixTimestampArrayResponse{}, err
	}
	return HeaderClientUnixTimestampArrayResponse{}, nil
}

// unixTimestampArrayCreateRequest creates the UnixTimestampArray request.
func (client *HeaderClient) unixTimestampArrayCreateRequest(ctx context.Context, value []time.Time, options *HeaderClientUnixTimestampArrayOptions) (*policy.Request, error) {
	urlPath := "/encode/datetime/header/unix-timestamp-array"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["value"] = []string{strings.Join(func() []string {
		encodedValue := make([]string, len(value))
		for i := 0; i < len(value); i++ {
			encodedValue[i] = timeUnix(value[i]).String()
		}
		return encodedValue
	}(), ",")}
	return req, nil
}
