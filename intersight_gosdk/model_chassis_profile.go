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

// checks if the ChassisProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChassisProfile{}

// ChassisProfile A profile specifying configuration settings for a chassis.
type ChassisProfile struct {
	ChassisBaseProfile
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType          string                            `json:"ObjectType"`
	ConfigChangeContext NullablePolicyConfigChangeContext `json:"ConfigChangeContext,omitempty"`
	ConfigChanges       NullablePolicyConfigChange        `json:"ConfigChanges,omitempty"`
	// User label assigned to the chassis profile.
	UserLabel         *string                              `json:"UserLabel,omitempty" validate:"regexp=^[ !#$%&\\\\(\\\\)\\\\*\\\\+,\\\\-\\\\.\\/:;\\\\?@\\\\[\\\\]_\\\\{\\\\|\\\\}~a-zA-Z0-9]*$"`
	AssignedChassis   NullableEquipmentChassisRelationship `json:"AssignedChassis,omitempty"`
	AssociatedChassis NullableEquipmentChassisRelationship `json:"AssociatedChassis,omitempty"`
	// An array of relationships to chassisConfigChangeDetail resources.
	ConfigChangeDetails []ChassisConfigChangeDetailRelationship      `json:"ConfigChangeDetails,omitempty"`
	Organization        NullableOrganizationOrganizationRelationship `json:"Organization,omitempty"`
	// An array of relationships to workflowWorkflowInfo resources.
	RunningWorkflows     []WorkflowWorkflowInfoRelationship `json:"RunningWorkflows,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ChassisProfile ChassisProfile

// NewChassisProfile instantiates a new ChassisProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChassisProfile(classId string, objectType string) *ChassisProfile {
	this := ChassisProfile{}
	this.ClassId = classId
	this.ObjectType = objectType
	var type_ string = "instance"
	this.Type = &type_
	var action string = "No-op"
	this.Action = &action
	var targetPlatform string = "FIAttached"
	this.TargetPlatform = &targetPlatform
	return &this
}

// NewChassisProfileWithDefaults instantiates a new ChassisProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChassisProfileWithDefaults() *ChassisProfile {
	this := ChassisProfile{}
	var classId string = "chassis.Profile"
	this.ClassId = classId
	var objectType string = "chassis.Profile"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ChassisProfile) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ChassisProfile) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ChassisProfile) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "chassis.Profile" of the ClassId field.
func (o *ChassisProfile) GetDefaultClassId() interface{} {
	return "chassis.Profile"
}

// GetObjectType returns the ObjectType field value
func (o *ChassisProfile) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ChassisProfile) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ChassisProfile) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "chassis.Profile" of the ObjectType field.
func (o *ChassisProfile) GetDefaultObjectType() interface{} {
	return "chassis.Profile"
}

// GetConfigChangeContext returns the ConfigChangeContext field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChassisProfile) GetConfigChangeContext() PolicyConfigChangeContext {
	if o == nil || IsNil(o.ConfigChangeContext.Get()) {
		var ret PolicyConfigChangeContext
		return ret
	}
	return *o.ConfigChangeContext.Get()
}

// GetConfigChangeContextOk returns a tuple with the ConfigChangeContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChassisProfile) GetConfigChangeContextOk() (*PolicyConfigChangeContext, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigChangeContext.Get(), o.ConfigChangeContext.IsSet()
}

// HasConfigChangeContext returns a boolean if a field has been set.
func (o *ChassisProfile) HasConfigChangeContext() bool {
	if o != nil && o.ConfigChangeContext.IsSet() {
		return true
	}

	return false
}

// SetConfigChangeContext gets a reference to the given NullablePolicyConfigChangeContext and assigns it to the ConfigChangeContext field.
func (o *ChassisProfile) SetConfigChangeContext(v PolicyConfigChangeContext) {
	o.ConfigChangeContext.Set(&v)
}

// SetConfigChangeContextNil sets the value for ConfigChangeContext to be an explicit nil
func (o *ChassisProfile) SetConfigChangeContextNil() {
	o.ConfigChangeContext.Set(nil)
}

// UnsetConfigChangeContext ensures that no value is present for ConfigChangeContext, not even an explicit nil
func (o *ChassisProfile) UnsetConfigChangeContext() {
	o.ConfigChangeContext.Unset()
}

// GetConfigChanges returns the ConfigChanges field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChassisProfile) GetConfigChanges() PolicyConfigChange {
	if o == nil || IsNil(o.ConfigChanges.Get()) {
		var ret PolicyConfigChange
		return ret
	}
	return *o.ConfigChanges.Get()
}

// GetConfigChangesOk returns a tuple with the ConfigChanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChassisProfile) GetConfigChangesOk() (*PolicyConfigChange, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigChanges.Get(), o.ConfigChanges.IsSet()
}

// HasConfigChanges returns a boolean if a field has been set.
func (o *ChassisProfile) HasConfigChanges() bool {
	if o != nil && o.ConfigChanges.IsSet() {
		return true
	}

	return false
}

// SetConfigChanges gets a reference to the given NullablePolicyConfigChange and assigns it to the ConfigChanges field.
func (o *ChassisProfile) SetConfigChanges(v PolicyConfigChange) {
	o.ConfigChanges.Set(&v)
}

// SetConfigChangesNil sets the value for ConfigChanges to be an explicit nil
func (o *ChassisProfile) SetConfigChangesNil() {
	o.ConfigChanges.Set(nil)
}

// UnsetConfigChanges ensures that no value is present for ConfigChanges, not even an explicit nil
func (o *ChassisProfile) UnsetConfigChanges() {
	o.ConfigChanges.Unset()
}

// GetUserLabel returns the UserLabel field value if set, zero value otherwise.
func (o *ChassisProfile) GetUserLabel() string {
	if o == nil || IsNil(o.UserLabel) {
		var ret string
		return ret
	}
	return *o.UserLabel
}

// GetUserLabelOk returns a tuple with the UserLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChassisProfile) GetUserLabelOk() (*string, bool) {
	if o == nil || IsNil(o.UserLabel) {
		return nil, false
	}
	return o.UserLabel, true
}

// HasUserLabel returns a boolean if a field has been set.
func (o *ChassisProfile) HasUserLabel() bool {
	if o != nil && !IsNil(o.UserLabel) {
		return true
	}

	return false
}

// SetUserLabel gets a reference to the given string and assigns it to the UserLabel field.
func (o *ChassisProfile) SetUserLabel(v string) {
	o.UserLabel = &v
}

// GetAssignedChassis returns the AssignedChassis field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChassisProfile) GetAssignedChassis() EquipmentChassisRelationship {
	if o == nil || IsNil(o.AssignedChassis.Get()) {
		var ret EquipmentChassisRelationship
		return ret
	}
	return *o.AssignedChassis.Get()
}

// GetAssignedChassisOk returns a tuple with the AssignedChassis field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChassisProfile) GetAssignedChassisOk() (*EquipmentChassisRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssignedChassis.Get(), o.AssignedChassis.IsSet()
}

// HasAssignedChassis returns a boolean if a field has been set.
func (o *ChassisProfile) HasAssignedChassis() bool {
	if o != nil && o.AssignedChassis.IsSet() {
		return true
	}

	return false
}

// SetAssignedChassis gets a reference to the given NullableEquipmentChassisRelationship and assigns it to the AssignedChassis field.
func (o *ChassisProfile) SetAssignedChassis(v EquipmentChassisRelationship) {
	o.AssignedChassis.Set(&v)
}

// SetAssignedChassisNil sets the value for AssignedChassis to be an explicit nil
func (o *ChassisProfile) SetAssignedChassisNil() {
	o.AssignedChassis.Set(nil)
}

// UnsetAssignedChassis ensures that no value is present for AssignedChassis, not even an explicit nil
func (o *ChassisProfile) UnsetAssignedChassis() {
	o.AssignedChassis.Unset()
}

// GetAssociatedChassis returns the AssociatedChassis field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChassisProfile) GetAssociatedChassis() EquipmentChassisRelationship {
	if o == nil || IsNil(o.AssociatedChassis.Get()) {
		var ret EquipmentChassisRelationship
		return ret
	}
	return *o.AssociatedChassis.Get()
}

// GetAssociatedChassisOk returns a tuple with the AssociatedChassis field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChassisProfile) GetAssociatedChassisOk() (*EquipmentChassisRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssociatedChassis.Get(), o.AssociatedChassis.IsSet()
}

// HasAssociatedChassis returns a boolean if a field has been set.
func (o *ChassisProfile) HasAssociatedChassis() bool {
	if o != nil && o.AssociatedChassis.IsSet() {
		return true
	}

	return false
}

// SetAssociatedChassis gets a reference to the given NullableEquipmentChassisRelationship and assigns it to the AssociatedChassis field.
func (o *ChassisProfile) SetAssociatedChassis(v EquipmentChassisRelationship) {
	o.AssociatedChassis.Set(&v)
}

// SetAssociatedChassisNil sets the value for AssociatedChassis to be an explicit nil
func (o *ChassisProfile) SetAssociatedChassisNil() {
	o.AssociatedChassis.Set(nil)
}

// UnsetAssociatedChassis ensures that no value is present for AssociatedChassis, not even an explicit nil
func (o *ChassisProfile) UnsetAssociatedChassis() {
	o.AssociatedChassis.Unset()
}

// GetConfigChangeDetails returns the ConfigChangeDetails field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChassisProfile) GetConfigChangeDetails() []ChassisConfigChangeDetailRelationship {
	if o == nil {
		var ret []ChassisConfigChangeDetailRelationship
		return ret
	}
	return o.ConfigChangeDetails
}

// GetConfigChangeDetailsOk returns a tuple with the ConfigChangeDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChassisProfile) GetConfigChangeDetailsOk() ([]ChassisConfigChangeDetailRelationship, bool) {
	if o == nil || IsNil(o.ConfigChangeDetails) {
		return nil, false
	}
	return o.ConfigChangeDetails, true
}

// HasConfigChangeDetails returns a boolean if a field has been set.
func (o *ChassisProfile) HasConfigChangeDetails() bool {
	if o != nil && !IsNil(o.ConfigChangeDetails) {
		return true
	}

	return false
}

// SetConfigChangeDetails gets a reference to the given []ChassisConfigChangeDetailRelationship and assigns it to the ConfigChangeDetails field.
func (o *ChassisProfile) SetConfigChangeDetails(v []ChassisConfigChangeDetailRelationship) {
	o.ConfigChangeDetails = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChassisProfile) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || IsNil(o.Organization.Get()) {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization.Get()
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChassisProfile) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Organization.Get(), o.Organization.IsSet()
}

// HasOrganization returns a boolean if a field has been set.
func (o *ChassisProfile) HasOrganization() bool {
	if o != nil && o.Organization.IsSet() {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given NullableOrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *ChassisProfile) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization.Set(&v)
}

// SetOrganizationNil sets the value for Organization to be an explicit nil
func (o *ChassisProfile) SetOrganizationNil() {
	o.Organization.Set(nil)
}

// UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
func (o *ChassisProfile) UnsetOrganization() {
	o.Organization.Unset()
}

// GetRunningWorkflows returns the RunningWorkflows field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChassisProfile) GetRunningWorkflows() []WorkflowWorkflowInfoRelationship {
	if o == nil {
		var ret []WorkflowWorkflowInfoRelationship
		return ret
	}
	return o.RunningWorkflows
}

// GetRunningWorkflowsOk returns a tuple with the RunningWorkflows field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChassisProfile) GetRunningWorkflowsOk() ([]WorkflowWorkflowInfoRelationship, bool) {
	if o == nil || IsNil(o.RunningWorkflows) {
		return nil, false
	}
	return o.RunningWorkflows, true
}

// HasRunningWorkflows returns a boolean if a field has been set.
func (o *ChassisProfile) HasRunningWorkflows() bool {
	if o != nil && !IsNil(o.RunningWorkflows) {
		return true
	}

	return false
}

// SetRunningWorkflows gets a reference to the given []WorkflowWorkflowInfoRelationship and assigns it to the RunningWorkflows field.
func (o *ChassisProfile) SetRunningWorkflows(v []WorkflowWorkflowInfoRelationship) {
	o.RunningWorkflows = v
}

func (o ChassisProfile) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChassisProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedChassisBaseProfile, errChassisBaseProfile := json.Marshal(o.ChassisBaseProfile)
	if errChassisBaseProfile != nil {
		return map[string]interface{}{}, errChassisBaseProfile
	}
	errChassisBaseProfile = json.Unmarshal([]byte(serializedChassisBaseProfile), &toSerialize)
	if errChassisBaseProfile != nil {
		return map[string]interface{}{}, errChassisBaseProfile
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if o.ConfigChangeContext.IsSet() {
		toSerialize["ConfigChangeContext"] = o.ConfigChangeContext.Get()
	}
	if o.ConfigChanges.IsSet() {
		toSerialize["ConfigChanges"] = o.ConfigChanges.Get()
	}
	if !IsNil(o.UserLabel) {
		toSerialize["UserLabel"] = o.UserLabel
	}
	if o.AssignedChassis.IsSet() {
		toSerialize["AssignedChassis"] = o.AssignedChassis.Get()
	}
	if o.AssociatedChassis.IsSet() {
		toSerialize["AssociatedChassis"] = o.AssociatedChassis.Get()
	}
	if o.ConfigChangeDetails != nil {
		toSerialize["ConfigChangeDetails"] = o.ConfigChangeDetails
	}
	if o.Organization.IsSet() {
		toSerialize["Organization"] = o.Organization.Get()
	}
	if o.RunningWorkflows != nil {
		toSerialize["RunningWorkflows"] = o.RunningWorkflows
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ChassisProfile) UnmarshalJSON(data []byte) (err error) {
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
	type ChassisProfileWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType          string                            `json:"ObjectType"`
		ConfigChangeContext NullablePolicyConfigChangeContext `json:"ConfigChangeContext,omitempty"`
		ConfigChanges       NullablePolicyConfigChange        `json:"ConfigChanges,omitempty"`
		// User label assigned to the chassis profile.
		UserLabel         *string                              `json:"UserLabel,omitempty" validate:"regexp=^[ !#$%&\\\\(\\\\)\\\\*\\\\+,\\\\-\\\\.\\/:;\\\\?@\\\\[\\\\]_\\\\{\\\\|\\\\}~a-zA-Z0-9]*$"`
		AssignedChassis   NullableEquipmentChassisRelationship `json:"AssignedChassis,omitempty"`
		AssociatedChassis NullableEquipmentChassisRelationship `json:"AssociatedChassis,omitempty"`
		// An array of relationships to chassisConfigChangeDetail resources.
		ConfigChangeDetails []ChassisConfigChangeDetailRelationship      `json:"ConfigChangeDetails,omitempty"`
		Organization        NullableOrganizationOrganizationRelationship `json:"Organization,omitempty"`
		// An array of relationships to workflowWorkflowInfo resources.
		RunningWorkflows []WorkflowWorkflowInfoRelationship `json:"RunningWorkflows,omitempty"`
	}

	varChassisProfileWithoutEmbeddedStruct := ChassisProfileWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varChassisProfileWithoutEmbeddedStruct)
	if err == nil {
		varChassisProfile := _ChassisProfile{}
		varChassisProfile.ClassId = varChassisProfileWithoutEmbeddedStruct.ClassId
		varChassisProfile.ObjectType = varChassisProfileWithoutEmbeddedStruct.ObjectType
		varChassisProfile.ConfigChangeContext = varChassisProfileWithoutEmbeddedStruct.ConfigChangeContext
		varChassisProfile.ConfigChanges = varChassisProfileWithoutEmbeddedStruct.ConfigChanges
		varChassisProfile.UserLabel = varChassisProfileWithoutEmbeddedStruct.UserLabel
		varChassisProfile.AssignedChassis = varChassisProfileWithoutEmbeddedStruct.AssignedChassis
		varChassisProfile.AssociatedChassis = varChassisProfileWithoutEmbeddedStruct.AssociatedChassis
		varChassisProfile.ConfigChangeDetails = varChassisProfileWithoutEmbeddedStruct.ConfigChangeDetails
		varChassisProfile.Organization = varChassisProfileWithoutEmbeddedStruct.Organization
		varChassisProfile.RunningWorkflows = varChassisProfileWithoutEmbeddedStruct.RunningWorkflows
		*o = ChassisProfile(varChassisProfile)
	} else {
		return err
	}

	varChassisProfile := _ChassisProfile{}

	err = json.Unmarshal(data, &varChassisProfile)
	if err == nil {
		o.ChassisBaseProfile = varChassisProfile.ChassisBaseProfile
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ConfigChangeContext")
		delete(additionalProperties, "ConfigChanges")
		delete(additionalProperties, "UserLabel")
		delete(additionalProperties, "AssignedChassis")
		delete(additionalProperties, "AssociatedChassis")
		delete(additionalProperties, "ConfigChangeDetails")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "RunningWorkflows")

		// remove fields from embedded structs
		reflectChassisBaseProfile := reflect.ValueOf(o.ChassisBaseProfile)
		for i := 0; i < reflectChassisBaseProfile.Type().NumField(); i++ {
			t := reflectChassisBaseProfile.Type().Field(i)

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

type NullableChassisProfile struct {
	value *ChassisProfile
	isSet bool
}

func (v NullableChassisProfile) Get() *ChassisProfile {
	return v.value
}

func (v *NullableChassisProfile) Set(val *ChassisProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableChassisProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableChassisProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChassisProfile(val *ChassisProfile) *NullableChassisProfile {
	return &NullableChassisProfile{value: val, isSet: true}
}

func (v NullableChassisProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChassisProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
