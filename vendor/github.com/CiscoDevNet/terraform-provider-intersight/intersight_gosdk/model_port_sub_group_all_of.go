/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-11-20T05:29:54Z.
 *
 * API version: 1.0.9-2713
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// PortSubGroupAllOf Definition of the list of properties defined in 'port.SubGroup', excluding properties defined in parent classes.
type PortSubGroupAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Type of port sub-group. Values are Eth or Fc.
	Transport *string `json:"Transport,omitempty"`
	// An array of relationships to etherPhysicalPort resources.
	EthernetPorts []EtherPhysicalPortRelationship `json:"EthernetPorts,omitempty"`
	// An array of relationships to fcPhysicalPort resources.
	FcPorts              []FcPhysicalPortRelationship         `json:"FcPorts,omitempty"`
	InventoryDeviceInfo  *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
	PortGroup            *PortGroupRelationship               `json:"PortGroup,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PortSubGroupAllOf PortSubGroupAllOf

// NewPortSubGroupAllOf instantiates a new PortSubGroupAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortSubGroupAllOf(classId string, objectType string) *PortSubGroupAllOf {
	this := PortSubGroupAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewPortSubGroupAllOfWithDefaults instantiates a new PortSubGroupAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortSubGroupAllOfWithDefaults() *PortSubGroupAllOf {
	this := PortSubGroupAllOf{}
	var classId string = "port.SubGroup"
	this.ClassId = classId
	var objectType string = "port.SubGroup"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *PortSubGroupAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *PortSubGroupAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *PortSubGroupAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *PortSubGroupAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *PortSubGroupAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *PortSubGroupAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetTransport returns the Transport field value if set, zero value otherwise.
func (o *PortSubGroupAllOf) GetTransport() string {
	if o == nil || o.Transport == nil {
		var ret string
		return ret
	}
	return *o.Transport
}

// GetTransportOk returns a tuple with the Transport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortSubGroupAllOf) GetTransportOk() (*string, bool) {
	if o == nil || o.Transport == nil {
		return nil, false
	}
	return o.Transport, true
}

// HasTransport returns a boolean if a field has been set.
func (o *PortSubGroupAllOf) HasTransport() bool {
	if o != nil && o.Transport != nil {
		return true
	}

	return false
}

// SetTransport gets a reference to the given string and assigns it to the Transport field.
func (o *PortSubGroupAllOf) SetTransport(v string) {
	o.Transport = &v
}

// GetEthernetPorts returns the EthernetPorts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PortSubGroupAllOf) GetEthernetPorts() []EtherPhysicalPortRelationship {
	if o == nil {
		var ret []EtherPhysicalPortRelationship
		return ret
	}
	return o.EthernetPorts
}

// GetEthernetPortsOk returns a tuple with the EthernetPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PortSubGroupAllOf) GetEthernetPortsOk() (*[]EtherPhysicalPortRelationship, bool) {
	if o == nil || o.EthernetPorts == nil {
		return nil, false
	}
	return &o.EthernetPorts, true
}

// HasEthernetPorts returns a boolean if a field has been set.
func (o *PortSubGroupAllOf) HasEthernetPorts() bool {
	if o != nil && o.EthernetPorts != nil {
		return true
	}

	return false
}

// SetEthernetPorts gets a reference to the given []EtherPhysicalPortRelationship and assigns it to the EthernetPorts field.
func (o *PortSubGroupAllOf) SetEthernetPorts(v []EtherPhysicalPortRelationship) {
	o.EthernetPorts = v
}

// GetFcPorts returns the FcPorts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PortSubGroupAllOf) GetFcPorts() []FcPhysicalPortRelationship {
	if o == nil {
		var ret []FcPhysicalPortRelationship
		return ret
	}
	return o.FcPorts
}

// GetFcPortsOk returns a tuple with the FcPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PortSubGroupAllOf) GetFcPortsOk() (*[]FcPhysicalPortRelationship, bool) {
	if o == nil || o.FcPorts == nil {
		return nil, false
	}
	return &o.FcPorts, true
}

// HasFcPorts returns a boolean if a field has been set.
func (o *PortSubGroupAllOf) HasFcPorts() bool {
	if o != nil && o.FcPorts != nil {
		return true
	}

	return false
}

// SetFcPorts gets a reference to the given []FcPhysicalPortRelationship and assigns it to the FcPorts field.
func (o *PortSubGroupAllOf) SetFcPorts(v []FcPhysicalPortRelationship) {
	o.FcPorts = v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *PortSubGroupAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortSubGroupAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *PortSubGroupAllOf) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *PortSubGroupAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetPortGroup returns the PortGroup field value if set, zero value otherwise.
func (o *PortSubGroupAllOf) GetPortGroup() PortGroupRelationship {
	if o == nil || o.PortGroup == nil {
		var ret PortGroupRelationship
		return ret
	}
	return *o.PortGroup
}

// GetPortGroupOk returns a tuple with the PortGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortSubGroupAllOf) GetPortGroupOk() (*PortGroupRelationship, bool) {
	if o == nil || o.PortGroup == nil {
		return nil, false
	}
	return o.PortGroup, true
}

// HasPortGroup returns a boolean if a field has been set.
func (o *PortSubGroupAllOf) HasPortGroup() bool {
	if o != nil && o.PortGroup != nil {
		return true
	}

	return false
}

// SetPortGroup gets a reference to the given PortGroupRelationship and assigns it to the PortGroup field.
func (o *PortSubGroupAllOf) SetPortGroup(v PortGroupRelationship) {
	o.PortGroup = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *PortSubGroupAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortSubGroupAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *PortSubGroupAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *PortSubGroupAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o PortSubGroupAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Transport != nil {
		toSerialize["Transport"] = o.Transport
	}
	if o.EthernetPorts != nil {
		toSerialize["EthernetPorts"] = o.EthernetPorts
	}
	if o.FcPorts != nil {
		toSerialize["FcPorts"] = o.FcPorts
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.PortGroup != nil {
		toSerialize["PortGroup"] = o.PortGroup
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PortSubGroupAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varPortSubGroupAllOf := _PortSubGroupAllOf{}

	if err = json.Unmarshal(bytes, &varPortSubGroupAllOf); err == nil {
		*o = PortSubGroupAllOf(varPortSubGroupAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Transport")
		delete(additionalProperties, "EthernetPorts")
		delete(additionalProperties, "FcPorts")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "PortGroup")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePortSubGroupAllOf struct {
	value *PortSubGroupAllOf
	isSet bool
}

func (v NullablePortSubGroupAllOf) Get() *PortSubGroupAllOf {
	return v.value
}

func (v *NullablePortSubGroupAllOf) Set(val *PortSubGroupAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePortSubGroupAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePortSubGroupAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortSubGroupAllOf(val *PortSubGroupAllOf) *NullablePortSubGroupAllOf {
	return &NullablePortSubGroupAllOf{value: val, isSet: true}
}

func (v NullablePortSubGroupAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortSubGroupAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}