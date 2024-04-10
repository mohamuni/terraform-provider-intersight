/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-15830
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// EquipmentHybridDriveSlotAllOf Definition of the list of properties defined in 'equipment.HybridDriveSlot', excluding properties defined in parent classes.
type EquipmentHybridDriveSlotAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The Configured Mode of the Hybrid Drive slot. * `` - Hybrid Drive slot  mode is not applicable. * `RAID` - Hybrid Drive slot mode is RAID. * `Direct` - Hybrid Drive slot mode is Direct.
	CurrentMode *string `json:"CurrentMode,omitempty"`
	// The Requested Mode for the Hybrid Drive slot. * `` - Hybrid Drive slot  mode is not applicable. * `RAID` - Hybrid Drive slot mode is RAID. * `Direct` - Hybrid Drive slot mode is Direct.
	RequestedMode        *string                              `json:"RequestedMode,omitempty"`
	ComputeBlade         *ComputeBladeRelationship            `json:"ComputeBlade,omitempty"`
	ComputeBoard         *ComputeBoardRelationship            `json:"ComputeBoard,omitempty"`
	ComputeRackUnit      *ComputeRackUnitRelationship         `json:"ComputeRackUnit,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EquipmentHybridDriveSlotAllOf EquipmentHybridDriveSlotAllOf

// NewEquipmentHybridDriveSlotAllOf instantiates a new EquipmentHybridDriveSlotAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquipmentHybridDriveSlotAllOf(classId string, objectType string) *EquipmentHybridDriveSlotAllOf {
	this := EquipmentHybridDriveSlotAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewEquipmentHybridDriveSlotAllOfWithDefaults instantiates a new EquipmentHybridDriveSlotAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquipmentHybridDriveSlotAllOfWithDefaults() *EquipmentHybridDriveSlotAllOf {
	this := EquipmentHybridDriveSlotAllOf{}
	var classId string = "equipment.HybridDriveSlot"
	this.ClassId = classId
	var objectType string = "equipment.HybridDriveSlot"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *EquipmentHybridDriveSlotAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *EquipmentHybridDriveSlotAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *EquipmentHybridDriveSlotAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *EquipmentHybridDriveSlotAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *EquipmentHybridDriveSlotAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *EquipmentHybridDriveSlotAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCurrentMode returns the CurrentMode field value if set, zero value otherwise.
func (o *EquipmentHybridDriveSlotAllOf) GetCurrentMode() string {
	if o == nil || o.CurrentMode == nil {
		var ret string
		return ret
	}
	return *o.CurrentMode
}

// GetCurrentModeOk returns a tuple with the CurrentMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentHybridDriveSlotAllOf) GetCurrentModeOk() (*string, bool) {
	if o == nil || o.CurrentMode == nil {
		return nil, false
	}
	return o.CurrentMode, true
}

// HasCurrentMode returns a boolean if a field has been set.
func (o *EquipmentHybridDriveSlotAllOf) HasCurrentMode() bool {
	if o != nil && o.CurrentMode != nil {
		return true
	}

	return false
}

// SetCurrentMode gets a reference to the given string and assigns it to the CurrentMode field.
func (o *EquipmentHybridDriveSlotAllOf) SetCurrentMode(v string) {
	o.CurrentMode = &v
}

// GetRequestedMode returns the RequestedMode field value if set, zero value otherwise.
func (o *EquipmentHybridDriveSlotAllOf) GetRequestedMode() string {
	if o == nil || o.RequestedMode == nil {
		var ret string
		return ret
	}
	return *o.RequestedMode
}

// GetRequestedModeOk returns a tuple with the RequestedMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentHybridDriveSlotAllOf) GetRequestedModeOk() (*string, bool) {
	if o == nil || o.RequestedMode == nil {
		return nil, false
	}
	return o.RequestedMode, true
}

// HasRequestedMode returns a boolean if a field has been set.
func (o *EquipmentHybridDriveSlotAllOf) HasRequestedMode() bool {
	if o != nil && o.RequestedMode != nil {
		return true
	}

	return false
}

// SetRequestedMode gets a reference to the given string and assigns it to the RequestedMode field.
func (o *EquipmentHybridDriveSlotAllOf) SetRequestedMode(v string) {
	o.RequestedMode = &v
}

// GetComputeBlade returns the ComputeBlade field value if set, zero value otherwise.
func (o *EquipmentHybridDriveSlotAllOf) GetComputeBlade() ComputeBladeRelationship {
	if o == nil || o.ComputeBlade == nil {
		var ret ComputeBladeRelationship
		return ret
	}
	return *o.ComputeBlade
}

// GetComputeBladeOk returns a tuple with the ComputeBlade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentHybridDriveSlotAllOf) GetComputeBladeOk() (*ComputeBladeRelationship, bool) {
	if o == nil || o.ComputeBlade == nil {
		return nil, false
	}
	return o.ComputeBlade, true
}

// HasComputeBlade returns a boolean if a field has been set.
func (o *EquipmentHybridDriveSlotAllOf) HasComputeBlade() bool {
	if o != nil && o.ComputeBlade != nil {
		return true
	}

	return false
}

// SetComputeBlade gets a reference to the given ComputeBladeRelationship and assigns it to the ComputeBlade field.
func (o *EquipmentHybridDriveSlotAllOf) SetComputeBlade(v ComputeBladeRelationship) {
	o.ComputeBlade = &v
}

// GetComputeBoard returns the ComputeBoard field value if set, zero value otherwise.
func (o *EquipmentHybridDriveSlotAllOf) GetComputeBoard() ComputeBoardRelationship {
	if o == nil || o.ComputeBoard == nil {
		var ret ComputeBoardRelationship
		return ret
	}
	return *o.ComputeBoard
}

// GetComputeBoardOk returns a tuple with the ComputeBoard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentHybridDriveSlotAllOf) GetComputeBoardOk() (*ComputeBoardRelationship, bool) {
	if o == nil || o.ComputeBoard == nil {
		return nil, false
	}
	return o.ComputeBoard, true
}

// HasComputeBoard returns a boolean if a field has been set.
func (o *EquipmentHybridDriveSlotAllOf) HasComputeBoard() bool {
	if o != nil && o.ComputeBoard != nil {
		return true
	}

	return false
}

// SetComputeBoard gets a reference to the given ComputeBoardRelationship and assigns it to the ComputeBoard field.
func (o *EquipmentHybridDriveSlotAllOf) SetComputeBoard(v ComputeBoardRelationship) {
	o.ComputeBoard = &v
}

// GetComputeRackUnit returns the ComputeRackUnit field value if set, zero value otherwise.
func (o *EquipmentHybridDriveSlotAllOf) GetComputeRackUnit() ComputeRackUnitRelationship {
	if o == nil || o.ComputeRackUnit == nil {
		var ret ComputeRackUnitRelationship
		return ret
	}
	return *o.ComputeRackUnit
}

// GetComputeRackUnitOk returns a tuple with the ComputeRackUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentHybridDriveSlotAllOf) GetComputeRackUnitOk() (*ComputeRackUnitRelationship, bool) {
	if o == nil || o.ComputeRackUnit == nil {
		return nil, false
	}
	return o.ComputeRackUnit, true
}

// HasComputeRackUnit returns a boolean if a field has been set.
func (o *EquipmentHybridDriveSlotAllOf) HasComputeRackUnit() bool {
	if o != nil && o.ComputeRackUnit != nil {
		return true
	}

	return false
}

// SetComputeRackUnit gets a reference to the given ComputeRackUnitRelationship and assigns it to the ComputeRackUnit field.
func (o *EquipmentHybridDriveSlotAllOf) SetComputeRackUnit(v ComputeRackUnitRelationship) {
	o.ComputeRackUnit = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *EquipmentHybridDriveSlotAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentHybridDriveSlotAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *EquipmentHybridDriveSlotAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *EquipmentHybridDriveSlotAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o EquipmentHybridDriveSlotAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.CurrentMode != nil {
		toSerialize["CurrentMode"] = o.CurrentMode
	}
	if o.RequestedMode != nil {
		toSerialize["RequestedMode"] = o.RequestedMode
	}
	if o.ComputeBlade != nil {
		toSerialize["ComputeBlade"] = o.ComputeBlade
	}
	if o.ComputeBoard != nil {
		toSerialize["ComputeBoard"] = o.ComputeBoard
	}
	if o.ComputeRackUnit != nil {
		toSerialize["ComputeRackUnit"] = o.ComputeRackUnit
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EquipmentHybridDriveSlotAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varEquipmentHybridDriveSlotAllOf := _EquipmentHybridDriveSlotAllOf{}

	if err = json.Unmarshal(bytes, &varEquipmentHybridDriveSlotAllOf); err == nil {
		*o = EquipmentHybridDriveSlotAllOf(varEquipmentHybridDriveSlotAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CurrentMode")
		delete(additionalProperties, "RequestedMode")
		delete(additionalProperties, "ComputeBlade")
		delete(additionalProperties, "ComputeBoard")
		delete(additionalProperties, "ComputeRackUnit")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEquipmentHybridDriveSlotAllOf struct {
	value *EquipmentHybridDriveSlotAllOf
	isSet bool
}

func (v NullableEquipmentHybridDriveSlotAllOf) Get() *EquipmentHybridDriveSlotAllOf {
	return v.value
}

func (v *NullableEquipmentHybridDriveSlotAllOf) Set(val *EquipmentHybridDriveSlotAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEquipmentHybridDriveSlotAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEquipmentHybridDriveSlotAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquipmentHybridDriveSlotAllOf(val *EquipmentHybridDriveSlotAllOf) *NullableEquipmentHybridDriveSlotAllOf {
	return &NullableEquipmentHybridDriveSlotAllOf{value: val, isSet: true}
}

func (v NullableEquipmentHybridDriveSlotAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquipmentHybridDriveSlotAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}