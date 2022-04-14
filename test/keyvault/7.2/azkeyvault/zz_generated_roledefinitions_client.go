//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azkeyvault

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// RoleDefinitionsClient contains the methods for the RoleDefinitions group.
// Don't use this type directly, use NewRoleDefinitionsClient() instead.
type RoleDefinitionsClient struct {
	pl runtime.Pipeline
}

// NewRoleDefinitionsClient creates a new instance of RoleDefinitionsClient with the specified values.
// pl - the pipeline used for sending requests and handling responses.
func NewRoleDefinitionsClient(pl runtime.Pipeline) *RoleDefinitionsClient {
	client := &RoleDefinitionsClient{
		pl: pl,
	}
	return client
}

// CreateOrUpdate - Creates or updates a custom role definition.
// If the operation fails it returns an *azcore.ResponseError type.
// vaultBaseURL - The vault name, for example https://myvault.vault.azure.net.
// scope - The scope of the role definition to create or update. Managed HSM only supports '/'.
// roleDefinitionName - The name of the role definition to create or update. It can be any valid GUID.
// parameters - Parameters for the role definition.
// options - RoleDefinitionsClientCreateOrUpdateOptions contains the optional parameters for the RoleDefinitionsClient.CreateOrUpdate
// method.
func (client *RoleDefinitionsClient) CreateOrUpdate(ctx context.Context, vaultBaseURL string, scope string, roleDefinitionName string, parameters RoleDefinitionCreateParameters, options *RoleDefinitionsClientCreateOrUpdateOptions) (RoleDefinitionsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, vaultBaseURL, scope, roleDefinitionName, parameters, options)
	if err != nil {
		return RoleDefinitionsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RoleDefinitionsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusCreated) {
		return RoleDefinitionsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *RoleDefinitionsClient) createOrUpdateCreateRequest(ctx context.Context, vaultBaseURL string, scope string, roleDefinitionName string, parameters RoleDefinitionCreateParameters, options *RoleDefinitionsClientCreateOrUpdateOptions) (*policy.Request, error) {
	host := "{vaultBaseUrl}"
	host = strings.ReplaceAll(host, "{vaultBaseUrl}", vaultBaseURL)
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleDefinitions/{roleDefinitionName}"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if roleDefinitionName == "" {
		return nil, errors.New("parameter roleDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleDefinitionName}", url.PathEscape(roleDefinitionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "7.2")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *RoleDefinitionsClient) createOrUpdateHandleResponse(resp *http.Response) (RoleDefinitionsClientCreateOrUpdateResponse, error) {
	result := RoleDefinitionsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleDefinition); err != nil {
		return RoleDefinitionsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes a custom role definition.
// If the operation fails it returns an *azcore.ResponseError type.
// vaultBaseURL - The vault name, for example https://myvault.vault.azure.net.
// scope - The scope of the role definition to delete. Managed HSM only supports '/'.
// roleDefinitionName - The name (GUID) of the role definition to delete.
// options - RoleDefinitionsClientDeleteOptions contains the optional parameters for the RoleDefinitionsClient.Delete method.
func (client *RoleDefinitionsClient) Delete(ctx context.Context, vaultBaseURL string, scope string, roleDefinitionName string, options *RoleDefinitionsClientDeleteOptions) (RoleDefinitionsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, vaultBaseURL, scope, roleDefinitionName, options)
	if err != nil {
		return RoleDefinitionsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RoleDefinitionsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RoleDefinitionsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return client.deleteHandleResponse(resp)
}

// deleteCreateRequest creates the Delete request.
func (client *RoleDefinitionsClient) deleteCreateRequest(ctx context.Context, vaultBaseURL string, scope string, roleDefinitionName string, options *RoleDefinitionsClientDeleteOptions) (*policy.Request, error) {
	host := "{vaultBaseUrl}"
	host = strings.ReplaceAll(host, "{vaultBaseUrl}", vaultBaseURL)
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleDefinitions/{roleDefinitionName}"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if roleDefinitionName == "" {
		return nil, errors.New("parameter roleDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleDefinitionName}", url.PathEscape(roleDefinitionName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "7.2")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleResponse handles the Delete response.
func (client *RoleDefinitionsClient) deleteHandleResponse(resp *http.Response) (RoleDefinitionsClientDeleteResponse, error) {
	result := RoleDefinitionsClientDeleteResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleDefinition); err != nil {
		return RoleDefinitionsClientDeleteResponse{}, err
	}
	return result, nil
}

// Get - Get the specified role definition.
// If the operation fails it returns an *azcore.ResponseError type.
// vaultBaseURL - The vault name, for example https://myvault.vault.azure.net.
// scope - The scope of the role definition to get. Managed HSM only supports '/'.
// roleDefinitionName - The name of the role definition to get.
// options - RoleDefinitionsClientGetOptions contains the optional parameters for the RoleDefinitionsClient.Get method.
func (client *RoleDefinitionsClient) Get(ctx context.Context, vaultBaseURL string, scope string, roleDefinitionName string, options *RoleDefinitionsClientGetOptions) (RoleDefinitionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, vaultBaseURL, scope, roleDefinitionName, options)
	if err != nil {
		return RoleDefinitionsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RoleDefinitionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RoleDefinitionsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *RoleDefinitionsClient) getCreateRequest(ctx context.Context, vaultBaseURL string, scope string, roleDefinitionName string, options *RoleDefinitionsClientGetOptions) (*policy.Request, error) {
	host := "{vaultBaseUrl}"
	host = strings.ReplaceAll(host, "{vaultBaseUrl}", vaultBaseURL)
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleDefinitions/{roleDefinitionName}"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if roleDefinitionName == "" {
		return nil, errors.New("parameter roleDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleDefinitionName}", url.PathEscape(roleDefinitionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "7.2")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RoleDefinitionsClient) getHandleResponse(resp *http.Response) (RoleDefinitionsClientGetResponse, error) {
	result := RoleDefinitionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleDefinition); err != nil {
		return RoleDefinitionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Get all role definitions that are applicable at scope and above.
// If the operation fails it returns an *azcore.ResponseError type.
// vaultBaseURL - The vault name, for example https://myvault.vault.azure.net.
// scope - The scope of the role definition.
// options - RoleDefinitionsClientListOptions contains the optional parameters for the RoleDefinitionsClient.List method.
func (client *RoleDefinitionsClient) NewListPager(vaultBaseURL string, scope string, options *RoleDefinitionsClientListOptions) *runtime.Pager[RoleDefinitionsClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[RoleDefinitionsClientListResponse]{
		More: func(page RoleDefinitionsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RoleDefinitionsClientListResponse) (RoleDefinitionsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, vaultBaseURL, scope, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return RoleDefinitionsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return RoleDefinitionsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return RoleDefinitionsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *RoleDefinitionsClient) listCreateRequest(ctx context.Context, vaultBaseURL string, scope string, options *RoleDefinitionsClientListOptions) (*policy.Request, error) {
	host := "{vaultBaseUrl}"
	host = strings.ReplaceAll(host, "{vaultBaseUrl}", vaultBaseURL)
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleDefinitions"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "7.2")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *RoleDefinitionsClient) listHandleResponse(resp *http.Response) (RoleDefinitionsClientListResponse, error) {
	result := RoleDefinitionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleDefinitionListResult); err != nil {
		return RoleDefinitionsClientListResponse{}, err
	}
	return result, nil
}
