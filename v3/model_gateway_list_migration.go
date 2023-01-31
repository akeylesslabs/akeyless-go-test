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

// GatewayListMigration gatewayListMigration is a command that list migration
type GatewayListMigration struct {
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewGatewayListMigration instantiates a new GatewayListMigration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayListMigration() *GatewayListMigration {
	this := GatewayListMigration{}
	return &this
}

// NewGatewayListMigrationWithDefaults instantiates a new GatewayListMigration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayListMigrationWithDefaults() *GatewayListMigration {
	this := GatewayListMigration{}
	return &this
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *GatewayListMigration) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayListMigration) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *GatewayListMigration) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *GatewayListMigration) SetJson(v bool) {
	o.Json = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GatewayListMigration) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayListMigration) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GatewayListMigration) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GatewayListMigration) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *GatewayListMigration) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayListMigration) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *GatewayListMigration) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *GatewayListMigration) SetUidToken(v string) {
	o.UidToken = &v
}

func (o GatewayListMigration) MarshalJSON() ([]byte, error) {
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

type NullableGatewayListMigration struct {
	value *GatewayListMigration
	isSet bool
}

func (v NullableGatewayListMigration) Get() *GatewayListMigration {
	return v.value
}

func (v *NullableGatewayListMigration) Set(val *GatewayListMigration) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayListMigration) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayListMigration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayListMigration(val *GatewayListMigration) *NullableGatewayListMigration {
	return &NullableGatewayListMigration{value: val, isSet: true}
}

func (v NullableGatewayListMigration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayListMigration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


