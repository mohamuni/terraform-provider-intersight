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

// checks if the EquipmentFan type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EquipmentFan{}

// EquipmentFan Fan in a Fabric Interconnect / Chassis / RackUnit.
type EquipmentFan struct {
	EquipmentBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// This field is to provide description for the fan.
	Description *string `json:"Description,omitempty"`
	// This field acts as the identifier for this particular Fan, within the Fabric Interconnect.
	FanId *int64 `json:"FanId,omitempty"`
	// This field is used to identify the Fan Module to which this Fan belongs.
	FanModuleId *int64 `json:"FanModuleId,omitempty"`
	// Fan module Identifier for the fan.
	ModuleId   *int64   `json:"ModuleId,omitempty"`
	OperReason []string `json:"OperReason,omitempty"`
	// This field is used to indicate this fan unit's operational state.
	OperState *string `json:"OperState,omitempty"`
	// This field identifies the Part Number for this Fan Unit.
	PartNumber *string `json:"PartNumber,omitempty"`
	// This field identifies the Product ID for the fans.
	Pid *string `json:"Pid,omitempty"`
	// This field identifies the Stockkeeping Unit for this Fan Unit.
	Sku *string `json:"Sku,omitempty"`
	// Tray identifier for the fan module.
	TrayId *int64 `json:"TrayId,omitempty"`
	// This field identifies the Vendor ID for this Fan Unit.
	Vid                  *string                                     `json:"Vid,omitempty"`
	EquipmentFanModule   NullableEquipmentFanModuleRelationship      `json:"EquipmentFanModule,omitempty"`
	EquipmentFex         NullableEquipmentFexRelationship            `json:"EquipmentFex,omitempty"`
	InventoryDeviceInfo  NullableInventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
	RegisteredDevice     NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EquipmentFan EquipmentFan

// NewEquipmentFan instantiates a new EquipmentFan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquipmentFan(classId string, objectType string) *EquipmentFan {
	this := EquipmentFan{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewEquipmentFanWithDefaults instantiates a new EquipmentFan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquipmentFanWithDefaults() *EquipmentFan {
	this := EquipmentFan{}
	var classId string = "equipment.Fan"
	this.ClassId = classId
	var objectType string = "equipment.Fan"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *EquipmentFan) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *EquipmentFan) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *EquipmentFan) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "equipment.Fan" of the ClassId field.
func (o *EquipmentFan) GetDefaultClassId() interface{} {
	return "equipment.Fan"
}

// GetObjectType returns the ObjectType field value
func (o *EquipmentFan) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *EquipmentFan) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *EquipmentFan) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "equipment.Fan" of the ObjectType field.
func (o *EquipmentFan) GetDefaultObjectType() interface{} {
	return "equipment.Fan"
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EquipmentFan) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentFan) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EquipmentFan) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EquipmentFan) SetDescription(v string) {
	o.Description = &v
}

// GetFanId returns the FanId field value if set, zero value otherwise.
func (o *EquipmentFan) GetFanId() int64 {
	if o == nil || IsNil(o.FanId) {
		var ret int64
		return ret
	}
	return *o.FanId
}

// GetFanIdOk returns a tuple with the FanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentFan) GetFanIdOk() (*int64, bool) {
	if o == nil || IsNil(o.FanId) {
		return nil, false
	}
	return o.FanId, true
}

// HasFanId returns a boolean if a field has been set.
func (o *EquipmentFan) HasFanId() bool {
	if o != nil && !IsNil(o.FanId) {
		return true
	}

	return false
}

// SetFanId gets a reference to the given int64 and assigns it to the FanId field.
func (o *EquipmentFan) SetFanId(v int64) {
	o.FanId = &v
}

// GetFanModuleId returns the FanModuleId field value if set, zero value otherwise.
func (o *EquipmentFan) GetFanModuleId() int64 {
	if o == nil || IsNil(o.FanModuleId) {
		var ret int64
		return ret
	}
	return *o.FanModuleId
}

// GetFanModuleIdOk returns a tuple with the FanModuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentFan) GetFanModuleIdOk() (*int64, bool) {
	if o == nil || IsNil(o.FanModuleId) {
		return nil, false
	}
	return o.FanModuleId, true
}

