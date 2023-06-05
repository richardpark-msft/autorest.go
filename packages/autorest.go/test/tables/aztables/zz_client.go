//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package aztables

import (
	"context"
	"encoding/xml"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// Client contains the methods for the Table group.
// Don't use this type directly, use a constructor function instead.
type Client struct {
	internal *azcore.Client
	endpoint string
}

// Create - Creates a new table under the given account.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-02-02
//   - tableProperties - The Table properties.
//   - options - ClientCreateOptions contains the optional parameters for the Client.Create method.
func (client *Client) Create(ctx context.Context, tableProperties Properties, options *ClientCreateOptions) (ClientCreateResponse, error) {
	var err error
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "Client.Create")
	req, err := client.createCreateRequest(ctx, tableProperties, options)
	if err != nil {
		return ClientCreateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusCreated, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ClientCreateResponse{}, err
	}
	resp, err := client.createHandleResponse(httpResp)
	return resp, err
}

// createCreateRequest creates the Create request.
func (client *Client) createCreateRequest(ctx context.Context, tableProperties Properties, options *ClientCreateOptions) (*policy.Request, error) {
	urlPath := "/Tables"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Format != nil {
		reqQP.Set("$format", string(*options.Format))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["x-ms-version"] = []string{"2019-02-02"}
	if options != nil && options.RequestID != nil {
		req.Raw().Header["x-ms-client-request-id"] = []string{*options.RequestID}
	}
	req.Raw().Header["DataServiceVersion"] = []string{"3.0"}
	if options != nil && options.ResponsePreference != nil {
		req.Raw().Header["Prefer"] = []string{string(*options.ResponsePreference)}
	}
	req.Raw().Header["Accept"] = []string{"application/json;odata=minimalmetadata"}
	if err := runtime.MarshalAsJSON(req, tableProperties); err != nil {
		return nil, err
	}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *Client) createHandleResponse(resp *http.Response) (ClientCreateResponse, error) {
	result := ClientCreateResponse{}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return ClientCreateResponse{}, err
		}
		result.Date = &date
	}
	if val := resp.Header.Get("Preference-Applied"); val != "" {
		result.PreferenceApplied = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.Response); err != nil {
		return ClientCreateResponse{}, err
	}
	return result, nil
}

// Delete - Operation permanently deletes the specified table.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-02-02
//   - table - The name of the table.
//   - options - ClientDeleteOptions contains the optional parameters for the Client.Delete method.
func (client *Client) Delete(ctx context.Context, table string, options *ClientDeleteOptions) (ClientDeleteResponse, error) {
	var err error
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "Client.Delete")
	req, err := client.deleteCreateRequest(ctx, table, options)
	if err != nil {
		return ClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ClientDeleteResponse{}, err
	}
	resp, err := client.deleteHandleResponse(httpResp)
	return resp, err
}

// deleteCreateRequest creates the Delete request.
func (client *Client) deleteCreateRequest(ctx context.Context, table string, options *ClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/Tables('{table}')"
	if table == "" {
		return nil, errors.New("parameter table cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{table}", url.PathEscape(table))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["x-ms-version"] = []string{"2019-02-02"}
	if options != nil && options.RequestID != nil {
		req.Raw().Header["x-ms-client-request-id"] = []string{*options.RequestID}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// deleteHandleResponse handles the Delete response.
func (client *Client) deleteHandleResponse(resp *http.Response) (ClientDeleteResponse, error) {
	result := ClientDeleteResponse{}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return ClientDeleteResponse{}, err
		}
		result.Date = &date
	}
	return result, nil
}

// DeleteEntity - Deletes the specified entity in a table.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-02-02
//   - table - The name of the table.
//   - partitionKey - The partition key of the entity.
//   - rowKey - The row key of the entity.
//   - ifMatch - Match condition for an entity to be deleted. If specified and a matching entity is not found, an error will be
//     raised. To force an unconditional delete, set to the wildcard character (*).
//   - options - ClientDeleteEntityOptions contains the optional parameters for the Client.DeleteEntity method.
func (client *Client) DeleteEntity(ctx context.Context, table string, partitionKey string, rowKey string, ifMatch string, options *ClientDeleteEntityOptions) (ClientDeleteEntityResponse, error) {
	var err error
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "Client.DeleteEntity")
	req, err := client.deleteEntityCreateRequest(ctx, table, partitionKey, rowKey, ifMatch, options)
	if err != nil {
		return ClientDeleteEntityResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ClientDeleteEntityResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ClientDeleteEntityResponse{}, err
	}
	resp, err := client.deleteEntityHandleResponse(httpResp)
	return resp, err
}

