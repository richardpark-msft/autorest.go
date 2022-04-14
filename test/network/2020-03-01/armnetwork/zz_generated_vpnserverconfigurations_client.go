//go:build go1.18
// +build go1.18

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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// VPNServerConfigurationsClient contains the methods for the VPNServerConfigurations group.
// Don't use this type directly, use NewVPNServerConfigurationsClient() instead.
type VPNServerConfigurationsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewVPNServerConfigurationsClient creates a new instance of VPNServerConfigurationsClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewVPNServerConfigurationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*VPNServerConfigurationsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublicCloud.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &VPNServerConfigurationsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates a VpnServerConfiguration resource if it doesn't exist else updates the existing VpnServerConfiguration.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name of the VpnServerConfiguration.
// vpnServerConfigurationName - The name of the VpnServerConfiguration being created or updated.
// vpnServerConfigurationParameters - Parameters supplied to create or update VpnServerConfiguration.
// options - VPNServerConfigurationsClientBeginCreateOrUpdateOptions contains the optional parameters for the VPNServerConfigurationsClient.BeginCreateOrUpdate
// method.
func (client *VPNServerConfigurationsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, vpnServerConfigurationName string, vpnServerConfigurationParameters VPNServerConfiguration, options *VPNServerConfigurationsClientBeginCreateOrUpdateOptions) (*armruntime.Poller[VPNServerConfigurationsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, vpnServerConfigurationName, vpnServerConfigurationParameters, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller(resp, client.pl, &armruntime.NewPollerOptions[VPNServerConfigurationsClientCreateOrUpdateResponse]{
			FinalStateVia: armruntime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return armruntime.NewPollerFromResumeToken[VPNServerConfigurationsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates a VpnServerConfiguration resource if it doesn't exist else updates the existing VpnServerConfiguration.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *VPNServerConfigurationsClient) createOrUpdate(ctx context.Context, resourceGroupName string, vpnServerConfigurationName string, vpnServerConfigurationParameters VPNServerConfiguration, options *VPNServerConfigurationsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, vpnServerConfigurationName, vpnServerConfigurationParameters, options)
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
func (client *VPNServerConfigurationsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, vpnServerConfigurationName string, vpnServerConfigurationParameters VPNServerConfiguration, options *VPNServerConfigurationsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnServerConfigurations/{vpnServerConfigurationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vpnServerConfigurationName == "" {
		return nil, errors.New("parameter vpnServerConfigurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vpnServerConfigurationName}", url.PathEscape(vpnServerConfigurationName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, vpnServerConfigurationParameters)
}

// BeginDelete - Deletes a VpnServerConfiguration.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name of the VpnServerConfiguration.
// vpnServerConfigurationName - The name of the VpnServerConfiguration being deleted.
// options - VPNServerConfigurationsClientBeginDeleteOptions contains the optional parameters for the VPNServerConfigurationsClient.BeginDelete
// method.
func (client *VPNServerConfigurationsClient) BeginDelete(ctx context.Context, resourceGroupName string, vpnServerConfigurationName string, options *VPNServerConfigurationsClientBeginDeleteOptions) (*armruntime.Poller[VPNServerConfigurationsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, vpnServerConfigurationName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller(resp, client.pl, &armruntime.NewPollerOptions[VPNServerConfigurationsClientDeleteResponse]{
			FinalStateVia: armruntime.FinalStateViaLocation,
		})
	} else {
		return armruntime.NewPollerFromResumeToken[VPNServerConfigurationsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes a VpnServerConfiguration.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *VPNServerConfigurationsClient) deleteOperation(ctx context.Context, resourceGroupName string, vpnServerConfigurationName string, options *VPNServerConfigurationsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, vpnServerConfigurationName, options)
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
func (client *VPNServerConfigurationsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, vpnServerConfigurationName string, options *VPNServerConfigurationsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnServerConfigurations/{vpnServerConfigurationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vpnServerConfigurationName == "" {
		return nil, errors.New("parameter vpnServerConfigurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vpnServerConfigurationName}", url.PathEscape(vpnServerConfigurationName))
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

// Get - Retrieves the details of a VpnServerConfiguration.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name of the VpnServerConfiguration.
// vpnServerConfigurationName - The name of the VpnServerConfiguration being retrieved.
// options - VPNServerConfigurationsClientGetOptions contains the optional parameters for the VPNServerConfigurationsClient.Get
// method.
func (client *VPNServerConfigurationsClient) Get(ctx context.Context, resourceGroupName string, vpnServerConfigurationName string, options *VPNServerConfigurationsClientGetOptions) (VPNServerConfigurationsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, vpnServerConfigurationName, options)
	if err != nil {
		return VPNServerConfigurationsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VPNServerConfigurationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VPNServerConfigurationsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VPNServerConfigurationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, vpnServerConfigurationName string, options *VPNServerConfigurationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnServerConfigurations/{vpnServerConfigurationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vpnServerConfigurationName == "" {
		return nil, errors.New("parameter vpnServerConfigurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vpnServerConfigurationName}", url.PathEscape(vpnServerConfigurationName))
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
func (client *VPNServerConfigurationsClient) getHandleResponse(resp *http.Response) (VPNServerConfigurationsClientGetResponse, error) {
	result := VPNServerConfigurationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VPNServerConfiguration); err != nil {
		return VPNServerConfigurationsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists all the VpnServerConfigurations in a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - VPNServerConfigurationsClientListOptions contains the optional parameters for the VPNServerConfigurationsClient.List
// method.
func (client *VPNServerConfigurationsClient) NewListPager(options *VPNServerConfigurationsClientListOptions) *runtime.Pager[VPNServerConfigurationsClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[VPNServerConfigurationsClientListResponse]{
		More: func(page VPNServerConfigurationsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VPNServerConfigurationsClientListResponse) (VPNServerConfigurationsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return VPNServerConfigurationsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return VPNServerConfigurationsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return VPNServerConfigurationsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *VPNServerConfigurationsClient) listCreateRequest(ctx context.Context, options *VPNServerConfigurationsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/vpnServerConfigurations"
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
func (client *VPNServerConfigurationsClient) listHandleResponse(resp *http.Response) (VPNServerConfigurationsClientListResponse, error) {
	result := VPNServerConfigurationsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListVPNServerConfigurationsResult); err != nil {
		return VPNServerConfigurationsClientListResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Lists all the vpnServerConfigurations in a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name of the VpnServerConfiguration.
// options - VPNServerConfigurationsClientListByResourceGroupOptions contains the optional parameters for the VPNServerConfigurationsClient.ListByResourceGroup
// method.
func (client *VPNServerConfigurationsClient) NewListByResourceGroupPager(resourceGroupName string, options *VPNServerConfigurationsClientListByResourceGroupOptions) *runtime.Pager[VPNServerConfigurationsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PageProcessor[VPNServerConfigurationsClientListByResourceGroupResponse]{
		More: func(page VPNServerConfigurationsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VPNServerConfigurationsClientListByResourceGroupResponse) (VPNServerConfigurationsClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return VPNServerConfigurationsClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return VPNServerConfigurationsClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return VPNServerConfigurationsClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *VPNServerConfigurationsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *VPNServerConfigurationsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnServerConfigurations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
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
func (client *VPNServerConfigurationsClient) listByResourceGroupHandleResponse(resp *http.Response) (VPNServerConfigurationsClientListByResourceGroupResponse, error) {
	result := VPNServerConfigurationsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListVPNServerConfigurationsResult); err != nil {
		return VPNServerConfigurationsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// UpdateTags - Updates VpnServerConfiguration tags.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name of the VpnServerConfiguration.
// vpnServerConfigurationName - The name of the VpnServerConfiguration being updated.
// vpnServerConfigurationParameters - Parameters supplied to update VpnServerConfiguration tags.
// options - VPNServerConfigurationsClientUpdateTagsOptions contains the optional parameters for the VPNServerConfigurationsClient.UpdateTags
// method.
func (client *VPNServerConfigurationsClient) UpdateTags(ctx context.Context, resourceGroupName string, vpnServerConfigurationName string, vpnServerConfigurationParameters TagsObject, options *VPNServerConfigurationsClientUpdateTagsOptions) (VPNServerConfigurationsClientUpdateTagsResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, vpnServerConfigurationName, vpnServerConfigurationParameters, options)
	if err != nil {
		return VPNServerConfigurationsClientUpdateTagsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VPNServerConfigurationsClientUpdateTagsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VPNServerConfigurationsClientUpdateTagsResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateTagsHandleResponse(resp)
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *VPNServerConfigurationsClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, vpnServerConfigurationName string, vpnServerConfigurationParameters TagsObject, options *VPNServerConfigurationsClientUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnServerConfigurations/{vpnServerConfigurationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vpnServerConfigurationName == "" {
		return nil, errors.New("parameter vpnServerConfigurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vpnServerConfigurationName}", url.PathEscape(vpnServerConfigurationName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, vpnServerConfigurationParameters)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *VPNServerConfigurationsClient) updateTagsHandleResponse(resp *http.Response) (VPNServerConfigurationsClientUpdateTagsResponse, error) {
	result := VPNServerConfigurationsClientUpdateTagsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VPNServerConfiguration); err != nil {
		return VPNServerConfigurationsClientUpdateTagsResponse{}, err
	}
	return result, nil
}
