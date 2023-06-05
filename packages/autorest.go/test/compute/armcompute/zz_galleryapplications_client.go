//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// GalleryApplicationsClient contains the methods for the GalleryApplications group.
// Don't use this type directly, use NewGalleryApplicationsClient() instead.
type GalleryApplicationsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewGalleryApplicationsClient creates a new instance of GalleryApplicationsClient with the specified values.
//   - subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
//     part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewGalleryApplicationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*GalleryApplicationsClient, error) {
	cl, err := arm.NewClient(moduleName+".GalleryApplicationsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &GalleryApplicationsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or update a gallery Application Definition.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-10-01
//   - resourceGroupName - The name of the resource group.
//   - galleryName - The name of the Shared Application Gallery in which the Application Definition is to be created.
//   - galleryApplicationName - The name of the gallery Application Definition to be created or updated. The allowed characters
//     are alphabets and numbers with dots, dashes, and periods allowed in the middle. The maximum length is 80
//     characters.
//   - galleryApplication - Parameters supplied to the create or update gallery Application operation.
//   - options - GalleryApplicationsClientBeginCreateOrUpdateOptions contains the optional parameters for the GalleryApplicationsClient.BeginCreateOrUpdate
//     method.
func (client *GalleryApplicationsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplication GalleryApplication, options *GalleryApplicationsClientBeginCreateOrUpdateOptions) (*runtime.Poller[GalleryApplicationsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, galleryName, galleryApplicationName, galleryApplication, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[GalleryApplicationsClientCreateOrUpdateResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[GalleryApplicationsClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Create or update a gallery Application Definition.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-10-01
func (client *GalleryApplicationsClient) createOrUpdate(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplication GalleryApplication, options *GalleryApplicationsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "GalleryApplicationsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, galleryName, galleryApplicationName, galleryApplication, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *GalleryApplicationsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplication GalleryApplication, options *GalleryApplicationsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/applications/{galleryApplicationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if galleryName == "" {
		return nil, errors.New("parameter galleryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryName}", url.PathEscape(galleryName))
	if galleryApplicationName == "" {
		return nil, errors.New("parameter galleryApplicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryApplicationName}", url.PathEscape(galleryApplicationName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, galleryApplication); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete a gallery Application.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-10-01
//   - resourceGroupName - The name of the resource group.
//   - galleryName - The name of the Shared Application Gallery in which the Application Definition is to be deleted.
//   - galleryApplicationName - The name of the gallery Application Definition to be deleted.
//   - options - GalleryApplicationsClientBeginDeleteOptions contains the optional parameters for the GalleryApplicationsClient.BeginDelete
//     method.
func (client *GalleryApplicationsClient) BeginDelete(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, options *GalleryApplicationsClientBeginDeleteOptions) (*runtime.Poller[GalleryApplicationsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, galleryName, galleryApplicationName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[GalleryApplicationsClientDeleteResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[GalleryApplicationsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Delete a gallery Application.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-10-01
func (client *GalleryApplicationsClient) deleteOperation(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, options *GalleryApplicationsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "GalleryApplicationsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, galleryName, galleryApplicationName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *GalleryApplicationsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, options *GalleryApplicationsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/applications/{galleryApplicationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if galleryName == "" {
		return nil, errors.New("parameter galleryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryName}", url.PathEscape(galleryName))
	if galleryApplicationName == "" {
		return nil, errors.New("parameter galleryApplicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryApplicationName}", url.PathEscape(galleryApplicationName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Retrieves information about a gallery Application Definition.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-10-01
//   - resourceGroupName - The name of the resource group.
//   - galleryName - The name of the Shared Application Gallery from which the Application Definitions are to be retrieved.
//   - galleryApplicationName - The name of the gallery Application Definition to be retrieved.
//   - options - GalleryApplicationsClientGetOptions contains the optional parameters for the GalleryApplicationsClient.Get method.
func (client *GalleryApplicationsClient) Get(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, options *GalleryApplicationsClientGetOptions) (GalleryApplicationsClientGetResponse, error) {
	var err error
	const operationName = "GalleryApplicationsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, galleryName, galleryApplicationName, options)
	if err != nil {
		return GalleryApplicationsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return GalleryApplicationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return GalleryApplicationsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *GalleryApplicationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, options *GalleryApplicationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/applications/{galleryApplicationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if galleryName == "" {
		return nil, errors.New("parameter galleryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryName}", url.PathEscape(galleryName))
	if galleryApplicationName == "" {
		return nil, errors.New("parameter galleryApplicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryApplicationName}", url.PathEscape(galleryApplicationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *GalleryApplicationsClient) getHandleResponse(resp *http.Response) (GalleryApplicationsClientGetResponse, error) {
	result := GalleryApplicationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GalleryApplication); err != nil {
		return GalleryApplicationsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByGalleryPager - List gallery Application Definitions in a gallery.
//
// Generated from API version 2021-10-01
//   - resourceGroupName - The name of the resource group.
//   - galleryName - The name of the Shared Application Gallery from which Application Definitions are to be listed.
//   - options - GalleryApplicationsClientListByGalleryOptions contains the optional parameters for the GalleryApplicationsClient.NewListByGalleryPager
//     method.
func (client *GalleryApplicationsClient) NewListByGalleryPager(resourceGroupName string, galleryName string, options *GalleryApplicationsClientListByGalleryOptions) *runtime.Pager[GalleryApplicationsClientListByGalleryResponse] {
	return runtime.NewPager(runtime.PagingHandler[GalleryApplicationsClientListByGalleryResponse]{
		More: func(page GalleryApplicationsClientListByGalleryResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *GalleryApplicationsClientListByGalleryResponse) (GalleryApplicationsClientListByGalleryResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "GalleryApplicationsClient.NewListByGalleryPager")
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByGalleryCreateRequest(ctx, resourceGroupName, galleryName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return GalleryApplicationsClientListByGalleryResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return GalleryApplicationsClientListByGalleryResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return GalleryApplicationsClientListByGalleryResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByGalleryHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByGalleryCreateRequest creates the ListByGallery request.
func (client *GalleryApplicationsClient) listByGalleryCreateRequest(ctx context.Context, resourceGroupName string, galleryName string, options *GalleryApplicationsClientListByGalleryOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/applications"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if galleryName == "" {
		return nil, errors.New("parameter galleryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryName}", url.PathEscape(galleryName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByGalleryHandleResponse handles the ListByGallery response.
func (client *GalleryApplicationsClient) listByGalleryHandleResponse(resp *http.Response) (GalleryApplicationsClientListByGalleryResponse, error) {
	result := GalleryApplicationsClientListByGalleryResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GalleryApplicationList); err != nil {
		return GalleryApplicationsClientListByGalleryResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Update a gallery Application Definition.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-10-01
//   - resourceGroupName - The name of the resource group.
//   - galleryName - The name of the Shared Application Gallery in which the Application Definition is to be updated.
//   - galleryApplicationName - The name of the gallery Application Definition to be updated. The allowed characters are alphabets
//     and numbers with dots, dashes, and periods allowed in the middle. The maximum length is 80
//     characters.
//   - galleryApplication - Parameters supplied to the update gallery Application operation.
//   - options - GalleryApplicationsClientBeginUpdateOptions contains the optional parameters for the GalleryApplicationsClient.BeginUpdate
//     method.
func (client *GalleryApplicationsClient) BeginUpdate(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplication GalleryApplicationUpdate, options *GalleryApplicationsClientBeginUpdateOptions) (*runtime.Poller[GalleryApplicationsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, galleryName, galleryApplicationName, galleryApplication, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[GalleryApplicationsClientUpdateResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[GalleryApplicationsClientUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Update - Update a gallery Application Definition.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-10-01
func (client *GalleryApplicationsClient) update(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplication GalleryApplicationUpdate, options *GalleryApplicationsClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "GalleryApplicationsClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, galleryName, galleryApplicationName, galleryApplication, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// updateCreateRequest creates the Update request.
func (client *GalleryApplicationsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplication GalleryApplicationUpdate, options *GalleryApplicationsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/applications/{galleryApplicationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if galleryName == "" {
		return nil, errors.New("parameter galleryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryName}", url.PathEscape(galleryName))
	if galleryApplicationName == "" {
		return nil, errors.New("parameter galleryApplicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryApplicationName}", url.PathEscape(galleryApplicationName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, galleryApplication); err != nil {
		return nil, err
	}
	return req, nil
}
