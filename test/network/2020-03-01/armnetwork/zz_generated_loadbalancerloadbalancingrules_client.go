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

// LoadBalancerLoadBalancingRulesClient contains the methods for the LoadBalancerLoadBalancingRules group.
// Don't use this type directly, use NewLoadBalancerLoadBalancingRulesClient() instead.
type LoadBalancerLoadBalancingRulesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewLoadBalancerLoadBalancingRulesClient creates a new instance of LoadBalancerLoadBalancingRulesClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewLoadBalancerLoadBalancingRulesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *LoadBalancerLoadBalancingRulesClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &LoadBalancerLoadBalancingRulesClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// Get - Gets the specified load balancer load balancing rule.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// loadBalancerName - The name of the load balancer.
// loadBalancingRuleName - The name of the load balancing rule.
// options - LoadBalancerLoadBalancingRulesClientGetOptions contains the optional parameters for the LoadBalancerLoadBalancingRulesClient.Get
// method.
func (client *LoadBalancerLoadBalancingRulesClient) Get(ctx context.Context, resourceGroupName string, loadBalancerName string, loadBalancingRuleName string, options *LoadBalancerLoadBalancingRulesClientGetOptions) (LoadBalancerLoadBalancingRulesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, loadBalancerName, loadBalancingRuleName, options)
	if err != nil {
		return LoadBalancerLoadBalancingRulesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return LoadBalancerLoadBalancingRulesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LoadBalancerLoadBalancingRulesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *LoadBalancerLoadBalancingRulesClient) getCreateRequest(ctx context.Context, resourceGroupName string, loadBalancerName string, loadBalancingRuleName string, options *LoadBalancerLoadBalancingRulesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/loadBalancingRules/{loadBalancingRuleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if loadBalancerName == "" {
		return nil, errors.New("parameter loadBalancerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancerName}", url.PathEscape(loadBalancerName))
	if loadBalancingRuleName == "" {
		return nil, errors.New("parameter loadBalancingRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancingRuleName}", url.PathEscape(loadBalancingRuleName))
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

// getHandleResponse handles the Get response.
func (client *LoadBalancerLoadBalancingRulesClient) getHandleResponse(resp *http.Response) (LoadBalancerLoadBalancingRulesClientGetResponse, error) {
	result := LoadBalancerLoadBalancingRulesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LoadBalancingRule); err != nil {
		return LoadBalancerLoadBalancingRulesClientGetResponse{}, err
	}
	return result, nil
}

// List - Gets all the load balancing rules in a load balancer.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// loadBalancerName - The name of the load balancer.
// options - LoadBalancerLoadBalancingRulesClientListOptions contains the optional parameters for the LoadBalancerLoadBalancingRulesClient.List
// method.
func (client *LoadBalancerLoadBalancingRulesClient) List(resourceGroupName string, loadBalancerName string, options *LoadBalancerLoadBalancingRulesClientListOptions) *LoadBalancerLoadBalancingRulesClientListPager {
	return &LoadBalancerLoadBalancingRulesClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, loadBalancerName, options)
		},
		advancer: func(ctx context.Context, resp LoadBalancerLoadBalancingRulesClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.LoadBalancerLoadBalancingRuleListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *LoadBalancerLoadBalancingRulesClient) listCreateRequest(ctx context.Context, resourceGroupName string, loadBalancerName string, options *LoadBalancerLoadBalancingRulesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/loadBalancingRules"
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
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *LoadBalancerLoadBalancingRulesClient) listHandleResponse(resp *http.Response) (LoadBalancerLoadBalancingRulesClientListResponse, error) {
	result := LoadBalancerLoadBalancingRulesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LoadBalancerLoadBalancingRuleListResult); err != nil {
		return LoadBalancerLoadBalancingRulesClientListResponse{}, err
	}
	return result, nil
}
