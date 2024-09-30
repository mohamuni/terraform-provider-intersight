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

// checks if the PortMacBinding type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PortMacBinding{}

// PortMacBinding Establishes relationship between the ports and connected end points based on LLDP TLVs.
type PortMacBinding struct {
	InventoryBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Aggregate Port ID of the local Switch Interface.
	AggregatePortId *int64 `json:"AggregatePortId,omitempty"`
	// Chassis/FEX device idetifier that is local to a cluster.
	ChassisId *int64 `json:"ChassisId,omitempty"`
	// Chassis/Rack Model that is associated with the Switch/FEX interface.
	ChassisModel *string `json:"ChassisModel,omitempty"`
	// Chassis/Rack Serial that is associated with the Switch/FEX interface.
	ChassisSerial *string `json:"ChassisSerial,omitempty"`
	// Chassis/Rack Vendor that is associated with the Switch/FEX interface.
	ChassisVendor *string `json:"ChassisVendor,omitempty"`
	// Device ID value that is advertised and available as a part of LLDP TLV.
	DeviceMac *string `json:"DeviceMac,omitempty"`
	// IOM/SIOC/Adapter Mode that is associated with the Switch/FEX interface.
	ModuleMode *int64 `json:"ModuleMode,omitempty"`
	// IOM/SIOC/Adapter Model that is associated with the Switch/FEX interface.
	ModuleModel *string `json:"ModuleModel,omitempty"`
	// Uplink port identifier of the VIC that is associated with the Switch/FEX interface.
	ModulePortId *int64 `json:"ModulePortId,omitempty"`
	// IOM/SIOC/Adapter Serial that is associated with the Switch/FEX interface.
	ModuleSerial *string `json:"ModuleSerial,omitempty"`
	// IOM/SIOC/Adapter Side that is associated with the Switch/FEX interface.
	ModuleSide *int64 `json:"ModuleSide,omitempty"`
	// IOM/SIOC/Adapter Slot that is associated with the Switch/FEX interface.
	ModuleSlot *int64 `json:"ModuleSlot,omitempty"`
	// IOM/SIOC/Adapter Vendor that is associated with the Switch/FEX interface.
	ModuleVendor *string `json:"ModuleVendor,omitempty"`
	// Port ID of the local Switch Interface.
	PortId *int64 `json:"PortId,omitempty"`
	// Port ID value that is advertised and available as a part of LLDP TLV.
	PortMac *string `json:"PortMac,omitempty"`
	// Slot ID of the local Switch slot Interface.
	SlotId *int64 `json:"SlotId,omitempty"`
	// Switch Identifier that is local to a cluster.
	SwitchId             *int64                                      `json:"SwitchId,omitempty"`
	NetworkElement       NullableNetworkElementRelationship          `json:"NetworkElement,omitempty"`
	RegisteredDevice     NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PortMacBinding PortMacBinding

// NewPortMacBinding instantiates a new PortMacBinding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortMacBinding(classId string, objectType string) *PortMacBinding {
	this := PortMacBinding{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewPortMacBindingWithDefaults instantiates a new PortMacBinding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortMacBindingWithDefaults() *PortMacBinding {
	this := PortMacBinding{}
	var classId string = "port.MacBinding"
	this.ClassId = classId
	var objectType string = "port.MacBinding"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *PortMacBinding) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *PortMacBinding) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *PortMacBinding) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "port.MacBinding" of the ClassId field.
func (o *PortMacBinding) GetDefaultClassId() interface{} {
	return "port.MacBinding"
}

// GetObjectType returns the ObjectType field value
func (o *PortMacBinding) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *PortMacBinding) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *PortMacBinding) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "port.MacBinding" of the ObjectType field.
func (o *PortMacBinding) GetDefaultObjectType() interface{} {
	return "port.MacBinding"
}

// GetAggregatePortId returns the AggregatePortId field value if set, zero value otherwise.
func (o *PortMacBinding) GetAggregatePortId() int64 {
	if o == nil || IsNil(o.AggregatePortId) {
		var ret int64
		return ret
	}
	return *o.AggregatePortId
}

// GetAggregatePortIdOk returns a tuple with the AggregatePortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortMacBinding) GetAggregatePortIdOk() (*int64, bool) {
	if o == nil || IsNil(o.AggregatePortId) {
		return nil, false
	}
	return o.AggregatePortId, true
}

