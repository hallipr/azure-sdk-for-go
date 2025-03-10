//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstorageimportexport

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

// BitLockerKeysClient contains the methods for the BitLockerKeys group.
// Don't use this type directly, use NewBitLockerKeysClient() instead.
type BitLockerKeysClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewBitLockerKeysClient creates a new instance of BitLockerKeysClient with the specified values.
//   - subscriptionID - The subscription ID for the Azure user.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewBitLockerKeysClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*BitLockerKeysClient, error) {
	cl, err := arm.NewClient(moduleName+".BitLockerKeysClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &BitLockerKeysClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListPager - Returns the BitLocker Keys for all drives in the specified job.
//
// Generated from API version 2021-01-01
//   - jobName - The name of the import/export job.
//   - resourceGroupName - The resource group name uniquely identifies the resource group within the user subscription.
//   - options - BitLockerKeysClientListOptions contains the optional parameters for the BitLockerKeysClient.NewListPager method.
func (client *BitLockerKeysClient) NewListPager(jobName string, resourceGroupName string, options *BitLockerKeysClientListOptions) *runtime.Pager[BitLockerKeysClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[BitLockerKeysClientListResponse]{
		More: func(page BitLockerKeysClientListResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *BitLockerKeysClientListResponse) (BitLockerKeysClientListResponse, error) {
			req, err := client.listCreateRequest(ctx, jobName, resourceGroupName, options)
			if err != nil {
				return BitLockerKeysClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return BitLockerKeysClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return BitLockerKeysClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *BitLockerKeysClient) listCreateRequest(ctx context.Context, jobName string, resourceGroupName string, options *BitLockerKeysClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ImportExport/jobs/{jobName}/listBitLockerKeys"
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.AcceptLanguage != nil {
		req.Raw().Header["Accept-Language"] = []string{*options.AcceptLanguage}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *BitLockerKeysClient) listHandleResponse(resp *http.Response) (BitLockerKeysClientListResponse, error) {
	result := BitLockerKeysClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GetBitLockerKeysResponse); err != nil {
		return BitLockerKeysClientListResponse{}, err
	}
	return result, nil
}
