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

// AuthMethodAccessInfo struct for AuthMethodAccessInfo
type AuthMethodAccessInfo struct {
	AccessExpires *int64 `json:"access_expires,omitempty"`
	// for accounts where AccessId holds encrypted email this field will hold generated AccessId, for accounts based on regular AccessId it will be equal to accessId itself
	AccessIdAlias *string `json:"access_id_alias,omitempty"`
	ApiKeyAccessRules *APIKeyAccessRules `json:"api_key_access_rules,omitempty"`
	AwsIamAccessRules *AWSIAMAccessRules `json:"aws_iam_access_rules,omitempty"`
	AzureAdAccessRules *AzureADAccessRules `json:"azure_ad_access_rules,omitempty"`
	CertAccessRules *CertAccessRules `json:"cert_access_rules,omitempty"`
	CidrWhitelist *string `json:"cidr_whitelist,omitempty"`
	EmailPassAccessRules *EmailPassAccessRules `json:"email_pass_access_rules,omitempty"`
	// if true the role associated with this auth method must include sub claims
	ForceSubClaims *bool `json:"force_sub_claims,omitempty"`
	GcpAccessRules *GCPAccessRules `json:"gcp_access_rules,omitempty"`
	GwCidrWhitelist *string `json:"gw_cidr_whitelist,omitempty"`
	HuaweiAccessRules *HuaweiAccessRules `json:"huawei_access_rules,omitempty"`
	JwtTtl *int64 `json:"jwt_ttl,omitempty"`
	K8sAccessRules *KubernetesAccessRules `json:"k8s_access_rules,omitempty"`
	LdapAccessRules *LDAPAccessRules `json:"ldap_access_rules,omitempty"`
	Oauth2AccessRules *OAuth2AccessRules `json:"oauth2_access_rules,omitempty"`
	OidcAccessRules *OIDCAccessRules `json:"oidc_access_rules,omitempty"`
	RulesType *string `json:"rules_type,omitempty"`
	SamlAccessRules *SAMLAccessRules `json:"saml_access_rules,omitempty"`
	UniversalIdentityAccessRules *UniversalIdentityAccessRules `json:"universal_identity_access_rules,omitempty"`
}

// NewAuthMethodAccessInfo instantiates a new AuthMethodAccessInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthMethodAccessInfo() *AuthMethodAccessInfo {
	this := AuthMethodAccessInfo{}
	return &this
}

// NewAuthMethodAccessInfoWithDefaults instantiates a new AuthMethodAccessInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthMethodAccessInfoWithDefaults() *AuthMethodAccessInfo {
	this := AuthMethodAccessInfo{}
	return &this
}

// GetAccessExpires returns the AccessExpires field value if set, zero value otherwise.
func (o *AuthMethodAccessInfo) GetAccessExpires() int64 {
	if o == nil || o.AccessExpires == nil {
		var ret int64
		return ret
	}
	return *o.AccessExpires
}

// GetAccessExpiresOk returns a tuple with the AccessExpires field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthMethodAccessInfo) GetAccessExpiresOk() (*int64, bool) {
	if o == nil || o.AccessExpires == nil {
		return nil, false
	}
	return o.AccessExpires, true
}

// HasAccessExpires returns a boolean if a field has been set.
func (o *AuthMethodAccessInfo) HasAccessExpires() bool {
	if o != nil && o.AccessExpires != nil {
		return true
	}

	return false
}

// SetAccessExpires gets a reference to the given int64 and assigns it to the AccessExpires field.
func (o *AuthMethodAccessInfo) SetAccessExpires(v int64) {
	o.AccessExpires = &v
}

// GetAccessIdAlias returns the AccessIdAlias field value if set, zero value otherwise.
func (o *AuthMethodAccessInfo) GetAccessIdAlias() string {
	if o == nil || o.AccessIdAlias == nil {
		var ret string
		return ret
	}
	return *o.AccessIdAlias
}

// GetAccessIdAliasOk returns a tuple with the AccessIdAlias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthMethodAccessInfo) GetAccessIdAliasOk() (*string, bool) {
	if o == nil || o.AccessIdAlias == nil {
		return nil, false
	}
	return o.AccessIdAlias, true
}

