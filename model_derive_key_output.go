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

// DeriveKeyOutput struct for DeriveKeyOutput
type DeriveKeyOutput struct {
	Key *string `json:"Key,omitempty"`
	Salt *string `json:"Salt,omitempty"`
}

// NewDeriveKeyOutput instantiates a new DeriveKeyOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeriveKeyOutput() *DeriveKeyOutput {
	this := DeriveKeyOutput{}
	return &this
}

// NewDeriveKeyOutputWithDefaults instantiates a new DeriveKeyOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeriveKeyOutputWithDefaults() *DeriveKeyOutput {
	this := DeriveKeyOutput{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *DeriveKeyOutput) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeriveKeyOutput) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *DeriveKeyOutput) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *DeriveKeyOutput) SetKey(v string) {
	o.Key = &v
}

// GetSalt returns the Salt field value if set, zero value otherwise.
func (o *DeriveKeyOutput) GetSalt() string {
	if o == nil || o.Salt == nil {
		var ret string
		return ret
	}
	return *o.Salt
}

// GetSaltOk returns a tuple with the Salt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeriveKeyOutput) GetSaltOk() (*string, bool) {
	if o == nil || o.Salt == nil {
		return nil, false
	}
	return o.Salt, true
}

// HasSalt returns a boolean if a field has been set.
func (o *DeriveKeyOutput) HasSalt() bool {
	if o != nil && o.Salt != nil {
		return true
	}

	return false
}

// SetSalt gets a reference to the given string and assigns it to the Salt field.
func (o *DeriveKeyOutput) SetSalt(v string) {
	o.Salt = &v
}

func (o DeriveKeyOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Key != nil {
		toSerialize["Key"] = o.Key
	}
	if o.Salt != nil {
		toSerialize["Salt"] = o.Salt
	}
	return json.Marshal(toSerialize)
}

type NullableDeriveKeyOutput struct {
	value *DeriveKeyOutput
	isSet bool
}

func (v NullableDeriveKeyOutput) Get() *DeriveKeyOutput {
	return v.value
}

func (v *NullableDeriveKeyOutput) Set(val *DeriveKeyOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableDeriveKeyOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableDeriveKeyOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeriveKeyOutput(val *DeriveKeyOutput) *NullableDeriveKeyOutput {
	return &NullableDeriveKeyOutput{value: val, isSet: true}
}

func (v NullableDeriveKeyOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeriveKeyOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