// deleteEntityCreateRequest creates the DeleteEntity request.
func (client *Client) deleteEntityCreateRequest(ctx context.Context, table string, partitionKey string, rowKey string, ifMatch string, options *ClientDeleteEntityOptions) (*policy.Request, error) {
	urlPath := "/{table}(PartitionKey='{partitionKey}',RowKey='{rowKey}')"
	if table == "" {
		return nil, errors.New("parameter table cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{table}", url.PathEscape(table))
	if partitionKey == "" {
		return nil, errors.New("parameter partitionKey cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{partitionKey}", url.PathEscape(partitionKey))
	if rowKey == "" {
		return nil, errors.New("parameter rowKey cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{rowKey}", url.PathEscape(rowKey))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	if options != nil && options.Format != nil {
		reqQP.Set("$format", string(*options.Format))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["x-ms-version"] = []string{"2019-02-02"}
	if options != nil && options.RequestID != nil {
		req.Raw().Header["x-ms-client-request-id"] = []string{*options.RequestID}
	}
	req.Raw().Header["DataServiceVersion"] = []string{"3.0"}
	req.Raw().Header["If-Match"] = []string{ifMatch}
	req.Raw().Header["Accept"] = []string{"application/json;odata=minimalmetadata"}
	return req, nil
}

// deleteEntityHandleResponse handles the DeleteEntity response.
func (client *Client) deleteEntityHandleResponse(resp *http.Response) (ClientDeleteEntityResponse, error) {
	result := ClientDeleteEntityResponse{}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return ClientDeleteEntityResponse{}, err
		}
		result.Date = &date
	}
	return result, nil
}

// GetAccessPolicy - Retrieves details about any stored access policies specified on the table that may be used with Shared
// Access Signatures.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-02-02
//   - table - The name of the table.
//   - options - ClientGetAccessPolicyOptions contains the optional parameters for the Client.GetAccessPolicy method.
func (client *Client) GetAccessPolicy(ctx context.Context, table string, options *ClientGetAccessPolicyOptions) (ClientGetAccessPolicyResponse, error) {
	var err error
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "Client.GetAccessPolicy")
	req, err := client.getAccessPolicyCreateRequest(ctx, table, options)
	if err != nil {
		return ClientGetAccessPolicyResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ClientGetAccessPolicyResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ClientGetAccessPolicyResponse{}, err
	}
	resp, err := client.getAccessPolicyHandleResponse(httpResp)
	return resp, err
}

// getAccessPolicyCreateRequest creates the GetAccessPolicy request.
func (client *Client) getAccessPolicyCreateRequest(ctx context.Context, table string, options *ClientGetAccessPolicyOptions) (*policy.Request, error) {
	urlPath := "/{table}"
	if table == "" {
		return nil, errors.New("parameter table cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{table}", url.PathEscape(table))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	reqQP.Set("comp", "acl")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["x-ms-version"] = []string{"2019-02-02"}
	if options != nil && options.RequestID != nil {
		req.Raw().Header["x-ms-client-request-id"] = []string{*options.RequestID}
	}
	req.Raw().Header["Accept"] = []string{"application/xml"}
	return req, nil
}

// getAccessPolicyHandleResponse handles the GetAccessPolicy response.
func (client *Client) getAccessPolicyHandleResponse(resp *http.Response) (ClientGetAccessPolicyResponse, error) {
	result := ClientGetAccessPolicyResponse{}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return ClientGetAccessPolicyResponse{}, err
		}
		result.Date = &date
	}
	if err := runtime.UnmarshalAsXML(resp, &result); err != nil {
		return ClientGetAccessPolicyResponse{}, err
	}
	return result, nil
}

