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

// FcPhysicalPortAllOf Definition of the list of properties defined in 'fc.PhysicalPort', excluding properties defined in parent classes.
type FcPhysicalPortAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Administrator configured Speed applied on the port.
	AdminSpeed *string `json:"AdminSpeed,omitempty"`
	// Administratively configured state (enabled/disabled) for this port.
	AdminState *string `json:"AdminState,omitempty"`
	// Buffer to Buffer credits of FC port.
	B2bCredit *int64 `json:"B2bCredit,omitempty"`
	// Maximum Speed with which the port operates.
	MaxSpeed *string `json:"MaxSpeed,omitempty"`
	// Mode information N_proxy, F or E associated to the Fibre Channel port.
	Mode *string `json:"Mode,omitempty"`
	// Operational Speed with which the port operates.
	OperSpeed *string `json:"OperSpeed,omitempty"`
	// PeerDn for fibre channel physical port.
	PeerDn *string `json:"PeerDn,omitempty"`
	// Port channel id of FC port channel created on FI switch.
	PortChannelId *int64 `json:"PortChannelId,omitempty"`
	// Transceiver type of a Fibre Channel port.
	TransceiverType *string `json:"TransceiverType,omitempty"`
	// Virtual San that is associated to the port.
	Vsan *int64 `json:"Vsan,omitempty"`
	// World Wide Name of a Fibre Channel port.
	Wwn                  *string                              `json:"Wwn,omitempty"`
	InventoryDeviceInfo  *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
	PortGroup            *PortGroupRelationship               `json:"PortGroup,omitempty"`
	PortSubGroup         *PortSubGroupRelationship            `json:"PortSubGroup,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FcPhysicalPortAllOf FcPhysicalPortAllOf

// NewFcPhysicalPortAllOf instantiates a new FcPhysicalPortAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFcPhysicalPortAllOf(classId string, objectType string) *FcPhysicalPortAllOf {
	this := FcPhysicalPortAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewFcPhysicalPortAllOfWithDefaults instantiates a new FcPhysicalPortAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFcPhysicalPortAllOfWithDefaults() *FcPhysicalPortAllOf {
	this := FcPhysicalPortAllOf{}
	var classId string = "fc.PhysicalPort"
	this.ClassId = classId
	var objectType string = "fc.PhysicalPort"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *FcPhysicalPortAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FcPhysicalPortAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FcPhysicalPortAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FcPhysicalPortAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FcPhysicalPortAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FcPhysicalPortAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAdminSpeed returns the AdminSpeed field value if set, zero value otherwise.
func (o *FcPhysicalPortAllOf) GetAdminSpeed() string {
	if o == nil || o.AdminSpeed == nil {
		var ret string
		return ret
	}
	return *o.AdminSpeed
}

// GetAdminSpeedOk returns a tuple with the AdminSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcPhysicalPortAllOf) GetAdminSpeedOk() (*string, bool) {
	if o == nil || o.AdminSpeed == nil {
		return nil, false
	}
	return o.AdminSpeed, true
}

// HasAdminSpeed returns a boolean if a field has been set.
func (o *FcPhysicalPortAllOf) HasAdminSpeed() bool {
	if o != nil && o.AdminSpeed != nil {
		return true
	}

	return false
}

// SetAdminSpeed gets a reference to the given string and assigns it to the AdminSpeed field.
func (o *FcPhysicalPortAllOf) SetAdminSpeed(v string) {
	o.AdminSpeed = &v
}

// GetAdminState returns the AdminState field value if set, zero value otherwise.
func (o *FcPhysicalPortAllOf) GetAdminState() string {
	if o == nil || o.AdminState == nil {
		var ret string
		return ret
	}
	return *o.AdminState
}

// GetAdminStateOk returns a tuple with the AdminState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcPhysicalPortAllOf) GetAdminStateOk() (*string, bool) {
	if o == nil || o.AdminState == nil {
		return nil, false
	}
	return o.AdminState, true
}

// HasAdminState returns a boolean if a field has been set.
func (o *FcPhysicalPortAllOf) HasAdminState() bool {
	if o != nil && o.AdminState != nil {
		return true
	}

	return false
}

// SetAdminState gets a reference to the given string and assigns it to the AdminState field.
func (o *FcPhysicalPortAllOf) SetAdminState(v string) {
	o.AdminState = &v
}

// GetB2bCredit returns the B2bCredit field value if set, zero value otherwise.
func (o *FcPhysicalPortAllOf) GetB2bCredit() int64 {
	if o == nil || o.B2bCredit == nil {
		var ret int64
		return ret
	}
	return *o.B2bCredit
}

// GetB2bCreditOk returns a tuple with the B2bCredit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcPhysicalPortAllOf) GetB2bCreditOk() (*int64, bool) {
	if o == nil || o.B2bCredit == nil {
		return nil, false
	}
	return o.B2bCredit, true
}

// HasB2bCredit returns a boolean if a field has been set.
func (o *FcPhysicalPortAllOf) HasB2bCredit() bool {
	if o != nil && o.B2bCredit != nil {
		return true
	}

	return false
}

// SetB2bCredit gets a reference to the given int64 and assigns it to the B2bCredit field.
func (o *FcPhysicalPortAllOf) SetB2bCredit(v int64) {
	o.B2bCredit = &v
}

// GetMaxSpeed returns the MaxSpeed field value if set, zero value otherwise.
func (o *FcPhysicalPortAllOf) GetMaxSpeed() string {
	if o == nil || o.MaxSpeed == nil {
		var ret string
		return ret
	}
	return *o.MaxSpeed
}

// GetMaxSpeedOk returns a tuple with the MaxSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcPhysicalPortAllOf) GetMaxSpeedOk() (*string, bool) {
	if o == nil || o.MaxSpeed == nil {
		return nil, false
	}
	return o.MaxSpeed, true
}

// HasMaxSpeed returns a boolean if a field has been set.
func (o *FcPhysicalPortAllOf) HasMaxSpeed() bool {
	if o != nil && o.MaxSpeed != nil {
		return true
	}

	return false
}

// SetMaxSpeed gets a reference to the given string and assigns it to the MaxSpeed field.
func (o *FcPhysicalPortAllOf) SetMaxSpeed(v string) {
	o.MaxSpeed = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *FcPhysicalPortAllOf) GetMode() string {
	if o == nil || o.Mode == nil {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcPhysicalPortAllOf) GetModeOk() (*string, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *FcPhysicalPortAllOf) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *FcPhysicalPortAllOf) SetMode(v string) {
	o.Mode = &v
}

// GetOperSpeed returns the OperSpeed field value if set, zero value otherwise.
func (o *FcPhysicalPortAllOf) GetOperSpeed() string {
	if o == nil || o.OperSpeed == nil {
		var ret string
		return ret
	}
	return *o.OperSpeed
}

// GetOperSpeedOk returns a tuple with the OperSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcPhysicalPortAllOf) GetOperSpeedOk() (*string, bool) {
	if o == nil || o.OperSpeed == nil {
		return nil, false
	}
	return o.OperSpeed, true
}

// HasOperSpeed returns a boolean if a field has been set.
func (o *FcPhysicalPortAllOf) HasOperSpeed() bool {
	if o != nil && o.OperSpeed != nil {
		return true
	}

	return false
}

// SetOperSpeed gets a reference to the given string and assigns it to the OperSpeed field.
func (o *FcPhysicalPortAllOf) SetOperSpeed(v string) {
	o.OperSpeed = &v
}

// GetPeerDn returns the PeerDn field value if set, zero value otherwise.
func (o *FcPhysicalPortAllOf) GetPeerDn() string {
	if o == nil || o.PeerDn == nil {
		var ret string
		return ret
	}
	return *o.PeerDn
}

// GetPeerDnOk returns a tuple with the PeerDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcPhysicalPortAllOf) GetPeerDnOk() (*string, bool) {
	if o == nil || o.PeerDn == nil {
		return nil, false
	}
	return o.PeerDn, true
}

// HasPeerDn returns a boolean if a field has been set.
func (o *FcPhysicalPortAllOf) HasPeerDn() bool {
	if o != nil && o.PeerDn != nil {
		return true
	}

	return false
}

// SetPeerDn gets a reference to the given string and assigns it to the PeerDn field.
func (o *FcPhysicalPortAllOf) SetPeerDn(v string) {
	o.PeerDn = &v
}

// GetPortChannelId returns the PortChannelId field value if set, zero value otherwise.
func (o *FcPhysicalPortAllOf) GetPortChannelId() int64 {
	if o == nil || o.PortChannelId == nil {
		var ret int64
		return ret
	}
	return *o.PortChannelId
}

// GetPortChannelIdOk returns a tuple with the PortChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcPhysicalPortAllOf) GetPortChannelIdOk() (*int64, bool) {
	if o == nil || o.PortChannelId == nil {
		return nil, false
	}
	return o.PortChannelId, true
}

// HasPortChannelId returns a boolean if a field has been set.
func (o *FcPhysicalPortAllOf) HasPortChannelId() bool {
	if o != nil && o.PortChannelId != nil {
		return true
	}

	return false
}

// SetPortChannelId gets a reference to the given int64 and assigns it to the PortChannelId field.
func (o *FcPhysicalPortAllOf) SetPortChannelId(v int64) {
	o.PortChannelId = &v
}

// GetTransceiverType returns the TransceiverType field value if set, zero value otherwise.
func (o *FcPhysicalPortAllOf) GetTransceiverType() string {
	if o == nil || o.TransceiverType == nil {
		var ret string
		return ret
	}
	return *o.TransceiverType
}

// GetTransceiverTypeOk returns a tuple with the TransceiverType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcPhysicalPortAllOf) GetTransceiverTypeOk() (*string, bool) {
	if o == nil || o.TransceiverType == nil {
		return nil, false
	}
	return o.TransceiverType, true
}

// HasTransceiverType returns a boolean if a field has been set.
func (o *FcPhysicalPortAllOf) HasTransceiverType() bool {
	if o != nil && o.TransceiverType != nil {
		return true
	}

	return false
}

// SetTransceiverType gets a reference to the given string and assigns it to the TransceiverType field.
func (o *FcPhysicalPortAllOf) SetTransceiverType(v string) {
	o.TransceiverType = &v
}

// GetVsan returns the Vsan field value if set, zero value otherwise.
func (o *FcPhysicalPortAllOf) GetVsan() int64 {
	if o == nil || o.Vsan == nil {
		var ret int64
		return ret
	}
	return *o.Vsan
}

// GetVsanOk returns a tuple with the Vsan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcPhysicalPortAllOf) GetVsanOk() (*int64, bool) {
	if o == nil || o.Vsan == nil {
		return nil, false
	}
	return o.Vsan, true
}

// HasVsan returns a boolean if a field has been set.
func (o *FcPhysicalPortAllOf) HasVsan() bool {
	if o != nil && o.Vsan != nil {
		return true
	}

	return false
}

// SetVsan gets a reference to the given int64 and assigns it to the Vsan field.
func (o *FcPhysicalPortAllOf) SetVsan(v int64) {
	o.Vsan = &v
}

// GetWwn returns the Wwn field value if set, zero value otherwise.
func (o *FcPhysicalPortAllOf) GetWwn() string {
	if o == nil || o.Wwn == nil {
		var ret string
		return ret
	}
	return *o.Wwn
}

// GetWwnOk returns a tuple with the Wwn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcPhysicalPortAllOf) GetWwnOk() (*string, bool) {
	if o == nil || o.Wwn == nil {
		return nil, false
	}
	return o.Wwn, true
}

// HasWwn returns a boolean if a field has been set.
func (o *FcPhysicalPortAllOf) HasWwn() bool {
	if o != nil && o.Wwn != nil {
		return true
	}

	return false
}

// SetWwn gets a reference to the given string and assigns it to the Wwn field.
func (o *FcPhysicalPortAllOf) SetWwn(v string) {
	o.Wwn = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *FcPhysicalPortAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcPhysicalPortAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *FcPhysicalPortAllOf) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *FcPhysicalPortAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetPortGroup returns the PortGroup field value if set, zero value otherwise.
func (o *FcPhysicalPortAllOf) GetPortGroup() PortGroupRelationship {
	if o == nil || o.PortGroup == nil {
		var ret PortGroupRelationship
		return ret
	}
	return *o.PortGroup
}

// GetPortGroupOk returns a tuple with the PortGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcPhysicalPortAllOf) GetPortGroupOk() (*PortGroupRelationship, bool) {
	if o == nil || o.PortGroup == nil {
		return nil, false
	}
	return o.PortGroup, true
}

// HasPortGroup returns a boolean if a field has been set.
func (o *FcPhysicalPortAllOf) HasPortGroup() bool {
	if o != nil && o.PortGroup != nil {
		return true
	}

	return false
}

// SetPortGroup gets a reference to the given PortGroupRelationship and assigns it to the PortGroup field.
func (o *FcPhysicalPortAllOf) SetPortGroup(v PortGroupRelationship) {
	o.PortGroup = &v
}

// GetPortSubGroup returns the PortSubGroup field value if set, zero value otherwise.
func (o *FcPhysicalPortAllOf) GetPortSubGroup() PortSubGroupRelationship {
	if o == nil || o.PortSubGroup == nil {
		var ret PortSubGroupRelationship
		return ret
	}
	return *o.PortSubGroup
}

// GetPortSubGroupOk returns a tuple with the PortSubGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcPhysicalPortAllOf) GetPortSubGroupOk() (*PortSubGroupRelationship, bool) {
	if o == nil || o.PortSubGroup == nil {
		return nil, false
	}
	return o.PortSubGroup, true
}

// HasPortSubGroup returns a boolean if a field has been set.
func (o *FcPhysicalPortAllOf) HasPortSubGroup() bool {
	if o != nil && o.PortSubGroup != nil {
		return true
	}

	return false
}

// SetPortSubGroup gets a reference to the given PortSubGroupRelationship and assigns it to the PortSubGroup field.
func (o *FcPhysicalPortAllOf) SetPortSubGroup(v PortSubGroupRelationship) {
	o.PortSubGroup = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *FcPhysicalPortAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcPhysicalPortAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *FcPhysicalPortAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *FcPhysicalPortAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o FcPhysicalPortAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AdminSpeed != nil {
		toSerialize["AdminSpeed"] = o.AdminSpeed
	}
	if o.AdminState != nil {
		toSerialize["AdminState"] = o.AdminState
	}
	if o.B2bCredit != nil {
		toSerialize["B2bCredit"] = o.B2bCredit
	}
	if o.MaxSpeed != nil {
		toSerialize["MaxSpeed"] = o.MaxSpeed
	}
	if o.Mode != nil {
		toSerialize["Mode"] = o.Mode
	}
	if o.OperSpeed != nil {
		toSerialize["OperSpeed"] = o.OperSpeed
	}
	if o.PeerDn != nil {
		toSerialize["PeerDn"] = o.PeerDn
	}
	if o.PortChannelId != nil {
		toSerialize["PortChannelId"] = o.PortChannelId
	}
	if o.TransceiverType != nil {
		toSerialize["TransceiverType"] = o.TransceiverType
	}
	if o.Vsan != nil {
		toSerialize["Vsan"] = o.Vsan
	}
	if o.Wwn != nil {
		toSerialize["Wwn"] = o.Wwn
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.PortGroup != nil {
		toSerialize["PortGroup"] = o.PortGroup
	}
	if o.PortSubGroup != nil {
		toSerialize["PortSubGroup"] = o.PortSubGroup
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FcPhysicalPortAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varFcPhysicalPortAllOf := _FcPhysicalPortAllOf{}

	if err = json.Unmarshal(bytes, &varFcPhysicalPortAllOf); err == nil {
		*o = FcPhysicalPortAllOf(varFcPhysicalPortAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AdminSpeed")
		delete(additionalProperties, "AdminState")
		delete(additionalProperties, "B2bCredit")
		delete(additionalProperties, "MaxSpeed")
		delete(additionalProperties, "Mode")
		delete(additionalProperties, "OperSpeed")
		delete(additionalProperties, "PeerDn")
		delete(additionalProperties, "PortChannelId")
		delete(additionalProperties, "TransceiverType")
		delete(additionalProperties, "Vsan")
		delete(additionalProperties, "Wwn")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "PortGroup")
		delete(additionalProperties, "PortSubGroup")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFcPhysicalPortAllOf struct {
	value *FcPhysicalPortAllOf
	isSet bool
}

func (v NullableFcPhysicalPortAllOf) Get() *FcPhysicalPortAllOf {
	return v.value
}

func (v *NullableFcPhysicalPortAllOf) Set(val *FcPhysicalPortAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFcPhysicalPortAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFcPhysicalPortAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFcPhysicalPortAllOf(val *FcPhysicalPortAllOf) *NullableFcPhysicalPortAllOf {
	return &NullableFcPhysicalPortAllOf{value: val, isSet: true}
}

func (v NullableFcPhysicalPortAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFcPhysicalPortAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
