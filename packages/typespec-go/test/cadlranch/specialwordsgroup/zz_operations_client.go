//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package specialwordsgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// OperationsClient contains the methods for the SpecialWords group.
// Don't use this type directly, use a constructor function instead.
type OperationsClient struct {
	internal *azcore.Client
}

func (client *OperationsClient) And(ctx context.Context, options *OperationsClientAndOptions) (OperationsClientAndResponse, error) {
	var err error
	req, err := client.andCreateRequest(ctx, options)
	if err != nil {
		return OperationsClientAndResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OperationsClientAndResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return OperationsClientAndResponse{}, err
	}
	return OperationsClientAndResponse{}, nil
}

// andCreateRequest creates the And request.
func (client *OperationsClient) andCreateRequest(ctx context.Context, options *OperationsClientAndOptions) (*policy.Request, error) {
	urlPath := "/special-words/operations/and"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (client *OperationsClient) As(ctx context.Context, options *OperationsClientAsOptions) (OperationsClientAsResponse, error) {
	var err error
	req, err := client.asCreateRequest(ctx, options)
	if err != nil {
		return OperationsClientAsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OperationsClientAsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return OperationsClientAsResponse{}, err
	}
	return OperationsClientAsResponse{}, nil
}

// asCreateRequest creates the As request.
func (client *OperationsClient) asCreateRequest(ctx context.Context, options *OperationsClientAsOptions) (*policy.Request, error) {
	urlPath := "/special-words/operations/as"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (client *OperationsClient) Assert(ctx context.Context, options *OperationsClientAssertOptions) (OperationsClientAssertResponse, error) {
	var err error
	req, err := client.assertCreateRequest(ctx, options)
	if err != nil {
		return OperationsClientAssertResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OperationsClientAssertResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return OperationsClientAssertResponse{}, err
	}
	return OperationsClientAssertResponse{}, nil
}

// assertCreateRequest creates the Assert request.
func (client *OperationsClient) assertCreateRequest(ctx context.Context, options *OperationsClientAssertOptions) (*policy.Request, error) {
	urlPath := "/special-words/operations/assert"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (client *OperationsClient) Async(ctx context.Context, options *OperationsClientAsyncOptions) (OperationsClientAsyncResponse, error) {
	var err error
	req, err := client.asyncCreateRequest(ctx, options)
	if err != nil {
		return OperationsClientAsyncResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OperationsClientAsyncResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return OperationsClientAsyncResponse{}, err
	}
	return OperationsClientAsyncResponse{}, nil
}

// asyncCreateRequest creates the Async request.
func (client *OperationsClient) asyncCreateRequest(ctx context.Context, options *OperationsClientAsyncOptions) (*policy.Request, error) {
	urlPath := "/special-words/operations/async"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (client *OperationsClient) Await(ctx context.Context, options *OperationsClientAwaitOptions) (OperationsClientAwaitResponse, error) {
	var err error
	req, err := client.awaitCreateRequest(ctx, options)
	if err != nil {
		return OperationsClientAwaitResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OperationsClientAwaitResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return OperationsClientAwaitResponse{}, err
	}
	return OperationsClientAwaitResponse{}, nil
}

// awaitCreateRequest creates the Await request.
func (client *OperationsClient) awaitCreateRequest(ctx context.Context, options *OperationsClientAwaitOptions) (*policy.Request, error) {
	urlPath := "/special-words/operations/await"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (client *OperationsClient) Break(ctx context.Context, options *OperationsClientBreakOptions) (OperationsClientBreakResponse, error) {
	var err error
	req, err := client.breakCreateRequest(ctx, options)
	if err != nil {
		return OperationsClientBreakResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OperationsClientBreakResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return OperationsClientBreakResponse{}, err
	}
	return OperationsClientBreakResponse{}, nil
}

// breakCreateRequest creates the Break request.
func (client *OperationsClient) breakCreateRequest(ctx context.Context, options *OperationsClientBreakOptions) (*policy.Request, error) {
	urlPath := "/special-words/operations/break"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (client *OperationsClient) Class(ctx context.Context, options *OperationsClientClassOptions) (OperationsClientClassResponse, error) {
	var err error
	req, err := client.classCreateRequest(ctx, options)
	if err != nil {
		return OperationsClientClassResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OperationsClientClassResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return OperationsClientClassResponse{}, err
	}
	return OperationsClientClassResponse{}, nil
}

// classCreateRequest creates the Class request.
func (client *OperationsClient) classCreateRequest(ctx context.Context, options *OperationsClientClassOptions) (*policy.Request, error) {
	urlPath := "/special-words/operations/class"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (client *OperationsClient) Constructor(ctx context.Context, options *OperationsClientConstructorOptions) (OperationsClientConstructorResponse, error) {
	var err error
	req, err := client.constructorCreateRequest(ctx, options)
	if err != nil {
		return OperationsClientConstructorResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OperationsClientConstructorResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return OperationsClientConstructorResponse{}, err
	}
	return OperationsClientConstructorResponse{}, nil
}

// constructorCreateRequest creates the Constructor request.
func (client *OperationsClient) constructorCreateRequest(ctx context.Context, options *OperationsClientConstructorOptions) (*policy.Request, error) {
	urlPath := "/special-words/operations/constructor"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (client *OperationsClient) Continue(ctx context.Context, options *OperationsClientContinueOptions) (OperationsClientContinueResponse, error) {
	var err error
	req, err := client.continueCreateRequest(ctx, options)
	if err != nil {
		return OperationsClientContinueResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OperationsClientContinueResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return OperationsClientContinueResponse{}, err
	}
	return OperationsClientContinueResponse{}, nil
}

// continueCreateRequest creates the Continue request.
func (client *OperationsClient) continueCreateRequest(ctx context.Context, options *OperationsClientContinueOptions) (*policy.Request, error) {
	urlPath := "/special-words/operations/continue"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (client *OperationsClient) Def(ctx context.Context, options *OperationsClientDefOptions) (OperationsClientDefResponse, error) {
	var err error
	req, err := client.defCreateRequest(ctx, options)
	if err != nil {
		return OperationsClientDefResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OperationsClientDefResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return OperationsClientDefResponse{}, err
	}
	return OperationsClientDefResponse{}, nil
}

// defCreateRequest creates the Def request.
func (client *OperationsClient) defCreateRequest(ctx context.Context, options *OperationsClientDefOptions) (*policy.Request, error) {
	urlPath := "/special-words/operations/def"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (client *OperationsClient) Del(ctx context.Context, options *OperationsClientDelOptions) (OperationsClientDelResponse, error) {
	var err error
	req, err := client.delCreateRequest(ctx, options)
	if err != nil {
		return OperationsClientDelResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OperationsClientDelResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return OperationsClientDelResponse{}, err
	}
	return OperationsClientDelResponse{}, nil
}

// delCreateRequest creates the Del request.
func (client *OperationsClient) delCreateRequest(ctx context.Context, options *OperationsClientDelOptions) (*policy.Request, error) {
	urlPath := "/special-words/operations/del"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (client *OperationsClient) Elif(ctx context.Context, options *OperationsClientElifOptions) (OperationsClientElifResponse, error) {
	var err error
	req, err := client.elifCreateRequest(ctx, options)
	if err != nil {
		return OperationsClientElifResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OperationsClientElifResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return OperationsClientElifResponse{}, err
	}
	return OperationsClientElifResponse{}, nil
}

// elifCreateRequest creates the Elif request.
func (client *OperationsClient) elifCreateRequest(ctx context.Context, options *OperationsClientElifOptions) (*policy.Request, error) {
	urlPath := "/special-words/operations/elif"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (client *OperationsClient) Else(ctx context.Context, options *OperationsClientElseOptions) (OperationsClientElseResponse, error) {
	var err error
	req, err := client.elseCreateRequest(ctx, options)
	if err != nil {
		return OperationsClientElseResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OperationsClientElseResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return OperationsClientElseResponse{}, err
	}
	return OperationsClientElseResponse{}, nil
}

// elseCreateRequest creates the Else request.
func (client *OperationsClient) elseCreateRequest(ctx context.Context, options *OperationsClientElseOptions) (*policy.Request, error) {
	urlPath := "/special-words/operations/else"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (client *OperationsClient) Except(ctx context.Context, options *OperationsClientExceptOptions) (OperationsClientExceptResponse, error) {
	var err error
	req, err := client.exceptCreateRequest(ctx, options)
	if err != nil {
		return OperationsClientExceptResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OperationsClientExceptResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return OperationsClientExceptResponse{}, err
	}
	return OperationsClientExceptResponse{}, nil
}

// exceptCreateRequest creates the Except request.
func (client *OperationsClient) exceptCreateRequest(ctx context.Context, options *OperationsClientExceptOptions) (*policy.Request, error) {
	urlPath := "/special-words/operations/except"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (client *OperationsClient) Exec(ctx context.Context, options *OperationsClientExecOptions) (OperationsClientExecResponse, error) {
	var err error
	req, err := client.execCreateRequest(ctx, options)
	if err != nil {
		return OperationsClientExecResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OperationsClientExecResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return OperationsClientExecResponse{}, err
	}
	return OperationsClientExecResponse{}, nil
}

// execCreateRequest creates the Exec request.
func (client *OperationsClient) execCreateRequest(ctx context.Context, options *OperationsClientExecOptions) (*policy.Request, error) {
	urlPath := "/special-words/operations/exec"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (client *OperationsClient) Finally(ctx context.Context, options *OperationsClientFinallyOptions) (OperationsClientFinallyResponse, error) {
	var err error
	req, err := client.finallyCreateRequest(ctx, options)
	if err != nil {
		return OperationsClientFinallyResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OperationsClientFinallyResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return OperationsClientFinallyResponse{}, err
	}
	return OperationsClientFinallyResponse{}, nil
}

// finallyCreateRequest creates the Finally request.
func (client *OperationsClient) finallyCreateRequest(ctx context.Context, options *OperationsClientFinallyOptions) (*policy.Request, error) {
	urlPath := "/special-words/operations/finally"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (client *OperationsClient) For(ctx context.Context, options *OperationsClientForOptions) (OperationsClientForResponse, error) {
	var err error
	req, err := client.forCreateRequest(ctx, options)
	if err != nil {
		return OperationsClientForResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OperationsClientForResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return OperationsClientForResponse{}, err
	}
	return OperationsClientForResponse{}, nil
}

// forCreateRequest creates the For request.
func (client *OperationsClient) forCreateRequest(ctx context.Context, options *OperationsClientForOptions) (*policy.Request, error) {
	urlPath := "/special-words/operations/for"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (client *OperationsClient) From(ctx context.Context, options *OperationsClientFromOptions) (OperationsClientFromResponse, error) {
	var err error
	req, err := client.fromCreateRequest(ctx, options)
	if err != nil {
		return OperationsClientFromResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OperationsClientFromResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return OperationsClientFromResponse{}, err
	}
	return OperationsClientFromResponse{}, nil
}

// fromCreateRequest creates the From request.
func (client *OperationsClient) fromCreateRequest(ctx context.Context, options *OperationsClientFromOptions) (*policy.Request, error) {
	urlPath := "/special-words/operations/from"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (client *OperationsClient) Global(ctx context.Context, options *OperationsClientGlobalOptions) (OperationsClientGlobalResponse, error) {
	var err error
	req, err := client.globalCreateRequest(ctx, options)
	if err != nil {
		return OperationsClientGlobalResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OperationsClientGlobalResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return OperationsClientGlobalResponse{}, err
	}
	return OperationsClientGlobalResponse{}, nil
}

// globalCreateRequest creates the Global request.
func (client *OperationsClient) globalCreateRequest(ctx context.Context, options *OperationsClientGlobalOptions) (*policy.Request, error) {
	urlPath := "/special-words/operations/global"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (client *OperationsClient) If(ctx context.Context, options *OperationsClientIfOptions) (OperationsClientIfResponse, error) {
	var err error
	req, err := client.ifCreateRequest(ctx, options)
	if err != nil {
		return OperationsClientIfResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OperationsClientIfResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return OperationsClientIfResponse{}, err
	}
	return OperationsClientIfResponse{}, nil
}

// ifCreateRequest creates the If request.
func (client *OperationsClient) ifCreateRequest(ctx context.Context, options *OperationsClientIfOptions) (*policy.Request, error) {
	urlPath := "/special-words/operations/if"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (client *OperationsClient) Import(ctx context.Context, options *OperationsClientImportOptions) (OperationsClientImportResponse, error) {
	var err error
	req, err := client.importCreateRequest(ctx, options)
	if err != nil {
		return OperationsClientImportResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OperationsClientImportResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return OperationsClientImportResponse{}, err
	}
	return OperationsClientImportResponse{}, nil
}

// importCreateRequest creates the Import request.
func (client *OperationsClient) importCreateRequest(ctx context.Context, options *OperationsClientImportOptions) (*policy.Request, error) {
	urlPath := "/special-words/operations/import"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (client *OperationsClient) In(ctx context.Context, options *OperationsClientInOptions) (OperationsClientInResponse, error) {
	var err error
	req, err := client.inCreateRequest(ctx, options)
	if err != nil {
		return OperationsClientInResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OperationsClientInResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return OperationsClientInResponse{}, err
	}
	return OperationsClientInResponse{}, nil
}

// inCreateRequest creates the In request.
func (client *OperationsClient) inCreateRequest(ctx context.Context, options *OperationsClientInOptions) (*policy.Request, error) {
	urlPath := "/special-words/operations/in"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (client *OperationsClient) Is(ctx context.Context, options *OperationsClientIsOptions) (OperationsClientIsResponse, error) {
	var err error
	req, err := client.isCreateRequest(ctx, options)
	if err != nil {
		return OperationsClientIsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OperationsClientIsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return OperationsClientIsResponse{}, err
	}
	return OperationsClientIsResponse{}, nil
}

// isCreateRequest creates the Is request.
func (client *OperationsClient) isCreateRequest(ctx context.Context, options *OperationsClientIsOptions) (*policy.Request, error) {
	urlPath := "/special-words/operations/is"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (client *OperationsClient) Lambda(ctx context.Context, options *OperationsClientLambdaOptions) (OperationsClientLambdaResponse, error) {
	var err error
	req, err := client.lambdaCreateRequest(ctx, options)
	if err != nil {
		return OperationsClientLambdaResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OperationsClientLambdaResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return OperationsClientLambdaResponse{}, err
	}
	return OperationsClientLambdaResponse{}, nil
}

// lambdaCreateRequest creates the Lambda request.
func (client *OperationsClient) lambdaCreateRequest(ctx context.Context, options *OperationsClientLambdaOptions) (*policy.Request, error) {
	urlPath := "/special-words/operations/lambda"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (client *OperationsClient) Not(ctx context.Context, options *OperationsClientNotOptions) (OperationsClientNotResponse, error) {
	var err error
	req, err := client.notCreateRequest(ctx, options)
	if err != nil {
		return OperationsClientNotResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OperationsClientNotResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return OperationsClientNotResponse{}, err
	}
	return OperationsClientNotResponse{}, nil
}

// notCreateRequest creates the Not request.
func (client *OperationsClient) notCreateRequest(ctx context.Context, options *OperationsClientNotOptions) (*policy.Request, error) {
	urlPath := "/special-words/operations/not"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (client *OperationsClient) Or(ctx context.Context, options *OperationsClientOrOptions) (OperationsClientOrResponse, error) {
	var err error
	req, err := client.orCreateRequest(ctx, options)
	if err != nil {
		return OperationsClientOrResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OperationsClientOrResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return OperationsClientOrResponse{}, err
	}
	return OperationsClientOrResponse{}, nil
}

// orCreateRequest creates the Or request.
func (client *OperationsClient) orCreateRequest(ctx context.Context, options *OperationsClientOrOptions) (*policy.Request, error) {
	urlPath := "/special-words/operations/or"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (client *OperationsClient) Pass(ctx context.Context, options *OperationsClientPassOptions) (OperationsClientPassResponse, error) {
	var err error
	req, err := client.passCreateRequest(ctx, options)
	if err != nil {
		return OperationsClientPassResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OperationsClientPassResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return OperationsClientPassResponse{}, err
	}
	return OperationsClientPassResponse{}, nil
}

// passCreateRequest creates the Pass request.
func (client *OperationsClient) passCreateRequest(ctx context.Context, options *OperationsClientPassOptions) (*policy.Request, error) {
	urlPath := "/special-words/operations/pass"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (client *OperationsClient) Raise(ctx context.Context, options *OperationsClientRaiseOptions) (OperationsClientRaiseResponse, error) {
	var err error
	req, err := client.raiseCreateRequest(ctx, options)
	if err != nil {
		return OperationsClientRaiseResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OperationsClientRaiseResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return OperationsClientRaiseResponse{}, err
	}
	return OperationsClientRaiseResponse{}, nil
}

// raiseCreateRequest creates the Raise request.
func (client *OperationsClient) raiseCreateRequest(ctx context.Context, options *OperationsClientRaiseOptions) (*policy.Request, error) {
	urlPath := "/special-words/operations/raise"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (client *OperationsClient) Return(ctx context.Context, options *OperationsClientReturnOptions) (OperationsClientReturnResponse, error) {
	var err error
	req, err := client.returnCreateRequest(ctx, options)
	if err != nil {
		return OperationsClientReturnResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OperationsClientReturnResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return OperationsClientReturnResponse{}, err
	}
	return OperationsClientReturnResponse{}, nil
}

// returnCreateRequest creates the Return request.
func (client *OperationsClient) returnCreateRequest(ctx context.Context, options *OperationsClientReturnOptions) (*policy.Request, error) {
	urlPath := "/special-words/operations/return"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (client *OperationsClient) Try(ctx context.Context, options *OperationsClientTryOptions) (OperationsClientTryResponse, error) {
	var err error
	req, err := client.tryCreateRequest(ctx, options)
	if err != nil {
		return OperationsClientTryResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OperationsClientTryResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return OperationsClientTryResponse{}, err
	}
	return OperationsClientTryResponse{}, nil
}

// tryCreateRequest creates the Try request.
func (client *OperationsClient) tryCreateRequest(ctx context.Context, options *OperationsClientTryOptions) (*policy.Request, error) {
	urlPath := "/special-words/operations/try"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (client *OperationsClient) While(ctx context.Context, options *OperationsClientWhileOptions) (OperationsClientWhileResponse, error) {
	var err error
	req, err := client.whileCreateRequest(ctx, options)
	if err != nil {
		return OperationsClientWhileResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OperationsClientWhileResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return OperationsClientWhileResponse{}, err
	}
	return OperationsClientWhileResponse{}, nil
}

// whileCreateRequest creates the While request.
func (client *OperationsClient) whileCreateRequest(ctx context.Context, options *OperationsClientWhileOptions) (*policy.Request, error) {
	urlPath := "/special-words/operations/while"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (client *OperationsClient) With(ctx context.Context, options *OperationsClientWithOptions) (OperationsClientWithResponse, error) {
	var err error
	req, err := client.withCreateRequest(ctx, options)
	if err != nil {
		return OperationsClientWithResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OperationsClientWithResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return OperationsClientWithResponse{}, err
	}
	return OperationsClientWithResponse{}, nil
}

// withCreateRequest creates the With request.
func (client *OperationsClient) withCreateRequest(ctx context.Context, options *OperationsClientWithOptions) (*policy.Request, error) {
	urlPath := "/special-words/operations/with"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (client *OperationsClient) Yield(ctx context.Context, options *OperationsClientYieldOptions) (OperationsClientYieldResponse, error) {
	var err error
	req, err := client.yieldCreateRequest(ctx, options)
	if err != nil {
		return OperationsClientYieldResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OperationsClientYieldResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return OperationsClientYieldResponse{}, err
	}
	return OperationsClientYieldResponse{}, nil
}

// yieldCreateRequest creates the Yield request.
func (client *OperationsClient) yieldCreateRequest(ctx context.Context, options *OperationsClientYieldOptions) (*policy.Request, error) {
	urlPath := "/special-words/operations/yield"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}