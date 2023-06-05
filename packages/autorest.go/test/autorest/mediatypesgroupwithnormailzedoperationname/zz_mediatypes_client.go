//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package mediatypesgroupwithnormailzedoperationname

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/streaming"
	"io"
	"net/http"
	"strings"
)

// MediaTypesClient contains the methods for the MediaTypesClient group.
// Don't use this type directly, use a constructor function instead.
type MediaTypesClient struct {
	internal *azcore.Client
}

// AnalyzeBody - Analyze body, that could be different media types.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2.0-preview
//   - options - MediaTypesClientAnalyzeBodyOptions contains the optional parameters for the MediaTypesClient.AnalyzeBody method.
func (client *MediaTypesClient) AnalyzeBody(ctx context.Context, options *MediaTypesClientAnalyzeBodyOptions) (MediaTypesClientAnalyzeBodyResponse, error) {
	var err error
	const operationName = "MediaTypesClient.AnalyzeBody"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.analyzeBodyCreateRequest(ctx, options)
	if err != nil {
		return MediaTypesClientAnalyzeBodyResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MediaTypesClientAnalyzeBodyResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return MediaTypesClientAnalyzeBodyResponse{}, err
	}
	resp, err := client.analyzeBodyHandleResponse(httpResp)
	return resp, err
}

