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

// UpdateSalesforceTargetOutput struct for UpdateSalesforceTargetOutput
type UpdateSalesforceTargetOutput struct {
	TargetId *int64 `json:"target_id,omitempty"`
}

// NewUpdateSalesforceTargetOutput instantiates a new UpdateSalesforceTargetOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSalesforceTargetOutput() *UpdateSalesforceTargetOutput {
	this := UpdateSalesforceTargetOutput{}
	return &this
}

// NewUpdateSalesforceTargetOutputWithDefaults instantiates a new UpdateSalesforceTargetOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSalesforceTargetOutputWithDefaults() *UpdateSalesforceTargetOutput {
	this := UpdateSalesforceTargetOutput{}
	return &this
}

// GetTargetId returns the TargetId field value if set, zero value otherwise.
func (o *UpdateSalesforceTargetOutput) GetTargetId() int64 {
	if o == nil || o.TargetId == nil {
		var ret int64
		return ret
	}
	return *o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSalesforceTargetOutput) GetTargetIdOk() (*int64, bool) {
	if o == nil || o.TargetId == nil {
		return nil, false
	}
	return o.TargetId, true
}

// HasTargetId returns a boolean if a field has been set.
func (o *UpdateSalesforceTargetOutput) HasTargetId() bool {
	if o != nil && o.TargetId != nil {
		return true
	}

	return false
}

// SetTargetId gets a reference to the given int64 and assigns it to the TargetId field.
func (o *UpdateSalesforceTargetOutput) SetTargetId(v int64) {
	o.TargetId = &v
}

func (o UpdateSalesforceTargetOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TargetId != nil {
		toSerialize["target_id"] = o.TargetId
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateSalesforceTargetOutput struct {
	value *UpdateSalesforceTargetOutput
	isSet bool
}

func (v NullableUpdateSalesforceTargetOutput) Get() *UpdateSalesforceTargetOutput {
	return v.value
}

func (v *NullableUpdateSalesforceTargetOutput) Set(val *UpdateSalesforceTargetOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSalesforceTargetOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSalesforceTargetOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSalesforceTargetOutput(val *UpdateSalesforceTargetOutput) *NullableUpdateSalesforceTargetOutput {
	return &NullableUpdateSalesforceTargetOutput{value: val, isSet: true}
}

func (v NullableUpdateSalesforceTargetOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSalesforceTargetOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


