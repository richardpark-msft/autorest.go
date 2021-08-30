//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package httpinfrastructuregroup

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// HTTPRetryClient contains the methods for the HTTPRetry group.
// Don't use this type directly, use NewHTTPRetryClient() instead.
type HTTPRetryClient struct {
	con *Connection
}

// NewHTTPRetryClient creates a new instance of HTTPRetryClient with the specified values.
func NewHTTPRetryClient(con *Connection) *HTTPRetryClient {
	return &HTTPRetryClient{con: con}
}

// Delete503 - Return 503 status code, then 200 after retry
// If the operation fails it returns the *Error error type.
func (client *HTTPRetryClient) Delete503(ctx context.Context, options *HTTPRetryDelete503Options) (HTTPRetryDelete503Response, error) {
	req, err := client.delete503CreateRequest(ctx, options)
	if err != nil {
		return HTTPRetryDelete503Response{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return HTTPRetryDelete503Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HTTPRetryDelete503Response{}, client.delete503HandleError(resp)
	}
	return HTTPRetryDelete503Response{RawResponse: resp}, nil
}

// delete503CreateRequest creates the Delete503 request.
func (client *HTTPRetryClient) delete503CreateRequest(ctx context.Context, options *HTTPRetryDelete503Options) (*policy.Request, error) {
	urlPath := "/http/retry/503"
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, true)
}

