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

// checks if the ComputeVmedia type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ComputeVmedia{}

// ComputeVmedia Inventory of Virtual Media configuration and images uploaded.
type ComputeVmedia struct {
	InventoryBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// State of the Virtual Media service on the server.
	Enabled *bool `json:"Enabled,omitempty"`
	// If enabled, allows encryption of all Virtual Media communications.
	Encryption *bool `json:"Encryption,omitempty"`
	// If enabled, the virtual drives appear on the boot selection menu after mapping the image and rebooting the host.
	LowPowerUsb         *bool                                   `json:"LowPowerUsb,omitempty"`
	ComputePhysicalUnit NullableComputePhysicalRelationship     `json:"ComputePhysicalUnit,omitempty"`
	InventoryDeviceInfo NullableInventoryDeviceInfoRelationship `json:"InventoryDeviceInfo,omitempty"`
	// An array of relationships to computeMapping resources.
	Mappings             []ComputeMappingRelationship                `json:"Mappings,omitempty"`
	RegisteredDevice     NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ComputeVmedia ComputeVmedia

// NewComputeVmedia instantiates a new ComputeVmedia object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputeVmedia(classId string, objectType string) *ComputeVmedia {
	this := ComputeVmedia{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewComputeVmediaWithDefaults instantiates a new ComputeVmedia object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputeVmediaWithDefaults() *ComputeVmedia {
	this := ComputeVmedia{}
	var classId string = "compute.Vmedia"
	this.ClassId = classId
	var objectType string = "compute.Vmedia"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ComputeVmedia) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ComputeVmedia) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ComputeVmedia) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "compute.Vmedia" of the ClassId field.
func (o *ComputeVmedia) GetDefaultClassId() interface{} {
	return "compute.Vmedia"
}

// GetObjectType returns the ObjectType field value
func (o *ComputeVmedia) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ComputeVmedia) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ComputeVmedia) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "compute.Vmedia" of the ObjectType field.
func (o *ComputeVmedia) GetDefaultObjectType() interface{} {
	return "compute.Vmedia"
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ComputeVmedia) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeVmedia) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ComputeVmedia) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ComputeVmedia) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetEncryption returns the Encryption field value if set, zero value otherwise.
func (o *ComputeVmedia) GetEncryption() bool {
	if o == nil || IsNil(o.Encryption) {
		var ret bool
		return ret
	}
	return *o.Encryption
}

// GetEncryptionOk returns a tuple with the Encryption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeVmedia) GetEncryptionOk() (*bool, bool) {
	if o == nil || IsNil(o.Encryption) {
		return nil, false
	}
	return o.Encryption, true
}

// HasEncryption returns a boolean if a field has been set.
func (o *ComputeVmedia) HasEncryption() bool {
	if o != nil && !IsNil(o.Encryption) {
		return true
	}

	return false
}

// SetEncryption gets a reference to the given bool and assigns it to the Encryption field.
func (o *ComputeVmedia) SetEncryption(v bool) {
	o.Encryption = &v
}

// GetLowPowerUsb returns the LowPowerUsb field value if set, zero value otherwise.
func (o *ComputeVmedia) GetLowPowerUsb() bool {
	if o == nil || IsNil(o.LowPowerUsb) {
		var ret bool
		return ret
	}
	return *o.LowPowerUsb
}

// GetLowPowerUsbOk returns a tuple with the LowPowerUsb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeVmedia) GetLowPowerUsbOk() (*bool, bool) {
	if o == nil || IsNil(o.LowPowerUsb) {
		return nil, false
	}
	return o.LowPowerUsb, true
}

// HasLowPowerUsb returns a boolean if a field has been set.
func (o *ComputeVmedia) HasLowPowerUsb() bool {
	if o != nil && !IsNil(o.LowPowerUsb) {
		return true
	}

	return false
}

