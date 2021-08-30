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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// VPNSiteLinksClient contains the methods for the VPNSiteLinks group.
// Don't use this type directly, use NewVPNSiteLinksClient() instead.
type VPNSiteLinksClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewVPNSiteLinksClient creates a new instance of VPNSiteLinksClient with the specified values.
func NewVPNSiteLinksClient(con *arm.Connection, subscriptionID string) *VPNSiteLinksClient {
	return &VPNSiteLinksClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// Get - Retrieves the details of a VPN site link.
// If the operation fails it returns the *CloudError error type.
func (client *VPNSiteLinksClient) Get(ctx context.Context, resourceGroupName string, vpnSiteName string, vpnSiteLinkName string, options *VPNSiteLinksGetOptions) (VPNSiteLinksGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, vpnSiteName, vpnSiteLinkName, options)
	if err != nil {
		return VPNSiteLinksGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VPNSiteLinksGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VPNSiteLinksGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VPNSiteLinksClient) getCreateRequest(ctx context.Context, resourceGroupName string, vpnSiteName string, vpnSiteLinkName string, options *VPNSiteLinksGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnSites/{vpnSiteName}/vpnSiteLinks/{vpnSiteLinkName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vpnSiteName == "" {
		return nil, errors.New("parameter vpnSiteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vpnSiteName}", url.PathEscape(vpnSiteName))
	if vpnSiteLinkName == "" {
		return nil, errors.New("parameter vpnSiteLinkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vpnSiteLinkName}", url.PathEscape(vpnSiteLinkName))
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
func (client *VPNSiteLinksClient) getHandleResponse(resp *http.Response) (VPNSiteLinksGetResponse, error) {
	result := VPNSiteLinksGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VPNSiteLink); err != nil {
		return VPNSiteLinksGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *VPNSiteLinksClient) getHandleError(resp *http.Response) error {
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

// ListByVPNSite - Lists all the vpnSiteLinks in a resource group for a vpn site.
// If the operation fails it returns the *CloudError error type.
func (client *VPNSiteLinksClient) ListByVPNSite(resourceGroupName string, vpnSiteName string, options *VPNSiteLinksListByVPNSiteOptions) *VPNSiteLinksListByVPNSitePager {
	return &VPNSiteLinksListByVPNSitePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByVPNSiteCreateRequest(ctx, resourceGroupName, vpnSiteName, options)
		},
		advancer: func(ctx context.Context, resp VPNSiteLinksListByVPNSiteResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ListVPNSiteLinksResult.NextLink)
		},
	}
}

// listByVPNSiteCreateRequest creates the ListByVPNSite request.
func (client *VPNSiteLinksClient) listByVPNSiteCreateRequest(ctx context.Context, resourceGroupName string, vpnSiteName string, options *VPNSiteLinksListByVPNSiteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnSites/{vpnSiteName}/vpnSiteLinks"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vpnSiteName == "" {
		return nil, errors.New("parameter vpnSiteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vpnSiteName}", url.PathEscape(vpnSiteName))
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

// listByVPNSiteHandleResponse handles the ListByVPNSite response.
func (client *VPNSiteLinksClient) listByVPNSiteHandleResponse(resp *http.Response) (VPNSiteLinksListByVPNSiteResponse, error) {
	result := VPNSiteLinksListByVPNSiteResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListVPNSiteLinksResult); err != nil {
		return VPNSiteLinksListByVPNSiteResponse{}, err
	}
	return result, nil
}

// listByVPNSiteHandleError handles the ListByVPNSite error response.
func (client *VPNSiteLinksClient) listByVPNSiteHandleError(resp *http.Response) error {
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
