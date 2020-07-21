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

// PeerExpressRouteCircuitConnectionsOperations contains the methods for the PeerExpressRouteCircuitConnections group.
type PeerExpressRouteCircuitConnectionsOperations interface {
	// Get - Gets the specified Peer Express Route Circuit Connection from the specified express route circuit.
	Get(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, connectionName string) (*PeerExpressRouteCircuitConnectionResponse, error)
	// List - Gets all global reach peer connections associated with a private peering in an express route circuit.
	List(resourceGroupName string, circuitName string, peeringName string) (PeerExpressRouteCircuitConnectionListResultPager, error)
}

// peerExpressRouteCircuitConnectionsOperations implements the PeerExpressRouteCircuitConnectionsOperations interface.
type peerExpressRouteCircuitConnectionsOperations struct {
	*Client
	subscriptionID string
}

// Get - Gets the specified Peer Express Route Circuit Connection from the specified express route circuit.
func (client *peerExpressRouteCircuitConnectionsOperations) Get(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, connectionName string) (*PeerExpressRouteCircuitConnectionResponse, error) {
	req, err := client.getCreateRequest(resourceGroupName, circuitName, peeringName, connectionName)
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
func (client *peerExpressRouteCircuitConnectionsOperations) getCreateRequest(resourceGroupName string, circuitName string, peeringName string, connectionName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}/peerConnections/{connectionName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{circuitName}", url.PathEscape(circuitName))
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
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
func (client *peerExpressRouteCircuitConnectionsOperations) getHandleResponse(resp *azcore.Response) (*PeerExpressRouteCircuitConnectionResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getHandleError(resp)
	}
	result := PeerExpressRouteCircuitConnectionResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.PeerExpressRouteCircuitConnection)
}

// getHandleError handles the Get error response.
func (client *peerExpressRouteCircuitConnectionsOperations) getHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// List - Gets all global reach peer connections associated with a private peering in an express route circuit.
func (client *peerExpressRouteCircuitConnectionsOperations) List(resourceGroupName string, circuitName string, peeringName string) (PeerExpressRouteCircuitConnectionListResultPager, error) {
	req, err := client.listCreateRequest(resourceGroupName, circuitName, peeringName)
	if err != nil {
		return nil, err
	}
	return &peerExpressRouteCircuitConnectionListResultPager{
		pipeline:  client.p,
		request:   req,
		responder: client.listHandleResponse,
		advancer: func(resp *PeerExpressRouteCircuitConnectionListResultResponse) (*azcore.Request, error) {
			u, err := url.Parse(*resp.PeerExpressRouteCircuitConnectionListResult.NextLink)
			if err != nil {
				return nil, fmt.Errorf("invalid NextLink: %w", err)
			}
			if u.Scheme == "" {
				return nil, fmt.Errorf("no scheme detected in NextLink %s", *resp.PeerExpressRouteCircuitConnectionListResult.NextLink)
			}
			return azcore.NewRequest(http.MethodGet, *u), nil
		},
	}, nil
}

// listCreateRequest creates the List request.
func (client *peerExpressRouteCircuitConnectionsOperations) listCreateRequest(resourceGroupName string, circuitName string, peeringName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}/peerConnections"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{circuitName}", url.PathEscape(circuitName))
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

// listHandleResponse handles the List response.
func (client *peerExpressRouteCircuitConnectionsOperations) listHandleResponse(resp *azcore.Response) (*PeerExpressRouteCircuitConnectionListResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.listHandleError(resp)
	}
	result := PeerExpressRouteCircuitConnectionListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.PeerExpressRouteCircuitConnectionListResult)
}

// listHandleError handles the List error response.
func (client *peerExpressRouteCircuitConnectionsOperations) listHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
