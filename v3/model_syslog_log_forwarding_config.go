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

// SyslogLogForwardingConfig struct for SyslogLogForwardingConfig
type SyslogLogForwardingConfig struct {
	SyslogFormatter *string `json:"syslog_formatter,omitempty"`
	SyslogHost *string `json:"syslog_host,omitempty"`
	SyslogNetwork *string `json:"syslog_network,omitempty"`
	SyslogTargetTag *string `json:"syslog_target_tag,omitempty"`
}

// NewSyslogLogForwardingConfig instantiates a new SyslogLogForwardingConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyslogLogForwardingConfig() *SyslogLogForwardingConfig {
	this := SyslogLogForwardingConfig{}
	return &this
}

// NewSyslogLogForwardingConfigWithDefaults instantiates a new SyslogLogForwardingConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyslogLogForwardingConfigWithDefaults() *SyslogLogForwardingConfig {
	this := SyslogLogForwardingConfig{}
	return &this
}

// GetSyslogFormatter returns the SyslogFormatter field value if set, zero value otherwise.
func (o *SyslogLogForwardingConfig) GetSyslogFormatter() string {
	if o == nil || o.SyslogFormatter == nil {
		var ret string
		return ret
	}
	return *o.SyslogFormatter
}

// GetSyslogFormatterOk returns a tuple with the SyslogFormatter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyslogLogForwardingConfig) GetSyslogFormatterOk() (*string, bool) {
	if o == nil || o.SyslogFormatter == nil {
		return nil, false
	}
	return o.SyslogFormatter, true
}

// HasSyslogFormatter returns a boolean if a field has been set.
func (o *SyslogLogForwardingConfig) HasSyslogFormatter() bool {
	if o != nil && o.SyslogFormatter != nil {
		return true
	}

	return false
}

// SetSyslogFormatter gets a reference to the given string and assigns it to the SyslogFormatter field.
func (o *SyslogLogForwardingConfig) SetSyslogFormatter(v string) {
	o.SyslogFormatter = &v
}

// GetSyslogHost returns the SyslogHost field value if set, zero value otherwise.
func (o *SyslogLogForwardingConfig) GetSyslogHost() string {
	if o == nil || o.SyslogHost == nil {
		var ret string
		return ret
	}
	return *o.SyslogHost
}

// GetSyslogHostOk returns a tuple with the SyslogHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyslogLogForwardingConfig) GetSyslogHostOk() (*string, bool) {
	if o == nil || o.SyslogHost == nil {
		return nil, false
	}
	return o.SyslogHost, true
}

// HasSyslogHost returns a boolean if a field has been set.
func (o *SyslogLogForwardingConfig) HasSyslogHost() bool {
	if o != nil && o.SyslogHost != nil {
		return true
	}

	return false
}

// SetSyslogHost gets a reference to the given string and assigns it to the SyslogHost field.
func (o *SyslogLogForwardingConfig) SetSyslogHost(v string) {
	o.SyslogHost = &v
}

// GetSyslogNetwork returns the SyslogNetwork field value if set, zero value otherwise.
func (o *SyslogLogForwardingConfig) GetSyslogNetwork() string {
	if o == nil || o.SyslogNetwork == nil {
		var ret string
		return ret
	}
	return *o.SyslogNetwork
}

// GetSyslogNetworkOk returns a tuple with the SyslogNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyslogLogForwardingConfig) GetSyslogNetworkOk() (*string, bool) {
	if o == nil || o.SyslogNetwork == nil {
		return nil, false
	}
	return o.SyslogNetwork, true
}

// HasSyslogNetwork returns a boolean if a field has been set.
func (o *SyslogLogForwardingConfig) HasSyslogNetwork() bool {
	if o != nil && o.SyslogNetwork != nil {
		return true
	}

	return false
}

// SetSyslogNetwork gets a reference to the given string and assigns it to the SyslogNetwork field.
func (o *SyslogLogForwardingConfig) SetSyslogNetwork(v string) {
	o.SyslogNetwork = &v
}

// GetSyslogTargetTag returns the SyslogTargetTag field value if set, zero value otherwise.
func (o *SyslogLogForwardingConfig) GetSyslogTargetTag() string {
	if o == nil || o.SyslogTargetTag == nil {
		var ret string
		return ret
	}
	return *o.SyslogTargetTag
}

// GetSyslogTargetTagOk returns a tuple with the SyslogTargetTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyslogLogForwardingConfig) GetSyslogTargetTagOk() (*string, bool) {
	if o == nil || o.SyslogTargetTag == nil {
		return nil, false
	}
	return o.SyslogTargetTag, true
}

// HasSyslogTargetTag returns a boolean if a field has been set.
func (o *SyslogLogForwardingConfig) HasSyslogTargetTag() bool {
	if o != nil && o.SyslogTargetTag != nil {
		return true
	}

	return false
}

// SetSyslogTargetTag gets a reference to the given string and assigns it to the SyslogTargetTag field.
func (o *SyslogLogForwardingConfig) SetSyslogTargetTag(v string) {
	o.SyslogTargetTag = &v
}

func (o SyslogLogForwardingConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SyslogFormatter != nil {
		toSerialize["syslog_formatter"] = o.SyslogFormatter
	}
	if o.SyslogHost != nil {
		toSerialize["syslog_host"] = o.SyslogHost
	}
	if o.SyslogNetwork != nil {
		toSerialize["syslog_network"] = o.SyslogNetwork
	}
	if o.SyslogTargetTag != nil {
		toSerialize["syslog_target_tag"] = o.SyslogTargetTag
	}
	return json.Marshal(toSerialize)
}

type NullableSyslogLogForwardingConfig struct {
	value *SyslogLogForwardingConfig
	isSet bool
}

func (v NullableSyslogLogForwardingConfig) Get() *SyslogLogForwardingConfig {
	return v.value
}

func (v *NullableSyslogLogForwardingConfig) Set(val *SyslogLogForwardingConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableSyslogLogForwardingConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableSyslogLogForwardingConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyslogLogForwardingConfig(val *SyslogLogForwardingConfig) *NullableSyslogLogForwardingConfig {
	return &NullableSyslogLogForwardingConfig{value: val, isSet: true}
}

func (v NullableSyslogLogForwardingConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyslogLogForwardingConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

