/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-03-27T10:08:12Z.
 *
 * API version: 1.0.9-4136
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// EquipmentPsuControlAllOf Definition of the list of properties defined in 'equipment.PsuControl', excluding properties defined in parent classes.
type EquipmentPsuControlAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// This field identifies the cluster state of the psu redundancy.
	ClusterState *string `json:"ClusterState,omitempty"`
	// This field identifies the input power state of the psus.
	InputPowerState *string `json:"InputPowerState,omitempty"`
	// This field identifies the name of psu control object.
	Name *string `json:"Name,omitempty"`
	// This field identifies the operational qualifier for the psu redundancy.
	OperQualifier *string  `json:"OperQualifier,omitempty"`
	OperReason    []string `json:"OperReason,omitempty"`
	// This field identifies the operational state of the psu redundancy.
	OperState *string `json:"OperState,omitempty"`
	// This field identifies the output power state of the psus.
	OutputPowerState *string `json:"OutputPowerState,omitempty"`
	// This field identifies the redundancy state of the psus.
	Redundancy           *string                              `json:"Redundancy,omitempty"`
	EquipmentChassis     *EquipmentChassisRelationship        `json:"EquipmentChassis,omitempty"`
	InventoryDeviceInfo  *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EquipmentPsuControlAllOf EquipmentPsuControlAllOf

// NewEquipmentPsuControlAllOf instantiates a new EquipmentPsuControlAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquipmentPsuControlAllOf(classId string, objectType string) *EquipmentPsuControlAllOf {
	this := EquipmentPsuControlAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewEquipmentPsuControlAllOfWithDefaults instantiates a new EquipmentPsuControlAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquipmentPsuControlAllOfWithDefaults() *EquipmentPsuControlAllOf {
	this := EquipmentPsuControlAllOf{}
	var classId string = "equipment.PsuControl"
	this.ClassId = classId
	var objectType string = "equipment.PsuControl"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *EquipmentPsuControlAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *EquipmentPsuControlAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *EquipmentPsuControlAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *EquipmentPsuControlAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *EquipmentPsuControlAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *EquipmentPsuControlAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetClusterState returns the ClusterState field value if set, zero value otherwise.
func (o *EquipmentPsuControlAllOf) GetClusterState() string {
	if o == nil || o.ClusterState == nil {
		var ret string
		return ret
	}
	return *o.ClusterState
}

// GetClusterStateOk returns a tuple with the ClusterState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentPsuControlAllOf) GetClusterStateOk() (*string, bool) {
	if o == nil || o.ClusterState == nil {
		return nil, false
	}
	return o.ClusterState, true
}

// HasClusterState returns a boolean if a field has been set.
func (o *EquipmentPsuControlAllOf) HasClusterState() bool {
	if o != nil && o.ClusterState != nil {
		return true
	}

	return false
}

// SetClusterState gets a reference to the given string and assigns it to the ClusterState field.
func (o *EquipmentPsuControlAllOf) SetClusterState(v string) {
	o.ClusterState = &v
}

// GetInputPowerState returns the InputPowerState field value if set, zero value otherwise.
func (o *EquipmentPsuControlAllOf) GetInputPowerState() string {
	if o == nil || o.InputPowerState == nil {
		var ret string
		return ret
	}
	return *o.InputPowerState
}

// GetInputPowerStateOk returns a tuple with the InputPowerState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentPsuControlAllOf) GetInputPowerStateOk() (*string, bool) {
	if o == nil || o.InputPowerState == nil {
		return nil, false
	}
	return o.InputPowerState, true
}

// HasInputPowerState returns a boolean if a field has been set.
func (o *EquipmentPsuControlAllOf) HasInputPowerState() bool {
	if o != nil && o.InputPowerState != nil {
		return true
	}

	return false
}

// SetInputPowerState gets a reference to the given string and assigns it to the InputPowerState field.
func (o *EquipmentPsuControlAllOf) SetInputPowerState(v string) {
	o.InputPowerState = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EquipmentPsuControlAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentPsuControlAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EquipmentPsuControlAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EquipmentPsuControlAllOf) SetName(v string) {
	o.Name = &v
}

