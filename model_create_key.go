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

// CreateKey createKey is a command that creates a new key. [Deprecated: Use command create-dfc-key]
type CreateKey struct {
	// Key type; options: [AES128GCM, AES256GCM, AES128SIV, AES256SIV, AES128CBC, AES256CBC, RSA1024, RSA2048, RSA3072, RSA4096]
	Alg string `json:"alg"`
	// Common name for the generated certificate. Relevant only for generate-self-signed-certificate.
	CertificateCommonName *string `json:"certificate-common-name,omitempty"`
	// Country name for the generated certificate. Relevant only for generate-self-signed-certificate.
	CertificateCountry *string `json:"certificate-country,omitempty"`
	// Digest algorithm to be used for the certificate key signing. Currently, we support only \"sha256\" so we hide this option for CLI.
	CertificateDigestAlgo *string `json:"certificate-digest-algo,omitempty"`
	// Locality for the generated certificate. Relevant only for generate-self-signed-certificate.
	CertificateLocality *string `json:"certificate-locality,omitempty"`
	// Organization name for the generated certificate. Relevant only for generate-self-signed-certificate.
	CertificateOrganization *string `json:"certificate-organization,omitempty"`
	// Province name for the generated certificate. Relevant only for generate-self-signed-certificate.
	CertificateProvince *string `json:"certificate-province,omitempty"`
	// TTL in days for the generated certificate. Required only for generate-self-signed-certificate.
	CertificateTtl *int64 `json:"certificate-ttl,omitempty"`
	// The csr config data in base64 encoding
	ConfFileData *string `json:"conf-file-data,omitempty"`
	// The customer fragment ID that will be used to create the key (if empty, the key will be created independently of a customer fragment)
	CustomerFrgId *string `json:"customer-frg-id,omitempty"`
	// Protection from accidental deletion of this item [true/false]
	DeleteProtection *string `json:"delete_protection,omitempty"`
	// Description of the object
	Description *string `json:"description,omitempty"`
	// Whether to generate a self signed certificate with the key. If set, --certificate-ttl must be provided.
	GenerateSelfSignedCertificate *bool `json:"generate-self-signed-certificate,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Deprecated - use description
	Metadata *string `json:"metadata,omitempty"`
	// Key name
	Name string `json:"name"`
	// The number of fragments that the item will be split into (not includes customer fragment)
	SplitLevel *int64 `json:"split-level,omitempty"`
	// List of the tags attached to this key
	Tag *[]string `json:"tag,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewCreateKey instantiates a new CreateKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateKey(alg string, name string, ) *CreateKey {
	this := CreateKey{}
	this.Alg = alg
	var json bool = false
	this.Json = &json
	this.Name = name
	var splitLevel int64 = 3
	this.SplitLevel = &splitLevel
	return &this
}

// NewCreateKeyWithDefaults instantiates a new CreateKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateKeyWithDefaults() *CreateKey {
	this := CreateKey{}
	var json bool = false
	this.Json = &json
	var splitLevel int64 = 3
	this.SplitLevel = &splitLevel
	return &this
}

// GetAlg returns the Alg field value
func (o *CreateKey) GetAlg() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Alg
}

// GetAlgOk returns a tuple with the Alg field value
// and a boolean to check if the value has been set.
func (o *CreateKey) GetAlgOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Alg, true
}

// SetAlg sets field value
func (o *CreateKey) SetAlg(v string) {
	o.Alg = v
}

// GetCertificateCommonName returns the CertificateCommonName field value if set, zero value otherwise.
func (o *CreateKey) GetCertificateCommonName() string {
	if o == nil || o.CertificateCommonName == nil {
		var ret string
		return ret
	}
	return *o.CertificateCommonName
}

// GetCertificateCommonNameOk returns a tuple with the CertificateCommonName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateKey) GetCertificateCommonNameOk() (*string, bool) {
	if o == nil || o.CertificateCommonName == nil {
		return nil, false
	}
	return o.CertificateCommonName, true
}

// HasCertificateCommonName returns a boolean if a field has been set.
func (o *CreateKey) HasCertificateCommonName() bool {
	if o != nil && o.CertificateCommonName != nil {
		return true
	}

	return false
}

// SetCertificateCommonName gets a reference to the given string and assigns it to the CertificateCommonName field.
func (o *CreateKey) SetCertificateCommonName(v string) {
	o.CertificateCommonName = &v
}

// GetCertificateCountry returns the CertificateCountry field value if set, zero value otherwise.
func (o *CreateKey) GetCertificateCountry() string {
	if o == nil || o.CertificateCountry == nil {
		var ret string
		return ret
	}
	return *o.CertificateCountry
}