// HasAccessIdAlias returns a boolean if a field has been set.
func (o *AuthMethodAccessInfo) HasAccessIdAlias() bool {
	if o != nil && o.AccessIdAlias != nil {
		return true
	}

	return false
}

// SetAccessIdAlias gets a reference to the given string and assigns it to the AccessIdAlias field.
func (o *AuthMethodAccessInfo) SetAccessIdAlias(v string) {
	o.AccessIdAlias = &v
}

// GetApiKeyAccessRules returns the ApiKeyAccessRules field value if set, zero value otherwise.
func (o *AuthMethodAccessInfo) GetApiKeyAccessRules() APIKeyAccessRules {
	if o == nil || o.ApiKeyAccessRules == nil {
		var ret APIKeyAccessRules
		return ret
	}
	return *o.ApiKeyAccessRules
}

// GetApiKeyAccessRulesOk returns a tuple with the ApiKeyAccessRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthMethodAccessInfo) GetApiKeyAccessRulesOk() (*APIKeyAccessRules, bool) {
	if o == nil || o.ApiKeyAccessRules == nil {
		return nil, false
	}
	return o.ApiKeyAccessRules, true
}

// HasApiKeyAccessRules returns a boolean if a field has been set.
func (o *AuthMethodAccessInfo) HasApiKeyAccessRules() bool {
	if o != nil && o.ApiKeyAccessRules != nil {
		return true
	}

	return false
}

// SetApiKeyAccessRules gets a reference to the given APIKeyAccessRules and assigns it to the ApiKeyAccessRules field.
func (o *AuthMethodAccessInfo) SetApiKeyAccessRules(v APIKeyAccessRules) {
	o.ApiKeyAccessRules = &v
}

// GetAwsIamAccessRules returns the AwsIamAccessRules field value if set, zero value otherwise.
func (o *AuthMethodAccessInfo) GetAwsIamAccessRules() AWSIAMAccessRules {
	if o == nil || o.AwsIamAccessRules == nil {
		var ret AWSIAMAccessRules
		return ret
	}
	return *o.AwsIamAccessRules
}

// GetAwsIamAccessRulesOk returns a tuple with the AwsIamAccessRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthMethodAccessInfo) GetAwsIamAccessRulesOk() (*AWSIAMAccessRules, bool) {
	if o == nil || o.AwsIamAccessRules == nil {
		return nil, false
	}
	return o.AwsIamAccessRules, true
}

// HasAwsIamAccessRules returns a boolean if a field has been set.
func (o *AuthMethodAccessInfo) HasAwsIamAccessRules() bool {
	if o != nil && o.AwsIamAccessRules != nil {
		return true
	}

	return false
}

// SetAwsIamAccessRules gets a reference to the given AWSIAMAccessRules and assigns it to the AwsIamAccessRules field.
func (o *AuthMethodAccessInfo) SetAwsIamAccessRules(v AWSIAMAccessRules) {
	o.AwsIamAccessRules = &v
}

// GetAzureAdAccessRules returns the AzureAdAccessRules field value if set, zero value otherwise.
func (o *AuthMethodAccessInfo) GetAzureAdAccessRules() AzureADAccessRules {
	if o == nil || o.AzureAdAccessRules == nil {
		var ret AzureADAccessRules
		return ret
	}
	return *o.AzureAdAccessRules
}

// GetAzureAdAccessRulesOk returns a tuple with the AzureAdAccessRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthMethodAccessInfo) GetAzureAdAccessRulesOk() (*AzureADAccessRules, bool) {
	if o == nil || o.AzureAdAccessRules == nil {
		return nil, false
	}
	return o.AzureAdAccessRules, true
}

// HasAzureAdAccessRules returns a boolean if a field has been set.
func (o *AuthMethodAccessInfo) HasAzureAdAccessRules() bool {
	if o != nil && o.AzureAdAccessRules != nil {
		return true
	}

	return false
}

