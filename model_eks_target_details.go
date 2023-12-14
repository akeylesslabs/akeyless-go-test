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

// EKSTargetDetails EKSTargetDetails defines details related to connecting to a EKS (Elastic Container Service) target
type EKSTargetDetails struct {
	EksAccessKeyId *string `json:"eks_access_key_id,omitempty"`
	EksClusterCaCertificate *string `json:"eks_cluster_ca_certificate,omitempty"`
	EksClusterEndpoint *string `json:"eks_cluster_endpoint,omitempty"`
	EksClusterName *string `json:"eks_cluster_name,omitempty"`
	EksRegion *string `json:"eks_region,omitempty"`
	EksSecretAccessKey *string `json:"eks_secret_access_key,omitempty"`
	UseGwCloudIdentity *bool `json:"use_gw_cloud_identity,omitempty"`
}

// NewEKSTargetDetails instantiates a new EKSTargetDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEKSTargetDetails() *EKSTargetDetails {
	this := EKSTargetDetails{}
	return &this
}

// NewEKSTargetDetailsWithDefaults instantiates a new EKSTargetDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEKSTargetDetailsWithDefaults() *EKSTargetDetails {
	this := EKSTargetDetails{}
	return &this
}

// GetEksAccessKeyId returns the EksAccessKeyId field value if set, zero value otherwise.
func (o *EKSTargetDetails) GetEksAccessKeyId() string {
	if o == nil || o.EksAccessKeyId == nil {
		var ret string
		return ret
	}
	return *o.EksAccessKeyId
}

// GetEksAccessKeyIdOk returns a tuple with the EksAccessKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EKSTargetDetails) GetEksAccessKeyIdOk() (*string, bool) {
	if o == nil || o.EksAccessKeyId == nil {
		return nil, false
	}
	return o.EksAccessKeyId, true
}

// HasEksAccessKeyId returns a boolean if a field has been set.
func (o *EKSTargetDetails) HasEksAccessKeyId() bool {
	if o != nil && o.EksAccessKeyId != nil {
		return true
	}

	return false
}

// SetEksAccessKeyId gets a reference to the given string and assigns it to the EksAccessKeyId field.
func (o *EKSTargetDetails) SetEksAccessKeyId(v string) {
	o.EksAccessKeyId = &v
}

// GetEksClusterCaCertificate returns the EksClusterCaCertificate field value if set, zero value otherwise.
func (o *EKSTargetDetails) GetEksClusterCaCertificate() string {
	if o == nil || o.EksClusterCaCertificate == nil {
		var ret string
		return ret
	}
	return *o.EksClusterCaCertificate
}

// GetEksClusterCaCertificateOk returns a tuple with the EksClusterCaCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EKSTargetDetails) GetEksClusterCaCertificateOk() (*string, bool) {
	if o == nil || o.EksClusterCaCertificate == nil {
		return nil, false
	}
	return o.EksClusterCaCertificate, true
}

// HasEksClusterCaCertificate returns a boolean if a field has been set.
func (o *EKSTargetDetails) HasEksClusterCaCertificate() bool {
	if o != nil && o.EksClusterCaCertificate != nil {
		return true
	}

	return false
}

// SetEksClusterCaCertificate gets a reference to the given string and assigns it to the EksClusterCaCertificate field.
func (o *EKSTargetDetails) SetEksClusterCaCertificate(v string) {
	o.EksClusterCaCertificate = &v
}

// GetEksClusterEndpoint returns the EksClusterEndpoint field value if set, zero value otherwise.
func (o *EKSTargetDetails) GetEksClusterEndpoint() string {
	if o == nil || o.EksClusterEndpoint == nil {
		var ret string
		return ret
	}
	return *o.EksClusterEndpoint
}

// GetEksClusterEndpointOk returns a tuple with the EksClusterEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EKSTargetDetails) GetEksClusterEndpointOk() (*string, bool) {
	if o == nil || o.EksClusterEndpoint == nil {
		return nil, false
	}
	return o.EksClusterEndpoint, true
}

// HasEksClusterEndpoint returns a boolean if a field has been set.
func (o *EKSTargetDetails) HasEksClusterEndpoint() bool {
	if o != nil && o.EksClusterEndpoint != nil {
		return true
	}

	return false
}

// SetEksClusterEndpoint gets a reference to the given string and assigns it to the EksClusterEndpoint field.
func (o *EKSTargetDetails) SetEksClusterEndpoint(v string) {
	o.EksClusterEndpoint = &v
}

// GetEksClusterName returns the EksClusterName field value if set, zero value otherwise.
func (o *EKSTargetDetails) GetEksClusterName() string {
	if o == nil || o.EksClusterName == nil {
		var ret string
		return ret
	}
	return *o.EksClusterName
}

// GetEksClusterNameOk returns a tuple with the EksClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EKSTargetDetails) GetEksClusterNameOk() (*string, bool) {
	if o == nil || o.EksClusterName == nil {
		return nil, false
	}
	return o.EksClusterName, true
}

