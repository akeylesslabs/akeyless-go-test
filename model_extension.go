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

// Extension struct for Extension
type Extension struct {
	Critical *bool `json:"Critical,omitempty"`
	Name *string `json:"Name,omitempty"`
	Value *string `json:"Value,omitempty"`
}

// NewExtension instantiates a new Extension object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtension() *Extension {
	this := Extension{}
	return &this
}

// NewExtensionWithDefaults instantiates a new Extension object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtensionWithDefaults() *Extension {
	this := Extension{}
	return &this
}

// GetCritical returns the Critical field value if set, zero value otherwise.
func (o *Extension) GetCritical() bool {
	if o == nil || o.Critical == nil {
		var ret bool
		return ret
	}
	return *o.Critical
}

// GetCriticalOk returns a tuple with the Critical field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Extension) GetCriticalOk() (*bool, bool) {
	if o == nil || o.Critical == nil {
		return nil, false
	}
	return o.Critical, true
}

// HasCritical returns a boolean if a field has been set.
func (o *Extension) HasCritical() bool {
	if o != nil && o.Critical != nil {
		return true
	}

	return false
}

// SetCritical gets a reference to the given bool and assigns it to the Critical field.
func (o *Extension) SetCritical(v bool) {
	o.Critical = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Extension) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Extension) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Extension) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Extension) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *Extension) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Extension) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *Extension) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *Extension) SetValue(v string) {
	o.Value = &v
}

func (o Extension) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Critical != nil {
		toSerialize["Critical"] = o.Critical
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Value != nil {
		toSerialize["Value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableExtension struct {
	value *Extension
	isSet bool
}

func (v NullableExtension) Get() *Extension {
	return v.value
}

func (v *NullableExtension) Set(val *Extension) {
	v.value = val
	v.isSet = true
}

func (v NullableExtension) IsSet() bool {
	return v.isSet
}

func (v *NullableExtension) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtension(val *Extension) *NullableExtension {
	return &NullableExtension{value: val, isSet: true}
}

func (v NullableExtension) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtension) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


