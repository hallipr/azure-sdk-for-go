//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
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

// InterfaceIPConfigurationsServer is a fake server for instances of the armnetwork.InterfaceIPConfigurationsClient type.
type InterfaceIPConfigurationsServer struct {
	// Get is the fake for method InterfaceIPConfigurationsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, networkInterfaceName string, ipConfigurationName string, options *armnetwork.InterfaceIPConfigurationsClientGetOptions) (resp azfake.Responder[armnetwork.InterfaceIPConfigurationsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method InterfaceIPConfigurationsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, networkInterfaceName string, options *armnetwork.InterfaceIPConfigurationsClientListOptions) (resp azfake.PagerResponder[armnetwork.InterfaceIPConfigurationsClientListResponse])
}

// NewInterfaceIPConfigurationsServerTransport creates a new instance of InterfaceIPConfigurationsServerTransport with the provided implementation.
// The returned InterfaceIPConfigurationsServerTransport instance is connected to an instance of armnetwork.InterfaceIPConfigurationsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewInterfaceIPConfigurationsServerTransport(srv *InterfaceIPConfigurationsServer) *InterfaceIPConfigurationsServerTransport {
	return &InterfaceIPConfigurationsServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armnetwork.InterfaceIPConfigurationsClientListResponse]](),
	}
}

// InterfaceIPConfigurationsServerTransport connects instances of armnetwork.InterfaceIPConfigurationsClient to instances of InterfaceIPConfigurationsServer.
// Don't use this type directly, use NewInterfaceIPConfigurationsServerTransport instead.
type InterfaceIPConfigurationsServerTransport struct {
	srv          *InterfaceIPConfigurationsServer
	newListPager *tracker[azfake.PagerResponder[armnetwork.InterfaceIPConfigurationsClientListResponse]]
}

// Do implements the policy.Transporter interface for InterfaceIPConfigurationsServerTransport.
func (i *InterfaceIPConfigurationsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "InterfaceIPConfigurationsClient.Get":
		resp, err = i.dispatchGet(req)
	case "InterfaceIPConfigurationsClient.NewListPager":
		resp, err = i.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (i *InterfaceIPConfigurationsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if i.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/networkInterfaces/(?P<networkInterfaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/ipConfigurations/(?P<ipConfigurationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	networkInterfaceNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("networkInterfaceName")])
	if err != nil {
		return nil, err
	}
	ipConfigurationNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("ipConfigurationName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.Get(req.Context(), resourceGroupNameUnescaped, networkInterfaceNameUnescaped, ipConfigurationNameUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).InterfaceIPConfiguration, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *InterfaceIPConfigurationsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if i.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := i.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/networkInterfaces/(?P<networkInterfaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/ipConfigurations`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		networkInterfaceNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("networkInterfaceName")])
		if err != nil {
			return nil, err
		}
		resp := i.srv.NewListPager(resourceGroupNameUnescaped, networkInterfaceNameUnescaped, nil)
		newListPager = &resp
		i.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armnetwork.InterfaceIPConfigurationsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		i.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		i.newListPager.remove(req)
	}
	return resp, nil
}
