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

// CreateESM CreateESM is a command that creates an External Secrets Manager. [Deprecated: Use command create-usc]
type CreateESM struct {
	// Azure Key Vault name (Relevant only for Azure targets)
	AzureKvName *string `json:"azure-kv-name,omitempty"`
	// Protection from accidental deletion of this item [true/false]
	DeleteProtection *string `json:"delete_protection,omitempty"`
	// Description of the External Secrets Manager
	Description *string `json:"description,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// K8s namespace (Relevant to Kubernetes targets)
	K8sNamespace *string `json:"k8s-namespace,omitempty"`
	// External Secrets Manager name
	Name string `json:"name"`
	// List of the tags attached to this External Secrets Manager
	Tags *[]string `json:"tags,omitempty"`
	// Target External Secrets Manager to connect
	TargetToAssociate string `json:"target-to-associate"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewCreateESM instantiates a new CreateESM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateESM(name string, targetToAssociate string, ) *CreateESM {
	this := CreateESM{}
	var json bool = false
	this.Json = &json
	this.Name = name
	this.TargetToAssociate = targetToAssociate
	return &this
}

// NewCreateESMWithDefaults instantiates a new CreateESM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateESMWithDefaults() *CreateESM {
	this := CreateESM{}
	var json bool = false
	this.Json = &json
	return &this
}

// GetAzureKvName returns the AzureKvName field value if set, zero value otherwise.
func (o *CreateESM) GetAzureKvName() string {
	if o == nil || o.AzureKvName == nil {
		var ret string
		return ret
	}
	return *o.AzureKvName
}

// GetAzureKvNameOk returns a tuple with the AzureKvName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateESM) GetAzureKvNameOk() (*string, bool) {
	if o == nil || o.AzureKvName == nil {
		return nil, false
	}
	return o.AzureKvName, true
}

// HasAzureKvName returns a boolean if a field has been set.
func (o *CreateESM) HasAzureKvName() bool {
	if o != nil && o.AzureKvName != nil {
		return true
	}

	return false
}

// SetAzureKvName gets a reference to the given string and assigns it to the AzureKvName field.
func (o *CreateESM) SetAzureKvName(v string) {
	o.AzureKvName = &v
}

// GetDeleteProtection returns the DeleteProtection field value if set, zero value otherwise.
func (o *CreateESM) GetDeleteProtection() string {
	if o == nil || o.DeleteProtection == nil {
		var ret string
		return ret
	}
	return *o.DeleteProtection
}

// GetDeleteProtectionOk returns a tuple with the DeleteProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateESM) GetDeleteProtectionOk() (*string, bool) {
	if o == nil || o.DeleteProtection == nil {
		return nil, false
	}
	return o.DeleteProtection, true
}

// HasDeleteProtection returns a boolean if a field has been set.
func (o *CreateESM) HasDeleteProtection() bool {
	if o != nil && o.DeleteProtection != nil {
		return true
	}

	return false
}

// SetDeleteProtection gets a reference to the given string and assigns it to the DeleteProtection field.
func (o *CreateESM) SetDeleteProtection(v string) {
	o.DeleteProtection = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateESM) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateESM) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateESM) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateESM) SetDescription(v string) {
	o.Description = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *CreateESM) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateESM) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *CreateESM) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *CreateESM) SetJson(v bool) {
	o.Json = &v
}

// GetK8sNamespace returns the K8sNamespace field value if set, zero value otherwise.
func (o *CreateESM) GetK8sNamespace() string {
	if o == nil || o.K8sNamespace == nil {
		var ret string
		return ret
	}
	return *o.K8sNamespace
}

// GetK8sNamespaceOk returns a tuple with the K8sNamespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateESM) GetK8sNamespaceOk() (*string, bool) {
	if o == nil || o.K8sNamespace == nil {
		return nil, false
	}
	return o.K8sNamespace, true
}

// HasK8sNamespace returns a boolean if a field has been set.
func (o *CreateESM) HasK8sNamespace() bool {
	if o != nil && o.K8sNamespace != nil {
		return true
	}

	return false
}

// SetK8sNamespace gets a reference to the given string and assigns it to the K8sNamespace field.
func (o *CreateESM) SetK8sNamespace(v string) {
	o.K8sNamespace = &v
}

// GetName returns the Name field value
func (o *CreateESM) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateESM) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateESM) SetName(v string) {
	o.Name = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *CreateESM) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateESM) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CreateESM) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *CreateESM) SetTags(v []string) {
	o.Tags = &v
}

// GetTargetToAssociate returns the TargetToAssociate field value
func (o *CreateESM) GetTargetToAssociate() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.TargetToAssociate
}

// GetTargetToAssociateOk returns a tuple with the TargetToAssociate field value
// and a boolean to check if the value has been set.
func (o *CreateESM) GetTargetToAssociateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TargetToAssociate, true
}

// SetTargetToAssociate sets field value
func (o *CreateESM) SetTargetToAssociate(v string) {
	o.TargetToAssociate = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *CreateESM) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateESM) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *CreateESM) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *CreateESM) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *CreateESM) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateESM) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *CreateESM) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *CreateESM) SetUidToken(v string) {
	o.UidToken = &v
}

func (o CreateESM) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AzureKvName != nil {
		toSerialize["azure-kv-name"] = o.AzureKvName
	}
	if o.DeleteProtection != nil {
		toSerialize["delete_protection"] = o.DeleteProtection
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if o.K8sNamespace != nil {
		toSerialize["k8s-namespace"] = o.K8sNamespace
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if true {
		toSerialize["target-to-associate"] = o.TargetToAssociate
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableCreateESM struct {
	value *CreateESM
	isSet bool
}

func (v NullableCreateESM) Get() *CreateESM {
	return v.value
}

func (v *NullableCreateESM) Set(val *CreateESM) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateESM) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateESM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateESM(val *CreateESM) *NullableCreateESM {
	return &NullableCreateESM{value: val, isSet: true}
}

func (v NullableCreateESM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateESM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

