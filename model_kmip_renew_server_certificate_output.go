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

// KmipRenewServerCertificateOutput struct for KmipRenewServerCertificateOutput
type KmipRenewServerCertificateOutput struct {
	CaCert *string `json:"ca_cert,omitempty"`
}

// NewKmipRenewServerCertificateOutput instantiates a new KmipRenewServerCertificateOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKmipRenewServerCertificateOutput() *KmipRenewServerCertificateOutput {
	this := KmipRenewServerCertificateOutput{}
	return &this
}

// NewKmipRenewServerCertificateOutputWithDefaults instantiates a new KmipRenewServerCertificateOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKmipRenewServerCertificateOutputWithDefaults() *KmipRenewServerCertificateOutput {
	this := KmipRenewServerCertificateOutput{}
	return &this
}

// GetCaCert returns the CaCert field value if set, zero value otherwise.
func (o *KmipRenewServerCertificateOutput) GetCaCert() string {
	if o == nil || o.CaCert == nil {
		var ret string
		return ret
	}
	return *o.CaCert
}

// GetCaCertOk returns a tuple with the CaCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KmipRenewServerCertificateOutput) GetCaCertOk() (*string, bool) {
	if o == nil || o.CaCert == nil {
		return nil, false
	}
	return o.CaCert, true
}

// HasCaCert returns a boolean if a field has been set.
func (o *KmipRenewServerCertificateOutput) HasCaCert() bool {
	if o != nil && o.CaCert != nil {
		return true
	}

	return false
}

// SetCaCert gets a reference to the given string and assigns it to the CaCert field.
func (o *KmipRenewServerCertificateOutput) SetCaCert(v string) {
	o.CaCert = &v
}

func (o KmipRenewServerCertificateOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CaCert != nil {
		toSerialize["ca_cert"] = o.CaCert
	}
	return json.Marshal(toSerialize)
}

type NullableKmipRenewServerCertificateOutput struct {
	value *KmipRenewServerCertificateOutput
	isSet bool
}

func (v NullableKmipRenewServerCertificateOutput) Get() *KmipRenewServerCertificateOutput {
	return v.value
}

func (v *NullableKmipRenewServerCertificateOutput) Set(val *KmipRenewServerCertificateOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableKmipRenewServerCertificateOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableKmipRenewServerCertificateOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKmipRenewServerCertificateOutput(val *KmipRenewServerCertificateOutput) *NullableKmipRenewServerCertificateOutput {
	return &NullableKmipRenewServerCertificateOutput{value: val, isSet: true}
}

func (v NullableKmipRenewServerCertificateOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKmipRenewServerCertificateOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


