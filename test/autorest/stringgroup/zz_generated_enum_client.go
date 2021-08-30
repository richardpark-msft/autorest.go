//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package stringgroup

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// EnumClient contains the methods for the Enum group.
// Don't use this type directly, use NewEnumClient() instead.
type EnumClient struct {
	con *Connection
}

// NewEnumClient creates a new instance of EnumClient with the specified values.
func NewEnumClient(con *Connection) *EnumClient {
	return &EnumClient{con: con}
}

// GetNotExpandable - Get enum value 'red color' from enumeration of 'red color', 'green-color', 'blue_color'.
// If the operation fails it returns the *Error error type.
func (client *EnumClient) GetNotExpandable(ctx context.Context, options *EnumGetNotExpandableOptions) (EnumGetNotExpandableResponse, error) {
	req, err := client.getNotExpandableCreateRequest(ctx, options)
	if err != nil {
		return EnumGetNotExpandableResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return EnumGetNotExpandableResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return EnumGetNotExpandableResponse{}, client.getNotExpandableHandleError(resp)
	}
	return client.getNotExpandableHandleResponse(resp)
}

// getNotExpandableCreateRequest creates the GetNotExpandable request.
func (client *EnumClient) getNotExpandableCreateRequest(ctx context.Context, options *EnumGetNotExpandableOptions) (*policy.Request, error) {
	urlPath := "/string/enum/notExpandable"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getNotExpandableHandleResponse handles the GetNotExpandable response.
func (client *EnumClient) getNotExpandableHandleResponse(resp *http.Response) (EnumGetNotExpandableResponse, error) {
	result := EnumGetNotExpandableResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return EnumGetNotExpandableResponse{}, err
	}
	return result, nil
}

// getNotExpandableHandleError handles the GetNotExpandable error response.
func (client *EnumClient) getNotExpandableHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// GetReferenced - Get enum value 'red color' from enumeration of 'red color', 'green-color', 'blue_color'.
// If the operation fails it returns the *Error error type.
func (client *EnumClient) GetReferenced(ctx context.Context, options *EnumGetReferencedOptions) (EnumGetReferencedResponse, error) {
	req, err := client.getReferencedCreateRequest(ctx, options)
	if err != nil {
		return EnumGetReferencedResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return EnumGetReferencedResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return EnumGetReferencedResponse{}, client.getReferencedHandleError(resp)
	}
	return client.getReferencedHandleResponse(resp)
}

// getReferencedCreateRequest creates the GetReferenced request.
func (client *EnumClient) getReferencedCreateRequest(ctx context.Context, options *EnumGetReferencedOptions) (*policy.Request, error) {
	urlPath := "/string/enum/Referenced"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getReferencedHandleResponse handles the GetReferenced response.
func (client *EnumClient) getReferencedHandleResponse(resp *http.Response) (EnumGetReferencedResponse, error) {
	result := EnumGetReferencedResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return EnumGetReferencedResponse{}, err
	}
	return result, nil
}

// getReferencedHandleError handles the GetReferenced error response.
func (client *EnumClient) getReferencedHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// GetReferencedConstant - Get value 'green-color' from the constant.
// If the operation fails it returns the *Error error type.
func (client *EnumClient) GetReferencedConstant(ctx context.Context, options *EnumGetReferencedConstantOptions) (EnumGetReferencedConstantResponse, error) {
	req, err := client.getReferencedConstantCreateRequest(ctx, options)
	if err != nil {
		return EnumGetReferencedConstantResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return EnumGetReferencedConstantResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return EnumGetReferencedConstantResponse{}, client.getReferencedConstantHandleError(resp)
	}
	return client.getReferencedConstantHandleResponse(resp)
}

