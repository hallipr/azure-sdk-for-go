//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmobilenetwork

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

// PacketCapturesClient contains the methods for the PacketCaptures group.
// Don't use this type directly, use NewPacketCapturesClient() instead.
type PacketCapturesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewPacketCapturesClient creates a new instance of PacketCapturesClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewPacketCapturesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PacketCapturesClient, error) {
	cl, err := arm.NewClient(moduleName+".PacketCapturesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &PacketCapturesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a packet capture.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - packetCoreControlPlaneName - The name of the packet core control plane.
//   - packetCaptureName - The name of the packet capture session.
//   - parameters - Parameters supplied to the create or update packet capture operation.
//   - options - PacketCapturesClientBeginCreateOrUpdateOptions contains the optional parameters for the PacketCapturesClient.BeginCreateOrUpdate
//     method.
func (client *PacketCapturesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, packetCoreControlPlaneName string, packetCaptureName string, parameters PacketCapture, options *PacketCapturesClientBeginCreateOrUpdateOptions) (*runtime.Poller[PacketCapturesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, packetCoreControlPlaneName, packetCaptureName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[PacketCapturesClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[PacketCapturesClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Creates or updates a packet capture.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01
func (client *PacketCapturesClient) createOrUpdate(ctx context.Context, resourceGroupName string, packetCoreControlPlaneName string, packetCaptureName string, parameters PacketCapture, options *PacketCapturesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, packetCoreControlPlaneName, packetCaptureName, parameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *PacketCapturesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, packetCoreControlPlaneName string, packetCaptureName string, parameters PacketCapture, options *PacketCapturesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/{packetCoreControlPlaneName}/packetCaptures/{packetCaptureName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if packetCoreControlPlaneName == "" {
		return nil, errors.New("parameter packetCoreControlPlaneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packetCoreControlPlaneName}", url.PathEscape(packetCoreControlPlaneName))
	if packetCaptureName == "" {
		return nil, errors.New("parameter packetCaptureName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packetCaptureName}", url.PathEscape(packetCaptureName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes the specified packet capture.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - packetCoreControlPlaneName - The name of the packet core control plane.
//   - packetCaptureName - The name of the packet capture session.
//   - options - PacketCapturesClientBeginDeleteOptions contains the optional parameters for the PacketCapturesClient.BeginDelete
//     method.
func (client *PacketCapturesClient) BeginDelete(ctx context.Context, resourceGroupName string, packetCoreControlPlaneName string, packetCaptureName string, options *PacketCapturesClientBeginDeleteOptions) (*runtime.Poller[PacketCapturesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, packetCoreControlPlaneName, packetCaptureName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[PacketCapturesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[PacketCapturesClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes the specified packet capture.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01
func (client *PacketCapturesClient) deleteOperation(ctx context.Context, resourceGroupName string, packetCoreControlPlaneName string, packetCaptureName string, options *PacketCapturesClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, packetCoreControlPlaneName, packetCaptureName, options)
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
func (client *PacketCapturesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, packetCoreControlPlaneName string, packetCaptureName string, options *PacketCapturesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/{packetCoreControlPlaneName}/packetCaptures/{packetCaptureName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if packetCoreControlPlaneName == "" {
		return nil, errors.New("parameter packetCoreControlPlaneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packetCoreControlPlaneName}", url.PathEscape(packetCoreControlPlaneName))
	if packetCaptureName == "" {
		return nil, errors.New("parameter packetCaptureName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packetCaptureName}", url.PathEscape(packetCaptureName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets information about the specified packet capture session.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - packetCoreControlPlaneName - The name of the packet core control plane.
//   - packetCaptureName - The name of the packet capture session.
//   - options - PacketCapturesClientGetOptions contains the optional parameters for the PacketCapturesClient.Get method.
func (client *PacketCapturesClient) Get(ctx context.Context, resourceGroupName string, packetCoreControlPlaneName string, packetCaptureName string, options *PacketCapturesClientGetOptions) (PacketCapturesClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, packetCoreControlPlaneName, packetCaptureName, options)
	if err != nil {
		return PacketCapturesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PacketCapturesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PacketCapturesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *PacketCapturesClient) getCreateRequest(ctx context.Context, resourceGroupName string, packetCoreControlPlaneName string, packetCaptureName string, options *PacketCapturesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/{packetCoreControlPlaneName}/packetCaptures/{packetCaptureName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if packetCoreControlPlaneName == "" {
		return nil, errors.New("parameter packetCoreControlPlaneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packetCoreControlPlaneName}", url.PathEscape(packetCoreControlPlaneName))
	if packetCaptureName == "" {
		return nil, errors.New("parameter packetCaptureName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packetCaptureName}", url.PathEscape(packetCaptureName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PacketCapturesClient) getHandleResponse(resp *http.Response) (PacketCapturesClientGetResponse, error) {
	result := PacketCapturesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PacketCapture); err != nil {
		return PacketCapturesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByPacketCoreControlPlanePager - Lists all the packet capture sessions under a packet core control plane.
//
// Generated from API version 2023-06-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - packetCoreControlPlaneName - The name of the packet core control plane.
//   - options - PacketCapturesClientListByPacketCoreControlPlaneOptions contains the optional parameters for the PacketCapturesClient.NewListByPacketCoreControlPlanePager
//     method.
func (client *PacketCapturesClient) NewListByPacketCoreControlPlanePager(resourceGroupName string, packetCoreControlPlaneName string, options *PacketCapturesClientListByPacketCoreControlPlaneOptions) *runtime.Pager[PacketCapturesClientListByPacketCoreControlPlaneResponse] {
	return runtime.NewPager(runtime.PagingHandler[PacketCapturesClientListByPacketCoreControlPlaneResponse]{
		More: func(page PacketCapturesClientListByPacketCoreControlPlaneResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PacketCapturesClientListByPacketCoreControlPlaneResponse) (PacketCapturesClientListByPacketCoreControlPlaneResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByPacketCoreControlPlaneCreateRequest(ctx, resourceGroupName, packetCoreControlPlaneName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return PacketCapturesClientListByPacketCoreControlPlaneResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return PacketCapturesClientListByPacketCoreControlPlaneResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return PacketCapturesClientListByPacketCoreControlPlaneResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByPacketCoreControlPlaneHandleResponse(resp)
		},
	})
}

// listByPacketCoreControlPlaneCreateRequest creates the ListByPacketCoreControlPlane request.
func (client *PacketCapturesClient) listByPacketCoreControlPlaneCreateRequest(ctx context.Context, resourceGroupName string, packetCoreControlPlaneName string, options *PacketCapturesClientListByPacketCoreControlPlaneOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/{packetCoreControlPlaneName}/packetCaptures"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if packetCoreControlPlaneName == "" {
		return nil, errors.New("parameter packetCoreControlPlaneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packetCoreControlPlaneName}", url.PathEscape(packetCoreControlPlaneName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByPacketCoreControlPlaneHandleResponse handles the ListByPacketCoreControlPlane response.
func (client *PacketCapturesClient) listByPacketCoreControlPlaneHandleResponse(resp *http.Response) (PacketCapturesClientListByPacketCoreControlPlaneResponse, error) {
	result := PacketCapturesClientListByPacketCoreControlPlaneResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PacketCaptureListResult); err != nil {
		return PacketCapturesClientListByPacketCoreControlPlaneResponse{}, err
	}
	return result, nil
}

// BeginStop - Stop a packet capture session.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - packetCoreControlPlaneName - The name of the packet core control plane.
//   - packetCaptureName - The name of the packet capture session.
//   - options - PacketCapturesClientBeginStopOptions contains the optional parameters for the PacketCapturesClient.BeginStop
//     method.
func (client *PacketCapturesClient) BeginStop(ctx context.Context, resourceGroupName string, packetCoreControlPlaneName string, packetCaptureName string, options *PacketCapturesClientBeginStopOptions) (*runtime.Poller[PacketCapturesClientStopResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.stop(ctx, resourceGroupName, packetCoreControlPlaneName, packetCaptureName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[PacketCapturesClientStopResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[PacketCapturesClientStopResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Stop - Stop a packet capture session.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01
func (client *PacketCapturesClient) stop(ctx context.Context, resourceGroupName string, packetCoreControlPlaneName string, packetCaptureName string, options *PacketCapturesClientBeginStopOptions) (*http.Response, error) {
	var err error
	req, err := client.stopCreateRequest(ctx, resourceGroupName, packetCoreControlPlaneName, packetCaptureName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// stopCreateRequest creates the Stop request.
func (client *PacketCapturesClient) stopCreateRequest(ctx context.Context, resourceGroupName string, packetCoreControlPlaneName string, packetCaptureName string, options *PacketCapturesClientBeginStopOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/{packetCoreControlPlaneName}/packetCaptures/{packetCaptureName}/stop"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if packetCoreControlPlaneName == "" {
		return nil, errors.New("parameter packetCoreControlPlaneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packetCoreControlPlaneName}", url.PathEscape(packetCoreControlPlaneName))
	if packetCaptureName == "" {
		return nil, errors.New("parameter packetCaptureName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packetCaptureName}", url.PathEscape(packetCaptureName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}
