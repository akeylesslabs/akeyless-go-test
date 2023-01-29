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

// LastConfigChange struct for LastConfigChange
type LastConfigChange struct {
	LastK8sAuthsChange *K8SAuthsConfigLastChange `json:"last_k8s_auths_change,omitempty"`
	LastMigrationsChange *MigrationsConfigLastChange `json:"last_migrations_change,omitempty"`
}

// NewLastConfigChange instantiates a new LastConfigChange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLastConfigChange() *LastConfigChange {
	this := LastConfigChange{}
	return &this
}

// NewLastConfigChangeWithDefaults instantiates a new LastConfigChange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLastConfigChangeWithDefaults() *LastConfigChange {
	this := LastConfigChange{}
	return &this
}

// GetLastK8sAuthsChange returns the LastK8sAuthsChange field value if set, zero value otherwise.
func (o *LastConfigChange) GetLastK8sAuthsChange() K8SAuthsConfigLastChange {
	if o == nil || o.LastK8sAuthsChange == nil {
		var ret K8SAuthsConfigLastChange
		return ret
	}
	return *o.LastK8sAuthsChange
}

// GetLastK8sAuthsChangeOk returns a tuple with the LastK8sAuthsChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LastConfigChange) GetLastK8sAuthsChangeOk() (*K8SAuthsConfigLastChange, bool) {
	if o == nil || o.LastK8sAuthsChange == nil {
		return nil, false
	}
	return o.LastK8sAuthsChange, true
}

// HasLastK8sAuthsChange returns a boolean if a field has been set.
func (o *LastConfigChange) HasLastK8sAuthsChange() bool {
	if o != nil && o.LastK8sAuthsChange != nil {
		return true
	}

	return false
}

// SetLastK8sAuthsChange gets a reference to the given K8SAuthsConfigLastChange and assigns it to the LastK8sAuthsChange field.
func (o *LastConfigChange) SetLastK8sAuthsChange(v K8SAuthsConfigLastChange) {
	o.LastK8sAuthsChange = &v
}

// GetLastMigrationsChange returns the LastMigrationsChange field value if set, zero value otherwise.
func (o *LastConfigChange) GetLastMigrationsChange() MigrationsConfigLastChange {
	if o == nil || o.LastMigrationsChange == nil {
		var ret MigrationsConfigLastChange
		return ret
	}
	return *o.LastMigrationsChange
}

// GetLastMigrationsChangeOk returns a tuple with the LastMigrationsChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LastConfigChange) GetLastMigrationsChangeOk() (*MigrationsConfigLastChange, bool) {
	if o == nil || o.LastMigrationsChange == nil {
		return nil, false
	}
	return o.LastMigrationsChange, true
}

// HasLastMigrationsChange returns a boolean if a field has been set.
func (o *LastConfigChange) HasLastMigrationsChange() bool {
	if o != nil && o.LastMigrationsChange != nil {
		return true
	}

	return false
}

// SetLastMigrationsChange gets a reference to the given MigrationsConfigLastChange and assigns it to the LastMigrationsChange field.
func (o *LastConfigChange) SetLastMigrationsChange(v MigrationsConfigLastChange) {
	o.LastMigrationsChange = &v
}

func (o LastConfigChange) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LastK8sAuthsChange != nil {
		toSerialize["last_k8s_auths_change"] = o.LastK8sAuthsChange
	}
	if o.LastMigrationsChange != nil {
		toSerialize["last_migrations_change"] = o.LastMigrationsChange
	}
	return json.Marshal(toSerialize)
}

type NullableLastConfigChange struct {
	value *LastConfigChange
	isSet bool
}

func (v NullableLastConfigChange) Get() *LastConfigChange {
	return v.value
}

func (v *NullableLastConfigChange) Set(val *LastConfigChange) {
	v.value = val
	v.isSet = true
}

func (v NullableLastConfigChange) IsSet() bool {
	return v.isSet
}

func (v *NullableLastConfigChange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLastConfigChange(val *LastConfigChange) *NullableLastConfigChange {
	return &NullableLastConfigChange{value: val, isSet: true}
}

func (v NullableLastConfigChange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLastConfigChange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