// delete503HandleError handles the Delete503 error response.
func (client *HTTPRetryClient) delete503HandleError(resp *http.Response) error {
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

// Get502 - Return 502 status code, then 200 after retry
// If the operation fails it returns the *Error error type.
func (client *HTTPRetryClient) Get502(ctx context.Context, options *HTTPRetryGet502Options) (HTTPRetryGet502Response, error) {
	req, err := client.get502CreateRequest(ctx, options)
	if err != nil {
		return HTTPRetryGet502Response{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return HTTPRetryGet502Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HTTPRetryGet502Response{}, client.get502HandleError(resp)
	}
	return HTTPRetryGet502Response{RawResponse: resp}, nil
}

// get502CreateRequest creates the Get502 request.
func (client *HTTPRetryClient) get502CreateRequest(ctx context.Context, options *HTTPRetryGet502Options) (*policy.Request, error) {
	urlPath := "/http/retry/502"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// get502HandleError handles the Get502 error response.
func (client *HTTPRetryClient) get502HandleError(resp *http.Response) error {
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

// Head408 - Return 408 status code, then 200 after retry
// If the operation fails it returns the *Error error type.
func (client *HTTPRetryClient) Head408(ctx context.Context, options *HTTPRetryHead408Options) (HTTPRetryHead408Response, error) {
	req, err := client.head408CreateRequest(ctx, options)
	if err != nil {
		return HTTPRetryHead408Response{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return HTTPRetryHead408Response{}, err
	}
	result := HTTPRetryHead408Response{RawResponse: resp}
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		result.Success = true
	}
	return result, nil
}

// head408CreateRequest creates the Head408 request.
func (client *HTTPRetryClient) head408CreateRequest(ctx context.Context, options *HTTPRetryHead408Options) (*policy.Request, error) {
	urlPath := "/http/retry/408"
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Options502 - Return 502 status code, then 200 after retry
// If the operation fails it returns the *Error error type.
func (client *HTTPRetryClient) Options502(ctx context.Context, options *HTTPRetryOptions502Options) (HTTPRetryOptions502Response, error) {
	req, err := client.options502CreateRequest(ctx, options)
	if err != nil {
		return HTTPRetryOptions502Response{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return HTTPRetryOptions502Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HTTPRetryOptions502Response{}, client.options502HandleError(resp)
	}
	return client.options502HandleResponse(resp)
}

// options502CreateRequest creates the Options502 request.
func (client *HTTPRetryClient) options502CreateRequest(ctx context.Context, options *HTTPRetryOptions502Options) (*policy.Request, error) {
	urlPath := "/http/retry/502"
	req, err := runtime.NewRequest(ctx, http.MethodOptions, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// options502HandleResponse handles the Options502 response.
func (client *HTTPRetryClient) options502HandleResponse(resp *http.Response) (HTTPRetryOptions502Response, error) {
	result := HTTPRetryOptions502Response{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return HTTPRetryOptions502Response{}, err
	}
	return result, nil
}

// options502HandleError handles the Options502 error response.
func (client *HTTPRetryClient) options502HandleError(resp *http.Response) error {
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

// Patch500 - Return 500 status code, then 200 after retry
// If the operation fails it returns the *Error error type.
func (client *HTTPRetryClient) Patch500(ctx context.Context, options *HTTPRetryPatch500Options) (HTTPRetryPatch500Response, error) {
	req, err := client.patch500CreateRequest(ctx, options)
	if err != nil {
		return HTTPRetryPatch500Response{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return HTTPRetryPatch500Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HTTPRetryPatch500Response{}, client.patch500HandleError(resp)
	}
	return HTTPRetryPatch500Response{RawResponse: resp}, nil
}

// patch500CreateRequest creates the Patch500 request.
func (client *HTTPRetryClient) patch500CreateRequest(ctx context.Context, options *HTTPRetryPatch500Options) (*policy.Request, error) {
	urlPath := "/http/retry/500"
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, true)
}

// patch500HandleError handles the Patch500 error response.
func (client *HTTPRetryClient) patch500HandleError(resp *http.Response) error {
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

// Patch504 - Return 504 status code, then 200 after retry
// If the operation fails it returns the *Error error type.
func (client *HTTPRetryClient) Patch504(ctx context.Context, options *HTTPRetryPatch504Options) (HTTPRetryPatch504Response, error) {
	req, err := client.patch504CreateRequest(ctx, options)
	if err != nil {
		return HTTPRetryPatch504Response{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return HTTPRetryPatch504Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HTTPRetryPatch504Response{}, client.patch504HandleError(resp)
	}
	return HTTPRetryPatch504Response{RawResponse: resp}, nil
}

// patch504CreateRequest creates the Patch504 request.
func (client *HTTPRetryClient) patch504CreateRequest(ctx context.Context, options *HTTPRetryPatch504Options) (*policy.Request, error) {
	urlPath := "/http/retry/504"
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, true)
}

// patch504HandleError handles the Patch504 error response.
func (client *HTTPRetryClient) patch504HandleError(resp *http.Response) error {
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

// Post503 - Return 503 status code, then 200 after retry
// If the operation fails it returns the *Error error type.
func (client *HTTPRetryClient) Post503(ctx context.Context, options *HTTPRetryPost503Options) (HTTPRetryPost503Response, error) {
	req, err := client.post503CreateRequest(ctx, options)
	if err != nil {
		return HTTPRetryPost503Response{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return HTTPRetryPost503Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HTTPRetryPost503Response{}, client.post503HandleError(resp)
	}
	return HTTPRetryPost503Response{RawResponse: resp}, nil
}

// post503CreateRequest creates the Post503 request.
func (client *HTTPRetryClient) post503CreateRequest(ctx context.Context, options *HTTPRetryPost503Options) (*policy.Request, error) {
	urlPath := "/http/retry/503"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, true)
}

// post503HandleError handles the Post503 error response.
func (client *HTTPRetryClient) post503HandleError(resp *http.Response) error {
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

// Put500 - Return 500 status code, then 200 after retry
// If the operation fails it returns the *Error error type.
func (client *HTTPRetryClient) Put500(ctx context.Context, options *HTTPRetryPut500Options) (HTTPRetryPut500Response, error) {
	req, err := client.put500CreateRequest(ctx, options)
	if err != nil {
		return HTTPRetryPut500Response{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return HTTPRetryPut500Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HTTPRetryPut500Response{}, client.put500HandleError(resp)
	}
	return HTTPRetryPut500Response{RawResponse: resp}, nil
}

// put500CreateRequest creates the Put500 request.
func (client *HTTPRetryClient) put500CreateRequest(ctx context.Context, options *HTTPRetryPut500Options) (*policy.Request, error) {
	urlPath := "/http/retry/500"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, true)
}

// put500HandleError handles the Put500 error response.
func (client *HTTPRetryClient) put500HandleError(resp *http.Response) error {
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

// Put504 - Return 504 status code, then 200 after retry
// If the operation fails it returns the *Error error type.
func (client *HTTPRetryClient) Put504(ctx context.Context, options *HTTPRetryPut504Options) (HTTPRetryPut504Response, error) {
	req, err := client.put504CreateRequest(ctx, options)
	if err != nil {
		return HTTPRetryPut504Response{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return HTTPRetryPut504Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HTTPRetryPut504Response{}, client.put504HandleError(resp)
	}
	return HTTPRetryPut504Response{RawResponse: resp}, nil
}

// put504CreateRequest creates the Put504 request.
func (client *HTTPRetryClient) put504CreateRequest(ctx context.Context, options *HTTPRetryPut504Options) (*policy.Request, error) {
	urlPath := "/http/retry/504"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, true)
}

// put504HandleError handles the Put504 error response.
func (client *HTTPRetryClient) put504HandleError(resp *http.Response) error {
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
