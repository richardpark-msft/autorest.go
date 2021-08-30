//go:build go1.16
// +build go1.16

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
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type pipelineClient struct {
	con *connection
}

// BeginCreateOrUpdatePipeline - Creates or updates a pipeline.
// If the operation fails it returns the *CloudError error type.
func (client *pipelineClient) BeginCreateOrUpdatePipeline(ctx context.Context, pipelineName string, pipeline PipelineResource, options *PipelineBeginCreateOrUpdatePipelineOptions) (PipelineCreateOrUpdatePipelinePollerResponse, error) {
	resp, err := client.createOrUpdatePipeline(ctx, pipelineName, pipeline, options)
	if err != nil {
		return PipelineCreateOrUpdatePipelinePollerResponse{}, err
	}
	result := PipelineCreateOrUpdatePipelinePollerResponse{
		RawResponse: resp,
	}
	pt, err := runtime.NewPoller("pipelineClient.CreateOrUpdatePipeline", resp, client.con.Pipeline(), client.createOrUpdatePipelineHandleError)
	if err != nil {
		return PipelineCreateOrUpdatePipelinePollerResponse{}, err
	}
	result.Poller = &PipelineCreateOrUpdatePipelinePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdatePipeline - Creates or updates a pipeline.
// If the operation fails it returns the *CloudError error type.
func (client *pipelineClient) createOrUpdatePipeline(ctx context.Context, pipelineName string, pipeline PipelineResource, options *PipelineBeginCreateOrUpdatePipelineOptions) (*http.Response, error) {
	req, err := client.createOrUpdatePipelineCreateRequest(ctx, pipelineName, pipeline, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.createOrUpdatePipelineHandleError(resp)
	}
	return resp, nil
}

// createOrUpdatePipelineCreateRequest creates the CreateOrUpdatePipeline request.
func (client *pipelineClient) createOrUpdatePipelineCreateRequest(ctx context.Context, pipelineName string, pipeline PipelineResource, options *PipelineBeginCreateOrUpdatePipelineOptions) (*policy.Request, error) {
	urlPath := "/pipelines/{pipelineName}"
	if pipelineName == "" {
		return nil, errors.New("parameter pipelineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pipelineName}", url.PathEscape(pipelineName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header.Set("If-Match", *options.IfMatch)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, pipeline)
}

// createOrUpdatePipelineHandleError handles the CreateOrUpdatePipeline error response.
func (client *pipelineClient) createOrUpdatePipelineHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// CreatePipelineRun - Creates a run of a pipeline.
// If the operation fails it returns the *CloudError error type.
func (client *pipelineClient) CreatePipelineRun(ctx context.Context, pipelineName string, options *PipelineCreatePipelineRunOptions) (PipelineCreatePipelineRunResponse, error) {
	req, err := client.createPipelineRunCreateRequest(ctx, pipelineName, options)
	if err != nil {
		return PipelineCreatePipelineRunResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return PipelineCreatePipelineRunResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return PipelineCreatePipelineRunResponse{}, client.createPipelineRunHandleError(resp)
	}
	return client.createPipelineRunHandleResponse(resp)
}

// createPipelineRunCreateRequest creates the CreatePipelineRun request.
func (client *pipelineClient) createPipelineRunCreateRequest(ctx context.Context, pipelineName string, options *PipelineCreatePipelineRunOptions) (*policy.Request, error) {
	urlPath := "/pipelines/{pipelineName}/createRun"
	if pipelineName == "" {
		return nil, errors.New("parameter pipelineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pipelineName}", url.PathEscape(pipelineName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	if options != nil && options.ReferencePipelineRunID != nil {
		reqQP.Set("referencePipelineRunId", *options.ReferencePipelineRunID)
	}
	if options != nil && options.IsRecovery != nil {
		reqQP.Set("isRecovery", strconv.FormatBool(*options.IsRecovery))
	}
	if options != nil && options.StartActivityName != nil {
		reqQP.Set("startActivityName", *options.StartActivityName)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	if options != nil && options.Parameters != nil {
		return req, runtime.MarshalAsJSON(req, options.Parameters)
	}
	return req, nil
}

// createPipelineRunHandleResponse handles the CreatePipelineRun response.
func (client *pipelineClient) createPipelineRunHandleResponse(resp *http.Response) (PipelineCreatePipelineRunResponse, error) {
	result := PipelineCreatePipelineRunResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.CreateRunResponse); err != nil {
		return PipelineCreatePipelineRunResponse{}, err
	}
	return result, nil
}

// createPipelineRunHandleError handles the CreatePipelineRun error response.
func (client *pipelineClient) createPipelineRunHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginDeletePipeline - Deletes a pipeline.
// If the operation fails it returns the *CloudError error type.
func (client *pipelineClient) BeginDeletePipeline(ctx context.Context, pipelineName string, options *PipelineBeginDeletePipelineOptions) (PipelineDeletePipelinePollerResponse, error) {
	resp, err := client.deletePipeline(ctx, pipelineName, options)
	if err != nil {
		return PipelineDeletePipelinePollerResponse{}, err
	}
	result := PipelineDeletePipelinePollerResponse{
		RawResponse: resp,
	}
	pt, err := runtime.NewPoller("pipelineClient.DeletePipeline", resp, client.con.Pipeline(), client.deletePipelineHandleError)
	if err != nil {
		return PipelineDeletePipelinePollerResponse{}, err
	}
	result.Poller = &PipelineDeletePipelinePoller{
		pt: pt,
	}
	return result, nil
}

// DeletePipeline - Deletes a pipeline.
// If the operation fails it returns the *CloudError error type.
func (client *pipelineClient) deletePipeline(ctx context.Context, pipelineName string, options *PipelineBeginDeletePipelineOptions) (*http.Response, error) {
	req, err := client.deletePipelineCreateRequest(ctx, pipelineName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deletePipelineHandleError(resp)
	}
	return resp, nil
}

// deletePipelineCreateRequest creates the DeletePipeline request.
func (client *pipelineClient) deletePipelineCreateRequest(ctx context.Context, pipelineName string, options *PipelineBeginDeletePipelineOptions) (*policy.Request, error) {
	urlPath := "/pipelines/{pipelineName}"
	if pipelineName == "" {
		return nil, errors.New("parameter pipelineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pipelineName}", url.PathEscape(pipelineName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deletePipelineHandleError handles the DeletePipeline error response.
func (client *pipelineClient) deletePipelineHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// GetPipeline - Gets a pipeline.
// If the operation fails it returns the *CloudError error type.
func (client *pipelineClient) GetPipeline(ctx context.Context, pipelineName string, options *PipelineGetPipelineOptions) (PipelineGetPipelineResponse, error) {
	req, err := client.getPipelineCreateRequest(ctx, pipelineName, options)
	if err != nil {
		return PipelineGetPipelineResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return PipelineGetPipelineResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNotModified) {
		return PipelineGetPipelineResponse{}, client.getPipelineHandleError(resp)
	}
	return client.getPipelineHandleResponse(resp)
}

// getPipelineCreateRequest creates the GetPipeline request.
func (client *pipelineClient) getPipelineCreateRequest(ctx context.Context, pipelineName string, options *PipelineGetPipelineOptions) (*policy.Request, error) {
	urlPath := "/pipelines/{pipelineName}"
	if pipelineName == "" {
		return nil, errors.New("parameter pipelineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pipelineName}", url.PathEscape(pipelineName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfNoneMatch != nil {
		req.Raw().Header.Set("If-None-Match", *options.IfNoneMatch)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getPipelineHandleResponse handles the GetPipeline response.
func (client *pipelineClient) getPipelineHandleResponse(resp *http.Response) (PipelineGetPipelineResponse, error) {
	result := PipelineGetPipelineResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.PipelineResource); err != nil {
		return PipelineGetPipelineResponse{}, err
	}
	return result, nil
}

// getPipelineHandleError handles the GetPipeline error response.
func (client *pipelineClient) getPipelineHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// GetPipelinesByWorkspace - Lists pipelines.
// If the operation fails it returns the *CloudError error type.
func (client *pipelineClient) GetPipelinesByWorkspace(options *PipelineGetPipelinesByWorkspaceOptions) *PipelineGetPipelinesByWorkspacePager {
	return &PipelineGetPipelinesByWorkspacePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.getPipelinesByWorkspaceCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp PipelineGetPipelinesByWorkspaceResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.PipelineListResponse.NextLink)
		},
	}
}

// getPipelinesByWorkspaceCreateRequest creates the GetPipelinesByWorkspace request.
func (client *pipelineClient) getPipelinesByWorkspaceCreateRequest(ctx context.Context, options *PipelineGetPipelinesByWorkspaceOptions) (*policy.Request, error) {
	urlPath := "/pipelines"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getPipelinesByWorkspaceHandleResponse handles the GetPipelinesByWorkspace response.
func (client *pipelineClient) getPipelinesByWorkspaceHandleResponse(resp *http.Response) (PipelineGetPipelinesByWorkspaceResponse, error) {
	result := PipelineGetPipelinesByWorkspaceResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.PipelineListResponse); err != nil {
		return PipelineGetPipelinesByWorkspaceResponse{}, err
	}
	return result, nil
}

// getPipelinesByWorkspaceHandleError handles the GetPipelinesByWorkspace error response.
func (client *pipelineClient) getPipelinesByWorkspaceHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginRenamePipeline - Renames a pipeline.
// If the operation fails it returns the *CloudError error type.
func (client *pipelineClient) BeginRenamePipeline(ctx context.Context, pipelineName string, request ArtifactRenameRequest, options *PipelineBeginRenamePipelineOptions) (PipelineRenamePipelinePollerResponse, error) {
	resp, err := client.renamePipeline(ctx, pipelineName, request, options)
	if err != nil {
		return PipelineRenamePipelinePollerResponse{}, err
	}
	result := PipelineRenamePipelinePollerResponse{
		RawResponse: resp,
	}
	pt, err := runtime.NewPoller("pipelineClient.RenamePipeline", resp, client.con.Pipeline(), client.renamePipelineHandleError)
	if err != nil {
		return PipelineRenamePipelinePollerResponse{}, err
	}
	result.Poller = &PipelineRenamePipelinePoller{
		pt: pt,
	}
	return result, nil
}

// RenamePipeline - Renames a pipeline.
// If the operation fails it returns the *CloudError error type.
func (client *pipelineClient) renamePipeline(ctx context.Context, pipelineName string, request ArtifactRenameRequest, options *PipelineBeginRenamePipelineOptions) (*http.Response, error) {
	req, err := client.renamePipelineCreateRequest(ctx, pipelineName, request, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.renamePipelineHandleError(resp)
	}
	return resp, nil
}

// renamePipelineCreateRequest creates the RenamePipeline request.
func (client *pipelineClient) renamePipelineCreateRequest(ctx context.Context, pipelineName string, request ArtifactRenameRequest, options *PipelineBeginRenamePipelineOptions) (*policy.Request, error) {
	urlPath := "/pipelines/{pipelineName}/rename"
	if pipelineName == "" {
		return nil, errors.New("parameter pipelineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pipelineName}", url.PathEscape(pipelineName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, request)
}

// renamePipelineHandleError handles the RenamePipeline error response.
func (client *pipelineClient) renamePipelineHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
