// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package scalargroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// Decimal128VerifyClient - Decimal128 type verification
// Don't use this type directly, use [ScalarClient.NewDecimal128VerifyClient] instead.
type Decimal128VerifyClient struct {
	internal *azcore.Client
}

//   - options - Decimal128VerifyClientPrepareVerifyOptions contains the optional parameters for the Decimal128VerifyClient.PrepareVerify
//     method.
func (client *Decimal128VerifyClient) PrepareVerify(ctx context.Context, options *Decimal128VerifyClientPrepareVerifyOptions) (Decimal128VerifyClientPrepareVerifyResponse, error) {
	var err error
	const operationName = "Decimal128VerifyClient.PrepareVerify"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.prepareVerifyCreateRequest(ctx, options)
	if err != nil {
		return Decimal128VerifyClientPrepareVerifyResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return Decimal128VerifyClientPrepareVerifyResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return Decimal128VerifyClientPrepareVerifyResponse{}, err
	}
	resp, err := client.prepareVerifyHandleResponse(httpResp)
	return resp, err
}

// prepareVerifyCreateRequest creates the PrepareVerify request.
func (client *Decimal128VerifyClient) prepareVerifyCreateRequest(ctx context.Context, options *Decimal128VerifyClientPrepareVerifyOptions) (*policy.Request, error) {
	urlPath := "/type/scalar/decimal128/prepare_verify"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// prepareVerifyHandleResponse handles the PrepareVerify response.
func (client *Decimal128VerifyClient) prepareVerifyHandleResponse(resp *http.Response) (Decimal128VerifyClientPrepareVerifyResponse, error) {
	result := Decimal128VerifyClientPrepareVerifyResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return Decimal128VerifyClientPrepareVerifyResponse{}, err
	}
	return result, nil
}

// - options - Decimal128VerifyClientVerifyOptions contains the optional parameters for the Decimal128VerifyClient.Verify method.
func (client *Decimal128VerifyClient) Verify(ctx context.Context, body float64, options *Decimal128VerifyClientVerifyOptions) (Decimal128VerifyClientVerifyResponse, error) {
	var err error
	const operationName = "Decimal128VerifyClient.Verify"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.verifyCreateRequest(ctx, body, options)
	if err != nil {
		return Decimal128VerifyClientVerifyResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return Decimal128VerifyClientVerifyResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return Decimal128VerifyClientVerifyResponse{}, err
	}
	return Decimal128VerifyClientVerifyResponse{}, nil
}

// verifyCreateRequest creates the Verify request.
func (client *Decimal128VerifyClient) verifyCreateRequest(ctx context.Context, body float64, options *Decimal128VerifyClientVerifyOptions) (*policy.Request, error) {
	urlPath := "/type/scalar/decimal128/verify"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}
