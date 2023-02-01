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

// UpdateEKSTargetOutput struct for UpdateEKSTargetOutput
type UpdateEKSTargetOutput struct {
	TargetId *int64 `json:"target_id,omitempty"`
}

// NewUpdateEKSTargetOutput instantiates a new UpdateEKSTargetOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateEKSTargetOutput() *UpdateEKSTargetOutput {
	this := UpdateEKSTargetOutput{}
	return &this
}

// NewUpdateEKSTargetOutputWithDefaults instantiates a new UpdateEKSTargetOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateEKSTargetOutputWithDefaults() *UpdateEKSTargetOutput {
	this := UpdateEKSTargetOutput{}
	return &this
}

// GetTargetId returns the TargetId field value if set, zero value otherwise.
func (o *UpdateEKSTargetOutput) GetTargetId() int64 {
	if o == nil || o.TargetId == nil {
		var ret int64
		return ret
	}
	return *o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEKSTargetOutput) GetTargetIdOk() (*int64, bool) {
	if o == nil || o.TargetId == nil {
		return nil, false
	}
	return o.TargetId, true
}

// HasTargetId returns a boolean if a field has been set.
func (o *UpdateEKSTargetOutput) HasTargetId() bool {
	if o != nil && o.TargetId != nil {
		return true
	}

	return false
}

// SetTargetId gets a reference to the given int64 and assigns it to the TargetId field.
func (o *UpdateEKSTargetOutput) SetTargetId(v int64) {
	o.TargetId = &v
}

func (o UpdateEKSTargetOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TargetId != nil {
		toSerialize["target_id"] = o.TargetId
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateEKSTargetOutput struct {
	value *UpdateEKSTargetOutput
	isSet bool
}

func (v NullableUpdateEKSTargetOutput) Get() *UpdateEKSTargetOutput {
	return v.value
}

func (v *NullableUpdateEKSTargetOutput) Set(val *UpdateEKSTargetOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateEKSTargetOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateEKSTargetOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateEKSTargetOutput(val *UpdateEKSTargetOutput) *NullableUpdateEKSTargetOutput {
	return &NullableUpdateEKSTargetOutput{value: val, isSet: true}
}

func (v NullableUpdateEKSTargetOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateEKSTargetOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