// GetCertificateCountryOk returns a tuple with the CertificateCountry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateKey) GetCertificateCountryOk() (*string, bool) {
	if o == nil || o.CertificateCountry == nil {
		return nil, false
	}
	return o.CertificateCountry, true
}

// HasCertificateCountry returns a boolean if a field has been set.
func (o *CreateKey) HasCertificateCountry() bool {
	if o != nil && o.CertificateCountry != nil {
		return true
	}

	return false
}

// SetCertificateCountry gets a reference to the given string and assigns it to the CertificateCountry field.
func (o *CreateKey) SetCertificateCountry(v string) {
	o.CertificateCountry = &v
}

// GetCertificateDigestAlgo returns the CertificateDigestAlgo field value if set, zero value otherwise.
func (o *CreateKey) GetCertificateDigestAlgo() string {
	if o == nil || o.CertificateDigestAlgo == nil {
		var ret string
		return ret
	}
	return *o.CertificateDigestAlgo
}

// GetCertificateDigestAlgoOk returns a tuple with the CertificateDigestAlgo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateKey) GetCertificateDigestAlgoOk() (*string, bool) {
	if o == nil || o.CertificateDigestAlgo == nil {
		return nil, false
	}
	return o.CertificateDigestAlgo, true
}

// HasCertificateDigestAlgo returns a boolean if a field has been set.
func (o *CreateKey) HasCertificateDigestAlgo() bool {
	if o != nil && o.CertificateDigestAlgo != nil {
		return true
	}

	return false
}

// SetCertificateDigestAlgo gets a reference to the given string and assigns it to the CertificateDigestAlgo field.
func (o *CreateKey) SetCertificateDigestAlgo(v string) {
	o.CertificateDigestAlgo = &v
}

// GetCertificateLocality returns the CertificateLocality field value if set, zero value otherwise.
func (o *CreateKey) GetCertificateLocality() string {
	if o == nil || o.CertificateLocality == nil {
		var ret string
		return ret
	}
	return *o.CertificateLocality
}

// GetCertificateLocalityOk returns a tuple with the CertificateLocality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateKey) GetCertificateLocalityOk() (*string, bool) {
	if o == nil || o.CertificateLocality == nil {
		return nil, false
	}
	return o.CertificateLocality, true
}

// HasCertificateLocality returns a boolean if a field has been set.
func (o *CreateKey) HasCertificateLocality() bool {
	if o != nil && o.CertificateLocality != nil {
		return true
	}

	return false
}

// SetCertificateLocality gets a reference to the given string and assigns it to the CertificateLocality field.
func (o *CreateKey) SetCertificateLocality(v string) {
	o.CertificateLocality = &v
}

// GetCertificateOrganization returns the CertificateOrganization field value if set, zero value otherwise.
func (o *CreateKey) GetCertificateOrganization() string {
	if o == nil || o.CertificateOrganization == nil {
		var ret string
		return ret
	}
	return *o.CertificateOrganization
}

// GetCertificateOrganizationOk returns a tuple with the CertificateOrganization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateKey) GetCertificateOrganizationOk() (*string, bool) {
	if o == nil || o.CertificateOrganization == nil {
		return nil, false
	}
	return o.CertificateOrganization, true
}

// HasCertificateOrganization returns a boolean if a field has been set.
func (o *CreateKey) HasCertificateOrganization() bool {
	if o != nil && o.CertificateOrganization != nil {
		return true
	}

	return false
}

// SetCertificateOrganization gets a reference to the given string and assigns it to the CertificateOrganization field.
func (o *CreateKey) SetCertificateOrganization(v string) {
	o.CertificateOrganization = &v
}

// GetCertificateProvince returns the CertificateProvince field value if set, zero value otherwise.
func (o *CreateKey) GetCertificateProvince() string {
	if o == nil || o.CertificateProvince == nil {
		var ret string
		return ret
	}
	return *o.CertificateProvince
}

// GetCertificateProvinceOk returns a tuple with the CertificateProvince field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateKey) GetCertificateProvinceOk() (*string, bool) {
	if o == nil || o.CertificateProvince == nil {
		return nil, false
	}
	return o.CertificateProvince, true
}

// HasCertificateProvince returns a boolean if a field has been set.
func (o *CreateKey) HasCertificateProvince() bool {
	if o != nil && o.CertificateProvince != nil {
		return true
	}

	return false
}