// SetLowPowerUsb gets a reference to the given bool and assigns it to the LowPowerUsb field.
func (o *ComputeVmedia) SetLowPowerUsb(v bool) {
	o.LowPowerUsb = &v
}

// GetComputePhysicalUnit returns the ComputePhysicalUnit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComputeVmedia) GetComputePhysicalUnit() ComputePhysicalRelationship {
	if o == nil || IsNil(o.ComputePhysicalUnit.Get()) {
		var ret ComputePhysicalRelationship
		return ret
	}
	return *o.ComputePhysicalUnit.Get()
}

// GetComputePhysicalUnitOk returns a tuple with the ComputePhysicalUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComputeVmedia) GetComputePhysicalUnitOk() (*ComputePhysicalRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.ComputePhysicalUnit.Get(), o.ComputePhysicalUnit.IsSet()
}

// HasComputePhysicalUnit returns a boolean if a field has been set.
func (o *ComputeVmedia) HasComputePhysicalUnit() bool {
	if o != nil && o.ComputePhysicalUnit.IsSet() {
		return true
	}

	return false
}

// SetComputePhysicalUnit gets a reference to the given NullableComputePhysicalRelationship and assigns it to the ComputePhysicalUnit field.
func (o *ComputeVmedia) SetComputePhysicalUnit(v ComputePhysicalRelationship) {
	o.ComputePhysicalUnit.Set(&v)
}

// SetComputePhysicalUnitNil sets the value for ComputePhysicalUnit to be an explicit nil
func (o *ComputeVmedia) SetComputePhysicalUnitNil() {
	o.ComputePhysicalUnit.Set(nil)
}

// UnsetComputePhysicalUnit ensures that no value is present for ComputePhysicalUnit, not even an explicit nil
func (o *ComputeVmedia) UnsetComputePhysicalUnit() {
	o.ComputePhysicalUnit.Unset()
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComputeVmedia) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || IsNil(o.InventoryDeviceInfo.Get()) {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo.Get()
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComputeVmedia) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo.Get(), o.InventoryDeviceInfo.IsSet()
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *ComputeVmedia) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo.IsSet() {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given NullableInventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *ComputeVmedia) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo.Set(&v)
}

// SetInventoryDeviceInfoNil sets the value for InventoryDeviceInfo to be an explicit nil
func (o *ComputeVmedia) SetInventoryDeviceInfoNil() {
	o.InventoryDeviceInfo.Set(nil)
}

// UnsetInventoryDeviceInfo ensures that no value is present for InventoryDeviceInfo, not even an explicit nil
func (o *ComputeVmedia) UnsetInventoryDeviceInfo() {
	o.InventoryDeviceInfo.Unset()
}

// GetMappings returns the Mappings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComputeVmedia) GetMappings() []ComputeMappingRelationship {
	if o == nil {
		var ret []ComputeMappingRelationship
		return ret
	}
	return o.Mappings
}

// GetMappingsOk returns a tuple with the Mappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComputeVmedia) GetMappingsOk() ([]ComputeMappingRelationship, bool) {
	if o == nil || IsNil(o.Mappings) {
		return nil, false
	}
	return o.Mappings, true
}

// HasMappings returns a boolean if a field has been set.
func (o *ComputeVmedia) HasMappings() bool {
	if o != nil && !IsNil(o.Mappings) {
		return true
	}

	return false
}

// SetMappings gets a reference to the given []ComputeMappingRelationship and assigns it to the Mappings field.
func (o *ComputeVmedia) SetMappings(v []ComputeMappingRelationship) {
	o.Mappings = v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComputeVmedia) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComputeVmedia) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *ComputeVmedia) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *ComputeVmedia) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *ComputeVmedia) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *ComputeVmedia) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

