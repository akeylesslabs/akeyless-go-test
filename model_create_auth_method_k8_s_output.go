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

// CreateAuthMethodK8SOutput struct for CreateAuthMethodK8SOutput
type CreateAuthMethodK8SOutput struct {
	AccessId *string `json:"access_id,omitempty"`
	PrvKey *string `json:"prv_key,omitempty"`
}

// NewCreateAuthMethodK8SOutput instantiates a new CreateAuthMethodK8SOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAuthMethodK8SOutput() *CreateAuthMethodK8SOutput {
	this := CreateAuthMethodK8SOutput{}
	return &this
}

// NewCreateAuthMethodK8SOutputWithDefaults instantiates a new CreateAuthMethodK8SOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAuthMethodK8SOutputWithDefaults() *CreateAuthMethodK8SOutput {
	this := CreateAuthMethodK8SOutput{}
	return &this
}

// GetAccessId returns the AccessId field value if set, zero value otherwise.
func (o *CreateAuthMethodK8SOutput) GetAccessId() string {
	if o == nil || o.AccessId == nil {
		var ret string
		return ret
	}
	return *o.AccessId
}

// GetAccessIdOk returns a tuple with the AccessId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodK8SOutput) GetAccessIdOk() (*string, bool) {
	if o == nil || o.AccessId == nil {
		return nil, false
	}
	return o.AccessId, true
}

// HasAccessId returns a boolean if a field has been set.
func (o *CreateAuthMethodK8SOutput) HasAccessId() bool {
	if o != nil && o.AccessId != nil {
		return true
	}

	return false
}

// SetAccessId gets a reference to the given string and assigns it to the AccessId field.
func (o *CreateAuthMethodK8SOutput) SetAccessId(v string) {
	o.AccessId = &v
}

// GetPrvKey returns the PrvKey field value if set, zero value otherwise.
func (o *CreateAuthMethodK8SOutput) GetPrvKey() string {
	if o == nil || o.PrvKey == nil {
		var ret string
		return ret
	}
	return *o.PrvKey
}

// GetPrvKeyOk returns a tuple with the PrvKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodK8SOutput) GetPrvKeyOk() (*string, bool) {
	if o == nil || o.PrvKey == nil {
		return nil, false
	}
	return o.PrvKey, true
}

// HasPrvKey returns a boolean if a field has been set.
func (o *CreateAuthMethodK8SOutput) HasPrvKey() bool {
	if o != nil && o.PrvKey != nil {
		return true
	}

	return false
}

// SetPrvKey gets a reference to the given string and assigns it to the PrvKey field.
func (o *CreateAuthMethodK8SOutput) SetPrvKey(v string) {
	o.PrvKey = &v
}

func (o CreateAuthMethodK8SOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessId != nil {
		toSerialize["access_id"] = o.AccessId
	}
	if o.PrvKey != nil {
		toSerialize["prv_key"] = o.PrvKey
	}
	return json.Marshal(toSerialize)
}

type NullableCreateAuthMethodK8SOutput struct {
	value *CreateAuthMethodK8SOutput
	isSet bool
}

func (v NullableCreateAuthMethodK8SOutput) Get() *CreateAuthMethodK8SOutput {
	return v.value
}

func (v *NullableCreateAuthMethodK8SOutput) Set(val *CreateAuthMethodK8SOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAuthMethodK8SOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAuthMethodK8SOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAuthMethodK8SOutput(val *CreateAuthMethodK8SOutput) *NullableCreateAuthMethodK8SOutput {
	return &NullableCreateAuthMethodK8SOutput{value: val, isSet: true}
}

func (v NullableCreateAuthMethodK8SOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAuthMethodK8SOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


