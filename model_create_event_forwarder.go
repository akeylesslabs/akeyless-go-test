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

// CreateEventForwarder struct for CreateEventForwarder
type CreateEventForwarder struct {
	// Workstation Admin Name
	AdminName *string `json:"admin-name,omitempty"`
	// Workstation Admin password
	AdminPwd *string `json:"admin-pwd,omitempty"`
	// Deprecated - use description
	Comment *string `json:"comment,omitempty"`
	// Description of the object
	Description *string `json:"description,omitempty"`
	// A comma seperated list of email addresses to send event to (relevant only for \\\"email\\\" Event Forwarder)
	EmailTo *string `json:"email-to,omitempty"`
	// Event sources
	EventSourceLocations []string `json:"event-source-locations"`
	// Event Source type [item, target]
	EventSourceType *string `json:"event-source-type,omitempty"`
	// Event types
	EventTypes *[]string `json:"event-types,omitempty"`
	// Rate of periodic runner repetition in hours
	Every *string `json:"every,omitempty"`
	ForwarderType string `json:"forwarder-type"`
	// Workstation Host
	Host *string `json:"host,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// The name of a key that used to encrypt the EventForwarder secret value (if empty, the account default protectionKey key will be used)
	Key *string `json:"key,omitempty"`
	// EventForwarder name
	Name string `json:"name"`
	RunnerType string `json:"runner-type"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewCreateEventForwarder instantiates a new CreateEventForwarder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateEventForwarder(eventSourceLocations []string, forwarderType string, name string, runnerType string, ) *CreateEventForwarder {
	this := CreateEventForwarder{}
	this.EventSourceLocations = eventSourceLocations
	var eventSourceType string = "item"
	this.EventSourceType = &eventSourceType
	this.ForwarderType = forwarderType
	this.Name = name
	this.RunnerType = runnerType
	return &this
}

// NewCreateEventForwarderWithDefaults instantiates a new CreateEventForwarder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateEventForwarderWithDefaults() *CreateEventForwarder {
	this := CreateEventForwarder{}
	var eventSourceType string = "item"
	this.EventSourceType = &eventSourceType
	return &this
}

// GetAdminName returns the AdminName field value if set, zero value otherwise.
func (o *CreateEventForwarder) GetAdminName() string {
	if o == nil || o.AdminName == nil {
		var ret string
		return ret
	}
	return *o.AdminName
}

// GetAdminNameOk returns a tuple with the AdminName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEventForwarder) GetAdminNameOk() (*string, bool) {
	if o == nil || o.AdminName == nil {
		return nil, false
	}
	return o.AdminName, true
}

// HasAdminName returns a boolean if a field has been set.
func (o *CreateEventForwarder) HasAdminName() bool {
	if o != nil && o.AdminName != nil {
		return true
	}

	return false
}

// SetAdminName gets a reference to the given string and assigns it to the AdminName field.
func (o *CreateEventForwarder) SetAdminName(v string) {
	o.AdminName = &v
}

// GetAdminPwd returns the AdminPwd field value if set, zero value otherwise.
func (o *CreateEventForwarder) GetAdminPwd() string {
	if o == nil || o.AdminPwd == nil {
		var ret string
		return ret
	}
	return *o.AdminPwd
}

// GetAdminPwdOk returns a tuple with the AdminPwd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEventForwarder) GetAdminPwdOk() (*string, bool) {
	if o == nil || o.AdminPwd == nil {
		return nil, false
	}
	return o.AdminPwd, true
}

// HasAdminPwd returns a boolean if a field has been set.
func (o *CreateEventForwarder) HasAdminPwd() bool {
	if o != nil && o.AdminPwd != nil {
		return true
	}

	return false
}

