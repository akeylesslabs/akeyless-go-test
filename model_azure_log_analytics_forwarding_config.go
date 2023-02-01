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

// AzureLogAnalyticsForwardingConfig struct for AzureLogAnalyticsForwardingConfig
type AzureLogAnalyticsForwardingConfig struct {
	AzureWorkspaceId *string `json:"azure_workspace_id,omitempty"`
	AzureWorkspaceKey *string `json:"azure_workspace_key,omitempty"`
}

// NewAzureLogAnalyticsForwardingConfig instantiates a new AzureLogAnalyticsForwardingConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureLogAnalyticsForwardingConfig() *AzureLogAnalyticsForwardingConfig {
	this := AzureLogAnalyticsForwardingConfig{}
	return &this
}

// NewAzureLogAnalyticsForwardingConfigWithDefaults instantiates a new AzureLogAnalyticsForwardingConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureLogAnalyticsForwardingConfigWithDefaults() *AzureLogAnalyticsForwardingConfig {
	this := AzureLogAnalyticsForwardingConfig{}
	return &this
}

// GetAzureWorkspaceId returns the AzureWorkspaceId field value if set, zero value otherwise.
func (o *AzureLogAnalyticsForwardingConfig) GetAzureWorkspaceId() string {
	if o == nil || o.AzureWorkspaceId == nil {
		var ret string
		return ret
	}
	return *o.AzureWorkspaceId
}

// GetAzureWorkspaceIdOk returns a tuple with the AzureWorkspaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureLogAnalyticsForwardingConfig) GetAzureWorkspaceIdOk() (*string, bool) {
	if o == nil || o.AzureWorkspaceId == nil {
		return nil, false
	}
	return o.AzureWorkspaceId, true
}

// HasAzureWorkspaceId returns a boolean if a field has been set.
func (o *AzureLogAnalyticsForwardingConfig) HasAzureWorkspaceId() bool {
	if o != nil && o.AzureWorkspaceId != nil {
		return true
	}

	return false
}

// SetAzureWorkspaceId gets a reference to the given string and assigns it to the AzureWorkspaceId field.
func (o *AzureLogAnalyticsForwardingConfig) SetAzureWorkspaceId(v string) {
	o.AzureWorkspaceId = &v
}

// GetAzureWorkspaceKey returns the AzureWorkspaceKey field value if set, zero value otherwise.
func (o *AzureLogAnalyticsForwardingConfig) GetAzureWorkspaceKey() string {
	if o == nil || o.AzureWorkspaceKey == nil {
		var ret string
		return ret
	}
	return *o.AzureWorkspaceKey
}

// GetAzureWorkspaceKeyOk returns a tuple with the AzureWorkspaceKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureLogAnalyticsForwardingConfig) GetAzureWorkspaceKeyOk() (*string, bool) {
	if o == nil || o.AzureWorkspaceKey == nil {
		return nil, false
	}
	return o.AzureWorkspaceKey, true
}

// HasAzureWorkspaceKey returns a boolean if a field has been set.
func (o *AzureLogAnalyticsForwardingConfig) HasAzureWorkspaceKey() bool {
	if o != nil && o.AzureWorkspaceKey != nil {
		return true
	}

	return false
}

// SetAzureWorkspaceKey gets a reference to the given string and assigns it to the AzureWorkspaceKey field.
func (o *AzureLogAnalyticsForwardingConfig) SetAzureWorkspaceKey(v string) {
	o.AzureWorkspaceKey = &v
}

func (o AzureLogAnalyticsForwardingConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AzureWorkspaceId != nil {
		toSerialize["azure_workspace_id"] = o.AzureWorkspaceId
	}
	if o.AzureWorkspaceKey != nil {
		toSerialize["azure_workspace_key"] = o.AzureWorkspaceKey
	}
	return json.Marshal(toSerialize)
}

type NullableAzureLogAnalyticsForwardingConfig struct {
	value *AzureLogAnalyticsForwardingConfig
	isSet bool
}

func (v NullableAzureLogAnalyticsForwardingConfig) Get() *AzureLogAnalyticsForwardingConfig {
	return v.value
}

func (v *NullableAzureLogAnalyticsForwardingConfig) Set(val *AzureLogAnalyticsForwardingConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureLogAnalyticsForwardingConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureLogAnalyticsForwardingConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureLogAnalyticsForwardingConfig(val *AzureLogAnalyticsForwardingConfig) *NullableAzureLogAnalyticsForwardingConfig {
	return &NullableAzureLogAnalyticsForwardingConfig{value: val, isSet: true}
}

func (v NullableAzureLogAnalyticsForwardingConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureLogAnalyticsForwardingConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


