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

// ApplicationSecurityGroupsClient contains the methods for the ApplicationSecurityGroups group.
// Don't use this type directly, use NewApplicationSecurityGroupsClient() instead.
type ApplicationSecurityGroupsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewApplicationSecurityGroupsClient creates a new instance of ApplicationSecurityGroupsClient with the specified values.
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewApplicationSecurityGroupsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ApplicationSecurityGroupsClient, error) {
	cl, err := arm.NewClient(moduleName+".ApplicationSecurityGroupsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ApplicationSecurityGroupsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates an application security group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The name of the resource group.
//   - applicationSecurityGroupName - The name of the application security group.
//   - parameters - Parameters supplied to the create or update ApplicationSecurityGroup operation.
//   - options - ApplicationSecurityGroupsClientBeginCreateOrUpdateOptions contains the optional parameters for the ApplicationSecurityGroupsClient.BeginCreateOrUpdate
//     method.
func (client *ApplicationSecurityGroupsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, applicationSecurityGroupName string, parameters ApplicationSecurityGroup, options *ApplicationSecurityGroupsClientBeginCreateOrUpdateOptions) (*runtime.Poller[ApplicationSecurityGroupsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, applicationSecurityGroupName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ApplicationSecurityGroupsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[ApplicationSecurityGroupsClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Creates or updates an application security group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
func (client *ApplicationSecurityGroupsClient) createOrUpdate(ctx context.Context, resourceGroupName string, applicationSecurityGroupName string, parameters ApplicationSecurityGroup, options *ApplicationSecurityGroupsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "ApplicationSecurityGroupsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, applicationSecurityGroupName, parameters, options)
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
func (client *ApplicationSecurityGroupsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, applicationSecurityGroupName string, parameters ApplicationSecurityGroup, options *ApplicationSecurityGroupsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationSecurityGroups/{applicationSecurityGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationSecurityGroupName == "" {
		return nil, errors.New("parameter applicationSecurityGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationSecurityGroupName}", url.PathEscape(applicationSecurityGroupName))
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

// BeginDelete - Deletes the specified application security group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The name of the resource group.
//   - applicationSecurityGroupName - The name of the application security group.
//   - options - ApplicationSecurityGroupsClientBeginDeleteOptions contains the optional parameters for the ApplicationSecurityGroupsClient.BeginDelete
//     method.
func (client *ApplicationSecurityGroupsClient) BeginDelete(ctx context.Context, resourceGroupName string, applicationSecurityGroupName string, options *ApplicationSecurityGroupsClientBeginDeleteOptions) (*runtime.Poller[ApplicationSecurityGroupsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, applicationSecurityGroupName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ApplicationSecurityGroupsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[ApplicationSecurityGroupsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes the specified application security group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
func (client *ApplicationSecurityGroupsClient) deleteOperation(ctx context.Context, resourceGroupName string, applicationSecurityGroupName string, options *ApplicationSecurityGroupsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "ApplicationSecurityGroupsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, applicationSecurityGroupName, options)
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
func (client *ApplicationSecurityGroupsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, applicationSecurityGroupName string, options *ApplicationSecurityGroupsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationSecurityGroups/{applicationSecurityGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationSecurityGroupName == "" {
		return nil, errors.New("parameter applicationSecurityGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationSecurityGroupName}", url.PathEscape(applicationSecurityGroupName))
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

// Get - Gets information about the specified application security group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The name of the resource group.
//   - applicationSecurityGroupName - The name of the application security group.
//   - options - ApplicationSecurityGroupsClientGetOptions contains the optional parameters for the ApplicationSecurityGroupsClient.Get
//     method.
func (client *ApplicationSecurityGroupsClient) Get(ctx context.Context, resourceGroupName string, applicationSecurityGroupName string, options *ApplicationSecurityGroupsClientGetOptions) (ApplicationSecurityGroupsClientGetResponse, error) {
	var err error
	const operationName = "ApplicationSecurityGroupsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, applicationSecurityGroupName, options)
	if err != nil {
		return ApplicationSecurityGroupsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ApplicationSecurityGroupsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ApplicationSecurityGroupsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ApplicationSecurityGroupsClient) getCreateRequest(ctx context.Context, resourceGroupName string, applicationSecurityGroupName string, options *ApplicationSecurityGroupsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationSecurityGroups/{applicationSecurityGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationSecurityGroupName == "" {
		return nil, errors.New("parameter applicationSecurityGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationSecurityGroupName}", url.PathEscape(applicationSecurityGroupName))
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

// getHandleResponse handles the Get response.
func (client *ApplicationSecurityGroupsClient) getHandleResponse(resp *http.Response) (ApplicationSecurityGroupsClientGetResponse, error) {
	result := ApplicationSecurityGroupsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationSecurityGroup); err != nil {
		return ApplicationSecurityGroupsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets all the application security groups in a resource group.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The name of the resource group.
//   - options - ApplicationSecurityGroupsClientListOptions contains the optional parameters for the ApplicationSecurityGroupsClient.NewListPager
//     method.
func (client *ApplicationSecurityGroupsClient) NewListPager(resourceGroupName string, options *ApplicationSecurityGroupsClientListOptions) *runtime.Pager[ApplicationSecurityGroupsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ApplicationSecurityGroupsClientListResponse]{
		More: func(page ApplicationSecurityGroupsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ApplicationSecurityGroupsClientListResponse) (ApplicationSecurityGroupsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ApplicationSecurityGroupsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return ApplicationSecurityGroupsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *ApplicationSecurityGroupsClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *ApplicationSecurityGroupsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationSecurityGroups"
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
func (client *ApplicationSecurityGroupsClient) listHandleResponse(resp *http.Response) (ApplicationSecurityGroupsClientListResponse, error) {
	result := ApplicationSecurityGroupsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationSecurityGroupListResult); err != nil {
		return ApplicationSecurityGroupsClientListResponse{}, err
	}
	return result, nil
}

// NewListAllPager - Gets all application security groups in a subscription.
//
// Generated from API version 2022-09-01
//   - options - ApplicationSecurityGroupsClientListAllOptions contains the optional parameters for the ApplicationSecurityGroupsClient.NewListAllPager
//     method.
func (client *ApplicationSecurityGroupsClient) NewListAllPager(options *ApplicationSecurityGroupsClientListAllOptions) *runtime.Pager[ApplicationSecurityGroupsClientListAllResponse] {
	return runtime.NewPager(runtime.PagingHandler[ApplicationSecurityGroupsClientListAllResponse]{
		More: func(page ApplicationSecurityGroupsClientListAllResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ApplicationSecurityGroupsClientListAllResponse) (ApplicationSecurityGroupsClientListAllResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ApplicationSecurityGroupsClient.NewListAllPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listAllCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return ApplicationSecurityGroupsClientListAllResponse{}, err
			}
			return client.listAllHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listAllCreateRequest creates the ListAll request.
func (client *ApplicationSecurityGroupsClient) listAllCreateRequest(ctx context.Context, options *ApplicationSecurityGroupsClientListAllOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/applicationSecurityGroups"
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
func (client *ApplicationSecurityGroupsClient) listAllHandleResponse(resp *http.Response) (ApplicationSecurityGroupsClientListAllResponse, error) {
	result := ApplicationSecurityGroupsClientListAllResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationSecurityGroupListResult); err != nil {
		return ApplicationSecurityGroupsClientListAllResponse{}, err
	}
	return result, nil
}

// UpdateTags - Updates an application security group's tags.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The name of the resource group.
//   - applicationSecurityGroupName - The name of the application security group.
//   - parameters - Parameters supplied to update application security group tags.
//   - options - ApplicationSecurityGroupsClientUpdateTagsOptions contains the optional parameters for the ApplicationSecurityGroupsClient.UpdateTags
//     method.
func (client *ApplicationSecurityGroupsClient) UpdateTags(ctx context.Context, resourceGroupName string, applicationSecurityGroupName string, parameters TagsObject, options *ApplicationSecurityGroupsClientUpdateTagsOptions) (ApplicationSecurityGroupsClientUpdateTagsResponse, error) {
	var err error
	const operationName = "ApplicationSecurityGroupsClient.UpdateTags"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, applicationSecurityGroupName, parameters, options)
	if err != nil {
		return ApplicationSecurityGroupsClientUpdateTagsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ApplicationSecurityGroupsClientUpdateTagsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ApplicationSecurityGroupsClientUpdateTagsResponse{}, err
	}
	resp, err := client.updateTagsHandleResponse(httpResp)
	return resp, err
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *ApplicationSecurityGroupsClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, applicationSecurityGroupName string, parameters TagsObject, options *ApplicationSecurityGroupsClientUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationSecurityGroups/{applicationSecurityGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationSecurityGroupName == "" {
		return nil, errors.New("parameter applicationSecurityGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationSecurityGroupName}", url.PathEscape(applicationSecurityGroupName))
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
func (client *ApplicationSecurityGroupsClient) updateTagsHandleResponse(resp *http.Response) (ApplicationSecurityGroupsClientUpdateTagsResponse, error) {
	result := ApplicationSecurityGroupsClientUpdateTagsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationSecurityGroup); err != nil {
		return ApplicationSecurityGroupsClientUpdateTagsResponse{}, err
	}
	return result, nil
}
