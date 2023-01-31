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

// OnePasswordPayload struct for OnePasswordPayload
type OnePasswordPayload struct {
	Email *string `json:"email,omitempty"`
	Password *string `json:"password,omitempty"`
	SecretKey *string `json:"secret_key,omitempty"`
	Url *string `json:"url,omitempty"`
	Vaults *[]string `json:"vaults,omitempty"`
}

// NewOnePasswordPayload instantiates a new OnePasswordPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOnePasswordPayload() *OnePasswordPayload {
	this := OnePasswordPayload{}
	return &this
}

// NewOnePasswordPayloadWithDefaults instantiates a new OnePasswordPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOnePasswordPayloadWithDefaults() *OnePasswordPayload {
	this := OnePasswordPayload{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *OnePasswordPayload) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnePasswordPayload) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *OnePasswordPayload) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *OnePasswordPayload) SetEmail(v string) {
	o.Email = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *OnePasswordPayload) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnePasswordPayload) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *OnePasswordPayload) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *OnePasswordPayload) SetPassword(v string) {
	o.Password = &v
}

// GetSecretKey returns the SecretKey field value if set, zero value otherwise.
func (o *OnePasswordPayload) GetSecretKey() string {
	if o == nil || o.SecretKey == nil {
		var ret string
		return ret
	}
	return *o.SecretKey
}

// GetSecretKeyOk returns a tuple with the SecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnePasswordPayload) GetSecretKeyOk() (*string, bool) {
	if o == nil || o.SecretKey == nil {
		return nil, false
	}
	return o.SecretKey, true
}

// HasSecretKey returns a boolean if a field has been set.
func (o *OnePasswordPayload) HasSecretKey() bool {
	if o != nil && o.SecretKey != nil {
		return true
	}

	return false
}

// SetSecretKey gets a reference to the given string and assigns it to the SecretKey field.
func (o *OnePasswordPayload) SetSecretKey(v string) {
	o.SecretKey = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *OnePasswordPayload) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnePasswordPayload) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *OnePasswordPayload) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *OnePasswordPayload) SetUrl(v string) {
	o.Url = &v
}

// GetVaults returns the Vaults field value if set, zero value otherwise.
func (o *OnePasswordPayload) GetVaults() []string {
	if o == nil || o.Vaults == nil {
		var ret []string
		return ret
	}
	return *o.Vaults
}

// GetVaultsOk returns a tuple with the Vaults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnePasswordPayload) GetVaultsOk() (*[]string, bool) {
	if o == nil || o.Vaults == nil {
		return nil, false
	}
	return o.Vaults, true
}

// HasVaults returns a boolean if a field has been set.
func (o *OnePasswordPayload) HasVaults() bool {
	if o != nil && o.Vaults != nil {
		return true
	}

	return false
}

// SetVaults gets a reference to the given []string and assigns it to the Vaults field.
func (o *OnePasswordPayload) SetVaults(v []string) {
	o.Vaults = &v
}

func (o OnePasswordPayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.SecretKey != nil {
		toSerialize["secret_key"] = o.SecretKey
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.Vaults != nil {
		toSerialize["vaults"] = o.Vaults
	}
	return json.Marshal(toSerialize)
}

type NullableOnePasswordPayload struct {
	value *OnePasswordPayload
	isSet bool
}

func (v NullableOnePasswordPayload) Get() *OnePasswordPayload {
	return v.value
}

func (v *NullableOnePasswordPayload) Set(val *OnePasswordPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableOnePasswordPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableOnePasswordPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnePasswordPayload(val *OnePasswordPayload) *NullableOnePasswordPayload {
	return &NullableOnePasswordPayload{value: val, isSet: true}
}

func (v NullableOnePasswordPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnePasswordPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