// SetAdminPwd gets a reference to the given string and assigns it to the AdminPwd field.
func (o *CreateEventForwarder) SetAdminPwd(v string) {
	o.AdminPwd = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *CreateEventForwarder) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEventForwarder) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *CreateEventForwarder) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *CreateEventForwarder) SetComment(v string) {
	o.Comment = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateEventForwarder) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEventForwarder) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateEventForwarder) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateEventForwarder) SetDescription(v string) {
	o.Description = &v
}

// GetEmailTo returns the EmailTo field value if set, zero value otherwise.
func (o *CreateEventForwarder) GetEmailTo() string {
	if o == nil || o.EmailTo == nil {
		var ret string
		return ret
	}
	return *o.EmailTo
}

// GetEmailToOk returns a tuple with the EmailTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEventForwarder) GetEmailToOk() (*string, bool) {
	if o == nil || o.EmailTo == nil {
		return nil, false
	}
	return o.EmailTo, true
}

// HasEmailTo returns a boolean if a field has been set.
func (o *CreateEventForwarder) HasEmailTo() bool {
	if o != nil && o.EmailTo != nil {
		return true
	}

	return false
}

// SetEmailTo gets a reference to the given string and assigns it to the EmailTo field.
func (o *CreateEventForwarder) SetEmailTo(v string) {
	o.EmailTo = &v
}

// GetEventSourceLocations returns the EventSourceLocations field value
func (o *CreateEventForwarder) GetEventSourceLocations() []string {
	if o == nil  {
		var ret []string
		return ret
	}

	return o.EventSourceLocations
}

// GetEventSourceLocationsOk returns a tuple with the EventSourceLocations field value
// and a boolean to check if the value has been set.
func (o *CreateEventForwarder) GetEventSourceLocationsOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EventSourceLocations, true
}

// SetEventSourceLocations sets field value
func (o *CreateEventForwarder) SetEventSourceLocations(v []string) {
	o.EventSourceLocations = v
}

// GetEventSourceType returns the EventSourceType field value if set, zero value otherwise.
func (o *CreateEventForwarder) GetEventSourceType() string {
	if o == nil || o.EventSourceType == nil {
		var ret string
		return ret
	}
	return *o.EventSourceType
}

// GetEventSourceTypeOk returns a tuple with the EventSourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEventForwarder) GetEventSourceTypeOk() (*string, bool) {
	if o == nil || o.EventSourceType == nil {
		return nil, false
	}
	return o.EventSourceType, true
}

// HasEventSourceType returns a boolean if a field has been set.
func (o *CreateEventForwarder) HasEventSourceType() bool {
	if o != nil && o.EventSourceType != nil {
		return true
	}

	return false
}

// SetEventSourceType gets a reference to the given string and assigns it to the EventSourceType field.
func (o *CreateEventForwarder) SetEventSourceType(v string) {
	o.EventSourceType = &v
}

// GetEventTypes returns the EventTypes field value if set, zero value otherwise.
func (o *CreateEventForwarder) GetEventTypes() []string {
	if o == nil || o.EventTypes == nil {
		var ret []string
		return ret
	}
	return *o.EventTypes
}

// GetEventTypesOk returns a tuple with the EventTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEventForwarder) GetEventTypesOk() (*[]string, bool) {
	if o == nil || o.EventTypes == nil {
		return nil, false
	}
	return o.EventTypes, true
}

// HasEventTypes returns a boolean if a field has been set.
func (o *CreateEventForwarder) HasEventTypes() bool {
	if o != nil && o.EventTypes != nil {
		return true
	}

	return false
}

// SetEventTypes gets a reference to the given []string and assigns it to the EventTypes field.
func (o *CreateEventForwarder) SetEventTypes(v []string) {
	o.EventTypes = &v
}

// GetEvery returns the Every field value if set, zero value otherwise.
func (o *CreateEventForwarder) GetEvery() string {
	if o == nil || o.Every == nil {
		var ret string
		return ret
	}
	return *o.Every
}