// GetOperQualifier returns the OperQualifier field value if set, zero value otherwise.
func (o *EquipmentPsuControlAllOf) GetOperQualifier() string {
	if o == nil || o.OperQualifier == nil {
		var ret string
		return ret
	}
	return *o.OperQualifier
}

// GetOperQualifierOk returns a tuple with the OperQualifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentPsuControlAllOf) GetOperQualifierOk() (*string, bool) {
	if o == nil || o.OperQualifier == nil {
		return nil, false
	}
	return o.OperQualifier, true
}

// HasOperQualifier returns a boolean if a field has been set.
func (o *EquipmentPsuControlAllOf) HasOperQualifier() bool {
	if o != nil && o.OperQualifier != nil {
		return true
	}

	return false
}

// SetOperQualifier gets a reference to the given string and assigns it to the OperQualifier field.
func (o *EquipmentPsuControlAllOf) SetOperQualifier(v string) {
	o.OperQualifier = &v
}

// GetOperReason returns the OperReason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquipmentPsuControlAllOf) GetOperReason() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.OperReason
}

// GetOperReasonOk returns a tuple with the OperReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquipmentPsuControlAllOf) GetOperReasonOk() (*[]string, bool) {
	if o == nil || o.OperReason == nil {
		return nil, false
	}
	return &o.OperReason, true
}

// HasOperReason returns a boolean if a field has been set.
func (o *EquipmentPsuControlAllOf) HasOperReason() bool {
	if o != nil && o.OperReason != nil {
		return true
	}

	return false
}

// SetOperReason gets a reference to the given []string and assigns it to the OperReason field.
func (o *EquipmentPsuControlAllOf) SetOperReason(v []string) {
	o.OperReason = v
}

// GetOperState returns the OperState field value if set, zero value otherwise.
func (o *EquipmentPsuControlAllOf) GetOperState() string {
	if o == nil || o.OperState == nil {
		var ret string
		return ret
	}
	return *o.OperState
}

// GetOperStateOk returns a tuple with the OperState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentPsuControlAllOf) GetOperStateOk() (*string, bool) {
	if o == nil || o.OperState == nil {
		return nil, false
	}
	return o.OperState, true
}

// HasOperState returns a boolean if a field has been set.
func (o *EquipmentPsuControlAllOf) HasOperState() bool {
	if o != nil && o.OperState != nil {
		return true
	}

	return false
}

// SetOperState gets a reference to the given string and assigns it to the OperState field.
func (o *EquipmentPsuControlAllOf) SetOperState(v string) {
	o.OperState = &v
}

// GetOutputPowerState returns the OutputPowerState field value if set, zero value otherwise.
func (o *EquipmentPsuControlAllOf) GetOutputPowerState() string {
	if o == nil || o.OutputPowerState == nil {
		var ret string
		return ret
	}
	return *o.OutputPowerState
}

// GetOutputPowerStateOk returns a tuple with the OutputPowerState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentPsuControlAllOf) GetOutputPowerStateOk() (*string, bool) {
	if o == nil || o.OutputPowerState == nil {
		return nil, false
	}
	return o.OutputPowerState, true
}

// HasOutputPowerState returns a boolean if a field has been set.
func (o *EquipmentPsuControlAllOf) HasOutputPowerState() bool {
	if o != nil && o.OutputPowerState != nil {
		return true
	}

	return false
}

// SetOutputPowerState gets a reference to the given string and assigns it to the OutputPowerState field.
func (o *EquipmentPsuControlAllOf) SetOutputPowerState(v string) {
	o.OutputPowerState = &v
}

// GetRedundancy returns the Redundancy field value if set, zero value otherwise.
func (o *EquipmentPsuControlAllOf) GetRedundancy() string {
	if o == nil || o.Redundancy == nil {
		var ret string
		return ret
	}
	return *o.Redundancy
}

// GetRedundancyOk returns a tuple with the Redundancy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentPsuControlAllOf) GetRedundancyOk() (*string, bool) {
	if o == nil || o.Redundancy == nil {
		return nil, false
	}
	return o.Redundancy, true
}

