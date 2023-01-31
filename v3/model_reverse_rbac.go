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

// ReverseRBAC reverseRBAC is a command that shows which auth methods have access to a particular object.
type ReverseRBAC struct {
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Path to an object
	Path string `json:"path"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// Type of object (item, am, role)
	Type string `json:"type"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewReverseRBAC instantiates a new ReverseRBAC object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReverseRBAC(path string, type_ string, ) *ReverseRBAC {
	this := ReverseRBAC{}
	this.Path = path
	this.Type = type_
	return &this
}

// NewReverseRBACWithDefaults instantiates a new ReverseRBAC object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReverseRBACWithDefaults() *ReverseRBAC {
	this := ReverseRBAC{}
	return &this
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *ReverseRBAC) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReverseRBAC) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *ReverseRBAC) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *ReverseRBAC) SetJson(v bool) {
	o.Json = &v
}

// GetPath returns the Path field value
func (o *ReverseRBAC) GetPath() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *ReverseRBAC) GetPathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *ReverseRBAC) SetPath(v string) {
	o.Path = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *ReverseRBAC) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReverseRBAC) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *ReverseRBAC) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *ReverseRBAC) SetToken(v string) {
	o.Token = &v
}

// GetType returns the Type field value
func (o *ReverseRBAC) GetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ReverseRBAC) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ReverseRBAC) SetType(v string) {
	o.Type = v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *ReverseRBAC) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReverseRBAC) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *ReverseRBAC) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *ReverseRBAC) SetUidToken(v string) {
	o.UidToken = &v
}

func (o ReverseRBAC) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if true {
		toSerialize["path"] = o.Path
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableReverseRBAC struct {
	value *ReverseRBAC
	isSet bool
}

func (v NullableReverseRBAC) Get() *ReverseRBAC {
	return v.value
}

func (v *NullableReverseRBAC) Set(val *ReverseRBAC) {
	v.value = val
	v.isSet = true
}

func (v NullableReverseRBAC) IsSet() bool {
	return v.isSet
}

func (v *NullableReverseRBAC) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReverseRBAC(val *ReverseRBAC) *NullableReverseRBAC {
	return &NullableReverseRBAC{value: val, isSet: true}
}

func (v NullableReverseRBAC) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReverseRBAC) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


