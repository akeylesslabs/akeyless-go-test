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

// CustomerFullAddress struct for CustomerFullAddress
type CustomerFullAddress struct {
	City *string `json:"city,omitempty"`
	Country *string `json:"country,omitempty"`
	PostalCode *string `json:"postal_code,omitempty"`
	Street *string `json:"street,omitempty"`
}

// NewCustomerFullAddress instantiates a new CustomerFullAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerFullAddress() *CustomerFullAddress {
	this := CustomerFullAddress{}
	return &this
}

// NewCustomerFullAddressWithDefaults instantiates a new CustomerFullAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerFullAddressWithDefaults() *CustomerFullAddress {
	this := CustomerFullAddress{}
	return &this
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *CustomerFullAddress) GetCity() string {
	if o == nil || o.City == nil {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerFullAddress) GetCityOk() (*string, bool) {
	if o == nil || o.City == nil {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *CustomerFullAddress) HasCity() bool {
	if o != nil && o.City != nil {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *CustomerFullAddress) SetCity(v string) {
	o.City = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *CustomerFullAddress) GetCountry() string {
	if o == nil || o.Country == nil {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerFullAddress) GetCountryOk() (*string, bool) {
	if o == nil || o.Country == nil {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *CustomerFullAddress) HasCountry() bool {
	if o != nil && o.Country != nil {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *CustomerFullAddress) SetCountry(v string) {
	o.Country = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *CustomerFullAddress) GetPostalCode() string {
	if o == nil || o.PostalCode == nil {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerFullAddress) GetPostalCodeOk() (*string, bool) {
	if o == nil || o.PostalCode == nil {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *CustomerFullAddress) HasPostalCode() bool {
	if o != nil && o.PostalCode != nil {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *CustomerFullAddress) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetStreet returns the Street field value if set, zero value otherwise.
func (o *CustomerFullAddress) GetStreet() string {
	if o == nil || o.Street == nil {
		var ret string
		return ret
	}
	return *o.Street
}

// GetStreetOk returns a tuple with the Street field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerFullAddress) GetStreetOk() (*string, bool) {
	if o == nil || o.Street == nil {
		return nil, false
	}
	return o.Street, true
}

// HasStreet returns a boolean if a field has been set.
func (o *CustomerFullAddress) HasStreet() bool {
	if o != nil && o.Street != nil {
		return true
	}

	return false
}

// SetStreet gets a reference to the given string and assigns it to the Street field.
func (o *CustomerFullAddress) SetStreet(v string) {
	o.Street = &v
}

func (o CustomerFullAddress) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.City != nil {
		toSerialize["city"] = o.City
	}
	if o.Country != nil {
		toSerialize["country"] = o.Country
	}
	if o.PostalCode != nil {
		toSerialize["postal_code"] = o.PostalCode
	}
	if o.Street != nil {
		toSerialize["street"] = o.Street
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerFullAddress struct {
	value *CustomerFullAddress
	isSet bool
}

func (v NullableCustomerFullAddress) Get() *CustomerFullAddress {
	return v.value
}

func (v *NullableCustomerFullAddress) Set(val *CustomerFullAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerFullAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerFullAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerFullAddress(val *CustomerFullAddress) *NullableCustomerFullAddress {
	return &NullableCustomerFullAddress{value: val, isSet: true}
}

func (v NullableCustomerFullAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerFullAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


