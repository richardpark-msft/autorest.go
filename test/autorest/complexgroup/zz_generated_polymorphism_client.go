//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package complexgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// PolymorphismClient contains the methods for the Polymorphism group.
// Don't use this type directly, use NewPolymorphismClient() instead.
type PolymorphismClient struct {
	pl runtime.Pipeline
}

// NewPolymorphismClient creates a new instance of PolymorphismClient with the specified values.
// pl - the pipeline used for sending requests and handling responses.
func NewPolymorphismClient(pl runtime.Pipeline) *PolymorphismClient {
	client := &PolymorphismClient{
		pl: pl,
	}
	return client
}

// GetComplicated - Get complex types that are polymorphic, but not at the root of the hierarchy; also have additional properties
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2016-02-29
// options - PolymorphismClientGetComplicatedOptions contains the optional parameters for the PolymorphismClient.GetComplicated
// method.
func (client *PolymorphismClient) GetComplicated(ctx context.Context, options *PolymorphismClientGetComplicatedOptions) (PolymorphismClientGetComplicatedResponse, error) {
	req, err := client.getComplicatedCreateRequest(ctx, options)
	if err != nil {
		return PolymorphismClientGetComplicatedResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PolymorphismClientGetComplicatedResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PolymorphismClientGetComplicatedResponse{}, runtime.NewResponseError(resp)
	}
	return client.getComplicatedHandleResponse(resp)
}

// getComplicatedCreateRequest creates the GetComplicated request.
func (client *PolymorphismClient) getComplicatedCreateRequest(ctx context.Context, options *PolymorphismClientGetComplicatedOptions) (*policy.Request, error) {
	urlPath := "/complex/polymorphism/complicated"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getComplicatedHandleResponse handles the GetComplicated response.
func (client *PolymorphismClient) getComplicatedHandleResponse(resp *http.Response) (PolymorphismClientGetComplicatedResponse, error) {
	result := PolymorphismClientGetComplicatedResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return PolymorphismClientGetComplicatedResponse{}, err
	}
	return result, nil
}

// GetComposedWithDiscriminator - Get complex object composing a polymorphic scalar property and array property with polymorphic
// element type, with discriminator specified. Deserialization must NOT fail and use the discriminator type
// specified on the wire.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2016-02-29
// options - PolymorphismClientGetComposedWithDiscriminatorOptions contains the optional parameters for the PolymorphismClient.GetComposedWithDiscriminator
// method.
func (client *PolymorphismClient) GetComposedWithDiscriminator(ctx context.Context, options *PolymorphismClientGetComposedWithDiscriminatorOptions) (PolymorphismClientGetComposedWithDiscriminatorResponse, error) {
	req, err := client.getComposedWithDiscriminatorCreateRequest(ctx, options)
	if err != nil {
		return PolymorphismClientGetComposedWithDiscriminatorResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PolymorphismClientGetComposedWithDiscriminatorResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PolymorphismClientGetComposedWithDiscriminatorResponse{}, runtime.NewResponseError(resp)
	}
	return client.getComposedWithDiscriminatorHandleResponse(resp)
}

// getComposedWithDiscriminatorCreateRequest creates the GetComposedWithDiscriminator request.
func (client *PolymorphismClient) getComposedWithDiscriminatorCreateRequest(ctx context.Context, options *PolymorphismClientGetComposedWithDiscriminatorOptions) (*policy.Request, error) {
	urlPath := "/complex/polymorphism/composedWithDiscriminator"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getComposedWithDiscriminatorHandleResponse handles the GetComposedWithDiscriminator response.
func (client *PolymorphismClient) getComposedWithDiscriminatorHandleResponse(resp *http.Response) (PolymorphismClientGetComposedWithDiscriminatorResponse, error) {
	result := PolymorphismClientGetComposedWithDiscriminatorResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DotFishMarket); err != nil {
		return PolymorphismClientGetComposedWithDiscriminatorResponse{}, err
	}
	return result, nil
}

// GetComposedWithoutDiscriminator - Get complex object composing a polymorphic scalar property and array property with polymorphic
// element type, without discriminator specified on wire. Deserialization must NOT fail and use the explicit
// type of the property.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2016-02-29
// options - PolymorphismClientGetComposedWithoutDiscriminatorOptions contains the optional parameters for the PolymorphismClient.GetComposedWithoutDiscriminator
// method.
func (client *PolymorphismClient) GetComposedWithoutDiscriminator(ctx context.Context, options *PolymorphismClientGetComposedWithoutDiscriminatorOptions) (PolymorphismClientGetComposedWithoutDiscriminatorResponse, error) {
	req, err := client.getComposedWithoutDiscriminatorCreateRequest(ctx, options)
	if err != nil {
		return PolymorphismClientGetComposedWithoutDiscriminatorResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PolymorphismClientGetComposedWithoutDiscriminatorResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PolymorphismClientGetComposedWithoutDiscriminatorResponse{}, runtime.NewResponseError(resp)
	}
	return client.getComposedWithoutDiscriminatorHandleResponse(resp)
}

