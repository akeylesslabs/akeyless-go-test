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

// GatewayMigratePersonalItems gatewayMigratePersonalItems is a command that migrate personal items from external vault
type GatewayMigratePersonalItems struct {
	// 1Password user email to connect to the API
	Var1passwordEmail *string `json:"1password-email,omitempty"`
	// 1Password user password to connect to the API
	Var1passwordPassword *string `json:"1password-password,omitempty"`
	// 1Password user secret key to connect to the API
	Var1passwordSecretKey *string `json:"1password-secret-key,omitempty"`
	// 1Password api container url
	Var1passwordUrl *string `json:"1password-url,omitempty"`
	// 1Password list of vault to get the items from
	Var1passwordVaults *[]string `json:"1password-vaults,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// The name of a key that used to encrypt the secret value
	ProtectionKey *string `json:"protection-key,omitempty"`
	// Target location in your Akeyless personal folder for migrated secrets
	TargetLocation *string `json:"target-location,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// Migration type for now only 1password.
	Type *string `json:"type,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewGatewayMigratePersonalItems instantiates a new GatewayMigratePersonalItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayMigratePersonalItems() *GatewayMigratePersonalItems {
	this := GatewayMigratePersonalItems{}
	var type_ string = "1password"
	this.Type = &type_
	return &this
}

// NewGatewayMigratePersonalItemsWithDefaults instantiates a new GatewayMigratePersonalItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayMigratePersonalItemsWithDefaults() *GatewayMigratePersonalItems {
	this := GatewayMigratePersonalItems{}
	var type_ string = "1password"
	this.Type = &type_
	return &this
}

// GetVar1passwordEmail returns the Var1passwordEmail field value if set, zero value otherwise.
func (o *GatewayMigratePersonalItems) GetVar1passwordEmail() string {
	if o == nil || o.Var1passwordEmail == nil {
		var ret string
		return ret
	}
	return *o.Var1passwordEmail
}

// GetVar1passwordEmailOk returns a tuple with the Var1passwordEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayMigratePersonalItems) GetVar1passwordEmailOk() (*string, bool) {
	if o == nil || o.Var1passwordEmail == nil {
		return nil, false
	}
	return o.Var1passwordEmail, true
}

// HasVar1passwordEmail returns a boolean if a field has been set.
func (o *GatewayMigratePersonalItems) HasVar1passwordEmail() bool {
	if o != nil && o.Var1passwordEmail != nil {
		return true
	}

	return false
}

// SetVar1passwordEmail gets a reference to the given string and assigns it to the Var1passwordEmail field.
func (o *GatewayMigratePersonalItems) SetVar1passwordEmail(v string) {
	o.Var1passwordEmail = &v
}

// GetVar1passwordPassword returns the Var1passwordPassword field value if set, zero value otherwise.
func (o *GatewayMigratePersonalItems) GetVar1passwordPassword() string {
	if o == nil || o.Var1passwordPassword == nil {
		var ret string
		return ret
	}
	return *o.Var1passwordPassword
}

// GetVar1passwordPasswordOk returns a tuple with the Var1passwordPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayMigratePersonalItems) GetVar1passwordPasswordOk() (*string, bool) {
	if o == nil || o.Var1passwordPassword == nil {
		return nil, false
	}
	return o.Var1passwordPassword, true
}

// HasVar1passwordPassword returns a boolean if a field has been set.
func (o *GatewayMigratePersonalItems) HasVar1passwordPassword() bool {
	if o != nil && o.Var1passwordPassword != nil {
		return true
	}

	return false
}

// SetVar1passwordPassword gets a reference to the given string and assigns it to the Var1passwordPassword field.
func (o *GatewayMigratePersonalItems) SetVar1passwordPassword(v string) {
	o.Var1passwordPassword = &v
}

// GetVar1passwordSecretKey returns the Var1passwordSecretKey field value if set, zero value otherwise.
func (o *GatewayMigratePersonalItems) GetVar1passwordSecretKey() string {
	if o == nil || o.Var1passwordSecretKey == nil {
		var ret string
		return ret
	}
	return *o.Var1passwordSecretKey
}

// GetVar1passwordSecretKeyOk returns a tuple with the Var1passwordSecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayMigratePersonalItems) GetVar1passwordSecretKeyOk() (*string, bool) {
	if o == nil || o.Var1passwordSecretKey == nil {
		return nil, false
	}
	return o.Var1passwordSecretKey, true
}

// HasVar1passwordSecretKey returns a boolean if a field has been set.
func (o *GatewayMigratePersonalItems) HasVar1passwordSecretKey() bool {
	if o != nil && o.Var1passwordSecretKey != nil {
		return true
	}

	return false
}

// SetVar1passwordSecretKey gets a reference to the given string and assigns it to the Var1passwordSecretKey field.
func (o *GatewayMigratePersonalItems) SetVar1passwordSecretKey(v string) {
	o.Var1passwordSecretKey = &v
}

// GetVar1passwordUrl returns the Var1passwordUrl field value if set, zero value otherwise.
func (o *GatewayMigratePersonalItems) GetVar1passwordUrl() string {
	if o == nil || o.Var1passwordUrl == nil {
		var ret string
		return ret
	}
	return *o.Var1passwordUrl
}

// GetVar1passwordUrlOk returns a tuple with the Var1passwordUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayMigratePersonalItems) GetVar1passwordUrlOk() (*string, bool) {
	if o == nil || o.Var1passwordUrl == nil {
		return nil, false
	}
	return o.Var1passwordUrl, true
}

// HasVar1passwordUrl returns a boolean if a field has been set.
func (o *GatewayMigratePersonalItems) HasVar1passwordUrl() bool {
	if o != nil && o.Var1passwordUrl != nil {
		return true
	}

	return false
}

// SetVar1passwordUrl gets a reference to the given string and assigns it to the Var1passwordUrl field.
func (o *GatewayMigratePersonalItems) SetVar1passwordUrl(v string) {
	o.Var1passwordUrl = &v
}

// GetVar1passwordVaults returns the Var1passwordVaults field value if set, zero value otherwise.
func (o *GatewayMigratePersonalItems) GetVar1passwordVaults() []string {
	if o == nil || o.Var1passwordVaults == nil {
		var ret []string
		return ret
	}
	return *o.Var1passwordVaults
}

// GetVar1passwordVaultsOk returns a tuple with the Var1passwordVaults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayMigratePersonalItems) GetVar1passwordVaultsOk() (*[]string, bool) {
	if o == nil || o.Var1passwordVaults == nil {
		return nil, false
	}
	return o.Var1passwordVaults, true
}

// HasVar1passwordVaults returns a boolean if a field has been set.
func (o *GatewayMigratePersonalItems) HasVar1passwordVaults() bool {
	if o != nil && o.Var1passwordVaults != nil {
		return true
	}

	return false
}

// SetVar1passwordVaults gets a reference to the given []string and assigns it to the Var1passwordVaults field.
func (o *GatewayMigratePersonalItems) SetVar1passwordVaults(v []string) {
	o.Var1passwordVaults = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *GatewayMigratePersonalItems) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayMigratePersonalItems) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *GatewayMigratePersonalItems) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *GatewayMigratePersonalItems) SetJson(v bool) {
	o.Json = &v
}

// GetProtectionKey returns the ProtectionKey field value if set, zero value otherwise.
func (o *GatewayMigratePersonalItems) GetProtectionKey() string {
	if o == nil || o.ProtectionKey == nil {
		var ret string
		return ret
	}
	return *o.ProtectionKey
}

// GetProtectionKeyOk returns a tuple with the ProtectionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayMigratePersonalItems) GetProtectionKeyOk() (*string, bool) {
	if o == nil || o.ProtectionKey == nil {
		return nil, false
	}
	return o.ProtectionKey, true
}

// HasProtectionKey returns a boolean if a field has been set.
func (o *GatewayMigratePersonalItems) HasProtectionKey() bool {
	if o != nil && o.ProtectionKey != nil {
		return true
	}

	return false
}

// SetProtectionKey gets a reference to the given string and assigns it to the ProtectionKey field.
func (o *GatewayMigratePersonalItems) SetProtectionKey(v string) {
	o.ProtectionKey = &v
}

// GetTargetLocation returns the TargetLocation field value if set, zero value otherwise.
func (o *GatewayMigratePersonalItems) GetTargetLocation() string {
	if o == nil || o.TargetLocation == nil {
		var ret string
		return ret
	}
	return *o.TargetLocation
}

// GetTargetLocationOk returns a tuple with the TargetLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayMigratePersonalItems) GetTargetLocationOk() (*string, bool) {
	if o == nil || o.TargetLocation == nil {
		return nil, false
	}
	return o.TargetLocation, true
}

// HasTargetLocation returns a boolean if a field has been set.
func (o *GatewayMigratePersonalItems) HasTargetLocation() bool {
	if o != nil && o.TargetLocation != nil {
		return true
	}

	return false
}

// SetTargetLocation gets a reference to the given string and assigns it to the TargetLocation field.
func (o *GatewayMigratePersonalItems) SetTargetLocation(v string) {
	o.TargetLocation = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GatewayMigratePersonalItems) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayMigratePersonalItems) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GatewayMigratePersonalItems) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GatewayMigratePersonalItems) SetToken(v string) {
	o.Token = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GatewayMigratePersonalItems) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayMigratePersonalItems) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GatewayMigratePersonalItems) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GatewayMigratePersonalItems) SetType(v string) {
	o.Type = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *GatewayMigratePersonalItems) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayMigratePersonalItems) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *GatewayMigratePersonalItems) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *GatewayMigratePersonalItems) SetUidToken(v string) {
	o.UidToken = &v
}

func (o GatewayMigratePersonalItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Var1passwordEmail != nil {
		toSerialize["1password-email"] = o.Var1passwordEmail
	}
	if o.Var1passwordPassword != nil {
		toSerialize["1password-password"] = o.Var1passwordPassword
	}
	if o.Var1passwordSecretKey != nil {
		toSerialize["1password-secret-key"] = o.Var1passwordSecretKey
	}
	if o.Var1passwordUrl != nil {
		toSerialize["1password-url"] = o.Var1passwordUrl
	}
	if o.Var1passwordVaults != nil {
		toSerialize["1password-vaults"] = o.Var1passwordVaults
	}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if o.ProtectionKey != nil {
		toSerialize["protection-key"] = o.ProtectionKey
	}
	if o.TargetLocation != nil {
		toSerialize["target-location"] = o.TargetLocation
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableGatewayMigratePersonalItems struct {
	value *GatewayMigratePersonalItems
	isSet bool
}

func (v NullableGatewayMigratePersonalItems) Get() *GatewayMigratePersonalItems {
	return v.value
}

func (v *NullableGatewayMigratePersonalItems) Set(val *GatewayMigratePersonalItems) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayMigratePersonalItems) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayMigratePersonalItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayMigratePersonalItems(val *GatewayMigratePersonalItems) *NullableGatewayMigratePersonalItems {
	return &NullableGatewayMigratePersonalItems{value: val, isSet: true}
}

func (v NullableGatewayMigratePersonalItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayMigratePersonalItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

