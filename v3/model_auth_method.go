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
	"time"
)

// AuthMethod struct for AuthMethod
type AuthMethod struct {
	AccessDate *time.Time `json:"access_date,omitempty"`
	AccessInfo *AuthMethodAccessInfo `json:"access_info,omitempty"`
	AccountId *string `json:"account_id,omitempty"`
	AuthMethodAccessId *string `json:"auth_method_access_id,omitempty"`
	AuthMethodName *string `json:"auth_method_name,omitempty"`
	AuthMethodRolesAssoc *[]AuthMethodRoleAssociation `json:"auth_method_roles_assoc,omitempty"`
	ClientPermissions *[]string `json:"client_permissions,omitempty"`
	CreationDate *time.Time `json:"creation_date,omitempty"`
	ModificationDate *time.Time `json:"modification_date,omitempty"`
}

// NewAuthMethod instantiates a new AuthMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthMethod() *AuthMethod {
	this := AuthMethod{}
	return &this
}

// NewAuthMethodWithDefaults instantiates a new AuthMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthMethodWithDefaults() *AuthMethod {
	this := AuthMethod{}
	return &this
}

// GetAccessDate returns the AccessDate field value if set, zero value otherwise.
func (o *AuthMethod) GetAccessDate() time.Time {
	if o == nil || o.AccessDate == nil {
		var ret time.Time
		return ret
	}
	return *o.AccessDate
}

// GetAccessDateOk returns a tuple with the AccessDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthMethod) GetAccessDateOk() (*time.Time, bool) {
	if o == nil || o.AccessDate == nil {
		return nil, false
	}
	return o.AccessDate, true
}

// HasAccessDate returns a boolean if a field has been set.
func (o *AuthMethod) HasAccessDate() bool {
	if o != nil && o.AccessDate != nil {
		return true
	}

	return false
}

// SetAccessDate gets a reference to the given time.Time and assigns it to the AccessDate field.
func (o *AuthMethod) SetAccessDate(v time.Time) {
	o.AccessDate = &v
}

// GetAccessInfo returns the AccessInfo field value if set, zero value otherwise.
func (o *AuthMethod) GetAccessInfo() AuthMethodAccessInfo {
	if o == nil || o.AccessInfo == nil {
		var ret AuthMethodAccessInfo
		return ret
	}
	return *o.AccessInfo
}

// GetAccessInfoOk returns a tuple with the AccessInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthMethod) GetAccessInfoOk() (*AuthMethodAccessInfo, bool) {
	if o == nil || o.AccessInfo == nil {
		return nil, false
	}
	return o.AccessInfo, true
}

// HasAccessInfo returns a boolean if a field has been set.
func (o *AuthMethod) HasAccessInfo() bool {
	if o != nil && o.AccessInfo != nil {
		return true
	}

	return false
}

// SetAccessInfo gets a reference to the given AuthMethodAccessInfo and assigns it to the AccessInfo field.
func (o *AuthMethod) SetAccessInfo(v AuthMethodAccessInfo) {
	o.AccessInfo = &v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *AuthMethod) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthMethod) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *AuthMethod) HasAccountId() bool {
	if o != nil && o.AccountId != nil {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *AuthMethod) SetAccountId(v string) {
	o.AccountId = &v
}

// GetAuthMethodAccessId returns the AuthMethodAccessId field value if set, zero value otherwise.
func (o *AuthMethod) GetAuthMethodAccessId() string {
	if o == nil || o.AuthMethodAccessId == nil {
		var ret string
		return ret
	}
	return *o.AuthMethodAccessId
}

// GetAuthMethodAccessIdOk returns a tuple with the AuthMethodAccessId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthMethod) GetAuthMethodAccessIdOk() (*string, bool) {
	if o == nil || o.AuthMethodAccessId == nil {
		return nil, false
	}
	return o.AuthMethodAccessId, true
}

// HasAuthMethodAccessId returns a boolean if a field has been set.
func (o *AuthMethod) HasAuthMethodAccessId() bool {
	if o != nil && o.AuthMethodAccessId != nil {
		return true
	}

	return false
}

// SetAuthMethodAccessId gets a reference to the given string and assigns it to the AuthMethodAccessId field.
func (o *AuthMethod) SetAuthMethodAccessId(v string) {
	o.AuthMethodAccessId = &v
}

// GetAuthMethodName returns the AuthMethodName field value if set, zero value otherwise.
func (o *AuthMethod) GetAuthMethodName() string {
	if o == nil || o.AuthMethodName == nil {
		var ret string
		return ret
	}
	return *o.AuthMethodName
}

// GetAuthMethodNameOk returns a tuple with the AuthMethodName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthMethod) GetAuthMethodNameOk() (*string, bool) {
	if o == nil || o.AuthMethodName == nil {
		return nil, false
	}
	return o.AuthMethodName, true
}

// HasAuthMethodName returns a boolean if a field has been set.
func (o *AuthMethod) HasAuthMethodName() bool {
	if o != nil && o.AuthMethodName != nil {
		return true
	}

	return false
}

// SetAuthMethodName gets a reference to the given string and assigns it to the AuthMethodName field.
func (o *AuthMethod) SetAuthMethodName(v string) {
	o.AuthMethodName = &v
}

