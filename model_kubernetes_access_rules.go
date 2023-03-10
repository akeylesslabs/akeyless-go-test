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

// KubernetesAccessRules struct for KubernetesAccessRules
type KubernetesAccessRules struct {
	Alg *string `json:"alg,omitempty"`
	// Audience is an optional Kubernetes jwt claim to verify
	Audience *string `json:"audience,omitempty"`
	// A list of namespaces that the authentication is restricted to.
	BoundNamespaces *[]string `json:"bound_namespaces,omitempty"`
	// A list of pods names that the authentication is restricted to.
	BoundPodNames *[]string `json:"bound_pod_names,omitempty"`
	// A list of service account names that the authentication is restricted to.
	BoundServiceAccountNames *[]string `json:"bound_service_account_names,omitempty"`
	// Generate public/private key (the private key is required for the K8S Auth Config in the Akeyless Gateway)
	GenKeyPair *string `json:"gen_key_pair,omitempty"`
	// The public key value of the Kubernetes auth method configuration in the Akeyless Gateway.
	PubKey *string `json:"pub_key,omitempty"`
}

// NewKubernetesAccessRules instantiates a new KubernetesAccessRules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesAccessRules() *KubernetesAccessRules {
	this := KubernetesAccessRules{}
	return &this
}

// NewKubernetesAccessRulesWithDefaults instantiates a new KubernetesAccessRules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesAccessRulesWithDefaults() *KubernetesAccessRules {
	this := KubernetesAccessRules{}
	return &this
}

// GetAlg returns the Alg field value if set, zero value otherwise.
func (o *KubernetesAccessRules) GetAlg() string {
	if o == nil || o.Alg == nil {
		var ret string
		return ret
	}
	return *o.Alg
}

// GetAlgOk returns a tuple with the Alg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAccessRules) GetAlgOk() (*string, bool) {
	if o == nil || o.Alg == nil {
		return nil, false
	}
	return o.Alg, true
}

// HasAlg returns a boolean if a field has been set.
func (o *KubernetesAccessRules) HasAlg() bool {
	if o != nil && o.Alg != nil {
		return true
	}

	return false
}

// SetAlg gets a reference to the given string and assigns it to the Alg field.
func (o *KubernetesAccessRules) SetAlg(v string) {
	o.Alg = &v
}

// GetAudience returns the Audience field value if set, zero value otherwise.
func (o *KubernetesAccessRules) GetAudience() string {
	if o == nil || o.Audience == nil {
		var ret string
		return ret
	}
	return *o.Audience
}

// GetAudienceOk returns a tuple with the Audience field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAccessRules) GetAudienceOk() (*string, bool) {
	if o == nil || o.Audience == nil {
		return nil, false
	}
	return o.Audience, true
}

// HasAudience returns a boolean if a field has been set.
func (o *KubernetesAccessRules) HasAudience() bool {
	if o != nil && o.Audience != nil {
		return true
	}

	return false
}

// SetAudience gets a reference to the given string and assigns it to the Audience field.
func (o *KubernetesAccessRules) SetAudience(v string) {
	o.Audience = &v
}

// GetBoundNamespaces returns the BoundNamespaces field value if set, zero value otherwise.
func (o *KubernetesAccessRules) GetBoundNamespaces() []string {
	if o == nil || o.BoundNamespaces == nil {
		var ret []string
		return ret
	}
	return *o.BoundNamespaces
}

// GetBoundNamespacesOk returns a tuple with the BoundNamespaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAccessRules) GetBoundNamespacesOk() (*[]string, bool) {
	if o == nil || o.BoundNamespaces == nil {
		return nil, false
	}
	return o.BoundNamespaces, true
}

// HasBoundNamespaces returns a boolean if a field has been set.
func (o *KubernetesAccessRules) HasBoundNamespaces() bool {
	if o != nil && o.BoundNamespaces != nil {
		return true
	}

	return false
}

// SetBoundNamespaces gets a reference to the given []string and assigns it to the BoundNamespaces field.
func (o *KubernetesAccessRules) SetBoundNamespaces(v []string) {
	o.BoundNamespaces = &v
}

// GetBoundPodNames returns the BoundPodNames field value if set, zero value otherwise.
func (o *KubernetesAccessRules) GetBoundPodNames() []string {
	if o == nil || o.BoundPodNames == nil {
		var ret []string
		return ret
	}
	return *o.BoundPodNames
}

// GetBoundPodNamesOk returns a tuple with the BoundPodNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAccessRules) GetBoundPodNamesOk() (*[]string, bool) {
	if o == nil || o.BoundPodNames == nil {
		return nil, false
	}
	return o.BoundPodNames, true
}

