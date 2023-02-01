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

// PKICertificateIssueDetails struct for PKICertificateIssueDetails
type PKICertificateIssueDetails struct {
	AllowAnyName *bool `json:"allow_any_name,omitempty"`
	AllowSubdomains *bool `json:"allow_subdomains,omitempty"`
	AllowedDomainsList *[]string `json:"allowed_domains_list,omitempty"`
	AllowedUriSans *[]string `json:"allowed_uri_sans,omitempty"`
	BasicConstraintsValidForNonCa *bool `json:"basic_constraints_valid_for_non_ca,omitempty"`
	ClientFlag *bool `json:"client_flag,omitempty"`
	CodeSigningFlag *bool `json:"code_signing_flag,omitempty"`
	Country *[]string `json:"country,omitempty"`
	EnforceHostnames *bool `json:"enforce_hostnames,omitempty"`
	IsCa *bool `json:"is_ca,omitempty"`
	KeyBits *int64 `json:"key_bits,omitempty"`
	KeyType *string `json:"key_type,omitempty"`
	KeyUsageList *[]string `json:"key_usage_list,omitempty"`
	Locality *[]string `json:"locality,omitempty"`
	// A Duration represents the elapsed time between two instants as an int64 nanosecond count. The representation limits the largest representable duration to approximately 290 years.
	NotBeforeDuration *int64 `json:"not_before_duration,omitempty"`
	OrganizationList *[]string `json:"organization_list,omitempty"`
	OrganizationUnitList *[]string `json:"organization_unit_list,omitempty"`
	PostalCode *[]string `json:"postal_code,omitempty"`
	Province *[]string `json:"province,omitempty"`
	RequireCn *bool `json:"require_cn,omitempty"`
	ServerFlag *bool `json:"server_flag,omitempty"`
	StreetAddress *[]string `json:"street_address,omitempty"`
}

// NewPKICertificateIssueDetails instantiates a new PKICertificateIssueDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPKICertificateIssueDetails() *PKICertificateIssueDetails {
	this := PKICertificateIssueDetails{}
	return &this
}

// NewPKICertificateIssueDetailsWithDefaults instantiates a new PKICertificateIssueDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPKICertificateIssueDetailsWithDefaults() *PKICertificateIssueDetails {
	this := PKICertificateIssueDetails{}
	return &this
}

// GetAllowAnyName returns the AllowAnyName field value if set, zero value otherwise.
func (o *PKICertificateIssueDetails) GetAllowAnyName() bool {
	if o == nil || o.AllowAnyName == nil {
		var ret bool
		return ret
	}
	return *o.AllowAnyName
}

// GetAllowAnyNameOk returns a tuple with the AllowAnyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PKICertificateIssueDetails) GetAllowAnyNameOk() (*bool, bool) {
	if o == nil || o.AllowAnyName == nil {
		return nil, false
	}
	return o.AllowAnyName, true
}

// HasAllowAnyName returns a boolean if a field has been set.
func (o *PKICertificateIssueDetails) HasAllowAnyName() bool {
	if o != nil && o.AllowAnyName != nil {
		return true
	}

	return false
}

// SetAllowAnyName gets a reference to the given bool and assigns it to the AllowAnyName field.
func (o *PKICertificateIssueDetails) SetAllowAnyName(v bool) {
	o.AllowAnyName = &v
}

// GetAllowSubdomains returns the AllowSubdomains field value if set, zero value otherwise.
func (o *PKICertificateIssueDetails) GetAllowSubdomains() bool {
	if o == nil || o.AllowSubdomains == nil {
		var ret bool
		return ret
	}
	return *o.AllowSubdomains
}

// GetAllowSubdomainsOk returns a tuple with the AllowSubdomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PKICertificateIssueDetails) GetAllowSubdomainsOk() (*bool, bool) {
	if o == nil || o.AllowSubdomains == nil {
		return nil, false
	}
	return o.AllowSubdomains, true
}

// HasAllowSubdomains returns a boolean if a field has been set.
func (o *PKICertificateIssueDetails) HasAllowSubdomains() bool {
	if o != nil && o.AllowSubdomains != nil {
		return true
	}

	return false
}

// SetAllowSubdomains gets a reference to the given bool and assigns it to the AllowSubdomains field.
func (o *PKICertificateIssueDetails) SetAllowSubdomains(v bool) {
	o.AllowSubdomains = &v
}

