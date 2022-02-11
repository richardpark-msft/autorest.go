//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// AvailableDelegationsClient contains the methods for the AvailableDelegations group.
// Don't use this type directly, use NewAvailableDelegationsClient() instead.
type AvailableDelegationsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAvailableDelegationsClient creates a new instance of AvailableDelegationsClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAvailableDelegationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *AvailableDelegationsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &AvailableDelegationsClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// List - Gets all of the available subnet delegations for this subscription in this region.
// If the operation fails it returns an *azcore.ResponseError type.
// location - The location of the subnet.
// options - AvailableDelegationsClientListOptions contains the optional parameters for the AvailableDelegationsClient.List
// method.
func (client *AvailableDelegationsClient) List(location string, options *AvailableDelegationsClientListOptions) *AvailableDelegationsClientListPager {
	return &AvailableDelegationsClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, location, options)
		},
		advancer: func(ctx context.Context, resp AvailableDelegationsClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.AvailableDelegationsResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *AvailableDelegationsClient) listCreateRequest(ctx context.Context, location string, options *AvailableDelegationsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/locations/{location}/availableDelegations"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AvailableDelegationsClient) listHandleResponse(resp *http.Response) (AvailableDelegationsClientListResponse, error) {
	result := AvailableDelegationsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AvailableDelegationsResult); err != nil {
		return AvailableDelegationsClientListResponse{}, err
	}
	return result, nil
}
