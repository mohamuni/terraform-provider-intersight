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

// checks if the PciLink type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PciLink{}

// PciLink The PCI Switch Link connected to PCIe Switch.
type PciLink struct {
	EquipmentBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The name of the PCI device.
	Adapter *string `json:"Adapter,omitempty"`
	// The upstream link speed of the PCI device.
	LinkSpeed *string `json:"LinkSpeed,omitempty"`
	// The upstream link status of the PCI device.
	LinkStatus *string `json:"LinkStatus,omitempty"`
	// The upstream link width of the PCI device.
	LinkWidth *string `json:"LinkWidth,omitempty"`
	// The slot name of the PCI device.
	PciSlot *string `json:"PciSlot,omitempty"`
	// The health information of the PCI device.
	SlotStatus           *string                                     `json:"SlotStatus,omitempty"`
	InventoryDeviceInfo  NullableInventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
	PciSwitch            NullablePciSwitchRelationship               `json:"PciSwitch,omitempty"`
	RegisteredDevice     NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PciLink PciLink

// NewPciLink instantiates a new PciLink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPciLink(classId string, objectType string) *PciLink {
	this := PciLink{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewPciLinkWithDefaults instantiates a new PciLink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPciLinkWithDefaults() *PciLink {
	this := PciLink{}
	var classId string = "pci.Link"
	this.ClassId = classId
	var objectType string = "pci.Link"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *PciLink) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *PciLink) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *PciLink) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "pci.Link" of the ClassId field.
func (o *PciLink) GetDefaultClassId() interface{} {
	return "pci.Link"
}

// GetObjectType returns the ObjectType field value
func (o *PciLink) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *PciLink) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *PciLink) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "pci.Link" of the ObjectType field.
func (o *PciLink) GetDefaultObjectType() interface{} {
	return "pci.Link"
}

// GetAdapter returns the Adapter field value if set, zero value otherwise.
func (o *PciLink) GetAdapter() string {
	if o == nil || IsNil(o.Adapter) {
		var ret string
		return ret
	}
	return *o.Adapter
}

// GetAdapterOk returns a tuple with the Adapter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PciLink) GetAdapterOk() (*string, bool) {
	if o == nil || IsNil(o.Adapter) {
		return nil, false
	}
	return o.Adapter, true
}

// HasAdapter returns a boolean if a field has been set.
func (o *PciLink) HasAdapter() bool {
	if o != nil && !IsNil(o.Adapter) {
		return true
	}

	return false
}

// SetAdapter gets a reference to the given string and assigns it to the Adapter field.
func (o *PciLink) SetAdapter(v string) {
	o.Adapter = &v
}

// GetLinkSpeed returns the LinkSpeed field value if set, zero value otherwise.
func (o *PciLink) GetLinkSpeed() string {
	if o == nil || IsNil(o.LinkSpeed) {
		var ret string
		return ret
	}
	return *o.LinkSpeed
}

// GetLinkSpeedOk returns a tuple with the LinkSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PciLink) GetLinkSpeedOk() (*string, bool) {
	if o == nil || IsNil(o.LinkSpeed) {
		return nil, false
	}
	return o.LinkSpeed, true
}

// HasLinkSpeed returns a boolean if a field has been set.
func (o *PciLink) HasLinkSpeed() bool {
	if o != nil && !IsNil(o.LinkSpeed) {
		return true
	}

	return false
}

// SetLinkSpeed gets a reference to the given string and assigns it to the LinkSpeed field.
func (o *PciLink) SetLinkSpeed(v string) {
	o.LinkSpeed = &v
}

// GetLinkStatus returns the LinkStatus field value if set, zero value otherwise.
func (o *PciLink) GetLinkStatus() string {
	if o == nil || IsNil(o.LinkStatus) {
		var ret string
		return ret
	}
	return *o.LinkStatus
}

// GetLinkStatusOk returns a tuple with the LinkStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PciLink) GetLinkStatusOk() (*string, bool) {
	if o == nil || IsNil(o.LinkStatus) {
		return nil, false
	}
	return o.LinkStatus, true
}

// HasLinkStatus returns a boolean if a field has been set.
func (o *PciLink) HasLinkStatus() bool {
	if o != nil && !IsNil(o.LinkStatus) {
		return true
	}

	return false
}

// SetLinkStatus gets a reference to the given string and assigns it to the LinkStatus field.
func (o *PciLink) SetLinkStatus(v string) {
	o.LinkStatus = &v
}

// GetLinkWidth returns the LinkWidth field value if set, zero value otherwise.
func (o *PciLink) GetLinkWidth() string {
	if o == nil || IsNil(o.LinkWidth) {
		var ret string
		return ret
	}
	return *o.LinkWidth
}

