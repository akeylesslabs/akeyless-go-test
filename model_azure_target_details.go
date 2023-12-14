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

// AzureTargetDetails struct for AzureTargetDetails
type AzureTargetDetails struct {
	AzureClientId *string `json:"azure_client_id,omitempty"`
	AzureClientSecret *string `json:"azure_client_secret,omitempty"`
	AzureResourceGroupName *string `json:"azure_resource_group_name,omitempty"`
	AzureResourceName *string `json:"azure_resource_name,omitempty"`
	AzureSubscriptionId *string `json:"azure_subscription_id,omitempty"`
	AzureTenantId *string `json:"azure_tenant_id,omitempty"`
	UseGwCloudIdentity *bool `json:"use_gw_cloud_identity,omitempty"`
}

// NewAzureTargetDetails instantiates a new AzureTargetDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureTargetDetails() *AzureTargetDetails {
	this := AzureTargetDetails{}
	return &this
}

// NewAzureTargetDetailsWithDefaults instantiates a new AzureTargetDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureTargetDetailsWithDefaults() *AzureTargetDetails {
	this := AzureTargetDetails{}
	return &this
}

// GetAzureClientId returns the AzureClientId field value if set, zero value otherwise.
func (o *AzureTargetDetails) GetAzureClientId() string {
	if o == nil || o.AzureClientId == nil {
		var ret string
		return ret
	}
	return *o.AzureClientId
}

// GetAzureClientIdOk returns a tuple with the AzureClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureTargetDetails) GetAzureClientIdOk() (*string, bool) {
	if o == nil || o.AzureClientId == nil {
		return nil, false
	}
	return o.AzureClientId, true
}

// HasAzureClientId returns a boolean if a field has been set.
func (o *AzureTargetDetails) HasAzureClientId() bool {
	if o != nil && o.AzureClientId != nil {
		return true
	}

	return false
}

// SetAzureClientId gets a reference to the given string and assigns it to the AzureClientId field.
func (o *AzureTargetDetails) SetAzureClientId(v string) {
	o.AzureClientId = &v
}

// GetAzureClientSecret returns the AzureClientSecret field value if set, zero value otherwise.
func (o *AzureTargetDetails) GetAzureClientSecret() string {
	if o == nil || o.AzureClientSecret == nil {
		var ret string
		return ret
	}
	return *o.AzureClientSecret
}

// GetAzureClientSecretOk returns a tuple with the AzureClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureTargetDetails) GetAzureClientSecretOk() (*string, bool) {
	if o == nil || o.AzureClientSecret == nil {
		return nil, false
	}
	return o.AzureClientSecret, true
}

// HasAzureClientSecret returns a boolean if a field has been set.
func (o *AzureTargetDetails) HasAzureClientSecret() bool {
	if o != nil && o.AzureClientSecret != nil {
		return true
	}

	return false
}

// SetAzureClientSecret gets a reference to the given string and assigns it to the AzureClientSecret field.
func (o *AzureTargetDetails) SetAzureClientSecret(v string) {
	o.AzureClientSecret = &v
}

// GetAzureResourceGroupName returns the AzureResourceGroupName field value if set, zero value otherwise.
func (o *AzureTargetDetails) GetAzureResourceGroupName() string {
	if o == nil || o.AzureResourceGroupName == nil {
		var ret string
		return ret
	}
	return *o.AzureResourceGroupName
}

// GetAzureResourceGroupNameOk returns a tuple with the AzureResourceGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureTargetDetails) GetAzureResourceGroupNameOk() (*string, bool) {
	if o == nil || o.AzureResourceGroupName == nil {
		return nil, false
	}
	return o.AzureResourceGroupName, true
}

// HasAzureResourceGroupName returns a boolean if a field has been set.
func (o *AzureTargetDetails) HasAzureResourceGroupName() bool {
	if o != nil && o.AzureResourceGroupName != nil {
		return true
	}

	return false
}

// SetAzureResourceGroupName gets a reference to the given string and assigns it to the AzureResourceGroupName field.
func (o *AzureTargetDetails) SetAzureResourceGroupName(v string) {
	o.AzureResourceGroupName = &v
}

// GetAzureResourceName returns the AzureResourceName field value if set, zero value otherwise.
func (o *AzureTargetDetails) GetAzureResourceName() string {
	if o == nil || o.AzureResourceName == nil {
		var ret string
		return ret
	}
	return *o.AzureResourceName
}

// GetAzureResourceNameOk returns a tuple with the AzureResourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureTargetDetails) GetAzureResourceNameOk() (*string, bool) {
	if o == nil || o.AzureResourceName == nil {
		return nil, false
	}
	return o.AzureResourceName, true
}

// HasAzureResourceName returns a boolean if a field has been set.
func (o *AzureTargetDetails) HasAzureResourceName() bool {
	if o != nil && o.AzureResourceName != nil {
		return true
	}

	return false
}

