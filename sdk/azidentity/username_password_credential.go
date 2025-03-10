//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azidentity

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
)

const credNameUserPassword = "UsernamePasswordCredential"

// UsernamePasswordCredentialOptions contains optional parameters for UsernamePasswordCredential.
type UsernamePasswordCredentialOptions struct {
	azcore.ClientOptions

	// AdditionallyAllowedTenants specifies additional tenants for which the credential may acquire tokens.
	// Add the wildcard value "*" to allow the credential to acquire tokens for any tenant in which the
	// application is registered.
	AdditionallyAllowedTenants []string

	// AuthenticationRecord returned by a call to a credential's Authenticate method. Set this option
	// to enable the credential to use data from a previous authentication.
	AuthenticationRecord AuthenticationRecord

	// DisableInstanceDiscovery should be set true only by applications authenticating in disconnected clouds, or
	// private clouds such as Azure Stack. It determines whether the credential requests Azure AD instance metadata
	// from https://login.microsoft.com before authenticating. Setting this to true will skip this request, making
	// the application responsible for ensuring the configured authority is valid and trustworthy.
	DisableInstanceDiscovery bool

	// TokenCachePersistenceOptions enables persistent token caching when not nil.
	TokenCachePersistenceOptions *TokenCachePersistenceOptions
}

// UsernamePasswordCredential authenticates a user with a password. Microsoft doesn't recommend this kind of authentication,
// because it's less secure than other authentication flows. This credential is not interactive, so it isn't compatible
// with any form of multi-factor authentication, and the application must already have user or admin consent.
// This credential can only authenticate work and school accounts; it can't authenticate Microsoft accounts.
type UsernamePasswordCredential struct {
	client *publicClient
}

// NewUsernamePasswordCredential creates a UsernamePasswordCredential. clientID is the ID of the application the user
// will authenticate to. Pass nil for options to accept defaults.
func NewUsernamePasswordCredential(tenantID string, clientID string, username string, password string, options *UsernamePasswordCredentialOptions) (*UsernamePasswordCredential, error) {
	if options == nil {
		options = &UsernamePasswordCredentialOptions{}
	}
	opts := publicClientOptions{
		AdditionallyAllowedTenants:   options.AdditionallyAllowedTenants,
		ClientOptions:                options.ClientOptions,
		DisableInstanceDiscovery:     options.DisableInstanceDiscovery,
		Password:                     password,
		Record:                       options.AuthenticationRecord,
		TokenCachePersistenceOptions: options.TokenCachePersistenceOptions,
		Username:                     username,
	}
	c, err := newPublicClient(tenantID, clientID, credNameUserPassword, opts)
	if err != nil {
		return nil, err
	}
	return &UsernamePasswordCredential{client: c}, err
}

// Authenticate the user. Subsequent calls to GetToken will automatically use the returned AuthenticationRecord.
func (c *UsernamePasswordCredential) Authenticate(ctx context.Context, opts *policy.TokenRequestOptions) (AuthenticationRecord, error) {
	return c.client.Authenticate(ctx, opts)
}

// GetToken requests an access token from Azure Active Directory. This method is called automatically by Azure SDK clients.
func (c *UsernamePasswordCredential) GetToken(ctx context.Context, opts policy.TokenRequestOptions) (azcore.AccessToken, error) {
	return c.client.GetToken(ctx, opts)
}

var _ azcore.TokenCredential = (*UsernamePasswordCredential)(nil)
