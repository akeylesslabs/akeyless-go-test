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

// BastionsList struct for BastionsList
type BastionsList struct {
	Clusters *[]BastionListEntry `json:"clusters,omitempty"`
}

// NewBastionsList instantiates a new BastionsList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBastionsList() *BastionsList {
	this := BastionsList{}
	return &this
}

// NewBastionsListWithDefaults instantiates a new BastionsList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBastionsListWithDefaults() *BastionsList {
	this := BastionsList{}
	return &this
}

// GetClusters returns the Clusters field value if set, zero value otherwise.
func (o *BastionsList) GetClusters() []BastionListEntry {
	if o == nil || o.Clusters == nil {
		var ret []BastionListEntry
		return ret
	}
	return *o.Clusters
}

// GetClustersOk returns a tuple with the Clusters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BastionsList) GetClustersOk() (*[]BastionListEntry, bool) {
	if o == nil || o.Clusters == nil {
		return nil, false
	}
	return o.Clusters, true
}

// HasClusters returns a boolean if a field has been set.
func (o *BastionsList) HasClusters() bool {
	if o != nil && o.Clusters != nil {
		return true
	}

	return false
}

// SetClusters gets a reference to the given []BastionListEntry and assigns it to the Clusters field.
func (o *BastionsList) SetClusters(v []BastionListEntry) {
	o.Clusters = &v
}

func (o BastionsList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Clusters != nil {
		toSerialize["clusters"] = o.Clusters
	}
	return json.Marshal(toSerialize)
}

type NullableBastionsList struct {
	value *BastionsList
	isSet bool
}

func (v NullableBastionsList) Get() *BastionsList {
	return v.value
}

func (v *NullableBastionsList) Set(val *BastionsList) {
	v.value = val
	v.isSet = true
}

func (v NullableBastionsList) IsSet() bool {
	return v.isSet
}

func (v *NullableBastionsList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBastionsList(val *BastionsList) *NullableBastionsList {
	return &NullableBastionsList{value: val, isSet: true}
}

func (v NullableBastionsList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBastionsList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


