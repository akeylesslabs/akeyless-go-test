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

// ConfigHash struct for ConfigHash
type ConfigHash struct {
	Admins *string `json:"admins,omitempty"`
	Cache *string `json:"cache,omitempty"`
	CustomerFragements *string `json:"customer_fragements,omitempty"`
	General *string `json:"general,omitempty"`
	K8sAuths *string `json:"k8s_auths,omitempty"`
	Kmip *string `json:"kmip,omitempty"`
	Ldap *string `json:"ldap,omitempty"`
	Leadership *string `json:"leadership,omitempty"`
	LogForwarding *string `json:"log_forwarding,omitempty"`
	MQueue *string `json:"m_queue,omitempty"`
	MigrationStatus *string `json:"migration_status,omitempty"`
	Migrations *string `json:"migrations,omitempty"`
	Producers *map[string]interface{} `json:"producers,omitempty"`
	ProducersStatus *string `json:"producers_status,omitempty"`
	Rotators *map[string]interface{} `json:"rotators,omitempty"`
	Saml *string `json:"saml,omitempty"`
	UniversalIdentity *string `json:"universal_identity,omitempty"`
}

// NewConfigHash instantiates a new ConfigHash object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigHash() *ConfigHash {
	this := ConfigHash{}
	return &this
}

// NewConfigHashWithDefaults instantiates a new ConfigHash object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigHashWithDefaults() *ConfigHash {
	this := ConfigHash{}
	return &this
}

// GetAdmins returns the Admins field value if set, zero value otherwise.
func (o *ConfigHash) GetAdmins() string {
	if o == nil || o.Admins == nil {
		var ret string
		return ret
	}
	return *o.Admins
}

// GetAdminsOk returns a tuple with the Admins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigHash) GetAdminsOk() (*string, bool) {
	if o == nil || o.Admins == nil {
		return nil, false
	}
	return o.Admins, true
}

// HasAdmins returns a boolean if a field has been set.
func (o *ConfigHash) HasAdmins() bool {
	if o != nil && o.Admins != nil {
		return true
	}

	return false
}

// SetAdmins gets a reference to the given string and assigns it to the Admins field.
func (o *ConfigHash) SetAdmins(v string) {
	o.Admins = &v
}

// GetCache returns the Cache field value if set, zero value otherwise.
func (o *ConfigHash) GetCache() string {
	if o == nil || o.Cache == nil {
		var ret string
		return ret
	}
	return *o.Cache
}

// GetCacheOk returns a tuple with the Cache field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigHash) GetCacheOk() (*string, bool) {
	if o == nil || o.Cache == nil {
		return nil, false
	}
	return o.Cache, true
}

// HasCache returns a boolean if a field has been set.
func (o *ConfigHash) HasCache() bool {
	if o != nil && o.Cache != nil {
		return true
	}

	return false
}

// SetCache gets a reference to the given string and assigns it to the Cache field.
func (o *ConfigHash) SetCache(v string) {
	o.Cache = &v
}

// GetCustomerFragements returns the CustomerFragements field value if set, zero value otherwise.
func (o *ConfigHash) GetCustomerFragements() string {
	if o == nil || o.CustomerFragements == nil {
		var ret string
		return ret
	}
	return *o.CustomerFragements
}

// GetCustomerFragementsOk returns a tuple with the CustomerFragements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigHash) GetCustomerFragementsOk() (*string, bool) {
	if o == nil || o.CustomerFragements == nil {
		return nil, false
	}
	return o.CustomerFragements, true
}

// HasCustomerFragements returns a boolean if a field has been set.
func (o *ConfigHash) HasCustomerFragements() bool {
	if o != nil && o.CustomerFragements != nil {
		return true
	}

	return false
}

// SetCustomerFragements gets a reference to the given string and assigns it to the CustomerFragements field.
func (o *ConfigHash) SetCustomerFragements(v string) {
	o.CustomerFragements = &v
}

// GetGeneral returns the General field value if set, zero value otherwise.
func (o *ConfigHash) GetGeneral() string {
	if o == nil || o.General == nil {
		var ret string
		return ret
	}
	return *o.General
}

// GetGeneralOk returns a tuple with the General field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigHash) GetGeneralOk() (*string, bool) {
	if o == nil || o.General == nil {
		return nil, false
	}
	return o.General, true
}