func (o ComputeVmedia) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ComputeVmedia) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Enabled) {
		toSerialize["Enabled"] = o.Enabled
	}
	if !IsNil(o.Encryption) {
		toSerialize["Encryption"] = o.Encryption
	}
	if !IsNil(o.LowPowerUsb) {
		toSerialize["LowPowerUsb"] = o.LowPowerUsb
	}
	if o.ComputePhysicalUnit.IsSet() {
		toSerialize["ComputePhysicalUnit"] = o.ComputePhysicalUnit.Get()
	}
	if o.InventoryDeviceInfo.IsSet() {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo.Get()
	}
	if o.Mappings != nil {
		toSerialize["Mappings"] = o.Mappings
	}
	if o.RegisteredDevice.IsSet() {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ComputeVmedia) UnmarshalJSON(data []byte) (err error) {
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
	type ComputeVmediaWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// State of the Virtual Media service on the server.
		Enabled *bool `json:"Enabled,omitempty"`
		// If enabled, allows encryption of all Virtual Media communications.
		Encryption *bool `json:"Encryption,omitempty"`
		// If enabled, the virtual drives appear on the boot selection menu after mapping the image and rebooting the host.
		LowPowerUsb         *bool                                   `json:"LowPowerUsb,omitempty"`
		ComputePhysicalUnit NullableComputePhysicalRelationship     `json:"ComputePhysicalUnit,omitempty"`
		InventoryDeviceInfo NullableInventoryDeviceInfoRelationship `json:"InventoryDeviceInfo,omitempty"`
		// An array of relationships to computeMapping resources.
		Mappings         []ComputeMappingRelationship                `json:"Mappings,omitempty"`
		RegisteredDevice NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varComputeVmediaWithoutEmbeddedStruct := ComputeVmediaWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varComputeVmediaWithoutEmbeddedStruct)
	if err == nil {
		varComputeVmedia := _ComputeVmedia{}
		varComputeVmedia.ClassId = varComputeVmediaWithoutEmbeddedStruct.ClassId
		varComputeVmedia.ObjectType = varComputeVmediaWithoutEmbeddedStruct.ObjectType
		varComputeVmedia.Enabled = varComputeVmediaWithoutEmbeddedStruct.Enabled
		varComputeVmedia.Encryption = varComputeVmediaWithoutEmbeddedStruct.Encryption
		varComputeVmedia.LowPowerUsb = varComputeVmediaWithoutEmbeddedStruct.LowPowerUsb
		varComputeVmedia.ComputePhysicalUnit = varComputeVmediaWithoutEmbeddedStruct.ComputePhysicalUnit
		varComputeVmedia.InventoryDeviceInfo = varComputeVmediaWithoutEmbeddedStruct.InventoryDeviceInfo
		varComputeVmedia.Mappings = varComputeVmediaWithoutEmbeddedStruct.Mappings
		varComputeVmedia.RegisteredDevice = varComputeVmediaWithoutEmbeddedStruct.RegisteredDevice
		*o = ComputeVmedia(varComputeVmedia)
	} else {
		return err
	}

	varComputeVmedia := _ComputeVmedia{}

	err = json.Unmarshal(data, &varComputeVmedia)
	if err == nil {
		o.InventoryBase = varComputeVmedia.InventoryBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Enabled")
		delete(additionalProperties, "Encryption")
		delete(additionalProperties, "LowPowerUsb")
		delete(additionalProperties, "ComputePhysicalUnit")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "Mappings")
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

type NullableComputeVmedia struct {
	value *ComputeVmedia
	isSet bool
}

func (v NullableComputeVmedia) Get() *ComputeVmedia {
	return v.value
}

func (v *NullableComputeVmedia) Set(val *ComputeVmedia) {
	v.value = val
	v.isSet = true
}

func (v NullableComputeVmedia) IsSet() bool {
	return v.isSet
}

func (v *NullableComputeVmedia) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputeVmedia(val *ComputeVmedia) *NullableComputeVmedia {
	return &NullableComputeVmedia{value: val, isSet: true}
}

func (v NullableComputeVmedia) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputeVmedia) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
