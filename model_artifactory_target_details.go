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

// ArtifactoryTargetDetails struct for ArtifactoryTargetDetails
type ArtifactoryTargetDetails struct {
	ArtifactoryAdminApikey *string `json:"artifactory_admin_apikey,omitempty"`
	ArtifactoryAdminUsername *string `json:"artifactory_admin_username,omitempty"`
	ArtifactoryBaseUrl *string `json:"artifactory_base_url,omitempty"`
}

// NewArtifactoryTargetDetails instantiates a new ArtifactoryTargetDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArtifactoryTargetDetails() *ArtifactoryTargetDetails {
	this := ArtifactoryTargetDetails{}
	return &this
}

// NewArtifactoryTargetDetailsWithDefaults instantiates a new ArtifactoryTargetDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArtifactoryTargetDetailsWithDefaults() *ArtifactoryTargetDetails {
	this := ArtifactoryTargetDetails{}
	return &this
}

// GetArtifactoryAdminApikey returns the ArtifactoryAdminApikey field value if set, zero value otherwise.
func (o *ArtifactoryTargetDetails) GetArtifactoryAdminApikey() string {
	if o == nil || o.ArtifactoryAdminApikey == nil {
		var ret string
		return ret
	}
	return *o.ArtifactoryAdminApikey
}

// GetArtifactoryAdminApikeyOk returns a tuple with the ArtifactoryAdminApikey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactoryTargetDetails) GetArtifactoryAdminApikeyOk() (*string, bool) {
	if o == nil || o.ArtifactoryAdminApikey == nil {
		return nil, false
	}
	return o.ArtifactoryAdminApikey, true
}

// HasArtifactoryAdminApikey returns a boolean if a field has been set.
func (o *ArtifactoryTargetDetails) HasArtifactoryAdminApikey() bool {
	if o != nil && o.ArtifactoryAdminApikey != nil {
		return true
	}

	return false
}

// SetArtifactoryAdminApikey gets a reference to the given string and assigns it to the ArtifactoryAdminApikey field.
func (o *ArtifactoryTargetDetails) SetArtifactoryAdminApikey(v string) {
	o.ArtifactoryAdminApikey = &v
}

// GetArtifactoryAdminUsername returns the ArtifactoryAdminUsername field value if set, zero value otherwise.
func (o *ArtifactoryTargetDetails) GetArtifactoryAdminUsername() string {
	if o == nil || o.ArtifactoryAdminUsername == nil {
		var ret string
		return ret
	}
	return *o.ArtifactoryAdminUsername
}

// GetArtifactoryAdminUsernameOk returns a tuple with the ArtifactoryAdminUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactoryTargetDetails) GetArtifactoryAdminUsernameOk() (*string, bool) {
	if o == nil || o.ArtifactoryAdminUsername == nil {
		return nil, false
	}
	return o.ArtifactoryAdminUsername, true
}

// HasArtifactoryAdminUsername returns a boolean if a field has been set.
func (o *ArtifactoryTargetDetails) HasArtifactoryAdminUsername() bool {
	if o != nil && o.ArtifactoryAdminUsername != nil {
		return true
	}

	return false
}

// SetArtifactoryAdminUsername gets a reference to the given string and assigns it to the ArtifactoryAdminUsername field.
func (o *ArtifactoryTargetDetails) SetArtifactoryAdminUsername(v string) {
	o.ArtifactoryAdminUsername = &v
}

// GetArtifactoryBaseUrl returns the ArtifactoryBaseUrl field value if set, zero value otherwise.
func (o *ArtifactoryTargetDetails) GetArtifactoryBaseUrl() string {
	if o == nil || o.ArtifactoryBaseUrl == nil {
		var ret string
		return ret
	}
	return *o.ArtifactoryBaseUrl
}

// GetArtifactoryBaseUrlOk returns a tuple with the ArtifactoryBaseUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactoryTargetDetails) GetArtifactoryBaseUrlOk() (*string, bool) {
	if o == nil || o.ArtifactoryBaseUrl == nil {
		return nil, false
	}
	return o.ArtifactoryBaseUrl, true
}

// HasArtifactoryBaseUrl returns a boolean if a field has been set.
func (o *ArtifactoryTargetDetails) HasArtifactoryBaseUrl() bool {
	if o != nil && o.ArtifactoryBaseUrl != nil {
		return true
	}

	return false
}

// SetArtifactoryBaseUrl gets a reference to the given string and assigns it to the ArtifactoryBaseUrl field.
func (o *ArtifactoryTargetDetails) SetArtifactoryBaseUrl(v string) {
	o.ArtifactoryBaseUrl = &v
}

func (o ArtifactoryTargetDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ArtifactoryAdminApikey != nil {
		toSerialize["artifactory_admin_apikey"] = o.ArtifactoryAdminApikey
	}
	if o.ArtifactoryAdminUsername != nil {
		toSerialize["artifactory_admin_username"] = o.ArtifactoryAdminUsername
	}
	if o.ArtifactoryBaseUrl != nil {
		toSerialize["artifactory_base_url"] = o.ArtifactoryBaseUrl
	}
	return json.Marshal(toSerialize)
}

type NullableArtifactoryTargetDetails struct {
	value *ArtifactoryTargetDetails
	isSet bool
}

func (v NullableArtifactoryTargetDetails) Get() *ArtifactoryTargetDetails {
	return v.value
}

func (v *NullableArtifactoryTargetDetails) Set(val *ArtifactoryTargetDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableArtifactoryTargetDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableArtifactoryTargetDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArtifactoryTargetDetails(val *ArtifactoryTargetDetails) *NullableArtifactoryTargetDetails {
	return &NullableArtifactoryTargetDetails{value: val, isSet: true}
}

func (v NullableArtifactoryTargetDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArtifactoryTargetDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

