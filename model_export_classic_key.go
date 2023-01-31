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

// ExportClassicKey ExportClassicKey is a command that returns the classic key material
type ExportClassicKey struct {
	// Ignore Cache Retrieve the Secret value without checking the Gateway's cache. This flag is only relevant when using the RestAPI
	IgnoreCache *string `json:"ignore-cache,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// ClassicKey name
	Name string `json:"name"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	// Classic key version
	Version *int32 `json:"version,omitempty"`
}

// NewExportClassicKey instantiates a new ExportClassicKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExportClassicKey(name string, ) *ExportClassicKey {
	this := ExportClassicKey{}
	var ignoreCache string = "false"
	this.IgnoreCache = &ignoreCache
	this.Name = name
	return &this
}

// NewExportClassicKeyWithDefaults instantiates a new ExportClassicKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExportClassicKeyWithDefaults() *ExportClassicKey {
	this := ExportClassicKey{}
	var ignoreCache string = "false"
	this.IgnoreCache = &ignoreCache
	return &this
}

// GetIgnoreCache returns the IgnoreCache field value if set, zero value otherwise.
func (o *ExportClassicKey) GetIgnoreCache() string {
	if o == nil || o.IgnoreCache == nil {
		var ret string
		return ret
	}
	return *o.IgnoreCache
}

// GetIgnoreCacheOk returns a tuple with the IgnoreCache field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportClassicKey) GetIgnoreCacheOk() (*string, bool) {
	if o == nil || o.IgnoreCache == nil {
		return nil, false
	}
	return o.IgnoreCache, true
}

// HasIgnoreCache returns a boolean if a field has been set.
func (o *ExportClassicKey) HasIgnoreCache() bool {
	if o != nil && o.IgnoreCache != nil {
		return true
	}

	return false
}

// SetIgnoreCache gets a reference to the given string and assigns it to the IgnoreCache field.
func (o *ExportClassicKey) SetIgnoreCache(v string) {
	o.IgnoreCache = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *ExportClassicKey) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportClassicKey) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *ExportClassicKey) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *ExportClassicKey) SetJson(v bool) {
	o.Json = &v
}

// GetName returns the Name field value
func (o *ExportClassicKey) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ExportClassicKey) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ExportClassicKey) SetName(v string) {
	o.Name = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *ExportClassicKey) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportClassicKey) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *ExportClassicKey) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *ExportClassicKey) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *ExportClassicKey) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportClassicKey) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *ExportClassicKey) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *ExportClassicKey) SetUidToken(v string) {
	o.UidToken = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ExportClassicKey) GetVersion() int32 {
	if o == nil || o.Version == nil {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportClassicKey) GetVersionOk() (*int32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ExportClassicKey) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *ExportClassicKey) SetVersion(v int32) {
	o.Version = &v
}

func (o ExportClassicKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IgnoreCache != nil {
		toSerialize["ignore-cache"] = o.IgnoreCache
	}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableExportClassicKey struct {
	value *ExportClassicKey
	isSet bool
}

func (v NullableExportClassicKey) Get() *ExportClassicKey {
	return v.value
}

func (v *NullableExportClassicKey) Set(val *ExportClassicKey) {
	v.value = val
	v.isSet = true
}

func (v NullableExportClassicKey) IsSet() bool {
	return v.isSet
}

func (v *NullableExportClassicKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExportClassicKey(val *ExportClassicKey) *NullableExportClassicKey {
	return &NullableExportClassicKey{value: val, isSet: true}
}

func (v NullableExportClassicKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExportClassicKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

