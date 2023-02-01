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

// GatewayUpdateProducerChef gatewayUpdateProducerChef is a command that updates chef producer
type GatewayUpdateProducerChef struct {
	// Organizations
	ChefOrgs *string `json:"chef-orgs,omitempty"`
	// Server key
	ChefServerKey *string `json:"chef-server-key,omitempty"`
	// Server URL
	ChefServerUrl *string `json:"chef-server-url,omitempty"`
	// Server username
	ChefServerUsername *string `json:"chef-server-username,omitempty"`
	// Protection from accidental deletion of this item
	DeleteProtection *string `json:"delete_protection,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Producer name
	Name string `json:"name"`
	// Producer name
	NewName *string `json:"new-name,omitempty"`
	// Dynamic producer encryption key
	ProducerEncryptionKeyName *string `json:"producer-encryption-key-name,omitempty"`
	// Skip SSL
	SkipSsl *bool `json:"skip-ssl,omitempty"`
	// List of the tags attached to this secret
	Tags *[]string `json:"tags,omitempty"`
	// Target name
	TargetName *string `json:"target-name,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	// User TTL
	UserTtl *string `json:"user-ttl,omitempty"`
}

// NewGatewayUpdateProducerChef instantiates a new GatewayUpdateProducerChef object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayUpdateProducerChef(name string, ) *GatewayUpdateProducerChef {
	this := GatewayUpdateProducerChef{}
	this.Name = name
	var skipSsl bool = true
	this.SkipSsl = &skipSsl
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// NewGatewayUpdateProducerChefWithDefaults instantiates a new GatewayUpdateProducerChef object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayUpdateProducerChefWithDefaults() *GatewayUpdateProducerChef {
	this := GatewayUpdateProducerChef{}
	var skipSsl bool = true
	this.SkipSsl = &skipSsl
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// GetChefOrgs returns the ChefOrgs field value if set, zero value otherwise.
func (o *GatewayUpdateProducerChef) GetChefOrgs() string {
	if o == nil || o.ChefOrgs == nil {
		var ret string
		return ret
	}
	return *o.ChefOrgs
}

// GetChefOrgsOk returns a tuple with the ChefOrgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerChef) GetChefOrgsOk() (*string, bool) {
	if o == nil || o.ChefOrgs == nil {
		return nil, false
	}
	return o.ChefOrgs, true
}

// HasChefOrgs returns a boolean if a field has been set.
func (o *GatewayUpdateProducerChef) HasChefOrgs() bool {
	if o != nil && o.ChefOrgs != nil {
		return true
	}

	return false
}

// SetChefOrgs gets a reference to the given string and assigns it to the ChefOrgs field.
func (o *GatewayUpdateProducerChef) SetChefOrgs(v string) {
	o.ChefOrgs = &v
}

// GetChefServerKey returns the ChefServerKey field value if set, zero value otherwise.
func (o *GatewayUpdateProducerChef) GetChefServerKey() string {
	if o == nil || o.ChefServerKey == nil {
		var ret string
		return ret
	}
	return *o.ChefServerKey
}

// GetChefServerKeyOk returns a tuple with the ChefServerKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerChef) GetChefServerKeyOk() (*string, bool) {
	if o == nil || o.ChefServerKey == nil {
		return nil, false
	}
	return o.ChefServerKey, true
}

// HasChefServerKey returns a boolean if a field has been set.
func (o *GatewayUpdateProducerChef) HasChefServerKey() bool {
	if o != nil && o.ChefServerKey != nil {
		return true
	}

	return false
}

// SetChefServerKey gets a reference to the given string and assigns it to the ChefServerKey field.
func (o *GatewayUpdateProducerChef) SetChefServerKey(v string) {
	o.ChefServerKey = &v
}

// GetChefServerUrl returns the ChefServerUrl field value if set, zero value otherwise.
func (o *GatewayUpdateProducerChef) GetChefServerUrl() string {
	if o == nil || o.ChefServerUrl == nil {
		var ret string
		return ret
	}
	return *o.ChefServerUrl
}

// GetChefServerUrlOk returns a tuple with the ChefServerUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerChef) GetChefServerUrlOk() (*string, bool) {
	if o == nil || o.ChefServerUrl == nil {
		return nil, false
	}
	return o.ChefServerUrl, true
}

// HasChefServerUrl returns a boolean if a field has been set.
func (o *GatewayUpdateProducerChef) HasChefServerUrl() bool {
	if o != nil && o.ChefServerUrl != nil {
		return true
	}

	return false
}