// HasAggregatePortId returns a boolean if a field has been set.
func (o *PortMacBinding) HasAggregatePortId() bool {
	if o != nil && !IsNil(o.AggregatePortId) {
		return true
	}

	return false
}

// SetAggregatePortId gets a reference to the given int64 and assigns it to the AggregatePortId field.
func (o *PortMacBinding) SetAggregatePortId(v int64) {
	o.AggregatePortId = &v
}

// GetChassisId returns the ChassisId field value if set, zero value otherwise.
func (o *PortMacBinding) GetChassisId() int64 {
	if o == nil || IsNil(o.ChassisId) {
		var ret int64
		return ret
	}
	return *o.ChassisId
}

// GetChassisIdOk returns a tuple with the ChassisId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortMacBinding) GetChassisIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ChassisId) {
		return nil, false
	}
	return o.ChassisId, true
}

// HasChassisId returns a boolean if a field has been set.
func (o *PortMacBinding) HasChassisId() bool {
	if o != nil && !IsNil(o.ChassisId) {
		return true
	}

	return false
}

// SetChassisId gets a reference to the given int64 and assigns it to the ChassisId field.
func (o *PortMacBinding) SetChassisId(v int64) {
	o.ChassisId = &v
}

// GetChassisModel returns the ChassisModel field value if set, zero value otherwise.
func (o *PortMacBinding) GetChassisModel() string {
	if o == nil || IsNil(o.ChassisModel) {
		var ret string
		return ret
	}
	return *o.ChassisModel
}

// GetChassisModelOk returns a tuple with the ChassisModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortMacBinding) GetChassisModelOk() (*string, bool) {
	if o == nil || IsNil(o.ChassisModel) {
		return nil, false
	}
	return o.ChassisModel, true
}

// HasChassisModel returns a boolean if a field has been set.
func (o *PortMacBinding) HasChassisModel() bool {
	if o != nil && !IsNil(o.ChassisModel) {
		return true
	}

	return false
}

// SetChassisModel gets a reference to the given string and assigns it to the ChassisModel field.
func (o *PortMacBinding) SetChassisModel(v string) {
	o.ChassisModel = &v
}

// GetChassisSerial returns the ChassisSerial field value if set, zero value otherwise.
func (o *PortMacBinding) GetChassisSerial() string {
	if o == nil || IsNil(o.ChassisSerial) {
		var ret string
		return ret
	}
	return *o.ChassisSerial
}

// GetChassisSerialOk returns a tuple with the ChassisSerial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortMacBinding) GetChassisSerialOk() (*string, bool) {
	if o == nil || IsNil(o.ChassisSerial) {
		return nil, false
	}
	return o.ChassisSerial, true
}

// HasChassisSerial returns a boolean if a field has been set.
func (o *PortMacBinding) HasChassisSerial() bool {
	if o != nil && !IsNil(o.ChassisSerial) {
		return true
	}

	return false
}

// SetChassisSerial gets a reference to the given string and assigns it to the ChassisSerial field.
func (o *PortMacBinding) SetChassisSerial(v string) {
	o.ChassisSerial = &v
}

// GetChassisVendor returns the ChassisVendor field value if set, zero value otherwise.
func (o *PortMacBinding) GetChassisVendor() string {
	if o == nil || IsNil(o.ChassisVendor) {
		var ret string
		return ret
	}
	return *o.ChassisVendor
}

// GetChassisVendorOk returns a tuple with the ChassisVendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortMacBinding) GetChassisVendorOk() (*string, bool) {
	if o == nil || IsNil(o.ChassisVendor) {
		return nil, false
	}
	return o.ChassisVendor, true
}

// HasChassisVendor returns a boolean if a field has been set.
func (o *PortMacBinding) HasChassisVendor() bool {
	if o != nil && !IsNil(o.ChassisVendor) {
		return true
	}

	return false
}

// SetChassisVendor gets a reference to the given string and assigns it to the ChassisVendor field.
func (o *PortMacBinding) SetChassisVendor(v string) {
	o.ChassisVendor = &v
}

