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

// UpdateNativeK8STarget struct for UpdateNativeK8STarget
type UpdateNativeK8STarget struct {
	// Deprecated - use description
	Comment *string `json:"comment,omitempty"`
	// Description of the object
	Description *string `json:"description,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// K8S cluster CA certificate
	K8sClusterCaCert string `json:"k8s-cluster-ca-cert"`
	// K8S cluster URL endpoint
	K8sClusterEndpoint string `json:"k8s-cluster-endpoint"`
	// K8S cluster Bearer token
	K8sClusterToken string `json:"k8s-cluster-token"`
	KeepPrevVersion *string `json:"keep-prev-version,omitempty"`
	// The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used)
	Key *string `json:"key,omitempty"`
	// Target name
	Name string `json:"name"`
	// New target name
	NewName *string `json:"new-name,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	// Deprecated
	UpdateVersion *bool `json:"update-version,omitempty"`
}

// NewUpdateNativeK8STarget instantiates a new UpdateNativeK8STarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNativeK8STarget(k8sClusterCaCert string, k8sClusterEndpoint string, k8sClusterToken string, name string, ) *UpdateNativeK8STarget {
	this := UpdateNativeK8STarget{}
	this.K8sClusterCaCert = k8sClusterCaCert
	this.K8sClusterEndpoint = k8sClusterEndpoint
	this.K8sClusterToken = k8sClusterToken
	this.Name = name
	return &this
}

// NewUpdateNativeK8STargetWithDefaults instantiates a new UpdateNativeK8STarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNativeK8STargetWithDefaults() *UpdateNativeK8STarget {
	this := UpdateNativeK8STarget{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *UpdateNativeK8STarget) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNativeK8STarget) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *UpdateNativeK8STarget) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *UpdateNativeK8STarget) SetComment(v string) {
	o.Comment = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateNativeK8STarget) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNativeK8STarget) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateNativeK8STarget) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateNativeK8STarget) SetDescription(v string) {
	o.Description = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *UpdateNativeK8STarget) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNativeK8STarget) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *UpdateNativeK8STarget) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *UpdateNativeK8STarget) SetJson(v bool) {
	o.Json = &v
}

// GetK8sClusterCaCert returns the K8sClusterCaCert field value
func (o *UpdateNativeK8STarget) GetK8sClusterCaCert() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.K8sClusterCaCert
}

// GetK8sClusterCaCertOk returns a tuple with the K8sClusterCaCert field value
// and a boolean to check if the value has been set.
func (o *UpdateNativeK8STarget) GetK8sClusterCaCertOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.K8sClusterCaCert, true
}

// SetK8sClusterCaCert sets field value
func (o *UpdateNativeK8STarget) SetK8sClusterCaCert(v string) {
	o.K8sClusterCaCert = v
}

// GetK8sClusterEndpoint returns the K8sClusterEndpoint field value
func (o *UpdateNativeK8STarget) GetK8sClusterEndpoint() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.K8sClusterEndpoint
}

// GetK8sClusterEndpointOk returns a tuple with the K8sClusterEndpoint field value
// and a boolean to check if the value has been set.
func (o *UpdateNativeK8STarget) GetK8sClusterEndpointOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.K8sClusterEndpoint, true
}

// SetK8sClusterEndpoint sets field value
func (o *UpdateNativeK8STarget) SetK8sClusterEndpoint(v string) {
	o.K8sClusterEndpoint = v
}

// GetK8sClusterToken returns the K8sClusterToken field value
func (o *UpdateNativeK8STarget) GetK8sClusterToken() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.K8sClusterToken
}

// GetK8sClusterTokenOk returns a tuple with the K8sClusterToken field value
// and a boolean to check if the value has been set.
func (o *UpdateNativeK8STarget) GetK8sClusterTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.K8sClusterToken, true
}

// SetK8sClusterToken sets field value
func (o *UpdateNativeK8STarget) SetK8sClusterToken(v string) {
	o.K8sClusterToken = v
}

// GetKeepPrevVersion returns the KeepPrevVersion field value if set, zero value otherwise.
func (o *UpdateNativeK8STarget) GetKeepPrevVersion() string {
	if o == nil || o.KeepPrevVersion == nil {
		var ret string
		return ret
	}
	return *o.KeepPrevVersion
}

