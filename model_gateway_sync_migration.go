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

// GatewaySyncMigration gatewaySyncMigration is a command that sync migration
type GatewaySyncMigration struct {
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Migration name
	Name string `json:"name"`
	// true, for starting synchronization, false for stopping
	StartSync *bool `json:"start-sync,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewGatewaySyncMigration instantiates a new GatewaySyncMigration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewaySyncMigration(name string, ) *GatewaySyncMigration {
	this := GatewaySyncMigration{}
	this.Name = name
	return &this
}

// NewGatewaySyncMigrationWithDefaults instantiates a new GatewaySyncMigration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewaySyncMigrationWithDefaults() *GatewaySyncMigration {
	this := GatewaySyncMigration{}
	return &this
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *GatewaySyncMigration) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewaySyncMigration) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *GatewaySyncMigration) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *GatewaySyncMigration) SetJson(v bool) {
	o.Json = &v
}

// GetName returns the Name field value
func (o *GatewaySyncMigration) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GatewaySyncMigration) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GatewaySyncMigration) SetName(v string) {
	o.Name = v
}

// GetStartSync returns the StartSync field value if set, zero value otherwise.
func (o *GatewaySyncMigration) GetStartSync() bool {
	if o == nil || o.StartSync == nil {
		var ret bool
		return ret
	}
	return *o.StartSync
}

// GetStartSyncOk returns a tuple with the StartSync field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewaySyncMigration) GetStartSyncOk() (*bool, bool) {
	if o == nil || o.StartSync == nil {
		return nil, false
	}
	return o.StartSync, true
}

// HasStartSync returns a boolean if a field has been set.
func (o *GatewaySyncMigration) HasStartSync() bool {
	if o != nil && o.StartSync != nil {
		return true
	}

	return false
}

// SetStartSync gets a reference to the given bool and assigns it to the StartSync field.
func (o *GatewaySyncMigration) SetStartSync(v bool) {
	o.StartSync = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GatewaySyncMigration) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewaySyncMigration) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GatewaySyncMigration) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GatewaySyncMigration) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *GatewaySyncMigration) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewaySyncMigration) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *GatewaySyncMigration) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *GatewaySyncMigration) SetUidToken(v string) {
	o.UidToken = &v
}

func (o GatewaySyncMigration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.StartSync != nil {
		toSerialize["start-sync"] = o.StartSync
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableGatewaySyncMigration struct {
	value *GatewaySyncMigration
	isSet bool
}

func (v NullableGatewaySyncMigration) Get() *GatewaySyncMigration {
	return v.value
}

func (v *NullableGatewaySyncMigration) Set(val *GatewaySyncMigration) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewaySyncMigration) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewaySyncMigration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewaySyncMigration(val *GatewaySyncMigration) *NullableGatewaySyncMigration {
	return &NullableGatewaySyncMigration{value: val, isSet: true}
}

func (v NullableGatewaySyncMigration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewaySyncMigration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