// HasFanModuleId returns a boolean if a field has been set.
func (o *EquipmentFan) HasFanModuleId() bool {
	if o != nil && !IsNil(o.FanModuleId) {
		return true
	}

	return false
}

// SetFanModuleId gets a reference to the given int64 and assigns it to the FanModuleId field.
func (o *EquipmentFan) SetFanModuleId(v int64) {
	o.FanModuleId = &v
}

// GetModuleId returns the ModuleId field value if set, zero value otherwise.
func (o *EquipmentFan) GetModuleId() int64 {
	if o == nil || IsNil(o.ModuleId) {
		var ret int64
		return ret
	}
	return *o.ModuleId
}

// GetModuleIdOk returns a tuple with the ModuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentFan) GetModuleIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ModuleId) {
		return nil, false
	}
	return o.ModuleId, true
}

// HasModuleId returns a boolean if a field has been set.
func (o *EquipmentFan) HasModuleId() bool {
	if o != nil && !IsNil(o.ModuleId) {
		return true
	}

	return false
}

// SetModuleId gets a reference to the given int64 and assigns it to the ModuleId field.
func (o *EquipmentFan) SetModuleId(v int64) {
	o.ModuleId = &v
}

// GetOperReason returns the OperReason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquipmentFan) GetOperReason() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.OperReason
}

// GetOperReasonOk returns a tuple with the OperReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquipmentFan) GetOperReasonOk() ([]string, bool) {
	if o == nil || IsNil(o.OperReason) {
		return nil, false
	}
	return o.OperReason, true
}

// HasOperReason returns a boolean if a field has been set.
func (o *EquipmentFan) HasOperReason() bool {
	if o != nil && !IsNil(o.OperReason) {
		return true
	}

	return false
}

// SetOperReason gets a reference to the given []string and assigns it to the OperReason field.
func (o *EquipmentFan) SetOperReason(v []string) {
	o.OperReason = v
}

// GetOperState returns the OperState field value if set, zero value otherwise.
func (o *EquipmentFan) GetOperState() string {
	if o == nil || IsNil(o.OperState) {
		var ret string
		return ret
	}
	return *o.OperState
}

// GetOperStateOk returns a tuple with the OperState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentFan) GetOperStateOk() (*string, bool) {
	if o == nil || IsNil(o.OperState) {
		return nil, false
	}
	return o.OperState, true
}

// HasOperState returns a boolean if a field has been set.
func (o *EquipmentFan) HasOperState() bool {
	if o != nil && !IsNil(o.OperState) {
		return true
	}

	return false
}

// SetOperState gets a reference to the given string and assigns it to the OperState field.
func (o *EquipmentFan) SetOperState(v string) {
	o.OperState = &v
}

// GetPartNumber returns the PartNumber field value if set, zero value otherwise.
func (o *EquipmentFan) GetPartNumber() string {
	if o == nil || IsNil(o.PartNumber) {
		var ret string
		return ret
	}
	return *o.PartNumber
}

// GetPartNumberOk returns a tuple with the PartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentFan) GetPartNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PartNumber) {
		return nil, false
	}
	return o.PartNumber, true
}

// HasPartNumber returns a boolean if a field has been set.
func (o *EquipmentFan) HasPartNumber() bool {
	if o != nil && !IsNil(o.PartNumber) {
		return true
	}

	return false
}

// SetPartNumber gets a reference to the given string and assigns it to the PartNumber field.
func (o *EquipmentFan) SetPartNumber(v string) {
	o.PartNumber = &v
}

// GetPid returns the Pid field value if set, zero value otherwise.
func (o *EquipmentFan) GetPid() string {
	if o == nil || IsNil(o.Pid) {
		var ret string
		return ret
	}
	return *o.Pid
}

// GetPidOk returns a tuple with the Pid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentFan) GetPidOk() (*string, bool) {
	if o == nil || IsNil(o.Pid) {
		return nil, false
	}
	return o.Pid, true
}

// HasPid returns a boolean if a field has been set.
func (o *EquipmentFan) HasPid() bool {
	if o != nil && !IsNil(o.Pid) {
		return true
	}

	return false
}

