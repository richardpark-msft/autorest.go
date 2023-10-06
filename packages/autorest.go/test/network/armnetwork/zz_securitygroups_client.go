//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// SecurityGroupsClient contains the methods for the NetworkSecurityGroups group.
// Don't use this type directly, use NewSecurityGroupsClient() instead.
type SecurityGroupsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSecurityGroupsClient creates a new instance of SecurityGroupsClient with the specified values.
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSecurityGroupsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SecurityGroupsClient, error) {
	cl, err := arm.NewClient(moduleName+".SecurityGroupsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SecurityGroupsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a network security group in the specified resource group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The name of the resource group.
//   - networkSecurityGroupName - The name of the network security group.
//   - parameters - Parameters supplied to the create or update network security group operation.
//   - options - SecurityGroupsClientBeginCreateOrUpdateOptions contains the optional parameters for the SecurityGroupsClient.BeginCreateOrUpdate
//     method.
func (client *SecurityGroupsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, parameters SecurityGroup, options *SecurityGroupsClientBeginCreateOrUpdateOptions) (*runtime.Poller[SecurityGroupsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, networkSecurityGroupName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[SecurityGroupsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[SecurityGroupsClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Creates or updates a network security group in the specified resource group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
func (client *SecurityGroupsClient) createOrUpdate(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, parameters SecurityGroup, options *SecurityGroupsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "SecurityGroupsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, networkSecurityGroupName, parameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *SecurityGroupsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, parameters SecurityGroup, options *SecurityGroupsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkSecurityGroupName == "" {
		return nil, errors.New("parameter networkSecurityGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkSecurityGroupName}", url.PathEscape(networkSecurityGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes the specified network security group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The name of the resource group.
//   - networkSecurityGroupName - The name of the network security group.
//   - options - SecurityGroupsClientBeginDeleteOptions contains the optional parameters for the SecurityGroupsClient.BeginDelete
//     method.
func (client *SecurityGroupsClient) BeginDelete(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, options *SecurityGroupsClientBeginDeleteOptions) (*runtime.Poller[SecurityGroupsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, networkSecurityGroupName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[SecurityGroupsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[SecurityGroupsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes the specified network security group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
func (client *SecurityGroupsClient) deleteOperation(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, options *SecurityGroupsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "SecurityGroupsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, networkSecurityGroupName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *SecurityGroupsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, options *SecurityGroupsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkSecurityGroupName == "" {
		return nil, errors.New("parameter networkSecurityGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkSecurityGroupName}", url.PathEscape(networkSecurityGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the specified network security group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The name of the resource group.
//   - networkSecurityGroupName - The name of the network security group.
//   - options - SecurityGroupsClientGetOptions contains the optional parameters for the SecurityGroupsClient.Get method.
func (client *SecurityGroupsClient) Get(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, options *SecurityGroupsClientGetOptions) (SecurityGroupsClientGetResponse, error) {
	var err error
	const operationName = "SecurityGroupsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, networkSecurityGroupName, options)
	if err != nil {
		return SecurityGroupsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SecurityGroupsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SecurityGroupsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *SecurityGroupsClient) getCreateRequest(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, options *SecurityGroupsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkSecurityGroupName == "" {
		return nil, errors.New("parameter networkSecurityGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkSecurityGroupName}", url.PathEscape(networkSecurityGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SecurityGroupsClient) getHandleResponse(resp *http.Response) (SecurityGroupsClientGetResponse, error) {
	result := SecurityGroupsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecurityGroup); err != nil {
		return SecurityGroupsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets all network security groups in a resource group.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The name of the resource group.
//   - options - SecurityGroupsClientListOptions contains the optional parameters for the SecurityGroupsClient.NewListPager method.
func (client *SecurityGroupsClient) NewListPager(resourceGroupName string, options *SecurityGroupsClientListOptions) *runtime.Pager[SecurityGroupsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[SecurityGroupsClientListResponse]{
		More: func(page SecurityGroupsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SecurityGroupsClientListResponse) (SecurityGroupsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "SecurityGroupsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return SecurityGroupsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *SecurityGroupsClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *SecurityGroupsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SecurityGroupsClient) listHandleResponse(resp *http.Response) (SecurityGroupsClientListResponse, error) {
	result := SecurityGroupsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecurityGroupListResult); err != nil {
		return SecurityGroupsClientListResponse{}, err
	}
	return result, nil
}

// NewListAllPager - Gets all network security groups in a subscription.
//
// Generated from API version 2022-09-01
//   - options - SecurityGroupsClientListAllOptions contains the optional parameters for the SecurityGroupsClient.NewListAllPager
//     method.
func (client *SecurityGroupsClient) NewListAllPager(options *SecurityGroupsClientListAllOptions) *runtime.Pager[SecurityGroupsClientListAllResponse] {
	return runtime.NewPager(runtime.PagingHandler[SecurityGroupsClientListAllResponse]{
		More: func(page SecurityGroupsClientListAllResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SecurityGroupsClientListAllResponse) (SecurityGroupsClientListAllResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "SecurityGroupsClient.NewListAllPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listAllCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return SecurityGroupsClientListAllResponse{}, err
			}
			return client.listAllHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listAllCreateRequest creates the ListAll request.
func (client *SecurityGroupsClient) listAllCreateRequest(ctx context.Context, options *SecurityGroupsClientListAllOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/networkSecurityGroups"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listAllHandleResponse handles the ListAll response.
func (client *SecurityGroupsClient) listAllHandleResponse(resp *http.Response) (SecurityGroupsClientListAllResponse, error) {
	result := SecurityGroupsClientListAllResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecurityGroupListResult); err != nil {
		return SecurityGroupsClientListAllResponse{}, err
	}
	return result, nil
}

// UpdateTags - Updates a network security group tags.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The name of the resource group.
//   - networkSecurityGroupName - The name of the network security group.
//   - parameters - Parameters supplied to update network security group tags.
//   - options - SecurityGroupsClientUpdateTagsOptions contains the optional parameters for the SecurityGroupsClient.UpdateTags
//     method.
func (client *SecurityGroupsClient) UpdateTags(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, parameters TagsObject, options *SecurityGroupsClientUpdateTagsOptions) (SecurityGroupsClientUpdateTagsResponse, error) {
	var err error
	const operationName = "SecurityGroupsClient.UpdateTags"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, networkSecurityGroupName, parameters, options)
	if err != nil {
		return SecurityGroupsClientUpdateTagsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SecurityGroupsClientUpdateTagsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SecurityGroupsClientUpdateTagsResponse{}, err
	}
	resp, err := client.updateTagsHandleResponse(httpResp)
	return resp, err
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *SecurityGroupsClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, parameters TagsObject, options *SecurityGroupsClientUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkSecurityGroupName == "" {
		return nil, errors.New("parameter networkSecurityGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkSecurityGroupName}", url.PathEscape(networkSecurityGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *SecurityGroupsClient) updateTagsHandleResponse(resp *http.Response) (SecurityGroupsClientUpdateTagsResponse, error) {
	result := SecurityGroupsClientUpdateTagsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecurityGroup); err != nil {
		return SecurityGroupsClientUpdateTagsResponse{}, err
	}
	return result, nil
}
