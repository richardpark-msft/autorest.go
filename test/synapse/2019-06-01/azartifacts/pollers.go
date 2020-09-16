// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azartifacts

// TODO replace this code with data-plane specific definitions
import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// CreateDataFlowDebugSessionResponsePoller provides polling facilities until the operation completes
type CreateDataFlowDebugSessionResponsePoller interface {
	Done() bool
	Poll(ctx context.Context) (*http.Response, error)
	FinalResponse(ctx context.Context) (*CreateDataFlowDebugSessionResponseResponse, error)
	ResumeToken() (string, error)
}

type createDataFlowDebugSessionResponsePoller struct {
	// the client for making the request
	pipeline azcore.Pipeline
}

// Done returns true if there was an error or polling has reached a terminal state
func (p *createDataFlowDebugSessionResponsePoller) Done() bool {
	return false
}

// Poll will send poll the service endpoint and return an http.Response or error received from the service
func (p *createDataFlowDebugSessionResponsePoller) Poll(ctx context.Context) (*http.Response, error) {
	return nil, nil
}

func (p *createDataFlowDebugSessionResponsePoller) FinalResponse(ctx context.Context) (*CreateDataFlowDebugSessionResponseResponse, error) {
	return nil, nil
}

// ResumeToken generates the string token that can be used with the ResumeCreateDataFlowDebugSessionResponsePoller method
// on the client to create a new poller from the data held in the current poller type
func (p *createDataFlowDebugSessionResponsePoller) ResumeToken() (string, error) {
	return "", nil
}

// DataFlowDebugCommandResponsePoller provides polling facilities until the operation completes
type DataFlowDebugCommandResponsePoller interface {
	Done() bool
	Poll(ctx context.Context) (*http.Response, error)
	FinalResponse(ctx context.Context) (*DataFlowDebugCommandResponseResponse, error)
	ResumeToken() (string, error)
}

type dataFlowDebugCommandResponsePoller struct {
	// the client for making the request
	pipeline azcore.Pipeline
}

// Done returns true if there was an error or polling has reached a terminal state
func (p *dataFlowDebugCommandResponsePoller) Done() bool {
	return false
}

// Poll will send poll the service endpoint and return an http.Response or error received from the service
func (p *dataFlowDebugCommandResponsePoller) Poll(ctx context.Context) (*http.Response, error) {
	return nil, nil
}

func (p *dataFlowDebugCommandResponsePoller) FinalResponse(ctx context.Context) (*DataFlowDebugCommandResponseResponse, error) {
	return nil, nil
}

// ResumeToken generates the string token that can be used with the ResumeDataFlowDebugCommandResponsePoller method
// on the client to create a new poller from the data held in the current poller type
func (p *dataFlowDebugCommandResponsePoller) ResumeToken() (string, error) {
	return "", nil
}

// DataFlowResourcePoller provides polling facilities until the operation completes
type DataFlowResourcePoller interface {
	Done() bool
	Poll(ctx context.Context) (*http.Response, error)
	FinalResponse(ctx context.Context) (*DataFlowResourceResponse, error)
	ResumeToken() (string, error)
}

type dataFlowResourcePoller struct {
	// the client for making the request
	pipeline azcore.Pipeline
}

// Done returns true if there was an error or polling has reached a terminal state
func (p *dataFlowResourcePoller) Done() bool {
	return false
}

// Poll will send poll the service endpoint and return an http.Response or error received from the service
func (p *dataFlowResourcePoller) Poll(ctx context.Context) (*http.Response, error) {
	return nil, nil
}

func (p *dataFlowResourcePoller) FinalResponse(ctx context.Context) (*DataFlowResourceResponse, error) {
	return nil, nil
}

// ResumeToken generates the string token that can be used with the ResumeDataFlowResourcePoller method
// on the client to create a new poller from the data held in the current poller type
func (p *dataFlowResourcePoller) ResumeToken() (string, error) {
	return "", nil
}

// DatasetResourcePoller provides polling facilities until the operation completes
type DatasetResourcePoller interface {
	Done() bool
	Poll(ctx context.Context) (*http.Response, error)
	FinalResponse(ctx context.Context) (*DatasetResourceResponse, error)
	ResumeToken() (string, error)
}