// SetPid gets a reference to the given string and assigns it to the Pid field.
func (o *EquipmentFan) SetPid(v string) {
	o.Pid = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *EquipmentFan) GetSku() string {
	if o == nil || IsNil(o.Sku) {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentFan) GetSkuOk() (*string, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *EquipmentFan) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *EquipmentFan) SetSku(v string) {
	o.Sku = &v
}

// GetTrayId returns the TrayId field value if set, zero value otherwise.
func (o *EquipmentFan) GetTrayId() int64 {
	if o == nil || IsNil(o.TrayId) {
		var ret int64
		return ret
	}
	return *o.TrayId
}

// GetTrayIdOk returns a tuple with the TrayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentFan) GetTrayIdOk() (*int64, bool) {
	if o == nil || IsNil(o.TrayId) {
		return nil, false
	}
	return o.TrayId, true
}

// HasTrayId returns a boolean if a field has been set.
func (o *EquipmentFan) HasTrayId() bool {
	if o != nil && !IsNil(o.TrayId) {
		return true
	}

	return false
}

// SetTrayId gets a reference to the given int64 and assigns it to the TrayId field.
func (o *EquipmentFan) SetTrayId(v int64) {
	o.TrayId = &v
}

// GetVid returns the Vid field value if set, zero value otherwise.
func (o *EquipmentFan) GetVid() string {
	if o == nil || IsNil(o.Vid) {
		var ret string
		return ret
	}
	return *o.Vid
}

// GetVidOk returns a tuple with the Vid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentFan) GetVidOk() (*string, bool) {
	if o == nil || IsNil(o.Vid) {
		return nil, false
	}
	return o.Vid, true
}

// HasVid returns a boolean if a field has been set.
func (o *EquipmentFan) HasVid() bool {
	if o != nil && !IsNil(o.Vid) {
		return true
	}

	return false
}

// SetVid gets a reference to the given string and assigns it to the Vid field.
func (o *EquipmentFan) SetVid(v string) {
	o.Vid = &v
}

// GetEquipmentFanModule returns the EquipmentFanModule field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquipmentFan) GetEquipmentFanModule() EquipmentFanModuleRelationship {
	if o == nil || IsNil(o.EquipmentFanModule.Get()) {
		var ret EquipmentFanModuleRelationship
		return ret
	}
	return *o.EquipmentFanModule.Get()
}

// GetEquipmentFanModuleOk returns a tuple with the EquipmentFanModule field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquipmentFan) GetEquipmentFanModuleOk() (*EquipmentFanModuleRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.EquipmentFanModule.Get(), o.EquipmentFanModule.IsSet()
}

// HasEquipmentFanModule returns a boolean if a field has been set.
func (o *EquipmentFan) HasEquipmentFanModule() bool {
	if o != nil && o.EquipmentFanModule.IsSet() {
		return true
	}

	return false
}

// SetEquipmentFanModule gets a reference to the given NullableEquipmentFanModuleRelationship and assigns it to the EquipmentFanModule field.
func (o *EquipmentFan) SetEquipmentFanModule(v EquipmentFanModuleRelationship) {
	o.EquipmentFanModule.Set(&v)
}

// SetEquipmentFanModuleNil sets the value for EquipmentFanModule to be an explicit nil
func (o *EquipmentFan) SetEquipmentFanModuleNil() {
	o.EquipmentFanModule.Set(nil)
}

// UnsetEquipmentFanModule ensures that no value is present for EquipmentFanModule, not even an explicit nil
func (o *EquipmentFan) UnsetEquipmentFanModule() {
	o.EquipmentFanModule.Unset()
}

// GetEquipmentFex returns the EquipmentFex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquipmentFan) GetEquipmentFex() EquipmentFexRelationship {
	if o == nil || IsNil(o.EquipmentFex.Get()) {
		var ret EquipmentFexRelationship
		return ret
	}
	return *o.EquipmentFex.Get()
}

// GetEquipmentFexOk returns a tuple with the EquipmentFex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquipmentFan) GetEquipmentFexOk() (*EquipmentFexRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.EquipmentFex.Get(), o.EquipmentFex.IsSet()
}

