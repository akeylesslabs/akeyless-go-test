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

// CreatePKICertIssuer struct for CreatePKICertIssuer
type CreatePKICertIssuer struct {
	// If set, clients can request certificates for any CN
	AllowAnyName *bool `json:"allow-any-name,omitempty"`
	// If set, clients can request certificates for subdomains and wildcard subdomains of the allowed domains
	AllowSubdomains *bool `json:"allow-subdomains,omitempty"`
	// A list of the allowed domains that clients can request to be included in the certificate (in a comma-delimited list)
	AllowedDomains *string `json:"allowed-domains,omitempty"`
	// A list of the allowed URIs that clients can request to be included in the certificate as part of the URI Subject Alternative Names (in a comma-delimited list)
	AllowedUriSans *string `json:"allowed-uri-sans,omitempty"`
	// If set, certificates will be flagged for client auth use
	ClientFlag *bool `json:"client-flag,omitempty"`
	// If set, certificates will be flagged for code signing use
	CodeSigningFlag *bool `json:"code-signing-flag,omitempty"`
	// A comma-separated list of countries that will be set in the issued certificate
	Country *string `json:"country,omitempty"`
	// Protection from accidental deletion of this item
	DeleteProtection *string `json:"delete_protection,omitempty"`
	// Description of the object
	Description *string `json:"description,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// key-usage
	KeyUsage *string `json:"key-usage,omitempty"`
	// A comma-separated list of localities that will be set in the issued certificate
	Locality *string `json:"locality,omitempty"`
	// Deprecated - use description
	Metadata *string `json:"metadata,omitempty"`
	// PKI certificate issuer name
	Name string `json:"name"`
	// If set, any names are allowed for CN and SANs in the certificate and not only a valid host name
	NotEnforceHostnames *bool `json:"not-enforce-hostnames,omitempty"`
	// If set, clients can request certificates without a CN
	NotRequireCn *bool `json:"not-require-cn,omitempty"`
	// A comma-separated list of organizational units (OU) that will be set in the issued certificate
	OrganizationalUnits *string `json:"organizational-units,omitempty"`
	// A comma-separated list of organizations (O) that will be set in the issued certificate
	Organizations *string `json:"organizations,omitempty"`
	// A comma-separated list of postal codes that will be set in the issued certificate
	PostalCode *string `json:"postal-code,omitempty"`
	// A comma-separated list of provinces that will be set in the issued certificate
	Province *string `json:"province,omitempty"`
	// If set, certificates will be flagged for server auth use
	ServerFlag *bool `json:"server-flag,omitempty"`
	// A key to sign the certificate with
	SignerKeyName string `json:"signer-key-name"`
	// A comma-separated list of street addresses that will be set in the issued certificate
	StreetAddress *string `json:"street-address,omitempty"`
	// List of the tags attached to this key
	Tag *[]string `json:"tag,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// he requested Time To Live for the certificate, in seconds
	Ttl int64 `json:"ttl"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewCreatePKICertIssuer instantiates a new CreatePKICertIssuer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreatePKICertIssuer(name string, signerKeyName string, ttl int64, ) *CreatePKICertIssuer {
	this := CreatePKICertIssuer{}
	var keyUsage string = "DigitalSignature,KeyAgreement,KeyEncipherment"
	this.KeyUsage = &keyUsage
	this.Name = name
	this.SignerKeyName = signerKeyName
	this.Ttl = ttl
	return &this
}

// NewCreatePKICertIssuerWithDefaults instantiates a new CreatePKICertIssuer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreatePKICertIssuerWithDefaults() *CreatePKICertIssuer {
	this := CreatePKICertIssuer{}
	var keyUsage string = "DigitalSignature,KeyAgreement,KeyEncipherment"
	this.KeyUsage = &keyUsage
	return &this
}

// GetAllowAnyName returns the AllowAnyName field value if set, zero value otherwise.
func (o *CreatePKICertIssuer) GetAllowAnyName() bool {
	if o == nil || o.AllowAnyName == nil {
		var ret bool
		return ret
	}
	return *o.AllowAnyName
}

// GetAllowAnyNameOk returns a tuple with the AllowAnyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePKICertIssuer) GetAllowAnyNameOk() (*bool, bool) {
	if o == nil || o.AllowAnyName == nil {
		return nil, false
	}
	return o.AllowAnyName, true
}

// HasAllowAnyName returns a boolean if a field has been set.
func (o *CreatePKICertIssuer) HasAllowAnyName() bool {
	if o != nil && o.AllowAnyName != nil {
		return true
	}

	return false
}

// SetAllowAnyName gets a reference to the given bool and assigns it to the AllowAnyName field.
func (o *CreatePKICertIssuer) SetAllowAnyName(v bool) {
	o.AllowAnyName = &v
}

// GetAllowSubdomains returns the AllowSubdomains field value if set, zero value otherwise.
func (o *CreatePKICertIssuer) GetAllowSubdomains() bool {
	if o == nil || o.AllowSubdomains == nil {
		var ret bool
		return ret
	}
	return *o.AllowSubdomains
}

// GetAllowSubdomainsOk returns a tuple with the AllowSubdomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePKICertIssuer) GetAllowSubdomainsOk() (*bool, bool) {
	if o == nil || o.AllowSubdomains == nil {
		return nil, false
	}
	return o.AllowSubdomains, true
}

// HasAllowSubdomains returns a boolean if a field has been set.
func (o *CreatePKICertIssuer) HasAllowSubdomains() bool {
	if o != nil && o.AllowSubdomains != nil {
		return true
	}

	return false
}

// SetAllowSubdomains gets a reference to the given bool and assigns it to the AllowSubdomains field.
func (o *CreatePKICertIssuer) SetAllowSubdomains(v bool) {
	o.AllowSubdomains = &v
}

// GetAllowedDomains returns the AllowedDomains field value if set, zero value otherwise.
func (o *CreatePKICertIssuer) GetAllowedDomains() string {
	if o == nil || o.AllowedDomains == nil {
		var ret string
		return ret
	}
	return *o.AllowedDomains
}

// GetAllowedDomainsOk returns a tuple with the AllowedDomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePKICertIssuer) GetAllowedDomainsOk() (*string, bool) {
	if o == nil || o.AllowedDomains == nil {
		return nil, false
	}
	return o.AllowedDomains, true
}

// HasAllowedDomains returns a boolean if a field has been set.
func (o *CreatePKICertIssuer) HasAllowedDomains() bool {
	if o != nil && o.AllowedDomains != nil {
		return true
	}

	return false
}

// SetAllowedDomains gets a reference to the given string and assigns it to the AllowedDomains field.
func (o *CreatePKICertIssuer) SetAllowedDomains(v string) {
	o.AllowedDomains = &v
}

// GetAllowedUriSans returns the AllowedUriSans field value if set, zero value otherwise.
func (o *CreatePKICertIssuer) GetAllowedUriSans() string {
	if o == nil || o.AllowedUriSans == nil {
		var ret string
		return ret
	}
	return *o.AllowedUriSans
}

// GetAllowedUriSansOk returns a tuple with the AllowedUriSans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePKICertIssuer) GetAllowedUriSansOk() (*string, bool) {
	if o == nil || o.AllowedUriSans == nil {
		return nil, false
	}
	return o.AllowedUriSans, true
}

// HasAllowedUriSans returns a boolean if a field has been set.
func (o *CreatePKICertIssuer) HasAllowedUriSans() bool {
	if o != nil && o.AllowedUriSans != nil {
		return true
	}

	return false
}

// SetAllowedUriSans gets a reference to the given string and assigns it to the AllowedUriSans field.
func (o *CreatePKICertIssuer) SetAllowedUriSans(v string) {
	o.AllowedUriSans = &v
}

// GetClientFlag returns the ClientFlag field value if set, zero value otherwise.
func (o *CreatePKICertIssuer) GetClientFlag() bool {
	if o == nil || o.ClientFlag == nil {
		var ret bool
		return ret
	}
	return *o.ClientFlag
}

// GetClientFlagOk returns a tuple with the ClientFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePKICertIssuer) GetClientFlagOk() (*bool, bool) {
	if o == nil || o.ClientFlag == nil {
		return nil, false
	}
	return o.ClientFlag, true
}

// HasClientFlag returns a boolean if a field has been set.
func (o *CreatePKICertIssuer) HasClientFlag() bool {
	if o != nil && o.ClientFlag != nil {
		return true
	}

	return false
}

// SetClientFlag gets a reference to the given bool and assigns it to the ClientFlag field.
func (o *CreatePKICertIssuer) SetClientFlag(v bool) {
	o.ClientFlag = &v
}

// GetCodeSigningFlag returns the CodeSigningFlag field value if set, zero value otherwise.
func (o *CreatePKICertIssuer) GetCodeSigningFlag() bool {
	if o == nil || o.CodeSigningFlag == nil {
		var ret bool
		return ret
	}
	return *o.CodeSigningFlag
}

// GetCodeSigningFlagOk returns a tuple with the CodeSigningFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePKICertIssuer) GetCodeSigningFlagOk() (*bool, bool) {
	if o == nil || o.CodeSigningFlag == nil {
		return nil, false
	}
	return o.CodeSigningFlag, true
}

// HasCodeSigningFlag returns a boolean if a field has been set.
func (o *CreatePKICertIssuer) HasCodeSigningFlag() bool {
	if o != nil && o.CodeSigningFlag != nil {
		return true
	}

	return false
}

// SetCodeSigningFlag gets a reference to the given bool and assigns it to the CodeSigningFlag field.
func (o *CreatePKICertIssuer) SetCodeSigningFlag(v bool) {
	o.CodeSigningFlag = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *CreatePKICertIssuer) GetCountry() string {
	if o == nil || o.Country == nil {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePKICertIssuer) GetCountryOk() (*string, bool) {
	if o == nil || o.Country == nil {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *CreatePKICertIssuer) HasCountry() bool {
	if o != nil && o.Country != nil {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *CreatePKICertIssuer) SetCountry(v string) {
	o.Country = &v
}

// GetDeleteProtection returns the DeleteProtection field value if set, zero value otherwise.
func (o *CreatePKICertIssuer) GetDeleteProtection() string {
	if o == nil || o.DeleteProtection == nil {
		var ret string
		return ret
	}
	return *o.DeleteProtection
}

// GetDeleteProtectionOk returns a tuple with the DeleteProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePKICertIssuer) GetDeleteProtectionOk() (*string, bool) {
	if o == nil || o.DeleteProtection == nil {
		return nil, false
	}
	return o.DeleteProtection, true
}

// HasDeleteProtection returns a boolean if a field has been set.
func (o *CreatePKICertIssuer) HasDeleteProtection() bool {
	if o != nil && o.DeleteProtection != nil {
		return true
	}

	return false
}

// SetDeleteProtection gets a reference to the given string and assigns it to the DeleteProtection field.
func (o *CreatePKICertIssuer) SetDeleteProtection(v string) {
	o.DeleteProtection = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreatePKICertIssuer) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePKICertIssuer) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreatePKICertIssuer) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreatePKICertIssuer) SetDescription(v string) {
	o.Description = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *CreatePKICertIssuer) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePKICertIssuer) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *CreatePKICertIssuer) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *CreatePKICertIssuer) SetJson(v bool) {
	o.Json = &v
}

// GetKeyUsage returns the KeyUsage field value if set, zero value otherwise.
func (o *CreatePKICertIssuer) GetKeyUsage() string {
	if o == nil || o.KeyUsage == nil {
		var ret string
		return ret
	}
	return *o.KeyUsage
}

// GetKeyUsageOk returns a tuple with the KeyUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePKICertIssuer) GetKeyUsageOk() (*string, bool) {
	if o == nil || o.KeyUsage == nil {
		return nil, false
	}
	return o.KeyUsage, true
}

// HasKeyUsage returns a boolean if a field has been set.
func (o *CreatePKICertIssuer) HasKeyUsage() bool {
	if o != nil && o.KeyUsage != nil {
		return true
	}

	return false
}

// SetKeyUsage gets a reference to the given string and assigns it to the KeyUsage field.
func (o *CreatePKICertIssuer) SetKeyUsage(v string) {
	o.KeyUsage = &v
}

// GetLocality returns the Locality field value if set, zero value otherwise.
func (o *CreatePKICertIssuer) GetLocality() string {
	if o == nil || o.Locality == nil {
		var ret string
		return ret
	}
	return *o.Locality
}

// GetLocalityOk returns a tuple with the Locality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePKICertIssuer) GetLocalityOk() (*string, bool) {
	if o == nil || o.Locality == nil {
		return nil, false
	}
	return o.Locality, true
}

// HasLocality returns a boolean if a field has been set.
func (o *CreatePKICertIssuer) HasLocality() bool {
	if o != nil && o.Locality != nil {
		return true
	}

	return false
}

// SetLocality gets a reference to the given string and assigns it to the Locality field.
func (o *CreatePKICertIssuer) SetLocality(v string) {
	o.Locality = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *CreatePKICertIssuer) GetMetadata() string {
	if o == nil || o.Metadata == nil {
		var ret string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePKICertIssuer) GetMetadataOk() (*string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CreatePKICertIssuer) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given string and assigns it to the Metadata field.
func (o *CreatePKICertIssuer) SetMetadata(v string) {
	o.Metadata = &v
}

// GetName returns the Name field value
func (o *CreatePKICertIssuer) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreatePKICertIssuer) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreatePKICertIssuer) SetName(v string) {
	o.Name = v
}

// GetNotEnforceHostnames returns the NotEnforceHostnames field value if set, zero value otherwise.
func (o *CreatePKICertIssuer) GetNotEnforceHostnames() bool {
	if o == nil || o.NotEnforceHostnames == nil {
		var ret bool
		return ret
	}
	return *o.NotEnforceHostnames
}

// GetNotEnforceHostnamesOk returns a tuple with the NotEnforceHostnames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePKICertIssuer) GetNotEnforceHostnamesOk() (*bool, bool) {
	if o == nil || o.NotEnforceHostnames == nil {
		return nil, false
	}
	return o.NotEnforceHostnames, true
}

// HasNotEnforceHostnames returns a boolean if a field has been set.
func (o *CreatePKICertIssuer) HasNotEnforceHostnames() bool {
	if o != nil && o.NotEnforceHostnames != nil {
		return true
	}

	return false
}

// SetNotEnforceHostnames gets a reference to the given bool and assigns it to the NotEnforceHostnames field.
func (o *CreatePKICertIssuer) SetNotEnforceHostnames(v bool) {
	o.NotEnforceHostnames = &v
}

// GetNotRequireCn returns the NotRequireCn field value if set, zero value otherwise.
func (o *CreatePKICertIssuer) GetNotRequireCn() bool {
	if o == nil || o.NotRequireCn == nil {
		var ret bool
		return ret
	}
	return *o.NotRequireCn
}

// GetNotRequireCnOk returns a tuple with the NotRequireCn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePKICertIssuer) GetNotRequireCnOk() (*bool, bool) {
	if o == nil || o.NotRequireCn == nil {
		return nil, false
	}
	return o.NotRequireCn, true
}

// HasNotRequireCn returns a boolean if a field has been set.
func (o *CreatePKICertIssuer) HasNotRequireCn() bool {
	if o != nil && o.NotRequireCn != nil {
		return true
	}

	return false
}

// SetNotRequireCn gets a reference to the given bool and assigns it to the NotRequireCn field.
func (o *CreatePKICertIssuer) SetNotRequireCn(v bool) {
	o.NotRequireCn = &v
}

// GetOrganizationalUnits returns the OrganizationalUnits field value if set, zero value otherwise.
func (o *CreatePKICertIssuer) GetOrganizationalUnits() string {
	if o == nil || o.OrganizationalUnits == nil {
		var ret string
		return ret
	}
	return *o.OrganizationalUnits
}

// GetOrganizationalUnitsOk returns a tuple with the OrganizationalUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePKICertIssuer) GetOrganizationalUnitsOk() (*string, bool) {
	if o == nil || o.OrganizationalUnits == nil {
		return nil, false
	}
	return o.OrganizationalUnits, true
}

// HasOrganizationalUnits returns a boolean if a field has been set.
func (o *CreatePKICertIssuer) HasOrganizationalUnits() bool {
	if o != nil && o.OrganizationalUnits != nil {
		return true
	}

	return false
}

// SetOrganizationalUnits gets a reference to the given string and assigns it to the OrganizationalUnits field.
func (o *CreatePKICertIssuer) SetOrganizationalUnits(v string) {
	o.OrganizationalUnits = &v
}

// GetOrganizations returns the Organizations field value if set, zero value otherwise.
func (o *CreatePKICertIssuer) GetOrganizations() string {
	if o == nil || o.Organizations == nil {
		var ret string
		return ret
	}
	return *o.Organizations
}

// GetOrganizationsOk returns a tuple with the Organizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePKICertIssuer) GetOrganizationsOk() (*string, bool) {
	if o == nil || o.Organizations == nil {
		return nil, false
	}
	return o.Organizations, true
}

// HasOrganizations returns a boolean if a field has been set.
func (o *CreatePKICertIssuer) HasOrganizations() bool {
	if o != nil && o.Organizations != nil {
		return true
	}

	return false
}

// SetOrganizations gets a reference to the given string and assigns it to the Organizations field.
func (o *CreatePKICertIssuer) SetOrganizations(v string) {
	o.Organizations = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *CreatePKICertIssuer) GetPostalCode() string {
	if o == nil || o.PostalCode == nil {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePKICertIssuer) GetPostalCodeOk() (*string, bool) {
	if o == nil || o.PostalCode == nil {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *CreatePKICertIssuer) HasPostalCode() bool {
	if o != nil && o.PostalCode != nil {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *CreatePKICertIssuer) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetProvince returns the Province field value if set, zero value otherwise.
func (o *CreatePKICertIssuer) GetProvince() string {
	if o == nil || o.Province == nil {
		var ret string
		return ret
	}
	return *o.Province
}

// GetProvinceOk returns a tuple with the Province field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePKICertIssuer) GetProvinceOk() (*string, bool) {
	if o == nil || o.Province == nil {
		return nil, false
	}
	return o.Province, true
}

// HasProvince returns a boolean if a field has been set.
func (o *CreatePKICertIssuer) HasProvince() bool {
	if o != nil && o.Province != nil {
		return true
	}

	return false
}

// SetProvince gets a reference to the given string and assigns it to the Province field.
func (o *CreatePKICertIssuer) SetProvince(v string) {
	o.Province = &v
}

// GetServerFlag returns the ServerFlag field value if set, zero value otherwise.
func (o *CreatePKICertIssuer) GetServerFlag() bool {
	if o == nil || o.ServerFlag == nil {
		var ret bool
		return ret
	}
	return *o.ServerFlag
}

// GetServerFlagOk returns a tuple with the ServerFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePKICertIssuer) GetServerFlagOk() (*bool, bool) {
	if o == nil || o.ServerFlag == nil {
		return nil, false
	}
	return o.ServerFlag, true
}

// HasServerFlag returns a boolean if a field has been set.
func (o *CreatePKICertIssuer) HasServerFlag() bool {
	if o != nil && o.ServerFlag != nil {
		return true
	}

	return false
}

// SetServerFlag gets a reference to the given bool and assigns it to the ServerFlag field.
func (o *CreatePKICertIssuer) SetServerFlag(v bool) {
	o.ServerFlag = &v
}

// GetSignerKeyName returns the SignerKeyName field value
func (o *CreatePKICertIssuer) GetSignerKeyName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.SignerKeyName
}

// GetSignerKeyNameOk returns a tuple with the SignerKeyName field value
// and a boolean to check if the value has been set.
func (o *CreatePKICertIssuer) GetSignerKeyNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SignerKeyName, true
}

// SetSignerKeyName sets field value
func (o *CreatePKICertIssuer) SetSignerKeyName(v string) {
	o.SignerKeyName = v
}

// GetStreetAddress returns the StreetAddress field value if set, zero value otherwise.
func (o *CreatePKICertIssuer) GetStreetAddress() string {
	if o == nil || o.StreetAddress == nil {
		var ret string
		return ret
	}
	return *o.StreetAddress
}

// GetStreetAddressOk returns a tuple with the StreetAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePKICertIssuer) GetStreetAddressOk() (*string, bool) {
	if o == nil || o.StreetAddress == nil {
		return nil, false
	}
	return o.StreetAddress, true
}

// HasStreetAddress returns a boolean if a field has been set.
func (o *CreatePKICertIssuer) HasStreetAddress() bool {
	if o != nil && o.StreetAddress != nil {
		return true
	}

	return false
}

// SetStreetAddress gets a reference to the given string and assigns it to the StreetAddress field.
func (o *CreatePKICertIssuer) SetStreetAddress(v string) {
	o.StreetAddress = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *CreatePKICertIssuer) GetTag() []string {
	if o == nil || o.Tag == nil {
		var ret []string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePKICertIssuer) GetTagOk() (*[]string, bool) {
	if o == nil || o.Tag == nil {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *CreatePKICertIssuer) HasTag() bool {
	if o != nil && o.Tag != nil {
		return true
	}

	return false
}

// SetTag gets a reference to the given []string and assigns it to the Tag field.
func (o *CreatePKICertIssuer) SetTag(v []string) {
	o.Tag = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *CreatePKICertIssuer) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePKICertIssuer) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *CreatePKICertIssuer) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *CreatePKICertIssuer) SetToken(v string) {
	o.Token = &v
}

// GetTtl returns the Ttl field value
func (o *CreatePKICertIssuer) GetTtl() int64 {
	if o == nil  {
		var ret int64
		return ret
	}

	return o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value
// and a boolean to check if the value has been set.
func (o *CreatePKICertIssuer) GetTtlOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Ttl, true
}

// SetTtl sets field value
func (o *CreatePKICertIssuer) SetTtl(v int64) {
	o.Ttl = v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *CreatePKICertIssuer) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePKICertIssuer) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *CreatePKICertIssuer) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *CreatePKICertIssuer) SetUidToken(v string) {
	o.UidToken = &v
}

func (o CreatePKICertIssuer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AllowAnyName != nil {
		toSerialize["allow-any-name"] = o.AllowAnyName
	}
	if o.AllowSubdomains != nil {
		toSerialize["allow-subdomains"] = o.AllowSubdomains
	}
	if o.AllowedDomains != nil {
		toSerialize["allowed-domains"] = o.AllowedDomains
	}
	if o.AllowedUriSans != nil {
		toSerialize["allowed-uri-sans"] = o.AllowedUriSans
	}
	if o.ClientFlag != nil {
		toSerialize["client-flag"] = o.ClientFlag
	}
	if o.CodeSigningFlag != nil {
		toSerialize["code-signing-flag"] = o.CodeSigningFlag
	}
	if o.Country != nil {
		toSerialize["country"] = o.Country
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
	if o.KeyUsage != nil {
		toSerialize["key-usage"] = o.KeyUsage
	}
	if o.Locality != nil {
		toSerialize["locality"] = o.Locality
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.NotEnforceHostnames != nil {
		toSerialize["not-enforce-hostnames"] = o.NotEnforceHostnames
	}
	if o.NotRequireCn != nil {
		toSerialize["not-require-cn"] = o.NotRequireCn
	}
	if o.OrganizationalUnits != nil {
		toSerialize["organizational-units"] = o.OrganizationalUnits
	}
	if o.Organizations != nil {
		toSerialize["organizations"] = o.Organizations
	}
	if o.PostalCode != nil {
		toSerialize["postal-code"] = o.PostalCode
	}
	if o.Province != nil {
		toSerialize["province"] = o.Province
	}
	if o.ServerFlag != nil {
		toSerialize["server-flag"] = o.ServerFlag
	}
	if true {
		toSerialize["signer-key-name"] = o.SignerKeyName
	}
	if o.StreetAddress != nil {
		toSerialize["street-address"] = o.StreetAddress
	}
	if o.Tag != nil {
		toSerialize["tag"] = o.Tag
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if true {
		toSerialize["ttl"] = o.Ttl
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableCreatePKICertIssuer struct {
	value *CreatePKICertIssuer
	isSet bool
}

func (v NullableCreatePKICertIssuer) Get() *CreatePKICertIssuer {
	return v.value
}

func (v *NullableCreatePKICertIssuer) Set(val *CreatePKICertIssuer) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatePKICertIssuer) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatePKICertIssuer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatePKICertIssuer(val *CreatePKICertIssuer) *NullableCreatePKICertIssuer {
	return &NullableCreatePKICertIssuer{value: val, isSet: true}
}

func (v NullableCreatePKICertIssuer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreatePKICertIssuer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


