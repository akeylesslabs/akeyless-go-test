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

// DatadogForwardingConfig struct for DatadogForwardingConfig
type DatadogForwardingConfig struct {
	DatadogApiKey *string `json:"datadog_api_key,omitempty"`
	DatadogHost *string `json:"datadog_host,omitempty"`
	DatadogLogService *string `json:"datadog_log_service,omitempty"`
	DatadogLogSource *string `json:"datadog_log_source,omitempty"`
	DatadogLogTags *string `json:"datadog_log_tags,omitempty"`
}

// NewDatadogForwardingConfig instantiates a new DatadogForwardingConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatadogForwardingConfig() *DatadogForwardingConfig {
	this := DatadogForwardingConfig{}
	return &this
}

// NewDatadogForwardingConfigWithDefaults instantiates a new DatadogForwardingConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatadogForwardingConfigWithDefaults() *DatadogForwardingConfig {
	this := DatadogForwardingConfig{}
	return &this
}

// GetDatadogApiKey returns the DatadogApiKey field value if set, zero value otherwise.
func (o *DatadogForwardingConfig) GetDatadogApiKey() string {
	if o == nil || o.DatadogApiKey == nil {
		var ret string
		return ret
	}
	return *o.DatadogApiKey
}

// GetDatadogApiKeyOk returns a tuple with the DatadogApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatadogForwardingConfig) GetDatadogApiKeyOk() (*string, bool) {
	if o == nil || o.DatadogApiKey == nil {
		return nil, false
	}
	return o.DatadogApiKey, true
}

// HasDatadogApiKey returns a boolean if a field has been set.
func (o *DatadogForwardingConfig) HasDatadogApiKey() bool {
	if o != nil && o.DatadogApiKey != nil {
		return true
	}

	return false
}

// SetDatadogApiKey gets a reference to the given string and assigns it to the DatadogApiKey field.
func (o *DatadogForwardingConfig) SetDatadogApiKey(v string) {
	o.DatadogApiKey = &v
}

// GetDatadogHost returns the DatadogHost field value if set, zero value otherwise.
func (o *DatadogForwardingConfig) GetDatadogHost() string {
	if o == nil || o.DatadogHost == nil {
		var ret string
		return ret
	}
	return *o.DatadogHost
}

// GetDatadogHostOk returns a tuple with the DatadogHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatadogForwardingConfig) GetDatadogHostOk() (*string, bool) {
	if o == nil || o.DatadogHost == nil {
		return nil, false
	}
	return o.DatadogHost, true
}

// HasDatadogHost returns a boolean if a field has been set.
func (o *DatadogForwardingConfig) HasDatadogHost() bool {
	if o != nil && o.DatadogHost != nil {
		return true
	}

	return false
}

// SetDatadogHost gets a reference to the given string and assigns it to the DatadogHost field.
func (o *DatadogForwardingConfig) SetDatadogHost(v string) {
	o.DatadogHost = &v
}

// GetDatadogLogService returns the DatadogLogService field value if set, zero value otherwise.
func (o *DatadogForwardingConfig) GetDatadogLogService() string {
	if o == nil || o.DatadogLogService == nil {
		var ret string
		return ret
	}
	return *o.DatadogLogService
}

// GetDatadogLogServiceOk returns a tuple with the DatadogLogService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatadogForwardingConfig) GetDatadogLogServiceOk() (*string, bool) {
	if o == nil || o.DatadogLogService == nil {
		return nil, false
	}
	return o.DatadogLogService, true
}

// HasDatadogLogService returns a boolean if a field has been set.
func (o *DatadogForwardingConfig) HasDatadogLogService() bool {
	if o != nil && o.DatadogLogService != nil {
		return true
	}

	return false
}

// SetDatadogLogService gets a reference to the given string and assigns it to the DatadogLogService field.
func (o *DatadogForwardingConfig) SetDatadogLogService(v string) {
	o.DatadogLogService = &v
}

// GetDatadogLogSource returns the DatadogLogSource field value if set, zero value otherwise.
func (o *DatadogForwardingConfig) GetDatadogLogSource() string {
	if o == nil || o.DatadogLogSource == nil {
		var ret string
		return ret
	}
	return *o.DatadogLogSource
}

// GetDatadogLogSourceOk returns a tuple with the DatadogLogSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatadogForwardingConfig) GetDatadogLogSourceOk() (*string, bool) {
	if o == nil || o.DatadogLogSource == nil {
		return nil, false
	}
	return o.DatadogLogSource, true
}

// HasDatadogLogSource returns a boolean if a field has been set.
func (o *DatadogForwardingConfig) HasDatadogLogSource() bool {
	if o != nil && o.DatadogLogSource != nil {
		return true
	}

	return false
}

// SetDatadogLogSource gets a reference to the given string and assigns it to the DatadogLogSource field.
func (o *DatadogForwardingConfig) SetDatadogLogSource(v string) {
	o.DatadogLogSource = &v
}

// GetDatadogLogTags returns the DatadogLogTags field value if set, zero value otherwise.
func (o *DatadogForwardingConfig) GetDatadogLogTags() string {
	if o == nil || o.DatadogLogTags == nil {
		var ret string
		return ret
	}
	return *o.DatadogLogTags
}

// GetDatadogLogTagsOk returns a tuple with the DatadogLogTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatadogForwardingConfig) GetDatadogLogTagsOk() (*string, bool) {
	if o == nil || o.DatadogLogTags == nil {
		return nil, false
	}
	return o.DatadogLogTags, true
}

// HasDatadogLogTags returns a boolean if a field has been set.
func (o *DatadogForwardingConfig) HasDatadogLogTags() bool {
	if o != nil && o.DatadogLogTags != nil {
		return true
	}

	return false
}

// SetDatadogLogTags gets a reference to the given string and assigns it to the DatadogLogTags field.
func (o *DatadogForwardingConfig) SetDatadogLogTags(v string) {
	o.DatadogLogTags = &v
}

func (o DatadogForwardingConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DatadogApiKey != nil {
		toSerialize["datadog_api_key"] = o.DatadogApiKey
	}
	if o.DatadogHost != nil {
		toSerialize["datadog_host"] = o.DatadogHost
	}
	if o.DatadogLogService != nil {
		toSerialize["datadog_log_service"] = o.DatadogLogService
	}
	if o.DatadogLogSource != nil {
		toSerialize["datadog_log_source"] = o.DatadogLogSource
	}
	if o.DatadogLogTags != nil {
		toSerialize["datadog_log_tags"] = o.DatadogLogTags
	}
	return json.Marshal(toSerialize)
}

type NullableDatadogForwardingConfig struct {
	value *DatadogForwardingConfig
	isSet bool
}

func (v NullableDatadogForwardingConfig) Get() *DatadogForwardingConfig {
	return v.value
}

func (v *NullableDatadogForwardingConfig) Set(val *DatadogForwardingConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableDatadogForwardingConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableDatadogForwardingConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatadogForwardingConfig(val *DatadogForwardingConfig) *NullableDatadogForwardingConfig {
	return &NullableDatadogForwardingConfig{value: val, isSet: true}
}

func (v NullableDatadogForwardingConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatadogForwardingConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

