//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package accessgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// PublicOperationClient contains the methods for the _Specs_.Azure.ClientGenerator.Core.Access group.
// Don't use this type directly, use a constructor function instead.
type PublicOperationClient struct {
	internal *azcore.Client
}

func (client *PublicOperationClient) NoDecoratorInPublic(ctx context.Context, name string, options *PublicOperationClientNoDecoratorInPublicOptions) (PublicOperationClientNoDecoratorInPublicResponse, error) {
	var err error
	req, err := client.noDecoratorInPublicCreateRequest(ctx, name, options)
	if err != nil {
		return PublicOperationClientNoDecoratorInPublicResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PublicOperationClientNoDecoratorInPublicResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PublicOperationClientNoDecoratorInPublicResponse{}, err
	}
	resp, err := client.noDecoratorInPublicHandleResponse(httpResp)
	return resp, err
}

// noDecoratorInPublicCreateRequest creates the NoDecoratorInPublic request.
func (client *PublicOperationClient) noDecoratorInPublicCreateRequest(ctx context.Context, name string, options *PublicOperationClientNoDecoratorInPublicOptions) (*policy.Request, error) {
	urlPath := "/azure/client-generator-core/access/publicOperation/noDecoratorInPublic"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("name", name)
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// noDecoratorInPublicHandleResponse handles the NoDecoratorInPublic response.
func (client *PublicOperationClient) noDecoratorInPublicHandleResponse(resp *http.Response) (PublicOperationClientNoDecoratorInPublicResponse, error) {
	result := PublicOperationClientNoDecoratorInPublicResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NoDecoratorModelInPublic); err != nil {
		return PublicOperationClientNoDecoratorInPublicResponse{}, err
	}
	return result, nil
}

func (client *PublicOperationClient) PublicDecoratorInPublic(ctx context.Context, name string, options *PublicOperationClientPublicDecoratorInPublicOptions) (PublicOperationClientPublicDecoratorInPublicResponse, error) {
	var err error
	req, err := client.publicDecoratorInPublicCreateRequest(ctx, name, options)
	if err != nil {
		return PublicOperationClientPublicDecoratorInPublicResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PublicOperationClientPublicDecoratorInPublicResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PublicOperationClientPublicDecoratorInPublicResponse{}, err
	}
	resp, err := client.publicDecoratorInPublicHandleResponse(httpResp)
	return resp, err
}

// publicDecoratorInPublicCreateRequest creates the PublicDecoratorInPublic request.
func (client *PublicOperationClient) publicDecoratorInPublicCreateRequest(ctx context.Context, name string, options *PublicOperationClientPublicDecoratorInPublicOptions) (*policy.Request, error) {
	urlPath := "/azure/client-generator-core/access/publicOperation/publicDecoratorInPublic"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("name", name)
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// publicDecoratorInPublicHandleResponse handles the PublicDecoratorInPublic response.
func (client *PublicOperationClient) publicDecoratorInPublicHandleResponse(resp *http.Response) (PublicOperationClientPublicDecoratorInPublicResponse, error) {
	result := PublicOperationClientPublicDecoratorInPublicResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PublicDecoratorModelInPublic); err != nil {
		return PublicOperationClientPublicDecoratorInPublicResponse{}, err
	}
	return result, nil
}
