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

// EmailPassAccessRules struct for EmailPassAccessRules
type EmailPassAccessRules struct {
	Alg *string `json:"alg,omitempty"`
	// The Email value
	Email *string `json:"email,omitempty"`
	// The password value
	HashPass *string `json:"hash_pass,omitempty"`
}

// NewEmailPassAccessRules instantiates a new EmailPassAccessRules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailPassAccessRules() *EmailPassAccessRules {
	this := EmailPassAccessRules{}
	return &this
}

// NewEmailPassAccessRulesWithDefaults instantiates a new EmailPassAccessRules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailPassAccessRulesWithDefaults() *EmailPassAccessRules {
	this := EmailPassAccessRules{}
	return &this
}

// GetAlg returns the Alg field value if set, zero value otherwise.
func (o *EmailPassAccessRules) GetAlg() string {
	if o == nil || o.Alg == nil {
		var ret string
		return ret
	}
	return *o.Alg
}

// GetAlgOk returns a tuple with the Alg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailPassAccessRules) GetAlgOk() (*string, bool) {
	if o == nil || o.Alg == nil {
		return nil, false
	}
	return o.Alg, true
}

// HasAlg returns a boolean if a field has been set.
func (o *EmailPassAccessRules) HasAlg() bool {
	if o != nil && o.Alg != nil {
		return true
	}

	return false
}

// SetAlg gets a reference to the given string and assigns it to the Alg field.
func (o *EmailPassAccessRules) SetAlg(v string) {
	o.Alg = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *EmailPassAccessRules) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailPassAccessRules) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *EmailPassAccessRules) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *EmailPassAccessRules) SetEmail(v string) {
	o.Email = &v
}

// GetHashPass returns the HashPass field value if set, zero value otherwise.
func (o *EmailPassAccessRules) GetHashPass() string {
	if o == nil || o.HashPass == nil {
		var ret string
		return ret
	}
	return *o.HashPass
}

// GetHashPassOk returns a tuple with the HashPass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailPassAccessRules) GetHashPassOk() (*string, bool) {
	if o == nil || o.HashPass == nil {
		return nil, false
	}
	return o.HashPass, true
}

// HasHashPass returns a boolean if a field has been set.
func (o *EmailPassAccessRules) HasHashPass() bool {
	if o != nil && o.HashPass != nil {
		return true
	}

	return false
}

// SetHashPass gets a reference to the given string and assigns it to the HashPass field.
func (o *EmailPassAccessRules) SetHashPass(v string) {
	o.HashPass = &v
}

func (o EmailPassAccessRules) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Alg != nil {
		toSerialize["alg"] = o.Alg
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.HashPass != nil {
		toSerialize["hash_pass"] = o.HashPass
	}
	return json.Marshal(toSerialize)
}

type NullableEmailPassAccessRules struct {
	value *EmailPassAccessRules
	isSet bool
}

func (v NullableEmailPassAccessRules) Get() *EmailPassAccessRules {
	return v.value
}

func (v *NullableEmailPassAccessRules) Set(val *EmailPassAccessRules) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailPassAccessRules) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailPassAccessRules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailPassAccessRules(val *EmailPassAccessRules) *NullableEmailPassAccessRules {
	return &NullableEmailPassAccessRules{value: val, isSet: true}
}

func (v NullableEmailPassAccessRules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailPassAccessRules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


