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

// SSHCertificateIssueDetails struct for SSHCertificateIssueDetails
type SSHCertificateIssueDetails struct {
	// Relevant for host certificate
	AllowedDomains *[]string `json:"allowed_domains,omitempty"`
	AllowedUserKeyLengths *map[string]int64 `json:"allowed_user_key_lengths,omitempty"`
	// Relevant for user certificate
	AllowedUsers *[]string `json:"allowed_users,omitempty"`
	CertType *int32 `json:"cert_type,omitempty"`
	CriticalOptions *map[string]string `json:"critical_options,omitempty"`
	Extensions *map[string]string `json:"extensions,omitempty"`
	Principals *[]string `json:"principals,omitempty"`
	// In case it is empty, the key ID will be combination of user identifiers and a random string
	StaticKeyId *string `json:"static_key_id,omitempty"`
}

// NewSSHCertificateIssueDetails instantiates a new SSHCertificateIssueDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSSHCertificateIssueDetails() *SSHCertificateIssueDetails {
	this := SSHCertificateIssueDetails{}
	return &this
}

// NewSSHCertificateIssueDetailsWithDefaults instantiates a new SSHCertificateIssueDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSSHCertificateIssueDetailsWithDefaults() *SSHCertificateIssueDetails {
	this := SSHCertificateIssueDetails{}
	return &this
}

// GetAllowedDomains returns the AllowedDomains field value if set, zero value otherwise.
func (o *SSHCertificateIssueDetails) GetAllowedDomains() []string {
	if o == nil || o.AllowedDomains == nil {
		var ret []string
		return ret
	}
	return *o.AllowedDomains
}

// GetAllowedDomainsOk returns a tuple with the AllowedDomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSHCertificateIssueDetails) GetAllowedDomainsOk() (*[]string, bool) {
	if o == nil || o.AllowedDomains == nil {
		return nil, false
	}
	return o.AllowedDomains, true
}

// HasAllowedDomains returns a boolean if a field has been set.
func (o *SSHCertificateIssueDetails) HasAllowedDomains() bool {
	if o != nil && o.AllowedDomains != nil {
		return true
	}

	return false
}

// SetAllowedDomains gets a reference to the given []string and assigns it to the AllowedDomains field.
func (o *SSHCertificateIssueDetails) SetAllowedDomains(v []string) {
	o.AllowedDomains = &v
}

// GetAllowedUserKeyLengths returns the AllowedUserKeyLengths field value if set, zero value otherwise.
func (o *SSHCertificateIssueDetails) GetAllowedUserKeyLengths() map[string]int64 {
	if o == nil || o.AllowedUserKeyLengths == nil {
		var ret map[string]int64
		return ret
	}
	return *o.AllowedUserKeyLengths
}

// GetAllowedUserKeyLengthsOk returns a tuple with the AllowedUserKeyLengths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSHCertificateIssueDetails) GetAllowedUserKeyLengthsOk() (*map[string]int64, bool) {
	if o == nil || o.AllowedUserKeyLengths == nil {
		return nil, false
	}
	return o.AllowedUserKeyLengths, true
}

// HasAllowedUserKeyLengths returns a boolean if a field has been set.
func (o *SSHCertificateIssueDetails) HasAllowedUserKeyLengths() bool {
	if o != nil && o.AllowedUserKeyLengths != nil {
		return true
	}

	return false
}

// SetAllowedUserKeyLengths gets a reference to the given map[string]int64 and assigns it to the AllowedUserKeyLengths field.
func (o *SSHCertificateIssueDetails) SetAllowedUserKeyLengths(v map[string]int64) {
	o.AllowedUserKeyLengths = &v
}

// GetAllowedUsers returns the AllowedUsers field value if set, zero value otherwise.
func (o *SSHCertificateIssueDetails) GetAllowedUsers() []string {
	if o == nil || o.AllowedUsers == nil {
		var ret []string
		return ret
	}
	return *o.AllowedUsers
}

// GetAllowedUsersOk returns a tuple with the AllowedUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSHCertificateIssueDetails) GetAllowedUsersOk() (*[]string, bool) {
	if o == nil || o.AllowedUsers == nil {
		return nil, false
	}
	return o.AllowedUsers, true
}

// HasAllowedUsers returns a boolean if a field has been set.
func (o *SSHCertificateIssueDetails) HasAllowedUsers() bool {
	if o != nil && o.AllowedUsers != nil {
		return true
	}

	return false
}

// SetAllowedUsers gets a reference to the given []string and assigns it to the AllowedUsers field.
func (o *SSHCertificateIssueDetails) SetAllowedUsers(v []string) {
	o.AllowedUsers = &v
}

// GetCertType returns the CertType field value if set, zero value otherwise.
func (o *SSHCertificateIssueDetails) GetCertType() int32 {
	if o == nil || o.CertType == nil {
		var ret int32
		return ret
	}
	return *o.CertType
}

// GetCertTypeOk returns a tuple with the CertType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSHCertificateIssueDetails) GetCertTypeOk() (*int32, bool) {
	if o == nil || o.CertType == nil {
		return nil, false
	}
	return o.CertType, true
}

// HasCertType returns a boolean if a field has been set.
func (o *SSHCertificateIssueDetails) HasCertType() bool {
	if o != nil && o.CertType != nil {
		return true
	}

	return false
}