// GetAllowedDomainsList returns the AllowedDomainsList field value if set, zero value otherwise.
func (o *PKICertificateIssueDetails) GetAllowedDomainsList() []string {
	if o == nil || o.AllowedDomainsList == nil {
		var ret []string
		return ret
	}
	return *o.AllowedDomainsList
}

// GetAllowedDomainsListOk returns a tuple with the AllowedDomainsList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PKICertificateIssueDetails) GetAllowedDomainsListOk() (*[]string, bool) {
	if o == nil || o.AllowedDomainsList == nil {
		return nil, false
	}
	return o.AllowedDomainsList, true
}

// HasAllowedDomainsList returns a boolean if a field has been set.
func (o *PKICertificateIssueDetails) HasAllowedDomainsList() bool {
	if o != nil && o.AllowedDomainsList != nil {
		return true
	}

	return false
}

// SetAllowedDomainsList gets a reference to the given []string and assigns it to the AllowedDomainsList field.
func (o *PKICertificateIssueDetails) SetAllowedDomainsList(v []string) {
	o.AllowedDomainsList = &v
}

// GetAllowedUriSans returns the AllowedUriSans field value if set, zero value otherwise.
func (o *PKICertificateIssueDetails) GetAllowedUriSans() []string {
	if o == nil || o.AllowedUriSans == nil {
		var ret []string
		return ret
	}
	return *o.AllowedUriSans
}

// GetAllowedUriSansOk returns a tuple with the AllowedUriSans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PKICertificateIssueDetails) GetAllowedUriSansOk() (*[]string, bool) {
	if o == nil || o.AllowedUriSans == nil {
		return nil, false
	}
	return o.AllowedUriSans, true
}

// HasAllowedUriSans returns a boolean if a field has been set.
func (o *PKICertificateIssueDetails) HasAllowedUriSans() bool {
	if o != nil && o.AllowedUriSans != nil {
		return true
	}

	return false
}

// SetAllowedUriSans gets a reference to the given []string and assigns it to the AllowedUriSans field.
func (o *PKICertificateIssueDetails) SetAllowedUriSans(v []string) {
	o.AllowedUriSans = &v
}

// GetBasicConstraintsValidForNonCa returns the BasicConstraintsValidForNonCa field value if set, zero value otherwise.
func (o *PKICertificateIssueDetails) GetBasicConstraintsValidForNonCa() bool {
	if o == nil || o.BasicConstraintsValidForNonCa == nil {
		var ret bool
		return ret
	}
	return *o.BasicConstraintsValidForNonCa
}

// GetBasicConstraintsValidForNonCaOk returns a tuple with the BasicConstraintsValidForNonCa field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PKICertificateIssueDetails) GetBasicConstraintsValidForNonCaOk() (*bool, bool) {
	if o == nil || o.BasicConstraintsValidForNonCa == nil {
		return nil, false
	}
	return o.BasicConstraintsValidForNonCa, true
}

// HasBasicConstraintsValidForNonCa returns a boolean if a field has been set.
func (o *PKICertificateIssueDetails) HasBasicConstraintsValidForNonCa() bool {
	if o != nil && o.BasicConstraintsValidForNonCa != nil {
		return true
	}

	return false
}

// SetBasicConstraintsValidForNonCa gets a reference to the given bool and assigns it to the BasicConstraintsValidForNonCa field.
func (o *PKICertificateIssueDetails) SetBasicConstraintsValidForNonCa(v bool) {
	o.BasicConstraintsValidForNonCa = &v
}

// GetClientFlag returns the ClientFlag field value if set, zero value otherwise.
func (o *PKICertificateIssueDetails) GetClientFlag() bool {
	if o == nil || o.ClientFlag == nil {
		var ret bool
		return ret
	}
	return *o.ClientFlag
}

// GetClientFlagOk returns a tuple with the ClientFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PKICertificateIssueDetails) GetClientFlagOk() (*bool, bool) {
	if o == nil || o.ClientFlag == nil {
		return nil, false
	}
	return o.ClientFlag, true
}

// HasClientFlag returns a boolean if a field has been set.
func (o *PKICertificateIssueDetails) HasClientFlag() bool {
	if o != nil && o.ClientFlag != nil {
		return true
	}

	return false
}

// SetClientFlag gets a reference to the given bool and assigns it to the ClientFlag field.
func (o *PKICertificateIssueDetails) SetClientFlag(v bool) {
	o.ClientFlag = &v
}

