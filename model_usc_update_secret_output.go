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

// UscUpdateSecretOutput struct for UscUpdateSecretOutput
type UscUpdateSecretOutput struct {
	Name *string `json:"name,omitempty"`
	SecretId *string `json:"secret_id,omitempty"`
	VersionId *string `json:"version_id,omitempty"`
}

// NewUscUpdateSecretOutput instantiates a new UscUpdateSecretOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUscUpdateSecretOutput() *UscUpdateSecretOutput {
	this := UscUpdateSecretOutput{}
	return &this
}

// NewUscUpdateSecretOutputWithDefaults instantiates a new UscUpdateSecretOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUscUpdateSecretOutputWithDefaults() *UscUpdateSecretOutput {
	this := UscUpdateSecretOutput{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UscUpdateSecretOutput) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UscUpdateSecretOutput) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UscUpdateSecretOutput) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UscUpdateSecretOutput) SetName(v string) {
	o.Name = &v
}

// GetSecretId returns the SecretId field value if set, zero value otherwise.
func (o *UscUpdateSecretOutput) GetSecretId() string {
	if o == nil || o.SecretId == nil {
		var ret string
		return ret
	}
	return *o.SecretId
}

// GetSecretIdOk returns a tuple with the SecretId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UscUpdateSecretOutput) GetSecretIdOk() (*string, bool) {
	if o == nil || o.SecretId == nil {
		return nil, false
	}
	return o.SecretId, true
}

// HasSecretId returns a boolean if a field has been set.
func (o *UscUpdateSecretOutput) HasSecretId() bool {
	if o != nil && o.SecretId != nil {
		return true
	}

	return false
}

// SetSecretId gets a reference to the given string and assigns it to the SecretId field.
func (o *UscUpdateSecretOutput) SetSecretId(v string) {
	o.SecretId = &v
}

// GetVersionId returns the VersionId field value if set, zero value otherwise.
func (o *UscUpdateSecretOutput) GetVersionId() string {
	if o == nil || o.VersionId == nil {
		var ret string
		return ret
	}
	return *o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UscUpdateSecretOutput) GetVersionIdOk() (*string, bool) {
	if o == nil || o.VersionId == nil {
		return nil, false
	}
	return o.VersionId, true
}

// HasVersionId returns a boolean if a field has been set.
func (o *UscUpdateSecretOutput) HasVersionId() bool {
	if o != nil && o.VersionId != nil {
		return true
	}

	return false
}

// SetVersionId gets a reference to the given string and assigns it to the VersionId field.
func (o *UscUpdateSecretOutput) SetVersionId(v string) {
	o.VersionId = &v
}

func (o UscUpdateSecretOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.SecretId != nil {
		toSerialize["secret_id"] = o.SecretId
	}
	if o.VersionId != nil {
		toSerialize["version_id"] = o.VersionId
	}
	return json.Marshal(toSerialize)
}

type NullableUscUpdateSecretOutput struct {
	value *UscUpdateSecretOutput
	isSet bool
}

func (v NullableUscUpdateSecretOutput) Get() *UscUpdateSecretOutput {
	return v.value
}

func (v *NullableUscUpdateSecretOutput) Set(val *UscUpdateSecretOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableUscUpdateSecretOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableUscUpdateSecretOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUscUpdateSecretOutput(val *UscUpdateSecretOutput) *NullableUscUpdateSecretOutput {
	return &NullableUscUpdateSecretOutput{value: val, isSet: true}
}

func (v NullableUscUpdateSecretOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUscUpdateSecretOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

