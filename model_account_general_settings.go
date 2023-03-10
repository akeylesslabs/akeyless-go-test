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

// AccountGeneralSettings AccountGeneralSettings describes general settings for an account
type AccountGeneralSettings struct {
	DataProtectionSection *DataProtectionSection `json:"data_protection_section,omitempty"`
	EnableRequestForAccess *bool `json:"enable_request_for_access,omitempty"`
	PasswordPolicy *PasswordPolicyInfo `json:"password_policy,omitempty"`
}

// NewAccountGeneralSettings instantiates a new AccountGeneralSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountGeneralSettings() *AccountGeneralSettings {
	this := AccountGeneralSettings{}
	return &this
}

// NewAccountGeneralSettingsWithDefaults instantiates a new AccountGeneralSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountGeneralSettingsWithDefaults() *AccountGeneralSettings {
	this := AccountGeneralSettings{}
	return &this
}

// GetDataProtectionSection returns the DataProtectionSection field value if set, zero value otherwise.
func (o *AccountGeneralSettings) GetDataProtectionSection() DataProtectionSection {
	if o == nil || o.DataProtectionSection == nil {
		var ret DataProtectionSection
		return ret
	}
	return *o.DataProtectionSection
}

// GetDataProtectionSectionOk returns a tuple with the DataProtectionSection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountGeneralSettings) GetDataProtectionSectionOk() (*DataProtectionSection, bool) {
	if o == nil || o.DataProtectionSection == nil {
		return nil, false
	}
	return o.DataProtectionSection, true
}

// HasDataProtectionSection returns a boolean if a field has been set.
func (o *AccountGeneralSettings) HasDataProtectionSection() bool {
	if o != nil && o.DataProtectionSection != nil {
		return true
	}

	return false
}

// SetDataProtectionSection gets a reference to the given DataProtectionSection and assigns it to the DataProtectionSection field.
func (o *AccountGeneralSettings) SetDataProtectionSection(v DataProtectionSection) {
	o.DataProtectionSection = &v
}

// GetEnableRequestForAccess returns the EnableRequestForAccess field value if set, zero value otherwise.
func (o *AccountGeneralSettings) GetEnableRequestForAccess() bool {
	if o == nil || o.EnableRequestForAccess == nil {
		var ret bool
		return ret
	}
	return *o.EnableRequestForAccess
}

// GetEnableRequestForAccessOk returns a tuple with the EnableRequestForAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountGeneralSettings) GetEnableRequestForAccessOk() (*bool, bool) {
	if o == nil || o.EnableRequestForAccess == nil {
		return nil, false
	}
	return o.EnableRequestForAccess, true
}

// HasEnableRequestForAccess returns a boolean if a field has been set.
func (o *AccountGeneralSettings) HasEnableRequestForAccess() bool {
	if o != nil && o.EnableRequestForAccess != nil {
		return true
	}

	return false
}

// SetEnableRequestForAccess gets a reference to the given bool and assigns it to the EnableRequestForAccess field.
func (o *AccountGeneralSettings) SetEnableRequestForAccess(v bool) {
	o.EnableRequestForAccess = &v
}

// GetPasswordPolicy returns the PasswordPolicy field value if set, zero value otherwise.
func (o *AccountGeneralSettings) GetPasswordPolicy() PasswordPolicyInfo {
	if o == nil || o.PasswordPolicy == nil {
		var ret PasswordPolicyInfo
		return ret
	}
	return *o.PasswordPolicy
}

// GetPasswordPolicyOk returns a tuple with the PasswordPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountGeneralSettings) GetPasswordPolicyOk() (*PasswordPolicyInfo, bool) {
	if o == nil || o.PasswordPolicy == nil {
		return nil, false
	}
	return o.PasswordPolicy, true
}

// HasPasswordPolicy returns a boolean if a field has been set.
func (o *AccountGeneralSettings) HasPasswordPolicy() bool {
	if o != nil && o.PasswordPolicy != nil {
		return true
	}

	return false
}

// SetPasswordPolicy gets a reference to the given PasswordPolicyInfo and assigns it to the PasswordPolicy field.
func (o *AccountGeneralSettings) SetPasswordPolicy(v PasswordPolicyInfo) {
	o.PasswordPolicy = &v
}

func (o AccountGeneralSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DataProtectionSection != nil {
		toSerialize["data_protection_section"] = o.DataProtectionSection
	}
	if o.EnableRequestForAccess != nil {
		toSerialize["enable_request_for_access"] = o.EnableRequestForAccess
	}
	if o.PasswordPolicy != nil {
		toSerialize["password_policy"] = o.PasswordPolicy
	}
	return json.Marshal(toSerialize)
}

type NullableAccountGeneralSettings struct {
	value *AccountGeneralSettings
	isSet bool
}

func (v NullableAccountGeneralSettings) Get() *AccountGeneralSettings {
	return v.value
}

func (v *NullableAccountGeneralSettings) Set(val *AccountGeneralSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountGeneralSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountGeneralSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountGeneralSettings(val *AccountGeneralSettings) *NullableAccountGeneralSettings {
	return &NullableAccountGeneralSettings{value: val, isSet: true}
}

func (v NullableAccountGeneralSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountGeneralSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


