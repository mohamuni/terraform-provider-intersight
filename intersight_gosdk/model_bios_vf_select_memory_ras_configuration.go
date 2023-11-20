/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-14430
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// BiosVfSelectMemoryRasConfiguration Memory Reliability, availability and serviceability (RAS) configuration.
type BiosVfSelectMemoryRasConfiguration struct {
	InventoryBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Parent server serial number.
	Serial *string `json:"Serial,omitempty"`
	// The actual BIOS memory RAS configuration as reported by the platform BIOS. Possible values are \"maximum-performance\", \"mirror-mode-1lm\", \"adddc-sparing\", \"platform-default\", \"lockstep\", \"sparing\", \"mirroring\".
	VpSelectMemoryRasConfiguration *string                              `json:"VpSelectMemoryRasConfiguration,omitempty"`
	ComputeBlade                   *ComputeBladeRelationship            `json:"ComputeBlade,omitempty"`
	ComputeRackUnit                *ComputeRackUnitRelationship         `json:"ComputeRackUnit,omitempty"`
	InventoryDeviceInfo            *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
	RegisteredDevice               *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties           map[string]interface{}
}

type _BiosVfSelectMemoryRasConfiguration BiosVfSelectMemoryRasConfiguration

// NewBiosVfSelectMemoryRasConfiguration instantiates a new BiosVfSelectMemoryRasConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBiosVfSelectMemoryRasConfiguration(classId string, objectType string) *BiosVfSelectMemoryRasConfiguration {
	this := BiosVfSelectMemoryRasConfiguration{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewBiosVfSelectMemoryRasConfigurationWithDefaults instantiates a new BiosVfSelectMemoryRasConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBiosVfSelectMemoryRasConfigurationWithDefaults() *BiosVfSelectMemoryRasConfiguration {
	this := BiosVfSelectMemoryRasConfiguration{}
	var classId string = "bios.VfSelectMemoryRasConfiguration"
	this.ClassId = classId
	var objectType string = "bios.VfSelectMemoryRasConfiguration"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *BiosVfSelectMemoryRasConfiguration) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *BiosVfSelectMemoryRasConfiguration) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *BiosVfSelectMemoryRasConfiguration) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *BiosVfSelectMemoryRasConfiguration) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *BiosVfSelectMemoryRasConfiguration) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *BiosVfSelectMemoryRasConfiguration) SetObjectType(v string) {
	o.ObjectType = v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *BiosVfSelectMemoryRasConfiguration) GetSerial() string {
	if o == nil || o.Serial == nil {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BiosVfSelectMemoryRasConfiguration) GetSerialOk() (*string, bool) {
	if o == nil || o.Serial == nil {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *BiosVfSelectMemoryRasConfiguration) HasSerial() bool {
	if o != nil && o.Serial != nil {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *BiosVfSelectMemoryRasConfiguration) SetSerial(v string) {
	o.Serial = &v
}

// GetVpSelectMemoryRasConfiguration returns the VpSelectMemoryRasConfiguration field value if set, zero value otherwise.
func (o *BiosVfSelectMemoryRasConfiguration) GetVpSelectMemoryRasConfiguration() string {
	if o == nil || o.VpSelectMemoryRasConfiguration == nil {
		var ret string
		return ret
	}
	return *o.VpSelectMemoryRasConfiguration
}

// GetVpSelectMemoryRasConfigurationOk returns a tuple with the VpSelectMemoryRasConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BiosVfSelectMemoryRasConfiguration) GetVpSelectMemoryRasConfigurationOk() (*string, bool) {
	if o == nil || o.VpSelectMemoryRasConfiguration == nil {
		return nil, false
	}
	return o.VpSelectMemoryRasConfiguration, true
}

// HasVpSelectMemoryRasConfiguration returns a boolean if a field has been set.
func (o *BiosVfSelectMemoryRasConfiguration) HasVpSelectMemoryRasConfiguration() bool {
	if o != nil && o.VpSelectMemoryRasConfiguration != nil {
		return true
	}

	return false
}

// SetVpSelectMemoryRasConfiguration gets a reference to the given string and assigns it to the VpSelectMemoryRasConfiguration field.
func (o *BiosVfSelectMemoryRasConfiguration) SetVpSelectMemoryRasConfiguration(v string) {
	o.VpSelectMemoryRasConfiguration = &v
}

// GetComputeBlade returns the ComputeBlade field value if set, zero value otherwise.
func (o *BiosVfSelectMemoryRasConfiguration) GetComputeBlade() ComputeBladeRelationship {
	if o == nil || o.ComputeBlade == nil {
		var ret ComputeBladeRelationship
		return ret
	}
	return *o.ComputeBlade
}

// GetComputeBladeOk returns a tuple with the ComputeBlade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BiosVfSelectMemoryRasConfiguration) GetComputeBladeOk() (*ComputeBladeRelationship, bool) {
	if o == nil || o.ComputeBlade == nil {
		return nil, false
	}
	return o.ComputeBlade, true
}

// HasComputeBlade returns a boolean if a field has been set.
func (o *BiosVfSelectMemoryRasConfiguration) HasComputeBlade() bool {
	if o != nil && o.ComputeBlade != nil {
		return true
	}

	return false
}

// SetComputeBlade gets a reference to the given ComputeBladeRelationship and assigns it to the ComputeBlade field.
func (o *BiosVfSelectMemoryRasConfiguration) SetComputeBlade(v ComputeBladeRelationship) {
	o.ComputeBlade = &v
}

// GetComputeRackUnit returns the ComputeRackUnit field value if set, zero value otherwise.
func (o *BiosVfSelectMemoryRasConfiguration) GetComputeRackUnit() ComputeRackUnitRelationship {
	if o == nil || o.ComputeRackUnit == nil {
		var ret ComputeRackUnitRelationship
		return ret
	}
	return *o.ComputeRackUnit
}

// GetComputeRackUnitOk returns a tuple with the ComputeRackUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BiosVfSelectMemoryRasConfiguration) GetComputeRackUnitOk() (*ComputeRackUnitRelationship, bool) {
	if o == nil || o.ComputeRackUnit == nil {
		return nil, false
	}
	return o.ComputeRackUnit, true
}

