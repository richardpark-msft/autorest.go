//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package stringgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/streaming"
	"net/http"
	"strings"
)

// StringClient contains the methods for the String group.
// Don't use this type directly, use a constructor function instead.
type StringClient struct {
	internal *azcore.Client
}

// GetBase64Encoded - Get value that is base64 encoded
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - StringClientGetBase64EncodedOptions contains the optional parameters for the StringClient.GetBase64Encoded method.
func (client *StringClient) GetBase64Encoded(ctx context.Context, options *StringClientGetBase64EncodedOptions) (StringClientGetBase64EncodedResponse, error) {
	var err error
	const operationName = "StringClient.GetBase64Encoded"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getBase64EncodedCreateRequest(ctx, options)
	if err != nil {
		return StringClientGetBase64EncodedResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return StringClientGetBase64EncodedResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return StringClientGetBase64EncodedResponse{}, err
	}
	resp, err := client.getBase64EncodedHandleResponse(httpResp)
	return resp, err
}

// getBase64EncodedCreateRequest creates the GetBase64Encoded request.
func (client *StringClient) getBase64EncodedCreateRequest(ctx context.Context, options *StringClientGetBase64EncodedOptions) (*policy.Request, error) {
	urlPath := "/string/base64Encoding"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getBase64EncodedHandleResponse handles the GetBase64Encoded response.
func (client *StringClient) getBase64EncodedHandleResponse(resp *http.Response) (StringClientGetBase64EncodedResponse, error) {
	result := StringClientGetBase64EncodedResponse{}
	if err := runtime.UnmarshalAsByteArray(resp, &result.Value, runtime.Base64StdFormat); err != nil {
		return StringClientGetBase64EncodedResponse{}, err
	}
	return result, nil
}

// GetBase64URLEncoded - Get value that is base64url encoded
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - StringClientGetBase64URLEncodedOptions contains the optional parameters for the StringClient.GetBase64URLEncoded
//     method.
func (client *StringClient) GetBase64URLEncoded(ctx context.Context, options *StringClientGetBase64URLEncodedOptions) (StringClientGetBase64URLEncodedResponse, error) {
	var err error
	const operationName = "StringClient.GetBase64URLEncoded"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getBase64URLEncodedCreateRequest(ctx, options)
	if err != nil {
		return StringClientGetBase64URLEncodedResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return StringClientGetBase64URLEncodedResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return StringClientGetBase64URLEncodedResponse{}, err
	}
	resp, err := client.getBase64URLEncodedHandleResponse(httpResp)
	return resp, err
}

// getBase64URLEncodedCreateRequest creates the GetBase64URLEncoded request.
func (client *StringClient) getBase64URLEncodedCreateRequest(ctx context.Context, options *StringClientGetBase64URLEncodedOptions) (*policy.Request, error) {
	urlPath := "/string/base64UrlEncoding"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getBase64URLEncodedHandleResponse handles the GetBase64URLEncoded response.
func (client *StringClient) getBase64URLEncodedHandleResponse(resp *http.Response) (StringClientGetBase64URLEncodedResponse, error) {
	result := StringClientGetBase64URLEncodedResponse{}
	if err := runtime.UnmarshalAsByteArray(resp, &result.Value, runtime.Base64URLFormat); err != nil {
		return StringClientGetBase64URLEncodedResponse{}, err
	}
	return result, nil
}

// GetEmpty - Get empty string value value ”
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - StringClientGetEmptyOptions contains the optional parameters for the StringClient.GetEmpty method.
func (client *StringClient) GetEmpty(ctx context.Context, options *StringClientGetEmptyOptions) (StringClientGetEmptyResponse, error) {
	var err error
	const operationName = "StringClient.GetEmpty"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getEmptyCreateRequest(ctx, options)
	if err != nil {
		return StringClientGetEmptyResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return StringClientGetEmptyResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return StringClientGetEmptyResponse{}, err
	}
	resp, err := client.getEmptyHandleResponse(httpResp)
	return resp, err
}

// getEmptyCreateRequest creates the GetEmpty request.
func (client *StringClient) getEmptyCreateRequest(ctx context.Context, options *StringClientGetEmptyOptions) (*policy.Request, error) {
	urlPath := "/string/empty"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getEmptyHandleResponse handles the GetEmpty response.
func (client *StringClient) getEmptyHandleResponse(resp *http.Response) (StringClientGetEmptyResponse, error) {
	result := StringClientGetEmptyResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return StringClientGetEmptyResponse{}, err
	}
	return result, nil
}

// GetMBCS - Get mbcs string value '啊齄丂狛狜隣郎隣兀﨩ˊ〞〡￤℡㈱‐ー﹡﹢﹫、〓ⅰⅹ⒈€㈠㈩ⅠⅫ！￣ぁんァヶΑ︴АЯаяāɡㄅㄩ─╋︵﹄︻︱︳︴ⅰⅹɑɡ〇〾⿻⺁䜣€'
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - StringClientGetMBCSOptions contains the optional parameters for the StringClient.GetMBCS method.
func (client *StringClient) GetMBCS(ctx context.Context, options *StringClientGetMBCSOptions) (StringClientGetMBCSResponse, error) {
	var err error
	const operationName = "StringClient.GetMBCS"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getMBCSCreateRequest(ctx, options)
	if err != nil {
		return StringClientGetMBCSResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return StringClientGetMBCSResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return StringClientGetMBCSResponse{}, err
	}
	resp, err := client.getMBCSHandleResponse(httpResp)
	return resp, err
}

// getMBCSCreateRequest creates the GetMBCS request.
func (client *StringClient) getMBCSCreateRequest(ctx context.Context, options *StringClientGetMBCSOptions) (*policy.Request, error) {
	urlPath := "/string/mbcs"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getMBCSHandleResponse handles the GetMBCS response.
func (client *StringClient) getMBCSHandleResponse(resp *http.Response) (StringClientGetMBCSResponse, error) {
	result := StringClientGetMBCSResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return StringClientGetMBCSResponse{}, err
	}
	return result, nil
}

// GetNotProvided - Get String value when no string value is sent in response payload
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - StringClientGetNotProvidedOptions contains the optional parameters for the StringClient.GetNotProvided method.
func (client *StringClient) GetNotProvided(ctx context.Context, options *StringClientGetNotProvidedOptions) (StringClientGetNotProvidedResponse, error) {
	var err error
	const operationName = "StringClient.GetNotProvided"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getNotProvidedCreateRequest(ctx, options)
	if err != nil {
		return StringClientGetNotProvidedResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return StringClientGetNotProvidedResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return StringClientGetNotProvidedResponse{}, err
	}
	resp, err := client.getNotProvidedHandleResponse(httpResp)
	return resp, err
}

// getNotProvidedCreateRequest creates the GetNotProvided request.
func (client *StringClient) getNotProvidedCreateRequest(ctx context.Context, options *StringClientGetNotProvidedOptions) (*policy.Request, error) {
	urlPath := "/string/notProvided"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getNotProvidedHandleResponse handles the GetNotProvided response.
func (client *StringClient) getNotProvidedHandleResponse(resp *http.Response) (StringClientGetNotProvidedResponse, error) {
	result := StringClientGetNotProvidedResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return StringClientGetNotProvidedResponse{}, err
	}
	return result, nil
}

// GetNull - Get null string value value
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - StringClientGetNullOptions contains the optional parameters for the StringClient.GetNull method.
func (client *StringClient) GetNull(ctx context.Context, options *StringClientGetNullOptions) (StringClientGetNullResponse, error) {
	var err error
	const operationName = "StringClient.GetNull"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getNullCreateRequest(ctx, options)
	if err != nil {
		return StringClientGetNullResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return StringClientGetNullResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return StringClientGetNullResponse{}, err
	}
	resp, err := client.getNullHandleResponse(httpResp)
	return resp, err
}

// getNullCreateRequest creates the GetNull request.
func (client *StringClient) getNullCreateRequest(ctx context.Context, options *StringClientGetNullOptions) (*policy.Request, error) {
	urlPath := "/string/null"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getNullHandleResponse handles the GetNull response.
func (client *StringClient) getNullHandleResponse(resp *http.Response) (StringClientGetNullResponse, error) {
	result := StringClientGetNullResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return StringClientGetNullResponse{}, err
	}
	return result, nil
}

// GetNullBase64URLEncoded - Get null value that is expected to be base64url encoded
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - StringClientGetNullBase64URLEncodedOptions contains the optional parameters for the StringClient.GetNullBase64URLEncoded
//     method.
func (client *StringClient) GetNullBase64URLEncoded(ctx context.Context, options *StringClientGetNullBase64URLEncodedOptions) (StringClientGetNullBase64URLEncodedResponse, error) {
	var err error
	const operationName = "StringClient.GetNullBase64URLEncoded"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getNullBase64URLEncodedCreateRequest(ctx, options)
	if err != nil {
		return StringClientGetNullBase64URLEncodedResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return StringClientGetNullBase64URLEncodedResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return StringClientGetNullBase64URLEncodedResponse{}, err
	}
	resp, err := client.getNullBase64URLEncodedHandleResponse(httpResp)
	return resp, err
}

// getNullBase64URLEncodedCreateRequest creates the GetNullBase64URLEncoded request.
func (client *StringClient) getNullBase64URLEncodedCreateRequest(ctx context.Context, options *StringClientGetNullBase64URLEncodedOptions) (*policy.Request, error) {
	urlPath := "/string/nullBase64UrlEncoding"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getNullBase64URLEncodedHandleResponse handles the GetNullBase64URLEncoded response.
func (client *StringClient) getNullBase64URLEncodedHandleResponse(resp *http.Response) (StringClientGetNullBase64URLEncodedResponse, error) {
	result := StringClientGetNullBase64URLEncodedResponse{}
	if err := runtime.UnmarshalAsByteArray(resp, &result.Value, runtime.Base64URLFormat); err != nil {
		return StringClientGetNullBase64URLEncodedResponse{}, err
	}
	return result, nil
}

// GetWhitespace - Get string value with leading and trailing whitespace 'Now is the time for all good men to come to the
// aid of their country'
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - StringClientGetWhitespaceOptions contains the optional parameters for the StringClient.GetWhitespace method.
func (client *StringClient) GetWhitespace(ctx context.Context, options *StringClientGetWhitespaceOptions) (StringClientGetWhitespaceResponse, error) {
	var err error
	const operationName = "StringClient.GetWhitespace"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getWhitespaceCreateRequest(ctx, options)
	if err != nil {
		return StringClientGetWhitespaceResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return StringClientGetWhitespaceResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return StringClientGetWhitespaceResponse{}, err
	}
	resp, err := client.getWhitespaceHandleResponse(httpResp)
	return resp, err
}

// getWhitespaceCreateRequest creates the GetWhitespace request.
func (client *StringClient) getWhitespaceCreateRequest(ctx context.Context, options *StringClientGetWhitespaceOptions) (*policy.Request, error) {
	urlPath := "/string/whitespace"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getWhitespaceHandleResponse handles the GetWhitespace response.
func (client *StringClient) getWhitespaceHandleResponse(resp *http.Response) (StringClientGetWhitespaceResponse, error) {
	result := StringClientGetWhitespaceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return StringClientGetWhitespaceResponse{}, err
	}
	return result, nil
}

// PutBase64URLEncoded - Put value that is base64url encoded
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - stringBody - string body
//   - options - StringClientPutBase64URLEncodedOptions contains the optional parameters for the StringClient.PutBase64URLEncoded
//     method.
func (client *StringClient) PutBase64URLEncoded(ctx context.Context, stringBody []byte, options *StringClientPutBase64URLEncodedOptions) (StringClientPutBase64URLEncodedResponse, error) {
	var err error
	const operationName = "StringClient.PutBase64URLEncoded"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putBase64URLEncodedCreateRequest(ctx, stringBody, options)
	if err != nil {
		return StringClientPutBase64URLEncodedResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return StringClientPutBase64URLEncodedResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return StringClientPutBase64URLEncodedResponse{}, err
	}
	return StringClientPutBase64URLEncodedResponse{}, nil
}

// putBase64URLEncodedCreateRequest creates the PutBase64URLEncoded request.
func (client *StringClient) putBase64URLEncodedCreateRequest(ctx context.Context, stringBody []byte, options *StringClientPutBase64URLEncodedOptions) (*policy.Request, error) {
	urlPath := "/string/base64UrlEncoding"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsByteArray(req, stringBody, runtime.Base64URLFormat); err != nil {
		return nil, err
	}
	return req, nil
}

// PutEmpty - Set string value empty ”
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - StringClientPutEmptyOptions contains the optional parameters for the StringClient.PutEmpty method.
func (client *StringClient) PutEmpty(ctx context.Context, options *StringClientPutEmptyOptions) (StringClientPutEmptyResponse, error) {
	var err error
	const operationName = "StringClient.PutEmpty"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putEmptyCreateRequest(ctx, options)
	if err != nil {
		return StringClientPutEmptyResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return StringClientPutEmptyResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return StringClientPutEmptyResponse{}, err
	}
	return StringClientPutEmptyResponse{}, nil
}

// putEmptyCreateRequest creates the PutEmpty request.
func (client *StringClient) putEmptyCreateRequest(ctx context.Context, options *StringClientPutEmptyOptions) (*policy.Request, error) {
	urlPath := "/string/empty"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, ""); err != nil {
		return nil, err
	}
	return req, nil
}

// PutMBCS - Set string value mbcs '啊齄丂狛狜隣郎隣兀﨩ˊ〞〡￤℡㈱‐ー﹡﹢﹫、〓ⅰⅹ⒈€㈠㈩ⅠⅫ！￣ぁんァヶΑ︴АЯаяāɡㄅㄩ─╋︵﹄︻︱︳︴ⅰⅹɑɡ〇〾⿻⺁䜣€'
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - StringClientPutMBCSOptions contains the optional parameters for the StringClient.PutMBCS method.
func (client *StringClient) PutMBCS(ctx context.Context, options *StringClientPutMBCSOptions) (StringClientPutMBCSResponse, error) {
	var err error
	const operationName = "StringClient.PutMBCS"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putMBCSCreateRequest(ctx, options)
	if err != nil {
		return StringClientPutMBCSResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return StringClientPutMBCSResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return StringClientPutMBCSResponse{}, err
	}
	return StringClientPutMBCSResponse{}, nil
}

// putMBCSCreateRequest creates the PutMBCS request.
func (client *StringClient) putMBCSCreateRequest(ctx context.Context, options *StringClientPutMBCSOptions) (*policy.Request, error) {
	urlPath := "/string/mbcs"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, "啊齄丂狛狜隣郎隣兀﨩ˊ〞〡￤℡㈱‐ー﹡﹢﹫、〓ⅰⅹ⒈€㈠㈩ⅠⅫ！￣ぁんァヶΑ︴АЯаяāɡㄅㄩ─╋︵﹄︻︱︳︴ⅰⅹɑɡ〇〾⿻⺁䜣€"); err != nil {
		return nil, err
	}
	return req, nil
}

// PutNull - Set string value null
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - stringBody - string body
//   - options - StringClientPutNullOptions contains the optional parameters for the StringClient.PutNull method.
func (client *StringClient) PutNull(ctx context.Context, stringBody string, options *StringClientPutNullOptions) (StringClientPutNullResponse, error) {
	var err error
	const operationName = "StringClient.PutNull"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putNullCreateRequest(ctx, stringBody, options)
	if err != nil {
		return StringClientPutNullResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return StringClientPutNullResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return StringClientPutNullResponse{}, err
	}
	return StringClientPutNullResponse{}, nil
}

// putNullCreateRequest creates the PutNull request.
func (client *StringClient) putNullCreateRequest(ctx context.Context, stringBody string, options *StringClientPutNullOptions) (*policy.Request, error) {
	urlPath := "/string/null"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	body := streaming.NopCloser(strings.NewReader(stringBody))
	if err := req.SetBody(body, "application/json"); err != nil {
		return nil, err
	}

	return req, nil
}

// PutWhitespace - Set String value with leading and trailing whitespace 'Now is the time for all good men to come to the
// aid of their country'
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - StringClientPutWhitespaceOptions contains the optional parameters for the StringClient.PutWhitespace method.
func (client *StringClient) PutWhitespace(ctx context.Context, options *StringClientPutWhitespaceOptions) (StringClientPutWhitespaceResponse, error) {
	var err error
	const operationName = "StringClient.PutWhitespace"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putWhitespaceCreateRequest(ctx, options)
	if err != nil {
		return StringClientPutWhitespaceResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return StringClientPutWhitespaceResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return StringClientPutWhitespaceResponse{}, err
	}
	return StringClientPutWhitespaceResponse{}, nil
}

// putWhitespaceCreateRequest creates the PutWhitespace request.
func (client *StringClient) putWhitespaceCreateRequest(ctx context.Context, options *StringClientPutWhitespaceOptions) (*policy.Request, error) {
	urlPath := "/string/whitespace"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, "    Now is the time for all good men to come to the aid of their country    "); err != nil {
		return nil, err
	}
	return req, nil
}