// HasBoundPodNames returns a boolean if a field has been set.
func (o *KubernetesAccessRules) HasBoundPodNames() bool {
	if o != nil && o.BoundPodNames != nil {
		return true
	}

	return false
}

// SetBoundPodNames gets a reference to the given []string and assigns it to the BoundPodNames field.
func (o *KubernetesAccessRules) SetBoundPodNames(v []string) {
	o.BoundPodNames = &v
}

// GetBoundServiceAccountNames returns the BoundServiceAccountNames field value if set, zero value otherwise.
func (o *KubernetesAccessRules) GetBoundServiceAccountNames() []string {
	if o == nil || o.BoundServiceAccountNames == nil {
		var ret []string
		return ret
	}
	return *o.BoundServiceAccountNames
}

// GetBoundServiceAccountNamesOk returns a tuple with the BoundServiceAccountNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAccessRules) GetBoundServiceAccountNamesOk() (*[]string, bool) {
	if o == nil || o.BoundServiceAccountNames == nil {
		return nil, false
	}
	return o.BoundServiceAccountNames, true
}

// HasBoundServiceAccountNames returns a boolean if a field has been set.
func (o *KubernetesAccessRules) HasBoundServiceAccountNames() bool {
	if o != nil && o.BoundServiceAccountNames != nil {
		return true
	}

	return false
}

// SetBoundServiceAccountNames gets a reference to the given []string and assigns it to the BoundServiceAccountNames field.
func (o *KubernetesAccessRules) SetBoundServiceAccountNames(v []string) {
	o.BoundServiceAccountNames = &v
}

// GetGenKeyPair returns the GenKeyPair field value if set, zero value otherwise.
func (o *KubernetesAccessRules) GetGenKeyPair() string {
	if o == nil || o.GenKeyPair == nil {
		var ret string
		return ret
	}
	return *o.GenKeyPair
}

// GetGenKeyPairOk returns a tuple with the GenKeyPair field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAccessRules) GetGenKeyPairOk() (*string, bool) {
	if o == nil || o.GenKeyPair == nil {
		return nil, false
	}
	return o.GenKeyPair, true
}

// HasGenKeyPair returns a boolean if a field has been set.
func (o *KubernetesAccessRules) HasGenKeyPair() bool {
	if o != nil && o.GenKeyPair != nil {
		return true
	}

	return false
}

// SetGenKeyPair gets a reference to the given string and assigns it to the GenKeyPair field.
func (o *KubernetesAccessRules) SetGenKeyPair(v string) {
	o.GenKeyPair = &v
}

// GetPubKey returns the PubKey field value if set, zero value otherwise.
func (o *KubernetesAccessRules) GetPubKey() string {
	if o == nil || o.PubKey == nil {
		var ret string
		return ret
	}
	return *o.PubKey
}

// GetPubKeyOk returns a tuple with the PubKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAccessRules) GetPubKeyOk() (*string, bool) {
	if o == nil || o.PubKey == nil {
		return nil, false
	}
	return o.PubKey, true
}

// HasPubKey returns a boolean if a field has been set.
func (o *KubernetesAccessRules) HasPubKey() bool {
	if o != nil && o.PubKey != nil {
		return true
	}

	return false
}

// SetPubKey gets a reference to the given string and assigns it to the PubKey field.
func (o *KubernetesAccessRules) SetPubKey(v string) {
	o.PubKey = &v
}

func (o KubernetesAccessRules) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Alg != nil {
		toSerialize["alg"] = o.Alg
	}
	if o.Audience != nil {
		toSerialize["audience"] = o.Audience
	}
	if o.BoundNamespaces != nil {
		toSerialize["bound_namespaces"] = o.BoundNamespaces
	}
	if o.BoundPodNames != nil {
		toSerialize["bound_pod_names"] = o.BoundPodNames
	}
	if o.BoundServiceAccountNames != nil {
		toSerialize["bound_service_account_names"] = o.BoundServiceAccountNames
	}
	if o.GenKeyPair != nil {
		toSerialize["gen_key_pair"] = o.GenKeyPair
	}
	if o.PubKey != nil {
		toSerialize["pub_key"] = o.PubKey
	}
	return json.Marshal(toSerialize)
}

type NullableKubernetesAccessRules struct {
	value *KubernetesAccessRules
	isSet bool
}

func (v NullableKubernetesAccessRules) Get() *KubernetesAccessRules {
	return v.value
}

func (v *NullableKubernetesAccessRules) Set(val *KubernetesAccessRules) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesAccessRules) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesAccessRules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesAccessRules(val *KubernetesAccessRules) *NullableKubernetesAccessRules {
	return &NullableKubernetesAccessRules{value: val, isSet: true}
}

func (v NullableKubernetesAccessRules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesAccessRules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


