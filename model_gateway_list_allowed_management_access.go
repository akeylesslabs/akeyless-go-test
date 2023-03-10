/*
 * Akeyless API
 *
 * The purpose of this application is to provide access to Akeyless API.
 *
 * API version: 2.0
 * Contact: support@akeyless.io
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package akeyless

import (
	"encoding/json"
)

// GatewayListAllowedManagementAccess gatewayListAllowedManagementAccess is a command that returns list sub admins
type GatewayListAllowedManagementAccess struct {
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewGatewayListAllowedManagementAccess instantiates a new GatewayListAllowedManagementAccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayListAllowedManagementAccess() *GatewayListAllowedManagementAccess {
	this := GatewayListAllowedManagementAccess{}
	return &this
}

// NewGatewayListAllowedManagementAccessWithDefaults instantiates a new GatewayListAllowedManagementAccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayListAllowedManagementAccessWithDefaults() *GatewayListAllowedManagementAccess {
	this := GatewayListAllowedManagementAccess{}
	return &this
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *GatewayListAllowedManagementAccess) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayListAllowedManagementAccess) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *GatewayListAllowedManagementAccess) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *GatewayListAllowedManagementAccess) SetJson(v bool) {
	o.Json = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GatewayListAllowedManagementAccess) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayListAllowedManagementAccess) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GatewayListAllowedManagementAccess) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GatewayListAllowedManagementAccess) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *GatewayListAllowedManagementAccess) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayListAllowedManagementAccess) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *GatewayListAllowedManagementAccess) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *GatewayListAllowedManagementAccess) SetUidToken(v string) {
	o.UidToken = &v
}

func (o GatewayListAllowedManagementAccess) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableGatewayListAllowedManagementAccess struct {
	value *GatewayListAllowedManagementAccess
	isSet bool
}

func (v NullableGatewayListAllowedManagementAccess) Get() *GatewayListAllowedManagementAccess {
	return v.value
}

func (v *NullableGatewayListAllowedManagementAccess) Set(val *GatewayListAllowedManagementAccess) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayListAllowedManagementAccess) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayListAllowedManagementAccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayListAllowedManagementAccess(val *GatewayListAllowedManagementAccess) *NullableGatewayListAllowedManagementAccess {
	return &NullableGatewayListAllowedManagementAccess{value: val, isSet: true}
}

func (v NullableGatewayListAllowedManagementAccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayListAllowedManagementAccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


