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

// checks if the BootDeviceBootMode type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BootDeviceBootMode{}

// BootDeviceBootMode Boot mode of the devices that BIOS uses to boot them.
type BootDeviceBootMode struct {
	InventoryBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The user desired BIOS boot mode as configured in the boot policy.
	ConfiguredBootMode   *string                                     `json:"ConfiguredBootMode,omitempty"`
	ComputeBlade         NullableComputeBladeRelationship            `json:"ComputeBlade,omitempty"`
	ComputeRackUnit      NullableComputeRackUnitRelationship         `json:"ComputeRackUnit,omitempty"`
	InventoryDeviceInfo  NullableInventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
	RegisteredDevice     NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BootDeviceBootMode BootDeviceBootMode

// NewBootDeviceBootMode instantiates a new BootDeviceBootMode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBootDeviceBootMode(classId string, objectType string) *BootDeviceBootMode {
	this := BootDeviceBootMode{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewBootDeviceBootModeWithDefaults instantiates a new BootDeviceBootMode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBootDeviceBootModeWithDefaults() *BootDeviceBootMode {
	this := BootDeviceBootMode{}
	var classId string = "boot.DeviceBootMode"
	this.ClassId = classId
	var objectType string = "boot.DeviceBootMode"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *BootDeviceBootMode) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *BootDeviceBootMode) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *BootDeviceBootMode) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "boot.DeviceBootMode" of the ClassId field.
func (o *BootDeviceBootMode) GetDefaultClassId() interface{} {
	return "boot.DeviceBootMode"
}

// GetObjectType returns the ObjectType field value
func (o *BootDeviceBootMode) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *BootDeviceBootMode) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *BootDeviceBootMode) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "boot.DeviceBootMode" of the ObjectType field.
func (o *BootDeviceBootMode) GetDefaultObjectType() interface{} {
	return "boot.DeviceBootMode"
}

// GetConfiguredBootMode returns the ConfiguredBootMode field value if set, zero value otherwise.
func (o *BootDeviceBootMode) GetConfiguredBootMode() string {
	if o == nil || IsNil(o.ConfiguredBootMode) {
		var ret string
		return ret
	}
	return *o.ConfiguredBootMode
}

// GetConfiguredBootModeOk returns a tuple with the ConfiguredBootMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootDeviceBootMode) GetConfiguredBootModeOk() (*string, bool) {
	if o == nil || IsNil(o.ConfiguredBootMode) {
		return nil, false
	}
	return o.ConfiguredBootMode, true
}

// HasConfiguredBootMode returns a boolean if a field has been set.
func (o *BootDeviceBootMode) HasConfiguredBootMode() bool {
	if o != nil && !IsNil(o.ConfiguredBootMode) {
		return true
	}

	return false
}

// SetConfiguredBootMode gets a reference to the given string and assigns it to the ConfiguredBootMode field.
func (o *BootDeviceBootMode) SetConfiguredBootMode(v string) {
	o.ConfiguredBootMode = &v
}

// GetComputeBlade returns the ComputeBlade field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BootDeviceBootMode) GetComputeBlade() ComputeBladeRelationship {
	if o == nil || IsNil(o.ComputeBlade.Get()) {
		var ret ComputeBladeRelationship
		return ret
	}
	return *o.ComputeBlade.Get()
}

// GetComputeBladeOk returns a tuple with the ComputeBlade field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BootDeviceBootMode) GetComputeBladeOk() (*ComputeBladeRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.ComputeBlade.Get(), o.ComputeBlade.IsSet()
}

// HasComputeBlade returns a boolean if a field has been set.
func (o *BootDeviceBootMode) HasComputeBlade() bool {
	if o != nil && o.ComputeBlade.IsSet() {
		return true
	}

	return false
}

// SetComputeBlade gets a reference to the given NullableComputeBladeRelationship and assigns it to the ComputeBlade field.
func (o *BootDeviceBootMode) SetComputeBlade(v ComputeBladeRelationship) {
	o.ComputeBlade.Set(&v)
}

// SetComputeBladeNil sets the value for ComputeBlade to be an explicit nil
func (o *BootDeviceBootMode) SetComputeBladeNil() {
	o.ComputeBlade.Set(nil)
}

// UnsetComputeBlade ensures that no value is present for ComputeBlade, not even an explicit nil
func (o *BootDeviceBootMode) UnsetComputeBlade() {
	o.ComputeBlade.Unset()
}

// GetComputeRackUnit returns the ComputeRackUnit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BootDeviceBootMode) GetComputeRackUnit() ComputeRackUnitRelationship {
	if o == nil || IsNil(o.ComputeRackUnit.Get()) {
		var ret ComputeRackUnitRelationship
		return ret
	}
	return *o.ComputeRackUnit.Get()
}

// GetComputeRackUnitOk returns a tuple with the ComputeRackUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BootDeviceBootMode) GetComputeRackUnitOk() (*ComputeRackUnitRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.ComputeRackUnit.Get(), o.ComputeRackUnit.IsSet()
}

// HasComputeRackUnit returns a boolean if a field has been set.
func (o *BootDeviceBootMode) HasComputeRackUnit() bool {
	if o != nil && o.ComputeRackUnit.IsSet() {
		return true
	}

	return false
}

// SetComputeRackUnit gets a reference to the given NullableComputeRackUnitRelationship and assigns it to the ComputeRackUnit field.
func (o *BootDeviceBootMode) SetComputeRackUnit(v ComputeRackUnitRelationship) {
	o.ComputeRackUnit.Set(&v)
}

