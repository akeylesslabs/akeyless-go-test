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

// GetCertificateValueOutput struct for GetCertificateValueOutput
type GetCertificateValueOutput struct {
	CertificatePem *string `json:"certificate_pem,omitempty"`
	PrivateKeyPem *string `json:"private_key_pem,omitempty"`
}

// NewGetCertificateValueOutput instantiates a new GetCertificateValueOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCertificateValueOutput() *GetCertificateValueOutput {
	this := GetCertificateValueOutput{}
	return &this
}

// NewGetCertificateValueOutputWithDefaults instantiates a new GetCertificateValueOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCertificateValueOutputWithDefaults() *GetCertificateValueOutput {
	this := GetCertificateValueOutput{}
	return &this
}

// GetCertificatePem returns the CertificatePem field value if set, zero value otherwise.
func (o *GetCertificateValueOutput) GetCertificatePem() string {
	if o == nil || o.CertificatePem == nil {
		var ret string
		return ret
	}
	return *o.CertificatePem
}

// GetCertificatePemOk returns a tuple with the CertificatePem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCertificateValueOutput) GetCertificatePemOk() (*string, bool) {
	if o == nil || o.CertificatePem == nil {
		return nil, false
	}
	return o.CertificatePem, true
}

// HasCertificatePem returns a boolean if a field has been set.
func (o *GetCertificateValueOutput) HasCertificatePem() bool {
	if o != nil && o.CertificatePem != nil {
		return true
	}

	return false
}

// SetCertificatePem gets a reference to the given string and assigns it to the CertificatePem field.
func (o *GetCertificateValueOutput) SetCertificatePem(v string) {
	o.CertificatePem = &v
}

// GetPrivateKeyPem returns the PrivateKeyPem field value if set, zero value otherwise.
func (o *GetCertificateValueOutput) GetPrivateKeyPem() string {
	if o == nil || o.PrivateKeyPem == nil {
		var ret string
		return ret
	}
	return *o.PrivateKeyPem
}

// GetPrivateKeyPemOk returns a tuple with the PrivateKeyPem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCertificateValueOutput) GetPrivateKeyPemOk() (*string, bool) {
	if o == nil || o.PrivateKeyPem == nil {
		return nil, false
	}
	return o.PrivateKeyPem, true
}

// HasPrivateKeyPem returns a boolean if a field has been set.
func (o *GetCertificateValueOutput) HasPrivateKeyPem() bool {
	if o != nil && o.PrivateKeyPem != nil {
		return true
	}

	return false
}

// SetPrivateKeyPem gets a reference to the given string and assigns it to the PrivateKeyPem field.
func (o *GetCertificateValueOutput) SetPrivateKeyPem(v string) {
	o.PrivateKeyPem = &v
}

func (o GetCertificateValueOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CertificatePem != nil {
		toSerialize["certificate_pem"] = o.CertificatePem
	}
	if o.PrivateKeyPem != nil {
		toSerialize["private_key_pem"] = o.PrivateKeyPem
	}
	return json.Marshal(toSerialize)
}

type NullableGetCertificateValueOutput struct {
	value *GetCertificateValueOutput
	isSet bool
}

func (v NullableGetCertificateValueOutput) Get() *GetCertificateValueOutput {
	return v.value
}

func (v *NullableGetCertificateValueOutput) Set(val *GetCertificateValueOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCertificateValueOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCertificateValueOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCertificateValueOutput(val *GetCertificateValueOutput) *NullableGetCertificateValueOutput {
	return &NullableGetCertificateValueOutput{value: val, isSet: true}
}

func (v NullableGetCertificateValueOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCertificateValueOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


