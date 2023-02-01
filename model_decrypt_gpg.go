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

// DecryptGPG struct for DecryptGPG
type DecryptGPG struct {
	// Ciphertext to be decrypted in base64 encoded format
	Ciphertext string `json:"ciphertext"`
	// The display id of the key to use in the decryption process
	DisplayId *string `json:"display-id,omitempty"`
	// The item id of the key to use in the decryption process
	ItemId *int64 `json:"item-id,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// The name of the key to use in the decryption process
	KeyName string `json:"key-name"`
	// If specified, the output will be formatted accordingly. options: [base64]
	OutputFormat *string `json:"output-format,omitempty"`
	// Passphrase that was used to generate the key
	Passphrase *string `json:"passphrase,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewDecryptGPG instantiates a new DecryptGPG object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDecryptGPG(ciphertext string, keyName string, ) *DecryptGPG {
	this := DecryptGPG{}
	this.Ciphertext = ciphertext
	this.KeyName = keyName
	return &this
}

// NewDecryptGPGWithDefaults instantiates a new DecryptGPG object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDecryptGPGWithDefaults() *DecryptGPG {
	this := DecryptGPG{}
	return &this
}

// GetCiphertext returns the Ciphertext field value
func (o *DecryptGPG) GetCiphertext() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Ciphertext
}

// GetCiphertextOk returns a tuple with the Ciphertext field value
// and a boolean to check if the value has been set.
func (o *DecryptGPG) GetCiphertextOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Ciphertext, true
}

// SetCiphertext sets field value
func (o *DecryptGPG) SetCiphertext(v string) {
	o.Ciphertext = v
}

// GetDisplayId returns the DisplayId field value if set, zero value otherwise.
func (o *DecryptGPG) GetDisplayId() string {
	if o == nil || o.DisplayId == nil {
		var ret string
		return ret
	}
	return *o.DisplayId
}

// GetDisplayIdOk returns a tuple with the DisplayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecryptGPG) GetDisplayIdOk() (*string, bool) {
	if o == nil || o.DisplayId == nil {
		return nil, false
	}
	return o.DisplayId, true
}

// HasDisplayId returns a boolean if a field has been set.
func (o *DecryptGPG) HasDisplayId() bool {
	if o != nil && o.DisplayId != nil {
		return true
	}

	return false
}

// SetDisplayId gets a reference to the given string and assigns it to the DisplayId field.
func (o *DecryptGPG) SetDisplayId(v string) {
	o.DisplayId = &v
}

// GetItemId returns the ItemId field value if set, zero value otherwise.
func (o *DecryptGPG) GetItemId() int64 {
	if o == nil || o.ItemId == nil {
		var ret int64
		return ret
	}
	return *o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecryptGPG) GetItemIdOk() (*int64, bool) {
	if o == nil || o.ItemId == nil {
		return nil, false
	}
	return o.ItemId, true
}

// HasItemId returns a boolean if a field has been set.
func (o *DecryptGPG) HasItemId() bool {
	if o != nil && o.ItemId != nil {
		return true
	}

	return false
}

// SetItemId gets a reference to the given int64 and assigns it to the ItemId field.
func (o *DecryptGPG) SetItemId(v int64) {
	o.ItemId = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *DecryptGPG) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecryptGPG) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *DecryptGPG) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *DecryptGPG) SetJson(v bool) {
	o.Json = &v
}

// GetKeyName returns the KeyName field value
func (o *DecryptGPG) GetKeyName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.KeyName
}

// GetKeyNameOk returns a tuple with the KeyName field value
// and a boolean to check if the value has been set.
func (o *DecryptGPG) GetKeyNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.KeyName, true
}

// SetKeyName sets field value
func (o *DecryptGPG) SetKeyName(v string) {
	o.KeyName = v
}

// GetOutputFormat returns the OutputFormat field value if set, zero value otherwise.
func (o *DecryptGPG) GetOutputFormat() string {
	if o == nil || o.OutputFormat == nil {
		var ret string
		return ret
	}
	return *o.OutputFormat
}

// GetOutputFormatOk returns a tuple with the OutputFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecryptGPG) GetOutputFormatOk() (*string, bool) {
	if o == nil || o.OutputFormat == nil {
		return nil, false
	}
	return o.OutputFormat, true
}

// HasOutputFormat returns a boolean if a field has been set.
func (o *DecryptGPG) HasOutputFormat() bool {
	if o != nil && o.OutputFormat != nil {
		return true
	}

	return false
}

// SetOutputFormat gets a reference to the given string and assigns it to the OutputFormat field.
func (o *DecryptGPG) SetOutputFormat(v string) {
	o.OutputFormat = &v
}

// GetPassphrase returns the Passphrase field value if set, zero value otherwise.
func (o *DecryptGPG) GetPassphrase() string {
	if o == nil || o.Passphrase == nil {
		var ret string
		return ret
	}
	return *o.Passphrase
}

// GetPassphraseOk returns a tuple with the Passphrase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecryptGPG) GetPassphraseOk() (*string, bool) {
	if o == nil || o.Passphrase == nil {
		return nil, false
	}
	return o.Passphrase, true
}

// HasPassphrase returns a boolean if a field has been set.
func (o *DecryptGPG) HasPassphrase() bool {
	if o != nil && o.Passphrase != nil {
		return true
	}

	return false
}

// SetPassphrase gets a reference to the given string and assigns it to the Passphrase field.
func (o *DecryptGPG) SetPassphrase(v string) {
	o.Passphrase = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *DecryptGPG) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecryptGPG) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *DecryptGPG) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *DecryptGPG) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *DecryptGPG) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecryptGPG) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *DecryptGPG) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *DecryptGPG) SetUidToken(v string) {
	o.UidToken = &v
}

func (o DecryptGPG) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ciphertext"] = o.Ciphertext
	}
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
	if o.OutputFormat != nil {
		toSerialize["output-format"] = o.OutputFormat
	}
	if o.Passphrase != nil {
		toSerialize["passphrase"] = o.Passphrase
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableDecryptGPG struct {
	value *DecryptGPG
	isSet bool
}

func (v NullableDecryptGPG) Get() *DecryptGPG {
	return v.value
}

func (v *NullableDecryptGPG) Set(val *DecryptGPG) {
	v.value = val
	v.isSet = true
}

func (v NullableDecryptGPG) IsSet() bool {
	return v.isSet
}

func (v *NullableDecryptGPG) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDecryptGPG(val *DecryptGPG) *NullableDecryptGPG {
	return &NullableDecryptGPG{value: val, isSet: true}
}

func (v NullableDecryptGPG) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDecryptGPG) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


