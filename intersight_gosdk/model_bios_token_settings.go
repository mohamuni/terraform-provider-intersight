/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-18534
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

// checks if the BiosTokenSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BiosTokenSettings{}

// BiosTokenSettings Token settings for Memory Reliability, availability and serviceability (RAS) configuration.
type BiosTokenSettings struct {
	InventoryBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Value that represents if the BIOS configuration is active. Possible values are \"yes\" and \"no\".
	IsAssigned *string `json:"IsAssigned,omitempty"`
	// Parent server serial number.
	Serial *string `json:"Serial,omitempty"`
	// BIOS configuration name as found in dn. Possible values are \"ADDDC-Sparing\", \"Maximum-Performance\", \"Partial-Mirror-mode-1LM\", \"Mirror-Mode-1LM\", \"Mirroring\", \"Lockstep\", \"Sparing\", \"Platform-Default\".
	SettingsMoRn         *string                                     `json:"SettingsMoRn,omitempty"`
	ComputeBlade         NullableComputeBladeRelationship            `json:"ComputeBlade,omitempty"`
	ComputeRackUnit      NullableComputeRackUnitRelationship         `json:"ComputeRackUnit,omitempty"`
	InventoryDeviceInfo  NullableInventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
	RegisteredDevice     NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BiosTokenSettings BiosTokenSettings

// NewBiosTokenSettings instantiates a new BiosTokenSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBiosTokenSettings(classId string, objectType string) *BiosTokenSettings {
	this := BiosTokenSettings{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewBiosTokenSettingsWithDefaults instantiates a new BiosTokenSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBiosTokenSettingsWithDefaults() *BiosTokenSettings {
	this := BiosTokenSettings{}
	var classId string = "bios.TokenSettings"
	this.ClassId = classId
	var objectType string = "bios.TokenSettings"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *BiosTokenSettings) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *BiosTokenSettings) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *BiosTokenSettings) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "bios.TokenSettings" of the ClassId field.
func (o *BiosTokenSettings) GetDefaultClassId() interface{} {
	return "bios.TokenSettings"
}

// GetObjectType returns the ObjectType field value
func (o *BiosTokenSettings) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *BiosTokenSettings) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *BiosTokenSettings) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "bios.TokenSettings" of the ObjectType field.
func (o *BiosTokenSettings) GetDefaultObjectType() interface{} {
	return "bios.TokenSettings"
}

// GetIsAssigned returns the IsAssigned field value if set, zero value otherwise.
func (o *BiosTokenSettings) GetIsAssigned() string {
	if o == nil || IsNil(o.IsAssigned) {
		var ret string
		return ret
	}
	return *o.IsAssigned
}

// GetIsAssignedOk returns a tuple with the IsAssigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BiosTokenSettings) GetIsAssignedOk() (*string, bool) {
	if o == nil || IsNil(o.IsAssigned) {
		return nil, false
	}
	return o.IsAssigned, true
}

// HasIsAssigned returns a boolean if a field has been set.
func (o *BiosTokenSettings) HasIsAssigned() bool {
	if o != nil && !IsNil(o.IsAssigned) {
		return true
	}

	return false
}

// SetIsAssigned gets a reference to the given string and assigns it to the IsAssigned field.
func (o *BiosTokenSettings) SetIsAssigned(v string) {
	o.IsAssigned = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *BiosTokenSettings) GetSerial() string {
	if o == nil || IsNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BiosTokenSettings) GetSerialOk() (*string, bool) {
	if o == nil || IsNil(o.Serial) {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *BiosTokenSettings) HasSerial() bool {
	if o != nil && !IsNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *BiosTokenSettings) SetSerial(v string) {
	o.Serial = &v
}

// GetSettingsMoRn returns the SettingsMoRn field value if set, zero value otherwise.
func (o *BiosTokenSettings) GetSettingsMoRn() string {
	if o == nil || IsNil(o.SettingsMoRn) {
		var ret string
		return ret
	}
	return *o.SettingsMoRn
}

// GetSettingsMoRnOk returns a tuple with the SettingsMoRn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BiosTokenSettings) GetSettingsMoRnOk() (*string, bool) {
	if o == nil || IsNil(o.SettingsMoRn) {
		return nil, false
	}
	return o.SettingsMoRn, true
}