// SetCertificateProvince gets a reference to the given string and assigns it to the CertificateProvince field.
func (o *CreateKey) SetCertificateProvince(v string) {
	o.CertificateProvince = &v
}

// GetCertificateTtl returns the CertificateTtl field value if set, zero value otherwise.
func (o *CreateKey) GetCertificateTtl() int64 {
	if o == nil || o.CertificateTtl == nil {
		var ret int64
		return ret
	}
	return *o.CertificateTtl
}

// GetCertificateTtlOk returns a tuple with the CertificateTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateKey) GetCertificateTtlOk() (*int64, bool) {
	if o == nil || o.CertificateTtl == nil {
		return nil, false
	}
	return o.CertificateTtl, true
}

// HasCertificateTtl returns a boolean if a field has been set.
func (o *CreateKey) HasCertificateTtl() bool {
	if o != nil && o.CertificateTtl != nil {
		return true
	}

	return false
}

// SetCertificateTtl gets a reference to the given int64 and assigns it to the CertificateTtl field.
func (o *CreateKey) SetCertificateTtl(v int64) {
	o.CertificateTtl = &v
}

// GetConfFileData returns the ConfFileData field value if set, zero value otherwise.
func (o *CreateKey) GetConfFileData() string {
	if o == nil || o.ConfFileData == nil {
		var ret string
		return ret
	}
	return *o.ConfFileData
}

// GetConfFileDataOk returns a tuple with the ConfFileData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateKey) GetConfFileDataOk() (*string, bool) {
	if o == nil || o.ConfFileData == nil {
		return nil, false
	}
	return o.ConfFileData, true
}

// HasConfFileData returns a boolean if a field has been set.
func (o *CreateKey) HasConfFileData() bool {
	if o != nil && o.ConfFileData != nil {
		return true
	}

	return false
}

// SetConfFileData gets a reference to the given string and assigns it to the ConfFileData field.
func (o *CreateKey) SetConfFileData(v string) {
	o.ConfFileData = &v
}

// GetCustomerFrgId returns the CustomerFrgId field value if set, zero value otherwise.
func (o *CreateKey) GetCustomerFrgId() string {
	if o == nil || o.CustomerFrgId == nil {
		var ret string
		return ret
	}
	return *o.CustomerFrgId
}

// GetCustomerFrgIdOk returns a tuple with the CustomerFrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateKey) GetCustomerFrgIdOk() (*string, bool) {
	if o == nil || o.CustomerFrgId == nil {
		return nil, false
	}
	return o.CustomerFrgId, true
}

// HasCustomerFrgId returns a boolean if a field has been set.
func (o *CreateKey) HasCustomerFrgId() bool {
	if o != nil && o.CustomerFrgId != nil {
		return true
	}

	return false
}

// SetCustomerFrgId gets a reference to the given string and assigns it to the CustomerFrgId field.
func (o *CreateKey) SetCustomerFrgId(v string) {
	o.CustomerFrgId = &v
}

// GetDeleteProtection returns the DeleteProtection field value if set, zero value otherwise.
func (o *CreateKey) GetDeleteProtection() string {
	if o == nil || o.DeleteProtection == nil {
		var ret string
		return ret
	}
	return *o.DeleteProtection
}

// GetDeleteProtectionOk returns a tuple with the DeleteProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateKey) GetDeleteProtectionOk() (*string, bool) {
	if o == nil || o.DeleteProtection == nil {
		return nil, false
	}
	return o.DeleteProtection, true
}

// HasDeleteProtection returns a boolean if a field has been set.
func (o *CreateKey) HasDeleteProtection() bool {
	if o != nil && o.DeleteProtection != nil {
		return true
	}

	return false
}

// SetDeleteProtection gets a reference to the given string and assigns it to the DeleteProtection field.
func (o *CreateKey) SetDeleteProtection(v string) {
	o.DeleteProtection = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateKey) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateKey) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateKey) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateKey) SetDescription(v string) {
	o.Description = &v
}

// GetGenerateSelfSignedCertificate returns the GenerateSelfSignedCertificate field value if set, zero value otherwise.
func (o *CreateKey) GetGenerateSelfSignedCertificate() bool {
	if o == nil || o.GenerateSelfSignedCertificate == nil {
		var ret bool
		return ret
	}
	return *o.GenerateSelfSignedCertificate
}

// GetGenerateSelfSignedCertificateOk returns a tuple with the GenerateSelfSignedCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateKey) GetGenerateSelfSignedCertificateOk() (*bool, bool) {
	if o == nil || o.GenerateSelfSignedCertificate == nil {
		return nil, false
	}
	return o.GenerateSelfSignedCertificate, true
}

