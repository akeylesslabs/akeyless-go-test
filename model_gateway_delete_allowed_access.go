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

// GatewayDeleteAllowedAccess gatewayDeleteAllowedAccess is a command that deletes allowed access from gateway
type GatewayDeleteAllowedAccess struct {
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Allowed access name to delete
	Name string `json:"name"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewGatewayDeleteAllowedAccess instantiates a new GatewayDeleteAllowedAccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayDeleteAllowedAccess(name string, ) *GatewayDeleteAllowedAccess {
	this := GatewayDeleteAllowedAccess{}
	var json bool = false
	this.Json = &json
	this.Name = name
	return &this
}

// NewGatewayDeleteAllowedAccessWithDefaults instantiates a new GatewayDeleteAllowedAccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayDeleteAllowedAccessWithDefaults() *GatewayDeleteAllowedAccess {
	this := GatewayDeleteAllowedAccess{}
	var json bool = false
	this.Json = &json
	return &this
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *GatewayDeleteAllowedAccess) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayDeleteAllowedAccess) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *GatewayDeleteAllowedAccess) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *GatewayDeleteAllowedAccess) SetJson(v bool) {
	o.Json = &v
}

// GetName returns the Name field value
func (o *GatewayDeleteAllowedAccess) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GatewayDeleteAllowedAccess) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GatewayDeleteAllowedAccess) SetName(v string) {
	o.Name = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GatewayDeleteAllowedAccess) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayDeleteAllowedAccess) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GatewayDeleteAllowedAccess) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GatewayDeleteAllowedAccess) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *GatewayDeleteAllowedAccess) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayDeleteAllowedAccess) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *GatewayDeleteAllowedAccess) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *GatewayDeleteAllowedAccess) SetUidToken(v string) {
	o.UidToken = &v
}

func (o GatewayDeleteAllowedAccess) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableGatewayDeleteAllowedAccess struct {
	value *GatewayDeleteAllowedAccess
	isSet bool
}

func (v NullableGatewayDeleteAllowedAccess) Get() *GatewayDeleteAllowedAccess {
	return v.value
}

func (v *NullableGatewayDeleteAllowedAccess) Set(val *GatewayDeleteAllowedAccess) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayDeleteAllowedAccess) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayDeleteAllowedAccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayDeleteAllowedAccess(val *GatewayDeleteAllowedAccess) *NullableGatewayDeleteAllowedAccess {
	return &NullableGatewayDeleteAllowedAccess{value: val, isSet: true}
}

func (v NullableGatewayDeleteAllowedAccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayDeleteAllowedAccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


