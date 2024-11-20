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

// checks if the RecoveryOnDemandBackup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecoveryOnDemandBackup{}

// RecoveryOnDemandBackup Handles requests for on demand backup for a given endpoint.
type RecoveryOnDemandBackup struct {
	RecoveryAbstractBackupConfig
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                                       `json:"ObjectType"`
	ConfigResult         NullableRecoveryConfigResultRelationship     `json:"ConfigResult,omitempty"`
	DeviceId             NullableAssetDeviceRegistrationRelationship  `json:"DeviceId,omitempty"`
	Organization         NullableOrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RecoveryOnDemandBackup RecoveryOnDemandBackup

// NewRecoveryOnDemandBackup instantiates a new RecoveryOnDemandBackup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoveryOnDemandBackup(classId string, objectType string) *RecoveryOnDemandBackup {
	this := RecoveryOnDemandBackup{}
	this.ClassId = classId
	this.ObjectType = objectType
	var locationType string = "Network Share"
	this.LocationType = &locationType
	var protocol string = "SCP"
	this.Protocol = &protocol
	var retentionCount int64 = 10
	this.RetentionCount = &retentionCount
	return &this
}

// NewRecoveryOnDemandBackupWithDefaults instantiates a new RecoveryOnDemandBackup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoveryOnDemandBackupWithDefaults() *RecoveryOnDemandBackup {
	this := RecoveryOnDemandBackup{}
	var classId string = "recovery.OnDemandBackup"
	this.ClassId = classId
	var objectType string = "recovery.OnDemandBackup"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *RecoveryOnDemandBackup) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *RecoveryOnDemandBackup) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *RecoveryOnDemandBackup) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "recovery.OnDemandBackup" of the ClassId field.
func (o *RecoveryOnDemandBackup) GetDefaultClassId() interface{} {
	return "recovery.OnDemandBackup"
}

// GetObjectType returns the ObjectType field value
func (o *RecoveryOnDemandBackup) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *RecoveryOnDemandBackup) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *RecoveryOnDemandBackup) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "recovery.OnDemandBackup" of the ObjectType field.
func (o *RecoveryOnDemandBackup) GetDefaultObjectType() interface{} {
	return "recovery.OnDemandBackup"
}

// GetConfigResult returns the ConfigResult field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoveryOnDemandBackup) GetConfigResult() RecoveryConfigResultRelationship {
	if o == nil || IsNil(o.ConfigResult.Get()) {
		var ret RecoveryConfigResultRelationship
		return ret
	}
	return *o.ConfigResult.Get()
}

// GetConfigResultOk returns a tuple with the ConfigResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoveryOnDemandBackup) GetConfigResultOk() (*RecoveryConfigResultRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigResult.Get(), o.ConfigResult.IsSet()
}

// HasConfigResult returns a boolean if a field has been set.
func (o *RecoveryOnDemandBackup) HasConfigResult() bool {
	if o != nil && o.ConfigResult.IsSet() {
		return true
	}

	return false
}

// SetConfigResult gets a reference to the given NullableRecoveryConfigResultRelationship and assigns it to the ConfigResult field.
func (o *RecoveryOnDemandBackup) SetConfigResult(v RecoveryConfigResultRelationship) {
	o.ConfigResult.Set(&v)
}

// SetConfigResultNil sets the value for ConfigResult to be an explicit nil
func (o *RecoveryOnDemandBackup) SetConfigResultNil() {
	o.ConfigResult.Set(nil)
}

// UnsetConfigResult ensures that no value is present for ConfigResult, not even an explicit nil
func (o *RecoveryOnDemandBackup) UnsetConfigResult() {
	o.ConfigResult.Unset()
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoveryOnDemandBackup) GetDeviceId() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.DeviceId.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.DeviceId.Get()
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoveryOnDemandBackup) GetDeviceIdOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeviceId.Get(), o.DeviceId.IsSet()
}

// HasDeviceId returns a boolean if a field has been set.
func (o *RecoveryOnDemandBackup) HasDeviceId() bool {
	if o != nil && o.DeviceId.IsSet() {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the DeviceId field.
func (o *RecoveryOnDemandBackup) SetDeviceId(v AssetDeviceRegistrationRelationship) {
	o.DeviceId.Set(&v)
}

// SetDeviceIdNil sets the value for DeviceId to be an explicit nil
func (o *RecoveryOnDemandBackup) SetDeviceIdNil() {
	o.DeviceId.Set(nil)
}

// UnsetDeviceId ensures that no value is present for DeviceId, not even an explicit nil
func (o *RecoveryOnDemandBackup) UnsetDeviceId() {
	o.DeviceId.Unset()
}

// GetOrganization returns the Organization field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoveryOnDemandBackup) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || IsNil(o.Organization.Get()) {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization.Get()
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoveryOnDemandBackup) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Organization.Get(), o.Organization.IsSet()
}