// getComposedWithoutDiscriminatorCreateRequest creates the GetComposedWithoutDiscriminator request.
func (client *PolymorphismClient) getComposedWithoutDiscriminatorCreateRequest(ctx context.Context, options *PolymorphismClientGetComposedWithoutDiscriminatorOptions) (*policy.Request, error) {
	urlPath := "/complex/polymorphism/composedWithoutDiscriminator"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getComposedWithoutDiscriminatorHandleResponse handles the GetComposedWithoutDiscriminator response.
func (client *PolymorphismClient) getComposedWithoutDiscriminatorHandleResponse(resp *http.Response) (PolymorphismClientGetComposedWithoutDiscriminatorResponse, error) {
	result := PolymorphismClientGetComposedWithoutDiscriminatorResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DotFishMarket); err != nil {
		return PolymorphismClientGetComposedWithoutDiscriminatorResponse{}, err
	}
	return result, nil
}

// GetDotSyntax - Get complex types that are polymorphic, JSON key contains a dot
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2016-02-29
// options - PolymorphismClientGetDotSyntaxOptions contains the optional parameters for the PolymorphismClient.GetDotSyntax
// method.
func (client *PolymorphismClient) GetDotSyntax(ctx context.Context, options *PolymorphismClientGetDotSyntaxOptions) (PolymorphismClientGetDotSyntaxResponse, error) {
	req, err := client.getDotSyntaxCreateRequest(ctx, options)
	if err != nil {
		return PolymorphismClientGetDotSyntaxResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PolymorphismClientGetDotSyntaxResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PolymorphismClientGetDotSyntaxResponse{}, runtime.NewResponseError(resp)
	}
	return client.getDotSyntaxHandleResponse(resp)
}

// getDotSyntaxCreateRequest creates the GetDotSyntax request.
func (client *PolymorphismClient) getDotSyntaxCreateRequest(ctx context.Context, options *PolymorphismClientGetDotSyntaxOptions) (*policy.Request, error) {
	urlPath := "/complex/polymorphism/dotsyntax"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getDotSyntaxHandleResponse handles the GetDotSyntax response.
func (client *PolymorphismClient) getDotSyntaxHandleResponse(resp *http.Response) (PolymorphismClientGetDotSyntaxResponse, error) {
	result := PolymorphismClientGetDotSyntaxResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return PolymorphismClientGetDotSyntaxResponse{}, err
	}
	return result, nil
}

// GetValid - Get complex types that are polymorphic
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2016-02-29
// options - PolymorphismClientGetValidOptions contains the optional parameters for the PolymorphismClient.GetValid method.
func (client *PolymorphismClient) GetValid(ctx context.Context, options *PolymorphismClientGetValidOptions) (PolymorphismClientGetValidResponse, error) {
	req, err := client.getValidCreateRequest(ctx, options)
	if err != nil {
		return PolymorphismClientGetValidResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PolymorphismClientGetValidResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PolymorphismClientGetValidResponse{}, runtime.NewResponseError(resp)
	}
	return client.getValidHandleResponse(resp)
}

// getValidCreateRequest creates the GetValid request.
func (client *PolymorphismClient) getValidCreateRequest(ctx context.Context, options *PolymorphismClientGetValidOptions) (*policy.Request, error) {
	urlPath := "/complex/polymorphism/valid"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getValidHandleResponse handles the GetValid response.
func (client *PolymorphismClient) getValidHandleResponse(resp *http.Response) (PolymorphismClientGetValidResponse, error) {
	result := PolymorphismClientGetValidResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return PolymorphismClientGetValidResponse{}, err
	}
	return result, nil
}

// PutComplicated - Put complex types that are polymorphic, but not at the root of the hierarchy; also have additional properties
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2016-02-29
// options - PolymorphismClientPutComplicatedOptions contains the optional parameters for the PolymorphismClient.PutComplicated
// method.
func (client *PolymorphismClient) PutComplicated(ctx context.Context, complexBody SalmonClassification, options *PolymorphismClientPutComplicatedOptions) (PolymorphismClientPutComplicatedResponse, error) {
	req, err := client.putComplicatedCreateRequest(ctx, complexBody, options)
	if err != nil {
		return PolymorphismClientPutComplicatedResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PolymorphismClientPutComplicatedResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PolymorphismClientPutComplicatedResponse{}, runtime.NewResponseError(resp)
	}
	return PolymorphismClientPutComplicatedResponse{}, nil
}

// putComplicatedCreateRequest creates the PutComplicated request.
func (client *PolymorphismClient) putComplicatedCreateRequest(ctx context.Context, complexBody SalmonClassification, options *PolymorphismClientPutComplicatedOptions) (*policy.Request, error) {
	urlPath := "/complex/polymorphism/complicated"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, complexBody)
}

