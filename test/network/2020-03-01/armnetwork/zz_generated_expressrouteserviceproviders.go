// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// ExpressRouteServiceProvidersOperations contains the methods for the ExpressRouteServiceProviders group.
type ExpressRouteServiceProvidersOperations interface {
	// List - Gets all the available express route service providers.
	List() ExpressRouteServiceProviderListResultPager
}

// ExpressRouteServiceProvidersClient implements the ExpressRouteServiceProvidersOperations interface.
// Don't use this type directly, use NewExpressRouteServiceProvidersClient() instead.
type ExpressRouteServiceProvidersClient struct {
	*Client
	subscriptionID string
}

// NewExpressRouteServiceProvidersClient creates a new instance of ExpressRouteServiceProvidersClient with the specified values.
func NewExpressRouteServiceProvidersClient(c *Client, subscriptionID string) ExpressRouteServiceProvidersOperations {
	return &ExpressRouteServiceProvidersClient{Client: c, subscriptionID: subscriptionID}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *ExpressRouteServiceProvidersClient) Do(req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(req)
}

// List - Gets all the available express route service providers.
func (client *ExpressRouteServiceProvidersClient) List() ExpressRouteServiceProviderListResultPager {
	return &expressRouteServiceProviderListResultPager{
		pipeline: client.p,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListCreateRequest(ctx)
		},
		responder: client.ListHandleResponse,
		advancer: func(ctx context.Context, resp *ExpressRouteServiceProviderListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ExpressRouteServiceProviderListResult.NextLink)
		},
	}
}

// ListCreateRequest creates the List request.
func (client *ExpressRouteServiceProvidersClient) ListCreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/expressRouteServiceProviders"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// ListHandleResponse handles the List response.
func (client *ExpressRouteServiceProvidersClient) ListHandleResponse(resp *azcore.Response) (*ExpressRouteServiceProviderListResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.ListHandleError(resp)
	}
	result := ExpressRouteServiceProviderListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ExpressRouteServiceProviderListResult)
}

// ListHandleError handles the List error response.
func (client *ExpressRouteServiceProvidersClient) ListHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}