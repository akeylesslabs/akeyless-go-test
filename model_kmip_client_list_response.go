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

// KMIPClientListResponse struct for KMIPClientListResponse
type KMIPClientListResponse struct {
	Clients *[]KMIPClient `json:"clients,omitempty"`
}

// NewKMIPClientListResponse instantiates a new KMIPClientListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKMIPClientListResponse() *KMIPClientListResponse {
	this := KMIPClientListResponse{}
	return &this
}

// NewKMIPClientListResponseWithDefaults instantiates a new KMIPClientListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKMIPClientListResponseWithDefaults() *KMIPClientListResponse {
	this := KMIPClientListResponse{}
	return &this
}

// GetClients returns the Clients field value if set, zero value otherwise.
func (o *KMIPClientListResponse) GetClients() []KMIPClient {
	if o == nil || o.Clients == nil {
		var ret []KMIPClient
		return ret
	}
	return *o.Clients
}

// GetClientsOk returns a tuple with the Clients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KMIPClientListResponse) GetClientsOk() (*[]KMIPClient, bool) {
	if o == nil || o.Clients == nil {
		return nil, false
	}
	return o.Clients, true
}

// HasClients returns a boolean if a field has been set.
func (o *KMIPClientListResponse) HasClients() bool {
	if o != nil && o.Clients != nil {
		return true
	}

	return false
}

// SetClients gets a reference to the given []KMIPClient and assigns it to the Clients field.
func (o *KMIPClientListResponse) SetClients(v []KMIPClient) {
	o.Clients = &v
}

func (o KMIPClientListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Clients != nil {
		toSerialize["clients"] = o.Clients
	}
	return json.Marshal(toSerialize)
}

type NullableKMIPClientListResponse struct {
	value *KMIPClientListResponse
	isSet bool
}

func (v NullableKMIPClientListResponse) Get() *KMIPClientListResponse {
	return v.value
}

func (v *NullableKMIPClientListResponse) Set(val *KMIPClientListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKMIPClientListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKMIPClientListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKMIPClientListResponse(val *KMIPClientListResponse) *NullableKMIPClientListResponse {
	return &NullableKMIPClientListResponse{value: val, isSet: true}
}

func (v NullableKMIPClientListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKMIPClientListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


