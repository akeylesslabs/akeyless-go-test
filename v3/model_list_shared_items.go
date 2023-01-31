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

// ListSharedItems listSharedItems is a command to list all the items been shared
type ListSharedItems struct {
	// for personal password manager
	Accessibility *string `json:"accessibility,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewListSharedItems instantiates a new ListSharedItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListSharedItems() *ListSharedItems {
	this := ListSharedItems{}
	var accessibility string = "regular"
	this.Accessibility = &accessibility
	return &this
}

// NewListSharedItemsWithDefaults instantiates a new ListSharedItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListSharedItemsWithDefaults() *ListSharedItems {
	this := ListSharedItems{}
	var accessibility string = "regular"
	this.Accessibility = &accessibility
	return &this
}

// GetAccessibility returns the Accessibility field value if set, zero value otherwise.
func (o *ListSharedItems) GetAccessibility() string {
	if o == nil || o.Accessibility == nil {
		var ret string
		return ret
	}
	return *o.Accessibility
}

// GetAccessibilityOk returns a tuple with the Accessibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSharedItems) GetAccessibilityOk() (*string, bool) {
	if o == nil || o.Accessibility == nil {
		return nil, false
	}
	return o.Accessibility, true
}

// HasAccessibility returns a boolean if a field has been set.
func (o *ListSharedItems) HasAccessibility() bool {
	if o != nil && o.Accessibility != nil {
		return true
	}

	return false
}

// SetAccessibility gets a reference to the given string and assigns it to the Accessibility field.
func (o *ListSharedItems) SetAccessibility(v string) {
	o.Accessibility = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *ListSharedItems) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSharedItems) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *ListSharedItems) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *ListSharedItems) SetJson(v bool) {
	o.Json = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *ListSharedItems) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSharedItems) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *ListSharedItems) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *ListSharedItems) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *ListSharedItems) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSharedItems) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *ListSharedItems) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *ListSharedItems) SetUidToken(v string) {
	o.UidToken = &v
}

func (o ListSharedItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Accessibility != nil {
		toSerialize["accessibility"] = o.Accessibility
	}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableListSharedItems struct {
	value *ListSharedItems
	isSet bool
}

func (v NullableListSharedItems) Get() *ListSharedItems {
	return v.value
}

func (v *NullableListSharedItems) Set(val *ListSharedItems) {
	v.value = val
	v.isSet = true
}

func (v NullableListSharedItems) IsSet() bool {
	return v.isSet
}

func (v *NullableListSharedItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListSharedItems(val *ListSharedItems) *NullableListSharedItems {
	return &NullableListSharedItems{value: val, isSet: true}
}

func (v NullableListSharedItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListSharedItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


