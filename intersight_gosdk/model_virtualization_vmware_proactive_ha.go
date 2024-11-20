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
	"time"
)

// checks if the VirtualizationVmwareProactiveHa type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VirtualizationVmwareProactiveHa{}

// VirtualizationVmwareProactiveHa Vmware vCenter has a functionality to support 'proactive HA' in clusters. Common attributes of 'HA Provider' and custom alarms added for a provider in a vCenter.
type VirtualizationVmwareProactiveHa struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Last recognized alarm time for a proactive HA alarm instance in a vCenter.
	LastAcknowledgedAlarmTime *time.Time `json:"LastAcknowledgedAlarmTime,omitempty"`
	// Time at which the last alarm was sent from cloud to the device connector.
	LastSentAlarmTime *time.Time `json:"LastSentAlarmTime,omitempty"`
	// An array of relationships to condAlarmDefinition resources.
	AlarmDefinitions     []CondAlarmDefinitionRelationship           `json:"AlarmDefinitions,omitempty"`
	RegisteredDevice     NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationVmwareProactiveHa VirtualizationVmwareProactiveHa

// NewVirtualizationVmwareProactiveHa instantiates a new VirtualizationVmwareProactiveHa object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationVmwareProactiveHa(classId string, objectType string) *VirtualizationVmwareProactiveHa {
	this := VirtualizationVmwareProactiveHa{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVirtualizationVmwareProactiveHaWithDefaults instantiates a new VirtualizationVmwareProactiveHa object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationVmwareProactiveHaWithDefaults() *VirtualizationVmwareProactiveHa {
	this := VirtualizationVmwareProactiveHa{}
	var classId string = "virtualization.VmwareProactiveHa"
	this.ClassId = classId
	var objectType string = "virtualization.VmwareProactiveHa"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationVmwareProactiveHa) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareProactiveHa) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationVmwareProactiveHa) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "virtualization.VmwareProactiveHa" of the ClassId field.
func (o *VirtualizationVmwareProactiveHa) GetDefaultClassId() interface{} {
	return "virtualization.VmwareProactiveHa"
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationVmwareProactiveHa) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareProactiveHa) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationVmwareProactiveHa) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "virtualization.VmwareProactiveHa" of the ObjectType field.
func (o *VirtualizationVmwareProactiveHa) GetDefaultObjectType() interface{} {
	return "virtualization.VmwareProactiveHa"
}

// GetLastAcknowledgedAlarmTime returns the LastAcknowledgedAlarmTime field value if set, zero value otherwise.
func (o *VirtualizationVmwareProactiveHa) GetLastAcknowledgedAlarmTime() time.Time {
	if o == nil || IsNil(o.LastAcknowledgedAlarmTime) {
		var ret time.Time
		return ret
	}
	return *o.LastAcknowledgedAlarmTime
}

// GetLastAcknowledgedAlarmTimeOk returns a tuple with the LastAcknowledgedAlarmTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareProactiveHa) GetLastAcknowledgedAlarmTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastAcknowledgedAlarmTime) {
		return nil, false
	}
	return o.LastAcknowledgedAlarmTime, true
}

// HasLastAcknowledgedAlarmTime returns a boolean if a field has been set.
func (o *VirtualizationVmwareProactiveHa) HasLastAcknowledgedAlarmTime() bool {
	if o != nil && !IsNil(o.LastAcknowledgedAlarmTime) {
		return true
	}

	return false
}

// SetLastAcknowledgedAlarmTime gets a reference to the given time.Time and assigns it to the LastAcknowledgedAlarmTime field.
func (o *VirtualizationVmwareProactiveHa) SetLastAcknowledgedAlarmTime(v time.Time) {
	o.LastAcknowledgedAlarmTime = &v
}

// GetLastSentAlarmTime returns the LastSentAlarmTime field value if set, zero value otherwise.
func (o *VirtualizationVmwareProactiveHa) GetLastSentAlarmTime() time.Time {
	if o == nil || IsNil(o.LastSentAlarmTime) {
		var ret time.Time
		return ret
	}
	return *o.LastSentAlarmTime
}

// GetLastSentAlarmTimeOk returns a tuple with the LastSentAlarmTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareProactiveHa) GetLastSentAlarmTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastSentAlarmTime) {
		return nil, false
	}
	return o.LastSentAlarmTime, true
}

// HasLastSentAlarmTime returns a boolean if a field has been set.
func (o *VirtualizationVmwareProactiveHa) HasLastSentAlarmTime() bool {
	if o != nil && !IsNil(o.LastSentAlarmTime) {
		return true
	}

	return false
}

// SetLastSentAlarmTime gets a reference to the given time.Time and assigns it to the LastSentAlarmTime field.
func (o *VirtualizationVmwareProactiveHa) SetLastSentAlarmTime(v time.Time) {
	o.LastSentAlarmTime = &v
}

// GetAlarmDefinitions returns the AlarmDefinitions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationVmwareProactiveHa) GetAlarmDefinitions() []CondAlarmDefinitionRelationship {
	if o == nil {
		var ret []CondAlarmDefinitionRelationship
		return ret
	}
	return o.AlarmDefinitions
}