// HasGeneral returns a boolean if a field has been set.
func (o *ConfigHash) HasGeneral() bool {
	if o != nil && o.General != nil {
		return true
	}

	return false
}

// SetGeneral gets a reference to the given string and assigns it to the General field.
func (o *ConfigHash) SetGeneral(v string) {
	o.General = &v
}

// GetK8sAuths returns the K8sAuths field value if set, zero value otherwise.
func (o *ConfigHash) GetK8sAuths() string {
	if o == nil || o.K8sAuths == nil {
		var ret string
		return ret
	}
	return *o.K8sAuths
}

// GetK8sAuthsOk returns a tuple with the K8sAuths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigHash) GetK8sAuthsOk() (*string, bool) {
	if o == nil || o.K8sAuths == nil {
		return nil, false
	}
	return o.K8sAuths, true
}

// HasK8sAuths returns a boolean if a field has been set.
func (o *ConfigHash) HasK8sAuths() bool {
	if o != nil && o.K8sAuths != nil {
		return true
	}

	return false
}

// SetK8sAuths gets a reference to the given string and assigns it to the K8sAuths field.
func (o *ConfigHash) SetK8sAuths(v string) {
	o.K8sAuths = &v
}

// GetKmip returns the Kmip field value if set, zero value otherwise.
func (o *ConfigHash) GetKmip() string {
	if o == nil || o.Kmip == nil {
		var ret string
		return ret
	}
	return *o.Kmip
}

// GetKmipOk returns a tuple with the Kmip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigHash) GetKmipOk() (*string, bool) {
	if o == nil || o.Kmip == nil {
		return nil, false
	}
	return o.Kmip, true
}

// HasKmip returns a boolean if a field has been set.
func (o *ConfigHash) HasKmip() bool {
	if o != nil && o.Kmip != nil {
		return true
	}

	return false
}

// SetKmip gets a reference to the given string and assigns it to the Kmip field.
func (o *ConfigHash) SetKmip(v string) {
	o.Kmip = &v
}

// GetLdap returns the Ldap field value if set, zero value otherwise.
func (o *ConfigHash) GetLdap() string {
	if o == nil || o.Ldap == nil {
		var ret string
		return ret
	}
	return *o.Ldap
}

// GetLdapOk returns a tuple with the Ldap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigHash) GetLdapOk() (*string, bool) {
	if o == nil || o.Ldap == nil {
		return nil, false
	}
	return o.Ldap, true
}

// HasLdap returns a boolean if a field has been set.
func (o *ConfigHash) HasLdap() bool {
	if o != nil && o.Ldap != nil {
		return true
	}

	return false
}

// SetLdap gets a reference to the given string and assigns it to the Ldap field.
func (o *ConfigHash) SetLdap(v string) {
	o.Ldap = &v
}

// GetLeadership returns the Leadership field value if set, zero value otherwise.
func (o *ConfigHash) GetLeadership() string {
	if o == nil || o.Leadership == nil {
		var ret string
		return ret
	}
	return *o.Leadership
}

// GetLeadershipOk returns a tuple with the Leadership field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigHash) GetLeadershipOk() (*string, bool) {
	if o == nil || o.Leadership == nil {
		return nil, false
	}
	return o.Leadership, true
}

// HasLeadership returns a boolean if a field has been set.
func (o *ConfigHash) HasLeadership() bool {
	if o != nil && o.Leadership != nil {
		return true
	}

	return false
}

// SetLeadership gets a reference to the given string and assigns it to the Leadership field.
func (o *ConfigHash) SetLeadership(v string) {
	o.Leadership = &v
}

// GetLogForwarding returns the LogForwarding field value if set, zero value otherwise.
func (o *ConfigHash) GetLogForwarding() string {
	if o == nil || o.LogForwarding == nil {
		var ret string
		return ret
	}
	return *o.LogForwarding
}

// GetLogForwardingOk returns a tuple with the LogForwarding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigHash) GetLogForwardingOk() (*string, bool) {
	if o == nil || o.LogForwarding == nil {
		return nil, false
	}
	return o.LogForwarding, true
}

// HasLogForwarding returns a boolean if a field has been set.
func (o *ConfigHash) HasLogForwarding() bool {
	if o != nil && o.LogForwarding != nil {
		return true
	}

	return false
}

// SetLogForwarding gets a reference to the given string and assigns it to the LogForwarding field.
func (o *ConfigHash) SetLogForwarding(v string) {
	o.LogForwarding = &v
}

