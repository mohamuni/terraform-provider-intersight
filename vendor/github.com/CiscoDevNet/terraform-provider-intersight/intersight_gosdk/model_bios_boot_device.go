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

// checks if the BiosBootDevice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BiosBootDevice{}

// BiosBootDevice Actual boot devices of the system as enumerated by BIOS.
type BiosBootDevice struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Name of the Configured Boot Device.
	DeviceName *string `json:"DeviceName,omitempty"`
	// Type of the Configured Boot Device.
	DeviceType           *string                                     `json:"DeviceType,omitempty"`
	BiosSystemBootOrder  NullableBiosSystemBootOrderRelationship     `json:"BiosSystemBootOrder,omitempty"`
	RegisteredDevice     NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BiosBootDevice BiosBootDevice

// NewBiosBootDevice instantiates a new BiosBootDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBiosBootDevice(classId string, objectType string) *BiosBootDevice {
	this := BiosBootDevice{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewBiosBootDeviceWithDefaults instantiates a new BiosBootDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBiosBootDeviceWithDefaults() *BiosBootDevice {
	this := BiosBootDevice{}
	var classId string = "bios.BootDevice"
	this.ClassId = classId
	var objectType string = "bios.BootDevice"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *BiosBootDevice) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *BiosBootDevice) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *BiosBootDevice) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "bios.BootDevice" of the ClassId field.
func (o *BiosBootDevice) GetDefaultClassId() interface{} {
	return "bios.BootDevice"
}

// GetObjectType returns the ObjectType field value
func (o *BiosBootDevice) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *BiosBootDevice) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *BiosBootDevice) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "bios.BootDevice" of the ObjectType field.
func (o *BiosBootDevice) GetDefaultObjectType() interface{} {
	return "bios.BootDevice"
}

// GetDeviceName returns the DeviceName field value if set, zero value otherwise.
func (o *BiosBootDevice) GetDeviceName() string {
	if o == nil || IsNil(o.DeviceName) {
		var ret string
		return ret
	}
	return *o.DeviceName
}

// GetDeviceNameOk returns a tuple with the DeviceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BiosBootDevice) GetDeviceNameOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceName) {
		return nil, false
	}
	return o.DeviceName, true
}

// HasDeviceName returns a boolean if a field has been set.
func (o *BiosBootDevice) HasDeviceName() bool {
	if o != nil && !IsNil(o.DeviceName) {
		return true
	}

	return false
}

// SetDeviceName gets a reference to the given string and assigns it to the DeviceName field.
func (o *BiosBootDevice) SetDeviceName(v string) {
	o.DeviceName = &v
}

// GetDeviceType returns the DeviceType field value if set, zero value otherwise.
func (o *BiosBootDevice) GetDeviceType() string {
	if o == nil || IsNil(o.DeviceType) {
		var ret string
		return ret
	}
	return *o.DeviceType
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BiosBootDevice) GetDeviceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceType) {
		return nil, false
	}
	return o.DeviceType, true
}

// HasDeviceType returns a boolean if a field has been set.
func (o *BiosBootDevice) HasDeviceType() bool {
	if o != nil && !IsNil(o.DeviceType) {
		return true
	}

	return false
}

// SetDeviceType gets a reference to the given string and assigns it to the DeviceType field.
func (o *BiosBootDevice) SetDeviceType(v string) {
	o.DeviceType = &v
}

// GetBiosSystemBootOrder returns the BiosSystemBootOrder field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BiosBootDevice) GetBiosSystemBootOrder() BiosSystemBootOrderRelationship {
	if o == nil || IsNil(o.BiosSystemBootOrder.Get()) {
		var ret BiosSystemBootOrderRelationship
		return ret
	}
	return *o.BiosSystemBootOrder.Get()
}

// GetBiosSystemBootOrderOk returns a tuple with the BiosSystemBootOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BiosBootDevice) GetBiosSystemBootOrderOk() (*BiosSystemBootOrderRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.BiosSystemBootOrder.Get(), o.BiosSystemBootOrder.IsSet()
}

