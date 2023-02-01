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

// UidGenerateTokenOutput struct for UidGenerateTokenOutput
type UidGenerateTokenOutput struct {
	Token *string `json:"token,omitempty"`
}

// NewUidGenerateTokenOutput instantiates a new UidGenerateTokenOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUidGenerateTokenOutput() *UidGenerateTokenOutput {
	this := UidGenerateTokenOutput{}
	return &this
}

// NewUidGenerateTokenOutputWithDefaults instantiates a new UidGenerateTokenOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUidGenerateTokenOutputWithDefaults() *UidGenerateTokenOutput {
	this := UidGenerateTokenOutput{}
	return &this
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *UidGenerateTokenOutput) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UidGenerateTokenOutput) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *UidGenerateTokenOutput) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *UidGenerateTokenOutput) SetToken(v string) {
	o.Token = &v
}

func (o UidGenerateTokenOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	return json.Marshal(toSerialize)
}

type NullableUidGenerateTokenOutput struct {
	value *UidGenerateTokenOutput
	isSet bool
}

func (v NullableUidGenerateTokenOutput) Get() *UidGenerateTokenOutput {
	return v.value
}

func (v *NullableUidGenerateTokenOutput) Set(val *UidGenerateTokenOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableUidGenerateTokenOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableUidGenerateTokenOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUidGenerateTokenOutput(val *UidGenerateTokenOutput) *NullableUidGenerateTokenOutput {
	return &NullableUidGenerateTokenOutput{value: val, isSet: true}
}

func (v NullableUidGenerateTokenOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUidGenerateTokenOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


