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

// GetCloudIdentity getCloudIdentity is a command that gets Cloud Identity Token (relevant only for access-type=azure_ad, aws_iam, gcp).
type GetCloudIdentity struct {
	// Azure Active Directory ObjectId (relevant only for access-type=azure_ad)
	AzureAdObjectId *string `json:"azure_ad_object_id,omitempty"`
	Debug *bool `json:"debug,omitempty"`
	// GCP JWT audience
	GcpAudience *string `json:"gcp-audience,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Escapes the token so it can be safely placed inside a URL query
	UrlSafe *bool `json:"url_safe,omitempty"`
}

// NewGetCloudIdentity instantiates a new GetCloudIdentity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCloudIdentity() *GetCloudIdentity {
	this := GetCloudIdentity{}
	return &this
}

// NewGetCloudIdentityWithDefaults instantiates a new GetCloudIdentity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCloudIdentityWithDefaults() *GetCloudIdentity {
	this := GetCloudIdentity{}
	return &this
}

// GetAzureAdObjectId returns the AzureAdObjectId field value if set, zero value otherwise.
func (o *GetCloudIdentity) GetAzureAdObjectId() string {
	if o == nil || o.AzureAdObjectId == nil {
		var ret string
		return ret
	}
	return *o.AzureAdObjectId
}

// GetAzureAdObjectIdOk returns a tuple with the AzureAdObjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCloudIdentity) GetAzureAdObjectIdOk() (*string, bool) {
	if o == nil || o.AzureAdObjectId == nil {
		return nil, false
	}
	return o.AzureAdObjectId, true
}

// HasAzureAdObjectId returns a boolean if a field has been set.
func (o *GetCloudIdentity) HasAzureAdObjectId() bool {
	if o != nil && o.AzureAdObjectId != nil {
		return true
	}

	return false
}

// SetAzureAdObjectId gets a reference to the given string and assigns it to the AzureAdObjectId field.
func (o *GetCloudIdentity) SetAzureAdObjectId(v string) {
	o.AzureAdObjectId = &v
}

// GetDebug returns the Debug field value if set, zero value otherwise.
func (o *GetCloudIdentity) GetDebug() bool {
	if o == nil || o.Debug == nil {
		var ret bool
		return ret
	}
	return *o.Debug
}

// GetDebugOk returns a tuple with the Debug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCloudIdentity) GetDebugOk() (*bool, bool) {
	if o == nil || o.Debug == nil {
		return nil, false
	}
	return o.Debug, true
}

// HasDebug returns a boolean if a field has been set.
func (o *GetCloudIdentity) HasDebug() bool {
	if o != nil && o.Debug != nil {
		return true
	}

	return false
}

// SetDebug gets a reference to the given bool and assigns it to the Debug field.
func (o *GetCloudIdentity) SetDebug(v bool) {
	o.Debug = &v
}

// GetGcpAudience returns the GcpAudience field value if set, zero value otherwise.
func (o *GetCloudIdentity) GetGcpAudience() string {
	if o == nil || o.GcpAudience == nil {
		var ret string
		return ret
	}
	return *o.GcpAudience
}

// GetGcpAudienceOk returns a tuple with the GcpAudience field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCloudIdentity) GetGcpAudienceOk() (*string, bool) {
	if o == nil || o.GcpAudience == nil {
		return nil, false
	}
	return o.GcpAudience, true
}

// HasGcpAudience returns a boolean if a field has been set.
func (o *GetCloudIdentity) HasGcpAudience() bool {
	if o != nil && o.GcpAudience != nil {
		return true
	}

	return false
}

// SetGcpAudience gets a reference to the given string and assigns it to the GcpAudience field.
func (o *GetCloudIdentity) SetGcpAudience(v string) {
	o.GcpAudience = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *GetCloudIdentity) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCloudIdentity) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *GetCloudIdentity) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *GetCloudIdentity) SetJson(v bool) {
	o.Json = &v
}

// GetUrlSafe returns the UrlSafe field value if set, zero value otherwise.
func (o *GetCloudIdentity) GetUrlSafe() bool {
	if o == nil || o.UrlSafe == nil {
		var ret bool
		return ret
	}
	return *o.UrlSafe
}

// GetUrlSafeOk returns a tuple with the UrlSafe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCloudIdentity) GetUrlSafeOk() (*bool, bool) {
	if o == nil || o.UrlSafe == nil {
		return nil, false
	}
	return o.UrlSafe, true
}

// HasUrlSafe returns a boolean if a field has been set.
func (o *GetCloudIdentity) HasUrlSafe() bool {
	if o != nil && o.UrlSafe != nil {
		return true
	}

	return false
}

// SetUrlSafe gets a reference to the given bool and assigns it to the UrlSafe field.
func (o *GetCloudIdentity) SetUrlSafe(v bool) {
	o.UrlSafe = &v
}

func (o GetCloudIdentity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AzureAdObjectId != nil {
		toSerialize["azure_ad_object_id"] = o.AzureAdObjectId
	}
	if o.Debug != nil {
		toSerialize["debug"] = o.Debug
	}
	if o.GcpAudience != nil {
		toSerialize["gcp-audience"] = o.GcpAudience
	}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if o.UrlSafe != nil {
		toSerialize["url_safe"] = o.UrlSafe
	}
	return json.Marshal(toSerialize)
}

type NullableGetCloudIdentity struct {
	value *GetCloudIdentity
	isSet bool
}

func (v NullableGetCloudIdentity) Get() *GetCloudIdentity {
	return v.value
}

func (v *NullableGetCloudIdentity) Set(val *GetCloudIdentity) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCloudIdentity) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCloudIdentity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCloudIdentity(val *GetCloudIdentity) *NullableGetCloudIdentity {
	return &NullableGetCloudIdentity{value: val, isSet: true}
}

func (v NullableGetCloudIdentity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCloudIdentity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


