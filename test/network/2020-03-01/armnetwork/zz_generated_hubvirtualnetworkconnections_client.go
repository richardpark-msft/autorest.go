//go:build go1.18
// +build go1.18

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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// HubVirtualNetworkConnectionsClient contains the methods for the HubVirtualNetworkConnections group.
// Don't use this type directly, use NewHubVirtualNetworkConnectionsClient() instead.
type HubVirtualNetworkConnectionsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewHubVirtualNetworkConnectionsClient creates a new instance of HubVirtualNetworkConnectionsClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewHubVirtualNetworkConnectionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*HubVirtualNetworkConnectionsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublicCloud.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &HubVirtualNetworkConnectionsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Get - Retrieves the details of a HubVirtualNetworkConnection.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-03-01
// resourceGroupName - The resource group name of the VirtualHub.
// virtualHubName - The name of the VirtualHub.
// connectionName - The name of the vpn connection.
// options - HubVirtualNetworkConnectionsClientGetOptions contains the optional parameters for the HubVirtualNetworkConnectionsClient.Get
// method.
func (client *HubVirtualNetworkConnectionsClient) Get(ctx context.Context, resourceGroupName string, virtualHubName string, connectionName string, options *HubVirtualNetworkConnectionsClientGetOptions) (HubVirtualNetworkConnectionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, virtualHubName, connectionName, options)
	if err != nil {
		return HubVirtualNetworkConnectionsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HubVirtualNetworkConnectionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HubVirtualNetworkConnectionsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *HubVirtualNetworkConnectionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, virtualHubName string, connectionName string, options *HubVirtualNetworkConnectionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/hubVirtualNetworkConnections/{connectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualHubName == "" {
		return nil, errors.New("parameter virtualHubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualHubName}", url.PathEscape(virtualHubName))
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *HubVirtualNetworkConnectionsClient) getHandleResponse(resp *http.Response) (HubVirtualNetworkConnectionsClientGetResponse, error) {
	result := HubVirtualNetworkConnectionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HubVirtualNetworkConnection); err != nil {
		return HubVirtualNetworkConnectionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Retrieves the details of all HubVirtualNetworkConnections.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-03-01
// resourceGroupName - The resource group name of the VirtualHub.
// virtualHubName - The name of the VirtualHub.
// options - HubVirtualNetworkConnectionsClientListOptions contains the optional parameters for the HubVirtualNetworkConnectionsClient.List
// method.
func (client *HubVirtualNetworkConnectionsClient) NewListPager(resourceGroupName string, virtualHubName string, options *HubVirtualNetworkConnectionsClientListOptions) *runtime.Pager[HubVirtualNetworkConnectionsClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[HubVirtualNetworkConnectionsClientListResponse]{
		More: func(page HubVirtualNetworkConnectionsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *HubVirtualNetworkConnectionsClientListResponse) (HubVirtualNetworkConnectionsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, virtualHubName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return HubVirtualNetworkConnectionsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return HubVirtualNetworkConnectionsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return HubVirtualNetworkConnectionsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *HubVirtualNetworkConnectionsClient) listCreateRequest(ctx context.Context, resourceGroupName string, virtualHubName string, options *HubVirtualNetworkConnectionsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/hubVirtualNetworkConnections"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualHubName == "" {
		return nil, errors.New("parameter virtualHubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualHubName}", url.PathEscape(virtualHubName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *HubVirtualNetworkConnectionsClient) listHandleResponse(resp *http.Response) (HubVirtualNetworkConnectionsClientListResponse, error) {
	result := HubVirtualNetworkConnectionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListHubVirtualNetworkConnectionsResult); err != nil {
		return HubVirtualNetworkConnectionsClientListResponse{}, err
	}
	return result, nil
}
