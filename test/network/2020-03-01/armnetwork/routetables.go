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
	"time"
)

// RouteTablesOperations contains the methods for the RouteTables group.
type RouteTablesOperations interface {
	// BeginCreateOrUpdate - Create or updates a route table in a specified resource group.
	BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, routeTableName string, parameters RouteTable) (*RouteTablePollerResponse, error)
	// ResumeCreateOrUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeCreateOrUpdate(token string) (RouteTablePoller, error)
	// BeginDelete - Deletes the specified route table.
	BeginDelete(ctx context.Context, resourceGroupName string, routeTableName string) (*HTTPPollerResponse, error)
	// ResumeDelete - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDelete(token string) (HTTPPoller, error)
	// Get - Gets the specified route table.
	Get(ctx context.Context, resourceGroupName string, routeTableName string, routeTablesGetOptions *RouteTablesGetOptions) (*RouteTableResponse, error)
	// List - Gets all route tables in a resource group.
	List(resourceGroupName string) (RouteTableListResultPager, error)
	// ListAll - Gets all route tables in a subscription.
	ListAll() (RouteTableListResultPager, error)
	// UpdateTags - Updates a route table tags.
	UpdateTags(ctx context.Context, resourceGroupName string, routeTableName string, parameters TagsObject) (*RouteTableResponse, error)
}

// routeTablesOperations implements the RouteTablesOperations interface.
type routeTablesOperations struct {
	*Client
	subscriptionID string
}

