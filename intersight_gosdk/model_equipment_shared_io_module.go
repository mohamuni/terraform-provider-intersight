/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-2024100405
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

// checks if the EquipmentSharedIoModule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EquipmentSharedIoModule{}

// EquipmentSharedIoModule I/O Controller present inside SIOC to provide data path from S-series server to FI.
type EquipmentSharedIoModule struct {
	EquipmentBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// This field identifies the configuration state for this SIOM Unit.
	ConfigState *string `json:"ConfigState,omitempty"`
	// This field identifies the discovery state of SIOM.
	Discovery *string `json:"Discovery,omitempty"`
	// This field identifies the MAC of IOM-A side.
	MacOfSharedIomAside *string `json:"MacOfSharedIomAside,omitempty"`
	// This field identifies the MAC of IOM-B side.
	MacOfSharedIomBside *string `json:"MacOfSharedIomBside,omitempty"`
	// This field identifies the SIOM operational state.
	OperState *string `json:"OperState,omitempty"`
	// This field identifies the Part Number for this SIOM Unit.
	PartNumber *string `json:"PartNumber,omitempty"`
	// This field identifies the reachability to FI-A and B side.
	Reachability *string `json:"Reachability,omitempty"`
	// User label configured for the SIOM.
	UsrLbl *string `json:"UsrLbl,omitempty"`
	// This field identifies the vendor id for this SIOM Unit.
	Vid                         *string                                         `json:"Vid,omitempty"`
	Controller                  NullableManagementControllerRelationship        `json:"Controller,omitempty"`
	EquipmentSystemIoController NullableEquipmentSystemIoControllerRelationship `json:"EquipmentSystemIoController,omitempty"`
	InventoryDeviceInfo         NullableInventoryDeviceInfoRelationship         `json:"InventoryDeviceInfo,omitempty"`
	// An array of relationships to portGroup resources.
	PortGroups           []PortGroupRelationship                     `json:"PortGroups,omitempty"`
	RegisteredDevice     NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EquipmentSharedIoModule EquipmentSharedIoModule

// NewEquipmentSharedIoModule instantiates a new EquipmentSharedIoModule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquipmentSharedIoModule(classId string, objectType string) *EquipmentSharedIoModule {
	this := EquipmentSharedIoModule{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewEquipmentSharedIoModuleWithDefaults instantiates a new EquipmentSharedIoModule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquipmentSharedIoModuleWithDefaults() *EquipmentSharedIoModule {
	this := EquipmentSharedIoModule{}
	var classId string = "equipment.SharedIoModule"
	this.ClassId = classId
	var objectType string = "equipment.SharedIoModule"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *EquipmentSharedIoModule) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *EquipmentSharedIoModule) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *EquipmentSharedIoModule) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "equipment.SharedIoModule" of the ClassId field.
func (o *EquipmentSharedIoModule) GetDefaultClassId() interface{} {
	return "equipment.SharedIoModule"
}

// GetObjectType returns the ObjectType field value
func (o *EquipmentSharedIoModule) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *EquipmentSharedIoModule) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *EquipmentSharedIoModule) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "equipment.SharedIoModule" of the ObjectType field.
func (o *EquipmentSharedIoModule) GetDefaultObjectType() interface{} {
	return "equipment.SharedIoModule"
}

// GetConfigState returns the ConfigState field value if set, zero value otherwise.
func (o *EquipmentSharedIoModule) GetConfigState() string {
	if o == nil || IsNil(o.ConfigState) {
		var ret string
		return ret
	}
	return *o.ConfigState
}

// GetConfigStateOk returns a tuple with the ConfigState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSharedIoModule) GetConfigStateOk() (*string, bool) {
	if o == nil || IsNil(o.ConfigState) {
		return nil, false
	}
	return o.ConfigState, true
}

// HasConfigState returns a boolean if a field has been set.
func (o *EquipmentSharedIoModule) HasConfigState() bool {
	if o != nil && !IsNil(o.ConfigState) {
		return true
	}

	return false
}

