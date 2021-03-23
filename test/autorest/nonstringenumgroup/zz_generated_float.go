// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package nonstringenumgroup

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"io/ioutil"
	"net/http"
)

// FloatClient contains the methods for the Float group.
// Don't use this type directly, use NewFloatClient() instead.
type FloatClient struct {
	con *Connection
}

// NewFloatClient creates a new instance of FloatClient with the specified values.
func NewFloatClient(con *Connection) *FloatClient {
	return &FloatClient{con: con}
}

// Get - Get a float enum
func (client *FloatClient) Get(ctx context.Context, options *FloatGetOptions) (FloatEnumResponse, error) {
	req, err := client.getCreateRequest(ctx, options)
	if err != nil {
		return FloatEnumResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return FloatEnumResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return FloatEnumResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *FloatClient) getCreateRequest(ctx context.Context, options *FloatGetOptions) (*azcore.Request, error) {
	urlPath := "/nonStringEnums/float/get"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *FloatClient) getHandleResponse(resp *azcore.Response) (FloatEnumResponse, error) {
	var val *FloatEnum
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return FloatEnumResponse{}, err
	}
	return FloatEnumResponse{RawResponse: resp.Response, Value: val}, nil
}

// getHandleError handles the Get error response.
func (client *FloatClient) getHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// Put - Put a float enum
func (client *FloatClient) Put(ctx context.Context, options *FloatPutOptions) (StringResponse, error) {
	req, err := client.putCreateRequest(ctx, options)
	if err != nil {
		return StringResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return StringResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return StringResponse{}, client.putHandleError(resp)
	}
	return client.putHandleResponse(resp)
}

// putCreateRequest creates the Put request.
func (client *FloatClient) putCreateRequest(ctx context.Context, options *FloatPutOptions) (*azcore.Request, error) {
	urlPath := "/nonStringEnums/float/put"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	if options != nil && options.Input != nil {
		return req, req.MarshalAsJSON(options.Input)
	}
	return req, nil
}

// putHandleResponse handles the Put response.
func (client *FloatClient) putHandleResponse(resp *azcore.Response) (StringResponse, error) {
	var val *string
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return StringResponse{}, err
	}
	return StringResponse{RawResponse: resp.Response, Value: val}, nil
}

// putHandleError handles the Put error response.
func (client *FloatClient) putHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}