// GetLinkWidthOk returns a tuple with the LinkWidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PciLink) GetLinkWidthOk() (*string, bool) {
	if o == nil || IsNil(o.LinkWidth) {
		return nil, false
	}
	return o.LinkWidth, true
}

// HasLinkWidth returns a boolean if a field has been set.
func (o *PciLink) HasLinkWidth() bool {
	if o != nil && !IsNil(o.LinkWidth) {
		return true
	}

	return false
}

// SetLinkWidth gets a reference to the given string and assigns it to the LinkWidth field.
func (o *PciLink) SetLinkWidth(v string) {
	o.LinkWidth = &v
}

// GetPciSlot returns the PciSlot field value if set, zero value otherwise.
func (o *PciLink) GetPciSlot() string {
	if o == nil || IsNil(o.PciSlot) {
		var ret string
		return ret
	}
	return *o.PciSlot
}

// GetPciSlotOk returns a tuple with the PciSlot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PciLink) GetPciSlotOk() (*string, bool) {
	if o == nil || IsNil(o.PciSlot) {
		return nil, false
	}
	return o.PciSlot, true
}

// HasPciSlot returns a boolean if a field has been set.
func (o *PciLink) HasPciSlot() bool {
	if o != nil && !IsNil(o.PciSlot) {
		return true
	}

	return false
}

// SetPciSlot gets a reference to the given string and assigns it to the PciSlot field.
func (o *PciLink) SetPciSlot(v string) {
	o.PciSlot = &v
}

// GetSlotStatus returns the SlotStatus field value if set, zero value otherwise.
func (o *PciLink) GetSlotStatus() string {
	if o == nil || IsNil(o.SlotStatus) {
		var ret string
		return ret
	}
	return *o.SlotStatus
}

// GetSlotStatusOk returns a tuple with the SlotStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PciLink) GetSlotStatusOk() (*string, bool) {
	if o == nil || IsNil(o.SlotStatus) {
		return nil, false
	}
	return o.SlotStatus, true
}

// HasSlotStatus returns a boolean if a field has been set.
func (o *PciLink) HasSlotStatus() bool {
	if o != nil && !IsNil(o.SlotStatus) {
		return true
	}

	return false
}

// SetSlotStatus gets a reference to the given string and assigns it to the SlotStatus field.
func (o *PciLink) SetSlotStatus(v string) {
	o.SlotStatus = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PciLink) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || IsNil(o.InventoryDeviceInfo.Get()) {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo.Get()
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PciLink) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo.Get(), o.InventoryDeviceInfo.IsSet()
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *PciLink) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo.IsSet() {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given NullableInventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *PciLink) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo.Set(&v)
}

// SetInventoryDeviceInfoNil sets the value for InventoryDeviceInfo to be an explicit nil
func (o *PciLink) SetInventoryDeviceInfoNil() {
	o.InventoryDeviceInfo.Set(nil)
}

// UnsetInventoryDeviceInfo ensures that no value is present for InventoryDeviceInfo, not even an explicit nil
func (o *PciLink) UnsetInventoryDeviceInfo() {
	o.InventoryDeviceInfo.Unset()
}

// GetPciSwitch returns the PciSwitch field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PciLink) GetPciSwitch() PciSwitchRelationship {
	if o == nil || IsNil(o.PciSwitch.Get()) {
		var ret PciSwitchRelationship
		return ret
	}
	return *o.PciSwitch.Get()
}

// GetPciSwitchOk returns a tuple with the PciSwitch field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PciLink) GetPciSwitchOk() (*PciSwitchRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.PciSwitch.Get(), o.PciSwitch.IsSet()
}

// HasPciSwitch returns a boolean if a field has been set.
func (o *PciLink) HasPciSwitch() bool {
	if o != nil && o.PciSwitch.IsSet() {
		return true
	}

	return false
}

// SetPciSwitch gets a reference to the given NullablePciSwitchRelationship and assigns it to the PciSwitch field.
func (o *PciLink) SetPciSwitch(v PciSwitchRelationship) {
	o.PciSwitch.Set(&v)
}

// SetPciSwitchNil sets the value for PciSwitch to be an explicit nil
func (o *PciLink) SetPciSwitchNil() {
	o.PciSwitch.Set(nil)
}

// UnsetPciSwitch ensures that no value is present for PciSwitch, not even an explicit nil
func (o *PciLink) UnsetPciSwitch() {
	o.PciSwitch.Unset()
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PciLink) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PciLink) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *PciLink) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *PciLink) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *PciLink) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *PciLink) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

