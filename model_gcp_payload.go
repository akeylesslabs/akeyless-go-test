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

// GCPPayload struct for GCPPayload
type GCPPayload struct {
	GcpCredentialsJson *string `json:"gcp_credentials_json,omitempty"`
}

// NewGCPPayload instantiates a new GCPPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGCPPayload() *GCPPayload {
	this := GCPPayload{}
	return &this
}

// NewGCPPayloadWithDefaults instantiates a new GCPPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGCPPayloadWithDefaults() *GCPPayload {
	this := GCPPayload{}
	return &this
}

// GetGcpCredentialsJson returns the GcpCredentialsJson field value if set, zero value otherwise.
func (o *GCPPayload) GetGcpCredentialsJson() string {
	if o == nil || o.GcpCredentialsJson == nil {
		var ret string
		return ret
	}
	return *o.GcpCredentialsJson
}

// GetGcpCredentialsJsonOk returns a tuple with the GcpCredentialsJson field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPPayload) GetGcpCredentialsJsonOk() (*string, bool) {
	if o == nil || o.GcpCredentialsJson == nil {
		return nil, false
	}
	return o.GcpCredentialsJson, true
}

// HasGcpCredentialsJson returns a boolean if a field has been set.
func (o *GCPPayload) HasGcpCredentialsJson() bool {
	if o != nil && o.GcpCredentialsJson != nil {
		return true
	}

	return false
}

// SetGcpCredentialsJson gets a reference to the given string and assigns it to the GcpCredentialsJson field.
func (o *GCPPayload) SetGcpCredentialsJson(v string) {
	o.GcpCredentialsJson = &v
}

func (o GCPPayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GcpCredentialsJson != nil {
		toSerialize["gcp_credentials_json"] = o.GcpCredentialsJson
	}
	return json.Marshal(toSerialize)
}

type NullableGCPPayload struct {
	value *GCPPayload
	isSet bool
}

func (v NullableGCPPayload) Get() *GCPPayload {
	return v.value
}

func (v *NullableGCPPayload) Set(val *GCPPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableGCPPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableGCPPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGCPPayload(val *GCPPayload) *NullableGCPPayload {
	return &NullableGCPPayload{value: val, isSet: true}
}

func (v NullableGCPPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGCPPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


