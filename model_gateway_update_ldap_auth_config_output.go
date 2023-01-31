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

// GatewayUpdateLdapAuthConfigOutput struct for GatewayUpdateLdapAuthConfigOutput
type GatewayUpdateLdapAuthConfigOutput struct {
	Updated *bool `json:"updated,omitempty"`
}

// NewGatewayUpdateLdapAuthConfigOutput instantiates a new GatewayUpdateLdapAuthConfigOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayUpdateLdapAuthConfigOutput() *GatewayUpdateLdapAuthConfigOutput {
	this := GatewayUpdateLdapAuthConfigOutput{}
	return &this
}

// NewGatewayUpdateLdapAuthConfigOutputWithDefaults instantiates a new GatewayUpdateLdapAuthConfigOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayUpdateLdapAuthConfigOutputWithDefaults() *GatewayUpdateLdapAuthConfigOutput {
	this := GatewayUpdateLdapAuthConfigOutput{}
	return &this
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *GatewayUpdateLdapAuthConfigOutput) GetUpdated() bool {
	if o == nil || o.Updated == nil {
		var ret bool
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateLdapAuthConfigOutput) GetUpdatedOk() (*bool, bool) {
	if o == nil || o.Updated == nil {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *GatewayUpdateLdapAuthConfigOutput) HasUpdated() bool {
	if o != nil && o.Updated != nil {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given bool and assigns it to the Updated field.
func (o *GatewayUpdateLdapAuthConfigOutput) SetUpdated(v bool) {
	o.Updated = &v
}

func (o GatewayUpdateLdapAuthConfigOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Updated != nil {
		toSerialize["updated"] = o.Updated
	}
	return json.Marshal(toSerialize)
}

type NullableGatewayUpdateLdapAuthConfigOutput struct {
	value *GatewayUpdateLdapAuthConfigOutput
	isSet bool
}

func (v NullableGatewayUpdateLdapAuthConfigOutput) Get() *GatewayUpdateLdapAuthConfigOutput {
	return v.value
}

func (v *NullableGatewayUpdateLdapAuthConfigOutput) Set(val *GatewayUpdateLdapAuthConfigOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayUpdateLdapAuthConfigOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayUpdateLdapAuthConfigOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayUpdateLdapAuthConfigOutput(val *GatewayUpdateLdapAuthConfigOutput) *NullableGatewayUpdateLdapAuthConfigOutput {
	return &NullableGatewayUpdateLdapAuthConfigOutput{value: val, isSet: true}
}

func (v NullableGatewayUpdateLdapAuthConfigOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayUpdateLdapAuthConfigOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