// analyzeBodyCreateRequest creates the AnalyzeBody request.
func (client *MediaTypesClient) analyzeBodyCreateRequest(ctx context.Context, options *MediaTypesClientAnalyzeBodyOptions) (*policy.Request, error) {
	urlPath := "/mediatypes/analyze"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Input != nil {
		if err := runtime.MarshalAsJSON(req, *options.Input); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// analyzeBodyHandleResponse handles the AnalyzeBody response.
func (client *MediaTypesClient) analyzeBodyHandleResponse(resp *http.Response) (MediaTypesClientAnalyzeBodyResponse, error) {
	result := MediaTypesClientAnalyzeBodyResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return MediaTypesClientAnalyzeBodyResponse{}, err
	}
	return result, nil
}

// AnalyzeBodyNoAcceptHeader - Analyze body, that could be different media types. Adds to AnalyzeBody by not having an accept
// type.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2.0-preview
//   - options - MediaTypesClientAnalyzeBodyNoAcceptHeaderOptions contains the optional parameters for the MediaTypesClient.AnalyzeBodyNoAcceptHeader
//     method.
func (client *MediaTypesClient) AnalyzeBodyNoAcceptHeader(ctx context.Context, options *MediaTypesClientAnalyzeBodyNoAcceptHeaderOptions) (MediaTypesClientAnalyzeBodyNoAcceptHeaderResponse, error) {
	var err error
	const operationName = "MediaTypesClient.AnalyzeBodyNoAcceptHeader"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.analyzeBodyNoAcceptHeaderCreateRequest(ctx, options)
	if err != nil {
		return MediaTypesClientAnalyzeBodyNoAcceptHeaderResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MediaTypesClientAnalyzeBodyNoAcceptHeaderResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return MediaTypesClientAnalyzeBodyNoAcceptHeaderResponse{}, err
	}
	return MediaTypesClientAnalyzeBodyNoAcceptHeaderResponse{}, nil
}

// analyzeBodyNoAcceptHeaderCreateRequest creates the AnalyzeBodyNoAcceptHeader request.
func (client *MediaTypesClient) analyzeBodyNoAcceptHeaderCreateRequest(ctx context.Context, options *MediaTypesClientAnalyzeBodyNoAcceptHeaderOptions) (*policy.Request, error) {
	urlPath := "/mediatypes/analyzeNoAccept"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	if options != nil && options.Input != nil {
		if err := runtime.MarshalAsJSON(req, *options.Input); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// AnalyzeBodyNoAcceptHeaderWithBinary - Analyze body, that could be different media types. Adds to AnalyzeBody by not having
// an accept type.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2.0-preview
//   - contentType - Upload file type
//   - options - MediaTypesClientAnalyzeBodyNoAcceptHeaderWithBinaryOptions contains the optional parameters for the MediaTypesClient.AnalyzeBodyNoAcceptHeaderWithBinary
//     method.
func (client *MediaTypesClient) AnalyzeBodyNoAcceptHeaderWithBinary(ctx context.Context, contentType ContentType, options *MediaTypesClientAnalyzeBodyNoAcceptHeaderWithBinaryOptions) (MediaTypesClientAnalyzeBodyNoAcceptHeaderWithBinaryResponse, error) {
	var err error
	const operationName = "MediaTypesClient.AnalyzeBodyNoAcceptHeaderWithBinary"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.analyzeBodyNoAcceptHeaderWithBinaryCreateRequest(ctx, contentType, options)
	if err != nil {
		return MediaTypesClientAnalyzeBodyNoAcceptHeaderWithBinaryResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MediaTypesClientAnalyzeBodyNoAcceptHeaderWithBinaryResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return MediaTypesClientAnalyzeBodyNoAcceptHeaderWithBinaryResponse{}, err
	}
	return MediaTypesClientAnalyzeBodyNoAcceptHeaderWithBinaryResponse{}, nil
}

// analyzeBodyNoAcceptHeaderWithBinaryCreateRequest creates the AnalyzeBodyNoAcceptHeaderWithBinary request.
func (client *MediaTypesClient) analyzeBodyNoAcceptHeaderWithBinaryCreateRequest(ctx context.Context, contentType ContentType, options *MediaTypesClientAnalyzeBodyNoAcceptHeaderWithBinaryOptions) (*policy.Request, error) {
	urlPath := "/mediatypes/analyzeNoAccept"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Content-Type"] = []string{string(contentType)}
	if options != nil && options.Input != nil {
		if err := req.SetBody(options.Input, string(contentType)); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// AnalyzeBodyWithBinary - Analyze body, that could be different media types.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2.0-preview
//   - contentType - Upload file type
//   - options - MediaTypesClientAnalyzeBodyWithBinaryOptions contains the optional parameters for the MediaTypesClient.AnalyzeBodyWithBinary
//     method.
func (client *MediaTypesClient) AnalyzeBodyWithBinary(ctx context.Context, contentType ContentType, options *MediaTypesClientAnalyzeBodyWithBinaryOptions) (MediaTypesClientAnalyzeBodyWithBinaryResponse, error) {
	var err error
	const operationName = "MediaTypesClient.AnalyzeBodyWithBinary"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.analyzeBodyWithBinaryCreateRequest(ctx, contentType, options)
	if err != nil {
		return MediaTypesClientAnalyzeBodyWithBinaryResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MediaTypesClientAnalyzeBodyWithBinaryResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return MediaTypesClientAnalyzeBodyWithBinaryResponse{}, err
	}
	resp, err := client.analyzeBodyWithBinaryHandleResponse(httpResp)
	return resp, err
}

// analyzeBodyWithBinaryCreateRequest creates the AnalyzeBodyWithBinary request.
func (client *MediaTypesClient) analyzeBodyWithBinaryCreateRequest(ctx context.Context, contentType ContentType, options *MediaTypesClientAnalyzeBodyWithBinaryOptions) (*policy.Request, error) {
	urlPath := "/mediatypes/analyze"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Content-Type"] = []string{string(contentType)}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Input != nil {
		if err := req.SetBody(options.Input, string(contentType)); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// analyzeBodyWithBinaryHandleResponse handles the AnalyzeBodyWithBinary response.
func (client *MediaTypesClient) analyzeBodyWithBinaryHandleResponse(resp *http.Response) (MediaTypesClientAnalyzeBodyWithBinaryResponse, error) {
	result := MediaTypesClientAnalyzeBodyWithBinaryResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return MediaTypesClientAnalyzeBodyWithBinaryResponse{}, err
	}
	return result, nil
}

// BinaryBodyWithThreeContentTypesWithBinary - Binary body with three content types. Pass in string 'hello, world' with content
// type 'text/plain', {'hello': world'} with content type 'application/json' and a byte string for
// 'application/octet-stream'.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2.0-preview
//   - contentType - Upload file type
//   - message - The payload body.
//   - options - MediaTypesClientBinaryBodyWithThreeContentTypesWithBinaryOptions contains the optional parameters for the MediaTypesClient.BinaryBodyWithThreeContentTypesWithBinary
//     method.
func (client *MediaTypesClient) BinaryBodyWithThreeContentTypesWithBinary(ctx context.Context, contentType ContentType1AutoGenerated, message io.ReadSeekCloser, options *MediaTypesClientBinaryBodyWithThreeContentTypesWithBinaryOptions) (MediaTypesClientBinaryBodyWithThreeContentTypesWithBinaryResponse, error) {
	var err error
	const operationName = "MediaTypesClient.BinaryBodyWithThreeContentTypesWithBinary"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.binaryBodyWithThreeContentTypesWithBinaryCreateRequest(ctx, contentType, message, options)
	if err != nil {
		return MediaTypesClientBinaryBodyWithThreeContentTypesWithBinaryResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MediaTypesClientBinaryBodyWithThreeContentTypesWithBinaryResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return MediaTypesClientBinaryBodyWithThreeContentTypesWithBinaryResponse{}, err
	}
	resp, err := client.binaryBodyWithThreeContentTypesWithBinaryHandleResponse(httpResp)
	return resp, err
}

// binaryBodyWithThreeContentTypesWithBinaryCreateRequest creates the BinaryBodyWithThreeContentTypesWithBinary request.
func (client *MediaTypesClient) binaryBodyWithThreeContentTypesWithBinaryCreateRequest(ctx context.Context, contentType ContentType1AutoGenerated, message io.ReadSeekCloser, options *MediaTypesClientBinaryBodyWithThreeContentTypesWithBinaryOptions) (*policy.Request, error) {
	urlPath := "/mediatypes/binaryBodyThreeContentTypes"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Content-Type"] = []string{string(contentType)}
	req.Raw().Header["Accept"] = []string{"text/plain"}
	if err := req.SetBody(message, string(contentType)); err != nil {
		return nil, err
	}
	return req, nil
}

// binaryBodyWithThreeContentTypesWithBinaryHandleResponse handles the BinaryBodyWithThreeContentTypesWithBinary response.
func (client *MediaTypesClient) binaryBodyWithThreeContentTypesWithBinaryHandleResponse(resp *http.Response) (MediaTypesClientBinaryBodyWithThreeContentTypesWithBinaryResponse, error) {
	result := MediaTypesClientBinaryBodyWithThreeContentTypesWithBinaryResponse{}
	body, err := runtime.Payload(resp)
	if err != nil {
		return MediaTypesClientBinaryBodyWithThreeContentTypesWithBinaryResponse{}, err
	}
	txt := string(body)
	result.Value = &txt
	return result, nil
}

// BinaryBodyWithTwoContentTypesWithBinary - Binary body with two content types. Pass in of {'hello': 'world'} for the application/json
// content type, and a byte stream of 'hello, world!' for application/octet-stream.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2.0-preview
//   - contentType - Upload file type
//   - message - The payload body.
//   - options - MediaTypesClientBinaryBodyWithTwoContentTypesWithBinaryOptions contains the optional parameters for the MediaTypesClient.BinaryBodyWithTwoContentTypesWithBinary
//     method.
func (client *MediaTypesClient) BinaryBodyWithTwoContentTypesWithBinary(ctx context.Context, contentType ContentType1, message io.ReadSeekCloser, options *MediaTypesClientBinaryBodyWithTwoContentTypesWithBinaryOptions) (MediaTypesClientBinaryBodyWithTwoContentTypesWithBinaryResponse, error) {
	var err error
	const operationName = "MediaTypesClient.BinaryBodyWithTwoContentTypesWithBinary"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.binaryBodyWithTwoContentTypesWithBinaryCreateRequest(ctx, contentType, message, options)
	if err != nil {
		return MediaTypesClientBinaryBodyWithTwoContentTypesWithBinaryResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MediaTypesClientBinaryBodyWithTwoContentTypesWithBinaryResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return MediaTypesClientBinaryBodyWithTwoContentTypesWithBinaryResponse{}, err
	}
	resp, err := client.binaryBodyWithTwoContentTypesWithBinaryHandleResponse(httpResp)
	return resp, err
}

// binaryBodyWithTwoContentTypesWithBinaryCreateRequest creates the BinaryBodyWithTwoContentTypesWithBinary request.
func (client *MediaTypesClient) binaryBodyWithTwoContentTypesWithBinaryCreateRequest(ctx context.Context, contentType ContentType1, message io.ReadSeekCloser, options *MediaTypesClientBinaryBodyWithTwoContentTypesWithBinaryOptions) (*policy.Request, error) {
	urlPath := "/mediatypes/binaryBodyTwoContentTypes"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Content-Type"] = []string{string(contentType)}
	req.Raw().Header["Accept"] = []string{"text/plain"}
	if err := req.SetBody(message, string(contentType)); err != nil {
		return nil, err
	}
	return req, nil
}

// binaryBodyWithTwoContentTypesWithBinaryHandleResponse handles the BinaryBodyWithTwoContentTypesWithBinary response.
func (client *MediaTypesClient) binaryBodyWithTwoContentTypesWithBinaryHandleResponse(resp *http.Response) (MediaTypesClientBinaryBodyWithTwoContentTypesWithBinaryResponse, error) {
	result := MediaTypesClientBinaryBodyWithTwoContentTypesWithBinaryResponse{}
	body, err := runtime.Payload(resp)
	if err != nil {
		return MediaTypesClientBinaryBodyWithTwoContentTypesWithBinaryResponse{}, err
	}
	txt := string(body)
	result.Value = &txt
	return result, nil
}

// ContentTypeWithEncodingWithText - Pass in contentType 'text/plain; charset=UTF-8' to pass test. Value for input does not
// matter
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2.0-preview
//   - options - MediaTypesClientContentTypeWithEncodingWithTextOptions contains the optional parameters for the MediaTypesClient.ContentTypeWithEncodingWithText
//     method.
func (client *MediaTypesClient) ContentTypeWithEncodingWithText(ctx context.Context, options *MediaTypesClientContentTypeWithEncodingWithTextOptions) (MediaTypesClientContentTypeWithEncodingWithTextResponse, error) {
	var err error
	const operationName = "MediaTypesClient.ContentTypeWithEncodingWithText"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.contentTypeWithEncodingWithTextCreateRequest(ctx, options)
	if err != nil {
		return MediaTypesClientContentTypeWithEncodingWithTextResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MediaTypesClientContentTypeWithEncodingWithTextResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return MediaTypesClientContentTypeWithEncodingWithTextResponse{}, err
	}
	resp, err := client.contentTypeWithEncodingWithTextHandleResponse(httpResp)
	return resp, err
}

// contentTypeWithEncodingWithTextCreateRequest creates the ContentTypeWithEncodingWithText request.
func (client *MediaTypesClient) contentTypeWithEncodingWithTextCreateRequest(ctx context.Context, options *MediaTypesClientContentTypeWithEncodingWithTextOptions) (*policy.Request, error) {
	urlPath := "/mediatypes/contentTypeWithEncoding"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Input != nil {
		body := streaming.NopCloser(strings.NewReader(*options.Input))
		if err := req.SetBody(body, "text/plain; charset=UTF-8"); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// contentTypeWithEncodingWithTextHandleResponse handles the ContentTypeWithEncodingWithText response.
func (client *MediaTypesClient) contentTypeWithEncodingWithTextHandleResponse(resp *http.Response) (MediaTypesClientContentTypeWithEncodingWithTextResponse, error) {
	result := MediaTypesClientContentTypeWithEncodingWithTextResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return MediaTypesClientContentTypeWithEncodingWithTextResponse{}, err
	}
	return result, nil
}

// PutTextAndJSONBodyWithText - Body that's either text/plain or application/json
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2.0-preview
//   - contentType - Upload file type
//   - message - The payload body.
//   - options - MediaTypesClientPutTextAndJSONBodyWithTextOptions contains the optional parameters for the MediaTypesClient.PutTextAndJSONBodyWithText
//     method.
func (client *MediaTypesClient) PutTextAndJSONBodyWithText(ctx context.Context, contentType ContentType1AutoGenerated2, message string, options *MediaTypesClientPutTextAndJSONBodyWithTextOptions) (MediaTypesClientPutTextAndJSONBodyWithTextResponse, error) {
	var err error
	const operationName = "MediaTypesClient.PutTextAndJSONBodyWithText"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putTextAndJSONBodyWithTextCreateRequest(ctx, contentType, message, options)
	if err != nil {
		return MediaTypesClientPutTextAndJSONBodyWithTextResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MediaTypesClientPutTextAndJSONBodyWithTextResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return MediaTypesClientPutTextAndJSONBodyWithTextResponse{}, err
	}
	resp, err := client.putTextAndJSONBodyWithTextHandleResponse(httpResp)
	return resp, err
}

// putTextAndJSONBodyWithTextCreateRequest creates the PutTextAndJSONBodyWithText request.
func (client *MediaTypesClient) putTextAndJSONBodyWithTextCreateRequest(ctx context.Context, contentType ContentType1AutoGenerated2, message string, options *MediaTypesClientPutTextAndJSONBodyWithTextOptions) (*policy.Request, error) {
	urlPath := "/mediatypes/textAndJson"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Content-Type"] = []string{string(contentType)}
	req.Raw().Header["Accept"] = []string{"text/plain"}
	body := streaming.NopCloser(strings.NewReader(message))
	if err := req.SetBody(body, "application/json"); err != nil {
		return nil, err
	}

	return req, nil
}

// putTextAndJSONBodyWithTextHandleResponse handles the PutTextAndJSONBodyWithText response.
func (client *MediaTypesClient) putTextAndJSONBodyWithTextHandleResponse(resp *http.Response) (MediaTypesClientPutTextAndJSONBodyWithTextResponse, error) {
	result := MediaTypesClientPutTextAndJSONBodyWithTextResponse{}
	body, err := runtime.Payload(resp)
	if err != nil {
		return MediaTypesClientPutTextAndJSONBodyWithTextResponse{}, err
	}
	txt := string(body)
	result.Value = &txt
	return result, nil
}