// SetConfigState gets a reference to the given string and assigns it to the ConfigState field.
func (o *EquipmentSharedIoModule) SetConfigState(v string) {
	o.ConfigState = &v
}

// GetDiscovery returns the Discovery field value if set, zero value otherwise.
func (o *EquipmentSharedIoModule) GetDiscovery() string {
	if o == nil || IsNil(o.Discovery) {
		var ret string
		return ret
	}
	return *o.Discovery
}

// GetDiscoveryOk returns a tuple with the Discovery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSharedIoModule) GetDiscoveryOk() (*string, bool) {
	if o == nil || IsNil(o.Discovery) {
		return nil, false
	}
	return o.Discovery, true
}

// HasDiscovery returns a boolean if a field has been set.
func (o *EquipmentSharedIoModule) HasDiscovery() bool {
	if o != nil && !IsNil(o.Discovery) {
		return true
	}

	return false
}

// SetDiscovery gets a reference to the given string and assigns it to the Discovery field.
func (o *EquipmentSharedIoModule) SetDiscovery(v string) {
	o.Discovery = &v
}

// GetMacOfSharedIomAside returns the MacOfSharedIomAside field value if set, zero value otherwise.
func (o *EquipmentSharedIoModule) GetMacOfSharedIomAside() string {
	if o == nil || IsNil(o.MacOfSharedIomAside) {
		var ret string
		return ret
	}
	return *o.MacOfSharedIomAside
}

// GetMacOfSharedIomAsideOk returns a tuple with the MacOfSharedIomAside field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSharedIoModule) GetMacOfSharedIomAsideOk() (*string, bool) {
	if o == nil || IsNil(o.MacOfSharedIomAside) {
		return nil, false
	}
	return o.MacOfSharedIomAside, true
}

// HasMacOfSharedIomAside returns a boolean if a field has been set.
func (o *EquipmentSharedIoModule) HasMacOfSharedIomAside() bool {
	if o != nil && !IsNil(o.MacOfSharedIomAside) {
		return true
	}

	return false
}

// SetMacOfSharedIomAside gets a reference to the given string and assigns it to the MacOfSharedIomAside field.
func (o *EquipmentSharedIoModule) SetMacOfSharedIomAside(v string) {
	o.MacOfSharedIomAside = &v
}

// GetMacOfSharedIomBside returns the MacOfSharedIomBside field value if set, zero value otherwise.
func (o *EquipmentSharedIoModule) GetMacOfSharedIomBside() string {
	if o == nil || IsNil(o.MacOfSharedIomBside) {
		var ret string
		return ret
	}
	return *o.MacOfSharedIomBside
}

// GetMacOfSharedIomBsideOk returns a tuple with the MacOfSharedIomBside field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSharedIoModule) GetMacOfSharedIomBsideOk() (*string, bool) {
	if o == nil || IsNil(o.MacOfSharedIomBside) {
		return nil, false
	}
	return o.MacOfSharedIomBside, true
}

// HasMacOfSharedIomBside returns a boolean if a field has been set.
func (o *EquipmentSharedIoModule) HasMacOfSharedIomBside() bool {
	if o != nil && !IsNil(o.MacOfSharedIomBside) {
		return true
	}

	return false
}

// SetMacOfSharedIomBside gets a reference to the given string and assigns it to the MacOfSharedIomBside field.
func (o *EquipmentSharedIoModule) SetMacOfSharedIomBside(v string) {
	o.MacOfSharedIomBside = &v
}

// GetOperState returns the OperState field value if set, zero value otherwise.
func (o *EquipmentSharedIoModule) GetOperState() string {
	if o == nil || IsNil(o.OperState) {
		var ret string
		return ret
	}
	return *o.OperState
}

// GetOperStateOk returns a tuple with the OperState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSharedIoModule) GetOperStateOk() (*string, bool) {
	if o == nil || IsNil(o.OperState) {
		return nil, false
	}
	return o.OperState, true
}

// HasOperState returns a boolean if a field has been set.
func (o *EquipmentSharedIoModule) HasOperState() bool {
	if o != nil && !IsNil(o.OperState) {
		return true
	}

	return false
}

