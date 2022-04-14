//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azartifacts

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type libraryClient struct {
	endpoint string
	pl       runtime.Pipeline
}

// newLibraryClient creates a new instance of libraryClient with the specified values.
// endpoint - The workspace development endpoint, for example https://myworkspace.dev.azuresynapse.net.
// pl - the pipeline used for sending requests and handling responses.
func newLibraryClient(endpoint string, pl runtime.Pipeline) *libraryClient {
	client := &libraryClient{
		endpoint: endpoint,
		pl:       pl,
	}
	return client
}

// Append - Append the content to the library resource created using the create operation. The maximum content size is 4MiB.
// Content larger than 4MiB must be appended in 4MiB chunks
// If the operation fails it returns an *azcore.ResponseError type.
// libraryName - file name to upload. Minimum length of the filename should be 1 excluding the extension length.
// content - Library file chunk.
// options - libraryClientAppendOptions contains the optional parameters for the libraryClient.Append method.
func (client *libraryClient) Append(ctx context.Context, libraryName string, content io.ReadSeekCloser, options *libraryClientAppendOptions) (libraryClientAppendResponse, error) {
	req, err := client.appendCreateRequest(ctx, libraryName, content, options)
	if err != nil {
		return libraryClientAppendResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return libraryClientAppendResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusCreated) {
		return libraryClientAppendResponse{}, runtime.NewResponseError(resp)
	}
	return libraryClientAppendResponse{}, nil
}