// HasGenerateSelfSignedCertificate returns a boolean if a field has been set.
func (o *CreateKey) HasGenerateSelfSignedCertificate() bool {
	if o != nil && o.GenerateSelfSignedCertificate != nil {
		return true
	}

	return false
}

// SetGenerateSelfSignedCertificate gets a reference to the given bool and assigns it to the GenerateSelfSignedCertificate field.
func (o *CreateKey) SetGenerateSelfSignedCertificate(v bool) {
	o.GenerateSelfSignedCertificate = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *CreateKey) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateKey) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *CreateKey) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *CreateKey) SetJson(v bool) {
	o.Json = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *CreateKey) GetMetadata() string {
	if o == nil || o.Metadata == nil {
		var ret string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateKey) GetMetadataOk() (*string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CreateKey) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given string and assigns it to the Metadata field.
func (o *CreateKey) SetMetadata(v string) {
	o.Metadata = &v
}

// GetName returns the Name field value
func (o *CreateKey) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateKey) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateKey) SetName(v string) {
	o.Name = v
}

// GetSplitLevel returns the SplitLevel field value if set, zero value otherwise.
func (o *CreateKey) GetSplitLevel() int64 {
	if o == nil || o.SplitLevel == nil {
		var ret int64
		return ret
	}
	return *o.SplitLevel
}

// GetSplitLevelOk returns a tuple with the SplitLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateKey) GetSplitLevelOk() (*int64, bool) {
	if o == nil || o.SplitLevel == nil {
		return nil, false
	}
	return o.SplitLevel, true
}

// HasSplitLevel returns a boolean if a field has been set.
func (o *CreateKey) HasSplitLevel() bool {
	if o != nil && o.SplitLevel != nil {
		return true
	}

	return false
}

// SetSplitLevel gets a reference to the given int64 and assigns it to the SplitLevel field.
func (o *CreateKey) SetSplitLevel(v int64) {
	o.SplitLevel = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *CreateKey) GetTag() []string {
	if o == nil || o.Tag == nil {
		var ret []string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateKey) GetTagOk() (*[]string, bool) {
	if o == nil || o.Tag == nil {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *CreateKey) HasTag() bool {
	if o != nil && o.Tag != nil {
		return true
	}

	return false
}

// SetTag gets a reference to the given []string and assigns it to the Tag field.
func (o *CreateKey) SetTag(v []string) {
	o.Tag = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *CreateKey) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateKey) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *CreateKey) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *CreateKey) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *CreateKey) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateKey) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *CreateKey) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *CreateKey) SetUidToken(v string) {
	o.UidToken = &v
}

func (o CreateKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["alg"] = o.Alg
	}
	if o.CertificateCommonName != nil {
		toSerialize["certificate-common-name"] = o.CertificateCommonName
	}
	if o.CertificateCountry != nil {
		toSerialize["certificate-country"] = o.CertificateCountry
	}
	if o.CertificateDigestAlgo != nil {
		toSerialize["certificate-digest-algo"] = o.CertificateDigestAlgo
	}
	if o.CertificateLocality != nil {
		toSerialize["certificate-locality"] = o.CertificateLocality
	}
	if o.CertificateOrganization != nil {
		toSerialize["certificate-organization"] = o.CertificateOrganization
	}
	if o.CertificateProvince != nil {
		toSerialize["certificate-province"] = o.CertificateProvince
	}
	if o.CertificateTtl != nil {
		toSerialize["certificate-ttl"] = o.CertificateTtl
	}
	if o.ConfFileData != nil {
		toSerialize["conf-file-data"] = o.ConfFileData
	}
	if o.CustomerFrgId != nil {
		toSerialize["customer-frg-id"] = o.CustomerFrgId
	}
	if o.DeleteProtection != nil {
		toSerialize["delete_protection"] = o.DeleteProtection
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.GenerateSelfSignedCertificate != nil {
		toSerialize["generate-self-signed-certificate"] = o.GenerateSelfSignedCertificate
	}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.SplitLevel != nil {
		toSerialize["split-level"] = o.SplitLevel
	}
	if o.Tag != nil {
		toSerialize["tag"] = o.Tag
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableCreateKey struct {
	value *CreateKey
	isSet bool
}

func (v NullableCreateKey) Get() *CreateKey {
	return v.value
}

func (v *NullableCreateKey) Set(val *CreateKey) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateKey) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateKey(val *CreateKey) *NullableCreateKey {
	return &NullableCreateKey{value: val, isSet: true}
}

func (v NullableCreateKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


