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

// VerifyGPG struct for VerifyGPG
type VerifyGPG struct {
	// The display id of the key to use in the encryption process
	DisplayId *string `json:"display-id,omitempty"`
	// The item id of the key to use in the encryption process
	ItemId *int64 `json:"item-id,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// The name of the key to use in the encryption process
	KeyName string `json:"key-name"`
	// Passphrase that was used to generate the key
	Passphrase *string `json:"passphrase,omitempty"`
	// The signature to be verified in base64 format
	Signature string `json:"signature"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewVerifyGPG instantiates a new VerifyGPG object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerifyGPG(keyName string, signature string, ) *VerifyGPG {
	this := VerifyGPG{}
	this.KeyName = keyName
	this.Signature = signature
	return &this
}

// NewVerifyGPGWithDefaults instantiates a new VerifyGPG object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerifyGPGWithDefaults() *VerifyGPG {
	this := VerifyGPG{}
	return &this
}

// GetDisplayId returns the DisplayId field value if set, zero value otherwise.
func (o *VerifyGPG) GetDisplayId() string {
	if o == nil || o.DisplayId == nil {
		var ret string
		return ret
	}
	return *o.DisplayId
}

// GetDisplayIdOk returns a tuple with the DisplayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyGPG) GetDisplayIdOk() (*string, bool) {
	if o == nil || o.DisplayId == nil {
		return nil, false
	}
	return o.DisplayId, true
}

// HasDisplayId returns a boolean if a field has been set.
func (o *VerifyGPG) HasDisplayId() bool {
	if o != nil && o.DisplayId != nil {
		return true
	}

	return false
}

// SetDisplayId gets a reference to the given string and assigns it to the DisplayId field.
func (o *VerifyGPG) SetDisplayId(v string) {
	o.DisplayId = &v
}

// GetItemId returns the ItemId field value if set, zero value otherwise.
func (o *VerifyGPG) GetItemId() int64 {
	if o == nil || o.ItemId == nil {
		var ret int64
		return ret
	}
	return *o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyGPG) GetItemIdOk() (*int64, bool) {
	if o == nil || o.ItemId == nil {
		return nil, false
	}
	return o.ItemId, true
}

// HasItemId returns a boolean if a field has been set.
func (o *VerifyGPG) HasItemId() bool {
	if o != nil && o.ItemId != nil {
		return true
	}

	return false
}

// SetItemId gets a reference to the given int64 and assigns it to the ItemId field.
func (o *VerifyGPG) SetItemId(v int64) {
	o.ItemId = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *VerifyGPG) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyGPG) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *VerifyGPG) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *VerifyGPG) SetJson(v bool) {
	o.Json = &v
}

// GetKeyName returns the KeyName field value
func (o *VerifyGPG) GetKeyName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.KeyName
}

// GetKeyNameOk returns a tuple with the KeyName field value
// and a boolean to check if the value has been set.
func (o *VerifyGPG) GetKeyNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.KeyName, true
}

// SetKeyName sets field value
func (o *VerifyGPG) SetKeyName(v string) {
	o.KeyName = v
}

// GetPassphrase returns the Passphrase field value if set, zero value otherwise.
func (o *VerifyGPG) GetPassphrase() string {
	if o == nil || o.Passphrase == nil {
		var ret string
		return ret
	}
	return *o.Passphrase
}

// GetPassphraseOk returns a tuple with the Passphrase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyGPG) GetPassphraseOk() (*string, bool) {
	if o == nil || o.Passphrase == nil {
		return nil, false
	}
	return o.Passphrase, true
}

// HasPassphrase returns a boolean if a field has been set.
func (o *VerifyGPG) HasPassphrase() bool {
	if o != nil && o.Passphrase != nil {
		return true
	}

	return false
}

// SetPassphrase gets a reference to the given string and assigns it to the Passphrase field.
func (o *VerifyGPG) SetPassphrase(v string) {
	o.Passphrase = &v
}

// GetSignature returns the Signature field value
func (o *VerifyGPG) GetSignature() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value
// and a boolean to check if the value has been set.
func (o *VerifyGPG) GetSignatureOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Signature, true
}

// SetSignature sets field value
func (o *VerifyGPG) SetSignature(v string) {
	o.Signature = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *VerifyGPG) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyGPG) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *VerifyGPG) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *VerifyGPG) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *VerifyGPG) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyGPG) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *VerifyGPG) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *VerifyGPG) SetUidToken(v string) {
	o.UidToken = &v
}

func (o VerifyGPG) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DisplayId != nil {
		toSerialize["display-id"] = o.DisplayId
	}
	if o.ItemId != nil {
		toSerialize["item-id"] = o.ItemId
	}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if true {
		toSerialize["key-name"] = o.KeyName
	}
	if o.Passphrase != nil {
		toSerialize["passphrase"] = o.Passphrase
	}
	if true {
		toSerialize["signature"] = o.Signature
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableVerifyGPG struct {
	value *VerifyGPG
	isSet bool
}

func (v NullableVerifyGPG) Get() *VerifyGPG {
	return v.value
}

func (v *NullableVerifyGPG) Set(val *VerifyGPG) {
	v.value = val
	v.isSet = true
}

func (v NullableVerifyGPG) IsSet() bool {
	return v.isSet
}

func (v *NullableVerifyGPG) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerifyGPG(val *VerifyGPG) *NullableVerifyGPG {
	return &NullableVerifyGPG{value: val, isSet: true}
}

func (v NullableVerifyGPG) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerifyGPG) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

