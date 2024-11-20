/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-2024101709
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

// checks if the CloudTfcOrganization type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudTfcOrganization{}

// CloudTfcOrganization Organizations are a shared space for teams to collaborate on workspaces in Terraform Cloud.
type CloudTfcOrganization struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The max number of active agents allowed in this organization.
	AgentCeiling *int64 `json:"AgentCeiling,omitempty"`
	// The admin email address associated with a user under this organization.
	Email *string `json:"Email,omitempty"`
	// The ID of the organization.
	Identity *string `json:"Identity,omitempty"`
	// The name of the organization.
	Name *string `json:"Name,omitempty"`
	// The number of teams under this organization.
	NumTeams *int64 `json:"NumTeams,omitempty"`
	// The number of users in this organization.
	NumUsers *int64 `json:"NumUsers,omitempty"`
	// The max number of simultaneous runs allowed in this organization.
	RunCeiling *int64 `json:"RunCeiling,omitempty"`
	// Total number of VCS providers in the organization.
	VcsProviders         *int64                          `json:"VcsProviders,omitempty"`
	Target               NullableAssetTargetRelationship `json:"Target,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CloudTfcOrganization CloudTfcOrganization

// NewCloudTfcOrganization instantiates a new CloudTfcOrganization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudTfcOrganization(classId string, objectType string) *CloudTfcOrganization {
	this := CloudTfcOrganization{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCloudTfcOrganizationWithDefaults instantiates a new CloudTfcOrganization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudTfcOrganizationWithDefaults() *CloudTfcOrganization {
	this := CloudTfcOrganization{}
	var classId string = "cloud.TfcOrganization"
	this.ClassId = classId
	var objectType string = "cloud.TfcOrganization"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CloudTfcOrganization) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CloudTfcOrganization) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CloudTfcOrganization) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "cloud.TfcOrganization" of the ClassId field.
func (o *CloudTfcOrganization) GetDefaultClassId() interface{} {
	return "cloud.TfcOrganization"
}

// GetObjectType returns the ObjectType field value
func (o *CloudTfcOrganization) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CloudTfcOrganization) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CloudTfcOrganization) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "cloud.TfcOrganization" of the ObjectType field.
func (o *CloudTfcOrganization) GetDefaultObjectType() interface{} {
	return "cloud.TfcOrganization"
}

// GetAgentCeiling returns the AgentCeiling field value if set, zero value otherwise.
func (o *CloudTfcOrganization) GetAgentCeiling() int64 {
	if o == nil || IsNil(o.AgentCeiling) {
		var ret int64
		return ret
	}
	return *o.AgentCeiling
}

// GetAgentCeilingOk returns a tuple with the AgentCeiling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudTfcOrganization) GetAgentCeilingOk() (*int64, bool) {
	if o == nil || IsNil(o.AgentCeiling) {
		return nil, false
	}
	return o.AgentCeiling, true
}

// HasAgentCeiling returns a boolean if a field has been set.
func (o *CloudTfcOrganization) HasAgentCeiling() bool {
	if o != nil && !IsNil(o.AgentCeiling) {
		return true
	}

	return false
}

// SetAgentCeiling gets a reference to the given int64 and assigns it to the AgentCeiling field.
func (o *CloudTfcOrganization) SetAgentCeiling(v int64) {
	o.AgentCeiling = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *CloudTfcOrganization) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudTfcOrganization) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *CloudTfcOrganization) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *CloudTfcOrganization) SetEmail(v string) {
	o.Email = &v
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *CloudTfcOrganization) GetIdentity() string {
	if o == nil || IsNil(o.Identity) {
		var ret string
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudTfcOrganization) GetIdentityOk() (*string, bool) {
	if o == nil || IsNil(o.Identity) {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *CloudTfcOrganization) HasIdentity() bool {
	if o != nil && !IsNil(o.Identity) {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given string and assigns it to the Identity field.
func (o *CloudTfcOrganization) SetIdentity(v string) {
	o.Identity = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CloudTfcOrganization) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudTfcOrganization) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CloudTfcOrganization) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CloudTfcOrganization) SetName(v string) {
	o.Name = &v
}

// GetNumTeams returns the NumTeams field value if set, zero value otherwise.
func (o *CloudTfcOrganization) GetNumTeams() int64 {
	if o == nil || IsNil(o.NumTeams) {
		var ret int64
		return ret
	}
	return *o.NumTeams
}

// GetNumTeamsOk returns a tuple with the NumTeams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudTfcOrganization) GetNumTeamsOk() (*int64, bool) {
	if o == nil || IsNil(o.NumTeams) {
		return nil, false
	}
	return o.NumTeams, true
}

// HasNumTeams returns a boolean if a field has been set.
func (o *CloudTfcOrganization) HasNumTeams() bool {
	if o != nil && !IsNil(o.NumTeams) {
		return true
	}

	return false
}

// SetNumTeams gets a reference to the given int64 and assigns it to the NumTeams field.
func (o *CloudTfcOrganization) SetNumTeams(v int64) {
	o.NumTeams = &v
}

// GetNumUsers returns the NumUsers field value if set, zero value otherwise.
func (o *CloudTfcOrganization) GetNumUsers() int64 {
	if o == nil || IsNil(o.NumUsers) {
		var ret int64
		return ret
	}
	return *o.NumUsers
}

// GetNumUsersOk returns a tuple with the NumUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudTfcOrganization) GetNumUsersOk() (*int64, bool) {
	if o == nil || IsNil(o.NumUsers) {
		return nil, false
	}
	return o.NumUsers, true
}

// HasNumUsers returns a boolean if a field has been set.
func (o *CloudTfcOrganization) HasNumUsers() bool {
	if o != nil && !IsNil(o.NumUsers) {
		return true
	}

	return false
}

// SetNumUsers gets a reference to the given int64 and assigns it to the NumUsers field.
func (o *CloudTfcOrganization) SetNumUsers(v int64) {
	o.NumUsers = &v
}

// GetRunCeiling returns the RunCeiling field value if set, zero value otherwise.
func (o *CloudTfcOrganization) GetRunCeiling() int64 {
	if o == nil || IsNil(o.RunCeiling) {
		var ret int64
		return ret
	}
	return *o.RunCeiling
}

// GetRunCeilingOk returns a tuple with the RunCeiling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudTfcOrganization) GetRunCeilingOk() (*int64, bool) {
	if o == nil || IsNil(o.RunCeiling) {
		return nil, false
	}
	return o.RunCeiling, true
}

// HasRunCeiling returns a boolean if a field has been set.
func (o *CloudTfcOrganization) HasRunCeiling() bool {
	if o != nil && !IsNil(o.RunCeiling) {
		return true
	}

	return false
}

// SetRunCeiling gets a reference to the given int64 and assigns it to the RunCeiling field.
func (o *CloudTfcOrganization) SetRunCeiling(v int64) {
	o.RunCeiling = &v
}

// GetVcsProviders returns the VcsProviders field value if set, zero value otherwise.
func (o *CloudTfcOrganization) GetVcsProviders() int64 {
	if o == nil || IsNil(o.VcsProviders) {
		var ret int64
		return ret
	}
	return *o.VcsProviders
}

// GetVcsProvidersOk returns a tuple with the VcsProviders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudTfcOrganization) GetVcsProvidersOk() (*int64, bool) {
	if o == nil || IsNil(o.VcsProviders) {
		return nil, false
	}
	return o.VcsProviders, true
}

// HasVcsProviders returns a boolean if a field has been set.
func (o *CloudTfcOrganization) HasVcsProviders() bool {
	if o != nil && !IsNil(o.VcsProviders) {
		return true
	}

	return false
}

// SetVcsProviders gets a reference to the given int64 and assigns it to the VcsProviders field.
func (o *CloudTfcOrganization) SetVcsProviders(v int64) {
	o.VcsProviders = &v
}

// GetTarget returns the Target field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudTfcOrganization) GetTarget() AssetTargetRelationship {
	if o == nil || IsNil(o.Target.Get()) {
		var ret AssetTargetRelationship
		return ret
	}
	return *o.Target.Get()
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudTfcOrganization) GetTargetOk() (*AssetTargetRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Target.Get(), o.Target.IsSet()
}

// HasTarget returns a boolean if a field has been set.
func (o *CloudTfcOrganization) HasTarget() bool {
	if o != nil && o.Target.IsSet() {
		return true
	}

	return false
}

// SetTarget gets a reference to the given NullableAssetTargetRelationship and assigns it to the Target field.
func (o *CloudTfcOrganization) SetTarget(v AssetTargetRelationship) {
	o.Target.Set(&v)
}

// SetTargetNil sets the value for Target to be an explicit nil
func (o *CloudTfcOrganization) SetTargetNil() {
	o.Target.Set(nil)
}

// UnsetTarget ensures that no value is present for Target, not even an explicit nil
func (o *CloudTfcOrganization) UnsetTarget() {
	o.Target.Unset()
}

func (o CloudTfcOrganization) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudTfcOrganization) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.AgentCeiling) {
		toSerialize["AgentCeiling"] = o.AgentCeiling
	}
	if !IsNil(o.Email) {
		toSerialize["Email"] = o.Email
	}
	if !IsNil(o.Identity) {
		toSerialize["Identity"] = o.Identity
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.NumTeams) {
		toSerialize["NumTeams"] = o.NumTeams
	}
	if !IsNil(o.NumUsers) {
		toSerialize["NumUsers"] = o.NumUsers
	}
	if !IsNil(o.RunCeiling) {
		toSerialize["RunCeiling"] = o.RunCeiling
	}
	if !IsNil(o.VcsProviders) {
		toSerialize["VcsProviders"] = o.VcsProviders
	}
	if o.Target.IsSet() {
		toSerialize["Target"] = o.Target.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CloudTfcOrganization) UnmarshalJSON(data []byte) (err error) {
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
	type CloudTfcOrganizationWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The max number of active agents allowed in this organization.
		AgentCeiling *int64 `json:"AgentCeiling,omitempty"`
		// The admin email address associated with a user under this organization.
		Email *string `json:"Email,omitempty"`
		// The ID of the organization.
		Identity *string `json:"Identity,omitempty"`
		// The name of the organization.
		Name *string `json:"Name,omitempty"`
		// The number of teams under this organization.
		NumTeams *int64 `json:"NumTeams,omitempty"`
		// The number of users in this organization.
		NumUsers *int64 `json:"NumUsers,omitempty"`
		// The max number of simultaneous runs allowed in this organization.
		RunCeiling *int64 `json:"RunCeiling,omitempty"`
		// Total number of VCS providers in the organization.
		VcsProviders *int64                          `json:"VcsProviders,omitempty"`
		Target       NullableAssetTargetRelationship `json:"Target,omitempty"`
	}

	varCloudTfcOrganizationWithoutEmbeddedStruct := CloudTfcOrganizationWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varCloudTfcOrganizationWithoutEmbeddedStruct)
	if err == nil {
		varCloudTfcOrganization := _CloudTfcOrganization{}
		varCloudTfcOrganization.ClassId = varCloudTfcOrganizationWithoutEmbeddedStruct.ClassId
		varCloudTfcOrganization.ObjectType = varCloudTfcOrganizationWithoutEmbeddedStruct.ObjectType
		varCloudTfcOrganization.AgentCeiling = varCloudTfcOrganizationWithoutEmbeddedStruct.AgentCeiling
		varCloudTfcOrganization.Email = varCloudTfcOrganizationWithoutEmbeddedStruct.Email
		varCloudTfcOrganization.Identity = varCloudTfcOrganizationWithoutEmbeddedStruct.Identity
		varCloudTfcOrganization.Name = varCloudTfcOrganizationWithoutEmbeddedStruct.Name
		varCloudTfcOrganization.NumTeams = varCloudTfcOrganizationWithoutEmbeddedStruct.NumTeams
		varCloudTfcOrganization.NumUsers = varCloudTfcOrganizationWithoutEmbeddedStruct.NumUsers
		varCloudTfcOrganization.RunCeiling = varCloudTfcOrganizationWithoutEmbeddedStruct.RunCeiling
		varCloudTfcOrganization.VcsProviders = varCloudTfcOrganizationWithoutEmbeddedStruct.VcsProviders
		varCloudTfcOrganization.Target = varCloudTfcOrganizationWithoutEmbeddedStruct.Target
		*o = CloudTfcOrganization(varCloudTfcOrganization)
	} else {
		return err
	}

	varCloudTfcOrganization := _CloudTfcOrganization{}

	err = json.Unmarshal(data, &varCloudTfcOrganization)
	if err == nil {
		o.MoBaseMo = varCloudTfcOrganization.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AgentCeiling")
		delete(additionalProperties, "Email")
		delete(additionalProperties, "Identity")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "NumTeams")
		delete(additionalProperties, "NumUsers")
		delete(additionalProperties, "RunCeiling")
		delete(additionalProperties, "VcsProviders")
		delete(additionalProperties, "Target")

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

type NullableCloudTfcOrganization struct {
	value *CloudTfcOrganization
	isSet bool
}

func (v NullableCloudTfcOrganization) Get() *CloudTfcOrganization {
	return v.value
}

func (v *NullableCloudTfcOrganization) Set(val *CloudTfcOrganization) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudTfcOrganization) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudTfcOrganization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudTfcOrganization(val *CloudTfcOrganization) *NullableCloudTfcOrganization {
	return &NullableCloudTfcOrganization{value: val, isSet: true}
}

func (v NullableCloudTfcOrganization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudTfcOrganization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
