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

// UpdatePKICertIssuer struct for UpdatePKICertIssuer
type UpdatePKICertIssuer struct {
	// List of the new tags that will be attached to this item
	AddTag *[]string `json:"add-tag,omitempty"`
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
	// A comma-separated list of the country that will be set in the issued certificate
	Country *string `json:"country,omitempty"`
	// Description of the object
	Description *string `json:"description,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// key-usage
	KeyUsage *string `json:"key-usage,omitempty"`
	// A comma-separated list of the locality that will be set in the issued certificate
	Locality *string `json:"locality,omitempty"`
	// Deprecated - use description
	Metadata *string `json:"metadata,omitempty"`
	// PKI certificate issuer name
	Name string `json:"name"`
	// New item name
	NewName *string `json:"new-name,omitempty"`
	// If set, any names are allowed for CN and SANs in the certificate and not only a valid host name
	NotEnforceHostnames *bool `json:"not-enforce-hostnames,omitempty"`
	// If set, clients can request certificates without a CN
	NotRequireCn *bool `json:"not-require-cn,omitempty"`
	// A comma-separated list of organizational units (OU) that will be set in the issued certificate
	OrganizationalUnits *string `json:"organizational-units,omitempty"`
	// A comma-separated list of organizations (O) that will be set in the issued certificate
	Organizations *string `json:"organizations,omitempty"`
	// A comma-separated list of the postal code that will be set in the issued certificate
	PostalCode *string `json:"postal-code,omitempty"`
	// A comma-separated list of the province that will be set in the issued certificate
	Province *string `json:"province,omitempty"`
	// List of the existent tags that will be removed from this item
	RmTag *[]string `json:"rm-tag,omitempty"`
	// If set, certificates will be flagged for server auth use
	ServerFlag *bool `json:"server-flag,omitempty"`
	// A key to sign the certificate with
	SignerKeyName string `json:"signer-key-name"`
	// A comma-separated list of the street address that will be set in the issued certificate
	StreetAddress *string `json:"street-address,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// he requested Time To Live for the certificate, in seconds
	Ttl int64 `json:"ttl"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewUpdatePKICertIssuer instantiates a new UpdatePKICertIssuer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdatePKICertIssuer(name string, signerKeyName string, ttl int64, ) *UpdatePKICertIssuer {
	this := UpdatePKICertIssuer{}
	var keyUsage string = "DigitalSignature,KeyAgreement,KeyEncipherment"
	this.KeyUsage = &keyUsage
	this.Name = name
	this.SignerKeyName = signerKeyName
	this.Ttl = ttl
	return &this
}

// NewUpdatePKICertIssuerWithDefaults instantiates a new UpdatePKICertIssuer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdatePKICertIssuerWithDefaults() *UpdatePKICertIssuer {
	this := UpdatePKICertIssuer{}
	var keyUsage string = "DigitalSignature,KeyAgreement,KeyEncipherment"
	this.KeyUsage = &keyUsage
	return &this
}

// GetAddTag returns the AddTag field value if set, zero value otherwise.
func (o *UpdatePKICertIssuer) GetAddTag() []string {
	if o == nil || o.AddTag == nil {
		var ret []string
		return ret
	}
	return *o.AddTag
}

// GetAddTagOk returns a tuple with the AddTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePKICertIssuer) GetAddTagOk() (*[]string, bool) {
	if o == nil || o.AddTag == nil {
		return nil, false
	}
	return o.AddTag, true
}

// HasAddTag returns a boolean if a field has been set.
func (o *UpdatePKICertIssuer) HasAddTag() bool {
	if o != nil && o.AddTag != nil {
		return true
	}

	return false
}

// SetAddTag gets a reference to the given []string and assigns it to the AddTag field.
func (o *UpdatePKICertIssuer) SetAddTag(v []string) {
	o.AddTag = &v
}

// GetAllowAnyName returns the AllowAnyName field value if set, zero value otherwise.
func (o *UpdatePKICertIssuer) GetAllowAnyName() bool {
	if o == nil || o.AllowAnyName == nil {
		var ret bool
		return ret
	}
	return *o.AllowAnyName
}

// GetAllowAnyNameOk returns a tuple with the AllowAnyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePKICertIssuer) GetAllowAnyNameOk() (*bool, bool) {
	if o == nil || o.AllowAnyName == nil {
		return nil, false
	}
	return o.AllowAnyName, true
}

// HasAllowAnyName returns a boolean if a field has been set.
func (o *UpdatePKICertIssuer) HasAllowAnyName() bool {
	if o != nil && o.AllowAnyName != nil {
		return true
	}

	return false
}

// SetAllowAnyName gets a reference to the given bool and assigns it to the AllowAnyName field.
func (o *UpdatePKICertIssuer) SetAllowAnyName(v bool) {
	o.AllowAnyName = &v
}

// GetAllowSubdomains returns the AllowSubdomains field value if set, zero value otherwise.
func (o *UpdatePKICertIssuer) GetAllowSubdomains() bool {
	if o == nil || o.AllowSubdomains == nil {
		var ret bool
		return ret
	}
	return *o.AllowSubdomains
}

// GetAllowSubdomainsOk returns a tuple with the AllowSubdomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePKICertIssuer) GetAllowSubdomainsOk() (*bool, bool) {
	if o == nil || o.AllowSubdomains == nil {
		return nil, false
	}
	return o.AllowSubdomains, true
}

// HasAllowSubdomains returns a boolean if a field has been set.
func (o *UpdatePKICertIssuer) HasAllowSubdomains() bool {
	if o != nil && o.AllowSubdomains != nil {
		return true
	}

	return false
}

// SetAllowSubdomains gets a reference to the given bool and assigns it to the AllowSubdomains field.
func (o *UpdatePKICertIssuer) SetAllowSubdomains(v bool) {
	o.AllowSubdomains = &v
}

// GetAllowedDomains returns the AllowedDomains field value if set, zero value otherwise.
func (o *UpdatePKICertIssuer) GetAllowedDomains() string {
	if o == nil || o.AllowedDomains == nil {
		var ret string
		return ret
	}
	return *o.AllowedDomains
}

// GetAllowedDomainsOk returns a tuple with the AllowedDomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePKICertIssuer) GetAllowedDomainsOk() (*string, bool) {
	if o == nil || o.AllowedDomains == nil {
		return nil, false
	}
	return o.AllowedDomains, true
}

// HasAllowedDomains returns a boolean if a field has been set.
func (o *UpdatePKICertIssuer) HasAllowedDomains() bool {
	if o != nil && o.AllowedDomains != nil {
		return true
	}

	return false
}

// SetAllowedDomains gets a reference to the given string and assigns it to the AllowedDomains field.
func (o *UpdatePKICertIssuer) SetAllowedDomains(v string) {
	o.AllowedDomains = &v
}

// GetAllowedUriSans returns the AllowedUriSans field value if set, zero value otherwise.
func (o *UpdatePKICertIssuer) GetAllowedUriSans() string {
	if o == nil || o.AllowedUriSans == nil {
		var ret string
		return ret
	}
	return *o.AllowedUriSans
}

// GetAllowedUriSansOk returns a tuple with the AllowedUriSans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePKICertIssuer) GetAllowedUriSansOk() (*string, bool) {
	if o == nil || o.AllowedUriSans == nil {
		return nil, false
	}
	return o.AllowedUriSans, true
}

// HasAllowedUriSans returns a boolean if a field has been set.
func (o *UpdatePKICertIssuer) HasAllowedUriSans() bool {
	if o != nil && o.AllowedUriSans != nil {
		return true
	}

	return false
}

// SetAllowedUriSans gets a reference to the given string and assigns it to the AllowedUriSans field.
func (o *UpdatePKICertIssuer) SetAllowedUriSans(v string) {
	o.AllowedUriSans = &v
}

// GetClientFlag returns the ClientFlag field value if set, zero value otherwise.
func (o *UpdatePKICertIssuer) GetClientFlag() bool {
	if o == nil || o.ClientFlag == nil {
		var ret bool
		return ret
	}
	return *o.ClientFlag
}

// GetClientFlagOk returns a tuple with the ClientFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePKICertIssuer) GetClientFlagOk() (*bool, bool) {
	if o == nil || o.ClientFlag == nil {
		return nil, false
	}
	return o.ClientFlag, true
}

// HasClientFlag returns a boolean if a field has been set.
func (o *UpdatePKICertIssuer) HasClientFlag() bool {
	if o != nil && o.ClientFlag != nil {
		return true
	}

	return false
}

// SetClientFlag gets a reference to the given bool and assigns it to the ClientFlag field.
func (o *UpdatePKICertIssuer) SetClientFlag(v bool) {
	o.ClientFlag = &v
}

// GetCodeSigningFlag returns the CodeSigningFlag field value if set, zero value otherwise.
func (o *UpdatePKICertIssuer) GetCodeSigningFlag() bool {
	if o == nil || o.CodeSigningFlag == nil {
		var ret bool
		return ret
	}
	return *o.CodeSigningFlag
}

// GetCodeSigningFlagOk returns a tuple with the CodeSigningFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePKICertIssuer) GetCodeSigningFlagOk() (*bool, bool) {
	if o == nil || o.CodeSigningFlag == nil {
		return nil, false
	}
	return o.CodeSigningFlag, true
}

// HasCodeSigningFlag returns a boolean if a field has been set.
func (o *UpdatePKICertIssuer) HasCodeSigningFlag() bool {
	if o != nil && o.CodeSigningFlag != nil {
		return true
	}

	return false
}

// SetCodeSigningFlag gets a reference to the given bool and assigns it to the CodeSigningFlag field.
func (o *UpdatePKICertIssuer) SetCodeSigningFlag(v bool) {
	o.CodeSigningFlag = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *UpdatePKICertIssuer) GetCountry() string {
	if o == nil || o.Country == nil {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePKICertIssuer) GetCountryOk() (*string, bool) {
	if o == nil || o.Country == nil {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *UpdatePKICertIssuer) HasCountry() bool {
	if o != nil && o.Country != nil {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *UpdatePKICertIssuer) SetCountry(v string) {
	o.Country = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdatePKICertIssuer) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePKICertIssuer) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdatePKICertIssuer) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdatePKICertIssuer) SetDescription(v string) {
	o.Description = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *UpdatePKICertIssuer) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePKICertIssuer) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *UpdatePKICertIssuer) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *UpdatePKICertIssuer) SetJson(v bool) {
	o.Json = &v
}

// GetKeyUsage returns the KeyUsage field value if set, zero value otherwise.
func (o *UpdatePKICertIssuer) GetKeyUsage() string {
	if o == nil || o.KeyUsage == nil {
		var ret string
		return ret
	}
	return *o.KeyUsage
}

// GetKeyUsageOk returns a tuple with the KeyUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePKICertIssuer) GetKeyUsageOk() (*string, bool) {
	if o == nil || o.KeyUsage == nil {
		return nil, false
	}
	return o.KeyUsage, true
}

// HasKeyUsage returns a boolean if a field has been set.
func (o *UpdatePKICertIssuer) HasKeyUsage() bool {
	if o != nil && o.KeyUsage != nil {
		return true
	}

	return false
}

// SetKeyUsage gets a reference to the given string and assigns it to the KeyUsage field.
func (o *UpdatePKICertIssuer) SetKeyUsage(v string) {
	o.KeyUsage = &v
}

// GetLocality returns the Locality field value if set, zero value otherwise.
func (o *UpdatePKICertIssuer) GetLocality() string {
	if o == nil || o.Locality == nil {
		var ret string
		return ret
	}
	return *o.Locality
}

// GetLocalityOk returns a tuple with the Locality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePKICertIssuer) GetLocalityOk() (*string, bool) {
	if o == nil || o.Locality == nil {
		return nil, false
	}
	return o.Locality, true
}

// HasLocality returns a boolean if a field has been set.
func (o *UpdatePKICertIssuer) HasLocality() bool {
	if o != nil && o.Locality != nil {
		return true
	}

	return false
}

// SetLocality gets a reference to the given string and assigns it to the Locality field.
func (o *UpdatePKICertIssuer) SetLocality(v string) {
	o.Locality = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *UpdatePKICertIssuer) GetMetadata() string {
	if o == nil || o.Metadata == nil {
		var ret string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePKICertIssuer) GetMetadataOk() (*string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *UpdatePKICertIssuer) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given string and assigns it to the Metadata field.
func (o *UpdatePKICertIssuer) SetMetadata(v string) {
	o.Metadata = &v
}

// GetName returns the Name field value
func (o *UpdatePKICertIssuer) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdatePKICertIssuer) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdatePKICertIssuer) SetName(v string) {
	o.Name = v
}

// GetNewName returns the NewName field value if set, zero value otherwise.
func (o *UpdatePKICertIssuer) GetNewName() string {
	if o == nil || o.NewName == nil {
		var ret string
		return ret
	}
	return *o.NewName
}

// GetNewNameOk returns a tuple with the NewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePKICertIssuer) GetNewNameOk() (*string, bool) {
	if o == nil || o.NewName == nil {
		return nil, false
	}
	return o.NewName, true
}

// HasNewName returns a boolean if a field has been set.
func (o *UpdatePKICertIssuer) HasNewName() bool {
	if o != nil && o.NewName != nil {
		return true
	}

	return false
}

// SetNewName gets a reference to the given string and assigns it to the NewName field.
func (o *UpdatePKICertIssuer) SetNewName(v string) {
	o.NewName = &v
}

// GetNotEnforceHostnames returns the NotEnforceHostnames field value if set, zero value otherwise.
func (o *UpdatePKICertIssuer) GetNotEnforceHostnames() bool {
	if o == nil || o.NotEnforceHostnames == nil {
		var ret bool
		return ret
	}
	return *o.NotEnforceHostnames
}

// GetNotEnforceHostnamesOk returns a tuple with the NotEnforceHostnames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePKICertIssuer) GetNotEnforceHostnamesOk() (*bool, bool) {
	if o == nil || o.NotEnforceHostnames == nil {
		return nil, false
	}
	return o.NotEnforceHostnames, true
}

// HasNotEnforceHostnames returns a boolean if a field has been set.
func (o *UpdatePKICertIssuer) HasNotEnforceHostnames() bool {
	if o != nil && o.NotEnforceHostnames != nil {
		return true
	}

	return false
}

// SetNotEnforceHostnames gets a reference to the given bool and assigns it to the NotEnforceHostnames field.
func (o *UpdatePKICertIssuer) SetNotEnforceHostnames(v bool) {
	o.NotEnforceHostnames = &v
}

// GetNotRequireCn returns the NotRequireCn field value if set, zero value otherwise.
func (o *UpdatePKICertIssuer) GetNotRequireCn() bool {
	if o == nil || o.NotRequireCn == nil {
		var ret bool
		return ret
	}
	return *o.NotRequireCn
}

// GetNotRequireCnOk returns a tuple with the NotRequireCn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePKICertIssuer) GetNotRequireCnOk() (*bool, bool) {
	if o == nil || o.NotRequireCn == nil {
		return nil, false
	}
	return o.NotRequireCn, true
}

// HasNotRequireCn returns a boolean if a field has been set.
func (o *UpdatePKICertIssuer) HasNotRequireCn() bool {
	if o != nil && o.NotRequireCn != nil {
		return true
	}

	return false
}

// SetNotRequireCn gets a reference to the given bool and assigns it to the NotRequireCn field.
func (o *UpdatePKICertIssuer) SetNotRequireCn(v bool) {
	o.NotRequireCn = &v
}

// GetOrganizationalUnits returns the OrganizationalUnits field value if set, zero value otherwise.
func (o *UpdatePKICertIssuer) GetOrganizationalUnits() string {
	if o == nil || o.OrganizationalUnits == nil {
		var ret string
		return ret
	}
	return *o.OrganizationalUnits
}

// GetOrganizationalUnitsOk returns a tuple with the OrganizationalUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePKICertIssuer) GetOrganizationalUnitsOk() (*string, bool) {
	if o == nil || o.OrganizationalUnits == nil {
		return nil, false
	}
	return o.OrganizationalUnits, true
}

// HasOrganizationalUnits returns a boolean if a field has been set.
func (o *UpdatePKICertIssuer) HasOrganizationalUnits() bool {
	if o != nil && o.OrganizationalUnits != nil {
		return true
	}

	return false
}

// SetOrganizationalUnits gets a reference to the given string and assigns it to the OrganizationalUnits field.
func (o *UpdatePKICertIssuer) SetOrganizationalUnits(v string) {
	o.OrganizationalUnits = &v
}

// GetOrganizations returns the Organizations field value if set, zero value otherwise.
func (o *UpdatePKICertIssuer) GetOrganizations() string {
	if o == nil || o.Organizations == nil {
		var ret string
		return ret
	}
	return *o.Organizations
}

// GetOrganizationsOk returns a tuple with the Organizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePKICertIssuer) GetOrganizationsOk() (*string, bool) {
	if o == nil || o.Organizations == nil {
		return nil, false
	}
	return o.Organizations, true
}

// HasOrganizations returns a boolean if a field has been set.
func (o *UpdatePKICertIssuer) HasOrganizations() bool {
	if o != nil && o.Organizations != nil {
		return true
	}

	return false
}

// SetOrganizations gets a reference to the given string and assigns it to the Organizations field.
func (o *UpdatePKICertIssuer) SetOrganizations(v string) {
	o.Organizations = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *UpdatePKICertIssuer) GetPostalCode() string {
	if o == nil || o.PostalCode == nil {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePKICertIssuer) GetPostalCodeOk() (*string, bool) {
	if o == nil || o.PostalCode == nil {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *UpdatePKICertIssuer) HasPostalCode() bool {
	if o != nil && o.PostalCode != nil {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *UpdatePKICertIssuer) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetProvince returns the Province field value if set, zero value otherwise.
func (o *UpdatePKICertIssuer) GetProvince() string {
	if o == nil || o.Province == nil {
		var ret string
		return ret
	}
	return *o.Province
}

// GetProvinceOk returns a tuple with the Province field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePKICertIssuer) GetProvinceOk() (*string, bool) {
	if o == nil || o.Province == nil {
		return nil, false
	}
	return o.Province, true
}

// HasProvince returns a boolean if a field has been set.
func (o *UpdatePKICertIssuer) HasProvince() bool {
	if o != nil && o.Province != nil {
		return true
	}

	return false
}

// SetProvince gets a reference to the given string and assigns it to the Province field.
func (o *UpdatePKICertIssuer) SetProvince(v string) {
	o.Province = &v
}

// GetRmTag returns the RmTag field value if set, zero value otherwise.
func (o *UpdatePKICertIssuer) GetRmTag() []string {
	if o == nil || o.RmTag == nil {
		var ret []string
		return ret
	}
	return *o.RmTag
}

// GetRmTagOk returns a tuple with the RmTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePKICertIssuer) GetRmTagOk() (*[]string, bool) {
	if o == nil || o.RmTag == nil {
		return nil, false
	}
	return o.RmTag, true
}

// HasRmTag returns a boolean if a field has been set.
func (o *UpdatePKICertIssuer) HasRmTag() bool {
	if o != nil && o.RmTag != nil {
		return true
	}

	return false
}

// SetRmTag gets a reference to the given []string and assigns it to the RmTag field.
func (o *UpdatePKICertIssuer) SetRmTag(v []string) {
	o.RmTag = &v
}

// GetServerFlag returns the ServerFlag field value if set, zero value otherwise.
func (o *UpdatePKICertIssuer) GetServerFlag() bool {
	if o == nil || o.ServerFlag == nil {
		var ret bool
		return ret
	}
	return *o.ServerFlag
}

// GetServerFlagOk returns a tuple with the ServerFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePKICertIssuer) GetServerFlagOk() (*bool, bool) {
	if o == nil || o.ServerFlag == nil {
		return nil, false
	}
	return o.ServerFlag, true
}

// HasServerFlag returns a boolean if a field has been set.
func (o *UpdatePKICertIssuer) HasServerFlag() bool {
	if o != nil && o.ServerFlag != nil {
		return true
	}

	return false
}

// SetServerFlag gets a reference to the given bool and assigns it to the ServerFlag field.
func (o *UpdatePKICertIssuer) SetServerFlag(v bool) {
	o.ServerFlag = &v
}

// GetSignerKeyName returns the SignerKeyName field value
func (o *UpdatePKICertIssuer) GetSignerKeyName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.SignerKeyName
}

// GetSignerKeyNameOk returns a tuple with the SignerKeyName field value
// and a boolean to check if the value has been set.
func (o *UpdatePKICertIssuer) GetSignerKeyNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SignerKeyName, true
}

// SetSignerKeyName sets field value
func (o *UpdatePKICertIssuer) SetSignerKeyName(v string) {
	o.SignerKeyName = v
}

// GetStreetAddress returns the StreetAddress field value if set, zero value otherwise.
func (o *UpdatePKICertIssuer) GetStreetAddress() string {
	if o == nil || o.StreetAddress == nil {
		var ret string
		return ret
	}
	return *o.StreetAddress
}

// GetStreetAddressOk returns a tuple with the StreetAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePKICertIssuer) GetStreetAddressOk() (*string, bool) {
	if o == nil || o.StreetAddress == nil {
		return nil, false
	}
	return o.StreetAddress, true
}

// HasStreetAddress returns a boolean if a field has been set.
func (o *UpdatePKICertIssuer) HasStreetAddress() bool {
	if o != nil && o.StreetAddress != nil {
		return true
	}

	return false
}

// SetStreetAddress gets a reference to the given string and assigns it to the StreetAddress field.
func (o *UpdatePKICertIssuer) SetStreetAddress(v string) {
	o.StreetAddress = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *UpdatePKICertIssuer) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePKICertIssuer) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *UpdatePKICertIssuer) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *UpdatePKICertIssuer) SetToken(v string) {
	o.Token = &v
}

// GetTtl returns the Ttl field value
func (o *UpdatePKICertIssuer) GetTtl() int64 {
	if o == nil  {
		var ret int64
		return ret
	}

	return o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value
// and a boolean to check if the value has been set.
func (o *UpdatePKICertIssuer) GetTtlOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Ttl, true
}

// SetTtl sets field value
func (o *UpdatePKICertIssuer) SetTtl(v int64) {
	o.Ttl = v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *UpdatePKICertIssuer) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePKICertIssuer) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *UpdatePKICertIssuer) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *UpdatePKICertIssuer) SetUidToken(v string) {
	o.UidToken = &v
}

func (o UpdatePKICertIssuer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AddTag != nil {
		toSerialize["add-tag"] = o.AddTag
	}
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
	if o.NewName != nil {
		toSerialize["new-name"] = o.NewName
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
	if o.RmTag != nil {
		toSerialize["rm-tag"] = o.RmTag
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

type NullableUpdatePKICertIssuer struct {
	value *UpdatePKICertIssuer
	isSet bool
}

func (v NullableUpdatePKICertIssuer) Get() *UpdatePKICertIssuer {
	return v.value
}

func (v *NullableUpdatePKICertIssuer) Set(val *UpdatePKICertIssuer) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdatePKICertIssuer) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdatePKICertIssuer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdatePKICertIssuer(val *UpdatePKICertIssuer) *NullableUpdatePKICertIssuer {
	return &NullableUpdatePKICertIssuer{value: val, isSet: true}
}

func (v NullableUpdatePKICertIssuer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdatePKICertIssuer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


