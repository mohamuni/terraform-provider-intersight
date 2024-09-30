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

// checks if the BootNvmeDevice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BootNvmeDevice{}

// BootNvmeDevice Nvme Boot Device configured on the server.
type BootNvmeDevice struct {
	BootConfiguredDevice
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                                      `json:"ObjectType"`
	ComputePhysical      NullableComputePhysicalRelationship         `json:"ComputePhysical,omitempty"`
	InventoryDeviceInfo  NullableInventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
	RegisteredDevice     NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BootNvmeDevice BootNvmeDevice

// NewBootNvmeDevice instantiates a new BootNvmeDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBootNvmeDevice(classId string, objectType string) *BootNvmeDevice {
	this := BootNvmeDevice{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewBootNvmeDeviceWithDefaults instantiates a new BootNvmeDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBootNvmeDeviceWithDefaults() *BootNvmeDevice {
	this := BootNvmeDevice{}
	var classId string = "boot.NvmeDevice"
	this.ClassId = classId
	var objectType string = "boot.NvmeDevice"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *BootNvmeDevice) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *BootNvmeDevice) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *BootNvmeDevice) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "boot.NvmeDevice" of the ClassId field.
func (o *BootNvmeDevice) GetDefaultClassId() interface{} {
	return "boot.NvmeDevice"
}

// GetObjectType returns the ObjectType field value
func (o *BootNvmeDevice) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *BootNvmeDevice) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *BootNvmeDevice) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "boot.NvmeDevice" of the ObjectType field.
func (o *BootNvmeDevice) GetDefaultObjectType() interface{} {
	return "boot.NvmeDevice"
}

// GetComputePhysical returns the ComputePhysical field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BootNvmeDevice) GetComputePhysical() ComputePhysicalRelationship {
	if o == nil || IsNil(o.ComputePhysical.Get()) {
		var ret ComputePhysicalRelationship
		return ret
	}
	return *o.ComputePhysical.Get()
}

// GetComputePhysicalOk returns a tuple with the ComputePhysical field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BootNvmeDevice) GetComputePhysicalOk() (*ComputePhysicalRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.ComputePhysical.Get(), o.ComputePhysical.IsSet()
}

// HasComputePhysical returns a boolean if a field has been set.
func (o *BootNvmeDevice) HasComputePhysical() bool {
	if o != nil && o.ComputePhysical.IsSet() {
		return true
	}

	return false
}

// SetComputePhysical gets a reference to the given NullableComputePhysicalRelationship and assigns it to the ComputePhysical field.
func (o *BootNvmeDevice) SetComputePhysical(v ComputePhysicalRelationship) {
	o.ComputePhysical.Set(&v)
}

// SetComputePhysicalNil sets the value for ComputePhysical to be an explicit nil
func (o *BootNvmeDevice) SetComputePhysicalNil() {
	o.ComputePhysical.Set(nil)
}

// UnsetComputePhysical ensures that no value is present for ComputePhysical, not even an explicit nil
func (o *BootNvmeDevice) UnsetComputePhysical() {
	o.ComputePhysical.Unset()
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BootNvmeDevice) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || IsNil(o.InventoryDeviceInfo.Get()) {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo.Get()
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BootNvmeDevice) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo.Get(), o.InventoryDeviceInfo.IsSet()
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *BootNvmeDevice) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo.IsSet() {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given NullableInventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *BootNvmeDevice) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo.Set(&v)
}

// SetInventoryDeviceInfoNil sets the value for InventoryDeviceInfo to be an explicit nil
func (o *BootNvmeDevice) SetInventoryDeviceInfoNil() {
	o.InventoryDeviceInfo.Set(nil)
}

// UnsetInventoryDeviceInfo ensures that no value is present for InventoryDeviceInfo, not even an explicit nil
func (o *BootNvmeDevice) UnsetInventoryDeviceInfo() {
	o.InventoryDeviceInfo.Unset()
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BootNvmeDevice) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BootNvmeDevice) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *BootNvmeDevice) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *BootNvmeDevice) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *BootNvmeDevice) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *BootNvmeDevice) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

func (o BootNvmeDevice) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BootNvmeDevice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedBootConfiguredDevice, errBootConfiguredDevice := json.Marshal(o.BootConfiguredDevice)
	if errBootConfiguredDevice != nil {
		return map[string]interface{}{}, errBootConfiguredDevice
	}
	errBootConfiguredDevice = json.Unmarshal([]byte(serializedBootConfiguredDevice), &toSerialize)
	if errBootConfiguredDevice != nil {
		return map[string]interface{}{}, errBootConfiguredDevice
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if o.ComputePhysical.IsSet() {
		toSerialize["ComputePhysical"] = o.ComputePhysical.Get()
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

func (o *BootNvmeDevice) UnmarshalJSON(data []byte) (err error) {
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
	type BootNvmeDeviceWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType          string                                      `json:"ObjectType"`
		ComputePhysical     NullableComputePhysicalRelationship         `json:"ComputePhysical,omitempty"`
		InventoryDeviceInfo NullableInventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
		RegisteredDevice    NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varBootNvmeDeviceWithoutEmbeddedStruct := BootNvmeDeviceWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varBootNvmeDeviceWithoutEmbeddedStruct)
	if err == nil {
		varBootNvmeDevice := _BootNvmeDevice{}
		varBootNvmeDevice.ClassId = varBootNvmeDeviceWithoutEmbeddedStruct.ClassId
		varBootNvmeDevice.ObjectType = varBootNvmeDeviceWithoutEmbeddedStruct.ObjectType
		varBootNvmeDevice.ComputePhysical = varBootNvmeDeviceWithoutEmbeddedStruct.ComputePhysical
		varBootNvmeDevice.InventoryDeviceInfo = varBootNvmeDeviceWithoutEmbeddedStruct.InventoryDeviceInfo
		varBootNvmeDevice.RegisteredDevice = varBootNvmeDeviceWithoutEmbeddedStruct.RegisteredDevice
		*o = BootNvmeDevice(varBootNvmeDevice)
	} else {
		return err
	}

	varBootNvmeDevice := _BootNvmeDevice{}

	err = json.Unmarshal(data, &varBootNvmeDevice)
	if err == nil {
		o.BootConfiguredDevice = varBootNvmeDevice.BootConfiguredDevice
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ComputePhysical")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "RegisteredDevice")

		// remove fields from embedded structs
		reflectBootConfiguredDevice := reflect.ValueOf(o.BootConfiguredDevice)
		for i := 0; i < reflectBootConfiguredDevice.Type().NumField(); i++ {
			t := reflectBootConfiguredDevice.Type().Field(i)

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

type NullableBootNvmeDevice struct {
	value *BootNvmeDevice
	isSet bool
}

func (v NullableBootNvmeDevice) Get() *BootNvmeDevice {
	return v.value
}

func (v *NullableBootNvmeDevice) Set(val *BootNvmeDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableBootNvmeDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableBootNvmeDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBootNvmeDevice(val *BootNvmeDevice) *NullableBootNvmeDevice {
	return &NullableBootNvmeDevice{value: val, isSet: true}
}

func (v NullableBootNvmeDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBootNvmeDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