// PutMissingDiscriminator - Put complex types that are polymorphic, omitting the discriminator
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2016-02-29
// options - PolymorphismClientPutMissingDiscriminatorOptions contains the optional parameters for the PolymorphismClient.PutMissingDiscriminator
// method.
func (client *PolymorphismClient) PutMissingDiscriminator(ctx context.Context, complexBody SalmonClassification, options *PolymorphismClientPutMissingDiscriminatorOptions) (PolymorphismClientPutMissingDiscriminatorResponse, error) {
	req, err := client.putMissingDiscriminatorCreateRequest(ctx, complexBody, options)
	if err != nil {
		return PolymorphismClientPutMissingDiscriminatorResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PolymorphismClientPutMissingDiscriminatorResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PolymorphismClientPutMissingDiscriminatorResponse{}, runtime.NewResponseError(resp)
	}
	return client.putMissingDiscriminatorHandleResponse(resp)
}

// putMissingDiscriminatorCreateRequest creates the PutMissingDiscriminator request.
func (client *PolymorphismClient) putMissingDiscriminatorCreateRequest(ctx context.Context, complexBody SalmonClassification, options *PolymorphismClientPutMissingDiscriminatorOptions) (*policy.Request, error) {
	urlPath := "/complex/polymorphism/missingdiscriminator"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, complexBody)
}

// putMissingDiscriminatorHandleResponse handles the PutMissingDiscriminator response.
func (client *PolymorphismClient) putMissingDiscriminatorHandleResponse(resp *http.Response) (PolymorphismClientPutMissingDiscriminatorResponse, error) {
	result := PolymorphismClientPutMissingDiscriminatorResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return PolymorphismClientPutMissingDiscriminatorResponse{}, err
	}
	return result, nil
}

// PutValid - Put complex types that are polymorphic
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2016-02-29
// complexBody - Please put a salmon that looks like this: { 'fishtype':'Salmon', 'location':'alaska', 'iswild':true, 'species':'king',
// 'length':1.0, 'siblings':[ { 'fishtype':'Shark', 'age':6, 'birthday':
// '2012-01-05T01:00:00Z', 'length':20.0, 'species':'predator', }, { 'fishtype':'Sawshark', 'age':105, 'birthday': '1900-01-05T01:00:00Z',
// 'length':10.0, 'picture': new Buffer([255, 255, 255, 255,
// 254]).toString('base64'), 'species':'dangerous', }, { 'fishtype': 'goblin', 'age': 1, 'birthday': '2015-08-08T00:00:00Z',
// 'length': 30.0, 'species': 'scary', 'jawsize': 5 } ] };
// options - PolymorphismClientPutValidOptions contains the optional parameters for the PolymorphismClient.PutValid method.
func (client *PolymorphismClient) PutValid(ctx context.Context, complexBody FishClassification, options *PolymorphismClientPutValidOptions) (PolymorphismClientPutValidResponse, error) {
	req, err := client.putValidCreateRequest(ctx, complexBody, options)
	if err != nil {
		return PolymorphismClientPutValidResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PolymorphismClientPutValidResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PolymorphismClientPutValidResponse{}, runtime.NewResponseError(resp)
	}
	return PolymorphismClientPutValidResponse{}, nil
}

// putValidCreateRequest creates the PutValid request.
func (client *PolymorphismClient) putValidCreateRequest(ctx context.Context, complexBody FishClassification, options *PolymorphismClientPutValidOptions) (*policy.Request, error) {
	urlPath := "/complex/polymorphism/valid"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, complexBody)
}

// PutValidMissingRequired - Put complex types that are polymorphic, attempting to omit required 'birthday' field - the request
// should not be allowed from the client
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2016-02-29
// complexBody - Please attempt put a sawshark that looks like this, the client should not allow this data to be sent: { "fishtype":
// "sawshark", "species": "snaggle toothed", "length": 18.5, "age": 2, "birthday":
// "2013-06-01T01:00:00Z", "location": "alaska", "picture": base64(FF FF FF FF FE), "siblings": [ { "fishtype": "shark", "species":
// "predator", "birthday": "2012-01-05T01:00:00Z", "length": 20, "age": 6
// }, { "fishtype": "sawshark", "species": "dangerous", "picture": base64(FF FF FF FF FE), "length": 10, "age": 105 } ] }
// options - PolymorphismClientPutValidMissingRequiredOptions contains the optional parameters for the PolymorphismClient.PutValidMissingRequired
// method.
func (client *PolymorphismClient) PutValidMissingRequired(ctx context.Context, complexBody FishClassification, options *PolymorphismClientPutValidMissingRequiredOptions) (PolymorphismClientPutValidMissingRequiredResponse, error) {
	req, err := client.putValidMissingRequiredCreateRequest(ctx, complexBody, options)
	if err != nil {
		return PolymorphismClientPutValidMissingRequiredResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PolymorphismClientPutValidMissingRequiredResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PolymorphismClientPutValidMissingRequiredResponse{}, runtime.NewResponseError(resp)
	}
	return PolymorphismClientPutValidMissingRequiredResponse{}, nil
}

// putValidMissingRequiredCreateRequest creates the PutValidMissingRequired request.
func (client *PolymorphismClient) putValidMissingRequiredCreateRequest(ctx context.Context, complexBody FishClassification, options *PolymorphismClientPutValidMissingRequiredOptions) (*policy.Request, error) {
	urlPath := "/complex/polymorphism/missingrequired/invalid"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, complexBody)
}