// SetAzureAdAccessRules gets a reference to the given AzureADAccessRules and assigns it to the AzureAdAccessRules field.
func (o *AuthMethodAccessInfo) SetAzureAdAccessRules(v AzureADAccessRules) {
	o.AzureAdAccessRules = &v
}

// GetCertAccessRules returns the CertAccessRules field value if set, zero value otherwise.
func (o *AuthMethodAccessInfo) GetCertAccessRules() CertAccessRules {
	if o == nil || o.CertAccessRules == nil {
		var ret CertAccessRules
		return ret
	}
	return *o.CertAccessRules
}

// GetCertAccessRulesOk returns a tuple with the CertAccessRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthMethodAccessInfo) GetCertAccessRulesOk() (*CertAccessRules, bool) {
	if o == nil || o.CertAccessRules == nil {
		return nil, false
	}
	return o.CertAccessRules, true
}

// HasCertAccessRules returns a boolean if a field has been set.
func (o *AuthMethodAccessInfo) HasCertAccessRules() bool {
	if o != nil && o.CertAccessRules != nil {
		return true
	}

	return false
}

// SetCertAccessRules gets a reference to the given CertAccessRules and assigns it to the CertAccessRules field.
func (o *AuthMethodAccessInfo) SetCertAccessRules(v CertAccessRules) {
	o.CertAccessRules = &v
}

// GetCidrWhitelist returns the CidrWhitelist field value if set, zero value otherwise.
func (o *AuthMethodAccessInfo) GetCidrWhitelist() string {
	if o == nil || o.CidrWhitelist == nil {
		var ret string
		return ret
	}
	return *o.CidrWhitelist
}

// GetCidrWhitelistOk returns a tuple with the CidrWhitelist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthMethodAccessInfo) GetCidrWhitelistOk() (*string, bool) {
	if o == nil || o.CidrWhitelist == nil {
		return nil, false
	}
	return o.CidrWhitelist, true
}

// HasCidrWhitelist returns a boolean if a field has been set.
func (o *AuthMethodAccessInfo) HasCidrWhitelist() bool {
	if o != nil && o.CidrWhitelist != nil {
		return true
	}

	return false
}

// SetCidrWhitelist gets a reference to the given string and assigns it to the CidrWhitelist field.
func (o *AuthMethodAccessInfo) SetCidrWhitelist(v string) {
	o.CidrWhitelist = &v
}

// GetEmailPassAccessRules returns the EmailPassAccessRules field value if set, zero value otherwise.
func (o *AuthMethodAccessInfo) GetEmailPassAccessRules() EmailPassAccessRules {
	if o == nil || o.EmailPassAccessRules == nil {
		var ret EmailPassAccessRules
		return ret
	}
	return *o.EmailPassAccessRules
}

// GetEmailPassAccessRulesOk returns a tuple with the EmailPassAccessRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthMethodAccessInfo) GetEmailPassAccessRulesOk() (*EmailPassAccessRules, bool) {
	if o == nil || o.EmailPassAccessRules == nil {
		return nil, false
	}
	return o.EmailPassAccessRules, true
}

// HasEmailPassAccessRules returns a boolean if a field has been set.
func (o *AuthMethodAccessInfo) HasEmailPassAccessRules() bool {
	if o != nil && o.EmailPassAccessRules != nil {
		return true
	}

	return false
}

// SetEmailPassAccessRules gets a reference to the given EmailPassAccessRules and assigns it to the EmailPassAccessRules field.
func (o *AuthMethodAccessInfo) SetEmailPassAccessRules(v EmailPassAccessRules) {
	o.EmailPassAccessRules = &v
}

// GetForceSubClaims returns the ForceSubClaims field value if set, zero value otherwise.
func (o *AuthMethodAccessInfo) GetForceSubClaims() bool {
	if o == nil || o.ForceSubClaims == nil {
		var ret bool
		return ret
	}
	return *o.ForceSubClaims
}

// GetForceSubClaimsOk returns a tuple with the ForceSubClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthMethodAccessInfo) GetForceSubClaimsOk() (*bool, bool) {
	if o == nil || o.ForceSubClaims == nil {
		return nil, false
	}
	return o.ForceSubClaims, true
}