// HasComputeRackUnit returns a boolean if a field has been set.
func (o *BiosVfSelectMemoryRasConfiguration) HasComputeRackUnit() bool {
	if o != nil && o.ComputeRackUnit != nil {
		return true
	}

	return false
}

// SetComputeRackUnit gets a reference to the given ComputeRackUnitRelationship and assigns it to the ComputeRackUnit field.
func (o *BiosVfSelectMemoryRasConfiguration) SetComputeRackUnit(v ComputeRackUnitRelationship) {
	o.ComputeRackUnit = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *BiosVfSelectMemoryRasConfiguration) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BiosVfSelectMemoryRasConfiguration) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *BiosVfSelectMemoryRasConfiguration) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *BiosVfSelectMemoryRasConfiguration) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *BiosVfSelectMemoryRasConfiguration) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BiosVfSelectMemoryRasConfiguration) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *BiosVfSelectMemoryRasConfiguration) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *BiosVfSelectMemoryRasConfiguration) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o BiosVfSelectMemoryRasConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedInventoryBase, errInventoryBase := json.Marshal(o.InventoryBase)
	if errInventoryBase != nil {
		return []byte{}, errInventoryBase
	}
	errInventoryBase = json.Unmarshal([]byte(serializedInventoryBase), &toSerialize)
	if errInventoryBase != nil {
		return []byte{}, errInventoryBase
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Serial != nil {
		toSerialize["Serial"] = o.Serial
	}
	if o.VpSelectMemoryRasConfiguration != nil {
		toSerialize["VpSelectMemoryRasConfiguration"] = o.VpSelectMemoryRasConfiguration
	}
	if o.ComputeBlade != nil {
		toSerialize["ComputeBlade"] = o.ComputeBlade
	}
	if o.ComputeRackUnit != nil {
		toSerialize["ComputeRackUnit"] = o.ComputeRackUnit
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BiosVfSelectMemoryRasConfiguration) UnmarshalJSON(bytes []byte) (err error) {
	type BiosVfSelectMemoryRasConfigurationWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Parent server serial number.
		Serial *string `json:"Serial,omitempty"`
		// The actual BIOS memory RAS configuration as reported by the platform BIOS. Possible values are \"maximum-performance\", \"mirror-mode-1lm\", \"adddc-sparing\", \"platform-default\", \"lockstep\", \"sparing\", \"mirroring\".
		VpSelectMemoryRasConfiguration *string                              `json:"VpSelectMemoryRasConfiguration,omitempty"`
		ComputeBlade                   *ComputeBladeRelationship            `json:"ComputeBlade,omitempty"`
		ComputeRackUnit                *ComputeRackUnitRelationship         `json:"ComputeRackUnit,omitempty"`
		InventoryDeviceInfo            *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
		RegisteredDevice               *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varBiosVfSelectMemoryRasConfigurationWithoutEmbeddedStruct := BiosVfSelectMemoryRasConfigurationWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varBiosVfSelectMemoryRasConfigurationWithoutEmbeddedStruct)
	if err == nil {
		varBiosVfSelectMemoryRasConfiguration := _BiosVfSelectMemoryRasConfiguration{}
		varBiosVfSelectMemoryRasConfiguration.ClassId = varBiosVfSelectMemoryRasConfigurationWithoutEmbeddedStruct.ClassId
		varBiosVfSelectMemoryRasConfiguration.ObjectType = varBiosVfSelectMemoryRasConfigurationWithoutEmbeddedStruct.ObjectType
		varBiosVfSelectMemoryRasConfiguration.Serial = varBiosVfSelectMemoryRasConfigurationWithoutEmbeddedStruct.Serial
		varBiosVfSelectMemoryRasConfiguration.VpSelectMemoryRasConfiguration = varBiosVfSelectMemoryRasConfigurationWithoutEmbeddedStruct.VpSelectMemoryRasConfiguration
		varBiosVfSelectMemoryRasConfiguration.ComputeBlade = varBiosVfSelectMemoryRasConfigurationWithoutEmbeddedStruct.ComputeBlade
		varBiosVfSelectMemoryRasConfiguration.ComputeRackUnit = varBiosVfSelectMemoryRasConfigurationWithoutEmbeddedStruct.ComputeRackUnit
		varBiosVfSelectMemoryRasConfiguration.InventoryDeviceInfo = varBiosVfSelectMemoryRasConfigurationWithoutEmbeddedStruct.InventoryDeviceInfo
		varBiosVfSelectMemoryRasConfiguration.RegisteredDevice = varBiosVfSelectMemoryRasConfigurationWithoutEmbeddedStruct.RegisteredDevice
		*o = BiosVfSelectMemoryRasConfiguration(varBiosVfSelectMemoryRasConfiguration)
	} else {
		return err
	}

	varBiosVfSelectMemoryRasConfiguration := _BiosVfSelectMemoryRasConfiguration{}

	err = json.Unmarshal(bytes, &varBiosVfSelectMemoryRasConfiguration)
	if err == nil {
		o.InventoryBase = varBiosVfSelectMemoryRasConfiguration.InventoryBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Serial")
		delete(additionalProperties, "VpSelectMemoryRasConfiguration")
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

type NullableBiosVfSelectMemoryRasConfiguration struct {
	value *BiosVfSelectMemoryRasConfiguration
	isSet bool
}

func (v NullableBiosVfSelectMemoryRasConfiguration) Get() *BiosVfSelectMemoryRasConfiguration {
	return v.value
}

func (v *NullableBiosVfSelectMemoryRasConfiguration) Set(val *BiosVfSelectMemoryRasConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableBiosVfSelectMemoryRasConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableBiosVfSelectMemoryRasConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBiosVfSelectMemoryRasConfiguration(val *BiosVfSelectMemoryRasConfiguration) *NullableBiosVfSelectMemoryRasConfiguration {
	return &NullableBiosVfSelectMemoryRasConfiguration{value: val, isSet: true}
}

func (v NullableBiosVfSelectMemoryRasConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBiosVfSelectMemoryRasConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
