//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsql

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

// DataMaskingRulesClient contains the methods for the DataMaskingRules group.
// Don't use this type directly, use NewDataMaskingRulesClient() instead.
type DataMaskingRulesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDataMaskingRulesClient creates a new instance of DataMaskingRulesClient with the specified values.
//   - subscriptionID - The subscription ID that identifies an Azure subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDataMaskingRulesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DataMaskingRulesClient, error) {
	cl, err := arm.NewClient(moduleName+".DataMaskingRulesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DataMaskingRulesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates a database data masking rule.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2014-04-01
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - databaseName - The name of the database.
//   - dataMaskingRuleName - The name of the data masking rule.
//   - parameters - The required parameters for creating or updating a data masking rule.
//   - options - DataMaskingRulesClientCreateOrUpdateOptions contains the optional parameters for the DataMaskingRulesClient.CreateOrUpdate
//     method.
func (client *DataMaskingRulesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, databaseName string, dataMaskingRuleName string, parameters DataMaskingRule, options *DataMaskingRulesClientCreateOrUpdateOptions) (DataMaskingRulesClientCreateOrUpdateResponse, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serverName, databaseName, dataMaskingRuleName, parameters, options)
	if err != nil {
		return DataMaskingRulesClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DataMaskingRulesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return DataMaskingRulesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DataMaskingRulesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, dataMaskingRuleName string, parameters DataMaskingRule, options *DataMaskingRulesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/dataMaskingPolicies/{dataMaskingPolicyName}/rules/{dataMaskingRuleName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	urlPath = strings.ReplaceAll(urlPath, "{dataMaskingPolicyName}", url.PathEscape("Default"))
	if dataMaskingRuleName == "" {
		return nil, errors.New("parameter dataMaskingRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataMaskingRuleName}", url.PathEscape(dataMaskingRuleName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2014-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *DataMaskingRulesClient) createOrUpdateHandleResponse(resp *http.Response) (DataMaskingRulesClientCreateOrUpdateResponse, error) {
	result := DataMaskingRulesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataMaskingRule); err != nil {
		return DataMaskingRulesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// NewListByDatabasePager - Gets a list of database data masking rules.
//
// Generated from API version 2014-04-01
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - databaseName - The name of the database.
//   - options - DataMaskingRulesClientListByDatabaseOptions contains the optional parameters for the DataMaskingRulesClient.NewListByDatabasePager
//     method.
func (client *DataMaskingRulesClient) NewListByDatabasePager(resourceGroupName string, serverName string, databaseName string, options *DataMaskingRulesClientListByDatabaseOptions) *runtime.Pager[DataMaskingRulesClientListByDatabaseResponse] {
	return runtime.NewPager(runtime.PagingHandler[DataMaskingRulesClientListByDatabaseResponse]{
		More: func(page DataMaskingRulesClientListByDatabaseResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *DataMaskingRulesClientListByDatabaseResponse) (DataMaskingRulesClientListByDatabaseResponse, error) {
			req, err := client.listByDatabaseCreateRequest(ctx, resourceGroupName, serverName, databaseName, options)
			if err != nil {
				return DataMaskingRulesClientListByDatabaseResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return DataMaskingRulesClientListByDatabaseResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DataMaskingRulesClientListByDatabaseResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByDatabaseHandleResponse(resp)
		},
	})
}

// listByDatabaseCreateRequest creates the ListByDatabase request.
func (client *DataMaskingRulesClient) listByDatabaseCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, options *DataMaskingRulesClientListByDatabaseOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/dataMaskingPolicies/{dataMaskingPolicyName}/rules"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	urlPath = strings.ReplaceAll(urlPath, "{dataMaskingPolicyName}", url.PathEscape("Default"))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2014-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByDatabaseHandleResponse handles the ListByDatabase response.
func (client *DataMaskingRulesClient) listByDatabaseHandleResponse(resp *http.Response) (DataMaskingRulesClientListByDatabaseResponse, error) {
	result := DataMaskingRulesClientListByDatabaseResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataMaskingRuleListResult); err != nil {
		return DataMaskingRulesClientListByDatabaseResponse{}, err
	}
	return result, nil
}