// GetCodeSigningFlag returns the CodeSigningFlag field value if set, zero value otherwise.
func (o *PKICertificateIssueDetails) GetCodeSigningFlag() bool {
	if o == nil || o.CodeSigningFlag == nil {
		var ret bool
		return ret
	}
	return *o.CodeSigningFlag
}

// GetCodeSigningFlagOk returns a tuple with the CodeSigningFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PKICertificateIssueDetails) GetCodeSigningFlagOk() (*bool, bool) {
	if o == nil || o.CodeSigningFlag == nil {
		return nil, false
	}
	return o.CodeSigningFlag, true
}

// HasCodeSigningFlag returns a boolean if a field has been set.
func (o *PKICertificateIssueDetails) HasCodeSigningFlag() bool {
	if o != nil && o.CodeSigningFlag != nil {
		return true
	}

	return false
}

// SetCodeSigningFlag gets a reference to the given bool and assigns it to the CodeSigningFlag field.
func (o *PKICertificateIssueDetails) SetCodeSigningFlag(v bool) {
	o.CodeSigningFlag = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *PKICertificateIssueDetails) GetCountry() []string {
	if o == nil || o.Country == nil {
		var ret []string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PKICertificateIssueDetails) GetCountryOk() (*[]string, bool) {
	if o == nil || o.Country == nil {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *PKICertificateIssueDetails) HasCountry() bool {
	if o != nil && o.Country != nil {
		return true
	}

	return false
}

// SetCountry gets a reference to the given []string and assigns it to the Country field.
func (o *PKICertificateIssueDetails) SetCountry(v []string) {
	o.Country = &v
}

// GetEnforceHostnames returns the EnforceHostnames field value if set, zero value otherwise.
func (o *PKICertificateIssueDetails) GetEnforceHostnames() bool {
	if o == nil || o.EnforceHostnames == nil {
		var ret bool
		return ret
	}
	return *o.EnforceHostnames
}

// GetEnforceHostnamesOk returns a tuple with the EnforceHostnames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PKICertificateIssueDetails) GetEnforceHostnamesOk() (*bool, bool) {
	if o == nil || o.EnforceHostnames == nil {
		return nil, false
	}
	return o.EnforceHostnames, true
}

// HasEnforceHostnames returns a boolean if a field has been set.
func (o *PKICertificateIssueDetails) HasEnforceHostnames() bool {
	if o != nil && o.EnforceHostnames != nil {
		return true
	}

	return false
}

// SetEnforceHostnames gets a reference to the given bool and assigns it to the EnforceHostnames field.
func (o *PKICertificateIssueDetails) SetEnforceHostnames(v bool) {
	o.EnforceHostnames = &v
}

// GetIsCa returns the IsCa field value if set, zero value otherwise.
func (o *PKICertificateIssueDetails) GetIsCa() bool {
	if o == nil || o.IsCa == nil {
		var ret bool
		return ret
	}
	return *o.IsCa
}

// GetIsCaOk returns a tuple with the IsCa field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PKICertificateIssueDetails) GetIsCaOk() (*bool, bool) {
	if o == nil || o.IsCa == nil {
		return nil, false
	}
	return o.IsCa, true
}

// HasIsCa returns a boolean if a field has been set.
func (o *PKICertificateIssueDetails) HasIsCa() bool {
	if o != nil && o.IsCa != nil {
		return true
	}

	return false
}

// SetIsCa gets a reference to the given bool and assigns it to the IsCa field.
func (o *PKICertificateIssueDetails) SetIsCa(v bool) {
	o.IsCa = &v
}

// GetKeyBits returns the KeyBits field value if set, zero value otherwise.
func (o *PKICertificateIssueDetails) GetKeyBits() int64 {
	if o == nil || o.KeyBits == nil {
		var ret int64
		return ret
	}
	return *o.KeyBits
}

// GetKeyBitsOk returns a tuple with the KeyBits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PKICertificateIssueDetails) GetKeyBitsOk() (*int64, bool) {
	if o == nil || o.KeyBits == nil {
		return nil, false
	}
	return o.KeyBits, true
}

// HasKeyBits returns a boolean if a field has been set.
func (o *PKICertificateIssueDetails) HasKeyBits() bool {
	if o != nil && o.KeyBits != nil {
		return true
	}

	return false
}

// SetKeyBits gets a reference to the given int64 and assigns it to the KeyBits field.
func (o *PKICertificateIssueDetails) SetKeyBits(v int64) {
	o.KeyBits = &v
}