// HasForceSubClaims returns a boolean if a field has been set.
func (o *AuthMethodAccessInfo) HasForceSubClaims() bool {
	if o != nil && o.ForceSubClaims != nil {
		return true
	}

	return false
}

// SetForceSubClaims gets a reference to the given bool and assigns it to the ForceSubClaims field.
func (o *AuthMethodAccessInfo) SetForceSubClaims(v bool) {
	o.ForceSubClaims = &v
}

// GetGcpAccessRules returns the GcpAccessRules field value if set, zero value otherwise.
func (o *AuthMethodAccessInfo) GetGcpAccessRules() GCPAccessRules {
	if o == nil || o.GcpAccessRules == nil {
		var ret GCPAccessRules
		return ret
	}
	return *o.GcpAccessRules
}

// GetGcpAccessRulesOk returns a tuple with the GcpAccessRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthMethodAccessInfo) GetGcpAccessRulesOk() (*GCPAccessRules, bool) {
	if o == nil || o.GcpAccessRules == nil {
		return nil, false
	}
	return o.GcpAccessRules, true
}

// HasGcpAccessRules returns a boolean if a field has been set.
func (o *AuthMethodAccessInfo) HasGcpAccessRules() bool {
	if o != nil && o.GcpAccessRules != nil {
		return true
	}

	return false
}

// SetGcpAccessRules gets a reference to the given GCPAccessRules and assigns it to the GcpAccessRules field.
func (o *AuthMethodAccessInfo) SetGcpAccessRules(v GCPAccessRules) {
	o.GcpAccessRules = &v
}

// GetGwCidrWhitelist returns the GwCidrWhitelist field value if set, zero value otherwise.
func (o *AuthMethodAccessInfo) GetGwCidrWhitelist() string {
	if o == nil || o.GwCidrWhitelist == nil {
		var ret string
		return ret
	}
	return *o.GwCidrWhitelist
}

// GetGwCidrWhitelistOk returns a tuple with the GwCidrWhitelist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthMethodAccessInfo) GetGwCidrWhitelistOk() (*string, bool) {
	if o == nil || o.GwCidrWhitelist == nil {
		return nil, false
	}
	return o.GwCidrWhitelist, true
}

// HasGwCidrWhitelist returns a boolean if a field has been set.
func (o *AuthMethodAccessInfo) HasGwCidrWhitelist() bool {
	if o != nil && o.GwCidrWhitelist != nil {
		return true
	}

	return false
}

// SetGwCidrWhitelist gets a reference to the given string and assigns it to the GwCidrWhitelist field.
func (o *AuthMethodAccessInfo) SetGwCidrWhitelist(v string) {
	o.GwCidrWhitelist = &v
}

// GetHuaweiAccessRules returns the HuaweiAccessRules field value if set, zero value otherwise.
func (o *AuthMethodAccessInfo) GetHuaweiAccessRules() HuaweiAccessRules {
	if o == nil || o.HuaweiAccessRules == nil {
		var ret HuaweiAccessRules
		return ret
	}
	return *o.HuaweiAccessRules
}

// GetHuaweiAccessRulesOk returns a tuple with the HuaweiAccessRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthMethodAccessInfo) GetHuaweiAccessRulesOk() (*HuaweiAccessRules, bool) {
	if o == nil || o.HuaweiAccessRules == nil {
		return nil, false
	}
	return o.HuaweiAccessRules, true
}

// HasHuaweiAccessRules returns a boolean if a field has been set.
func (o *AuthMethodAccessInfo) HasHuaweiAccessRules() bool {
	if o != nil && o.HuaweiAccessRules != nil {
		return true
	}

	return false
}

// SetHuaweiAccessRules gets a reference to the given HuaweiAccessRules and assigns it to the HuaweiAccessRules field.
func (o *AuthMethodAccessInfo) SetHuaweiAccessRules(v HuaweiAccessRules) {
	o.HuaweiAccessRules = &v
}