// HasRedundancy returns a boolean if a field has been set.
func (o *EquipmentPsuControlAllOf) HasRedundancy() bool {
	if o != nil && o.Redundancy != nil {
		return true
	}

	return false
}

// SetRedundancy gets a reference to the given string and assigns it to the Redundancy field.
func (o *EquipmentPsuControlAllOf) SetRedundancy(v string) {
	o.Redundancy = &v
}

// GetEquipmentChassis returns the EquipmentChassis field value if set, zero value otherwise.
func (o *EquipmentPsuControlAllOf) GetEquipmentChassis() EquipmentChassisRelationship {
	if o == nil || o.EquipmentChassis == nil {
		var ret EquipmentChassisRelationship
		return ret
	}
	return *o.EquipmentChassis
}

// GetEquipmentChassisOk returns a tuple with the EquipmentChassis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentPsuControlAllOf) GetEquipmentChassisOk() (*EquipmentChassisRelationship, bool) {
	if o == nil || o.EquipmentChassis == nil {
		return nil, false
	}
	return o.EquipmentChassis, true
}

// HasEquipmentChassis returns a boolean if a field has been set.
func (o *EquipmentPsuControlAllOf) HasEquipmentChassis() bool {
	if o != nil && o.EquipmentChassis != nil {
		return true
	}

	return false
}

// SetEquipmentChassis gets a reference to the given EquipmentChassisRelationship and assigns it to the EquipmentChassis field.
func (o *EquipmentPsuControlAllOf) SetEquipmentChassis(v EquipmentChassisRelationship) {
	o.EquipmentChassis = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *EquipmentPsuControlAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentPsuControlAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *EquipmentPsuControlAllOf) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *EquipmentPsuControlAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *EquipmentPsuControlAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentPsuControlAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *EquipmentPsuControlAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *EquipmentPsuControlAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o EquipmentPsuControlAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ClusterState != nil {
		toSerialize["ClusterState"] = o.ClusterState
	}
	if o.InputPowerState != nil {
		toSerialize["InputPowerState"] = o.InputPowerState
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.OperQualifier != nil {
		toSerialize["OperQualifier"] = o.OperQualifier
	}
	if o.OperReason != nil {
		toSerialize["OperReason"] = o.OperReason
	}
	if o.OperState != nil {
		toSerialize["OperState"] = o.OperState
	}
	if o.OutputPowerState != nil {
		toSerialize["OutputPowerState"] = o.OutputPowerState
	}
	if o.Redundancy != nil {
		toSerialize["Redundancy"] = o.Redundancy
	}
	if o.EquipmentChassis != nil {
		toSerialize["EquipmentChassis"] = o.EquipmentChassis
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

func (o *EquipmentPsuControlAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varEquipmentPsuControlAllOf := _EquipmentPsuControlAllOf{}

	if err = json.Unmarshal(bytes, &varEquipmentPsuControlAllOf); err == nil {
		*o = EquipmentPsuControlAllOf(varEquipmentPsuControlAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ClusterState")
		delete(additionalProperties, "InputPowerState")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "OperQualifier")
		delete(additionalProperties, "OperReason")
		delete(additionalProperties, "OperState")
		delete(additionalProperties, "OutputPowerState")
		delete(additionalProperties, "Redundancy")
		delete(additionalProperties, "EquipmentChassis")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEquipmentPsuControlAllOf struct {
	value *EquipmentPsuControlAllOf
	isSet bool
}

func (v NullableEquipmentPsuControlAllOf) Get() *EquipmentPsuControlAllOf {
	return v.value
}

func (v *NullableEquipmentPsuControlAllOf) Set(val *EquipmentPsuControlAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEquipmentPsuControlAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEquipmentPsuControlAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquipmentPsuControlAllOf(val *EquipmentPsuControlAllOf) *NullableEquipmentPsuControlAllOf {
	return &NullableEquipmentPsuControlAllOf{value: val, isSet: true}
}

func (v NullableEquipmentPsuControlAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquipmentPsuControlAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
