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

// GatewayMigrationSyncOutput struct for GatewayMigrationSyncOutput
type GatewayMigrationSyncOutput struct {
	MigrationName *string `json:"migration_name,omitempty"`
}

// NewGatewayMigrationSyncOutput instantiates a new GatewayMigrationSyncOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayMigrationSyncOutput() *GatewayMigrationSyncOutput {
	this := GatewayMigrationSyncOutput{}
	return &this
}

// NewGatewayMigrationSyncOutputWithDefaults instantiates a new GatewayMigrationSyncOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayMigrationSyncOutputWithDefaults() *GatewayMigrationSyncOutput {
	this := GatewayMigrationSyncOutput{}
	return &this
}

// GetMigrationName returns the MigrationName field value if set, zero value otherwise.
func (o *GatewayMigrationSyncOutput) GetMigrationName() string {
	if o == nil || o.MigrationName == nil {
		var ret string
		return ret
	}
	return *o.MigrationName
}

// GetMigrationNameOk returns a tuple with the MigrationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayMigrationSyncOutput) GetMigrationNameOk() (*string, bool) {
	if o == nil || o.MigrationName == nil {
		return nil, false
	}
	return o.MigrationName, true
}

// HasMigrationName returns a boolean if a field has been set.
func (o *GatewayMigrationSyncOutput) HasMigrationName() bool {
	if o != nil && o.MigrationName != nil {
		return true
	}

	return false
}

// SetMigrationName gets a reference to the given string and assigns it to the MigrationName field.
func (o *GatewayMigrationSyncOutput) SetMigrationName(v string) {
	o.MigrationName = &v
}

func (o GatewayMigrationSyncOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MigrationName != nil {
		toSerialize["migration_name"] = o.MigrationName
	}
	return json.Marshal(toSerialize)
}

type NullableGatewayMigrationSyncOutput struct {
	value *GatewayMigrationSyncOutput
	isSet bool
}

func (v NullableGatewayMigrationSyncOutput) Get() *GatewayMigrationSyncOutput {
	return v.value
}

func (v *NullableGatewayMigrationSyncOutput) Set(val *GatewayMigrationSyncOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayMigrationSyncOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayMigrationSyncOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayMigrationSyncOutput(val *GatewayMigrationSyncOutput) *NullableGatewayMigrationSyncOutput {
	return &NullableGatewayMigrationSyncOutput{value: val, isSet: true}
}

func (v NullableGatewayMigrationSyncOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayMigrationSyncOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


