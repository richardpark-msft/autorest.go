//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// CloudServicesUpdateDomainClient contains the methods for the CloudServicesUpdateDomain group.
// Don't use this type directly, use NewCloudServicesUpdateDomainClient() instead.
type CloudServicesUpdateDomainClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewCloudServicesUpdateDomainClient creates a new instance of CloudServicesUpdateDomainClient with the specified values.
//   - subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
//     part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewCloudServicesUpdateDomainClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*CloudServicesUpdateDomainClient, error) {
	cl, err := arm.NewClient(moduleName+".CloudServicesUpdateDomainClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &CloudServicesUpdateDomainClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// GetUpdateDomain - Gets the specified update domain of a cloud service. Use nextLink property in the response to get the
// next page of update domains. Do this till nextLink is null to fetch all the update domains.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-03-01
//   - resourceGroupName - Name of the resource group.
//   - cloudServiceName - Name of the cloud service.
//   - updateDomain - Specifies an integer value that identifies the update domain. Update domains are identified with a zero-based
//     index: the first update domain has an ID of 0, the second has an ID of 1, and so on.
//   - options - CloudServicesUpdateDomainClientGetUpdateDomainOptions contains the optional parameters for the CloudServicesUpdateDomainClient.GetUpdateDomain
//     method.
func (client *CloudServicesUpdateDomainClient) GetUpdateDomain(ctx context.Context, resourceGroupName string, cloudServiceName string, updateDomain int32, options *CloudServicesUpdateDomainClientGetUpdateDomainOptions) (CloudServicesUpdateDomainClientGetUpdateDomainResponse, error) {
	var err error
	const operationName = "CloudServicesUpdateDomainClient.GetUpdateDomain"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getUpdateDomainCreateRequest(ctx, resourceGroupName, cloudServiceName, updateDomain, options)
	if err != nil {
		return CloudServicesUpdateDomainClientGetUpdateDomainResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CloudServicesUpdateDomainClientGetUpdateDomainResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return CloudServicesUpdateDomainClientGetUpdateDomainResponse{}, err
	}
	resp, err := client.getUpdateDomainHandleResponse(httpResp)
	return resp, err
}

// getUpdateDomainCreateRequest creates the GetUpdateDomain request.
func (client *CloudServicesUpdateDomainClient) getUpdateDomainCreateRequest(ctx context.Context, resourceGroupName string, cloudServiceName string, updateDomain int32, options *CloudServicesUpdateDomainClientGetUpdateDomainOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/cloudServices/{cloudServiceName}/updateDomains/{updateDomain}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if cloudServiceName == "" {
		return nil, errors.New("parameter cloudServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cloudServiceName}", url.PathEscape(cloudServiceName))
	urlPath = strings.ReplaceAll(urlPath, "{updateDomain}", url.PathEscape(strconv.FormatInt(int64(updateDomain), 10)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getUpdateDomainHandleResponse handles the GetUpdateDomain response.
func (client *CloudServicesUpdateDomainClient) getUpdateDomainHandleResponse(resp *http.Response) (CloudServicesUpdateDomainClientGetUpdateDomainResponse, error) {
	result := CloudServicesUpdateDomainClientGetUpdateDomainResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.UpdateDomain); err != nil {
		return CloudServicesUpdateDomainClientGetUpdateDomainResponse{}, err
	}
	return result, nil
}

// NewListUpdateDomainsPager - Gets a list of all update domains in a cloud service.
//
// Generated from API version 2021-03-01
//   - resourceGroupName - Name of the resource group.
//   - cloudServiceName - Name of the cloud service.
//   - options - CloudServicesUpdateDomainClientListUpdateDomainsOptions contains the optional parameters for the CloudServicesUpdateDomainClient.NewListUpdateDomainsPager
//     method.
func (client *CloudServicesUpdateDomainClient) NewListUpdateDomainsPager(resourceGroupName string, cloudServiceName string, options *CloudServicesUpdateDomainClientListUpdateDomainsOptions) *runtime.Pager[CloudServicesUpdateDomainClientListUpdateDomainsResponse] {
	return runtime.NewPager(runtime.PagingHandler[CloudServicesUpdateDomainClientListUpdateDomainsResponse]{
		More: func(page CloudServicesUpdateDomainClientListUpdateDomainsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *CloudServicesUpdateDomainClientListUpdateDomainsResponse) (CloudServicesUpdateDomainClientListUpdateDomainsResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "CloudServicesUpdateDomainClient.NewListUpdateDomainsPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listUpdateDomainsCreateRequest(ctx, resourceGroupName, cloudServiceName, options)
			}, nil)
			if err != nil {
				return CloudServicesUpdateDomainClientListUpdateDomainsResponse{}, err
			}
			return client.listUpdateDomainsHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listUpdateDomainsCreateRequest creates the ListUpdateDomains request.
func (client *CloudServicesUpdateDomainClient) listUpdateDomainsCreateRequest(ctx context.Context, resourceGroupName string, cloudServiceName string, options *CloudServicesUpdateDomainClientListUpdateDomainsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/cloudServices/{cloudServiceName}/updateDomains"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if cloudServiceName == "" {
		return nil, errors.New("parameter cloudServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cloudServiceName}", url.PathEscape(cloudServiceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listUpdateDomainsHandleResponse handles the ListUpdateDomains response.
func (client *CloudServicesUpdateDomainClient) listUpdateDomainsHandleResponse(resp *http.Response) (CloudServicesUpdateDomainClientListUpdateDomainsResponse, error) {
	result := CloudServicesUpdateDomainClientListUpdateDomainsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.UpdateDomainListResult); err != nil {
		return CloudServicesUpdateDomainClientListUpdateDomainsResponse{}, err
	}
	return result, nil
}

// BeginWalkUpdateDomain - Updates the role instances in the specified update domain.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-03-01
//   - resourceGroupName - Name of the resource group.
//   - cloudServiceName - Name of the cloud service.
//   - updateDomain - Specifies an integer value that identifies the update domain. Update domains are identified with a zero-based
//     index: the first update domain has an ID of 0, the second has an ID of 1, and so on.
//   - parameters - The update domain object.
//   - options - CloudServicesUpdateDomainClientBeginWalkUpdateDomainOptions contains the optional parameters for the CloudServicesUpdateDomainClient.BeginWalkUpdateDomain
//     method.
func (client *CloudServicesUpdateDomainClient) BeginWalkUpdateDomain(ctx context.Context, resourceGroupName string, cloudServiceName string, updateDomain int32, parameters UpdateDomain, options *CloudServicesUpdateDomainClientBeginWalkUpdateDomainOptions) (*runtime.Poller[CloudServicesUpdateDomainClientWalkUpdateDomainResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.walkUpdateDomain(ctx, resourceGroupName, cloudServiceName, updateDomain, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[CloudServicesUpdateDomainClientWalkUpdateDomainResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[CloudServicesUpdateDomainClientWalkUpdateDomainResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// WalkUpdateDomain - Updates the role instances in the specified update domain.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-03-01
func (client *CloudServicesUpdateDomainClient) walkUpdateDomain(ctx context.Context, resourceGroupName string, cloudServiceName string, updateDomain int32, parameters UpdateDomain, options *CloudServicesUpdateDomainClientBeginWalkUpdateDomainOptions) (*http.Response, error) {
	var err error
	const operationName = "CloudServicesUpdateDomainClient.BeginWalkUpdateDomain"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.walkUpdateDomainCreateRequest(ctx, resourceGroupName, cloudServiceName, updateDomain, parameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// walkUpdateDomainCreateRequest creates the WalkUpdateDomain request.
func (client *CloudServicesUpdateDomainClient) walkUpdateDomainCreateRequest(ctx context.Context, resourceGroupName string, cloudServiceName string, updateDomain int32, parameters UpdateDomain, options *CloudServicesUpdateDomainClientBeginWalkUpdateDomainOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/cloudServices/{cloudServiceName}/updateDomains/{updateDomain}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if cloudServiceName == "" {
		return nil, errors.New("parameter cloudServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cloudServiceName}", url.PathEscape(cloudServiceName))
	urlPath = strings.ReplaceAll(urlPath, "{updateDomain}", url.PathEscape(strconv.FormatInt(int64(updateDomain), 10)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}
