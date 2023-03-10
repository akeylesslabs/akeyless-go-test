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

// GCPAccessRules struct for GCPAccessRules
type GCPAccessRules struct {
	// The audience in the JWT
	Audience *string `json:"audience,omitempty"`
	// A map of GCP labels formatted as \"key:value\" strings that must be set on authorized GCE instances. TODO: Because GCP labels are not currently ACL'd ....
	BoundLabels *map[string]string `json:"bound_labels,omitempty"`
	// Human and Machine authentication section Array of GCP project IDs. Only entities belonging to any of the provided projects can authenticate.
	BoundProjects *[]string `json:"bound_projects,omitempty"`
	// List of regions that a GCE instance must belong to in order to be authenticated. TODO: If bound_instance_groups is provided, it is assumed to be a regional group and the group must belong to this region. If bound_zones are provided, this attribute is ignored.
	BoundRegions *[]string `json:"bound_regions,omitempty"`
	// List of service accounts the service account must be part of in order to be authenticated
	BoundServiceAccounts *[]string `json:"bound_service_accounts,omitempty"`
	// === Machine authentication section === List of zones that a GCE instance must belong to in order to be authenticated. TODO: If bound_instance_groups is provided, it is assumed to be a zonal group and the group must belong to this zone.
	BoundZones *[]string `json:"bound_zones,omitempty"`
	// ServiceAccount holds the credentials file contents to be used by Akeyless to validate IAM (Human) and GCE (Machine) logins against GCP base64 encoded string
	ServiceAccount *string `json:"service_account,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewGCPAccessRules instantiates a new GCPAccessRules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGCPAccessRules() *GCPAccessRules {
	this := GCPAccessRules{}
	var audience string = "akeyless.io"
	this.Audience = &audience
	return &this
}

// NewGCPAccessRulesWithDefaults instantiates a new GCPAccessRules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGCPAccessRulesWithDefaults() *GCPAccessRules {
	this := GCPAccessRules{}
	var audience string = "akeyless.io"
	this.Audience = &audience
	return &this
}

// GetAudience returns the Audience field value if set, zero value otherwise.
func (o *GCPAccessRules) GetAudience() string {
	if o == nil || o.Audience == nil {
		var ret string
		return ret
	}
	return *o.Audience
}

// GetAudienceOk returns a tuple with the Audience field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPAccessRules) GetAudienceOk() (*string, bool) {
	if o == nil || o.Audience == nil {
		return nil, false
	}
	return o.Audience, true
}

// HasAudience returns a boolean if a field has been set.
func (o *GCPAccessRules) HasAudience() bool {
	if o != nil && o.Audience != nil {
		return true
	}

	return false
}

// SetAudience gets a reference to the given string and assigns it to the Audience field.
func (o *GCPAccessRules) SetAudience(v string) {
	o.Audience = &v
}

// GetBoundLabels returns the BoundLabels field value if set, zero value otherwise.
func (o *GCPAccessRules) GetBoundLabels() map[string]string {
	if o == nil || o.BoundLabels == nil {
		var ret map[string]string
		return ret
	}
	return *o.BoundLabels
}

// GetBoundLabelsOk returns a tuple with the BoundLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPAccessRules) GetBoundLabelsOk() (*map[string]string, bool) {
	if o == nil || o.BoundLabels == nil {
		return nil, false
	}
	return o.BoundLabels, true
}

// HasBoundLabels returns a boolean if a field has been set.
func (o *GCPAccessRules) HasBoundLabels() bool {
	if o != nil && o.BoundLabels != nil {
		return true
	}

	return false
}

// SetBoundLabels gets a reference to the given map[string]string and assigns it to the BoundLabels field.
func (o *GCPAccessRules) SetBoundLabels(v map[string]string) {
	o.BoundLabels = &v
}

// GetBoundProjects returns the BoundProjects field value if set, zero value otherwise.
func (o *GCPAccessRules) GetBoundProjects() []string {
	if o == nil || o.BoundProjects == nil {
		var ret []string
		return ret
	}
	return *o.BoundProjects
}

// GetBoundProjectsOk returns a tuple with the BoundProjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPAccessRules) GetBoundProjectsOk() (*[]string, bool) {
	if o == nil || o.BoundProjects == nil {
		return nil, false
	}
	return o.BoundProjects, true
}

// HasBoundProjects returns a boolean if a field has been set.
func (o *GCPAccessRules) HasBoundProjects() bool {
	if o != nil && o.BoundProjects != nil {
		return true
	}

	return false
}

// SetBoundProjects gets a reference to the given []string and assigns it to the BoundProjects field.
func (o *GCPAccessRules) SetBoundProjects(v []string) {
	o.BoundProjects = &v
}

// GetBoundRegions returns the BoundRegions field value if set, zero value otherwise.
func (o *GCPAccessRules) GetBoundRegions() []string {
	if o == nil || o.BoundRegions == nil {
		var ret []string
		return ret
	}
	return *o.BoundRegions
}

// GetBoundRegionsOk returns a tuple with the BoundRegions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPAccessRules) GetBoundRegionsOk() (*[]string, bool) {
	if o == nil || o.BoundRegions == nil {
		return nil, false
	}
	return o.BoundRegions, true
}

// HasBoundRegions returns a boolean if a field has been set.
func (o *GCPAccessRules) HasBoundRegions() bool {
	if o != nil && o.BoundRegions != nil {
		return true
	}

	return false
}

// SetBoundRegions gets a reference to the given []string and assigns it to the BoundRegions field.
func (o *GCPAccessRules) SetBoundRegions(v []string) {
	o.BoundRegions = &v
}

// GetBoundServiceAccounts returns the BoundServiceAccounts field value if set, zero value otherwise.
func (o *GCPAccessRules) GetBoundServiceAccounts() []string {
	if o == nil || o.BoundServiceAccounts == nil {
		var ret []string
		return ret
	}
	return *o.BoundServiceAccounts
}

// GetBoundServiceAccountsOk returns a tuple with the BoundServiceAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPAccessRules) GetBoundServiceAccountsOk() (*[]string, bool) {
	if o == nil || o.BoundServiceAccounts == nil {
		return nil, false
	}
	return o.BoundServiceAccounts, true
}

// HasBoundServiceAccounts returns a boolean if a field has been set.
func (o *GCPAccessRules) HasBoundServiceAccounts() bool {
	if o != nil && o.BoundServiceAccounts != nil {
		return true
	}

	return false
}

// SetBoundServiceAccounts gets a reference to the given []string and assigns it to the BoundServiceAccounts field.
func (o *GCPAccessRules) SetBoundServiceAccounts(v []string) {
	o.BoundServiceAccounts = &v
}

// GetBoundZones returns the BoundZones field value if set, zero value otherwise.
func (o *GCPAccessRules) GetBoundZones() []string {
	if o == nil || o.BoundZones == nil {
		var ret []string
		return ret
	}
	return *o.BoundZones
}

// GetBoundZonesOk returns a tuple with the BoundZones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPAccessRules) GetBoundZonesOk() (*[]string, bool) {
	if o == nil || o.BoundZones == nil {
		return nil, false
	}
	return o.BoundZones, true
}

// HasBoundZones returns a boolean if a field has been set.
func (o *GCPAccessRules) HasBoundZones() bool {
	if o != nil && o.BoundZones != nil {
		return true
	}

	return false
}

// SetBoundZones gets a reference to the given []string and assigns it to the BoundZones field.
func (o *GCPAccessRules) SetBoundZones(v []string) {
	o.BoundZones = &v
}

// GetServiceAccount returns the ServiceAccount field value if set, zero value otherwise.
func (o *GCPAccessRules) GetServiceAccount() string {
	if o == nil || o.ServiceAccount == nil {
		var ret string
		return ret
	}
	return *o.ServiceAccount
}

// GetServiceAccountOk returns a tuple with the ServiceAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPAccessRules) GetServiceAccountOk() (*string, bool) {
	if o == nil || o.ServiceAccount == nil {
		return nil, false
	}
	return o.ServiceAccount, true
}

// HasServiceAccount returns a boolean if a field has been set.
func (o *GCPAccessRules) HasServiceAccount() bool {
	if o != nil && o.ServiceAccount != nil {
		return true
	}

	return false
}

// SetServiceAccount gets a reference to the given string and assigns it to the ServiceAccount field.
func (o *GCPAccessRules) SetServiceAccount(v string) {
	o.ServiceAccount = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GCPAccessRules) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPAccessRules) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GCPAccessRules) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GCPAccessRules) SetType(v string) {
	o.Type = &v
}

func (o GCPAccessRules) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Audience != nil {
		toSerialize["audience"] = o.Audience
	}
	if o.BoundLabels != nil {
		toSerialize["bound_labels"] = o.BoundLabels
	}
	if o.BoundProjects != nil {
		toSerialize["bound_projects"] = o.BoundProjects
	}
	if o.BoundRegions != nil {
		toSerialize["bound_regions"] = o.BoundRegions
	}
	if o.BoundServiceAccounts != nil {
		toSerialize["bound_service_accounts"] = o.BoundServiceAccounts
	}
	if o.BoundZones != nil {
		toSerialize["bound_zones"] = o.BoundZones
	}
	if o.ServiceAccount != nil {
		toSerialize["service_account"] = o.ServiceAccount
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableGCPAccessRules struct {
	value *GCPAccessRules
	isSet bool
}

func (v NullableGCPAccessRules) Get() *GCPAccessRules {
	return v.value
}

func (v *NullableGCPAccessRules) Set(val *GCPAccessRules) {
	v.value = val
	v.isSet = true
}

func (v NullableGCPAccessRules) IsSet() bool {
	return v.isSet
}

func (v *NullableGCPAccessRules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGCPAccessRules(val *GCPAccessRules) *NullableGCPAccessRules {
	return &NullableGCPAccessRules{value: val, isSet: true}
}

func (v NullableGCPAccessRules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGCPAccessRules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


