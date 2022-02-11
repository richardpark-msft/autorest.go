//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

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

// DedicatedHostsClient contains the methods for the DedicatedHosts group.
// Don't use this type directly, use NewDedicatedHostsClient() instead.
type DedicatedHostsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewDedicatedHostsClient creates a new instance of DedicatedHostsClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewDedicatedHostsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *DedicatedHostsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &DedicatedHostsClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// BeginCreateOrUpdate - Create or update a dedicated host .
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// hostGroupName - The name of the dedicated host group.
// hostName - The name of the dedicated host .
// parameters - Parameters supplied to the Create Dedicated Host.
// options - DedicatedHostsClientBeginCreateOrUpdateOptions contains the optional parameters for the DedicatedHostsClient.BeginCreateOrUpdate
// method.
func (client *DedicatedHostsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, hostGroupName string, hostName string, parameters DedicatedHost, options *DedicatedHostsClientBeginCreateOrUpdateOptions) (DedicatedHostsClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, hostGroupName, hostName, parameters, options)
	if err != nil {
		return DedicatedHostsClientCreateOrUpdatePollerResponse{}, err
	}
	result := DedicatedHostsClientCreateOrUpdatePollerResponse{}
	pt, err := armruntime.NewPoller("DedicatedHostsClient.CreateOrUpdate", "", resp, client.pl)
	if err != nil {
		return DedicatedHostsClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &DedicatedHostsClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Create or update a dedicated host .
// If the operation fails it returns an *azcore.ResponseError type.
func (client *DedicatedHostsClient) createOrUpdate(ctx context.Context, resourceGroupName string, hostGroupName string, hostName string, parameters DedicatedHost, options *DedicatedHostsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, hostGroupName, hostName, parameters, options)
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
func (client *DedicatedHostsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, hostGroupName string, hostName string, parameters DedicatedHost, options *DedicatedHostsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/hostGroups/{hostGroupName}/hosts/{hostName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hostGroupName == "" {
		return nil, errors.New("parameter hostGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostGroupName}", url.PathEscape(hostGroupName))
	if hostName == "" {
		return nil, errors.New("parameter hostName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostName}", url.PathEscape(hostName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Delete a dedicated host.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// hostGroupName - The name of the dedicated host group.
// hostName - The name of the dedicated host.
// options - DedicatedHostsClientBeginDeleteOptions contains the optional parameters for the DedicatedHostsClient.BeginDelete
// method.
func (client *DedicatedHostsClient) BeginDelete(ctx context.Context, resourceGroupName string, hostGroupName string, hostName string, options *DedicatedHostsClientBeginDeleteOptions) (DedicatedHostsClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, hostGroupName, hostName, options)
	if err != nil {
		return DedicatedHostsClientDeletePollerResponse{}, err
	}
	result := DedicatedHostsClientDeletePollerResponse{}
	pt, err := armruntime.NewPoller("DedicatedHostsClient.Delete", "", resp, client.pl)
	if err != nil {
		return DedicatedHostsClientDeletePollerResponse{}, err
	}
	result.Poller = &DedicatedHostsClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Delete a dedicated host.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *DedicatedHostsClient) deleteOperation(ctx context.Context, resourceGroupName string, hostGroupName string, hostName string, options *DedicatedHostsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, hostGroupName, hostName, options)
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
func (client *DedicatedHostsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, hostGroupName string, hostName string, options *DedicatedHostsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/hostGroups/{hostGroupName}/hosts/{hostName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hostGroupName == "" {
		return nil, errors.New("parameter hostGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostGroupName}", url.PathEscape(hostGroupName))
	if hostName == "" {
		return nil, errors.New("parameter hostName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostName}", url.PathEscape(hostName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Retrieves information about a dedicated host.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// hostGroupName - The name of the dedicated host group.
// hostName - The name of the dedicated host.
// options - DedicatedHostsClientGetOptions contains the optional parameters for the DedicatedHostsClient.Get method.
func (client *DedicatedHostsClient) Get(ctx context.Context, resourceGroupName string, hostGroupName string, hostName string, options *DedicatedHostsClientGetOptions) (DedicatedHostsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, hostGroupName, hostName, options)
	if err != nil {
		return DedicatedHostsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DedicatedHostsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DedicatedHostsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DedicatedHostsClient) getCreateRequest(ctx context.Context, resourceGroupName string, hostGroupName string, hostName string, options *DedicatedHostsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/hostGroups/{hostGroupName}/hosts/{hostName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hostGroupName == "" {
		return nil, errors.New("parameter hostGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostGroupName}", url.PathEscape(hostGroupName))
	if hostName == "" {
		return nil, errors.New("parameter hostName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostName}", url.PathEscape(hostName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", "instanceView")
	}
	reqQP.Set("api-version", "2019-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DedicatedHostsClient) getHandleResponse(resp *http.Response) (DedicatedHostsClientGetResponse, error) {
	result := DedicatedHostsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DedicatedHost); err != nil {
		return DedicatedHostsClientGetResponse{}, err
	}
	return result, nil
}

// ListByHostGroup - Lists all of the dedicated hosts in the specified dedicated host group. Use the nextLink property in
// the response to get the next page of dedicated hosts.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// hostGroupName - The name of the dedicated host group.
// options - DedicatedHostsClientListByHostGroupOptions contains the optional parameters for the DedicatedHostsClient.ListByHostGroup
// method.
func (client *DedicatedHostsClient) ListByHostGroup(resourceGroupName string, hostGroupName string, options *DedicatedHostsClientListByHostGroupOptions) *DedicatedHostsClientListByHostGroupPager {
	return &DedicatedHostsClientListByHostGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByHostGroupCreateRequest(ctx, resourceGroupName, hostGroupName, options)
		},
		advancer: func(ctx context.Context, resp DedicatedHostsClientListByHostGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.DedicatedHostListResult.NextLink)
		},
	}
}

// listByHostGroupCreateRequest creates the ListByHostGroup request.
func (client *DedicatedHostsClient) listByHostGroupCreateRequest(ctx context.Context, resourceGroupName string, hostGroupName string, options *DedicatedHostsClientListByHostGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/hostGroups/{hostGroupName}/hosts"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hostGroupName == "" {
		return nil, errors.New("parameter hostGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostGroupName}", url.PathEscape(hostGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByHostGroupHandleResponse handles the ListByHostGroup response.
func (client *DedicatedHostsClient) listByHostGroupHandleResponse(resp *http.Response) (DedicatedHostsClientListByHostGroupResponse, error) {
	result := DedicatedHostsClientListByHostGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DedicatedHostListResult); err != nil {
		return DedicatedHostsClientListByHostGroupResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Update an dedicated host .
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// hostGroupName - The name of the dedicated host group.
// hostName - The name of the dedicated host .
// parameters - Parameters supplied to the Update Dedicated Host operation.
// options - DedicatedHostsClientBeginUpdateOptions contains the optional parameters for the DedicatedHostsClient.BeginUpdate
// method.
func (client *DedicatedHostsClient) BeginUpdate(ctx context.Context, resourceGroupName string, hostGroupName string, hostName string, parameters DedicatedHostUpdate, options *DedicatedHostsClientBeginUpdateOptions) (DedicatedHostsClientUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, hostGroupName, hostName, parameters, options)
	if err != nil {
		return DedicatedHostsClientUpdatePollerResponse{}, err
	}
	result := DedicatedHostsClientUpdatePollerResponse{}
	pt, err := armruntime.NewPoller("DedicatedHostsClient.Update", "", resp, client.pl)
	if err != nil {
		return DedicatedHostsClientUpdatePollerResponse{}, err
	}
	result.Poller = &DedicatedHostsClientUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - Update an dedicated host .
// If the operation fails it returns an *azcore.ResponseError type.
func (client *DedicatedHostsClient) update(ctx context.Context, resourceGroupName string, hostGroupName string, hostName string, parameters DedicatedHostUpdate, options *DedicatedHostsClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, hostGroupName, hostName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *DedicatedHostsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, hostGroupName string, hostName string, parameters DedicatedHostUpdate, options *DedicatedHostsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/hostGroups/{hostGroupName}/hosts/{hostName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hostGroupName == "" {
		return nil, errors.New("parameter hostGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostGroupName}", url.PathEscape(hostGroupName))
	if hostName == "" {
		return nil, errors.New("parameter hostName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostName}", url.PathEscape(hostName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}
