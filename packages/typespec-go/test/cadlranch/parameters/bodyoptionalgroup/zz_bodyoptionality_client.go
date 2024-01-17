//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package bodyoptionalgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// BodyOptionalityClient contains the methods for the Parameters.BodyOptionality group.
// Don't use this type directly, use a constructor function instead.
type BodyOptionalityClient struct {
	internal *azcore.Client
}

func (client *BodyOptionalityClient) RequiredExplicit(ctx context.Context, body BodyModel, options *BodyOptionalityClientRequiredExplicitOptions) (BodyOptionalityClientRequiredExplicitResponse, error) {
	var err error
	req, err := client.requiredExplicitCreateRequest(ctx, body, options)
	if err != nil {
		return BodyOptionalityClientRequiredExplicitResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BodyOptionalityClientRequiredExplicitResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return BodyOptionalityClientRequiredExplicitResponse{}, err
	}
	return BodyOptionalityClientRequiredExplicitResponse{}, nil
}

// requiredExplicitCreateRequest creates the RequiredExplicit request.
func (client *BodyOptionalityClient) requiredExplicitCreateRequest(ctx context.Context, body BodyModel, options *BodyOptionalityClientRequiredExplicitOptions) (*policy.Request, error) {
	urlPath := "/parameters/body-optionality/required-explicit"
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

func (client *BodyOptionalityClient) RequiredImplicit(ctx context.Context, body BodyModel, options *BodyOptionalityClientRequiredImplicitOptions) (BodyOptionalityClientRequiredImplicitResponse, error) {
	var err error
	req, err := client.requiredImplicitCreateRequest(ctx, body, options)
	if err != nil {
		return BodyOptionalityClientRequiredImplicitResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BodyOptionalityClientRequiredImplicitResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return BodyOptionalityClientRequiredImplicitResponse{}, err
	}
	return BodyOptionalityClientRequiredImplicitResponse{}, nil
}

// requiredImplicitCreateRequest creates the RequiredImplicit request.
func (client *BodyOptionalityClient) requiredImplicitCreateRequest(ctx context.Context, body BodyModel, options *BodyOptionalityClientRequiredImplicitOptions) (*policy.Request, error) {
	urlPath := "/parameters/body-optionality/required-implicit"
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