// SetChefServerUrl gets a reference to the given string and assigns it to the ChefServerUrl field.
func (o *GatewayUpdateProducerChef) SetChefServerUrl(v string) {
	o.ChefServerUrl = &v
}

// GetChefServerUsername returns the ChefServerUsername field value if set, zero value otherwise.
func (o *GatewayUpdateProducerChef) GetChefServerUsername() string {
	if o == nil || o.ChefServerUsername == nil {
		var ret string
		return ret
	}
	return *o.ChefServerUsername
}

// GetChefServerUsernameOk returns a tuple with the ChefServerUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerChef) GetChefServerUsernameOk() (*string, bool) {
	if o == nil || o.ChefServerUsername == nil {
		return nil, false
	}
	return o.ChefServerUsername, true
}

// HasChefServerUsername returns a boolean if a field has been set.
func (o *GatewayUpdateProducerChef) HasChefServerUsername() bool {
	if o != nil && o.ChefServerUsername != nil {
		return true
	}

	return false
}

// SetChefServerUsername gets a reference to the given string and assigns it to the ChefServerUsername field.
func (o *GatewayUpdateProducerChef) SetChefServerUsername(v string) {
	o.ChefServerUsername = &v
}

// GetDeleteProtection returns the DeleteProtection field value if set, zero value otherwise.
func (o *GatewayUpdateProducerChef) GetDeleteProtection() string {
	if o == nil || o.DeleteProtection == nil {
		var ret string
		return ret
	}
	return *o.DeleteProtection
}

// GetDeleteProtectionOk returns a tuple with the DeleteProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerChef) GetDeleteProtectionOk() (*string, bool) {
	if o == nil || o.DeleteProtection == nil {
		return nil, false
	}
	return o.DeleteProtection, true
}

// HasDeleteProtection returns a boolean if a field has been set.
func (o *GatewayUpdateProducerChef) HasDeleteProtection() bool {
	if o != nil && o.DeleteProtection != nil {
		return true
	}

	return false
}

// SetDeleteProtection gets a reference to the given string and assigns it to the DeleteProtection field.
func (o *GatewayUpdateProducerChef) SetDeleteProtection(v string) {
	o.DeleteProtection = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *GatewayUpdateProducerChef) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerChef) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *GatewayUpdateProducerChef) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *GatewayUpdateProducerChef) SetJson(v bool) {
	o.Json = &v
}

// GetName returns the Name field value
func (o *GatewayUpdateProducerChef) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerChef) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GatewayUpdateProducerChef) SetName(v string) {
	o.Name = v
}

// GetNewName returns the NewName field value if set, zero value otherwise.
func (o *GatewayUpdateProducerChef) GetNewName() string {
	if o == nil || o.NewName == nil {
		var ret string
		return ret
	}
	return *o.NewName
}

// GetNewNameOk returns a tuple with the NewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerChef) GetNewNameOk() (*string, bool) {
	if o == nil || o.NewName == nil {
		return nil, false
	}
	return o.NewName, true
}

// HasNewName returns a boolean if a field has been set.
func (o *GatewayUpdateProducerChef) HasNewName() bool {
	if o != nil && o.NewName != nil {
		return true
	}

	return false
}

// SetNewName gets a reference to the given string and assigns it to the NewName field.
func (o *GatewayUpdateProducerChef) SetNewName(v string) {
	o.NewName = &v
}

// GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field value if set, zero value otherwise.
func (o *GatewayUpdateProducerChef) GetProducerEncryptionKeyName() string {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		var ret string
		return ret
	}
	return *o.ProducerEncryptionKeyName
}

// GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerChef) GetProducerEncryptionKeyNameOk() (*string, bool) {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		return nil, false
	}
	return o.ProducerEncryptionKeyName, true
}

// HasProducerEncryptionKeyName returns a boolean if a field has been set.
func (o *GatewayUpdateProducerChef) HasProducerEncryptionKeyName() bool {
	if o != nil && o.ProducerEncryptionKeyName != nil {
		return true
	}

	return false
}

// SetProducerEncryptionKeyName gets a reference to the given string and assigns it to the ProducerEncryptionKeyName field.
func (o *GatewayUpdateProducerChef) SetProducerEncryptionKeyName(v string) {
	o.ProducerEncryptionKeyName = &v
}

// GetSkipSsl returns the SkipSsl field value if set, zero value otherwise.
func (o *GatewayUpdateProducerChef) GetSkipSsl() bool {
	if o == nil || o.SkipSsl == nil {
		var ret bool
		return ret
	}
	return *o.SkipSsl
}