// GetKeyType returns the KeyType field value if set, zero value otherwise.
func (o *PKICertificateIssueDetails) GetKeyType() string {
	if o == nil || o.KeyType == nil {
		var ret string
		return ret
	}
	return *o.KeyType
}

// GetKeyTypeOk returns a tuple with the KeyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PKICertificateIssueDetails) GetKeyTypeOk() (*string, bool) {
	if o == nil || o.KeyType == nil {
		return nil, false
	}
	return o.KeyType, true
}

// HasKeyType returns a boolean if a field has been set.
func (o *PKICertificateIssueDetails) HasKeyType() bool {
	if o != nil && o.KeyType != nil {
		return true
	}

	return false
}

// SetKeyType gets a reference to the given string and assigns it to the KeyType field.
func (o *PKICertificateIssueDetails) SetKeyType(v string) {
	o.KeyType = &v
}

// GetKeyUsageList returns the KeyUsageList field value if set, zero value otherwise.
func (o *PKICertificateIssueDetails) GetKeyUsageList() []string {
	if o == nil || o.KeyUsageList == nil {
		var ret []string
		return ret
	}
	return *o.KeyUsageList
}

// GetKeyUsageListOk returns a tuple with the KeyUsageList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PKICertificateIssueDetails) GetKeyUsageListOk() (*[]string, bool) {
	if o == nil || o.KeyUsageList == nil {
		return nil, false
	}
	return o.KeyUsageList, true
}

// HasKeyUsageList returns a boolean if a field has been set.
func (o *PKICertificateIssueDetails) HasKeyUsageList() bool {
	if o != nil && o.KeyUsageList != nil {
		return true
	}

	return false
}

// SetKeyUsageList gets a reference to the given []string and assigns it to the KeyUsageList field.
func (o *PKICertificateIssueDetails) SetKeyUsageList(v []string) {
	o.KeyUsageList = &v
}

// GetLocality returns the Locality field value if set, zero value otherwise.
func (o *PKICertificateIssueDetails) GetLocality() []string {
	if o == nil || o.Locality == nil {
		var ret []string
		return ret
	}
	return *o.Locality
}

// GetLocalityOk returns a tuple with the Locality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PKICertificateIssueDetails) GetLocalityOk() (*[]string, bool) {
	if o == nil || o.Locality == nil {
		return nil, false
	}
	return o.Locality, true
}

// HasLocality returns a boolean if a field has been set.
func (o *PKICertificateIssueDetails) HasLocality() bool {
	if o != nil && o.Locality != nil {
		return true
	}

	return false
}

// SetLocality gets a reference to the given []string and assigns it to the Locality field.
func (o *PKICertificateIssueDetails) SetLocality(v []string) {
	o.Locality = &v
}

// GetNotBeforeDuration returns the NotBeforeDuration field value if set, zero value otherwise.
func (o *PKICertificateIssueDetails) GetNotBeforeDuration() int64 {
	if o == nil || o.NotBeforeDuration == nil {
		var ret int64
		return ret
	}
	return *o.NotBeforeDuration
}

// GetNotBeforeDurationOk returns a tuple with the NotBeforeDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PKICertificateIssueDetails) GetNotBeforeDurationOk() (*int64, bool) {
	if o == nil || o.NotBeforeDuration == nil {
		return nil, false
	}
	return o.NotBeforeDuration, true
}

// HasNotBeforeDuration returns a boolean if a field has been set.
func (o *PKICertificateIssueDetails) HasNotBeforeDuration() bool {
	if o != nil && o.NotBeforeDuration != nil {
		return true
	}

	return false
}

// SetNotBeforeDuration gets a reference to the given int64 and assigns it to the NotBeforeDuration field.
func (o *PKICertificateIssueDetails) SetNotBeforeDuration(v int64) {
	o.NotBeforeDuration = &v
}

// GetOrganizationList returns the OrganizationList field value if set, zero value otherwise.
func (o *PKICertificateIssueDetails) GetOrganizationList() []string {
	if o == nil || o.OrganizationList == nil {
		var ret []string
		return ret
	}
	return *o.OrganizationList
}

// GetOrganizationListOk returns a tuple with the OrganizationList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PKICertificateIssueDetails) GetOrganizationListOk() (*[]string, bool) {
	if o == nil || o.OrganizationList == nil {
		return nil, false
	}
	return o.OrganizationList, true
}