// HasOrganization returns a boolean if a field has been set.
func (o *RecoveryOnDemandBackup) HasOrganization() bool {
	if o != nil && o.Organization.IsSet() {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given NullableOrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *RecoveryOnDemandBackup) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization.Set(&v)
}

// SetOrganizationNil sets the value for Organization to be an explicit nil
func (o *RecoveryOnDemandBackup) SetOrganizationNil() {
	o.Organization.Set(nil)
}

// UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
func (o *RecoveryOnDemandBackup) UnsetOrganization() {
	o.Organization.Unset()
}

func (o RecoveryOnDemandBackup) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecoveryOnDemandBackup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedRecoveryAbstractBackupConfig, errRecoveryAbstractBackupConfig := json.Marshal(o.RecoveryAbstractBackupConfig)
	if errRecoveryAbstractBackupConfig != nil {
		return map[string]interface{}{}, errRecoveryAbstractBackupConfig
	}
	errRecoveryAbstractBackupConfig = json.Unmarshal([]byte(serializedRecoveryAbstractBackupConfig), &toSerialize)
	if errRecoveryAbstractBackupConfig != nil {
		return map[string]interface{}{}, errRecoveryAbstractBackupConfig
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if o.ConfigResult.IsSet() {
		toSerialize["ConfigResult"] = o.ConfigResult.Get()
	}
	if o.DeviceId.IsSet() {
		toSerialize["DeviceId"] = o.DeviceId.Get()
	}
	if o.Organization.IsSet() {
		toSerialize["Organization"] = o.Organization.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RecoveryOnDemandBackup) UnmarshalJSON(data []byte) (err error) {
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
	type RecoveryOnDemandBackupWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType   string                                       `json:"ObjectType"`
		ConfigResult NullableRecoveryConfigResultRelationship     `json:"ConfigResult,omitempty"`
		DeviceId     NullableAssetDeviceRegistrationRelationship  `json:"DeviceId,omitempty"`
		Organization NullableOrganizationOrganizationRelationship `json:"Organization,omitempty"`
	}

	varRecoveryOnDemandBackupWithoutEmbeddedStruct := RecoveryOnDemandBackupWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varRecoveryOnDemandBackupWithoutEmbeddedStruct)
	if err == nil {
		varRecoveryOnDemandBackup := _RecoveryOnDemandBackup{}
		varRecoveryOnDemandBackup.ClassId = varRecoveryOnDemandBackupWithoutEmbeddedStruct.ClassId
		varRecoveryOnDemandBackup.ObjectType = varRecoveryOnDemandBackupWithoutEmbeddedStruct.ObjectType
		varRecoveryOnDemandBackup.ConfigResult = varRecoveryOnDemandBackupWithoutEmbeddedStruct.ConfigResult
		varRecoveryOnDemandBackup.DeviceId = varRecoveryOnDemandBackupWithoutEmbeddedStruct.DeviceId
		varRecoveryOnDemandBackup.Organization = varRecoveryOnDemandBackupWithoutEmbeddedStruct.Organization
		*o = RecoveryOnDemandBackup(varRecoveryOnDemandBackup)
	} else {
		return err
	}

	varRecoveryOnDemandBackup := _RecoveryOnDemandBackup{}

	err = json.Unmarshal(data, &varRecoveryOnDemandBackup)
	if err == nil {
		o.RecoveryAbstractBackupConfig = varRecoveryOnDemandBackup.RecoveryAbstractBackupConfig
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ConfigResult")
		delete(additionalProperties, "DeviceId")
		delete(additionalProperties, "Organization")

		// remove fields from embedded structs
		reflectRecoveryAbstractBackupConfig := reflect.ValueOf(o.RecoveryAbstractBackupConfig)
		for i := 0; i < reflectRecoveryAbstractBackupConfig.Type().NumField(); i++ {
			t := reflectRecoveryAbstractBackupConfig.Type().Field(i)

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

type NullableRecoveryOnDemandBackup struct {
	value *RecoveryOnDemandBackup
	isSet bool
}

func (v NullableRecoveryOnDemandBackup) Get() *RecoveryOnDemandBackup {
	return v.value
}

func (v *NullableRecoveryOnDemandBackup) Set(val *RecoveryOnDemandBackup) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoveryOnDemandBackup) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoveryOnDemandBackup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoveryOnDemandBackup(val *RecoveryOnDemandBackup) *NullableRecoveryOnDemandBackup {
	return &NullableRecoveryOnDemandBackup{value: val, isSet: true}
}

func (v NullableRecoveryOnDemandBackup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoveryOnDemandBackup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
