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

// KmipRenewClientCertificateOutput struct for KmipRenewClientCertificateOutput
type KmipRenewClientCertificateOutput struct {
	Certificate *string `json:"certificate,omitempty"`
	Id *string `json:"id,omitempty"`
	Key *string `json:"key,omitempty"`
}

// NewKmipRenewClientCertificateOutput instantiates a new KmipRenewClientCertificateOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKmipRenewClientCertificateOutput() *KmipRenewClientCertificateOutput {
	this := KmipRenewClientCertificateOutput{}
	return &this
}

// NewKmipRenewClientCertificateOutputWithDefaults instantiates a new KmipRenewClientCertificateOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKmipRenewClientCertificateOutputWithDefaults() *KmipRenewClientCertificateOutput {
	this := KmipRenewClientCertificateOutput{}
	return &this
}

// GetCertificate returns the Certificate field value if set, zero value otherwise.
func (o *KmipRenewClientCertificateOutput) GetCertificate() string {
	if o == nil || o.Certificate == nil {
		var ret string
		return ret
	}
	return *o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KmipRenewClientCertificateOutput) GetCertificateOk() (*string, bool) {
	if o == nil || o.Certificate == nil {
		return nil, false
	}
	return o.Certificate, true
}

// HasCertificate returns a boolean if a field has been set.
func (o *KmipRenewClientCertificateOutput) HasCertificate() bool {
	if o != nil && o.Certificate != nil {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given string and assigns it to the Certificate field.
func (o *KmipRenewClientCertificateOutput) SetCertificate(v string) {
	o.Certificate = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *KmipRenewClientCertificateOutput) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KmipRenewClientCertificateOutput) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *KmipRenewClientCertificateOutput) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *KmipRenewClientCertificateOutput) SetId(v string) {
	o.Id = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *KmipRenewClientCertificateOutput) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KmipRenewClientCertificateOutput) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *KmipRenewClientCertificateOutput) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *KmipRenewClientCertificateOutput) SetKey(v string) {
	o.Key = &v
}

func (o KmipRenewClientCertificateOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Certificate != nil {
		toSerialize["certificate"] = o.Certificate
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	return json.Marshal(toSerialize)
}

type NullableKmipRenewClientCertificateOutput struct {
	value *KmipRenewClientCertificateOutput
	isSet bool
}

func (v NullableKmipRenewClientCertificateOutput) Get() *KmipRenewClientCertificateOutput {
	return v.value
}

func (v *NullableKmipRenewClientCertificateOutput) Set(val *KmipRenewClientCertificateOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableKmipRenewClientCertificateOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableKmipRenewClientCertificateOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKmipRenewClientCertificateOutput(val *KmipRenewClientCertificateOutput) *NullableKmipRenewClientCertificateOutput {
	return &NullableKmipRenewClientCertificateOutput{value: val, isSet: true}
}

func (v NullableKmipRenewClientCertificateOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKmipRenewClientCertificateOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


