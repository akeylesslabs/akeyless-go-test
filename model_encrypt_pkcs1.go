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

// EncryptPKCS1 struct for EncryptPKCS1
type EncryptPKCS1 struct {
	// The display id of the key to use in the encryption process
	DisplayId *string `json:"display-id,omitempty"`
	// The item id of the key to use in the encryption process
	ItemId *int64 `json:"item-id,omitempty"`
	// The name of the key to use in the encryption process
	KeyName string `json:"key-name"`
	// Data to be encrypted
	Plaintext string `json:"plaintext"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewEncryptPKCS1 instantiates a new EncryptPKCS1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEncryptPKCS1(keyName string, plaintext string, ) *EncryptPKCS1 {
	this := EncryptPKCS1{}
	this.KeyName = keyName
	this.Plaintext = plaintext
	return &this
}

// NewEncryptPKCS1WithDefaults instantiates a new EncryptPKCS1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEncryptPKCS1WithDefaults() *EncryptPKCS1 {
	this := EncryptPKCS1{}
	return &this
}

// GetDisplayId returns the DisplayId field value if set, zero value otherwise.
func (o *EncryptPKCS1) GetDisplayId() string {
	if o == nil || o.DisplayId == nil {
		var ret string
		return ret
	}
	return *o.DisplayId
}

// GetDisplayIdOk returns a tuple with the DisplayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncryptPKCS1) GetDisplayIdOk() (*string, bool) {
	if o == nil || o.DisplayId == nil {
		return nil, false
	}
	return o.DisplayId, true
}

// HasDisplayId returns a boolean if a field has been set.
func (o *EncryptPKCS1) HasDisplayId() bool {
	if o != nil && o.DisplayId != nil {
		return true
	}

	return false
}

// SetDisplayId gets a reference to the given string and assigns it to the DisplayId field.
func (o *EncryptPKCS1) SetDisplayId(v string) {
	o.DisplayId = &v
}

// GetItemId returns the ItemId field value if set, zero value otherwise.
func (o *EncryptPKCS1) GetItemId() int64 {
	if o == nil || o.ItemId == nil {
		var ret int64
		return ret
	}
	return *o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncryptPKCS1) GetItemIdOk() (*int64, bool) {
	if o == nil || o.ItemId == nil {
		return nil, false
	}
	return o.ItemId, true
}

// HasItemId returns a boolean if a field has been set.
func (o *EncryptPKCS1) HasItemId() bool {
	if o != nil && o.ItemId != nil {
		return true
	}

	return false
}

// SetItemId gets a reference to the given int64 and assigns it to the ItemId field.
func (o *EncryptPKCS1) SetItemId(v int64) {
	o.ItemId = &v
}

// GetKeyName returns the KeyName field value
func (o *EncryptPKCS1) GetKeyName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.KeyName
}

// GetKeyNameOk returns a tuple with the KeyName field value
// and a boolean to check if the value has been set.
func (o *EncryptPKCS1) GetKeyNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.KeyName, true
}

// SetKeyName sets field value
func (o *EncryptPKCS1) SetKeyName(v string) {
	o.KeyName = v
}

// GetPlaintext returns the Plaintext field value
func (o *EncryptPKCS1) GetPlaintext() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Plaintext
}

// GetPlaintextOk returns a tuple with the Plaintext field value
// and a boolean to check if the value has been set.
func (o *EncryptPKCS1) GetPlaintextOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Plaintext, true
}

// SetPlaintext sets field value
func (o *EncryptPKCS1) SetPlaintext(v string) {
	o.Plaintext = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *EncryptPKCS1) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncryptPKCS1) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *EncryptPKCS1) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *EncryptPKCS1) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *EncryptPKCS1) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncryptPKCS1) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *EncryptPKCS1) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *EncryptPKCS1) SetUidToken(v string) {
	o.UidToken = &v
}

func (o EncryptPKCS1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DisplayId != nil {
		toSerialize["display-id"] = o.DisplayId
	}
	if o.ItemId != nil {
		toSerialize["item-id"] = o.ItemId
	}
	if true {
		toSerialize["key-name"] = o.KeyName
	}
	if true {
		toSerialize["plaintext"] = o.Plaintext
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableEncryptPKCS1 struct {
	value *EncryptPKCS1
	isSet bool
}

func (v NullableEncryptPKCS1) Get() *EncryptPKCS1 {
	return v.value
}

func (v *NullableEncryptPKCS1) Set(val *EncryptPKCS1) {
	v.value = val
	v.isSet = true
}

func (v NullableEncryptPKCS1) IsSet() bool {
	return v.isSet
}

func (v *NullableEncryptPKCS1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEncryptPKCS1(val *EncryptPKCS1) *NullableEncryptPKCS1 {
	return &NullableEncryptPKCS1{value: val, isSet: true}
}

func (v NullableEncryptPKCS1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEncryptPKCS1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


