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

// UpdateRotatedSecretOutput struct for UpdateRotatedSecretOutput
type UpdateRotatedSecretOutput struct {
	Name *string `json:"name,omitempty"`
}

// NewUpdateRotatedSecretOutput instantiates a new UpdateRotatedSecretOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateRotatedSecretOutput() *UpdateRotatedSecretOutput {
	this := UpdateRotatedSecretOutput{}
	return &this
}

// NewUpdateRotatedSecretOutputWithDefaults instantiates a new UpdateRotatedSecretOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateRotatedSecretOutputWithDefaults() *UpdateRotatedSecretOutput {
	this := UpdateRotatedSecretOutput{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateRotatedSecretOutput) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRotatedSecretOutput) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateRotatedSecretOutput) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateRotatedSecretOutput) SetName(v string) {
	o.Name = &v
}

func (o UpdateRotatedSecretOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateRotatedSecretOutput struct {
	value *UpdateRotatedSecretOutput
	isSet bool
}

func (v NullableUpdateRotatedSecretOutput) Get() *UpdateRotatedSecretOutput {
	return v.value
}

func (v *NullableUpdateRotatedSecretOutput) Set(val *UpdateRotatedSecretOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateRotatedSecretOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateRotatedSecretOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateRotatedSecretOutput(val *UpdateRotatedSecretOutput) *NullableUpdateRotatedSecretOutput {
	return &NullableUpdateRotatedSecretOutput{value: val, isSet: true}
}

func (v NullableUpdateRotatedSecretOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateRotatedSecretOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


