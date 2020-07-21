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

// NetworkInterfaceIPConfigurationsOperations contains the methods for the NetworkInterfaceIPConfigurations group.
type NetworkInterfaceIPConfigurationsOperations interface {
	// Get - Gets the specified network interface ip configuration.
	Get(ctx context.Context, resourceGroupName string, networkInterfaceName string, ipConfigurationName string) (*NetworkInterfaceIPConfigurationResponse, error)
	// List - Get all ip configurations in a network interface.
	List(resourceGroupName string, networkInterfaceName string) (NetworkInterfaceIPConfigurationListResultPager, error)
}

// networkInterfaceIPConfigurationsOperations implements the NetworkInterfaceIPConfigurationsOperations interface.
type networkInterfaceIPConfigurationsOperations struct {
	*Client
	subscriptionID string
}

// Get - Gets the specified network interface ip configuration.
func (client *networkInterfaceIPConfigurationsOperations) Get(ctx context.Context, resourceGroupName string, networkInterfaceName string, ipConfigurationName string) (*NetworkInterfaceIPConfigurationResponse, error) {
	req, err := client.getCreateRequest(resourceGroupName, networkInterfaceName, ipConfigurationName)
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
func (client *networkInterfaceIPConfigurationsOperations) getCreateRequest(resourceGroupName string, networkInterfaceName string, ipConfigurationName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}/ipConfigurations/{ipConfigurationName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{networkInterfaceName}", url.PathEscape(networkInterfaceName))
	urlPath = strings.ReplaceAll(urlPath, "{ipConfigurationName}", url.PathEscape(ipConfigurationName))
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
func (client *networkInterfaceIPConfigurationsOperations) getHandleResponse(resp *azcore.Response) (*NetworkInterfaceIPConfigurationResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getHandleError(resp)
	}
	result := NetworkInterfaceIPConfigurationResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.NetworkInterfaceIPConfiguration)
}

// getHandleError handles the Get error response.
func (client *networkInterfaceIPConfigurationsOperations) getHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// List - Get all ip configurations in a network interface.
func (client *networkInterfaceIPConfigurationsOperations) List(resourceGroupName string, networkInterfaceName string) (NetworkInterfaceIPConfigurationListResultPager, error) {
	req, err := client.listCreateRequest(resourceGroupName, networkInterfaceName)
	if err != nil {
		return nil, err
	}
	return &networkInterfaceIPConfigurationListResultPager{
		pipeline:  client.p,
		request:   req,
		responder: client.listHandleResponse,
		advancer: func(resp *NetworkInterfaceIPConfigurationListResultResponse) (*azcore.Request, error) {
			u, err := url.Parse(*resp.NetworkInterfaceIPConfigurationListResult.NextLink)
			if err != nil {
				return nil, fmt.Errorf("invalid NextLink: %w", err)
			}
			if u.Scheme == "" {
				return nil, fmt.Errorf("no scheme detected in NextLink %s", *resp.NetworkInterfaceIPConfigurationListResult.NextLink)
			}
			return azcore.NewRequest(http.MethodGet, *u), nil
		},
	}, nil
}

// listCreateRequest creates the List request.
func (client *networkInterfaceIPConfigurationsOperations) listCreateRequest(resourceGroupName string, networkInterfaceName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}/ipConfigurations"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{networkInterfaceName}", url.PathEscape(networkInterfaceName))
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
func (client *networkInterfaceIPConfigurationsOperations) listHandleResponse(resp *azcore.Response) (*NetworkInterfaceIPConfigurationListResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.listHandleError(resp)
	}
	result := NetworkInterfaceIPConfigurationListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.NetworkInterfaceIPConfigurationListResult)
}

// listHandleError handles the List error response.
func (client *networkInterfaceIPConfigurationsOperations) listHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
