//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v4"
	"net/http"
	"net/url"
	"regexp"
)

// ApplicationGatewayPrivateLinkResourcesServer is a fake server for instances of the armnetwork.ApplicationGatewayPrivateLinkResourcesClient type.
type ApplicationGatewayPrivateLinkResourcesServer struct {
	// NewListPager is the fake for method ApplicationGatewayPrivateLinkResourcesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, applicationGatewayName string, options *armnetwork.ApplicationGatewayPrivateLinkResourcesClientListOptions) (resp azfake.PagerResponder[armnetwork.ApplicationGatewayPrivateLinkResourcesClientListResponse])
}

// NewApplicationGatewayPrivateLinkResourcesServerTransport creates a new instance of ApplicationGatewayPrivateLinkResourcesServerTransport with the provided implementation.
// The returned ApplicationGatewayPrivateLinkResourcesServerTransport instance is connected to an instance of armnetwork.ApplicationGatewayPrivateLinkResourcesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewApplicationGatewayPrivateLinkResourcesServerTransport(srv *ApplicationGatewayPrivateLinkResourcesServer) *ApplicationGatewayPrivateLinkResourcesServerTransport {
	return &ApplicationGatewayPrivateLinkResourcesServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armnetwork.ApplicationGatewayPrivateLinkResourcesClientListResponse]](),
	}
}

// ApplicationGatewayPrivateLinkResourcesServerTransport connects instances of armnetwork.ApplicationGatewayPrivateLinkResourcesClient to instances of ApplicationGatewayPrivateLinkResourcesServer.
// Don't use this type directly, use NewApplicationGatewayPrivateLinkResourcesServerTransport instead.
type ApplicationGatewayPrivateLinkResourcesServerTransport struct {
	srv          *ApplicationGatewayPrivateLinkResourcesServer
	newListPager *tracker[azfake.PagerResponder[armnetwork.ApplicationGatewayPrivateLinkResourcesClientListResponse]]
}

// Do implements the policy.Transporter interface for ApplicationGatewayPrivateLinkResourcesServerTransport.
func (a *ApplicationGatewayPrivateLinkResourcesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ApplicationGatewayPrivateLinkResourcesClient.NewListPager":
		resp, err = a.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *ApplicationGatewayPrivateLinkResourcesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := a.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/applicationGateways/(?P<applicationGatewayName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateLinkResources`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		applicationGatewayNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("applicationGatewayName")])
		if err != nil {
			return nil, err
		}
		resp := a.srv.NewListPager(resourceGroupNameUnescaped, applicationGatewayNameUnescaped, nil)
		newListPager = &resp
		a.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armnetwork.ApplicationGatewayPrivateLinkResourcesClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		a.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		a.newListPager.remove(req)
	}
	return resp, nil
}
