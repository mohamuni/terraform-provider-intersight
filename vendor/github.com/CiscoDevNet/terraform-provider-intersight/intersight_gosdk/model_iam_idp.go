/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-2024100405
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

// checks if the IamIdp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IamIdp{}

// IamIdp The SAML identity provider such as Cisco, that has been used to log in to Intersight.
type IamIdp struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Email domain name of the user for this IdP. When a user enters an email during login in the Intersight home page, the IdP is picked by matching this domain name with the email domain name for authentication.
	// Deprecated
	DomainName  *string  `json:"DomainName,omitempty" validate:"regexp=^$|^[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"`
	DomainNames []string `json:"DomainNames,omitempty"`
	// Setting that indicates whether 'Single Logout (SLO)' has been enabled for this IdP.
	EnableSingleLogout *bool `json:"EnableSingleLogout,omitempty"`
	// The Entity ID of the IdP. In SAML, the entity ID uniquely identifies the IdP or Service Provider.
	IdpEntityId *string `json:"IdpEntityId,omitempty"`
	// SAML metadata of the IdP.
	Metadata *string `json:"Metadata,omitempty"`
	// The name of the Identity Provider, for example Cisco, Okta, or OneID.
	Name *string `json:"Name,omitempty"`
	// When users attempt the Account URL login with an unverified Domain Name, they get a warning stating that they are logging in using an unverified Domain Name. Enable the slider if you do not wish to see the warning message.
	SkipWarning *bool `json:"SkipWarning,omitempty"`
	// Authentication protocol used by the IdP. * `saml` - Use SAML as the authentication protocol for sign-on. * `oidc` - Open ID connect to be used as an authentication protocol for sign-on. * `local` - The local authentication method to be used for sign-on. Local type is set to default for the Intersight Appliance IdP.
	Type       *string                           `json:"Type,omitempty"`
	Account    NullableIamAccountRelationship    `json:"Account,omitempty"`
	LdapPolicy NullableIamLdapPolicyRelationship `json:"LdapPolicy,omitempty"`
	System     NullableIamSystemRelationship     `json:"System,omitempty"`
	// An array of relationships to iamUserPreference resources.
	UserPreferences []IamUserPreferenceRelationship `json:"UserPreferences,omitempty"`
	// An array of relationships to iamUserSetting resources.
	UserSettings []IamUserSettingRelationship `json:"UserSettings,omitempty"`
	// An array of relationships to iamUserGroup resources.
	Usergroups []IamUserGroupRelationship `json:"Usergroups,omitempty"`
	// An array of relationships to iamUser resources.
	Users                []IamUserRelationship `json:"Users,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamIdp IamIdp

// NewIamIdp instantiates a new IamIdp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamIdp(classId string, objectType string) *IamIdp {
	this := IamIdp{}
	this.ClassId = classId
	this.ObjectType = objectType
	var type_ string = "saml"
	this.Type = &type_
	return &this
}

// NewIamIdpWithDefaults instantiates a new IamIdp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamIdpWithDefaults() *IamIdp {
	this := IamIdp{}
	var classId string = "iam.Idp"
	this.ClassId = classId
	var objectType string = "iam.Idp"
	this.ObjectType = objectType
	var type_ string = "saml"
	this.Type = &type_
	return &this
}

// GetClassId returns the ClassId field value
func (o *IamIdp) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IamIdp) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IamIdp) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "iam.Idp" of the ClassId field.
func (o *IamIdp) GetDefaultClassId() interface{} {
	return "iam.Idp"
}

// GetObjectType returns the ObjectType field value
func (o *IamIdp) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IamIdp) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IamIdp) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "iam.Idp" of the ObjectType field.
func (o *IamIdp) GetDefaultObjectType() interface{} {
	return "iam.Idp"
}

// GetDomainName returns the DomainName field value if set, zero value otherwise.
// Deprecated
func (o *IamIdp) GetDomainName() string {
	if o == nil || IsNil(o.DomainName) {
		var ret string
		return ret
	}
	return *o.DomainName
}

// GetDomainNameOk returns a tuple with the DomainName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *IamIdp) GetDomainNameOk() (*string, bool) {
	if o == nil || IsNil(o.DomainName) {
		return nil, false
	}
	return o.DomainName, true
}

// HasDomainName returns a boolean if a field has been set.
func (o *IamIdp) HasDomainName() bool {
	if o != nil && !IsNil(o.DomainName) {
		return true
	}

	return false
}

// SetDomainName gets a reference to the given string and assigns it to the DomainName field.
// Deprecated
func (o *IamIdp) SetDomainName(v string) {
	o.DomainName = &v
}

// GetDomainNames returns the DomainNames field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamIdp) GetDomainNames() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.DomainNames
}

// GetDomainNamesOk returns a tuple with the DomainNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamIdp) GetDomainNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.DomainNames) {
		return nil, false
	}
	return o.DomainNames, true
}

// HasDomainNames returns a boolean if a field has been set.
func (o *IamIdp) HasDomainNames() bool {
	if o != nil && !IsNil(o.DomainNames) {
		return true
	}

	return false
}

// SetDomainNames gets a reference to the given []string and assigns it to the DomainNames field.
func (o *IamIdp) SetDomainNames(v []string) {
	o.DomainNames = v
}

// GetEnableSingleLogout returns the EnableSingleLogout field value if set, zero value otherwise.
func (o *IamIdp) GetEnableSingleLogout() bool {
	if o == nil || IsNil(o.EnableSingleLogout) {
		var ret bool
		return ret
	}
	return *o.EnableSingleLogout
}

// GetEnableSingleLogoutOk returns a tuple with the EnableSingleLogout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamIdp) GetEnableSingleLogoutOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableSingleLogout) {
		return nil, false
	}
	return o.EnableSingleLogout, true
}

// HasEnableSingleLogout returns a boolean if a field has been set.
func (o *IamIdp) HasEnableSingleLogout() bool {
	if o != nil && !IsNil(o.EnableSingleLogout) {
		return true
	}

	return false
}

// SetEnableSingleLogout gets a reference to the given bool and assigns it to the EnableSingleLogout field.
func (o *IamIdp) SetEnableSingleLogout(v bool) {
	o.EnableSingleLogout = &v
}

// GetIdpEntityId returns the IdpEntityId field value if set, zero value otherwise.
func (o *IamIdp) GetIdpEntityId() string {
	if o == nil || IsNil(o.IdpEntityId) {
		var ret string
		return ret
	}
	return *o.IdpEntityId
}

// GetIdpEntityIdOk returns a tuple with the IdpEntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamIdp) GetIdpEntityIdOk() (*string, bool) {
	if o == nil || IsNil(o.IdpEntityId) {
		return nil, false
	}
	return o.IdpEntityId, true
}

// HasIdpEntityId returns a boolean if a field has been set.
func (o *IamIdp) HasIdpEntityId() bool {
	if o != nil && !IsNil(o.IdpEntityId) {
		return true
	}

	return false
}

// SetIdpEntityId gets a reference to the given string and assigns it to the IdpEntityId field.
func (o *IamIdp) SetIdpEntityId(v string) {
	o.IdpEntityId = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *IamIdp) GetMetadata() string {
	if o == nil || IsNil(o.Metadata) {
		var ret string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamIdp) GetMetadataOk() (*string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *IamIdp) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given string and assigns it to the Metadata field.
func (o *IamIdp) SetMetadata(v string) {
	o.Metadata = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IamIdp) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamIdp) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IamIdp) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IamIdp) SetName(v string) {
	o.Name = &v
}

// GetSkipWarning returns the SkipWarning field value if set, zero value otherwise.
func (o *IamIdp) GetSkipWarning() bool {
	if o == nil || IsNil(o.SkipWarning) {
		var ret bool
		return ret
	}
	return *o.SkipWarning
}

// GetSkipWarningOk returns a tuple with the SkipWarning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamIdp) GetSkipWarningOk() (*bool, bool) {
	if o == nil || IsNil(o.SkipWarning) {
		return nil, false
	}
	return o.SkipWarning, true
}

// HasSkipWarning returns a boolean if a field has been set.
func (o *IamIdp) HasSkipWarning() bool {
	if o != nil && !IsNil(o.SkipWarning) {
		return true
	}

	return false
}

// SetSkipWarning gets a reference to the given bool and assigns it to the SkipWarning field.
func (o *IamIdp) SetSkipWarning(v bool) {
	o.SkipWarning = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IamIdp) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamIdp) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IamIdp) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *IamIdp) SetType(v string) {
	o.Type = &v
}

// GetAccount returns the Account field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamIdp) GetAccount() IamAccountRelationship {
	if o == nil || IsNil(o.Account.Get()) {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account.Get()
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamIdp) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Account.Get(), o.Account.IsSet()
}

// HasAccount returns a boolean if a field has been set.
func (o *IamIdp) HasAccount() bool {
	if o != nil && o.Account.IsSet() {
		return true
	}

	return false
}

// SetAccount gets a reference to the given NullableIamAccountRelationship and assigns it to the Account field.
func (o *IamIdp) SetAccount(v IamAccountRelationship) {
	o.Account.Set(&v)
}

// SetAccountNil sets the value for Account to be an explicit nil
func (o *IamIdp) SetAccountNil() {
	o.Account.Set(nil)
}

// UnsetAccount ensures that no value is present for Account, not even an explicit nil
func (o *IamIdp) UnsetAccount() {
	o.Account.Unset()
}

// GetLdapPolicy returns the LdapPolicy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamIdp) GetLdapPolicy() IamLdapPolicyRelationship {
	if o == nil || IsNil(o.LdapPolicy.Get()) {
		var ret IamLdapPolicyRelationship
		return ret
	}
	return *o.LdapPolicy.Get()
}

// GetLdapPolicyOk returns a tuple with the LdapPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamIdp) GetLdapPolicyOk() (*IamLdapPolicyRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.LdapPolicy.Get(), o.LdapPolicy.IsSet()
}

// HasLdapPolicy returns a boolean if a field has been set.
func (o *IamIdp) HasLdapPolicy() bool {
	if o != nil && o.LdapPolicy.IsSet() {
		return true
	}

	return false
}

// SetLdapPolicy gets a reference to the given NullableIamLdapPolicyRelationship and assigns it to the LdapPolicy field.
func (o *IamIdp) SetLdapPolicy(v IamLdapPolicyRelationship) {
	o.LdapPolicy.Set(&v)
}

// SetLdapPolicyNil sets the value for LdapPolicy to be an explicit nil
func (o *IamIdp) SetLdapPolicyNil() {
	o.LdapPolicy.Set(nil)
}

// UnsetLdapPolicy ensures that no value is present for LdapPolicy, not even an explicit nil
func (o *IamIdp) UnsetLdapPolicy() {
	o.LdapPolicy.Unset()
}

// GetSystem returns the System field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamIdp) GetSystem() IamSystemRelationship {
	if o == nil || IsNil(o.System.Get()) {
		var ret IamSystemRelationship
		return ret
	}
	return *o.System.Get()
}

// GetSystemOk returns a tuple with the System field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamIdp) GetSystemOk() (*IamSystemRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.System.Get(), o.System.IsSet()
}

// HasSystem returns a boolean if a field has been set.
func (o *IamIdp) HasSystem() bool {
	if o != nil && o.System.IsSet() {
		return true
	}

	return false
}

// SetSystem gets a reference to the given NullableIamSystemRelationship and assigns it to the System field.
func (o *IamIdp) SetSystem(v IamSystemRelationship) {
	o.System.Set(&v)
}

// SetSystemNil sets the value for System to be an explicit nil
func (o *IamIdp) SetSystemNil() {
	o.System.Set(nil)
}

// UnsetSystem ensures that no value is present for System, not even an explicit nil
func (o *IamIdp) UnsetSystem() {
	o.System.Unset()
}

// GetUserPreferences returns the UserPreferences field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamIdp) GetUserPreferences() []IamUserPreferenceRelationship {
	if o == nil {
		var ret []IamUserPreferenceRelationship
		return ret
	}
	return o.UserPreferences
}

// GetUserPreferencesOk returns a tuple with the UserPreferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamIdp) GetUserPreferencesOk() ([]IamUserPreferenceRelationship, bool) {
	if o == nil || IsNil(o.UserPreferences) {
		return nil, false
	}
	return o.UserPreferences, true
}

// HasUserPreferences returns a boolean if a field has been set.
func (o *IamIdp) HasUserPreferences() bool {
	if o != nil && !IsNil(o.UserPreferences) {
		return true
	}

	return false
}

// SetUserPreferences gets a reference to the given []IamUserPreferenceRelationship and assigns it to the UserPreferences field.
func (o *IamIdp) SetUserPreferences(v []IamUserPreferenceRelationship) {
	o.UserPreferences = v
}

// GetUserSettings returns the UserSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamIdp) GetUserSettings() []IamUserSettingRelationship {
	if o == nil {
		var ret []IamUserSettingRelationship
		return ret
	}
	return o.UserSettings
}

// GetUserSettingsOk returns a tuple with the UserSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamIdp) GetUserSettingsOk() ([]IamUserSettingRelationship, bool) {
	if o == nil || IsNil(o.UserSettings) {
		return nil, false
	}
	return o.UserSettings, true
}

// HasUserSettings returns a boolean if a field has been set.
func (o *IamIdp) HasUserSettings() bool {
	if o != nil && !IsNil(o.UserSettings) {
		return true
	}

	return false
}

// SetUserSettings gets a reference to the given []IamUserSettingRelationship and assigns it to the UserSettings field.
func (o *IamIdp) SetUserSettings(v []IamUserSettingRelationship) {
	o.UserSettings = v
}

// GetUsergroups returns the Usergroups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamIdp) GetUsergroups() []IamUserGroupRelationship {
	if o == nil {
		var ret []IamUserGroupRelationship
		return ret
	}
	return o.Usergroups
}

// GetUsergroupsOk returns a tuple with the Usergroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamIdp) GetUsergroupsOk() ([]IamUserGroupRelationship, bool) {
	if o == nil || IsNil(o.Usergroups) {
		return nil, false
	}
	return o.Usergroups, true
}

// HasUsergroups returns a boolean if a field has been set.
func (o *IamIdp) HasUsergroups() bool {
	if o != nil && !IsNil(o.Usergroups) {
		return true
	}

	return false
}

// SetUsergroups gets a reference to the given []IamUserGroupRelationship and assigns it to the Usergroups field.
func (o *IamIdp) SetUsergroups(v []IamUserGroupRelationship) {
	o.Usergroups = v
}

// GetUsers returns the Users field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamIdp) GetUsers() []IamUserRelationship {
	if o == nil {
		var ret []IamUserRelationship
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamIdp) GetUsersOk() ([]IamUserRelationship, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *IamIdp) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []IamUserRelationship and assigns it to the Users field.
func (o *IamIdp) SetUsers(v []IamUserRelationship) {
	o.Users = v
}

func (o IamIdp) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IamIdp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.DomainName) {
		toSerialize["DomainName"] = o.DomainName
	}
	if o.DomainNames != nil {
		toSerialize["DomainNames"] = o.DomainNames
	}
	if !IsNil(o.EnableSingleLogout) {
		toSerialize["EnableSingleLogout"] = o.EnableSingleLogout
	}
	if !IsNil(o.IdpEntityId) {
		toSerialize["IdpEntityId"] = o.IdpEntityId
	}
	if !IsNil(o.Metadata) {
		toSerialize["Metadata"] = o.Metadata
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.SkipWarning) {
		toSerialize["SkipWarning"] = o.SkipWarning
	}
	if !IsNil(o.Type) {
		toSerialize["Type"] = o.Type
	}
	if o.Account.IsSet() {
		toSerialize["Account"] = o.Account.Get()
	}
	if o.LdapPolicy.IsSet() {
		toSerialize["LdapPolicy"] = o.LdapPolicy.Get()
	}
	if o.System.IsSet() {
		toSerialize["System"] = o.System.Get()
	}
	if o.UserPreferences != nil {
		toSerialize["UserPreferences"] = o.UserPreferences
	}
	if o.UserSettings != nil {
		toSerialize["UserSettings"] = o.UserSettings
	}
	if o.Usergroups != nil {
		toSerialize["Usergroups"] = o.Usergroups
	}
	if o.Users != nil {
		toSerialize["Users"] = o.Users
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IamIdp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ClassId",
		"ObjectType",
	}

	// defaultValueFuncMap captures the default values for required properties.
	// These values are used when required properties are missing from the payload.
	defaultValueFuncMap := map[string]func() interface{}{
		"ClassId":    o.GetDefaultClassId,
		"ObjectType": o.GetDefaultObjectType,
	}
	var defaultValueApplied bool
	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if value, exists := allProperties[requiredProperty]; !exists || value == "" {
			if _, ok := defaultValueFuncMap[requiredProperty]; ok {
				allProperties[requiredProperty] = defaultValueFuncMap[requiredProperty]()
				defaultValueApplied = true
			}
		}
		if value, exists := allProperties[requiredProperty]; !exists || value == "" {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	if defaultValueApplied {
		data, err = json.Marshal(allProperties)
		if err != nil {
			return err
		}
	}
	type IamIdpWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Email domain name of the user for this IdP. When a user enters an email during login in the Intersight home page, the IdP is picked by matching this domain name with the email domain name for authentication.
		// Deprecated
		DomainName  *string  `json:"DomainName,omitempty" validate:"regexp=^$|^[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"`
		DomainNames []string `json:"DomainNames,omitempty"`
		// Setting that indicates whether 'Single Logout (SLO)' has been enabled for this IdP.
		EnableSingleLogout *bool `json:"EnableSingleLogout,omitempty"`
		// The Entity ID of the IdP. In SAML, the entity ID uniquely identifies the IdP or Service Provider.
		IdpEntityId *string `json:"IdpEntityId,omitempty"`
		// SAML metadata of the IdP.
		Metadata *string `json:"Metadata,omitempty"`
		// The name of the Identity Provider, for example Cisco, Okta, or OneID.
		Name *string `json:"Name,omitempty"`
		// When users attempt the Account URL login with an unverified Domain Name, they get a warning stating that they are logging in using an unverified Domain Name. Enable the slider if you do not wish to see the warning message.
		SkipWarning *bool `json:"SkipWarning,omitempty"`
		// Authentication protocol used by the IdP. * `saml` - Use SAML as the authentication protocol for sign-on. * `oidc` - Open ID connect to be used as an authentication protocol for sign-on. * `local` - The local authentication method to be used for sign-on. Local type is set to default for the Intersight Appliance IdP.
		Type       *string                           `json:"Type,omitempty"`
		Account    NullableIamAccountRelationship    `json:"Account,omitempty"`
		LdapPolicy NullableIamLdapPolicyRelationship `json:"LdapPolicy,omitempty"`
		System     NullableIamSystemRelationship     `json:"System,omitempty"`
		// An array of relationships to iamUserPreference resources.
		UserPreferences []IamUserPreferenceRelationship `json:"UserPreferences,omitempty"`
		// An array of relationships to iamUserSetting resources.
		UserSettings []IamUserSettingRelationship `json:"UserSettings,omitempty"`
		// An array of relationships to iamUserGroup resources.
		Usergroups []IamUserGroupRelationship `json:"Usergroups,omitempty"`
		// An array of relationships to iamUser resources.
		Users []IamUserRelationship `json:"Users,omitempty"`
	}

	varIamIdpWithoutEmbeddedStruct := IamIdpWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varIamIdpWithoutEmbeddedStruct)
	if err == nil {
		varIamIdp := _IamIdp{}
		varIamIdp.ClassId = varIamIdpWithoutEmbeddedStruct.ClassId
		varIamIdp.ObjectType = varIamIdpWithoutEmbeddedStruct.ObjectType
		varIamIdp.DomainName = varIamIdpWithoutEmbeddedStruct.DomainName
		varIamIdp.DomainNames = varIamIdpWithoutEmbeddedStruct.DomainNames
		varIamIdp.EnableSingleLogout = varIamIdpWithoutEmbeddedStruct.EnableSingleLogout
		varIamIdp.IdpEntityId = varIamIdpWithoutEmbeddedStruct.IdpEntityId
		varIamIdp.Metadata = varIamIdpWithoutEmbeddedStruct.Metadata
		varIamIdp.Name = varIamIdpWithoutEmbeddedStruct.Name
		varIamIdp.SkipWarning = varIamIdpWithoutEmbeddedStruct.SkipWarning
		varIamIdp.Type = varIamIdpWithoutEmbeddedStruct.Type
		varIamIdp.Account = varIamIdpWithoutEmbeddedStruct.Account
		varIamIdp.LdapPolicy = varIamIdpWithoutEmbeddedStruct.LdapPolicy
		varIamIdp.System = varIamIdpWithoutEmbeddedStruct.System
		varIamIdp.UserPreferences = varIamIdpWithoutEmbeddedStruct.UserPreferences
		varIamIdp.UserSettings = varIamIdpWithoutEmbeddedStruct.UserSettings
		varIamIdp.Usergroups = varIamIdpWithoutEmbeddedStruct.Usergroups
		varIamIdp.Users = varIamIdpWithoutEmbeddedStruct.Users
		*o = IamIdp(varIamIdp)
	} else {
		return err
	}

	varIamIdp := _IamIdp{}

	err = json.Unmarshal(data, &varIamIdp)
	if err == nil {
		o.MoBaseMo = varIamIdp.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DomainName")
		delete(additionalProperties, "DomainNames")
		delete(additionalProperties, "EnableSingleLogout")
		delete(additionalProperties, "IdpEntityId")
		delete(additionalProperties, "Metadata")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "SkipWarning")
		delete(additionalProperties, "Type")
		delete(additionalProperties, "Account")
		delete(additionalProperties, "LdapPolicy")
		delete(additionalProperties, "System")
		delete(additionalProperties, "UserPreferences")
		delete(additionalProperties, "UserSettings")
		delete(additionalProperties, "Usergroups")
		delete(additionalProperties, "Users")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIamIdp struct {
	value *IamIdp
	isSet bool
}

func (v NullableIamIdp) Get() *IamIdp {
	return v.value
}

func (v *NullableIamIdp) Set(val *IamIdp) {
	v.value = val
	v.isSet = true
}

func (v NullableIamIdp) IsSet() bool {
	return v.isSet
}

func (v *NullableIamIdp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamIdp(val *IamIdp) *NullableIamIdp {
	return &NullableIamIdp{value: val, isSet: true}
}

func (v NullableIamIdp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamIdp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
