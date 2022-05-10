//go:build go1.18
// +build go1.18

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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// DiagnosticSettingsClient contains the methods for the DiagnosticSettings group.
// Don't use this type directly, use NewDiagnosticSettingsClient() instead.
type DiagnosticSettingsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewDiagnosticSettingsClient creates a new instance of DiagnosticSettingsClient with the specified values.
// subscriptionID - The subscription ID.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewDiagnosticSettingsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DiagnosticSettingsClient, error) {
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
	client := &DiagnosticSettingsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// GetDiagnosticProactiveLogCollectionSettings - Gets the proactive log collection settings of the specified Data Box Edge/Data
// Box Gateway device.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-02-01
// deviceName - The device name.
// resourceGroupName - The resource group name.
// options - DiagnosticSettingsClientGetDiagnosticProactiveLogCollectionSettingsOptions contains the optional parameters for
// the DiagnosticSettingsClient.GetDiagnosticProactiveLogCollectionSettings method.
func (client *DiagnosticSettingsClient) GetDiagnosticProactiveLogCollectionSettings(ctx context.Context, deviceName string, resourceGroupName string, options *DiagnosticSettingsClientGetDiagnosticProactiveLogCollectionSettingsOptions) (DiagnosticSettingsClientGetDiagnosticProactiveLogCollectionSettingsResponse, error) {
	req, err := client.getDiagnosticProactiveLogCollectionSettingsCreateRequest(ctx, deviceName, resourceGroupName, options)
	if err != nil {
		return DiagnosticSettingsClientGetDiagnosticProactiveLogCollectionSettingsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DiagnosticSettingsClientGetDiagnosticProactiveLogCollectionSettingsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DiagnosticSettingsClientGetDiagnosticProactiveLogCollectionSettingsResponse{}, runtime.NewResponseError(resp)
	}
	return client.getDiagnosticProactiveLogCollectionSettingsHandleResponse(resp)
}