func (o PciLink) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PciLink) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedEquipmentBase, errEquipmentBase := json.Marshal(o.EquipmentBase)
	if errEquipmentBase != nil {
		return map[string]interface{}{}, errEquipmentBase
	}
	errEquipmentBase = json.Unmarshal([]byte(serializedEquipmentBase), &toSerialize)
	if errEquipmentBase != nil {
		return map[string]interface{}{}, errEquipmentBase
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.Adapter) {
		toSerialize["Adapter"] = o.Adapter
	}
	if !IsNil(o.LinkSpeed) {
		toSerialize["LinkSpeed"] = o.LinkSpeed
	}
	if !IsNil(o.LinkStatus) {
		toSerialize["LinkStatus"] = o.LinkStatus
	}
	if !IsNil(o.LinkWidth) {
		toSerialize["LinkWidth"] = o.LinkWidth
	}
	if !IsNil(o.PciSlot) {
		toSerialize["PciSlot"] = o.PciSlot
	}
	if !IsNil(o.SlotStatus) {
		toSerialize["SlotStatus"] = o.SlotStatus
	}
	if o.InventoryDeviceInfo.IsSet() {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo.Get()
	}
	if o.PciSwitch.IsSet() {
		toSerialize["PciSwitch"] = o.PciSwitch.Get()
	}
	if o.RegisteredDevice.IsSet() {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PciLink) UnmarshalJSON(data []byte) (err error) {
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
	type PciLinkWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The name of the PCI device.
		Adapter *string `json:"Adapter,omitempty"`
		// The upstream link speed of the PCI device.
		LinkSpeed *string `json:"LinkSpeed,omitempty"`
		// The upstream link status of the PCI device.
		LinkStatus *string `json:"LinkStatus,omitempty"`
		// The upstream link width of the PCI device.
		LinkWidth *string `json:"LinkWidth,omitempty"`
		// The slot name of the PCI device.
		PciSlot *string `json:"PciSlot,omitempty"`
		// The health information of the PCI device.
		SlotStatus          *string                                     `json:"SlotStatus,omitempty"`
		InventoryDeviceInfo NullableInventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
		PciSwitch           NullablePciSwitchRelationship               `json:"PciSwitch,omitempty"`
		RegisteredDevice    NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varPciLinkWithoutEmbeddedStruct := PciLinkWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varPciLinkWithoutEmbeddedStruct)
	if err == nil {
		varPciLink := _PciLink{}
		varPciLink.ClassId = varPciLinkWithoutEmbeddedStruct.ClassId
		varPciLink.ObjectType = varPciLinkWithoutEmbeddedStruct.ObjectType
		varPciLink.Adapter = varPciLinkWithoutEmbeddedStruct.Adapter
		varPciLink.LinkSpeed = varPciLinkWithoutEmbeddedStruct.LinkSpeed
		varPciLink.LinkStatus = varPciLinkWithoutEmbeddedStruct.LinkStatus
		varPciLink.LinkWidth = varPciLinkWithoutEmbeddedStruct.LinkWidth
		varPciLink.PciSlot = varPciLinkWithoutEmbeddedStruct.PciSlot
		varPciLink.SlotStatus = varPciLinkWithoutEmbeddedStruct.SlotStatus
		varPciLink.InventoryDeviceInfo = varPciLinkWithoutEmbeddedStruct.InventoryDeviceInfo
		varPciLink.PciSwitch = varPciLinkWithoutEmbeddedStruct.PciSwitch
		varPciLink.RegisteredDevice = varPciLinkWithoutEmbeddedStruct.RegisteredDevice
		*o = PciLink(varPciLink)
	} else {
		return err
	}

	varPciLink := _PciLink{}

	err = json.Unmarshal(data, &varPciLink)
	if err == nil {
		o.EquipmentBase = varPciLink.EquipmentBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Adapter")
		delete(additionalProperties, "LinkSpeed")
		delete(additionalProperties, "LinkStatus")
		delete(additionalProperties, "LinkWidth")
		delete(additionalProperties, "PciSlot")
		delete(additionalProperties, "SlotStatus")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "PciSwitch")
		delete(additionalProperties, "RegisteredDevice")

		// remove fields from embedded structs
		reflectEquipmentBase := reflect.ValueOf(o.EquipmentBase)
		for i := 0; i < reflectEquipmentBase.Type().NumField(); i++ {
			t := reflectEquipmentBase.Type().Field(i)

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

type NullablePciLink struct {
	value *PciLink
	isSet bool
}

func (v NullablePciLink) Get() *PciLink {
	return v.value
}

func (v *NullablePciLink) Set(val *PciLink) {
	v.value = val
	v.isSet = true
}

func (v NullablePciLink) IsSet() bool {
	return v.isSet
}

func (v *NullablePciLink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePciLink(val *PciLink) *NullablePciLink {
	return &NullablePciLink{value: val, isSet: true}
}

func (v NullablePciLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePciLink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