// HasSettingsMoRn returns a boolean if a field has been set.
func (o *BiosTokenSettings) HasSettingsMoRn() bool {
	if o != nil && !IsNil(o.SettingsMoRn) {
		return true
	}

	return false
}

// SetSettingsMoRn gets a reference to the given string and assigns it to the SettingsMoRn field.
func (o *BiosTokenSettings) SetSettingsMoRn(v string) {
	o.SettingsMoRn = &v
}

// GetComputeBlade returns the ComputeBlade field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BiosTokenSettings) GetComputeBlade() ComputeBladeRelationship {
	if o == nil || IsNil(o.ComputeBlade.Get()) {
		var ret ComputeBladeRelationship
		return ret
	}
	return *o.ComputeBlade.Get()
}

// GetComputeBladeOk returns a tuple with the ComputeBlade field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BiosTokenSettings) GetComputeBladeOk() (*ComputeBladeRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.ComputeBlade.Get(), o.ComputeBlade.IsSet()
}

// HasComputeBlade returns a boolean if a field has been set.
func (o *BiosTokenSettings) HasComputeBlade() bool {
	if o != nil && o.ComputeBlade.IsSet() {
		return true
	}

	return false
}

// SetComputeBlade gets a reference to the given NullableComputeBladeRelationship and assigns it to the ComputeBlade field.
func (o *BiosTokenSettings) SetComputeBlade(v ComputeBladeRelationship) {
	o.ComputeBlade.Set(&v)
}

// SetComputeBladeNil sets the value for ComputeBlade to be an explicit nil
func (o *BiosTokenSettings) SetComputeBladeNil() {
	o.ComputeBlade.Set(nil)
}

// UnsetComputeBlade ensures that no value is present for ComputeBlade, not even an explicit nil
func (o *BiosTokenSettings) UnsetComputeBlade() {
	o.ComputeBlade.Unset()
}

// GetComputeRackUnit returns the ComputeRackUnit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BiosTokenSettings) GetComputeRackUnit() ComputeRackUnitRelationship {
	if o == nil || IsNil(o.ComputeRackUnit.Get()) {
		var ret ComputeRackUnitRelationship
		return ret
	}
	return *o.ComputeRackUnit.Get()
}

// GetComputeRackUnitOk returns a tuple with the ComputeRackUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BiosTokenSettings) GetComputeRackUnitOk() (*ComputeRackUnitRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.ComputeRackUnit.Get(), o.ComputeRackUnit.IsSet()
}

// HasComputeRackUnit returns a boolean if a field has been set.
func (o *BiosTokenSettings) HasComputeRackUnit() bool {
	if o != nil && o.ComputeRackUnit.IsSet() {
		return true
	}

	return false
}

// SetComputeRackUnit gets a reference to the given NullableComputeRackUnitRelationship and assigns it to the ComputeRackUnit field.
func (o *BiosTokenSettings) SetComputeRackUnit(v ComputeRackUnitRelationship) {
	o.ComputeRackUnit.Set(&v)
}

// SetComputeRackUnitNil sets the value for ComputeRackUnit to be an explicit nil
func (o *BiosTokenSettings) SetComputeRackUnitNil() {
	o.ComputeRackUnit.Set(nil)
}

// UnsetComputeRackUnit ensures that no value is present for ComputeRackUnit, not even an explicit nil
func (o *BiosTokenSettings) UnsetComputeRackUnit() {
	o.ComputeRackUnit.Unset()
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BiosTokenSettings) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || IsNil(o.InventoryDeviceInfo.Get()) {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo.Get()
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BiosTokenSettings) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo.Get(), o.InventoryDeviceInfo.IsSet()
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *BiosTokenSettings) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo.IsSet() {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given NullableInventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *BiosTokenSettings) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo.Set(&v)
}

// SetInventoryDeviceInfoNil sets the value for InventoryDeviceInfo to be an explicit nil
func (o *BiosTokenSettings) SetInventoryDeviceInfoNil() {
	o.InventoryDeviceInfo.Set(nil)
}

