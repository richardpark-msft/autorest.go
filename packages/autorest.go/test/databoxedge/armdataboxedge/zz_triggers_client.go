//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdataboxedge

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

// TriggersClient contains the methods for the Triggers group.
// Don't use this type directly, use NewTriggersClient() instead.
type TriggersClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewTriggersClient creates a new instance of TriggersClient with the specified values.
//   - subscriptionID - The subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewTriggersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*TriggersClient, error) {
	cl, err := arm.NewClient(moduleName+".TriggersClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &TriggersClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a trigger.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-02-01
//   - deviceName - Creates or updates a trigger
//   - name - The trigger name.
//   - resourceGroupName - The resource group name.
//   - trigger - The trigger.
//   - options - TriggersClientBeginCreateOrUpdateOptions contains the optional parameters for the TriggersClient.BeginCreateOrUpdate
//     method.
func (client *TriggersClient) BeginCreateOrUpdate(ctx context.Context, deviceName string, name string, resourceGroupName string, trigger TriggerClassification, options *TriggersClientBeginCreateOrUpdateOptions) (*runtime.Poller[TriggersClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, deviceName, name, resourceGroupName, trigger, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[TriggersClientCreateOrUpdateResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[TriggersClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Creates or updates a trigger.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-02-01
func (client *TriggersClient) createOrUpdate(ctx context.Context, deviceName string, name string, resourceGroupName string, trigger TriggerClassification, options *TriggersClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "TriggersClient.BeginCreateOrUpdate")
	req, err := client.createOrUpdateCreateRequest(ctx, deviceName, name, resourceGroupName, trigger, options)
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

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *TriggersClient) createOrUpdateCreateRequest(ctx context.Context, deviceName string, name string, resourceGroupName string, trigger TriggerClassification, options *TriggersClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/triggers/{name}"
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, trigger); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes the trigger on the gateway device.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-02-01
//   - deviceName - The device name.
//   - name - The trigger name.
//   - resourceGroupName - The resource group name.
//   - options - TriggersClientBeginDeleteOptions contains the optional parameters for the TriggersClient.BeginDelete method.
func (client *TriggersClient) BeginDelete(ctx context.Context, deviceName string, name string, resourceGroupName string, options *TriggersClientBeginDeleteOptions) (*runtime.Poller[TriggersClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, deviceName, name, resourceGroupName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[TriggersClientDeleteResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[TriggersClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes the trigger on the gateway device.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-02-01
func (client *TriggersClient) deleteOperation(ctx context.Context, deviceName string, name string, resourceGroupName string, options *TriggersClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "TriggersClient.BeginDelete")
	req, err := client.deleteCreateRequest(ctx, deviceName, name, resourceGroupName, options)
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
func (client *TriggersClient) deleteCreateRequest(ctx context.Context, deviceName string, name string, resourceGroupName string, options *TriggersClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/triggers/{name}"
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a specific trigger by name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-02-01
//   - deviceName - The device name.
//   - name - The trigger name.
//   - resourceGroupName - The resource group name.
//   - options - TriggersClientGetOptions contains the optional parameters for the TriggersClient.Get method.
func (client *TriggersClient) Get(ctx context.Context, deviceName string, name string, resourceGroupName string, options *TriggersClientGetOptions) (TriggersClientGetResponse, error) {
	var err error
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "TriggersClient.Get")
	req, err := client.getCreateRequest(ctx, deviceName, name, resourceGroupName, options)
	if err != nil {
		return TriggersClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TriggersClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return TriggersClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *TriggersClient) getCreateRequest(ctx context.Context, deviceName string, name string, resourceGroupName string, options *TriggersClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/triggers/{name}"
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *TriggersClient) getHandleResponse(resp *http.Response) (TriggersClientGetResponse, error) {
	result := TriggersClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return TriggersClientGetResponse{}, err
	}
	return result, nil
}

// NewListByDataBoxEdgeDevicePager - Lists all the triggers configured in the device.
//
// Generated from API version 2021-02-01
//   - deviceName - The device name.
//   - resourceGroupName - The resource group name.
//   - options - TriggersClientListByDataBoxEdgeDeviceOptions contains the optional parameters for the TriggersClient.NewListByDataBoxEdgeDevicePager
//     method.
func (client *TriggersClient) NewListByDataBoxEdgeDevicePager(deviceName string, resourceGroupName string, options *TriggersClientListByDataBoxEdgeDeviceOptions) *runtime.Pager[TriggersClientListByDataBoxEdgeDeviceResponse] {
	return runtime.NewPager(runtime.PagingHandler[TriggersClientListByDataBoxEdgeDeviceResponse]{
		More: func(page TriggersClientListByDataBoxEdgeDeviceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *TriggersClientListByDataBoxEdgeDeviceResponse) (TriggersClientListByDataBoxEdgeDeviceResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "TriggersClient.NewListByDataBoxEdgeDevicePager")
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByDataBoxEdgeDeviceCreateRequest(ctx, deviceName, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return TriggersClientListByDataBoxEdgeDeviceResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return TriggersClientListByDataBoxEdgeDeviceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return TriggersClientListByDataBoxEdgeDeviceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByDataBoxEdgeDeviceHandleResponse(resp)
		},
	})
}

// listByDataBoxEdgeDeviceCreateRequest creates the ListByDataBoxEdgeDevice request.
func (client *TriggersClient) listByDataBoxEdgeDeviceCreateRequest(ctx context.Context, deviceName string, resourceGroupName string, options *TriggersClientListByDataBoxEdgeDeviceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/triggers"
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByDataBoxEdgeDeviceHandleResponse handles the ListByDataBoxEdgeDevice response.
func (client *TriggersClient) listByDataBoxEdgeDeviceHandleResponse(resp *http.Response) (TriggersClientListByDataBoxEdgeDeviceResponse, error) {
	result := TriggersClientListByDataBoxEdgeDeviceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TriggerList); err != nil {
		return TriggersClientListByDataBoxEdgeDeviceResponse{}, err
	}
	return result, nil
}
