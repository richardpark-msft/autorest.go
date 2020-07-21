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

// ExpressRouteLinksOperations contains the methods for the ExpressRouteLinks group.
type ExpressRouteLinksOperations interface {
	// Get - Retrieves the specified ExpressRouteLink resource.
	Get(ctx context.Context, resourceGroupName string, expressRoutePortName string, linkName string) (*ExpressRouteLinkResponse, error)
	// List - Retrieve the ExpressRouteLink sub-resources of the specified ExpressRoutePort resource.
	List(resourceGroupName string, expressRoutePortName string) (ExpressRouteLinkListResultPager, error)
}

// expressRouteLinksOperations implements the ExpressRouteLinksOperations interface.
type expressRouteLinksOperations struct {
	*Client
	subscriptionID string
}

// Get - Retrieves the specified ExpressRouteLink resource.
func (client *expressRouteLinksOperations) Get(ctx context.Context, resourceGroupName string, expressRoutePortName string, linkName string) (*ExpressRouteLinkResponse, error) {
	req, err := client.getCreateRequest(resourceGroupName, expressRoutePortName, linkName)
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
func (client *expressRouteLinksOperations) getCreateRequest(resourceGroupName string, expressRoutePortName string, linkName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ExpressRoutePorts/{expressRoutePortName}/links/{linkName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{expressRoutePortName}", url.PathEscape(expressRoutePortName))
	urlPath = strings.ReplaceAll(urlPath, "{linkName}", url.PathEscape(linkName))
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
func (client *expressRouteLinksOperations) getHandleResponse(resp *azcore.Response) (*ExpressRouteLinkResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getHandleError(resp)
	}
	result := ExpressRouteLinkResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ExpressRouteLink)
}

// getHandleError handles the Get error response.
func (client *expressRouteLinksOperations) getHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// List - Retrieve the ExpressRouteLink sub-resources of the specified ExpressRoutePort resource.
func (client *expressRouteLinksOperations) List(resourceGroupName string, expressRoutePortName string) (ExpressRouteLinkListResultPager, error) {
	req, err := client.listCreateRequest(resourceGroupName, expressRoutePortName)
	if err != nil {
		return nil, err
	}
	return &expressRouteLinkListResultPager{
		pipeline:  client.p,
		request:   req,
		responder: client.listHandleResponse,
		advancer: func(resp *ExpressRouteLinkListResultResponse) (*azcore.Request, error) {
			u, err := url.Parse(*resp.ExpressRouteLinkListResult.NextLink)
			if err != nil {
				return nil, fmt.Errorf("invalid NextLink: %w", err)
			}
			if u.Scheme == "" {
				return nil, fmt.Errorf("no scheme detected in NextLink %s", *resp.ExpressRouteLinkListResult.NextLink)
			}
			return azcore.NewRequest(http.MethodGet, *u), nil
		},
	}, nil
}

// listCreateRequest creates the List request.
func (client *expressRouteLinksOperations) listCreateRequest(resourceGroupName string, expressRoutePortName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ExpressRoutePorts/{expressRoutePortName}/links"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{expressRoutePortName}", url.PathEscape(expressRoutePortName))
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

// listHandleResponse handles the List response.
func (client *expressRouteLinksOperations) listHandleResponse(resp *azcore.Response) (*ExpressRouteLinkListResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.listHandleError(resp)
	}
	result := ExpressRouteLinkListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ExpressRouteLinkListResult)
}

// listHandleError handles the List error response.
func (client *expressRouteLinksOperations) listHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