// HasOrganizationList returns a boolean if a field has been set.
func (o *PKICertificateIssueDetails) HasOrganizationList() bool {
	if o != nil && o.OrganizationList != nil {
		return true
	}

	return false
}

// SetOrganizationList gets a reference to the given []string and assigns it to the OrganizationList field.
func (o *PKICertificateIssueDetails) SetOrganizationList(v []string) {
	o.OrganizationList = &v
}

// GetOrganizationUnitList returns the OrganizationUnitList field value if set, zero value otherwise.
func (o *PKICertificateIssueDetails) GetOrganizationUnitList() []string {
	if o == nil || o.OrganizationUnitList == nil {
		var ret []string
		return ret
	}
	return *o.OrganizationUnitList
}

// GetOrganizationUnitListOk returns a tuple with the OrganizationUnitList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PKICertificateIssueDetails) GetOrganizationUnitListOk() (*[]string, bool) {
	if o == nil || o.OrganizationUnitList == nil {
		return nil, false
	}
	return o.OrganizationUnitList, true
}

// HasOrganizationUnitList returns a boolean if a field has been set.
func (o *PKICertificateIssueDetails) HasOrganizationUnitList() bool {
	if o != nil && o.OrganizationUnitList != nil {
		return true
	}

	return false
}

// SetOrganizationUnitList gets a reference to the given []string and assigns it to the OrganizationUnitList field.
func (o *PKICertificateIssueDetails) SetOrganizationUnitList(v []string) {
	o.OrganizationUnitList = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *PKICertificateIssueDetails) GetPostalCode() []string {
	if o == nil || o.PostalCode == nil {
		var ret []string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PKICertificateIssueDetails) GetPostalCodeOk() (*[]string, bool) {
	if o == nil || o.PostalCode == nil {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *PKICertificateIssueDetails) HasPostalCode() bool {
	if o != nil && o.PostalCode != nil {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given []string and assigns it to the PostalCode field.
func (o *PKICertificateIssueDetails) SetPostalCode(v []string) {
	o.PostalCode = &v
}

// GetProvince returns the Province field value if set, zero value otherwise.
func (o *PKICertificateIssueDetails) GetProvince() []string {
	if o == nil || o.Province == nil {
		var ret []string
		return ret
	}
	return *o.Province
}

// GetProvinceOk returns a tuple with the Province field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PKICertificateIssueDetails) GetProvinceOk() (*[]string, bool) {
	if o == nil || o.Province == nil {
		return nil, false
	}
	return o.Province, true
}

// HasProvince returns a boolean if a field has been set.
func (o *PKICertificateIssueDetails) HasProvince() bool {
	if o != nil && o.Province != nil {
		return true
	}

	return false
}

// SetProvince gets a reference to the given []string and assigns it to the Province field.
func (o *PKICertificateIssueDetails) SetProvince(v []string) {
	o.Province = &v
}

// GetRequireCn returns the RequireCn field value if set, zero value otherwise.
func (o *PKICertificateIssueDetails) GetRequireCn() bool {
	if o == nil || o.RequireCn == nil {
		var ret bool
		return ret
	}
	return *o.RequireCn
}

// GetRequireCnOk returns a tuple with the RequireCn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PKICertificateIssueDetails) GetRequireCnOk() (*bool, bool) {
	if o == nil || o.RequireCn == nil {
		return nil, false
	}
	return o.RequireCn, true
}

// HasRequireCn returns a boolean if a field has been set.
func (o *PKICertificateIssueDetails) HasRequireCn() bool {
	if o != nil && o.RequireCn != nil {
		return true
	}

	return false
}

// SetRequireCn gets a reference to the given bool and assigns it to the RequireCn field.
func (o *PKICertificateIssueDetails) SetRequireCn(v bool) {
	o.RequireCn = &v
}

// GetServerFlag returns the ServerFlag field value if set, zero value otherwise.
func (o *PKICertificateIssueDetails) GetServerFlag() bool {
	if o == nil || o.ServerFlag == nil {
		var ret bool
		return ret
	}
	return *o.ServerFlag
}

// GetServerFlagOk returns a tuple with the ServerFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PKICertificateIssueDetails) GetServerFlagOk() (*bool, bool) {
	if o == nil || o.ServerFlag == nil {
		return nil, false
	}
	return o.ServerFlag, true
}

// HasServerFlag returns a boolean if a field has been set.
func (o *PKICertificateIssueDetails) HasServerFlag() bool {
	if o != nil && o.ServerFlag != nil {
		return true
	}

	return false
}

// SetServerFlag gets a reference to the given bool and assigns it to the ServerFlag field.
func (o *PKICertificateIssueDetails) SetServerFlag(v bool) {
	o.ServerFlag = &v
}

// GetStreetAddress returns the StreetAddress field value if set, zero value otherwise.
func (o *PKICertificateIssueDetails) GetStreetAddress() []string {
	if o == nil || o.StreetAddress == nil {
		var ret []string
		return ret
	}
	return *o.StreetAddress
}

// GetStreetAddressOk returns a tuple with the StreetAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PKICertificateIssueDetails) GetStreetAddressOk() (*[]string, bool) {
	if o == nil || o.StreetAddress == nil {
		return nil, false
	}
	return o.StreetAddress, true
}