// HasEquipmentFex returns a boolean if a field has been set.
func (o *EquipmentFan) HasEquipmentFex() bool {
	if o != nil && o.EquipmentFex.IsSet() {
		return true
	}

	return false
}

// SetEquipmentFex gets a reference to the given NullableEquipmentFexRelationship and assigns it to the EquipmentFex field.
func (o *EquipmentFan) SetEquipmentFex(v EquipmentFexRelationship) {
	o.EquipmentFex.Set(&v)
}

// SetEquipmentFexNil sets the value for EquipmentFex to be an explicit nil
func (o *EquipmentFan) SetEquipmentFexNil() {
	o.EquipmentFex.Set(nil)
}

// UnsetEquipmentFex ensures that no value is present for EquipmentFex, not even an explicit nil
func (o *EquipmentFan) UnsetEquipmentFex() {
	o.EquipmentFex.Unset()
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquipmentFan) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || IsNil(o.InventoryDeviceInfo.Get()) {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo.Get()
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquipmentFan) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo.Get(), o.InventoryDeviceInfo.IsSet()
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *EquipmentFan) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo.IsSet() {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given NullableInventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *EquipmentFan) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo.Set(&v)
}

// SetInventoryDeviceInfoNil sets the value for InventoryDeviceInfo to be an explicit nil
func (o *EquipmentFan) SetInventoryDeviceInfoNil() {
	o.InventoryDeviceInfo.Set(nil)
}

// UnsetInventoryDeviceInfo ensures that no value is present for InventoryDeviceInfo, not even an explicit nil
func (o *EquipmentFan) UnsetInventoryDeviceInfo() {
	o.InventoryDeviceInfo.Unset()
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquipmentFan) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquipmentFan) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *EquipmentFan) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *EquipmentFan) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *EquipmentFan) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *EquipmentFan) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

