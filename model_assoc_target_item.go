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

// AssocTargetItem assocTargetItem is a command that creates an association between target and item.
type AssocTargetItem struct {
	// Automatically disable previous key version (required for azure targets)
	DisablePreviousKeyVersion *bool `json:"disable-previous-key-version,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// A list of allowed operations for the key (required for azure targets)
	KeyOperations *[]string `json:"key-operations,omitempty"`
	// Keyring name of the GCP KMS (required for gcp targets)
	KeyringName *string `json:"keyring-name,omitempty"`
	// Algorithm of the key in GCP KMS (required for gcp targets)
	KmsAlgorithm *string `json:"kms-algorithm,omitempty"`
	// Location id of the GCP KMS (required for gcp targets)
	LocationId *string `json:"location-id,omitempty"`
	// Set to 'true' to create a multi region managed key (relevant for aws targets)
	MultiRegion *string `json:"multi-region,omitempty"`
	// The item to associate
	Name string `json:"name"`
	// Project id of the GCP KMS (required for gcp targets)
	ProjectId *string `json:"project-id,omitempty"`
	// Purpose of the key in GCP KMS (required for gcp targets)
	Purpose *string `json:"purpose,omitempty"`
	// The list of regions to create a copy of the key in (relevant for aws targets)
	Regions *[]string `json:"regions,omitempty"`
	// The target to associate
	TargetName string `json:"target-name"`
	// The tenant secret type [Data/SearchIndex/Analytics] (required for salesforce targets)
	TenantSecretType *string `json:"tenant-secret-type,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	// Name of the vault used (required for azure targets)
	VaultName *string `json:"vault-name,omitempty"`
}

// NewAssocTargetItem instantiates a new AssocTargetItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssocTargetItem(name string, targetName string, ) *AssocTargetItem {
	this := AssocTargetItem{}
	var multiRegion string = "false"
	this.MultiRegion = &multiRegion
	this.Name = name
	this.TargetName = targetName
	return &this
}

// NewAssocTargetItemWithDefaults instantiates a new AssocTargetItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssocTargetItemWithDefaults() *AssocTargetItem {
	this := AssocTargetItem{}
	var multiRegion string = "false"
	this.MultiRegion = &multiRegion
	return &this
}

// GetDisablePreviousKeyVersion returns the DisablePreviousKeyVersion field value if set, zero value otherwise.
func (o *AssocTargetItem) GetDisablePreviousKeyVersion() bool {
	if o == nil || o.DisablePreviousKeyVersion == nil {
		var ret bool
		return ret
	}
	return *o.DisablePreviousKeyVersion
}

// GetDisablePreviousKeyVersionOk returns a tuple with the DisablePreviousKeyVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssocTargetItem) GetDisablePreviousKeyVersionOk() (*bool, bool) {
	if o == nil || o.DisablePreviousKeyVersion == nil {
		return nil, false
	}
	return o.DisablePreviousKeyVersion, true
}

// HasDisablePreviousKeyVersion returns a boolean if a field has been set.
func (o *AssocTargetItem) HasDisablePreviousKeyVersion() bool {
	if o != nil && o.DisablePreviousKeyVersion != nil {
		return true
	}

	return false
}

// SetDisablePreviousKeyVersion gets a reference to the given bool and assigns it to the DisablePreviousKeyVersion field.
func (o *AssocTargetItem) SetDisablePreviousKeyVersion(v bool) {
	o.DisablePreviousKeyVersion = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *AssocTargetItem) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssocTargetItem) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *AssocTargetItem) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *AssocTargetItem) SetJson(v bool) {
	o.Json = &v
}

// GetKeyOperations returns the KeyOperations field value if set, zero value otherwise.
func (o *AssocTargetItem) GetKeyOperations() []string {
	if o == nil || o.KeyOperations == nil {
		var ret []string
		return ret
	}
	return *o.KeyOperations
}

// GetKeyOperationsOk returns a tuple with the KeyOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssocTargetItem) GetKeyOperationsOk() (*[]string, bool) {
	if o == nil || o.KeyOperations == nil {
		return nil, false
	}
	return o.KeyOperations, true
}

// HasKeyOperations returns a boolean if a field has been set.
func (o *AssocTargetItem) HasKeyOperations() bool {
	if o != nil && o.KeyOperations != nil {
		return true
	}

	return false
}

// SetKeyOperations gets a reference to the given []string and assigns it to the KeyOperations field.
func (o *AssocTargetItem) SetKeyOperations(v []string) {
	o.KeyOperations = &v
}

// GetKeyringName returns the KeyringName field value if set, zero value otherwise.
func (o *AssocTargetItem) GetKeyringName() string {
	if o == nil || o.KeyringName == nil {
		var ret string
		return ret
	}
	return *o.KeyringName
}