// GetDeviceMac returns the DeviceMac field value if set, zero value otherwise.
func (o *PortMacBinding) GetDeviceMac() string {
	if o == nil || IsNil(o.DeviceMac) {
		var ret string
		return ret
	}
	return *o.DeviceMac
}

// GetDeviceMacOk returns a tuple with the DeviceMac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortMacBinding) GetDeviceMacOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceMac) {
		return nil, false
	}
	return o.DeviceMac, true
}

// HasDeviceMac returns a boolean if a field has been set.
func (o *PortMacBinding) HasDeviceMac() bool {
	if o != nil && !IsNil(o.DeviceMac) {
		return true
	}

	return false
}

// SetDeviceMac gets a reference to the given string and assigns it to the DeviceMac field.
func (o *PortMacBinding) SetDeviceMac(v string) {
	o.DeviceMac = &v
}

// GetModuleMode returns the ModuleMode field value if set, zero value otherwise.
func (o *PortMacBinding) GetModuleMode() int64 {
	if o == nil || IsNil(o.ModuleMode) {
		var ret int64
		return ret
	}
	return *o.ModuleMode
}

// GetModuleModeOk returns a tuple with the ModuleMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortMacBinding) GetModuleModeOk() (*int64, bool) {
	if o == nil || IsNil(o.ModuleMode) {
		return nil, false
	}
	return o.ModuleMode, true
}

// HasModuleMode returns a boolean if a field has been set.
func (o *PortMacBinding) HasModuleMode() bool {
	if o != nil && !IsNil(o.ModuleMode) {
		return true
	}

	return false
}

// SetModuleMode gets a reference to the given int64 and assigns it to the ModuleMode field.
func (o *PortMacBinding) SetModuleMode(v int64) {
	o.ModuleMode = &v
}

// GetModuleModel returns the ModuleModel field value if set, zero value otherwise.
func (o *PortMacBinding) GetModuleModel() string {
	if o == nil || IsNil(o.ModuleModel) {
		var ret string
		return ret
	}
	return *o.ModuleModel
}

// GetModuleModelOk returns a tuple with the ModuleModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortMacBinding) GetModuleModelOk() (*string, bool) {
	if o == nil || IsNil(o.ModuleModel) {
		return nil, false
	}
	return o.ModuleModel, true
}

// HasModuleModel returns a boolean if a field has been set.
func (o *PortMacBinding) HasModuleModel() bool {
	if o != nil && !IsNil(o.ModuleModel) {
		return true
	}

	return false
}

// SetModuleModel gets a reference to the given string and assigns it to the ModuleModel field.
func (o *PortMacBinding) SetModuleModel(v string) {
	o.ModuleModel = &v
}

// GetModulePortId returns the ModulePortId field value if set, zero value otherwise.
func (o *PortMacBinding) GetModulePortId() int64 {
	if o == nil || IsNil(o.ModulePortId) {
		var ret int64
		return ret
	}
	return *o.ModulePortId
}

// GetModulePortIdOk returns a tuple with the ModulePortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortMacBinding) GetModulePortIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ModulePortId) {
		return nil, false
	}
	return o.ModulePortId, true
}

// HasModulePortId returns a boolean if a field has been set.
func (o *PortMacBinding) HasModulePortId() bool {
	if o != nil && !IsNil(o.ModulePortId) {
		return true
	}

	return false
}

// SetModulePortId gets a reference to the given int64 and assigns it to the ModulePortId field.
func (o *PortMacBinding) SetModulePortId(v int64) {
	o.ModulePortId = &v
}

// GetModuleSerial returns the ModuleSerial field value if set, zero value otherwise.
func (o *PortMacBinding) GetModuleSerial() string {
	if o == nil || IsNil(o.ModuleSerial) {
		var ret string
		return ret
	}
	return *o.ModuleSerial
}

// GetModuleSerialOk returns a tuple with the ModuleSerial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortMacBinding) GetModuleSerialOk() (*string, bool) {
	if o == nil || IsNil(o.ModuleSerial) {
		return nil, false
	}
	return o.ModuleSerial, true
}

// HasModuleSerial returns a boolean if a field has been set.
func (o *PortMacBinding) HasModuleSerial() bool {
	if o != nil && !IsNil(o.ModuleSerial) {
		return true
	}

	return false
}

// SetModuleSerial gets a reference to the given string and assigns it to the ModuleSerial field.
func (o *PortMacBinding) SetModuleSerial(v string) {
	o.ModuleSerial = &v
}

