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

// GetKubeExecCredsOutput struct for GetKubeExecCredsOutput
type GetKubeExecCredsOutput struct {
	ApiVersion *string `json:"apiVersion,omitempty"`
	Kind *string `json:"kind,omitempty"`
	Status *ClientData `json:"status,omitempty"`
}

// NewGetKubeExecCredsOutput instantiates a new GetKubeExecCredsOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetKubeExecCredsOutput() *GetKubeExecCredsOutput {
	this := GetKubeExecCredsOutput{}
	return &this
}

// NewGetKubeExecCredsOutputWithDefaults instantiates a new GetKubeExecCredsOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetKubeExecCredsOutputWithDefaults() *GetKubeExecCredsOutput {
	this := GetKubeExecCredsOutput{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *GetKubeExecCredsOutput) GetApiVersion() string {
	if o == nil || o.ApiVersion == nil {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetKubeExecCredsOutput) GetApiVersionOk() (*string, bool) {
	if o == nil || o.ApiVersion == nil {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *GetKubeExecCredsOutput) HasApiVersion() bool {
	if o != nil && o.ApiVersion != nil {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *GetKubeExecCredsOutput) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *GetKubeExecCredsOutput) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetKubeExecCredsOutput) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *GetKubeExecCredsOutput) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *GetKubeExecCredsOutput) SetKind(v string) {
	o.Kind = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetKubeExecCredsOutput) GetStatus() ClientData {
	if o == nil || o.Status == nil {
		var ret ClientData
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetKubeExecCredsOutput) GetStatusOk() (*ClientData, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetKubeExecCredsOutput) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ClientData and assigns it to the Status field.
func (o *GetKubeExecCredsOutput) SetStatus(v ClientData) {
	o.Status = &v
}

func (o GetKubeExecCredsOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiVersion != nil {
		toSerialize["apiVersion"] = o.ApiVersion
	}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableGetKubeExecCredsOutput struct {
	value *GetKubeExecCredsOutput
	isSet bool
}

func (v NullableGetKubeExecCredsOutput) Get() *GetKubeExecCredsOutput {
	return v.value
}

func (v *NullableGetKubeExecCredsOutput) Set(val *GetKubeExecCredsOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableGetKubeExecCredsOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableGetKubeExecCredsOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetKubeExecCredsOutput(val *GetKubeExecCredsOutput) *NullableGetKubeExecCredsOutput {
	return &NullableGetKubeExecCredsOutput{value: val, isSet: true}
}

func (v NullableGetKubeExecCredsOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetKubeExecCredsOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

