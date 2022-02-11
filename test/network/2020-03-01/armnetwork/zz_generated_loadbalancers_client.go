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

// LoadBalancersClient contains the methods for the LoadBalancers group.
// Don't use this type directly, use NewLoadBalancersClient() instead.
type LoadBalancersClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewLoadBalancersClient creates a new instance of LoadBalancersClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewLoadBalancersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *LoadBalancersClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &LoadBalancersClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// BeginCreateOrUpdate - Creates or updates a load balancer.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// loadBalancerName - The name of the load balancer.
// parameters - Parameters supplied to the create or update load balancer operation.
// options - LoadBalancersClientBeginCreateOrUpdateOptions contains the optional parameters for the LoadBalancersClient.BeginCreateOrUpdate
// method.
func (client *LoadBalancersClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, loadBalancerName string, parameters LoadBalancer, options *LoadBalancersClientBeginCreateOrUpdateOptions) (LoadBalancersClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, loadBalancerName, parameters, options)
	if err != nil {
		return LoadBalancersClientCreateOrUpdatePollerResponse{}, err
	}
	result := LoadBalancersClientCreateOrUpdatePollerResponse{}
	pt, err := armruntime.NewPoller("LoadBalancersClient.CreateOrUpdate", "azure-async-operation", resp, client.pl)
	if err != nil {
		return LoadBalancersClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &LoadBalancersClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates a load balancer.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *LoadBalancersClient) createOrUpdate(ctx context.Context, resourceGroupName string, loadBalancerName string, parameters LoadBalancer, options *LoadBalancersClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, loadBalancerName, parameters, options)
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
func (client *LoadBalancersClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, loadBalancerName string, parameters LoadBalancer, options *LoadBalancersClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if loadBalancerName == "" {
		return nil, errors.New("parameter loadBalancerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancerName}", url.PathEscape(loadBalancerName))
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

// BeginDelete - Deletes the specified load balancer.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// loadBalancerName - The name of the load balancer.
// options - LoadBalancersClientBeginDeleteOptions contains the optional parameters for the LoadBalancersClient.BeginDelete
// method.
func (client *LoadBalancersClient) BeginDelete(ctx context.Context, resourceGroupName string, loadBalancerName string, options *LoadBalancersClientBeginDeleteOptions) (LoadBalancersClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, loadBalancerName, options)
	if err != nil {
		return LoadBalancersClientDeletePollerResponse{}, err
	}
	result := LoadBalancersClientDeletePollerResponse{}
	pt, err := armruntime.NewPoller("LoadBalancersClient.Delete", "location", resp, client.pl)
	if err != nil {
		return LoadBalancersClientDeletePollerResponse{}, err
	}
	result.Poller = &LoadBalancersClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes the specified load balancer.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *LoadBalancersClient) deleteOperation(ctx context.Context, resourceGroupName string, loadBalancerName string, options *LoadBalancersClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, loadBalancerName, options)
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
func (client *LoadBalancersClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, loadBalancerName string, options *LoadBalancersClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if loadBalancerName == "" {
		return nil, errors.New("parameter loadBalancerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancerName}", url.PathEscape(loadBalancerName))
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

// Get - Gets the specified load balancer.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// loadBalancerName - The name of the load balancer.
// options - LoadBalancersClientGetOptions contains the optional parameters for the LoadBalancersClient.Get method.
func (client *LoadBalancersClient) Get(ctx context.Context, resourceGroupName string, loadBalancerName string, options *LoadBalancersClientGetOptions) (LoadBalancersClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, loadBalancerName, options)
	if err != nil {
		return LoadBalancersClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return LoadBalancersClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LoadBalancersClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *LoadBalancersClient) getCreateRequest(ctx context.Context, resourceGroupName string, loadBalancerName string, options *LoadBalancersClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if loadBalancerName == "" {
		return nil, errors.New("parameter loadBalancerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancerName}", url.PathEscape(loadBalancerName))
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
func (client *LoadBalancersClient) getHandleResponse(resp *http.Response) (LoadBalancersClientGetResponse, error) {
	result := LoadBalancersClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LoadBalancer); err != nil {
		return LoadBalancersClientGetResponse{}, err
	}
	return result, nil
}

// List - Gets all the load balancers in a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// options - LoadBalancersClientListOptions contains the optional parameters for the LoadBalancersClient.List method.
func (client *LoadBalancersClient) List(resourceGroupName string, options *LoadBalancersClientListOptions) *LoadBalancersClientListPager {
	return &LoadBalancersClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp LoadBalancersClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.LoadBalancerListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *LoadBalancersClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *LoadBalancersClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers"
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

// listHandleResponse handles the List response.
func (client *LoadBalancersClient) listHandleResponse(resp *http.Response) (LoadBalancersClientListResponse, error) {
	result := LoadBalancersClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LoadBalancerListResult); err != nil {
		return LoadBalancersClientListResponse{}, err
	}
	return result, nil
}

// ListAll - Gets all the load balancers in a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - LoadBalancersClientListAllOptions contains the optional parameters for the LoadBalancersClient.ListAll method.
func (client *LoadBalancersClient) ListAll(options *LoadBalancersClientListAllOptions) *LoadBalancersClientListAllPager {
	return &LoadBalancersClientListAllPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listAllCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp LoadBalancersClientListAllResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.LoadBalancerListResult.NextLink)
		},
	}
}

// listAllCreateRequest creates the ListAll request.
func (client *LoadBalancersClient) listAllCreateRequest(ctx context.Context, options *LoadBalancersClientListAllOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/loadBalancers"
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

// listAllHandleResponse handles the ListAll response.
func (client *LoadBalancersClient) listAllHandleResponse(resp *http.Response) (LoadBalancersClientListAllResponse, error) {
	result := LoadBalancersClientListAllResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LoadBalancerListResult); err != nil {
		return LoadBalancersClientListAllResponse{}, err
	}
	return result, nil
}

// UpdateTags - Updates a load balancer tags.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// loadBalancerName - The name of the load balancer.
// parameters - Parameters supplied to update load balancer tags.
// options - LoadBalancersClientUpdateTagsOptions contains the optional parameters for the LoadBalancersClient.UpdateTags
// method.
func (client *LoadBalancersClient) UpdateTags(ctx context.Context, resourceGroupName string, loadBalancerName string, parameters TagsObject, options *LoadBalancersClientUpdateTagsOptions) (LoadBalancersClientUpdateTagsResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, loadBalancerName, parameters, options)
	if err != nil {
		return LoadBalancersClientUpdateTagsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return LoadBalancersClientUpdateTagsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LoadBalancersClientUpdateTagsResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateTagsHandleResponse(resp)
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *LoadBalancersClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, loadBalancerName string, parameters TagsObject, options *LoadBalancersClientUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if loadBalancerName == "" {
		return nil, errors.New("parameter loadBalancerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancerName}", url.PathEscape(loadBalancerName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
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
func (client *LoadBalancersClient) updateTagsHandleResponse(resp *http.Response) (LoadBalancersClientUpdateTagsResponse, error) {
	result := LoadBalancersClientUpdateTagsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LoadBalancer); err != nil {
		return LoadBalancersClientUpdateTagsResponse{}, err
	}
	return result, nil
}