// GetModuleSide returns the ModuleSide field value if set, zero value otherwise.
func (o *PortMacBinding) GetModuleSide() int64 {
	if o == nil || IsNil(o.ModuleSide) {
		var ret int64
		return ret
	}
	return *o.ModuleSide
}

// GetModuleSideOk returns a tuple with the ModuleSide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortMacBinding) GetModuleSideOk() (*int64, bool) {
	if o == nil || IsNil(o.ModuleSide) {
		return nil, false
	}
	return o.ModuleSide, true
}

// HasModuleSide returns a boolean if a field has been set.
func (o *PortMacBinding) HasModuleSide() bool {
	if o != nil && !IsNil(o.ModuleSide) {
		return true
	}

	return false
}

// SetModuleSide gets a reference to the given int64 and assigns it to the ModuleSide field.
func (o *PortMacBinding) SetModuleSide(v int64) {
	o.ModuleSide = &v
}

// GetModuleSlot returns the ModuleSlot field value if set, zero value otherwise.
func (o *PortMacBinding) GetModuleSlot() int64 {
	if o == nil || IsNil(o.ModuleSlot) {
		var ret int64
		return ret
	}
	return *o.ModuleSlot
}

// GetModuleSlotOk returns a tuple with the ModuleSlot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortMacBinding) GetModuleSlotOk() (*int64, bool) {
	if o == nil || IsNil(o.ModuleSlot) {
		return nil, false
	}
	return o.ModuleSlot, true
}

// HasModuleSlot returns a boolean if a field has been set.
func (o *PortMacBinding) HasModuleSlot() bool {
	if o != nil && !IsNil(o.ModuleSlot) {
		return true
	}

	return false
}

// SetModuleSlot gets a reference to the given int64 and assigns it to the ModuleSlot field.
func (o *PortMacBinding) SetModuleSlot(v int64) {
	o.ModuleSlot = &v
}

// GetModuleVendor returns the ModuleVendor field value if set, zero value otherwise.
func (o *PortMacBinding) GetModuleVendor() string {
	if o == nil || IsNil(o.ModuleVendor) {
		var ret string
		return ret
	}
	return *o.ModuleVendor
}

// GetModuleVendorOk returns a tuple with the ModuleVendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortMacBinding) GetModuleVendorOk() (*string, bool) {
	if o == nil || IsNil(o.ModuleVendor) {
		return nil, false
	}
	return o.ModuleVendor, true
}

// HasModuleVendor returns a boolean if a field has been set.
func (o *PortMacBinding) HasModuleVendor() bool {
	if o != nil && !IsNil(o.ModuleVendor) {
		return true
	}

	return false
}

// SetModuleVendor gets a reference to the given string and assigns it to the ModuleVendor field.
func (o *PortMacBinding) SetModuleVendor(v string) {
	o.ModuleVendor = &v
}

// GetPortId returns the PortId field value if set, zero value otherwise.
func (o *PortMacBinding) GetPortId() int64 {
	if o == nil || IsNil(o.PortId) {
		var ret int64
		return ret
	}
	return *o.PortId
}

// GetPortIdOk returns a tuple with the PortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortMacBinding) GetPortIdOk() (*int64, bool) {
	if o == nil || IsNil(o.PortId) {
		return nil, false
	}
	return o.PortId, true
}

// HasPortId returns a boolean if a field has been set.
func (o *PortMacBinding) HasPortId() bool {
	if o != nil && !IsNil(o.PortId) {
		return true
	}

	return false
}

// SetPortId gets a reference to the given int64 and assigns it to the PortId field.
func (o *PortMacBinding) SetPortId(v int64) {
	o.PortId = &v
}

// GetPortMac returns the PortMac field value if set, zero value otherwise.
func (o *PortMacBinding) GetPortMac() string {
	if o == nil || IsNil(o.PortMac) {
		var ret string
		return ret
	}
	return *o.PortMac
}

// GetPortMacOk returns a tuple with the PortMac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortMacBinding) GetPortMacOk() (*string, bool) {
	if o == nil || IsNil(o.PortMac) {
		return nil, false
	}
	return o.PortMac, true
}

