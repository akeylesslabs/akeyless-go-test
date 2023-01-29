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

// GatewayDeleteProducerOutput struct for GatewayDeleteProducerOutput
type GatewayDeleteProducerOutput struct {
	ProducerName *string `json:"producer_name,omitempty"`
}

// NewGatewayDeleteProducerOutput instantiates a new GatewayDeleteProducerOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayDeleteProducerOutput() *GatewayDeleteProducerOutput {
	this := GatewayDeleteProducerOutput{}
	return &this
}

// NewGatewayDeleteProducerOutputWithDefaults instantiates a new GatewayDeleteProducerOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayDeleteProducerOutputWithDefaults() *GatewayDeleteProducerOutput {
	this := GatewayDeleteProducerOutput{}
	return &this
}

// GetProducerName returns the ProducerName field value if set, zero value otherwise.
func (o *GatewayDeleteProducerOutput) GetProducerName() string {
	if o == nil || o.ProducerName == nil {
		var ret string
		return ret
	}
	return *o.ProducerName
}

// GetProducerNameOk returns a tuple with the ProducerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayDeleteProducerOutput) GetProducerNameOk() (*string, bool) {
	if o == nil || o.ProducerName == nil {
		return nil, false
	}
	return o.ProducerName, true
}

// HasProducerName returns a boolean if a field has been set.
func (o *GatewayDeleteProducerOutput) HasProducerName() bool {
	if o != nil && o.ProducerName != nil {
		return true
	}

	return false
}

// SetProducerName gets a reference to the given string and assigns it to the ProducerName field.
func (o *GatewayDeleteProducerOutput) SetProducerName(v string) {
	o.ProducerName = &v
}

func (o GatewayDeleteProducerOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ProducerName != nil {
		toSerialize["producer_name"] = o.ProducerName
	}
	return json.Marshal(toSerialize)
}

type NullableGatewayDeleteProducerOutput struct {
	value *GatewayDeleteProducerOutput
	isSet bool
}

func (v NullableGatewayDeleteProducerOutput) Get() *GatewayDeleteProducerOutput {
	return v.value
}

func (v *NullableGatewayDeleteProducerOutput) Set(val *GatewayDeleteProducerOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayDeleteProducerOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayDeleteProducerOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayDeleteProducerOutput(val *GatewayDeleteProducerOutput) *NullableGatewayDeleteProducerOutput {
	return &NullableGatewayDeleteProducerOutput{value: val, isSet: true}
}

func (v NullableGatewayDeleteProducerOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayDeleteProducerOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


