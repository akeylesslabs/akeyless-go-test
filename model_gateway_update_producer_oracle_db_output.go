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

// GatewayUpdateProducerOracleDbOutput struct for GatewayUpdateProducerOracleDbOutput
type GatewayUpdateProducerOracleDbOutput struct {
	ProducerDetails *DSProducerDetails `json:"producer_details,omitempty"`
}

// NewGatewayUpdateProducerOracleDbOutput instantiates a new GatewayUpdateProducerOracleDbOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayUpdateProducerOracleDbOutput() *GatewayUpdateProducerOracleDbOutput {
	this := GatewayUpdateProducerOracleDbOutput{}
	return &this
}

// NewGatewayUpdateProducerOracleDbOutputWithDefaults instantiates a new GatewayUpdateProducerOracleDbOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayUpdateProducerOracleDbOutputWithDefaults() *GatewayUpdateProducerOracleDbOutput {
	this := GatewayUpdateProducerOracleDbOutput{}
	return &this
}

// GetProducerDetails returns the ProducerDetails field value if set, zero value otherwise.
func (o *GatewayUpdateProducerOracleDbOutput) GetProducerDetails() DSProducerDetails {
	if o == nil || o.ProducerDetails == nil {
		var ret DSProducerDetails
		return ret
	}
	return *o.ProducerDetails
}

// GetProducerDetailsOk returns a tuple with the ProducerDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerOracleDbOutput) GetProducerDetailsOk() (*DSProducerDetails, bool) {
	if o == nil || o.ProducerDetails == nil {
		return nil, false
	}
	return o.ProducerDetails, true
}

// HasProducerDetails returns a boolean if a field has been set.
func (o *GatewayUpdateProducerOracleDbOutput) HasProducerDetails() bool {
	if o != nil && o.ProducerDetails != nil {
		return true
	}

	return false
}

// SetProducerDetails gets a reference to the given DSProducerDetails and assigns it to the ProducerDetails field.
func (o *GatewayUpdateProducerOracleDbOutput) SetProducerDetails(v DSProducerDetails) {
	o.ProducerDetails = &v
}

func (o GatewayUpdateProducerOracleDbOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ProducerDetails != nil {
		toSerialize["producer_details"] = o.ProducerDetails
	}
	return json.Marshal(toSerialize)
}

type NullableGatewayUpdateProducerOracleDbOutput struct {
	value *GatewayUpdateProducerOracleDbOutput
	isSet bool
}

func (v NullableGatewayUpdateProducerOracleDbOutput) Get() *GatewayUpdateProducerOracleDbOutput {
	return v.value
}

func (v *NullableGatewayUpdateProducerOracleDbOutput) Set(val *GatewayUpdateProducerOracleDbOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayUpdateProducerOracleDbOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayUpdateProducerOracleDbOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayUpdateProducerOracleDbOutput(val *GatewayUpdateProducerOracleDbOutput) *NullableGatewayUpdateProducerOracleDbOutput {
	return &NullableGatewayUpdateProducerOracleDbOutput{value: val, isSet: true}
}

func (v NullableGatewayUpdateProducerOracleDbOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayUpdateProducerOracleDbOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


