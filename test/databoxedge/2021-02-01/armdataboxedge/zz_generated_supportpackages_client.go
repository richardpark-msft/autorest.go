//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdataboxedge

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

// SupportPackagesClient contains the methods for the SupportPackages group.
// Don't use this type directly, use NewSupportPackagesClient() instead.
type SupportPackagesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewSupportPackagesClient creates a new instance of SupportPackagesClient with the specified values.
// subscriptionID - The subscription ID.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewSupportPackagesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *SupportPackagesClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &SupportPackagesClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// BeginTriggerSupportPackage - Triggers support package on the device
// If the operation fails it returns an *azcore.ResponseError type.
// deviceName - The device name.
// resourceGroupName - The resource group name.
// triggerSupportPackageRequest - The trigger support package request object
// options - SupportPackagesClientBeginTriggerSupportPackageOptions contains the optional parameters for the SupportPackagesClient.BeginTriggerSupportPackage
// method.
func (client *SupportPackagesClient) BeginTriggerSupportPackage(ctx context.Context, deviceName string, resourceGroupName string, triggerSupportPackageRequest TriggerSupportPackageRequest, options *SupportPackagesClientBeginTriggerSupportPackageOptions) (SupportPackagesClientTriggerSupportPackagePollerResponse, error) {
	resp, err := client.triggerSupportPackage(ctx, deviceName, resourceGroupName, triggerSupportPackageRequest, options)
	if err != nil {
		return SupportPackagesClientTriggerSupportPackagePollerResponse{}, err
	}
	result := SupportPackagesClientTriggerSupportPackagePollerResponse{}
	pt, err := armruntime.NewPoller("SupportPackagesClient.TriggerSupportPackage", "", resp, client.pl)
	if err != nil {
		return SupportPackagesClientTriggerSupportPackagePollerResponse{}, err
	}
	result.Poller = &SupportPackagesClientTriggerSupportPackagePoller{
		pt: pt,
	}
	return result, nil
}

// TriggerSupportPackage - Triggers support package on the device
// If the operation fails it returns an *azcore.ResponseError type.
func (client *SupportPackagesClient) triggerSupportPackage(ctx context.Context, deviceName string, resourceGroupName string, triggerSupportPackageRequest TriggerSupportPackageRequest, options *SupportPackagesClientBeginTriggerSupportPackageOptions) (*http.Response, error) {
	req, err := client.triggerSupportPackageCreateRequest(ctx, deviceName, resourceGroupName, triggerSupportPackageRequest, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// triggerSupportPackageCreateRequest creates the TriggerSupportPackage request.
func (client *SupportPackagesClient) triggerSupportPackageCreateRequest(ctx context.Context, deviceName string, resourceGroupName string, triggerSupportPackageRequest TriggerSupportPackageRequest, options *SupportPackagesClientBeginTriggerSupportPackageOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/triggerSupportPackage"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, triggerSupportPackageRequest)
}
