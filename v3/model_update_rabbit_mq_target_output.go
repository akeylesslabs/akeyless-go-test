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

// UpdateRabbitMQTargetOutput struct for UpdateRabbitMQTargetOutput
type UpdateRabbitMQTargetOutput struct {
	TargetId *int64 `json:"target_id,omitempty"`
}

// NewUpdateRabbitMQTargetOutput instantiates a new UpdateRabbitMQTargetOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateRabbitMQTargetOutput() *UpdateRabbitMQTargetOutput {
	this := UpdateRabbitMQTargetOutput{}
	return &this
}

// NewUpdateRabbitMQTargetOutputWithDefaults instantiates a new UpdateRabbitMQTargetOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateRabbitMQTargetOutputWithDefaults() *UpdateRabbitMQTargetOutput {
	this := UpdateRabbitMQTargetOutput{}
	return &this
}

// GetTargetId returns the TargetId field value if set, zero value otherwise.
func (o *UpdateRabbitMQTargetOutput) GetTargetId() int64 {
	if o == nil || o.TargetId == nil {
		var ret int64
		return ret
	}
	return *o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRabbitMQTargetOutput) GetTargetIdOk() (*int64, bool) {
	if o == nil || o.TargetId == nil {
		return nil, false
	}
	return o.TargetId, true
}

// HasTargetId returns a boolean if a field has been set.
func (o *UpdateRabbitMQTargetOutput) HasTargetId() bool {
	if o != nil && o.TargetId != nil {
		return true
	}

	return false
}

// SetTargetId gets a reference to the given int64 and assigns it to the TargetId field.
func (o *UpdateRabbitMQTargetOutput) SetTargetId(v int64) {
	o.TargetId = &v
}

func (o UpdateRabbitMQTargetOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TargetId != nil {
		toSerialize["target_id"] = o.TargetId
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateRabbitMQTargetOutput struct {
	value *UpdateRabbitMQTargetOutput
	isSet bool
}

func (v NullableUpdateRabbitMQTargetOutput) Get() *UpdateRabbitMQTargetOutput {
	return v.value
}

func (v *NullableUpdateRabbitMQTargetOutput) Set(val *UpdateRabbitMQTargetOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateRabbitMQTargetOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateRabbitMQTargetOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateRabbitMQTargetOutput(val *UpdateRabbitMQTargetOutput) *NullableUpdateRabbitMQTargetOutput {
	return &NullableUpdateRabbitMQTargetOutput{value: val, isSet: true}
}

func (v NullableUpdateRabbitMQTargetOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateRabbitMQTargetOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


