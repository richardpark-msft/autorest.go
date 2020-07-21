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

// VirtualRouterPeeringsOperations contains the methods for the VirtualRouterPeerings group.
type VirtualRouterPeeringsOperations interface {
	// BeginCreateOrUpdate - Creates or updates the specified Virtual Router Peering.
	BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, virtualRouterName string, peeringName string, parameters VirtualRouterPeering) (*VirtualRouterPeeringPollerResponse, error)
	// ResumeCreateOrUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeCreateOrUpdate(token string) (VirtualRouterPeeringPoller, error)
	// BeginDelete - Deletes the specified peering from a Virtual Router.
	BeginDelete(ctx context.Context, resourceGroupName string, virtualRouterName string, peeringName string) (*HTTPPollerResponse, error)
	// ResumeDelete - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDelete(token string) (HTTPPoller, error)
	// Get - Gets the specified Virtual Router Peering.
	Get(ctx context.Context, resourceGroupName string, virtualRouterName string, peeringName string) (*VirtualRouterPeeringResponse, error)
	// List - Lists all Virtual Router Peerings in a Virtual Router resource.
	List(resourceGroupName string, virtualRouterName string) (VirtualRouterPeeringListResultPager, error)
}

// virtualRouterPeeringsOperations implements the VirtualRouterPeeringsOperations interface.
type virtualRouterPeeringsOperations struct {
	*Client
	subscriptionID string
}

// CreateOrUpdate - Creates or updates the specified Virtual Router Peering.
func (client *virtualRouterPeeringsOperations) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, virtualRouterName string, peeringName string, parameters VirtualRouterPeering) (*VirtualRouterPeeringPollerResponse, error) {
	req, err := client.createOrUpdateCreateRequest(resourceGroupName, virtualRouterName, peeringName, parameters)
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
	pt, err := createPollingTracker("virtualRouterPeeringsOperations.CreateOrUpdate", "azure-async-operation", resp, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &virtualRouterPeeringPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*VirtualRouterPeeringResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *virtualRouterPeeringsOperations) ResumeCreateOrUpdate(token string) (VirtualRouterPeeringPoller, error) {
	pt, err := resumePollingTracker("virtualRouterPeeringsOperations.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &virtualRouterPeeringPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *virtualRouterPeeringsOperations) createOrUpdateCreateRequest(resourceGroupName string, virtualRouterName string, peeringName string, parameters VirtualRouterPeering) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualRouters/{virtualRouterName}/peerings/{peeringName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualRouterName}", url.PathEscape(virtualRouterName))
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
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
func (client *virtualRouterPeeringsOperations) createOrUpdateHandleResponse(resp *azcore.Response) (*VirtualRouterPeeringPollerResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated, http.StatusNoContent) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return &VirtualRouterPeeringPollerResponse{RawResponse: resp.Response}, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *virtualRouterPeeringsOperations) createOrUpdateHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Delete - Deletes the specified peering from a Virtual Router.
func (client *virtualRouterPeeringsOperations) BeginDelete(ctx context.Context, resourceGroupName string, virtualRouterName string, peeringName string) (*HTTPPollerResponse, error) {
	req, err := client.deleteCreateRequest(resourceGroupName, virtualRouterName, peeringName)
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
	pt, err := createPollingTracker("virtualRouterPeeringsOperations.Delete", "location", resp, client.deleteHandleError)
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

func (client *virtualRouterPeeringsOperations) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := resumePollingTracker("virtualRouterPeeringsOperations.Delete", token, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *virtualRouterPeeringsOperations) deleteCreateRequest(resourceGroupName string, virtualRouterName string, peeringName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualRouters/{virtualRouterName}/peerings/{peeringName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualRouterName}", url.PathEscape(virtualRouterName))
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
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
func (client *virtualRouterPeeringsOperations) deleteHandleResponse(resp *azcore.Response) (*HTTPPollerResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return &HTTPPollerResponse{RawResponse: resp.Response}, nil
}

// deleteHandleError handles the Delete error response.
func (client *virtualRouterPeeringsOperations) deleteHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Get - Gets the specified Virtual Router Peering.
func (client *virtualRouterPeeringsOperations) Get(ctx context.Context, resourceGroupName string, virtualRouterName string, peeringName string) (*VirtualRouterPeeringResponse, error) {
	req, err := client.getCreateRequest(resourceGroupName, virtualRouterName, peeringName)
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
func (client *virtualRouterPeeringsOperations) getCreateRequest(resourceGroupName string, virtualRouterName string, peeringName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualRouters/{virtualRouterName}/peerings/{peeringName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualRouterName}", url.PathEscape(virtualRouterName))
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
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
func (client *virtualRouterPeeringsOperations) getHandleResponse(resp *azcore.Response) (*VirtualRouterPeeringResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getHandleError(resp)
	}
	result := VirtualRouterPeeringResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VirtualRouterPeering)
}

// getHandleError handles the Get error response.
func (client *virtualRouterPeeringsOperations) getHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// List - Lists all Virtual Router Peerings in a Virtual Router resource.
func (client *virtualRouterPeeringsOperations) List(resourceGroupName string, virtualRouterName string) (VirtualRouterPeeringListResultPager, error) {
	req, err := client.listCreateRequest(resourceGroupName, virtualRouterName)
	if err != nil {
		return nil, err
	}
	return &virtualRouterPeeringListResultPager{
		pipeline:  client.p,
		request:   req,
		responder: client.listHandleResponse,
		advancer: func(resp *VirtualRouterPeeringListResultResponse) (*azcore.Request, error) {
			u, err := url.Parse(*resp.VirtualRouterPeeringListResult.NextLink)
			if err != nil {
				return nil, fmt.Errorf("invalid NextLink: %w", err)
			}
			if u.Scheme == "" {
				return nil, fmt.Errorf("no scheme detected in NextLink %s", *resp.VirtualRouterPeeringListResult.NextLink)
			}
			return azcore.NewRequest(http.MethodGet, *u), nil
		},
	}, nil
}

// listCreateRequest creates the List request.
func (client *virtualRouterPeeringsOperations) listCreateRequest(resourceGroupName string, virtualRouterName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualRouters/{virtualRouterName}/peerings"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualRouterName}", url.PathEscape(virtualRouterName))
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
func (client *virtualRouterPeeringsOperations) listHandleResponse(resp *azcore.Response) (*VirtualRouterPeeringListResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.listHandleError(resp)
	}
	result := VirtualRouterPeeringListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VirtualRouterPeeringListResult)
}

// listHandleError handles the List error response.
func (client *virtualRouterPeeringsOperations) listHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