// HasBiosSystemBootOrder returns a boolean if a field has been set.
func (o *BiosBootDevice) HasBiosSystemBootOrder() bool {
	if o != nil && o.BiosSystemBootOrder.IsSet() {
		return true
	}

	return false
}

// SetBiosSystemBootOrder gets a reference to the given NullableBiosSystemBootOrderRelationship and assigns it to the BiosSystemBootOrder field.
func (o *BiosBootDevice) SetBiosSystemBootOrder(v BiosSystemBootOrderRelationship) {
	o.BiosSystemBootOrder.Set(&v)
}

// SetBiosSystemBootOrderNil sets the value for BiosSystemBootOrder to be an explicit nil
func (o *BiosBootDevice) SetBiosSystemBootOrderNil() {
	o.BiosSystemBootOrder.Set(nil)
}

// UnsetBiosSystemBootOrder ensures that no value is present for BiosSystemBootOrder, not even an explicit nil
func (o *BiosBootDevice) UnsetBiosSystemBootOrder() {
	o.BiosSystemBootOrder.Unset()
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BiosBootDevice) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BiosBootDevice) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *BiosBootDevice) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *BiosBootDevice) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *BiosBootDevice) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *BiosBootDevice) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

func (o BiosBootDevice) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BiosBootDevice) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.DeviceName) {
		toSerialize["DeviceName"] = o.DeviceName
	}
	if !IsNil(o.DeviceType) {
		toSerialize["DeviceType"] = o.DeviceType
	}
	if o.BiosSystemBootOrder.IsSet() {
		toSerialize["BiosSystemBootOrder"] = o.BiosSystemBootOrder.Get()
	}
	if o.RegisteredDevice.IsSet() {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BiosBootDevice) UnmarshalJSON(data []byte) (err error) {
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
	type BiosBootDeviceWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Name of the Configured Boot Device.
		DeviceName *string `json:"DeviceName,omitempty"`
		// Type of the Configured Boot Device.
		DeviceType          *string                                     `json:"DeviceType,omitempty"`
		BiosSystemBootOrder NullableBiosSystemBootOrderRelationship     `json:"BiosSystemBootOrder,omitempty"`
		RegisteredDevice    NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varBiosBootDeviceWithoutEmbeddedStruct := BiosBootDeviceWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varBiosBootDeviceWithoutEmbeddedStruct)
	if err == nil {
		varBiosBootDevice := _BiosBootDevice{}
		varBiosBootDevice.ClassId = varBiosBootDeviceWithoutEmbeddedStruct.ClassId
		varBiosBootDevice.ObjectType = varBiosBootDeviceWithoutEmbeddedStruct.ObjectType
		varBiosBootDevice.DeviceName = varBiosBootDeviceWithoutEmbeddedStruct.DeviceName
		varBiosBootDevice.DeviceType = varBiosBootDeviceWithoutEmbeddedStruct.DeviceType
		varBiosBootDevice.BiosSystemBootOrder = varBiosBootDeviceWithoutEmbeddedStruct.BiosSystemBootOrder
		varBiosBootDevice.RegisteredDevice = varBiosBootDeviceWithoutEmbeddedStruct.RegisteredDevice
		*o = BiosBootDevice(varBiosBootDevice)
	} else {
		return err
	}

	varBiosBootDevice := _BiosBootDevice{}

	err = json.Unmarshal(data, &varBiosBootDevice)
	if err == nil {
		o.MoBaseMo = varBiosBootDevice.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DeviceName")
		delete(additionalProperties, "DeviceType")
		delete(additionalProperties, "BiosSystemBootOrder")
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

type NullableBiosBootDevice struct {
	value *BiosBootDevice
	isSet bool
}

func (v NullableBiosBootDevice) Get() *BiosBootDevice {
	return v.value
}

func (v *NullableBiosBootDevice) Set(val *BiosBootDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableBiosBootDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableBiosBootDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBiosBootDevice(val *BiosBootDevice) *NullableBiosBootDevice {
	return &NullableBiosBootDevice{value: val, isSet: true}
}

func (v NullableBiosBootDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBiosBootDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
