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

// UpdateAuthMethodGCP updateAuthMethodGCP is a command that updates a new auth method that will be able to authenticate using GCP IAM Service Account credentials or GCE instance credentials.
type UpdateAuthMethodGCP struct {
	// Access expiration date in Unix timestamp (select 0 for access without expiry date)
	AccessExpires *int64 `json:"access-expires,omitempty"`
	// The audience to verify in the JWT received by the client
	Audience string `json:"audience"`
	// A CIDR whitelist with the IPs that the access is restricted to
	BoundIps *[]string `json:"bound-ips,omitempty"`
	// A comma-separated list of GCP labels formatted as \"key:value\" strings that must be set on authorized GCE instances. TODO: Because GCP labels are not currently ACL'd ....
	BoundLabels *[]string `json:"bound-labels,omitempty"`
	// === Human and Machine authentication section === Array of GCP project IDs. Only entities belonging to any of the provided projects can authenticate.
	BoundProjects *[]string `json:"bound-projects,omitempty"`
	// List of regions that a GCE instance must belong to in order to be authenticated. TODO: If bound_instance_groups is provided, it is assumed to be a regional group and the group must belong to this region. If bound_zones are provided, this attribute is ignored.
	BoundRegions *[]string `json:"bound-regions,omitempty"`
	// List of service accounts the service account must be part of in order to be authenticated.
	BoundServiceAccounts *[]string `json:"bound-service-accounts,omitempty"`
	// === Machine authentication section === List of zones that a GCE instance must belong to in order to be authenticated. TODO: If bound_instance_groups is provided, it is assumed to be a zonal group and the group must belong to this zone.
	BoundZones *[]string `json:"bound-zones,omitempty"`
	// if true: enforce role-association must include sub claims
	ForceSubClaims *bool `json:"force-sub-claims,omitempty"`
	// A CIDR whitelist with the GW IPs that the access is restricted to
	GwBoundIps *[]string `json:"gw-bound-ips,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Jwt TTL
	JwtTtl *int64 `json:"jwt-ttl,omitempty"`
	// Auth Method name
	Name string `json:"name"`
	// Auth Method new name
	NewName *string `json:"new-name,omitempty"`
	// ServiceAccount credentials data instead of giving a file path, base64 encoded
	ServiceAccountCredsData *string `json:"service-account-creds-data,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// Type of the GCP Access Rules
	Type string `json:"type"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewUpdateAuthMethodGCP instantiates a new UpdateAuthMethodGCP object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateAuthMethodGCP(audience string, name string, type_ string, ) *UpdateAuthMethodGCP {
	this := UpdateAuthMethodGCP{}
	var accessExpires int64 = 0
	this.AccessExpires = &accessExpires
	this.Audience = audience
	this.Name = name
	this.Type = type_
	return &this
}

// NewUpdateAuthMethodGCPWithDefaults instantiates a new UpdateAuthMethodGCP object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateAuthMethodGCPWithDefaults() *UpdateAuthMethodGCP {
	this := UpdateAuthMethodGCP{}
	var accessExpires int64 = 0
	this.AccessExpires = &accessExpires
	var audience string = "akeyless.io"
	this.Audience = audience
	return &this
}

// GetAccessExpires returns the AccessExpires field value if set, zero value otherwise.
func (o *UpdateAuthMethodGCP) GetAccessExpires() int64 {
	if o == nil || o.AccessExpires == nil {
		var ret int64
		return ret
	}
	return *o.AccessExpires
}

// GetAccessExpiresOk returns a tuple with the AccessExpires field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodGCP) GetAccessExpiresOk() (*int64, bool) {
	if o == nil || o.AccessExpires == nil {
		return nil, false
	}
	return o.AccessExpires, true
}

// HasAccessExpires returns a boolean if a field has been set.
func (o *UpdateAuthMethodGCP) HasAccessExpires() bool {
	if o != nil && o.AccessExpires != nil {
		return true
	}

	return false
}

// SetAccessExpires gets a reference to the given int64 and assigns it to the AccessExpires field.
func (o *UpdateAuthMethodGCP) SetAccessExpires(v int64) {
	o.AccessExpires = &v
}

// GetAudience returns the Audience field value
func (o *UpdateAuthMethodGCP) GetAudience() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Audience
}

// GetAudienceOk returns a tuple with the Audience field value
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodGCP) GetAudienceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Audience, true
}

// SetAudience sets field value
func (o *UpdateAuthMethodGCP) SetAudience(v string) {
	o.Audience = v
}

// GetBoundIps returns the BoundIps field value if set, zero value otherwise.
func (o *UpdateAuthMethodGCP) GetBoundIps() []string {
	if o == nil || o.BoundIps == nil {
		var ret []string
		return ret
	}
	return *o.BoundIps
}

// GetBoundIpsOk returns a tuple with the BoundIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodGCP) GetBoundIpsOk() (*[]string, bool) {
	if o == nil || o.BoundIps == nil {
		return nil, false
	}
	return o.BoundIps, true
}

// HasBoundIps returns a boolean if a field has been set.
func (o *UpdateAuthMethodGCP) HasBoundIps() bool {
	if o != nil && o.BoundIps != nil {
		return true
	}

	return false
}

// SetBoundIps gets a reference to the given []string and assigns it to the BoundIps field.
func (o *UpdateAuthMethodGCP) SetBoundIps(v []string) {
	o.BoundIps = &v
}

// GetBoundLabels returns the BoundLabels field value if set, zero value otherwise.
func (o *UpdateAuthMethodGCP) GetBoundLabels() []string {
	if o == nil || o.BoundLabels == nil {
		var ret []string
		return ret
	}
	return *o.BoundLabels
}

// GetBoundLabelsOk returns a tuple with the BoundLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodGCP) GetBoundLabelsOk() (*[]string, bool) {
	if o == nil || o.BoundLabels == nil {
		return nil, false
	}
	return o.BoundLabels, true
}

// HasBoundLabels returns a boolean if a field has been set.
func (o *UpdateAuthMethodGCP) HasBoundLabels() bool {
	if o != nil && o.BoundLabels != nil {
		return true
	}

	return false
}

// SetBoundLabels gets a reference to the given []string and assigns it to the BoundLabels field.
func (o *UpdateAuthMethodGCP) SetBoundLabels(v []string) {
	o.BoundLabels = &v
}

// GetBoundProjects returns the BoundProjects field value if set, zero value otherwise.
func (o *UpdateAuthMethodGCP) GetBoundProjects() []string {
	if o == nil || o.BoundProjects == nil {
		var ret []string
		return ret
	}
	return *o.BoundProjects
}

// GetBoundProjectsOk returns a tuple with the BoundProjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodGCP) GetBoundProjectsOk() (*[]string, bool) {
	if o == nil || o.BoundProjects == nil {
		return nil, false
	}
	return o.BoundProjects, true
}

// HasBoundProjects returns a boolean if a field has been set.
func (o *UpdateAuthMethodGCP) HasBoundProjects() bool {
	if o != nil && o.BoundProjects != nil {
		return true
	}

	return false
}

// SetBoundProjects gets a reference to the given []string and assigns it to the BoundProjects field.
func (o *UpdateAuthMethodGCP) SetBoundProjects(v []string) {
	o.BoundProjects = &v
}

// GetBoundRegions returns the BoundRegions field value if set, zero value otherwise.
func (o *UpdateAuthMethodGCP) GetBoundRegions() []string {
	if o == nil || o.BoundRegions == nil {
		var ret []string
		return ret
	}
	return *o.BoundRegions
}

// GetBoundRegionsOk returns a tuple with the BoundRegions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodGCP) GetBoundRegionsOk() (*[]string, bool) {
	if o == nil || o.BoundRegions == nil {
		return nil, false
	}
	return o.BoundRegions, true
}

// HasBoundRegions returns a boolean if a field has been set.
func (o *UpdateAuthMethodGCP) HasBoundRegions() bool {
	if o != nil && o.BoundRegions != nil {
		return true
	}

	return false
}

// SetBoundRegions gets a reference to the given []string and assigns it to the BoundRegions field.
func (o *UpdateAuthMethodGCP) SetBoundRegions(v []string) {
	o.BoundRegions = &v
}

// GetBoundServiceAccounts returns the BoundServiceAccounts field value if set, zero value otherwise.
func (o *UpdateAuthMethodGCP) GetBoundServiceAccounts() []string {
	if o == nil || o.BoundServiceAccounts == nil {
		var ret []string
		return ret
	}
	return *o.BoundServiceAccounts
}

// GetBoundServiceAccountsOk returns a tuple with the BoundServiceAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodGCP) GetBoundServiceAccountsOk() (*[]string, bool) {
	if o == nil || o.BoundServiceAccounts == nil {
		return nil, false
	}
	return o.BoundServiceAccounts, true
}

// HasBoundServiceAccounts returns a boolean if a field has been set.
func (o *UpdateAuthMethodGCP) HasBoundServiceAccounts() bool {
	if o != nil && o.BoundServiceAccounts != nil {
		return true
	}

	return false
}

// SetBoundServiceAccounts gets a reference to the given []string and assigns it to the BoundServiceAccounts field.
func (o *UpdateAuthMethodGCP) SetBoundServiceAccounts(v []string) {
	o.BoundServiceAccounts = &v
}

// GetBoundZones returns the BoundZones field value if set, zero value otherwise.
func (o *UpdateAuthMethodGCP) GetBoundZones() []string {
	if o == nil || o.BoundZones == nil {
		var ret []string
		return ret
	}
	return *o.BoundZones
}

// GetBoundZonesOk returns a tuple with the BoundZones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodGCP) GetBoundZonesOk() (*[]string, bool) {
	if o == nil || o.BoundZones == nil {
		return nil, false
	}
	return o.BoundZones, true
}

// HasBoundZones returns a boolean if a field has been set.
func (o *UpdateAuthMethodGCP) HasBoundZones() bool {
	if o != nil && o.BoundZones != nil {
		return true
	}

	return false
}

// SetBoundZones gets a reference to the given []string and assigns it to the BoundZones field.
func (o *UpdateAuthMethodGCP) SetBoundZones(v []string) {
	o.BoundZones = &v
}

// GetForceSubClaims returns the ForceSubClaims field value if set, zero value otherwise.
func (o *UpdateAuthMethodGCP) GetForceSubClaims() bool {
	if o == nil || o.ForceSubClaims == nil {
		var ret bool
		return ret
	}
	return *o.ForceSubClaims
}

// GetForceSubClaimsOk returns a tuple with the ForceSubClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodGCP) GetForceSubClaimsOk() (*bool, bool) {
	if o == nil || o.ForceSubClaims == nil {
		return nil, false
	}
	return o.ForceSubClaims, true
}

// HasForceSubClaims returns a boolean if a field has been set.
func (o *UpdateAuthMethodGCP) HasForceSubClaims() bool {
	if o != nil && o.ForceSubClaims != nil {
		return true
	}

	return false
}

// SetForceSubClaims gets a reference to the given bool and assigns it to the ForceSubClaims field.
func (o *UpdateAuthMethodGCP) SetForceSubClaims(v bool) {
	o.ForceSubClaims = &v
}

// GetGwBoundIps returns the GwBoundIps field value if set, zero value otherwise.
func (o *UpdateAuthMethodGCP) GetGwBoundIps() []string {
	if o == nil || o.GwBoundIps == nil {
		var ret []string
		return ret
	}
	return *o.GwBoundIps
}

// GetGwBoundIpsOk returns a tuple with the GwBoundIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodGCP) GetGwBoundIpsOk() (*[]string, bool) {
	if o == nil || o.GwBoundIps == nil {
		return nil, false
	}
	return o.GwBoundIps, true
}

// HasGwBoundIps returns a boolean if a field has been set.
func (o *UpdateAuthMethodGCP) HasGwBoundIps() bool {
	if o != nil && o.GwBoundIps != nil {
		return true
	}

	return false
}

// SetGwBoundIps gets a reference to the given []string and assigns it to the GwBoundIps field.
func (o *UpdateAuthMethodGCP) SetGwBoundIps(v []string) {
	o.GwBoundIps = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *UpdateAuthMethodGCP) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodGCP) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *UpdateAuthMethodGCP) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *UpdateAuthMethodGCP) SetJson(v bool) {
	o.Json = &v
}

// GetJwtTtl returns the JwtTtl field value if set, zero value otherwise.
func (o *UpdateAuthMethodGCP) GetJwtTtl() int64 {
	if o == nil || o.JwtTtl == nil {
		var ret int64
		return ret
	}
	return *o.JwtTtl
}

// GetJwtTtlOk returns a tuple with the JwtTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodGCP) GetJwtTtlOk() (*int64, bool) {
	if o == nil || o.JwtTtl == nil {
		return nil, false
	}
	return o.JwtTtl, true
}

// HasJwtTtl returns a boolean if a field has been set.
func (o *UpdateAuthMethodGCP) HasJwtTtl() bool {
	if o != nil && o.JwtTtl != nil {
		return true
	}

	return false
}

// SetJwtTtl gets a reference to the given int64 and assigns it to the JwtTtl field.
func (o *UpdateAuthMethodGCP) SetJwtTtl(v int64) {
	o.JwtTtl = &v
}

// GetName returns the Name field value
func (o *UpdateAuthMethodGCP) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodGCP) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateAuthMethodGCP) SetName(v string) {
	o.Name = v
}

// GetNewName returns the NewName field value if set, zero value otherwise.
func (o *UpdateAuthMethodGCP) GetNewName() string {
	if o == nil || o.NewName == nil {
		var ret string
		return ret
	}
	return *o.NewName
}

// GetNewNameOk returns a tuple with the NewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodGCP) GetNewNameOk() (*string, bool) {
	if o == nil || o.NewName == nil {
		return nil, false
	}
	return o.NewName, true
}

// HasNewName returns a boolean if a field has been set.
func (o *UpdateAuthMethodGCP) HasNewName() bool {
	if o != nil && o.NewName != nil {
		return true
	}

	return false
}

// SetNewName gets a reference to the given string and assigns it to the NewName field.
func (o *UpdateAuthMethodGCP) SetNewName(v string) {
	o.NewName = &v
}

// GetServiceAccountCredsData returns the ServiceAccountCredsData field value if set, zero value otherwise.
func (o *UpdateAuthMethodGCP) GetServiceAccountCredsData() string {
	if o == nil || o.ServiceAccountCredsData == nil {
		var ret string
		return ret
	}
	return *o.ServiceAccountCredsData
}

// GetServiceAccountCredsDataOk returns a tuple with the ServiceAccountCredsData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodGCP) GetServiceAccountCredsDataOk() (*string, bool) {
	if o == nil || o.ServiceAccountCredsData == nil {
		return nil, false
	}
	return o.ServiceAccountCredsData, true
}

// HasServiceAccountCredsData returns a boolean if a field has been set.
func (o *UpdateAuthMethodGCP) HasServiceAccountCredsData() bool {
	if o != nil && o.ServiceAccountCredsData != nil {
		return true
	}

	return false
}

// SetServiceAccountCredsData gets a reference to the given string and assigns it to the ServiceAccountCredsData field.
func (o *UpdateAuthMethodGCP) SetServiceAccountCredsData(v string) {
	o.ServiceAccountCredsData = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *UpdateAuthMethodGCP) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodGCP) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *UpdateAuthMethodGCP) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *UpdateAuthMethodGCP) SetToken(v string) {
	o.Token = &v
}

// GetType returns the Type field value
func (o *UpdateAuthMethodGCP) GetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodGCP) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *UpdateAuthMethodGCP) SetType(v string) {
	o.Type = v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *UpdateAuthMethodGCP) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodGCP) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *UpdateAuthMethodGCP) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *UpdateAuthMethodGCP) SetUidToken(v string) {
	o.UidToken = &v
}

func (o UpdateAuthMethodGCP) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessExpires != nil {
		toSerialize["access-expires"] = o.AccessExpires
	}
	if true {
		toSerialize["audience"] = o.Audience
	}
	if o.BoundIps != nil {
		toSerialize["bound-ips"] = o.BoundIps
	}
	if o.BoundLabels != nil {
		toSerialize["bound-labels"] = o.BoundLabels
	}
	if o.BoundProjects != nil {
		toSerialize["bound-projects"] = o.BoundProjects
	}
	if o.BoundRegions != nil {
		toSerialize["bound-regions"] = o.BoundRegions
	}
	if o.BoundServiceAccounts != nil {
		toSerialize["bound-service-accounts"] = o.BoundServiceAccounts
	}
	if o.BoundZones != nil {
		toSerialize["bound-zones"] = o.BoundZones
	}
	if o.ForceSubClaims != nil {
		toSerialize["force-sub-claims"] = o.ForceSubClaims
	}
	if o.GwBoundIps != nil {
		toSerialize["gw-bound-ips"] = o.GwBoundIps
	}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if o.JwtTtl != nil {
		toSerialize["jwt-ttl"] = o.JwtTtl
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.NewName != nil {
		toSerialize["new-name"] = o.NewName
	}
	if o.ServiceAccountCredsData != nil {
		toSerialize["service-account-creds-data"] = o.ServiceAccountCredsData
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateAuthMethodGCP struct {
	value *UpdateAuthMethodGCP
	isSet bool
}

func (v NullableUpdateAuthMethodGCP) Get() *UpdateAuthMethodGCP {
	return v.value
}

func (v *NullableUpdateAuthMethodGCP) Set(val *UpdateAuthMethodGCP) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateAuthMethodGCP) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateAuthMethodGCP) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateAuthMethodGCP(val *UpdateAuthMethodGCP) *NullableUpdateAuthMethodGCP {
	return &NullableUpdateAuthMethodGCP{value: val, isSet: true}
}

func (v NullableUpdateAuthMethodGCP) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateAuthMethodGCP) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