// getDiagnosticProactiveLogCollectionSettingsCreateRequest creates the GetDiagnosticProactiveLogCollectionSettings request.
func (client *DiagnosticSettingsClient) getDiagnosticProactiveLogCollectionSettingsCreateRequest(ctx context.Context, deviceName string, resourceGroupName string, options *DiagnosticSettingsClientGetDiagnosticProactiveLogCollectionSettingsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/diagnosticProactiveLogCollectionSettings/default"
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
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
	reqQP.Set("api-version", "2021-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getDiagnosticProactiveLogCollectionSettingsHandleResponse handles the GetDiagnosticProactiveLogCollectionSettings response.
func (client *DiagnosticSettingsClient) getDiagnosticProactiveLogCollectionSettingsHandleResponse(resp *http.Response) (DiagnosticSettingsClientGetDiagnosticProactiveLogCollectionSettingsResponse, error) {
	result := DiagnosticSettingsClientGetDiagnosticProactiveLogCollectionSettingsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DiagnosticProactiveLogCollectionSettings); err != nil {
		return DiagnosticSettingsClientGetDiagnosticProactiveLogCollectionSettingsResponse{}, err
	}
	return result, nil
}

// GetDiagnosticRemoteSupportSettings - Gets the diagnostic remote support settings of the specified Data Box Edge/Data Box
// Gateway device.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-02-01
// deviceName - The device name.
// resourceGroupName - The resource group name.
// options - DiagnosticSettingsClientGetDiagnosticRemoteSupportSettingsOptions contains the optional parameters for the DiagnosticSettingsClient.GetDiagnosticRemoteSupportSettings
// method.
func (client *DiagnosticSettingsClient) GetDiagnosticRemoteSupportSettings(ctx context.Context, deviceName string, resourceGroupName string, options *DiagnosticSettingsClientGetDiagnosticRemoteSupportSettingsOptions) (DiagnosticSettingsClientGetDiagnosticRemoteSupportSettingsResponse, error) {
	req, err := client.getDiagnosticRemoteSupportSettingsCreateRequest(ctx, deviceName, resourceGroupName, options)
	if err != nil {
		return DiagnosticSettingsClientGetDiagnosticRemoteSupportSettingsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DiagnosticSettingsClientGetDiagnosticRemoteSupportSettingsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DiagnosticSettingsClientGetDiagnosticRemoteSupportSettingsResponse{}, runtime.NewResponseError(resp)
	}
	return client.getDiagnosticRemoteSupportSettingsHandleResponse(resp)
}

// getDiagnosticRemoteSupportSettingsCreateRequest creates the GetDiagnosticRemoteSupportSettings request.
func (client *DiagnosticSettingsClient) getDiagnosticRemoteSupportSettingsCreateRequest(ctx context.Context, deviceName string, resourceGroupName string, options *DiagnosticSettingsClientGetDiagnosticRemoteSupportSettingsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/diagnosticRemoteSupportSettings/default"
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
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
	reqQP.Set("api-version", "2021-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getDiagnosticRemoteSupportSettingsHandleResponse handles the GetDiagnosticRemoteSupportSettings response.
func (client *DiagnosticSettingsClient) getDiagnosticRemoteSupportSettingsHandleResponse(resp *http.Response) (DiagnosticSettingsClientGetDiagnosticRemoteSupportSettingsResponse, error) {
	result := DiagnosticSettingsClientGetDiagnosticRemoteSupportSettingsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DiagnosticRemoteSupportSettings); err != nil {
		return DiagnosticSettingsClientGetDiagnosticRemoteSupportSettingsResponse{}, err
	}
	return result, nil
}

// BeginUpdateDiagnosticProactiveLogCollectionSettings - Updates the proactive log collection settings on a Data Box Edge/Data
// Box Gateway device.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-02-01
// deviceName - The device name.
// resourceGroupName - The resource group name.
// proactiveLogCollectionSettings - The proactive log collection settings.
// options - DiagnosticSettingsClientBeginUpdateDiagnosticProactiveLogCollectionSettingsOptions contains the optional parameters
// for the DiagnosticSettingsClient.BeginUpdateDiagnosticProactiveLogCollectionSettings method.
func (client *DiagnosticSettingsClient) BeginUpdateDiagnosticProactiveLogCollectionSettings(ctx context.Context, deviceName string, resourceGroupName string, proactiveLogCollectionSettings DiagnosticProactiveLogCollectionSettings, options *DiagnosticSettingsClientBeginUpdateDiagnosticProactiveLogCollectionSettingsOptions) (*armruntime.Poller[DiagnosticSettingsClientUpdateDiagnosticProactiveLogCollectionSettingsResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.updateDiagnosticProactiveLogCollectionSettings(ctx, deviceName, resourceGroupName, proactiveLogCollectionSettings, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[DiagnosticSettingsClientUpdateDiagnosticProactiveLogCollectionSettingsResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[DiagnosticSettingsClientUpdateDiagnosticProactiveLogCollectionSettingsResponse](options.ResumeToken, client.pl, nil)
	}
}

// UpdateDiagnosticProactiveLogCollectionSettings - Updates the proactive log collection settings on a Data Box Edge/Data
// Box Gateway device.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-02-01
func (client *DiagnosticSettingsClient) updateDiagnosticProactiveLogCollectionSettings(ctx context.Context, deviceName string, resourceGroupName string, proactiveLogCollectionSettings DiagnosticProactiveLogCollectionSettings, options *DiagnosticSettingsClientBeginUpdateDiagnosticProactiveLogCollectionSettingsOptions) (*http.Response, error) {
	req, err := client.updateDiagnosticProactiveLogCollectionSettingsCreateRequest(ctx, deviceName, resourceGroupName, proactiveLogCollectionSettings, options)
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

// updateDiagnosticProactiveLogCollectionSettingsCreateRequest creates the UpdateDiagnosticProactiveLogCollectionSettings request.
func (client *DiagnosticSettingsClient) updateDiagnosticProactiveLogCollectionSettingsCreateRequest(ctx context.Context, deviceName string, resourceGroupName string, proactiveLogCollectionSettings DiagnosticProactiveLogCollectionSettings, options *DiagnosticSettingsClientBeginUpdateDiagnosticProactiveLogCollectionSettingsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/diagnosticProactiveLogCollectionSettings/default"
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, proactiveLogCollectionSettings)
}

// BeginUpdateDiagnosticRemoteSupportSettings - Updates the diagnostic remote support settings on a Data Box Edge/Data Box
// Gateway device.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-02-01
// deviceName - The device name.
// resourceGroupName - The resource group name.
// diagnosticRemoteSupportSettings - The diagnostic remote support settings.
// options - DiagnosticSettingsClientBeginUpdateDiagnosticRemoteSupportSettingsOptions contains the optional parameters for
// the DiagnosticSettingsClient.BeginUpdateDiagnosticRemoteSupportSettings method.
func (client *DiagnosticSettingsClient) BeginUpdateDiagnosticRemoteSupportSettings(ctx context.Context, deviceName string, resourceGroupName string, diagnosticRemoteSupportSettings DiagnosticRemoteSupportSettings, options *DiagnosticSettingsClientBeginUpdateDiagnosticRemoteSupportSettingsOptions) (*armruntime.Poller[DiagnosticSettingsClientUpdateDiagnosticRemoteSupportSettingsResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.updateDiagnosticRemoteSupportSettings(ctx, deviceName, resourceGroupName, diagnosticRemoteSupportSettings, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[DiagnosticSettingsClientUpdateDiagnosticRemoteSupportSettingsResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[DiagnosticSettingsClientUpdateDiagnosticRemoteSupportSettingsResponse](options.ResumeToken, client.pl, nil)
	}
}

// UpdateDiagnosticRemoteSupportSettings - Updates the diagnostic remote support settings on a Data Box Edge/Data Box Gateway
// device.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-02-01
func (client *DiagnosticSettingsClient) updateDiagnosticRemoteSupportSettings(ctx context.Context, deviceName string, resourceGroupName string, diagnosticRemoteSupportSettings DiagnosticRemoteSupportSettings, options *DiagnosticSettingsClientBeginUpdateDiagnosticRemoteSupportSettingsOptions) (*http.Response, error) {
	req, err := client.updateDiagnosticRemoteSupportSettingsCreateRequest(ctx, deviceName, resourceGroupName, diagnosticRemoteSupportSettings, options)
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

// updateDiagnosticRemoteSupportSettingsCreateRequest creates the UpdateDiagnosticRemoteSupportSettings request.
func (client *DiagnosticSettingsClient) updateDiagnosticRemoteSupportSettingsCreateRequest(ctx context.Context, deviceName string, resourceGroupName string, diagnosticRemoteSupportSettings DiagnosticRemoteSupportSettings, options *DiagnosticSettingsClientBeginUpdateDiagnosticRemoteSupportSettingsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/diagnosticRemoteSupportSettings/default"
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, diagnosticRemoteSupportSettings)
}