// SetOperState gets a reference to the given string and assigns it to the OperState field.
func (o *EquipmentSharedIoModule) SetOperState(v string) {
	o.OperState = &v
}

// GetPartNumber returns the PartNumber field value if set, zero value otherwise.
func (o *EquipmentSharedIoModule) GetPartNumber() string {
	if o == nil || IsNil(o.PartNumber) {
		var ret string
		return ret
	}
	return *o.PartNumber
}

// GetPartNumberOk returns a tuple with the PartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSharedIoModule) GetPartNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PartNumber) {
		return nil, false
	}
	return o.PartNumber, true
}

// HasPartNumber returns a boolean if a field has been set.
func (o *EquipmentSharedIoModule) HasPartNumber() bool {
	if o != nil && !IsNil(o.PartNumber) {
		return true
	}

	return false
}

// SetPartNumber gets a reference to the given string and assigns it to the PartNumber field.
func (o *EquipmentSharedIoModule) SetPartNumber(v string) {
	o.PartNumber = &v
}

// GetReachability returns the Reachability field value if set, zero value otherwise.
func (o *EquipmentSharedIoModule) GetReachability() string {
	if o == nil || IsNil(o.Reachability) {
		var ret string
		return ret
	}
	return *o.Reachability
}

// GetReachabilityOk returns a tuple with the Reachability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSharedIoModule) GetReachabilityOk() (*string, bool) {
	if o == nil || IsNil(o.Reachability) {
		return nil, false
	}
	return o.Reachability, true
}

// HasReachability returns a boolean if a field has been set.
func (o *EquipmentSharedIoModule) HasReachability() bool {
	if o != nil && !IsNil(o.Reachability) {
		return true
	}

	return false
}

// SetReachability gets a reference to the given string and assigns it to the Reachability field.
func (o *EquipmentSharedIoModule) SetReachability(v string) {
	o.Reachability = &v
}

// GetUsrLbl returns the UsrLbl field value if set, zero value otherwise.
func (o *EquipmentSharedIoModule) GetUsrLbl() string {
	if o == nil || IsNil(o.UsrLbl) {
		var ret string
		return ret
	}
	return *o.UsrLbl
}

// GetUsrLblOk returns a tuple with the UsrLbl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSharedIoModule) GetUsrLblOk() (*string, bool) {
	if o == nil || IsNil(o.UsrLbl) {
		return nil, false
	}
	return o.UsrLbl, true
}

// HasUsrLbl returns a boolean if a field has been set.
func (o *EquipmentSharedIoModule) HasUsrLbl() bool {
	if o != nil && !IsNil(o.UsrLbl) {
		return true
	}

	return false
}

// SetUsrLbl gets a reference to the given string and assigns it to the UsrLbl field.
func (o *EquipmentSharedIoModule) SetUsrLbl(v string) {
	o.UsrLbl = &v
}

// GetVid returns the Vid field value if set, zero value otherwise.
func (o *EquipmentSharedIoModule) GetVid() string {
	if o == nil || IsNil(o.Vid) {
		var ret string
		return ret
	}
	return *o.Vid
}

// GetVidOk returns a tuple with the Vid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSharedIoModule) GetVidOk() (*string, bool) {
	if o == nil || IsNil(o.Vid) {
		return nil, false
	}
	return o.Vid, true
}

// HasVid returns a boolean if a field has been set.
func (o *EquipmentSharedIoModule) HasVid() bool {
	if o != nil && !IsNil(o.Vid) {
		return true
	}

	return false
}

// SetVid gets a reference to the given string and assigns it to the Vid field.
func (o *EquipmentSharedIoModule) SetVid(v string) {
	o.Vid = &v
}

// GetController returns the Controller field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquipmentSharedIoModule) GetController() ManagementControllerRelationship {
	if o == nil || IsNil(o.Controller.Get()) {
		var ret ManagementControllerRelationship
		return ret
	}
	return *o.Controller.Get()
}

// GetControllerOk returns a tuple with the Controller field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquipmentSharedIoModule) GetControllerOk() (*ManagementControllerRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Controller.Get(), o.Controller.IsSet()
}

