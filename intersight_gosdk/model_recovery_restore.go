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

// checks if the RecoveryRestore type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecoveryRestore{}

// RecoveryRestore Triggers a restore operation on the target endpoint.
type RecoveryRestore struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                                         `json:"ObjectType"`
	ConfigParams         NullableRecoveryConfigParams                   `json:"ConfigParams,omitempty"`
	BackupInfo           NullableRecoveryAbstractBackupInfoRelationship `json:"BackupInfo,omitempty"`
	Device               NullableAssetDeviceRegistrationRelationship    `json:"Device,omitempty"`
	Organization         NullableOrganizationOrganizationRelationship   `json:"Organization,omitempty"`
	Workflow             NullableWorkflowWorkflowInfoRelationship       `json:"Workflow,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RecoveryRestore RecoveryRestore

// NewRecoveryRestore instantiates a new RecoveryRestore object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoveryRestore(classId string, objectType string) *RecoveryRestore {
	this := RecoveryRestore{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewRecoveryRestoreWithDefaults instantiates a new RecoveryRestore object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoveryRestoreWithDefaults() *RecoveryRestore {
	this := RecoveryRestore{}
	var classId string = "recovery.Restore"
	this.ClassId = classId
	var objectType string = "recovery.Restore"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *RecoveryRestore) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *RecoveryRestore) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *RecoveryRestore) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "recovery.Restore" of the ClassId field.
func (o *RecoveryRestore) GetDefaultClassId() interface{} {
	return "recovery.Restore"
}

// GetObjectType returns the ObjectType field value
func (o *RecoveryRestore) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *RecoveryRestore) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *RecoveryRestore) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "recovery.Restore" of the ObjectType field.
func (o *RecoveryRestore) GetDefaultObjectType() interface{} {
	return "recovery.Restore"
}

// GetConfigParams returns the ConfigParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoveryRestore) GetConfigParams() RecoveryConfigParams {
	if o == nil || IsNil(o.ConfigParams.Get()) {
		var ret RecoveryConfigParams
		return ret
	}
	return *o.ConfigParams.Get()
}

// GetConfigParamsOk returns a tuple with the ConfigParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoveryRestore) GetConfigParamsOk() (*RecoveryConfigParams, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigParams.Get(), o.ConfigParams.IsSet()
}

// HasConfigParams returns a boolean if a field has been set.
func (o *RecoveryRestore) HasConfigParams() bool {
	if o != nil && o.ConfigParams.IsSet() {
		return true
	}

	return false
}

// SetConfigParams gets a reference to the given NullableRecoveryConfigParams and assigns it to the ConfigParams field.
func (o *RecoveryRestore) SetConfigParams(v RecoveryConfigParams) {
	o.ConfigParams.Set(&v)
}

// SetConfigParamsNil sets the value for ConfigParams to be an explicit nil
func (o *RecoveryRestore) SetConfigParamsNil() {
	o.ConfigParams.Set(nil)
}

// UnsetConfigParams ensures that no value is present for ConfigParams, not even an explicit nil
func (o *RecoveryRestore) UnsetConfigParams() {
	o.ConfigParams.Unset()
}

// GetBackupInfo returns the BackupInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoveryRestore) GetBackupInfo() RecoveryAbstractBackupInfoRelationship {
	if o == nil || IsNil(o.BackupInfo.Get()) {
		var ret RecoveryAbstractBackupInfoRelationship
		return ret
	}
	return *o.BackupInfo.Get()
}

// GetBackupInfoOk returns a tuple with the BackupInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoveryRestore) GetBackupInfoOk() (*RecoveryAbstractBackupInfoRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.BackupInfo.Get(), o.BackupInfo.IsSet()
}

// HasBackupInfo returns a boolean if a field has been set.
func (o *RecoveryRestore) HasBackupInfo() bool {
	if o != nil && o.BackupInfo.IsSet() {
		return true
	}

	return false
}

// SetBackupInfo gets a reference to the given NullableRecoveryAbstractBackupInfoRelationship and assigns it to the BackupInfo field.
func (o *RecoveryRestore) SetBackupInfo(v RecoveryAbstractBackupInfoRelationship) {
	o.BackupInfo.Set(&v)
}

// SetBackupInfoNil sets the value for BackupInfo to be an explicit nil
func (o *RecoveryRestore) SetBackupInfoNil() {
	o.BackupInfo.Set(nil)
}

// UnsetBackupInfo ensures that no value is present for BackupInfo, not even an explicit nil
func (o *RecoveryRestore) UnsetBackupInfo() {
	o.BackupInfo.Unset()
}

// GetDevice returns the Device field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoveryRestore) GetDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.Device.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.Device.Get()
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoveryRestore) GetDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Device.Get(), o.Device.IsSet()
}

// HasDevice returns a boolean if a field has been set.
func (o *RecoveryRestore) HasDevice() bool {
	if o != nil && o.Device.IsSet() {
		return true
	}

	return false
}

// SetDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the Device field.
func (o *RecoveryRestore) SetDevice(v AssetDeviceRegistrationRelationship) {
	o.Device.Set(&v)
}

// SetDeviceNil sets the value for Device to be an explicit nil
func (o *RecoveryRestore) SetDeviceNil() {
	o.Device.Set(nil)
}

// UnsetDevice ensures that no value is present for Device, not even an explicit nil
func (o *RecoveryRestore) UnsetDevice() {
	o.Device.Unset()
}

// GetOrganization returns the Organization field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoveryRestore) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || IsNil(o.Organization.Get()) {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization.Get()
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoveryRestore) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Organization.Get(), o.Organization.IsSet()
}

// HasOrganization returns a boolean if a field has been set.
func (o *RecoveryRestore) HasOrganization() bool {
	if o != nil && o.Organization.IsSet() {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given NullableOrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *RecoveryRestore) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization.Set(&v)
}

// SetOrganizationNil sets the value for Organization to be an explicit nil
func (o *RecoveryRestore) SetOrganizationNil() {
	o.Organization.Set(nil)
}

// UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
func (o *RecoveryRestore) UnsetOrganization() {
	o.Organization.Unset()
}

// GetWorkflow returns the Workflow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoveryRestore) GetWorkflow() WorkflowWorkflowInfoRelationship {
	if o == nil || IsNil(o.Workflow.Get()) {
		var ret WorkflowWorkflowInfoRelationship
		return ret
	}
	return *o.Workflow.Get()
}

// GetWorkflowOk returns a tuple with the Workflow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoveryRestore) GetWorkflowOk() (*WorkflowWorkflowInfoRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Workflow.Get(), o.Workflow.IsSet()
}

// HasWorkflow returns a boolean if a field has been set.
func (o *RecoveryRestore) HasWorkflow() bool {
	if o != nil && o.Workflow.IsSet() {
		return true
	}

	return false
}

// SetWorkflow gets a reference to the given NullableWorkflowWorkflowInfoRelationship and assigns it to the Workflow field.
func (o *RecoveryRestore) SetWorkflow(v WorkflowWorkflowInfoRelationship) {
	o.Workflow.Set(&v)
}

// SetWorkflowNil sets the value for Workflow to be an explicit nil
func (o *RecoveryRestore) SetWorkflowNil() {
	o.Workflow.Set(nil)
}

// UnsetWorkflow ensures that no value is present for Workflow, not even an explicit nil
func (o *RecoveryRestore) UnsetWorkflow() {
	o.Workflow.Unset()
}

func (o RecoveryRestore) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecoveryRestore) ToMap() (map[string]interface{}, error) {
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
	if o.ConfigParams.IsSet() {
		toSerialize["ConfigParams"] = o.ConfigParams.Get()
	}
	if o.BackupInfo.IsSet() {
		toSerialize["BackupInfo"] = o.BackupInfo.Get()
	}
	if o.Device.IsSet() {
		toSerialize["Device"] = o.Device.Get()
	}
	if o.Organization.IsSet() {
		toSerialize["Organization"] = o.Organization.Get()
	}
	if o.Workflow.IsSet() {
		toSerialize["Workflow"] = o.Workflow.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RecoveryRestore) UnmarshalJSON(data []byte) (err error) {
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
	type RecoveryRestoreWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType   string                                         `json:"ObjectType"`
		ConfigParams NullableRecoveryConfigParams                   `json:"ConfigParams,omitempty"`
		BackupInfo   NullableRecoveryAbstractBackupInfoRelationship `json:"BackupInfo,omitempty"`
		Device       NullableAssetDeviceRegistrationRelationship    `json:"Device,omitempty"`
		Organization NullableOrganizationOrganizationRelationship   `json:"Organization,omitempty"`
		Workflow     NullableWorkflowWorkflowInfoRelationship       `json:"Workflow,omitempty"`
	}

	varRecoveryRestoreWithoutEmbeddedStruct := RecoveryRestoreWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varRecoveryRestoreWithoutEmbeddedStruct)
	if err == nil {
		varRecoveryRestore := _RecoveryRestore{}
		varRecoveryRestore.ClassId = varRecoveryRestoreWithoutEmbeddedStruct.ClassId
		varRecoveryRestore.ObjectType = varRecoveryRestoreWithoutEmbeddedStruct.ObjectType
		varRecoveryRestore.ConfigParams = varRecoveryRestoreWithoutEmbeddedStruct.ConfigParams
		varRecoveryRestore.BackupInfo = varRecoveryRestoreWithoutEmbeddedStruct.BackupInfo
		varRecoveryRestore.Device = varRecoveryRestoreWithoutEmbeddedStruct.Device
		varRecoveryRestore.Organization = varRecoveryRestoreWithoutEmbeddedStruct.Organization
		varRecoveryRestore.Workflow = varRecoveryRestoreWithoutEmbeddedStruct.Workflow
		*o = RecoveryRestore(varRecoveryRestore)
	} else {
		return err
	}

	varRecoveryRestore := _RecoveryRestore{}

	err = json.Unmarshal(data, &varRecoveryRestore)
	if err == nil {
		o.MoBaseMo = varRecoveryRestore.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ConfigParams")
		delete(additionalProperties, "BackupInfo")
		delete(additionalProperties, "Device")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "Workflow")

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

type NullableRecoveryRestore struct {
	value *RecoveryRestore
	isSet bool
}

func (v NullableRecoveryRestore) Get() *RecoveryRestore {
	return v.value
}

func (v *NullableRecoveryRestore) Set(val *RecoveryRestore) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoveryRestore) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoveryRestore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoveryRestore(val *RecoveryRestore) *NullableRecoveryRestore {
	return &NullableRecoveryRestore{value: val, isSet: true}
}

func (v NullableRecoveryRestore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoveryRestore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
