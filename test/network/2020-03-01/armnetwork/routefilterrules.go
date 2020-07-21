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

// RouteFilterRulesOperations contains the methods for the RouteFilterRules group.
type RouteFilterRulesOperations interface {
	// BeginCreateOrUpdate - Creates or updates a route in the specified route filter.
	BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, routeFilterName string, ruleName string, routeFilterRuleParameters RouteFilterRule) (*RouteFilterRulePollerResponse, error)
	// ResumeCreateOrUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeCreateOrUpdate(token string) (RouteFilterRulePoller, error)
	// BeginDelete - Deletes the specified rule from a route filter.
	BeginDelete(ctx context.Context, resourceGroupName string, routeFilterName string, ruleName string) (*HTTPPollerResponse, error)
	// ResumeDelete - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDelete(token string) (HTTPPoller, error)
	// Get - Gets the specified rule from a route filter.
	Get(ctx context.Context, resourceGroupName string, routeFilterName string, ruleName string) (*RouteFilterRuleResponse, error)
	// ListByRouteFilter - Gets all RouteFilterRules in a route filter.
	ListByRouteFilter(resourceGroupName string, routeFilterName string) (RouteFilterRuleListResultPager, error)
}

// routeFilterRulesOperations implements the RouteFilterRulesOperations interface.
type routeFilterRulesOperations struct {
	*Client
	subscriptionID string
}

// CreateOrUpdate - Creates or updates a route in the specified route filter.
func (client *routeFilterRulesOperations) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, routeFilterName string, ruleName string, routeFilterRuleParameters RouteFilterRule) (*RouteFilterRulePollerResponse, error) {
	req, err := client.createOrUpdateCreateRequest(resourceGroupName, routeFilterName, ruleName, routeFilterRuleParameters)
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
	pt, err := createPollingTracker("routeFilterRulesOperations.CreateOrUpdate", "azure-async-operation", resp, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &routeFilterRulePoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*RouteFilterRuleResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *routeFilterRulesOperations) ResumeCreateOrUpdate(token string) (RouteFilterRulePoller, error) {
	pt, err := resumePollingTracker("routeFilterRulesOperations.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &routeFilterRulePoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *routeFilterRulesOperations) createOrUpdateCreateRequest(resourceGroupName string, routeFilterName string, ruleName string, routeFilterRuleParameters RouteFilterRule) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}/routeFilterRules/{ruleName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{routeFilterName}", url.PathEscape(routeFilterName))
	urlPath = strings.ReplaceAll(urlPath, "{ruleName}", url.PathEscape(ruleName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(path.Join(client.u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPut, *u)
	return req, req.MarshalAsJSON(routeFilterRuleParameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *routeFilterRulesOperations) createOrUpdateHandleResponse(resp *azcore.Response) (*RouteFilterRulePollerResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated, http.StatusNoContent) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return &RouteFilterRulePollerResponse{RawResponse: resp.Response}, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *routeFilterRulesOperations) createOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Delete - Deletes the specified rule from a route filter.
func (client *routeFilterRulesOperations) BeginDelete(ctx context.Context, resourceGroupName string, routeFilterName string, ruleName string) (*HTTPPollerResponse, error) {
	req, err := client.deleteCreateRequest(resourceGroupName, routeFilterName, ruleName)
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
	pt, err := createPollingTracker("routeFilterRulesOperations.Delete", "location", resp, client.deleteHandleError)
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

func (client *routeFilterRulesOperations) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := resumePollingTracker("routeFilterRulesOperations.Delete", token, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *routeFilterRulesOperations) deleteCreateRequest(resourceGroupName string, routeFilterName string, ruleName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}/routeFilterRules/{ruleName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{routeFilterName}", url.PathEscape(routeFilterName))
	urlPath = strings.ReplaceAll(urlPath, "{ruleName}", url.PathEscape(ruleName))
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
func (client *routeFilterRulesOperations) deleteHandleResponse(resp *azcore.Response) (*HTTPPollerResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return &HTTPPollerResponse{RawResponse: resp.Response}, nil
}

// deleteHandleError handles the Delete error response.
func (client *routeFilterRulesOperations) deleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Get - Gets the specified rule from a route filter.
func (client *routeFilterRulesOperations) Get(ctx context.Context, resourceGroupName string, routeFilterName string, ruleName string) (*RouteFilterRuleResponse, error) {
	req, err := client.getCreateRequest(resourceGroupName, routeFilterName, ruleName)
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
func (client *routeFilterRulesOperations) getCreateRequest(resourceGroupName string, routeFilterName string, ruleName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}/routeFilterRules/{ruleName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{routeFilterName}", url.PathEscape(routeFilterName))
	urlPath = strings.ReplaceAll(urlPath, "{ruleName}", url.PathEscape(ruleName))
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

// getHandleResponse handles the Get response.
func (client *routeFilterRulesOperations) getHandleResponse(resp *azcore.Response) (*RouteFilterRuleResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getHandleError(resp)
	}
	result := RouteFilterRuleResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.RouteFilterRule)
}

// getHandleError handles the Get error response.
func (client *routeFilterRulesOperations) getHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// ListByRouteFilter - Gets all RouteFilterRules in a route filter.
func (client *routeFilterRulesOperations) ListByRouteFilter(resourceGroupName string, routeFilterName string) (RouteFilterRuleListResultPager, error) {
	req, err := client.listByRouteFilterCreateRequest(resourceGroupName, routeFilterName)
	if err != nil {
		return nil, err
	}
	return &routeFilterRuleListResultPager{
		pipeline:  client.p,
		request:   req,
		responder: client.listByRouteFilterHandleResponse,
		advancer: func(resp *RouteFilterRuleListResultResponse) (*azcore.Request, error) {
			u, err := url.Parse(*resp.RouteFilterRuleListResult.NextLink)
			if err != nil {
				return nil, fmt.Errorf("invalid NextLink: %w", err)
			}
			if u.Scheme == "" {
				return nil, fmt.Errorf("no scheme detected in NextLink %s", *resp.RouteFilterRuleListResult.NextLink)
			}
			return azcore.NewRequest(http.MethodGet, *u), nil
		},
	}, nil
}

// listByRouteFilterCreateRequest creates the ListByRouteFilter request.
func (client *routeFilterRulesOperations) listByRouteFilterCreateRequest(resourceGroupName string, routeFilterName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}/routeFilterRules"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{routeFilterName}", url.PathEscape(routeFilterName))
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

// listByRouteFilterHandleResponse handles the ListByRouteFilter response.
func (client *routeFilterRulesOperations) listByRouteFilterHandleResponse(resp *azcore.Response) (*RouteFilterRuleListResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.listByRouteFilterHandleError(resp)
	}
	result := RouteFilterRuleListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.RouteFilterRuleListResult)
}

// listByRouteFilterHandleError handles the ListByRouteFilter error response.
func (client *routeFilterRulesOperations) listByRouteFilterHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