// HasEksClusterName returns a boolean if a field has been set.
func (o *EKSTargetDetails) HasEksClusterName() bool {
	if o != nil && o.EksClusterName != nil {
		return true
	}

	return false
}

// SetEksClusterName gets a reference to the given string and assigns it to the EksClusterName field.
func (o *EKSTargetDetails) SetEksClusterName(v string) {
	o.EksClusterName = &v
}

// GetEksRegion returns the EksRegion field value if set, zero value otherwise.
func (o *EKSTargetDetails) GetEksRegion() string {
	if o == nil || o.EksRegion == nil {
		var ret string
		return ret
	}
	return *o.EksRegion
}

// GetEksRegionOk returns a tuple with the EksRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EKSTargetDetails) GetEksRegionOk() (*string, bool) {
	if o == nil || o.EksRegion == nil {
		return nil, false
	}
	return o.EksRegion, true
}

// HasEksRegion returns a boolean if a field has been set.
func (o *EKSTargetDetails) HasEksRegion() bool {
	if o != nil && o.EksRegion != nil {
		return true
	}

	return false
}

// SetEksRegion gets a reference to the given string and assigns it to the EksRegion field.
func (o *EKSTargetDetails) SetEksRegion(v string) {
	o.EksRegion = &v
}

// GetEksSecretAccessKey returns the EksSecretAccessKey field value if set, zero value otherwise.
func (o *EKSTargetDetails) GetEksSecretAccessKey() string {
	if o == nil || o.EksSecretAccessKey == nil {
		var ret string
		return ret
	}
	return *o.EksSecretAccessKey
}

// GetEksSecretAccessKeyOk returns a tuple with the EksSecretAccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EKSTargetDetails) GetEksSecretAccessKeyOk() (*string, bool) {
	if o == nil || o.EksSecretAccessKey == nil {
		return nil, false
	}
	return o.EksSecretAccessKey, true
}

// HasEksSecretAccessKey returns a boolean if a field has been set.
func (o *EKSTargetDetails) HasEksSecretAccessKey() bool {
	if o != nil && o.EksSecretAccessKey != nil {
		return true
	}

	return false
}

// SetEksSecretAccessKey gets a reference to the given string and assigns it to the EksSecretAccessKey field.
func (o *EKSTargetDetails) SetEksSecretAccessKey(v string) {
	o.EksSecretAccessKey = &v
}

// GetUseGwCloudIdentity returns the UseGwCloudIdentity field value if set, zero value otherwise.
func (o *EKSTargetDetails) GetUseGwCloudIdentity() bool {
	if o == nil || o.UseGwCloudIdentity == nil {
		var ret bool
		return ret
	}
	return *o.UseGwCloudIdentity
}

// GetUseGwCloudIdentityOk returns a tuple with the UseGwCloudIdentity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EKSTargetDetails) GetUseGwCloudIdentityOk() (*bool, bool) {
	if o == nil || o.UseGwCloudIdentity == nil {
		return nil, false
	}
	return o.UseGwCloudIdentity, true
}

// HasUseGwCloudIdentity returns a boolean if a field has been set.
func (o *EKSTargetDetails) HasUseGwCloudIdentity() bool {
	if o != nil && o.UseGwCloudIdentity != nil {
		return true
	}

	return false
}

// SetUseGwCloudIdentity gets a reference to the given bool and assigns it to the UseGwCloudIdentity field.
func (o *EKSTargetDetails) SetUseGwCloudIdentity(v bool) {
	o.UseGwCloudIdentity = &v
}

func (o EKSTargetDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EksAccessKeyId != nil {
		toSerialize["eks_access_key_id"] = o.EksAccessKeyId
	}
	if o.EksClusterCaCertificate != nil {
		toSerialize["eks_cluster_ca_certificate"] = o.EksClusterCaCertificate
	}
	if o.EksClusterEndpoint != nil {
		toSerialize["eks_cluster_endpoint"] = o.EksClusterEndpoint
	}
	if o.EksClusterName != nil {
		toSerialize["eks_cluster_name"] = o.EksClusterName
	}
	if o.EksRegion != nil {
		toSerialize["eks_region"] = o.EksRegion
	}
	if o.EksSecretAccessKey != nil {
		toSerialize["eks_secret_access_key"] = o.EksSecretAccessKey
	}
	if o.UseGwCloudIdentity != nil {
		toSerialize["use_gw_cloud_identity"] = o.UseGwCloudIdentity
	}
	return json.Marshal(toSerialize)
}

type NullableEKSTargetDetails struct {
	value *EKSTargetDetails
	isSet bool
}

func (v NullableEKSTargetDetails) Get() *EKSTargetDetails {
	return v.value
}

func (v *NullableEKSTargetDetails) Set(val *EKSTargetDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableEKSTargetDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableEKSTargetDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEKSTargetDetails(val *EKSTargetDetails) *NullableEKSTargetDetails {
	return &NullableEKSTargetDetails{value: val, isSet: true}
}

func (v NullableEKSTargetDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEKSTargetDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

