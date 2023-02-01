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

// VaultlessTokenizerInfo struct for VaultlessTokenizerInfo
type VaultlessTokenizerInfo struct {
	EmailTokenizerInfo *EmailTokenizerInfo `json:"email_tokenizer_info,omitempty"`
	KeyName *string `json:"key_name,omitempty"`
	RegexpTokenizerInfo *RegexpTokenizerInfo `json:"regexp_tokenizer_info,omitempty"`
	TemplateType *string `json:"template_type,omitempty"`
	// Tweak used in the case of internal tweak type
	Tweak *string `json:"tweak,omitempty"`
	TweakType *string `json:"tweak_type,omitempty"`
}

// NewVaultlessTokenizerInfo instantiates a new VaultlessTokenizerInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVaultlessTokenizerInfo() *VaultlessTokenizerInfo {
	this := VaultlessTokenizerInfo{}
	return &this
}

// NewVaultlessTokenizerInfoWithDefaults instantiates a new VaultlessTokenizerInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVaultlessTokenizerInfoWithDefaults() *VaultlessTokenizerInfo {
	this := VaultlessTokenizerInfo{}
	return &this
}

// GetEmailTokenizerInfo returns the EmailTokenizerInfo field value if set, zero value otherwise.
func (o *VaultlessTokenizerInfo) GetEmailTokenizerInfo() EmailTokenizerInfo {
	if o == nil || o.EmailTokenizerInfo == nil {
		var ret EmailTokenizerInfo
		return ret
	}
	return *o.EmailTokenizerInfo
}

// GetEmailTokenizerInfoOk returns a tuple with the EmailTokenizerInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VaultlessTokenizerInfo) GetEmailTokenizerInfoOk() (*EmailTokenizerInfo, bool) {
	if o == nil || o.EmailTokenizerInfo == nil {
		return nil, false
	}
	return o.EmailTokenizerInfo, true
}

// HasEmailTokenizerInfo returns a boolean if a field has been set.
func (o *VaultlessTokenizerInfo) HasEmailTokenizerInfo() bool {
	if o != nil && o.EmailTokenizerInfo != nil {
		return true
	}

	return false
}

// SetEmailTokenizerInfo gets a reference to the given EmailTokenizerInfo and assigns it to the EmailTokenizerInfo field.
func (o *VaultlessTokenizerInfo) SetEmailTokenizerInfo(v EmailTokenizerInfo) {
	o.EmailTokenizerInfo = &v
}

// GetKeyName returns the KeyName field value if set, zero value otherwise.
func (o *VaultlessTokenizerInfo) GetKeyName() string {
	if o == nil || o.KeyName == nil {
		var ret string
		return ret
	}
	return *o.KeyName
}

// GetKeyNameOk returns a tuple with the KeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VaultlessTokenizerInfo) GetKeyNameOk() (*string, bool) {
	if o == nil || o.KeyName == nil {
		return nil, false
	}
	return o.KeyName, true
}

// HasKeyName returns a boolean if a field has been set.
func (o *VaultlessTokenizerInfo) HasKeyName() bool {
	if o != nil && o.KeyName != nil {
		return true
	}

	return false
}

// SetKeyName gets a reference to the given string and assigns it to the KeyName field.
func (o *VaultlessTokenizerInfo) SetKeyName(v string) {
	o.KeyName = &v
}

// GetRegexpTokenizerInfo returns the RegexpTokenizerInfo field value if set, zero value otherwise.
func (o *VaultlessTokenizerInfo) GetRegexpTokenizerInfo() RegexpTokenizerInfo {
	if o == nil || o.RegexpTokenizerInfo == nil {
		var ret RegexpTokenizerInfo
		return ret
	}
	return *o.RegexpTokenizerInfo
}

// GetRegexpTokenizerInfoOk returns a tuple with the RegexpTokenizerInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VaultlessTokenizerInfo) GetRegexpTokenizerInfoOk() (*RegexpTokenizerInfo, bool) {
	if o == nil || o.RegexpTokenizerInfo == nil {
		return nil, false
	}
	return o.RegexpTokenizerInfo, true
}

// HasRegexpTokenizerInfo returns a boolean if a field has been set.
func (o *VaultlessTokenizerInfo) HasRegexpTokenizerInfo() bool {
	if o != nil && o.RegexpTokenizerInfo != nil {
		return true
	}

	return false
}