// SetComputeRackUnitNil sets the value for ComputeRackUnit to be an explicit nil
func (o *BootDeviceBootMode) SetComputeRackUnitNil() {
	o.ComputeRackUnit.Set(nil)
}

// UnsetComputeRackUnit ensures that no value is present for ComputeRackUnit, not even an explicit nil
func (o *BootDeviceBootMode) UnsetComputeRackUnit() {
	o.ComputeRackUnit.Unset()
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BootDeviceBootMode) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || IsNil(o.InventoryDeviceInfo.Get()) {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo.Get()
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BootDeviceBootMode) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo.Get(), o.InventoryDeviceInfo.IsSet()
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *BootDeviceBootMode) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo.IsSet() {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given NullableInventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *BootDeviceBootMode) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo.Set(&v)
}

// SetInventoryDeviceInfoNil sets the value for InventoryDeviceInfo to be an explicit nil
func (o *BootDeviceBootMode) SetInventoryDeviceInfoNil() {
	o.InventoryDeviceInfo.Set(nil)
}

// UnsetInventoryDeviceInfo ensures that no value is present for InventoryDeviceInfo, not even an explicit nil
func (o *BootDeviceBootMode) UnsetInventoryDeviceInfo() {
	o.InventoryDeviceInfo.Unset()
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BootDeviceBootMode) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BootDeviceBootMode) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *BootDeviceBootMode) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *BootDeviceBootMode) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *BootDeviceBootMode) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *BootDeviceBootMode) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

func (o BootDeviceBootMode) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BootDeviceBootMode) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedInventoryBase, errInventoryBase := json.Marshal(o.InventoryBase)
	if errInventoryBase != nil {
		return map[string]interface{}{}, errInventoryBase
	}
	errInventoryBase = json.Unmarshal([]byte(serializedInventoryBase), &toSerialize)
	if errInventoryBase != nil {
		return map[string]interface{}{}, errInventoryBase
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.ConfiguredBootMode) {
		toSerialize["ConfiguredBootMode"] = o.ConfiguredBootMode
	}
	if o.ComputeBlade.IsSet() {
		toSerialize["ComputeBlade"] = o.ComputeBlade.Get()
	}
	if o.ComputeRackUnit.IsSet() {
		toSerialize["ComputeRackUnit"] = o.ComputeRackUnit.Get()
	}
	if o.InventoryDeviceInfo.IsSet() {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo.Get()
	}
	if o.RegisteredDevice.IsSet() {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BootDeviceBootMode) UnmarshalJSON(data []byte) (err error) {
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
	type BootDeviceBootModeWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The user desired BIOS boot mode as configured in the boot policy.
		ConfiguredBootMode  *string                                     `json:"ConfiguredBootMode,omitempty"`
		ComputeBlade        NullableComputeBladeRelationship            `json:"ComputeBlade,omitempty"`
		ComputeRackUnit     NullableComputeRackUnitRelationship         `json:"ComputeRackUnit,omitempty"`
		InventoryDeviceInfo NullableInventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
		RegisteredDevice    NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varBootDeviceBootModeWithoutEmbeddedStruct := BootDeviceBootModeWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varBootDeviceBootModeWithoutEmbeddedStruct)
	if err == nil {
		varBootDeviceBootMode := _BootDeviceBootMode{}
		varBootDeviceBootMode.ClassId = varBootDeviceBootModeWithoutEmbeddedStruct.ClassId
		varBootDeviceBootMode.ObjectType = varBootDeviceBootModeWithoutEmbeddedStruct.ObjectType
		varBootDeviceBootMode.ConfiguredBootMode = varBootDeviceBootModeWithoutEmbeddedStruct.ConfiguredBootMode
		varBootDeviceBootMode.ComputeBlade = varBootDeviceBootModeWithoutEmbeddedStruct.ComputeBlade
		varBootDeviceBootMode.ComputeRackUnit = varBootDeviceBootModeWithoutEmbeddedStruct.ComputeRackUnit
		varBootDeviceBootMode.InventoryDeviceInfo = varBootDeviceBootModeWithoutEmbeddedStruct.InventoryDeviceInfo
		varBootDeviceBootMode.RegisteredDevice = varBootDeviceBootModeWithoutEmbeddedStruct.RegisteredDevice
		*o = BootDeviceBootMode(varBootDeviceBootMode)
	} else {
		return err
	}

	varBootDeviceBootMode := _BootDeviceBootMode{}

	err = json.Unmarshal(data, &varBootDeviceBootMode)
	if err == nil {
		o.InventoryBase = varBootDeviceBootMode.InventoryBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ConfiguredBootMode")
		delete(additionalProperties, "ComputeBlade")
		delete(additionalProperties, "ComputeRackUnit")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "RegisteredDevice")

		// remove fields from embedded structs
		reflectInventoryBase := reflect.ValueOf(o.InventoryBase)
		for i := 0; i < reflectInventoryBase.Type().NumField(); i++ {
			t := reflectInventoryBase.Type().Field(i)

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

type NullableBootDeviceBootMode struct {
	value *BootDeviceBootMode
	isSet bool
}

func (v NullableBootDeviceBootMode) Get() *BootDeviceBootMode {
	return v.value
}

func (v *NullableBootDeviceBootMode) Set(val *BootDeviceBootMode) {
	v.value = val
	v.isSet = true
}

func (v NullableBootDeviceBootMode) IsSet() bool {
	return v.isSet
}

func (v *NullableBootDeviceBootMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBootDeviceBootMode(val *BootDeviceBootMode) *NullableBootDeviceBootMode {
	return &NullableBootDeviceBootMode{value: val, isSet: true}
}

func (v NullableBootDeviceBootMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBootDeviceBootMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