// HasController returns a boolean if a field has been set.
func (o *EquipmentSharedIoModule) HasController() bool {
	if o != nil && o.Controller.IsSet() {
		return true
	}

	return false
}

// SetController gets a reference to the given NullableManagementControllerRelationship and assigns it to the Controller field.
func (o *EquipmentSharedIoModule) SetController(v ManagementControllerRelationship) {
	o.Controller.Set(&v)
}

// SetControllerNil sets the value for Controller to be an explicit nil
func (o *EquipmentSharedIoModule) SetControllerNil() {
	o.Controller.Set(nil)
}

// UnsetController ensures that no value is present for Controller, not even an explicit nil
func (o *EquipmentSharedIoModule) UnsetController() {
	o.Controller.Unset()
}

// GetEquipmentSystemIoController returns the EquipmentSystemIoController field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquipmentSharedIoModule) GetEquipmentSystemIoController() EquipmentSystemIoControllerRelationship {
	if o == nil || IsNil(o.EquipmentSystemIoController.Get()) {
		var ret EquipmentSystemIoControllerRelationship
		return ret
	}
	return *o.EquipmentSystemIoController.Get()
}

// GetEquipmentSystemIoControllerOk returns a tuple with the EquipmentSystemIoController field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquipmentSharedIoModule) GetEquipmentSystemIoControllerOk() (*EquipmentSystemIoControllerRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.EquipmentSystemIoController.Get(), o.EquipmentSystemIoController.IsSet()
}

// HasEquipmentSystemIoController returns a boolean if a field has been set.
func (o *EquipmentSharedIoModule) HasEquipmentSystemIoController() bool {
	if o != nil && o.EquipmentSystemIoController.IsSet() {
		return true
	}

	return false
}

// SetEquipmentSystemIoController gets a reference to the given NullableEquipmentSystemIoControllerRelationship and assigns it to the EquipmentSystemIoController field.
func (o *EquipmentSharedIoModule) SetEquipmentSystemIoController(v EquipmentSystemIoControllerRelationship) {
	o.EquipmentSystemIoController.Set(&v)
}

// SetEquipmentSystemIoControllerNil sets the value for EquipmentSystemIoController to be an explicit nil
func (o *EquipmentSharedIoModule) SetEquipmentSystemIoControllerNil() {
	o.EquipmentSystemIoController.Set(nil)
}

// UnsetEquipmentSystemIoController ensures that no value is present for EquipmentSystemIoController, not even an explicit nil
func (o *EquipmentSharedIoModule) UnsetEquipmentSystemIoController() {
	o.EquipmentSystemIoController.Unset()
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquipmentSharedIoModule) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || IsNil(o.InventoryDeviceInfo.Get()) {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo.Get()
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquipmentSharedIoModule) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo.Get(), o.InventoryDeviceInfo.IsSet()
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *EquipmentSharedIoModule) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo.IsSet() {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given NullableInventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *EquipmentSharedIoModule) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo.Set(&v)
}

// SetInventoryDeviceInfoNil sets the value for InventoryDeviceInfo to be an explicit nil
func (o *EquipmentSharedIoModule) SetInventoryDeviceInfoNil() {
	o.InventoryDeviceInfo.Set(nil)
}

// UnsetInventoryDeviceInfo ensures that no value is present for InventoryDeviceInfo, not even an explicit nil
func (o *EquipmentSharedIoModule) UnsetInventoryDeviceInfo() {
	o.InventoryDeviceInfo.Unset()
}

// GetPortGroups returns the PortGroups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquipmentSharedIoModule) GetPortGroups() []PortGroupRelationship {
	if o == nil {
		var ret []PortGroupRelationship
		return ret
	}
	return o.PortGroups
}

// GetPortGroupsOk returns a tuple with the PortGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquipmentSharedIoModule) GetPortGroupsOk() ([]PortGroupRelationship, bool) {
	if o == nil || IsNil(o.PortGroups) {
		return nil, false
	}
	return o.PortGroups, true
}

// HasPortGroups returns a boolean if a field has been set.
func (o *EquipmentSharedIoModule) HasPortGroups() bool {
	if o != nil && !IsNil(o.PortGroups) {
		return true
	}

	return false
}

