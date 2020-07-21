// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azurespecialsgroup

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"io/ioutil"
	"net/http"
	"path"
)

// XMSClientRequestIDOperations contains the methods for the XMSClientRequestID group.
type XMSClientRequestIDOperations interface {
	// Get - Get method that overwrites x-ms-client-request header with value 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0.
	Get(ctx context.Context) (*http.Response, error)
	// ParamGet - Get method that overwrites x-ms-client-request header with value 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0.
	ParamGet(ctx context.Context, xMSClientRequestId string) (*http.Response, error)
}

// xmsClientRequestIdoperations implements the XMSClientRequestIDOperations interface.
type xmsClientRequestIdoperations struct {
	*Client
}

// Get - Get method that overwrites x-ms-client-request header with value 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0.
func (client *xmsClientRequestIdoperations) Get(ctx context.Context) (*http.Response, error) {
	req, err := client.getCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getCreateRequest creates the Get request.
func (client *xmsClientRequestIdoperations) getCreateRequest() (*azcore.Request, error) {
	urlPath := "/azurespecials/overwrite/x-ms-client-request-id/method/"
	u, err := client.u.Parse(path.Join(client.u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *xmsClientRequestIdoperations) getHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getHandleError(resp)
	}
	return resp.Response, nil
}

// getHandleError handles the Get error response.
func (client *xmsClientRequestIdoperations) getHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return errors.New(resp.Status)
	}
	return errors.New(string(body))
}

// ParamGet - Get method that overwrites x-ms-client-request header with value 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0.
func (client *xmsClientRequestIdoperations) ParamGet(ctx context.Context, xMSClientRequestId string) (*http.Response, error) {
	req, err := client.paramGetCreateRequest(xMSClientRequestId)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.paramGetHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// paramGetCreateRequest creates the ParamGet request.
func (client *xmsClientRequestIdoperations) paramGetCreateRequest(xMSClientRequestId string) (*azcore.Request, error) {
	urlPath := "/azurespecials/overwrite/x-ms-client-request-id/via-param/method/"
	u, err := client.u.Parse(path.Join(client.u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	req.Header.Set("x-ms-client-request-id", xMSClientRequestId)
	return req, nil
}

// paramGetHandleResponse handles the ParamGet response.
func (client *xmsClientRequestIdoperations) paramGetHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.paramGetHandleError(resp)
	}
	return resp.Response, nil
}

// paramGetHandleError handles the ParamGet error response.
func (client *xmsClientRequestIdoperations) paramGetHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