// HasPortMac returns a boolean if a field has been set.
func (o *PortMacBinding) HasPortMac() bool {
	if o != nil && !IsNil(o.PortMac) {
		return true
	}

	return false
}

// SetPortMac gets a reference to the given string and assigns it to the PortMac field.
func (o *PortMacBinding) SetPortMac(v string) {
	o.PortMac = &v
}

// GetSlotId returns the SlotId field value if set, zero value otherwise.
func (o *PortMacBinding) GetSlotId() int64 {
	if o == nil || IsNil(o.SlotId) {
		var ret int64
		return ret
	}
	return *o.SlotId
}

// GetSlotIdOk returns a tuple with the SlotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortMacBinding) GetSlotIdOk() (*int64, bool) {
	if o == nil || IsNil(o.SlotId) {
		return nil, false
	}
	return o.SlotId, true
}

// HasSlotId returns a boolean if a field has been set.
func (o *PortMacBinding) HasSlotId() bool {
	if o != nil && !IsNil(o.SlotId) {
		return true
	}

	return false
}

// SetSlotId gets a reference to the given int64 and assigns it to the SlotId field.
func (o *PortMacBinding) SetSlotId(v int64) {
	o.SlotId = &v
}

// GetSwitchId returns the SwitchId field value if set, zero value otherwise.
func (o *PortMacBinding) GetSwitchId() int64 {
	if o == nil || IsNil(o.SwitchId) {
		var ret int64
		return ret
	}
	return *o.SwitchId
}

// GetSwitchIdOk returns a tuple with the SwitchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortMacBinding) GetSwitchIdOk() (*int64, bool) {
	if o == nil || IsNil(o.SwitchId) {
		return nil, false
	}
	return o.SwitchId, true
}

// HasSwitchId returns a boolean if a field has been set.
func (o *PortMacBinding) HasSwitchId() bool {
	if o != nil && !IsNil(o.SwitchId) {
		return true
	}

	return false
}

// SetSwitchId gets a reference to the given int64 and assigns it to the SwitchId field.
func (o *PortMacBinding) SetSwitchId(v int64) {
	o.SwitchId = &v
}

// GetNetworkElement returns the NetworkElement field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PortMacBinding) GetNetworkElement() NetworkElementRelationship {
	if o == nil || IsNil(o.NetworkElement.Get()) {
		var ret NetworkElementRelationship
		return ret
	}
	return *o.NetworkElement.Get()
}

// GetNetworkElementOk returns a tuple with the NetworkElement field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PortMacBinding) GetNetworkElementOk() (*NetworkElementRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.NetworkElement.Get(), o.NetworkElement.IsSet()
}

// HasNetworkElement returns a boolean if a field has been set.
func (o *PortMacBinding) HasNetworkElement() bool {
	if o != nil && o.NetworkElement.IsSet() {
		return true
	}

	return false
}

// SetNetworkElement gets a reference to the given NullableNetworkElementRelationship and assigns it to the NetworkElement field.
func (o *PortMacBinding) SetNetworkElement(v NetworkElementRelationship) {
	o.NetworkElement.Set(&v)
}

// SetNetworkElementNil sets the value for NetworkElement to be an explicit nil
func (o *PortMacBinding) SetNetworkElementNil() {
	o.NetworkElement.Set(nil)
}

// UnsetNetworkElement ensures that no value is present for NetworkElement, not even an explicit nil
func (o *PortMacBinding) UnsetNetworkElement() {
	o.NetworkElement.Unset()
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PortMacBinding) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PortMacBinding) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *PortMacBinding) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *PortMacBinding) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *PortMacBinding) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *PortMacBinding) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