type datasetResourcePoller struct {
	// the client for making the request
	pipeline azcore.Pipeline
}

// Done returns true if there was an error or polling has reached a terminal state
func (p *datasetResourcePoller) Done() bool {
	return false
}

// Poll will send poll the service endpoint and return an http.Response or error received from the service
func (p *datasetResourcePoller) Poll(ctx context.Context) (*http.Response, error) {
	return nil, nil
}

func (p *datasetResourcePoller) FinalResponse(ctx context.Context) (*DatasetResourceResponse, error) {
	return nil, nil
}

// ResumeToken generates the string token that can be used with the ResumeDatasetResourcePoller method
// on the client to create a new poller from the data held in the current poller type
func (p *datasetResourcePoller) ResumeToken() (string, error) {
	return "", nil
}

// HTTPPoller provides polling facilities until the operation completes
type HTTPPoller interface {
	Done() bool
	Poll(ctx context.Context) (*http.Response, error)
	FinalResponse() *http.Response
	ResumeToken() (string, error)
}

type httpPoller struct {
	// the client for making the request
	pipeline azcore.Pipeline
}

// Done returns true if there was an error or polling has reached a terminal state
func (p *httpPoller) Done() bool {
	return false
}

// Poll will send poll the service endpoint and return an http.Response or error received from the service
func (p *httpPoller) Poll(ctx context.Context) (*http.Response, error) {
	return nil, nil
}

func (p *httpPoller) FinalResponse() *http.Response {
	return nil
}

// ResumeToken generates the string token that can be used with the ResumeHTTPPoller method
// on the client to create a new poller from the data held in the current poller type
func (p *httpPoller) ResumeToken() (string, error) {
	return "", nil
}

// LinkedServiceResourcePoller provides polling facilities until the operation completes
type LinkedServiceResourcePoller interface {
	Done() bool
	Poll(ctx context.Context) (*http.Response, error)
	FinalResponse(ctx context.Context) (*LinkedServiceResourceResponse, error)
	ResumeToken() (string, error)
}

type linkedServiceResourcePoller struct {
	// the client for making the request
	pipeline azcore.Pipeline
}

// Done returns true if there was an error or polling has reached a terminal state
func (p *linkedServiceResourcePoller) Done() bool {
	return false
}

// Poll will send poll the service endpoint and return an http.Response or error received from the service
func (p *linkedServiceResourcePoller) Poll(ctx context.Context) (*http.Response, error) {
	return nil, nil
}

func (p *linkedServiceResourcePoller) FinalResponse(ctx context.Context) (*LinkedServiceResourceResponse, error) {
	return nil, nil
}

// ResumeToken generates the string token that can be used with the ResumeLinkedServiceResourcePoller method
// on the client to create a new poller from the data held in the current poller type
func (p *linkedServiceResourcePoller) ResumeToken() (string, error) {
	return "", nil
}

// NotebookResourcePoller provides polling facilities until the operation completes
type NotebookResourcePoller interface {
	Done() bool
	Poll(ctx context.Context) (*http.Response, error)
	FinalResponse(ctx context.Context) (*NotebookResourceResponse, error)
	ResumeToken() (string, error)
}

type notebookResourcePoller struct {
	// the client for making the request
	pipeline azcore.Pipeline
}

// Done returns true if there was an error or polling has reached a terminal state
func (p *notebookResourcePoller) Done() bool {
	return false
}

// Poll will send poll the service endpoint and return an http.Response or error received from the service
func (p *notebookResourcePoller) Poll(ctx context.Context) (*http.Response, error) {
	return nil, nil
}

func (p *notebookResourcePoller) FinalResponse(ctx context.Context) (*NotebookResourceResponse, error) {
	return nil, nil
}

// ResumeToken generates the string token that can be used with the ResumeNotebookResourcePoller method
// on the client to create a new poller from the data held in the current poller type
func (p *notebookResourcePoller) ResumeToken() (string, error) {
	return "", nil
}

// PipelineResourcePoller provides polling facilities until the operation completes
type PipelineResourcePoller interface {
	Done() bool
	Poll(ctx context.Context) (*http.Response, error)
	FinalResponse(ctx context.Context) (*PipelineResourceResponse, error)
	ResumeToken() (string, error)
}

type pipelineResourcePoller struct {
	// the client for making the request
	pipeline azcore.Pipeline
}

