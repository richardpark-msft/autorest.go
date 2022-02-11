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

// VPNConnectionsClient contains the methods for the VPNConnections group.
// Don't use this type directly, use NewVPNConnectionsClient() instead.
type VPNConnectionsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewVPNConnectionsClient creates a new instance of VPNConnectionsClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewVPNConnectionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *VPNConnectionsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &VPNConnectionsClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// BeginCreateOrUpdate - Creates a vpn connection to a scalable vpn gateway if it doesn't exist else updates the existing
// connection.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name of the VpnGateway.
// gatewayName - The name of the gateway.
// connectionName - The name of the connection.
// vpnConnectionParameters - Parameters supplied to create or Update a VPN Connection.
// options - VPNConnectionsClientBeginCreateOrUpdateOptions contains the optional parameters for the VPNConnectionsClient.BeginCreateOrUpdate
// method.
func (client *VPNConnectionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, vpnConnectionParameters VPNConnection, options *VPNConnectionsClientBeginCreateOrUpdateOptions) (VPNConnectionsClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, gatewayName, connectionName, vpnConnectionParameters, options)
	if err != nil {
		return VPNConnectionsClientCreateOrUpdatePollerResponse{}, err
	}
	result := VPNConnectionsClientCreateOrUpdatePollerResponse{}
	pt, err := armruntime.NewPoller("VPNConnectionsClient.CreateOrUpdate", "azure-async-operation", resp, client.pl)
	if err != nil {
		return VPNConnectionsClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &VPNConnectionsClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates a vpn connection to a scalable vpn gateway if it doesn't exist else updates the existing connection.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *VPNConnectionsClient) createOrUpdate(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, vpnConnectionParameters VPNConnection, options *VPNConnectionsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, gatewayName, connectionName, vpnConnectionParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *VPNConnectionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, vpnConnectionParameters VPNConnection, options *VPNConnectionsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/vpnConnections/{connectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, vpnConnectionParameters)
}

// BeginDelete - Deletes a vpn connection.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name of the VpnGateway.
// gatewayName - The name of the gateway.
// connectionName - The name of the connection.
// options - VPNConnectionsClientBeginDeleteOptions contains the optional parameters for the VPNConnectionsClient.BeginDelete
// method.
func (client *VPNConnectionsClient) BeginDelete(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, options *VPNConnectionsClientBeginDeleteOptions) (VPNConnectionsClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, gatewayName, connectionName, options)
	if err != nil {
		return VPNConnectionsClientDeletePollerResponse{}, err
	}
	result := VPNConnectionsClientDeletePollerResponse{}
	pt, err := armruntime.NewPoller("VPNConnectionsClient.Delete", "location", resp, client.pl)
	if err != nil {
		return VPNConnectionsClientDeletePollerResponse{}, err
	}
	result.Poller = &VPNConnectionsClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes a vpn connection.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *VPNConnectionsClient) deleteOperation(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, options *VPNConnectionsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, gatewayName, connectionName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *VPNConnectionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, options *VPNConnectionsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/vpnConnections/{connectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Retrieves the details of a vpn connection.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name of the VpnGateway.
// gatewayName - The name of the gateway.
// connectionName - The name of the vpn connection.
// options - VPNConnectionsClientGetOptions contains the optional parameters for the VPNConnectionsClient.Get method.
func (client *VPNConnectionsClient) Get(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, options *VPNConnectionsClientGetOptions) (VPNConnectionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, gatewayName, connectionName, options)
	if err != nil {
		return VPNConnectionsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VPNConnectionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VPNConnectionsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VPNConnectionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, options *VPNConnectionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/vpnConnections/{connectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
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
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VPNConnectionsClient) getHandleResponse(resp *http.Response) (VPNConnectionsClientGetResponse, error) {
	result := VPNConnectionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VPNConnection); err != nil {
		return VPNConnectionsClientGetResponse{}, err
	}
	return result, nil
}

// ListByVPNGateway - Retrieves all vpn connections for a particular virtual wan vpn gateway.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name of the VpnGateway.
// gatewayName - The name of the gateway.
// options - VPNConnectionsClientListByVPNGatewayOptions contains the optional parameters for the VPNConnectionsClient.ListByVPNGateway
// method.
func (client *VPNConnectionsClient) ListByVPNGateway(resourceGroupName string, gatewayName string, options *VPNConnectionsClientListByVPNGatewayOptions) *VPNConnectionsClientListByVPNGatewayPager {
	return &VPNConnectionsClientListByVPNGatewayPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByVPNGatewayCreateRequest(ctx, resourceGroupName, gatewayName, options)
		},
		advancer: func(ctx context.Context, resp VPNConnectionsClientListByVPNGatewayResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ListVPNConnectionsResult.NextLink)
		},
	}
}

// listByVPNGatewayCreateRequest creates the ListByVPNGateway request.
func (client *VPNConnectionsClient) listByVPNGatewayCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, options *VPNConnectionsClientListByVPNGatewayOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/vpnConnections"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
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

// listByVPNGatewayHandleResponse handles the ListByVPNGateway response.
func (client *VPNConnectionsClient) listByVPNGatewayHandleResponse(resp *http.Response) (VPNConnectionsClientListByVPNGatewayResponse, error) {
	result := VPNConnectionsClientListByVPNGatewayResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListVPNConnectionsResult); err != nil {
		return VPNConnectionsClientListByVPNGatewayResponse{}, err
	}
	return result, nil
}