// GetMQueue returns the MQueue field value if set, zero value otherwise.
func (o *ConfigHash) GetMQueue() string {
	if o == nil || o.MQueue == nil {
		var ret string
		return ret
	}
	return *o.MQueue
}

// GetMQueueOk returns a tuple with the MQueue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigHash) GetMQueueOk() (*string, bool) {
	if o == nil || o.MQueue == nil {
		return nil, false
	}
	return o.MQueue, true
}

// HasMQueue returns a boolean if a field has been set.
func (o *ConfigHash) HasMQueue() bool {
	if o != nil && o.MQueue != nil {
		return true
	}

	return false
}

// SetMQueue gets a reference to the given string and assigns it to the MQueue field.
func (o *ConfigHash) SetMQueue(v string) {
	o.MQueue = &v
}

// GetMigrationStatus returns the MigrationStatus field value if set, zero value otherwise.
func (o *ConfigHash) GetMigrationStatus() string {
	if o == nil || o.MigrationStatus == nil {
		var ret string
		return ret
	}
	return *o.MigrationStatus
}

// GetMigrationStatusOk returns a tuple with the MigrationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigHash) GetMigrationStatusOk() (*string, bool) {
	if o == nil || o.MigrationStatus == nil {
		return nil, false
	}
	return o.MigrationStatus, true
}

// HasMigrationStatus returns a boolean if a field has been set.
func (o *ConfigHash) HasMigrationStatus() bool {
	if o != nil && o.MigrationStatus != nil {
		return true
	}

	return false
}

// SetMigrationStatus gets a reference to the given string and assigns it to the MigrationStatus field.
func (o *ConfigHash) SetMigrationStatus(v string) {
	o.MigrationStatus = &v
}

// GetMigrations returns the Migrations field value if set, zero value otherwise.
func (o *ConfigHash) GetMigrations() string {
	if o == nil || o.Migrations == nil {
		var ret string
		return ret
	}
	return *o.Migrations
}

// GetMigrationsOk returns a tuple with the Migrations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigHash) GetMigrationsOk() (*string, bool) {
	if o == nil || o.Migrations == nil {
		return nil, false
	}
	return o.Migrations, true
}

// HasMigrations returns a boolean if a field has been set.
func (o *ConfigHash) HasMigrations() bool {
	if o != nil && o.Migrations != nil {
		return true
	}

	return false
}

// SetMigrations gets a reference to the given string and assigns it to the Migrations field.
func (o *ConfigHash) SetMigrations(v string) {
	o.Migrations = &v
}

// GetProducers returns the Producers field value if set, zero value otherwise.
func (o *ConfigHash) GetProducers() map[string]interface{} {
	if o == nil || o.Producers == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Producers
}

// GetProducersOk returns a tuple with the Producers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigHash) GetProducersOk() (*map[string]interface{}, bool) {
	if o == nil || o.Producers == nil {
		return nil, false
	}
	return o.Producers, true
}

// HasProducers returns a boolean if a field has been set.
func (o *ConfigHash) HasProducers() bool {
	if o != nil && o.Producers != nil {
		return true
	}

	return false
}

// SetProducers gets a reference to the given map[string]interface{} and assigns it to the Producers field.
func (o *ConfigHash) SetProducers(v map[string]interface{}) {
	o.Producers = &v
}

// GetProducersStatus returns the ProducersStatus field value if set, zero value otherwise.
func (o *ConfigHash) GetProducersStatus() string {
	if o == nil || o.ProducersStatus == nil {
		var ret string
		return ret
	}
	return *o.ProducersStatus
}

// GetProducersStatusOk returns a tuple with the ProducersStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigHash) GetProducersStatusOk() (*string, bool) {
	if o == nil || o.ProducersStatus == nil {
		return nil, false
	}
	return o.ProducersStatus, true
}

// HasProducersStatus returns a boolean if a field has been set.
func (o *ConfigHash) HasProducersStatus() bool {
	if o != nil && o.ProducersStatus != nil {
		return true
	}

	return false
}

// SetProducersStatus gets a reference to the given string and assigns it to the ProducersStatus field.
func (o *ConfigHash) SetProducersStatus(v string) {
	o.ProducersStatus = &v
}

// GetRotators returns the Rotators field value if set, zero value otherwise.
func (o *ConfigHash) GetRotators() map[string]interface{} {
	if o == nil || o.Rotators == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Rotators
}