// HasStreetAddress returns a boolean if a field has been set.
func (o *PKICertificateIssueDetails) HasStreetAddress() bool {
	if o != nil && o.StreetAddress != nil {
		return true
	}

	return false
}

// SetStreetAddress gets a reference to the given []string and assigns it to the StreetAddress field.
func (o *PKICertificateIssueDetails) SetStreetAddress(v []string) {
	o.StreetAddress = &v
}

func (o PKICertificateIssueDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AllowAnyName != nil {
		toSerialize["allow_any_name"] = o.AllowAnyName
	}
	if o.AllowSubdomains != nil {
		toSerialize["allow_subdomains"] = o.AllowSubdomains
	}
	if o.AllowedDomainsList != nil {
		toSerialize["allowed_domains_list"] = o.AllowedDomainsList
	}
	if o.AllowedUriSans != nil {
		toSerialize["allowed_uri_sans"] = o.AllowedUriSans
	}
	if o.BasicConstraintsValidForNonCa != nil {
		toSerialize["basic_constraints_valid_for_non_ca"] = o.BasicConstraintsValidForNonCa
	}
	if o.ClientFlag != nil {
		toSerialize["client_flag"] = o.ClientFlag
	}
	if o.CodeSigningFlag != nil {
		toSerialize["code_signing_flag"] = o.CodeSigningFlag
	}
	if o.Country != nil {
		toSerialize["country"] = o.Country
	}
	if o.EnforceHostnames != nil {
		toSerialize["enforce_hostnames"] = o.EnforceHostnames
	}
	if o.IsCa != nil {
		toSerialize["is_ca"] = o.IsCa
	}
	if o.KeyBits != nil {
		toSerialize["key_bits"] = o.KeyBits
	}
	if o.KeyType != nil {
		toSerialize["key_type"] = o.KeyType
	}
	if o.KeyUsageList != nil {
		toSerialize["key_usage_list"] = o.KeyUsageList
	}
	if o.Locality != nil {
		toSerialize["locality"] = o.Locality
	}
	if o.NotBeforeDuration != nil {
		toSerialize["not_before_duration"] = o.NotBeforeDuration
	}
	if o.OrganizationList != nil {
		toSerialize["organization_list"] = o.OrganizationList
	}
	if o.OrganizationUnitList != nil {
		toSerialize["organization_unit_list"] = o.OrganizationUnitList
	}
	if o.PostalCode != nil {
		toSerialize["postal_code"] = o.PostalCode
	}
	if o.Province != nil {
		toSerialize["province"] = o.Province
	}
	if o.RequireCn != nil {
		toSerialize["require_cn"] = o.RequireCn
	}
	if o.ServerFlag != nil {
		toSerialize["server_flag"] = o.ServerFlag
	}
	if o.StreetAddress != nil {
		toSerialize["street_address"] = o.StreetAddress
	}
	return json.Marshal(toSerialize)
}

type NullablePKICertificateIssueDetails struct {
	value *PKICertificateIssueDetails
	isSet bool
}

func (v NullablePKICertificateIssueDetails) Get() *PKICertificateIssueDetails {
	return v.value
}

func (v *NullablePKICertificateIssueDetails) Set(val *PKICertificateIssueDetails) {
	v.value = val
	v.isSet = true
}

func (v NullablePKICertificateIssueDetails) IsSet() bool {
	return v.isSet
}

func (v *NullablePKICertificateIssueDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePKICertificateIssueDetails(val *PKICertificateIssueDetails) *NullablePKICertificateIssueDetails {
	return &NullablePKICertificateIssueDetails{value: val, isSet: true}
}

func (v NullablePKICertificateIssueDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePKICertificateIssueDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