// GetJwtTtl returns the JwtTtl field value if set, zero value otherwise.
func (o *AuthMethodAccessInfo) GetJwtTtl() int64 {
	if o == nil || o.JwtTtl == nil {
		var ret int64
		return ret
	}
	return *o.JwtTtl
}

// GetJwtTtlOk returns a tuple with the JwtTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthMethodAccessInfo) GetJwtTtlOk() (*int64, bool) {
	if o == nil || o.JwtTtl == nil {
		return nil, false
	}
	return o.JwtTtl, true
}

// HasJwtTtl returns a boolean if a field has been set.
func (o *AuthMethodAccessInfo) HasJwtTtl() bool {
	if o != nil && o.JwtTtl != nil {
		return true
	}

	return false
}

// SetJwtTtl gets a reference to the given int64 and assigns it to the JwtTtl field.
func (o *AuthMethodAccessInfo) SetJwtTtl(v int64) {
	o.JwtTtl = &v
}

// GetK8sAccessRules returns the K8sAccessRules field value if set, zero value otherwise.
func (o *AuthMethodAccessInfo) GetK8sAccessRules() KubernetesAccessRules {
	if o == nil || o.K8sAccessRules == nil {
		var ret KubernetesAccessRules
		return ret
	}
	return *o.K8sAccessRules
}

// GetK8sAccessRulesOk returns a tuple with the K8sAccessRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthMethodAccessInfo) GetK8sAccessRulesOk() (*KubernetesAccessRules, bool) {
	if o == nil || o.K8sAccessRules == nil {
		return nil, false
	}
	return o.K8sAccessRules, true
}

// HasK8sAccessRules returns a boolean if a field has been set.
func (o *AuthMethodAccessInfo) HasK8sAccessRules() bool {
	if o != nil && o.K8sAccessRules != nil {
		return true
	}

	return false
}

// SetK8sAccessRules gets a reference to the given KubernetesAccessRules and assigns it to the K8sAccessRules field.
func (o *AuthMethodAccessInfo) SetK8sAccessRules(v KubernetesAccessRules) {
	o.K8sAccessRules = &v
}

// GetLdapAccessRules returns the LdapAccessRules field value if set, zero value otherwise.
func (o *AuthMethodAccessInfo) GetLdapAccessRules() LDAPAccessRules {
	if o == nil || o.LdapAccessRules == nil {
		var ret LDAPAccessRules
		return ret
	}
	return *o.LdapAccessRules
}

// GetLdapAccessRulesOk returns a tuple with the LdapAccessRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthMethodAccessInfo) GetLdapAccessRulesOk() (*LDAPAccessRules, bool) {
	if o == nil || o.LdapAccessRules == nil {
		return nil, false
	}
	return o.LdapAccessRules, true
}

// HasLdapAccessRules returns a boolean if a field has been set.
func (o *AuthMethodAccessInfo) HasLdapAccessRules() bool {
	if o != nil && o.LdapAccessRules != nil {
		return true
	}

	return false
}

// SetLdapAccessRules gets a reference to the given LDAPAccessRules and assigns it to the LdapAccessRules field.
func (o *AuthMethodAccessInfo) SetLdapAccessRules(v LDAPAccessRules) {
	o.LdapAccessRules = &v
}

// GetOauth2AccessRules returns the Oauth2AccessRules field value if set, zero value otherwise.
func (o *AuthMethodAccessInfo) GetOauth2AccessRules() OAuth2AccessRules {
	if o == nil || o.Oauth2AccessRules == nil {
		var ret OAuth2AccessRules
		return ret
	}
	return *o.Oauth2AccessRules
}

// GetOauth2AccessRulesOk returns a tuple with the Oauth2AccessRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthMethodAccessInfo) GetOauth2AccessRulesOk() (*OAuth2AccessRules, bool) {
	if o == nil || o.Oauth2AccessRules == nil {
		return nil, false
	}
	return o.Oauth2AccessRules, true
}

// HasOauth2AccessRules returns a boolean if a field has been set.
func (o *AuthMethodAccessInfo) HasOauth2AccessRules() bool {
	if o != nil && o.Oauth2AccessRules != nil {
		return true
	}

	return false
}