// SetCertType gets a reference to the given int32 and assigns it to the CertType field.
func (o *SSHCertificateIssueDetails) SetCertType(v int32) {
	o.CertType = &v
}

// GetCriticalOptions returns the CriticalOptions field value if set, zero value otherwise.
func (o *SSHCertificateIssueDetails) GetCriticalOptions() map[string]string {
	if o == nil || o.CriticalOptions == nil {
		var ret map[string]string
		return ret
	}
	return *o.CriticalOptions
}

// GetCriticalOptionsOk returns a tuple with the CriticalOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSHCertificateIssueDetails) GetCriticalOptionsOk() (*map[string]string, bool) {
	if o == nil || o.CriticalOptions == nil {
		return nil, false
	}
	return o.CriticalOptions, true
}

// HasCriticalOptions returns a boolean if a field has been set.
func (o *SSHCertificateIssueDetails) HasCriticalOptions() bool {
	if o != nil && o.CriticalOptions != nil {
		return true
	}

	return false
}

// SetCriticalOptions gets a reference to the given map[string]string and assigns it to the CriticalOptions field.
func (o *SSHCertificateIssueDetails) SetCriticalOptions(v map[string]string) {
	o.CriticalOptions = &v
}

// GetExtensions returns the Extensions field value if set, zero value otherwise.
func (o *SSHCertificateIssueDetails) GetExtensions() map[string]string {
	if o == nil || o.Extensions == nil {
		var ret map[string]string
		return ret
	}
	return *o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSHCertificateIssueDetails) GetExtensionsOk() (*map[string]string, bool) {
	if o == nil || o.Extensions == nil {
		return nil, false
	}
	return o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *SSHCertificateIssueDetails) HasExtensions() bool {
	if o != nil && o.Extensions != nil {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given map[string]string and assigns it to the Extensions field.
func (o *SSHCertificateIssueDetails) SetExtensions(v map[string]string) {
	o.Extensions = &v
}

// GetPrincipals returns the Principals field value if set, zero value otherwise.
func (o *SSHCertificateIssueDetails) GetPrincipals() []string {
	if o == nil || o.Principals == nil {
		var ret []string
		return ret
	}
	return *o.Principals
}

// GetPrincipalsOk returns a tuple with the Principals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSHCertificateIssueDetails) GetPrincipalsOk() (*[]string, bool) {
	if o == nil || o.Principals == nil {
		return nil, false
	}
	return o.Principals, true
}

// HasPrincipals returns a boolean if a field has been set.
func (o *SSHCertificateIssueDetails) HasPrincipals() bool {
	if o != nil && o.Principals != nil {
		return true
	}

	return false
}

// SetPrincipals gets a reference to the given []string and assigns it to the Principals field.
func (o *SSHCertificateIssueDetails) SetPrincipals(v []string) {
	o.Principals = &v
}

// GetStaticKeyId returns the StaticKeyId field value if set, zero value otherwise.
func (o *SSHCertificateIssueDetails) GetStaticKeyId() string {
	if o == nil || o.StaticKeyId == nil {
		var ret string
		return ret
	}
	return *o.StaticKeyId
}

// GetStaticKeyIdOk returns a tuple with the StaticKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSHCertificateIssueDetails) GetStaticKeyIdOk() (*string, bool) {
	if o == nil || o.StaticKeyId == nil {
		return nil, false
	}
	return o.StaticKeyId, true
}

// HasStaticKeyId returns a boolean if a field has been set.
func (o *SSHCertificateIssueDetails) HasStaticKeyId() bool {
	if o != nil && o.StaticKeyId != nil {
		return true
	}

	return false
}

// SetStaticKeyId gets a reference to the given string and assigns it to the StaticKeyId field.
func (o *SSHCertificateIssueDetails) SetStaticKeyId(v string) {
	o.StaticKeyId = &v
}

func (o SSHCertificateIssueDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AllowedDomains != nil {
		toSerialize["allowed_domains"] = o.AllowedDomains
	}
	if o.AllowedUserKeyLengths != nil {
		toSerialize["allowed_user_key_lengths"] = o.AllowedUserKeyLengths
	}
	if o.AllowedUsers != nil {
		toSerialize["allowed_users"] = o.AllowedUsers
	}
	if o.CertType != nil {
		toSerialize["cert_type"] = o.CertType
	}
	if o.CriticalOptions != nil {
		toSerialize["critical_options"] = o.CriticalOptions
	}
	if o.Extensions != nil {
		toSerialize["extensions"] = o.Extensions
	}
	if o.Principals != nil {
		toSerialize["principals"] = o.Principals
	}
	if o.StaticKeyId != nil {
		toSerialize["static_key_id"] = o.StaticKeyId
	}
	return json.Marshal(toSerialize)
}

type NullableSSHCertificateIssueDetails struct {
	value *SSHCertificateIssueDetails
	isSet bool
}

func (v NullableSSHCertificateIssueDetails) Get() *SSHCertificateIssueDetails {
	return v.value
}

func (v *NullableSSHCertificateIssueDetails) Set(val *SSHCertificateIssueDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableSSHCertificateIssueDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableSSHCertificateIssueDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSSHCertificateIssueDetails(val *SSHCertificateIssueDetails) *NullableSSHCertificateIssueDetails {
	return &NullableSSHCertificateIssueDetails{value: val, isSet: true}
}

func (v NullableSSHCertificateIssueDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSSHCertificateIssueDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


