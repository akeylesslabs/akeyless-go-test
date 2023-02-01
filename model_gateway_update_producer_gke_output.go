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

// GatewayUpdateProducerGkeOutput struct for GatewayUpdateProducerGkeOutput
type GatewayUpdateProducerGkeOutput struct {
	ProducerDetails *DSProducerDetails `json:"producer_details,omitempty"`
}

// NewGatewayUpdateProducerGkeOutput instantiates a new GatewayUpdateProducerGkeOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayUpdateProducerGkeOutput() *GatewayUpdateProducerGkeOutput {
	this := GatewayUpdateProducerGkeOutput{}
	return &this
}

// NewGatewayUpdateProducerGkeOutputWithDefaults instantiates a new GatewayUpdateProducerGkeOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayUpdateProducerGkeOutputWithDefaults() *GatewayUpdateProducerGkeOutput {
	this := GatewayUpdateProducerGkeOutput{}
	return &this
}

// GetProducerDetails returns the ProducerDetails field value if set, zero value otherwise.
func (o *GatewayUpdateProducerGkeOutput) GetProducerDetails() DSProducerDetails {
	if o == nil || o.ProducerDetails == nil {
		var ret DSProducerDetails
		return ret
	}
	return *o.ProducerDetails
}

// GetProducerDetailsOk returns a tuple with the ProducerDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerGkeOutput) GetProducerDetailsOk() (*DSProducerDetails, bool) {
	if o == nil || o.ProducerDetails == nil {
		return nil, false
	}
	return o.ProducerDetails, true
}

// HasProducerDetails returns a boolean if a field has been set.
func (o *GatewayUpdateProducerGkeOutput) HasProducerDetails() bool {
	if o != nil && o.ProducerDetails != nil {
		return true
	}

	return false
}

// SetProducerDetails gets a reference to the given DSProducerDetails and assigns it to the ProducerDetails field.
func (o *GatewayUpdateProducerGkeOutput) SetProducerDetails(v DSProducerDetails) {
	o.ProducerDetails = &v
}

func (o GatewayUpdateProducerGkeOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ProducerDetails != nil {
		toSerialize["producer_details"] = o.ProducerDetails
	}
	return json.Marshal(toSerialize)
}

type NullableGatewayUpdateProducerGkeOutput struct {
	value *GatewayUpdateProducerGkeOutput
	isSet bool
}

func (v NullableGatewayUpdateProducerGkeOutput) Get() *GatewayUpdateProducerGkeOutput {
	return v.value
}

func (v *NullableGatewayUpdateProducerGkeOutput) Set(val *GatewayUpdateProducerGkeOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayUpdateProducerGkeOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayUpdateProducerGkeOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayUpdateProducerGkeOutput(val *GatewayUpdateProducerGkeOutput) *NullableGatewayUpdateProducerGkeOutput {
	return &NullableGatewayUpdateProducerGkeOutput{value: val, isSet: true}
}

func (v NullableGatewayUpdateProducerGkeOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayUpdateProducerGkeOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


