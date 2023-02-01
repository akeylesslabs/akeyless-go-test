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

// KmipMoveServerOutput struct for KmipMoveServerOutput
type KmipMoveServerOutput struct {
	NewRoot *string `json:"new_root,omitempty"`
	OldRoot *string `json:"old_root,omitempty"`
}

// NewKmipMoveServerOutput instantiates a new KmipMoveServerOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKmipMoveServerOutput() *KmipMoveServerOutput {
	this := KmipMoveServerOutput{}
	return &this
}

// NewKmipMoveServerOutputWithDefaults instantiates a new KmipMoveServerOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKmipMoveServerOutputWithDefaults() *KmipMoveServerOutput {
	this := KmipMoveServerOutput{}
	return &this
}

// GetNewRoot returns the NewRoot field value if set, zero value otherwise.
func (o *KmipMoveServerOutput) GetNewRoot() string {
	if o == nil || o.NewRoot == nil {
		var ret string
		return ret
	}
	return *o.NewRoot
}

// GetNewRootOk returns a tuple with the NewRoot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KmipMoveServerOutput) GetNewRootOk() (*string, bool) {
	if o == nil || o.NewRoot == nil {
		return nil, false
	}
	return o.NewRoot, true
}

// HasNewRoot returns a boolean if a field has been set.
func (o *KmipMoveServerOutput) HasNewRoot() bool {
	if o != nil && o.NewRoot != nil {
		return true
	}

	return false
}

// SetNewRoot gets a reference to the given string and assigns it to the NewRoot field.
func (o *KmipMoveServerOutput) SetNewRoot(v string) {
	o.NewRoot = &v
}

// GetOldRoot returns the OldRoot field value if set, zero value otherwise.
func (o *KmipMoveServerOutput) GetOldRoot() string {
	if o == nil || o.OldRoot == nil {
		var ret string
		return ret
	}
	return *o.OldRoot
}

// GetOldRootOk returns a tuple with the OldRoot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KmipMoveServerOutput) GetOldRootOk() (*string, bool) {
	if o == nil || o.OldRoot == nil {
		return nil, false
	}
	return o.OldRoot, true
}

// HasOldRoot returns a boolean if a field has been set.
func (o *KmipMoveServerOutput) HasOldRoot() bool {
	if o != nil && o.OldRoot != nil {
		return true
	}

	return false
}

// SetOldRoot gets a reference to the given string and assigns it to the OldRoot field.
func (o *KmipMoveServerOutput) SetOldRoot(v string) {
	o.OldRoot = &v
}

func (o KmipMoveServerOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NewRoot != nil {
		toSerialize["new_root"] = o.NewRoot
	}
	if o.OldRoot != nil {
		toSerialize["old_root"] = o.OldRoot
	}
	return json.Marshal(toSerialize)
}

type NullableKmipMoveServerOutput struct {
	value *KmipMoveServerOutput
	isSet bool
}

func (v NullableKmipMoveServerOutput) Get() *KmipMoveServerOutput {
	return v.value
}

func (v *NullableKmipMoveServerOutput) Set(val *KmipMoveServerOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableKmipMoveServerOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableKmipMoveServerOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKmipMoveServerOutput(val *KmipMoveServerOutput) *NullableKmipMoveServerOutput {
	return &NullableKmipMoveServerOutput{value: val, isSet: true}
}

func (v NullableKmipMoveServerOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKmipMoveServerOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