// GetAuthMethodRolesAssoc returns the AuthMethodRolesAssoc field value if set, zero value otherwise.
func (o *AuthMethod) GetAuthMethodRolesAssoc() []AuthMethodRoleAssociation {
	if o == nil || o.AuthMethodRolesAssoc == nil {
		var ret []AuthMethodRoleAssociation
		return ret
	}
	return *o.AuthMethodRolesAssoc
}

// GetAuthMethodRolesAssocOk returns a tuple with the AuthMethodRolesAssoc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthMethod) GetAuthMethodRolesAssocOk() (*[]AuthMethodRoleAssociation, bool) {
	if o == nil || o.AuthMethodRolesAssoc == nil {
		return nil, false
	}
	return o.AuthMethodRolesAssoc, true
}

// HasAuthMethodRolesAssoc returns a boolean if a field has been set.
func (o *AuthMethod) HasAuthMethodRolesAssoc() bool {
	if o != nil && o.AuthMethodRolesAssoc != nil {
		return true
	}

	return false
}

// SetAuthMethodRolesAssoc gets a reference to the given []AuthMethodRoleAssociation and assigns it to the AuthMethodRolesAssoc field.
func (o *AuthMethod) SetAuthMethodRolesAssoc(v []AuthMethodRoleAssociation) {
	o.AuthMethodRolesAssoc = &v
}

// GetClientPermissions returns the ClientPermissions field value if set, zero value otherwise.
func (o *AuthMethod) GetClientPermissions() []string {
	if o == nil || o.ClientPermissions == nil {
		var ret []string
		return ret
	}
	return *o.ClientPermissions
}

// GetClientPermissionsOk returns a tuple with the ClientPermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthMethod) GetClientPermissionsOk() (*[]string, bool) {
	if o == nil || o.ClientPermissions == nil {
		return nil, false
	}
	return o.ClientPermissions, true
}

// HasClientPermissions returns a boolean if a field has been set.
func (o *AuthMethod) HasClientPermissions() bool {
	if o != nil && o.ClientPermissions != nil {
		return true
	}

	return false
}

// SetClientPermissions gets a reference to the given []string and assigns it to the ClientPermissions field.
func (o *AuthMethod) SetClientPermissions(v []string) {
	o.ClientPermissions = &v
}

// GetCreationDate returns the CreationDate field value if set, zero value otherwise.
func (o *AuthMethod) GetCreationDate() time.Time {
	if o == nil || o.CreationDate == nil {
		var ret time.Time
		return ret
	}
	return *o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthMethod) GetCreationDateOk() (*time.Time, bool) {
	if o == nil || o.CreationDate == nil {
		return nil, false
	}
	return o.CreationDate, true
}

// HasCreationDate returns a boolean if a field has been set.
func (o *AuthMethod) HasCreationDate() bool {
	if o != nil && o.CreationDate != nil {
		return true
	}

	return false
}

// SetCreationDate gets a reference to the given time.Time and assigns it to the CreationDate field.
func (o *AuthMethod) SetCreationDate(v time.Time) {
	o.CreationDate = &v
}

// GetModificationDate returns the ModificationDate field value if set, zero value otherwise.
func (o *AuthMethod) GetModificationDate() time.Time {
	if o == nil || o.ModificationDate == nil {
		var ret time.Time
		return ret
	}
	return *o.ModificationDate
}

// GetModificationDateOk returns a tuple with the ModificationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthMethod) GetModificationDateOk() (*time.Time, bool) {
	if o == nil || o.ModificationDate == nil {
		return nil, false
	}
	return o.ModificationDate, true
}

// HasModificationDate returns a boolean if a field has been set.
func (o *AuthMethod) HasModificationDate() bool {
	if o != nil && o.ModificationDate != nil {
		return true
	}

	return false
}

// SetModificationDate gets a reference to the given time.Time and assigns it to the ModificationDate field.
func (o *AuthMethod) SetModificationDate(v time.Time) {
	o.ModificationDate = &v
}

func (o AuthMethod) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessDate != nil {
		toSerialize["access_date"] = o.AccessDate
	}
	if o.AccessInfo != nil {
		toSerialize["access_info"] = o.AccessInfo
	}
	if o.AccountId != nil {
		toSerialize["account_id"] = o.AccountId
	}
	if o.AuthMethodAccessId != nil {
		toSerialize["auth_method_access_id"] = o.AuthMethodAccessId
	}
	if o.AuthMethodName != nil {
		toSerialize["auth_method_name"] = o.AuthMethodName
	}
	if o.AuthMethodRolesAssoc != nil {
		toSerialize["auth_method_roles_assoc"] = o.AuthMethodRolesAssoc
	}
	if o.ClientPermissions != nil {
		toSerialize["client_permissions"] = o.ClientPermissions
	}
	if o.CreationDate != nil {
		toSerialize["creation_date"] = o.CreationDate
	}
	if o.ModificationDate != nil {
		toSerialize["modification_date"] = o.ModificationDate
	}
	return json.Marshal(toSerialize)
}

type NullableAuthMethod struct {
	value *AuthMethod
	isSet bool
}

func (v NullableAuthMethod) Get() *AuthMethod {
	return v.value
}

func (v *NullableAuthMethod) Set(val *AuthMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthMethod(val *AuthMethod) *NullableAuthMethod {
	return &NullableAuthMethod{value: val, isSet: true}
}

func (v NullableAuthMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