// appendCreateRequest creates the Append request.
func (client *libraryClient) appendCreateRequest(ctx context.Context, libraryName string, content io.ReadSeekCloser, options *libraryClientAppendOptions) (*policy.Request, error) {
	urlPath := "/libraries/{libraryName}"
	if libraryName == "" {
		return nil, errors.New("parameter libraryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{libraryName}", url.PathEscape(libraryName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.XMSBlobConditionAppendpos != nil {
		req.Raw().Header.Set("x-ms-blob-condition-appendpos", strconv.FormatInt(*options.XMSBlobConditionAppendpos, 10))
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, req.SetBody(content, "application/octet-stream")
}

// BeginCreate - Creates a library with the library name.
// If the operation fails it returns an *azcore.ResponseError type.
// libraryName - file name to upload. Minimum length of the filename should be 1 excluding the extension length.
// options - libraryClientBeginCreateOptions contains the optional parameters for the libraryClient.BeginCreate method.
func (client *libraryClient) BeginCreate(ctx context.Context, libraryName string, options *libraryClientBeginCreateOptions) (*runtime.Poller[libraryClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, libraryName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[libraryClientCreateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[libraryClientCreateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Create - Creates a library with the library name.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *libraryClient) create(ctx context.Context, libraryName string, options *libraryClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, libraryName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *libraryClient) createCreateRequest(ctx context.Context, libraryName string, options *libraryClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/libraries/{libraryName}"
	if libraryName == "" {
		return nil, errors.New("parameter libraryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{libraryName}", url.PathEscape(libraryName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// BeginDelete - Delete Library
// If the operation fails it returns an *azcore.ResponseError type.
// libraryName - file name to upload. Minimum length of the filename should be 1 excluding the extension length.
// options - libraryClientBeginDeleteOptions contains the optional parameters for the libraryClient.BeginDelete method.
func (client *libraryClient) BeginDelete(ctx context.Context, libraryName string, options *libraryClientBeginDeleteOptions) (*runtime.Poller[libraryClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, libraryName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[libraryClientDeleteResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[libraryClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Delete Library
// If the operation fails it returns an *azcore.ResponseError type.
func (client *libraryClient) deleteOperation(ctx context.Context, libraryName string, options *libraryClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, libraryName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusConflict) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *libraryClient) deleteCreateRequest(ctx context.Context, libraryName string, options *libraryClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/libraries/{libraryName}"
	if libraryName == "" {
		return nil, errors.New("parameter libraryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{libraryName}", url.PathEscape(libraryName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// BeginFlush - Flush Library
// If the operation fails it returns an *azcore.ResponseError type.
// libraryName - file name to upload. Minimum length of the filename should be 1 excluding the extension length.
// options - libraryClientBeginFlushOptions contains the optional parameters for the libraryClient.BeginFlush method.
func (client *libraryClient) BeginFlush(ctx context.Context, libraryName string, options *libraryClientBeginFlushOptions) (*runtime.Poller[libraryClientFlushResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.flush(ctx, libraryName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[libraryClientFlushResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[libraryClientFlushResponse](options.ResumeToken, client.pl, nil)
	}
}

// Flush - Flush Library
// If the operation fails it returns an *azcore.ResponseError type.
func (client *libraryClient) flush(ctx context.Context, libraryName string, options *libraryClientBeginFlushOptions) (*http.Response, error) {
	req, err := client.flushCreateRequest(ctx, libraryName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// flushCreateRequest creates the Flush request.
func (client *libraryClient) flushCreateRequest(ctx context.Context, libraryName string, options *libraryClientBeginFlushOptions) (*policy.Request, error) {
	urlPath := "/libraries/{libraryName}/flush"
	if libraryName == "" {
		return nil, errors.New("parameter libraryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{libraryName}", url.PathEscape(libraryName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Get Library
// If the operation fails it returns an *azcore.ResponseError type.
// libraryName - file name to upload. Minimum length of the filename should be 1 excluding the extension length.
// options - libraryClientGetOptions contains the optional parameters for the libraryClient.Get method.
func (client *libraryClient) Get(ctx context.Context, libraryName string, options *libraryClientGetOptions) (libraryClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, libraryName, options)
	if err != nil {
		return libraryClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return libraryClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNotModified) {
		return libraryClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *libraryClient) getCreateRequest(ctx context.Context, libraryName string, options *libraryClientGetOptions) (*policy.Request, error) {
	urlPath := "/libraries/{libraryName}"
	if libraryName == "" {
		return nil, errors.New("parameter libraryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{libraryName}", url.PathEscape(libraryName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *libraryClient) getHandleResponse(resp *http.Response) (libraryClientGetResponse, error) {
	result := libraryClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LibraryResource); err != nil {
		return libraryClientGetResponse{}, err
	}
	return result, nil
}

// GetOperationResult - Get Operation result for Library
// If the operation fails it returns an *azcore.ResponseError type.
// operationID - operation id for which status is requested
// options - libraryClientGetOperationResultOptions contains the optional parameters for the libraryClient.GetOperationResult
// method.
func (client *libraryClient) GetOperationResult(ctx context.Context, operationID string, options *libraryClientGetOperationResultOptions) (libraryClientGetOperationResultResponse, error) {
	req, err := client.getOperationResultCreateRequest(ctx, operationID, options)
	if err != nil {
		return libraryClientGetOperationResultResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return libraryClientGetOperationResultResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return libraryClientGetOperationResultResponse{}, runtime.NewResponseError(resp)
	}
	return client.getOperationResultHandleResponse(resp)
}

// getOperationResultCreateRequest creates the GetOperationResult request.
func (client *libraryClient) getOperationResultCreateRequest(ctx context.Context, operationID string, options *libraryClientGetOperationResultOptions) (*policy.Request, error) {
	urlPath := "/libraryOperationResults/{operationId}"
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getOperationResultHandleResponse handles the GetOperationResult response.
func (client *libraryClient) getOperationResultHandleResponse(resp *http.Response) (libraryClientGetOperationResultResponse, error) {
	result := libraryClientGetOperationResultResponse{}
	switch resp.StatusCode {
	case http.StatusOK:
		var val LibraryResource
		if err := runtime.UnmarshalAsJSON(resp, &val); err != nil {
			return libraryClientGetOperationResultResponse{}, err
		}
		result.Value = val
	case http.StatusAccepted:
		var val OperationResult
		if err := runtime.UnmarshalAsJSON(resp, &val); err != nil {
			return libraryClientGetOperationResultResponse{}, err
		}
		result.Value = val
	default:
		return libraryClientGetOperationResultResponse{}, fmt.Errorf("unhandled HTTP status code %d", resp.StatusCode)
	}
	return result, nil
}

// NewListPager - Lists Library.
// If the operation fails it returns an *azcore.ResponseError type.
// options - libraryClientListOptions contains the optional parameters for the libraryClient.List method.
func (client *libraryClient) NewListPager(options *libraryClientListOptions) *runtime.Pager[libraryClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[libraryClientListResponse]{
		More: func(page libraryClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *libraryClientListResponse) (libraryClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return libraryClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return libraryClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return libraryClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *libraryClient) listCreateRequest(ctx context.Context, options *libraryClientListOptions) (*policy.Request, error) {
	urlPath := "/libraries"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *libraryClient) listHandleResponse(resp *http.Response) (libraryClientListResponse, error) {
	result := libraryClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LibraryListResponse); err != nil {
		return libraryClientListResponse{}, err
	}
	return result, nil
}