// getReferencedConstantCreateRequest creates the GetReferencedConstant request.
func (client *EnumClient) getReferencedConstantCreateRequest(ctx context.Context, options *EnumGetReferencedConstantOptions) (*policy.Request, error) {
	urlPath := "/string/enum/ReferencedConstant"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getReferencedConstantHandleResponse handles the GetReferencedConstant response.
func (client *EnumClient) getReferencedConstantHandleResponse(resp *http.Response) (EnumGetReferencedConstantResponse, error) {
	result := EnumGetReferencedConstantResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RefColorConstant); err != nil {
		return EnumGetReferencedConstantResponse{}, err
	}
	return result, nil
}

// getReferencedConstantHandleError handles the GetReferencedConstant error response.
func (client *EnumClient) getReferencedConstantHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// PutNotExpandable - Sends value 'red color' from enumeration of 'red color', 'green-color', 'blue_color'
// If the operation fails it returns the *Error error type.
func (client *EnumClient) PutNotExpandable(ctx context.Context, stringBody Colors, options *EnumPutNotExpandableOptions) (EnumPutNotExpandableResponse, error) {
	req, err := client.putNotExpandableCreateRequest(ctx, stringBody, options)
	if err != nil {
		return EnumPutNotExpandableResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return EnumPutNotExpandableResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return EnumPutNotExpandableResponse{}, client.putNotExpandableHandleError(resp)
	}
	return EnumPutNotExpandableResponse{RawResponse: resp}, nil
}

// putNotExpandableCreateRequest creates the PutNotExpandable request.
func (client *EnumClient) putNotExpandableCreateRequest(ctx context.Context, stringBody Colors, options *EnumPutNotExpandableOptions) (*policy.Request, error) {
	urlPath := "/string/enum/notExpandable"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, stringBody)
}

// putNotExpandableHandleError handles the PutNotExpandable error response.
func (client *EnumClient) putNotExpandableHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// PutReferenced - Sends value 'red color' from enumeration of 'red color', 'green-color', 'blue_color'
// If the operation fails it returns the *Error error type.
func (client *EnumClient) PutReferenced(ctx context.Context, enumStringBody Colors, options *EnumPutReferencedOptions) (EnumPutReferencedResponse, error) {
	req, err := client.putReferencedCreateRequest(ctx, enumStringBody, options)
	if err != nil {
		return EnumPutReferencedResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return EnumPutReferencedResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return EnumPutReferencedResponse{}, client.putReferencedHandleError(resp)
	}
	return EnumPutReferencedResponse{RawResponse: resp}, nil
}

// putReferencedCreateRequest creates the PutReferenced request.
func (client *EnumClient) putReferencedCreateRequest(ctx context.Context, enumStringBody Colors, options *EnumPutReferencedOptions) (*policy.Request, error) {
	urlPath := "/string/enum/Referenced"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, enumStringBody)
}

// putReferencedHandleError handles the PutReferenced error response.
func (client *EnumClient) putReferencedHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// PutReferencedConstant - Sends value 'green-color' from a constant
// If the operation fails it returns the *Error error type.
func (client *EnumClient) PutReferencedConstant(ctx context.Context, enumStringBody RefColorConstant, options *EnumPutReferencedConstantOptions) (EnumPutReferencedConstantResponse, error) {
	req, err := client.putReferencedConstantCreateRequest(ctx, enumStringBody, options)
	if err != nil {
		return EnumPutReferencedConstantResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return EnumPutReferencedConstantResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return EnumPutReferencedConstantResponse{}, client.putReferencedConstantHandleError(resp)
	}
	return EnumPutReferencedConstantResponse{RawResponse: resp}, nil
}

// putReferencedConstantCreateRequest creates the PutReferencedConstant request.
func (client *EnumClient) putReferencedConstantCreateRequest(ctx context.Context, enumStringBody RefColorConstant, options *EnumPutReferencedConstantOptions) (*policy.Request, error) {
	urlPath := "/string/enum/ReferencedConstant"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, enumStringBody)
}

// putReferencedConstantHandleError handles the PutReferencedConstant error response.
func (client *EnumClient) putReferencedConstantHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
