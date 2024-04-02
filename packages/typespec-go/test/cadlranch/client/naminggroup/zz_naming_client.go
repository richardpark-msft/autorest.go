// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package naminggroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// NamingClient - Describe changing names of types in a client with `@clientName`
// Don't use this type directly, use a constructor function instead.
type NamingClient struct {
	internal *azcore.Client
}

// NewNamingModelClient creates a new instance of [NamingModelClient].
func (client *NamingClient) NewNamingModelClient() *NamingModelClient {
	return &NamingModelClient{
		internal: client.internal,
	}
}

// NewNamingUnionEnumClient creates a new instance of [NamingUnionEnumClient].
func (client *NamingClient) NewNamingUnionEnumClient() *NamingUnionEnumClient {
	return &NamingUnionEnumClient{
		internal: client.internal,
	}
}

// - options - NamingClientClientOptions contains the optional parameters for the NamingClient.Client method.
func (client *NamingClient) Client(ctx context.Context, clientNameModel ClientNameModel, options *NamingClientClientOptions) (NamingClientClientResponse, error) {
	var err error
	const operationName = "NamingClient.Client"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.clientCreateRequest(ctx, clientNameModel, options)
	if err != nil {
		return NamingClientClientResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NamingClientClientResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return NamingClientClientResponse{}, err
	}
	return NamingClientClientResponse{}, nil
}

// clientCreateRequest creates the Client request.
func (client *NamingClient) clientCreateRequest(ctx context.Context, clientNameModel ClientNameModel, options *NamingClientClientOptions) (*policy.Request, error) {
	urlPath := "/client/naming/property/client"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, clientNameModel); err != nil {
		return nil, err
	}
	return req, nil
}

// - options - NamingClientClientNameOptions contains the optional parameters for the NamingClient.ClientName method.
func (client *NamingClient) ClientName(ctx context.Context, options *NamingClientClientNameOptions) (NamingClientClientNameResponse, error) {
	var err error
	const operationName = "NamingClient.ClientName"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.clientNameCreateRequest(ctx, options)
	if err != nil {
		return NamingClientClientNameResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NamingClientClientNameResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return NamingClientClientNameResponse{}, err
	}
	return NamingClientClientNameResponse{}, nil
}

// clientNameCreateRequest creates the ClientName request.
func (client *NamingClient) clientNameCreateRequest(ctx context.Context, options *NamingClientClientNameOptions) (*policy.Request, error) {
	urlPath := "/client/naming/operation"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

//   - options - NamingClientCompatibleWithEncodedNameOptions contains the optional parameters for the NamingClient.CompatibleWithEncodedName
//     method.
func (client *NamingClient) CompatibleWithEncodedName(ctx context.Context, clientNameAndJSONEncodedNameModel ClientNameAndJSONEncodedNameModel, options *NamingClientCompatibleWithEncodedNameOptions) (NamingClientCompatibleWithEncodedNameResponse, error) {
	var err error
	const operationName = "NamingClient.CompatibleWithEncodedName"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.compatibleWithEncodedNameCreateRequest(ctx, clientNameAndJSONEncodedNameModel, options)
	if err != nil {
		return NamingClientCompatibleWithEncodedNameResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NamingClientCompatibleWithEncodedNameResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return NamingClientCompatibleWithEncodedNameResponse{}, err
	}
	return NamingClientCompatibleWithEncodedNameResponse{}, nil
}

// compatibleWithEncodedNameCreateRequest creates the CompatibleWithEncodedName request.
func (client *NamingClient) compatibleWithEncodedNameCreateRequest(ctx context.Context, clientNameAndJSONEncodedNameModel ClientNameAndJSONEncodedNameModel, options *NamingClientCompatibleWithEncodedNameOptions) (*policy.Request, error) {
	urlPath := "/client/naming/property/compatible-with-encoded-name"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, clientNameAndJSONEncodedNameModel); err != nil {
		return nil, err
	}
	return req, nil
}

// - options - NamingClientLanguageOptions contains the optional parameters for the NamingClient.Language method.
func (client *NamingClient) Language(ctx context.Context, languageClientNameModel LanguageClientNameModel, options *NamingClientLanguageOptions) (NamingClientLanguageResponse, error) {
	var err error
	const operationName = "NamingClient.Language"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.languageCreateRequest(ctx, languageClientNameModel, options)
	if err != nil {
		return NamingClientLanguageResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NamingClientLanguageResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return NamingClientLanguageResponse{}, err
	}
	return NamingClientLanguageResponse{}, nil
}

// languageCreateRequest creates the Language request.
func (client *NamingClient) languageCreateRequest(ctx context.Context, languageClientNameModel LanguageClientNameModel, options *NamingClientLanguageOptions) (*policy.Request, error) {
	urlPath := "/client/naming/property/language"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, languageClientNameModel); err != nil {
		return nil, err
	}
	return req, nil
}

// - options - NamingClientParameterOptions contains the optional parameters for the NamingClient.Parameter method.
func (client *NamingClient) Parameter(ctx context.Context, clientName string, options *NamingClientParameterOptions) (NamingClientParameterResponse, error) {
	var err error
	const operationName = "NamingClient.Parameter"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.parameterCreateRequest(ctx, clientName, options)
	if err != nil {
		return NamingClientParameterResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NamingClientParameterResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return NamingClientParameterResponse{}, err
	}
	return NamingClientParameterResponse{}, nil
}

// parameterCreateRequest creates the Parameter request.
func (client *NamingClient) parameterCreateRequest(ctx context.Context, clientName string, options *NamingClientParameterOptions) (*policy.Request, error) {
	urlPath := "/client/naming/parameter"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("defaultName", clientName)
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// - options - NamingClientRequestOptions contains the optional parameters for the NamingClient.Request method.
func (client *NamingClient) Request(ctx context.Context, clientName string, options *NamingClientRequestOptions) (NamingClientRequestResponse, error) {
	var err error
	const operationName = "NamingClient.Request"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.requestCreateRequest(ctx, clientName, options)
	if err != nil {
		return NamingClientRequestResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NamingClientRequestResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return NamingClientRequestResponse{}, err
	}
	return NamingClientRequestResponse{}, nil
}

// requestCreateRequest creates the Request request.
func (client *NamingClient) requestCreateRequest(ctx context.Context, clientName string, options *NamingClientRequestOptions) (*policy.Request, error) {
	urlPath := "/client/naming/header"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["default-name"] = []string{clientName}
	return req, nil
}

// - options - NamingClientResponseOptions contains the optional parameters for the NamingClient.Response method.
func (client *NamingClient) Response(ctx context.Context, options *NamingClientResponseOptions) (NamingClientResponseResponse, error) {
	var err error
	const operationName = "NamingClient.Response"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.responseCreateRequest(ctx, options)
	if err != nil {
		return NamingClientResponseResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NamingClientResponseResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return NamingClientResponseResponse{}, err
	}
	resp, err := client.responseHandleResponse(httpResp)
	return resp, err
}

// responseCreateRequest creates the Response request.
func (client *NamingClient) responseCreateRequest(ctx context.Context, options *NamingClientResponseOptions) (*policy.Request, error) {
	urlPath := "/client/naming/header"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

// responseHandleResponse handles the Response response.
func (client *NamingClient) responseHandleResponse(resp *http.Response) (NamingClientResponseResponse, error) {
	result := NamingClientResponseResponse{}
	if val := resp.Header.Get("default-name"); val != "" {
		result.DefaultName = &val
	}
	return result, nil
}