// GetSkipSslOk returns a tuple with the SkipSsl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerChef) GetSkipSslOk() (*bool, bool) {
	if o == nil || o.SkipSsl == nil {
		return nil, false
	}
	return o.SkipSsl, true
}

// HasSkipSsl returns a boolean if a field has been set.
func (o *GatewayUpdateProducerChef) HasSkipSsl() bool {
	if o != nil && o.SkipSsl != nil {
		return true
	}

	return false
}

// SetSkipSsl gets a reference to the given bool and assigns it to the SkipSsl field.
func (o *GatewayUpdateProducerChef) SetSkipSsl(v bool) {
	o.SkipSsl = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *GatewayUpdateProducerChef) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerChef) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *GatewayUpdateProducerChef) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *GatewayUpdateProducerChef) SetTags(v []string) {
	o.Tags = &v
}

// GetTargetName returns the TargetName field value if set, zero value otherwise.
func (o *GatewayUpdateProducerChef) GetTargetName() string {
	if o == nil || o.TargetName == nil {
		var ret string
		return ret
	}
	return *o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerChef) GetTargetNameOk() (*string, bool) {
	if o == nil || o.TargetName == nil {
		return nil, false
	}
	return o.TargetName, true
}

// HasTargetName returns a boolean if a field has been set.
func (o *GatewayUpdateProducerChef) HasTargetName() bool {
	if o != nil && o.TargetName != nil {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given string and assigns it to the TargetName field.
func (o *GatewayUpdateProducerChef) SetTargetName(v string) {
	o.TargetName = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GatewayUpdateProducerChef) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerChef) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GatewayUpdateProducerChef) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GatewayUpdateProducerChef) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *GatewayUpdateProducerChef) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerChef) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *GatewayUpdateProducerChef) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *GatewayUpdateProducerChef) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUserTtl returns the UserTtl field value if set, zero value otherwise.
func (o *GatewayUpdateProducerChef) GetUserTtl() string {
	if o == nil || o.UserTtl == nil {
		var ret string
		return ret
	}
	return *o.UserTtl
}

// GetUserTtlOk returns a tuple with the UserTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerChef) GetUserTtlOk() (*string, bool) {
	if o == nil || o.UserTtl == nil {
		return nil, false
	}
	return o.UserTtl, true
}

// HasUserTtl returns a boolean if a field has been set.
func (o *GatewayUpdateProducerChef) HasUserTtl() bool {
	if o != nil && o.UserTtl != nil {
		return true
	}

	return false
}

// SetUserTtl gets a reference to the given string and assigns it to the UserTtl field.
func (o *GatewayUpdateProducerChef) SetUserTtl(v string) {
	o.UserTtl = &v
}

func (o GatewayUpdateProducerChef) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ChefOrgs != nil {
		toSerialize["chef-orgs"] = o.ChefOrgs
	}
	if o.ChefServerKey != nil {
		toSerialize["chef-server-key"] = o.ChefServerKey
	}
	if o.ChefServerUrl != nil {
		toSerialize["chef-server-url"] = o.ChefServerUrl
	}
	if o.ChefServerUsername != nil {
		toSerialize["chef-server-username"] = o.ChefServerUsername
	}
	if o.DeleteProtection != nil {
		toSerialize["delete_protection"] = o.DeleteProtection
	}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.NewName != nil {
		toSerialize["new-name"] = o.NewName
	}
	if o.ProducerEncryptionKeyName != nil {
		toSerialize["producer-encryption-key-name"] = o.ProducerEncryptionKeyName
	}
	if o.SkipSsl != nil {
		toSerialize["skip-ssl"] = o.SkipSsl
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.TargetName != nil {
		toSerialize["target-name"] = o.TargetName
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	if o.UserTtl != nil {
		toSerialize["user-ttl"] = o.UserTtl
	}
	return json.Marshal(toSerialize)
}

type NullableGatewayUpdateProducerChef struct {
	value *GatewayUpdateProducerChef
	isSet bool
}

func (v NullableGatewayUpdateProducerChef) Get() *GatewayUpdateProducerChef {
	return v.value
}

func (v *NullableGatewayUpdateProducerChef) Set(val *GatewayUpdateProducerChef) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayUpdateProducerChef) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayUpdateProducerChef) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayUpdateProducerChef(val *GatewayUpdateProducerChef) *NullableGatewayUpdateProducerChef {
	return &NullableGatewayUpdateProducerChef{value: val, isSet: true}
}

func (v NullableGatewayUpdateProducerChef) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayUpdateProducerChef) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