// SetPortGroups gets a reference to the given []PortGroupRelationship and assigns it to the PortGroups field.
func (o *EquipmentSharedIoModule) SetPortGroups(v []PortGroupRelationship) {
	o.PortGroups = v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquipmentSharedIoModule) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquipmentSharedIoModule) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *EquipmentSharedIoModule) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *EquipmentSharedIoModule) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *EquipmentSharedIoModule) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *EquipmentSharedIoModule) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

func (o EquipmentSharedIoModule) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EquipmentSharedIoModule) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.ConfigState) {
		toSerialize["ConfigState"] = o.ConfigState
	}
	if !IsNil(o.Discovery) {
		toSerialize["Discovery"] = o.Discovery
	}
	if !IsNil(o.MacOfSharedIomAside) {
		toSerialize["MacOfSharedIomAside"] = o.MacOfSharedIomAside
	}
	if !IsNil(o.MacOfSharedIomBside) {
		toSerialize["MacOfSharedIomBside"] = o.MacOfSharedIomBside
	}
	if !IsNil(o.OperState) {
		toSerialize["OperState"] = o.OperState
	}
	if !IsNil(o.PartNumber) {
		toSerialize["PartNumber"] = o.PartNumber
	}
	if !IsNil(o.Reachability) {
		toSerialize["Reachability"] = o.Reachability
	}
	if !IsNil(o.UsrLbl) {
		toSerialize["UsrLbl"] = o.UsrLbl
	}
	if !IsNil(o.Vid) {
		toSerialize["Vid"] = o.Vid
	}
	if o.Controller.IsSet() {
		toSerialize["Controller"] = o.Controller.Get()
	}
	if o.EquipmentSystemIoController.IsSet() {
		toSerialize["EquipmentSystemIoController"] = o.EquipmentSystemIoController.Get()
	}
	if o.InventoryDeviceInfo.IsSet() {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo.Get()
	}
	if o.PortGroups != nil {
		toSerialize["PortGroups"] = o.PortGroups
	}
	if o.RegisteredDevice.IsSet() {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EquipmentSharedIoModule) UnmarshalJSON(data []byte) (err error) {
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
	type EquipmentSharedIoModuleWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// This field identifies the configuration state for this SIOM Unit.
		ConfigState *string `json:"ConfigState,omitempty"`
		// This field identifies the discovery state of SIOM.
		Discovery *string `json:"Discovery,omitempty"`
		// This field identifies the MAC of IOM-A side.
		MacOfSharedIomAside *string `json:"MacOfSharedIomAside,omitempty"`
		// This field identifies the MAC of IOM-B side.
		MacOfSharedIomBside *string `json:"MacOfSharedIomBside,omitempty"`
		// This field identifies the SIOM operational state.
		OperState *string `json:"OperState,omitempty"`
		// This field identifies the Part Number for this SIOM Unit.
		PartNumber *string `json:"PartNumber,omitempty"`
		// This field identifies the reachability to FI-A and B side.
		Reachability *string `json:"Reachability,omitempty"`
		// User label configured for the SIOM.
		UsrLbl *string `json:"UsrLbl,omitempty"`
		// This field identifies the vendor id for this SIOM Unit.
		Vid                         *string                                         `json:"Vid,omitempty"`
		Controller                  NullableManagementControllerRelationship        `json:"Controller,omitempty"`
		EquipmentSystemIoController NullableEquipmentSystemIoControllerRelationship `json:"EquipmentSystemIoController,omitempty"`
		InventoryDeviceInfo         NullableInventoryDeviceInfoRelationship         `json:"InventoryDeviceInfo,omitempty"`
		// An array of relationships to portGroup resources.
		PortGroups       []PortGroupRelationship                     `json:"PortGroups,omitempty"`
		RegisteredDevice NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varEquipmentSharedIoModuleWithoutEmbeddedStruct := EquipmentSharedIoModuleWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varEquipmentSharedIoModuleWithoutEmbeddedStruct)
	if err == nil {
		varEquipmentSharedIoModule := _EquipmentSharedIoModule{}
		varEquipmentSharedIoModule.ClassId = varEquipmentSharedIoModuleWithoutEmbeddedStruct.ClassId
		varEquipmentSharedIoModule.ObjectType = varEquipmentSharedIoModuleWithoutEmbeddedStruct.ObjectType
		varEquipmentSharedIoModule.ConfigState = varEquipmentSharedIoModuleWithoutEmbeddedStruct.ConfigState
		varEquipmentSharedIoModule.Discovery = varEquipmentSharedIoModuleWithoutEmbeddedStruct.Discovery
		varEquipmentSharedIoModule.MacOfSharedIomAside = varEquipmentSharedIoModuleWithoutEmbeddedStruct.MacOfSharedIomAside
		varEquipmentSharedIoModule.MacOfSharedIomBside = varEquipmentSharedIoModuleWithoutEmbeddedStruct.MacOfSharedIomBside
		varEquipmentSharedIoModule.OperState = varEquipmentSharedIoModuleWithoutEmbeddedStruct.OperState
		varEquipmentSharedIoModule.PartNumber = varEquipmentSharedIoModuleWithoutEmbeddedStruct.PartNumber
		varEquipmentSharedIoModule.Reachability = varEquipmentSharedIoModuleWithoutEmbeddedStruct.Reachability
		varEquipmentSharedIoModule.UsrLbl = varEquipmentSharedIoModuleWithoutEmbeddedStruct.UsrLbl
		varEquipmentSharedIoModule.Vid = varEquipmentSharedIoModuleWithoutEmbeddedStruct.Vid
		varEquipmentSharedIoModule.Controller = varEquipmentSharedIoModuleWithoutEmbeddedStruct.Controller
		varEquipmentSharedIoModule.EquipmentSystemIoController = varEquipmentSharedIoModuleWithoutEmbeddedStruct.EquipmentSystemIoController
		varEquipmentSharedIoModule.InventoryDeviceInfo = varEquipmentSharedIoModuleWithoutEmbeddedStruct.InventoryDeviceInfo
		varEquipmentSharedIoModule.PortGroups = varEquipmentSharedIoModuleWithoutEmbeddedStruct.PortGroups
		varEquipmentSharedIoModule.RegisteredDevice = varEquipmentSharedIoModuleWithoutEmbeddedStruct.RegisteredDevice
		*o = EquipmentSharedIoModule(varEquipmentSharedIoModule)
	} else {
		return err
	}

	varEquipmentSharedIoModule := _EquipmentSharedIoModule{}

	err = json.Unmarshal(data, &varEquipmentSharedIoModule)
	if err == nil {
		o.EquipmentBase = varEquipmentSharedIoModule.EquipmentBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ConfigState")
		delete(additionalProperties, "Discovery")
		delete(additionalProperties, "MacOfSharedIomAside")
		delete(additionalProperties, "MacOfSharedIomBside")
		delete(additionalProperties, "OperState")
		delete(additionalProperties, "PartNumber")
		delete(additionalProperties, "Reachability")
		delete(additionalProperties, "UsrLbl")
		delete(additionalProperties, "Vid")
		delete(additionalProperties, "Controller")
		delete(additionalProperties, "EquipmentSystemIoController")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "PortGroups")
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

type NullableEquipmentSharedIoModule struct {
	value *EquipmentSharedIoModule
	isSet bool
}

func (v NullableEquipmentSharedIoModule) Get() *EquipmentSharedIoModule {
	return v.value
}

func (v *NullableEquipmentSharedIoModule) Set(val *EquipmentSharedIoModule) {
	v.value = val
	v.isSet = true
}

func (v NullableEquipmentSharedIoModule) IsSet() bool {
	return v.isSet
}

func (v *NullableEquipmentSharedIoModule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquipmentSharedIoModule(val *EquipmentSharedIoModule) *NullableEquipmentSharedIoModule {
	return &NullableEquipmentSharedIoModule{value: val, isSet: true}
}

func (v NullableEquipmentSharedIoModule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquipmentSharedIoModule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