// SetOauth2AccessRules gets a reference to the given OAuth2AccessRules and assigns it to the Oauth2AccessRules field.
func (o *AuthMethodAccessInfo) SetOauth2AccessRules(v OAuth2AccessRules) {
	o.Oauth2AccessRules = &v
}

// GetOidcAccessRules returns the OidcAccessRules field value if set, zero value otherwise.
func (o *AuthMethodAccessInfo) GetOidcAccessRules() OIDCAccessRules {
	if o == nil || o.OidcAccessRules == nil {
		var ret OIDCAccessRules
		return ret
	}
	return *o.OidcAccessRules
}

// GetOidcAccessRulesOk returns a tuple with the OidcAccessRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthMethodAccessInfo) GetOidcAccessRulesOk() (*OIDCAccessRules, bool) {
	if o == nil || o.OidcAccessRules == nil {
		return nil, false
	}
	return o.OidcAccessRules, true
}

// HasOidcAccessRules returns a boolean if a field has been set.
func (o *AuthMethodAccessInfo) HasOidcAccessRules() bool {
	if o != nil && o.OidcAccessRules != nil {
		return true
	}

	return false
}

// SetOidcAccessRules gets a reference to the given OIDCAccessRules and assigns it to the OidcAccessRules field.
func (o *AuthMethodAccessInfo) SetOidcAccessRules(v OIDCAccessRules) {
	o.OidcAccessRules = &v
}

// GetRulesType returns the RulesType field value if set, zero value otherwise.
func (o *AuthMethodAccessInfo) GetRulesType() string {
	if o == nil || o.RulesType == nil {
		var ret string
		return ret
	}
	return *o.RulesType
}

// GetRulesTypeOk returns a tuple with the RulesType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthMethodAccessInfo) GetRulesTypeOk() (*string, bool) {
	if o == nil || o.RulesType == nil {
		return nil, false
	}
	return o.RulesType, true
}

// HasRulesType returns a boolean if a field has been set.
func (o *AuthMethodAccessInfo) HasRulesType() bool {
	if o != nil && o.RulesType != nil {
		return true
	}

	return false
}

// SetRulesType gets a reference to the given string and assigns it to the RulesType field.
func (o *AuthMethodAccessInfo) SetRulesType(v string) {
	o.RulesType = &v
}

// GetSamlAccessRules returns the SamlAccessRules field value if set, zero value otherwise.
func (o *AuthMethodAccessInfo) GetSamlAccessRules() SAMLAccessRules {
	if o == nil || o.SamlAccessRules == nil {
		var ret SAMLAccessRules
		return ret
	}
	return *o.SamlAccessRules
}

// GetSamlAccessRulesOk returns a tuple with the SamlAccessRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthMethodAccessInfo) GetSamlAccessRulesOk() (*SAMLAccessRules, bool) {
	if o == nil || o.SamlAccessRules == nil {
		return nil, false
	}
	return o.SamlAccessRules, true
}

// HasSamlAccessRules returns a boolean if a field has been set.
func (o *AuthMethodAccessInfo) HasSamlAccessRules() bool {
	if o != nil && o.SamlAccessRules != nil {
		return true
	}

	return false
}

// SetSamlAccessRules gets a reference to the given SAMLAccessRules and assigns it to the SamlAccessRules field.
func (o *AuthMethodAccessInfo) SetSamlAccessRules(v SAMLAccessRules) {
	o.SamlAccessRules = &v
}

// GetUniversalIdentityAccessRules returns the UniversalIdentityAccessRules field value if set, zero value otherwise.
func (o *AuthMethodAccessInfo) GetUniversalIdentityAccessRules() UniversalIdentityAccessRules {
	if o == nil || o.UniversalIdentityAccessRules == nil {
		var ret UniversalIdentityAccessRules
		return ret
	}
	return *o.UniversalIdentityAccessRules
}

