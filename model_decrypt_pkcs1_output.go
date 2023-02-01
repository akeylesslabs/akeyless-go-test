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

// DecryptPKCS1Output struct for DecryptPKCS1Output
type DecryptPKCS1Output struct {
	Plaintext *string `json:"plaintext,omitempty"`
}

// NewDecryptPKCS1Output instantiates a new DecryptPKCS1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDecryptPKCS1Output() *DecryptPKCS1Output {
	this := DecryptPKCS1Output{}
	return &this
}

// NewDecryptPKCS1OutputWithDefaults instantiates a new DecryptPKCS1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDecryptPKCS1OutputWithDefaults() *DecryptPKCS1Output {
	this := DecryptPKCS1Output{}
	return &this
}

// GetPlaintext returns the Plaintext field value if set, zero value otherwise.
func (o *DecryptPKCS1Output) GetPlaintext() string {
	if o == nil || o.Plaintext == nil {
		var ret string
		return ret
	}
	return *o.Plaintext
}

// GetPlaintextOk returns a tuple with the Plaintext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecryptPKCS1Output) GetPlaintextOk() (*string, bool) {
	if o == nil || o.Plaintext == nil {
		return nil, false
	}
	return o.Plaintext, true
}

// HasPlaintext returns a boolean if a field has been set.
func (o *DecryptPKCS1Output) HasPlaintext() bool {
	if o != nil && o.Plaintext != nil {
		return true
	}

	return false
}

// SetPlaintext gets a reference to the given string and assigns it to the Plaintext field.
func (o *DecryptPKCS1Output) SetPlaintext(v string) {
	o.Plaintext = &v
}

func (o DecryptPKCS1Output) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Plaintext != nil {
		toSerialize["plaintext"] = o.Plaintext
	}
	return json.Marshal(toSerialize)
}

type NullableDecryptPKCS1Output struct {
	value *DecryptPKCS1Output
	isSet bool
}

func (v NullableDecryptPKCS1Output) Get() *DecryptPKCS1Output {
	return v.value
}

func (v *NullableDecryptPKCS1Output) Set(val *DecryptPKCS1Output) {
	v.value = val
	v.isSet = true
}

func (v NullableDecryptPKCS1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableDecryptPKCS1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDecryptPKCS1Output(val *DecryptPKCS1Output) *NullableDecryptPKCS1Output {
	return &NullableDecryptPKCS1Output{value: val, isSet: true}
}

func (v NullableDecryptPKCS1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDecryptPKCS1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


