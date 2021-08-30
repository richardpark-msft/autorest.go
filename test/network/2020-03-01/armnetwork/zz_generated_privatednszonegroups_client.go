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
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// PrivateDNSZoneGroupsClient contains the methods for the PrivateDNSZoneGroups group.
// Don't use this type directly, use NewPrivateDNSZoneGroupsClient() instead.
type PrivateDNSZoneGroupsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewPrivateDNSZoneGroupsClient creates a new instance of PrivateDNSZoneGroupsClient with the specified values.
func NewPrivateDNSZoneGroupsClient(con *arm.Connection, subscriptionID string) *PrivateDNSZoneGroupsClient {
	return &PrivateDNSZoneGroupsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates or updates a private dns zone group in the specified private endpoint.
// If the operation fails it returns the *CloudError error type.
func (client *PrivateDNSZoneGroupsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, privateEndpointName string, privateDNSZoneGroupName string, parameters PrivateDNSZoneGroup, options *PrivateDNSZoneGroupsBeginCreateOrUpdateOptions) (PrivateDNSZoneGroupsCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, privateEndpointName, privateDNSZoneGroupName, parameters, options)
	if err != nil {
		return PrivateDNSZoneGroupsCreateOrUpdatePollerResponse{}, err
	}
	result := PrivateDNSZoneGroupsCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("PrivateDNSZoneGroupsClient.CreateOrUpdate", "azure-async-operation", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return PrivateDNSZoneGroupsCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &PrivateDNSZoneGroupsCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates a private dns zone group in the specified private endpoint.
// If the operation fails it returns the *CloudError error type.
func (client *PrivateDNSZoneGroupsClient) createOrUpdate(ctx context.Context, resourceGroupName string, privateEndpointName string, privateDNSZoneGroupName string, parameters PrivateDNSZoneGroup, options *PrivateDNSZoneGroupsBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, privateEndpointName, privateDNSZoneGroupName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *PrivateDNSZoneGroupsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, privateEndpointName string, privateDNSZoneGroupName string, parameters PrivateDNSZoneGroup, options *PrivateDNSZoneGroupsBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateEndpoints/{privateEndpointName}/privateDnsZoneGroups/{privateDnsZoneGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateEndpointName == "" {
		return nil, errors.New("parameter privateEndpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointName}", url.PathEscape(privateEndpointName))
	if privateDNSZoneGroupName == "" {
		return nil, errors.New("parameter privateDNSZoneGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateDnsZoneGroupName}", url.PathEscape(privateDNSZoneGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *PrivateDNSZoneGroupsClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginDelete - Deletes the specified private dns zone group.
// If the operation fails it returns the *CloudError error type.
func (client *PrivateDNSZoneGroupsClient) BeginDelete(ctx context.Context, resourceGroupName string, privateEndpointName string, privateDNSZoneGroupName string, options *PrivateDNSZoneGroupsBeginDeleteOptions) (PrivateDNSZoneGroupsDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, privateEndpointName, privateDNSZoneGroupName, options)
	if err != nil {
		return PrivateDNSZoneGroupsDeletePollerResponse{}, err
	}
	result := PrivateDNSZoneGroupsDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("PrivateDNSZoneGroupsClient.Delete", "location", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return PrivateDNSZoneGroupsDeletePollerResponse{}, err
	}
	result.Poller = &PrivateDNSZoneGroupsDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes the specified private dns zone group.
// If the operation fails it returns the *CloudError error type.
func (client *PrivateDNSZoneGroupsClient) deleteOperation(ctx context.Context, resourceGroupName string, privateEndpointName string, privateDNSZoneGroupName string, options *PrivateDNSZoneGroupsBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, privateEndpointName, privateDNSZoneGroupName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *PrivateDNSZoneGroupsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, privateEndpointName string, privateDNSZoneGroupName string, options *PrivateDNSZoneGroupsBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateEndpoints/{privateEndpointName}/privateDnsZoneGroups/{privateDnsZoneGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateEndpointName == "" {
		return nil, errors.New("parameter privateEndpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointName}", url.PathEscape(privateEndpointName))
	if privateDNSZoneGroupName == "" {
		return nil, errors.New("parameter privateDNSZoneGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateDnsZoneGroupName}", url.PathEscape(privateDNSZoneGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *PrivateDNSZoneGroupsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Gets the private dns zone group resource by specified private dns zone group name.
// If the operation fails it returns the *CloudError error type.
func (client *PrivateDNSZoneGroupsClient) Get(ctx context.Context, resourceGroupName string, privateEndpointName string, privateDNSZoneGroupName string, options *PrivateDNSZoneGroupsGetOptions) (PrivateDNSZoneGroupsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, privateEndpointName, privateDNSZoneGroupName, options)
	if err != nil {
		return PrivateDNSZoneGroupsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PrivateDNSZoneGroupsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PrivateDNSZoneGroupsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PrivateDNSZoneGroupsClient) getCreateRequest(ctx context.Context, resourceGroupName string, privateEndpointName string, privateDNSZoneGroupName string, options *PrivateDNSZoneGroupsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateEndpoints/{privateEndpointName}/privateDnsZoneGroups/{privateDnsZoneGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateEndpointName == "" {
		return nil, errors.New("parameter privateEndpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointName}", url.PathEscape(privateEndpointName))
	if privateDNSZoneGroupName == "" {
		return nil, errors.New("parameter privateDNSZoneGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateDnsZoneGroupName}", url.PathEscape(privateDNSZoneGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
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
func (client *PrivateDNSZoneGroupsClient) getHandleResponse(resp *http.Response) (PrivateDNSZoneGroupsGetResponse, error) {
	result := PrivateDNSZoneGroupsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateDNSZoneGroup); err != nil {
		return PrivateDNSZoneGroupsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *PrivateDNSZoneGroupsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - Gets all private dns zone groups in a private endpoint.
// If the operation fails it returns the *Error error type.
func (client *PrivateDNSZoneGroupsClient) List(privateEndpointName string, resourceGroupName string, options *PrivateDNSZoneGroupsListOptions) *PrivateDNSZoneGroupsListPager {
	return &PrivateDNSZoneGroupsListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, privateEndpointName, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp PrivateDNSZoneGroupsListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.PrivateDNSZoneGroupListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *PrivateDNSZoneGroupsClient) listCreateRequest(ctx context.Context, privateEndpointName string, resourceGroupName string, options *PrivateDNSZoneGroupsListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateEndpoints/{privateEndpointName}/privateDnsZoneGroups"
	if privateEndpointName == "" {
		return nil, errors.New("parameter privateEndpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointName}", url.PathEscape(privateEndpointName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
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
func (client *PrivateDNSZoneGroupsClient) listHandleResponse(resp *http.Response) (PrivateDNSZoneGroupsListResponse, error) {
	result := PrivateDNSZoneGroupsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateDNSZoneGroupListResult); err != nil {
		return PrivateDNSZoneGroupsListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *PrivateDNSZoneGroupsClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
