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

// GetTargetDetailsOutput struct for GetTargetDetailsOutput
type GetTargetDetailsOutput struct {
	Target *Target `json:"target,omitempty"`
	Value *TargetTypeDetailsInput `json:"value,omitempty"`
}

// NewGetTargetDetailsOutput instantiates a new GetTargetDetailsOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTargetDetailsOutput() *GetTargetDetailsOutput {
	this := GetTargetDetailsOutput{}
	return &this
}

// NewGetTargetDetailsOutputWithDefaults instantiates a new GetTargetDetailsOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTargetDetailsOutputWithDefaults() *GetTargetDetailsOutput {
	this := GetTargetDetailsOutput{}
	return &this
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *GetTargetDetailsOutput) GetTarget() Target {
	if o == nil || o.Target == nil {
		var ret Target
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTargetDetailsOutput) GetTargetOk() (*Target, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *GetTargetDetailsOutput) HasTarget() bool {
	if o != nil && o.Target != nil {
		return true
	}

	return false
}

// SetTarget gets a reference to the given Target and assigns it to the Target field.
func (o *GetTargetDetailsOutput) SetTarget(v Target) {
	o.Target = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *GetTargetDetailsOutput) GetValue() TargetTypeDetailsInput {
	if o == nil || o.Value == nil {
		var ret TargetTypeDetailsInput
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTargetDetailsOutput) GetValueOk() (*TargetTypeDetailsInput, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *GetTargetDetailsOutput) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given TargetTypeDetailsInput and assigns it to the Value field.
func (o *GetTargetDetailsOutput) SetValue(v TargetTypeDetailsInput) {
	o.Value = &v
}

func (o GetTargetDetailsOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableGetTargetDetailsOutput struct {
	value *GetTargetDetailsOutput
	isSet bool
}

func (v NullableGetTargetDetailsOutput) Get() *GetTargetDetailsOutput {
	return v.value
}

func (v *NullableGetTargetDetailsOutput) Set(val *GetTargetDetailsOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTargetDetailsOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTargetDetailsOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTargetDetailsOutput(val *GetTargetDetailsOutput) *NullableGetTargetDetailsOutput {
	return &NullableGetTargetDetailsOutput{value: val, isSet: true}
}

func (v NullableGetTargetDetailsOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTargetDetailsOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

