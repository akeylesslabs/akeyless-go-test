/*
 * Akeyless API
 *
 * The purpose of this application is to provide access to Akeyless API.
 *
 * API version: 3.0
 * Contact: support@akeyless.io
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package akeyless

import (
	"encoding/json"
)

// OIDCAccessRules OIDCAccessRules contains access rules specific to Open Id Connect authentication method.
type OIDCAccessRules struct {
	// Allowed redirect URIs after the authentication
	AllowedRedirectURIs *[]string `json:"allowed_redirect_URIs,omitempty"`
	// Audience claim to be used as part of the authentication flow. In case set, it must match the one configured on the Identity Provider's Application
	Audience *string `json:"audience,omitempty"`
	// The claims that login is restricted to.
	BoundClaims *[]OIDCCustomClaim `json:"bound_claims,omitempty"`
	// Client ID
	ClientId *string `json:"client_id,omitempty"`
	// Client Secret
	ClientSecret *string `json:"client_secret,omitempty"`
	// IsInternal indicates whether this is an internal Auth Method where the client has no control over it, or it was created by the client e.g - Sign In with Google will create an OIDC Auth Method with IsInternal=true
	IsInternal *bool `json:"is_internal,omitempty"`
	// Issuer URL
	Issuer *string `json:"issuer,omitempty"`
	// A list of required scopes to request from the oidc provider, and to check on the token
	RequiredScopes *[]string `json:"required_scopes,omitempty"`
	// A prefix to add to the required scopes (for example, azures' Application ID URI)
	RequiredScopesPrefix *string `json:"required_scopes_prefix,omitempty"`
	// A unique identifier to distinguish different users
	UniqueIdentifier *string `json:"unique_identifier,omitempty"`
}

// NewOIDCAccessRules instantiates a new OIDCAccessRules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOIDCAccessRules() *OIDCAccessRules {
	this := OIDCAccessRules{}
	return &this
}

// NewOIDCAccessRulesWithDefaults instantiates a new OIDCAccessRules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOIDCAccessRulesWithDefaults() *OIDCAccessRules {
	this := OIDCAccessRules{}
	return &this
}

// GetAllowedRedirectURIs returns the AllowedRedirectURIs field value if set, zero value otherwise.
func (o *OIDCAccessRules) GetAllowedRedirectURIs() []string {
	if o == nil || o.AllowedRedirectURIs == nil {
		var ret []string
		return ret
	}
	return *o.AllowedRedirectURIs
}

// GetAllowedRedirectURIsOk returns a tuple with the AllowedRedirectURIs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDCAccessRules) GetAllowedRedirectURIsOk() (*[]string, bool) {
	if o == nil || o.AllowedRedirectURIs == nil {
		return nil, false
	}
	return o.AllowedRedirectURIs, true
}

// HasAllowedRedirectURIs returns a boolean if a field has been set.
func (o *OIDCAccessRules) HasAllowedRedirectURIs() bool {
	if o != nil && o.AllowedRedirectURIs != nil {
		return true
	}

	return false
}

// SetAllowedRedirectURIs gets a reference to the given []string and assigns it to the AllowedRedirectURIs field.
func (o *OIDCAccessRules) SetAllowedRedirectURIs(v []string) {
	o.AllowedRedirectURIs = &v
}

// GetAudience returns the Audience field value if set, zero value otherwise.
func (o *OIDCAccessRules) GetAudience() string {
	if o == nil || o.Audience == nil {
		var ret string
		return ret
	}
	return *o.Audience
}

// GetAudienceOk returns a tuple with the Audience field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDCAccessRules) GetAudienceOk() (*string, bool) {
	if o == nil || o.Audience == nil {
		return nil, false
	}
	return o.Audience, true
}

// HasAudience returns a boolean if a field has been set.
func (o *OIDCAccessRules) HasAudience() bool {
	if o != nil && o.Audience != nil {
		return true
	}

	return false
}

// SetAudience gets a reference to the given string and assigns it to the Audience field.
func (o *OIDCAccessRules) SetAudience(v string) {
	o.Audience = &v
}

// GetBoundClaims returns the BoundClaims field value if set, zero value otherwise.
func (o *OIDCAccessRules) GetBoundClaims() []OIDCCustomClaim {
	if o == nil || o.BoundClaims == nil {
		var ret []OIDCCustomClaim
		return ret
	}
	return *o.BoundClaims
}

// GetBoundClaimsOk returns a tuple with the BoundClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDCAccessRules) GetBoundClaimsOk() (*[]OIDCCustomClaim, bool) {
	if o == nil || o.BoundClaims == nil {
		return nil, false
	}
	return o.BoundClaims, true
}

// HasBoundClaims returns a boolean if a field has been set.
func (o *OIDCAccessRules) HasBoundClaims() bool {
	if o != nil && o.BoundClaims != nil {
		return true
	}

	return false
}

// SetBoundClaims gets a reference to the given []OIDCCustomClaim and assigns it to the BoundClaims field.
func (o *OIDCAccessRules) SetBoundClaims(v []OIDCCustomClaim) {
	o.BoundClaims = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *OIDCAccessRules) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDCAccessRules) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *OIDCAccessRules) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *OIDCAccessRules) SetClientId(v string) {
	o.ClientId = &v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *OIDCAccessRules) GetClientSecret() string {
	if o == nil || o.ClientSecret == nil {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDCAccessRules) GetClientSecretOk() (*string, bool) {
	if o == nil || o.ClientSecret == nil {
		return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *OIDCAccessRules) HasClientSecret() bool {
	if o != nil && o.ClientSecret != nil {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *OIDCAccessRules) SetClientSecret(v string) {
	o.ClientSecret = &v
}

// GetIsInternal returns the IsInternal field value if set, zero value otherwise.
func (o *OIDCAccessRules) GetIsInternal() bool {
	if o == nil || o.IsInternal == nil {
		var ret bool
		return ret
	}
	return *o.IsInternal
}

// GetIsInternalOk returns a tuple with the IsInternal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDCAccessRules) GetIsInternalOk() (*bool, bool) {
	if o == nil || o.IsInternal == nil {
		return nil, false
	}
	return o.IsInternal, true
}

// HasIsInternal returns a boolean if a field has been set.
func (o *OIDCAccessRules) HasIsInternal() bool {
	if o != nil && o.IsInternal != nil {
		return true
	}

	return false
}

// SetIsInternal gets a reference to the given bool and assigns it to the IsInternal field.
func (o *OIDCAccessRules) SetIsInternal(v bool) {
	o.IsInternal = &v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *OIDCAccessRules) GetIssuer() string {
	if o == nil || o.Issuer == nil {
		var ret string
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDCAccessRules) GetIssuerOk() (*string, bool) {
	if o == nil || o.Issuer == nil {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *OIDCAccessRules) HasIssuer() bool {
	if o != nil && o.Issuer != nil {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given string and assigns it to the Issuer field.
func (o *OIDCAccessRules) SetIssuer(v string) {
	o.Issuer = &v
}

// GetRequiredScopes returns the RequiredScopes field value if set, zero value otherwise.
func (o *OIDCAccessRules) GetRequiredScopes() []string {
	if o == nil || o.RequiredScopes == nil {
		var ret []string
		return ret
	}
	return *o.RequiredScopes
}

// GetRequiredScopesOk returns a tuple with the RequiredScopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDCAccessRules) GetRequiredScopesOk() (*[]string, bool) {
	if o == nil || o.RequiredScopes == nil {
		return nil, false
	}
	return o.RequiredScopes, true
}

// HasRequiredScopes returns a boolean if a field has been set.
func (o *OIDCAccessRules) HasRequiredScopes() bool {
	if o != nil && o.RequiredScopes != nil {
		return true
	}

	return false
}

// SetRequiredScopes gets a reference to the given []string and assigns it to the RequiredScopes field.
func (o *OIDCAccessRules) SetRequiredScopes(v []string) {
	o.RequiredScopes = &v
}

// GetRequiredScopesPrefix returns the RequiredScopesPrefix field value if set, zero value otherwise.
func (o *OIDCAccessRules) GetRequiredScopesPrefix() string {
	if o == nil || o.RequiredScopesPrefix == nil {
		var ret string
		return ret
	}
	return *o.RequiredScopesPrefix
}

// GetRequiredScopesPrefixOk returns a tuple with the RequiredScopesPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDCAccessRules) GetRequiredScopesPrefixOk() (*string, bool) {
	if o == nil || o.RequiredScopesPrefix == nil {
		return nil, false
	}
	return o.RequiredScopesPrefix, true
}

// HasRequiredScopesPrefix returns a boolean if a field has been set.
func (o *OIDCAccessRules) HasRequiredScopesPrefix() bool {
	if o != nil && o.RequiredScopesPrefix != nil {
		return true
	}

	return false
}

// SetRequiredScopesPrefix gets a reference to the given string and assigns it to the RequiredScopesPrefix field.
func (o *OIDCAccessRules) SetRequiredScopesPrefix(v string) {
	o.RequiredScopesPrefix = &v
}

// GetUniqueIdentifier returns the UniqueIdentifier field value if set, zero value otherwise.
func (o *OIDCAccessRules) GetUniqueIdentifier() string {
	if o == nil || o.UniqueIdentifier == nil {
		var ret string
		return ret
	}
	return *o.UniqueIdentifier
}

// GetUniqueIdentifierOk returns a tuple with the UniqueIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDCAccessRules) GetUniqueIdentifierOk() (*string, bool) {
	if o == nil || o.UniqueIdentifier == nil {
		return nil, false
	}
	return o.UniqueIdentifier, true
}

// HasUniqueIdentifier returns a boolean if a field has been set.
func (o *OIDCAccessRules) HasUniqueIdentifier() bool {
	if o != nil && o.UniqueIdentifier != nil {
		return true
	}

	return false
}

// SetUniqueIdentifier gets a reference to the given string and assigns it to the UniqueIdentifier field.
func (o *OIDCAccessRules) SetUniqueIdentifier(v string) {
	o.UniqueIdentifier = &v
}

func (o OIDCAccessRules) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AllowedRedirectURIs != nil {
		toSerialize["allowed_redirect_URIs"] = o.AllowedRedirectURIs
	}
	if o.Audience != nil {
		toSerialize["audience"] = o.Audience
	}
	if o.BoundClaims != nil {
		toSerialize["bound_claims"] = o.BoundClaims
	}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.ClientSecret != nil {
		toSerialize["client_secret"] = o.ClientSecret
	}
	if o.IsInternal != nil {
		toSerialize["is_internal"] = o.IsInternal
	}
	if o.Issuer != nil {
		toSerialize["issuer"] = o.Issuer
	}
	if o.RequiredScopes != nil {
		toSerialize["required_scopes"] = o.RequiredScopes
	}
	if o.RequiredScopesPrefix != nil {
		toSerialize["required_scopes_prefix"] = o.RequiredScopesPrefix
	}
	if o.UniqueIdentifier != nil {
		toSerialize["unique_identifier"] = o.UniqueIdentifier
	}
	return json.Marshal(toSerialize)
}

type NullableOIDCAccessRules struct {
	value *OIDCAccessRules
	isSet bool
}

func (v NullableOIDCAccessRules) Get() *OIDCAccessRules {
	return v.value
}

func (v *NullableOIDCAccessRules) Set(val *OIDCAccessRules) {
	v.value = val
	v.isSet = true
}

func (v NullableOIDCAccessRules) IsSet() bool {
	return v.isSet
}

func (v *NullableOIDCAccessRules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOIDCAccessRules(val *OIDCAccessRules) *NullableOIDCAccessRules {
	return &NullableOIDCAccessRules{value: val, isSet: true}
}

func (v NullableOIDCAccessRules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOIDCAccessRules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


