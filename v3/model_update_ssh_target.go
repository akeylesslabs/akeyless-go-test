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

// UpdateSSHTarget struct for UpdateSSHTarget
type UpdateSSHTarget struct {
	// Deprecated - use description
	Comment *string `json:"comment,omitempty"`
	// Description of the object
	Description *string `json:"description,omitempty"`
	Host *string `json:"host,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	KeepPrevVersion *string `json:"keep-prev-version,omitempty"`
	// The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used)
	Key *string `json:"key,omitempty"`
	// Target name
	Name string `json:"name"`
	// New target name
	NewName *string `json:"new-name,omitempty"`
	Port *string `json:"port,omitempty"`
	PrivateKey *string `json:"private-key,omitempty"`
	PrivateKeyPassword *string `json:"private-key-password,omitempty"`
	SshPassword *string `json:"ssh-password,omitempty"`
	SshUsername *string `json:"ssh-username,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	// Deprecated
	UpdateVersion *bool `json:"update-version,omitempty"`
}

// NewUpdateSSHTarget instantiates a new UpdateSSHTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSSHTarget(name string, ) *UpdateSSHTarget {
	this := UpdateSSHTarget{}
	this.Name = name
	return &this
}

// NewUpdateSSHTargetWithDefaults instantiates a new UpdateSSHTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSSHTargetWithDefaults() *UpdateSSHTarget {
	this := UpdateSSHTarget{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *UpdateSSHTarget) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSSHTarget) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *UpdateSSHTarget) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *UpdateSSHTarget) SetComment(v string) {
	o.Comment = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateSSHTarget) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSSHTarget) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateSSHTarget) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateSSHTarget) SetDescription(v string) {
	o.Description = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *UpdateSSHTarget) GetHost() string {
	if o == nil || o.Host == nil {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSSHTarget) GetHostOk() (*string, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *UpdateSSHTarget) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *UpdateSSHTarget) SetHost(v string) {
	o.Host = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *UpdateSSHTarget) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSSHTarget) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *UpdateSSHTarget) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *UpdateSSHTarget) SetJson(v bool) {
	o.Json = &v
}

// GetKeepPrevVersion returns the KeepPrevVersion field value if set, zero value otherwise.
func (o *UpdateSSHTarget) GetKeepPrevVersion() string {
	if o == nil || o.KeepPrevVersion == nil {
		var ret string
		return ret
	}
	return *o.KeepPrevVersion
}

// GetKeepPrevVersionOk returns a tuple with the KeepPrevVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSSHTarget) GetKeepPrevVersionOk() (*string, bool) {
	if o == nil || o.KeepPrevVersion == nil {
		return nil, false
	}
	return o.KeepPrevVersion, true
}

// HasKeepPrevVersion returns a boolean if a field has been set.
func (o *UpdateSSHTarget) HasKeepPrevVersion() bool {
	if o != nil && o.KeepPrevVersion != nil {
		return true
	}

	return false
}

// SetKeepPrevVersion gets a reference to the given string and assigns it to the KeepPrevVersion field.
func (o *UpdateSSHTarget) SetKeepPrevVersion(v string) {
	o.KeepPrevVersion = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *UpdateSSHTarget) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSSHTarget) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *UpdateSSHTarget) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *UpdateSSHTarget) SetKey(v string) {
	o.Key = &v
}

// GetName returns the Name field value
func (o *UpdateSSHTarget) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateSSHTarget) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateSSHTarget) SetName(v string) {
	o.Name = v
}

// GetNewName returns the NewName field value if set, zero value otherwise.
func (o *UpdateSSHTarget) GetNewName() string {
	if o == nil || o.NewName == nil {
		var ret string
		return ret
	}
	return *o.NewName
}

// GetNewNameOk returns a tuple with the NewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSSHTarget) GetNewNameOk() (*string, bool) {
	if o == nil || o.NewName == nil {
		return nil, false
	}
	return o.NewName, true
}

// HasNewName returns a boolean if a field has been set.
func (o *UpdateSSHTarget) HasNewName() bool {
	if o != nil && o.NewName != nil {
		return true
	}

	return false
}

// SetNewName gets a reference to the given string and assigns it to the NewName field.
func (o *UpdateSSHTarget) SetNewName(v string) {
	o.NewName = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *UpdateSSHTarget) GetPort() string {
	if o == nil || o.Port == nil {
		var ret string
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSSHTarget) GetPortOk() (*string, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *UpdateSSHTarget) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given string and assigns it to the Port field.
func (o *UpdateSSHTarget) SetPort(v string) {
	o.Port = &v
}

// GetPrivateKey returns the PrivateKey field value if set, zero value otherwise.
func (o *UpdateSSHTarget) GetPrivateKey() string {
	if o == nil || o.PrivateKey == nil {
		var ret string
		return ret
	}
	return *o.PrivateKey
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSSHTarget) GetPrivateKeyOk() (*string, bool) {
	if o == nil || o.PrivateKey == nil {
		return nil, false
	}
	return o.PrivateKey, true
}

// HasPrivateKey returns a boolean if a field has been set.
func (o *UpdateSSHTarget) HasPrivateKey() bool {
	if o != nil && o.PrivateKey != nil {
		return true
	}

	return false
}

// SetPrivateKey gets a reference to the given string and assigns it to the PrivateKey field.
func (o *UpdateSSHTarget) SetPrivateKey(v string) {
	o.PrivateKey = &v
}

// GetPrivateKeyPassword returns the PrivateKeyPassword field value if set, zero value otherwise.
func (o *UpdateSSHTarget) GetPrivateKeyPassword() string {
	if o == nil || o.PrivateKeyPassword == nil {
		var ret string
		return ret
	}
	return *o.PrivateKeyPassword
}

// GetPrivateKeyPasswordOk returns a tuple with the PrivateKeyPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSSHTarget) GetPrivateKeyPasswordOk() (*string, bool) {
	if o == nil || o.PrivateKeyPassword == nil {
		return nil, false
	}
	return o.PrivateKeyPassword, true
}