// Done returns true if there was an error or polling has reached a terminal state
func (p *pipelineResourcePoller) Done() bool {
	return false
}

// Poll will send poll the service endpoint and return an http.Response or error received from the service
func (p *pipelineResourcePoller) Poll(ctx context.Context) (*http.Response, error) {
	return nil, nil
}

func (p *pipelineResourcePoller) FinalResponse(ctx context.Context) (*PipelineResourceResponse, error) {
	return nil, nil
}

// ResumeToken generates the string token that can be used with the ResumePipelineResourcePoller method
// on the client to create a new poller from the data held in the current poller type
func (p *pipelineResourcePoller) ResumeToken() (string, error) {
	return "", nil
}

// SparkBatchJobPoller provides polling facilities until the operation completes
type SparkBatchJobPoller interface {
	Done() bool
	Poll(ctx context.Context) (*http.Response, error)
	FinalResponse(ctx context.Context) (*SparkBatchJobResponse, error)
	ResumeToken() (string, error)
}

type sparkBatchJobPoller struct {
	// the client for making the request
	pipeline azcore.Pipeline
}

// Done returns true if there was an error or polling has reached a terminal state
func (p *sparkBatchJobPoller) Done() bool {
	return false
}

// Poll will send poll the service endpoint and return an http.Response or error received from the service
func (p *sparkBatchJobPoller) Poll(ctx context.Context) (*http.Response, error) {
	return nil, nil
}

func (p *sparkBatchJobPoller) FinalResponse(ctx context.Context) (*SparkBatchJobResponse, error) {
	return nil, nil
}

// ResumeToken generates the string token that can be used with the ResumeSparkBatchJobPoller method
// on the client to create a new poller from the data held in the current poller type
func (p *sparkBatchJobPoller) ResumeToken() (string, error) {
	return "", nil
}

// TriggerResourcePoller provides polling facilities until the operation completes
type TriggerResourcePoller interface {
	Done() bool
	Poll(ctx context.Context) (*http.Response, error)
	FinalResponse(ctx context.Context) (*TriggerResourceResponse, error)
	ResumeToken() (string, error)
}

type triggerResourcePoller struct {
	// the client for making the request
	pipeline azcore.Pipeline
}

// Done returns true if there was an error or polling has reached a terminal state
func (p *triggerResourcePoller) Done() bool {
	return false
}

// Poll will send poll the service endpoint and return an http.Response or error received from the service
func (p *triggerResourcePoller) Poll(ctx context.Context) (*http.Response, error) {
	return nil, nil
}

func (p *triggerResourcePoller) FinalResponse(ctx context.Context) (*TriggerResourceResponse, error) {
	return nil, nil
}

// ResumeToken generates the string token that can be used with the ResumeTriggerResourcePoller method
// on the client to create a new poller from the data held in the current poller type
func (p *triggerResourcePoller) ResumeToken() (string, error) {
	return "", nil
}

// TriggerSubscriptionOperationStatusPoller provides polling facilities until the operation completes
type TriggerSubscriptionOperationStatusPoller interface {
	Done() bool
	Poll(ctx context.Context) (*http.Response, error)
	FinalResponse(ctx context.Context) (*TriggerSubscriptionOperationStatusResponse, error)
	ResumeToken() (string, error)
}

type triggerSubscriptionOperationStatusPoller struct {
	// the client for making the request
	pipeline azcore.Pipeline
}

// Done returns true if there was an error or polling has reached a terminal state
func (p *triggerSubscriptionOperationStatusPoller) Done() bool {
	return false
}

// Poll will send poll the service endpoint and return an http.Response or error received from the service
func (p *triggerSubscriptionOperationStatusPoller) Poll(ctx context.Context) (*http.Response, error) {
	return nil, nil
}

func (p *triggerSubscriptionOperationStatusPoller) FinalResponse(ctx context.Context) (*TriggerSubscriptionOperationStatusResponse, error) {
	return nil, nil
}

// ResumeToken generates the string token that can be used with the ResumeTriggerSubscriptionOperationStatusPoller method
// on the client to create a new poller from the data held in the current poller type
func (p *triggerSubscriptionOperationStatusPoller) ResumeToken() (string, error) {
	return "", nil
}
