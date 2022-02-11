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

// VirtualAppliancesClient contains the methods for the NetworkVirtualAppliances group.
// Don't use this type directly, use NewVirtualAppliancesClient() instead.
type VirtualAppliancesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewVirtualAppliancesClient creates a new instance of VirtualAppliancesClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewVirtualAppliancesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *VirtualAppliancesClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &VirtualAppliancesClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// BeginCreateOrUpdate - Creates or updates the specified Network Virtual Appliance.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// networkVirtualApplianceName - The name of Network Virtual Appliance.
// parameters - Parameters supplied to the create or update Network Virtual Appliance.
// options - VirtualAppliancesClientBeginCreateOrUpdateOptions contains the optional parameters for the VirtualAppliancesClient.BeginCreateOrUpdate
// method.
func (client *VirtualAppliancesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, parameters VirtualAppliance, options *VirtualAppliancesClientBeginCreateOrUpdateOptions) (VirtualAppliancesClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, networkVirtualApplianceName, parameters, options)
	if err != nil {
		return VirtualAppliancesClientCreateOrUpdatePollerResponse{}, err
	}
	result := VirtualAppliancesClientCreateOrUpdatePollerResponse{}
	pt, err := armruntime.NewPoller("VirtualAppliancesClient.CreateOrUpdate", "azure-async-operation", resp, client.pl)
	if err != nil {
		return VirtualAppliancesClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &VirtualAppliancesClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates the specified Network Virtual Appliance.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *VirtualAppliancesClient) createOrUpdate(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, parameters VirtualAppliance, options *VirtualAppliancesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, networkVirtualApplianceName, parameters, options)
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
func (client *VirtualAppliancesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, parameters VirtualAppliance, options *VirtualAppliancesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkVirtualAppliances/{networkVirtualApplianceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkVirtualApplianceName == "" {
		return nil, errors.New("parameter networkVirtualApplianceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkVirtualApplianceName}", url.PathEscape(networkVirtualApplianceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes the specified Network Virtual Appliance.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// networkVirtualApplianceName - The name of Network Virtual Appliance.
// options - VirtualAppliancesClientBeginDeleteOptions contains the optional parameters for the VirtualAppliancesClient.BeginDelete
// method.
func (client *VirtualAppliancesClient) BeginDelete(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, options *VirtualAppliancesClientBeginDeleteOptions) (VirtualAppliancesClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, networkVirtualApplianceName, options)
	if err != nil {
		return VirtualAppliancesClientDeletePollerResponse{}, err
	}
	result := VirtualAppliancesClientDeletePollerResponse{}
	pt, err := armruntime.NewPoller("VirtualAppliancesClient.Delete", "location", resp, client.pl)
	if err != nil {
		return VirtualAppliancesClientDeletePollerResponse{}, err
	}
	result.Poller = &VirtualAppliancesClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes the specified Network Virtual Appliance.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *VirtualAppliancesClient) deleteOperation(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, options *VirtualAppliancesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, networkVirtualApplianceName, options)
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
func (client *VirtualAppliancesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, options *VirtualAppliancesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkVirtualAppliances/{networkVirtualApplianceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkVirtualApplianceName == "" {
		return nil, errors.New("parameter networkVirtualApplianceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkVirtualApplianceName}", url.PathEscape(networkVirtualApplianceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
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

// Get - Gets the specified Network Virtual Appliance.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// networkVirtualApplianceName - The name of Network Virtual Appliance.
// options - VirtualAppliancesClientGetOptions contains the optional parameters for the VirtualAppliancesClient.Get method.
func (client *VirtualAppliancesClient) Get(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, options *VirtualAppliancesClientGetOptions) (VirtualAppliancesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, networkVirtualApplianceName, options)
	if err != nil {
		return VirtualAppliancesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VirtualAppliancesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualAppliancesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VirtualAppliancesClient) getCreateRequest(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, options *VirtualAppliancesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkVirtualAppliances/{networkVirtualApplianceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkVirtualApplianceName == "" {
		return nil, errors.New("parameter networkVirtualApplianceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkVirtualApplianceName}", url.PathEscape(networkVirtualApplianceName))
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
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VirtualAppliancesClient) getHandleResponse(resp *http.Response) (VirtualAppliancesClientGetResponse, error) {
	result := VirtualAppliancesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualAppliance); err != nil {
		return VirtualAppliancesClientGetResponse{}, err
	}
	return result, nil
}

// List - Gets all Network Virtual Appliances in a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - VirtualAppliancesClientListOptions contains the optional parameters for the VirtualAppliancesClient.List method.
func (client *VirtualAppliancesClient) List(options *VirtualAppliancesClientListOptions) *VirtualAppliancesClientListPager {
	return &VirtualAppliancesClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp VirtualAppliancesClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.VirtualApplianceListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *VirtualAppliancesClient) listCreateRequest(ctx context.Context, options *VirtualAppliancesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/networkVirtualAppliances"
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
func (client *VirtualAppliancesClient) listHandleResponse(resp *http.Response) (VirtualAppliancesClientListResponse, error) {
	result := VirtualAppliancesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualApplianceListResult); err != nil {
		return VirtualAppliancesClientListResponse{}, err
	}
	return result, nil
}

// ListByResourceGroup - Lists all Network Virtual Appliances in a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// options - VirtualAppliancesClientListByResourceGroupOptions contains the optional parameters for the VirtualAppliancesClient.ListByResourceGroup
// method.
func (client *VirtualAppliancesClient) ListByResourceGroup(resourceGroupName string, options *VirtualAppliancesClientListByResourceGroupOptions) *VirtualAppliancesClientListByResourceGroupPager {
	return &VirtualAppliancesClientListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp VirtualAppliancesClientListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.VirtualApplianceListResult.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *VirtualAppliancesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *VirtualAppliancesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkVirtualAppliances"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
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

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *VirtualAppliancesClient) listByResourceGroupHandleResponse(resp *http.Response) (VirtualAppliancesClientListByResourceGroupResponse, error) {
	result := VirtualAppliancesClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualApplianceListResult); err != nil {
		return VirtualAppliancesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// UpdateTags - Updates a Network Virtual Appliance.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name of Network Virtual Appliance.
// networkVirtualApplianceName - The name of Network Virtual Appliance being updated.
// parameters - Parameters supplied to Update Network Virtual Appliance Tags.
// options - VirtualAppliancesClientUpdateTagsOptions contains the optional parameters for the VirtualAppliancesClient.UpdateTags
// method.
func (client *VirtualAppliancesClient) UpdateTags(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, parameters TagsObject, options *VirtualAppliancesClientUpdateTagsOptions) (VirtualAppliancesClientUpdateTagsResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, networkVirtualApplianceName, parameters, options)
	if err != nil {
		return VirtualAppliancesClientUpdateTagsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VirtualAppliancesClientUpdateTagsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualAppliancesClientUpdateTagsResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateTagsHandleResponse(resp)
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *VirtualAppliancesClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, parameters TagsObject, options *VirtualAppliancesClientUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkVirtualAppliances/{networkVirtualApplianceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkVirtualApplianceName == "" {
		return nil, errors.New("parameter networkVirtualApplianceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkVirtualApplianceName}", url.PathEscape(networkVirtualApplianceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *VirtualAppliancesClient) updateTagsHandleResponse(resp *http.Response) (VirtualAppliancesClientUpdateTagsResponse, error) {
	result := VirtualAppliancesClientUpdateTagsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualAppliance); err != nil {
		return VirtualAppliancesClientUpdateTagsResponse{}, err
	}
	return result, nil
}