// GetKeyringNameOk returns a tuple with the KeyringName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssocTargetItem) GetKeyringNameOk() (*string, bool) {
	if o == nil || o.KeyringName == nil {
		return nil, false
	}
	return o.KeyringName, true
}

// HasKeyringName returns a boolean if a field has been set.
func (o *AssocTargetItem) HasKeyringName() bool {
	if o != nil && o.KeyringName != nil {
		return true
	}

	return false
}

// SetKeyringName gets a reference to the given string and assigns it to the KeyringName field.
func (o *AssocTargetItem) SetKeyringName(v string) {
	o.KeyringName = &v
}

// GetKmsAlgorithm returns the KmsAlgorithm field value if set, zero value otherwise.
func (o *AssocTargetItem) GetKmsAlgorithm() string {
	if o == nil || o.KmsAlgorithm == nil {
		var ret string
		return ret
	}
	return *o.KmsAlgorithm
}

// GetKmsAlgorithmOk returns a tuple with the KmsAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssocTargetItem) GetKmsAlgorithmOk() (*string, bool) {
	if o == nil || o.KmsAlgorithm == nil {
		return nil, false
	}
	return o.KmsAlgorithm, true
}

// HasKmsAlgorithm returns a boolean if a field has been set.
func (o *AssocTargetItem) HasKmsAlgorithm() bool {
	if o != nil && o.KmsAlgorithm != nil {
		return true
	}

	return false
}

// SetKmsAlgorithm gets a reference to the given string and assigns it to the KmsAlgorithm field.
func (o *AssocTargetItem) SetKmsAlgorithm(v string) {
	o.KmsAlgorithm = &v
}

// GetLocationId returns the LocationId field value if set, zero value otherwise.
func (o *AssocTargetItem) GetLocationId() string {
	if o == nil || o.LocationId == nil {
		var ret string
		return ret
	}
	return *o.LocationId
}

// GetLocationIdOk returns a tuple with the LocationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssocTargetItem) GetLocationIdOk() (*string, bool) {
	if o == nil || o.LocationId == nil {
		return nil, false
	}
	return o.LocationId, true
}

// HasLocationId returns a boolean if a field has been set.
func (o *AssocTargetItem) HasLocationId() bool {
	if o != nil && o.LocationId != nil {
		return true
	}

	return false
}

// SetLocationId gets a reference to the given string and assigns it to the LocationId field.
func (o *AssocTargetItem) SetLocationId(v string) {
	o.LocationId = &v
}

// GetMultiRegion returns the MultiRegion field value if set, zero value otherwise.
func (o *AssocTargetItem) GetMultiRegion() string {
	if o == nil || o.MultiRegion == nil {
		var ret string
		return ret
	}
	return *o.MultiRegion
}

// GetMultiRegionOk returns a tuple with the MultiRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssocTargetItem) GetMultiRegionOk() (*string, bool) {
	if o == nil || o.MultiRegion == nil {
		return nil, false
	}
	return o.MultiRegion, true
}

// HasMultiRegion returns a boolean if a field has been set.
func (o *AssocTargetItem) HasMultiRegion() bool {
	if o != nil && o.MultiRegion != nil {
		return true
	}

	return false
}

// SetMultiRegion gets a reference to the given string and assigns it to the MultiRegion field.
func (o *AssocTargetItem) SetMultiRegion(v string) {
	o.MultiRegion = &v
}

// GetName returns the Name field value
func (o *AssocTargetItem) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AssocTargetItem) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AssocTargetItem) SetName(v string) {
	o.Name = v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *AssocTargetItem) GetProjectId() string {
	if o == nil || o.ProjectId == nil {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssocTargetItem) GetProjectIdOk() (*string, bool) {
	if o == nil || o.ProjectId == nil {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *AssocTargetItem) HasProjectId() bool {
	if o != nil && o.ProjectId != nil {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *AssocTargetItem) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetPurpose returns the Purpose field value if set, zero value otherwise.
func (o *AssocTargetItem) GetPurpose() string {
	if o == nil || o.Purpose == nil {
		var ret string
		return ret
	}
	return *o.Purpose
}

// GetPurposeOk returns a tuple with the Purpose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssocTargetItem) GetPurposeOk() (*string, bool) {
	if o == nil || o.Purpose == nil {
		return nil, false
	}
	return o.Purpose, true
}

// HasPurpose returns a boolean if a field has been set.
func (o *AssocTargetItem) HasPurpose() bool {
	if o != nil && o.Purpose != nil {
		return true
	}

	return false
}

// SetPurpose gets a reference to the given string and assigns it to the Purpose field.
func (o *AssocTargetItem) SetPurpose(v string) {
	o.Purpose = &v
}

// GetRegions returns the Regions field value if set, zero value otherwise.
func (o *AssocTargetItem) GetRegions() []string {
	if o == nil || o.Regions == nil {
		var ret []string
		return ret
	}
	return *o.Regions
}

// GetRegionsOk returns a tuple with the Regions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssocTargetItem) GetRegionsOk() (*[]string, bool) {
	if o == nil || o.Regions == nil {
		return nil, false
	}
	return o.Regions, true
}

// HasRegions returns a boolean if a field has been set.
func (o *AssocTargetItem) HasRegions() bool {
	if o != nil && o.Regions != nil {
		return true
	}

	return false
}

// SetRegions gets a reference to the given []string and assigns it to the Regions field.
func (o *AssocTargetItem) SetRegions(v []string) {
	o.Regions = &v
}

// GetTargetName returns the TargetName field value
func (o *AssocTargetItem) GetTargetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value
// and a boolean to check if the value has been set.
func (o *AssocTargetItem) GetTargetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TargetName, true
}

