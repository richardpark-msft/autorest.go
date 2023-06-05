//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package additionalpropsgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// PetsClient contains the methods for the Pets group.
// Don't use this type directly, use a constructor function instead.
type PetsClient struct {
	internal *azcore.Client
}

// CreateAPInProperties - Create a Pet which contains more properties than what is defined.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - PetsClientCreateAPInPropertiesOptions contains the optional parameters for the PetsClient.CreateAPInProperties
//     method.
func (client *PetsClient) CreateAPInProperties(ctx context.Context, createParameters PetAPInProperties, options *PetsClientCreateAPInPropertiesOptions) (PetsClientCreateAPInPropertiesResponse, error) {
	var err error
	const operationName = "PetsClient.CreateAPInProperties"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createAPInPropertiesCreateRequest(ctx, createParameters, options)
	if err != nil {
		return PetsClientCreateAPInPropertiesResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PetsClientCreateAPInPropertiesResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PetsClientCreateAPInPropertiesResponse{}, err
	}
	resp, err := client.createAPInPropertiesHandleResponse(httpResp)
	return resp, err
}

// createAPInPropertiesCreateRequest creates the CreateAPInProperties request.
func (client *PetsClient) createAPInPropertiesCreateRequest(ctx context.Context, createParameters PetAPInProperties, options *PetsClientCreateAPInPropertiesOptions) (*policy.Request, error) {
	urlPath := "/additionalProperties/in/properties"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, createParameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createAPInPropertiesHandleResponse handles the CreateAPInProperties response.
func (client *PetsClient) createAPInPropertiesHandleResponse(resp *http.Response) (PetsClientCreateAPInPropertiesResponse, error) {
	result := PetsClientCreateAPInPropertiesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PetAPInProperties); err != nil {
		return PetsClientCreateAPInPropertiesResponse{}, err
	}
	return result, nil
}

// CreateAPInPropertiesWithAPString - Create a Pet which contains more properties than what is defined.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - PetsClientCreateAPInPropertiesWithAPStringOptions contains the optional parameters for the PetsClient.CreateAPInPropertiesWithAPString
//     method.
func (client *PetsClient) CreateAPInPropertiesWithAPString(ctx context.Context, createParameters PetAPInPropertiesWithAPString, options *PetsClientCreateAPInPropertiesWithAPStringOptions) (PetsClientCreateAPInPropertiesWithAPStringResponse, error) {
	var err error
	const operationName = "PetsClient.CreateAPInPropertiesWithAPString"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createAPInPropertiesWithAPStringCreateRequest(ctx, createParameters, options)
	if err != nil {
		return PetsClientCreateAPInPropertiesWithAPStringResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PetsClientCreateAPInPropertiesWithAPStringResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PetsClientCreateAPInPropertiesWithAPStringResponse{}, err
	}
	resp, err := client.createAPInPropertiesWithAPStringHandleResponse(httpResp)
	return resp, err
}

// createAPInPropertiesWithAPStringCreateRequest creates the CreateAPInPropertiesWithAPString request.
func (client *PetsClient) createAPInPropertiesWithAPStringCreateRequest(ctx context.Context, createParameters PetAPInPropertiesWithAPString, options *PetsClientCreateAPInPropertiesWithAPStringOptions) (*policy.Request, error) {
	urlPath := "/additionalProperties/in/properties/with/additionalProperties/string"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, createParameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createAPInPropertiesWithAPStringHandleResponse handles the CreateAPInPropertiesWithAPString response.
func (client *PetsClient) createAPInPropertiesWithAPStringHandleResponse(resp *http.Response) (PetsClientCreateAPInPropertiesWithAPStringResponse, error) {
	result := PetsClientCreateAPInPropertiesWithAPStringResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PetAPInPropertiesWithAPString); err != nil {
		return PetsClientCreateAPInPropertiesWithAPStringResponse{}, err
	}
	return result, nil
}

// CreateAPObject - Create a Pet which contains more properties than what is defined.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - PetsClientCreateAPObjectOptions contains the optional parameters for the PetsClient.CreateAPObject method.
func (client *PetsClient) CreateAPObject(ctx context.Context, createParameters PetAPObject, options *PetsClientCreateAPObjectOptions) (PetsClientCreateAPObjectResponse, error) {
	var err error
	const operationName = "PetsClient.CreateAPObject"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createAPObjectCreateRequest(ctx, createParameters, options)
	if err != nil {
		return PetsClientCreateAPObjectResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PetsClientCreateAPObjectResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PetsClientCreateAPObjectResponse{}, err
	}
	resp, err := client.createAPObjectHandleResponse(httpResp)
	return resp, err
}

