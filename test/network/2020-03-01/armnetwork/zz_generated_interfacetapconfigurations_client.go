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

// InterfaceTapConfigurationsClient contains the methods for the NetworkInterfaceTapConfigurations group.
// Don't use this type directly, use NewInterfaceTapConfigurationsClient() instead.
type InterfaceTapConfigurationsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewInterfaceTapConfigurationsClient creates a new instance of InterfaceTapConfigurationsClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewInterfaceTapConfigurationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*InterfaceTapConfigurationsClient, error) {
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
	client := &InterfaceTapConfigurationsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a Tap configuration in the specified NetworkInterface.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-03-01
// resourceGroupName - The name of the resource group.
// networkInterfaceName - The name of the network interface.
// tapConfigurationName - The name of the tap configuration.
// tapConfigurationParameters - Parameters supplied to the create or update tap configuration operation.
// options - InterfaceTapConfigurationsClientBeginCreateOrUpdateOptions contains the optional parameters for the InterfaceTapConfigurationsClient.BeginCreateOrUpdate
// method.
func (client *InterfaceTapConfigurationsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, networkInterfaceName string, tapConfigurationName string, tapConfigurationParameters InterfaceTapConfiguration, options *InterfaceTapConfigurationsClientBeginCreateOrUpdateOptions) (*armruntime.Poller[InterfaceTapConfigurationsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, networkInterfaceName, tapConfigurationName, tapConfigurationParameters, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller(resp, client.pl, &armruntime.NewPollerOptions[InterfaceTapConfigurationsClientCreateOrUpdateResponse]{
			FinalStateVia: armruntime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return armruntime.NewPollerFromResumeToken[InterfaceTapConfigurationsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates or updates a Tap configuration in the specified NetworkInterface.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-03-01
func (client *InterfaceTapConfigurationsClient) createOrUpdate(ctx context.Context, resourceGroupName string, networkInterfaceName string, tapConfigurationName string, tapConfigurationParameters InterfaceTapConfiguration, options *InterfaceTapConfigurationsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, networkInterfaceName, tapConfigurationName, tapConfigurationParameters, options)
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
func (client *InterfaceTapConfigurationsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, networkInterfaceName string, tapConfigurationName string, tapConfigurationParameters InterfaceTapConfiguration, options *InterfaceTapConfigurationsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}/tapConfigurations/{tapConfigurationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkInterfaceName == "" {
		return nil, errors.New("parameter networkInterfaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkInterfaceName}", url.PathEscape(networkInterfaceName))
	if tapConfigurationName == "" {
		return nil, errors.New("parameter tapConfigurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tapConfigurationName}", url.PathEscape(tapConfigurationName))
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
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, tapConfigurationParameters)
}

// BeginDelete - Deletes the specified tap configuration from the NetworkInterface.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-03-01
// resourceGroupName - The name of the resource group.
// networkInterfaceName - The name of the network interface.
// tapConfigurationName - The name of the tap configuration.
// options - InterfaceTapConfigurationsClientBeginDeleteOptions contains the optional parameters for the InterfaceTapConfigurationsClient.BeginDelete
// method.
func (client *InterfaceTapConfigurationsClient) BeginDelete(ctx context.Context, resourceGroupName string, networkInterfaceName string, tapConfigurationName string, options *InterfaceTapConfigurationsClientBeginDeleteOptions) (*armruntime.Poller[InterfaceTapConfigurationsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, networkInterfaceName, tapConfigurationName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller(resp, client.pl, &armruntime.NewPollerOptions[InterfaceTapConfigurationsClientDeleteResponse]{
			FinalStateVia: armruntime.FinalStateViaLocation,
		})
	} else {
		return armruntime.NewPollerFromResumeToken[InterfaceTapConfigurationsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes the specified tap configuration from the NetworkInterface.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-03-01
func (client *InterfaceTapConfigurationsClient) deleteOperation(ctx context.Context, resourceGroupName string, networkInterfaceName string, tapConfigurationName string, options *InterfaceTapConfigurationsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, networkInterfaceName, tapConfigurationName, options)
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
func (client *InterfaceTapConfigurationsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, networkInterfaceName string, tapConfigurationName string, options *InterfaceTapConfigurationsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}/tapConfigurations/{tapConfigurationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkInterfaceName == "" {
		return nil, errors.New("parameter networkInterfaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkInterfaceName}", url.PathEscape(networkInterfaceName))
	if tapConfigurationName == "" {
		return nil, errors.New("parameter tapConfigurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tapConfigurationName}", url.PathEscape(tapConfigurationName))
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
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get the specified tap configuration on a network interface.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-03-01
// resourceGroupName - The name of the resource group.
// networkInterfaceName - The name of the network interface.
// tapConfigurationName - The name of the tap configuration.
// options - InterfaceTapConfigurationsClientGetOptions contains the optional parameters for the InterfaceTapConfigurationsClient.Get
// method.
func (client *InterfaceTapConfigurationsClient) Get(ctx context.Context, resourceGroupName string, networkInterfaceName string, tapConfigurationName string, options *InterfaceTapConfigurationsClientGetOptions) (InterfaceTapConfigurationsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, networkInterfaceName, tapConfigurationName, options)
	if err != nil {
		return InterfaceTapConfigurationsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return InterfaceTapConfigurationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return InterfaceTapConfigurationsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *InterfaceTapConfigurationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, networkInterfaceName string, tapConfigurationName string, options *InterfaceTapConfigurationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}/tapConfigurations/{tapConfigurationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkInterfaceName == "" {
		return nil, errors.New("parameter networkInterfaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkInterfaceName}", url.PathEscape(networkInterfaceName))
	if tapConfigurationName == "" {
		return nil, errors.New("parameter tapConfigurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tapConfigurationName}", url.PathEscape(tapConfigurationName))
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
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *InterfaceTapConfigurationsClient) getHandleResponse(resp *http.Response) (InterfaceTapConfigurationsClientGetResponse, error) {
	result := InterfaceTapConfigurationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.InterfaceTapConfiguration); err != nil {
		return InterfaceTapConfigurationsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Get all Tap configurations in a network interface.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-03-01
// resourceGroupName - The name of the resource group.
// networkInterfaceName - The name of the network interface.
// options - InterfaceTapConfigurationsClientListOptions contains the optional parameters for the InterfaceTapConfigurationsClient.List
// method.
func (client *InterfaceTapConfigurationsClient) NewListPager(resourceGroupName string, networkInterfaceName string, options *InterfaceTapConfigurationsClientListOptions) *runtime.Pager[InterfaceTapConfigurationsClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[InterfaceTapConfigurationsClientListResponse]{
		More: func(page InterfaceTapConfigurationsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *InterfaceTapConfigurationsClientListResponse) (InterfaceTapConfigurationsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, networkInterfaceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return InterfaceTapConfigurationsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return InterfaceTapConfigurationsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return InterfaceTapConfigurationsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *InterfaceTapConfigurationsClient) listCreateRequest(ctx context.Context, resourceGroupName string, networkInterfaceName string, options *InterfaceTapConfigurationsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}/tapConfigurations"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkInterfaceName == "" {
		return nil, errors.New("parameter networkInterfaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkInterfaceName}", url.PathEscape(networkInterfaceName))
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
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *InterfaceTapConfigurationsClient) listHandleResponse(resp *http.Response) (InterfaceTapConfigurationsClientListResponse, error) {
	result := InterfaceTapConfigurationsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.InterfaceTapConfigurationListResult); err != nil {
		return InterfaceTapConfigurationsClientListResponse{}, err
	}
	return result, nil
}