func (o EquipmentFan) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EquipmentFan) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	if !IsNil(o.FanId) {
		toSerialize["FanId"] = o.FanId
	}
	if !IsNil(o.FanModuleId) {
		toSerialize["FanModuleId"] = o.FanModuleId
	}
	if !IsNil(o.ModuleId) {
		toSerialize["ModuleId"] = o.ModuleId
	}
	if o.OperReason != nil {
		toSerialize["OperReason"] = o.OperReason
	}
	if !IsNil(o.OperState) {
		toSerialize["OperState"] = o.OperState
	}
	if !IsNil(o.PartNumber) {
		toSerialize["PartNumber"] = o.PartNumber
	}
	if !IsNil(o.Pid) {
		toSerialize["Pid"] = o.Pid
	}
	if !IsNil(o.Sku) {
		toSerialize["Sku"] = o.Sku
	}
	if !IsNil(o.TrayId) {
		toSerialize["TrayId"] = o.TrayId
	}
	if !IsNil(o.Vid) {
		toSerialize["Vid"] = o.Vid
	}
	if o.EquipmentFanModule.IsSet() {
		toSerialize["EquipmentFanModule"] = o.EquipmentFanModule.Get()
	}
	if o.EquipmentFex.IsSet() {
		toSerialize["EquipmentFex"] = o.EquipmentFex.Get()
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

func (o *EquipmentFan) UnmarshalJSON(data []byte) (err error) {
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
	type EquipmentFanWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// This field is to provide description for the fan.
		Description *string `json:"Description,omitempty"`
		// This field acts as the identifier for this particular Fan, within the Fabric Interconnect.
		FanId *int64 `json:"FanId,omitempty"`
		// This field is used to identify the Fan Module to which this Fan belongs.
		FanModuleId *int64 `json:"FanModuleId,omitempty"`
		// Fan module Identifier for the fan.
		ModuleId   *int64   `json:"ModuleId,omitempty"`
		OperReason []string `json:"OperReason,omitempty"`
		// This field is used to indicate this fan unit's operational state.
		OperState *string `json:"OperState,omitempty"`
		// This field identifies the Part Number for this Fan Unit.
		PartNumber *string `json:"PartNumber,omitempty"`
		// This field identifies the Product ID for the fans.
		Pid *string `json:"Pid,omitempty"`
		// This field identifies the Stockkeeping Unit for this Fan Unit.
		Sku *string `json:"Sku,omitempty"`
		// Tray identifier for the fan module.
		TrayId *int64 `json:"TrayId,omitempty"`
		// This field identifies the Vendor ID for this Fan Unit.
		Vid                 *string                                     `json:"Vid,omitempty"`
		EquipmentFanModule  NullableEquipmentFanModuleRelationship      `json:"EquipmentFanModule,omitempty"`
		EquipmentFex        NullableEquipmentFexRelationship            `json:"EquipmentFex,omitempty"`
		InventoryDeviceInfo NullableInventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
		RegisteredDevice    NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varEquipmentFanWithoutEmbeddedStruct := EquipmentFanWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varEquipmentFanWithoutEmbeddedStruct)
	if err == nil {
		varEquipmentFan := _EquipmentFan{}
		varEquipmentFan.ClassId = varEquipmentFanWithoutEmbeddedStruct.ClassId
		varEquipmentFan.ObjectType = varEquipmentFanWithoutEmbeddedStruct.ObjectType
		varEquipmentFan.Description = varEquipmentFanWithoutEmbeddedStruct.Description
		varEquipmentFan.FanId = varEquipmentFanWithoutEmbeddedStruct.FanId
		varEquipmentFan.FanModuleId = varEquipmentFanWithoutEmbeddedStruct.FanModuleId
		varEquipmentFan.ModuleId = varEquipmentFanWithoutEmbeddedStruct.ModuleId
		varEquipmentFan.OperReason = varEquipmentFanWithoutEmbeddedStruct.OperReason
		varEquipmentFan.OperState = varEquipmentFanWithoutEmbeddedStruct.OperState
		varEquipmentFan.PartNumber = varEquipmentFanWithoutEmbeddedStruct.PartNumber
		varEquipmentFan.Pid = varEquipmentFanWithoutEmbeddedStruct.Pid
		varEquipmentFan.Sku = varEquipmentFanWithoutEmbeddedStruct.Sku
		varEquipmentFan.TrayId = varEquipmentFanWithoutEmbeddedStruct.TrayId
		varEquipmentFan.Vid = varEquipmentFanWithoutEmbeddedStruct.Vid
		varEquipmentFan.EquipmentFanModule = varEquipmentFanWithoutEmbeddedStruct.EquipmentFanModule
		varEquipmentFan.EquipmentFex = varEquipmentFanWithoutEmbeddedStruct.EquipmentFex
		varEquipmentFan.InventoryDeviceInfo = varEquipmentFanWithoutEmbeddedStruct.InventoryDeviceInfo
		varEquipmentFan.RegisteredDevice = varEquipmentFanWithoutEmbeddedStruct.RegisteredDevice
		*o = EquipmentFan(varEquipmentFan)
	} else {
		return err
	}

	varEquipmentFan := _EquipmentFan{}

	err = json.Unmarshal(data, &varEquipmentFan)
	if err == nil {
		o.EquipmentBase = varEquipmentFan.EquipmentBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "FanId")
		delete(additionalProperties, "FanModuleId")
		delete(additionalProperties, "ModuleId")
		delete(additionalProperties, "OperReason")
		delete(additionalProperties, "OperState")
		delete(additionalProperties, "PartNumber")
		delete(additionalProperties, "Pid")
		delete(additionalProperties, "Sku")
		delete(additionalProperties, "TrayId")
		delete(additionalProperties, "Vid")
		delete(additionalProperties, "EquipmentFanModule")
		delete(additionalProperties, "EquipmentFex")
		delete(additionalProperties, "InventoryDeviceInfo")
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

type NullableEquipmentFan struct {
	value *EquipmentFan
	isSet bool
}

func (v NullableEquipmentFan) Get() *EquipmentFan {
	return v.value
}

func (v *NullableEquipmentFan) Set(val *EquipmentFan) {
	v.value = val
	v.isSet = true
}

func (v NullableEquipmentFan) IsSet() bool {
	return v.isSet
}

func (v *NullableEquipmentFan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquipmentFan(val *EquipmentFan) *NullableEquipmentFan {
	return &NullableEquipmentFan{value: val, isSet: true}
}

func (v NullableEquipmentFan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquipmentFan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