// InsertEntity - Insert entity in a table.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-02-02
//   - table - The name of the table.
//   - options - ClientInsertEntityOptions contains the optional parameters for the Client.InsertEntity method.
func (client *Client) InsertEntity(ctx context.Context, table string, options *ClientInsertEntityOptions) (ClientInsertEntityResponse, error) {
	var err error
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "Client.InsertEntity")
	req, err := client.insertEntityCreateRequest(ctx, table, options)
	if err != nil {
		return ClientInsertEntityResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ClientInsertEntityResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusCreated, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ClientInsertEntityResponse{}, err
	}
	resp, err := client.insertEntityHandleResponse(httpResp)
	return resp, err
}

// insertEntityCreateRequest creates the InsertEntity request.
func (client *Client) insertEntityCreateRequest(ctx context.Context, table string, options *ClientInsertEntityOptions) (*policy.Request, error) {
	urlPath := "/{table}"
	if table == "" {
		return nil, errors.New("parameter table cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{table}", url.PathEscape(table))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	if options != nil && options.Format != nil {
		reqQP.Set("$format", string(*options.Format))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["x-ms-version"] = []string{"2019-02-02"}
	if options != nil && options.RequestID != nil {
		req.Raw().Header["x-ms-client-request-id"] = []string{*options.RequestID}
	}
	req.Raw().Header["DataServiceVersion"] = []string{"3.0"}
	if options != nil && options.ResponsePreference != nil {
		req.Raw().Header["Prefer"] = []string{string(*options.ResponsePreference)}
	}
	req.Raw().Header["Accept"] = []string{"application/json;odata=minimalmetadata"}
	if options != nil && options.TableEntityProperties != nil {
		if err := runtime.MarshalAsJSON(req, options.TableEntityProperties); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// insertEntityHandleResponse handles the InsertEntity response.
func (client *Client) insertEntityHandleResponse(resp *http.Response) (ClientInsertEntityResponse, error) {
	result := ClientInsertEntityResponse{}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return ClientInsertEntityResponse{}, err
		}
		result.Date = &date
	}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if val := resp.Header.Get("Preference-Applied"); val != "" {
		result.PreferenceApplied = &val
	}
	if val := resp.Header.Get("Content-Type"); val != "" {
		result.ContentType = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return ClientInsertEntityResponse{}, err
	}
	return result, nil
}

// MergeEntity - Merge entity in a table.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-02-02
//   - table - The name of the table.
//   - partitionKey - The partition key of the entity.
//   - rowKey - The row key of the entity.
//   - tableEntityProperties - The properties for the table entity.
//   - options - ClientMergeEntityOptions contains the optional parameters for the Client.MergeEntity method.
func (client *Client) MergeEntity(ctx context.Context, table string, partitionKey string, rowKey string, tableEntityProperties map[string]any, options *ClientMergeEntityOptions) (ClientMergeEntityResponse, error) {
	var err error
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "Client.MergeEntity")
	req, err := client.mergeEntityCreateRequest(ctx, table, partitionKey, rowKey, tableEntityProperties, options)
	if err != nil {
		return ClientMergeEntityResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ClientMergeEntityResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ClientMergeEntityResponse{}, err
	}
	resp, err := client.mergeEntityHandleResponse(httpResp)
	return resp, err
}

// mergeEntityCreateRequest creates the MergeEntity request.
func (client *Client) mergeEntityCreateRequest(ctx context.Context, table string, partitionKey string, rowKey string, tableEntityProperties map[string]any, options *ClientMergeEntityOptions) (*policy.Request, error) {
	urlPath := "/{table}(PartitionKey='{partitionKey}',RowKey='{rowKey}')"
	if table == "" {
		return nil, errors.New("parameter table cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{table}", url.PathEscape(table))
	if partitionKey == "" {
		return nil, errors.New("parameter partitionKey cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{partitionKey}", url.PathEscape(partitionKey))
	if rowKey == "" {
		return nil, errors.New("parameter rowKey cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{rowKey}", url.PathEscape(rowKey))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	if options != nil && options.Format != nil {
		reqQP.Set("$format", string(*options.Format))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["x-ms-version"] = []string{"2019-02-02"}
	if options != nil && options.RequestID != nil {
		req.Raw().Header["x-ms-client-request-id"] = []string{*options.RequestID}
	}
	req.Raw().Header["DataServiceVersion"] = []string{"3.0"}
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, tableEntityProperties); err != nil {
		return nil, err
	}
	return req, nil
}

// mergeEntityHandleResponse handles the MergeEntity response.
func (client *Client) mergeEntityHandleResponse(resp *http.Response) (ClientMergeEntityResponse, error) {
	result := ClientMergeEntityResponse{}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return ClientMergeEntityResponse{}, err
		}
		result.Date = &date
	}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	return result, nil
}

// Query - Queries tables under the given account.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-02-02
//   - options - ClientQueryOptions contains the optional parameters for the Client.Query method.
func (client *Client) Query(ctx context.Context, options *ClientQueryOptions) (ClientQueryResponse, error) {
	var err error
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "Client.Query")
	req, err := client.queryCreateRequest(ctx, options)
	if err != nil {
		return ClientQueryResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ClientQueryResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ClientQueryResponse{}, err
	}
	resp, err := client.queryHandleResponse(httpResp)
	return resp, err
}

// queryCreateRequest creates the Query request.
func (client *Client) queryCreateRequest(ctx context.Context, options *ClientQueryOptions) (*policy.Request, error) {
	urlPath := "/Tables"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Format != nil {
		reqQP.Set("$format", string(*options.Format))
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Select != nil {
		reqQP.Set("$select", *options.Select)
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.NextTableName != nil {
		reqQP.Set("NextTableName", *options.NextTableName)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["x-ms-version"] = []string{"2019-02-02"}
	if options != nil && options.RequestID != nil {
		req.Raw().Header["x-ms-client-request-id"] = []string{*options.RequestID}
	}
	req.Raw().Header["DataServiceVersion"] = []string{"3.0"}
	req.Raw().Header["Accept"] = []string{"application/json;odata=minimalmetadata"}
	return req, nil
}

// queryHandleResponse handles the Query response.
func (client *Client) queryHandleResponse(resp *http.Response) (ClientQueryResponse, error) {
	result := ClientQueryResponse{}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return ClientQueryResponse{}, err
		}
		result.Date = &date
	}
	if val := resp.Header.Get("x-ms-continuation-NextTableName"); val != "" {
		result.XMSContinuationNextTableName = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.QueryResponse); err != nil {
		return ClientQueryResponse{}, err
	}
	return result, nil
}

// QueryEntities - Queries entities in a table.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-02-02
//   - table - The name of the table.
//   - options - ClientQueryEntitiesOptions contains the optional parameters for the Client.QueryEntities method.
func (client *Client) QueryEntities(ctx context.Context, table string, options *ClientQueryEntitiesOptions) (ClientQueryEntitiesResponse, error) {
	var err error
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "Client.QueryEntities")
	req, err := client.queryEntitiesCreateRequest(ctx, table, options)
	if err != nil {
		return ClientQueryEntitiesResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ClientQueryEntitiesResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ClientQueryEntitiesResponse{}, err
	}
	resp, err := client.queryEntitiesHandleResponse(httpResp)
	return resp, err
}

// queryEntitiesCreateRequest creates the QueryEntities request.
func (client *Client) queryEntitiesCreateRequest(ctx context.Context, table string, options *ClientQueryEntitiesOptions) (*policy.Request, error) {
	urlPath := "/{table}()"
	if table == "" {
		return nil, errors.New("parameter table cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{table}", url.PathEscape(table))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	if options != nil && options.Format != nil {
		reqQP.Set("$format", string(*options.Format))
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Select != nil {
		reqQP.Set("$select", *options.Select)
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.NextPartitionKey != nil {
		reqQP.Set("NextPartitionKey", *options.NextPartitionKey)
	}
	if options != nil && options.NextRowKey != nil {
		reqQP.Set("NextRowKey", *options.NextRowKey)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["x-ms-version"] = []string{"2019-02-02"}
	if options != nil && options.RequestID != nil {
		req.Raw().Header["x-ms-client-request-id"] = []string{*options.RequestID}
	}
	req.Raw().Header["DataServiceVersion"] = []string{"3.0"}
	req.Raw().Header["Accept"] = []string{"application/json;odata=minimalmetadata"}
	return req, nil
}

// queryEntitiesHandleResponse handles the QueryEntities response.
func (client *Client) queryEntitiesHandleResponse(resp *http.Response) (ClientQueryEntitiesResponse, error) {
	result := ClientQueryEntitiesResponse{}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return ClientQueryEntitiesResponse{}, err
		}
		result.Date = &date
	}
	if val := resp.Header.Get("x-ms-continuation-NextPartitionKey"); val != "" {
		result.XMSContinuationNextPartitionKey = &val
	}
	if val := resp.Header.Get("x-ms-continuation-NextRowKey"); val != "" {
		result.XMSContinuationNextRowKey = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.EntityQueryResponse); err != nil {
		return ClientQueryEntitiesResponse{}, err
	}
	return result, nil
}

// QueryEntityWithPartitionAndRowKey - Queries a single entity in a table.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-02-02
//   - table - The name of the table.
//   - partitionKey - The partition key of the entity.
//   - rowKey - The row key of the entity.
//   - options - ClientQueryEntityWithPartitionAndRowKeyOptions contains the optional parameters for the Client.QueryEntityWithPartitionAndRowKey
//     method.
func (client *Client) QueryEntityWithPartitionAndRowKey(ctx context.Context, table string, partitionKey string, rowKey string, options *ClientQueryEntityWithPartitionAndRowKeyOptions) (ClientQueryEntityWithPartitionAndRowKeyResponse, error) {
	var err error
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "Client.QueryEntityWithPartitionAndRowKey")
	req, err := client.queryEntityWithPartitionAndRowKeyCreateRequest(ctx, table, partitionKey, rowKey, options)
	if err != nil {
		return ClientQueryEntityWithPartitionAndRowKeyResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ClientQueryEntityWithPartitionAndRowKeyResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ClientQueryEntityWithPartitionAndRowKeyResponse{}, err
	}
	resp, err := client.queryEntityWithPartitionAndRowKeyHandleResponse(httpResp)
	return resp, err
}

// queryEntityWithPartitionAndRowKeyCreateRequest creates the QueryEntityWithPartitionAndRowKey request.
func (client *Client) queryEntityWithPartitionAndRowKeyCreateRequest(ctx context.Context, table string, partitionKey string, rowKey string, options *ClientQueryEntityWithPartitionAndRowKeyOptions) (*policy.Request, error) {
	urlPath := "/{table}(PartitionKey='{partitionKey}',RowKey='{rowKey}')"
	if table == "" {
		return nil, errors.New("parameter table cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{table}", url.PathEscape(table))
	if partitionKey == "" {
		return nil, errors.New("parameter partitionKey cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{partitionKey}", url.PathEscape(partitionKey))
	if rowKey == "" {
		return nil, errors.New("parameter rowKey cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{rowKey}", url.PathEscape(rowKey))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	if options != nil && options.Format != nil {
		reqQP.Set("$format", string(*options.Format))
	}
	if options != nil && options.Select != nil {
		reqQP.Set("$select", *options.Select)
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["x-ms-version"] = []string{"2019-02-02"}
	if options != nil && options.RequestID != nil {
		req.Raw().Header["x-ms-client-request-id"] = []string{*options.RequestID}
	}
	req.Raw().Header["DataServiceVersion"] = []string{"3.0"}
	req.Raw().Header["Accept"] = []string{"application/json;odata=minimalmetadata"}
	return req, nil
}

// queryEntityWithPartitionAndRowKeyHandleResponse handles the QueryEntityWithPartitionAndRowKey response.
func (client *Client) queryEntityWithPartitionAndRowKeyHandleResponse(resp *http.Response) (ClientQueryEntityWithPartitionAndRowKeyResponse, error) {
	result := ClientQueryEntityWithPartitionAndRowKeyResponse{}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return ClientQueryEntityWithPartitionAndRowKeyResponse{}, err
		}
		result.Date = &date
	}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if val := resp.Header.Get("x-ms-continuation-NextPartitionKey"); val != "" {
		result.XMSContinuationNextPartitionKey = &val
	}
	if val := resp.Header.Get("x-ms-continuation-NextRowKey"); val != "" {
		result.XMSContinuationNextRowKey = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return ClientQueryEntityWithPartitionAndRowKeyResponse{}, err
	}
	return result, nil
}

// SetAccessPolicy - Sets stored access policies for the table that may be used with Shared Access Signatures.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-02-02
//   - table - The name of the table.
//   - tableACL - The acls for the table.
//   - options - ClientSetAccessPolicyOptions contains the optional parameters for the Client.SetAccessPolicy method.
func (client *Client) SetAccessPolicy(ctx context.Context, table string, tableACL []*SignedIdentifier, options *ClientSetAccessPolicyOptions) (ClientSetAccessPolicyResponse, error) {
	var err error
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "Client.SetAccessPolicy")
	req, err := client.setAccessPolicyCreateRequest(ctx, table, tableACL, options)
	if err != nil {
		return ClientSetAccessPolicyResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ClientSetAccessPolicyResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ClientSetAccessPolicyResponse{}, err
	}
	resp, err := client.setAccessPolicyHandleResponse(httpResp)
	return resp, err
}

// setAccessPolicyCreateRequest creates the SetAccessPolicy request.
func (client *Client) setAccessPolicyCreateRequest(ctx context.Context, table string, tableACL []*SignedIdentifier, options *ClientSetAccessPolicyOptions) (*policy.Request, error) {
	urlPath := "/{table}"
	if table == "" {
		return nil, errors.New("parameter table cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{table}", url.PathEscape(table))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	reqQP.Set("comp", "acl")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["x-ms-version"] = []string{"2019-02-02"}
	if options != nil && options.RequestID != nil {
		req.Raw().Header["x-ms-client-request-id"] = []string{*options.RequestID}
	}
	req.Raw().Header["Accept"] = []string{"application/xml"}
	type wrapper struct {
		XMLName  xml.Name             `xml:"SignedIdentifiers"`
		TableACL *[]*SignedIdentifier `xml:"SignedIdentifier"`
	}
	if err := runtime.MarshalAsXML(req, wrapper{TableACL: &tableACL}); err != nil {
		return nil, err
	}
	return req, nil
}

// setAccessPolicyHandleResponse handles the SetAccessPolicy response.
func (client *Client) setAccessPolicyHandleResponse(resp *http.Response) (ClientSetAccessPolicyResponse, error) {
	result := ClientSetAccessPolicyResponse{}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return ClientSetAccessPolicyResponse{}, err
		}
		result.Date = &date
	}
	return result, nil
}

// UpdateEntity - Update entity in a table.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-02-02
//   - table - The name of the table.
//   - partitionKey - The partition key of the entity.
//   - rowKey - The row key of the entity.
//   - tableEntityProperties - The properties for the table entity.
//   - options - ClientUpdateEntityOptions contains the optional parameters for the Client.UpdateEntity method.
func (client *Client) UpdateEntity(ctx context.Context, table string, partitionKey string, rowKey string, tableEntityProperties map[string]any, options *ClientUpdateEntityOptions) (ClientUpdateEntityResponse, error) {
	var err error
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "Client.UpdateEntity")
	req, err := client.updateEntityCreateRequest(ctx, table, partitionKey, rowKey, tableEntityProperties, options)
	if err != nil {
		return ClientUpdateEntityResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ClientUpdateEntityResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ClientUpdateEntityResponse{}, err
	}
	resp, err := client.updateEntityHandleResponse(httpResp)
	return resp, err
}

// updateEntityCreateRequest creates the UpdateEntity request.
func (client *Client) updateEntityCreateRequest(ctx context.Context, table string, partitionKey string, rowKey string, tableEntityProperties map[string]any, options *ClientUpdateEntityOptions) (*policy.Request, error) {
	urlPath := "/{table}(PartitionKey='{partitionKey}',RowKey='{rowKey}')"
	if table == "" {
		return nil, errors.New("parameter table cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{table}", url.PathEscape(table))
	if partitionKey == "" {
		return nil, errors.New("parameter partitionKey cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{partitionKey}", url.PathEscape(partitionKey))
	if rowKey == "" {
		return nil, errors.New("parameter rowKey cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{rowKey}", url.PathEscape(rowKey))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	if options != nil && options.Format != nil {
		reqQP.Set("$format", string(*options.Format))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["x-ms-version"] = []string{"2019-02-02"}
	if options != nil && options.RequestID != nil {
		req.Raw().Header["x-ms-client-request-id"] = []string{*options.RequestID}
	}
	req.Raw().Header["DataServiceVersion"] = []string{"3.0"}
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, tableEntityProperties); err != nil {
		return nil, err
	}
	return req, nil
}

// updateEntityHandleResponse handles the UpdateEntity response.
func (client *Client) updateEntityHandleResponse(resp *http.Response) (ClientUpdateEntityResponse, error) {
	result := ClientUpdateEntityResponse{}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return ClientUpdateEntityResponse{}, err
		}
		result.Date = &date
	}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	return result, nil
}
