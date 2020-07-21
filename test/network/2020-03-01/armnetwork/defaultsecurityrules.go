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

// DefaultSecurityRulesOperations contains the methods for the DefaultSecurityRules group.
type DefaultSecurityRulesOperations interface {
	// Get - Get the specified default network security rule.
	Get(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, defaultSecurityRuleName string) (*SecurityRuleResponse, error)
	// List - Gets all default security rules in a network security group.
	List(resourceGroupName string, networkSecurityGroupName string) (SecurityRuleListResultPager, error)
}

// defaultSecurityRulesOperations implements the DefaultSecurityRulesOperations interface.
type defaultSecurityRulesOperations struct {
	*Client
	subscriptionID string
}

// Get - Get the specified default network security rule.
func (client *defaultSecurityRulesOperations) Get(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, defaultSecurityRuleName string) (*SecurityRuleResponse, error) {
	req, err := client.getCreateRequest(resourceGroupName, networkSecurityGroupName, defaultSecurityRuleName)
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
func (client *defaultSecurityRulesOperations) getCreateRequest(resourceGroupName string, networkSecurityGroupName string, defaultSecurityRuleName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}/defaultSecurityRules/{defaultSecurityRuleName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{networkSecurityGroupName}", url.PathEscape(networkSecurityGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{defaultSecurityRuleName}", url.PathEscape(defaultSecurityRuleName))
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
func (client *defaultSecurityRulesOperations) getHandleResponse(resp *azcore.Response) (*SecurityRuleResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getHandleError(resp)
	}
	result := SecurityRuleResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.SecurityRule)
}

// getHandleError handles the Get error response.
func (client *defaultSecurityRulesOperations) getHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// List - Gets all default security rules in a network security group.
func (client *defaultSecurityRulesOperations) List(resourceGroupName string, networkSecurityGroupName string) (SecurityRuleListResultPager, error) {
	req, err := client.listCreateRequest(resourceGroupName, networkSecurityGroupName)
	if err != nil {
		return nil, err
	}
	return &securityRuleListResultPager{
		pipeline:  client.p,
		request:   req,
		responder: client.listHandleResponse,
		advancer: func(resp *SecurityRuleListResultResponse) (*azcore.Request, error) {
			u, err := url.Parse(*resp.SecurityRuleListResult.NextLink)
			if err != nil {
				return nil, fmt.Errorf("invalid NextLink: %w", err)
			}
			if u.Scheme == "" {
				return nil, fmt.Errorf("no scheme detected in NextLink %s", *resp.SecurityRuleListResult.NextLink)
			}
			return azcore.NewRequest(http.MethodGet, *u), nil
		},
	}, nil
}

// listCreateRequest creates the List request.
func (client *defaultSecurityRulesOperations) listCreateRequest(resourceGroupName string, networkSecurityGroupName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}/defaultSecurityRules"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{networkSecurityGroupName}", url.PathEscape(networkSecurityGroupName))
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
func (client *defaultSecurityRulesOperations) listHandleResponse(resp *azcore.Response) (*SecurityRuleListResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.listHandleError(resp)
	}
	result := SecurityRuleListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.SecurityRuleListResult)
}

// listHandleError handles the List error response.
func (client *defaultSecurityRulesOperations) listHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
