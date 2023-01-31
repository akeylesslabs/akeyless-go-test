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

// StaticCredsAuthOutput struct for StaticCredsAuthOutput
type StaticCredsAuthOutput struct {
	Token *string `json:"token,omitempty"`
}

// NewStaticCredsAuthOutput instantiates a new StaticCredsAuthOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStaticCredsAuthOutput() *StaticCredsAuthOutput {
	this := StaticCredsAuthOutput{}
	return &this
}

// NewStaticCredsAuthOutputWithDefaults instantiates a new StaticCredsAuthOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStaticCredsAuthOutputWithDefaults() *StaticCredsAuthOutput {
	this := StaticCredsAuthOutput{}
	return &this
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *StaticCredsAuthOutput) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaticCredsAuthOutput) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *StaticCredsAuthOutput) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *StaticCredsAuthOutput) SetToken(v string) {
	o.Token = &v
}

func (o StaticCredsAuthOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	return json.Marshal(toSerialize)
}

type NullableStaticCredsAuthOutput struct {
	value *StaticCredsAuthOutput
	isSet bool
}

func (v NullableStaticCredsAuthOutput) Get() *StaticCredsAuthOutput {
	return v.value
}

func (v *NullableStaticCredsAuthOutput) Set(val *StaticCredsAuthOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableStaticCredsAuthOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableStaticCredsAuthOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStaticCredsAuthOutput(val *StaticCredsAuthOutput) *NullableStaticCredsAuthOutput {
	return &NullableStaticCredsAuthOutput{value: val, isSet: true}
}

func (v NullableStaticCredsAuthOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStaticCredsAuthOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