// SetRegexpTokenizerInfo gets a reference to the given RegexpTokenizerInfo and assigns it to the RegexpTokenizerInfo field.
func (o *VaultlessTokenizerInfo) SetRegexpTokenizerInfo(v RegexpTokenizerInfo) {
	o.RegexpTokenizerInfo = &v
}

// GetTemplateType returns the TemplateType field value if set, zero value otherwise.
func (o *VaultlessTokenizerInfo) GetTemplateType() string {
	if o == nil || o.TemplateType == nil {
		var ret string
		return ret
	}
	return *o.TemplateType
}

// GetTemplateTypeOk returns a tuple with the TemplateType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VaultlessTokenizerInfo) GetTemplateTypeOk() (*string, bool) {
	if o == nil || o.TemplateType == nil {
		return nil, false
	}
	return o.TemplateType, true
}

// HasTemplateType returns a boolean if a field has been set.
func (o *VaultlessTokenizerInfo) HasTemplateType() bool {
	if o != nil && o.TemplateType != nil {
		return true
	}

	return false
}

// SetTemplateType gets a reference to the given string and assigns it to the TemplateType field.
func (o *VaultlessTokenizerInfo) SetTemplateType(v string) {
	o.TemplateType = &v
}

// GetTweak returns the Tweak field value if set, zero value otherwise.
func (o *VaultlessTokenizerInfo) GetTweak() string {
	if o == nil || o.Tweak == nil {
		var ret string
		return ret
	}
	return *o.Tweak
}

// GetTweakOk returns a tuple with the Tweak field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VaultlessTokenizerInfo) GetTweakOk() (*string, bool) {
	if o == nil || o.Tweak == nil {
		return nil, false
	}
	return o.Tweak, true
}

// HasTweak returns a boolean if a field has been set.
func (o *VaultlessTokenizerInfo) HasTweak() bool {
	if o != nil && o.Tweak != nil {
		return true
	}

	return false
}

// SetTweak gets a reference to the given string and assigns it to the Tweak field.
func (o *VaultlessTokenizerInfo) SetTweak(v string) {
	o.Tweak = &v
}

// GetTweakType returns the TweakType field value if set, zero value otherwise.
func (o *VaultlessTokenizerInfo) GetTweakType() string {
	if o == nil || o.TweakType == nil {
		var ret string
		return ret
	}
	return *o.TweakType
}

// GetTweakTypeOk returns a tuple with the TweakType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VaultlessTokenizerInfo) GetTweakTypeOk() (*string, bool) {
	if o == nil || o.TweakType == nil {
		return nil, false
	}
	return o.TweakType, true
}

// HasTweakType returns a boolean if a field has been set.
func (o *VaultlessTokenizerInfo) HasTweakType() bool {
	if o != nil && o.TweakType != nil {
		return true
	}

	return false
}

// SetTweakType gets a reference to the given string and assigns it to the TweakType field.
func (o *VaultlessTokenizerInfo) SetTweakType(v string) {
	o.TweakType = &v
}

func (o VaultlessTokenizerInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EmailTokenizerInfo != nil {
		toSerialize["email_tokenizer_info"] = o.EmailTokenizerInfo
	}
	if o.KeyName != nil {
		toSerialize["key_name"] = o.KeyName
	}
	if o.RegexpTokenizerInfo != nil {
		toSerialize["regexp_tokenizer_info"] = o.RegexpTokenizerInfo
	}
	if o.TemplateType != nil {
		toSerialize["template_type"] = o.TemplateType
	}
	if o.Tweak != nil {
		toSerialize["tweak"] = o.Tweak
	}
	if o.TweakType != nil {
		toSerialize["tweak_type"] = o.TweakType
	}
	return json.Marshal(toSerialize)
}

type NullableVaultlessTokenizerInfo struct {
	value *VaultlessTokenizerInfo
	isSet bool
}

func (v NullableVaultlessTokenizerInfo) Get() *VaultlessTokenizerInfo {
	return v.value
}

func (v *NullableVaultlessTokenizerInfo) Set(val *VaultlessTokenizerInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVaultlessTokenizerInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVaultlessTokenizerInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVaultlessTokenizerInfo(val *VaultlessTokenizerInfo) *NullableVaultlessTokenizerInfo {
	return &NullableVaultlessTokenizerInfo{value: val, isSet: true}
}

func (v NullableVaultlessTokenizerInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVaultlessTokenizerInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