// CreateOrUpdate - Create or updates a route table in a specified resource group.
func (client *routeTablesOperations) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, routeTableName string, parameters RouteTable) (*RouteTablePollerResponse, error) {
	req, err := client.createOrUpdateCreateRequest(resourceGroupName, routeTableName, parameters)
	if err != nil {
		return nil, err
	}
	// send the first request to initialize the poller
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.createOrUpdateHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	pt, err := createPollingTracker("routeTablesOperations.CreateOrUpdate", "azure-async-operation", resp, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &routeTablePoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*RouteTableResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *routeTablesOperations) ResumeCreateOrUpdate(token string) (RouteTablePoller, error) {
	pt, err := resumePollingTracker("routeTablesOperations.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &routeTablePoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *routeTablesOperations) createOrUpdateCreateRequest(resourceGroupName string, routeTableName string, parameters RouteTable) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeTables/{routeTableName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{routeTableName}", url.PathEscape(routeTableName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(path.Join(client.u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPut, *u)
	return req, req.MarshalAsJSON(parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *routeTablesOperations) createOrUpdateHandleResponse(resp *azcore.Response) (*RouteTablePollerResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated, http.StatusNoContent) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return &RouteTablePollerResponse{RawResponse: resp.Response}, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *routeTablesOperations) createOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Delete - Deletes the specified route table.
func (client *routeTablesOperations) BeginDelete(ctx context.Context, resourceGroupName string, routeTableName string) (*HTTPPollerResponse, error) {
	req, err := client.deleteCreateRequest(resourceGroupName, routeTableName)
	if err != nil {
		return nil, err
	}
	// send the first request to initialize the poller
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.deleteHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	pt, err := createPollingTracker("routeTablesOperations.Delete", "location", resp, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	poller := &httpPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *routeTablesOperations) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := resumePollingTracker("routeTablesOperations.Delete", token, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *routeTablesOperations) deleteCreateRequest(resourceGroupName string, routeTableName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeTables/{routeTableName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{routeTableName}", url.PathEscape(routeTableName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(path.Join(client.u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodDelete, *u)
	return req, nil
}

// deleteHandleResponse handles the Delete response.
func (client *routeTablesOperations) deleteHandleResponse(resp *azcore.Response) (*HTTPPollerResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return &HTTPPollerResponse{RawResponse: resp.Response}, nil
}

// deleteHandleError handles the Delete error response.
func (client *routeTablesOperations) deleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Get - Gets the specified route table.
func (client *routeTablesOperations) Get(ctx context.Context, resourceGroupName string, routeTableName string, routeTablesGetOptions *RouteTablesGetOptions) (*RouteTableResponse, error) {
	req, err := client.getCreateRequest(resourceGroupName, routeTableName, routeTablesGetOptions)
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
func (client *routeTablesOperations) getCreateRequest(resourceGroupName string, routeTableName string, routeTablesGetOptions *RouteTablesGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeTables/{routeTableName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{routeTableName}", url.PathEscape(routeTableName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(path.Join(client.u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	if routeTablesGetOptions != nil && routeTablesGetOptions.Expand != nil {
		query.Set("$expand", *routeTablesGetOptions.Expand)
	}
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *routeTablesOperations) getHandleResponse(resp *azcore.Response) (*RouteTableResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getHandleError(resp)
	}
	result := RouteTableResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.RouteTable)
}

// getHandleError handles the Get error response.
func (client *routeTablesOperations) getHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// List - Gets all route tables in a resource group.
func (client *routeTablesOperations) List(resourceGroupName string) (RouteTableListResultPager, error) {
	req, err := client.listCreateRequest(resourceGroupName)
	if err != nil {
		return nil, err
	}
	return &routeTableListResultPager{
		pipeline:  client.p,
		request:   req,
		responder: client.listHandleResponse,
		advancer: func(resp *RouteTableListResultResponse) (*azcore.Request, error) {
			u, err := url.Parse(*resp.RouteTableListResult.NextLink)
			if err != nil {
				return nil, fmt.Errorf("invalid NextLink: %w", err)
			}
			if u.Scheme == "" {
				return nil, fmt.Errorf("no scheme detected in NextLink %s", *resp.RouteTableListResult.NextLink)
			}
			return azcore.NewRequest(http.MethodGet, *u), nil
		},
	}, nil
}

// listCreateRequest creates the List request.
func (client *routeTablesOperations) listCreateRequest(resourceGroupName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeTables"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
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
func (client *routeTablesOperations) listHandleResponse(resp *azcore.Response) (*RouteTableListResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.listHandleError(resp)
	}
	result := RouteTableListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.RouteTableListResult)
}

// listHandleError handles the List error response.
func (client *routeTablesOperations) listHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// ListAll - Gets all route tables in a subscription.
func (client *routeTablesOperations) ListAll() (RouteTableListResultPager, error) {
	req, err := client.listAllCreateRequest()
	if err != nil {
		return nil, err
	}
	return &routeTableListResultPager{
		pipeline:  client.p,
		request:   req,
		responder: client.listAllHandleResponse,
		advancer: func(resp *RouteTableListResultResponse) (*azcore.Request, error) {
			u, err := url.Parse(*resp.RouteTableListResult.NextLink)
			if err != nil {
				return nil, fmt.Errorf("invalid NextLink: %w", err)
			}
			if u.Scheme == "" {
				return nil, fmt.Errorf("no scheme detected in NextLink %s", *resp.RouteTableListResult.NextLink)
			}
			return azcore.NewRequest(http.MethodGet, *u), nil
		},
	}, nil
}

// listAllCreateRequest creates the ListAll request.
func (client *routeTablesOperations) listAllCreateRequest() (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/routeTables"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
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

// listAllHandleResponse handles the ListAll response.
func (client *routeTablesOperations) listAllHandleResponse(resp *azcore.Response) (*RouteTableListResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.listAllHandleError(resp)
	}
	result := RouteTableListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.RouteTableListResult)
}

// listAllHandleError handles the ListAll error response.
func (client *routeTablesOperations) listAllHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// UpdateTags - Updates a route table tags.
func (client *routeTablesOperations) UpdateTags(ctx context.Context, resourceGroupName string, routeTableName string, parameters TagsObject) (*RouteTableResponse, error) {
	req, err := client.updateTagsCreateRequest(resourceGroupName, routeTableName, parameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.updateTagsHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *routeTablesOperations) updateTagsCreateRequest(resourceGroupName string, routeTableName string, parameters TagsObject) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeTables/{routeTableName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{routeTableName}", url.PathEscape(routeTableName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(path.Join(client.u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPatch, *u)
	return req, req.MarshalAsJSON(parameters)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *routeTablesOperations) updateTagsHandleResponse(resp *azcore.Response) (*RouteTableResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.updateTagsHandleError(resp)
	}
	result := RouteTableResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.RouteTable)
}

// updateTagsHandleError handles the UpdateTags error response.
func (client *routeTablesOperations) updateTagsHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