// GetKeepPrevVersionOk returns a tuple with the KeepPrevVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNativeK8STarget) GetKeepPrevVersionOk() (*string, bool) {
	if o == nil || o.KeepPrevVersion == nil {
		return nil, false
	}
	return o.KeepPrevVersion, true
}

// HasKeepPrevVersion returns a boolean if a field has been set.
func (o *UpdateNativeK8STarget) HasKeepPrevVersion() bool {
	if o != nil && o.KeepPrevVersion != nil {
		return true
	}

	return false
}

// SetKeepPrevVersion gets a reference to the given string and assigns it to the KeepPrevVersion field.
func (o *UpdateNativeK8STarget) SetKeepPrevVersion(v string) {
	o.KeepPrevVersion = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *UpdateNativeK8STarget) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNativeK8STarget) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *UpdateNativeK8STarget) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *UpdateNativeK8STarget) SetKey(v string) {
	o.Key = &v
}

// GetName returns the Name field value
func (o *UpdateNativeK8STarget) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateNativeK8STarget) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateNativeK8STarget) SetName(v string) {
	o.Name = v
}

// GetNewName returns the NewName field value if set, zero value otherwise.
func (o *UpdateNativeK8STarget) GetNewName() string {
	if o == nil || o.NewName == nil {
		var ret string
		return ret
	}
	return *o.NewName
}

// GetNewNameOk returns a tuple with the NewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNativeK8STarget) GetNewNameOk() (*string, bool) {
	if o == nil || o.NewName == nil {
		return nil, false
	}
	return o.NewName, true
}

// HasNewName returns a boolean if a field has been set.
func (o *UpdateNativeK8STarget) HasNewName() bool {
	if o != nil && o.NewName != nil {
		return true
	}

	return false
}

// SetNewName gets a reference to the given string and assigns it to the NewName field.
func (o *UpdateNativeK8STarget) SetNewName(v string) {
	o.NewName = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *UpdateNativeK8STarget) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNativeK8STarget) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *UpdateNativeK8STarget) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *UpdateNativeK8STarget) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *UpdateNativeK8STarget) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNativeK8STarget) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *UpdateNativeK8STarget) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *UpdateNativeK8STarget) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUpdateVersion returns the UpdateVersion field value if set, zero value otherwise.
func (o *UpdateNativeK8STarget) GetUpdateVersion() bool {
	if o == nil || o.UpdateVersion == nil {
		var ret bool
		return ret
	}
	return *o.UpdateVersion
}

// GetUpdateVersionOk returns a tuple with the UpdateVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNativeK8STarget) GetUpdateVersionOk() (*bool, bool) {
	if o == nil || o.UpdateVersion == nil {
		return nil, false
	}
	return o.UpdateVersion, true
}

// HasUpdateVersion returns a boolean if a field has been set.
func (o *UpdateNativeK8STarget) HasUpdateVersion() bool {
	if o != nil && o.UpdateVersion != nil {
		return true
	}

	return false
}

// SetUpdateVersion gets a reference to the given bool and assigns it to the UpdateVersion field.
func (o *UpdateNativeK8STarget) SetUpdateVersion(v bool) {
	o.UpdateVersion = &v
}

func (o UpdateNativeK8STarget) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Comment != nil {
		toSerialize["comment"] = o.Comment
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if true {
		toSerialize["k8s-cluster-ca-cert"] = o.K8sClusterCaCert
	}
	if true {
		toSerialize["k8s-cluster-endpoint"] = o.K8sClusterEndpoint
	}
	if true {
		toSerialize["k8s-cluster-token"] = o.K8sClusterToken
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

type NullableUpdateNativeK8STarget struct {
	value *UpdateNativeK8STarget
	isSet bool
}

func (v NullableUpdateNativeK8STarget) Get() *UpdateNativeK8STarget {
	return v.value
}

func (v *NullableUpdateNativeK8STarget) Set(val *UpdateNativeK8STarget) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNativeK8STarget) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNativeK8STarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNativeK8STarget(val *UpdateNativeK8STarget) *NullableUpdateNativeK8STarget {
	return &NullableUpdateNativeK8STarget{value: val, isSet: true}
}

func (v NullableUpdateNativeK8STarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNativeK8STarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


