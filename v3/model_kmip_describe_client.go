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

// KmipDescribeClient struct for KmipDescribeClient
type KmipDescribeClient struct {
	ClientId *string `json:"client-id,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	Name *string `json:"name,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewKmipDescribeClient instantiates a new KmipDescribeClient object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKmipDescribeClient() *KmipDescribeClient {
	this := KmipDescribeClient{}
	return &this
}

// NewKmipDescribeClientWithDefaults instantiates a new KmipDescribeClient object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKmipDescribeClientWithDefaults() *KmipDescribeClient {
	this := KmipDescribeClient{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *KmipDescribeClient) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KmipDescribeClient) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *KmipDescribeClient) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *KmipDescribeClient) SetClientId(v string) {
	o.ClientId = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *KmipDescribeClient) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KmipDescribeClient) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *KmipDescribeClient) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *KmipDescribeClient) SetJson(v bool) {
	o.Json = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *KmipDescribeClient) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KmipDescribeClient) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *KmipDescribeClient) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *KmipDescribeClient) SetName(v string) {
	o.Name = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *KmipDescribeClient) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KmipDescribeClient) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *KmipDescribeClient) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *KmipDescribeClient) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *KmipDescribeClient) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KmipDescribeClient) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *KmipDescribeClient) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *KmipDescribeClient) SetUidToken(v string) {
	o.UidToken = &v
}

func (o KmipDescribeClient) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client-id"] = o.ClientId
	}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableKmipDescribeClient struct {
	value *KmipDescribeClient
	isSet bool
}

func (v NullableKmipDescribeClient) Get() *KmipDescribeClient {
	return v.value
}

func (v *NullableKmipDescribeClient) Set(val *KmipDescribeClient) {
	v.value = val
	v.isSet = true
}

func (v NullableKmipDescribeClient) IsSet() bool {
	return v.isSet
}

func (v *NullableKmipDescribeClient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKmipDescribeClient(val *KmipDescribeClient) *NullableKmipDescribeClient {
	return &NullableKmipDescribeClient{value: val, isSet: true}
}

func (v NullableKmipDescribeClient) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKmipDescribeClient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