// GetEveryOk returns a tuple with the Every field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEventForwarder) GetEveryOk() (*string, bool) {
	if o == nil || o.Every == nil {
		return nil, false
	}
	return o.Every, true
}

// HasEvery returns a boolean if a field has been set.
func (o *CreateEventForwarder) HasEvery() bool {
	if o != nil && o.Every != nil {
		return true
	}

	return false
}

// SetEvery gets a reference to the given string and assigns it to the Every field.
func (o *CreateEventForwarder) SetEvery(v string) {
	o.Every = &v
}

// GetForwarderType returns the ForwarderType field value
func (o *CreateEventForwarder) GetForwarderType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ForwarderType
}

// GetForwarderTypeOk returns a tuple with the ForwarderType field value
// and a boolean to check if the value has been set.
func (o *CreateEventForwarder) GetForwarderTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ForwarderType, true
}

// SetForwarderType sets field value
func (o *CreateEventForwarder) SetForwarderType(v string) {
	o.ForwarderType = v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *CreateEventForwarder) GetHost() string {
	if o == nil || o.Host == nil {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEventForwarder) GetHostOk() (*string, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *CreateEventForwarder) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *CreateEventForwarder) SetHost(v string) {
	o.Host = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *CreateEventForwarder) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEventForwarder) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *CreateEventForwarder) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *CreateEventForwarder) SetJson(v bool) {
	o.Json = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *CreateEventForwarder) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEventForwarder) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *CreateEventForwarder) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *CreateEventForwarder) SetKey(v string) {
	o.Key = &v
}

// GetName returns the Name field value
func (o *CreateEventForwarder) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateEventForwarder) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateEventForwarder) SetName(v string) {
	o.Name = v
}

// GetRunnerType returns the RunnerType field value
func (o *CreateEventForwarder) GetRunnerType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.RunnerType
}

// GetRunnerTypeOk returns a tuple with the RunnerType field value
// and a boolean to check if the value has been set.
func (o *CreateEventForwarder) GetRunnerTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RunnerType, true
}

// SetRunnerType sets field value
func (o *CreateEventForwarder) SetRunnerType(v string) {
	o.RunnerType = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *CreateEventForwarder) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEventForwarder) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *CreateEventForwarder) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *CreateEventForwarder) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *CreateEventForwarder) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEventForwarder) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *CreateEventForwarder) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *CreateEventForwarder) SetUidToken(v string) {
	o.UidToken = &v
}

func (o CreateEventForwarder) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AdminName != nil {
		toSerialize["admin-name"] = o.AdminName
	}
	if o.AdminPwd != nil {
		toSerialize["admin-pwd"] = o.AdminPwd
	}
	if o.Comment != nil {
		toSerialize["comment"] = o.Comment
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.EmailTo != nil {
		toSerialize["email-to"] = o.EmailTo
	}
	if true {
		toSerialize["event-source-locations"] = o.EventSourceLocations
	}
	if o.EventSourceType != nil {
		toSerialize["event-source-type"] = o.EventSourceType
	}
	if o.EventTypes != nil {
		toSerialize["event-types"] = o.EventTypes
	}
	if o.Every != nil {
		toSerialize["every"] = o.Every
	}
	if true {
		toSerialize["forwarder-type"] = o.ForwarderType
	}
	if o.Host != nil {
		toSerialize["host"] = o.Host
	}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["runner-type"] = o.RunnerType
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableCreateEventForwarder struct {
	value *CreateEventForwarder
	isSet bool
}

func (v NullableCreateEventForwarder) Get() *CreateEventForwarder {
	return v.value
}

func (v *NullableCreateEventForwarder) Set(val *CreateEventForwarder) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateEventForwarder) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateEventForwarder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateEventForwarder(val *CreateEventForwarder) *NullableCreateEventForwarder {
	return &NullableCreateEventForwarder{value: val, isSet: true}
}

func (v NullableCreateEventForwarder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateEventForwarder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


