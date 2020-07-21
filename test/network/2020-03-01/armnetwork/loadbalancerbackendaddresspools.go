// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"path"
	"strings"
)

// LoadBalancerBackendAddressPoolsOperations contains the methods for the LoadBalancerBackendAddressPools group.
type LoadBalancerBackendAddressPoolsOperations interface {
	// Get - Gets load balancer backend address pool.
	Get(ctx context.Context, resourceGroupName string, loadBalancerName string, backendAddressPoolName string) (*BackendAddressPoolResponse, error)
	// List - Gets all the load balancer backed address pools.
	List(resourceGroupName string, loadBalancerName string) (LoadBalancerBackendAddressPoolListResultPager, error)
}

// loadBalancerBackendAddressPoolsOperations implements the LoadBalancerBackendAddressPoolsOperations interface.
type loadBalancerBackendAddressPoolsOperations struct {
	*Client
	subscriptionID string
}

// Get - Gets load balancer backend address pool.
func (client *loadBalancerBackendAddressPoolsOperations) Get(ctx context.Context, resourceGroupName string, loadBalancerName string, backendAddressPoolName string) (*BackendAddressPoolResponse, error) {
	req, err := client.getCreateRequest(resourceGroupName, loadBalancerName, backendAddressPoolName)
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
func (client *loadBalancerBackendAddressPoolsOperations) getCreateRequest(resourceGroupName string, loadBalancerName string, backendAddressPoolName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/backendAddressPools/{backendAddressPoolName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancerName}", url.PathEscape(loadBalancerName))
	urlPath = strings.ReplaceAll(urlPath, "{backendAddressPoolName}", url.PathEscape(backendAddressPoolName))
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

// getHandleResponse handles the Get response.
func (client *loadBalancerBackendAddressPoolsOperations) getHandleResponse(resp *azcore.Response) (*BackendAddressPoolResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getHandleError(resp)
	}
	result := BackendAddressPoolResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.BackendAddressPool)
}

// getHandleError handles the Get error response.
func (client *loadBalancerBackendAddressPoolsOperations) getHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// List - Gets all the load balancer backed address pools.
func (client *loadBalancerBackendAddressPoolsOperations) List(resourceGroupName string, loadBalancerName string) (LoadBalancerBackendAddressPoolListResultPager, error) {
	req, err := client.listCreateRequest(resourceGroupName, loadBalancerName)
	if err != nil {
		return nil, err
	}
	return &loadBalancerBackendAddressPoolListResultPager{
		pipeline:  client.p,
		request:   req,
		responder: client.listHandleResponse,
		advancer: func(resp *LoadBalancerBackendAddressPoolListResultResponse) (*azcore.Request, error) {
			u, err := url.Parse(*resp.LoadBalancerBackendAddressPoolListResult.NextLink)
			if err != nil {
				return nil, fmt.Errorf("invalid NextLink: %w", err)
			}
			if u.Scheme == "" {
				return nil, fmt.Errorf("no scheme detected in NextLink %s", *resp.LoadBalancerBackendAddressPoolListResult.NextLink)
			}
			return azcore.NewRequest(http.MethodGet, *u), nil
		},
	}, nil
}

// listCreateRequest creates the List request.
func (client *loadBalancerBackendAddressPoolsOperations) listCreateRequest(resourceGroupName string, loadBalancerName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/backendAddressPools"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancerName}", url.PathEscape(loadBalancerName))
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
func (client *loadBalancerBackendAddressPoolsOperations) listHandleResponse(resp *azcore.Response) (*LoadBalancerBackendAddressPoolListResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.listHandleError(resp)
	}
	result := LoadBalancerBackendAddressPoolListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.LoadBalancerBackendAddressPoolListResult)
}

// listHandleError handles the List error response.
func (client *loadBalancerBackendAddressPoolsOperations) listHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