// UnsetInventoryDeviceInfo ensures that no value is present for InventoryDeviceInfo, not even an explicit nil
func (o *BiosTokenSettings) UnsetInventoryDeviceInfo() {
	o.InventoryDeviceInfo.Unset()
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BiosTokenSettings) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BiosTokenSettings) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *BiosTokenSettings) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *BiosTokenSettings) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *BiosTokenSettings) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *BiosTokenSettings) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

func (o BiosTokenSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BiosTokenSettings) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.IsAssigned) {
		toSerialize["IsAssigned"] = o.IsAssigned
	}
	if !IsNil(o.Serial) {
		toSerialize["Serial"] = o.Serial
	}
	if !IsNil(o.SettingsMoRn) {
		toSerialize["SettingsMoRn"] = o.SettingsMoRn
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

func (o *BiosTokenSettings) UnmarshalJSON(data []byte) (err error) {
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
	type BiosTokenSettingsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Value that represents if the BIOS configuration is active. Possible values are \"yes\" and \"no\".
		IsAssigned *string `json:"IsAssigned,omitempty"`
		// Parent server serial number.
		Serial *string `json:"Serial,omitempty"`
		// BIOS configuration name as found in dn. Possible values are \"ADDDC-Sparing\", \"Maximum-Performance\", \"Partial-Mirror-mode-1LM\", \"Mirror-Mode-1LM\", \"Mirroring\", \"Lockstep\", \"Sparing\", \"Platform-Default\".
		SettingsMoRn        *string                                     `json:"SettingsMoRn,omitempty"`
		ComputeBlade        NullableComputeBladeRelationship            `json:"ComputeBlade,omitempty"`
		ComputeRackUnit     NullableComputeRackUnitRelationship         `json:"ComputeRackUnit,omitempty"`
		InventoryDeviceInfo NullableInventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
		RegisteredDevice    NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varBiosTokenSettingsWithoutEmbeddedStruct := BiosTokenSettingsWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varBiosTokenSettingsWithoutEmbeddedStruct)
	if err == nil {
		varBiosTokenSettings := _BiosTokenSettings{}
		varBiosTokenSettings.ClassId = varBiosTokenSettingsWithoutEmbeddedStruct.ClassId
		varBiosTokenSettings.ObjectType = varBiosTokenSettingsWithoutEmbeddedStruct.ObjectType
		varBiosTokenSettings.IsAssigned = varBiosTokenSettingsWithoutEmbeddedStruct.IsAssigned
		varBiosTokenSettings.Serial = varBiosTokenSettingsWithoutEmbeddedStruct.Serial
		varBiosTokenSettings.SettingsMoRn = varBiosTokenSettingsWithoutEmbeddedStruct.SettingsMoRn
		varBiosTokenSettings.ComputeBlade = varBiosTokenSettingsWithoutEmbeddedStruct.ComputeBlade
		varBiosTokenSettings.ComputeRackUnit = varBiosTokenSettingsWithoutEmbeddedStruct.ComputeRackUnit
		varBiosTokenSettings.InventoryDeviceInfo = varBiosTokenSettingsWithoutEmbeddedStruct.InventoryDeviceInfo
		varBiosTokenSettings.RegisteredDevice = varBiosTokenSettingsWithoutEmbeddedStruct.RegisteredDevice
		*o = BiosTokenSettings(varBiosTokenSettings)
	} else {
		return err
	}

	varBiosTokenSettings := _BiosTokenSettings{}

	err = json.Unmarshal(data, &varBiosTokenSettings)
	if err == nil {
		o.InventoryBase = varBiosTokenSettings.InventoryBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "IsAssigned")
		delete(additionalProperties, "Serial")
		delete(additionalProperties, "SettingsMoRn")
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

type NullableBiosTokenSettings struct {
	value *BiosTokenSettings
	isSet bool
}

func (v NullableBiosTokenSettings) Get() *BiosTokenSettings {
	return v.value
}

func (v *NullableBiosTokenSettings) Set(val *BiosTokenSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableBiosTokenSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableBiosTokenSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBiosTokenSettings(val *BiosTokenSettings) *NullableBiosTokenSettings {
	return &NullableBiosTokenSettings{value: val, isSet: true}
}

func (v NullableBiosTokenSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBiosTokenSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