// SetTargetName sets field value
func (o *AssocTargetItem) SetTargetName(v string) {
	o.TargetName = v
}

// GetTenantSecretType returns the TenantSecretType field value if set, zero value otherwise.
func (o *AssocTargetItem) GetTenantSecretType() string {
	if o == nil || o.TenantSecretType == nil {
		var ret string
		return ret
	}
	return *o.TenantSecretType
}

// GetTenantSecretTypeOk returns a tuple with the TenantSecretType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssocTargetItem) GetTenantSecretTypeOk() (*string, bool) {
	if o == nil || o.TenantSecretType == nil {
		return nil, false
	}
	return o.TenantSecretType, true
}

// HasTenantSecretType returns a boolean if a field has been set.
func (o *AssocTargetItem) HasTenantSecretType() bool {
	if o != nil && o.TenantSecretType != nil {
		return true
	}

	return false
}

// SetTenantSecretType gets a reference to the given string and assigns it to the TenantSecretType field.
func (o *AssocTargetItem) SetTenantSecretType(v string) {
	o.TenantSecretType = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *AssocTargetItem) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssocTargetItem) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *AssocTargetItem) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *AssocTargetItem) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *AssocTargetItem) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssocTargetItem) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *AssocTargetItem) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *AssocTargetItem) SetUidToken(v string) {
	o.UidToken = &v
}

// GetVaultName returns the VaultName field value if set, zero value otherwise.
func (o *AssocTargetItem) GetVaultName() string {
	if o == nil || o.VaultName == nil {
		var ret string
		return ret
	}
	return *o.VaultName
}

// GetVaultNameOk returns a tuple with the VaultName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssocTargetItem) GetVaultNameOk() (*string, bool) {
	if o == nil || o.VaultName == nil {
		return nil, false
	}
	return o.VaultName, true
}

// HasVaultName returns a boolean if a field has been set.
func (o *AssocTargetItem) HasVaultName() bool {
	if o != nil && o.VaultName != nil {
		return true
	}

	return false
}

// SetVaultName gets a reference to the given string and assigns it to the VaultName field.
func (o *AssocTargetItem) SetVaultName(v string) {
	o.VaultName = &v
}

func (o AssocTargetItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DisablePreviousKeyVersion != nil {
		toSerialize["disable-previous-key-version"] = o.DisablePreviousKeyVersion
	}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if o.KeyOperations != nil {
		toSerialize["key-operations"] = o.KeyOperations
	}
	if o.KeyringName != nil {
		toSerialize["keyring-name"] = o.KeyringName
	}
	if o.KmsAlgorithm != nil {
		toSerialize["kms-algorithm"] = o.KmsAlgorithm
	}
	if o.LocationId != nil {
		toSerialize["location-id"] = o.LocationId
	}
	if o.MultiRegion != nil {
		toSerialize["multi-region"] = o.MultiRegion
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.ProjectId != nil {
		toSerialize["project-id"] = o.ProjectId
	}
	if o.Purpose != nil {
		toSerialize["purpose"] = o.Purpose
	}
	if o.Regions != nil {
		toSerialize["regions"] = o.Regions
	}
	if true {
		toSerialize["target-name"] = o.TargetName
	}
	if o.TenantSecretType != nil {
		toSerialize["tenant-secret-type"] = o.TenantSecretType
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	if o.VaultName != nil {
		toSerialize["vault-name"] = o.VaultName
	}
	return json.Marshal(toSerialize)
}

type NullableAssocTargetItem struct {
	value *AssocTargetItem
	isSet bool
}

func (v NullableAssocTargetItem) Get() *AssocTargetItem {
	return v.value
}

func (v *NullableAssocTargetItem) Set(val *AssocTargetItem) {
	v.value = val
	v.isSet = true
}

func (v NullableAssocTargetItem) IsSet() bool {
	return v.isSet
}

func (v *NullableAssocTargetItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssocTargetItem(val *AssocTargetItem) *NullableAssocTargetItem {
	return &NullableAssocTargetItem{value: val, isSet: true}
}

func (v NullableAssocTargetItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssocTargetItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