// GetRotatorsOk returns a tuple with the Rotators field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigHash) GetRotatorsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Rotators == nil {
		return nil, false
	}
	return o.Rotators, true
}

// HasRotators returns a boolean if a field has been set.
func (o *ConfigHash) HasRotators() bool {
	if o != nil && o.Rotators != nil {
		return true
	}

	return false
}

// SetRotators gets a reference to the given map[string]interface{} and assigns it to the Rotators field.
func (o *ConfigHash) SetRotators(v map[string]interface{}) {
	o.Rotators = &v
}

// GetSaml returns the Saml field value if set, zero value otherwise.
func (o *ConfigHash) GetSaml() string {
	if o == nil || o.Saml == nil {
		var ret string
		return ret
	}
	return *o.Saml
}

// GetSamlOk returns a tuple with the Saml field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigHash) GetSamlOk() (*string, bool) {
	if o == nil || o.Saml == nil {
		return nil, false
	}
	return o.Saml, true
}

// HasSaml returns a boolean if a field has been set.
func (o *ConfigHash) HasSaml() bool {
	if o != nil && o.Saml != nil {
		return true
	}

	return false
}

// SetSaml gets a reference to the given string and assigns it to the Saml field.
func (o *ConfigHash) SetSaml(v string) {
	o.Saml = &v
}

// GetUniversalIdentity returns the UniversalIdentity field value if set, zero value otherwise.
func (o *ConfigHash) GetUniversalIdentity() string {
	if o == nil || o.UniversalIdentity == nil {
		var ret string
		return ret
	}
	return *o.UniversalIdentity
}

// GetUniversalIdentityOk returns a tuple with the UniversalIdentity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigHash) GetUniversalIdentityOk() (*string, bool) {
	if o == nil || o.UniversalIdentity == nil {
		return nil, false
	}
	return o.UniversalIdentity, true
}

// HasUniversalIdentity returns a boolean if a field has been set.
func (o *ConfigHash) HasUniversalIdentity() bool {
	if o != nil && o.UniversalIdentity != nil {
		return true
	}

	return false
}

// SetUniversalIdentity gets a reference to the given string and assigns it to the UniversalIdentity field.
func (o *ConfigHash) SetUniversalIdentity(v string) {
	o.UniversalIdentity = &v
}

func (o ConfigHash) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Admins != nil {
		toSerialize["admins"] = o.Admins
	}
	if o.Cache != nil {
		toSerialize["cache"] = o.Cache
	}
	if o.CustomerFragements != nil {
		toSerialize["customer_fragements"] = o.CustomerFragements
	}
	if o.General != nil {
		toSerialize["general"] = o.General
	}
	if o.K8sAuths != nil {
		toSerialize["k8s_auths"] = o.K8sAuths
	}
	if o.Kmip != nil {
		toSerialize["kmip"] = o.Kmip
	}
	if o.Ldap != nil {
		toSerialize["ldap"] = o.Ldap
	}
	if o.Leadership != nil {
		toSerialize["leadership"] = o.Leadership
	}
	if o.LogForwarding != nil {
		toSerialize["log_forwarding"] = o.LogForwarding
	}
	if o.MQueue != nil {
		toSerialize["m_queue"] = o.MQueue
	}
	if o.MigrationStatus != nil {
		toSerialize["migration_status"] = o.MigrationStatus
	}
	if o.Migrations != nil {
		toSerialize["migrations"] = o.Migrations
	}
	if o.Producers != nil {
		toSerialize["producers"] = o.Producers
	}
	if o.ProducersStatus != nil {
		toSerialize["producers_status"] = o.ProducersStatus
	}
	if o.Rotators != nil {
		toSerialize["rotators"] = o.Rotators
	}
	if o.Saml != nil {
		toSerialize["saml"] = o.Saml
	}
	if o.UniversalIdentity != nil {
		toSerialize["universal_identity"] = o.UniversalIdentity
	}
	return json.Marshal(toSerialize)
}

type NullableConfigHash struct {
	value *ConfigHash
	isSet bool
}

func (v NullableConfigHash) Get() *ConfigHash {
	return v.value
}

func (v *NullableConfigHash) Set(val *ConfigHash) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigHash) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigHash) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigHash(val *ConfigHash) *NullableConfigHash {
	return &NullableConfigHash{value: val, isSet: true}
}

func (v NullableConfigHash) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigHash) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


