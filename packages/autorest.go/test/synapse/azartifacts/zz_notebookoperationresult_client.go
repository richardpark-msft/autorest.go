//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azartifacts

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// NotebookOperationResultClient contains the methods for the NotebookOperationResult group.
// Don't use this type directly, use a constructor function instead.
type NotebookOperationResultClient struct {
	internal *azcore.Client
	endpoint string
}

// Get - Get notebook operation result
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-12-01
//   - operationID - Operation ID.
//   - options - NotebookOperationResultClientGetOptions contains the optional parameters for the NotebookOperationResultClient.Get
//     method.
func (client *NotebookOperationResultClient) Get(ctx context.Context, operationID string, options *NotebookOperationResultClientGetOptions) (NotebookOperationResultClientGetResponse, error) {
	var err error
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "NotebookOperationResultClient.Get")
	req, err := client.getCreateRequest(ctx, operationID, options)
	if err != nil {
		return NotebookOperationResultClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NotebookOperationResultClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return NotebookOperationResultClientGetResponse{}, err
	}
	return NotebookOperationResultClientGetResponse{}, nil
}

// getCreateRequest creates the Get request.
func (client *NotebookOperationResultClient) getCreateRequest(ctx context.Context, operationID string, options *NotebookOperationResultClientGetOptions) (*policy.Request, error) {
	urlPath := "/notebookOperationResults/{operationId}"
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}