// HasPrivateKeyPassword returns a boolean if a field has been set.
func (o *UpdateSSHTarget) HasPrivateKeyPassword() bool {
	if o != nil && o.PrivateKeyPassword != nil {
		return true
	}

	return false
}

// SetPrivateKeyPassword gets a reference to the given string and assigns it to the PrivateKeyPassword field.
func (o *UpdateSSHTarget) SetPrivateKeyPassword(v string) {
	o.PrivateKeyPassword = &v
}

// GetSshPassword returns the SshPassword field value if set, zero value otherwise.
func (o *UpdateSSHTarget) GetSshPassword() string {
	if o == nil || o.SshPassword == nil {
		var ret string
		return ret
	}
	return *o.SshPassword
}

// GetSshPasswordOk returns a tuple with the SshPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSSHTarget) GetSshPasswordOk() (*string, bool) {
	if o == nil || o.SshPassword == nil {
		return nil, false
	}
	return o.SshPassword, true
}

// HasSshPassword returns a boolean if a field has been set.
func (o *UpdateSSHTarget) HasSshPassword() bool {
	if o != nil && o.SshPassword != nil {
		return true
	}

	return false
}

// SetSshPassword gets a reference to the given string and assigns it to the SshPassword field.
func (o *UpdateSSHTarget) SetSshPassword(v string) {
	o.SshPassword = &v
}

// GetSshUsername returns the SshUsername field value if set, zero value otherwise.
func (o *UpdateSSHTarget) GetSshUsername() string {
	if o == nil || o.SshUsername == nil {
		var ret string
		return ret
	}
	return *o.SshUsername
}

// GetSshUsernameOk returns a tuple with the SshUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSSHTarget) GetSshUsernameOk() (*string, bool) {
	if o == nil || o.SshUsername == nil {
		return nil, false
	}
	return o.SshUsername, true
}

// HasSshUsername returns a boolean if a field has been set.
func (o *UpdateSSHTarget) HasSshUsername() bool {
	if o != nil && o.SshUsername != nil {
		return true
	}

	return false
}

// SetSshUsername gets a reference to the given string and assigns it to the SshUsername field.
func (o *UpdateSSHTarget) SetSshUsername(v string) {
	o.SshUsername = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *UpdateSSHTarget) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSSHTarget) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *UpdateSSHTarget) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *UpdateSSHTarget) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *UpdateSSHTarget) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSSHTarget) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *UpdateSSHTarget) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *UpdateSSHTarget) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUpdateVersion returns the UpdateVersion field value if set, zero value otherwise.
func (o *UpdateSSHTarget) GetUpdateVersion() bool {
	if o == nil || o.UpdateVersion == nil {
		var ret bool
		return ret
	}
	return *o.UpdateVersion
}

// GetUpdateVersionOk returns a tuple with the UpdateVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSSHTarget) GetUpdateVersionOk() (*bool, bool) {
	if o == nil || o.UpdateVersion == nil {
		return nil, false
	}
	return o.UpdateVersion, true
}

// HasUpdateVersion returns a boolean if a field has been set.
func (o *UpdateSSHTarget) HasUpdateVersion() bool {
	if o != nil && o.UpdateVersion != nil {
		return true
	}

	return false
}

// SetUpdateVersion gets a reference to the given bool and assigns it to the UpdateVersion field.
func (o *UpdateSSHTarget) SetUpdateVersion(v bool) {
	o.UpdateVersion = &v
}

func (o UpdateSSHTarget) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Comment != nil {
		toSerialize["comment"] = o.Comment
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Host != nil {
		toSerialize["host"] = o.Host
	}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if o.KeepPrevVersion != nil {
		toSerialize["keep-prev-version"] = o.KeepPrevVersion
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.NewName != nil {
		toSerialize["new-name"] = o.NewName
	}
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}
	if o.PrivateKey != nil {
		toSerialize["private-key"] = o.PrivateKey
	}
	if o.PrivateKeyPassword != nil {
		toSerialize["private-key-password"] = o.PrivateKeyPassword
	}
	if o.SshPassword != nil {
		toSerialize["ssh-password"] = o.SshPassword
	}
	if o.SshUsername != nil {
		toSerialize["ssh-username"] = o.SshUsername
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	if o.UpdateVersion != nil {
		toSerialize["update-version"] = o.UpdateVersion
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateSSHTarget struct {
	value *UpdateSSHTarget
	isSet bool
}

func (v NullableUpdateSSHTarget) Get() *UpdateSSHTarget {
	return v.value
}

func (v *NullableUpdateSSHTarget) Set(val *UpdateSSHTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSSHTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSSHTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSSHTarget(val *UpdateSSHTarget) *NullableUpdateSSHTarget {
	return &NullableUpdateSSHTarget{value: val, isSet: true}
}

func (v NullableUpdateSSHTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSSHTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


