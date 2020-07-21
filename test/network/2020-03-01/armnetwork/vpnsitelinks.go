// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"path"
	"strings"
)

// VpnSiteLinksOperations contains the methods for the VpnSiteLinks group.
type VpnSiteLinksOperations interface {
	// Get - Retrieves the details of a VPN site link.
	Get(ctx context.Context, resourceGroupName string, vpnSiteName string, vpnSiteLinkName string) (*VpnSiteLinkResponse, error)
	// ListByVpnSite - Lists all the vpnSiteLinks in a resource group for a vpn site.
	ListByVpnSite(resourceGroupName string, vpnSiteName string) (ListVpnSiteLinksResultPager, error)
}

// vpnSiteLinksOperations implements the VpnSiteLinksOperations interface.
type vpnSiteLinksOperations struct {
	*Client
	subscriptionID string
}

// Get - Retrieves the details of a VPN site link.
func (client *vpnSiteLinksOperations) Get(ctx context.Context, resourceGroupName string, vpnSiteName string, vpnSiteLinkName string) (*VpnSiteLinkResponse, error) {
	req, err := client.getCreateRequest(resourceGroupName, vpnSiteName, vpnSiteLinkName)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getCreateRequest creates the Get request.
func (client *vpnSiteLinksOperations) getCreateRequest(resourceGroupName string, vpnSiteName string, vpnSiteLinkName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnSites/{vpnSiteName}/vpnSiteLinks/{vpnSiteLinkName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{vpnSiteName}", url.PathEscape(vpnSiteName))
	urlPath = strings.ReplaceAll(urlPath, "{vpnSiteLinkName}", url.PathEscape(vpnSiteLinkName))
	u, err := client.u.Parse(path.Join(client.u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *vpnSiteLinksOperations) getHandleResponse(resp *azcore.Response) (*VpnSiteLinkResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getHandleError(resp)
	}
	result := VpnSiteLinkResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VpnSiteLink)
}

// getHandleError handles the Get error response.
func (client *vpnSiteLinksOperations) getHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// ListByVpnSite - Lists all the vpnSiteLinks in a resource group for a vpn site.
func (client *vpnSiteLinksOperations) ListByVpnSite(resourceGroupName string, vpnSiteName string) (ListVpnSiteLinksResultPager, error) {
	req, err := client.listByVpnSiteCreateRequest(resourceGroupName, vpnSiteName)
	if err != nil {
		return nil, err
	}
	return &listVpnSiteLinksResultPager{
		pipeline:  client.p,
		request:   req,
		responder: client.listByVpnSiteHandleResponse,
		advancer: func(resp *ListVpnSiteLinksResultResponse) (*azcore.Request, error) {
			u, err := url.Parse(*resp.ListVpnSiteLinksResult.NextLink)
			if err != nil {
				return nil, fmt.Errorf("invalid NextLink: %w", err)
			}
			if u.Scheme == "" {
				return nil, fmt.Errorf("no scheme detected in NextLink %s", *resp.ListVpnSiteLinksResult.NextLink)
			}
			return azcore.NewRequest(http.MethodGet, *u), nil
		},
	}, nil
}

// listByVpnSiteCreateRequest creates the ListByVpnSite request.
func (client *vpnSiteLinksOperations) listByVpnSiteCreateRequest(resourceGroupName string, vpnSiteName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnSites/{vpnSiteName}/vpnSiteLinks"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{vpnSiteName}", url.PathEscape(vpnSiteName))
	u, err := client.u.Parse(path.Join(client.u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// listByVpnSiteHandleResponse handles the ListByVpnSite response.
func (client *vpnSiteLinksOperations) listByVpnSiteHandleResponse(resp *azcore.Response) (*ListVpnSiteLinksResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.listByVpnSiteHandleError(resp)
	}
	result := ListVpnSiteLinksResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ListVpnSiteLinksResult)
}

// listByVpnSiteHandleError handles the ListByVpnSite error response.
func (client *vpnSiteLinksOperations) listByVpnSiteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
