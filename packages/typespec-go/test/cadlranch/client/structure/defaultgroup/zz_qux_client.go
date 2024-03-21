// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package defaultgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
)

// QuxClient contains the methods for the Client.Structure.Service namespace.
// Don't use this type directly, use [ServiceClient.NewQuxClient] instead.
type QuxClient struct {
	internal *azcore.Client
	endpoint string
	client   ClientType
}

// NewBarClient creates a new instance of [BarClient].
func (client *QuxClient) NewBarClient() *BarClient {
	return &BarClient{
		internal: client.internal,
		endpoint: client.endpoint,
		client:   client.client,
	}
}

// - options - QuxClientEightOptions contains the optional parameters for the QuxClient.Eight method.
func (client *QuxClient) Eight(ctx context.Context, options *QuxClientEightOptions) (QuxClientEightResponse, error) {
	var err error
	const operationName = "QuxClient.Eight"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.eightCreateRequest(ctx, options)
	if err != nil {
		return QuxClientEightResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return QuxClientEightResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return QuxClientEightResponse{}, err
	}
	return QuxClientEightResponse{}, nil
}

// eightCreateRequest creates the Eight request.
func (client *QuxClient) eightCreateRequest(ctx context.Context, options *QuxClientEightOptions) (*policy.Request, error) {
	host := "{endpoint}/client/structure/{client}"
	host = strings.ReplaceAll(host, "{endpoint}", client.endpoint)
	host = strings.ReplaceAll(host, "{client}", string(client.client))
	urlPath := "/eight"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}