// GetUniversalIdentityAccessRulesOk returns a tuple with the UniversalIdentityAccessRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthMethodAccessInfo) GetUniversalIdentityAccessRulesOk() (*UniversalIdentityAccessRules, bool) {
	if o == nil || o.UniversalIdentityAccessRules == nil {
		return nil, false
	}
	return o.UniversalIdentityAccessRules, true
}

// HasUniversalIdentityAccessRules returns a boolean if a field has been set.
func (o *AuthMethodAccessInfo) HasUniversalIdentityAccessRules() bool {
	if o != nil && o.UniversalIdentityAccessRules != nil {
		return true
	}

	return false
}

// SetUniversalIdentityAccessRules gets a reference to the given UniversalIdentityAccessRules and assigns it to the UniversalIdentityAccessRules field.
func (o *AuthMethodAccessInfo) SetUniversalIdentityAccessRules(v UniversalIdentityAccessRules) {
	o.UniversalIdentityAccessRules = &v
}

func (o AuthMethodAccessInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessExpires != nil {
		toSerialize["access_expires"] = o.AccessExpires
	}
	if o.AccessIdAlias != nil {
		toSerialize["access_id_alias"] = o.AccessIdAlias
	}
	if o.ApiKeyAccessRules != nil {
		toSerialize["api_key_access_rules"] = o.ApiKeyAccessRules
	}
	if o.AwsIamAccessRules != nil {
		toSerialize["aws_iam_access_rules"] = o.AwsIamAccessRules
	}
	if o.AzureAdAccessRules != nil {
		toSerialize["azure_ad_access_rules"] = o.AzureAdAccessRules
	}
	if o.CertAccessRules != nil {
		toSerialize["cert_access_rules"] = o.CertAccessRules
	}
	if o.CidrWhitelist != nil {
		toSerialize["cidr_whitelist"] = o.CidrWhitelist
	}
	if o.EmailPassAccessRules != nil {
		toSerialize["email_pass_access_rules"] = o.EmailPassAccessRules
	}
	if o.ForceSubClaims != nil {
		toSerialize["force_sub_claims"] = o.ForceSubClaims
	}
	if o.GcpAccessRules != nil {
		toSerialize["gcp_access_rules"] = o.GcpAccessRules
	}
	if o.GwCidrWhitelist != nil {
		toSerialize["gw_cidr_whitelist"] = o.GwCidrWhitelist
	}
	if o.HuaweiAccessRules != nil {
		toSerialize["huawei_access_rules"] = o.HuaweiAccessRules
	}
	if o.JwtTtl != nil {
		toSerialize["jwt_ttl"] = o.JwtTtl
	}
	if o.K8sAccessRules != nil {
		toSerialize["k8s_access_rules"] = o.K8sAccessRules
	}
	if o.LdapAccessRules != nil {
		toSerialize["ldap_access_rules"] = o.LdapAccessRules
	}
	if o.Oauth2AccessRules != nil {
		toSerialize["oauth2_access_rules"] = o.Oauth2AccessRules
	}
	if o.OidcAccessRules != nil {
		toSerialize["oidc_access_rules"] = o.OidcAccessRules
	}
	if o.RulesType != nil {
		toSerialize["rules_type"] = o.RulesType
	}
	if o.SamlAccessRules != nil {
		toSerialize["saml_access_rules"] = o.SamlAccessRules
	}
	if o.UniversalIdentityAccessRules != nil {
		toSerialize["universal_identity_access_rules"] = o.UniversalIdentityAccessRules
	}
	return json.Marshal(toSerialize)
}

type NullableAuthMethodAccessInfo struct {
	value *AuthMethodAccessInfo
	isSet bool
}

func (v NullableAuthMethodAccessInfo) Get() *AuthMethodAccessInfo {
	return v.value
}

func (v *NullableAuthMethodAccessInfo) Set(val *AuthMethodAccessInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthMethodAccessInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthMethodAccessInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthMethodAccessInfo(val *AuthMethodAccessInfo) *NullableAuthMethodAccessInfo {
	return &NullableAuthMethodAccessInfo{value: val, isSet: true}
}

func (v NullableAuthMethodAccessInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthMethodAccessInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


