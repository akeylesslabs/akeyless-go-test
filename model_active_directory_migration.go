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

// ActiveDirectoryMigration struct for ActiveDirectoryMigration
type ActiveDirectoryMigration struct {
	General *MigrationGeneral `json:"general,omitempty"`
	Payload *ActiveDirectoryPayload `json:"payload,omitempty"`
}

// NewActiveDirectoryMigration instantiates a new ActiveDirectoryMigration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActiveDirectoryMigration() *ActiveDirectoryMigration {
	this := ActiveDirectoryMigration{}
	return &this
}

// NewActiveDirectoryMigrationWithDefaults instantiates a new ActiveDirectoryMigration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActiveDirectoryMigrationWithDefaults() *ActiveDirectoryMigration {
	this := ActiveDirectoryMigration{}
	return &this
}

// GetGeneral returns the General field value if set, zero value otherwise.
func (o *ActiveDirectoryMigration) GetGeneral() MigrationGeneral {
	if o == nil || o.General == nil {
		var ret MigrationGeneral
		return ret
	}
	return *o.General
}

// GetGeneralOk returns a tuple with the General field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveDirectoryMigration) GetGeneralOk() (*MigrationGeneral, bool) {
	if o == nil || o.General == nil {
		return nil, false
	}
	return o.General, true
}

// HasGeneral returns a boolean if a field has been set.
func (o *ActiveDirectoryMigration) HasGeneral() bool {
	if o != nil && o.General != nil {
		return true
	}

	return false
}

// SetGeneral gets a reference to the given MigrationGeneral and assigns it to the General field.
func (o *ActiveDirectoryMigration) SetGeneral(v MigrationGeneral) {
	o.General = &v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *ActiveDirectoryMigration) GetPayload() ActiveDirectoryPayload {
	if o == nil || o.Payload == nil {
		var ret ActiveDirectoryPayload
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveDirectoryMigration) GetPayloadOk() (*ActiveDirectoryPayload, bool) {
	if o == nil || o.Payload == nil {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *ActiveDirectoryMigration) HasPayload() bool {
	if o != nil && o.Payload != nil {
		return true
	}

	return false
}

// SetPayload gets a reference to the given ActiveDirectoryPayload and assigns it to the Payload field.
func (o *ActiveDirectoryMigration) SetPayload(v ActiveDirectoryPayload) {
	o.Payload = &v
}

func (o ActiveDirectoryMigration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.General != nil {
		toSerialize["general"] = o.General
	}
	if o.Payload != nil {
		toSerialize["payload"] = o.Payload
	}
	return json.Marshal(toSerialize)
}

type NullableActiveDirectoryMigration struct {
	value *ActiveDirectoryMigration
	isSet bool
}

func (v NullableActiveDirectoryMigration) Get() *ActiveDirectoryMigration {
	return v.value
}

func (v *NullableActiveDirectoryMigration) Set(val *ActiveDirectoryMigration) {
	v.value = val
	v.isSet = true
}

func (v NullableActiveDirectoryMigration) IsSet() bool {
	return v.isSet
}

func (v *NullableActiveDirectoryMigration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActiveDirectoryMigration(val *ActiveDirectoryMigration) *NullableActiveDirectoryMigration {
	return &NullableActiveDirectoryMigration{value: val, isSet: true}
}

func (v NullableActiveDirectoryMigration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActiveDirectoryMigration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


