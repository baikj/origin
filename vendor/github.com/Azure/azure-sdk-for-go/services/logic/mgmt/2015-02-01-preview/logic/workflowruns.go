package logic

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// WorkflowRunsClient is the REST API for Azure Logic Apps.
type WorkflowRunsClient struct {
	BaseClient
}

// NewWorkflowRunsClient creates an instance of the WorkflowRunsClient client.
func NewWorkflowRunsClient(subscriptionID string) WorkflowRunsClient {
	return NewWorkflowRunsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewWorkflowRunsClientWithBaseURI creates an instance of the WorkflowRunsClient client.
func NewWorkflowRunsClientWithBaseURI(baseURI string, subscriptionID string) WorkflowRunsClient {
	return WorkflowRunsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Cancel cancels a workflow run.
//
// resourceGroupName is the resource group name. workflowName is the workflow name. runName is the workflow run name.
func (client WorkflowRunsClient) Cancel(ctx context.Context, resourceGroupName string, workflowName string, runName string) (result autorest.Response, err error) {
	req, err := client.CancelPreparer(ctx, resourceGroupName, workflowName, runName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.WorkflowRunsClient", "Cancel", nil, "Failure preparing request")
		return
	}

	resp, err := client.CancelSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "logic.WorkflowRunsClient", "Cancel", resp, "Failure sending request")
		return
	}

	result, err = client.CancelResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.WorkflowRunsClient", "Cancel", resp, "Failure responding to request")
	}

	return
}

// CancelPreparer prepares the Cancel request.
func (client WorkflowRunsClient) CancelPreparer(ctx context.Context, resourceGroupName string, workflowName string, runName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"runName":           autorest.Encode("path", runName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workflowName":      autorest.Encode("path", workflowName),
	}

	const APIVersion = "2015-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/runs/{runName}/cancel", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CancelSender sends the Cancel request. The method will close the
// http.Response Body if it receives an error.
func (client WorkflowRunsClient) CancelSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// CancelResponder handles the response to the Cancel request. The method always
// closes the http.Response Body.
func (client WorkflowRunsClient) CancelResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets a workflow run.
//
// resourceGroupName is the resource group name. workflowName is the workflow name. runName is the workflow run name.
func (client WorkflowRunsClient) Get(ctx context.Context, resourceGroupName string, workflowName string, runName string) (result WorkflowRun, err error) {
	req, err := client.GetPreparer(ctx, resourceGroupName, workflowName, runName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.WorkflowRunsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "logic.WorkflowRunsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.WorkflowRunsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client WorkflowRunsClient) GetPreparer(ctx context.Context, resourceGroupName string, workflowName string, runName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"runName":           autorest.Encode("path", runName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workflowName":      autorest.Encode("path", workflowName),
	}

	const APIVersion = "2015-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/runs/{runName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client WorkflowRunsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client WorkflowRunsClient) GetResponder(resp *http.Response) (result WorkflowRun, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets a list of workflow runs.
//
// resourceGroupName is the resource group name. workflowName is the workflow name. top is the number of items to be
// included in the result. filter is the filter to apply on the operation.
func (client WorkflowRunsClient) List(ctx context.Context, resourceGroupName string, workflowName string, top *int32, filter string) (result WorkflowRunListResultPage, err error) {
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, workflowName, top, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.WorkflowRunsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.wrlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "logic.WorkflowRunsClient", "List", resp, "Failure sending request")
		return
	}

	result.wrlr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.WorkflowRunsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client WorkflowRunsClient) ListPreparer(ctx context.Context, resourceGroupName string, workflowName string, top *int32, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workflowName":      autorest.Encode("path", workflowName),
	}

	const APIVersion = "2015-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/runs", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client WorkflowRunsClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client WorkflowRunsClient) ListResponder(resp *http.Response) (result WorkflowRunListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client WorkflowRunsClient) listNextResults(lastResults WorkflowRunListResult) (result WorkflowRunListResult, err error) {
	req, err := lastResults.workflowRunListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic.WorkflowRunsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "logic.WorkflowRunsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.WorkflowRunsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client WorkflowRunsClient) ListComplete(ctx context.Context, resourceGroupName string, workflowName string, top *int32, filter string) (result WorkflowRunListResultIterator, err error) {
	result.page, err = client.List(ctx, resourceGroupName, workflowName, top, filter)
	return
}