func (o PortMacBinding) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PortMacBinding) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.AggregatePortId) {
		toSerialize["AggregatePortId"] = o.AggregatePortId
	}
	if !IsNil(o.ChassisId) {
		toSerialize["ChassisId"] = o.ChassisId
	}
	if !IsNil(o.ChassisModel) {
		toSerialize["ChassisModel"] = o.ChassisModel
	}
	if !IsNil(o.ChassisSerial) {
		toSerialize["ChassisSerial"] = o.ChassisSerial
	}
	if !IsNil(o.ChassisVendor) {
		toSerialize["ChassisVendor"] = o.ChassisVendor
	}
	if !IsNil(o.DeviceMac) {
		toSerialize["DeviceMac"] = o.DeviceMac
	}
	if !IsNil(o.ModuleMode) {
		toSerialize["ModuleMode"] = o.ModuleMode
	}
	if !IsNil(o.ModuleModel) {
		toSerialize["ModuleModel"] = o.ModuleModel
	}
	if !IsNil(o.ModulePortId) {
		toSerialize["ModulePortId"] = o.ModulePortId
	}
	if !IsNil(o.ModuleSerial) {
		toSerialize["ModuleSerial"] = o.ModuleSerial
	}
	if !IsNil(o.ModuleSide) {
		toSerialize["ModuleSide"] = o.ModuleSide
	}
	if !IsNil(o.ModuleSlot) {
		toSerialize["ModuleSlot"] = o.ModuleSlot
	}
	if !IsNil(o.ModuleVendor) {
		toSerialize["ModuleVendor"] = o.ModuleVendor
	}
	if !IsNil(o.PortId) {
		toSerialize["PortId"] = o.PortId
	}
	if !IsNil(o.PortMac) {
		toSerialize["PortMac"] = o.PortMac
	}
	if !IsNil(o.SlotId) {
		toSerialize["SlotId"] = o.SlotId
	}
	if !IsNil(o.SwitchId) {
		toSerialize["SwitchId"] = o.SwitchId
	}
	if o.NetworkElement.IsSet() {
		toSerialize["NetworkElement"] = o.NetworkElement.Get()
	}
	if o.RegisteredDevice.IsSet() {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PortMacBinding) UnmarshalJSON(data []byte) (err error) {
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
	type PortMacBindingWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Aggregate Port ID of the local Switch Interface.
		AggregatePortId *int64 `json:"AggregatePortId,omitempty"`
		// Chassis/FEX device idetifier that is local to a cluster.
		ChassisId *int64 `json:"ChassisId,omitempty"`
		// Chassis/Rack Model that is associated with the Switch/FEX interface.
		ChassisModel *string `json:"ChassisModel,omitempty"`
		// Chassis/Rack Serial that is associated with the Switch/FEX interface.
		ChassisSerial *string `json:"ChassisSerial,omitempty"`
		// Chassis/Rack Vendor that is associated with the Switch/FEX interface.
		ChassisVendor *string `json:"ChassisVendor,omitempty"`
		// Device ID value that is advertised and available as a part of LLDP TLV.
		DeviceMac *string `json:"DeviceMac,omitempty"`
		// IOM/SIOC/Adapter Mode that is associated with the Switch/FEX interface.
		ModuleMode *int64 `json:"ModuleMode,omitempty"`
		// IOM/SIOC/Adapter Model that is associated with the Switch/FEX interface.
		ModuleModel *string `json:"ModuleModel,omitempty"`
		// Uplink port identifier of the VIC that is associated with the Switch/FEX interface.
		ModulePortId *int64 `json:"ModulePortId,omitempty"`
		// IOM/SIOC/Adapter Serial that is associated with the Switch/FEX interface.
		ModuleSerial *string `json:"ModuleSerial,omitempty"`
		// IOM/SIOC/Adapter Side that is associated with the Switch/FEX interface.
		ModuleSide *int64 `json:"ModuleSide,omitempty"`
		// IOM/SIOC/Adapter Slot that is associated with the Switch/FEX interface.
		ModuleSlot *int64 `json:"ModuleSlot,omitempty"`
		// IOM/SIOC/Adapter Vendor that is associated with the Switch/FEX interface.
		ModuleVendor *string `json:"ModuleVendor,omitempty"`
		// Port ID of the local Switch Interface.
		PortId *int64 `json:"PortId,omitempty"`
		// Port ID value that is advertised and available as a part of LLDP TLV.
		PortMac *string `json:"PortMac,omitempty"`
		// Slot ID of the local Switch slot Interface.
		SlotId *int64 `json:"SlotId,omitempty"`
		// Switch Identifier that is local to a cluster.
		SwitchId         *int64                                      `json:"SwitchId,omitempty"`
		NetworkElement   NullableNetworkElementRelationship          `json:"NetworkElement,omitempty"`
		RegisteredDevice NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varPortMacBindingWithoutEmbeddedStruct := PortMacBindingWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varPortMacBindingWithoutEmbeddedStruct)
	if err == nil {
		varPortMacBinding := _PortMacBinding{}
		varPortMacBinding.ClassId = varPortMacBindingWithoutEmbeddedStruct.ClassId
		varPortMacBinding.ObjectType = varPortMacBindingWithoutEmbeddedStruct.ObjectType
		varPortMacBinding.AggregatePortId = varPortMacBindingWithoutEmbeddedStruct.AggregatePortId
		varPortMacBinding.ChassisId = varPortMacBindingWithoutEmbeddedStruct.ChassisId
		varPortMacBinding.ChassisModel = varPortMacBindingWithoutEmbeddedStruct.ChassisModel
		varPortMacBinding.ChassisSerial = varPortMacBindingWithoutEmbeddedStruct.ChassisSerial
		varPortMacBinding.ChassisVendor = varPortMacBindingWithoutEmbeddedStruct.ChassisVendor
		varPortMacBinding.DeviceMac = varPortMacBindingWithoutEmbeddedStruct.DeviceMac
		varPortMacBinding.ModuleMode = varPortMacBindingWithoutEmbeddedStruct.ModuleMode
		varPortMacBinding.ModuleModel = varPortMacBindingWithoutEmbeddedStruct.ModuleModel
		varPortMacBinding.ModulePortId = varPortMacBindingWithoutEmbeddedStruct.ModulePortId
		varPortMacBinding.ModuleSerial = varPortMacBindingWithoutEmbeddedStruct.ModuleSerial
		varPortMacBinding.ModuleSide = varPortMacBindingWithoutEmbeddedStruct.ModuleSide
		varPortMacBinding.ModuleSlot = varPortMacBindingWithoutEmbeddedStruct.ModuleSlot
		varPortMacBinding.ModuleVendor = varPortMacBindingWithoutEmbeddedStruct.ModuleVendor
		varPortMacBinding.PortId = varPortMacBindingWithoutEmbeddedStruct.PortId
		varPortMacBinding.PortMac = varPortMacBindingWithoutEmbeddedStruct.PortMac
		varPortMacBinding.SlotId = varPortMacBindingWithoutEmbeddedStruct.SlotId
		varPortMacBinding.SwitchId = varPortMacBindingWithoutEmbeddedStruct.SwitchId
		varPortMacBinding.NetworkElement = varPortMacBindingWithoutEmbeddedStruct.NetworkElement
		varPortMacBinding.RegisteredDevice = varPortMacBindingWithoutEmbeddedStruct.RegisteredDevice
		*o = PortMacBinding(varPortMacBinding)
	} else {
		return err
	}

	varPortMacBinding := _PortMacBinding{}

	err = json.Unmarshal(data, &varPortMacBinding)
	if err == nil {
		o.InventoryBase = varPortMacBinding.InventoryBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AggregatePortId")
		delete(additionalProperties, "ChassisId")
		delete(additionalProperties, "ChassisModel")
		delete(additionalProperties, "ChassisSerial")
		delete(additionalProperties, "ChassisVendor")
		delete(additionalProperties, "DeviceMac")
		delete(additionalProperties, "ModuleMode")
		delete(additionalProperties, "ModuleModel")
		delete(additionalProperties, "ModulePortId")
		delete(additionalProperties, "ModuleSerial")
		delete(additionalProperties, "ModuleSide")
		delete(additionalProperties, "ModuleSlot")
		delete(additionalProperties, "ModuleVendor")
		delete(additionalProperties, "PortId")
		delete(additionalProperties, "PortMac")
		delete(additionalProperties, "SlotId")
		delete(additionalProperties, "SwitchId")
		delete(additionalProperties, "NetworkElement")
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

type NullablePortMacBinding struct {
	value *PortMacBinding
	isSet bool
}

func (v NullablePortMacBinding) Get() *PortMacBinding {
	return v.value
}

func (v *NullablePortMacBinding) Set(val *PortMacBinding) {
	v.value = val
	v.isSet = true
}

func (v NullablePortMacBinding) IsSet() bool {
	return v.isSet
}

func (v *NullablePortMacBinding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortMacBinding(val *PortMacBinding) *NullablePortMacBinding {
	return &NullablePortMacBinding{value: val, isSet: true}
}

func (v NullablePortMacBinding) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortMacBinding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
