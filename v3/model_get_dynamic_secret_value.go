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

// GetDynamicSecretValue struct for GetDynamicSecretValue
type GetDynamicSecretValue struct {
	// Optional arguments as key=value pairs or JSON strings, e.g - \\\"--args=csr=base64_encoded_csr --args=common_name=bar\\\" or args='{\\\"csr\\\":\\\"base64_encoded_csr\\\"}. It is possible to combine both formats.'
	Args *[]string `json:"args,omitempty"`
	// Host
	Host *string `json:"host,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Dynamic secret name
	Name string `json:"name"`
	// Target Name
	Target *string `json:"target,omitempty"`
	// Timeout in seconds
	Timeout *int64 `json:"timeout,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewGetDynamicSecretValue instantiates a new GetDynamicSecretValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDynamicSecretValue(name string, ) *GetDynamicSecretValue {
	this := GetDynamicSecretValue{}
	this.Name = name
	var timeout int64 = 15
	this.Timeout = &timeout
	return &this
}

// NewGetDynamicSecretValueWithDefaults instantiates a new GetDynamicSecretValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDynamicSecretValueWithDefaults() *GetDynamicSecretValue {
	this := GetDynamicSecretValue{}
	var timeout int64 = 15
	this.Timeout = &timeout
	return &this
}

// GetArgs returns the Args field value if set, zero value otherwise.
func (o *GetDynamicSecretValue) GetArgs() []string {
	if o == nil || o.Args == nil {
		var ret []string
		return ret
	}
	return *o.Args
}

// GetArgsOk returns a tuple with the Args field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDynamicSecretValue) GetArgsOk() (*[]string, bool) {
	if o == nil || o.Args == nil {
		return nil, false
	}
	return o.Args, true
}

// HasArgs returns a boolean if a field has been set.
func (o *GetDynamicSecretValue) HasArgs() bool {
	if o != nil && o.Args != nil {
		return true
	}

	return false
}

// SetArgs gets a reference to the given []string and assigns it to the Args field.
func (o *GetDynamicSecretValue) SetArgs(v []string) {
	o.Args = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *GetDynamicSecretValue) GetHost() string {
	if o == nil || o.Host == nil {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDynamicSecretValue) GetHostOk() (*string, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *GetDynamicSecretValue) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *GetDynamicSecretValue) SetHost(v string) {
	o.Host = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *GetDynamicSecretValue) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDynamicSecretValue) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *GetDynamicSecretValue) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *GetDynamicSecretValue) SetJson(v bool) {
	o.Json = &v
}

// GetName returns the Name field value
func (o *GetDynamicSecretValue) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetDynamicSecretValue) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetDynamicSecretValue) SetName(v string) {
	o.Name = v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *GetDynamicSecretValue) GetTarget() string {
	if o == nil || o.Target == nil {
		var ret string
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDynamicSecretValue) GetTargetOk() (*string, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *GetDynamicSecretValue) HasTarget() bool {
	if o != nil && o.Target != nil {
		return true
	}

	return false
}

// SetTarget gets a reference to the given string and assigns it to the Target field.
func (o *GetDynamicSecretValue) SetTarget(v string) {
	o.Target = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *GetDynamicSecretValue) GetTimeout() int64 {
	if o == nil || o.Timeout == nil {
		var ret int64
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDynamicSecretValue) GetTimeoutOk() (*int64, bool) {
	if o == nil || o.Timeout == nil {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *GetDynamicSecretValue) HasTimeout() bool {
	if o != nil && o.Timeout != nil {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int64 and assigns it to the Timeout field.
func (o *GetDynamicSecretValue) SetTimeout(v int64) {
	o.Timeout = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GetDynamicSecretValue) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDynamicSecretValue) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GetDynamicSecretValue) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GetDynamicSecretValue) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *GetDynamicSecretValue) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDynamicSecretValue) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *GetDynamicSecretValue) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *GetDynamicSecretValue) SetUidToken(v string) {
	o.UidToken = &v
}

func (o GetDynamicSecretValue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Args != nil {
		toSerialize["args"] = o.Args
	}
	if o.Host != nil {
		toSerialize["host"] = o.Host
	}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}
	if o.Timeout != nil {
		toSerialize["timeout"] = o.Timeout
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableGetDynamicSecretValue struct {
	value *GetDynamicSecretValue
	isSet bool
}

func (v NullableGetDynamicSecretValue) Get() *GetDynamicSecretValue {
	return v.value
}

func (v *NullableGetDynamicSecretValue) Set(val *GetDynamicSecretValue) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDynamicSecretValue) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDynamicSecretValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDynamicSecretValue(val *GetDynamicSecretValue) *NullableGetDynamicSecretValue {
	return &NullableGetDynamicSecretValue{value: val, isSet: true}
}

func (v NullableGetDynamicSecretValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDynamicSecretValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