// createAPObjectCreateRequest creates the CreateAPObject request.
func (client *PetsClient) createAPObjectCreateRequest(ctx context.Context, createParameters PetAPObject, options *PetsClientCreateAPObjectOptions) (*policy.Request, error) {
	urlPath := "/additionalProperties/type/object"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, createParameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createAPObjectHandleResponse handles the CreateAPObject response.
func (client *PetsClient) createAPObjectHandleResponse(resp *http.Response) (PetsClientCreateAPObjectResponse, error) {
	result := PetsClientCreateAPObjectResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PetAPObject); err != nil {
		return PetsClientCreateAPObjectResponse{}, err
	}
	return result, nil
}

// CreateAPString - Create a Pet which contains more properties than what is defined.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - PetsClientCreateAPStringOptions contains the optional parameters for the PetsClient.CreateAPString method.
func (client *PetsClient) CreateAPString(ctx context.Context, createParameters PetAPString, options *PetsClientCreateAPStringOptions) (PetsClientCreateAPStringResponse, error) {
	var err error
	const operationName = "PetsClient.CreateAPString"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createAPStringCreateRequest(ctx, createParameters, options)
	if err != nil {
		return PetsClientCreateAPStringResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PetsClientCreateAPStringResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PetsClientCreateAPStringResponse{}, err
	}
	resp, err := client.createAPStringHandleResponse(httpResp)
	return resp, err
}

// createAPStringCreateRequest creates the CreateAPString request.
func (client *PetsClient) createAPStringCreateRequest(ctx context.Context, createParameters PetAPString, options *PetsClientCreateAPStringOptions) (*policy.Request, error) {
	urlPath := "/additionalProperties/type/string"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, createParameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createAPStringHandleResponse handles the CreateAPString response.
func (client *PetsClient) createAPStringHandleResponse(resp *http.Response) (PetsClientCreateAPStringResponse, error) {
	result := PetsClientCreateAPStringResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PetAPString); err != nil {
		return PetsClientCreateAPStringResponse{}, err
	}
	return result, nil
}

// CreateAPTrue - Create a Pet which contains more properties than what is defined.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - PetsClientCreateAPTrueOptions contains the optional parameters for the PetsClient.CreateAPTrue method.
func (client *PetsClient) CreateAPTrue(ctx context.Context, createParameters PetAPTrue, options *PetsClientCreateAPTrueOptions) (PetsClientCreateAPTrueResponse, error) {
	var err error
	const operationName = "PetsClient.CreateAPTrue"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createAPTrueCreateRequest(ctx, createParameters, options)
	if err != nil {
		return PetsClientCreateAPTrueResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PetsClientCreateAPTrueResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PetsClientCreateAPTrueResponse{}, err
	}
	resp, err := client.createAPTrueHandleResponse(httpResp)
	return resp, err
}

// createAPTrueCreateRequest creates the CreateAPTrue request.
func (client *PetsClient) createAPTrueCreateRequest(ctx context.Context, createParameters PetAPTrue, options *PetsClientCreateAPTrueOptions) (*policy.Request, error) {
	urlPath := "/additionalProperties/true"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, createParameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createAPTrueHandleResponse handles the CreateAPTrue response.
func (client *PetsClient) createAPTrueHandleResponse(resp *http.Response) (PetsClientCreateAPTrueResponse, error) {
	result := PetsClientCreateAPTrueResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PetAPTrue); err != nil {
		return PetsClientCreateAPTrueResponse{}, err
	}
	return result, nil
}

// CreateCatAPTrue - Create a CatAPTrue which contains more properties than what is defined.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - PetsClientCreateCatAPTrueOptions contains the optional parameters for the PetsClient.CreateCatAPTrue method.
func (client *PetsClient) CreateCatAPTrue(ctx context.Context, createParameters CatAPTrue, options *PetsClientCreateCatAPTrueOptions) (PetsClientCreateCatAPTrueResponse, error) {
	var err error
	const operationName = "PetsClient.CreateCatAPTrue"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCatAPTrueCreateRequest(ctx, createParameters, options)
	if err != nil {
		return PetsClientCreateCatAPTrueResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PetsClientCreateCatAPTrueResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PetsClientCreateCatAPTrueResponse{}, err
	}
	resp, err := client.createCatAPTrueHandleResponse(httpResp)
	return resp, err
}

// createCatAPTrueCreateRequest creates the CreateCatAPTrue request.
func (client *PetsClient) createCatAPTrueCreateRequest(ctx context.Context, createParameters CatAPTrue, options *PetsClientCreateCatAPTrueOptions) (*policy.Request, error) {
	urlPath := "/additionalProperties/true-subclass"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, createParameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createCatAPTrueHandleResponse handles the CreateCatAPTrue response.
func (client *PetsClient) createCatAPTrueHandleResponse(resp *http.Response) (PetsClientCreateCatAPTrueResponse, error) {
	result := PetsClientCreateCatAPTrueResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CatAPTrue); err != nil {
		return PetsClientCreateCatAPTrueResponse{}, err
	}
	return result, nil
}