// GetAlarmDefinitionsOk returns a tuple with the AlarmDefinitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationVmwareProactiveHa) GetAlarmDefinitionsOk() ([]CondAlarmDefinitionRelationship, bool) {
	if o == nil || IsNil(o.AlarmDefinitions) {
		return nil, false
	}
	return o.AlarmDefinitions, true
}

// HasAlarmDefinitions returns a boolean if a field has been set.
func (o *VirtualizationVmwareProactiveHa) HasAlarmDefinitions() bool {
	if o != nil && !IsNil(o.AlarmDefinitions) {
		return true
	}

	return false
}

// SetAlarmDefinitions gets a reference to the given []CondAlarmDefinitionRelationship and assigns it to the AlarmDefinitions field.
func (o *VirtualizationVmwareProactiveHa) SetAlarmDefinitions(v []CondAlarmDefinitionRelationship) {
	o.AlarmDefinitions = v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationVmwareProactiveHa) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationVmwareProactiveHa) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *VirtualizationVmwareProactiveHa) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *VirtualizationVmwareProactiveHa) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *VirtualizationVmwareProactiveHa) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *VirtualizationVmwareProactiveHa) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

func (o VirtualizationVmwareProactiveHa) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VirtualizationVmwareProactiveHa) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.LastAcknowledgedAlarmTime) {
		toSerialize["LastAcknowledgedAlarmTime"] = o.LastAcknowledgedAlarmTime
	}
	if !IsNil(o.LastSentAlarmTime) {
		toSerialize["LastSentAlarmTime"] = o.LastSentAlarmTime
	}
	if o.AlarmDefinitions != nil {
		toSerialize["AlarmDefinitions"] = o.AlarmDefinitions
	}
	if o.RegisteredDevice.IsSet() {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VirtualizationVmwareProactiveHa) UnmarshalJSON(data []byte) (err error) {
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
	type VirtualizationVmwareProactiveHaWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Last recognized alarm time for a proactive HA alarm instance in a vCenter.
		LastAcknowledgedAlarmTime *time.Time `json:"LastAcknowledgedAlarmTime,omitempty"`
		// Time at which the last alarm was sent from cloud to the device connector.
		LastSentAlarmTime *time.Time `json:"LastSentAlarmTime,omitempty"`
		// An array of relationships to condAlarmDefinition resources.
		AlarmDefinitions []CondAlarmDefinitionRelationship           `json:"AlarmDefinitions,omitempty"`
		RegisteredDevice NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varVirtualizationVmwareProactiveHaWithoutEmbeddedStruct := VirtualizationVmwareProactiveHaWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varVirtualizationVmwareProactiveHaWithoutEmbeddedStruct)
	if err == nil {
		varVirtualizationVmwareProactiveHa := _VirtualizationVmwareProactiveHa{}
		varVirtualizationVmwareProactiveHa.ClassId = varVirtualizationVmwareProactiveHaWithoutEmbeddedStruct.ClassId
		varVirtualizationVmwareProactiveHa.ObjectType = varVirtualizationVmwareProactiveHaWithoutEmbeddedStruct.ObjectType
		varVirtualizationVmwareProactiveHa.LastAcknowledgedAlarmTime = varVirtualizationVmwareProactiveHaWithoutEmbeddedStruct.LastAcknowledgedAlarmTime
		varVirtualizationVmwareProactiveHa.LastSentAlarmTime = varVirtualizationVmwareProactiveHaWithoutEmbeddedStruct.LastSentAlarmTime
		varVirtualizationVmwareProactiveHa.AlarmDefinitions = varVirtualizationVmwareProactiveHaWithoutEmbeddedStruct.AlarmDefinitions
		varVirtualizationVmwareProactiveHa.RegisteredDevice = varVirtualizationVmwareProactiveHaWithoutEmbeddedStruct.RegisteredDevice
		*o = VirtualizationVmwareProactiveHa(varVirtualizationVmwareProactiveHa)
	} else {
		return err
	}

	varVirtualizationVmwareProactiveHa := _VirtualizationVmwareProactiveHa{}

	err = json.Unmarshal(data, &varVirtualizationVmwareProactiveHa)
	if err == nil {
		o.MoBaseMo = varVirtualizationVmwareProactiveHa.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "LastAcknowledgedAlarmTime")
		delete(additionalProperties, "LastSentAlarmTime")
		delete(additionalProperties, "AlarmDefinitions")
		delete(additionalProperties, "RegisteredDevice")

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

type NullableVirtualizationVmwareProactiveHa struct {
	value *VirtualizationVmwareProactiveHa
	isSet bool
}

func (v NullableVirtualizationVmwareProactiveHa) Get() *VirtualizationVmwareProactiveHa {
	return v.value
}

func (v *NullableVirtualizationVmwareProactiveHa) Set(val *VirtualizationVmwareProactiveHa) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationVmwareProactiveHa) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationVmwareProactiveHa) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationVmwareProactiveHa(val *VirtualizationVmwareProactiveHa) *NullableVirtualizationVmwareProactiveHa {
	return &NullableVirtualizationVmwareProactiveHa{value: val, isSet: true}
}

func (v NullableVirtualizationVmwareProactiveHa) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationVmwareProactiveHa) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
