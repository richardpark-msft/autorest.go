// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"path"
	"strings"
)

// AvailableEndpointServicesOperations contains the methods for the AvailableEndpointServices group.
type AvailableEndpointServicesOperations interface {
	// List - List what values of endpoint services are available for use.
	List(location string) (EndpointServicesListResultPager, error)
}

// availableEndpointServicesOperations implements the AvailableEndpointServicesOperations interface.
type availableEndpointServicesOperations struct {
	*Client
	subscriptionID string
}

// List - List what values of endpoint services are available for use.
func (client *availableEndpointServicesOperations) List(location string) (EndpointServicesListResultPager, error) {
	req, err := client.listCreateRequest(location)
	if err != nil {
		return nil, err
	}
	return &endpointServicesListResultPager{
		pipeline:  client.p,
		request:   req,
		responder: client.listHandleResponse,
		advancer: func(resp *EndpointServicesListResultResponse) (*azcore.Request, error) {
			u, err := url.Parse(*resp.EndpointServicesListResult.NextLink)
			if err != nil {
				return nil, fmt.Errorf("invalid NextLink: %w", err)
			}
			if u.Scheme == "" {
				return nil, fmt.Errorf("no scheme detected in NextLink %s", *resp.EndpointServicesListResult.NextLink)
			}
			return azcore.NewRequest(http.MethodGet, *u), nil
		},
	}, nil
}

// listCreateRequest creates the List request.
func (client *availableEndpointServicesOperations) listCreateRequest(location string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/locations/{location}/virtualNetworkAvailableEndpointServices"
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(path.Join(client.u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// listHandleResponse handles the List response.
func (client *availableEndpointServicesOperations) listHandleResponse(resp *azcore.Response) (*EndpointServicesListResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.listHandleError(resp)
	}
	result := EndpointServicesListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.EndpointServicesListResult)
}

// listHandleError handles the List error response.
func (client *availableEndpointServicesOperations) listHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
