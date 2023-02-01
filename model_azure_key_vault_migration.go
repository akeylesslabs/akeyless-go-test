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

// AzureKeyVaultMigration struct for AzureKeyVaultMigration
type AzureKeyVaultMigration struct {
	General *MigrationGeneral `json:"general,omitempty"`
	Payload *AzurePayload `json:"payload,omitempty"`
}

// NewAzureKeyVaultMigration instantiates a new AzureKeyVaultMigration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureKeyVaultMigration() *AzureKeyVaultMigration {
	this := AzureKeyVaultMigration{}
	return &this
}

// NewAzureKeyVaultMigrationWithDefaults instantiates a new AzureKeyVaultMigration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureKeyVaultMigrationWithDefaults() *AzureKeyVaultMigration {
	this := AzureKeyVaultMigration{}
	return &this
}

// GetGeneral returns the General field value if set, zero value otherwise.
func (o *AzureKeyVaultMigration) GetGeneral() MigrationGeneral {
	if o == nil || o.General == nil {
		var ret MigrationGeneral
		return ret
	}
	return *o.General
}

// GetGeneralOk returns a tuple with the General field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultMigration) GetGeneralOk() (*MigrationGeneral, bool) {
	if o == nil || o.General == nil {
		return nil, false
	}
	return o.General, true
}

// HasGeneral returns a boolean if a field has been set.
func (o *AzureKeyVaultMigration) HasGeneral() bool {
	if o != nil && o.General != nil {
		return true
	}

	return false
}

// SetGeneral gets a reference to the given MigrationGeneral and assigns it to the General field.
func (o *AzureKeyVaultMigration) SetGeneral(v MigrationGeneral) {
	o.General = &v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *AzureKeyVaultMigration) GetPayload() AzurePayload {
	if o == nil || o.Payload == nil {
		var ret AzurePayload
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultMigration) GetPayloadOk() (*AzurePayload, bool) {
	if o == nil || o.Payload == nil {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *AzureKeyVaultMigration) HasPayload() bool {
	if o != nil && o.Payload != nil {
		return true
	}

	return false
}

// SetPayload gets a reference to the given AzurePayload and assigns it to the Payload field.
func (o *AzureKeyVaultMigration) SetPayload(v AzurePayload) {
	o.Payload = &v
}

func (o AzureKeyVaultMigration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.General != nil {
		toSerialize["general"] = o.General
	}
	if o.Payload != nil {
		toSerialize["payload"] = o.Payload
	}
	return json.Marshal(toSerialize)
}

type NullableAzureKeyVaultMigration struct {
	value *AzureKeyVaultMigration
	isSet bool
}

func (v NullableAzureKeyVaultMigration) Get() *AzureKeyVaultMigration {
	return v.value
}

func (v *NullableAzureKeyVaultMigration) Set(val *AzureKeyVaultMigration) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureKeyVaultMigration) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureKeyVaultMigration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureKeyVaultMigration(val *AzureKeyVaultMigration) *NullableAzureKeyVaultMigration {
	return &NullableAzureKeyVaultMigration{value: val, isSet: true}
}

func (v NullableAzureKeyVaultMigration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureKeyVaultMigration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


