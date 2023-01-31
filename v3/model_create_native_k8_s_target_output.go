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

// CreateNativeK8STargetOutput struct for CreateNativeK8STargetOutput
type CreateNativeK8STargetOutput struct {
	TargetId *int64 `json:"target_id,omitempty"`
}

// NewCreateNativeK8STargetOutput instantiates a new CreateNativeK8STargetOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNativeK8STargetOutput() *CreateNativeK8STargetOutput {
	this := CreateNativeK8STargetOutput{}
	return &this
}

// NewCreateNativeK8STargetOutputWithDefaults instantiates a new CreateNativeK8STargetOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNativeK8STargetOutputWithDefaults() *CreateNativeK8STargetOutput {
	this := CreateNativeK8STargetOutput{}
	return &this
}

// GetTargetId returns the TargetId field value if set, zero value otherwise.
func (o *CreateNativeK8STargetOutput) GetTargetId() int64 {
	if o == nil || o.TargetId == nil {
		var ret int64
		return ret
	}
	return *o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNativeK8STargetOutput) GetTargetIdOk() (*int64, bool) {
	if o == nil || o.TargetId == nil {
		return nil, false
	}
	return o.TargetId, true
}

// HasTargetId returns a boolean if a field has been set.
func (o *CreateNativeK8STargetOutput) HasTargetId() bool {
	if o != nil && o.TargetId != nil {
		return true
	}

	return false
}

// SetTargetId gets a reference to the given int64 and assigns it to the TargetId field.
func (o *CreateNativeK8STargetOutput) SetTargetId(v int64) {
	o.TargetId = &v
}

func (o CreateNativeK8STargetOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TargetId != nil {
		toSerialize["target_id"] = o.TargetId
	}
	return json.Marshal(toSerialize)
}

type NullableCreateNativeK8STargetOutput struct {
	value *CreateNativeK8STargetOutput
	isSet bool
}

func (v NullableCreateNativeK8STargetOutput) Get() *CreateNativeK8STargetOutput {
	return v.value
}

func (v *NullableCreateNativeK8STargetOutput) Set(val *CreateNativeK8STargetOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNativeK8STargetOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNativeK8STargetOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNativeK8STargetOutput(val *CreateNativeK8STargetOutput) *NullableCreateNativeK8STargetOutput {
	return &NullableCreateNativeK8STargetOutput{value: val, isSet: true}
}

func (v NullableCreateNativeK8STargetOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNativeK8STargetOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