// SetAzureResourceName gets a reference to the given string and assigns it to the AzureResourceName field.
func (o *AzureTargetDetails) SetAzureResourceName(v string) {
	o.AzureResourceName = &v
}

// GetAzureSubscriptionId returns the AzureSubscriptionId field value if set, zero value otherwise.
func (o *AzureTargetDetails) GetAzureSubscriptionId() string {
	if o == nil || o.AzureSubscriptionId == nil {
		var ret string
		return ret
	}
	return *o.AzureSubscriptionId
}

// GetAzureSubscriptionIdOk returns a tuple with the AzureSubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureTargetDetails) GetAzureSubscriptionIdOk() (*string, bool) {
	if o == nil || o.AzureSubscriptionId == nil {
		return nil, false
	}
	return o.AzureSubscriptionId, true
}

// HasAzureSubscriptionId returns a boolean if a field has been set.
func (o *AzureTargetDetails) HasAzureSubscriptionId() bool {
	if o != nil && o.AzureSubscriptionId != nil {
		return true
	}

	return false
}

// SetAzureSubscriptionId gets a reference to the given string and assigns it to the AzureSubscriptionId field.
func (o *AzureTargetDetails) SetAzureSubscriptionId(v string) {
	o.AzureSubscriptionId = &v
}

// GetAzureTenantId returns the AzureTenantId field value if set, zero value otherwise.
func (o *AzureTargetDetails) GetAzureTenantId() string {
	if o == nil || o.AzureTenantId == nil {
		var ret string
		return ret
	}
	return *o.AzureTenantId
}

// GetAzureTenantIdOk returns a tuple with the AzureTenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureTargetDetails) GetAzureTenantIdOk() (*string, bool) {
	if o == nil || o.AzureTenantId == nil {
		return nil, false
	}
	return o.AzureTenantId, true
}

// HasAzureTenantId returns a boolean if a field has been set.
func (o *AzureTargetDetails) HasAzureTenantId() bool {
	if o != nil && o.AzureTenantId != nil {
		return true
	}

	return false
}

// SetAzureTenantId gets a reference to the given string and assigns it to the AzureTenantId field.
func (o *AzureTargetDetails) SetAzureTenantId(v string) {
	o.AzureTenantId = &v
}

// GetUseGwCloudIdentity returns the UseGwCloudIdentity field value if set, zero value otherwise.
func (o *AzureTargetDetails) GetUseGwCloudIdentity() bool {
	if o == nil || o.UseGwCloudIdentity == nil {
		var ret bool
		return ret
	}
	return *o.UseGwCloudIdentity
}

// GetUseGwCloudIdentityOk returns a tuple with the UseGwCloudIdentity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureTargetDetails) GetUseGwCloudIdentityOk() (*bool, bool) {
	if o == nil || o.UseGwCloudIdentity == nil {
		return nil, false
	}
	return o.UseGwCloudIdentity, true
}

// HasUseGwCloudIdentity returns a boolean if a field has been set.
func (o *AzureTargetDetails) HasUseGwCloudIdentity() bool {
	if o != nil && o.UseGwCloudIdentity != nil {
		return true
	}

	return false
}

// SetUseGwCloudIdentity gets a reference to the given bool and assigns it to the UseGwCloudIdentity field.
func (o *AzureTargetDetails) SetUseGwCloudIdentity(v bool) {
	o.UseGwCloudIdentity = &v
}

func (o AzureTargetDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AzureClientId != nil {
		toSerialize["azure_client_id"] = o.AzureClientId
	}
	if o.AzureClientSecret != nil {
		toSerialize["azure_client_secret"] = o.AzureClientSecret
	}
	if o.AzureResourceGroupName != nil {
		toSerialize["azure_resource_group_name"] = o.AzureResourceGroupName
	}
	if o.AzureResourceName != nil {
		toSerialize["azure_resource_name"] = o.AzureResourceName
	}
	if o.AzureSubscriptionId != nil {
		toSerialize["azure_subscription_id"] = o.AzureSubscriptionId
	}
	if o.AzureTenantId != nil {
		toSerialize["azure_tenant_id"] = o.AzureTenantId
	}
	if o.UseGwCloudIdentity != nil {
		toSerialize["use_gw_cloud_identity"] = o.UseGwCloudIdentity
	}
	return json.Marshal(toSerialize)
}

type NullableAzureTargetDetails struct {
	value *AzureTargetDetails
	isSet bool
}

func (v NullableAzureTargetDetails) Get() *AzureTargetDetails {
	return v.value
}

func (v *NullableAzureTargetDetails) Set(val *AzureTargetDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureTargetDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureTargetDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureTargetDetails(val *AzureTargetDetails) *NullableAzureTargetDetails {
	return &NullableAzureTargetDetails{value: val, isSet: true}
}

func (v NullableAzureTargetDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureTargetDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

